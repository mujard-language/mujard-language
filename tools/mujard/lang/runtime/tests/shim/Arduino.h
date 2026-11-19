// A fake Arduino for testing the Mujard runtime on a PC: time, pins, the serial port and tones are all simulated and
// can be read back by the tests. Only what mujard_rt uses.
#pragma once
#include <cctype>
#include <cmath>
#include <cstdint>
#include <cstdio>
#include <cstdlib>
#include <cstring>
#include <ctime>
#include <string>
#include <vector>

#if defined(_WIN32)
// MinGW knows localtime_s / gmtime_s (Windows argument order), not the POSIX _r spellings the cores use.
#define localtime_r(t, out) (localtime_s(out, t), out)
#define gmtime_r(t, out) (gmtime_s(out, t), out)
#endif

#define HIGH 1
#define LOW 0
#define INPUT 0
#define OUTPUT 1
#define INPUT_PULLUP 2
#define INPUT_PULLDOWN 3
#define LED_BUILTIN 13
#define MSBFIRST 1

namespace fake {
extern uint64_t us;               // the simulated clock
extern int mode[64], level[64];   // pinMode / digitalWrite results
extern int input[64];             // what digitalRead returns
extern int analogIn[64];          // what analogRead returns
extern int pwm[64];               // the last analogWrite value
struct ToneCall {
    int pin;
    unsigned freq;
    unsigned long ms;  // 0 = until noTone
    uint64_t at;
};
extern std::vector<ToneCall> tones;
extern int noTones;
extern unsigned long pulseResult;  // what pulseIn returns
extern std::string serialOut, serialIn;
void reset();
}  // namespace fake

inline unsigned long micros() { return (unsigned long)(uint32_t)fake::us; }
inline unsigned long millis() { return (unsigned long)(fake::us / 1000); }
inline void delay(unsigned long ms) { fake::us += (uint64_t)ms * 1000; }
inline void delayMicroseconds(unsigned int u) { fake::us += u; }
inline void pinMode(uint8_t p, uint8_t m) {
    fake::mode[p] = m;
    if (m == INPUT_PULLUP) fake::input[p] = HIGH;  // nothing pulls it down yet: the pull-up holds it high
}
inline void digitalWrite(uint8_t p, uint8_t v) { fake::level[p] = v; }
inline int digitalRead(uint8_t p) { return fake::input[p]; }
inline int analogRead(uint8_t p) { return fake::analogIn[p]; }
inline void analogWrite(uint8_t p, int v) { fake::pwm[p] = v; }
inline void tone(uint8_t p, unsigned int f, unsigned long ms = 0) { fake::tones.push_back({p, f, ms, fake::us}); }
inline void noTone(uint8_t) { fake::noTones++; }
inline unsigned long pulseIn(uint8_t, uint8_t, unsigned long) { return fake::pulseResult; }
inline long random(long lo, long hi) { return lo + std::rand() % (hi - lo); }
inline void randomSeed(unsigned long s) { std::srand((unsigned)s); }
inline char* ltoa(long v, char* b, int) {
    std::sprintf(b, "%ld", v);
    return b;
}
inline char* ultoa(unsigned long v, char* b, int) {
    std::sprintf(b, "%lu", v);
    return b;
}
inline void shiftOut(uint8_t, uint8_t, uint8_t, uint8_t) {}
inline void noInterrupts() {}
inline void interrupts() {}

// interrupt pins ('on pin.x fast' events): the fake board calls the registered handler
// from test_pretend_pin when the wire changes, the same way real silicon calls the ISR
// on the matching edge.
#define CHANGE 1
#define RISING 2
#define FALLING 3
#define digitalPinToInterrupt(p) (p)
namespace fake { extern void (*isr[64])(void); extern int isrMode[64]; }
inline void attachInterrupt(uint8_t num, void (*fn)(void), int mode) {
    fake::isr[num] = fn;
    fake::isrMode[num] = mode;
}
inline void detachInterrupt(uint8_t num) { fake::isr[num] = nullptr; }

struct FakeSerial {
    void begin(unsigned long) {}
    void end() {}
    int available() { return (int)fake::serialIn.size(); }
    int read() {
        if (fake::serialIn.empty()) return -1;
        int c = (unsigned char)fake::serialIn[0];
        fake::serialIn.erase(0, 1);
        return c;
    }
    void println(const char* s) {
        fake::serialOut += s;
        fake::serialOut += "\n";
    }
    size_t write(char c) {
        if (c != '\r') fake::serialOut += c;  // like a terminal: line ends read as "\n"
        return 1;
    }
    size_t write(uint8_t c) { return write((char)c); }
    size_t print(const char* s) {
        size_t n = 0;
        for (; *s; s++) n += write(*s);
        return n;
    }
};
extern FakeSerial Serial;
