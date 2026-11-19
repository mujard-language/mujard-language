// A fake flash-backed EEPROM (begin / read / write / commit) for the Mujard runtime's PC tests.
#pragma once
#include <cstdint>

struct FakeEEPROM {
    uint8_t mem[1024];
    uint8_t committed[1024];
    int commits = 0;
    FakeEEPROM() {
        for (int i = 0; i < 1024; i++) mem[i] = committed[i] = 0xFF;
    }
    void begin(int) {
        for (int i = 0; i < 1024; i++) mem[i] = committed[i];  // what a reboot would see
    }
    uint8_t read(int at) const { return mem[at]; }
    void write(int at, uint8_t v) { mem[at] = v; }
    int length() const { return 1024; }
    void commit() {
        commits++;
        for (int i = 0; i < 1024; i++) committed[i] = mem[i];
    }
};
extern FakeEEPROM EEPROM;
