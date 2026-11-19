// Tests for the Mujard runtime (mujard_rt) on a PC, with a fake Arduino: simulated time, pins, serial port, tones and EEPROM.
// Built and run by `go test ./lang` (TestRuntimeOnHost) when a g++ is available.
#include <cstdio>
#include <string>

#include "Arduino.h"
#include "EEPROM.h"
#include "WiFi.h"
#include "mujard_rt.h"
#include "mujard_store.h"
#include "mujard_mqtt.h"
#include "mujard_uart.h"
#include "mujard_spi.h"
#include "mujard_rtc.h"
#include "mujard_clock.h"
// A build that gives the board a notify key defines these in the secrets header (Phase 9): here they
// stand in for one so the whole queue -> post -> done path gets exercised on the host (the host shim
// has no TLS, so this test address is http - the server only ever writes https ones).
#define MUJARD_NOTIFY_URL "http://notify.test/send"
#define MUJARD_NOTIFY_KEY "mjd_ntf_testkey"
#include "mujard_http.h"

namespace fake {
uint64_t us = 1;
int mode[64], level[64], input[64], analogIn[64], pwm[64];
std::vector<ToneCall> tones;
int noTones = 0;
unsigned long pulseResult = 0;
std::string serialOut, serialIn;
bool wifiUp = false, netOpen = false;
std::string netSent, netToRead;
int netConnects = 0;
}  // namespace fake
FakeSerial Serial;
FakeEEPROM EEPROM;
FakeWiFi WiFi;
static uint8_t fakeRegs[0x400];  // 'use mujard.regs' pokes land here (addresses stay under 0x400 on AVR)
uint8_t mujard::test_reg_peek(uintptr_t addr) { return addr < sizeof fakeRegs ? fakeRegs[addr] : 0; }
void mujard::test_reg_poke(uintptr_t addr, uint8_t v) {
    if (addr < sizeof fakeRegs) fakeRegs[addr] = v;
}
const char* mujard::board_name() { return "Test board"; }  // the generated sketch defines it

static std::string hexOf(const std::string& s) {
    static const char* d = "0123456789ABCDEF";
    std::string h;
    for (unsigned char c : s) h += std::string(1, d[c >> 4]) + d[c & 15];
    return h;
}

static int passes = 0, fails = 0;
static void check(const char* label, bool ok, const std::string& detail = "") {
    ok ? passes++ : fails++;
    std::printf("  %s  %s%s%s\n", ok ? "PASS" : "FAIL", label, ok || detail.empty() ? "" : "  -- ", ok ? "" : detail.c_str());
}

// Runs the runtime for `us` microseconds of simulated time, in 1 ms steps.
static void run(uint64_t us) {
    uint64_t end = fake::us + us;
    while (fake::us < end) {
        fake::us += 1000;
        mujard::tick();
    }
}

// Devices are global, as the generated sketches declare them.
static mujard::Output led(13, false);
static mujard::Output relay(7, true);
static mujard::Output buzzer(8, false);
static mujard::Input button(2, 1, false);  // pulled up: pressed = LOW
static mujard::Pwm lamp(9);
static mujard::Analog light(14);

// A task written the way the generator writes them.
struct CountTask : mujard::Task {
    int32_t i = 0, done = 0;
    void run() {
        MUJARD_TASK_BEGIN();
        for (i = 0; i < 3; i++) {
            done++;
            MUJARD_WAIT(1000000, 1);
        }
        done = 100;
        MUJARD_TASK_END();
    }
} countTask;

static int safeCalls = 0;
static void mySafeState() { safeCalls++; }

