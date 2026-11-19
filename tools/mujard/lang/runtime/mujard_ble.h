// use mujard.ble - the board advertises as a BLE peripheral and exchanges text lines with a phone or another
// board (Nordic-UART shape, spec 12). ESP32 family via BLEDevice, Nano 33 BLE via ArduinoBLE; PC test builds
// keep a fake link, like uart.pretend.
#pragma once
#include "mujard_rt.h"
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
#include <BLEDevice.h>
#include <BLE2902.h>  // the notify descriptor; BLEDevice.h doesn't bring it in (esp32 core 3.x)
#endif
#if defined(ARDUINO_ARCH_MBED) && defined(ARDUINO_ARDUINO_NANO33BLE) && !defined(MUJARD_HOST_TEST)
#include <ArduinoBLE.h>
#endif
#if defined(ARDUINO_ARCH_RP2040) && defined(PICO_CYW43_SUPPORTED) && !defined(MUJARD_HOST_TEST)
#include <BLE.h>
#endif

namespace mujard {

/** The board's BLE link: advertises under a name, carries lines of text each way. */
class Ble : public Ticked {
public:
    /** ble.start("name"): advertise under name so phones and boards can connect. */
    void start(const char* name);
    /** ble.send("line"): tell whatever is connected; false when nothing is listening. */
    bool send(const Text& line);
    bool connected() { return central_; }
    /** One line for 'on ble.line as l:' - like serial.line. */
    bool line(Text& out);
    /** ble.read_line(): the next line that arrived, or none. */
    Opt<Text> read_line();
    void tick(Time t) override;
#ifdef MUJARD_HOST_TEST
    void pretend(const Text& line) { pushRx(line); }
    List<Text, 4> sent_;
#endif
    /** What a received character callback hands in (board side only). */
    void gotLine(const char* line);
    /** One app protocol line (LF included) to the phone, in packets every central accepts. */
    void sendApp(const char* s, uint16_t n);
    bool central_ = false;  // the BLE callbacks flip it; public so the callback class can write
private:
    /** The app's lines (0x1F first) go to the app protocol on the main loop, never to 'on ble.line'. */
    void routeApp();
    void packet(const uint8_t* p, uint16_t n);
    /** Queue a received line; a full queue drops it and counts the loss instead of failing inside a BLE callback. */
    void pushRx(const Text& line) {
        if (rx_.n >= 4) { dropped_events++; return; }
        rx_.add(line);
    }
    List<Text, 4> rx_;
    bool started_ = false;
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    BLECharacteristic* tx_ = nullptr;
#endif
};

static Ble ble_rt;

#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
// Nordic UART service shape: the central writes to RX, we notify on TX.
static const char* kBleService = "6e400001-b5a3-f393-e0a9-e50e24dcca9e";
static const char* kBleRx = "6e400002-b5a3-f393-e0a9-e50e24dcca9e";
static const char* kBleTx = "6e400003-b5a3-f393-e0a9-e50e24dcca9e";

class BleServerCb : public BLEServerCallbacks {
    void onConnect(BLEServer*) override { ble_rt.central_ = true; }
    void onDisconnect(BLEServer*) override { ble_rt.central_ = false; }
};
class BleRxCb : public BLECharacteristicCallbacks {
    void onWrite(BLECharacteristic* c) override {
        String v = c->getValue();
        ble_rt.gotLine(v.c_str());
    }
};
static BleServerCb ble_server_cb;
static BleRxCb ble_rx_cb;
#elif defined(ARDUINO_ARCH_MBED) && defined(ARDUINO_ARDUINO_NANO33BLE) && !defined(MUJARD_HOST_TEST)
static BLEService ble_nus("6e400001-b5a3-f393-e0a9-e50e24dcca9e");
static BLECharacteristic ble_rx("6e400002-b5a3-f393-e0a9-e50e24dcca9e", BLEWrite, 64);
static BLECharacteristic ble_tx("6e400003-b5a3-f393-e0a9-e50e24dcca9e", BLENotify, 64);
#elif defined(ARDUINO_ARCH_RP2040) && defined(PICO_CYW43_SUPPORTED) && !defined(MUJARD_HOST_TEST)
// Pico W's CYW43: the core's UART service is the same Nordic shape as the ESP32's (BLEServiceUART).
static BLEServiceUART ble_uart;
#endif

inline void Ble::gotLine(const char* line) {
    pushRx(Text(line));
}

inline void Ble::routeApp() {
#ifdef MUJARD_APP_MIRROR
    for (uint8_t i = 0; i < rx_.n;) {
        if (rx_.items[i].s[0] != 0x1F) { i++; continue; }
        Text t = rx_.items[i];
        rx_.remove(i);
        while (t.n > 0 && (t.s[t.n - 1] == '\n' || t.s[t.n - 1] == '\r')) t.s[--t.n] = 0;
        app_input_line(t.s);
    }
#endif
}

inline void Ble::packet(const uint8_t* p, uint16_t n) {
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    tx_->setValue((uint8_t*)p, n);
    tx_->notify();
#elif defined(ARDUINO_ARCH_MBED) && defined(ARDUINO_ARDUINO_NANO33BLE) && !defined(MUJARD_HOST_TEST)
    ble_tx.writeValue(p, n);
#elif defined(ARDUINO_ARCH_RP2040) && defined(PICO_CYW43_SUPPORTED) && !defined(MUJARD_HOST_TEST)
    ble_uart.write(p, n);
#else
    (void)p; (void)n;
#endif
}

inline void Ble::sendApp(const char* s, uint16_t n) {
    if (!central_) return;
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    if (tx_ == nullptr) return;
#endif
#ifdef MUJARD_HOST_TEST
    Text t;
    for (uint16_t i = 0; i < n && s[i] != '\n' && t.n < MUJARD_TEXT_MAX; i++) t.s[t.n++] = s[i];
    t.s[t.n] = 0;
    if (sent_.n < 4) sent_.add(t);
#else
    // 20 bytes is what every central takes before MTU negotiation; the phone joins packets up to the LF.
    for (uint16_t i = 0; i < n; i += 20) packet((const uint8_t*)s + i, n - i < 20 ? n - i : 20);
#endif
}

#ifdef MUJARD_APP_MIRROR
static void bleAppMirror(const char* s, uint16_t n) { ble_rt.sendApp(s, n); }
#endif

inline void Ble::start(const char* name) {
    if (started_) return;
    started_ = true;
#ifdef MUJARD_APP_MIRROR
    app_mirror_add(bleAppMirror);  // the dashboard's show / notify / control messages go over BLE too
#endif
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    BLEDevice::init(name);
    BLEServer* s = BLEDevice::createServer();
    s->setCallbacks(&ble_server_cb);
    BLEService* svc = s->createService(kBleService);
    BLECharacteristic* rx = svc->createCharacteristic(kBleRx, BLECharacteristic::PROPERTY_WRITE);
    rx->setCallbacks(&ble_rx_cb);
    tx_ = svc->createCharacteristic(kBleTx, BLECharacteristic::PROPERTY_NOTIFY);
    tx_->addDescriptor(new BLE2902());
    svc->start();
    BLEAdvertising* adv = BLEDevice::getAdvertising();
    adv->addServiceUUID(kBleService);
    adv->setScanResponse(true);
    adv->start();
#elif defined(ARDUINO_ARCH_MBED) && defined(ARDUINO_ARDUINO_NANO33BLE) && !defined(MUJARD_HOST_TEST)
    if (!BLE.begin()) return;
    BLELocalDevice().setLocalName(name);
    BLELocalDevice().setDeviceName(name);
    BLE.setAdvertisedService(ble_nus);
    ble_nus.addCharacteristic(ble_rx);
    ble_nus.addCharacteristic(ble_tx);
    BLE.addService(ble_nus);
    BLE.advertise();
#elif defined(ARDUINO_ARCH_RP2040) && defined(PICO_CYW43_SUPPORTED) && !defined(MUJARD_HOST_TEST)
    BLE.begin(name);
    BLE.server()->addService(&ble_uart);
    BLE.startAdvertising();
    ble_uart.setAutoflush(50);  // partial writes flush on their own
#else
    central_ = started_;  // the fake link is "up" once started
#endif
}

inline bool Ble::send(const Text& text) {
    // Like print: a program's text never carries the app marker, so relayed data can't pose as a notify or a control.
    Text line = text;
    for (uint8_t i = 0; i < line.n; i++) if (line.s[i] == 0x1F) line.s[i] = '?';
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    if (!central_ || tx_ == nullptr) return false;
    tx_->setValue((uint8_t*)line.s, line.n);
    tx_->notify();
    return true;
#elif defined(ARDUINO_ARCH_MBED) && defined(ARDUINO_ARDUINO_NANO33BLE) && !defined(MUJARD_HOST_TEST)
    BLEDevice c = BLE.central();
    if (!c.connected()) return false;
    ble_tx.writeValue((const uint8_t*)line.s, line.n);
    return true;
#elif defined(ARDUINO_ARCH_RP2040) && defined(PICO_CYW43_SUPPORTED) && !defined(MUJARD_HOST_TEST)
    if (!central_) return false;
    ble_uart.write((const uint8_t*)line.s, line.n);
    return true;
#else
    if (sent_.n >= 4) return false;
    sent_.add(line);
    return central_;
#endif
}

inline bool Ble::line(Text& out) {
    routeApp();
    if (rx_.n == 0) return false;
    out = rx_.items[0];
    rx_.remove(0);
    return true;
}

inline Opt<Text> Ble::read_line() {
    Text t;
    if (!line(t)) return Opt<Text>();
    return Opt<Text>(t);
}

inline void Ble::tick(Time) {
    routeApp();
#if defined(ARDUINO_ARCH_MBED) && defined(ARDUINO_ARDUINO_NANO33BLE) && !defined(MUJARD_HOST_TEST)
    BLE.poll();
    central_ = BLE.central().connected();
    if (central_ && ble_rx.written()) {
        char buf[65];
        int n = ble_rx.valueLength() > 64 ? 64 : ble_rx.valueLength();
        memcpy(buf, ble_rx.value(), n);
        buf[n] = 0;
        gotLine(buf);
        ble_rx.setValue("", 0);
    }
#elif defined(ARDUINO_ARCH_RP2040) && defined(PICO_CYW43_SUPPORTED) && !defined(MUJARD_HOST_TEST)
    central_ = (bool)ble_uart;
    // The UART is a stream, not one write per line: join it at LF like the phone's packets.
    static char buf[128];
    static uint8_t bn = 0;
    while (ble_uart.available()) {
        char c = (char)ble_uart.read();
        if (bn >= sizeof(buf) - 1) { bn = 0; continue; }  // a runaway line is dropped
        buf[bn++] = c;
        if (c == '\n') {
            buf[bn] = 0;
            gotLine(buf);
            bn = 0;
        }
    }
#endif
}

}  // namespace mujard
