// use mujard.timer - a hardware timer object (spec 12). t.every/t.freq/t.once set a precise trigger - a real
// timer interrupt, not the cooperative 'every' - and 'on t.tick' runs on the scheduler while 'on t.tick fast'
// runs in the ISR itself. AVR's Timer1, ESP32's hw_timer, RP2040's repeating_timer; PC test builds fake it.
#pragma once
#include "mujard_rt.h"
#if !defined(MUJARD_HOST_TEST)
#if defined(ARDUINO_ARCH_AVR) || defined(ARDUINO_ARCH_SAM)
#define MUJARD_TIMER_AVR 1
#elif defined(ARDUINO_ARCH_ESP32)
#define MUJARD_TIMER_ESP32 1
#elif defined(ARDUINO_ARCH_RP2040)
#include "pico/time.h"
#define MUJARD_TIMER_RP2040 1
#endif
#endif

namespace mujard {

/** A hardware timer: arm it, and it fires tick() at ISR speed (or sets the scheduler's 'due' flag). */
class Timer : public Ticked {
public:
    /** t.every(1 s): repeat at period. */
    void every(Time period) { arm(period, false); }
    /** t.freq(1 khz): every() by frequency. */
    void freq(float hz) {
        if (hz <= 0) return;
        arm(Time(1000000.0 / hz), false);
    }
    /** t.once(500 ms): fire once. */
    void once(Time period) { arm(period, true); }
    void start() { running_ = true; hwStart(); }
    void stop() { running_ = false; hwStop(); }
    /** µs the timer has been running; 0 when stopped. */
    Time elapsed() const { return running_ ? now() - at_ : 0; }
    /** 'on t.tick': true in the scheduler pass after the ISR fired. */
    bool due() {
        bool d = due_;
        due_ = false;
        return d;
    }
    /** 'on t.tick fast' runs the given fn in the ISR (a flag on a PC test). */
    void onTick(void (*isr)()) { isr_ = isr; }
    /** The ISR entry point: arm/rearm and set the scheduler flag (or call the fast handler). */
    void fire() {
        due_ = true;
        if (isr_ != nullptr) isr_();
        if (once_) stop();
        else if (running_) at_ += period_;  // drift-free: next due is one period after the last
    }
    void tick(Time) override {
#ifdef MUJARD_HOST_TEST
        if (!running_ || period_ <= 0) return;
        Time t = now();
        if (once_) {
            if (t - at_ >= period_) fire();
        } else {
            while (t - at_ >= period_) fire();
        }
#endif
    }
#ifdef MUJARD_HOST_TEST
    void pretend() { fire(); }
#endif
    static void hwFire(void* self) { ((Timer*)self)->fire(); }

private:
    void arm(Time period, bool once) {
        period_ = period;
        once_ = once;
        at_ = now();
        running_ = true;
        hwArm(period, once);
    }
    // hardware side: start/stop/arm for the board's timer backend (a no-op on a PC test).
    void hwArm(Time period, bool once);
    void hwStart();
    void hwStop();

    Time period_ = 0, at_ = 0;
    bool running_ = false, once_ = false, due_ = false;
    void (*isr_)() = nullptr;
#ifdef MUJARD_TIMER_ESP32
    hw_timer_t* hwt_ = nullptr;
#endif
#ifdef MUJARD_TIMER_RP2040
    repeating_timer_t rt_;
#endif
};

// ---- board backends: AVR's Timer1 CTC, ESP32's hw_timer alarm, RP2040's repeating_timer ----
#if defined(MUJARD_TIMER_AVR)
// One Timer1 per program for now: the ISR dispatches to the one armed Timer.
static Timer* avr_timer = nullptr;
ISR(TIMER1_COMPA_vect) {
    if (avr_timer != nullptr) avr_timer->fire();
}
inline void Timer::hwArm(Time period, bool) {
    avr_timer = this;
    // CTC on OCR1A at F_CPU/8: ticks = 2 us at 16 MHz. period_us/2, capped at the 16-bit count.
    uint32_t ticks = uint32_t(period) / 2;
    if (ticks > 65535) ticks = 65535;
    if (ticks < 1) ticks = 1;
    TCCR1A = 0;
    TCCR1B = _BV(WGM12) | _BV(CS11);
    OCR1A = uint16_t(ticks - 1);
    TIMSK1 |= _BV(OCIE1A);
    sei();
}
inline void Timer::hwStart() { avr_timer = this; }
inline void Timer::hwStop() { avr_timer = nullptr; TIMSK1 &= ~_BV(OCIE1A); }
#elif defined(MUJARD_TIMER_ESP32)
static bool esp32_timer_cb(hw_timer_t* t) {
    // find the Timer whose hwt_ this is; a program has the Timer objects it made
    // (one armed timer per hw timer for now).
    extern Timer* esp32_timer_for(hw_timer_t*);
    if (Timer* self = esp32_timer_for(t)) self->fire();
    return true;
}
static hw_timer_t* esp32_hwt = nullptr;
static Timer* esp32_timer = nullptr;
Timer* esp32_timer_for(hw_timer_t* t) { return t == esp32_hwt ? esp32_timer : nullptr; }
inline void Timer::hwArm(Time period, bool once) {
    if (hwt_ == nullptr) {
        hwt_ = timerBegin(1, 80, true);
        timerAttachInterrupt(hwt_, &esp32_timer_cb, true);
        esp32_hwt = hwt_;
    }
    esp32_timer = this;
    timerAlarmWrite(hwt_, (uint64_t)period, !once);
}
inline void Timer::hwStart() { timerAlarmEnable(hwt_); }
inline void Timer::hwStop() { timerAlarmDisable(hwt_); }
#elif defined(MUJARD_TIMER_RP2040)
static Timer* rp2040_timer = nullptr;
static bool rp2040_timer_cb(repeating_timer_t*) {
    if (rp2040_timer != nullptr) rp2040_timer->fire();
    return rp2040_timer != nullptr && !rp2040_timer->once_;  // once cancels the repeat
}
inline void Timer::hwArm(Time period, bool) {
    add_repeating_timer_us(-(int64_t)period, rp2040_timer_cb, nullptr, &rt_);
}
inline void Timer::hwStart() { rp2040_timer = this; }
inline void Timer::hwStop() { rp2040_timer = nullptr; cancel_repeating_timer(&rt_); }
#else
inline void Timer::hwArm(Time, bool) {}
inline void Timer::hwStart() {}
inline void Timer::hwStop() {}
#endif

}  // namespace mujard
