// store.save(name, value) / store.load(name, default:) - named values saved when the program chooses (spec Appendix A
// 69-70). Included only in programs that use the store: its RAM copy costs STORE_SLOTS * 24 bytes.
//
// The last STORE_AREA bytes of the EEPROM / flash hold STORE_SLOTS slots: [0xA5][name hash:2][type][size][value:19].
// Saves go to the RAM copy at once and to storage at most once per storage_interval() (60 s, 5 s on the ESP32), and
// before sleep, restart and halt - repeated saves can't wear the storage out. Text keeps up to 18 characters.
#pragma once
#include "mujard_rt.h"

namespace mujard {

class Store : public Ticked {
public:
    static const uint8_t SLOTS = STORE_AREA / 24;
    static const uint8_t VALUE = 19;

    void init() override {
        base_ = store_area_start();
        if (base_ >= 0)
            for (uint16_t i = 0; i < SLOTS * 24; i++) ((uint8_t*)slots_)[i] = storage_read((uint16_t)(base_ + i));
        store_flush_hook = flushNow;
        self() = this;
    }
    void tick(Time t) override {
        if (dirty_ && t >= next_) {
            next_ = t + storage_interval();
            flush();
        }
    }

    void save(const char* name, int32_t v) { put(name, 1, &v, sizeof v); }
    void save(const char* name, float v) { put(name, 2, &v, sizeof v); }
    void save(const char* name, bool v) { put(name, 3, &v, sizeof v); }
    void save(const char* name, int64_t v) { put(name, 4, &v, sizeof v); }  // time, frequency
    void save(const char* name, const Text& v) {
        uint8_t n = v.n < VALUE - 1 ? v.n : VALUE - 1;
        char b[VALUE] = {0};
        memcpy(b, v.s, n);
        put(name, 5, b, (uint8_t)(n + 1));
    }
    template <int N>
    void save(const char* name, const TextN<N>& v) {
        save(name, Text(v.s));
    }
    template <class E>
    void save(const char* name, E v) {  // enums
        put(name, 6, &v, sizeof v);
    }

    int32_t load(const char* name, int32_t d) {
        get(name, 1, &d, sizeof d);
        return d;
    }
    float load(const char* name, float d) {
        get(name, 2, &d, sizeof d);
        return d;
    }
    bool load(const char* name, bool d) {
        get(name, 3, &d, sizeof d);
        return d;
    }
    int64_t load(const char* name, int64_t d) {
        get(name, 4, &d, sizeof d);
        return d;
    }
    Text load(const char* name, const Text& d) {
        char b[VALUE] = {0};
        if (!get(name, 5, b, VALUE)) return d;
        b[VALUE - 1] = 0;
        return Text(b);
    }
    template <int N>
    TextN<N> load(const char* name, const TextN<N>& d) {
        Text v = load(name, Text(d.s));
        return TextN<N>(v.s);
    }
    template <class E>
    E load(const char* name, E d) {
        get(name, 6, &d, sizeof d);
        return d;
    }

    // store.load(name) without a default gives T? - the type comes from where the value goes
    // ('var x: int? = store.load("k")'), and the generator passes it as a null pointer of that type.
    Opt<int32_t> load(const char* name, int32_t*) {
        Opt<int32_t> o;
        o.has = get(name, 1, &o.v, sizeof o.v);
        return o;
    }
    Opt<float> load(const char* name, float*) {
        Opt<float> o;
        o.has = get(name, 2, &o.v, sizeof o.v);
        return o;
    }
    Opt<bool> load(const char* name, bool*) {
        Opt<bool> o;
        o.has = get(name, 3, &o.v, sizeof o.v);
        return o;
    }
    Opt<int64_t> load(const char* name, int64_t*) {
        Opt<int64_t> o;
        o.has = get(name, 4, &o.v, sizeof o.v);
        return o;
    }
    Opt<Text> load(const char* name, Text*) {
        Opt<Text> o;
        char b[VALUE] = {0};
        if (get(name, 5, b, VALUE)) {
            b[VALUE - 1] = 0;
            o.v = Text(b);
            o.has = true;
        }
        return o;
    }
    template <int N>
    Opt<TextN<N>> load(const char* name, TextN<N>*) {
        Opt<TextN<N>> o;
        Opt<Text> v = load(name, (Text*)nullptr);
        o.has = v.has;
        if (v.has) o.v = TextN<N>(v.v.s);
        return o;
    }
    template <class E>
    Opt<E> load(const char* name, E*) {  // enums
        Opt<E> o;
        o.has = get(name, 6, &o.v, sizeof o.v);
        return o;
    }

