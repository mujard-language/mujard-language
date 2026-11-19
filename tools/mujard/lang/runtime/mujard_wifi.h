// wifi.connect / wifi.connected / on wifi.joined / on wifi.lost (spec Appendix A 77-78), for the ESP32, ESP8266 and
// Pico W. Connecting happens in the background: connect() returns at once, and the link reconnects by itself when the
// network drops (every 10 s, on top of what the core does).
//
// On a board without Wi-Fi (Uno, Nano, Mega) there is no link at all: the checker keeps wifi.* calls out, and the
// stub below just makes the headers that build on this one (clock, http, mqtt) compile.
#pragma once
#include "mujard_rt.h"
#if defined(ARDUINO_ARCH_ESP8266)
#include <ESP8266WiFi.h>
#define MUJARD_WIFI_HW 1
#elif defined(ARDUINO_ARCH_ESP32) || defined(ARDUINO_ARCH_RP2040) || defined(MUJARD_HOST_TEST)
#include <WiFi.h>
#define MUJARD_WIFI_HW 1
#endif
#if defined(MUJARD_WIFI_HW) && !defined(MUJARD_HOST_TEST)
#if defined(ARDUINO_ARCH_ESP32)
#include <ESPmDNS.h>
#elif defined(ARDUINO_ARCH_ESP8266)
#include <ESP8266mDNS.h>
#elif defined(ARDUINO_ARCH_RP2040)
#include <LEAmDNS.h>
#endif
#endif

namespace mujard {

#ifdef MUJARD_APP_MIRROR
static void wifiAppMirror(const char* s, uint16_t n);  // defined after wifi_link below
#endif

#ifdef MUJARD_WIFI_HW
class WifiLink : public Ticked {
public:
    /** ssid and password are the build's secret macros (spec 11.1): string literals that live for the whole program. */
    void connect(const char* ssid, const char* password) {
        ssid_ = ssid;
        pass_ = password;
        wanted_ = true;
        WiFi.mode(WIFI_STA);
        WiFi.begin(ssid_, pass_);
        lastTry_ = now();
    }
    bool connected() const { return WiFi.status() == WL_CONNECTED; }
    /** on wifi.joined / on wifi.lost: true in the pass the link came up / went down. */
    bool joined() const { return joined_; }
    bool lost() const { return lost_; }

    /** wifi.serve("name"): advertise _mujard._tcp on mDNS and listen on TCP 8266 for the app's lines
        (docs/lang/MUJARD_APP_PROTOCOL.md) - the same dashboard as USB or BLE, over the home network.
        wifi.serve("name", "1234") asks each connecting phone for that PIN first ('M1 pair-required',
        then 'M1 pair <pin>'): until 'M1 paired' the socket reads and watches but its set/answer
        lines are dropped, so only someone holding the PIN can drive it. */
    void serve(const char* name, const char* pin = "") {
        // The PIN may be set or changed on any call; the name and the listener start once.
        uint8_t o = 0;
        for (const char* p = pin; *p && o < sizeof(pin_) - 1; p++) pin_[o++] = *p;
        pin_[o] = 0;
        fails_ = 0;
#ifdef MUJARD_HOST_TEST
        paired_ = pin_[0] == 0;
#endif
        if (serveWanted_) return;
        serveWanted_ = true;
        hostify(name);
#ifdef MUJARD_APP_MIRROR
        app_mirror_add(wifiAppMirror);
#endif
#ifdef MUJARD_HOST_TEST
        clientUp_ = true;   // the fake socket is "up" once serving
#endif
    }
    /** wifi.client: a phone is on the socket right now. */
    bool client() const { return clientUp_; }

    /** The app mirror target: one finished app line, LF included. */
    void sendApp(const char* s, uint16_t n) {
#ifdef MUJARD_HOST_TEST
        Text t;
        for (uint16_t i = 0; i < n && s[i] != '\n' && t.n < MUJARD_TEXT_MAX; i++) t.s[t.n++] = s[i];
        t.s[t.n] = 0;
        if (sent_.n < 4) sent_.add(t);
#else
        if (clientUp_ && client_.connected()) client_.write((const uint8_t*)s, n);
#endif
    }

#ifdef MUJARD_HOST_TEST
    void pretend(const Text& line) { pushRx(line); }
    List<Text, 4> sent_;
#endif

