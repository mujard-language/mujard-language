# Described libraries

`use "Name"` (or `use arduino "Name"`) brings a library in; the checker verifies its calls against the shipped description. `use mujard.<name>` libraries are built in and listed under Built-ins.

## `AccelStepper`

- include: `AccelStepper.h`

### `AccelStepper`

- new: `AccelStepper(interface: int, pin1: int, pin2: int, pin3: int?, pin4: int?, enable: bool?) -> AccelStepper`
- `currentPosition() -> int`
- `disableOutputs()`
- `distanceToGo() -> int`
- `enableOutputs()`
- `isRunning() -> bool`
- `maxSpeed() -> float`
- `move(relative: int)`
- `moveTo(absolute: int)`
- `run() -> bool`
- `runSpeed() -> bool`
- `runSpeedToPosition() -> bool`
- `runToNewPosition(position: int)`
- `runToPosition()`
- `setAcceleration(acceleration: float)`
- `setCurrentPosition(position: int)`
- `setEnablePin(pin: int)`
- `setMaxSpeed(speed: float)`
- `setMinPulseWidth(min_width: int)`
- `setPinsInverted(direction: bool, step: bool, enable: bool)`
- `setSpeed(speed: float)`
- `speed() -> float`
- `stop()`
- `targetPosition() -> int`

### `MultiStepper`

- new: `MultiStepper() -> MultiStepper`
- `addStepper(stepper: any) -> bool`
- `moveTo(absolute: any)`
- `run() -> bool`
- `runSpeedToPosition()`

## `Adafruit ADS1X15`

- include: `Adafruit_ADS1X15.h`
- takes over: i2c

### `Adafruit_ADS1015`

- extends `Adafruit_ADS1X15` (everything it has, this has too)
- new: `Adafruit_ADS1015() -> Adafruit_ADS1015`

### `Adafruit_ADS1115`

- extends `Adafruit_ADS1X15` (everything it has, this has too)
- new: `Adafruit_ADS1115() -> Adafruit_ADS1115`

### `Adafruit_ADS1X15`

- a base class: make one of the classes that extend it
- `begin(i2c_addr: int?) -> bool`
- `computeVolts(counts: int) -> float`
- `conversionComplete() -> bool`
- `getDataRate() -> int`
- `getFsRange() -> float`
- `getGain() -> adsGain_t`
- `getLastConversionResults() -> int`
- `readADC_Differential_0_1() -> int`
- `readADC_Differential_0_3() -> int`
- `readADC_Differential_1_3() -> int`
- `readADC_Differential_2_3() -> int`
- `readADC_SingleEnded(channel: int) -> int`
- `setDataRate(rate: int)`
- `setGain(gain: adsGain_t)`
- `startADCReading(mux: int, continuous: bool)`
- `startComparator_SingleEnded(channel: int, threshold: int)`

## `Adafruit BME280 Library`

- include: `Adafruit_BME280.h`
- takes over: i2c

### `Adafruit_BME280`

- new: `Adafruit_BME280() -> Adafruit_BME280`
- `begin(addr: int?, wire: any?) -> bool`
- `getTemperatureCompensation() -> float`
- `init() -> bool`
- `readAltitude(seaLevel: float) -> float`
- `readHumidity() -> float`
- `readPressure() -> float`
- `readTemperature() -> float`
- `seaLevelForAltitude(altitude: float, pressure: float) -> float`
- `sensorID() -> int`
- `setSampling(mode: int?, tempSampling: int?, pressSampling: int?, humSampling: int?, filter: int?, duration: int?)`
- `setTemperatureCompensation(adjustment: float)`
- `takeForcedMeasurement()`

## `Adafruit BMP280 Library`

- include: `Adafruit_BMP280.h`
- takes over: i2c

### `Adafruit_BMP280`

- new: `Adafruit_BMP280() -> Adafruit_BMP280`
- `begin(addr: int?, chipid: int?) -> bool`
- `getStatus() -> int`
- `readAltitude(seaLevelhPa: float?) -> float`
- `readPressure() -> float`
- `readTemperature() -> float`
- `reset()`
- `seaLevelForAltitude(altitude: float, atmospheric: float) -> float`
- `sensorID() -> int`
- `takeForcedMeasurement() -> bool`
- `waterBoilingPoint(pressure: float) -> float`

## `Adafruit GFX Library`

- include: `Adafruit_GFX.h`

### `Adafruit_GFX`

- a base class: make one of the classes that extend it
- `drawBitmap(x: int, y: int, bitmap: any, w: int, h: int, color: int, bg: int?)`
- `drawChar(x: int, y: int, c: int, color: int, bg: int, size: int)`
- `drawCircle(x0: int, y0: int, r: int, color: int)`
- `drawFastHLine(x: int, y: int, w: int, color: int)`
- `drawFastVLine(x: int, y: int, h: int, color: int)`
- `drawLine(x0: int, y0: int, x1: int, y1: int, color: int)`
- `drawPixel(x: int, y: int, color: int)`
- `drawRect(x: int, y: int, w: int, h: int, color: int)`
- `drawRoundRect(x: int, y: int, w: int, h: int, r: int, color: int)`
- `drawTriangle(x0: int, y0: int, x1: int, y1: int, x2: int, y2: int, color: int)`
- `fillCircle(x0: int, y0: int, r: int, color: int)`
- `fillRect(x: int, y: int, w: int, h: int, color: int)`
- `fillRoundRect(x: int, y: int, w: int, h: int, r: int, color: int)`
- `fillScreen(color: int)`
- `fillTriangle(x0: int, y0: int, x1: int, y1: int, x2: int, y2: int, color: int)`
- `getCursorX() -> int`
- `getCursorY() -> int`
- `getRotation() -> int`
- `height() -> int`
- `print(value: any, base: any?) -> int`
- `println(value: any?) -> int`
- `setCursor(x: int, y: int)`
- `setFont(font: any)`
- `setRotation(rot: int)`
- `setTextColor(color: int, bg: int?)`
- `setTextSize(s: int, sy: int?)`
- `setTextWrap(wrap: bool)`
- `width() -> int`

## `Adafruit ILI9341`

- include: `Adafruit_ILI9341.h`
- takes over: spi

### `Adafruit_ILI9341`

- new: `Adafruit_ILI9341(cs: int, dc: int, rst: int?) -> Adafruit_ILI9341`
- new: `Adafruit_ILI9341(cs: int, dc: int, mosi: int, sclk: int, rst: int, miso: int) -> Adafruit_ILI9341`
- `begin(freq: int?, spi: any?)`
- `drawCircle(x0: int, y0: int, r: int, color: int)`
- `drawLine(x0: int, y0: int, x1: int, y1: int, color: int)`
- `drawPixel(x: int, y: int, color: int)`
- `drawRect(x: int, y: int, w: int, h: int, color: int)`
- `drawRoundRect(x: int, y: int, w: int, h: int, r: int, color: int)`
- `drawTriangle(x0: int, y0: int, x1: int, y1: int, x2: int, y2: int, color: int)`
- `fillCircle(x0: int, y0: int, r: int, color: int)`
- `fillRect(x: int, y: int, w: int, h: int, color: int)`
- `fillRoundRect(x: int, y: int, w: int, h: int, r: int, color: int)`
- `fillScreen(color: int)`
- `fillTriangle(x0: int, y0: int, x1: int, y1: int, x2: int, y2: int, color: int)`
- `getRotation() -> int`
- `height() -> int`
- `invertDisplay(invert: bool)`
- `print(value: any, base: any?) -> int`
- `println(value: any?) -> int`
- `readcommand8(command: int, index: int?) -> int`
- `scrollTo(y: int)`
- `setCursor(x: int, y: int)`
- `setFont(font: any)`
- `setRotation(rot: int)`
- `setSPISpeed(freq: int)`
- `setScrollMargins(top: int, bottom: int)`
- `setTextColor(color: int, bg: int?)`
- `setTextSize(s: int, sy: int?)`
- `setTextWrap(wrap: bool)`
- `width() -> int`