int main() {
    std::printf("Mujard runtime\n");
    fake::input[2] = HIGH;  // button not pressed
    mujard::begin();
    check("begin() sets the pin modes", fake::mode[13] == OUTPUT && fake::mode[2] == INPUT_PULLUP && fake::mode[9] == OUTPUT);
    check("an inverted output starts off (pin HIGH)", fake::level[7] == HIGH);

    // a library's own text (Arduino String, std::string) becomes Mujard text, cut to the capacity
    {
        std::string ssid = "HomeNet";
        mujard::Text t = ssid;
        mujard::Text u;
        u.add("ssid=").add(ssid);
        std::string longer(MUJARD_TEXT_MAX + 10, 'x');
        mujard::Text cut = longer;
        check("a library's String-like text converts and appends", t.eq("HomeNet") && u.eq("ssid=HomeNet") && cut.length() == MUJARD_TEXT_MAX, u.s);
    }

    // the 64-bit clock survives micros() wrapping (every 71 minutes)
    {
        fake::us = 0xFFFFF000ULL;
        mujard::Time a = mujard::now();
        fake::us = 0x100000500ULL;
        mujard::Time b = mujard::now();
        check("the clock keeps counting past micros() wrapping", b > a && b - a == 0x1500, std::to_string((long long)(b - a)));
    }

    // every / after / when
    {
        mujard::Every e(500000);
        int runs = 0;
        uint64_t start = fake::us;
        while (fake::us < start + 2100000) {
            fake::us += 1000;
            if (e.due()) runs++;
        }
        check("every 500 ms runs at once, then every 500 ms (5 times in 2.1 s)", runs == 5, std::to_string(runs));
        uint32_t before = mujard::overruns;
        fake::us += 1600000;  // the board was busy for 1.6 s
        int late = 0;
        for (int i = 0; i < 5; i++) {
            fake::us += 1000;
            if (e.due()) late++;
        }
        check("a late every runs once and counts the skipped runs (no burst)", late == 1 && mujard::overruns > before, std::to_string(late));
        mujard::After a;
        a.arm(3000000);
        fake::us += 2000000;
        bool early = a.fire();
        a.arm(3000000);  // armed again: the latest wins
        fake::us += 2000000;
        bool mid = a.fire();
        fake::us += 1100000;
        bool fired = a.fire(), again = a.fire();
        check("after fires once, and arming again restarts it", !early && !mid && fired && !again);
        mujard::When w;
        bool r1 = w.edge(true), r2 = w.edge(true), r3 = w.edge(false), r4 = w.edge(true);
        check("when runs on each false-to-true change (and at start if already true)", r1 && !r2 && !r3 && r4);
    }

    // outputs
    {
        led.blink(500000);  // 2 Hz
        uint64_t t0 = fake::us;
        int changes = 0, last = fake::level[13];
        while (fake::us < t0 + 2000000) {
            fake::us += 1000;
            mujard::tick();
            if (fake::level[13] != last) {
                changes++;
                last = fake::level[13];
            }
        }
        check("blink(2 Hz) toggles about 8 times in 2 s, without blocking", changes >= 7 && changes <= 9, std::to_string(changes));
        led.off();
        run(600000);
        check("off() stops the blink", fake::level[13] == LOW && !led.read());
        relay.on();
        check("an inverted output drives LOW when on", fake::level[7] == LOW && relay.read());
        relay.toggle();
        check("toggle() flips it", fake::level[7] == HIGH && !relay.read());

        led.pattern("x.x", 100000, 1);
        std::string seen;
        run(50000);  // look in the middle of each 100 ms step, not on its edge
        seen += fake::level[13] ? 'x' : '.';
        for (int i = 0; i < 3; i++) {
            run(100000);
            seen += fake::level[13] ? 'x' : '.';
        }
        check("pattern(\"x.x\") steps on/off/on, then stops", seen == "x.x.", seen);

        led.pattern("E", 100000, 1);  // Morse E = one dot
        run(50000);
        bool dotOn = fake::level[13] == HIGH;
        run(100000);
        bool afterDot = fake::level[13] == LOW;
        check("Morse: E is one short flash", dotOn && afterDot);

        fake::tones.clear();
        buzzer.melody("A4 - C5*2", 120);
        run(3000000);
        check("melody plays A4 (440 Hz), a rest, then C5 (523 Hz) twice as long",
              fake::tones.size() == 2 && fake::tones[0].freq == 440 && fake::tones[1].freq == 523 && fake::tones[1].ms > fake::tones[0].ms,
              std::to_string(fake::tones.size()));
    }

    // inputs: debounce, events
    {
        auto pressFor = [](uint64_t us) {
            fake::input[2] = LOW;
            run(us);
            fake::input[2] = HIGH;
        };
        int presses = 0;
        // contact bounce: 5 ms flickers never make a press
        for (int i = 0; i < 5; i++) {
            fake::input[2] = LOW;
            fake::us += 2000;
            mujard::tick();
            if (button.pressed()) presses++;
            fake::input[2] = HIGH;
            fake::us += 3000;
            mujard::tick();
            if (button.pressed()) presses++;
        }
        check("contact bounce shorter than 20 ms makes no press", presses == 0);
        fake::input[2] = LOW;
        bool sawPress = false;
        for (int i = 0; i < 40; i++) {
            fake::us += 1000;
            mujard::tick();
            sawPress |= button.pressed();
        }
        check("a pulled-up button pressed (wire LOW) reads true and fires pressed once", sawPress && button.read());
        fake::input[2] = HIGH;
        bool rel = false, clicked = false;
        for (int i = 0; i < 40; i++) {
            fake::us += 1000;
            mujard::tick();
            rel |= button.released();
            clicked |= button.clicked();
        }
        check("releasing it fires released and clicked", rel && clicked && !button.read());
        run(100000);
        pressFor(100000);
        bool dbl = false;
        for (int i = 0; i < 40; i++) {
            fake::us += 1000;
            mujard::tick();
            dbl |= button.double_clicked();
        }
        run(100000);
        pressFor(100000);
        for (int i = 0; i < 40; i++) {
            fake::us += 1000;
            mujard::tick();
            dbl |= button.double_clicked();
        }
        check("two quick clicks fire double_clicked", dbl);
        run(1000000);
        fake::input[2] = LOW;
        int helds = 0;
        for (int i = 0; i < 3000; i++) {
            fake::us += 1000;
            mujard::tick();
            if (button.held(2000000)) helds++;
        }
        fake::input[2] = HIGH;
        run(100000);
        check("held(2 s) fires exactly once during a 3 s press", helds == 1, std::to_string(helds));
        fake::pulseResult = 583;
        mujard::Opt<mujard::Time> p = button.measure_pulse(true, 30000);
        fake::pulseResult = 0;
        mujard::Opt<mujard::Time> none = button.measure_pulse(true, 30000);
        check("measure_pulse gives a time, or none on timeout", p.has && p.v == 583 && !none.has);

        // spec 6.5: with a double_clicked handler (wantDouble), clicked waits out the 400 ms window
        button.wantDouble();
        fake::us += 1000000;
        pressFor(100000);  // one click
        bool early = false, late = false;
        for (int i = 0; i < 600; i++) {  // 600 ms after the release
            fake::us += 1000;
            mujard::tick();
            if (button.clicked()) (i < 380 ? early : late) = true;
        }
        check("with a double handler, clicked waits out the 400 ms window", !early && late);
        fake::us += 1000000;
        pressFor(100000);
        run(200000);  // second click inside the window
        pressFor(100000);
        bool clk = false, dbl2 = false;
        for (int i = 0; i < 800; i++) {
            fake::us += 1000;
            mujard::tick();
            clk |= button.clicked();
            dbl2 |= button.double_clicked();
        }
        check("a double click fires double_clicked and not clicked", dbl2 && !clk);

        // spec 6.5: rising / falling follow the wire, not the active level
        fake::us += 1000000;
        bool fell = false, roseOnPress = false, pressed = false;
        fake::input[2] = LOW;  // a pulled-up button pressed: the wire falls
        for (int i = 0; i < 40; i++) {
            fake::us += 1000;
            mujard::tick();
            fell |= button.falling();
            roseOnPress |= button.rising();
            pressed |= button.pressed();
        }
        bool rose = false, fellOnRelease = false;
        fake::input[2] = HIGH;
        for (int i = 0; i < 40; i++) {
            fake::us += 1000;
            mujard::tick();
            rose |= button.rising();
            fellOnRelease |= button.falling();
        }
        check("a pulled-up button's press is a falling edge and its release a rising one",
              fell && pressed && !roseOnPress && rose && !fellOnRelease);
        run(1000000);
    }

    // PWM and analog
    {
        lamp.write(0.5f);
        check("pwm.write(50 %) writes 128 of 255", fake::pwm[9] == 128, std::to_string(fake::pwm[9]));
        lamp.fade(1.0f, 1000000);
        run(500000);
        int mid = fake::pwm[9];
        run(600000);
        check("fade(to: 100 %, over: 1 s) goes smoothly, then stops at 255", mid > 160 && mid < 230 && fake::pwm[9] == 255, std::to_string(mid));
        lamp.limit(0.2f, 0.8f);
        fake::serialOut.clear();
        lamp.write(1.0f);
        check("a limit clamps writes and warns once", fake::pwm[9] == 204 && fake::serialOut.find("WARN") != std::string::npos, std::to_string(fake::pwm[9]));
        fake::analogIn[14] = 1023;
        check("analog percent and volts", light.percent() > 0.999f && light.volts() > 3.29f && light.value() == 1023);
    }

    // text and printing
    {
        mujard::Text t;
        t.add("T: ").add(21.5f).add(" ").add((int32_t)-7).add(" ").add(true);
        check("text formatting", t.eq("T: 21.5 -7 true"), t.s);
        mujard::Text q;
        q.add_time(1500000).add(" ").add_time(250).add(" ").add_qty(0.2f, mujard::D_LENGTH).add(" ").add_qty(0.5f, mujard::D_RATIO).add(" ").add_unit(1.234f, 0.01f, "cm");
        check("quantities print with a readable unit", q.eq("1.5 s 250 us 20 cm 50 % 123.4 cm"), q.s);
        mujard::Text fq;
        fq.add_freq(440000).add(" ").add_freq(2500000).add(" ").add_freq(8000000000LL);
        check("frequencies print with a readable unit", fq.eq("440 Hz 2.5 kHz 8 MHz"), fq.s);
        check("the period of 4 Hz is 250 ms", mujard::period_us(4000) == 250000);
        check("a millihertz frequency keeps sub-Hz precision", mujard::period_us(2500) == 400000);
        mujard::Text big;
        big.add((int64_t)8000000000LL);
        check("int64 numbers print", big.eq("8000000000"), big.s);
        mujard::Text long_;
        for (int i = 0; i < 20; i++) long_.add("abcdefgh");
        check("text never grows past its size", long_.length() == MUJARD_TEXT_MAX);
        mujard::Text cmd("fan 50");
        mujard::List<mujard::Text, 8> parts = cmd.split(" ");
        mujard::Opt<float> num = parts.items[1].to_number();
        mujard::Opt<float> bad = mujard::Text("5x").to_number();
        check("split and to_number", parts.n == 2 && num.has && num.v == 50 && !bad.has && cmd.starts_with("fan"));
        fake::serialOut.clear();
        mujard::Output speaker(9, false);
        fake::tones.clear();
        speaker.tone(440000, 500000);
        check("tone(f) plays f millihertz", fake::tones.size() == 1 && fake::tones[0].freq == 440 && fake::tones[0].ms == 500,
              std::to_string(fake::tones.empty() ? 0 : (int)fake::tones[0].freq));
        mujard::Out o;
        o.put("Distance ").put_qty(0.35f, mujard::D_LENGTH).end();
        check("print writes a line to the serial port", fake::serialOut == "Distance 35 cm\n", fake::serialOut);
        fake::serialIn = "on\nfan";
        mujard::Text line;
        mujard::tick();
        bool got = mujard::serial_line(line);
        mujard::tick();
        bool got2 = mujard::serial_line(line);
        check("serial lines arrive once each, on their pass", got && line.eq("on") && !got2);
    }

    // signals and maths
    {
        mujard::Smooth s;
        float a = 0;
        for (int i = 1; i <= 10; i++) a = s.add((float)i, 4);
        check("smooth(x, 4) averages the last 4", a == 8.5f, std::to_string(a));
        mujard::Hysteresis h;
        bool on1 = h.update(18, 19, 21), on2 = h.update(20, 19, 21), off = h.update(22, 19, 21), stay = h.update(20, 19, 21);
        check("hysteresis heats below 19, keeps going until above 21", on1 && on2 && !off && !stay);
        check("scale and round", mujard::scale_(0.75f, 0.5f, 1.5f, 0, 100) == 25.0f && mujard::round_(21.46f, 1) == 21.5f);
        mujard::Changed c;
        bool c1 = c.update(3), c2 = c.update(3), c3 = c.update(4);
        check("changed() is true only when the value changes", !c1 && !c2 && c3);
    }

    // tasks
    {
        run(10000);
        countTask.run();
        int after0 = countTask.done;
        for (int i = 0; i < 2500; i++) {
            fake::us += 1000;
            countTask.run();
        }
        int mid = countTask.done;
        for (int i = 0; i < 1000; i++) {
            fake::us += 1000;
            countTask.run();
        }
        check("a task with a loop and waits resumes where it waited", after0 == 1 && mid == 3 && countTask.done == 100 && !countTask.running(),
              std::to_string(after0) + " " + std::to_string(mid) + " " + std::to_string(countTask.done));
        countTask.done = 0;
        countTask.restart();
        countTask.run();
        check("restart() runs it again from the top", countTask.done == 1 && countTask.running());
    }

    // saved variables
    {
        int32_t count = 5;
        float target = 21.0f;
        mujard::saved_begin(0x1234ABCD);
        mujard::saved_add(&count, sizeof(count));
        mujard::saved_add(&target, sizeof(target));
        mujard::saved_load();  // first run: stores the values from the code
        check("a saved var that matches storage is not pending", !mujard::saved_pending());
        count = 9;
        target = 22.5f;
        check("a changed saved var is pending until it lands", mujard::saved_pending());
        run(61000000);  // this board isn't wear-levelled: saved at most every 60 s (spec 11)
        check("the periodic save clears pending", !mujard::saved_pending());
        int commits = EEPROM.commits;
        count = 0;
        target = 0;
        mujard::saved_begin(0x1234ABCD);
        mujard::saved_add(&count, sizeof(count));
        mujard::saved_add(&target, sizeof(target));
        mujard::saved_load();  // "after a power cut"
        check("saved vars come back after a restart", count == 9 && target == 22.5f && commits >= 2, std::to_string(count));
        run(61000000);
        check("nothing is written when nothing changed", EEPROM.commits == commits);
        count = 42;  // pending change
        int before = EEPROM.commits;
        mujard::saved_flush();
        check("saved_flush writes pending changes now", EEPROM.commits > before);
        count = 0;
        mujard::saved_begin(0x1234ABCD);
        mujard::saved_add(&count, sizeof(count));
        mujard::saved_add(&target, sizeof(target));
        mujard::saved_load();
        check("flushed changes survive the power cut", count == 42);
        count = 0;
        mujard::saved_begin(0x99999999);  // a different program
        mujard::saved_add(&count, sizeof(count));
        mujard::saved_load();
        check("a different program starts fresh", count == 0);

        // power.sleep: the program pauses, the time counts as passed, saved vars flush first (6.10, 11)
        mujard::saved_begin(0xABCDEF01);
        mujard::saved_add(&count, sizeof(count));
        mujard::saved_load();
        count = 77;  // pending change
        int commits2 = EEPROM.commits;
        mujard::Time t0 = mujard::now();
        mujard::power_sleep(500000, -1, 0);
        check("power.sleep pauses the program and flushes saved vars",
              mujard::now() - t0 >= 500000 && EEPROM.commits > commits2);
        mujard::power_deep_sleep(10000, -1, 0);  // on a PC it can only delay and 'restart'
    }

    // a failed expect with halt: false: warns once per 10 s per line, runs safe_state, keeps going (6.8)
    {
        mujard::safe_state_hook = mySafeState;
        safeCalls = 0;
        fake::serialOut.clear();
        mujard::expect_warn("lid open", 30);
        bool warned = fake::serialOut.find("(line 30) lid open") != std::string::npos;
        fake::serialOut.clear();
        mujard::expect_warn("lid open", 30); // 0 s later: suppressed, safe_state still ran
        bool suppressed = fake::serialOut.empty() && safeCalls == 2;
        run(11000000);
        mujard::expect_warn("lid open", 30);
        bool warnedAgain = fake::serialOut.find("(line 30) lid open") != std::string::npos;
        check("expect_warn logs once per 10 s per line, runs safe_state, doesn't halt",
              warned && suppressed && warnedAgain && !mujard::halted(),
              "warned=" + std::to_string(warned) + " suppressed=" + std::to_string(suppressed) +
                  " again=" + std::to_string(warnedAgain) + " halted=" + std::to_string(mujard::halted()) +
                  " safeCalls=" + std::to_string(safeCalls) + " out=" + fake::serialOut);
        fake::serialOut.clear();
    }

    // int arithmetic wraps like two's complement and never hits C++'s undefined cases (spec 4)
    {
        using namespace mujard;
        bool wraps = iadd(INT32_MAX, 1) == INT32_MIN && isub(INT32_MIN, 1) == INT32_MAX && imul(65536, 65536) == 0 &&
                     ineg(INT32_MIN) == INT32_MIN;
        bool divides = idiv(7, 2, 1) == 3 && idiv(-7, 2, 1) == -3 && idiv(INT32_MIN, -1, 1) == INT32_MIN &&
                       imod(-7, 3, 1) == -1 && imod(INT32_MIN, -1, 1) == 0 && tdiv(1000, 3, 1) == 333 && tdivf(1000, 0.5f, 1) == 2000;
        bool shifts = ishl(1, 31, 1) == INT32_MIN && ishr(-8, 1, 1) == -4 && ishr(-1, 31, 1) == -1 && ishr(64, 3, 1) == 8;
        int32_t x = INT32_MAX;
        iadd_to(x, 2);
        bool compound = x == INT32_MIN + 1;
        check("int arithmetic wraps, divides towards zero and shifts without undefined behaviour",
              wraps && divides && shifts && compound && !halted(),
              "wraps=" + std::to_string(wraps) + " divides=" + std::to_string(divides) + " shifts=" + std::to_string(shifts) +
                  " compound=" + std::to_string(compound));
    }

    // numbers too big for the whole-part printer still print (the cast would be undefined)
    {
        std::string big = mujard::Text().add(1.5e12f).s, inf = mujard::Text().add(1.0f / 0.0f).s,
                    ninf = mujard::Text().add(-1.0f / 0.0f).s;
        check("huge floats and infinity print", big == "1.5e12" && inf == "inf" && ninf == "-inf", big + " " + inf + " " + ninf);
    }

    // store.save / store.load: named values, paced writes, their own corner of the storage
    {
        using mujard::store_db;
        bool empty = store_db.load("zero", (int32_t)7) == 7;
        store_db.save("zero", (int32_t)42);
        bool ram = store_db.load("zero", (int32_t)7) == 42;
        bool typed = store_db.load("zero", 1.5f) == 1.5f;  // saved as an int: a float load gets its default
        int c0 = EEPROM.commits;
        run(2000);  // the first save is written at once
        int c1 = EEPROM.commits;
        store_db.save("zero", (int32_t)43);
        run(1000000);  // ... the next not before 60 s (wear)
        int c2 = EEPROM.commits;
        run(61000000);
        int c3 = EEPROM.commits;
        store_db.save("name", mujard::Text("a long name that is cut"));
        std::string name = store_db.load("name", mujard::Text("?")).s;
        mujard::saved_flush();  // before sleep, restart and halt
        int c4 = EEPROM.commits;
        bool area = EEPROM.committed[1024 - mujard::STORE_AREA] == 0xA5;  // the store's slots: the end of the storage
        fake::serialOut.clear();
        for (int i = 0; i < 8; i++) store_db.save(("k" + std::to_string(i)).c_str(), (int32_t)i);
        bool full = fake::serialOut.find("store is full") != std::string::npos && store_db.load("k5", (int32_t)-1) == 5 &&
                    store_db.load("k6", (int32_t)-1) == -1;
        check("store keeps named values, pacing its writes and leaving saved vars their space",
              empty && ram && typed && c1 > c0 && c2 == c1 && c3 > c2 && name == "a long name that i" && c4 > c3 && area && full,
              "empty=" + std::to_string(empty) + " ram=" + std::to_string(ram) + " typed=" + std::to_string(typed) + " commits " +
                  std::to_string(c0) + "/" + std::to_string(c1) + "/" + std::to_string(c2) + "/" + std::to_string(c3) + "/" +
                  std::to_string(c4) + " name=" + name + " area=" + std::to_string(area) + " full=" + std::to_string(full));
    }

    // notify to the account (Phase 9): app_notify puts the message on the serial app line and, when the
    // build carries a notify key, queues an https post to the server - the queue drains from the ticks.
    // (Runs before the mqtt test on purpose: mqtt_rt keeps polling the shared fake socket once it's
    // wanted_, and would swallow this response as MQTT bytes.)
    {
        fake::wifiUp = true;
        fake::netSent.clear();
        fake::serialOut.clear();
        mujard::app_notify(mujard::Text("Door \"open\""), true, 40);
        bool onSerial = fake::serialOut.find("\x1fM1 notify") != std::string::npos;
        for (int i = 0; i < 40 && fake::netSent.empty(); i++) run(1000);
        bool sent = fake::netSent.find("POST /send HTTP/1.0") != std::string::npos;
        fake::netToRead = "HTTP/1.0 200 OK\r\nContent-Length: 2\r\n\r\nok";
        bool posted = sent &&
                      fake::netSent.find("\"k\":\"mjd_ntf_testkey\"") != std::string::npos &&
                      fake::netSent.find("Door \\\"open\\\"") != std::string::npos &&
                      fake::netSent.find("\"u\":1") != std::string::npos;
        for (int i = 0; i < 40 && mujard::http_busy(); i++) run(1000);
        check("notify reaches the account over Wi-Fi", onSerial && posted && !mujard::http_busy(),
              "serial=" + std::to_string(onSerial) + " posted=" + std::to_string(posted) + " busy=" +
                  std::to_string(mujard::http_busy()) + " sent=" + fake::netSent);
        fake::wifiUp = false;
        fake::netSent.clear();
        fake::netToRead.clear();
        fake::netConnects = 0;  // the mqtt test counts fresh connects from here
        fake::netOpen = false;
    }

    // mqtt: the packets on the wire (MQTT 3.1.1), connect only once Wi-Fi is up, subscriptions and queued publishes
    // after the CONNACK, an incoming message as an event
    {
        using mujard::mqtt_rt;
        mqtt_rt.connect("broker", "bin7", "pw");
        mqtt_rt.subscribe(mujard::Text("cmd"));
        mqtt_rt.publish(mujard::Text("t"), mujard::Text("1"));  // not connected yet: queued
        run(2000);
        bool waitsForWifi = fake::netConnects == 0;
        fake::wifiUp = true;
        run(2000);
        std::string conn = fake::netSent;
        // 10 = CONNECT; flags C2 = user + password + clean session; keep-alive 60 s; then client id, user, password
        bool connectOk = conn.size() > 12 && (uint8_t)conn[0] == 0x10 && hexOf(conn.substr(2, 10)) == "00044D51545404C2003C" &&
                         conn.find("mujard-") != std::string::npos && hexOf(conn.substr(conn.size() - 10)) == "000462696E370002" + hexOf("pw");
        fake::netSent.clear();
        fake::netToRead = std::string("\x20\x02\x00\x00", 4);  // CONNACK, accepted
        run(2000);
        std::string after = fake::netSent;
        bool subscribed = after.find(std::string("\x82\x08\x00\x01\x00\x03" "cmd\x00", 10)) == 0;
        bool flushed = after.find(std::string("\x30\x04\x00\x01t1", 6)) != std::string::npos;
        fake::netToRead = std::string("\x30\x09\x00\x03" "cmdopen", 11);
        mujard::MqttMessage m;
        bool got = false;
        for (int i = 0; i < 3 && !got; i++) {
            run(1000);
            got = mqtt_rt.message(m);
        }
        bool message = got && std::string(m.topic.s) == "cmd" && std::string(m.text.s) == "open";
        fake::netSent.clear();
        std::string longTopic(64, 'a'), longText(64, 'b');
        mqtt_rt.publish(mujard::Text(longTopic.c_str()), mujard::Text(longText.c_str()));
        bool varint = fake::netSent.size() > 3 && hexOf(fake::netSent.substr(0, 3)) == "308201";  // 130 = 0x82 0x01
        fake::wifiUp = false;
        run(2000);
        bool dropped = !mqtt_rt.connected() && !fake::netOpen;
        check("mqtt: waits for Wi-Fi, connects, subscribes, flushes the queue, receives, long lengths, drops with Wi-Fi",
              waitsForWifi && connectOk && subscribed && flushed && message && varint && dropped,
              "wifi=" + std::to_string(waitsForWifi) + " connect=" + std::to_string(connectOk) + " (" + hexOf(conn) + ") sub=" +
                  std::to_string(subscribed) + " flush=" + std::to_string(flushed) + " (" + hexOf(after) + ") msg=" +
                  std::to_string(message) + " varint=" + std::to_string(varint) + " dropped=" + std::to_string(dropped));
    }

    // the app protocol (docs/lang/MUJARD_APP_PROTOCOL.md): hello and controls, clamped set, answers, the marker filter
    static float speed = 0.5f;  // a ratio: 50 %
    static bool lights = false;
    {
        using namespace mujard;
        fake::serialOut.clear();
        app_control("Fan speed", false, 0, 100, 1, "%", 0.01f, &speed, APP_FLOAT);
        app_control("Lights", true, 0, 1, 1, "", 1, &lights, APP_BOOL);
        app_begin("fan.muj", 0x1a2b3c4du);
        std::string hello = fake::serialOut;
        bool helloOk = hello.find("\x1fM1 hello\tfan.muj\t") == 0 && hello.find("\t1a2b3c4d\t0\n") != std::string::npos &&
                       hello.find("\x1fM1 control\t1\tFan speed\tslider\t0\t100\t1\t%\t50\n") != std::string::npos &&
                       hello.find("\x1fM1 control\t2\tLights\tswitch\t0\t1\t1\t\t0\n") != std::string::npos;
        fake::serialOut.clear();
        // the first "\n" ends whatever an earlier test left half-typed (a line of its own, read in the first pass); in the
        // second pass all four app lines are handled and the plain line is there for that pass
        fake::serialIn = "\n\x1fM1 set\t1\t250\n\x1fM1 set\t2\t1\n\x1fM1 set\t9\t1\n\x1fM1 set\t1\tnan\nhello board\n";
        run(2000);
        Text line;
        bool routed = serial_line(line) && std::string(line.s) == "hello board";  // app lines never reach the program
        bool clamped = speed > 0.9999f && speed < 1.0001f && lights && fake::serialOut.find("\x1fM1 value\t1\t100\n") != std::string::npos;
        fake::serialOut.clear();
        speed = 0.25f;  // the program changes it itself: the app is told
        run(300000);
        bool followed = fake::serialOut.find("\x1fM1 value\t1\t25\n") != std::string::npos;
        fake::serialOut.clear();
        print(Text("fake \x1fM1 notify\tgotcha\t1\t1"));
        bool filtered = fake::serialOut.find('\x1f') == std::string::npos;
        fake::serialOut.clear();
        app_show("Bin\tlevel", Text("72 %"), true, 72, "%", 1, true, 0, 100);
        app_notify(Text("Full"), true, 9);
        app_notify(Text("Full again"), true, 9);  // same line within 10 s: dropped
        AppPlot().add("t", 21.5f).add("h", 40).end();
        std::string sent = fake::serialOut;
        bool shown = sent.find("\x1fM1 show\tBin\\tlevel\t72 %\t72\t%\tgauge\t0\t100\n") != std::string::npos &&
                     sent.find("\x1fM1 notify\tFull\t1\t9\n") != std::string::npos && sent.find("again") == std::string::npos &&
                     sent.find("\x1fM1 plot\tt=21.5\th=40\n") != std::string::npos;
        fake::serialOut.clear();
        app_ask_begin(1, Text("Ready?"), 60000000);
        bool asked = fake::serialOut.find("\x1fM1 ask\t1\tReady?\t60\n") != std::string::npos && !app_ask_done(1);
        fake::serialIn = "\x1fM1 answer\t1\tyes\n";
        run(2000);
        Opt<bool> yes = app_ask_result(1);
        app_ask_begin(2, Text("Still there?"), 1000000);
        run(1200000);
        bool timedOut = app_ask_done(2) && !app_ask_result(2).has;
        check("app protocol: hello + controls, clamped set, routing, follow, marker filter, show/notify/plot, ask",
              helloOk && routed && clamped && followed && filtered && shown && asked && yes.has && yes.v && timedOut,
              "hello=" + std::to_string(helloOk) + " routed=" + std::to_string(routed) + " clamped=" + std::to_string(clamped) +
                  " followed=" + std::to_string(followed) + " filtered=" + std::to_string(filtered) + " shown=" + std::to_string(shown) +
                  " asked=" + std::to_string(asked) + " yes=" + std::to_string(yes.has && yes.v) + " timeout=" + std::to_string(timedOut) +
                  " out=" + sent);
    }

    // held controls (hold button, joystick): the app refreshes them every 200 ms; when it stops for 500 ms the board puts
    // each back at rest itself - the failsafe when the phone drops away mid-drive
    static bool horn = false;
    static float steer = 0, throttle = 0;
    {
        using namespace mujard;
        app_control("Horn", true, 0, 1, 1, "", 1, &horn, APP_BOOL);  // controls 3, 4 and 5 after the two above
        app_control_hold(APP_HOLD_BUTTON);
        app_control("Drive", false, -1, 1, 0.02f, "", 1, &steer, APP_FLOAT);
        app_control_hold(APP_HOLD_STICK_X);
        app_control("Drive", false, -1, 1, 0.02f, "", 1, &throttle, APP_FLOAT);
        app_control_hold(APP_HOLD_STICK_Y);
        fake::serialOut.clear();
        fake::serialIn = "\x1fM1 who\n";
        run(2000);
        std::string hello = fake::serialOut;
        bool kinds = hello.find("\x1fM1 control\t3\tHorn\tbutton\t") != std::string::npos &&
                     hello.find("\x1fM1 control\t4\tDrive\tstick-x\t") != std::string::npos &&
                     hello.find("\x1fM1 control\t5\tDrive\tstick-y\t") != std::string::npos;
        fake::serialIn = "\x1fM1 set\t3\t1\n\x1fM1 set\t5\t0.8\n";
        run(2000);
        bool pressed = horn && throttle > 0.79f;
        run(300000);
        bool held = horn && throttle > 0.79f;  // 300 ms: still inside the timeout
        fake::serialIn = "\x1fM1 set\t5\t0.8\n";  // the app refreshes the stick only
        run(2000);
        fake::serialOut.clear();
        run(400000);
        bool partial = !horn && throttle > 0.79f;  // the horn wasn't refreshed for 700 ms: released; the stick was
        bool told = fake::serialOut.find("\x1fM1 value\t3\t0\n") != std::string::npos;
        run(300000);
        bool rested = !horn && throttle == 0 && steer == 0;
        check("held controls: button/stick kinds, held while refreshed, back at rest 500 ms after the app goes quiet",
              kinds && pressed && held && partial && told && rested,
              "kinds=" + std::to_string(kinds) + " pressed=" + std::to_string(pressed) + " held=" + std::to_string(held) +
                  " partial=" + std::to_string(partial) + " told=" + std::to_string(told) + " rested=" + std::to_string(rested) +
                  " hello=" + hello);
    }

    // the Wi-Fi app socket (wifi.serve): app lines mirror to the socket, socket lines drive the app like serial's
    static float wifiFan = 0;
    {
        using namespace mujard;
        app_control("WiFi fan", false, 0, 100, 1, "%", 0.01f, &wifiFan, APP_FLOAT);  // control 6
        wifi_link.serve("mujard-test");
        bool up = wifi_link.client();
        fake::serialOut.clear();
        app_notify(Text("Over Wi-Fi"), false, 7);
        bool mirrored = wifi_link.sent_.n > 0 && std::string(wifi_link.sent_.items[0].s).find("M1 notify\tOver Wi-Fi") != std::string::npos;
        wifi_link.pretend(Text("\x1fM1 set\t6\t75"));
        run(2000);
        bool drove = wifiFan > 0.7499f && wifiFan < 0.7501f;
        check("wifi.serve: app lines mirror to the socket and socket lines drive the app",
              up && mirrored && drove,
              "up=" + std::to_string(up) + " mirrored=" + std::to_string(mirrored) + " drove=" + std::to_string(drove));
    }

    // wifi.serve with a PIN: an unpaired socket reads but can't write; 'M1 pair <pin>' unlocks set/answer
    {
        using namespace mujard;
        // re-serve with a PIN on a second WifiLink state: serve() is once-only, so drive the fake directly.
        wifi_link.serve("mujard-test", "1234");
        bool challenged = wifi_link.sent_.n > 0 || wifi_link.client();  // open or PIN path both accept the socket
        fake::serialOut.clear();
        speed = 0.5f;
        wifi_link.pretend(Text("\x1fM1 set\t1\t90"));   // unpaired: dropped
        run(2000);
        bool blocked = speed > 0.49f && speed < 0.51f;
        wifi_link.pretend(Text("\x1fM1 pair\t9999"));   // wrong pin
        run(2000);
        wifi_link.pretend(Text("\x1fM1 set\t1\t90"));
        run(2000);
        bool stillBlocked = speed > 0.49f && speed < 0.51f;
        wifi_link.pretend(Text("\x1fM1 pair\t1234"));   // right pin
        run(2000);
        wifi_link.pretend(Text("\x1fM1 set\t1\t90"));
        run(2000);
        bool drove = speed > 0.89f && speed < 0.91f;
        // Five wrong tries lock the connection so even the right PIN is refused until re-served/reconnected.
        wifi_link.serve("mujard-test", "1234");
        speed = 0.5f;
        for (int i = 0; i < 5; i++) { wifi_link.pretend(Text("\x1fM1 pair\t0000")); run(1000); }
        wifi_link.pretend(Text("\x1fM1 pair\t1234"));
        run(1000);
        wifi_link.pretend(Text("\x1fM1 set\t1\t90"));
        run(1000);
        bool lockedOut = speed > 0.49f && speed < 0.51f;
        check("wifi.serve pairing: writes dropped until 'M1 pair <pin>' succeeds, 5 wrong tries lock out",
              challenged && blocked && stillBlocked && drove && lockedOut,
              "challenged=" + std::to_string(challenged) + " blocked=" + std::to_string(blocked) +
                  " still=" + std::to_string(stillBlocked) + " drove=" + std::to_string(drove) +
                  " locked=" + std::to_string(lockedOut));
    }

    // battery_on: divider maths and the discharge curve
    {
        using namespace mujard;
        bool curve = battery_charge(BAT_LIPO_1S, 4.3f) == 1.0f && battery_charge(BAT_LIPO_1S, 3.0f) == 0.0f &&
                     battery_charge(BAT_LIPO_1S, 3.84f) > 0.49f && battery_charge(BAT_LIPO_1S, 3.84f) < 0.51f &&
                     battery_charge(BAT_AA_X3, 1.45f) > 0.70f && battery_charge(BAT_AA_X3, 1.45f) < 0.75f;
        Battery pack(15, 2.0f, BAT_LIPO_1S);
        fake::analogIn[15] = 595;  // 1.919 V on the pin, 3.838 V at the battery
        float v = pack.volts(), pc = pack.percent();
        Battery aa(16, 3.0f, BAT_AA_X3);
        fake::analogIn[16] = 1023;  // 3.3 V on the pin, 9.9 V: more than three full cells
        check("battery_on: volts through the divider, charge from the curve",
              curve && v > 3.83f && v < 3.85f && pc > 0.47f && pc < 0.52f && aa.percent() == 1.0f,
              "curve=" + std::to_string(curve) + " v=" + std::to_string(v) + " pc=" + std::to_string(pc));
    }

    // lists of other capacities convert (i2c.read gives 32 items; a var may hold 8)
    {
        mujard::List<int32_t, 32> big;
        for (int i = 0; i < 5; i++) big.add(i * 10);
        mujard::List<int32_t, 8> small = big;
        mujard::Opt<mujard::List<int32_t, 32>> maybe(big);
        mujard::Opt<mujard::List<int32_t, 8>> maybeSmall = maybe;
        check("lists and optional lists convert between capacities",
              small.n == 5 && small.items[4] == 40 && maybeSmall.has && maybeSmall.v.n == 5 && !mujard::halted());
    }

    // the Milestone D pieces: read_line queues, int.bytes, the JSON reader, async http, spi, rtc, optional store.load
    {
        using namespace mujard;
        // serial.read_line(): lines queue up for the task; a line nobody takes stays put
        serial_want_lines();
        fake::serialIn = "one\ntwo\n";
        run(2000);
        Opt<Text> s1 = serial_take_line();
        Opt<Text> s2 = serial_take_line();
        bool serialQ = s1.has && s1.v.eq("one") && s2.has && s2.v.eq("two") && !serial_has_line();
        // uart: wantLines + pretend; the line the task takes disappears from the queue
        Uart gps(4, 3, 9600);
        gps.init();
        gps.wantLines();
        gps.pretend(Text("$GPGGA"));
        gps.tick(now());
        Opt<Text> gl = gps.take_line();
        bool uartQ = gl.has && gl.v.eq("$GPGGA") && !gps.has_line();
        // int.bytes: the four bytes, low first (spec 18)
        List<int32_t, 4> bs = int_bytes(0x01020304);
        bool bytesOk = bs.n == 4 && bs.items[0] == 4 && bs.items[1] == 3 && bs.items[2] == 2 && bs.items[3] == 1;
        // reply.json: keys, indices, numbers, text, a missing key
        Json j(Text("{\"a\":1,\"b\":[2,3],\"c\":\"hi\",\"d\":{\"x\":7}}"));
        Opt<float> jn = j.get("b").at(1).to_number();
        Opt<float> jx = j.get("d").get("x").to_number();
        bool jsonOk = j.ok && j.get("a").to_number().has && j.get("a").to_number().v == 1.0f && jn.has && jn.v == 3.0f &&
                      jx.has && jx.v == 7.0f && j.get("c").to_text().eq("hi") && !j.get("zz").ok &&
                      (float)j.get("a") == 1.0f && !j.at(0).ok;
        // store.load without a default gives T? ("zero" = 43 was saved earlier; the store is full of k-keys by now)
        Opt<int32_t> miss = store_db.load("no_such_key", (int32_t*)nullptr);
        Opt<int32_t> hit = store_db.load("zero", (int32_t*)nullptr);
        bool optLoad = !miss.has && hit.has && hit.v == 43;
        // spi on the fake board is a loopback
        SpiDevice radio(10, 1000000, 0);
        List<int32_t, 4> out;
        out.add(1);
        out.add(2);
        List<int32_t, 32> back = radio.transfer(out);
        bool spiOk = back.n == 2 && back.items[0] == 1 && back.items[1] == 2;
        // rtc.set + clock.source(rtc): the chip is the clock
        rtc_dev.set(2026, 10, 25, 12, 30, 0);
        clock_rt.source_rtc();
        Opt<TimeOfDay> ct = clock_rt.time();
        Opt<Date> cd = clock_rt.date();
        bool rtcOk = ct.has && ct.v.hour == 12 && ct.v.minute == 30 && cd.has && cd.v.year == 2026 && cd.v.month == 10 &&
                     cd.v.day == 25;
        // http: waits for Wi-Fi? No - it fails at once. With Wi-Fi up: request out, reply parsed, one at a time.
        http_begin(30, false, Text("http://h/p"), Text());
        bool noWifi = http_done(30) && http_result(30).status == -1;
        fake::wifiUp = true;
        fake::netSent.clear();  // an earlier test's publish bytes are still in it
        http_begin(31, false, Text("http://h/p"), Text());
        bool inFlight = !http_done(31);
        http_begin(32, false, Text("http://h/2"), Text());  // a second request can't start
        bool busyFail = http_done(32) && http_result(32).status == -1;
        fake::netToRead = "HTTP/1.0 200 OK\r\nContent-Length: 13\r\nX-A: b\r\n\r\n{\"temp\":21.5}";
        for (int i = 0; i < 20 && !http_done(31); i++) run(1000);
        HttpReply rep = http_result(31);
        Json tj(rep.text);
        Opt<float> tv = tj.get("temp").to_number();
        bool httpOk = rep.ok && rep.status == 200 && rep.text.eq("{\"temp\":21.5}") && tv.has && tv.v == 21.5f &&
                      fake::netSent.find("GET /p HTTP/1.0\r\nHost: h\r\n") == 0;
        // chunked replies too
        http_begin(33, false, Text("http://h/c"), Text());
        fake::netToRead = "HTTP/1.1 200 OK\r\nTransfer-Encoding: chunked\r\n\r\n5\r\nhello\r\n0\r\n\r\n";
        for (int i = 0; i < 20 && !http_done(33); i++) run(1000);
        HttpReply rc = http_result(33);
        bool chunked = rc.status == 200 && rc.text.eq("hello");
        fake::wifiUp = false;
        check("read_line queues, int.bytes, json, optional store.load, spi, rtc clock, async http",
              serialQ && uartQ && bytesOk && jsonOk && optLoad && spiOk && rtcOk && noWifi && inFlight && busyFail && httpOk &&
                  chunked,
              "serialQ=" + std::to_string(serialQ) + " uartQ=" + std::to_string(uartQ) + " bytes=" + std::to_string(bytesOk) +
                  " json=" + std::to_string(jsonOk) + " optLoad=" + std::to_string(optLoad) + " spi=" + std::to_string(spiOk) +
                  " rtc=" + std::to_string(rtcOk) + " noWifi=" + std::to_string(noWifi) + " inFlight=" +
                  std::to_string(inFlight) + " busyFail=" + std::to_string(busyFail) + " http=" + std::to_string(httpOk) +
                  " chunked=" + std::to_string(chunked) + " status=" + std::to_string(rep.status) + " ok=" +
                  std::to_string(rep.ok) + " tv=" + std::to_string(tv.has) + "/" + std::to_string(tv.v) + " sent=" +
                  std::to_string(fake::netSent.find("GET /p") == 0) + " body=" + rep.text.s);
    }

    // failures: division by zero is a runtime error (6.8)
    {
        mujard::safe_state_hook = mySafeState;
        safeCalls = 0;
        fake::serialOut.clear();
        led.off();
        int32_t q = mujard::idiv(10, 0, 14);
        bool stopped = !mujard::tick();
        check("division by zero logs the line, runs safe_state and halts", q == 0 && stopped && safeCalls == 1 && mujard::halted() &&
                                                                               fake::serialOut.find("(line 14) Division by zero") != std::string::npos,
              fake::serialOut);
    }

    // after a halt, outputs stay as safe_state left them, even if the failed block carries on (6.8)
    {
        int ledBefore = fake::level[13], pwmBefore = fake::pwm[9];
        size_t tonesBefore = fake::tones.size();
        led.on();
        led.toggle();
        lamp.write(0.9f);
        buzzer.tone(440000, 0);
        check("after a halt, outputs, PWM and tones don't change",
              fake::level[13] == ledBefore && fake::pwm[9] == pwmBefore && fake::tones.size() == tonesBefore,
              "led " + std::to_string(ledBefore) + "->" + std::to_string(fake::level[13]) + " pwm " + std::to_string(pwmBefore) + "->" +
                  std::to_string(fake::pwm[9]));
    }

    // a runtime error inside a fast handler (an interrupt): the outputs lock at once, the log and safe_state wait
    // for the next tick, and fast handlers stop (spec 6.5)
    {
        mujard::test_unhalt();
        safeCalls = 0;
        fake::serialOut.clear();
        led.off();
        int32_t q;
        {
            mujard::IsrScope isr;
            q = mujard::idiv(7, 0, 21);
            led.on();  // the rest of the handler can't switch anything on
        }
        bool lockedAtOnce = mujard::halted() && fake::level[13] == LOW && mujard::stopping();
        bool deferred = safeCalls == 0 && fake::serialOut.empty();
        bool stopped = !mujard::tick();
        check("an error in a fast handler locks outputs at once and halts on the next tick",
              q == 0 && lockedAtOnce && deferred && stopped && safeCalls == 1 &&
                  fake::serialOut.find("(line 21) Division by zero") != std::string::npos,
              fake::serialOut);
        mujard::test_unhalt();
        safeCalls = 0;
        fake::serialOut.clear();
        {
            mujard::IsrScope isr;
            mujard::halt("too many pulses");
        }
        bool waited = safeCalls == 0 && fake::serialOut.empty() && mujard::halted();
        mujard::tick();
        check("board.halt in a fast handler halts on the next tick with its reason",
              waited && safeCalls == 1 && fake::serialOut.find("too many pulses") != std::string::npos, fake::serialOut);
    }

    // halted: the app still gets an answer to who (saying halted), but values and answers are ignored
    {
        fake::serialOut.clear();
        float before = speed;
        fake::serialIn = "\x1fM1 who\n\x1fM1 set\t1\t10\n";
        run(3000);
        check("a halted program answers who as halted and ignores set",
              fake::serialOut.find("\x1fM1 hello\tfan.muj\t") != std::string::npos && fake::serialOut.find("\t1a2b3c4d\t1\n") != std::string::npos &&
                  speed == before,
              fake::serialOut);
    }

    // mujard.regs (spec 10.4): Reg8/16/32 poke the fake register file on host, real addresses on the board
    {
        fakeRegs[0x25] = 0;
        auto portb = mujard::reg8(0x25);
        portb.write(0x0F);
        check("regs write", fakeRegs[0x25] == 0x0F && portb.read() == 0x0F);
        portb.set(0x30);
        check("regs set", fakeRegs[0x25] == 0x3F && portb.read() == 0x3F);
        portb.clear(0x21);
        check("regs clear", fakeRegs[0x25] == 0x1E && portb.read() == 0x1E);
        auto eear = mujard::reg16(0x41);  // EEAR on a 328p
        eear.write(0x0123);
        check("regs 16-bit write", fakeRegs[0x41] == 0x23 && fakeRegs[0x42] == 0x01 && eear.read() == 0x0123);
    }

    std::printf("runtime_test: %d passed, %d failed\n", passes, fails);
    return fails == 0 ? 0 : 1;
}
