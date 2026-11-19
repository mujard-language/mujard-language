// http.get / http.post (spec Appendix A 79), used inside a task. The request runs in the background: http_begin
// starts it, the task waits on http_done, and http_result gives the reply - so the rest of the program keeps
// running meanwhile. One request at a time: a second call while one is in flight answers status -1 at once.
//
// https:// is done only with the server's certificate checked: the ESP32 uses the core's CA bundle, the ESP8266
// and the Pico W the roots built in below. Where that check can't run (a board without Wi-Fi, the PC tests) an
// https:// address is refused with status -1 rather than sent unchecked.
#pragma once
#include "mujard_wifi.h"

// The secure client is picked by the core's own version, never by probing for headers: the build server refuses
// sketches that probe which files exist.
#if defined(ARDUINO_ARCH_ESP32)
#if defined(ESP_ARDUINO_VERSION_MAJOR) && ESP_ARDUINO_VERSION_MAJOR >= 3
// Core 3.x: the full Mozilla bundle is built into the prebuilt IDF libraries
// (CONFIG_MBEDTLS_CERTIFICATE_BUNDLE_DEFAULT_FULL), so useBuiltinCACertBundle() is the check.
// The WiFiClientSecure name stays through the core's compat header.
#include <NetworkClientSecure.h>
#include <WiFiClientSecure.h>
#define MUJARD_TLS_BUNDLE_BUILTIN 1
#else
#include <WiFiClientSecure.h>
extern "C" const uint8_t x509_crt_bundle[] __attribute__((weak));  // the core's CA bundle, when enabled
#define MUJARD_TLS_BUNDLE 1
#endif
#elif defined(ARDUINO_ARCH_ESP8266) || defined(ARDUINO_ARCH_RP2040)
#include <WiFiClientSecureBearSSL.h>  // both cores ship it (ESP8266 2.5+, arduino-pico)
#define MUJARD_TLS_BEARSSL 1
#endif