## `Adafruit INA219`

- include: `Adafruit_INA219.h`
- takes over: i2c

### `Adafruit_INA219`

- new: `Adafruit_INA219(addr: int?) -> Adafruit_INA219`
- `begin() -> bool`
- `getBusVoltage_V() -> float`
- `getCurrent_mA() -> float`
- `getPower_mW() -> float`
- `getShuntVoltage_mV() -> float`
- `powerSave(on: bool)`
- `setCalibration_16V_400mA()`
- `setCalibration_32V_1A()`
- `setCalibration_32V_2A()`
- `success() -> bool`

## `Adafruit MCP23017 Arduino Library`

- include: `Adafruit_MCP23X17.h`
- takes over: i2c

### `Adafruit_MCP23X17`

- extends `Adafruit_MCP23XXX` (everything it has, this has too)
- new: `Adafruit_MCP23X17() -> Adafruit_MCP23X17`
- `enableAddrPins()`
- `readGPIOA() -> int`
- `readGPIOAB() -> int`
- `readGPIOB() -> int`
- `writeGPIOA(value: int)`
- `writeGPIOAB(value: int)`
- `writeGPIOB(value: int)`

### `Adafruit_MCP23XXX`

- a base class: make one of the classes that extend it
- `begin_I2C(i2c_addr: int?) -> bool`
- `clearInterrupts()`
- `digitalRead(pin: int) -> int`
- `digitalWrite(pin: int, value: int)`
- `disableInterruptPin(pin: int)`
- `getCapturedInterrupt() -> int`
- `getLastInterruptPin() -> int`
- `pinMode(pin: int, mode: int)`
- `readGPIO(port: int?) -> int`
- `setupInterruptPin(pin: int, mode: int?)`
- `setupInterrupts(mirroring: bool, openDrain: bool, polarity: int)`
- `writeGPIO(value: int, port: int?)`

## `Adafruit MLX90614 Library`

- include: `Adafruit_MLX90614.h`
- takes over: i2c

### `Adafruit_MLX90614`

- new: `Adafruit_MLX90614() -> Adafruit_MLX90614`
- `begin(addr: int?) -> bool`
- `readAmbientTempC() -> float`
- `readAmbientTempF() -> float`
- `readEmissivity() -> float`
- `readEmissivityReg() -> int`
- `readObjectTempC() -> float`
- `readObjectTempF() -> float`
- `writeEmissivity(emissivity: float)`
- `writeEmissivityReg(ereg: int)`

## `Adafruit MPU6050`

- include: `Adafruit_MPU6050.h`
- takes over: i2c

### `Adafruit_MPU6050`

- new: `Adafruit_MPU6050() -> Adafruit_MPU6050`
- `begin(addr: int?, wire: any?, sensorId: int?) -> bool`
- `enableCycle(enable: bool)`
- `enableSleep(enable: bool)`
- `getAccelerometerRange() -> int`
- `getAccelerometerSensor() -> any`
- `getCycleRate() -> int`
- `getEvent(accel: any, gyro: any, temp: any) -> bool`
- `getFilterBandwidth() -> int`
- `getGyroRange() -> int`
- `getGyroSensor() -> any`
- `getHighPassFilter() -> int`
- `getMotionInterruptStatus() -> bool`
- `getSampleRateDivisor() -> int`
- `getTemperatureSensor() -> any`
- `reset()`
- `setAccelerometerRange(range: int)`
- `setCycleRate(rate: int)`
- `setFilterBandwidth(bandwidth: int)`
- `setGyroRange(range: int)`
- `setHighPassFilter(bandwidth: int)`
- `setMotionDetectionDuration(value: int)`
- `setMotionDetectionThreshold(value: int)`
- `setMotionInterrupt(active: bool)`
- `setSampleRateDivisor(value: int)`

## `Adafruit NeoPixel`

- include: `Adafruit_NeoPixel.h`

### `Adafruit_NeoPixel`

- new: `Adafruit_NeoPixel(n: int, pin: int?, type: int?) -> Adafruit_NeoPixel`
- `Color(r: int, g: int, b: int, w: int?) -> int`
- `ColorHSV(hue: int, sat: int?, val: int?) -> int`
- `begin()`
- `clear()`
- `fill(color: int?, first: int?, count: int?)`
- `gamma32(x: int) -> int`
- `getBrightness() -> int`
- `getPin() -> int`
- `getPixelColor(n: int) -> int`
- `numPixels() -> int`
- `setBrightness(b: int)`
- `setPin(pin: int)`
- `setPixelColor(n: int, r_or_color: int, g: int?, b: int?, w: int?)`
- `show()`

## `Adafruit PWM Servo Driver Library`

- include: `Adafruit_PWMServoDriver.h`
- takes over: i2c

### `Adafruit_PWMServoDriver`

- new: `Adafruit_PWMServoDriver(addr: int?) -> Adafruit_PWMServoDriver`
- `begin(prescale: int?) -> bool`
- `getOscillatorFrequency() -> int`
- `getPWM(num: int, off: bool?) -> int`
- `readPrescale() -> int`
- `reset()`
- `setExtClk(prescale: int)`
- `setOscillatorFrequency(freq: int)`
- `setOutputMode(totempole: bool)`
- `setPWM(num: int, on: int, off: int) -> int`
- `setPWMFreq(freq: float)`
- `setPin(num: int, val: int, invert: bool?)`
- `sleep()`
- `wakeup()`
- `writeMicroseconds(num: int, microseconds: int)`

## `Adafruit SHT31 Library`

- include: `Adafruit_SHT31.h`
- takes over: i2c

### `Adafruit_SHT31`

- new: `Adafruit_SHT31() -> Adafruit_SHT31`
- `begin(i2caddr: int?) -> bool`
- `heater(h: bool)`
- `isHeaterEnabled() -> bool`
- `readHumidity() -> float`
- `readStatus() -> int`
- `readTemperature() -> float`
- `reset()`

## `Adafruit SSD1306`

- include: `Adafruit_SSD1306.h`

### `Adafruit_SSD1306`

- new: `Adafruit_SSD1306(w: int, h: int, wire_or_mosi: any?, rst_or_clk: any?, dc: int?, rst: int?, cs: int?) -> Adafruit_SSD1306`
- `begin(vcs: int?, addr: int?, reset: bool?, periphBegin: bool?) -> bool`
- `clearDisplay()`
- `dim(dim: bool)`
- `display()`
- `drawCircle(x0: int, y0: int, r: int, color: int)`
- `drawFastHLine(x: int, y: int, w: int, color: int)`
- `drawFastVLine(x: int, y: int, h: int, color: int)`
- `drawLine(x0: int, y0: int, x1: int, y1: int, color: int)`
- `drawPixel(x: int, y: int, color: int)`
- `drawRect(x: int, y: int, w: int, h: int, color: int)`
- `drawRoundRect(x: int, y: int, w: int, h: int, r: int, color: int)`
- `drawTriangle(x0: int, y0: int, x1: int, y1: int, x2: int, y2: int, color: int)`
- `fillCircle(x0: int, y0: int, r: int, color: int)`
- `fillRect(x: int, y: int, w: int, h: int, color: int)`
- `fillRoundRect(x: int, y: int, w: int, h: int, r: int, color: int)`
- `fillScreen(color: int)`
- `fillTriangle(x0: int, y0: int, x1: int, y1: int, x2: int, y2: int, color: int)`
- `getRotation() -> int`
- `height() -> int`
- `invertDisplay(i: bool)`
- `print(value: any, base: any?) -> int`
- `println(value: any?) -> int`
- `setCursor(x: int, y: int)`
- `setRotation(rot: int)`
- `setTextColor(color: int, bg: int?)`
- `setTextSize(s: int, sy: int?)`
- `setTextWrap(wrap: bool)`
- `ssd1306_command(c: int)`
- `width() -> int`

