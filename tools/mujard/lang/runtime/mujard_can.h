// use mujard.can - a CAN bus through the ESP32 family's TWAI controller (spec 12). On other boards the generator
// refuses with C320; PC test builds keep a fake bus, like i2c.pretend.
#pragma once
#include "mujard_rt.h"
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
#include <driver/twai.h>
#endif

namespace mujard {

/** One CAN frame: a message id, up to 8 data bytes and a remote-request flag. */
struct CanMessage {
    int32_t id = 0;
    List<int32_t, 8> data;
    bool rtr = false;
};

/** The bus: a Ticked that polls the driver each scheduler turn so 'on bus.message' and bus.read() see what came. */
class CanBus : public Ticked {
public:
    CanBus(uint8_t tx, uint8_t rx, int32_t baud) : tx_(tx), rx_(rx), baud_(baud) {}
    void init() override;
    void tick(Time t) override { poll(); }
    /** Queue a frame to send; false when the driver queue is full. */
    bool send(int32_t id, const List<int32_t, 8>& data);
    /** The next received frame, or none. */
    Opt<CanMessage> read() {
        poll();
        Opt<CanMessage> m = next_;
        next_.has = false;
        return m;
    }
    /** One frame for an 'on bus.message' handler - like serial.line. */
    bool message(CanMessage& m) {
        if (!next_.has) return false;
        m = next_.v;
        next_.has = false;
        return true;
    }
#ifdef MUJARD_HOST_TEST
    /** A test's bus.pretend: the next frame the bus "receives" (spec 12). A full queue drops the frame and counts
     *  the loss, like the TWAI driver's rx queue filling up on the board. */
    void pretend(int32_t id, const List<int32_t, 8>& data) {
        if (rxq_.n >= 8) { dropped_events++; return; }
        CanMessage m;
        m.id = id;
        m.data = data;
        rxq_.add(m);
    }
    List<CanMessage, 8> sent_;
#endif

private:
    void poll();
    uint8_t tx_, rx_;
    int32_t baud_;
    Opt<CanMessage> next_;
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    uint32_t missed_ = 0; // the TWAI driver's missed-message count we last folded into dropped_events
#endif
#ifdef MUJARD_HOST_TEST
    List<CanMessage, 8> rxq_;
#endif
};

#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
// Standard CAN rates get the driver's exact timing macros; anything else lands on the nearest one.
static twai_timing_config_t canTiming(int32_t baud) {
    if (baud <= 37500) return TWAI_TIMING_CONFIG_25KBITS();
    if (baud <= 75000) return TWAI_TIMING_CONFIG_50KBITS();
    if (baud <= 112500) return TWAI_TIMING_CONFIG_100KBITS();
    if (baud <= 187500) return TWAI_TIMING_CONFIG_125KBITS();
    if (baud <= 375000) return TWAI_TIMING_CONFIG_250KBITS();
    if (baud <= 650000) return TWAI_TIMING_CONFIG_500KBITS();
    if (baud <= 900000) return TWAI_TIMING_CONFIG_800KBITS();
    return TWAI_TIMING_CONFIG_1MBITS();
}
#endif

inline void CanBus::init() {
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    twai_general_config_t g = TWAI_GENERAL_CONFIG_DEFAULT((gpio_num_t)tx_, (gpio_num_t)rx_, TWAI_MODE_NORMAL);
    twai_timing_config_t t = canTiming(baud_);
    twai_filter_config_t f = TWAI_FILTER_CONFIG_ACCEPT_ALL();
    if (twai_driver_install(&g, &t, &f) == ESP_OK) (void)twai_start();
#endif
}

inline bool CanBus::send(int32_t id, const List<int32_t, 8>& data) {
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    twai_message_t m;
    m.flags = TWAI_MSG_FLAG_NONE;
    m.identifier = (uint32_t)id;
    m.data_length_code = data.n > 8 ? 8 : (uint8_t)data.n;
    for (int i = 0; i < m.data_length_code; i++) m.data[i] = (uint8_t)data.items[i];
    return twai_transmit(&m, 0) == ESP_OK;
#elif defined(MUJARD_HOST_TEST)
    if (sent_.n >= 8) return false;
    CanMessage m;
    m.id = id;
    m.data = data;
    sent_.add(m);
    return true;
#else
    (void)id;
    (void)data;
    return false;
#endif
}

inline void CanBus::poll() {
    if (next_.has) return;
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    // The TWAI driver counts frames it had to drop before we could read them (its own rx queue filled): fold the
    // delta into dropped_events so board.dropped_events sees driver-level loss, not only our own queue's.
    twai_status_info_t s;
    if (twai_get_status_info(&s) == ESP_OK && s.msgs_missed > missed_) {
        dropped_events += s.msgs_missed - missed_;
        missed_ = s.msgs_missed;
    }
    twai_message_t m;
    if (twai_receive(&m, 0) != ESP_OK) return;
    CanMessage cm;
    cm.id = (int32_t)m.identifier;
    cm.rtr = (m.flags & TWAI_MSG_FLAG_RTR) != 0;
    for (int i = 0; i < m.data_length_code && i < 8; i++) cm.data.add(m.data[i]);
    next_ = Opt<CanMessage>(cm);
#elif defined(MUJARD_HOST_TEST)
    if (rxq_.n > 0) {
        next_ = Opt<CanMessage>(rxq_.items[0]);
        rxq_.remove(0);
    }
#endif
}

}  // namespace mujard
