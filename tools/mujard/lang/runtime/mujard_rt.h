// mujard_rt - the runtime of Mujard programs (docs/lang/MUJARD_LANGUAGE_SPEC.md). The Mujard compiler copies this file and
// mujard_rt.cpp next to the sketch it generates. Plain C++11 on top of the Arduino API, for AVR, ESP32/ESP8266, RP2040
// and CMSIS ARM cores.
//
// Units are stored in fixed base units: time as int64 microseconds (mujard::Time); every other quantity as a float in
// hertz, metres, volts, amperes, degrees, degrees Celsius, metres per second, bytes, or a ratio (0.5 = 50 %).
#pragma once
#include <Arduino.h>
// std::initializer_list: AVR's libstdc++ doesn't ship the header. The type is a compiler builtin the
// implementation constructs with (begin, size), so on AVR the standard-required shape is declared by hand -
// without it nothing with a list literal compiles on AVR.
#if defined(ARDUINO_ARCH_AVR) && !defined(MUJARD_HOST_TEST)
namespace std {
template <class T>
class initializer_list {
public:
    typedef T value_type;
    typedef const T& reference;
    typedef const T& const_reference;
    typedef size_t size_type;
    typedef const T* iterator;
    typedef const T* const_iterator;

private:
    iterator b_;
    size_type n_;
    constexpr initializer_list(const_iterator a, size_type l) : b_(a), n_(l) {}

public:
    constexpr initializer_list() : b_(0), n_(0) {}
    constexpr size_type size() const { return n_; }
    constexpr const_iterator begin() const { return b_; }
    constexpr const_iterator end() const { return b_ + n_; }
};
}
#else
#include <initializer_list>
#endif
#include <time.h>

// mujard_config.h: written by the compiler next to this file with every program (empty when nothing can go), it switches
// off the runtime parts the program provably doesn't use (MUJARD_NO_APP, MUJARD_NO_SERIAL_IN, MUJARD_NO_BLINK,
// MUJARD_NO_PATTERN, MUJARD_NO_MELODY, MUJARD_NO_TONE), so they take no flash or RAM. Only ever "off". The parts it
// switches off are left undefined, not stubbed, so a program that did use one fails to link instead of quietly losing
// it. PC test builds always use the whole runtime. (Included plainly: the build server refuses sketches that probe for
// files, so the generator always writes it.)
#if !defined(MUJARD_HOST_TEST)
#include "mujard_config.h"
#endif
#if defined(MUJARD_NO_SERIAL_IN) && !defined(MUJARD_NO_APP)
#error "MUJARD_NO_SERIAL_IN needs MUJARD_NO_APP: the app's messages arrive as serial input"
#endif

// Direct pin access (plan Phase 13): on AVR an output or input finds its port register and bit once, at start-up, and
// then reads and writes the register itself - a fraction of digitalWrite/digitalRead's time, with the same interrupt
// safety. PC test builds keep digitalWrite/digitalRead, which the test shim records.
#if defined(ARDUINO_ARCH_AVR) && !defined(MUJARD_HOST_TEST) && defined(portOutputRegister) && defined(portInputRegister)
#define MUJARD_DIRECT_PINS 1
#endif

#ifndef MUJARD_TEXT_MAX
#define MUJARD_TEXT_MAX 64
#endif
static_assert(MUJARD_TEXT_MAX >= 1 && MUJARD_TEXT_MAX <= 255, "MUJARD_TEXT_MAX must be 1..255 (Text counts in a byte)");