## `Adafruit ST7735 and ST7789 Library`

- include: `Adafruit_ST7735.h`
- takes over: spi

### `Adafruit_ST7735`

- new: `Adafruit_ST7735(cs: int, dc: int, rst: int) -> Adafruit_ST7735`
- new: `Adafruit_ST7735(cs: int, dc: int, mosi: int, sclk: int, rst: int) -> Adafruit_ST7735`
- `drawCircle(x0: int, y0: int, r: int, color: int)`
- `drawLine(x0: int, y0: int, x1: int, y1: int, color: int)`
- `drawPixel(x: int, y: int, color: int)`
- `drawRect(x: int, y: int, w: int, h: int, color: int)`
- `drawRoundRect(x: int, y: int, w: int, h: int, r: int, color: int)`
- `drawTriangle(x0: int, y0: int, x1: int, y1: int, x2: int, y2: int, color: int)`
- `fillCircle(x0: int, y0: int, r: int, color: int)`
- `fillRect(x: int, y: int, w: int, h: int, color: int)`
- `fillRoundRect(x: int, y: int, w: int, h: int, r: int, color: int)`
- `fillScreen(color: int)`
- `fillTriangle(x0: int, y0: int, x1: int, y1: int, x2: int, y2: int, color: int)`
- `getRotation() -> int`
- `height() -> int`
- `initB(spi: any?)`
- `initR(options: int?, spi: any?)`
- `invertDisplay(invert: bool)`
- `print(value: any, base: any?) -> int`
- `println(value: any?) -> int`
- `setAddrWindow(x: int, y: int, w: int, h: int)`
- `setCursor(x: int, y: int)`
- `setFont(font: any)`
- `setRotation(rot: int)`
- `setTextColor(color: int, bg: int?)`
- `setTextSize(s: int, sy: int?)`
- `setTextWrap(wrap: bool)`
- `width() -> int`

## `Adafruit Si7021 Library`

- include: `Adafruit_Si7021.h`
- takes over: i2c

### `Adafruit_Si7021`

- new: `Adafruit_Si7021() -> Adafruit_Si7021`
- `begin() -> bool`
- `getRevision() -> int`
- `heater(h: bool)`
- `isHeaterEnabled() -> bool`
- `readHumidity() -> float`
- `readSerialNumber()`
- `readTemperature() -> float`
- `reset()`
- `setHeatLevel(level: int)`
- `sernum_a: int`
- `sernum_b: int`

## `ArduinoJson`

- include: `ArduinoJson.h`

### `DynamicJsonDocument`

- new: `DynamicJsonDocument(capacity: int) -> DynamicJsonDocument`
- `capacity() -> int`
- `clear()`
- `isNull() -> bool`
- `overflowed() -> bool`
- `size() -> int`

### `JsonDocument`

- new: `JsonDocument() -> JsonDocument`
- `clear()`
- `isNull() -> bool`
- `nesting() -> int`
- `overflowed() -> bool`
- `shrinkToFit()`
- `size() -> int`

- `JSON_ARRAY_SIZE(elements: int) -> int`
- `JSON_OBJECT_SIZE(members: int) -> int`
- `JSON_STRING_SIZE(length: int) -> int`
- `deserializeJson(doc: any, input: any) -> any`
- `deserializeMsgPack(doc: any, input: any) -> any`
- `measureJson(doc: any) -> int`
- `measureJsonPretty(doc: any) -> int`
- `serializeJson(doc: any, output: any) -> int`
- `serializeJsonPretty(doc: any, output: any) -> int`

## `BH1750`

- include: `BH1750.h`
- architectures: avr, sam, esp8266, esp32, stm32
- takes over: i2c

### `BH1750`

- new: `BH1750(addr: int?) -> BH1750`
- `begin() -> bool`
- `measurementReady(maxWait: bool?) -> bool`
- `readLightLevel() -> float`
- `setMTreg(MTreg: int) -> bool`

## `Bounce2`

- include: `Bounce2.h`

### `Bounce`

- new: `Bounce() -> Bounce`
- `attach(pin: int, interval: int?)`
- `changed() -> bool`
- `currentDuration() -> int`
- `duration() -> int`
- `fell() -> bool`
- `interval(interval: int)`
- `previousDuration() -> int`
- `read() -> int`
- `rose() -> bool`
- `update() -> bool`

## `CapacitiveSensor`

- include: `CapacitiveSensor.h`

### `CapacitiveSensor`

- new: `CapacitiveSensor(sendPin: int, receivePin: int) -> CapacitiveSensor`
- `capacitiveSensor(samples: int) -> int`
- `capacitiveSensorRaw(samples: int) -> int`
- `reset_CS_AutocaL()`
- `set_CS_AutocaL_Millis(ms: int)`
- `set_CS_Timeout_Millis(ms: int)`

## `DHT sensor library`

- include: `DHT.h`

### `DHT`

- new: `DHT(pin: int, type: int, count: int?) -> DHT`
- `begin(usec: int?)`
- `computeHeatIndex(temperature: float, percent_humidity: float, fahrenheit: bool?) -> float`
- `convertCtoF(c: float) -> float`
- `convertFtoC(f: float) -> float`
- `read(force: bool?) -> bool`
- `readHumidity(force: bool?) -> float`
- `readTemperature(fahrenheit: bool?, force: bool?) -> float`

## `DallasTemperature`

- include: `DallasTemperature.h`

### `DallasTemperature`

- new: `DallasTemperature(wire: any) -> DallasTemperature`
- `alarmSearch(addr: any) -> bool`
- `begin()`
- `getAddress(addr: any, index: int) -> bool`
- `getDS18Count() -> int`
- `getDeviceCount() -> int`
- `getResolution(addr: any?) -> int`
- `getTempC(addr: any) -> float`
- `getTempCByIndex(index: int) -> float`
- `getTempF(addr: any) -> float`
- `getTempFByIndex(index: int) -> float`
- `getUserData(addr: any) -> int`
- `getWaitForConversion() -> bool`
- `hasAlarm(addr: any?) -> bool`
- `isConnected(addr: any) -> bool`
- `isParasitePowerMode() -> bool`
- `processAlarms()`
- `requestTemperatures() -> bool`
- `requestTemperaturesByAddress(addr: any) -> bool`
- `requestTemperaturesByIndex(index: int) -> bool`
- `resetAlarmSearch()`
- `setCheckForConversion(flag: bool)`
- `setHighAlarmTemp(addr: any, celsius: int)`
- `setLowAlarmTemp(addr: any, celsius: int)`
- `setResolution(res: any, addr: any?)`
- `setUserData(addr: any, data: int)`
- `setWaitForConversion(flag: bool)`
- `toCelsius(fahrenheit: float) -> float`
- `toFahrenheit(celsius: float) -> float`

## `ESP32Servo`

- include: `ESP32Servo.h`
- architectures: esp32
- takes over: servo

### `Servo`

- new: `Servo() -> Servo`
- `attach(pin: int, min: int?, max: int?, channel: int?) -> int`
- `attached() -> bool`
- `detach()`
- `read() -> int`
- `readMicroseconds() -> int`
- `readTicks() -> int`
- `setPeriodHertz(hertz: int)`
- `setTimerWidth(bits: int)`
- `write(value: int)`
- `writeMicroseconds(value: int)`
- `writeTicks(value: int)`

## `Encoder`

- include: `Encoder.h`

### `Encoder`

- new: `Encoder(pin1: int, pin2: int) -> Encoder`
- `read() -> int`
- `readAndReset() -> int`
- `write(p: int)`

## `FastLED`

- include: `FastLED.h`

### `CFastLED`

