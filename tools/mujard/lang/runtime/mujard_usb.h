// use mujard.usb - the board acts as a USB keyboard+mouse to whatever it's plugged into (spec 12): usb.type
// types text, usb.key taps a key, usb.mouse_move and usb.click drive the pointer. The standard Arduino HID API
// on 32U4, RP2040/RP2350, SAMD21 and RA4M1; PC test builds record what would have been sent.
#pragma once
#include "mujard_rt.h"
#if !defined(MUJARD_HOST_TEST) && \
    (defined(ARDUINO_AVR_LEONARDO) || defined(ARDUINO_AVR_MICRO) || defined(ARDUINO_AVR_YUN) || \
     defined(ARDUINO_AVR_LEONARDO_ETH) || defined(ARDUINO_ARCH_RP2040) || defined(ARDUINO_ARCH_SAMD) || \
     defined(ARDUINO_ARCH_RENESAS))
#include <Keyboard.h>
#include <Mouse.h>
#define MUJARD_USB_HW 1
#endif

namespace mujard {

/** The board's USB-device role: a keyboard and mouse the host sees. */
class Usb : public Ticked {
public:
    void init() override {
#ifdef MUJARD_USB_HW
        Keyboard.begin();
        Mouse.begin();
#endif
        started_ = true;
    }
    /** usb.type("text"): type the text as keystrokes; false when the HID device isn't up. */
    bool type(const Text& s) {
#ifdef MUJARD_USB_HW
        if (!connected()) return false;
        Keyboard.print(s.s);
        return true;
#else
        if (sent_.n < 8) sent_.add(s); // the record is host-test bookkeeping; keep sending past its cap
        return connected();
#endif
    }
    /** usb.key("a"): tap one key. */
    bool key(const Text& k) {
#ifdef MUJARD_USB_HW
        if (!connected() || k.n < 1) return false;
        Keyboard.write((uint8_t)k.s[0]);
        return true;
#else
        if (sent_.n < 8) sent_.add(k); // the record is host-test bookkeeping; keep sending past its cap
        return connected();
#endif
    }
    bool mouse_move(int32_t x, int32_t y) {
#ifdef MUJARD_USB_HW
        if (!connected()) return false;
        Mouse.move((int8_t)x, (int8_t)y);
        return true;
#else
        moved_ = true;
        return connected();
#endif
    }
    bool click() {
#ifdef MUJARD_USB_HW
        if (!connected()) return false;
        Mouse.click();
        return true;
#else
        clicked_ = true;
        return connected();
#endif
    }
    /** Whether the HID device is up (started on a board; pretend(true) in a test). */
    bool connected() const {
#ifdef MUJARD_USB_HW
        return started_;
#else
        return central_;
#endif
    }
#ifdef MUJARD_HOST_TEST
    void pretend(bool up) { central_ = up; }
    List<Text, 8> sent_;
    bool moved_ = false, clicked_ = false;
#endif
private:
    bool started_ = false;
#ifdef MUJARD_HOST_TEST
    bool central_ = false;
#endif
};

static Usb usb_rt;

}  // namespace mujard