namespace mujard {

typedef int64_t Time;  // microseconds
typedef int64_t Freq;  // millihertz (spec 5.1)

// ---- time

Time now();
inline float secs(Time t) { return (float)t / 1000000.0f; }
inline Time from_secs(float s) { return (Time)(s * 1000000.0f); }
/** The period of a frequency, in microseconds. */
inline Time period_us(Freq f) { return f > 0 ? 1000000000LL / f : 0; }

// ---- the program's life

/** Called first in setup(). */
void begin();
/** Called first in every loop() pass: the clock, inputs, devices, saved values, the watchdog. False once halted. */
bool tick();
/** The program's safe_state block (the generated sketch sets it; empty by default). */
extern void (*safe_state_hook)();
/** A runtime error or board.halt(): log it, run safe_state, stop everything (USB and uploads keep working). */
void fail(const char* message, int line);
/** A failed expect with halt: false: log it (at most once per 10 s per line), run safe_state, carry on. */
void expect_warn(const char* message, int line);
void halt(const char* reason);
/** True once the program has halted. From then on every output stays as safe_state left it: pins, PWM, tones,
 *  servos, shift registers and I2C writes are ignored, so code still running in the failed block can't switch
 *  anything back on (spec 6.8). */
bool halted();
extern uint32_t overruns;
/** Queued input lines dropped because the bounded queue was already full. */
extern uint32_t dropped_events;
/** A fast handler (spec 6.5) runs inside an IsrScope: a halt or runtime error in it locks the outputs at once and the
 *  rest (the log, safe_state) runs on the next tick(). stopping() is true from the moment a halt begins; fast
 *  handlers then return without doing anything. */
void isr_enter();
void isr_leave();
bool stopping();
struct IsrScope {
    IsrScope() { isr_enter(); }
    ~IsrScope() { isr_leave(); }
};
#ifdef MUJARD_HOST_TEST
/** The runtime's own tests only: run on after a halt. */
void test_unhalt();
/** The runtime's own tests only: register pokes land here instead of real addresses. */
uint8_t test_reg_peek(uintptr_t addr);
void test_reg_poke(uintptr_t addr, uint8_t v);
#endif

// ---- mujard.regs (spec 10.4): typed register access. reg8(addr) is a small value object; read/write/set/clear
// poke the chip register directly. Register writes are NOT covered by the halt output lock (same as cpp:).
// On MUJARD_HOST_TEST the pokes go to test_reg_poke/peek so tests see them.
template <class T>
struct Reg {
    uintptr_t addr;
    int32_t read() const {
#ifdef MUJARD_HOST_TEST
        int32_t v = 0;
        for (unsigned i = 0; i < sizeof(T); i++) v |= (int32_t)test_reg_peek(addr + i) << (8 * i);
        return v;
#else
        return *reinterpret_cast<volatile T*>(addr);
#endif
    }
    void write(int32_t v) const {
#ifdef MUJARD_HOST_TEST
        for (unsigned i = 0; i < sizeof(T); i++) test_reg_poke(addr + i, (uint8_t)(v >> (8 * i)));
#else
        *reinterpret_cast<volatile T*>(addr) = (T)v;
#endif
    }
    void set(int32_t bits) const { write(read() | bits); }
    void clear(int32_t bits) const { write(read() & ~bits); }
};
inline Reg<uint8_t> reg8(uintptr_t a) { return Reg<uint8_t>{a}; }
inline Reg<uint16_t> reg16(uintptr_t a) { return Reg<uint16_t>{a}; }
inline Reg<uint32_t> reg32(uintptr_t a) { return Reg<uint32_t>{a}; }

// ---- int arithmetic (spec 4, 6.8): wraps on overflow like two's complement (C++ leaves signed overflow undefined,
// so the generated code never uses the plain operators on ints); division or remainder by zero and a shift by less
// than 0 or more than 31 are runtime errors.

inline int32_t iadd(int32_t a, int32_t b) { return (int32_t)((uint32_t)a + (uint32_t)b); }
inline int32_t isub(int32_t a, int32_t b) { return (int32_t)((uint32_t)a - (uint32_t)b); }
inline int32_t imul(int32_t a, int32_t b) { return (int32_t)((uint32_t)a * (uint32_t)b); }
inline int32_t ineg(int32_t a) { return (int32_t)(0u - (uint32_t)a); }
int32_t idiv(int32_t a, int32_t b, int line);
int32_t imod(int32_t a, int32_t b, int line);
int32_t ishl(int32_t a, int32_t n, int line);
int32_t ishr(int32_t a, int32_t n, int line);
/** A time divided by a whole number / by a decimal number. */
Time tdiv(Time a, int32_t b, int line);
Time tdivf(Time a, float b, int line);
// x += y and friends; the volatile versions are for variables a fast handler also changes (spec 6.5)
inline void iadd_to(int32_t& t, int32_t v) { t = iadd(t, v); }
inline void isub_to(int32_t& t, int32_t v) { t = isub(t, v); }
inline void imul_to(int32_t& t, int32_t v) { t = imul(t, v); }
inline void idiv_to(int32_t& t, int32_t v, int line) { t = idiv(t, v, line); }
inline void iadd_to(volatile int32_t& t, int32_t v) { t = iadd(t, v); }
inline void isub_to(volatile int32_t& t, int32_t v) { t = isub(t, v); }
inline void imul_to(volatile int32_t& t, int32_t v) { t = imul(t, v); }
inline void idiv_to(volatile int32_t& t, int32_t v, int line) { t = idiv(t, v, line); }

/** A variable wider than a byte that a fast handler writes is read with interrupts off (spec 6.5). */
template <class T>
T atomic_read(const volatile T& v) {
    T t;
    noInterrupts();
    t = v;
    interrupts();
    return t;
}

// ---- optional values

template <class T>
struct Opt {
    bool has;
    T v;
    Opt() : has(false), v() {}
    Opt(const T& x) : has(true), v(x) {}
    /** An optional of a convertible type (a 32-byte list? into an 8-item list?). */
    template <class U>
    Opt(const Opt<U>& o) : has(o.has), v() {
        if (has) v = o.v;
    }
    const T& value(int line) const {
        if (!has) fail("A value was none", line);
        return v;
    }
};

// ---- dimensions (for printing quantities with a sensible unit)

enum Dim : uint8_t { D_FREQ = 1, D_LENGTH, D_VOLTAGE, D_CURRENT, D_ANGLE, D_TEMP, D_RATIO, D_SPEED, D_MEMORY, D_BAUD };

// ---- clock values (clock.time, clock.date; set over Wi-Fi in mujard_wifi.h)

struct TimeOfDay {
    int32_t hour, minute, second;
};
struct Date {
    int32_t year, month, day, weekday;  // weekday: 1 = Monday ... 7 = Sunday
};
/** An RTC chip (mujard_rtc.h) registers its reader here, so the clock (mujard_clock.h) can ask it for the
 *  time of day without the two headers knowing each other. False when the chip can't answer. */
extern bool (*rtc_clock_hook)(struct tm& out);

// ---- text

template <class T, int N>
struct List;
template <int N>
struct TextN;

/** Text of fixed capacity (spec 4.5): never grows past MUJARD_TEXT_MAX characters; longer results are cut. */
struct Text {
    char s[MUJARD_TEXT_MAX + 1];
    uint8_t n;
    Text() : n(0) { s[0] = 0; }
    Text(const char* c) : n(0) {
        s[0] = 0;
        add(c);
    }
    /** A library's own text (Arduino String, std::string - anything with c_str()), as a C++ library method
     *  described as returning text hands it back. Cut to the capacity like any other text. */
    template <class S, class = decltype(((const S*)0)->c_str())>
    Text(const S& v) : n(0) {
        s[0] = 0;
        add((const char*)v.c_str());
    }
    template <class S, class = decltype(((const S*)0)->c_str())>
    Text& add(const S& v) {
        return add((const char*)v.c_str());
    }
    /** A different fixed-capacity text (or a library string) assigned into the default 64-character one. */
    template <class S, class = decltype(((const S*)0)->c_str())>
    Text& operator=(const S& v) {
        n = 0;
        s[0] = 0;
        return add(v);
    }
#if !defined(MUJARD_HOST_TEST)
    /** A library's text kept in program memory (F("..."): MFRC522's PICC_GetTypeName and many others), read out of it.
     *  Cut to the capacity like any other text. */
    Text(const __FlashStringHelper* f) : n(0) {
        s[0] = 0;
        add(f);
    }
    Text& add(const __FlashStringHelper* f) {
        PGM_P p = reinterpret_cast<PGM_P>(f);
        for (char c; p && n < MUJARD_TEXT_MAX && (c = (char)pgm_read_byte(p)) != 0; p++) {
            s[n++] = c;
            s[n] = 0;
        }
        return *this;
    }
#endif
    Text& add(const char* c);
    Text& add(const Text& t) { return add(t.s); }
    Text& add(int32_t v);
    Text& add(int64_t v);
    Text& add(float v);
    Text& add(double v) { return add((float)v); }
    Text& add(bool v) { return add(v ? "true" : "false"); }
    Text& add(const TimeOfDay& t);  // 07:30:05
    Text& add(const Date& d);       // 2026-10-22
    Text& add_time(Time t);
    Text& add_freq(Freq f);
    Text& add_qty(float v, Dim d);
    Text& add_unit(float v, float per, const char* unit);  // v in base units shown in a fixed unit (x as cm)
    Text& add_none() { return add("none"); }
    template <class T>
    Text& add_opt(const Opt<T>& o) {
        if (!o.has) return add_none();
        return add(o.v);
    }
    Text& add_opt_time(const Opt<Time>& o) { return o.has ? add_time(o.v) : add_none(); }
    Text& add_opt_freq(const Opt<Freq>& o) { return o.has ? add_freq(o.v) : add_none(); }
    Text& add_opt_qty(const Opt<float>& o, Dim d) { return o.has ? add_qty(o.v, d) : add_none(); }
    template <class T, int N>
    Text& add_list(const List<T, N>& l);
    int32_t length() const { return n; }
    /** t[i]: the character as a one-character text; out of range is a runtime error, like list.at (spec 4). */
    Text at(int32_t i, int line) const;
    bool eq(const char* c) const { return strcmp(s, c) == 0; }
    bool eq(const Text& t) const { return strcmp(s, t.s) == 0; }
    template <int M>
    bool eq(const TextN<M>& t) const { return eq(t.s); }
    bool contains(const char* c) const { return strstr(s, c) != nullptr; }
    bool starts_with(const char* c) const { return strncmp(s, c, strlen(c)) == 0; }
    bool ends_with(const char* c) const {
        size_t k = strlen(c);
        return k <= n && strcmp(s + n - k, c) == 0;
    }
    Text upper() const;
    Text lower() const;
    Opt<float> to_number() const;
    List<Text, 8> split(const char* sep) const;
    operator const char*() const { return s; }
};

inline bool operator==(const Text& a, const Text& b) { return a.eq(b); }
inline bool operator!=(const Text& a, const Text& b) { return !a.eq(b); }

/** Builds a line for print / log and writes it to the serial port. */
struct Out {
    Out(const char* prefix = nullptr);
    template <class T>
    Out& put(const T& v) {
        t.add(v);
        return *this;
    }
    Out& put_time(Time v) {
        t.add_time(v);
        return *this;
    }
    Out& put_qty(float v, Dim d) {
        t.add_qty(v, d);
        return *this;
    }
    Out& put_unit(float v, float per, const char* unit) {
        t.add_unit(v, per, unit);
        return *this;
    }
    template <class T>
    Out& put_opt(const Opt<T>& o) {
        t.add_opt(o);
        return *this;
    }
    void end();
    Text t;
};

void serial_begin(uint32_t baud);
/** True on the loop pass when a complete line arrived on the serial port; the line goes to out. */
bool serial_line(Text& out);
/** serial.read_line() in a task: incoming lines are kept until one is taken (the generated setup() calls this). */
void serial_want_lines();
bool serial_has_line();
/** The oldest buffered line, or none. A task waits on serial_has_line() first, so the value is there. */
Opt<Text> serial_take_line();

// ---- lists

template <class T, int N>
struct List {
    T items[N];
    int32_t n;
    List() : items(), n(0) {}
    /** [1, 2, 3] written inside an expression: each item is add()ed; more than fit is a runtime error, like add() on a
     *  full list. */
    List(std::initializer_list<T> il) : items(), n(0) {
        for (const T& v : il) add(v);
    }
    /** A list of another capacity (i2c.read gives 32 bytes, a var may hold 8): copies the items; more than fit is a
     *  runtime error, like add() on a full list. */
    template <class U, int M>
    List(const List<U, M>& o) : items(), n(0) {
        copy(o);
    }
    template <class U, int M>
    List& operator=(const List<U, M>& o) {
        n = 0;
        copy(o);
        return *this;
    }
    template <class U, int M>
    void copy(const List<U, M>& o) {
        for (int32_t i = 0; i < o.n; i++) add(o.items[i]);
    }
    int32_t length() const { return n; }
    void add(const T& x) {
        if (n >= N) {
            fail("The list is full", 0);
            return;
        }
        items[n++] = x;
    }
    void remove(int32_t i) {
        if (i < 0 || i >= n) {
            fail("List index out of range", 0);
            return;
        }
        for (int32_t k = i; k + 1 < n; k++) items[k] = items[k + 1];
        n--;
    }
    void clear() { n = 0; }
    T& at(int32_t i, int line) {
        if (i < 0 || i >= n) {
            fail("List index out of range", line);
            return items[0];
        }
        return items[i];
    }
    /** x in list (spec 9). */
    bool contains(const T& x) const {
        for (int32_t i = 0; i < n; i++) {
            if (items[i] == x) return true;
        }
        return false;
    }
};

template <class T, int N>
Text& Text::add_list(const List<T, N>& l) {
    add("[");
    for (int32_t i = 0; i < l.n; i++) {
        if (i) add(", ");
        add(l.items[i]);
    }
    return add("]");
}

/** Text of another fixed capacity. Text is the 64-character API type; text[N] gives a variable this smaller or
 *  larger buffer. Building and comparing go through the same truncating rules (spec 4.5). */
template <int N>
struct TextN {
    static_assert(N >= 1 && N <= 255, "text size must be 1..255");
    char s[N + 1];
    uint8_t n;
    TextN() : n(0) { s[0] = 0; }
    TextN(const char* c) : n(0) {
        s[0] = 0;
        add(c);
    }
    TextN(const Text& t) : n(0) {
        s[0] = 0;
        add(t.s);
    }
    template <int M>
    TextN(const TextN<M>& t) : n(0) {
        s[0] = 0;
        add(t.s);
    }
    template <class S, class = decltype(((const S*)0)->c_str())>
    TextN(const S& v) : n(0) {
        s[0] = 0;
        add((const char*)v.c_str());
    }
    TextN& operator=(const TextN& t) = default;
    TextN& operator=(const char* c) {
        n = 0;
        s[0] = 0;
        return add(c);
    }
    TextN& operator=(const Text& t) { return *this = t.s; }
    template <int M>
    TextN& operator=(const TextN<M>& t) { return *this = t.s; }
    template <class S, class = decltype(((const S*)0)->c_str())>
    TextN& operator=(const S& v) {
        n = 0;
        s[0] = 0;
        return add((const char*)v.c_str());
    }
    TextN& add(const char* c) {
        while (c && *c && n < N) s[n++] = *c++;
        s[n] = 0;
        return *this;
    }
    TextN& add(const Text& t) { return add(t.s); }
    template <int M>
    TextN& add(const TextN<M>& t) { return add(t.s); }
    template <class T>
    TextN& add(const T& v) {
        Text t;
        t.add(v);
        return add(t.s);
    }
#if !defined(MUJARD_HOST_TEST)
    TextN(const __FlashStringHelper* f) : n(0) {
        s[0] = 0;
        Text t(f);
        add(t.s);
    }
    TextN& add(const __FlashStringHelper* f) {
        Text t;
        t.add(f);
        return add(t.s);
    }
#endif
    TextN& add_time(Time t) {
        Text v;
        v.add_time(t);
        return add(v.s);
    }
    TextN& add_freq(Freq f) {
        Text v;
        v.add_freq(f);
        return add(v.s);
    }
    TextN& add_qty(float v, Dim d) {
        Text t;
        t.add_qty(v, d);
        return add(t.s);
    }
    TextN& add_unit(float v, float per, const char* unit) {
        Text t;
        t.add_unit(v, per, unit);
        return add(t.s);
    }
    TextN& add_none() { return add("none"); }
    template <class T>
    TextN& add_opt(const Opt<T>& o) {
        if (!o.has) return add_none();
        return add(o.v);
    }
    TextN& add_opt_time(const Opt<Time>& o) { return o.has ? add_time(o.v) : add_none(); }
    TextN& add_opt_freq(const Opt<Freq>& o) { return o.has ? add_freq(o.v) : add_none(); }
    TextN& add_opt_qty(const Opt<float>& o, Dim d) { return o.has ? add_qty(o.v, d) : add_none(); }
    template <class T, int M>
    TextN& add_list(const List<T, M>& l) {
        add("[");
        for (int32_t i = 0; i < l.n; i++) {
            if (i) add(", ");
            add(l.items[i]);
        }
        return add("]");
    }
    int32_t length() const { return n; }
    TextN at(int32_t i, int line) const {
        if (i < 0 || i >= n) {
            fail("Text index out of range", line);
            return TextN();
        }
        char b[2] = {s[i], 0};
        return TextN(b);
    }
    const char* c_str() const { return s; }
    bool eq(const char* c) const { return strcmp(s, c) == 0; }
    bool eq(const Text& t) const { return strcmp(s, t.s) == 0; }
    template <int M>
    bool eq(const TextN<M>& t) const { return strcmp(s, t.s) == 0; }
    bool contains(const char* c) const { return strstr(s, c) != nullptr; }
    bool contains(const Text& t) const { return contains(t.s); }
    template <int M>
    bool contains(const TextN<M>& t) const { return contains(t.s); }
    bool starts_with(const char* c) const { return strncmp(s, c, strlen(c)) == 0; }
    bool starts_with(const Text& t) const { return starts_with(t.s); }
    template <int M>
    bool starts_with(const TextN<M>& t) const { return starts_with(t.s); }
    bool ends_with(const char* c) const {
        size_t k = strlen(c);
        return k <= n && strcmp(s + n - k, c) == 0;
    }
    bool ends_with(const Text& t) const { return ends_with(t.s); }
    template <int M>
    bool ends_with(const TextN<M>& t) const { return ends_with(t.s); }
    TextN upper() const {
        TextN t(*this);
        for (uint8_t i = 0; i < t.n; i++) t.s[i] = (char)toupper(t.s[i]);
        return t;
    }
    TextN lower() const {
        TextN t(*this);
        for (uint8_t i = 0; i < t.n; i++) t.s[i] = (char)tolower(t.s[i]);
        return t;
    }
    Opt<float> to_number() const {
        if (n == 0) return Opt<float>();
        char* end = nullptr;
        float v = (float)strtod(s, &end);
        while (end && *end == ' ') end++;
        if (end == s || (end && *end)) return Opt<float>();
        return Opt<float>(v);
    }
    List<TextN, 8> split(const char* sep) const {
        List<TextN, 8> out;
        size_t k = strlen(sep);
        const char* p = s;
        while (true) {
            const char* q = k ? strstr(p, sep) : nullptr;
            TextN part;
            size_t len = q ? (size_t)(q - p) : strlen(p);
            for (size_t i = 0; i < len && part.n < N; i++) part.s[part.n++] = p[i];
            part.s[part.n] = 0;
            if (out.n < 8) out.items[out.n++] = part;
            if (!q) break;
            p = q + k;
        }
        return out;
    }
    operator const char*() const { return s; }
    operator Text() const { return Text(s); }
};

template <int A, int B>
inline bool operator==(const TextN<A>& a, const TextN<B>& b) { return a.eq(b); }
template <int N>
inline bool operator==(const TextN<N>& a, const Text& b) { return a.eq(b); }
template <int N>
inline bool operator==(const Text& a, const TextN<N>& b) { return a.eq(b.s); }
template <int A, int B>
inline bool operator!=(const TextN<A>& a, const TextN<B>& b) { return !a.eq(b); }
template <int N>
inline bool operator!=(const TextN<N>& a, const Text& b) { return !a.eq(b); }
template <int N>
inline bool operator!=(const Text& a, const TextN<N>& b) { return !a.eq(b.s); }

/** int.bytes: the number's four bytes, the low byte first (spec 18). */
inline List<int32_t, 4> int_bytes(int32_t v) {
    List<int32_t, 4> b;
    for (int i = 0; i < 4; i++) b.add((int32_t)(((uint32_t)v >> (8 * i)) & 0xFF));
    return b;
}

/** print: one line on the serial port. */
void print(const Text& t);
inline void print(const char* s) { print(Text(s)); }
template <int N>
inline void print(const TextN<N>& t) { print(t.s); }
/** log.info / log.warn / log.error: a line with the time and the level. */
void log_line(const char* level, const Text& t);
inline void log_line(const char* level, const char* s) { log_line(level, Text(s)); }
template <int N>
inline void log_line(const char* level, const TextN<N>& t) { log_line(level, t.s); }

// ---- maths and smart signals

template <class T>
T abs_(T x) { return x < 0 ? -x : x; }
template <class T>
T min_(T a, T b) { return b < a ? b : a; }
template <class T>
T max_(T a, T b) { return a < b ? b : a; }
template <class T>
T clamp_(T x, T lo, T hi) { return x < lo ? lo : (hi < x ? hi : x); }
float scale_(float x, float a, float b, float c, float d);
float round_(float x, int32_t decimals);
inline float sin_deg(float d) { return sinf(d * 0.01745329252f); }
inline float cos_deg(float d) { return cosf(d * 0.01745329252f); }
inline float tan_deg(float d) { return tanf(d * 0.01745329252f); }
inline float atan2_deg(float y, float x) { return atan2f(y, x) * 57.2957795f; }
int32_t random_range(int32_t lo, int32_t hi_inclusive);
float random_range(float lo, float hi);

/** Average of the last count readings (count up to 32), one per call site. */
struct Smooth {
    float buf[32];
    uint8_t i, filled;
    Smooth() : buf(), i(0), filled(0) {}
    float add(float x, int32_t count);
};

struct Hysteresis {
    bool state;
    Hysteresis() : state(false) {}
    bool update(float x, float on, float off);
};

struct Changed {
    float last;
    bool first;
    Changed() : last(0), first(true) {}
    bool update(float x);
};

struct Edge {
    bool prev;
    Edge() : prev(false) {}
    bool rising(bool c) {
        bool e = c && !prev;
        prev = c;
        return e;
    }
    bool falling(bool c) {
        bool e = !c && prev;
        prev = c;
        return e;
    }
};

struct Peak {
    float best;
    Time at;
    bool any;
    Peak() : best(0), at(0), any(false) {}
    float update(float x, Time window);
};

// ---- blocks and tasks

/** every P: runs on the first pass, then every P on its schedule; missed runs are skipped (spec 6.2). */
struct Every {
    Time period, next;
    bool on;
    explicit Every(Time p) : period(p), next(0), on(true) {}
    bool due();
    void start() {
        on = true;
        next = 0;
    }
    void stop() { on = false; }
    void restart() { start(); }
    bool running() const { return on; }
};

/** after D: one shot; arming again restarts it (spec 6.3). */
struct After {
    Time at;
    bool armed;
    After() : at(0), armed(false) {}
    void arm(Time d) {
        at = now() + d;
        armed = true;
    }
    bool fire() {
        if (armed && now() >= at) {
            armed = false;
            return true;
        }
        return false;
    }
    void start() {}
    void stop() { armed = false; }
    void restart() {}
    bool running() const { return armed; }
};

/** when C: runs each time C goes from false to true; starts as false (spec 6.4). */
struct When {
    bool prev, on;
    When() : prev(false), on(true) {}
    bool edge(bool c) {
        bool e = c && !prev && on;
        prev = c;
        return e;
    }
    void start() { on = true; }
    void stop() { on = false; }
    void restart() { on = true; }
    bool running() const { return on; }
};

/** A task: a resumable function. The generator writes its body between MUJARD_TASK_BEGIN and MUJARD_TASK_END. */
struct Task {
    int16_t line_;
    Time wake_;
    bool running_;
    Task() : line_(0), wake_(0), running_(true) {}
    void start() { running_ = true; }
    void stop() { running_ = false; }
    void restart() {
        line_ = 0;
        running_ = true;
    }
    bool running() const { return running_; }
};

#define MUJARD_TASK_BEGIN() \
    if (!running_) return; \
    switch (line_) {       \
        case 0:
#define MUJARD_WAIT(t, n)                       \
    do {                                     \
        wake_ = ::mujard::now() + (t);          \
        line_ = (n);                         \
        return;                              \
        case (n):                            \
            if (::mujard::now() < wake_) return; \
    } while (0)
// wait until c: resume each pass and check c again (spec 6.6).
#define MUJARD_WAIT_UNTIL(c, n)                 \
    do {                                     \
        line_ = (n);                         \
        return;                              \
        case (n):                            \
            if (!(c)) return;                \
    } while (0)
#define MUJARD_TASK_END() \
    }                  \
    running_ = false;  \
    line_ = 0;

struct Stopwatch {
    Time started, total;
    bool on;
    Stopwatch() : started(0), total(0), on(false) {}
    void start() {
        if (!on) {
            started = now();
            on = true;
        }
    }
    void stop() {
        if (on) {
            total += now() - started;
            on = false;
        }
    }
    void reset() {
        total = 0;
        started = now();
    }
    Time elapsed() const { return on ? total + (now() - started) : total; }
    bool running() const { return on; }
};

// ---- devices: everything with work to do in the background registers itself

struct Ticked {
    Ticked* next_;
    Ticked();
    Ticked(const Ticked&);  // a copy registers itself too
    Ticked& operator=(const Ticked&) { return *this; }
    virtual ~Ticked();      // and every device unregisters when it goes away
    virtual void init() {}
    virtual void tick(Time t) { (void)t; }
};

class Output : public Ticked {
public:
    Output(uint8_t pin, bool inverted) : pin_(pin), inv_(inverted), state_(false), mode_(0) {}
    void init() override;
    void tick(Time t) override;
    void on() { stop(), set(true); }
    void off() { stop(), set(false); }
    void toggle() { stop(), set(!state_); }
    void write(bool v) { stop(), set(v); }
    bool read() const { return state_; }
    void blink(Time period, Time duration = 0, Time onTime = -1);
    void pattern(const char* p, Time step = 150000, int32_t repeat = -1);
    void pulse(Time length);
    void tone(Freq f, Time length = 0);
    void melody(const char* notes, int32_t tempo = 120);
    void stop();

private:
    void set(bool v);
    bool nextPatternStep(bool* on);
    uint8_t pin_;
    bool inv_, state_;
    uint8_t mode_;  // 0 still, 1 blink, 2 pattern, 3 melody
#ifdef MUJARD_DIRECT_PINS
    volatile uint8_t* port_ = nullptr;  // the pin's output register, once init() found it
    uint8_t mask_ = 0;
#endif
    Time t0_ = 0, period_ = 0, onTime_ = 0, until_ = 0, step_ = 0, stepEnd_ = 0;
    const char* text_ = nullptr;
    int32_t repeat_ = 0;
    uint16_t pos_ = 0;
    uint8_t sub_ = 0;
    int32_t tempo_ = 120;
};

class Input : public Ticked {
public:
    /** pull: 0 none, 1 pullup, 2 pulldown. A pulled-up input is active-low unless inverted (spec 7). */
    Input(uint8_t pin, uint8_t pull, bool inverted) : pin_(pin), pull_(pull), activeLow_((pull == 1) != inverted) {}
    void init() override;
    void tick(Time t) override;
    bool read() const { return stable_; }
    bool pressed() const { return ev_ & 1; }
    bool released() const { return ev_ & 2; }
    bool changed() const { return ev_ & 4; }
    bool clicked() const { return ev_ & 8; }
    bool double_clicked() const { return ev_ & 16; }
    /** The voltage on the wire went LOW to HIGH / HIGH to LOW, whatever pullup or inverted say (spec 6.5). */
    bool rising() const { return ev_ & 32; }
    bool falling() const { return ev_ & 64; }
    bool held(Time t) const { return stable_ && heldFor_ >= t && prevHeld_ < t; }
    /** The program has a double_clicked handler: clicked waits out the 400 ms window (spec 6.5). */
    void wantDouble() { wantsDouble_ = true; }
    void debounce(Time t) { debounce_ = t; }
    Opt<Time> measure_pulse(bool high = true, Time timeout = 1000000);
    bool raw() const;
    uint8_t pin() const { return pin_; }
    /** The wire is LOW when the input is active (a pulled-up button, spec 7). */
    bool activeLow() const { return activeLow_; }

private:
    uint8_t pin_, pull_;
    bool activeLow_, wantsDouble_ = false;
    bool stable_ = false, lastRaw_ = false;
    Time lastChange_ = 0, debounce_ = 20000, pressStart_ = 0, lastClick_ = -1000000000LL;
    Time pendingClick_ = -1;
    Time heldFor_ = 0, prevHeld_ = 0;
    uint8_t ev_ = 0;
#ifdef MUJARD_DIRECT_PINS
    volatile uint8_t* port_ = nullptr;  // the pin's input register, once init() found it
    uint8_t mask_ = 0;
#endif
};

class Pwm : public Ticked {
public:
    Pwm(uint8_t pin) : pin_(pin) {}
    void init() override;
    void tick(Time t) override;
    void write(float ratio);
    void on() { write(1.0f); }
    void off() { write(0.0f); }
    void stop() { fading_ = false; }
    void fade(float to, Time over);
    void frequency(Freq f);
    void limit(float lo, float hi) {
        lo_ = lo;
        hi_ = hi;
    }

private:
    void out(float r);
    uint8_t pin_;
    float level_ = 0, from_ = 0, to_ = 0, lo_ = 0, hi_ = 1;
    Time t0_ = 0, dur_ = 0;
    bool fading_ = false, warned_ = false;
};

class Analog : public Ticked {
public:
    Analog(uint8_t pin) : pin_(pin) {}
    int32_t value() const;
    float volts() const;
    float percent() const;
    void reference(float volts);
    uint8_t pin() const { return pin_; }
    /** The reading this pin gives for a voltage / a share of the range (what a test's pretend sets, spec 12). */
    int32_t rawFor(float volts) const;
    int32_t rawForRatio(float ratio) const;

private:
    uint8_t pin_;
    float ref_ = 0;
};

/** battery_on(analog pin, divider:, type:): a battery's voltage through a resistor divider, and its charge from the
 *  battery type's discharge curve (spec Appendix A, 84-85). volts is the battery's, not the pin's. */
enum BatteryType : uint8_t { BAT_LIPO_1S, BAT_AA_X3, BAT_AA_X4 };
class Battery {
public:
    uint8_t pin() const { return pin_; }
    /** What a test's 'battery.pretend' sets: the pin's raw reading for a voltage / share of range
     *  (the divider's input side - volts() then applies the divider + discharge curve as on a board). */
    int32_t rawFor(float volts) const;
    int32_t rawForRatio(float ratio) const;
    Battery(uint8_t pin, float divider, BatteryType type) : pin_(pin), divider_(divider > 0 ? divider : 1), type_(type) {}
    float volts() const;
    /** 0..1 (a ratio). */
    float percent() const;

private:
    uint8_t pin_;
    float divider_;
    BatteryType type_;
};
/** The charge (0..1) of one cell at this resting voltage: piecewise linear between the points of a discharge curve. */
float battery_charge(BatteryType type, float cellVolts);

void group_write(Output* items, int n, int32_t bits);

// ---- the board

/** Why this run started (spec 4.6). PC tests report the reason their injected reset came from; ESP32/ESP8266,
 *  Optiboot AVR, STM32F103, SAMD21, nRF52840 and Renesas RA4M1 builds read the chip's reset flags. */
enum ResetReason : uint8_t {
    RESET_POWER_ON,
    RESET_RESET_BUTTON,
    RESET_WATCHDOG,
    RESET_BROWNOUT,
    RESET_SOFTWARE,
    RESET_WAKE,
    RESET_CRASH,
};

const char* board_name();
const char* board_chip();
Text board_id();
ResetReason board_reset_reason();
int32_t free_memory();
/** The board's supply voltage where the chip can measure it, or none (spec Appendix A 16). */
Opt<float> board_voltage();
void restart();
/** Sleep the whole program for up to t, or until the wake pin's edge (spec 6.10). pin < 0: timed. */
void power_sleep(Time t, int8_t pin, uint8_t mode);
/** Flush saved vars and sleep until t (or the pin's edge); the board restarts on waking (spec 6.10). */
void power_deep_sleep(Time t, int8_t pin, uint8_t mode);
void watchdog(Time t);

// ---- saved variables (spec 11): kept in EEPROM / flash, written at most every 5 s when they change

// Storage shared with store.save (mujard_store.h): its named slots are the last STORE_AREA bytes.
static const int STORE_AREA = 192;
int storage_begin();  // the storage's size in bytes (0: none)
uint8_t storage_read(uint16_t at);
bool storage_write(uint16_t at, uint8_t v);  // writes only a changed byte; true if it wrote
// false: the commit couldn't run safely right now (a multicore lockout timed out) - the write stays staged.
bool storage_commit();
// Durability observability: storage_commit records every attempt's result so the program can tell a real
// commit from a silent retry - commits_failed counts the failures, commit_ok is the last attempt's result.
extern uint32_t storage_commits_failed;
extern bool storage_commit_ok;
// RP2040/RP2350: a flash write stalls the peer core's XIP fetch, so both cores opt in here so flash_safe_execute
// can freeze the peer during a commit. No-op on boards whose flash write is already driver-managed.
bool coreInitFlashSafe();
int store_area_start();  // -1 when the storage is too small for a store
Time storage_interval();  // the least time between writes (5 s on ESP32, 60 s elsewhere)
extern void (*store_flush_hook)();

// ---- the app: show, plot, notify, control, ask (docs/lang/MUJARD_APP_PROTOCOL.md) - lines on the serial port that
// start with the byte 0x1F, which print and log never send

enum AppVarType : uint8_t { APP_BOOL, APP_INT, APP_FLOAT, APP_I64 };
bool app_on();
// The app's messages can also travel over extra links (mujard.ble, the Wi-Fi socket): every finished message line
// goes to each transport registered with app_mirror_add, and app_input_line takes their lines back. Not on AVR
// (no radio there, and the RAM is scarce).
#if !defined(MUJARD_NO_APP) && !defined(ARDUINO_ARCH_AVR)
#define MUJARD_APP_MIRROR 1
void app_mirror_add(void (*fn)(const char* line, uint16_t n));
void app_input_line(const char* s);
#endif
/** control "Label", var, range: register before app_begin. per = base units per display unit (% -> 0.01). */
void app_control(const char* label, bool isSwitch, float min, float max, float step, const char* unit, float per, void* var, uint8_t type);
/** control "Horn", horn, hold / a joystick's axes: marks the control just registered as held. The app refreshes a held
 *  control while a finger is on it; when it stops for APP_HOLD_TIMEOUT the board puts it back at rest (false, or 0
 *  clamped into the range) itself - the failsafe when the phone drops away mid-drive. */
enum AppHoldKind : uint8_t { APP_HOLD_BUTTON = 2, APP_HOLD_STICK_X = 3, APP_HOLD_STICK_Y = 4 };
void app_control_hold(uint8_t kind);
/** Sends hello and the controls; the generated setup() calls it once, after saved vars are loaded. */
void app_begin(const char* program, uint32_t hash);
/** style: 0 number, 1 gauge, 2 light. number, min and max are in the display unit. */
void app_show(const char* label, const Text& text, bool hasNumber, float number, const char* unit, uint8_t style, bool hasRange, float min, float max);
void app_notify(const Text& message, bool urgent, int line);
/** notify over Wi-Fi (Phase 9): when the build carries MUJARD_NOTIFY_* the http library points this at
 *  its sender and app_notify calls it after the serial line, so a board with no app connected still
 *  reaches the account through the server. Null when the sketch has no such key. */
extern void (*notify_net_fn)(const char* message, bool urgent);
void app_ask_begin(int16_t id, const Text& question, Time timeout);
bool app_ask_done(int16_t id);
Opt<bool> app_ask_result(int16_t id);
/** plot a, b: up to 8 named values in one line. */
class AppPlot {
public:
    AppPlot();
    AppPlot& add(const char* name, float v);
    void end();

private:
    const char* names_[8];
    float values_[8];
    uint8_t count_ = 0;
};

void saved_begin(uint32_t layout);
void saved_add(void* p, uint16_t size);
void saved_load();
/** Write pending saved-var changes now; called before sleep, restart and halt (spec 11). */
void saved_flush();
/** A saved var's live value isn't in storage yet (spec 11's auto-save window). */
bool saved_pending();

}  // namespace mujard