- new: `CFastLED() -> CFastLED`
- `clear(writeData: bool?)`
- `clearData()`
- `count() -> int`
- `delay(ms: int)`
- `getBrightness() -> int`
- `getFPS() -> int`
- `setBrightness(scale: int)`
- `setCorrection(correction: any)`
- `setDither(dither: int)`
- `setMaxPowerInVoltsAndMilliamps(volts: int, milliamps: int)`
- `setMaxRefreshRate(rate: int, constrain: bool?)`
- `setTemperature(temperature: any)`
- `show(scale: int?)`
- `showColor(color: CRGB, scale: int?)`
- `size() -> int`

### `CHSV`

- new: `CHSV(h: int, s: int, v: int) -> CHSV`
- new: `CHSV() -> CHSV`
- `h: int`
- `hue: int`
- `s: int`
- `sat: int`
- `v: int`
- `val: int`

### `CRGB`

- new: `CRGB(r: int, g: int, b: int) -> CRGB`
- new: `CRGB() -> CRGB`
- new: `CRGB(color: int) -> CRGB`
- `fadeLightBy(fadefactor: int)`
- `fadeToBlackBy(fadefactor: int)`
- `getAverageLight() -> int`
- `getLuma() -> int`
- `getParity() -> int`
- `lerp8(other: CRGB, frac: int) -> CRGB`
- `maximizeBrightness()`
- `nscale8(scaledown: int)`
- `setHSV(h: int, s: int, v: int) -> CRGB`
- `setHue(hue: int) -> CRGB`
- `setRGB(r: int, g: int, b: int) -> CRGB`
- `b: int`
- `blue: int`
- `g: int`
- `green: int`
- `r: int`
- `red: int`

- `ColorFromPalette(pal: any, index: int, brightness: int?, blend: any?) -> CRGB`
- `beat16(bpm: int) -> int`
- `beat8(bpm: int) -> int`
- `beatsin16(bpm: int, lo: int?, hi: int?) -> int`
- `beatsin8(bpm: int, lo: int?, hi: int?) -> int`
- `cos8(theta: int) -> int`
- `fadeToBlackBy(leds: any, num: int, fadefactor: int)`
- `fill_gradient_RGB(leds: any, num: int, c1: CRGB, c2: CRGB)`
- `fill_rainbow(leds: any, num: int, firsthue: int, deltahue: int?)`
- `fill_solid(leds: any, num: int, color: CRGB)`
- `hsv2rgb_rainbow(hsv: CHSV, rgb: CRGB)`
- `map8(x: int, low: int, high: int) -> int`
- `random16(lim: int?) -> int`
- `random8(lim: int?) -> int`
- `scale8(i: int, scale: int) -> int`
- `sin8(theta: int) -> int`

## `HTTPClient`

- include: `HTTPClient.h`
- architectures: esp32

### `HTTPClient`

- new: `HTTPClient() -> HTTPClient`
- `GET() -> int`
- `PATCH(payload: any) -> int`
- `POST(payload: any) -> int`
- `PUT(payload: any) -> int`
- `addHeader(name: text, value: text, first: bool?, replace: bool?)`
- `begin(client_or_url: any, url: text?) -> bool`
- `collectHeaders(keys: any, count: int)`
- `connected() -> bool`
- `end()`
- `getLocation() -> text`
- `getSize() -> int`
- `getStream() -> any`
- `getString() -> text`
- `hasHeader(name: text) -> bool`
- `header(name: text) -> text`
- `sendRequest(type: text, payload: any?) -> int`
- `setAuthorization(user: text, password: text?)`
- `setFollowRedirects(mode: any)`
- `setRedirectLimit(limit: int)`
- `setReuse(reuse: bool)`
- `setTimeout(ms: int)`
- `setUserAgent(ua: text)`
- `writeToStream(stream: any) -> int`

## `HX711`

- include: `HX711.h`

### `HX711`

- new: `HX711() -> HX711`
- `begin(dataPin: int, clockPin: int, fastProcessor: bool?, doReset: bool?)`
- `calibrate_scale(weight: float, times: int?)`
- `get_gain() -> int`
- `get_mode() -> int`
- `get_offset() -> int`
- `get_price(times: int?) -> float`
- `get_rate() -> int`
- `get_scale() -> float`
- `get_tare() -> float`
- `get_unit_price() -> float`
- `get_units(times: int?) -> float`
- `get_value(times: int?) -> float`
- `is_ready() -> bool`
- `last_time_read() -> int`
- `power_down()`
- `power_up()`
- `read() -> float`
- `read_average(times: int?) -> float`
- `read_medavg(times: int?) -> float`
- `read_median(times: int?) -> float`
- `read_runavg(times: int?, alpha: float?) -> float`
- `reset()`
- `set_average_mode()`
- `set_gain(gain: int?, forced: bool?) -> bool`
- `set_medavg_mode()`
- `set_median_mode()`
- `set_offset(offset: int?)`
- `set_rate_10SPS()`
- `set_rate_80SPS()`
- `set_rate_pin(pin: int)`
- `set_raw_mode()`
- `set_runavg_mode()`
- `set_scale(scale: float?) -> bool`
- `set_unit_price(price: float?)`
- `tare(times: int?)`
- `tare_set() -> bool`
- `wait_ready(ms: int?)`
- `wait_ready_retry(retries: int?, ms: int?) -> bool`
- `wait_ready_timeout(timeout: int?, ms: int?) -> bool`

## `IRremote`

- include: `IRremote.h`

### `IRrecv`

- new: `IRrecv(pin: int, bufsize: int?) -> IRrecv`
- `blink13(enable: bool)`
- `decode(results: any?) -> bool`
- `disableIRIn()`
- `enableIRIn()`
- `isIdle() -> bool`
- `resume()`

### `IRsend`

- new: `IRsend(pin: int?) -> IRsend`
- `begin(pin: int, enable_led: bool?, led_pin: int?)`
- `enableIROut(khz: int)`
- `sendDenon(data: int, nbits: int)`
- `sendJVC(data: int, nbits: int, repeat: bool)`
- `sendLG(data: int, nbits: int)`
- `sendNEC(data: int, nbits: int, repeat: bool?)`
- `sendPanasonic(addr: int, data: int, nbits: int)`
- `sendRC5(data: int, nbits: int)`
- `sendRaw(buf: any, len: int, khz: int)`
- `sendSAMSUNG(data: int, nbits: int)`
- `sendSony(data: int, nbits: int, repeat: int?)`

### `IrReceiver`

- `available() -> bool`
- `begin(pin: int, enable_led: bool?, led_pin: int?, feedback: any?)`
- `decode() -> bool`
- `isIdle() -> bool`
- `printIRResultShort(output: any)`
- `resume()`
- `start(micros: int?)`
- `stop()`
- `decodedIRData: any`

### `IrSender`

- `begin(pin: int, enable_led: bool?, led_pin: int?)`
- `sendDenon(addr: int, command: int, repeats: int)`
- `sendJVC(addr: int, command: int, repeats: int)`
- `sendLG(addr: int, command: int, repeats: int)`
- `sendNEC(addr: int, command: int, repeats: int)`
- `sendPanasonic(addr: int, command: int, repeats: int)`
- `sendPulseDistanceWidthFromArray(khz: int, header_mark: int, header_space: int, one_mark: int, one_space: int, zero_mark: int, zero_space: int, data: any, bits: int, protocol: int, flags: int, repeats: int)`
- `sendRC5(addr: int, command: int, repeats: int, toggle: bool?)`
- `sendRC6(addr: int, command: int, repeats: int)`
- `sendRaw(buf: any, len: int, khz: int)`
- `sendSamsung(addr: int, command: int, repeats: int)`
- `sendSharp(addr: int, command: int, repeats: int)`
- `sendSony(addr: int, command: int, repeats: int, bits: int?)`

### `decode_results`

- `address: int`
- `bits: int`
- `command: int`
- `decode_type: int`
- `rawlen: int`
- `value: int`

