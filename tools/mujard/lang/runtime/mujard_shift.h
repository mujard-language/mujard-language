// use mujard.shift - shift registers such as the 74HC595 (more outputs from three pins).
#pragma once
#include "mujard_rt.h"

namespace mujard {

class ShiftOut : public Ticked {
public:
    ShiftOut(uint8_t data, uint8_t clock, uint8_t latch) : data_(data), clock_(clock), latch_(latch) {}
    void init() override {
        pinMode(data_, OUTPUT);
        pinMode(clock_, OUTPUT);
        pinMode(latch_, OUTPUT);
    }
    void write(int32_t bits) {
        if (halted()) return;
        digitalWrite(latch_, LOW);
        shiftOut(data_, clock_, MSBFIRST, (uint8_t)bits);
        digitalWrite(latch_, HIGH);
    }

private:
    uint8_t data_, clock_, latch_;
};

}  // namespace mujard
