// use mujard.servo - a hobby servo with smooth moves (Servo on AVR/RP2040/ESP8266, ESP32Servo on the ESP32).
#pragma once
#include "mujard_rt.h"
#if defined(ARDUINO_ARCH_ESP32)
#include <ESP32Servo.h>
#else
#include <Servo.h>
#endif

namespace mujard {

class ServoMotor : public Ticked {
public:
    explicit ServoMotor(uint8_t pin) : pin_(pin) {}
    void init() override {
        servo_.attach(pin_);
        servo_.write((int)angle_);
    }
    /** A program limit ('limit s to 0 deg..90 deg'): later moves clamp to it; the first clamp logs a warning once,
     *  like Pwm's limit (spec 11.3). */
    void limit(float lo, float hi) {
        lo_ = lo;
        hi_ = hi;
    }
    /** Turns to deg (0-180, or the program's limit), over a time if given (non-blocking). */
    void move(float deg, Time over = 0) {
        if (deg < lo_ || deg > hi_) {
            if (!warned_) {
                warned_ = true;
                Out o("WARN");
                o.put("Servo angle limited to ").put_qty(deg < lo_ ? lo_ : hi_, D_ANGLE);
                o.end();
            }
            deg = deg < lo_ ? lo_ : hi_;
        }
        if (deg < 0) deg = 0;
        if (deg > 180) deg = 180;
        from_ = angle_;
        to_ = deg;
        t0_ = now();
        dur_ = over;
        if (over <= 0) set(deg);
    }
    float angle() const { return angle_; }
    void tick(Time t) override {
        if (dur_ <= 0 || angle_ == to_) return;
        float k = (float)(t - t0_) / (float)dur_;
        set(k >= 1 ? to_ : from_ + (to_ - from_) * k);
        if (k >= 1) dur_ = 0;
    }

private:
    void set(float deg) {
        if (halted()) return;
        angle_ = deg;
        servo_.write((int)(deg + 0.5f));
    }
    Servo servo_;
    uint8_t pin_;
    float angle_ = 0, from_ = 0, to_ = 0, lo_ = 0, hi_ = 180;
    Time t0_ = 0, dur_ = 0;
    bool warned_ = false;
};

}  // namespace mujard
