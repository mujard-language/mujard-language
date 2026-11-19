// use mujard.uart - extra serial ports for GPS, Bluetooth and GSM modules: gps = Uart(rx: pin 4, tx: pin 3, speed: 9600 baud).
// A hardware port where the board has a spare one (ESP32, Pico, Mega, Leonardo/Micro on 0/1), a software port otherwise.
//
// gps.read_line() waits inside a task for a line; wantLines() (the generated setup calls it) keeps each incoming
// line in a small queue until it is taken, instead of only being there for the pass it arrived on.
#pragma once
#include "mujard_rt.h"
#ifdef MUJARD_HOST_TEST
#include <string>
#else
#if defined(ARDUINO_ARCH_AVR) || defined(ARDUINO_ARCH_ESP8266)
#include <SoftwareSerial.h>
#endif
#if defined(ARDUINO_ARCH_RP2040) && !defined(ARDUINO_ARCH_MBED)
#include <SerialPIO.h>
#endif
#endif

namespace mujard {

class Uart : public Ticked {
public:
    Uart(uint8_t rx, uint8_t tx, uint32_t baud) : rx_(rx), tx_(tx), baud_(baud ? baud : 9600) {}
    void init() override {
#ifdef MUJARD_HOST_TEST
        in_.clear();
#else
        port_ = open();
        if (!port_) fail("No serial port is left for this Uart on this board", 0);
#endif
    }
    /** Collects incoming characters; a finished line is there for one pass of the program (like serial.line),
     *  or queues up for read_line() when wantLines() ran. */
    void tick(Time t) override {
        (void)t;
        ready_ = false;
        int c;
        while ((c = next()) >= 0) {
            if (c == '\r') continue;
            if (c == '\n') {
                if (wantQ_) {
                    if (qCount_ < LINE_Q) {
                        q_[(uint8_t)((qHead_ + qCount_) % LINE_Q)] = buf_;
                        qCount_++;
                    } else {
                        dropped_events++;
                    }
                } else {
                    done_ = buf_;
                    ready_ = true;
                }
                buf_ = Text();
                if (!wantQ_) return;
                continue;
            }
            char one[2] = {(char)c, 0};
            buf_.add(one);
        }
    }
    /** Sends the text as it is (add "\n" for a line). A halted program sends nothing: a modem could be listening. */
    void write(const Text& t) {
        if (halted()) return;
#ifdef MUJARD_HOST_TEST
        sent_ += t.s;
#else
        if (port_) port_->print(t.s);
#endif
    }
    void write(const char* s) {
        if (halted()) return;
#ifdef MUJARD_HOST_TEST
        sent_ += s;
#else
        if (port_) port_->print(s);
#endif
    }
    template <int N>
    void write(const TextN<N>& t) { write(t.s); }
    /** on x.line: true in the pass a line arrived, with the line in out. */
    bool line(Text& out) {
        if (wantQ_) {  // the queue feeds the event too: one line per pass, none lost
            if (qCount_ == 0) return false;
            out = q_[qHead_];
            qHead_ = (uint8_t)((qHead_ + 1) % LINE_Q);
            qCount_--;
            return true;
        }
        if (!ready_) return false;
        out = done_;
        return true;
    }

    // read_line() in a task (spec Appendix A 76): the generated setup() calls wantLines() once.
    void wantLines() { wantQ_ = true; }
    bool has_line() const { return qCount_ > 0; }
    Opt<Text> take_line() {
        Opt<Text> out;
        if (qCount_ == 0) return out;
        out.has = true;
        out.v = q_[qHead_];
        qHead_ = (uint8_t)((qHead_ + 1) % LINE_Q);
        qCount_--;
        return out;
    }

#ifdef MUJARD_HOST_TEST
    /** test: pretend a line came in on the wire. */
    void pretend(const Text& t) {
        in_ += t.s;
        in_ += '\n';
    }
#endif

private:
    static const uint8_t LINE_Q = 4;

#ifdef MUJARD_HOST_TEST
    int next() {
        if (in_.empty()) return -1;
        int c = (uint8_t)in_[0];
        in_.erase(0, 1);
        return c;
    }
    std::string in_, sent_;
#else
    int next() { return port_ && port_->available() > 0 ? port_->read() : -1; }
    Stream* open() {
#if defined(ARDUINO_ARCH_ESP32)
        // UART0 is the USB serial port; the others take any pins.
        static uint8_t next = 1;
        if (next >= SOC_UART_NUM) return nullptr;
        HardwareSerial* s = new HardwareSerial(next++);
        s->begin(baud_, SERIAL_8N1, rx_, tx_);
        return s;
#elif defined(ARDUINO_ARCH_RP2040) && !defined(ARDUINO_ARCH_MBED)
        // Serial1 / Serial2 when the pins belong to UART0 / UART1, else a PIO serial port on any pins.
        static bool used1 = false, used2 = false;
        if (!used1 && Serial1.setRX(rx_) && Serial1.setTX(tx_)) {
            used1 = true;
            Serial1.begin(baud_);
            return &Serial1;
        }
        if (!used2 && Serial2.setRX(rx_) && Serial2.setTX(tx_)) {
            used2 = true;
            Serial2.begin(baud_);
            return &Serial2;
        }
        SerialPIO* p = new SerialPIO(tx_, rx_);
        p->begin(baud_);
        return p;
#elif defined(ARDUINO_ARCH_AVR)
#if defined(__AVR_ATmega32U4__)
        if (rx_ == 0 && tx_ == 1) {  // Leonardo / Micro: Serial is USB, Serial1 is pins 0 and 1
            Serial1.begin(baud_);
            return &Serial1;
        }
#elif defined(HAVE_HWSERIAL1)
        if (rx_ == 19 && tx_ == 18) {  // Mega
            Serial1.begin(baud_);
            return &Serial1;
        }
#endif
#if defined(HAVE_HWSERIAL2)
        if (rx_ == 17 && tx_ == 16) {
            Serial2.begin(baud_);
            return &Serial2;
        }
#endif
#if defined(HAVE_HWSERIAL3)
        if (rx_ == 15 && tx_ == 14) {
            Serial3.begin(baud_);
            return &Serial3;
        }
#endif
        SoftwareSerial* s = new SoftwareSerial(rx_, tx_);
        s->begin(baud_);
        return s;
#elif defined(ARDUINO_ARCH_ESP8266)
        SoftwareSerial* s = new SoftwareSerial(rx_, tx_);
        s->begin(baud_);
        return s;
#else
        return nullptr;
#endif
    }
    Stream* port_ = nullptr;
#endif

    uint8_t rx_, tx_;
    uint32_t baud_;
    Text buf_, done_;
    Text q_[LINE_Q];
    uint8_t qHead_ = 0, qCount_ = 0;
    bool ready_ = false, wantQ_ = false;
};

}  // namespace mujard
