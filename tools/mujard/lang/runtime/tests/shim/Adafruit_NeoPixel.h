// A fake Adafruit NeoPixel for runs on the PC (mujard test): it remembers the fill and what it was told to show.
#pragma once
#include "Arduino.h"

#define NEO_GRB 0
#define NEO_RGB 1
#define NEO_RGBW 2
#define NEO_GRBW 3
#define NEO_KHZ800 0
#define NEO_KHZ400 0x0100

namespace fake {
inline int& neoFill() { static int c = -1; return c; }
inline int& neoShown() { static int n = 0; return n; }
}  // namespace fake

class Adafruit_NeoPixel {
public:
    Adafruit_NeoPixel(int n, int pin = -1, int type = 0) : n_(n), pin_(pin) {}
    void begin() {}
    void show() { fake::neoShown()++; }
    void setPixelColor(int, int, int = 0, int = 0, int = 0) {}
    void fill(int color = 0, int = 0, int = 0) { fake::neoFill() = color; }
    void clear() { fake::neoFill() = 0; }
    void setBrightness(int b) { bright_ = b; }
    int getBrightness() const { return bright_; }
    void setPin(int p) { pin_ = p; }
    int getPin() const { return pin_; }
    int numPixels() const { return n_; }
    int getPixelColor(int) const { return fake::neoFill(); }
    static int Color(int r, int g, int b, int = 0) { return (r << 16) | (g << 8) | b; }
    static int ColorHSV(int hue, int = 255, int = 255) { return hue; }
    static int gamma32(int x) { return x; }

private:
    int n_, pin_, bright_ = 255;
};
