package lang

import "regexp"

// timeZones turns the names people write ("Africa/Lagos", from the IANA time zone list) into the POSIX TZ rules the
// boards' C libraries understand ("WAT-1"), at compile time: the boards carry no time zone database, and a misspelt
// zone is a compile error (C326) instead of a clock that is silently hours off. The rules are those in force in 2026.
var timeZones = map[string]string{
	"UTC": "UTC0", "Etc/UTC": "UTC0", "GMT": "GMT0",

	"Africa/Abidjan": "GMT0", "Africa/Accra": "GMT0", "Africa/Addis_Ababa": "EAT-3", "Africa/Algiers": "CET-1",
	"Africa/Cairo": "EET-2EEST,M4.5.5/0,M10.5.4/24", "Africa/Casablanca": "<+01>-1", "Africa/Dakar": "GMT0",
	"Africa/Dar_es_Salaam": "EAT-3", "Africa/Douala": "WAT-1", "Africa/Harare": "CAT-2", "Africa/Johannesburg": "SAST-2",
	"Africa/Kampala": "EAT-3", "Africa/Khartoum": "CAT-2", "Africa/Kigali": "CAT-2", "Africa/Kinshasa": "WAT-1",
	"Africa/Lagos": "WAT-1", "Africa/Luanda": "WAT-1", "Africa/Lusaka": "CAT-2", "Africa/Maputo": "CAT-2",
	"Africa/Nairobi": "EAT-3", "Africa/Tunis": "CET-1",

	"Europe/Amsterdam": "CET-1CEST,M3.5.0,M10.5.0/3", "Europe/Athens": "EET-2EEST,M3.5.0/3,M10.5.0/4",
	"Europe/Berlin": "CET-1CEST,M3.5.0,M10.5.0/3", "Europe/Brussels": "CET-1CEST,M3.5.0,M10.5.0/3",
	"Europe/Bucharest": "EET-2EEST,M3.5.0/3,M10.5.0/4", "Europe/Copenhagen": "CET-1CEST,M3.5.0,M10.5.0/3",
	"Europe/Dublin": "IST-1GMT0,M10.5.0,M3.5.0/1", "Europe/Helsinki": "EET-2EEST,M3.5.0/3,M10.5.0/4",
	"Europe/Istanbul": "<+03>-3", "Europe/Kyiv": "EET-2EEST,M3.5.0/3,M10.5.0/4", "Europe/Lisbon": "WET0WEST,M3.5.0/1,M10.5.0",
	"Europe/London": "GMT0BST,M3.5.0/1,M10.5.0", "Europe/Madrid": "CET-1CEST,M3.5.0,M10.5.0/3",
	"Europe/Moscow": "MSK-3", "Europe/Oslo": "CET-1CEST,M3.5.0,M10.5.0/3", "Europe/Paris": "CET-1CEST,M3.5.0,M10.5.0/3",
	"Europe/Prague": "CET-1CEST,M3.5.0,M10.5.0/3", "Europe/Rome": "CET-1CEST,M3.5.0,M10.5.0/3",
	"Europe/Stockholm": "CET-1CEST,M3.5.0,M10.5.0/3", "Europe/Vienna": "CET-1CEST,M3.5.0,M10.5.0/3",
	"Europe/Warsaw": "CET-1CEST,M3.5.0,M10.5.0/3", "Europe/Zurich": "CET-1CEST,M3.5.0,M10.5.0/3",

	"America/Anchorage": "AKST9AKDT,M3.2.0,M11.1.0", "America/Argentina/Buenos_Aires": "<-03>3",
	"America/Bogota": "<-05>5", "America/Chicago": "CST6CDT,M3.2.0,M11.1.0", "America/Denver": "MST7MDT,M3.2.0,M11.1.0",
	"America/Lima": "<-05>5", "America/Los_Angeles": "PST8PDT,M3.2.0,M11.1.0", "America/Mexico_City": "CST6",
	"America/New_York": "EST5EDT,M3.2.0,M11.1.0", "America/Phoenix": "MST7", "America/Santiago": "<-04>4<-03>,M9.1.6/24,M4.1.6/24",
	"America/Sao_Paulo": "<-03>3", "America/Toronto": "EST5EDT,M3.2.0,M11.1.0", "America/Vancouver": "PST8PDT,M3.2.0,M11.1.0",
	"Pacific/Honolulu": "HST10",

	"Asia/Bangkok": "<+07>-7", "Asia/Dhaka": "<+06>-6", "Asia/Dubai": "<+04>-4", "Asia/Hong_Kong": "HKT-8",
	"Asia/Jakarta": "WIB-7", "Asia/Jerusalem": "IST-2IDT,M3.4.4/26,M10.5.0", "Asia/Karachi": "PKT-5",
	"Asia/Kathmandu": "<+0545>-5:45", "Asia/Kolkata": "IST-5:30", "Asia/Manila": "PST-8", "Asia/Riyadh": "<+03>-3",
	"Asia/Seoul": "KST-9", "Asia/Shanghai": "CST-8", "Asia/Singapore": "<+08>-8", "Asia/Tehran": "<+0330>-3:30",
	"Asia/Tokyo": "JST-9",

	"Australia/Adelaide": "ACST-9:30ACDT,M10.1.0,M4.1.0/3", "Australia/Brisbane": "AEST-10",
	"Australia/Melbourne": "AEST-10AEDT,M10.1.0,M4.1.0/3", "Australia/Perth": "AWST-8",
	"Australia/Sydney": "AEST-10AEDT,M10.1.0,M4.1.0/3", "Pacific/Auckland": "NZST-12NZDT,M9.5.0,M4.1.0/3",
}

// posixTZ matches a TZ rule written out by hand ("WAT-1", "<+0545>-5:45", "CET-1CEST,M3.5.0,M10.5.0/3").
var posixTZ = regexp.MustCompile(`^(?:[A-Za-z]{3,}|<[+-]?[0-9A-Za-z]+>)[+-]?\d{1,2}(?::\d{2})?(?:(?:[A-Za-z]{3,}|<[+-]?[0-9A-Za-z]+>)(?:[+-]?\d{1,2}(?::\d{2})?)?(?:,M\d{1,2}\.\d\.\d(?:/-?\d{1,3}(?::\d{2})?)?){2})?$`)

// posixZone gives the POSIX rule for a zone name or a hand-written rule; ok is false for anything else.
func posixZone(name string) (string, bool) {
	if tz, ok := timeZones[name]; ok {
		return tz, true
	}
	if posixTZ.MatchString(name) {
		return name, true
	}
	return "", false
}
