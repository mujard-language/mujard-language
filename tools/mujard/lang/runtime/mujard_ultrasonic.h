// use mujard.ultrasonic - an HC-SR04-style distance sensor.
#pragma once
#include "mujard_rt.h"

namespace mujard {

class Ultrasonic : public Ticked {
public:
    Ultrasonic(uint8_t trig, uint8_t echo) : trig_(trig), echo_(echo) {}
    void init() override {
        pinMode(trig_, OUTPUT);
        digitalWrite(trig_, LOW);
        pinMode(echo_, INPUT);
    }
    /** The distance in metres, or none when nothing answers within about 5 m (30 ms). */
    Opt<float> distance() {
        digitalWrite(trig_, LOW);
        delayMicroseconds(2);
        digitalWrite(trig_, HIGH);
        delayMicroseconds(10);
        digitalWrite(trig_, LOW);
        unsigned long us = pulseIn(echo_, HIGH, 30000UL);
        if (us == 0) return Opt<float>();
        return Opt<float>((float)us * 343.0f / 2.0f / 1000000.0f);  // speed of sound, there and back
    }

private:
    uint8_t trig_, echo_;
};

}  // namespace mujard
