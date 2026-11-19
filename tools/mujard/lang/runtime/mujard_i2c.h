// use mujard.i2c - the I2C bus (Wire).
#pragma once
#include <Wire.h>
#include "mujard_rt.h"

namespace mujard {
namespace i2c {

inline void ensure() {
    static bool started = false;
    if (!started) {
        Wire.begin();
        started = true;
    }
}

/** The addresses that answer, like the classic "I2C scanner" sketch. */
inline List<int32_t, 16> scan() {
    ensure();
    List<int32_t, 16> found;
    for (uint8_t a = 1; a < 127; a++) {
        Wire.beginTransmission(a);
        if (Wire.endTransmission() == 0 && found.n < 16) found.items[found.n++] = a;
    }
    return found;
}

template <int N>
inline void write(int32_t address, const List<int32_t, N>& data) {
    if (halted()) return;  // a halted program sends nothing (a motor driver could be listening)
    ensure();
    Wire.beginTransmission((uint8_t)address);
    for (int32_t i = 0; i < data.n; i++) Wire.write((uint8_t)data.items[i]);
    Wire.endTransmission();
}

/** count bytes from a device, or none if it didn't answer. */
inline Opt<List<int32_t, 32> > read(int32_t address, int32_t count) {
    ensure();
    List<int32_t, 32> out;
    if (count > 32) count = 32;
    if (count < 1) return Opt<List<int32_t, 32> >();
    if (Wire.requestFrom((uint8_t)address, (uint8_t)count) != (uint8_t)count) return Opt<List<int32_t, 32> >();
    while (Wire.available() && out.n < count) out.items[out.n++] = Wire.read();
    return Opt<List<int32_t, 32> >(out);
}

}  // namespace i2c

template <class T, int N>
inline Text& add_list(Text& t, const List<T, N>& l) {
    t.add("[");
    for (int32_t i = 0; i < l.n; i++) {
        if (i) t.add(", ");
        t.add(l.items[i]);
    }
    return t.add("]");
}

}  // namespace mujard
