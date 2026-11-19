// A fake Servo library for runs on the PC (the runtime's tests and mujard test): it remembers what it was told.
#pragma once
#include "Arduino.h"

class Servo {
public:
    uint8_t attach(int pin) {
        pin_ = pin;
        return 1;
    }
    void detach() { pin_ = -1; }
    void write(int value) { value_ = value; }
    void writeMicroseconds(int us) { value_ = us; }
    int read() { return value_; }
    bool attached() { return pin_ >= 0; }

private:
    int pin_ = -1, value_ = 90;
};