## `Keypad`

- include: `Keypad.h`

### `Keypad`

- new: `Keypad(keymap: any, rowPins: any, colPins: any, rows: int, cols: int) -> Keypad`
- `addEventListener(listener: any)`
- `getKey() -> int`
- `getKeys() -> bool`
- `getState() -> int`
- `isPressed(keyChar: int) -> bool`
- `keyStateChanged() -> bool`
- `numKeys() -> int`
- `setDebounceTime(time: int)`
- `setHoldTime(time: int)`
- `waitForKey() -> int`
- `key: any`

## `LedControl`

- include: `LedControl.h`
- takes over: spi

### `LedControl`

- new: `LedControl(din: int, clk: int, cs: int, devices: int?) -> LedControl`
- `clearDisplay(addr: int)`
- `getDeviceCount() -> int`
- `setChar(addr: int, digit: int, value: int, dp: bool)`
- `setColumn(addr: int, col: int, value: int)`
- `setDigit(addr: int, digit: int, value: int, dp: bool)`
- `setIntensity(addr: int, intensity: int)`
- `setLed(addr: int, row: int, col: int, state: bool)`
- `setRow(addr: int, row: int, value: int)`
- `setScanLimit(addr: int, limit: int)`
- `shutdown(addr: int, status: bool)`

## `LiquidCrystal`

- include: `LiquidCrystal.h`

### `LiquidCrystal`

- new: `LiquidCrystal(rs: int, enable: int, d4: int, d5: int, d6: int, d7: int) -> LiquidCrystal`
- new: `LiquidCrystal(rs: int, rw: int, enable: int, d0: int, d1: int, d2: int, d3: int, d4: int, d5: int, d6: int, d7: int) -> LiquidCrystal`
- `autoscroll()`
- `begin(cols: int, rows: int, charsize: int?)`
- `blink()`
- `clear()`
- `createChar(location: int, charmap: any)`
- `cursor()`
- `display()`
- `home()`
- `leftToRight()`
- `noAutoscroll()`
- `noBlink()`
- `noCursor()`
- `noDisplay()`
- `print(value: any, base: any?) -> int`
- `println(value: any?) -> int`
- `rightToLeft()`
- `scrollDisplayLeft()`
- `scrollDisplayRight()`
- `setCursor(col: int, row: int)`
- `write(value: int) -> int`

## `LiquidCrystal I2C`

- include: `LiquidCrystal_I2C.h`
- takes over: i2c

### `LiquidCrystal_I2C`

- new: `LiquidCrystal_I2C(addr: int, cols: int, rows: int) -> LiquidCrystal_I2C`
- `autoscroll()`
- `backlight()`
- `begin(cols: int, rows: int, size: int?)`
- `blink()`
- `clear()`
- `createChar(location: int, charmap: any)`
- `cursor()`
- `display()`
- `home()`
- `init()`
- `leftToRight()`
- `noAutoscroll()`
- `noBacklight()`
- `noBlink()`
- `noCursor()`
- `noDisplay()`
- `print(value: any, base: any?) -> int`
- `rightToLeft()`
- `scrollDisplayLeft()`
- `scrollDisplayRight()`
- `setBacklight(value: int)`
- `setCursor(col: int, row: int)`
- `write(value: any) -> int`

## `LoRa`

- include: `LoRa.h`
- takes over: spi

### `LoRa`

- `available() -> int`
- `begin(frequency: int) -> int`
- `beginPacket(implicitHeader: bool?) -> int`
- `disableCrc()`
- `disableInvertIQ()`
- `enableCrc()`
- `enableInvertIQ()`
- `end()`
- `endPacket(async: bool?) -> int`
- `flush()`
- `idle()`
- `packetFrequencyError() -> int`
- `packetRssi() -> int`
- `packetSnr() -> float`
- `parsePacket(size: int?) -> int`
- `peek() -> int`
- `print(value: any, base: any?) -> int`
- `println(value: any?) -> int`
- `random() -> int`
- `read() -> int`
- `receive(size: int?)`
- `rssi() -> int`
- `setCodingRate4(denominator: int)`
- `setFrequency(frequency: int)`
- `setGain(gain: int)`
- `setOCP(mA: int)`
- `setPins(ss: int?, reset: int?, dio0: int?)`
- `setPreambleLength(length: int)`
- `setSPIFrequency(frequency: int)`
- `setSignalBandwidth(sbw: int)`
- `setSpreadingFactor(sf: int)`
- `setSyncWord(sw: int)`
- `setTxPower(level: int, outputPin: int?)`
- `sleep()`
- `write(byte: int) -> int`

## `MFRC522`

- include: `MFRC522.h`
- takes over: spi

### `MFRC522`

- new: `MFRC522(chipSelect: int, reset: int) -> MFRC522`
- new: `MFRC522(ss: int, sck: int, mosi: int, miso: int, irq: int, reset: int) -> MFRC522`
- `MIFARE_Decrement(block: int, delta: int) -> int`
- `MIFARE_GetValue(block: int, value: any) -> int`
- `MIFARE_Increment(block: int, delta: int) -> int`
- `MIFARE_Read(block: int, buffer: any, size: any) -> int`
- `MIFARE_Restore(block: int) -> int`
- `MIFARE_SetValue(block: int, value: int) -> int`
- `MIFARE_Transfer(block: int) -> int`
- `MIFARE_Ultralight_Write(page: int, buffer: any, size: int) -> int`
- `MIFARE_UnbrickUidSector(logging: bool) -> bool`
- `MIFARE_Write(block: int, buffer: any, size: int) -> int`
- `PCD_AntennaOff()`
- `PCD_AntennaOn()`
- `PCD_Authenticate(command: int, block: int, key: any, uid: any) -> int`
- `PCD_DumpVersionToSerial()`
- `PCD_GetAntennaGain() -> int`
- `PCD_Init()`
- `PCD_PerformSelfTest() -> bool`
- `PCD_Reset()`
- `PCD_SetAntennaGain(mask: int)`
- `PCD_StopCrypto1()`
- `PICC_DumpDetailsToSerial(uid: any)`
- `PICC_DumpToSerial(uid: any)`
- `PICC_GetType(sak: int) -> int`
- `PICC_GetTypeName(type: int) -> text`
- `PICC_HaltA() -> int`
- `PICC_IsNewCardPresent() -> bool`
- `PICC_ReadCardSerial() -> bool`
- `PICC_RequestA(buffer: any, size: any) -> int`
- `PICC_Select(uid: any, validBits: int?) -> int`
- `PICC_WakeupA(buffer: any, size: any) -> int`
- `uid: any`

## `NTPClient`

- include: `NTPClient.h`

### `NTPClient`

- new: `NTPClient(udp: any, offset: int?, interval: int?) -> NTPClient`
- `begin()`
- `end()`
- `forceUpdate() -> bool`
- `getDay() -> int`
- `getEpochTime() -> int`
- `getFormattedTime() -> text`
- `getHours() -> int`
- `getMinutes() -> int`
- `getSeconds() -> int`
- `isTimeSet() -> bool`
- `setPoolServerName(server: text)`
- `setRandomPort(min: int, max: int)`
- `setTimeOffset(offset: int)`
- `setUpdateInterval(interval: int)`
- `update() -> bool`

## `NewPing`

- include: `NewPing.h`
- takes over: newping

### `NewPing`

- new: `NewPing(trigger: int, echo: int, max_cm: int?) -> NewPing`
- `check_timer() -> bool`
- `convert_cm(echoTime: int) -> int`
- `convert_in(echoTime: int) -> int`
- `ping(max_cm: int?) -> int`
- `ping_cm(max_cm: int?) -> int`
- `ping_in(max_cm: int?) -> int`
- `ping_median(it: int, max_cm: int?) -> int`
- `ping_timer(callback: any, max_cm: int?)`
- `timer_stop()`
- `timer_us(frequency: int, callback: any)`

