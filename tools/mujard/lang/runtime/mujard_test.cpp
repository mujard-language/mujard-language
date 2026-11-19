// mujard_test.cpp - the PC side of a test build (spec 12): the fake board's state, the test helpers, and main(), which
// starts the program and runs it in 1 ms steps of simulated time until the selected test ends. The last line it prints is
// the result the runner reads:
//   MUJARD-TEST pass
//   MUJARD-TEST fail <line> <message>      (line 0: not a line of the test, e.g. the program halted)
#include <cstdio>
#include <cstdlib>
#include <string>

#include "Arduino.h"
#include "EEPROM.h"
#include "WiFi.h"
#include "Wire.h"
#include "mujard_test.h"

namespace fake {
uint64_t us = 1;
int mode[64], level[64], input[64], analogIn[64], pwm[64];
void (*isr[64])(void);
int isrMode[64];
std::vector<ToneCall> tones;
int noTones = 0;
unsigned long pulseResult = 0;
std::string serialOut, serialIn;
bool wifiUp = false, netOpen = false;
std::string netSent, netToRead;
int netConnects = 0;
}  // namespace fake
FakeSerial Serial;
FakeWire Wire;
FakeEEPROM EEPROM;
FakeWiFi WiFi;

void setup();
void loop();

namespace mujard {

int test_selected = 0;
static bool done = false, passed = false;
static int failLine = 0;
static std::string failMessage;

void test_pass() {
    if (!done) done = passed = true;
}

void test_fail(const char* message, int line) {
    if (done) return;
    done = true;
    failLine = line;
    failMessage = message ? message : "";
}

/** The program's own output lines: the app protocol's frames (they start with 0x1F) are not printed text. */
static bool lineHas(bool frames, const char* prefix, const Text& part) {
    size_t at = 0;
    const std::string& out = fake::serialOut;
    while (at < out.size()) {
        size_t end = out.find('\n', at);
        if (end == std::string::npos) end = out.size();
        std::string line = out.substr(at, end - at);
        at = end + 1;
        bool frame = !line.empty() && line[0] == 0x1F;
        if (frame != frames || line.compare(0, strlen(prefix), prefix) != 0) continue;
        if (line.find(part.s, strlen(prefix)) != std::string::npos) return true;
    }
    return false;
}

bool test_printed(const Text& part) { return lineHas(false, "", part); }

/* notified: what the program asked to send (a notification a phone would get), not the
 * serial frame it happens to travel in - app_notify records here under MUJARD_HOST_TEST. */
static std::string notifiedLog;
void test_notify_seen(const Text& message) { notifiedLog += message.s; notifiedLog += '\n'; }
bool test_notified(const Text& part) { return notifiedLog.find(part.s) != std::string::npos; }

static bool haltChecked = false;  // the test asked about a halt (expect stopped()) - an unobserved halt still fails it
bool test_stopped() {
    haltChecked = true;
    return mujard::stopping();
}

void test_pretend_pin(uint8_t pin, bool activeLow, bool active) {
    int was = fake::input[pin], now = (active != activeLow) ? HIGH : LOW;
    fake::input[pin] = now;
    void (*fn)(void) = fake::isr[pin];
    if (fn && was != now) {
        int edge = now == HIGH ? RISING : FALLING;
        if (fake::isrMode[pin] == CHANGE || fake::isrMode[pin] == edge) fn();
    }
}
void test_pretend_analog(uint8_t pin, int32_t raw) { fake::analogIn[pin] = raw; }
void test_pretend_echo(float metres) { fake::pulseResult = metres < 0 ? 0 : (unsigned long)(metres * 2.0f / 343.0f * 1e6f + 0.5f); }
void test_pretend_serial(const Text& line) {
    fake::serialIn += line.s;
    fake::serialIn += '\n';
}

void test_i2c_pretend_bytes(int32_t address, const int32_t* bytes, int32_t n) { Wire.pretend((uint8_t)address, bytes, n); }

void test_i2c_gone(int32_t address) { Wire.gone((uint8_t)address); }

// ---- power_cut() / brown_out(): what survives a reset, kept in a file between the two runs of a test

int test_resume_cut = 0;
static const char* statePath = nullptr;
static ResetReason resetReason = RESET_POWER_ON;

ResetReason test_reset_reason() { return resetReason; }

static void restartWith(ResetReason reason, int k) {
    std::FILE* f = statePath ? std::fopen(statePath, "wb") : nullptr;
    if (!f) {
        test_fail("a test restart needs the test runner (mujard test) to start the board again", 0);
        return;
    }
    uint8_t r = (uint8_t)reason;
    std::fwrite(&r, 1, 1, f);
    std::fwrite(EEPROM.committed, 1, sizeof(EEPROM.committed), f);  // written but not committed: lost, as on a board
    std::fwrite(fake::input, sizeof(fake::input), 1, f);
    std::fwrite(fake::analogIn, sizeof(fake::analogIn), 1, f);
    std::fwrite(&fake::pulseResult, sizeof(fake::pulseResult), 1, f);
    std::fwrite(Wire.dev, sizeof(Wire.dev), 1, f);
    std::fclose(f);
    std::fwrite(fake::serialOut.data(), 1, fake::serialOut.size(), stdout);
    std::printf("\nMUJARD-TEST restart %d %s\n", k, reason == RESET_BROWNOUT ? "brown-out" : "power-cut");
    std::fflush(stdout);
    std::exit(0);
}

void test_power_cut(int k) { restartWith(RESET_POWER_ON, k); }
void test_brown_out(int k) { restartWith(RESET_BROWNOUT, k); }

static void loadPowerCutState() {
    std::FILE* f = std::fopen(statePath, "rb");
    if (!f) return;
    uint8_t r = (uint8_t)RESET_POWER_ON;
    bool ok = std::fread(&r, 1, 1, f) == 1 &&
              std::fread(EEPROM.committed, 1, sizeof(EEPROM.committed), f) == sizeof(EEPROM.committed) &&
              std::fread(fake::input, sizeof(fake::input), 1, f) == 1 &&
              std::fread(fake::analogIn, sizeof(fake::analogIn), 1, f) == 1 &&
              std::fread(&fake::pulseResult, sizeof(fake::pulseResult), 1, f) == 1 &&
              std::fread(Wire.dev, sizeof(Wire.dev), 1, f) == 1;
    std::fclose(f);
    if (ok) {
        resetReason = (ResetReason)r;
        for (int i = 0; i < 1024; i++) EEPROM.mem[i] = EEPROM.committed[i];  // what the board reads after power-up
    }
}

}  // namespace mujard

