# Chip registers

`use mujard.regs` gives `regs.<NAME>`: `.read()`, `.write(v)`, `.set(bits)`, `.clear(bits)`, and `.<FIELD>` names a bit-field's mask. Register names are `PERIPHERAL_REGISTER` on SVD chips and the io-header names on AVR. The tables below are generated from the chip makers' description files.

| Chip | Registers | Page |
|---|---|---|
| `atmega2560` | 165 | [regs-atmega2560.md](regs-atmega2560.md) |
| `atmega328p` | 92 | [regs-atmega328p.md](regs-atmega328p.md) |
| `atmega32u4` | 167 | [regs-atmega32u4.md](regs-atmega32u4.md) |
| `esp32` | 2435 | [regs-esp32.md](regs-esp32.md) |
| `esp32c2` | 1098 | [regs-esp32c2.md](regs-esp32c2.md) |
| `esp32c3` | 4477 | [regs-esp32c3.md](regs-esp32c3.md) |
| `esp32c6` | 5743 | [regs-esp32c6.md](regs-esp32c6.md) |
| `esp32h2` | 4400 | [regs-esp32h2.md](regs-esp32h2.md) |
| `esp32s2` | 4917 | [regs-esp32s2.md](regs-esp32s2.md) |
| `esp32s3` | 7949 | [regs-esp32s3.md](regs-esp32s3.md) |
| `esp8266` | 204 | [regs-esp8266.md](regs-esp8266.md) |
| `nrf52840` | 1821 | [regs-nrf52840.md](regs-nrf52840.md) |
| `ra4m1` | 2318 | [regs-ra4m1.md](regs-ra4m1.md) |
| `rp2040` | 1180 | [regs-rp2040.md](regs-rp2040.md) |
| `samd21` | 962 | [regs-samd21.md](regs-samd21.md) |
| `stm32f103` | 722 | [regs-stm32f103.md](regs-stm32f103.md) |