## `OneWire`

- include: `OneWire.h`

### `OneWire`

- new: `OneWire(pin: int) -> OneWire`
- `check_crc16(input: any, len: int, inverted_crc: any, crc: int?) -> bool`
- `crc16(input: any, len: int, crc: int?) -> int`
- `crc8(addr: any, len: int) -> int`
- `depower()`
- `read() -> int`
- `read_bit() -> int`
- `read_bytes(buf: any, count: int)`
- `reset() -> int`
- `reset_search()`
- `search(addr: any, search_alarm: bool?) -> bool`
- `select(rom: any)`
- `skip()`
- `target_search(family_code: int)`
- `write(v: int, power: int?)`
- `write_bit(v: int)`
- `write_bytes(buf: any, count: int, power: int?)`

## `PubSubClient`

- include: `PubSubClient.h`

### `PubSubClient`

- new: `PubSubClient(client: any) -> PubSubClient`
- `connect(id: text, user: text?, pass: text?) -> bool`
- `connected() -> bool`
- `disconnect()`
- `loop() -> bool`
- `publish(topic: text, payload: any, retained: bool?) -> bool`
- `publish_P(topic: text, payload: any, retained: bool?) -> bool`
- `setBufferSize(size: int) -> bool`
- `setCallback(callback: any) -> any`
- `setKeepAlive(seconds: int) -> any`
- `setServer(server: any, port: int) -> any`
- `setSocketTimeout(seconds: int) -> any`
- `state() -> int`
- `subscribe(topic: text, qos: int?) -> bool`
- `unsubscribe(topic: text) -> bool`

## `QMC5883LCompass`

- include: `QMC5883LCompass.h`
- takes over: i2c

### `QMC5883LCompass`

- new: `QMC5883LCompass() -> QMC5883LCompass`
- `calibrate()`
- `clearCalibration()`
- `getAzimuth() -> int`
- `getBearing(azimuth: int) -> int`
- `getCalibrationOffset(index: int) -> float`
- `getCalibrationScale(index: int) -> float`
- `getX() -> int`
- `getY() -> int`
- `getZ() -> int`
- `init()`
- `read()`
- `setADDR(b: int)`
- `setCalibration(x_min: int, x_max: int, y_min: int, y_max: int, z_min: int, z_max: int)`
- `setCalibrationOffsets(x_offset: float, y_offset: float, z_offset: float)`
- `setCalibrationScales(x_scale: float, y_scale: float, z_scale: float)`
- `setMagneticDeclination(degrees: int, minutes: int)`
- `setMode(mode: int, odr: int, rng: int, osr: int)`
- `setReset()`
- `setSmoothing(steps: int, adv: bool)`

## `RTClib`

- include: `RTClib.h`
- takes over: i2c

### `DateTime`

- new: `DateTime(year: int, month: int, day: int, hour: int?, minute: int?, second: int?) -> DateTime`
- new: `DateTime(t: int) -> DateTime`
- `day() -> int`
- `dayOfTheWeek() -> int`
- `hour() -> int`
- `isValid() -> bool`
- `minute() -> int`
- `month() -> int`
- `second() -> int`
- `secondstime() -> int`
- `timestamp(format: int?) -> text`
- `toString(buffer: any)`
- `unixtime() -> int`
- `year() -> int`

### `RTC_DS1307`

- new: `RTC_DS1307(wire: any?) -> RTC_DS1307`
- `adjust(dt: DateTime)`
- `begin(wire: any?) -> bool`
- `isrunning() -> bool`
- `now() -> DateTime`
- `readnvram(address: int) -> int`
- `writenvram(address: int, data: int)`

### `RTC_DS3231`

- new: `RTC_DS3231(wire: any?) -> RTC_DS3231`
- `adjust(dt: DateTime)`
- `alarmFired(n: int) -> bool`
- `begin(wire: any?) -> bool`
- `clearAlarm(n: int)`
- `disable32K()`
- `disableAlarm(n: int)`
- `enable32K()`
- `getTemperature() -> float`
- `isEnabled32K() -> bool`
- `lostPower() -> bool`
- `now() -> DateTime`
- `readSqwPinMode() -> int`
- `setAlarm1(dt: DateTime, mode: int) -> bool`
- `setAlarm2(dt: DateTime, mode: int) -> bool`
- `writeSqwPinMode(mode: int)`

### `RTC_Micros`

- new: `RTC_Micros() -> RTC_Micros`
- `adjust(dt: DateTime)`
- `begin(dt: DateTime)`
- `now() -> DateTime`

### `RTC_Millis`

- new: `RTC_Millis() -> RTC_Millis`
- `adjust(dt: DateTime)`
- `begin(dt: DateTime)`
- `now() -> DateTime`

### `RTC_PCF8523`

- new: `RTC_PCF8523(wire: any?) -> RTC_PCF8523`
- `adjust(dt: DateTime)`
- `begin(wire: any?) -> bool`
- `calibrate(mode: int, offset: int)`
- `initialized() -> bool`
- `lostPower() -> bool`
- `now() -> DateTime`

### `RTC_PCF8563`

- new: `RTC_PCF8563(wire: any?) -> RTC_PCF8563`
- `adjust(dt: DateTime)`
- `begin(wire: any?) -> bool`
- `isrunning() -> bool`
- `lostPower() -> bool`
- `now() -> DateTime`

### `TimeSpan`

- new: `TimeSpan(seconds: int) -> TimeSpan`
- new: `TimeSpan(days: int, hours: int, minutes: int, seconds: int) -> TimeSpan`
- `days() -> int`
- `hours() -> int`
- `minutes() -> int`
- `seconds() -> int`
- `totalseconds() -> int`

## `SD`

- include: `SD.h`
- takes over: spi

### `File`

- `available() -> int`
- `close()`
- `flush()`
- `isDirectory() -> bool`
- `name() -> text`
- `openNextFile(mode: any?) -> File`
- `peek() -> int`
- `position() -> int`
- `print(value: any, base: any?) -> int`
- `println(value: any?) -> int`
- `read() -> int`
- `readBytes(buf: any, len: int) -> int`
- `readString() -> text`
- `rewindDirectory()`
- `seek(pos: int) -> bool`
- `size() -> int`
- `write(data: any, len: any?) -> int`

### `SD`

- `begin(cs_pin: int?, spi: any?) -> bool`
- `end()`
- `exists(path: text) -> bool`
- `mkdir(path: text) -> bool`
- `open(path: text, mode: any?) -> File`
- `remove(path: text) -> bool`
- `rmdir(path: text) -> bool`

## `Servo`

- include: `Servo.h`
- takes over: servo

### `Servo`

- new: `Servo() -> Servo`
- `attach(pin: int, min: int?, max: int?) -> int`
- `attached() -> bool`
- `detach()`
- `read() -> int`
- `readMicroseconds() -> int`
- `write(value: int)`
- `writeMicroseconds(value: int)`

## `SoftwareSerial`

- include: `SoftwareSerial.h`

### `SoftwareSerial`

- new: `SoftwareSerial(rx: int, tx: int, inverse: bool?) -> SoftwareSerial`
- `available() -> int`
- `begin(baud: int)`
- `end()`
- `flush()`
- `isListening() -> bool`
- `listen() -> bool`
- `overflow() -> bool`
- `peek() -> int`
- `print(value: any, base: any?) -> int`
- `println(value: any?) -> int`
- `read() -> int`
- `stopListening() -> bool`
- `write(value: int) -> int`

## `Stepper`

- include: `Stepper.h`

### `Stepper`

- new: `Stepper(steps: int, pin1: int, pin2: int, pin3: int?, pin4: int?) -> Stepper`
- `setSpeed(speed: int)`
- `step(steps: int)`
- `version() -> int`

## `TM1637`

- include: `TM1637Display.h`

