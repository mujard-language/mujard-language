// A fake Wi-Fi stack for the Mujard runtime's PC tests: a link that is up when the test says so, and a WiFiClient whose
// sent bytes are recorded and whose received bytes the test supplies.
#pragma once
#include <cstdint>
#include <string>

#define WL_CONNECTED 3
#define WIFI_STA 1

namespace fake {
extern bool wifiUp;
extern std::string netSent, netToRead;
extern bool netOpen;
extern int netConnects;
}  // namespace fake

struct FakeWiFi {
    void mode(int) {}
    void begin(const char*, const char*) {}
    void disconnect() {}
    int status() const { return fake::wifiUp ? WL_CONNECTED : 0; }
};
extern FakeWiFi WiFi;

struct WiFiClient {
    bool connect(const char*, uint16_t) {
        fake::netConnects++;
        fake::netOpen = true;
        return true;
    }
    bool connected() { return fake::netOpen; }
    int available() { return (int)fake::netToRead.size(); }
    int read() {
        if (fake::netToRead.empty()) return -1;
        int c = (uint8_t)fake::netToRead[0];
        fake::netToRead.erase(0, 1);
        return c;
    }
    size_t write(const uint8_t* b, size_t n) {
        fake::netSent.append((const char*)b, n);
        return n;
    }
    void stop() { fake::netOpen = false; }
};
