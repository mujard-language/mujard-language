// clock.sync / clock.time / clock.date / at HH:MM (spec Appendix A 22-24): the time of day from the internet (NTP)
// over Wi-Fi, or from an RTC chip (use mujard.rtc) when there is none - clock.source(rtc) picks the chip. The
// compiler turns the time zone name into a POSIX rule ("Africa/Lagos" -> "WAT-1"), so no zone database is needed
// on the board. Until a source gives a time the clock is none, and 'at HH:MM' blocks don't run.
//
// With an RTC in the program (mujard_rtc.h registers rtc_clock_hook), the chip is also the fallback while the
// internet time hasn't arrived yet.
#pragma once
#include <time.h>
#include "mujard_wifi.h"

namespace mujard {

class Clock {
public:
    void sync(const char* posixTz) {
#if defined(ARDUINO_ARCH_ESP32)
        configTzTime(posixTz, "pool.ntp.org", "time.nist.gov");
#elif defined(ARDUINO_ARCH_ESP8266)
        configTime(posixTz, "pool.ntp.org", "time.nist.gov");
#elif defined(ARDUINO_ARCH_RP2040) && !defined(ARDUINO_ARCH_MBED)
        setenv("TZ", posixTz, 1);
        tzset();
        NTP.begin("pool.ntp.org", "time.nist.gov");
#else
        (void)posixTz;  // no internet clock here (the checker wants Wi-Fi for clock.sync)
#endif
    }
    /** clock.source(rtc): read the time from the RTC chip (use mujard.rtc) instead of the internet. */
    void source_rtc() { useRtc_ = true; }
    /** True once the time came from the internet (anything before 2023 means "not yet"). */
    bool isSet() const { return ::time(nullptr) > 1700000000; }
    Opt<TimeOfDay> time() const {
        struct tm tm;
        if (!local(tm)) return Opt<TimeOfDay>();
        TimeOfDay t = {tm.tm_hour, tm.tm_min, tm.tm_sec};
        return Opt<TimeOfDay>(t);
    }
    Opt<Date> date() const {
        struct tm tm;
        if (!local(tm)) return Opt<Date>();
        Date d = {tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, (tm.tm_wday + 6) % 7 + 1};
        return Opt<Date>(d);
    }
    /** at HH:MM: true once a day, in the minute the local time is hour:minute. lastDay is the block's own memory. */
    bool at(int hour, int minute, int32_t& lastDay) const {
        struct tm tm;
        if (!local(tm) || tm.tm_hour != hour || tm.tm_min != minute || tm.tm_yday == lastDay) return false;
        lastDay = tm.tm_yday;
        return true;
    }

private:
    bool useRtc_ = false;
    bool local(struct tm& out) const {
        // the chip when it's the source, or while there's no internet time yet; then the internet time
        if (rtc_clock_hook && (useRtc_ || !isSet())) {
            if (rtc_clock_hook(out)) return true;
            if (useRtc_) return false;
        }
        if (!isSet()) return false;
        time_t now = ::time(nullptr);
        localtime_r(&now, &out);
        return true;
    }
};

static Clock clock_rt;

}  // namespace mujard