    void tick(Time t) override {
        joined_ = lost_ = false;
        pumpServe();   // the socket serves whether or not the program also calls wifi.connect
        if (!wanted_) return;
        bool up = connected();
        if (up != up_) {
            up_ = up;
            if (up) joined_ = true;
            else lost_ = true;
        }
        if (!up && t - lastTry_ > 10000000) {
            lastTry_ = t;
            WiFi.disconnect();
            WiFi.begin(ssid_, pass_);
        }
    }

private:
    /** The socket half: starts the listener once the link is up, pumps lines, drops it on link loss. */
    void pumpServe() {
#ifdef MUJARD_HOST_TEST
        routeApp();
#else
        if (!serveWanted_) return;
        if (connected() && !serverUp_) {
            server_.begin();
            if (host_[0] != 0) {
                MDNS.begin(host_);
                MDNS.addService("_mujard", "_tcp", 8266);
            }
            serverUp_ = true;
        }
        if (serverUp_ && !connected()) {
            client_.stop();
            server_.close();
            clientUp_ = serverUp_ = false;
            rxN_ = 0;
        }
        if (!serverUp_) return;
#if defined(ARDUINO_ARCH_ESP8266) || defined(ARDUINO_ARCH_RP2040)
        MDNS.update();
#endif
        // One client, like the BLE link's one central; a second is turned away.
        WiFiClient extra = server_.available();
        if (extra) {
            if (clientUp_) extra.stop();
            else {
                client_ = extra;
                clientUp_ = true;
                rxN_ = 0;
                // A PIN-gated socket challenges each new connection; an open one is paired from the start.
                paired_ = pin_[0] == 0;
                fails_ = 0;
                if (!paired_) client_.write("\x1fM1 pair-required\n");
            }
        }
        for (uint8_t budget = 128; clientUp_ && budget-- > 0 && client_.available() > 0;) {
            char c = (char)client_.read();
            if (rxN_ >= sizeof(rxBuf_) - 1) { rxN_ = 0; continue; }  // no LF before the cap: the runaway line is dropped
            rxBuf_[rxN_++] = c;
            if (c == '\n') {
                rxBuf_[rxN_] = 0;
                gotLine();
                rxN_ = 0;
            }
        }
        if (clientUp_ && !client_.connected()) { client_.stop(); clientUp_ = false; rxN_ = 0; }
#endif
    }