    // Durability status for store.pending / store.failed / store.committed / store.last_commit.
    // pending: writes are accepted but not yet committed. failed: commits that didn't land (shared with
    // saved vars - same storage_commit). committed/last_commit: the most recent commit's result - a save
    // isn't durable until a commit landed, and store.save is async.
    bool pending() const { return dirty_; }
    int32_t failed() const { return (int32_t)storage_commits_failed; }
    bool committed() const { return storage_commit_ok; }
    bool last_commit() const { return storage_commit_ok; }

private:
    struct Slot {
        uint8_t magic;
        uint8_t hash[2];
        uint8_t type;
        uint8_t size;
        uint8_t value[VALUE];
    };

    static Store*& self() {
        static Store* s = nullptr;
        return s;
    }
    static void flushNow() {
        if (self()) self()->flush();
    }
    static uint16_t hashOf(const char* name) {
        uint16_t h = 0x811C;  // FNV-1a, folded to 16 bits
        for (; *name; name++) h = (uint16_t)((h ^ (uint8_t)*name) * 0x0193);
        return h;
    }
    Slot* find(uint16_t h, bool create) {
        Slot* empty = nullptr;
        for (uint8_t i = 0; i < SLOTS; i++) {
            Slot& s = slots_[i];
            if (s.magic == 0xA5 && s.hash[0] == (uint8_t)h && s.hash[1] == (uint8_t)(h >> 8)) return &s;
            if (s.magic != 0xA5 && !empty) empty = &s;
        }
        if (create && !empty && !warnedFull_) {
            warnedFull_ = true;
            Out o("WARN");
            o.put("The store is full: it keeps ").put((int32_t)SLOTS).put(" names");
            o.end();
        }
        return create ? empty : nullptr;
    }
    void put(const char* name, uint8_t type, const void* v, uint8_t size) {
        uint16_t h = hashOf(name);
        Slot* s = find(h, true);
        if (!s || size > VALUE) return;
        Slot n = Slot();
        n.magic = 0xA5;
        n.hash[0] = (uint8_t)h;
        n.hash[1] = (uint8_t)(h >> 8);
        n.type = type;
        n.size = size;
        memcpy(n.value, v, size);
        if (memcmp(s, &n, sizeof n) == 0) return;
        *s = n;
        dirty_ = true;
    }
    bool get(const char* name, uint8_t type, void* out, uint8_t size) {
        Slot* s = find(hashOf(name), false);
        if (!s || s->type != type || s->size > size || s->size > VALUE) return false;
        memcpy(out, s->value, s->size);
        return true;
    }
    void flush() {
        if (!dirty_ || base_ < 0) return;
        bool any = false;
        for (uint16_t i = 0; i < SLOTS * 24; i++) any |= storage_write((uint16_t)(base_ + i), ((uint8_t*)slots_)[i]);
        if (any && !storage_commit()) return;  // a lockout timed out: stay dirty, a later flush retries
        dirty_ = false;
    }

    Slot slots_[SLOTS];
    int base_ = -1;
    bool dirty_ = false, warnedFull_ = false;
    Time next_ = 0;
};

static Store store_db;

}  // namespace mujard
