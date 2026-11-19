// use mujard.ethernet - the board's wired Ethernet link (the classic ESP32's EMAC + external PHY, spec 12):
// ethernet.start() brings it up, connected/ip read it, 'on ethernet.joined'/'lost' fire on link changes, and
// it feeds net_up() so http/mqtt/clock work over the wire too. PC test builds fake it with ethernet.pretend.
#pragma once
#include "mujard_wifi.h"
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
#include <ETH.h>
#endif

namespace mujard {

/** A wired Ethernet link: power the PHY, wait for the wire, take an address by DHCP. */
class Eth : public Ticked {
public:
    /** ethernet.start(): power up the EMAC/PHY and start DHCP. */
    void start() {
        if (started_) return;
        started_ = true;
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
        ETH.begin();
#endif
    }
    bool connected() const {
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
        return ETH.linkUp();
#else
        return fake_;
#endif
    }
    /** The address DHCP handed us, as text ("none" before the link is up). */
    Text ip() const {
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
        return Text(ETH.localIP().toString().c_str());
#else
        return connected() ? Text("192.168.1.50") : Text("none");
#endif
    }
    /** on ethernet.joined / on ethernet.lost: true in the pass the wire came up / went down. */
    bool joined() const { return joined_; }
    bool lost() const { return lost_; }
    void tick(Time) override {
        joined_ = lost_ = false;
        if (!started_) return;
        bool up = connected();
        if (up != up_) {
            up_ = up;
            if (up) joined_ = true;
            else lost_ = true;
        }
        ethernet_link_up_ = up_;
    }
#ifdef MUJARD_HOST_TEST
    /** A test's ethernet.pretend(up): the fake link comes up or drops on the next scheduler turn. */
    void pretend(bool up) { fake_ = up; }
#endif
private:
    bool started_ = false, up_ = false, joined_ = false, lost_ = false;
#ifdef MUJARD_HOST_TEST
    bool fake_ = false;
#endif
};

static Eth eth_rt;

}  // namespace mujard
