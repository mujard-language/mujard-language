// use mujard.mqtt - mqtt.connect / publish / subscribe / on mqtt.message (spec Appendix A 80-81). A small MQTT 3.1.1
// client (QoS 0) of our own, so a build needs no extra library. Nothing waits: connect() only records where to connect,
// the link comes up in the background once Wi-Fi is up, reconnects with a growing pause (5 s .. 60 s), sends a ping
// every 30 s, re-subscribes after every reconnect, and keeps up to 4 publishes made while it was down.
#pragma once
#include "mujard_wifi.h"

namespace mujard {

struct MqttMessage {
    Text topic, text;
};

class Mqtt : public Ticked {
public:
    static const uint8_t MAX_SUBS = 4, MAX_QUEUE = 4;
    static const uint16_t MAX_PACKET = 256;  // incoming packets above this are skipped

    void connect(const char* host, const char* user = nullptr, const char* password = nullptr, uint16_t port = 1883) {
        host_ = Text(host);
        copy(user_, sizeof user_, user);
        copy(pass_, sizeof pass_, password);
        port_ = port;
        wanted_ = true;
        nextTry_ = 0;
        backoff_ = 5000000;
    }
    bool connected() { return up_ && net_.connected(); }
    void publish(const Text& topic, const Text& payload) {
        if (halted()) return;  // a halted program sends nothing
        if (connected()) {
            sendPublish(topic, payload);
        } else if (queued_ < MAX_QUEUE) {
            queue_[queued_].topic = topic;
            queue_[queued_].text = payload;
            queued_++;
        }
    }
    void subscribe(const Text& topic) {
        for (uint8_t i = 0; i < subCount_; i++)
            if (subs_[i].eq(topic)) return;
        if (subCount_ < MAX_SUBS) subs_[subCount_++] = topic;
        if (connected()) sendSubscribe(topic);
    }
    /** on mqtt.message as m: true in the pass a message arrived. */
    bool message(MqttMessage& out) {
        if (!ready_) return false;
        out = msg_;
        return true;
    }