### `TM1637Display`

- new: `TM1637Display(clk: int, dio: int, bitDelay: int?) -> TM1637Display`
- `clear()`
- `encodeDigit(digit: int) -> int`
- `setBrightness(brightness: int, on: bool?)`
- `setSegments(segments: any, length: int?, pos: int?)`
- `showNumberDec(num: int, leading_zero: bool?, length: int?, pos: int?)`
- `showNumberDecEx(num: int, dots: int?, leading_zero: bool?, length: int?, pos: int?)`
- `showNumberHexEx(num: int, dots: int?, leading_zero: bool?, length: int?, pos: int?)`

## `Time`

- include: `TimeLib.h`

- `adjustTime(adjustment: int)`
- `breakTime(t: int, tm: any)`
- `day(t: int?) -> int`
- `dayShortStr(d: int) -> text`
- `dayStr(d: int) -> text`
- `hour(t: int?) -> int`
- `hourFormat12(t: int?) -> int`
- `isAM(t: int?) -> bool`
- `isPM(t: int?) -> bool`
- `makeTime(tm: any) -> int`
- `minute(t: int?) -> int`
- `month(t: int?) -> int`
- `monthShortStr(m: int) -> text`
- `monthStr(m: int) -> text`
- `now() -> int`
- `second(t: int?) -> int`
- `setSyncInterval(interval: int)`
- `setSyncProvider(fn: any)`
- `setTime(t: int)`
- `timeStatus() -> int`
- `weekday(t: int?) -> int`
- `year(t: int?) -> int`

## `TinyGPSPlus`

- include: `TinyGPSPlus.h`

### `TinyGPSAltitude`

- `age() -> int`
- `feet() -> float`
- `isUpdated() -> bool`
- `isValid() -> bool`
- `km() -> float`
- `meters() -> float`
- `miles() -> float`
- `value() -> float`

### `TinyGPSCourse`

- `age() -> int`
- `deg() -> float`
- `isUpdated() -> bool`
- `isValid() -> bool`
- `value() -> float`

### `TinyGPSDate`

- `age() -> int`
- `day() -> int`
- `isUpdated() -> bool`
- `isValid() -> bool`
- `month() -> int`
- `value() -> int`
- `year() -> int`

### `TinyGPSDecimal`

- `age() -> int`
- `hdop() -> float`
- `isUpdated() -> bool`
- `isValid() -> bool`
- `value() -> int`

### `TinyGPSInteger`

- `age() -> int`
- `isUpdated() -> bool`
- `isValid() -> bool`
- `value() -> int`

### `TinyGPSLocation`

- `age() -> int`
- `isUpdated() -> bool`
- `isValid() -> bool`
- `lat() -> float`
- `lng() -> float`

### `TinyGPSPlus`

- new: `TinyGPSPlus() -> TinyGPSPlus`
- `cardinal(course: float) -> text`
- `charsProcessed() -> int`
- `courseTo(lat1: float, lon1: float, lat2: float, lon2: float) -> float`
- `distanceBetween(lat1: float, lon1: float, lat2: float, lon2: float) -> float`
- `encode(c: int) -> bool`
- `failedChecksum() -> int`
- `passedChecksum() -> int`
- `sentencesWithFix() -> int`
- `altitude: TinyGPSAltitude`
- `course: TinyGPSCourse`
- `date: TinyGPSDate`
- `hdop: TinyGPSDecimal`
- `location: TinyGPSLocation`
- `satellites: TinyGPSInteger`
- `speed: TinyGPSSpeed`
- `time: TinyGPSTime`

### `TinyGPSSpeed`

- `age() -> int`
- `isUpdated() -> bool`
- `isValid() -> bool`
- `kmph() -> float`
- `knots() -> float`
- `mph() -> float`
- `mps() -> float`
- `value() -> float`

### `TinyGPSTime`

- `age() -> int`
- `centisecond() -> int`
- `hour() -> int`
- `isUpdated() -> bool`
- `isValid() -> bool`
- `minute() -> int`
- `second() -> int`
- `value() -> int`

## `U8g2`

- include: `U8g2lib.h`

### `U8G2`

- the classes `U8G2_<display>_F|1|2_<bus>` all extend `U8G2`; a misspelt display is an unknown name
- bus `2ND_4W_HW_SPI` (197 displays, e.g. `U8G2_CH1120_128X160_F_2ND_4W_HW_SPI`): `new(rotation: U8G2Rotation, cs: int, dc: int, reset: int?) -> U8G2`
- bus `2ND_HW_I2C` (120 displays, e.g. `U8G2_CH1120_128X160_F_2ND_HW_I2C`): `new(rotation: U8G2Rotation, reset: int?) -> U8G2`
- bus `2ND_HW_SPI` (6 displays, e.g. `U8G2_ST7920_128X32_F_2ND_HW_SPI`): `new(rotation: U8G2Rotation, cs: int, reset: int?) -> U8G2`
- bus `3W_HW_SPI` (148 displays, e.g. `U8G2_CH1120_128X160_F_3W_HW_SPI`): `new(rotation: U8G2Rotation, cs: int, reset: int?) -> U8G2`
- bus `3W_SW_SPI` (148 displays, e.g. `U8G2_CH1120_128X160_F_3W_SW_SPI`): `new(rotation: U8G2Rotation, clock: int, data: int, cs: int, reset: int?) -> U8G2`
- bus `4W_HW_SPI` (197 displays, e.g. `U8G2_CH1120_128X160_F_4W_HW_SPI`): `new(rotation: U8G2Rotation, cs: int, dc: int, reset: int?) -> U8G2`
- bus `4W_SW_SPI` (197 displays, e.g. `U8G2_CH1120_128X160_F_4W_SW_SPI`): `new(rotation: U8G2Rotation, clock: int, data: int, cs: int, dc: int, reset: int?) -> U8G2`
- bus `6800` (187 displays, e.g. `U8G2_CH1120_128X160_F_6800`): `new(rotation: U8G2Rotation, d0: int, d1: int, d2: int, d3: int, d4: int, d5: int, d6: int, d7: int, enable: int, cs: int, dc: int, reset: int?) -> U8G2`
- bus `8080` (190 displays, e.g. `U8G2_CH1120_128X160_F_8080`): `new(rotation: U8G2Rotation, d0: int, d1: int, d2: int, d3: int, d4: int, d5: int, d6: int, d7: int, enable: int, cs: int, dc: int, reset: int?) -> U8G2`
- bus `HW_I2C` (120 displays, e.g. `U8G2_CH1120_128X160_F_HW_I2C`): `new(rotation: U8G2Rotation, reset: int?, clock: int?, data: int?) -> U8G2`
- bus `HW_SPI` (6 displays, e.g. `U8G2_ST7920_128X32_F_HW_SPI`): `new(rotation: U8G2Rotation, cs: int, reset: int?) -> U8G2`
- bus `SW_I2C` (120 displays, e.g. `U8G2_CH1120_128X160_F_SW_I2C`): `new(rotation: U8G2Rotation, clock: int, data: int, reset: int?) -> U8G2`
- bus `SW_SPI` (6 displays, e.g. `U8G2_ST7920_128X32_F_SW_SPI`): `new(rotation: U8G2Rotation, clock: int, data: int, cs: int, reset: int?) -> U8G2`
- `begin(menu_select_pin: int?, menu_next_pin: int?, menu_prev_pin: int?, menu_up_pin: int?, menu_down_pin: int?, menu_home_pin: int?) -> bool`
- `beginSimple()`
- `clear()`
- `clearBuffer()`
- `clearDisplay()`
- `disableUTF8Print()`
- `display()`
- `drawArc(x0: int, y0: int, rad: int, start: int, end: int)`
- `drawBox(x: int, y: int, w: int, h: int)`
- `drawButtonUTF8(x: int, y: int, flags: int, width: int, padding_h: int, padding_v: int, text: text)`
- `drawCircle(x0: int, y0: int, rad: int, opt: int?)`
- `drawDisc(x0: int, y0: int, rad: int, opt: int?)`
- `drawEllipse(x0: int, y0: int, rx: int, ry: int, opt: int?)`
- `drawFilledEllipse(x0: int, y0: int, rx: int, ry: int, opt: int?)`
- `drawFrame(x: int, y: int, w: int, h: int)`
- `drawGlyph(x: int, y: int, encoding: int) -> int`
- `drawGlyphX2(x: int, y: int, encoding: int) -> int`
- `drawHLine(x: int, y: int, w: int)`
- `drawLine(x1: int, y1: int, x2: int, y2: int)`
- `drawPixel(x: int, y: int)`
- `drawRBox(x: int, y: int, w: int, h: int, r: int)`
- `drawRFrame(x: int, y: int, w: int, h: int, r: int)`
- `drawStr(x: int, y: int, s: text) -> int`
- `drawStrX2(x: int, y: int, s: text) -> int`
- `drawTriangle(x0: int, y0: int, x1: int, y1: int, x2: int, y2: int)`
- `drawUTF8(x: int, y: int, s: text) -> int`
- `drawUTF8X2(x: int, y: int, s: text) -> int`
- `drawVLine(x: int, y: int, h: int)`
- `enableUTF8Print()`
- `firstPage()`
- `getAscent() -> int`
- `getBufferSize() -> int`
- `getColorIndex() -> int`
- `getCols() -> int`
- `getCursorX() -> int`
- `getCursorY() -> int`
- `getDescent() -> int`
- `getDisplayHeight() -> int`
- `getDisplayWidth() -> int`
- `getDrawColor() -> int`
- `getFontAscent() -> int`
- `getFontDescent() -> int`
- `getHeight() -> int`
- `getMaxCharHeight() -> int`
- `getMaxCharWidth() -> int`
- `getRows() -> int`
- `getStrWidth(s: text) -> int`
- `getUTF8Width(s: text) -> int`
- `getWidth() -> int`
- `getXOffsetGlyph(encoding: int) -> int`
- `getXOffsetUTF8(utf8: text) -> int`
- `home()`
- `initDisplay()`
- `nextPage() -> int`
- `noDisplay()`
- `print(value: any, base: any?) -> int`
- `println(value: any?) -> int`
- `refreshDisplay()`
- `sendBuffer()`
- `setAutoPageClear(mode: int)`
- `setBitmapMode(is_transparent: int)`
- `setBusClock(clock_speed: int)`
- `setClipWindow(x0: int, y0: int, x1: int, y1: int)`
- `setColorIndex(color_index: int)`
- `setContrast(value: int)`
- `setCursor(x: int, y: int)`
- `setDisplayRotation(rotation: U8G2Rotation)`
- `setDrawColor(color_index: int)`
- `setFlipMode(mode: int)`
- `setFont(font: U8G2Font)`
- `setFontDirection(dir: int)`
- `setFontMode(is_transparent: int)`
- `setFontPosBaseline()`
- `setFontPosBottom()`
- `setFontPosCenter()`
- `setFontPosTop()`
- `setFontRefHeightAll()`
- `setFontRefHeightExtendedText()`
- `setFontRefHeightText()`
- `setI2CAddress(adr: int)`
- `setMaxClipWindow()`
- `setPowerSave(is_enable: int)`
- `sleepOff()`
- `sleepOn()`
- `updateDisplay()`
- `updateDisplayArea(tx: int, ty: int, tw: int, th: int)`
- `userInterfaceMessage(title1: text, title2: text, title3: text, buttons: text) -> int`
- `userInterfaceSelectionList(title: text, start_pos: int, sl: text) -> int`

