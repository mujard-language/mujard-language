#include "mujard_rt.h"

// MUJARD_HOST_TEST: the runtime's own tests on a PC (tools/mujard/lang/runtime/tests), with a fake Arduino.h and EEPROM.h.
#if defined(ARDUINO_ARCH_AVR) || defined(ARDUINO_ARCH_ESP32) || defined(ARDUINO_ARCH_ESP8266) || defined(ARDUINO_ARCH_RP2040) || defined(MUJARD_HOST_TEST)
#define MUJARD_HAS_EEPROM 1
#define MUJARD_HAS_TONE 1
#include <EEPROM.h>
#endif
#if defined(ARDUINO_ARCH_ESP32) || defined(ARDUINO_ARCH_ESP8266) || defined(ARDUINO_ARCH_RP2040) || defined(MUJARD_HOST_TEST)
#define MUJARD_EEPROM_COMMIT 1  // flash-backed EEPROM: begin(size) first, commit() to write
#endif
#ifdef MUJARD_NO_TONE
// No tone or melody in the program: the core's tone() isn't referenced at all, so its timer interrupt (Timer 2 on AVR)
// stays free for a library that needs it (NewPing's timer functions define the same interrupt).
#undef MUJARD_HAS_TONE
#endif
#if defined(ARDUINO_ARCH_AVR)
#include <avr/wdt.h>
#include <avr/sleep.h>
#elif defined(ARDUINO_ARCH_ESP32)
#include <esp_task_wdt.h>
#include <esp_sleep.h>
#include <esp_system.h>
#include <soc/soc_caps.h>
#elif defined(ARDUINO_ARCH_ESP8266)
#include <user_interface.h>
#elif defined(ARDUINO_ARCH_RP2040)
#include <pico/unique_id.h>
#include <pico/flash.h>
#elif defined(ARDUINO_ARCH_MBED) || defined(ARDUINO_ARCH_MBED_NANO)
#include <nrf.h>
#include <mbed_stats.h>
#endif

#if defined(ARDUINO_ARCH_RP2040)
// RP2040/RP2350: a flash write stalls the peer core's XIP fetch, so EEPROM.commit() must run inside the SDK's
// flash_safe_execute (its multicore lockout) or the peer faults mid-erase. The callback carries commit()'s own
// result back through the void* param.
static void flashCommit(void* ok) { *static_cast<bool*>(ok) = EEPROM.commit(); }
#endif

