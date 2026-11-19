#pragma once
// A fake TwoWire for PC test builds. The simulated board has no real I2C bus: by default nothing answers, so a scan
// finds nothing and reads return no data - honest "nothing connected" behavior. A test can put a pretend device on
// the bus (i2c.pretend(0x48, reply), spec 12): it then acknowledges its address and answers every read with its reply
// bytes, from the first, repeating if more are asked for. i2c.pretend(0x48, none) takes it off the bus again - a sensor
// that fails while the program runs.
#include <stdint.h>
#include <stddef.h>

struct FakeWire {
    struct Dev {
        bool present = false;
        uint8_t data[32] = {};
        uint8_t n = 0;
    };
    Dev dev[128];
    uint8_t addr = 0, left = 0, pos = 0;
    uint32_t writes = 0;  // bytes the program sent to present devices (a test can't see them yet; kept for the driver)

    void pretend(uint8_t a, const int32_t* bytes, int32_t n) {
        Dev& d = dev[a & 0x7F];
        d.present = true;
        d.n = (uint8_t)(n > 32 ? 32 : (n < 0 ? 0 : n));
        for (uint8_t i = 0; i < d.n; i++) d.data[i] = (uint8_t)bytes[i];
    }
    void gone(uint8_t a) { dev[a & 0x7F] = Dev(); }

    void begin() {}
    void beginTransmission(uint8_t a) { addr = a & 0x7F; }
    uint8_t endTransmission() { return dev[addr].present ? 0 : 2; }  // 2: the address got no acknowledge
    size_t write(uint8_t) {
        if (dev[addr].present) writes++;
        return 1;
    }
    uint8_t requestFrom(uint8_t a, uint8_t count) {
        addr = a & 0x7F;
        pos = 0;
        left = dev[addr].present && dev[addr].n > 0 ? count : 0;
        return left;
    }
    int available() { return left; }
    int read() {
        if (!left) return -1;
        left--;
        const Dev& d = dev[addr];
        return d.data[pos++ % d.n];
    }
};
extern FakeWire Wire;
