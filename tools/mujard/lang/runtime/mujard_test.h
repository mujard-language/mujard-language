// mujard_test.h - what a program's 'test' blocks use (spec 12). A test build runs on the PC against a fake Arduino
// (runtime/tests/shim) with simulated time, one fresh board per test; it is never compiled for a board.
//
// The program starts up as it would on the board (safe_state, globals, 'at start', tasks); the test then runs as a task
// beside it: 'advance(t)' lets the whole program run for t of simulated time, 'pretend' sets what a sensor or input
// reads from then on, and a failed 'expect' ends the test.
#pragma once
#include "mujard_rt.h"

namespace mujard {

/** Which test this run is for (the command line's first argument). */
extern int test_selected;

/** power_cut() / brown_out() (spec 12): which restart this run starts after (0: the test's first run).
 *  test_power_cut(k) / test_brown_out(k) save what a board keeps - the EEPROM/flash as last committed, and the
 *  outside world (input levels, analog readings, the echo, pretend I2C devices) - print "MUJARD-TEST restart k" and
 *  end the run; RunTests starts the test again on a fresh board with that state, after restart k. */
extern int test_resume_cut;
void test_power_cut(int k);
void test_brown_out(int k);
/** What board.reset_reason reports in this test build (the reason saved with the last restart). */
ResetReason test_reset_reason();

/** The test reached its end / a check failed: the result the runner reads. */
void test_pass();
void test_fail(const char* message, int line);

/** printed("x") / notified("x"): did the program print a line / send a notification containing x (spec 12)? */
bool test_printed(const Text& part);
bool test_notified(const Text& part);
/** True once the program is stopping or halted - and marks the halt as observed, so a test that ends having checked
 *  it isn't then failed by that halt. */
bool test_stopped();
/** app_notify calls this under MUJARD_HOST_TEST so 'expect notified' sees the ask, not the wire frame. */
void test_notify_seen(const Text& message);

/** pretend: what an input reads (its active level, spec 7), an analog pin's reading, an HC-SR04's distance in metres
 *  (negative: nothing answers). They set the fake board's wires, so debouncing and events work as on a board. */
void test_pretend_pin(uint8_t pin, bool activeLow, bool active);
void test_pretend_analog(uint8_t pin, int32_t raw);
void test_pretend_echo(float metres);
/** serial.pretend("x"): a line the program reads as if it arrived on the serial port (then events and
 *  serial.read_line() work as on a board). A Uart's own pretend() feeds its port the same way. */
void test_pretend_serial(const Text& line);
/** i2c.pretend(address, reply): a device on the fake bus that answers with reply's bytes; i2c.pretend(address, none)
 *  takes it off the bus (a sensor that fails mid-run). */
void test_i2c_pretend_bytes(int32_t address, const int32_t* bytes, int32_t n);
void test_i2c_gone(int32_t address);
template <int N>
inline void test_i2c_pretend(int32_t address, const List<int32_t, N>& reply) {
    test_i2c_pretend_bytes(address, reply.items, reply.n);
}
/** answer true / answer false: replies to the ask the program is waiting on; answer none leaves it to run out
 *  of time. Implemented in mujard_rt.cpp, next to the asks. */
void test_answer(int yes);

}  // namespace mujard