## `VL53L1X`

- include: `VL53L1X.h`
- takes over: i2c

### `VL53L1X`

- new: `VL53L1X() -> VL53L1X`
- `dataReady() -> bool`
- `getAddress() -> int`
- `getMeasurementTimingBudget() -> int`
- `getROICenter() -> int`
- `getTimeout() -> int`
- `init(io_2v8: bool?) -> bool`
- `read(blocking: bool?) -> int`
- `readRangeContinuousMillimeters(blocking: bool?) -> int`
- `readRangeSingleMillimeters(blocking: bool?) -> int`
- `readSingle(blocking: bool?) -> int`
- `setAddress(new_addr: int)`
- `setMeasurementTimingBudget(budget_us: int) -> bool`
- `setROICenter(spadNum: int)`
- `setROISize(width: int, height: int)`
- `setTimeout(timeout: int)`
- `startContinuous(period_ms: int)`
- `stopContinuous()`
- `timeoutOccurred() -> bool`

## `WiFi`

- include: `WiFi.h`
- architectures: esp32

### `WiFi`

- `BSSIDstr() -> text`
- `RSSI() -> int`
- `SSID() -> text`
- `begin(ssid: text?, password: text?) -> int`
- `channel() -> int`
- `disconnect(wifioff: bool?, eraseap: bool?) -> bool`
- `dnsIP(n: int?) -> any`
- `gatewayIP() -> any`
- `getHostname() -> text`
- `hostname(name: text) -> bool`
- `localIP() -> any`
- `macAddress() -> text`
- `mode(m: any) -> bool`
- `persistent(enable: bool)`
- `reconnect() -> bool`
- `scanComplete() -> int`
- `scanDelete()`
- `scanNetworks(async: bool?, hidden: bool?) -> int`
- `setSleep(enable: bool) -> bool`
- `softAP(ssid: text, password: text?, channel: int?) -> bool`
- `softAPIP() -> any`
- `softAPdisconnect(wifioff: bool?) -> bool`
- `status() -> int`
- `subnetMask() -> any`

### `WiFiClient`

- new: `WiFiClient() -> WiFiClient`
- `available() -> int`
- `connect(host: text, port: int) -> int`
- `connected(timeout: int?) -> bool`
- `flush()`
- `peek() -> int`
- `print(value: any, base: any?) -> int`
- `println(value: any?) -> int`
- `read() -> int`
- `remoteIP() -> any`
- `remotePort() -> int`
- `setTimeout(seconds: int)`
- `stop()`
- `write(data: any) -> int`

### `WiFiServer`

- new: `WiFiServer(port: int, max_clients: int?) -> WiFiServer`
- `accept() -> WiFiClient`
- `available() -> WiFiClient`
- `begin(port: int?, enable: int?)`
- `close()`
- `end()`
- `hasClient() -> bool`

## `WiFiManager`

- include: `WiFiManager.h`
- architectures: esp8266, esp32

### `WiFiManager`

- new: `WiFiManager() -> WiFiManager`
- `autoConnect(apName: text?, apPassword: text?) -> bool`
- `disconnect() -> bool`
- `getConfigPortalActive() -> bool`
- `getWiFiPass(persistent: bool?) -> text`
- `getWiFiSSID(persistent: bool?) -> text`
- `process() -> bool`
- `resetSettings()`
- `setConfigPortalBlocking(shouldBlock: bool)`
- `setConfigPortalTimeout(seconds: int)`
- `setConnectRetries(numRetries: int)`
- `setConnectTimeout(seconds: int)`
- `setDarkMode(enable: bool)`
- `setDebugOutput(debug: bool)`
- `setHostname(hostname: text) -> bool`
- `setTitle(title: text)`
- `startConfigPortal(apName: text?, apPassword: text?) -> bool`

## `ezButton`

- include: `ezButton.h`

### `ezButton`

- new: `ezButton(pin: int, mode: int?) -> ezButton`
- `getCount() -> int`
- `getState() -> int`
- `getStateRaw() -> int`
- `isPressed() -> bool`
- `isReleased() -> bool`
- `loop()`
- `resetCount()`
- `setCountMode(mode: int)`
- `setDebounceTime(time: int)`