namespace mujard {

// ---- time: micros() wraps every 71 minutes; this clock doesn't (it must be read at least that often, which tick() does)

static uint32_t lastMicros = 0;
static uint64_t high = 0;
static Time sleptUs = 0;  // power.sleep() time the chip's own clock didn't see (spec 6.10)

Time now() {
    uint32_t m = micros();
    if (m < lastMicros) high += 0x100000000ULL;
    lastMicros = m;
    return (Time)(high | m) + sleptUs;
}

// ---- registry of background work

static Ticked* devices = nullptr;
Ticked::Ticked() : next_(devices) { devices = this; }
Ticked::Ticked(const Ticked&) : next_(devices) { devices = this; }
Ticked::~Ticked() {
    for (Ticked** p = &devices; *p; p = &(*p)->next_) {
        if (*p == this) {
            *p = next_;
            return;
        }
    }
}

static void noSafeState() {}
void (*safe_state_hook)() = noSafeState;
uint32_t overruns = 0;
uint32_t dropped_events = 0;
static volatile bool halted_ = false;
static volatile bool stopping_ = false;  // a halt has begun: fast handlers return at once
// A halt or runtime error inside a fast handler (an interrupt) locks the outputs at once, but the log and safe_state
// wait for the next pass of the main loop: printing and the program's own code are not interrupt-safe (spec 6.5).
static volatile uint8_t inIsr_ = 0;
static volatile bool pending_ = false, pendingIsHalt_ = false;
static volatile int pendingLine_ = 0;
static char pendingMsg_[48];
static bool serialStarted = false;
static Time wdtPeriod = 0;

// ---- serial output serialization (multicore): two cores printing produce byte-interleaved lines unless
// each whole statement is written under a lock. On RP-family a critical_section, on ESP32-family a
// FreeRTOS mutex; a single-core board has nothing to race and pays nothing.
#if defined(ARDUINO_ARCH_RP2040)
#include "pico/critical_section.h"
static critical_section_t serialLock;
static bool serialLockReady = false;
static void serialOutLock() {
    if (!serialLockReady) { critical_section_init(&serialLock); serialLockReady = true; }
    critical_section_enter_blocking(&serialLock);
}
static void serialOutUnlock() { critical_section_exit(&serialLock); }
#elif defined(ARDUINO_ARCH_ESP32)
#include "freertos/FreeRTOS.h"
#include "freertos/semphr.h"
static SemaphoreHandle_t serialLock = nullptr;
static void serialOutLock() {
    if (!serialLock) serialLock = xSemaphoreCreateMutex();
    if (serialLock) xSemaphoreTake(serialLock, portMAX_DELAY);
}
static void serialOutUnlock() { if (serialLock) xSemaphoreGive(serialLock); }
#else
static void serialOutLock() {}
static void serialOutUnlock() {}
#endif

// ---- saved variables (spec 11)
// Each region in EEPROM holds [layout:4][sequence:1][data]. Every write goes to the next region,
// so the wear of repeated saves is spread over the free storage instead of one cell.

struct SavedSlot {
    void* p;
    uint16_t size;
    uint16_t offset;
};
static SavedSlot slots[24];
static uint8_t slotCount = 0;
static uint16_t dataSize = 0;     // saved bytes, not counting the region header
static uint16_t regionSize = 5;   // layout code + sequence byte + data
static uint8_t regionCount = 1;   // how many regions fit in the EEPROM, capped at 8
static uint8_t curRegion = 0;
static uint8_t curSeq = 0;
static uint32_t storeLayout = 0;
static Time nextSave = 0;

#if defined(ARDUINO_ARCH_ESP32)
static const Time saveEvery = 5000000;   // wear-levelled flash: every 5 s (spec 11)
#else
static const Time saveEvery = 60000000;  // EEPROM and the rest: once a minute
#endif

static uint8_t readByte(uint16_t at) {
#ifdef MUJARD_HAS_EEPROM
    return EEPROM.read(at);
#else
    (void)at;
    return 0xFF;
#endif
}

static bool writeByte(uint16_t at, uint8_t v) {
#ifdef MUJARD_HAS_EEPROM
    if (EEPROM.read(at) == v) return false;  // flash and EEPROM wear out: only write what changed
    EEPROM.write(at, v);
    return true;
#else
    (void)at;
    (void)v;
    return false;
#endif
}

void saved_begin(uint32_t layout) {
    storeLayout = layout;
    slotCount = 0;
    dataSize = 0;
}

void saved_add(void* p, uint16_t size) {
    if (slotCount >= sizeof(slots) / sizeof(slots[0])) return;
    slots[slotCount++] = SavedSlot{p, size, dataSize};
    dataSize += size;
}

// Durability observability: every storage_commit path (saved vars and store.save) funnels through
// storeCommit, so its result is recorded once here - the program can read store.failed / store.last_commit
// to tell a real commit from a silent retry.
uint32_t storage_commits_failed = 0;
bool storage_commit_ok = true;

static bool storeCommit() {
#ifdef MUJARD_EEPROM_COMMIT
#if defined(ARDUINO_ARCH_RP2040)
    bool done = false;
    // Lock out the peer core for the whole erase+program; PICO_ERROR_* (lockout failed) or commit() failing is
    // reported as false so the caller keeps the data dirty and retries, never a silent loss.
    bool ok = flash_safe_execute(flashCommit, &done, 1000) == 0 && done;
    storage_commit_ok = ok;
    if (!ok) storage_commits_failed++;
    return ok;
#else
    EEPROM.commit();
    storage_commit_ok = true;
#endif
    return true;  // a byte-write EEPROM (AVR), a commit with no result, or no storage
#endif
}

bool coreInitFlashSafe() {
#if defined(ARDUINO_ARCH_RP2040)
    static bool ready = flash_safe_execute_core_init();  // once per core: let the peer core lock this one out
    return ready;
#else
    return true;  // a single core, or a board whose flash write is driver-managed (ESP32 NVS)
#endif
}

// ---- storage shared by saved variables and store.save (spec 11): saved records from address 0, the store's named
// slots in the last STORE_AREA bytes, so the two never overlap.

void (*store_flush_hook)() = nullptr;

int storage_begin() {
    static int len = -1;
    if (len >= 0) return len;
#ifdef MUJARD_EEPROM_COMMIT
    EEPROM.begin(1024);  // flash-backed EEPROM: one fixed size for both users
    len = (int)EEPROM.length();
#elif defined(MUJARD_HAS_EEPROM)
    len = (int)EEPROM.length();
#else
    len = 0;
#endif
    return len;
}

uint8_t storage_read(uint16_t at) { return readByte(at); }
bool storage_write(uint16_t at, uint8_t v) { return writeByte(at, v); }
bool storage_commit() { return storeCommit(); }

int store_area_start() {
    int len = storage_begin();
    return len >= STORE_AREA + 64 ? len - STORE_AREA : -1;  // too small a storage: no store
}

Time storage_interval() { return saveEvery; }

// A saved var's live value not yet in storage = a pending write (spec 11's auto-save window),
// the same truth store.pending reports for store.save.
bool saved_pending() {
    if (slotCount == 0 || regionSize == 0) return false;
    uint16_t base = (uint16_t)curRegion * regionSize + 5;
    for (uint8_t s = 0; s < slotCount; s++)
        for (uint16_t b = 0; b < slots[s].size; b++)
            if (readByte(base + slots[s].offset + b) != ((uint8_t*)slots[s].p)[b]) return true;
    return false;
}

// Write the record to the next region, unless the newest region already holds these bytes.
static bool savedWrite(bool force) {
    if (slotCount == 0) return false;
    if (!force && !saved_pending()) return false;
    curRegion = (uint8_t)((curRegion + 1) % regionCount);
    curSeq++;
    uint16_t base = (uint16_t)curRegion * regionSize;
    bool any = false;
    for (int i = 0; i < 4; i++) any |= writeByte(base + i, (uint8_t)(storeLayout >> (8 * i)));
    any |= writeByte(base + 4, curSeq);
    for (uint8_t s = 0; s < slotCount; s++)
        for (uint16_t b = 0; b < slots[s].size; b++)
            any |= writeByte(base + 5 + slots[s].offset + b, ((uint8_t*)slots[s].p)[b]);
    if (any) storeCommit();
    return any;
}

void saved_load() {
    if (slotCount == 0) return;
    regionSize = (uint16_t)(5 + dataSize);
    int eepromLen = storage_begin();
    if (store_area_start() >= 0) eepromLen = store_area_start();  // the store's slots sit at the end
    regionCount = (uint8_t)(eepromLen / regionSize);
    if (regionCount > 8) regionCount = 8;
    if (regionCount < 1) regionCount = 1;
    // find the newest region whose layout matches (sequence bytes compare with wrap-around)
    int best = -1;
    uint8_t bestSeq = 0;
    for (uint8_t i = 0; i < regionCount; i++) {
        uint16_t base = (uint16_t)i * regionSize;
        uint32_t got = 0;
        for (int b = 0; b < 4; b++) got |= (uint32_t)readByte(base + b) << (8 * b);
        if (got != storeLayout) continue;
        uint8_t sq = readByte(base + 4);
        if (best < 0 || (int8_t)(sq - bestSeq) > 0) {
            best = i;
            bestSeq = sq;
        }
    }
    if (best < 0) {
        // A new program or a changed list of saved vars: start fresh with the values in the code.
        curRegion = (uint8_t)(regionCount - 1);
        savedWrite(true);
        return;
    }
    curRegion = (uint8_t)best;
    curSeq = bestSeq;
    uint16_t base = (uint16_t)best * regionSize + 5;
    for (uint8_t s = 0; s < slotCount; s++)
        for (uint16_t b = 0; b < slots[s].size; b++) ((uint8_t*)slots[s].p)[b] = readByte(base + slots[s].offset + b);
}

void saved_flush() {  // pending changes go to storage before sleep, restart and halt (spec 11)
    savedWrite(false);
    if (store_flush_hook) store_flush_hook();
}

static void savedSync(Time t) {
    if (slotCount == 0 || t < nextSave) return;
    nextSave = t + saveEvery;
    savedWrite(false);
}

// ---- life

void begin() {
    now();
    for (Ticked* d = devices; d; d = d->next_) d->init();
#if defined(ARDUINO_ARCH_ESP32)
    randomSeed(esp_random());
#elif defined(ARDUINO_ARCH_ESP8266)
    randomSeed(RANDOM_REG32);
#else
    randomSeed(analogRead(0) ^ (uint32_t)micros());
#endif
}

static void feedWatchdog() {
    if (wdtPeriod == 0) return;
#if defined(ARDUINO_ARCH_AVR)
    wdt_reset();
#elif defined(ARDUINO_ARCH_ESP32)
    esp_task_wdt_reset();
#elif defined(ARDUINO_ARCH_ESP8266)
    ESP.wdtFeed();
#elif defined(ARDUINO_ARCH_RP2040)
    rp2040.wdt_reset();
#endif
}

#ifndef MUJARD_NO_SERIAL_IN
static bool lineReady = false;
static Text lineBuf, lineDone;

// serial.read_line(): lines are kept in a small queue until a task takes one. With the queue on, lines the
// program doesn't take stay buffered; 'on serial.line' pops one a pass like before, through the same queue.
static bool wantLines = false;
#if defined(ARDUINO_ARCH_AVR)
static const uint8_t LINE_Q = 2;  // the Uno's RAM is tight
#else
static const uint8_t LINE_Q = 4;
#endif
static Text lineQ[LINE_Q];
static uint8_t qHead = 0, qCount = 0;
static void readSerial();
#endif
bool (*rtc_clock_hook)(struct tm& out) = nullptr;

#ifndef MUJARD_NO_APP
static void appInput(const char* s);
static void appTick(Time t);
#endif

bool tick() {
    Time t = now();
    feedWatchdog();
    if (pending_) {  // a fast handler halted: finish it here, outside the interrupt
        pending_ = false;
        halted_ = false;  // so safe_state can set the outputs; fast handlers stay stopped (stopping_)
        if (pendingIsHalt_) halt(pendingMsg_);
        else fail(pendingMsg_, pendingLine_);
    }
    if (halted_) {
#ifdef LED_BUILTIN
        digitalWrite(LED_BUILTIN, (t / 100000) % 2 ? HIGH : LOW);  // a fast blink says "halted"
#endif
#ifndef MUJARD_NO_APP
        if (app_on()) readSerial();  // the app may still ask who is there (the answer says "halted")
#endif
        return false;
    }
    for (Ticked* d = devices; d; d = d->next_) d->tick(t);
#ifndef MUJARD_NO_SERIAL_IN
    readSerial();
#endif
#ifndef MUJARD_NO_APP
    if (app_on()) appTick(t);
#endif
    savedSync(t);
    return true;
}

#ifndef MUJARD_NO_SERIAL_IN
// One line of serial input per pass: lines from the app (starting with 0x1F) go to the app protocol and never reach
// the program; anything else is serial.line.
static void readSerial() {
    lineReady = false;
    if (!serialStarted) return;
    while (Serial.available() > 0) {
        int c = Serial.read();
        if (c == '\r') continue;
        if (c == '\n') {
            if (lineBuf.s[0] == 0x1F) {  // the app's line: never the program's input
#ifndef MUJARD_NO_APP
                appInput(lineBuf.s);
#endif
                lineBuf = Text();
                continue;
            }
            lineDone = lineBuf;
            lineBuf = Text();
            if (wantLines) {
                if (qCount < LINE_Q) {
                    lineQ[(uint8_t)((qHead + qCount) % LINE_Q)] = lineDone;
                    qCount++;
                } else {
                    dropped_events++;
                }
                continue;  // the queue takes the rest of the lines this pass
            }
            lineReady = true;
            break;
        }
        char one[2] = {(char)c, 0};
        lineBuf.add(one);
    }
}

bool serial_line(Text& out) {
    if (!serialStarted) serial_begin(115200);
    if (wantLines) {  // the queue feeds the event too: one line per pass, none lost
        if (qCount == 0) return false;
        out = lineQ[qHead];
        qHead = (uint8_t)((qHead + 1) % LINE_Q);
        qCount--;
        return true;
    }
    if (!lineReady) return false;
    out = lineDone;
    return true;
}

void serial_want_lines() {
    if (!serialStarted) serial_begin(115200);
    wantLines = true;
}

bool serial_has_line() { return qCount > 0; }

Opt<Text> serial_take_line() {
    Opt<Text> out;
    if (qCount == 0) return out;
    out.has = true;
    out.v = lineQ[qHead];
    qHead = (uint8_t)((qHead + 1) % LINE_Q);
    qCount--;
    return out;
}
#endif  // MUJARD_NO_SERIAL_IN

void serial_begin(uint32_t baud) {
    if (serialStarted) Serial.end();
    Serial.begin(baud);
    serialStarted = true;
}

static void deferFromIsr(bool isHalt, const char* message, int line) {
    if (halted_) return;
    stopping_ = halted_ = true;
    pendingIsHalt_ = isHalt;
    pendingLine_ = line;
    strncpy(pendingMsg_, message ? message : "", sizeof pendingMsg_ - 1);
    pendingMsg_[sizeof pendingMsg_ - 1] = 0;
    pending_ = true;
}

void isr_enter() { inIsr_++; }
void isr_leave() { inIsr_--; }
bool stopping() { return stopping_; }

void fail(const char* message, int line) {
    if (halted_) return;
    if (inIsr_) {
        deferFromIsr(false, message, line);
        return;
    }
    stopping_ = true;
    Out o("ERROR");
    if (line > 0) o.put("(line ").put((int32_t)line).put(") ");
    o.put(message);
    o.end();
    safe_state_hook();
    saved_flush();
    halted_ = true;
#ifdef LED_BUILTIN
    pinMode(LED_BUILTIN, OUTPUT);
#endif
}

void expect_warn(const char* message, int line) {
    // spec 6.8: the log rate is capped at once per 10 s for each source line; safe_state runs every failure.
    static int16_t lines[8];
    static uint32_t last[8];
    static uint8_t nSlots = 0;
    int slot = -1;
    for (uint8_t i = 0; i < nSlots; i++) {
        if (lines[i] == line) slot = i;
    }
    if (slot < 0) {
        if (nSlots < 8) slot = nSlots++;
        else slot = 0; // more than 8 expect lines share the last slot
        lines[slot] = (int16_t)line;
        last[slot] = 0;
    }
    uint32_t t = millis();
    if (last[slot] == 0 || (uint32_t)(t - last[slot]) >= 10000UL) {
        last[slot] = t;
        Out o("WARN");
        if (line > 0) o.put("(line ").put((int32_t)line).put(") ");
        o.put(message);
        o.end();
    }
    safe_state_hook();
}

void halt(const char* reason) {
    if (halted_) return;
    if (inIsr_) {
        deferFromIsr(true, reason, 0);
        return;
    }
    stopping_ = true;
    Out o("HALT");
    o.put(reason);
    o.end();
    safe_state_hook();
    saved_flush();
    halted_ = true;
#ifdef LED_BUILTIN
    pinMode(LED_BUILTIN, OUTPUT);
#endif
}

bool halted() { return halted_; }

#ifdef MUJARD_HOST_TEST
void test_unhalt() { halted_ = stopping_ = pending_ = false; }
#endif

// ---- int arithmetic

int32_t idiv(int32_t a, int32_t b, int line) {
    if (b == 0) {
        fail("Division by zero", line);
        return 0;
    }
    if (a == INT32_MIN && b == -1) return INT32_MIN;  // the one quotient that doesn't fit: it wraps
    return a / b;
}

int32_t imod(int32_t a, int32_t b, int line) {
    if (b == 0) {
        fail("Remainder (mod) by zero", line);
        return 0;
    }
    if (b == -1) return 0;
    return a % b;
}

int32_t ishl(int32_t a, int32_t n, int line) {
    if (n < 0 || n > 31) {
        fail("Shift by less than 0 or more than 31", line);
        return 0;
    }
    return (int32_t)((uint32_t)a << n);
}

int32_t ishr(int32_t a, int32_t n, int line) {
    if (n < 0 || n > 31) {
        fail("Shift by less than 0 or more than 31", line);
        return 0;
    }
    return a >= 0 ? a >> n : ~(~a >> n);  // keeps the sign, on every compiler
}

Time tdiv(Time a, int32_t b, int line) {
    if (b == 0) {
        fail("Division by zero", line);
        return 0;
    }
    return a / b;
}

Time tdivf(Time a, float b, int line) {
    float q = b != 0 ? (float)a / b : 0;
    if (b == 0 || q != q || q > 9.2e18f || q < -9.2e18f) {  // the cast to a Time would be undefined
        fail(b == 0 ? "Division by zero" : "The time is too big", line);
        return 0;
    }
    return (Time)q;
}

// ---- text

Text& Text::add(const char* c) {
    while (c && *c && n < MUJARD_TEXT_MAX) s[n++] = *c++;
    s[n] = 0;
    return *this;
}

// ltoa/ultoa aren't in every core's libc (the mbed core lacks them): a local version keeps text working there
static char* uToStr(uint32_t v, char* b) {
    char tmp[10];
    int n = 0;
    do {
        tmp[n++] = (char)('0' + v % 10);
        v /= 10;
    } while (v);
    char* p = b;
    while (n) *p++ = tmp[--n];
    *p = 0;
    return b;
}
static char* iToStr(int32_t v, char* b) {
    if (v < 0) {
        *b = '-';
        uToStr((uint32_t)(-(v + 1)) + 1, b + 1);
        return b;
    }
    return uToStr((uint32_t)v, b);
}

Text& Text::add(int32_t v) {
    char b[14];
    iToStr(v, b);
    return add(b);
}

Text& Text::add(int64_t v) {
    char b[24];
    char* p = b + sizeof(b);
    *--p = 0;
    bool neg = v < 0;
    uint64_t u = neg ? (uint64_t)(-(v + 1)) + 1 : (uint64_t)v;
    do {
        *--p = (char)('0' + u % 10);
        u /= 10;
    } while (u);
    if (neg) *--p = '-';
    return add(p);
}

static void floatText(float v, int decimals, char* out) {
    if (v != v) {  // NaN (isnan() isn't declared the same way on every core)
        strcpy(out, "nan");
        return;
    }
    // dtostrf-like, without depending on it (not on every core): sign, whole part, decimals, trailing zeros removed
    bool neg = v < 0;
    if (neg) v = -v;
    if (v >= 4294967040.0f) {  // too big for the uint32 whole part (casting it would be undefined): 1.5e12, inf
        if (neg) *out++ = '-';
        if (v > 3.4028235e38f) {
            strcpy(out, "inf");
            return;
        }
        int exp = 0;
        while (v >= 10.0f) {
            v /= 10.0f;
            exp++;
        }
        floatText(v, decimals, out);
        out += strlen(out);
        *out++ = 'e';
        uToStr((uint32_t)exp, out);
        return;
    }
    float scale = 1;
    for (int i = 0; i < decimals; i++) scale *= 10;
    uint32_t whole = (uint32_t)v;
    uint32_t frac = (uint32_t)((v - (float)whole) * scale + 0.5f);
    if (frac >= (uint32_t)scale) {
        whole++;
        frac -= (uint32_t)scale;
    }
    char* p = out;
    if (neg) *p++ = '-';
    uToStr(whole, p);
    p += strlen(p);
    if (decimals > 0 && frac > 0) {
        *p++ = '.';
        char d[12];
        uToStr(frac, d);
        for (int pad = decimals - (int)strlen(d); pad > 0; pad--) *p++ = '0';
        strcpy(p, d);
        p += strlen(p);
        while (p[-1] == '0') *--p = 0;
    }
    *p = 0;
}

static void twoDigits(char* p, int32_t v) {
    p[0] = (char)('0' + (v / 10) % 10);
    p[1] = (char)('0' + v % 10);
}

Text& Text::add(const TimeOfDay& t) {
    char b[9] = "00:00:00";
    twoDigits(b, t.hour);
    twoDigits(b + 3, t.minute);
    twoDigits(b + 6, t.second);
    return add(b);
}

Text& Text::add(const Date& d) {
    char b[11] = "0000-00-00";
    twoDigits(b, d.year / 100);
    twoDigits(b + 2, d.year % 100);
    twoDigits(b + 5, d.month);
    twoDigits(b + 8, d.day);
    return add(b);
}

Text& Text::add(float v) {
    char b[24];
    floatText(v, 3, b);
    return add(b);
}

Text& Text::add_time(Time t) {
    Time a = t < 0 ? -t : t;
    if (a < 1000) return add((int32_t)t).add(" us");
    if (a < 1000000) return add((float)t / 1000.0f).add(" ms");
    if (a < 120000000) return add((float)t / 1000000.0f).add(" s");
    if (a < 7200000000LL) return add((float)t / 60000000.0f).add(" min");
    return add((float)t / 3600000000.0f).add(" h");
}

Text& Text::add_freq(Freq f) {
    Freq a = f < 0 ? -f : f;
    if (a >= 1000000000LL) return add((float)f / 1e9f).add(" MHz");
    if (a >= 1000000LL) return add((float)f / 1e6f).add(" kHz");
    return add((float)f / 1000.0f).add(" Hz");
}

Text& Text::add_qty(float v, Dim d) {
    float a = v < 0 ? -v : v;
    switch (d) {
        case D_FREQ:
            if (a >= 1e6f) return add(v / 1e6f).add(" MHz");
            if (a >= 1e3f) return add(v / 1e3f).add(" kHz");
            return add(v).add(" Hz");
        case D_LENGTH:
            if (a >= 1) return add(v).add(" m");
            if (a >= 0.01f) return add(v * 100).add(" cm");
            return add(v * 1000).add(" mm");
        case D_VOLTAGE:
            if (a < 1) return add(v * 1000).add(" mV");
            return add(v).add(" V");
        case D_CURRENT:
            if (a < 1) return add(v * 1000).add(" mA");
            return add(v).add(" A");
        case D_ANGLE: return add(v).add(" deg");
        case D_TEMP: return add(v).add(" C");
        case D_RATIO: return add(v * 100).add(" %");
        case D_SPEED: return add(v).add(" m/s");
        case D_MEMORY: return add((int32_t)v).add(" bytes");
        case D_BAUD: return add((int32_t)v).add(" baud");
    }
    return add(v);
}

Text& Text::add_unit(float v, float per, const char* unit) {
    add(v / per);
    if (strcmp(unit, "%") != 0) add(" ");
    return add(unit);
}

Text Text::at(int32_t i, int line) const {
    if (i < 0 || i >= n) {
        fail("Text index out of range", line);
        return Text();
    }
    char b[2] = {s[i], 0};
    return Text(b);
}

Text Text::upper() const {
    Text t(*this);
    for (uint8_t i = 0; i < t.n; i++) t.s[i] = (char)toupper(t.s[i]);
    return t;
}

Text Text::lower() const {
    Text t(*this);
    for (uint8_t i = 0; i < t.n; i++) t.s[i] = (char)tolower(t.s[i]);
    return t;
}

Opt<float> Text::to_number() const {
    if (n == 0) return Opt<float>();
    char* end = nullptr;
    float v = (float)strtod(s, &end);
    while (end && *end == ' ') end++;
    if (end == s || (end && *end)) return Opt<float>();
    return Opt<float>(v);
}

List<Text, 8> Text::split(const char* sep) const {
    List<Text, 8> out;
    size_t k = strlen(sep);
    const char* p = s;
    while (true) {
        const char* q = k ? strstr(p, sep) : nullptr;
        Text part;
        size_t len = q ? (size_t)(q - p) : strlen(p);
        for (size_t i = 0; i < len && part.n < MUJARD_TEXT_MAX; i++) part.s[part.n++] = p[i];
        part.s[part.n] = 0;
        if (out.n < 8) out.items[out.n++] = part;
        if (!q) break;
        p = q + k;
    }
    return out;
}

Out::Out(const char* prefix) {
    if (prefix) t.add("[").add(secs(now())).add(" s] ").add(prefix).add(" ");
}

// Printed text never carries the app protocol's marker byte (docs/lang/MUJARD_APP_PROTOCOL.md): a program that prints
// text it received (a GPS line, an MQTT message) can't fake a notify or a control that way.
static void printClean(const char* s) {
    for (; *s; s++) Serial.write(*s == 0x1F ? '?' : *s);
}

void Out::end() {
    if (!serialStarted) serial_begin(115200);
    serialOutLock();
    printClean(t.s);
    Serial.write('\r');
    Serial.write('\n');
    serialOutUnlock();
}

void print(const Text& t) {
    if (!serialStarted) serial_begin(115200);
    serialOutLock();
    printClean(t.s);
    Serial.write('\r');
    Serial.write('\n');
    serialOutUnlock();
}

void log_line(const char* level, const Text& t) {
    Out o(level);
    o.put(t);
    o.end();
}

// ---- maths and smart signals

float scale_(float x, float a, float b, float c, float d) {
    if (b == a) return c;
    return c + (x - a) * (d - c) / (b - a);
}

float round_(float x, int32_t decimals) {
    float m = 1;
    for (int32_t i = 0; i < decimals; i++) m *= 10;
    return floorf(x * m + 0.5f) / m;
}

int32_t random_range(int32_t lo, int32_t hi) { return lo <= hi ? (int32_t)random(lo, (long)hi + 1) : (int32_t)random(hi, (long)lo + 1); }
float random_range(float lo, float hi) { return lo + (hi - lo) * (float)random(0, 1000001) / 1000000.0f; }

float Smooth::add(float x, int32_t count) {
    if (count < 1) count = 1;
    if (count > 32) count = 32;
    buf[i] = x;
    i = (uint8_t)((i + 1) % count);
    if (filled < count) filled++;
    float sum = 0;
    for (uint8_t k = 0; k < filled; k++) sum += buf[k];
    return sum / filled;
}

bool Hysteresis::update(float x, float on, float off) {
    if (on < off) {  // heating: on below 'on', off above 'off'
        if (x < on) state = true;
        else if (x > off) state = false;
    } else {  // cooling: on above 'on', off below 'off'
        if (x > on) state = true;
        else if (x < off) state = false;
    }
    return state;
}

bool Changed::update(float x) {
    bool c = !first && x != last;
    first = false;
    last = x;
    return c;
}

float Peak::update(float x, Time window) {
    Time t = now();
    if (!any || x >= best || t - at > window) {
        best = x;
        at = t;
        any = true;
    }
    return best;
}

bool Every::due() {
    if (!on) return false;
    Time t = now();
    if (next == 0) {
        next = t + period;
        return true;
    }
    if (t < next) return false;
    next += period;
    if (t >= next) {  // late by a whole period or more: skip the missed runs
        overruns += (uint32_t)((t - next) / period + 1);
        next = t + period;
    }
    return true;
}

// ---- outputs

void Output::init() {
    pinMode(pin_, OUTPUT);
    set(state_);  // through digitalWrite: it also disconnects a PWM timer from the pin, once
#ifdef MUJARD_DIRECT_PINS
    uint8_t port = digitalPinToPort(pin_);
    if (port != NOT_A_PIN) {
        port_ = portOutputRegister(port);
        mask_ = digitalPinToBitMask(pin_);
    }
#endif
}

void Output::set(bool v) {
    if (halted_) return;  // after a halt, outputs stay as safe_state left them
    state_ = v;
    bool high = v != inv_;
#ifdef MUJARD_DIRECT_PINS
    if (port_) {
        uint8_t sreg = SREG;  // as digitalWrite does: a fast handler may write the same port in between
        cli();
        if (high) *port_ |= mask_;
        else *port_ &= (uint8_t)~mask_;
        SREG = sreg;
        return;
    }
#endif
    digitalWrite(pin_, high ? HIGH : LOW);
}

void Output::stop() {
#ifndef MUJARD_NO_MELODY
    if (mode_ == 3) noTone(pin_);
#endif
    mode_ = 0;
}

// The effects an output can run by itself (blink, pattern, melody): each one the program doesn't use is switched off by
// mujard_config.h, its method and its part of tick() together.
#ifndef MUJARD_NO_BLINK
void Output::blink(Time period, Time duration, Time onTime) {
    stop();
    if (period <= 0) return;
    mode_ = 1;
    period_ = period;
    onTime_ = onTime < 0 ? period / 2 : onTime;
    t0_ = now();
    until_ = duration > 0 ? t0_ + duration : 0;
    set(true);
}
#endif

#ifndef MUJARD_NO_PATTERN
static const char* const MORSE[36] = {".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.",
                                      "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..", "-----", ".----", "..---", "...--", "....-", ".....", "-....", "--...",
                                      "---..", "----."};

void Output::pattern(const char* p, Time step, int32_t repeat) {
    stop();
    if (!p || !*p) return;
    mode_ = 2;
    text_ = p;
    step_ = step > 0 ? step : 150000;
    repeat_ = repeat;
    pos_ = 0;
    sub_ = 0;
    stepEnd_ = now();
}

// Next step of a pattern: "x" on / "." off for step patterns; anything else is spelled in Morse code
// (dot 1 step on, dash 3 on, 1 off inside a letter, 3 between letters, 7 between words).
bool Output::nextPatternStep(bool* on) {
    bool steps = true;
    for (const char* c = text_; *c; c++)
        if (*c != 'x' && *c != '.') steps = false;
    if (steps) {
        if (!text_[pos_]) {
            if (repeat_ > 0 && --repeat_ == 0) return false;
            pos_ = 0;
        }
        *on = text_[pos_++] == 'x';
        stepEnd_ += step_;
        return true;
    }
    // Morse: sub_ walks the current letter's code; odd sub_ values are the gaps
    for (;;) {
        char ch = text_[pos_];
        if (!ch) {
            if (repeat_ > 0 && --repeat_ == 0) return false;
            pos_ = 0;
            sub_ = 0;
            *on = false;
            stepEnd_ += 7 * step_;
            return true;
        }
        if (ch == ' ') {
            pos_++;
            *on = false;
            stepEnd_ += 4 * step_;
            return true;
        }
        int idx = -1;
        char u = (char)toupper(ch);
        if (u >= 'A' && u <= 'Z') idx = u - 'A';
        if (u >= '0' && u <= '9') idx = 26 + (u - '0');
        if (idx < 0) {
            pos_++;
            continue;
        }
        const char* code = MORSE[idx];
        uint8_t element = sub_ / 2;
        if (!code[element]) {
            pos_++;
            sub_ = 0;
            *on = false;
            stepEnd_ += 2 * step_;  // plus the 1-step gap already taken = 3 between letters
            return true;
        }
        if (sub_ % 2 == 0) {
            *on = true;
            stepEnd_ += (code[element] == '-' ? 3 : 1) * step_;
        } else {
            *on = false;
            stepEnd_ += step_;
        }
        sub_++;
        return true;
    }
}
#endif  // MUJARD_NO_PATTERN

#ifndef MUJARD_NO_MELODY
static float noteHz(const char* p, uint8_t* used) {
    static const int8_t semis[7] = {9, 11, 0, 2, 4, 5, 7};  // A B C D E F G
    char n = (char)toupper(p[0]);
    if (n < 'A' || n > 'G') {
        *used = 0;
        return 0;
    }
    int s = semis[n - 'A'];
    uint8_t k = 1;
    if (p[k] == '#') s++, k++;
    else if (p[k] == 'b') s--, k++;
    int octave = 4;
    if (p[k] >= '0' && p[k] <= '9') octave = p[k++] - '0';
    *used = k;
    int midi = (octave + 1) * 12 + s;
    return 440.0f * powf(2.0f, (midi - 69) / 12.0f);
}

void Output::melody(const char* notes, int32_t tempo) {
    stop();
    if (!notes || !*notes) return;
    mode_ = 3;
    text_ = notes;
    pos_ = 0;
    tempo_ = tempo > 0 ? tempo : 120;
    stepEnd_ = now();
}
#endif  // MUJARD_NO_MELODY

#ifndef MUJARD_NO_TONE
void Output::tone(Freq f, Time length) {
    stop();
    if (halted_) return;
#ifdef MUJARD_HAS_TONE
    if (length > 0) ::tone(pin_, (unsigned int)(f / 1000), (unsigned long)(length / 1000));
    else ::tone(pin_, (unsigned int)(f / 1000));
#endif
}
#endif

void Output::pulse(Time length) {
    stop();
    set(true);
    if (length < 16000) delayMicroseconds((unsigned int)length);
    else delay((unsigned long)(length / 1000));
    set(false);
}

void Output::tick(Time t) {
    (void)t;
    switch (mode_) {
#ifndef MUJARD_NO_BLINK
        case 1: {
            if (until_ && t >= until_) {
                mode_ = 0;
                set(false);
                return;
            }
            Time phase = (t - t0_) % period_;
            bool want = phase < onTime_;
            if (want != state_) set(want);
            return;
        }
#endif
#ifndef MUJARD_NO_PATTERN
        case 2: {
            while (t >= stepEnd_) {
                bool on;
                if (!nextPatternStep(&on)) {
                    mode_ = 0;
                    set(false);
                    return;
                }
                set(on);
            }
            return;
        }
#endif
#ifndef MUJARD_NO_MELODY
        case 3: {
            if (t < stepEnd_) return;
            while (text_[pos_] == ' ') pos_++;
            if (!text_[pos_]) {
                noTone(pin_);
                mode_ = 0;
                return;
            }
            Time beat = 60000000LL / tempo_;
            float hz = 0;
            uint8_t used = 1;
            if (text_[pos_] != '-') hz = noteHz(text_ + pos_, &used);
            pos_ += used ? used : 1;
            Time length = beat;
            if (text_[pos_] == '*' || text_[pos_] == '/') {
                char op = text_[pos_++];
                int f = 0;
                while (text_[pos_] >= '0' && text_[pos_] <= '9') f = f * 10 + (text_[pos_++] - '0');
                if (f > 0) length = op == '*' ? beat * f : beat / f;
            }
#ifdef MUJARD_HAS_TONE
            if (hz > 0) ::tone(pin_, (unsigned int)hz, (unsigned long)(length * 9 / 10 / 1000));
            else noTone(pin_);
#endif
            stepEnd_ = t + length;
            return;
        }
#endif
    }
}

void group_write(Output* items, int n, int32_t bits) {
    for (int i = 0; i < n; i++) items[i].write((bits >> (n - 1 - i)) & 1);
}

// ---- inputs

void Input::init() {
    uint8_t mode = INPUT;
    if (pull_ == 1) mode = INPUT_PULLUP;
#if defined(INPUT_PULLDOWN)
    if (pull_ == 2) mode = INPUT_PULLDOWN;
#endif
    pinMode(pin_, mode);
#ifdef MUJARD_DIRECT_PINS
    uint8_t port = digitalPinToPort(pin_);
    if (port != NOT_A_PIN) {
        port_ = portInputRegister(port);
        mask_ = digitalPinToBitMask(pin_);
    }
#endif
    lastRaw_ = stable_ = raw();
}

bool Input::raw() const {
#ifdef MUJARD_DIRECT_PINS
    if (port_) return ((*port_ & mask_) != 0) != activeLow_;  // one register read: atomic on AVR
#endif
    return (digitalRead(pin_) == HIGH) != activeLow_;
}

void Input::tick(Time t) {
    ev_ = 0;
    bool r = raw();
    if (r != lastRaw_) {
        lastRaw_ = r;
        lastChange_ = t;
    }
    prevHeld_ = heldFor_;
    if (r != stable_ && t - lastChange_ >= debounce_) {
        stable_ = r;
        ev_ |= 4;
        ev_ |= (r != activeLow_) ? 32 : 64;  // rising / falling follow the wire, not the active level (spec 6.5)
        if (r) {
            ev_ |= 1;
            pressStart_ = t;
            prevHeld_ = heldFor_ = 0;
        } else {
            ev_ |= 2;
            if (t - pressStart_ < 500000) {  // a short press: a click, or the second of a double click
                if (t - lastClick_ < 400000) {
                    ev_ |= 16;
                    lastClick_ = -1000000000LL;
                    pendingClick_ = -1;  // the second click turned the pending single into a double
                } else {
                    lastClick_ = t;
                    if (wantsDouble_) pendingClick_ = t + 400000;
                    else ev_ |= 8;
                }
            }
        }
    }
    if (pendingClick_ >= 0 && t >= pendingClick_) {  // the window passed: it was a single click (6.5)
        pendingClick_ = -1;
        ev_ |= 8;
    }
    heldFor_ = stable_ ? t - pressStart_ : 0;
}

Opt<Time> Input::measure_pulse(bool high, Time timeout) {
    unsigned long us = pulseIn(pin_, high ? HIGH : LOW, (unsigned long)timeout);
    if (us == 0) return Opt<Time>();
    return Opt<Time>((Time)us);
}

// ---- PWM

#if defined(ARDUINO_ARCH_AVR)
static const float PWM_MAX = 255.0f;
#else
static const float PWM_MAX = 255.0f;  // analogWrite's default 8-bit range on the ESP32, ESP8266 and RP2040 cores
#endif

void Pwm::init() {
    pinMode(pin_, OUTPUT);
    out(level_);
}

void Pwm::out(float r) {
    if (r < lo_ || r > hi_) {
        if (!warned_) {
            warned_ = true;
            Out o("WARN");
            o.put("PWM level limited to ").put_qty(r < lo_ ? lo_ : hi_, D_RATIO);
            o.end();
        }
        r = r < lo_ ? lo_ : hi_;
    }
    if (halted_) return;
    level_ = r;
    analogWrite(pin_, (int)(r * PWM_MAX + 0.5f));
}

void Pwm::write(float ratio) {
    fading_ = false;
    out(ratio);
}

void Pwm::fade(float to, Time over) {
    from_ = level_;
    to_ = to;
    t0_ = now();
    dur_ = over > 0 ? over : 1;
    fading_ = true;
}

void Pwm::tick(Time t) {
    if (!fading_) return;
    float k = (float)(t - t0_) / (float)dur_;
    if (k >= 1) {
        fading_ = false;
        out(to_);
    } else {
        out(from_ + (to_ - from_) * k);
    }
}

void Pwm::frequency(Freq f) {
#if defined(ARDUINO_ARCH_ESP32)
    analogWriteFrequency(pin_, (uint32_t)(f / 1000));
#elif defined(ARDUINO_ARCH_ESP8266) || defined(ARDUINO_ARCH_RP2040)
    analogWriteFreq((uint32_t)(f / 1000));
#else
    (void)f;  // AVR: the timers' frequency is fixed; the checker refuses frequency() there
#endif
}

// ---- analog

#if defined(MUJARD_TEST_ADC_MAX)  // a test build on the PC behaves like the program's board (mujard test)
static const float ADC_MAX = MUJARD_TEST_ADC_MAX, VREF = MUJARD_TEST_VREF;
#elif defined(ARDUINO_ARCH_ESP32)
static const float ADC_MAX = 4095.0f, VREF = 3.3f;
#elif defined(ARDUINO_ARCH_AVR)
static const float ADC_MAX = 1023.0f, VREF = 5.0f;
#else
static const float ADC_MAX = 1023.0f, VREF = 3.3f;
#endif

int32_t Analog::value() const { return analogRead(pin_); }

float Analog::volts() const {
#if defined(ARDUINO_ARCH_ESP32)
    return analogReadMilliVolts(pin_) / 1000.0f;  // calibrated by the core
#else
    return (float)analogRead(pin_) * (ref_ > 0 ? ref_ : VREF) / ADC_MAX;
#endif
}

float Analog::percent() const { return (float)analogRead(pin_) / ADC_MAX; }

int32_t Analog::rawFor(float volts) const {
    float r = volts / (ref_ > 0 ? ref_ : VREF) * ADC_MAX;
    return (int32_t)(r < 0 ? 0 : (r > ADC_MAX ? ADC_MAX : r + 0.5f));
}

int32_t Analog::rawForRatio(float ratio) const {
    float r = ratio * ADC_MAX;
    return (int32_t)(r < 0 ? 0 : (r > ADC_MAX ? ADC_MAX : r + 0.5f));
}

int32_t Battery::rawFor(float volts) const {
    // the pin's raw reading that produces 'volts' on the divider's input side -
    // volts() then applies the divider and the discharge curve, the real path a test exercises
    float r = volts / VREF * ADC_MAX;
    return (int32_t)(r < 0 ? 0 : (r > ADC_MAX ? ADC_MAX : r + 0.5f));
}

int32_t Battery::rawForRatio(float ratio) const {
    float r = ratio * ADC_MAX;
    return (int32_t)(r < 0 ? 0 : (r > ADC_MAX ? ADC_MAX : r + 0.5f));
}

void Analog::reference(float volts) {
#if defined(ARDUINO_ARCH_AVR)
#if defined(INTERNAL1V1)  // the Mega names its internal references by voltage
    if (volts > 1.0f && volts < 1.2f) analogReference(INTERNAL1V1);
    else if (volts > 2.5f && volts < 2.6f) analogReference(INTERNAL2V56);
#else
    if (volts > 1.0f && volts < 1.2f) analogReference(INTERNAL);
#endif
    else analogReference(DEFAULT);
    ref_ = volts;
#else
    (void)volts;
#endif
}

// ---- battery

// Resting voltage of one cell -> charge, highest first. Typical curves: a 1S LiPo, and an alkaline AA cell.
struct CurvePoint {
    float volts, charge;
};
static const CurvePoint LIPO_CURVE[] = {{4.20f, 1.00f}, {4.11f, 0.90f}, {4.02f, 0.80f}, {3.95f, 0.70f}, {3.87f, 0.60f}, {3.84f, 0.50f},
                                        {3.80f, 0.40f}, {3.77f, 0.30f}, {3.73f, 0.20f}, {3.69f, 0.10f}, {3.27f, 0.00f}};
static const CurvePoint AA_CURVE[] = {{1.60f, 1.00f}, {1.50f, 0.85f}, {1.40f, 0.60f}, {1.30f, 0.35f}, {1.20f, 0.15f}, {1.10f, 0.05f}, {1.00f, 0.00f}};

float battery_charge(BatteryType type, float v) {
    const CurvePoint* c = type == BAT_LIPO_1S ? LIPO_CURVE : AA_CURVE;
    int n = type == BAT_LIPO_1S ? (int)(sizeof(LIPO_CURVE) / sizeof(LIPO_CURVE[0])) : (int)(sizeof(AA_CURVE) / sizeof(AA_CURVE[0]));
    if (v != v) return 0.0f;  // NaN: no reading, report empty
    if (v >= c[0].volts) return 1.0f;
    for (int i = 1; i < n; i++) {
        if (v >= c[i].volts) {
            float k = (v - c[i].volts) / (c[i - 1].volts - c[i].volts);
            return c[i].charge + k * (c[i - 1].charge - c[i].charge);
        }
    }
    return 0.0f;
}

float Battery::volts() const {
    // the average of 8 readings: the divider's high resistance makes single readings noisy
    float sum = 0;
    for (int i = 0; i < 8; i++) {
#if defined(ARDUINO_ARCH_ESP32)
        sum += analogReadMilliVolts(pin_) / 1000.0f;
#else
        sum += (float)analogRead(pin_) * VREF / ADC_MAX;
#endif
    }
    return sum / 8.0f * divider_;
}

float Battery::percent() const {
    int cells = type_ == BAT_AA_X3 ? 3 : type_ == BAT_AA_X4 ? 4 : 1;
    return battery_charge(type_, volts() / (float)cells);
}

// ---- the board

#if defined(ARDUINO_ARCH_AVR)
// Optiboot passes MCUSR to the application in r2. Catch it before avr-gcc's init code can reuse that register.
static uint8_t avrResetFlags __attribute__((section(".noinit")));
__attribute__((naked, used, section(".init0")))
static void avrCaptureResetFlags() { __asm__ volatile("sts %0, r2" : "=m"(avrResetFlags)); }
#endif

const char* board_chip() {
#if defined(ARDUINO_ARCH_ESP32)
    return ESP.getChipModel();
#elif defined(ARDUINO_ARCH_ESP8266)
    return "ESP8266";
#elif defined(ARDUINO_ARCH_RP2040)
    return "RP2040";
#elif defined(__AVR_ATmega2560__)
    return "ATmega2560";
#elif defined(__AVR_ATmega32U4__)
    return "ATmega32U4";
#else
    return "unknown";  // generated programs use the board catalog's compile-time Family instead
#endif
}

Text board_id() {
    Text t;
#if defined(ARDUINO_ARCH_ESP32)
    uint64_t mac = ESP.getEfuseMac();
    char b[20];
    snprintf(b, sizeof(b), "%04X-%08lX", (unsigned)(mac >> 32), (unsigned long)mac);
    t.add(b);
#elif defined(ARDUINO_ARCH_ESP8266)
    t.add((int32_t)ESP.getChipId());
#elif defined(ARDUINO_ARCH_RP2040)
    pico_unique_board_id_t id;
    pico_get_unique_board_id(&id);
    char b[17];
    char* p = b;
    for (uint8_t i = 0; i < PICO_UNIQUE_BOARD_ID_SIZE_BYTES; i++) p += sprintf(p, "%02X", id.id[i]);
    t.add(b);
#elif defined(ARDUINO_ARCH_STM32)
    char b[25];  // STM32F103's 96-bit unique device ID
    snprintf(b, sizeof(b), "%08lX%08lX%08lX", (unsigned long)HAL_GetUIDw0(), (unsigned long)HAL_GetUIDw1(), (unsigned long)HAL_GetUIDw2());
    t.add(b);
#elif defined(ARDUINO_ARCH_SAMD)
    // SAMD21's 128-bit serial number lives at these four addresses (the words aren't contiguous in address order).
    char b[33];
    snprintf(b, sizeof(b), "%08lX%08lX%08lX%08lX",
             (unsigned long)*(volatile uint32_t*)0x0080A00C, (unsigned long)*(volatile uint32_t*)0x0080A040,
             (unsigned long)*(volatile uint32_t*)0x0080A044, (unsigned long)*(volatile uint32_t*)0x0080A048);
    t.add(b);
#elif defined(ARDUINO_ARCH_MBED) || defined(ARDUINO_ARCH_MBED_NANO)
    char b[17];  // nRF52840 FICR DEVICEID is the factory 64-bit chip ID
    snprintf(b, sizeof(b), "%08lX%08lX", (unsigned long)NRF_FICR->DEVICEID[0], (unsigned long)NRF_FICR->DEVICEID[1]);
    t.add(b);
#elif defined(ARDUINO_ARCH_RENESAS_UNO)
    // RA4M1's four 32-bit unique-ID registers follow FMIFRT (0x407FB19C): UIDR0 is FMIFRT+0x14.
    const volatile uint32_t* uid = (const volatile uint32_t*)0x407FB1B0;
    char b[33];
    snprintf(b, sizeof(b), "%08lX%08lX%08lX%08lX",
             (unsigned long)uid[0], (unsigned long)uid[1], (unsigned long)uid[2], (unsigned long)uid[3]);
    t.add(b);
#else
    t.add("unknown");
#endif
    return t;
}

ResetReason board_reset_reason() {
#if defined(ARDUINO_ARCH_AVR)
    if (avrResetFlags & (1 << WDRF)) return RESET_WATCHDOG;
    if (avrResetFlags & (1 << BORF)) return RESET_BROWNOUT;
    if (avrResetFlags & (1 << EXTRF)) return RESET_RESET_BUTTON;
    if (avrResetFlags & (1 << PORF)) return RESET_POWER_ON;
    return RESET_POWER_ON;
#elif defined(ARDUINO_ARCH_ESP32)
    switch (esp_reset_reason()) {
        case ESP_RST_EXT: return RESET_RESET_BUTTON;
        case ESP_RST_SW: return RESET_SOFTWARE;
        case ESP_RST_PANIC: return RESET_CRASH;
        case ESP_RST_INT_WDT:
        case ESP_RST_TASK_WDT:
        case ESP_RST_WDT: return RESET_WATCHDOG;
        case ESP_RST_DEEPSLEEP: return RESET_WAKE;
        case ESP_RST_BROWNOUT: return RESET_BROWNOUT;
        case ESP_RST_POWERON:
        default: return RESET_POWER_ON;
    }
#elif defined(ARDUINO_ARCH_ESP8266)
    rst_info* i = ESP.getResetInfoPtr();
    if (i == nullptr) return RESET_POWER_ON;
    switch (i->reason) {
        case REASON_WDT_RST:
        case REASON_SOFT_WDT_RST: return RESET_WATCHDOG;
        case REASON_EXCEPTION_RST: return RESET_CRASH;
        case REASON_SOFT_RESTART: return RESET_SOFTWARE;
        case REASON_DEEP_SLEEP_AWAKE: return RESET_WAKE;
        case REASON_EXT_SYS_RST: return RESET_RESET_BUTTON;
        case REASON_DEFAULT_RST:
        default: return RESET_POWER_ON;
    }
#elif defined(ARDUINO_ARCH_STM32)
    static uint32_t flags;
    static bool haveFlags;
    if (!haveFlags) {
        flags = RCC->CSR;
        __HAL_RCC_CLEAR_RESET_FLAGS();
        haveFlags = true;
    }
    if (flags & RCC_CSR_PORRSTF) return RESET_POWER_ON;  // POR can also set PINRSTF on some parts
#ifdef RCC_CSR_BORRSTF
    if (flags & RCC_CSR_BORRSTF) return RESET_BROWNOUT;
#endif
    if (flags & (RCC_CSR_IWDGRSTF | RCC_CSR_WWDGRSTF)) return RESET_WATCHDOG;
    if (flags & RCC_CSR_SFTRSTF) return RESET_SOFTWARE;
    if (flags & RCC_CSR_PINRSTF) return RESET_RESET_BUTTON;
    return RESET_POWER_ON;
#elif defined(ARDUINO_ARCH_SAMD)
    uint8_t flags = PM->RCAUSE.reg;
    if (flags & PM_RCAUSE_WDT) return RESET_WATCHDOG;
    if (flags & PM_RCAUSE_SYST) return RESET_SOFTWARE;
    if (flags & PM_RCAUSE_EXT) return RESET_RESET_BUTTON;
    if (flags & (PM_RCAUSE_BOD12 | PM_RCAUSE_BOD33)) return RESET_BROWNOUT;
    if (flags & PM_RCAUSE_POR) return RESET_POWER_ON;
    return RESET_POWER_ON;
#elif defined(ARDUINO_ARCH_MBED) || defined(ARDUINO_ARCH_MBED_NANO)
    static uint32_t flags;
    static bool haveFlags;
    if (!haveFlags) {
        flags = NRF_POWER->RESETREAS;
        NRF_POWER->RESETREAS = 0xFFFFFFFFu;  // write-1-to-clear; keep this boot's reason stable for later reads
        haveFlags = true;
    }
    if (flags & POWER_RESETREAS_LOCKUP_Msk) return RESET_CRASH;
    if (flags & POWER_RESETREAS_DOG_Msk) return RESET_WATCHDOG;
    if (flags & POWER_RESETREAS_SREQ_Msk) return RESET_SOFTWARE;
    if (flags & (POWER_RESETREAS_OFF_Msk | POWER_RESETREAS_LPCOMP_Msk | POWER_RESETREAS_DIF_Msk |
                 POWER_RESETREAS_NFC_Msk | POWER_RESETREAS_VBUS_Msk)) return RESET_WAKE;
    if (flags & POWER_RESETREAS_RESETPIN_Msk) return RESET_RESET_BUTTON;
    return RESET_POWER_ON;
#elif defined(ARDUINO_ARCH_RENESAS_UNO)
    uint8_t r0 = R_SYSTEM->RSTSR0;
    uint16_t r1 = R_SYSTEM->RSTSR1;
    if (r0 & 0x0E) return RESET_BROWNOUT;  // LVD0/LVD1/LVD2 reset flags
    if (r1 & 0x03) return RESET_WATCHDOG;  // IWDT and WDT reset flags
    if (r1 & 0x04) return RESET_SOFTWARE;
    if (r1 & 0x1F00) return RESET_CRASH;   // bus and RAM error reset flags
    if (r0 & 0x01) return RESET_POWER_ON;  // PORF
    return RESET_RESET_BUTTON;             // RES# leaves no dedicated flag on RA4M1
#else
    return RESET_POWER_ON;  // only emitted where the generator has a real backend
#endif
}

// newlib's heap-break syscall: the samd/stm32/renesas cores link it but don't declare it, so declare it here once.
#if defined(ARDUINO_ARCH_RENESAS_UNO) || defined(ARDUINO_ARCH_SAMD) || defined(ARDUINO_ARCH_STM32)
extern "C" void* _sbrk(int incr);
#endif

int32_t free_memory() {
#if defined(MUJARD_HOST_TEST)
#ifdef MUJARD_TEST_FREE_MEMORY
    return (int32_t)MUJARD_TEST_FREE_MEMORY;
#else
    return 2048;  // standalone runtime tests without board test defines
#endif
#elif defined(ARDUINO_ARCH_AVR)
    extern int __heap_start, *__brkval;
    int v;
    return (int32_t)((int)&v - (__brkval == 0 ? (int)&__heap_start : (int)__brkval));
#elif defined(ARDUINO_ARCH_ESP32) || defined(ARDUINO_ARCH_ESP8266)
    return (int32_t)ESP.getFreeHeap();
#elif defined(ARDUINO_ARCH_RP2040)
    return (int32_t)rp2040.getFreeHeap();
#elif defined(ARDUINO_ARCH_RENESAS_UNO) || defined(ARDUINO_ARCH_SAMD) || defined(ARDUINO_ARCH_STM32)
    char top;  // stack top minus the heap break = the free bytes between them
    return (int32_t)((intptr_t)&top - (intptr_t)_sbrk(0));
#elif defined(ARDUINO_ARCH_MBED) || defined(ARDUINO_ARCH_MBED_NANO)
    mbed_stats_heap_t s;
    mbed_stats_heap_get(&s);
    return s.current_size >= s.reserved_size ? 0 : (int32_t)(s.reserved_size - s.current_size);
#else
    return -1;  // only emitted where the generator has a real backend
#endif
}

Opt<float> board_voltage() {
#if defined(MUJARD_HOST_TEST)
#ifdef MUJARD_TEST_BOARD_VOLTAGE
    return Opt<float>((float)MUJARD_TEST_BOARD_VOLTAGE);
#else
    return Opt<float>();
#endif
#elif defined(ARDUINO_ARCH_AVR)
    uint8_t old = ADMUX;
#if defined(__AVR_ATmega32U4__) || defined(__AVR_ATmega2560__) || defined(__AVR_ATmega1280__)
    const uint8_t mux = 30;  // the 1.1 V bandgap channel on the 32U4/2560
#else
    const uint8_t mux = 14;  // ATmega328P and family
#endif
    ADMUX = _BV(REFS0) | mux;  // Vcc as the reference, bandgap as the input
    delayMicroseconds(250);    // let the mux and reference settle
    ADCSRA |= _BV(ADSC);       // discard the first conversion after changing channels
    while (ADCSRA & _BV(ADSC)) {
    }
    (void)ADC;
    uint32_t sum = 0;
    for (uint8_t i = 0; i < 8; i++) {
        ADCSRA |= _BV(ADSC);
        while (ADCSRA & _BV(ADSC)) {
        }
        sum += ADC;
    }
    ADMUX = old;  // analogRead() can use the channel and reference the program chose
    float raw = (float)sum / 8.0f;
    if (raw <= 0) return Opt<float>();
    return Opt<float>(1.1f * 1023.0f / raw);
#elif defined(ARDUINO_ARCH_RP2040)
    // Pico / Pico W: GP29 is VSYS through a divider of three. RP2350-USB-A doesn't enter here: its GP29 is a
    // normal analog pin, so its generated program takes the 'none' path instead.
    return Opt<float>((float)analogRead(29) * 3.0f * VREF / ADC_MAX);
#elif defined(ARDUINO_ARCH_MBED) || defined(ARDUINO_ARCH_MBED_NANO)
    // nRF52840: SAADC's VDD input with the internal 0.6 V reference and 1/6 gain gives a 3.6 V range.
    volatile int16_t result = 0;
    uint32_t oldEnable = NRF_SAADC->ENABLE;
    uint32_t oldResolution = NRF_SAADC->RESOLUTION;
    uint32_t oldOversample = NRF_SAADC->OVERSAMPLE;
    uint32_t oldSamplerate = NRF_SAADC->SAMPLERATE;
    uint32_t oldConfig = NRF_SAADC->CH[0].CONFIG;
    uint32_t oldPselp = NRF_SAADC->CH[0].PSELP;
    uint32_t oldPseln = NRF_SAADC->CH[0].PSELN;
    uint32_t oldResultPtr = NRF_SAADC->RESULT.PTR;
    uint32_t oldResultMaxcnt = NRF_SAADC->RESULT.MAXCNT;

    NRF_SAADC->ENABLE = SAADC_ENABLE_ENABLE_Disabled << SAADC_ENABLE_ENABLE_Pos;
    NRF_SAADC->RESOLUTION = SAADC_RESOLUTION_VAL_14bit << SAADC_RESOLUTION_VAL_Pos;
    NRF_SAADC->OVERSAMPLE = SAADC_OVERSAMPLE_OVERSAMPLE_Bypass << SAADC_OVERSAMPLE_OVERSAMPLE_Pos;
    NRF_SAADC->SAMPLERATE = SAADC_SAMPLERATE_MODE_Task << SAADC_SAMPLERATE_MODE_Pos;
    NRF_SAADC->CH[0].CONFIG =
        (SAADC_CH_CONFIG_RESP_Bypass << SAADC_CH_CONFIG_RESP_Pos) |
        (SAADC_CH_CONFIG_RESN_Bypass << SAADC_CH_CONFIG_RESN_Pos) |
        (SAADC_CH_CONFIG_GAIN_Gain1_6 << SAADC_CH_CONFIG_GAIN_Pos) |
        (SAADC_CH_CONFIG_REFSEL_Internal << SAADC_CH_CONFIG_REFSEL_Pos) |
        (SAADC_CH_CONFIG_TACQ_10us << SAADC_CH_CONFIG_TACQ_Pos) |
        (SAADC_CH_CONFIG_MODE_SE << SAADC_CH_CONFIG_MODE_Pos) |
        (SAADC_CH_CONFIG_BURST_Disabled << SAADC_CH_CONFIG_BURST_Pos);
    NRF_SAADC->CH[0].PSELP = SAADC_CH_PSELP_PSELP_VDD << SAADC_CH_PSELP_PSELP_Pos;
    NRF_SAADC->CH[0].PSELN = SAADC_CH_PSELN_PSELN_NC << SAADC_CH_PSELN_PSELN_Pos;
    NRF_SAADC->RESULT.PTR = (uint32_t)(uintptr_t)&result;
    NRF_SAADC->RESULT.MAXCNT = 1;
    NRF_SAADC->EVENTS_STARTED = 0;
    NRF_SAADC->EVENTS_END = 0;
    NRF_SAADC->EVENTS_STOPPED = 0;
    NRF_SAADC->ENABLE = SAADC_ENABLE_ENABLE_Enabled << SAADC_ENABLE_ENABLE_Pos;
    NRF_SAADC->TASKS_START = 1;
    while (!NRF_SAADC->EVENTS_STARTED) {
    }
    NRF_SAADC->EVENTS_STARTED = 0;
    NRF_SAADC->TASKS_SAMPLE = 1;
    while (!NRF_SAADC->EVENTS_END) {
    }
    NRF_SAADC->TASKS_STOP = 1;
    while (!NRF_SAADC->EVENTS_STOPPED) {
    }

    float v = (float)result * 3.6f / 16383.0f;
    NRF_SAADC->ENABLE = SAADC_ENABLE_ENABLE_Disabled << SAADC_ENABLE_ENABLE_Pos;
    NRF_SAADC->RESULT.PTR = oldResultPtr;
    NRF_SAADC->RESULT.MAXCNT = oldResultMaxcnt;
    NRF_SAADC->CH[0].PSELN = oldPseln;
    NRF_SAADC->CH[0].PSELP = oldPselp;
    NRF_SAADC->CH[0].CONFIG = oldConfig;
    NRF_SAADC->SAMPLERATE = oldSamplerate;
    NRF_SAADC->OVERSAMPLE = oldOversample;
    NRF_SAADC->RESOLUTION = oldResolution;
    NRF_SAADC->ENABLE = oldEnable;
    return v > 0 ? Opt<float>(v) : Opt<float>();
#elif defined(ARDUINO_ARCH_SAMD)
    // SAMD21: SCALEDIOVCC is IOVCC/4. With the 1 V internal reference, VDD = raw * 4 / 4095.
    uint16_t oldCtrlA = ADC->CTRLA.reg;
    uint16_t oldCtrlB = ADC->CTRLB.reg;
    uint8_t oldRef = ADC->REFCTRL.reg;
    uint32_t oldInput = ADC->INPUTCTRL.reg;
    bool wasEnabled = (oldCtrlA & ADC_CTRLA_ENABLE) != 0;
    ADC->CTRLA.bit.ENABLE = 0;
    while (ADC->STATUS.bit.SYNCBUSY) {
    }
    ADC->CTRLB.reg = ADC_CTRLB_PRESCALER_DIV256 | ADC_CTRLB_RESSEL_12BIT;
    ADC->REFCTRL.reg = ADC_REFCTRL_REFSEL_INT1V;
    ADC->INPUTCTRL.reg = ADC_INPUTCTRL_GAIN_1X | ADC_INPUTCTRL_MUXNEG_GND |
                         ADC_INPUTCTRL_MUXPOS_SCALEDIOVCC;
    ADC->CTRLA.bit.ENABLE = 1;
    while (ADC->STATUS.bit.SYNCBUSY) {
    }
    uint32_t sum = 0;
    for (uint8_t i = 0; i < 9; i++) {
        ADC->INTFLAG.bit.RESRDY = 1;
        ADC->SWTRIG.bit.START = 1;
        while (!ADC->INTFLAG.bit.RESRDY) {
        }
        if (i > 0) sum += ADC->RESULT.reg;  // discard the first sample after changing channels/reference
    }
    float v = ((float)sum / 8.0f) * 4.0f / 4095.0f;
    ADC->CTRLA.bit.ENABLE = 0;
    while (ADC->STATUS.bit.SYNCBUSY) {
    }
    ADC->CTRLB.reg = oldCtrlB;
    ADC->REFCTRL.reg = oldRef;
    ADC->INPUTCTRL.reg = oldInput;
    if (wasEnabled) {
        ADC->CTRLA.bit.ENABLE = 1;
        while (ADC->STATUS.bit.SYNCBUSY) {
        }
    }
    return v > 0 ? Opt<float>(v) : Opt<float>();
#elif defined(ARDUINO_ARCH_STM32)
#ifdef AVREF
    int raw = analogRead(AVREF);  // STM32F1 VREFINT is nominally 1.2 V against VDDA
    if (raw <= 0) return Opt<float>();
    return Opt<float>(1.2f * (ADC_MAX + 1.0f) / (float)raw);
#else
    return Opt<float>();
#endif
#elif defined(ARDUINO_ARCH_RENESAS_UNO)
#ifdef AVCC_MEASURE_PIN
    analogReference(AR_DEFAULT);   // the core measures the AVCC pin in default-reference mode
    float v = analogReference();
    return v > 0 && v == v ? Opt<float>(v) : Opt<float>();
#else
    return Opt<float>();
#endif
#else
    return Opt<float>();
#endif
}

void restart() {
    saved_flush();
#if defined(ARDUINO_ARCH_ESP32) || defined(ARDUINO_ARCH_ESP8266)
    ESP.restart();
#elif defined(ARDUINO_ARCH_RP2040)
    rp2040.reboot();
#elif defined(ARDUINO_ARCH_AVR)
    wdt_enable(WDTO_15MS);
    for (;;) {
    }
#elif defined(__ARM_ARCH)
    NVIC_SystemReset();  // CMSIS reset on the Renesas, nRF52, SAMD and STM32 targets
    for (;;) {
    }
#endif
}

// ---- sleep (spec 6.10): stop the whole program for up to t, or until the wake pin's edge.

#if defined(ARDUINO_ARCH_AVR)
static volatile bool wakeFired = false;
static void wakeIsr() { wakeFired = true; }
ISR(WDT_vect) {}  // the interrupt only has to wake the chip

// One watchdog interval of power-down sleep. The timers stop, so micros() doesn't see this time.
static void sleepChunk(uint8_t wdp) {
    cli();
    wdt_reset();
    MCUSR &= ~(1 << WDRF);
    WDTCSR = (1 << WDCE) | (1 << WDE);
    WDTCSR = (1 << WDIE) | wdp;
    sei();
    set_sleep_mode(SLEEP_MODE_PWR_DOWN);
    sleep_enable();
    sleep_cpu();
    sleep_disable();
}

static Time avrSleep(Time t, int8_t pin, uint8_t mode) {
    // WDP 0..9: 16 ms, 32, 64, 128, 256, 512, 1 s, 2, 4, 8
    static const uint32_t chunkMs[10] = {16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192};
    wakeFired = false;
    int ip = -1;
#if defined(digitalPinToInterrupt) && defined(NOT_AN_INTERRUPT)
    if (pin >= 0) {
        ip = digitalPinToInterrupt((uint8_t)pin);
        if (ip == NOT_AN_INTERRUPT) ip = -1;
    }
#endif
    if (ip >= 0) attachInterrupt((uint8_t)ip, wakeIsr, mode);
    uint64_t slept = 0, want = (uint64_t)t;
    while (slept < want && !wakeFired) {
        int8_t w = 9;
        while (w > 0 && (uint64_t)chunkMs[w] * 1000 > want - slept) w--;
        sleepChunk((uint8_t)w);
        slept += (uint64_t)chunkMs[w] * 1000;
    }
    if (ip >= 0) detachInterrupt((uint8_t)ip);
    wdt_disable();
    return (Time)slept;
}
#endif

void power_sleep(Time t, int8_t pin, uint8_t mode) {
    saved_flush();  // pending saved-var changes are written before the sleep (spec 11)
    if (t <= 0) return;
#if defined(ARDUINO_ARCH_ESP32)
    esp_sleep_enable_timer_wakeup((uint64_t)t);
    if (pin >= 0 && (mode == HIGH || mode == LOW)) {
#if SOC_PM_SUPPORT_EXT0_WAKEUP
        esp_sleep_enable_ext0_wakeup((gpio_num_t)pin, mode);
#else
        // no ext0 on the C3/C6: the same level wakeup through the GPIO wakeup path
        gpio_wakeup_enable((gpio_num_t)pin, mode == LOW ? GPIO_INTR_LOW_LEVEL : GPIO_INTR_HIGH_LEVEL);
        esp_sleep_enable_gpio_wakeup();
#endif
    }
    esp_light_sleep_start();  // esp_timer keeps counting through light sleep: the clock sees it
#elif defined(ARDUINO_ARCH_AVR)
    sleptUs += avrSleep(t, pin, mode);
#else
    (void)pin;
    (void)mode;
    delay((unsigned long)(t / 1000));  // no light-sleep API on this core: the program still pauses
#endif
}

void power_deep_sleep(Time t, int8_t pin, uint8_t mode) {
    saved_flush();
#if defined(ARDUINO_ARCH_ESP32)
    esp_sleep_enable_timer_wakeup((uint64_t)t);
    if (pin >= 0 && (mode == HIGH || mode == LOW)) {
#if SOC_PM_SUPPORT_EXT0_WAKEUP
        esp_sleep_enable_ext0_wakeup((gpio_num_t)pin, mode);
#else
        // no ext0 on the C3/C6: their deep-sleep wakeup is a GPIO level mask instead
        esp_deep_sleep_enable_gpio_wakeup(1ULL << pin, mode == LOW ? ESP_GPIO_WAKEUP_GPIO_LOW : ESP_GPIO_WAKEUP_GPIO_HIGH);
#endif
    }
    esp_deep_sleep_start();  // does not return: the board restarts on waking
#elif defined(ARDUINO_ARCH_ESP8266)
    ESP.deepSleep((uint32_t)t);  // GPIO16->RST restarts the board on waking; if it can't sleep it returns
    for (;;) {
    }
#elif defined(ARDUINO_ARCH_AVR)
    avrSleep(t, pin, mode);
    restart();  // restart on waking (a watchdog reset here: no 'wake' reason on this chip)
#else
    (void)pin;
    (void)mode;
    delay((unsigned long)(t / 1000));
    restart();
#endif
}

void watchdog(Time t) {
    wdtPeriod = t;
#if defined(ARDUINO_ARCH_AVR)
    uint8_t to = WDTO_8S;
    if (t <= 250000) to = WDTO_250MS;
    else if (t <= 500000) to = WDTO_500MS;
    else if (t <= 1000000) to = WDTO_1S;
    else if (t <= 2000000) to = WDTO_2S;
    else if (t <= 4000000) to = WDTO_4S;
    wdt_enable(to);
#elif defined(ARDUINO_ARCH_ESP32)
    esp_task_wdt_config_t cfg = {};
    cfg.timeout_ms = (uint32_t)(t / 1000);
    cfg.trigger_panic = true;
    esp_task_wdt_reconfigure(&cfg);
    esp_task_wdt_add(NULL);
#elif defined(ARDUINO_ARCH_ESP8266)
    ESP.wdtEnable((uint32_t)(t / 1000));
#elif defined(ARDUINO_ARCH_RP2040)
    rp2040.wdt_begin((uint32_t)(t / 1000));
#endif
}

// ---- the app: show, plot, notify, control, ask (docs/lang/MUJARD_APP_PROTOCOL.md)

#ifdef MUJARD_NO_APP
bool app_on() { return false; }  // a program without app built-ins never talks to the app
#else
static bool appOn_ = false;
static const char* appProgram_ = "";
static uint32_t appHash_ = 0;

bool app_on() { return appOn_; }

static void appStart() {
    if (!serialStarted) serial_begin(115200);
    appOn_ = true;
}

#ifdef MUJARD_APP_MIRROR
// Extra links for the app's messages (mujard.ble, the Wi-Fi socket): each finished line, LF included, goes to each
// registered transport - BLE and Wi-Fi can both carry the dashboard at once.
static void (*appMirrors_[2])(const char* line, uint16_t n) = {nullptr, nullptr};
void app_mirror_add(void (*fn)(const char* line, uint16_t n)) {
    for (uint8_t i = 0; i < 2; i++) {
        if (appMirrors_[i] == nullptr) { appMirrors_[i] = fn; return; }
    }
}
static bool appMirrorOn() { return appMirrors_[0] != nullptr || appMirrors_[1] != nullptr; }
static char appMirror_[248];
#endif

/** One message: 0x1F "M1 kind", then TAB-separated escaped fields, cut at 240 bytes. */
struct AppOut {
    uint16_t n;
#ifdef MUJARD_APP_MIRROR
    uint16_t m = 0;
#endif
    void put(char c) {
        Serial.write(c);
#ifdef MUJARD_APP_MIRROR
        if (appMirrorOn() && m < sizeof(appMirror_)) appMirror_[m++] = c;
#endif
    }
    void puts(const char* s) {
        for (; *s; s++) put(*s);
    }
    explicit AppOut(const char* kind) : n(0) {
        appStart();
        serialOutLock();
        put((char)0x1F);
        puts("M1 ");
        puts(kind);
        n = (uint16_t)(4 + strlen(kind));
    }
    AppOut& field(const char* s) {
        put('\t');
        n++;
        for (; s && *s && n < 236; s++) {
            char c = *s;
            const char* e = c == '\\' ? "\\\\" : c == '\t' ? "\\t" : c == '\n' ? "\\n" : c == '\r' ? "\\r" : c == 0x1F ? "\\u" : nullptr;
            if (e) {
                puts(e);
                n += 2;
            } else {
                put(c);
                n++;
            }
        }
        return *this;
    }
    AppOut& num(float v) {
        char b[24];
        if (v != v || v > 3.4e38f || v < -3.4e38f) v = 0;  // only finite numbers on the wire
        floatText(v, 3, b);
        return field(b);
    }
    AppOut& whole(int32_t v) {
        char b[14];
        iToStr(v, b);
        return field(b);
    }
    void end() {
        put('\n');
        serialOutUnlock();
#ifdef MUJARD_APP_MIRROR
        if (m > 0 && appMirror_[m - 1] == '\n') {
            for (uint8_t i = 0; i < 2; i++) if (appMirrors_[i] != nullptr) appMirrors_[i](appMirror_, m);
        }
#endif
    }
};

// controls
struct AppControl {
    const char* label;
    const char* unit;
    void* var;
    float min, max, step, per, sent;
    uint8_t type;
    uint8_t kind;     // 0 switch, 1 slider, then the held ones: 2 button, 3 stick x, 4 stick y (AppHoldKind)
    uint32_t heldAt;  // a held control: when the app last refreshed it (ms, appMs_)
};
#if defined(ARDUINO_ARCH_AVR)
static const uint8_t APP_CONTROLS = 8;
#else
static const uint8_t APP_CONTROLS = 16;
#endif
static AppControl controls_[APP_CONTROLS];
static uint8_t controlCount_ = 0;
static Time nextValueCheck_ = 0;

void app_control(const char* label, bool isSwitch, float min, float max, float step, const char* unit, float per, void* var, uint8_t type) {
    if (controlCount_ >= APP_CONTROLS) return;
    AppControl& c = controls_[controlCount_++];
    c.label = label;
    c.kind = isSwitch ? 0 : 1;
    c.heldAt = 0;
    c.min = isSwitch ? 0 : min;
    c.max = isSwitch ? 1 : max;
    c.step = isSwitch ? 1 : step;
    c.unit = unit ? unit : "";
    c.per = per > 0 ? per : 1;
    c.var = var;
    c.type = type;
    c.sent = 0;
}

/** The variable's value in the control's display unit. */
static float controlValue(const AppControl& c) {
    switch (c.type) {
    case APP_BOOL:
        return *(bool*)c.var ? 1 : 0;
    case APP_INT:
        return (float)*(int32_t*)c.var / c.per;
    case APP_I64:
        return (float)*(int64_t*)c.var / c.per;
    default:
        return *(float*)c.var / c.per;
    }
}

static const char* const kControlKinds[] = {"switch", "slider", "button", "stick-x", "stick-y"};

static void sendControl(uint8_t i) {
    AppControl& c = controls_[i];
    c.sent = controlValue(c);
    AppOut o("control");
    o.whole(i + 1).field(c.label).field(kControlKinds[c.kind < 5 ? c.kind : 1]).num(c.min).num(c.max).num(c.step).field(c.unit).num(c.sent);
    o.end();
}

static void sendHello() {
    char hash[9];
    for (int k = 0; k < 8; k++) hash[k] = "0123456789abcdef"[(appHash_ >> (28 - 4 * k)) & 15];
    hash[8] = 0;
    AppOut o("hello");
    o.field(appProgram_).field(board_name()).field("M1").field(hash).field(halted_ ? "1" : "0");
    o.end();
    for (uint8_t i = 0; i < controlCount_; i++) sendControl(i);
}

void app_begin(const char* program, uint32_t hash) {
    appProgram_ = program;
    appHash_ = hash;
    appStart();
    sendHello();
}

void app_control_hold(uint8_t kind) {
    if (controlCount_ == 0 || kind < APP_HOLD_BUTTON || kind > APP_HOLD_STICK_Y) return;
    controls_[controlCount_ - 1].kind = kind;
}

static const uint32_t APP_HOLD_TIMEOUT_MS = 500;  // the app refreshes a held control every 200 ms
static uint32_t appMs_ = 0;                       // the app clock: the pass's time in ms, set by appTick

/** A held control's rest value in its display unit: 0, or the range end nearest to it. */
static float controlRest(const AppControl& c) { return c.min > 0 ? c.min : c.max < 0 ? c.max : 0; }

static void writeControl(uint8_t i, float d);

static void appSet(long id, double v) {
    if (id < 1 || id > controlCount_) return;
    AppControl& c = controls_[id - 1];
    if (c.kind >= APP_HOLD_BUTTON) c.heldAt = appMs_;
    writeControl((uint8_t)(id - 1), (float)v);
}

static void writeControl(uint8_t i, float d) {
    AppControl& c = controls_[i];
    long id = i + 1;
    if (d < c.min) d = c.min;  // the board clamps whatever the app sends (spec 11)
    if (d > c.max) d = c.max;
    float base = d * c.per;
    noInterrupts();  // a fast handler may read it
    switch (c.type) {
    case APP_BOOL:
        *(bool*)c.var = d != 0;
        break;
    case APP_INT:
        *(int32_t*)c.var = (int32_t)(base < 0 ? base - 0.5f : base + 0.5f);
        break;
    case APP_I64:
        *(int64_t*)c.var = (int64_t)(base < 0 ? base - 0.5f : base + 0.5f);
        break;
    default:
        *(float*)c.var = base;
    }
    interrupts();
    c.sent = controlValue(c);
    AppOut o("value");
    o.whole((int32_t)id).num(c.sent);
    o.end();
}

// questions
struct AppAsk {
    int16_t id;
    uint8_t state;  // 0 free, 1 waiting, 2 answered
    bool yes;
    Time until;
};
static AppAsk asks_[4];

void app_ask_begin(int16_t id, const Text& question, Time timeout) {
    AppAsk* slot = nullptr;
    for (AppAsk& a : asks_)
        if (a.state && a.id == id) slot = &a;
    for (AppAsk& a : asks_)
        if (!slot && !a.state) slot = &a;
    if (!slot) return;  // four questions open already: this one is none
    if (timeout <= 0) timeout = 300000000LL;  // 5 min (README 6)
    slot->id = id;
    slot->state = 1;
    slot->until = now() + timeout;
    AppOut o("ask");
    o.whole(id).field(question.s).whole((int32_t)(timeout / 1000000));
    o.end();
}

bool app_ask_done(int16_t id) {
    for (AppAsk& a : asks_)
        if (a.state && a.id == id) return a.state == 2 || now() >= a.until || halted_;
    return true;
}

Opt<bool> app_ask_result(int16_t id) {
    for (AppAsk& a : asks_) {
        if (a.state && a.id == id) {
            bool answered = a.state == 2;
            a.state = 0;
            return answered ? Opt<bool>(a.yes) : Opt<bool>();
        }
    }
    return Opt<bool>();
}

#ifdef MUJARD_HOST_TEST
/** test blocks: 'answer true' / 'answer false' answers the question the program waits on, 'answer none'
 *  leaves it to run out of time (spec 12). Nothing waiting: the answer is dropped. */
void test_answer(int yes) {
    for (AppAsk& a : asks_) {
        if (a.state == 1) {
            if (yes < 0)
                a.until = 0;  // out of time: app_ask_done ends the wait, the result is none
            else {
                a.state = 2;
                a.yes = yes != 0;
            }
            return;
        }
    }
}
#endif

static void appAnswer(long id, bool yes) {
    for (AppAsk& a : asks_) {
        if (a.state == 1 && a.id == id) {
            a.state = 2;
            a.yes = yes;
        }
    }
}

static void appInput(const char* s) {
    if (strncmp(s, "\x1fM1 ", 4) != 0) return;
    s += 4;
    if (strcmp(s, "who") == 0) {
        if (appOn_) sendHello();
        return;
    }
    if (halted_ || !appOn_) return;  // a halted program takes no values or answers
    char* end;
    if (strncmp(s, "set\t", 4) == 0) {
        long id = strtol(s + 4, &end, 10);
        if (*end != '\t' || end == s + 4) return;
        const char* v = end + 1;
        double d = strtod(v, &end);
        if (*end || end == v || d != d || d > 3.4e38 || d < -3.4e38) return;
        appSet(id, d);
    } else if (strncmp(s, "answer\t", 7) == 0) {
        long id = strtol(s + 7, &end, 10);
        if (*end != '\t' || end == s + 7) return;
        if (strcmp(end + 1, "yes") == 0) appAnswer(id, true);
        else if (strcmp(end + 1, "no") == 0) appAnswer(id, false);
    }
}

#ifdef MUJARD_APP_MIRROR
/** An app line that came over the second link (mujard.ble), handled exactly like one from the serial port. */
void app_input_line(const char* s) { appInput(s); }
#endif

static void appTick(Time t) {
    appMs_ = (uint32_t)(t / 1000);
    // The failsafe: a held control (hold button, joystick) the app stopped refreshing - the phone went away, the
    // link dropped, the screen locked - goes back to rest on the board itself, so a car or robot stops on its own.
    for (uint8_t i = 0; i < controlCount_; i++) {
        AppControl& c = controls_[i];
        if (c.kind >= APP_HOLD_BUTTON && c.heldAt != 0 && appMs_ - c.heldAt > APP_HOLD_TIMEOUT_MS) {
            c.heldAt = 0;
            if (controlValue(c) != controlRest(c)) writeControl(i, controlRest(c));
        }
    }
    // the program changed a controlled variable itself: move the control in the app (checked 5 times a second)
    if (t < nextValueCheck_) return;
    nextValueCheck_ = t + 200000;
    for (uint8_t i = 0; i < controlCount_; i++) {
        AppControl& c = controls_[i];
        float v = controlValue(c);
        if (v != c.sent) {
            c.sent = v;
            AppOut o("value");
            o.whole(i + 1).num(v);
            o.end();
        }
    }
}

// show / plot / notify

void app_show(const char* label, const Text& text, bool hasNumber, float number, const char* unit, uint8_t style, bool hasRange, float min, float max) {
    static const char* styles[] = {"number", "gauge", "light"};
    AppOut o("show");
    o.field(label).field(text.s);
    if (hasNumber) o.num(number);
    else o.field("");
    o.field(unit).field(styles[style < 3 ? style : 0]);
    if (hasRange) o.num(min).num(max);
    else o.field("").field("");
    o.end();
}

AppPlot::AppPlot() { appStart(); }
AppPlot& AppPlot::add(const char* name, float v) {
    if (count_ >= 8) return *this;
    names_[count_] = name;
    values_[count_++] = v;
    return *this;
}
void AppPlot::end() {
    AppOut o("plot");
    for (uint8_t i = 0; i < count_; i++) {
        char b[24];
        float v = values_[i];
        if (v != v || v > 3.4e38f || v < -3.4e38f) continue;
        floatText(v, 3, b);
        Text pair(names_[i]);
        pair.add("=").add(b);
        o.field(pair.s);
    }
    o.end();
}
#endif  // the app protocol continues with app_notify below

void (*notify_net_fn)(const char*, bool) = nullptr;  // mujard_http.h sets it, with or without the app protocol

#ifndef MUJARD_NO_APP
#ifdef MUJARD_HOST_TEST
void test_notify_seen(const Text&);   // the test build records what was asked, not the wire frame
// The runtime's own tests link this file without mujard_test.cpp: a no-op they can call; the real one wins if linked.
#if defined(__GNUC__)
__attribute__((weak)) void test_notify_seen(const Text&) {}
#endif
#endif

void app_notify(const Text& message, bool urgent, int line) {
#ifdef MUJARD_HOST_TEST
    test_notify_seen(message);
#endif
    // at most one per 10 s from each line of the program (spec Appendix A, 2); the app and server limit again
    static int16_t lines[8];
    static Time last[8];
    static uint8_t used = 0;
    Time t = now();
    int slot = -1;
    for (uint8_t i = 0; i < used; i++)
        if (lines[i] == line) slot = i;
    if (slot < 0) {
        slot = used < 8 ? used++ : (int)(line & 7);
        lines[slot] = (int16_t)line;
        last[slot] = -10000001;
    }
    if (t - last[slot] < 10000000) return;
    last[slot] = t;
    AppOut o("notify");
    o.field(message.s).field(urgent ? "1" : "0").whole(line);
    o.end();
    // The same message to the account when this build can reach the server over Wi-Fi (Phase 9).
    if (notify_net_fn) notify_net_fn(message.s, urgent);
}
#endif  // MUJARD_NO_APP

}  // namespace mujard