#ifdef MUJARD_TLS_BEARSSL
// The trust anchors ESP8266 / Pico programs check https certificates against: the roots behind most public
// sites (Let's Encrypt, DigiCert, Google Trust Services). Compiled into the sketch.
static const char mujard_ca_roots[] =
    "-----BEGIN CERTIFICATE-----\n"
    "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\n"
    "-----END CERTIFICATE-----\n"  // ISRG Root X1
    "-----BEGIN CERTIFICATE-----\n"
    "MIICGzCCAaGgAwIBAgIQQdKd0XLq7qeAwSxs6S+HUjAKBggqhkjOPQQDAzBPMQsw\n"
    "CQYDVQQGEwJVUzEpMCcGA1UEChMgSW50ZXJuZXQgU2VjdXJpdHkgUmVzZWFyY2gg\n"
    "R3JvdXAxFTATBgNVBAMTDElTUkcgUm9vdCBYMjAeFw0yMDA5MDQwMDAwMDBaFw00\n"
    "MDA5MTcxNjAwMDBaME8xCzAJBgNVBAYTAlVTMSkwJwYDVQQKEyBJbnRlcm5ldCBT\n"
    "ZWN1cml0eSBSZXNlYXJjaCBHcm91cDEVMBMGA1UEAxMMSVNSRyBSb290IFgyMHYw\n"
    "EAYHKoZIzj0CAQYFK4EEACIDYgAEzZvVn4CDCuwJSvMWSj5cz3es3mcFDR0HttwW\n"
    "+1qLFNvicWDEukWVEYmO6gbf9yoWHKS5xcUy4APgHoIYOIvXRdgKam7mAHf7AlF9\n"
    "ItgKbppbd9/w+kHsOdx1ymgHDB/qo0IwQDAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0T\n"
    "AQH/BAUwAwEB/zAdBgNVHQ4EFgQUfEKWrt5LSDv6kviejM9ti6lyN5UwCgYIKoZI\n"
    "zj0EAwMDaAAwZQIwe3lORlCEwkSHRhtFcP9Ymd70/aTSVaYgLXTWNLxBo1BfASdW\n"
    "tL4ndQavEi51mI38AjEAi/V3bNTIZargCyzuFJ0nN6T5U6VR5CmD1/iQMVtCnwr1\n"
    "/q4AaOeMSQ+2b1tbFfLn\n"
    "-----END CERTIFICATE-----\n"  // ISRG Root X2
    "-----BEGIN CERTIFICATE-----\n"
    "MIIDjjCCAnagAwIBAgIQAzrx5qcRqaC7KGSxHQn65TANBgkqhkiG9w0BAQsFADBh\n"
    "MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\n"
    "d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBH\n"
    "MjAeFw0xMzA4MDExMjAwMDBaFw0zODAxMTUxMjAwMDBaMGExCzAJBgNVBAYTAlVT\n"
    "MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\n"
    "b20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IEcyMIIBIjANBgkqhkiG\n"
    "9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuzfNNNx7a8myaJCtSnX/RrohCgiN9RlUyfuI\n"
    "2/Ou8jqJkTx65qsGGmvPrC3oXgkkRLpimn7Wo6h+4FR1IAWsULecYxpsMNzaHxmx\n"
    "1x7e/dfgy5SDN67sH0NO3Xss0r0upS/kqbitOtSZpLYl6ZtrAGCSYP9PIUkY92eQ\n"
    "q2EGnI/yuum06ZIya7XzV+hdG82MHauVBJVJ8zUtluNJbd134/tJS7SsVQepj5Wz\n"
    "tCO7TG1F8PapspUwtP1MVYwnSlcUfIKdzXOS0xZKBgyMUNGPHgm+F6HmIcr9g+UQ\n"
    "vIOlCsRnKPZzFBQ9RnbDhxSJITRNrw9FDKZJobq7nMWxM4MphQIDAQABo0IwQDAP\n"
    "BgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNVHQ4EFgQUTiJUIBiV\n"
    "5uNu5g/6+rkS7QYXjzkwDQYJKoZIhvcNAQELBQADggEBAGBnKJRvDkhj6zHd6mcY\n"
    "1Yl9PMWLSn/pvtsrF9+wX3N3KjITOYFnQoQj8kVnNeyIv/iPsGEMNKSuIEyExtv4\n"
    "NeF22d+mQrvHRAiGfzZ0JFrabA0UWTW98kndth/Jsw1HKj2ZL7tcu7XUIOGZX1NG\n"
    "Fdtom/DzMNU+MeKNhJ7jitralj41E6Vf8PlwUHBHQRFXGU7Aj64GxJUTFy8bJZ91\n"
    "8rGOmaFvE7FBcf6IKshPECBV1/MUReXgRPTqh5Uykw7+U0b6LJ3/iyK5S9kJRaTe\n"
    "pLiaWN0bfVKfjllDiIGknibVb63dDcY3fe0Dkhvld1927jyNxF1WW6LZZm6zNTfl\n"
    "MrY=\n"
    "-----END CERTIFICATE-----\n"  // DigiCert Global Root G2
    "-----BEGIN CERTIFICATE-----\n"
    "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\n"
    "bP6MvPJwNQzcmRk13NfIRmPVNnGuV/u3gm3c\n"
    "-----END CERTIFICATE-----\n";  // GTS Root R1
#endif

// A build that gives the board a notify key writes MUJARD_NOTIFY_URL / MUJARD_NOTIFY_KEY into the same per-build
// secrets header the program's own secret() values use. Every build writes that header (empty when there's nothing
// in it): the server, mujard build and the test builds alike. PC test builds have no notify path.
#if !defined(MUJARD_HOST_TEST)
#include "mujard_secrets.h"
#endif

namespace mujard {

struct HttpReply {
    bool ok = false;
    int32_t status = 0;  // the HTTP status, or -1 when there was no answer
    Text text;           // the body, cut to the text size
};

// reply.json - a little JSON reader. It keeps the value's own text; ["key"] and [i] find the part,
// .ok says it was there, to_number() makes a number of it, to_text() the text, and 'as <unit>' a quantity.
struct Json {
    Text raw;
    bool ok;
    Json() : ok(false) {}
    Json(const Text& src) : raw(src), ok(kind(src.s) != 0) {}

