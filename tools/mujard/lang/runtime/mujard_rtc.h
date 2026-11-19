// use mujard.rtc - a battery-backed clock chip (DS3231) on the I2C pins (spec Appendix A 24): it keeps the time
// when there's no Wi-Fi, and clock reads it once clock.source(rtc) is called - or by itself, while there's no
// internet time. rtc.set(year, month, day, hour, minute, second:) writes the chip once.
//
// The clock finds the chip through rtc_clock_hook (mujard_rt.h), so the two headers don't have to know each other.
#pragma once
#include "mujard_rt.h"
#ifndef MUJARD_HOST_TEST
#include <Wire.h>
#endif

namespace mujard {

class Rtc : public Ticked {
public:
    void init() override {
#ifndef MUJARD_HOST_TEST
        Wire.begin();  // the chip is at 0x68; the I2C pins are the board's own (W609 warns about them)
#endif
        rtc_clock_hook = &read;
    }

    /** Write the time to the chip once (usually at build time or after the internet gave it). */
    void set(int32_t year, int32_t month, int32_t day, int32_t hour, int32_t minute, int32_t second = 0) {
        if (halted()) return;  // a halted program writes nothing
#ifdef MUJARD_HOST_TEST
        setAt_ = epoch(year, month, day, hour, minute, second);
        setWhen_ = now();
#else
        Wire.beginTransmission(0x68);
        Wire.write((uint8_t)0);  // register 0
        Wire.write(bcd(second));
        Wire.write(bcd(minute));
        Wire.write(bcd(hour));  // 24-hour mode: bit 6 stays 0
        Wire.write(bcd(weekday(year, month, day)));
        Wire.write(bcd(day));
        Wire.write(bcd(month));
        Wire.write(bcd(year % 100));
        Wire.endTransmission();
#endif
    }

    /** The chip's time into out; false when there is no answer or it was never set (a year before 2023). */
    static bool read(struct tm& out) {
#ifdef MUJARD_HOST_TEST
        if (setAt_ == 0) return false;
        time_t t = (time_t)(setAt_ + (now() - setWhen_) / 1000000);
        gmtime_r(&t, &out);
#else
        Wire.beginTransmission(0x68);
        Wire.write((uint8_t)0);
        if (Wire.endTransmission(false) != 0) return false;
        if (Wire.requestFrom((uint8_t)0x68, (uint8_t)7) != 7) return false;
        out.tm_sec = unbcd(Wire.read());
        out.tm_min = unbcd(Wire.read());
        out.tm_hour = unbcd(Wire.read() & 0x3F);
        Wire.read();  // the chip's own day-of-week; the date's is computed below
        out.tm_mday = unbcd(Wire.read());
        out.tm_mon = unbcd(Wire.read() & 0x7F) - 1;
        out.tm_year = unbcd(Wire.read()) + 100;  // the chip counts from 2000
        out.tm_wday = weekday(out.tm_year + 1900, out.tm_mon + 1, out.tm_mday) % 7;
#endif
        return out.tm_year + 1900 >= 2023;
    }

private:
    static uint8_t bcd(int32_t v) {
        if (v < 0) v = 0;
        if (v > 99) v = 99;
        return (uint8_t)((v / 10) << 4 | (v % 10));
    }
    static int32_t unbcd(uint8_t v) { return (v >> 4) * 10 + (v & 15); }
    /** Day of the week for the date, 0 = Sunday (Sakamoto). */
    static int weekday(int y, int m, int d) {
        static const int t[] = {0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4};
        if (m < 3) y--;
        return (y + y / 4 - y / 100 + y / 400 + t[m - 1] + d) % 7;
    }
#ifdef MUJARD_HOST_TEST
    /** Days-since-1970 for a date (Hinnant's days_from_civil), so the fake chip keeps real calendar time. */
    static int64_t epoch(int y, int m, int d, int hh, int mm, int ss) {
        y -= m <= 2;
        int64_t era = (y >= 0 ? y : y - 399) / 400;
        int64_t yoe = y - era * 400;
        int64_t doy = (153 * (m + (m > 2 ? -3 : 9)) + 2) / 5 + d - 1;
        int64_t doe = yoe * 365 + yoe / 4 - yoe / 100 + doy;
        return (era * 146097 + doe - 719468) * 86400 + hh * 3600 + mm * 60 + ss;
    }
    static int64_t setAt_;
    static Time setWhen_;
#endif
};

#ifdef MUJARD_HOST_TEST
int64_t Rtc::setAt_ = 0;
Time Rtc::setWhen_ = 0;
#endif

static Rtc rtc_dev;

}  // namespace mujard
