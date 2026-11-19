// A bounded queue for handing values between blocks - including across cores, where a plain global races (W620).
//
// q = Queue(text, 8) -> mujard::Queue<Text, 8>. push is false and counts a dropped event when full; pop gives none
// when empty; available/full/n report the state. Nothing grows: the storage is N slots of the declared element.
//
// Backends, honest per platform:
//   ESP32-family:  a FreeRTOS queue - xQueueSend/xQueueReceive are already safe between cores (and from ISRs,
//                  though fast handlers can't call push anyway, C314).
//   RP2040/RP2350: the SDK's mutex_t around a fixed ring - multicore_launch_core1 cores share the same RAM and
//                  pico_sync's mutex is built for exactly this.
//   host test:     the same ring without a lock - one fake scheduler can't exercise real contention, and it
//                  doesn't pretend to.
#pragma once
#include "mujard_rt.h"
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    #include "freertos/FreeRTOS.h"
    #include "freertos/queue.h"
#elif defined(ARDUINO_ARCH_RP2040) && !defined(MUJARD_HOST_TEST)
    #include "pico/mutex.h"
    #define MUJARD_QUEUE_LOCKED 1
#endif

namespace mujard {

template <class T, int N>
class Queue {
public:
    Queue() {
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
        // The installed ESP32 cores set CONFIG_FREERTOS_SUPPORT_STATIC_ALLOCATION=y, so the queue's storage is
        // this object's own array - no heap, and the RAM the checker estimates is what the board really uses.
        h_ = xQueueCreateStatic(N, sizeof(T), storage_, &ctl_);
#elif defined(MUJARD_QUEUE_LOCKED)
        mutex_init(&m_);
#endif
    }

    /** push: copy a value in. False when the queue is full - the item is dropped and counted (board.dropped_events). */
    bool push(const T& x) {
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
        if (h_ == nullptr || xQueueSend(h_, &x, 0) != pdTRUE) {
            dropped_events++;
            return false;
        }
        return true;
#elif defined(MUJARD_QUEUE_LOCKED)
        mutex_enter_blocking(&m_);
        if (cnt_ >= N) {
            mutex_exit(&m_);
            dropped_events++;
            return false;
        }
        items_[(head_ + cnt_) % N] = x;
        cnt_++;
        mutex_exit(&m_);
        return true;
#else
        if (cnt_ >= N) {
            dropped_events++;
            return false;
        }
        items_[(head_ + cnt_) % N] = x;
        cnt_++;
        return true;
#endif
    }

    /** pop: the oldest value, or none when the queue is empty. */
    Opt<T> pop() {
        Opt<T> o;
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
        if (h_ != nullptr && xQueueReceive(h_, &o.v, 0) == pdTRUE) o.has = true;
#elif defined(MUJARD_QUEUE_LOCKED)
        mutex_enter_blocking(&m_);
        if (cnt_ > 0) {
            o.v = items_[head_];
            head_ = (head_ + 1) % N;
            cnt_--;
            o.has = true;
        }
        mutex_exit(&m_);
#else
        if (cnt_ > 0) {
            o.v = items_[head_];
            head_ = (head_ + 1) % N;
            cnt_--;
            o.has = true;
        }
#endif
        return o;
    }

    bool available() { return n() > 0; }
    bool full() { return n() >= N; }

    int32_t n() {
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
        return h_ ? (int32_t)uxQueueMessagesWaiting(h_) : 0;
#elif defined(MUJARD_QUEUE_LOCKED)
        mutex_enter_blocking(&m_);
        int32_t r = cnt_;
        mutex_exit(&m_);
        return r;
#else
        return cnt_;
#endif
    }

private:
#if defined(ARDUINO_ARCH_ESP32) && !defined(MUJARD_HOST_TEST)
    uint8_t storage_[N * sizeof(T)]; // static queue buffer (configSUPPORT_STATIC_ALLOCATION=y on ESP32 cores)
    StaticQueue_t ctl_;
    QueueHandle_t h_ = nullptr;
#else
    T items_[N];
    int head_ = 0, cnt_ = 0;
#if defined(MUJARD_QUEUE_LOCKED)
    mutex_t m_;
#endif
#endif
};

}  // namespace mujard