    /** j["key"]: the member's value, or a Json with ok false when the object hasn't got it. */
    Json get(const char* key) const {
        Json out;
        const char* s = ws(raw.s);
        if (*s != '{' || !key) return out;
        for (;;) {
            s = ws(s + 1);
            if (*s != '"') return out;  // '}' or a broken object: not found
            const char* ke = strEnd(s);
            if (!ke) return out;
            const char* colon = ws(ke);
            if (*colon != ':') return out;
            bool match = keyEq(s + 1, ke - 1, key);  // ke is past the closing quote
            const char* v = ws(colon + 1);
            const char* ve = valueEnd(v);
            if (!ve) return out;
            if (match) {
                out.raw = fragment(v, ve);
                out.ok = true;
                return out;
            }
            s = ws(ve);
            if (*s != ',') return out;
        }
    }
    /** j[i]: the element, or a Json with ok false when the array is shorter or isn't one. */
    Json at(int32_t i) const {
        Json out;
        if (i < 0) return out;
        const char* s = ws(raw.s);
        if (*s != '[') return out;
        s = ws(s + 1);
        if (*s == ']') return out;
        for (;;) {
            const char* e = valueEnd(s);
            if (!e) return out;
            if (i == 0) {
                out.raw = fragment(s, e);
                out.ok = true;
                return out;
            }
            i--;
            s = ws(e);
            if (*s != ',') return out;
            s = ws(s + 1);
            if (*s == ']') return out;
        }
    }
    /** The number the value holds (a bare number, or a number in quotes), or none. */
    Opt<float> to_number() const {
        char b[MUJARD_TEXT_MAX + 1];
        const char* s = raw.s;
        if (!ok) return Opt<float>();
        if (*ws(s) == '"') {
            unquote(s, b, sizeof b);
            s = b;
        }
        char* end;
        float v = strtof(s, &end);
        if (end == s) return Opt<float>();
        return Opt<float>(v);
    }
    /** The text the value holds: a string comes without its quotes, anything else stays JSON. */
    Text to_text() const {
        Text t;
        if (!ok) return t;
        if (*ws(raw.s) == '"') {
            char b[MUJARD_TEXT_MAX + 1];
            unquote(raw.s, b, sizeof b);
            t.add(b);
            return t;
        }
        return raw;
    }
    operator float() const {
        Opt<float> n = to_number();
        return n.has ? n.v : 0;
    }

private:
    static const char* ws(const char* s) {
        while (*s == ' ' || *s == '\t' || *s == '\r' || *s == '\n') s++;
        return s;
    }
    /** 0 when s doesn't start a JSON value, else its first character. */
    static char kind(const char* s) {
        s = ws(s);
        return *s && strchr("{[\"-0123456789tfn", *s) ? *s : 0;
    }
    /** Past the closing quote of the string s opens with, or nullptr. */
    static const char* strEnd(const char* s) {
        for (const char* p = s + 1; *p; p++) {
            if (*p == '\\') {
                if (!*++p) return nullptr;
                continue;
            }
            if (*p == '"') return p + 1;
        }
        return nullptr;
    }
    /** Just past the JSON value s starts with, or nullptr when it doesn't hold one. */
    static const char* valueEnd(const char* s) {
        s = ws(s);
        switch (*s) {
        case '"':
            return strEnd(s);
        case '{':
            for (s = ws(s + 1); *s != '}';) {
                if (*s != '"') return nullptr;
                s = strEnd(s);
                if (!s || *ws(s) != ':') return nullptr;
                s = valueEnd(ws(s) + 1);
                if (!s) return nullptr;
                s = ws(s);
                if (*s == ',') {
                    s = ws(s + 1);
                    continue;
                }
            }
            return s + 1;
        case '[':
            for (s = ws(s + 1); *s != ']';) {
                s = valueEnd(s);
                if (!s) return nullptr;
                s = ws(s);
                if (*s == ',') {
                    s = ws(s + 1);
                    continue;
                }
            }
            return s + 1;
        case 't':
            return strncmp(s, "true", 4) == 0 ? s + 4 : nullptr;
        case 'f':
            return strncmp(s, "false", 5) == 0 ? s + 5 : nullptr;
        case 'n':
            return strncmp(s, "null", 4) == 0 ? s + 4 : nullptr;
        default: {
            const char* e = s;
            if (*e == '-' || *e == '+') e++;
            bool any = false;
            while (isdigit((uint8_t)*e) || *e == '.' || *e == 'e' || *e == 'E' || *e == '+' || *e == '-') {
                e++;
                any = true;
            }
            return any ? e : nullptr;
        }
        }
    }
    /** A key span (no quotes) against a plain C string; keys with escape characters don't match. */
    static bool keyEq(const char* a, const char* b, const char* key) {
        while (a < b && *key) {
            if (*a != *key) return false;
            a++;
            key++;
        }
        return a == b && !*key;
    }
    /** The fragment a..b as a Text (cut to its size). */
    static Text fragment(const char* a, const char* b) {
        Text t;
        while (a < b && t.n < MUJARD_TEXT_MAX) t.s[t.n++] = *a++;
        t.s[t.n] = 0;
        return t;
    }
    /** The inside of the JSON string s, with the usual escapes, into out. */
    static void unquote(const char* s, char* out, size_t cap) {
        s = ws(s);
        size_t n = 0;
        for (const char* p = s + 1; *p && *p != '"' && n + 1 < cap; p++) {
            char c = *p;
            if (c == '\\' && p[1]) {
                c = *++p;
                if (c == 'n') c = '\n';
                else if (c == 't') c = '\t';
                else if (c == 'r') c = '\r';
                else if (c == 'u') {  // \uXXXX: only ASCII maps into a Mujard text
                    int v = 0;
                    for (int k = 0; k < 4 && isxdigit((uint8_t)p[1]); k++) {
                        p++;
                        v = v * 16 + (*p <= '9' ? *p - '0' : (*p & 31) + 9);
                    }
                    c = v > 0 && v < 127 ? (char)v : '?';
                }
            }
            out[n++] = c;
        }
        out[n] = 0;
    }
};

// ---- one request in the background, polled by http_done() on each pass of the waiting task

struct HttpJob {
    WiFiClient* client = nullptr;
    int32_t id = 0;      // the call site's number; 0 = no request in flight
    HttpReply reply;
    Time deadline = 0;
    bool finished = false;
    // the response parser
    Text line;           // the status line / the header line being read
    uint8_t phase = 0;   // 0 status line, 1 headers, 2 body by length, 3 chunk size, 4 chunk data, 5 chunk tail, 6 to close
    bool chunked = false;
    int32_t length = -1;  // content-length, -1 = the body's until the connection closes
    int32_t chunkLeft = 0;
};

static HttpJob http_job;

// A call that can't start (busy, bad address, ...) gets its -1 reply here: the task's wait ends and
// http_result(id) gives this text.
struct HttpFail {
    int32_t id;
    const char* why;
};
static HttpFail http_failed[4];
static uint8_t http_fails = 0;

static void http_fail_now(int32_t id, const char* why) {
    for (HttpFail& f : http_failed) {
        if (f.id == id) return;  // already noted
    }
    if (http_fails < 4) {
        http_failed[http_fails].id = id;
        http_failed[http_fails].why = why;
        http_fails++;
    }
}

#if defined(MUJARD_TLS_BUNDLE) || defined(MUJARD_TLS_BUNDLE_BUILTIN) || defined(MUJARD_TLS_BEARSSL)
/** Certificates carry validity dates, so a check needs the clock: start a plain UTC sync once if it isn't set
 *  yet. The first https call may still answer "the time isn't set yet" - later ones check properly. */
static void https_time() {
    if (::time(nullptr) > 1700000000) return;
#if defined(ARDUINO_ARCH_ESP32)
    configTzTime("UTC0", "pool.ntp.org", "time.nist.gov");
#elif defined(ARDUINO_ARCH_ESP8266)
    configTime("UTC0", "pool.ntp.org", "time.nist.gov");
#elif defined(ARDUINO_ARCH_RP2040) && !defined(ARDUINO_ARCH_MBED)
    setenv("TZ", "UTC0", 1);
    tzset();
    NTP.begin("pool.ntp.org", "time.nist.gov");
#endif
}
#endif

/** A plain or a certificate-checking connection for the address, or nullptr with the reason in why. */
static WiFiClient* http_client(bool secure, const char** why) {
#if defined(MUJARD_TLS_BUNDLE_BUILTIN)
    if (secure) {
        https_time();
        WiFiClientSecure* s = new WiFiClientSecure();
        s->useBuiltinCACertBundle();
        return s;
    }
#elif defined(MUJARD_TLS_BUNDLE)
    if (secure) {
        if (x509_crt_bundle) {
            https_time();
            WiFiClientSecure* s = new WiFiClientSecure();
            s->setCACertBundle(x509_crt_bundle);
            return s;
        }
        *why = "This build has no certificate bundle for https";
        return nullptr;
    }
#elif defined(MUJARD_TLS_BEARSSL)
    if (secure) {
        https_time();
        WiFiClientSecure* s = new WiFiClientSecure();
        static BearSSL::X509List* anchors = nullptr;
        if (!anchors) anchors = new BearSSL::X509List(mujard_ca_roots);
        s->setTrustAnchors(anchors);
        if (::time(nullptr) > 1700000000) s->setX509Time(::time(nullptr));  // certificates carry validity dates
        return s;
    }
#else
    if (secure) {
        *why = "https needs a Wi-Fi board with certificate checks (ESP32, ESP8266, Pico W)";
        return nullptr;
    }
#endif
    (void)why;
    return new WiFiClient();
}

void http_begin(int32_t id, bool post, const Text& url, const Text& body) {
    if (http_job.id == id) return;         // the call's own request is already going
    if (http_job.id != 0) {                // one at a time: the second waits its turn... 
        http_fail_now(id, "Another http request is in flight");
        return;                          // ...or rather, it fails: a task waiting twice would never match
    }
    HttpReply bad;
    bad.status = -1;
    const char* fail = nullptr;
    bool secure = strncmp(url.s, "https://", 8) == 0;
    if (halted()) {
        fail = "The program is halted";
    } else if (!secure && strncmp(url.s, "http://", 7) != 0) {
        fail = "The address must start with http:// or https://";
    } else if (!net_up()) {
        fail = "Not connected to Wi-Fi";
    }
    if (fail) {
        http_fail_now(id, fail);
        return;
    }
    // split the address: host[:port]/path
    const char* rest = url.s + (secure ? 8 : 7);
    const char* slash = strchr(rest, '/');
    const char* path = slash ? slash : "/";
    char host[MUJARD_TEXT_MAX + 1];
    uint16_t port = secure ? 443 : 80;
    size_t hl = slash ? (size_t)(slash - rest) : strlen(rest);
    const char* colon = (const char*)memchr(rest, ':', hl);
    if (colon) {
        port = (uint16_t)strtol(colon + 1, nullptr, 10);
        hl = (size_t)(colon - rest);
    }
    if (hl == 0 || hl > MUJARD_TEXT_MAX) {
        http_fail_now(id, "Bad address");
        return;
    }
    memcpy(host, rest, hl);
    host[hl] = 0;

    WiFiClient* client = http_client(secure, &fail);
    if (!client) {
        http_fail_now(id, fail);
        return;
    }
#if defined(MUJARD_HOST_TEST)
    if (!client->connect(host, port)) {
#elif defined(ARDUINO_ARCH_ESP32)
    if (!client->connect(host, port, 2000)) {  // the only blocking moment: 2 s to open the connection
#else
    // The ESP8266 and Pico W cores have no timeout argument on connect(): they use the stream timeout (in ms) for it.
    client->setTimeout(2000);
    if (!client->connect(host, port)) {
#endif
        delete client;
        http_fail_now(id, "Couldn't connect");
        return;
    }
    // the request: HTTP/1.0 with "Connection: close", so the answer ends by closing
    char req[MUJARD_TEXT_MAX + 48];
    int n = snprintf(req, sizeof req, "%s %s HTTP/1.0\r\nHost: %s\r\nConnection: close\r\n", post ? "POST" : "GET",
                     path, host);
    client->write((const uint8_t*)req, (size_t)n);
    if (post) {
        n = snprintf(req, sizeof req, "Content-Length: %d\r\n\r\n", (int)body.n);
        client->write((const uint8_t*)req, (size_t)n);
        client->write((const uint8_t*)body.s, body.n);
    } else {
        client->write((const uint8_t*)"\r\n", 2);
    }
    http_job = HttpJob();
    http_job.client = client;
    http_job.id = id;
    http_job.deadline = now() + 5000000;  // at most 5 s of waiting (spec Appendix A 79)
}

/** Parses the response one character at a time: status line, headers, then the body by content-length,
 *  chunked, or until the connection closes. */
static void http_feed(uint8_t c) {
    HttpJob& j = http_job;
    switch (j.phase) {
    case 0:  // "HTTP/1.1 200 OK"
        if (c == '\n') {
            const char* s = j.line.s;
            if (strncmp(s, "HTTP/", 5) == 0) {
                const char* sp = strchr(s, ' ');
                if (sp) j.reply.status = (int32_t)strtol(sp + 1, nullptr, 10);
            }
            j.line = Text();
            j.phase = 1;
            return;
        }
        if (c != '\r') {
            char one[2] = {(char)c, 0};
            j.line.add(one);
        }
        return;
    case 1:  // header lines until an empty one
        if (c == '\n') {
            Text& h = j.line;
            if (h.n == 0) {  // blank line: the body follows
                j.phase = j.chunked ? 3 : (j.length >= 0 ? 2 : 6);
                j.line = Text();
                if (j.length == 0) j.finished = true;
                return;
            }
            if (h.n >= 15 && strncasecmp(h.s, "Content-Length:", 15) == 0)
                j.length = (int32_t)strtol(h.s + 15, nullptr, 10);
            else if (h.n >= 19 && strncasecmp(h.s, "Transfer-Encoding:", 18) == 0 && strstr(h.s, "chunked"))
                j.chunked = true;
            h = Text();
            return;
        }
        if (c != '\r') {
            char one[2] = {(char)c, 0};
            j.line.add(one);
        }
        return;
    case 2:  // a body of known length (the text cuts itself at its size; the count keeps going)
        {
            char one[2] = {(char)c, 0};
            j.reply.text.add(one);
            if (--j.length <= 0) j.finished = true;
        }
        return;
    case 3:  // a chunk size line
        if (c == '\n') {
            j.chunkLeft = (int32_t)strtol(j.line.s, nullptr, 16);
            j.line = Text();
            j.phase = j.chunkLeft > 0 ? 4 : 5;
            if (j.phase == 5) j.finished = true;  // the closing "0" chunk ends it
            return;
        }
        if (c != '\r') {
            char one[2] = {(char)c, 0};
            j.line.add(one);
        }
        return;
    case 4:  // chunk data
        {
            char one[2] = {(char)c, 0};
            j.reply.text.add(one);
            if (--j.chunkLeft <= 0) j.phase = 5;
        }
        return;
    case 5:  // the CRLF after a chunk
        if (c == '\n') j.phase = 3;
        return;
    default:  // a body read until the connection closes
        {
            char one[2] = {(char)c, 0};
            j.reply.text.add(one);
        }
        return;
    }
}

static void http_finish() {
    HttpJob& j = http_job;
    j.finished = true;
    if (j.reply.status == 0) j.reply.status = -1;
    j.reply.ok = j.reply.status >= 200 && j.reply.status < 300;
    if (j.client) {
        j.client->stop();
        delete j.client;
        j.client = nullptr;
    }
}

bool http_done(int32_t id) {
    HttpJob& j = http_job;
    if (j.id != id) {
        for (HttpFail& f : http_failed)
            if (f.id == id) return true;
        return false;
    }
    while (j.client && j.client->available() > 0 && !j.finished) http_feed((uint8_t)j.client->read());
    if (j.finished) http_finish();  // the parser finished it: compute ok and let the connection go
    if (!j.finished && j.client && !j.client->connected() && j.client->available() == 0) http_finish();
    if (!j.finished && now() >= j.deadline) http_finish();
    return j.finished;
}

HttpReply http_result(int32_t id) {
    if (http_job.id == id) {
        HttpReply r = http_job.reply;
        if (r.status == 0) r.status = -1;
        http_job = HttpJob();
        return r;
    }
    HttpReply r;
    r.status = -1;
    r.text = Text("No answer for this request");
    for (uint8_t i = 0; i < http_fails; i++) {
        if (http_failed[i].id == id) {
            r.text = Text(http_failed[i].why);
            http_failed[i] = http_failed[--http_fails];
            break;
        }
    }
    return r;
}

/** True while a request (a program's http call or a notify post) is in flight. */
inline bool http_busy() { return http_job.id != 0; }

// ---- notify to the account, no app needed (Phase 9)
//
// When the build carried a notify key, app_notify hands each message here as well: it queues up to
// three and posts them one at a time through the same background machinery as http.post, driven by
// the scheduler's tick - so a notify never waits and the rest of the program never pauses for it.
#if defined(MUJARD_NOTIFY_URL) && defined(MUJARD_NOTIFY_KEY)
#define MUJARD_NOTIFY_ID 0x40000000  // call-site ids are small positive numbers; this can't collide

struct NotifySlot {
    Text text;
    bool urgent = false;
    bool used = false;
};
static NotifySlot notify_q[3];
static volatile bool notify_in_flight = false;

static void notify_net_impl(const char* msg, bool urgent) {
    for (NotifySlot& q : notify_q) {
        if (!q.used) {
            q.text = Text(msg);
            q.urgent = urgent;
            q.used = true;
            return;
        }
    }
    // all three are waiting already: the new one is dropped - the serial line still went out
}

/** The JSON body {"k":key,"t":message,"u":0|1}, with the message escaped for a JSON string. */
static void notify_body(Text& out, const char* msg, bool urgent) {
    out.add("{\"k\":\"").add(MUJARD_NOTIFY_KEY).add("\",\"t\":\"");
    for (const char* p = msg; *p; p++) {
        char c = *p;
        char one[2] = {c, 0};
        if (c == '"' || c == '\\') out.add("\\").add(one);
        else if ((unsigned char)c < 0x20) out.add(" ");
        else out.add(one);
    }
    out.add("\",\"u\":").add(urgent ? "1" : "0").add("}");
}

/** Drives the queue: finish the in-flight post, then start the oldest waiting one. */
struct NotifyNet : Ticked {
    void tick(Time t) override {
        (void)t;
        if (notify_in_flight) {
            if (!http_done(MUJARD_NOTIFY_ID)) return;
            http_result(MUJARD_NOTIFY_ID);  // frees the job; what the server answered doesn't matter
            notify_in_flight = false;
        }
        if (http_busy()) return;  // a program http call has the slot - try again on the next pass
        for (NotifySlot& q : notify_q) {
            if (!q.used) continue;
            Text body;
            notify_body(body, q.text.s, q.urgent);
            q.used = false;
            notify_in_flight = true;
            http_begin(MUJARD_NOTIFY_ID, true, Text(MUJARD_NOTIFY_URL), body);
            return;
        }
    }
};
static NotifyNet notify_net_dev;
static struct NotifyInit {
    NotifyInit() { notify_net_fn = &notify_net_impl; }
} notify_init_;
#endif

}  // namespace mujard