    /** The mDNS host name: letters, digits and hyphens only, at most 32. */
    void hostify(const char* name) {
#if !defined(MUJARD_HOST_TEST)
        uint8_t o = 0;
        for (const char* p = name; *p && o < sizeof(host_) - 1; p++) {
            char c = *p;
            if (c >= 'A' && c <= 'Z') c += 'a' - 'A';
            if ((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '-') host_[o++] = c;
        }
        host_[o] = 0;
#else
        (void)name;
#endif
    }

    /** One app line, whichever way it arrived: pair lines answer here, writes need pairing when a PIN is
        set, everything else goes to the app protocol like a serial line. */
    void feedLine(const char* s) {
#ifdef MUJARD_APP_MIRROR
        if (s[0] != 0x1F) return;
        if (strncmp(s + 1, "M1 pair\t", 8) == 0) {
            // Five wrong guesses on one connection lock it out and drop the socket to stop PIN brute-forcing.
            if (fails_ >= 5) {
                sendLine("\x1fM1 pair-locked\n");
                return;
            }
            // 'M1 pair <pin>': the pin is everything after the tab, trimmed of line ends.
            const char* given = s + 9;
            bool ok = pin_[0] != 0 && strncmp(given, pin_, sizeof(pin_) - 1) == 0 && strlen(given) == strlen(pin_);
            paired_ = ok;
            if (ok) {
                fails_ = 0;
                sendLine("\x1fM1 paired\n");
            } else {
                fails_++;
                if (fails_ >= 5) {
                    sendLine("\x1fM1 pair-locked\n");
#if defined(MUJARD_WIFI_HW) && !defined(MUJARD_HOST_TEST)
                    if (clientUp_) { client_.stop(); clientUp_ = false; }
#endif
                } else {
                    sendLine("\x1fM1 pair-failed\n");
                }
            }
            return;
        }
        if (pin_[0] != 0 && !paired_ &&
            (strncmp(s + 1, "M1 set\t", 7) == 0 || strncmp(s + 1, "M1 answer\t", 10) == 0)) return;
        app_input_line(s);
#endif
    }

    /** One protocol line to the client (answers to pair challenges). */
    void sendLine(const char* s) {
#if defined(MUJARD_WIFI_HW) && !defined(MUJARD_HOST_TEST)
        if (clientUp_ && client_.connected()) client_.write(s);
#else
        sendApp(s, (uint16_t)strlen(s));
#endif
    }

#ifndef MUJARD_HOST_TEST
    /** One line off the socket: the app protocol (0x1F first) goes to the app, anything else is ignored. */
    void gotLine() {
#ifdef MUJARD_APP_MIRROR
        if (rxBuf_[0] != 0x1F) return;
        while (rxN_ > 0 && (rxBuf_[rxN_ - 1] == '\n' || rxBuf_[rxN_ - 1] == '\r')) rxBuf_[--rxN_] = 0;
        feedLine(rxBuf_);
#endif
    }
#else
    /** The fake socket: queued lines; app lines reach the app protocol, the rest are ignored. */
    void routeApp() {
#ifdef MUJARD_APP_MIRROR
        for (uint8_t i = 0; i < rx_.n;) {
            if (rx_.items[i].s[0] != 0x1F) { i++; continue; }
            Text t = rx_.items[i];
            rx_.remove(i);
            while (t.n > 0 && (t.s[t.n - 1] == '\n' || t.s[t.n - 1] == '\r')) t.s[--t.n] = 0;
            feedLine(t.s);
        }
#endif
    }
    void pushRx(const Text& line) {
        if (rx_.n >= 4) { dropped_events++; return; }
        rx_.add(line);
    }
    List<Text, 4> rx_;
#endif

    const char* ssid_ = "";
    const char* pass_ = "";
    bool wanted_ = false, up_ = false, joined_ = false, lost_ = false;
    Time lastTry_ = 0;
    bool clientUp_ = false, serveWanted_ = false, paired_ = false;
    uint8_t fails_ = 0;
    char pin_[17] = "";
#if defined(MUJARD_WIFI_HW) && !defined(MUJARD_HOST_TEST)
    WiFiServer server_{8266};
    WiFiClient client_;
    char host_[33] = "";
    char rxBuf_[256];
    uint16_t rxN_ = 0;
    bool serverUp_ = false;
#endif
};
#else
class WifiLink {  // a board without Wi-Fi: never up, never joins or loses, never serves
public:
    void connect(const char* ssid, const char* password) {
        (void)ssid;
        (void)password;
    }
    bool connected() const { return false; }
    bool joined() const { return false; }
    bool lost() const { return false; }
    void serve(const char* name) { (void)name; }
    bool client() const { return false; }
    void sendApp(const char* s, uint16_t n) { (void)s; (void)n; }
};
#endif

static WifiLink wifi_link;

#ifdef MUJARD_APP_MIRROR
static void wifiAppMirror(const char* s, uint16_t n) { wifi_link.sendApp(s, n); }
#endif

// What a wired Ethernet link (mujard_ethernet.h) reports each scheduler turn; on a program without it this
// stays false and net_up() is just the Wi-Fi link.
static bool ethernet_link_up_ = false;

/** net_up(): either network link is up - Wi-Fi or Ethernet; http/mqtt/clock take whichever is connected. */
inline bool net_up() {
    return wifi_link.connected() || ethernet_link_up_;
}

}  // namespace mujard