/** The last error line the program logged ("ERROR (line 14) Division by zero"), for a test that ends in a halt. */
static std::string lastError() {
    const std::string& out = fake::serialOut;
    size_t at = out.rfind("ERROR");
    if (at == std::string::npos) at = out.rfind("HALT");
    if (at == std::string::npos) return "the program stopped";
    size_t end = out.find('\n', at);
    return out.substr(at, end == std::string::npos ? std::string::npos : end - at);
}

int main(int argc, char** argv) {
    mujard::test_selected = argc > 1 ? std::atoi(argv[1]) : 0;
    // argv[2]: the state file power_cut() writes; argv[3]: the cut this run starts after (power_cut, spec 12)
    if (argc > 2) mujard::statePath = argv[2];
    if (argc > 3) {
        mujard::test_resume_cut = std::atoi(argv[3]);
        mujard::loadPowerCutState();
    }
    const uint64_t limit = fake::us + 3600ULL * 1000000ULL;  // an hour of simulated time
    setup();
    std::string haltedNote;  // the halt a test's "stopped()" can still observe; context if it times out
    while (!mujard::done && fake::us < limit) {
        loop();
        if (mujard::halted() && haltedNote.empty()) haltedNote = lastError();
        fake::us += 1000;
    }
    if (!mujard::done) {
        std::string m = haltedNote.empty() ? "the test didn't finish within an hour of simulated time"
                                           : "the program halted: " + haltedNote;
        mujard::test_fail(m.c_str(), 0);
    }
    std::fwrite(fake::serialOut.data(), 1, fake::serialOut.size(), stdout);
    // A halt the test never checked with stopped()/halted() still fails it - the pass must mean the program behaved.
    if (mujard::passed && mujard::halted() && !mujard::haltChecked) {
        mujard::passed = false;
        mujard::failLine = 0;
        mujard::failMessage = "the program halted: " + haltedNote;
    }
    if (mujard::passed) {
        std::printf("\nMUJARD-TEST pass\n");
    } else {
        std::string m = mujard::failMessage;
        for (char& c : m) {
            if (c == '\n' || c == '\r') c = ' ';
        }
        std::printf("\nMUJARD-TEST fail %d %s\n", mujard::failLine, m.c_str());
    }
    return 0;
}