    void tick(Time t) override {
        ready_ = false;
        if (!wanted_) return;
        if (!net_up()) {
            if (up_ || net_.connected()) drop();
            return;
        }
        if (!net_.connected()) {
            if (up_) drop();
            if (t < nextTry_) return;
            nextTry_ = t + backoff_;
            if (backoff_ < 60000000) backoff_ *= 2;
            if (!net_.connect(host_.s, port_)) return;
            sendConnect();
            connackBy_ = t + 5000000;
            return;
        }
        while (net_.available() > 0 && !ready_) feed((uint8_t)net_.read());
        if (!up_ && t > connackBy_) {
            net_.stop();  // no CONNACK in time: try again later
            return;
        }
        if (up_ && t - lastSent_ > 30000000) {
            const uint8_t ping[2] = {0xC0, 0x00};
            send(ping, 2);
        }
    }

private:
    static void copy(char* to, size_t cap, const char* from) {
        size_t i = 0;
        for (; from && from[i] && i + 1 < cap; i++) to[i] = from[i];
        to[i] = 0;
    }
    void drop() {
        net_.stop();
        up_ = false;
        phase_ = 0;
    }
    void send(const uint8_t* b, size_t n) {
        net_.write(b, n);
        lastSent_ = now();
    }
    // The remaining length (MQTT's 1-4 byte varint) and a packet's fixed header.
    void header(uint8_t type, uint32_t len) {
        uint8_t b[5];
        size_t n = 0;
        b[n++] = type;
        do {
            uint8_t d = len % 128;
            len /= 128;
            b[n++] = len ? (uint8_t)(d | 0x80) : d;
        } while (len);
        send(b, n);
    }
    void str(const char* s, uint16_t n) {
        uint8_t l[2] = {(uint8_t)(n >> 8), (uint8_t)n};
        send(l, 2);
        send((const uint8_t*)s, n);
    }
    void sendConnect() {
        Text id("mujard-");
        id.add(board_id());
        uint16_t nu = (uint16_t)strlen(user_), np = (uint16_t)strlen(pass_);
        uint8_t flags = 0x02;  // clean session
        uint32_t len = 10 + 2 + id.n;
        if (nu) flags |= 0x80, len += 2 + nu;
        if (np) flags |= 0x40, len += 2 + np;
        header(0x10, len);
        const uint8_t var[10] = {0x00, 0x04, 'M', 'Q', 'T', 'T', 0x04, flags, 0x00, 60};  // MQTT 3.1.1, keep-alive 60 s
        send(var, 10);
        str(id.s, id.n);
        if (nu) str(user_, nu);
        if (np) str(pass_, np);
    }
    void sendPublish(const Text& topic, const Text& payload) {
        header(0x30, 2 + topic.n + payload.n);
        str(topic.s, topic.n);
        send((const uint8_t*)payload.s, payload.n);
    }
    void sendSubscribe(const Text& topic) {
        header(0x82, 2 + 2 + topic.n + 1);
        if (++packetId_ == 0) packetId_ = 1;  // 0 isn't a valid packet id
        const uint8_t pid[2] = {(uint8_t)(packetId_ >> 8), (uint8_t)packetId_};
        send(pid, 2);
        str(topic.s, topic.n);
        const uint8_t qos = 0;
        send(&qos, 1);
    }
    // Incoming bytes: fixed header, remaining length, body.
    void feed(uint8_t c) {
        switch (phase_) {
        case 0:
            type_ = c;
            len_ = 0;
            shift_ = 0;
            phase_ = 1;
            return;
        case 1:
            len_ |= (uint32_t)(c & 0x7F) << shift_;
            shift_ += 7;
            if (c & 0x80) {
                if (shift_ > 21) phase_ = 0;  // malformed: more than 4 length bytes
                return;
            }
            got_ = 0;
            phase_ = 2;
            if (len_ == 0) packet();
            return;
        default:
            if (got_ < MAX_PACKET) buf_[got_] = c;
            got_++;
            if (got_ >= len_) packet();
        }
    }
    void packet() {
        phase_ = 0;
        uint8_t kind = type_ & 0xF0;
        if (len_ > MAX_PACKET) return;  // too big for this board: skipped
        if (kind == 0x20 && len_ >= 2) {  // CONNACK
            if (buf_[1] != 0) {
                net_.stop();  // refused (bad user/password, ...): try again later
                return;
            }
            up_ = true;
            backoff_ = 5000000;
            for (uint8_t i = 0; i < subCount_; i++) sendSubscribe(subs_[i]);
            for (uint8_t i = 0; i < queued_; i++) sendPublish(queue_[i].topic, queue_[i].text);
            queued_ = 0;
        } else if (kind == 0x30 && len_ >= 2) {  // PUBLISH
            uint16_t tl = (uint16_t)((buf_[0] << 8) | buf_[1]);
            uint16_t at = (uint16_t)(2 + tl + ((type_ & 0x06) ? 2 : 0));  // QoS 1/2 carry a packet id
            if (at > len_) return;
            msg_.topic = Text();
            msg_.text = Text();
            for (uint16_t i = 0; i < tl; i++) addChar(msg_.topic, buf_[2 + i]);
            for (uint32_t i = at; i < len_; i++) addChar(msg_.text, buf_[i]);
            ready_ = true;
        }
    }
    static void addChar(Text& t, uint8_t c) {
        char one[2] = {(char)c, 0};
        if (c) t.add(one);
    }

    WiFiClient net_;
    Text host_;
    char user_[65] = "", pass_[161] = "";
    uint16_t port_ = 1883, packetId_ = 0;
    bool wanted_ = false, up_ = false, ready_ = false;
    Time nextTry_ = 0, backoff_ = 5000000, connackBy_ = 0, lastSent_ = 0;
    Text subs_[MAX_SUBS];
    uint8_t subCount_ = 0, queued_ = 0;
    MqttMessage queue_[MAX_QUEUE], msg_;
    uint8_t phase_ = 0, type_ = 0, shift_ = 0;
    uint32_t len_ = 0, got_ = 0;
    uint8_t buf_[MAX_PACKET];
};

static Mqtt mqtt_rt;

}  // namespace mujard
