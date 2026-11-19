// use mujard.spi - the SPI bus (spec Appendix A 75-76): rfid = spi.device(cs: pin 10, speed: 1 MHz, mode: 0) makes
// a device on its own chip-select pin, and rfid.transfer(list) writes the list's bytes while collecting the ones
// the device answers with at the same time. MOSI, MISO and SCK are the board's SPI pins (the checker warns about
// them, W609); SPI.begin() runs on the first transfer.
#pragma once
#include "mujard_rt.h"
#ifndef MUJARD_HOST_TEST
#include <SPI.h>
#endif

namespace mujard {

class SpiDevice {
public:
    /** cs: the chip-select pin; hz: the bus speed in hertz (Mujard gives it in millihertz - the generator
     *  divides); mode: SPI mode 0-3. */
    SpiDevice(int32_t cs, uint32_t hz, int32_t mode) : cs_((uint8_t)cs), mode_((uint8_t)(mode & 3)), hz_(hz ? hz : 1000000) {}

    /** Each byte of data goes out and its answer byte comes back; the result is as long as data, cut at 32. */
    template <int N>
    List<int32_t, 32> transfer(const List<int32_t, N>& data) {
        List<int32_t, 32> out;
        if (halted()) return out;  // a halted program sends nothing: a radio or a display could be listening
#ifdef MUJARD_HOST_TEST
        for (int32_t i = 0; i < data.n && out.n < 32; i++) out.add(data.items[i]);  // the fake bus is a loopback
#else
        ensure();
        SPI.beginTransaction(SPISettings(hz_, MSBFIRST, mode_));
        digitalWrite(cs_, LOW);
        for (int32_t i = 0; i < data.n && out.n < 32; i++) out.add((int32_t)SPI.transfer((uint8_t)data.items[i]));
        digitalWrite(cs_, HIGH);
        SPI.endTransaction();
#endif
        return out;
    }

private:
#ifndef MUJARD_HOST_TEST
    void ensure() {
        if (started_) return;
        started_ = true;
        SPI.begin();
        pinMode(cs_, OUTPUT);
        digitalWrite(cs_, HIGH);
    }
#endif
    uint8_t cs_, mode_;
    uint32_t hz_;
    bool started_ = false;
};

}  // namespace mujard
