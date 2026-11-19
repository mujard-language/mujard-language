# Registers - esp32h2

| Name | Address | Bits | Fields | Drives pins |
|---|---|---|---|---|
| `AES_AAD_BLOCK_NUM` | 0x6003A0A0 | 32 | AAD_BLOCK_NUM=0xFFFFFFFF | - |
| `AES_BLOCK_MODE` | 0x6003A094 | 32 | BLOCK_MODE=0x7 | - |
| `AES_BLOCK_NUM` | 0x6003A098 | 32 | BLOCK_NUM=0xFFFFFFFF | - |
| `AES_CONTINUE` | 0x6003A0A8 | 32 | CONTINUE=0x1 | - |
| `AES_DATE` | 0x6003A0B4 | 32 | DATE=0x3FFFFFFF | - |
| `AES_DMA_ENABLE` | 0x6003A090 | 32 | DMA_ENABLE=0x1 | - |
| `AES_DMA_EXIT` | 0x6003A0B8 | 32 | DMA_EXIT=0x1 | - |
| `AES_ENDIAN` | 0x6003A044 | 32 | ENDIAN=0x3F | - |
| `AES_H_MEM[0]` | 0x6003A060 | 8 |  | - |
| `AES_H_MEM[10]` | 0x6003A06A | 8 |  | - |
| `AES_H_MEM[11]` | 0x6003A06B | 8 |  | - |
| `AES_H_MEM[12]` | 0x6003A06C | 8 |  | - |
| `AES_H_MEM[13]` | 0x6003A06D | 8 |  | - |
| `AES_H_MEM[14]` | 0x6003A06E | 8 |  | - |
| `AES_H_MEM[15]` | 0x6003A06F | 8 |  | - |
| `AES_H_MEM[1]` | 0x6003A061 | 8 |  | - |
| `AES_H_MEM[2]` | 0x6003A062 | 8 |  | - |
| `AES_H_MEM[3]` | 0x6003A063 | 8 |  | - |
| `AES_H_MEM[4]` | 0x6003A064 | 8 |  | - |
| `AES_H_MEM[5]` | 0x6003A065 | 8 |  | - |
| `AES_H_MEM[6]` | 0x6003A066 | 8 |  | - |
| `AES_H_MEM[7]` | 0x6003A067 | 8 |  | - |
| `AES_H_MEM[8]` | 0x6003A068 | 8 |  | - |
| `AES_H_MEM[9]` | 0x6003A069 | 8 |  | - |
| `AES_INC_SEL` | 0x6003A09C | 32 | INC_SEL=0x1 | - |
| `AES_INT_CLEAR` | 0x6003A0AC | 32 | INT_CLEAR=0x1 | - |
| `AES_INT_ENA` | 0x6003A0B0 | 32 | INT_ENA=0x1 | - |
| `AES_IV_MEM[0]` | 0x6003A050 | 8 |  | - |
| `AES_IV_MEM[10]` | 0x6003A05A | 8 |  | - |
| `AES_IV_MEM[11]` | 0x6003A05B | 8 |  | - |
| `AES_IV_MEM[12]` | 0x6003A05C | 8 |  | - |
| `AES_IV_MEM[13]` | 0x6003A05D | 8 |  | - |
| `AES_IV_MEM[14]` | 0x6003A05E | 8 |  | - |
| `AES_IV_MEM[15]` | 0x6003A05F | 8 |  | - |
| `AES_IV_MEM[1]` | 0x6003A051 | 8 |  | - |
| `AES_IV_MEM[2]` | 0x6003A052 | 8 |  | - |
| `AES_IV_MEM[3]` | 0x6003A053 | 8 |  | - |
| `AES_IV_MEM[4]` | 0x6003A054 | 8 |  | - |
| `AES_IV_MEM[5]` | 0x6003A055 | 8 |  | - |
| `AES_IV_MEM[6]` | 0x6003A056 | 8 |  | - |
| `AES_IV_MEM[7]` | 0x6003A057 | 8 |  | - |
| `AES_IV_MEM[8]` | 0x6003A058 | 8 |  | - |
| `AES_IV_MEM[9]` | 0x6003A059 | 8 |  | - |
| `AES_J0_MEM[0]` | 0x6003A070 | 8 |  | - |
| `AES_J0_MEM[10]` | 0x6003A07A | 8 |  | - |
| `AES_J0_MEM[11]` | 0x6003A07B | 8 |  | - |
| `AES_J0_MEM[12]` | 0x6003A07C | 8 |  | - |
| `AES_J0_MEM[13]` | 0x6003A07D | 8 |  | - |
| `AES_J0_MEM[14]` | 0x6003A07E | 8 |  | - |
| `AES_J0_MEM[15]` | 0x6003A07F | 8 |  | - |
| `AES_J0_MEM[1]` | 0x6003A071 | 8 |  | - |
| `AES_J0_MEM[2]` | 0x6003A072 | 8 |  | - |
| `AES_J0_MEM[3]` | 0x6003A073 | 8 |  | - |
| `AES_J0_MEM[4]` | 0x6003A074 | 8 |  | - |
| `AES_J0_MEM[5]` | 0x6003A075 | 8 |  | - |
| `AES_J0_MEM[6]` | 0x6003A076 | 8 |  | - |
| `AES_J0_MEM[7]` | 0x6003A077 | 8 |  | - |
| `AES_J0_MEM[8]` | 0x6003A078 | 8 |  | - |
| `AES_J0_MEM[9]` | 0x6003A079 | 8 |  | - |
| `AES_KEY_0` | 0x6003A000 | 32 | KEY_0=0xFFFFFFFF | - |
| `AES_KEY_1` | 0x6003A004 | 32 | KEY_1=0xFFFFFFFF | - |
| `AES_KEY_2` | 0x6003A008 | 32 | KEY_2=0xFFFFFFFF | - |
| `AES_KEY_3` | 0x6003A00C | 32 | KEY_3=0xFFFFFFFF | - |
| `AES_KEY_4` | 0x6003A010 | 32 | KEY_4=0xFFFFFFFF | - |
| `AES_KEY_5` | 0x6003A014 | 32 | KEY_5=0xFFFFFFFF | - |
| `AES_KEY_6` | 0x6003A018 | 32 | KEY_6=0xFFFFFFFF | - |
| `AES_KEY_7` | 0x6003A01C | 32 | KEY_7=0xFFFFFFFF | - |
| `AES_MODE` | 0x6003A040 | 32 | MODE=0x7 | - |
| `AES_REMAINDER_BIT_NUM` | 0x6003A0A4 | 32 | REMAINDER_BIT_NUM=0x7F | - |
| `AES_STATE` | 0x6003A04C | 32 | STATE=0x3 | - |
| `AES_T0_MEM[0]` | 0x6003A080 | 8 |  | - |
| `AES_T0_MEM[10]` | 0x6003A08A | 8 |  | - |
| `AES_T0_MEM[11]` | 0x6003A08B | 8 |  | - |
| `AES_T0_MEM[12]` | 0x6003A08C | 8 |  | - |
| `AES_T0_MEM[13]` | 0x6003A08D | 8 |  | - |
| `AES_T0_MEM[14]` | 0x6003A08E | 8 |  | - |
| `AES_T0_MEM[15]` | 0x6003A08F | 8 |  | - |
| `AES_T0_MEM[1]` | 0x6003A081 | 8 |  | - |
| `AES_T0_MEM[2]` | 0x6003A082 | 8 |  | - |
| `AES_T0_MEM[3]` | 0x6003A083 | 8 |  | - |
| `AES_T0_MEM[4]` | 0x6003A084 | 8 |  | - |
| `AES_T0_MEM[5]` | 0x6003A085 | 8 |  | - |
| `AES_T0_MEM[6]` | 0x6003A086 | 8 |  | - |
| `AES_T0_MEM[7]` | 0x6003A087 | 8 |  | - |
| `AES_T0_MEM[8]` | 0x6003A088 | 8 |  | - |
| `AES_T0_MEM[9]` | 0x6003A089 | 8 |  | - |
| `AES_TEXT_IN_0` | 0x6003A020 | 32 | TEXT_IN_0=0xFFFFFFFF | - |
| `AES_TEXT_IN_1` | 0x6003A024 | 32 | TEXT_IN_1=0xFFFFFFFF | - |
| `AES_TEXT_IN_2` | 0x6003A028 | 32 | TEXT_IN_2=0xFFFFFFFF | - |
| `AES_TEXT_IN_3` | 0x6003A02C | 32 | TEXT_IN_3=0xFFFFFFFF | - |
| `AES_TEXT_OUT_0` | 0x6003A030 | 32 | TEXT_OUT_0=0xFFFFFFFF | - |
| `AES_TEXT_OUT_1` | 0x6003A034 | 32 | TEXT_OUT_1=0xFFFFFFFF | - |
| `AES_TEXT_OUT_2` | 0x6003A038 | 32 | TEXT_OUT_2=0xFFFFFFFF | - |
| `AES_TEXT_OUT_3` | 0x6003A03C | 32 | TEXT_OUT_3=0xFFFFFFFF | - |
| `AES_TRIGGER` | 0x6003A048 | 32 | TRIGGER=0x1 | - |
| `APB_SARADC_APB_TSENS_CTRL` | 0x60040058 | 32 | TSENS_CLK_DIV=0x3FC000 TSENS_IN_INV=0x2000 TSENS_OUT=0xFF TSENS_PU=0x400000 | - |
| `APB_SARADC_APB_TSENS_SAMPLE` | 0x60040068 | 32 | TSENS_SAMPLE_EN=0x10000 TSENS_SAMPLE_RATE=0xFFFF | - |
| `APB_SARADC_APB_TSENS_WAKE` | 0x60040064 | 32 | WAKEUP_EN=0x40000 WAKEUP_MODE=0x20000 WAKEUP_OVER_UPPER_TH=0x10000 WAKEUP_TH_HIGH=0xFF00 WAKEUP_TH_LOW=0xFF | - |
| `APB_SARADC_ARB_CTRL` | 0x60040024 | 32 | ADC_ARB_APB_FORCE=0x4 ADC_ARB_APB_PRIORITY=0xC0 ADC_ARB_FIX_PRIORITY=0x1000 ADC_ARB_GRANT_FORCE=0x20 ADC_ARB_RTC_FORCE=0x8 ADC_ARB_RTC_PRIORITY=0x300 ADC_ARB_WIFI_FORCE=0x10 ADC_ARB_WIFI_PRIORITY=0xC00 | - |
| `APB_SARADC_CALI` | 0x60040060 | 32 | APB_SARADC_CALI_CFG=0x1FFFF | - |
| `APB_SARADC_CLKM_CONF` | 0x60040054 | 32 | CLKM_DIV_A=0xFC000 CLKM_DIV_B=0x3F00 CLKM_DIV_NUM=0xFF CLK_EN=0x100000 CLK_SEL=0x600000 | - |
| `APB_SARADC_CTRL` | 0x60040000 | 32 | SARADC2_PWDET_DRV=0x20000000 SARADC_SAR_CLK_DIV=0x7F80 SARADC_SAR_CLK_GATED=0x40 SARADC_SAR_PATT_LEN=0x38000 SARADC_SAR_PATT_P_CLEAR=0x800000 SARADC_START=0x2 SARADC_START_FORCE=0x1 SARADC_WAIT_ARB_CYCLE=0xC0000000 SARADC_XPD_SAR_FORCE=0x18... | - |
| `APB_SARADC_CTRL2` | 0x60040004 | 32 | SARADC_MAX_MEAS_NUM=0x1FE SARADC_MEAS_NUM_LIMIT=0x1 SARADC_SAR1_INV=0x200 SARADC_SAR2_INV=0x400 SARADC_TIMER_EN=0x1000000 SARADC_TIMER_TARGET=0xFFF000 | - |
| `APB_SARADC_CTRL_DATE` | 0x600403FC | 32 | DATE=0xFFFFFFFF | - |
| `APB_SARADC_DMA_CONF` | 0x60040050 | 32 | APB_ADC_EOF_NUM=0xFFFF APB_ADC_RESET_FSM=0x40000000 APB_ADC_TRANS=0x80000000 | - |
| `APB_SARADC_FILTER_CTRL0` | 0x60040028 | 32 | APB_SARADC_FILTER_CHANNEL0=0x3C00000 APB_SARADC_FILTER_CHANNEL1=0x3C0000 APB_SARADC_FILTER_RESET=0x80000000 | - |
| `APB_SARADC_FILTER_CTRL1` | 0x60040008 | 32 | APB_SARADC_FILTER_FACTOR0=0xE0000000 APB_SARADC_FILTER_FACTOR1=0x1C000000 | - |
| `APB_SARADC_FSM_WAIT` | 0x6004000C | 32 | SARADC_RSTB_WAIT=0xFF00 SARADC_STANDBY_WAIT=0xFF0000 SARADC_XPD_WAIT=0xFF | - |
| `APB_SARADC_INT_CLR` | 0x6004004C | 32 | APB_SARADC1_DONE_INT_CLR=0x80000000 APB_SARADC2_DONE_INT_CLR=0x40000000 APB_SARADC_THRES0_HIGH_INT_CLR=0x20000000 APB_SARADC_THRES0_LOW_INT_CLR=0x8000000 APB_SARADC_THRES1_HIGH_INT_CLR=0x10000000 APB_SARADC_THRES1_LOW_INT_CLR=0x4000000 APB_... | - |
| `APB_SARADC_INT_ENA` | 0x60040040 | 32 | APB_SARADC1_DONE_INT_ENA=0x80000000 APB_SARADC2_DONE_INT_ENA=0x40000000 APB_SARADC_THRES0_HIGH_INT_ENA=0x20000000 APB_SARADC_THRES0_LOW_INT_ENA=0x8000000 APB_SARADC_THRES1_HIGH_INT_ENA=0x10000000 APB_SARADC_THRES1_LOW_INT_ENA=0x4000000 APB_... | - |
| `APB_SARADC_INT_RAW` | 0x60040044 | 32 | APB_SARADC1_DONE_INT_RAW=0x80000000 APB_SARADC2_DONE_INT_RAW=0x40000000 APB_SARADC_THRES0_HIGH_INT_RAW=0x20000000 APB_SARADC_THRES0_LOW_INT_RAW=0x8000000 APB_SARADC_THRES1_HIGH_INT_RAW=0x10000000 APB_SARADC_THRES1_LOW_INT_RAW=0x4000000 APB_... | - |
| `APB_SARADC_INT_ST` | 0x60040048 | 32 | APB_SARADC1_DONE_INT_ST=0x80000000 APB_SARADC2_DONE_INT_ST=0x40000000 APB_SARADC_THRES0_HIGH_INT_ST=0x20000000 APB_SARADC_THRES0_LOW_INT_ST=0x8000000 APB_SARADC_THRES1_HIGH_INT_ST=0x10000000 APB_SARADC_THRES1_LOW_INT_ST=0x4000000 APB_SARADC... | - |
| `APB_SARADC_ONETIME_SAMPLE` | 0x60040020 | 32 | SARADC1_ONETIME_SAMPLE=0x80000000 SARADC2_ONETIME_SAMPLE=0x40000000 SARADC_ONETIME_ATTEN=0x1800000 SARADC_ONETIME_CHANNEL=0x1E000000 SARADC_ONETIME_START=0x20000000 | - |
| `APB_SARADC_SAR1DATA_STATUS` | 0x6004002C | 32 | APB_SARADC1_DATA=0x1FFFF | - |
| `APB_SARADC_SAR1_STATUS` | 0x60040010 | 32 | SARADC_SAR1_STATUS=0xFFFFFFFF | - |
| `APB_SARADC_SAR2DATA_STATUS` | 0x60040030 | 32 | APB_SARADC2_DATA=0x1FFFF | - |
| `APB_SARADC_SAR2_STATUS` | 0x60040014 | 32 | SARADC_SAR2_STATUS=0xFFFFFFFF | - |
| `APB_SARADC_SAR_PATT_TAB1` | 0x60040018 | 32 | SARADC_SAR_PATT_TAB1=0xFFFFFF | - |
| `APB_SARADC_SAR_PATT_TAB2` | 0x6004001C | 32 | SARADC_SAR_PATT_TAB2=0xFFFFFF | - |
| `APB_SARADC_THRES0_CTRL` | 0x60040034 | 32 | APB_SARADC_THRES0_CHANNEL=0xF APB_SARADC_THRES0_HIGH=0x3FFE0 APB_SARADC_THRES0_LOW=0x7FFC0000 | - |
| `APB_SARADC_THRES1_CTRL` | 0x60040038 | 32 | APB_SARADC_THRES1_CHANNEL=0xF APB_SARADC_THRES1_HIGH=0x3FFE0 APB_SARADC_THRES1_LOW=0x7FFC0000 | - |
| `APB_SARADC_THRES_CTRL` | 0x6004003C | 32 | APB_SARADC_THRES0_EN=0x80000000 APB_SARADC_THRES1_EN=0x40000000 APB_SARADC_THRES_ALL_EN=0x8000000 | - |
| `APB_SARADC_TSENS_CTRL2` | 0x6004005C | 32 | TSENS_CLK_INV=0x4000 TSENS_CLK_SEL=0x8000 TSENS_XPD_FORCE=0x3000 TSENS_XPD_WAIT=0xFFF | - |
| `ASSIST_DEBUG_C0RE_0_DEBUG_MODE` | 0x600CE074 | 32 | CORE_0_DEBUG_MODE=0x1 CORE_0_DEBUG_MODULE_ACTIVE=0x2 | - |
| `ASSIST_DEBUG_C0RE_0_INTR_ENA` | 0x600CE000 | 32 | CORE_0_AREA_DRAM0_0_RD_ENA=0x1 CORE_0_AREA_DRAM0_0_WR_ENA=0x2 CORE_0_AREA_DRAM0_1_RD_ENA=0x4 CORE_0_AREA_DRAM0_1_WR_ENA=0x8 CORE_0_AREA_PIF_0_RD_ENA=0x10 CORE_0_AREA_PIF_0_WR_ENA=0x20 CORE_0_AREA_PIF_1_RD_ENA=0x40 CORE_0_AREA_PIF_1_WR_ENA=0... | - |
| `ASSIST_DEBUG_C0RE_0_LASTPC_BEFORE_EXCEPTION` | 0x600CE070 | 32 | CORE_0_LASTPC_BEFORE_EXC=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CLOCK_GATE` | 0x600CE078 | 32 | CLK_EN=0x1 | - |
| `ASSIST_DEBUG_CORE_0_AREA_DRAM0_0_MAX` | 0x600CE014 | 32 | CORE_0_AREA_DRAM0_0_MAX=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_AREA_DRAM0_0_MIN` | 0x600CE010 | 32 | CORE_0_AREA_DRAM0_0_MIN=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_AREA_DRAM0_1_MAX` | 0x600CE01C | 32 | CORE_0_AREA_DRAM0_1_MAX=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_AREA_DRAM0_1_MIN` | 0x600CE018 | 32 | CORE_0_AREA_DRAM0_1_MIN=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_AREA_PC` | 0x600CE030 | 32 | CORE_0_AREA_PC=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_AREA_PIF_0_MAX` | 0x600CE024 | 32 | CORE_0_AREA_PIF_0_MAX=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_AREA_PIF_0_MIN` | 0x600CE020 | 32 | CORE_0_AREA_PIF_0_MIN=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_AREA_PIF_1_MAX` | 0x600CE02C | 32 | CORE_0_AREA_PIF_1_MAX=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_AREA_PIF_1_MIN` | 0x600CE028 | 32 | CORE_0_AREA_PIF_1_MIN=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_AREA_SP` | 0x600CE034 | 32 | CORE_0_AREA_SP=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_DRAM0_EXCEPTION_MONITOR_0` | 0x600CE058 | 32 | CORE_0_DRAM0_RECORDING_ADDR_0=0xFFFFFF CORE_0_DRAM0_RECORDING_BYTEEN_0=0x1E000000 CORE_0_DRAM0_RECORDING_WR_0=0x1000000 | - |
| `ASSIST_DEBUG_CORE_0_DRAM0_EXCEPTION_MONITOR_1` | 0x600CE05C | 32 | CORE_0_DRAM0_RECORDING_PC_0=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_DRAM0_EXCEPTION_MONITOR_2` | 0x600CE060 | 32 | CORE_0_DRAM0_RECORDING_ADDR_1=0xFFFFFF CORE_0_DRAM0_RECORDING_BYTEEN_1=0x1E000000 CORE_0_DRAM0_RECORDING_WR_1=0x1000000 | - |
| `ASSIST_DEBUG_CORE_0_DRAM0_EXCEPTION_MONITOR_3` | 0x600CE064 | 32 | CORE_0_DRAM0_RECORDING_PC_1=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_INTR_CLR` | 0x600CE00C | 32 | CORE_0_AREA_DRAM0_0_RD_CLR=0x1 CORE_0_AREA_DRAM0_0_WR_CLR=0x2 CORE_0_AREA_DRAM0_1_RD_CLR=0x4 CORE_0_AREA_DRAM0_1_WR_CLR=0x8 CORE_0_AREA_PIF_0_RD_CLR=0x10 CORE_0_AREA_PIF_0_WR_CLR=0x20 CORE_0_AREA_PIF_1_RD_CLR=0x40 CORE_0_AREA_PIF_1_WR_CLR=0... | - |
| `ASSIST_DEBUG_CORE_0_INTR_RAW` | 0x600CE004 | 32 | CORE_0_AREA_DRAM0_0_RD_RAW=0x1 CORE_0_AREA_DRAM0_0_WR_RAW=0x2 CORE_0_AREA_DRAM0_1_RD_RAW=0x4 CORE_0_AREA_DRAM0_1_WR_RAW=0x8 CORE_0_AREA_PIF_0_RD_RAW=0x10 CORE_0_AREA_PIF_0_WR_RAW=0x20 CORE_0_AREA_PIF_1_RD_RAW=0x40 CORE_0_AREA_PIF_1_WR_RAW=0... | - |
| `ASSIST_DEBUG_CORE_0_INTR_RLS` | 0x600CE008 | 32 | CORE_0_AREA_DRAM0_0_RD_RLS=0x1 CORE_0_AREA_DRAM0_0_WR_RLS=0x2 CORE_0_AREA_DRAM0_1_RD_RLS=0x4 CORE_0_AREA_DRAM0_1_WR_RLS=0x8 CORE_0_AREA_PIF_0_RD_RLS=0x10 CORE_0_AREA_PIF_0_WR_RLS=0x20 CORE_0_AREA_PIF_1_RD_RLS=0x40 CORE_0_AREA_PIF_1_WR_RLS=0... | - |
| `ASSIST_DEBUG_CORE_0_IRAM0_EXCEPTION_MONITOR_0` | 0x600CE050 | 32 | CORE_0_IRAM0_RECORDING_ADDR_0=0xFFFFFF CORE_0_IRAM0_RECORDING_LOADSTORE_0=0x2000000 CORE_0_IRAM0_RECORDING_WR_0=0x1000000 | - |
| `ASSIST_DEBUG_CORE_0_IRAM0_EXCEPTION_MONITOR_1` | 0x600CE054 | 32 | CORE_0_IRAM0_RECORDING_ADDR_1=0xFFFFFF CORE_0_IRAM0_RECORDING_LOADSTORE_1=0x2000000 CORE_0_IRAM0_RECORDING_WR_1=0x1000000 | - |
| `ASSIST_DEBUG_CORE_0_RCD_EN` | 0x600CE044 | 32 | CORE_0_RCD_PDEBUGEN=0x2 CORE_0_RCD_RECORDEN=0x1 | - |
| `ASSIST_DEBUG_CORE_0_RCD_PDEBUGPC` | 0x600CE048 | 32 | CORE_0_RCD_PDEBUGPC=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_RCD_PDEBUGSP` | 0x600CE04C | 32 | CORE_0_RCD_PDEBUGSP=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_SP_MAX` | 0x600CE03C | 32 | CORE_0_SP_MAX=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_SP_MIN` | 0x600CE038 | 32 | CORE_0_SP_MIN=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_0_SP_PC` | 0x600CE040 | 32 | CORE_0_SP_PC=0xFFFFFFFF | - |
| `ASSIST_DEBUG_CORE_X_IRAM0_DRAM0_EXCEPTION_MONITOR_0` | 0x600CE068 | 32 | CORE_X_IRAM0_DRAM0_LIMIT_CYCLE_0=0xFFFFF | - |
| `ASSIST_DEBUG_CORE_X_IRAM0_DRAM0_EXCEPTION_MONITOR_1` | 0x600CE06C | 32 | CORE_X_IRAM0_DRAM0_LIMIT_CYCLE_1=0xFFFFF | - |
| `ASSIST_DEBUG_DATE` | 0x600CE3FC | 32 | ASSIST_DEBUG_DATE=0xFFFFFFF | - |
| `DS_BOX_MEM[0]` | 0x6008C600 | 8 |  | - |
| `DS_BOX_MEM[10]` | 0x6008C60A | 8 |  | - |
| `DS_BOX_MEM[11]` | 0x6008C60B | 8 |  | - |
| `DS_BOX_MEM[12]` | 0x6008C60C | 8 |  | - |
| `DS_BOX_MEM[13]` | 0x6008C60D | 8 |  | - |
| `DS_BOX_MEM[14]` | 0x6008C60E | 8 |  | - |
| `DS_BOX_MEM[15]` | 0x6008C60F | 8 |  | - |
| `DS_BOX_MEM[16]` | 0x6008C610 | 8 |  | - |
| `DS_BOX_MEM[17]` | 0x6008C611 | 8 |  | - |
| `DS_BOX_MEM[18]` | 0x6008C612 | 8 |  | - |
| `DS_BOX_MEM[19]` | 0x6008C613 | 8 |  | - |
| `DS_BOX_MEM[1]` | 0x6008C601 | 8 |  | - |
| `DS_BOX_MEM[20]` | 0x6008C614 | 8 |  | - |
| `DS_BOX_MEM[21]` | 0x6008C615 | 8 |  | - |
| `DS_BOX_MEM[22]` | 0x6008C616 | 8 |  | - |
| `DS_BOX_MEM[23]` | 0x6008C617 | 8 |  | - |
| `DS_BOX_MEM[24]` | 0x6008C618 | 8 |  | - |
| `DS_BOX_MEM[25]` | 0x6008C619 | 8 |  | - |
| `DS_BOX_MEM[26]` | 0x6008C61A | 8 |  | - |
| `DS_BOX_MEM[27]` | 0x6008C61B | 8 |  | - |
| `DS_BOX_MEM[28]` | 0x6008C61C | 8 |  | - |
| `DS_BOX_MEM[29]` | 0x6008C61D | 8 |  | - |
| `DS_BOX_MEM[2]` | 0x6008C602 | 8 |  | - |
| `DS_BOX_MEM[30]` | 0x6008C61E | 8 |  | - |
| `DS_BOX_MEM[31]` | 0x6008C61F | 8 |  | - |
| `DS_BOX_MEM[32]` | 0x6008C620 | 8 |  | - |
| `DS_BOX_MEM[33]` | 0x6008C621 | 8 |  | - |
| `DS_BOX_MEM[34]` | 0x6008C622 | 8 |  | - |
| `DS_BOX_MEM[35]` | 0x6008C623 | 8 |  | - |
| `DS_BOX_MEM[36]` | 0x6008C624 | 8 |  | - |
| `DS_BOX_MEM[37]` | 0x6008C625 | 8 |  | - |
| `DS_BOX_MEM[38]` | 0x6008C626 | 8 |  | - |
| `DS_BOX_MEM[39]` | 0x6008C627 | 8 |  | - |
| `DS_BOX_MEM[3]` | 0x6008C603 | 8 |  | - |
| `DS_BOX_MEM[40]` | 0x6008C628 | 8 |  | - |
| `DS_BOX_MEM[41]` | 0x6008C629 | 8 |  | - |
| `DS_BOX_MEM[42]` | 0x6008C62A | 8 |  | - |
| `DS_BOX_MEM[43]` | 0x6008C62B | 8 |  | - |
| `DS_BOX_MEM[44]` | 0x6008C62C | 8 |  | - |
| `DS_BOX_MEM[45]` | 0x6008C62D | 8 |  | - |
| `DS_BOX_MEM[46]` | 0x6008C62E | 8 |  | - |
| `DS_BOX_MEM[47]` | 0x6008C62F | 8 |  | - |
| `DS_BOX_MEM[4]` | 0x6008C604 | 8 |  | - |
| `DS_BOX_MEM[5]` | 0x6008C605 | 8 |  | - |
| `DS_BOX_MEM[6]` | 0x6008C606 | 8 |  | - |
| `DS_BOX_MEM[7]` | 0x6008C607 | 8 |  | - |
| `DS_BOX_MEM[8]` | 0x6008C608 | 8 |  | - |
| `DS_BOX_MEM[9]` | 0x6008C609 | 8 |  | - |
| `DS_DATE` | 0x6008CE20 | 32 | DATE=0x3FFFFFFF | - |
| `DS_IV_MEM[0]` | 0x6008C630 | 8 |  | - |
| `DS_IV_MEM[10]` | 0x6008C63A | 8 |  | - |
| `DS_IV_MEM[11]` | 0x6008C63B | 8 |  | - |
| `DS_IV_MEM[12]` | 0x6008C63C | 8 |  | - |
| `DS_IV_MEM[13]` | 0x6008C63D | 8 |  | - |
| `DS_IV_MEM[14]` | 0x6008C63E | 8 |  | - |
| `DS_IV_MEM[15]` | 0x6008C63F | 8 |  | - |
| `DS_IV_MEM[1]` | 0x6008C631 | 8 |  | - |
| `DS_IV_MEM[2]` | 0x6008C632 | 8 |  | - |
| `DS_IV_MEM[3]` | 0x6008C633 | 8 |  | - |
| `DS_IV_MEM[4]` | 0x6008C634 | 8 |  | - |
| `DS_IV_MEM[5]` | 0x6008C635 | 8 |  | - |
| `DS_IV_MEM[6]` | 0x6008C636 | 8 |  | - |
| `DS_IV_MEM[7]` | 0x6008C637 | 8 |  | - |
| `DS_IV_MEM[8]` | 0x6008C638 | 8 |  | - |
| `DS_IV_MEM[9]` | 0x6008C639 | 8 |  | - |
| `DS_M_MEM[0]` | 0x6008C200 | 8 |  | - |
| `DS_M_MEM[100]` | 0x6008C264 | 8 |  | - |
| `DS_M_MEM[101]` | 0x6008C265 | 8 |  | - |
| `DS_M_MEM[102]` | 0x6008C266 | 8 |  | - |
| `DS_M_MEM[103]` | 0x6008C267 | 8 |  | - |
| `DS_M_MEM[104]` | 0x6008C268 | 8 |  | - |
| `DS_M_MEM[105]` | 0x6008C269 | 8 |  | - |
| `DS_M_MEM[106]` | 0x6008C26A | 8 |  | - |
| `DS_M_MEM[107]` | 0x6008C26B | 8 |  | - |
| `DS_M_MEM[108]` | 0x6008C26C | 8 |  | - |
| `DS_M_MEM[109]` | 0x6008C26D | 8 |  | - |
| `DS_M_MEM[10]` | 0x6008C20A | 8 |  | - |
| `DS_M_MEM[110]` | 0x6008C26E | 8 |  | - |
| `DS_M_MEM[111]` | 0x6008C26F | 8 |  | - |
| `DS_M_MEM[112]` | 0x6008C270 | 8 |  | - |
| `DS_M_MEM[113]` | 0x6008C271 | 8 |  | - |
| `DS_M_MEM[114]` | 0x6008C272 | 8 |  | - |
| `DS_M_MEM[115]` | 0x6008C273 | 8 |  | - |
| `DS_M_MEM[116]` | 0x6008C274 | 8 |  | - |
| `DS_M_MEM[117]` | 0x6008C275 | 8 |  | - |
| `DS_M_MEM[118]` | 0x6008C276 | 8 |  | - |
| `DS_M_MEM[119]` | 0x6008C277 | 8 |  | - |
| `DS_M_MEM[11]` | 0x6008C20B | 8 |  | - |
| `DS_M_MEM[120]` | 0x6008C278 | 8 |  | - |
| `DS_M_MEM[121]` | 0x6008C279 | 8 |  | - |
| `DS_M_MEM[122]` | 0x6008C27A | 8 |  | - |
| `DS_M_MEM[123]` | 0x6008C27B | 8 |  | - |
| `DS_M_MEM[124]` | 0x6008C27C | 8 |  | - |
| `DS_M_MEM[125]` | 0x6008C27D | 8 |  | - |
| `DS_M_MEM[126]` | 0x6008C27E | 8 |  | - |
| `DS_M_MEM[127]` | 0x6008C27F | 8 |  | - |
| `DS_M_MEM[128]` | 0x6008C280 | 8 |  | - |
| `DS_M_MEM[129]` | 0x6008C281 | 8 |  | - |
| `DS_M_MEM[12]` | 0x6008C20C | 8 |  | - |
| `DS_M_MEM[130]` | 0x6008C282 | 8 |  | - |
| `DS_M_MEM[131]` | 0x6008C283 | 8 |  | - |
| `DS_M_MEM[132]` | 0x6008C284 | 8 |  | - |
| `DS_M_MEM[133]` | 0x6008C285 | 8 |  | - |
| `DS_M_MEM[134]` | 0x6008C286 | 8 |  | - |
| `DS_M_MEM[135]` | 0x6008C287 | 8 |  | - |
| `DS_M_MEM[136]` | 0x6008C288 | 8 |  | - |
| `DS_M_MEM[137]` | 0x6008C289 | 8 |  | - |
| `DS_M_MEM[138]` | 0x6008C28A | 8 |  | - |
| `DS_M_MEM[139]` | 0x6008C28B | 8 |  | - |
| `DS_M_MEM[13]` | 0x6008C20D | 8 |  | - |
| `DS_M_MEM[140]` | 0x6008C28C | 8 |  | - |
| `DS_M_MEM[141]` | 0x6008C28D | 8 |  | - |
| `DS_M_MEM[142]` | 0x6008C28E | 8 |  | - |
| `DS_M_MEM[143]` | 0x6008C28F | 8 |  | - |
| `DS_M_MEM[144]` | 0x6008C290 | 8 |  | - |
| `DS_M_MEM[145]` | 0x6008C291 | 8 |  | - |
| `DS_M_MEM[146]` | 0x6008C292 | 8 |  | - |
| `DS_M_MEM[147]` | 0x6008C293 | 8 |  | - |
| `DS_M_MEM[148]` | 0x6008C294 | 8 |  | - |
| `DS_M_MEM[149]` | 0x6008C295 | 8 |  | - |
| `DS_M_MEM[14]` | 0x6008C20E | 8 |  | - |
| `DS_M_MEM[150]` | 0x6008C296 | 8 |  | - |
| `DS_M_MEM[151]` | 0x6008C297 | 8 |  | - |
| `DS_M_MEM[152]` | 0x6008C298 | 8 |  | - |
| `DS_M_MEM[153]` | 0x6008C299 | 8 |  | - |
| `DS_M_MEM[154]` | 0x6008C29A | 8 |  | - |
| `DS_M_MEM[155]` | 0x6008C29B | 8 |  | - |
| `DS_M_MEM[156]` | 0x6008C29C | 8 |  | - |
| `DS_M_MEM[157]` | 0x6008C29D | 8 |  | - |
| `DS_M_MEM[158]` | 0x6008C29E | 8 |  | - |
| `DS_M_MEM[159]` | 0x6008C29F | 8 |  | - |
| `DS_M_MEM[15]` | 0x6008C20F | 8 |  | - |
| `DS_M_MEM[160]` | 0x6008C2A0 | 8 |  | - |
| `DS_M_MEM[161]` | 0x6008C2A1 | 8 |  | - |
| `DS_M_MEM[162]` | 0x6008C2A2 | 8 |  | - |
| `DS_M_MEM[163]` | 0x6008C2A3 | 8 |  | - |
| `DS_M_MEM[164]` | 0x6008C2A4 | 8 |  | - |
| `DS_M_MEM[165]` | 0x6008C2A5 | 8 |  | - |
| `DS_M_MEM[166]` | 0x6008C2A6 | 8 |  | - |
| `DS_M_MEM[167]` | 0x6008C2A7 | 8 |  | - |
| `DS_M_MEM[168]` | 0x6008C2A8 | 8 |  | - |
| `DS_M_MEM[169]` | 0x6008C2A9 | 8 |  | - |
| `DS_M_MEM[16]` | 0x6008C210 | 8 |  | - |
| `DS_M_MEM[170]` | 0x6008C2AA | 8 |  | - |
| `DS_M_MEM[171]` | 0x6008C2AB | 8 |  | - |
| `DS_M_MEM[172]` | 0x6008C2AC | 8 |  | - |
| `DS_M_MEM[173]` | 0x6008C2AD | 8 |  | - |
| `DS_M_MEM[174]` | 0x6008C2AE | 8 |  | - |
| `DS_M_MEM[175]` | 0x6008C2AF | 8 |  | - |
| `DS_M_MEM[176]` | 0x6008C2B0 | 8 |  | - |
| `DS_M_MEM[177]` | 0x6008C2B1 | 8 |  | - |
| `DS_M_MEM[178]` | 0x6008C2B2 | 8 |  | - |
| `DS_M_MEM[179]` | 0x6008C2B3 | 8 |  | - |
| `DS_M_MEM[17]` | 0x6008C211 | 8 |  | - |
| `DS_M_MEM[180]` | 0x6008C2B4 | 8 |  | - |
| `DS_M_MEM[181]` | 0x6008C2B5 | 8 |  | - |
| `DS_M_MEM[182]` | 0x6008C2B6 | 8 |  | - |
| `DS_M_MEM[183]` | 0x6008C2B7 | 8 |  | - |
| `DS_M_MEM[184]` | 0x6008C2B8 | 8 |  | - |
| `DS_M_MEM[185]` | 0x6008C2B9 | 8 |  | - |
| `DS_M_MEM[186]` | 0x6008C2BA | 8 |  | - |
| `DS_M_MEM[187]` | 0x6008C2BB | 8 |  | - |
| `DS_M_MEM[188]` | 0x6008C2BC | 8 |  | - |
| `DS_M_MEM[189]` | 0x6008C2BD | 8 |  | - |
| `DS_M_MEM[18]` | 0x6008C212 | 8 |  | - |
| `DS_M_MEM[190]` | 0x6008C2BE | 8 |  | - |
| `DS_M_MEM[191]` | 0x6008C2BF | 8 |  | - |
| `DS_M_MEM[192]` | 0x6008C2C0 | 8 |  | - |
| `DS_M_MEM[193]` | 0x6008C2C1 | 8 |  | - |
| `DS_M_MEM[194]` | 0x6008C2C2 | 8 |  | - |
| `DS_M_MEM[195]` | 0x6008C2C3 | 8 |  | - |
| `DS_M_MEM[196]` | 0x6008C2C4 | 8 |  | - |
| `DS_M_MEM[197]` | 0x6008C2C5 | 8 |  | - |
| `DS_M_MEM[198]` | 0x6008C2C6 | 8 |  | - |
| `DS_M_MEM[199]` | 0x6008C2C7 | 8 |  | - |
| `DS_M_MEM[19]` | 0x6008C213 | 8 |  | - |
| `DS_M_MEM[1]` | 0x6008C201 | 8 |  | - |
| `DS_M_MEM[200]` | 0x6008C2C8 | 8 |  | - |
| `DS_M_MEM[201]` | 0x6008C2C9 | 8 |  | - |
| `DS_M_MEM[202]` | 0x6008C2CA | 8 |  | - |
| `DS_M_MEM[203]` | 0x6008C2CB | 8 |  | - |
| `DS_M_MEM[204]` | 0x6008C2CC | 8 |  | - |
| `DS_M_MEM[205]` | 0x6008C2CD | 8 |  | - |
| `DS_M_MEM[206]` | 0x6008C2CE | 8 |  | - |
| `DS_M_MEM[207]` | 0x6008C2CF | 8 |  | - |
| `DS_M_MEM[208]` | 0x6008C2D0 | 8 |  | - |
| `DS_M_MEM[209]` | 0x6008C2D1 | 8 |  | - |
| `DS_M_MEM[20]` | 0x6008C214 | 8 |  | - |
| `DS_M_MEM[210]` | 0x6008C2D2 | 8 |  | - |
| `DS_M_MEM[211]` | 0x6008C2D3 | 8 |  | - |
| `DS_M_MEM[212]` | 0x6008C2D4 | 8 |  | - |
| `DS_M_MEM[213]` | 0x6008C2D5 | 8 |  | - |
| `DS_M_MEM[214]` | 0x6008C2D6 | 8 |  | - |
| `DS_M_MEM[215]` | 0x6008C2D7 | 8 |  | - |
| `DS_M_MEM[216]` | 0x6008C2D8 | 8 |  | - |
| `DS_M_MEM[217]` | 0x6008C2D9 | 8 |  | - |
| `DS_M_MEM[218]` | 0x6008C2DA | 8 |  | - |
| `DS_M_MEM[219]` | 0x6008C2DB | 8 |  | - |
| `DS_M_MEM[21]` | 0x6008C215 | 8 |  | - |
| `DS_M_MEM[220]` | 0x6008C2DC | 8 |  | - |
| `DS_M_MEM[221]` | 0x6008C2DD | 8 |  | - |
| `DS_M_MEM[222]` | 0x6008C2DE | 8 |  | - |
| `DS_M_MEM[223]` | 0x6008C2DF | 8 |  | - |
| `DS_M_MEM[224]` | 0x6008C2E0 | 8 |  | - |
| `DS_M_MEM[225]` | 0x6008C2E1 | 8 |  | - |
| `DS_M_MEM[226]` | 0x6008C2E2 | 8 |  | - |
| `DS_M_MEM[227]` | 0x6008C2E3 | 8 |  | - |
| `DS_M_MEM[228]` | 0x6008C2E4 | 8 |  | - |
| `DS_M_MEM[229]` | 0x6008C2E5 | 8 |  | - |
| `DS_M_MEM[22]` | 0x6008C216 | 8 |  | - |
| `DS_M_MEM[230]` | 0x6008C2E6 | 8 |  | - |
| `DS_M_MEM[231]` | 0x6008C2E7 | 8 |  | - |
| `DS_M_MEM[232]` | 0x6008C2E8 | 8 |  | - |
| `DS_M_MEM[233]` | 0x6008C2E9 | 8 |  | - |
| `DS_M_MEM[234]` | 0x6008C2EA | 8 |  | - |
| `DS_M_MEM[235]` | 0x6008C2EB | 8 |  | - |
| `DS_M_MEM[236]` | 0x6008C2EC | 8 |  | - |
| `DS_M_MEM[237]` | 0x6008C2ED | 8 |  | - |
| `DS_M_MEM[238]` | 0x6008C2EE | 8 |  | - |
| `DS_M_MEM[239]` | 0x6008C2EF | 8 |  | - |
| `DS_M_MEM[23]` | 0x6008C217 | 8 |  | - |
| `DS_M_MEM[240]` | 0x6008C2F0 | 8 |  | - |
| `DS_M_MEM[241]` | 0x6008C2F1 | 8 |  | - |
| `DS_M_MEM[242]` | 0x6008C2F2 | 8 |  | - |
| `DS_M_MEM[243]` | 0x6008C2F3 | 8 |  | - |
| `DS_M_MEM[244]` | 0x6008C2F4 | 8 |  | - |
| `DS_M_MEM[245]` | 0x6008C2F5 | 8 |  | - |
| `DS_M_MEM[246]` | 0x6008C2F6 | 8 |  | - |
| `DS_M_MEM[247]` | 0x6008C2F7 | 8 |  | - |
| `DS_M_MEM[248]` | 0x6008C2F8 | 8 |  | - |
| `DS_M_MEM[249]` | 0x6008C2F9 | 8 |  | - |
| `DS_M_MEM[24]` | 0x6008C218 | 8 |  | - |
| `DS_M_MEM[250]` | 0x6008C2FA | 8 |  | - |
| `DS_M_MEM[251]` | 0x6008C2FB | 8 |  | - |
| `DS_M_MEM[252]` | 0x6008C2FC | 8 |  | - |
| `DS_M_MEM[253]` | 0x6008C2FD | 8 |  | - |
| `DS_M_MEM[254]` | 0x6008C2FE | 8 |  | - |
| `DS_M_MEM[255]` | 0x6008C2FF | 8 |  | - |
| `DS_M_MEM[256]` | 0x6008C300 | 8 |  | - |
| `DS_M_MEM[257]` | 0x6008C301 | 8 |  | - |
| `DS_M_MEM[258]` | 0x6008C302 | 8 |  | - |
| `DS_M_MEM[259]` | 0x6008C303 | 8 |  | - |
| `DS_M_MEM[25]` | 0x6008C219 | 8 |  | - |
| `DS_M_MEM[260]` | 0x6008C304 | 8 |  | - |
| `DS_M_MEM[261]` | 0x6008C305 | 8 |  | - |
| `DS_M_MEM[262]` | 0x6008C306 | 8 |  | - |
| `DS_M_MEM[263]` | 0x6008C307 | 8 |  | - |
| `DS_M_MEM[264]` | 0x6008C308 | 8 |  | - |
| `DS_M_MEM[265]` | 0x6008C309 | 8 |  | - |
| `DS_M_MEM[266]` | 0x6008C30A | 8 |  | - |
| `DS_M_MEM[267]` | 0x6008C30B | 8 |  | - |
| `DS_M_MEM[268]` | 0x6008C30C | 8 |  | - |
| `DS_M_MEM[269]` | 0x6008C30D | 8 |  | - |
| `DS_M_MEM[26]` | 0x6008C21A | 8 |  | - |
| `DS_M_MEM[270]` | 0x6008C30E | 8 |  | - |
| `DS_M_MEM[271]` | 0x6008C30F | 8 |  | - |
| `DS_M_MEM[272]` | 0x6008C310 | 8 |  | - |
| `DS_M_MEM[273]` | 0x6008C311 | 8 |  | - |
| `DS_M_MEM[274]` | 0x6008C312 | 8 |  | - |
| `DS_M_MEM[275]` | 0x6008C313 | 8 |  | - |
| `DS_M_MEM[276]` | 0x6008C314 | 8 |  | - |
| `DS_M_MEM[277]` | 0x6008C315 | 8 |  | - |
| `DS_M_MEM[278]` | 0x6008C316 | 8 |  | - |
| `DS_M_MEM[279]` | 0x6008C317 | 8 |  | - |
| `DS_M_MEM[27]` | 0x6008C21B | 8 |  | - |
| `DS_M_MEM[280]` | 0x6008C318 | 8 |  | - |
| `DS_M_MEM[281]` | 0x6008C319 | 8 |  | - |
| `DS_M_MEM[282]` | 0x6008C31A | 8 |  | - |
| `DS_M_MEM[283]` | 0x6008C31B | 8 |  | - |
| `DS_M_MEM[284]` | 0x6008C31C | 8 |  | - |
| `DS_M_MEM[285]` | 0x6008C31D | 8 |  | - |
| `DS_M_MEM[286]` | 0x6008C31E | 8 |  | - |
| `DS_M_MEM[287]` | 0x6008C31F | 8 |  | - |
| `DS_M_MEM[288]` | 0x6008C320 | 8 |  | - |
| `DS_M_MEM[289]` | 0x6008C321 | 8 |  | - |
| `DS_M_MEM[28]` | 0x6008C21C | 8 |  | - |
| `DS_M_MEM[290]` | 0x6008C322 | 8 |  | - |
| `DS_M_MEM[291]` | 0x6008C323 | 8 |  | - |
| `DS_M_MEM[292]` | 0x6008C324 | 8 |  | - |
| `DS_M_MEM[293]` | 0x6008C325 | 8 |  | - |
| `DS_M_MEM[294]` | 0x6008C326 | 8 |  | - |
| `DS_M_MEM[295]` | 0x6008C327 | 8 |  | - |
| `DS_M_MEM[296]` | 0x6008C328 | 8 |  | - |
| `DS_M_MEM[297]` | 0x6008C329 | 8 |  | - |
| `DS_M_MEM[298]` | 0x6008C32A | 8 |  | - |
| `DS_M_MEM[299]` | 0x6008C32B | 8 |  | - |
| `DS_M_MEM[29]` | 0x6008C21D | 8 |  | - |
| `DS_M_MEM[2]` | 0x6008C202 | 8 |  | - |
| `DS_M_MEM[300]` | 0x6008C32C | 8 |  | - |
| `DS_M_MEM[301]` | 0x6008C32D | 8 |  | - |
| `DS_M_MEM[302]` | 0x6008C32E | 8 |  | - |
| `DS_M_MEM[303]` | 0x6008C32F | 8 |  | - |
| `DS_M_MEM[304]` | 0x6008C330 | 8 |  | - |
| `DS_M_MEM[305]` | 0x6008C331 | 8 |  | - |
| `DS_M_MEM[306]` | 0x6008C332 | 8 |  | - |
| `DS_M_MEM[307]` | 0x6008C333 | 8 |  | - |
| `DS_M_MEM[308]` | 0x6008C334 | 8 |  | - |
| `DS_M_MEM[309]` | 0x6008C335 | 8 |  | - |
| `DS_M_MEM[30]` | 0x6008C21E | 8 |  | - |
| `DS_M_MEM[310]` | 0x6008C336 | 8 |  | - |
| `DS_M_MEM[311]` | 0x6008C337 | 8 |  | - |
| `DS_M_MEM[312]` | 0x6008C338 | 8 |  | - |
| `DS_M_MEM[313]` | 0x6008C339 | 8 |  | - |
| `DS_M_MEM[314]` | 0x6008C33A | 8 |  | - |
| `DS_M_MEM[315]` | 0x6008C33B | 8 |  | - |
| `DS_M_MEM[316]` | 0x6008C33C | 8 |  | - |
| `DS_M_MEM[317]` | 0x6008C33D | 8 |  | - |
| `DS_M_MEM[318]` | 0x6008C33E | 8 |  | - |
| `DS_M_MEM[319]` | 0x6008C33F | 8 |  | - |
| `DS_M_MEM[31]` | 0x6008C21F | 8 |  | - |
| `DS_M_MEM[320]` | 0x6008C340 | 8 |  | - |
| `DS_M_MEM[321]` | 0x6008C341 | 8 |  | - |
| `DS_M_MEM[322]` | 0x6008C342 | 8 |  | - |
| `DS_M_MEM[323]` | 0x6008C343 | 8 |  | - |
| `DS_M_MEM[324]` | 0x6008C344 | 8 |  | - |
| `DS_M_MEM[325]` | 0x6008C345 | 8 |  | - |
| `DS_M_MEM[326]` | 0x6008C346 | 8 |  | - |
| `DS_M_MEM[327]` | 0x6008C347 | 8 |  | - |
| `DS_M_MEM[328]` | 0x6008C348 | 8 |  | - |
| `DS_M_MEM[329]` | 0x6008C349 | 8 |  | - |
| `DS_M_MEM[32]` | 0x6008C220 | 8 |  | - |
| `DS_M_MEM[330]` | 0x6008C34A | 8 |  | - |
| `DS_M_MEM[331]` | 0x6008C34B | 8 |  | - |
| `DS_M_MEM[332]` | 0x6008C34C | 8 |  | - |
| `DS_M_MEM[333]` | 0x6008C34D | 8 |  | - |
| `DS_M_MEM[334]` | 0x6008C34E | 8 |  | - |
| `DS_M_MEM[335]` | 0x6008C34F | 8 |  | - |
| `DS_M_MEM[336]` | 0x6008C350 | 8 |  | - |
| `DS_M_MEM[337]` | 0x6008C351 | 8 |  | - |
| `DS_M_MEM[338]` | 0x6008C352 | 8 |  | - |
| `DS_M_MEM[339]` | 0x6008C353 | 8 |  | - |
| `DS_M_MEM[33]` | 0x6008C221 | 8 |  | - |
| `DS_M_MEM[340]` | 0x6008C354 | 8 |  | - |
| `DS_M_MEM[341]` | 0x6008C355 | 8 |  | - |
| `DS_M_MEM[342]` | 0x6008C356 | 8 |  | - |
| `DS_M_MEM[343]` | 0x6008C357 | 8 |  | - |
| `DS_M_MEM[344]` | 0x6008C358 | 8 |  | - |
| `DS_M_MEM[345]` | 0x6008C359 | 8 |  | - |
| `DS_M_MEM[346]` | 0x6008C35A | 8 |  | - |
| `DS_M_MEM[347]` | 0x6008C35B | 8 |  | - |
| `DS_M_MEM[348]` | 0x6008C35C | 8 |  | - |
| `DS_M_MEM[349]` | 0x6008C35D | 8 |  | - |
| `DS_M_MEM[34]` | 0x6008C222 | 8 |  | - |
| `DS_M_MEM[350]` | 0x6008C35E | 8 |  | - |
| `DS_M_MEM[351]` | 0x6008C35F | 8 |  | - |
| `DS_M_MEM[352]` | 0x6008C360 | 8 |  | - |
| `DS_M_MEM[353]` | 0x6008C361 | 8 |  | - |
| `DS_M_MEM[354]` | 0x6008C362 | 8 |  | - |
| `DS_M_MEM[355]` | 0x6008C363 | 8 |  | - |
| `DS_M_MEM[356]` | 0x6008C364 | 8 |  | - |
| `DS_M_MEM[357]` | 0x6008C365 | 8 |  | - |
| `DS_M_MEM[358]` | 0x6008C366 | 8 |  | - |
| `DS_M_MEM[359]` | 0x6008C367 | 8 |  | - |
| `DS_M_MEM[35]` | 0x6008C223 | 8 |  | - |
| `DS_M_MEM[360]` | 0x6008C368 | 8 |  | - |
| `DS_M_MEM[361]` | 0x6008C369 | 8 |  | - |
| `DS_M_MEM[362]` | 0x6008C36A | 8 |  | - |
| `DS_M_MEM[363]` | 0x6008C36B | 8 |  | - |
| `DS_M_MEM[364]` | 0x6008C36C | 8 |  | - |
| `DS_M_MEM[365]` | 0x6008C36D | 8 |  | - |
| `DS_M_MEM[366]` | 0x6008C36E | 8 |  | - |
| `DS_M_MEM[367]` | 0x6008C36F | 8 |  | - |
| `DS_M_MEM[368]` | 0x6008C370 | 8 |  | - |
| `DS_M_MEM[369]` | 0x6008C371 | 8 |  | - |
| `DS_M_MEM[36]` | 0x6008C224 | 8 |  | - |
| `DS_M_MEM[370]` | 0x6008C372 | 8 |  | - |
| `DS_M_MEM[371]` | 0x6008C373 | 8 |  | - |
| `DS_M_MEM[372]` | 0x6008C374 | 8 |  | - |
| `DS_M_MEM[373]` | 0x6008C375 | 8 |  | - |
| `DS_M_MEM[374]` | 0x6008C376 | 8 |  | - |
| `DS_M_MEM[375]` | 0x6008C377 | 8 |  | - |
| `DS_M_MEM[376]` | 0x6008C378 | 8 |  | - |
| `DS_M_MEM[377]` | 0x6008C379 | 8 |  | - |
| `DS_M_MEM[378]` | 0x6008C37A | 8 |  | - |
| `DS_M_MEM[379]` | 0x6008C37B | 8 |  | - |
| `DS_M_MEM[37]` | 0x6008C225 | 8 |  | - |
| `DS_M_MEM[380]` | 0x6008C37C | 8 |  | - |
| `DS_M_MEM[381]` | 0x6008C37D | 8 |  | - |
| `DS_M_MEM[382]` | 0x6008C37E | 8 |  | - |
| `DS_M_MEM[383]` | 0x6008C37F | 8 |  | - |
| `DS_M_MEM[384]` | 0x6008C380 | 8 |  | - |
| `DS_M_MEM[385]` | 0x6008C381 | 8 |  | - |
| `DS_M_MEM[386]` | 0x6008C382 | 8 |  | - |
| `DS_M_MEM[387]` | 0x6008C383 | 8 |  | - |
| `DS_M_MEM[388]` | 0x6008C384 | 8 |  | - |
| `DS_M_MEM[389]` | 0x6008C385 | 8 |  | - |
| `DS_M_MEM[38]` | 0x6008C226 | 8 |  | - |
| `DS_M_MEM[390]` | 0x6008C386 | 8 |  | - |
| `DS_M_MEM[391]` | 0x6008C387 | 8 |  | - |
| `DS_M_MEM[392]` | 0x6008C388 | 8 |  | - |
| `DS_M_MEM[393]` | 0x6008C389 | 8 |  | - |
| `DS_M_MEM[394]` | 0x6008C38A | 8 |  | - |
| `DS_M_MEM[395]` | 0x6008C38B | 8 |  | - |
| `DS_M_MEM[396]` | 0x6008C38C | 8 |  | - |
| `DS_M_MEM[397]` | 0x6008C38D | 8 |  | - |
| `DS_M_MEM[398]` | 0x6008C38E | 8 |  | - |
| `DS_M_MEM[399]` | 0x6008C38F | 8 |  | - |
| `DS_M_MEM[39]` | 0x6008C227 | 8 |  | - |
| `DS_M_MEM[3]` | 0x6008C203 | 8 |  | - |
| `DS_M_MEM[400]` | 0x6008C390 | 8 |  | - |
| `DS_M_MEM[401]` | 0x6008C391 | 8 |  | - |
| `DS_M_MEM[402]` | 0x6008C392 | 8 |  | - |
| `DS_M_MEM[403]` | 0x6008C393 | 8 |  | - |
| `DS_M_MEM[404]` | 0x6008C394 | 8 |  | - |
| `DS_M_MEM[405]` | 0x6008C395 | 8 |  | - |
| `DS_M_MEM[406]` | 0x6008C396 | 8 |  | - |
| `DS_M_MEM[407]` | 0x6008C397 | 8 |  | - |
| `DS_M_MEM[408]` | 0x6008C398 | 8 |  | - |
| `DS_M_MEM[409]` | 0x6008C399 | 8 |  | - |
| `DS_M_MEM[40]` | 0x6008C228 | 8 |  | - |
| `DS_M_MEM[410]` | 0x6008C39A | 8 |  | - |
| `DS_M_MEM[411]` | 0x6008C39B | 8 |  | - |
| `DS_M_MEM[412]` | 0x6008C39C | 8 |  | - |
| `DS_M_MEM[413]` | 0x6008C39D | 8 |  | - |
| `DS_M_MEM[414]` | 0x6008C39E | 8 |  | - |
| `DS_M_MEM[415]` | 0x6008C39F | 8 |  | - |
| `DS_M_MEM[416]` | 0x6008C3A0 | 8 |  | - |
| `DS_M_MEM[417]` | 0x6008C3A1 | 8 |  | - |
| `DS_M_MEM[418]` | 0x6008C3A2 | 8 |  | - |
| `DS_M_MEM[419]` | 0x6008C3A3 | 8 |  | - |
| `DS_M_MEM[41]` | 0x6008C229 | 8 |  | - |
| `DS_M_MEM[420]` | 0x6008C3A4 | 8 |  | - |
| `DS_M_MEM[421]` | 0x6008C3A5 | 8 |  | - |
| `DS_M_MEM[422]` | 0x6008C3A6 | 8 |  | - |
| `DS_M_MEM[423]` | 0x6008C3A7 | 8 |  | - |
| `DS_M_MEM[424]` | 0x6008C3A8 | 8 |  | - |
| `DS_M_MEM[425]` | 0x6008C3A9 | 8 |  | - |
| `DS_M_MEM[426]` | 0x6008C3AA | 8 |  | - |
| `DS_M_MEM[427]` | 0x6008C3AB | 8 |  | - |
| `DS_M_MEM[428]` | 0x6008C3AC | 8 |  | - |
| `DS_M_MEM[429]` | 0x6008C3AD | 8 |  | - |
| `DS_M_MEM[42]` | 0x6008C22A | 8 |  | - |
| `DS_M_MEM[430]` | 0x6008C3AE | 8 |  | - |
| `DS_M_MEM[431]` | 0x6008C3AF | 8 |  | - |
| `DS_M_MEM[432]` | 0x6008C3B0 | 8 |  | - |
| `DS_M_MEM[433]` | 0x6008C3B1 | 8 |  | - |
| `DS_M_MEM[434]` | 0x6008C3B2 | 8 |  | - |
| `DS_M_MEM[435]` | 0x6008C3B3 | 8 |  | - |
| `DS_M_MEM[436]` | 0x6008C3B4 | 8 |  | - |
| `DS_M_MEM[437]` | 0x6008C3B5 | 8 |  | - |
| `DS_M_MEM[438]` | 0x6008C3B6 | 8 |  | - |
| `DS_M_MEM[439]` | 0x6008C3B7 | 8 |  | - |
| `DS_M_MEM[43]` | 0x6008C22B | 8 |  | - |
| `DS_M_MEM[440]` | 0x6008C3B8 | 8 |  | - |
| `DS_M_MEM[441]` | 0x6008C3B9 | 8 |  | - |
| `DS_M_MEM[442]` | 0x6008C3BA | 8 |  | - |
| `DS_M_MEM[443]` | 0x6008C3BB | 8 |  | - |
| `DS_M_MEM[444]` | 0x6008C3BC | 8 |  | - |
| `DS_M_MEM[445]` | 0x6008C3BD | 8 |  | - |
| `DS_M_MEM[446]` | 0x6008C3BE | 8 |  | - |
| `DS_M_MEM[447]` | 0x6008C3BF | 8 |  | - |
| `DS_M_MEM[448]` | 0x6008C3C0 | 8 |  | - |
| `DS_M_MEM[449]` | 0x6008C3C1 | 8 |  | - |
| `DS_M_MEM[44]` | 0x6008C22C | 8 |  | - |
| `DS_M_MEM[450]` | 0x6008C3C2 | 8 |  | - |
| `DS_M_MEM[451]` | 0x6008C3C3 | 8 |  | - |
| `DS_M_MEM[452]` | 0x6008C3C4 | 8 |  | - |
| `DS_M_MEM[453]` | 0x6008C3C5 | 8 |  | - |
| `DS_M_MEM[454]` | 0x6008C3C6 | 8 |  | - |
| `DS_M_MEM[455]` | 0x6008C3C7 | 8 |  | - |
| `DS_M_MEM[456]` | 0x6008C3C8 | 8 |  | - |
| `DS_M_MEM[457]` | 0x6008C3C9 | 8 |  | - |
| `DS_M_MEM[458]` | 0x6008C3CA | 8 |  | - |
| `DS_M_MEM[459]` | 0x6008C3CB | 8 |  | - |
| `DS_M_MEM[45]` | 0x6008C22D | 8 |  | - |
| `DS_M_MEM[460]` | 0x6008C3CC | 8 |  | - |
| `DS_M_MEM[461]` | 0x6008C3CD | 8 |  | - |
| `DS_M_MEM[462]` | 0x6008C3CE | 8 |  | - |
| `DS_M_MEM[463]` | 0x6008C3CF | 8 |  | - |
| `DS_M_MEM[464]` | 0x6008C3D0 | 8 |  | - |
| `DS_M_MEM[465]` | 0x6008C3D1 | 8 |  | - |
| `DS_M_MEM[466]` | 0x6008C3D2 | 8 |  | - |
| `DS_M_MEM[467]` | 0x6008C3D3 | 8 |  | - |
| `DS_M_MEM[468]` | 0x6008C3D4 | 8 |  | - |
| `DS_M_MEM[469]` | 0x6008C3D5 | 8 |  | - |
| `DS_M_MEM[46]` | 0x6008C22E | 8 |  | - |
| `DS_M_MEM[470]` | 0x6008C3D6 | 8 |  | - |
| `DS_M_MEM[471]` | 0x6008C3D7 | 8 |  | - |
| `DS_M_MEM[472]` | 0x6008C3D8 | 8 |  | - |
| `DS_M_MEM[473]` | 0x6008C3D9 | 8 |  | - |
| `DS_M_MEM[474]` | 0x6008C3DA | 8 |  | - |
| `DS_M_MEM[475]` | 0x6008C3DB | 8 |  | - |
| `DS_M_MEM[476]` | 0x6008C3DC | 8 |  | - |
| `DS_M_MEM[477]` | 0x6008C3DD | 8 |  | - |
| `DS_M_MEM[478]` | 0x6008C3DE | 8 |  | - |
| `DS_M_MEM[479]` | 0x6008C3DF | 8 |  | - |
| `DS_M_MEM[47]` | 0x6008C22F | 8 |  | - |
| `DS_M_MEM[480]` | 0x6008C3E0 | 8 |  | - |
| `DS_M_MEM[481]` | 0x6008C3E1 | 8 |  | - |
| `DS_M_MEM[482]` | 0x6008C3E2 | 8 |  | - |
| `DS_M_MEM[483]` | 0x6008C3E3 | 8 |  | - |
| `DS_M_MEM[484]` | 0x6008C3E4 | 8 |  | - |
| `DS_M_MEM[485]` | 0x6008C3E5 | 8 |  | - |
| `DS_M_MEM[486]` | 0x6008C3E6 | 8 |  | - |
| `DS_M_MEM[487]` | 0x6008C3E7 | 8 |  | - |
| `DS_M_MEM[488]` | 0x6008C3E8 | 8 |  | - |
| `DS_M_MEM[489]` | 0x6008C3E9 | 8 |  | - |
| `DS_M_MEM[48]` | 0x6008C230 | 8 |  | - |
| `DS_M_MEM[490]` | 0x6008C3EA | 8 |  | - |
| `DS_M_MEM[491]` | 0x6008C3EB | 8 |  | - |
| `DS_M_MEM[492]` | 0x6008C3EC | 8 |  | - |
| `DS_M_MEM[493]` | 0x6008C3ED | 8 |  | - |
| `DS_M_MEM[494]` | 0x6008C3EE | 8 |  | - |
| `DS_M_MEM[495]` | 0x6008C3EF | 8 |  | - |
| `DS_M_MEM[496]` | 0x6008C3F0 | 8 |  | - |
| `DS_M_MEM[497]` | 0x6008C3F1 | 8 |  | - |
| `DS_M_MEM[498]` | 0x6008C3F2 | 8 |  | - |
| `DS_M_MEM[499]` | 0x6008C3F3 | 8 |  | - |
| `DS_M_MEM[49]` | 0x6008C231 | 8 |  | - |
| `DS_M_MEM[4]` | 0x6008C204 | 8 |  | - |
| `DS_M_MEM[500]` | 0x6008C3F4 | 8 |  | - |
| `DS_M_MEM[501]` | 0x6008C3F5 | 8 |  | - |
| `DS_M_MEM[502]` | 0x6008C3F6 | 8 |  | - |
| `DS_M_MEM[503]` | 0x6008C3F7 | 8 |  | - |
| `DS_M_MEM[504]` | 0x6008C3F8 | 8 |  | - |
| `DS_M_MEM[505]` | 0x6008C3F9 | 8 |  | - |
| `DS_M_MEM[506]` | 0x6008C3FA | 8 |  | - |
| `DS_M_MEM[507]` | 0x6008C3FB | 8 |  | - |
| `DS_M_MEM[508]` | 0x6008C3FC | 8 |  | - |
| `DS_M_MEM[509]` | 0x6008C3FD | 8 |  | - |
| `DS_M_MEM[50]` | 0x6008C232 | 8 |  | - |
| `DS_M_MEM[510]` | 0x6008C3FE | 8 |  | - |
| `DS_M_MEM[511]` | 0x6008C3FF | 8 |  | - |
| `DS_M_MEM[51]` | 0x6008C233 | 8 |  | - |
| `DS_M_MEM[52]` | 0x6008C234 | 8 |  | - |
| `DS_M_MEM[53]` | 0x6008C235 | 8 |  | - |
| `DS_M_MEM[54]` | 0x6008C236 | 8 |  | - |
| `DS_M_MEM[55]` | 0x6008C237 | 8 |  | - |
| `DS_M_MEM[56]` | 0x6008C238 | 8 |  | - |
| `DS_M_MEM[57]` | 0x6008C239 | 8 |  | - |
| `DS_M_MEM[58]` | 0x6008C23A | 8 |  | - |
| `DS_M_MEM[59]` | 0x6008C23B | 8 |  | - |
| `DS_M_MEM[5]` | 0x6008C205 | 8 |  | - |
| `DS_M_MEM[60]` | 0x6008C23C | 8 |  | - |
| `DS_M_MEM[61]` | 0x6008C23D | 8 |  | - |
| `DS_M_MEM[62]` | 0x6008C23E | 8 |  | - |
| `DS_M_MEM[63]` | 0x6008C23F | 8 |  | - |
| `DS_M_MEM[64]` | 0x6008C240 | 8 |  | - |
| `DS_M_MEM[65]` | 0x6008C241 | 8 |  | - |
| `DS_M_MEM[66]` | 0x6008C242 | 8 |  | - |
| `DS_M_MEM[67]` | 0x6008C243 | 8 |  | - |
| `DS_M_MEM[68]` | 0x6008C244 | 8 |  | - |
| `DS_M_MEM[69]` | 0x6008C245 | 8 |  | - |
| `DS_M_MEM[6]` | 0x6008C206 | 8 |  | - |
| `DS_M_MEM[70]` | 0x6008C246 | 8 |  | - |
| `DS_M_MEM[71]` | 0x6008C247 | 8 |  | - |
| `DS_M_MEM[72]` | 0x6008C248 | 8 |  | - |
| `DS_M_MEM[73]` | 0x6008C249 | 8 |  | - |
| `DS_M_MEM[74]` | 0x6008C24A | 8 |  | - |
| `DS_M_MEM[75]` | 0x6008C24B | 8 |  | - |
| `DS_M_MEM[76]` | 0x6008C24C | 8 |  | - |
| `DS_M_MEM[77]` | 0x6008C24D | 8 |  | - |
| `DS_M_MEM[78]` | 0x6008C24E | 8 |  | - |
| `DS_M_MEM[79]` | 0x6008C24F | 8 |  | - |
| `DS_M_MEM[7]` | 0x6008C207 | 8 |  | - |
| `DS_M_MEM[80]` | 0x6008C250 | 8 |  | - |
| `DS_M_MEM[81]` | 0x6008C251 | 8 |  | - |
| `DS_M_MEM[82]` | 0x6008C252 | 8 |  | - |
| `DS_M_MEM[83]` | 0x6008C253 | 8 |  | - |
| `DS_M_MEM[84]` | 0x6008C254 | 8 |  | - |
| `DS_M_MEM[85]` | 0x6008C255 | 8 |  | - |
| `DS_M_MEM[86]` | 0x6008C256 | 8 |  | - |
| `DS_M_MEM[87]` | 0x6008C257 | 8 |  | - |
| `DS_M_MEM[88]` | 0x6008C258 | 8 |  | - |
| `DS_M_MEM[89]` | 0x6008C259 | 8 |  | - |
| `DS_M_MEM[8]` | 0x6008C208 | 8 |  | - |
| `DS_M_MEM[90]` | 0x6008C25A | 8 |  | - |
| `DS_M_MEM[91]` | 0x6008C25B | 8 |  | - |
| `DS_M_MEM[92]` | 0x6008C25C | 8 |  | - |
| `DS_M_MEM[93]` | 0x6008C25D | 8 |  | - |
| `DS_M_MEM[94]` | 0x6008C25E | 8 |  | - |
| `DS_M_MEM[95]` | 0x6008C25F | 8 |  | - |
| `DS_M_MEM[96]` | 0x6008C260 | 8 |  | - |
| `DS_M_MEM[97]` | 0x6008C261 | 8 |  | - |
| `DS_M_MEM[98]` | 0x6008C262 | 8 |  | - |
| `DS_M_MEM[99]` | 0x6008C263 | 8 |  | - |
| `DS_M_MEM[9]` | 0x6008C209 | 8 |  | - |
| `DS_QUERY_BUSY` | 0x6008CE0C | 32 | QUERY_BUSY=0x1 | - |
| `DS_QUERY_CHECK` | 0x6008CE14 | 32 | MD_ERROR=0x1 PADDING_BAD=0x2 | - |
| `DS_QUERY_KEY_WRONG` | 0x6008CE10 | 32 | QUERY_KEY_WRONG=0xF | - |
| `DS_RB_MEM[0]` | 0x6008C400 | 8 |  | - |
| `DS_RB_MEM[100]` | 0x6008C464 | 8 |  | - |
| `DS_RB_MEM[101]` | 0x6008C465 | 8 |  | - |
| `DS_RB_MEM[102]` | 0x6008C466 | 8 |  | - |
| `DS_RB_MEM[103]` | 0x6008C467 | 8 |  | - |
| `DS_RB_MEM[104]` | 0x6008C468 | 8 |  | - |
| `DS_RB_MEM[105]` | 0x6008C469 | 8 |  | - |
| `DS_RB_MEM[106]` | 0x6008C46A | 8 |  | - |
| `DS_RB_MEM[107]` | 0x6008C46B | 8 |  | - |
| `DS_RB_MEM[108]` | 0x6008C46C | 8 |  | - |
| `DS_RB_MEM[109]` | 0x6008C46D | 8 |  | - |
| `DS_RB_MEM[10]` | 0x6008C40A | 8 |  | - |
| `DS_RB_MEM[110]` | 0x6008C46E | 8 |  | - |
| `DS_RB_MEM[111]` | 0x6008C46F | 8 |  | - |
| `DS_RB_MEM[112]` | 0x6008C470 | 8 |  | - |
| `DS_RB_MEM[113]` | 0x6008C471 | 8 |  | - |
| `DS_RB_MEM[114]` | 0x6008C472 | 8 |  | - |
| `DS_RB_MEM[115]` | 0x6008C473 | 8 |  | - |
| `DS_RB_MEM[116]` | 0x6008C474 | 8 |  | - |
| `DS_RB_MEM[117]` | 0x6008C475 | 8 |  | - |
| `DS_RB_MEM[118]` | 0x6008C476 | 8 |  | - |
| `DS_RB_MEM[119]` | 0x6008C477 | 8 |  | - |
| `DS_RB_MEM[11]` | 0x6008C40B | 8 |  | - |
| `DS_RB_MEM[120]` | 0x6008C478 | 8 |  | - |
| `DS_RB_MEM[121]` | 0x6008C479 | 8 |  | - |
| `DS_RB_MEM[122]` | 0x6008C47A | 8 |  | - |
| `DS_RB_MEM[123]` | 0x6008C47B | 8 |  | - |
| `DS_RB_MEM[124]` | 0x6008C47C | 8 |  | - |
| `DS_RB_MEM[125]` | 0x6008C47D | 8 |  | - |
| `DS_RB_MEM[126]` | 0x6008C47E | 8 |  | - |
| `DS_RB_MEM[127]` | 0x6008C47F | 8 |  | - |
| `DS_RB_MEM[128]` | 0x6008C480 | 8 |  | - |
| `DS_RB_MEM[129]` | 0x6008C481 | 8 |  | - |
| `DS_RB_MEM[12]` | 0x6008C40C | 8 |  | - |
| `DS_RB_MEM[130]` | 0x6008C482 | 8 |  | - |
| `DS_RB_MEM[131]` | 0x6008C483 | 8 |  | - |
| `DS_RB_MEM[132]` | 0x6008C484 | 8 |  | - |
| `DS_RB_MEM[133]` | 0x6008C485 | 8 |  | - |
| `DS_RB_MEM[134]` | 0x6008C486 | 8 |  | - |
| `DS_RB_MEM[135]` | 0x6008C487 | 8 |  | - |
| `DS_RB_MEM[136]` | 0x6008C488 | 8 |  | - |
| `DS_RB_MEM[137]` | 0x6008C489 | 8 |  | - |
| `DS_RB_MEM[138]` | 0x6008C48A | 8 |  | - |
| `DS_RB_MEM[139]` | 0x6008C48B | 8 |  | - |
| `DS_RB_MEM[13]` | 0x6008C40D | 8 |  | - |
| `DS_RB_MEM[140]` | 0x6008C48C | 8 |  | - |
| `DS_RB_MEM[141]` | 0x6008C48D | 8 |  | - |
| `DS_RB_MEM[142]` | 0x6008C48E | 8 |  | - |
| `DS_RB_MEM[143]` | 0x6008C48F | 8 |  | - |
| `DS_RB_MEM[144]` | 0x6008C490 | 8 |  | - |
| `DS_RB_MEM[145]` | 0x6008C491 | 8 |  | - |
| `DS_RB_MEM[146]` | 0x6008C492 | 8 |  | - |
| `DS_RB_MEM[147]` | 0x6008C493 | 8 |  | - |
| `DS_RB_MEM[148]` | 0x6008C494 | 8 |  | - |
| `DS_RB_MEM[149]` | 0x6008C495 | 8 |  | - |
| `DS_RB_MEM[14]` | 0x6008C40E | 8 |  | - |
| `DS_RB_MEM[150]` | 0x6008C496 | 8 |  | - |
| `DS_RB_MEM[151]` | 0x6008C497 | 8 |  | - |
| `DS_RB_MEM[152]` | 0x6008C498 | 8 |  | - |
| `DS_RB_MEM[153]` | 0x6008C499 | 8 |  | - |
| `DS_RB_MEM[154]` | 0x6008C49A | 8 |  | - |
| `DS_RB_MEM[155]` | 0x6008C49B | 8 |  | - |
| `DS_RB_MEM[156]` | 0x6008C49C | 8 |  | - |
| `DS_RB_MEM[157]` | 0x6008C49D | 8 |  | - |
| `DS_RB_MEM[158]` | 0x6008C49E | 8 |  | - |
| `DS_RB_MEM[159]` | 0x6008C49F | 8 |  | - |
| `DS_RB_MEM[15]` | 0x6008C40F | 8 |  | - |
| `DS_RB_MEM[160]` | 0x6008C4A0 | 8 |  | - |
| `DS_RB_MEM[161]` | 0x6008C4A1 | 8 |  | - |
| `DS_RB_MEM[162]` | 0x6008C4A2 | 8 |  | - |
| `DS_RB_MEM[163]` | 0x6008C4A3 | 8 |  | - |
| `DS_RB_MEM[164]` | 0x6008C4A4 | 8 |  | - |
| `DS_RB_MEM[165]` | 0x6008C4A5 | 8 |  | - |
| `DS_RB_MEM[166]` | 0x6008C4A6 | 8 |  | - |
| `DS_RB_MEM[167]` | 0x6008C4A7 | 8 |  | - |
| `DS_RB_MEM[168]` | 0x6008C4A8 | 8 |  | - |
| `DS_RB_MEM[169]` | 0x6008C4A9 | 8 |  | - |
| `DS_RB_MEM[16]` | 0x6008C410 | 8 |  | - |
| `DS_RB_MEM[170]` | 0x6008C4AA | 8 |  | - |
| `DS_RB_MEM[171]` | 0x6008C4AB | 8 |  | - |
| `DS_RB_MEM[172]` | 0x6008C4AC | 8 |  | - |
| `DS_RB_MEM[173]` | 0x6008C4AD | 8 |  | - |
| `DS_RB_MEM[174]` | 0x6008C4AE | 8 |  | - |
| `DS_RB_MEM[175]` | 0x6008C4AF | 8 |  | - |
| `DS_RB_MEM[176]` | 0x6008C4B0 | 8 |  | - |
| `DS_RB_MEM[177]` | 0x6008C4B1 | 8 |  | - |
| `DS_RB_MEM[178]` | 0x6008C4B2 | 8 |  | - |
| `DS_RB_MEM[179]` | 0x6008C4B3 | 8 |  | - |
| `DS_RB_MEM[17]` | 0x6008C411 | 8 |  | - |
| `DS_RB_MEM[180]` | 0x6008C4B4 | 8 |  | - |
| `DS_RB_MEM[181]` | 0x6008C4B5 | 8 |  | - |
| `DS_RB_MEM[182]` | 0x6008C4B6 | 8 |  | - |
| `DS_RB_MEM[183]` | 0x6008C4B7 | 8 |  | - |
| `DS_RB_MEM[184]` | 0x6008C4B8 | 8 |  | - |
| `DS_RB_MEM[185]` | 0x6008C4B9 | 8 |  | - |
| `DS_RB_MEM[186]` | 0x6008C4BA | 8 |  | - |
| `DS_RB_MEM[187]` | 0x6008C4BB | 8 |  | - |
| `DS_RB_MEM[188]` | 0x6008C4BC | 8 |  | - |
| `DS_RB_MEM[189]` | 0x6008C4BD | 8 |  | - |
| `DS_RB_MEM[18]` | 0x6008C412 | 8 |  | - |
| `DS_RB_MEM[190]` | 0x6008C4BE | 8 |  | - |
| `DS_RB_MEM[191]` | 0x6008C4BF | 8 |  | - |
| `DS_RB_MEM[192]` | 0x6008C4C0 | 8 |  | - |
| `DS_RB_MEM[193]` | 0x6008C4C1 | 8 |  | - |
| `DS_RB_MEM[194]` | 0x6008C4C2 | 8 |  | - |
| `DS_RB_MEM[195]` | 0x6008C4C3 | 8 |  | - |
| `DS_RB_MEM[196]` | 0x6008C4C4 | 8 |  | - |
| `DS_RB_MEM[197]` | 0x6008C4C5 | 8 |  | - |
| `DS_RB_MEM[198]` | 0x6008C4C6 | 8 |  | - |
| `DS_RB_MEM[199]` | 0x6008C4C7 | 8 |  | - |
| `DS_RB_MEM[19]` | 0x6008C413 | 8 |  | - |
| `DS_RB_MEM[1]` | 0x6008C401 | 8 |  | - |
| `DS_RB_MEM[200]` | 0x6008C4C8 | 8 |  | - |
| `DS_RB_MEM[201]` | 0x6008C4C9 | 8 |  | - |
| `DS_RB_MEM[202]` | 0x6008C4CA | 8 |  | - |
| `DS_RB_MEM[203]` | 0x6008C4CB | 8 |  | - |
| `DS_RB_MEM[204]` | 0x6008C4CC | 8 |  | - |
| `DS_RB_MEM[205]` | 0x6008C4CD | 8 |  | - |
| `DS_RB_MEM[206]` | 0x6008C4CE | 8 |  | - |
| `DS_RB_MEM[207]` | 0x6008C4CF | 8 |  | - |
| `DS_RB_MEM[208]` | 0x6008C4D0 | 8 |  | - |
| `DS_RB_MEM[209]` | 0x6008C4D1 | 8 |  | - |
| `DS_RB_MEM[20]` | 0x6008C414 | 8 |  | - |
| `DS_RB_MEM[210]` | 0x6008C4D2 | 8 |  | - |
| `DS_RB_MEM[211]` | 0x6008C4D3 | 8 |  | - |
| `DS_RB_MEM[212]` | 0x6008C4D4 | 8 |  | - |
| `DS_RB_MEM[213]` | 0x6008C4D5 | 8 |  | - |
| `DS_RB_MEM[214]` | 0x6008C4D6 | 8 |  | - |
| `DS_RB_MEM[215]` | 0x6008C4D7 | 8 |  | - |
| `DS_RB_MEM[216]` | 0x6008C4D8 | 8 |  | - |
| `DS_RB_MEM[217]` | 0x6008C4D9 | 8 |  | - |
| `DS_RB_MEM[218]` | 0x6008C4DA | 8 |  | - |
| `DS_RB_MEM[219]` | 0x6008C4DB | 8 |  | - |
| `DS_RB_MEM[21]` | 0x6008C415 | 8 |  | - |
| `DS_RB_MEM[220]` | 0x6008C4DC | 8 |  | - |
| `DS_RB_MEM[221]` | 0x6008C4DD | 8 |  | - |
| `DS_RB_MEM[222]` | 0x6008C4DE | 8 |  | - |
| `DS_RB_MEM[223]` | 0x6008C4DF | 8 |  | - |
| `DS_RB_MEM[224]` | 0x6008C4E0 | 8 |  | - |
| `DS_RB_MEM[225]` | 0x6008C4E1 | 8 |  | - |
| `DS_RB_MEM[226]` | 0x6008C4E2 | 8 |  | - |
| `DS_RB_MEM[227]` | 0x6008C4E3 | 8 |  | - |
| `DS_RB_MEM[228]` | 0x6008C4E4 | 8 |  | - |
| `DS_RB_MEM[229]` | 0x6008C4E5 | 8 |  | - |
| `DS_RB_MEM[22]` | 0x6008C416 | 8 |  | - |
| `DS_RB_MEM[230]` | 0x6008C4E6 | 8 |  | - |
| `DS_RB_MEM[231]` | 0x6008C4E7 | 8 |  | - |
| `DS_RB_MEM[232]` | 0x6008C4E8 | 8 |  | - |
| `DS_RB_MEM[233]` | 0x6008C4E9 | 8 |  | - |
| `DS_RB_MEM[234]` | 0x6008C4EA | 8 |  | - |
| `DS_RB_MEM[235]` | 0x6008C4EB | 8 |  | - |
| `DS_RB_MEM[236]` | 0x6008C4EC | 8 |  | - |
| `DS_RB_MEM[237]` | 0x6008C4ED | 8 |  | - |
| `DS_RB_MEM[238]` | 0x6008C4EE | 8 |  | - |
| `DS_RB_MEM[239]` | 0x6008C4EF | 8 |  | - |
| `DS_RB_MEM[23]` | 0x6008C417 | 8 |  | - |
| `DS_RB_MEM[240]` | 0x6008C4F0 | 8 |  | - |
| `DS_RB_MEM[241]` | 0x6008C4F1 | 8 |  | - |
| `DS_RB_MEM[242]` | 0x6008C4F2 | 8 |  | - |
| `DS_RB_MEM[243]` | 0x6008C4F3 | 8 |  | - |
| `DS_RB_MEM[244]` | 0x6008C4F4 | 8 |  | - |
| `DS_RB_MEM[245]` | 0x6008C4F5 | 8 |  | - |
| `DS_RB_MEM[246]` | 0x6008C4F6 | 8 |  | - |
| `DS_RB_MEM[247]` | 0x6008C4F7 | 8 |  | - |
| `DS_RB_MEM[248]` | 0x6008C4F8 | 8 |  | - |
| `DS_RB_MEM[249]` | 0x6008C4F9 | 8 |  | - |
| `DS_RB_MEM[24]` | 0x6008C418 | 8 |  | - |
| `DS_RB_MEM[250]` | 0x6008C4FA | 8 |  | - |
| `DS_RB_MEM[251]` | 0x6008C4FB | 8 |  | - |
| `DS_RB_MEM[252]` | 0x6008C4FC | 8 |  | - |
| `DS_RB_MEM[253]` | 0x6008C4FD | 8 |  | - |
| `DS_RB_MEM[254]` | 0x6008C4FE | 8 |  | - |
| `DS_RB_MEM[255]` | 0x6008C4FF | 8 |  | - |
| `DS_RB_MEM[256]` | 0x6008C500 | 8 |  | - |
| `DS_RB_MEM[257]` | 0x6008C501 | 8 |  | - |
| `DS_RB_MEM[258]` | 0x6008C502 | 8 |  | - |
| `DS_RB_MEM[259]` | 0x6008C503 | 8 |  | - |
| `DS_RB_MEM[25]` | 0x6008C419 | 8 |  | - |
| `DS_RB_MEM[260]` | 0x6008C504 | 8 |  | - |
| `DS_RB_MEM[261]` | 0x6008C505 | 8 |  | - |
| `DS_RB_MEM[262]` | 0x6008C506 | 8 |  | - |
| `DS_RB_MEM[263]` | 0x6008C507 | 8 |  | - |
| `DS_RB_MEM[264]` | 0x6008C508 | 8 |  | - |
| `DS_RB_MEM[265]` | 0x6008C509 | 8 |  | - |
| `DS_RB_MEM[266]` | 0x6008C50A | 8 |  | - |
| `DS_RB_MEM[267]` | 0x6008C50B | 8 |  | - |
| `DS_RB_MEM[268]` | 0x6008C50C | 8 |  | - |
| `DS_RB_MEM[269]` | 0x6008C50D | 8 |  | - |
| `DS_RB_MEM[26]` | 0x6008C41A | 8 |  | - |
| `DS_RB_MEM[270]` | 0x6008C50E | 8 |  | - |
| `DS_RB_MEM[271]` | 0x6008C50F | 8 |  | - |
| `DS_RB_MEM[272]` | 0x6008C510 | 8 |  | - |
| `DS_RB_MEM[273]` | 0x6008C511 | 8 |  | - |
| `DS_RB_MEM[274]` | 0x6008C512 | 8 |  | - |
| `DS_RB_MEM[275]` | 0x6008C513 | 8 |  | - |
| `DS_RB_MEM[276]` | 0x6008C514 | 8 |  | - |
| `DS_RB_MEM[277]` | 0x6008C515 | 8 |  | - |
| `DS_RB_MEM[278]` | 0x6008C516 | 8 |  | - |
| `DS_RB_MEM[279]` | 0x6008C517 | 8 |  | - |
| `DS_RB_MEM[27]` | 0x6008C41B | 8 |  | - |
| `DS_RB_MEM[280]` | 0x6008C518 | 8 |  | - |
| `DS_RB_MEM[281]` | 0x6008C519 | 8 |  | - |
| `DS_RB_MEM[282]` | 0x6008C51A | 8 |  | - |
| `DS_RB_MEM[283]` | 0x6008C51B | 8 |  | - |
| `DS_RB_MEM[284]` | 0x6008C51C | 8 |  | - |
| `DS_RB_MEM[285]` | 0x6008C51D | 8 |  | - |
| `DS_RB_MEM[286]` | 0x6008C51E | 8 |  | - |
| `DS_RB_MEM[287]` | 0x6008C51F | 8 |  | - |
| `DS_RB_MEM[288]` | 0x6008C520 | 8 |  | - |
| `DS_RB_MEM[289]` | 0x6008C521 | 8 |  | - |
| `DS_RB_MEM[28]` | 0x6008C41C | 8 |  | - |
| `DS_RB_MEM[290]` | 0x6008C522 | 8 |  | - |
| `DS_RB_MEM[291]` | 0x6008C523 | 8 |  | - |
| `DS_RB_MEM[292]` | 0x6008C524 | 8 |  | - |
| `DS_RB_MEM[293]` | 0x6008C525 | 8 |  | - |
| `DS_RB_MEM[294]` | 0x6008C526 | 8 |  | - |
| `DS_RB_MEM[295]` | 0x6008C527 | 8 |  | - |
| `DS_RB_MEM[296]` | 0x6008C528 | 8 |  | - |
| `DS_RB_MEM[297]` | 0x6008C529 | 8 |  | - |
| `DS_RB_MEM[298]` | 0x6008C52A | 8 |  | - |
| `DS_RB_MEM[299]` | 0x6008C52B | 8 |  | - |
| `DS_RB_MEM[29]` | 0x6008C41D | 8 |  | - |
| `DS_RB_MEM[2]` | 0x6008C402 | 8 |  | - |
| `DS_RB_MEM[300]` | 0x6008C52C | 8 |  | - |
| `DS_RB_MEM[301]` | 0x6008C52D | 8 |  | - |
| `DS_RB_MEM[302]` | 0x6008C52E | 8 |  | - |
| `DS_RB_MEM[303]` | 0x6008C52F | 8 |  | - |
| `DS_RB_MEM[304]` | 0x6008C530 | 8 |  | - |
| `DS_RB_MEM[305]` | 0x6008C531 | 8 |  | - |
| `DS_RB_MEM[306]` | 0x6008C532 | 8 |  | - |
| `DS_RB_MEM[307]` | 0x6008C533 | 8 |  | - |
| `DS_RB_MEM[308]` | 0x6008C534 | 8 |  | - |
| `DS_RB_MEM[309]` | 0x6008C535 | 8 |  | - |
| `DS_RB_MEM[30]` | 0x6008C41E | 8 |  | - |
| `DS_RB_MEM[310]` | 0x6008C536 | 8 |  | - |
| `DS_RB_MEM[311]` | 0x6008C537 | 8 |  | - |
| `DS_RB_MEM[312]` | 0x6008C538 | 8 |  | - |
| `DS_RB_MEM[313]` | 0x6008C539 | 8 |  | - |
| `DS_RB_MEM[314]` | 0x6008C53A | 8 |  | - |
| `DS_RB_MEM[315]` | 0x6008C53B | 8 |  | - |
| `DS_RB_MEM[316]` | 0x6008C53C | 8 |  | - |
| `DS_RB_MEM[317]` | 0x6008C53D | 8 |  | - |
| `DS_RB_MEM[318]` | 0x6008C53E | 8 |  | - |
| `DS_RB_MEM[319]` | 0x6008C53F | 8 |  | - |
| `DS_RB_MEM[31]` | 0x6008C41F | 8 |  | - |
| `DS_RB_MEM[320]` | 0x6008C540 | 8 |  | - |
| `DS_RB_MEM[321]` | 0x6008C541 | 8 |  | - |
| `DS_RB_MEM[322]` | 0x6008C542 | 8 |  | - |
| `DS_RB_MEM[323]` | 0x6008C543 | 8 |  | - |
| `DS_RB_MEM[324]` | 0x6008C544 | 8 |  | - |
| `DS_RB_MEM[325]` | 0x6008C545 | 8 |  | - |
| `DS_RB_MEM[326]` | 0x6008C546 | 8 |  | - |
| `DS_RB_MEM[327]` | 0x6008C547 | 8 |  | - |
| `DS_RB_MEM[328]` | 0x6008C548 | 8 |  | - |
| `DS_RB_MEM[329]` | 0x6008C549 | 8 |  | - |
| `DS_RB_MEM[32]` | 0x6008C420 | 8 |  | - |
| `DS_RB_MEM[330]` | 0x6008C54A | 8 |  | - |
| `DS_RB_MEM[331]` | 0x6008C54B | 8 |  | - |
| `DS_RB_MEM[332]` | 0x6008C54C | 8 |  | - |
| `DS_RB_MEM[333]` | 0x6008C54D | 8 |  | - |
| `DS_RB_MEM[334]` | 0x6008C54E | 8 |  | - |
| `DS_RB_MEM[335]` | 0x6008C54F | 8 |  | - |
| `DS_RB_MEM[336]` | 0x6008C550 | 8 |  | - |
| `DS_RB_MEM[337]` | 0x6008C551 | 8 |  | - |
| `DS_RB_MEM[338]` | 0x6008C552 | 8 |  | - |
| `DS_RB_MEM[339]` | 0x6008C553 | 8 |  | - |
| `DS_RB_MEM[33]` | 0x6008C421 | 8 |  | - |
| `DS_RB_MEM[340]` | 0x6008C554 | 8 |  | - |
| `DS_RB_MEM[341]` | 0x6008C555 | 8 |  | - |
| `DS_RB_MEM[342]` | 0x6008C556 | 8 |  | - |
| `DS_RB_MEM[343]` | 0x6008C557 | 8 |  | - |
| `DS_RB_MEM[344]` | 0x6008C558 | 8 |  | - |
| `DS_RB_MEM[345]` | 0x6008C559 | 8 |  | - |
| `DS_RB_MEM[346]` | 0x6008C55A | 8 |  | - |
| `DS_RB_MEM[347]` | 0x6008C55B | 8 |  | - |
| `DS_RB_MEM[348]` | 0x6008C55C | 8 |  | - |
| `DS_RB_MEM[349]` | 0x6008C55D | 8 |  | - |
| `DS_RB_MEM[34]` | 0x6008C422 | 8 |  | - |
| `DS_RB_MEM[350]` | 0x6008C55E | 8 |  | - |
| `DS_RB_MEM[351]` | 0x6008C55F | 8 |  | - |
| `DS_RB_MEM[352]` | 0x6008C560 | 8 |  | - |
| `DS_RB_MEM[353]` | 0x6008C561 | 8 |  | - |
| `DS_RB_MEM[354]` | 0x6008C562 | 8 |  | - |
| `DS_RB_MEM[355]` | 0x6008C563 | 8 |  | - |
| `DS_RB_MEM[356]` | 0x6008C564 | 8 |  | - |
| `DS_RB_MEM[357]` | 0x6008C565 | 8 |  | - |
| `DS_RB_MEM[358]` | 0x6008C566 | 8 |  | - |
| `DS_RB_MEM[359]` | 0x6008C567 | 8 |  | - |
| `DS_RB_MEM[35]` | 0x6008C423 | 8 |  | - |
| `DS_RB_MEM[360]` | 0x6008C568 | 8 |  | - |
| `DS_RB_MEM[361]` | 0x6008C569 | 8 |  | - |
| `DS_RB_MEM[362]` | 0x6008C56A | 8 |  | - |
| `DS_RB_MEM[363]` | 0x6008C56B | 8 |  | - |
| `DS_RB_MEM[364]` | 0x6008C56C | 8 |  | - |
| `DS_RB_MEM[365]` | 0x6008C56D | 8 |  | - |
| `DS_RB_MEM[366]` | 0x6008C56E | 8 |  | - |
| `DS_RB_MEM[367]` | 0x6008C56F | 8 |  | - |
| `DS_RB_MEM[368]` | 0x6008C570 | 8 |  | - |
| `DS_RB_MEM[369]` | 0x6008C571 | 8 |  | - |
| `DS_RB_MEM[36]` | 0x6008C424 | 8 |  | - |
| `DS_RB_MEM[370]` | 0x6008C572 | 8 |  | - |
| `DS_RB_MEM[371]` | 0x6008C573 | 8 |  | - |
| `DS_RB_MEM[372]` | 0x6008C574 | 8 |  | - |
| `DS_RB_MEM[373]` | 0x6008C575 | 8 |  | - |
| `DS_RB_MEM[374]` | 0x6008C576 | 8 |  | - |
| `DS_RB_MEM[375]` | 0x6008C577 | 8 |  | - |
| `DS_RB_MEM[376]` | 0x6008C578 | 8 |  | - |
| `DS_RB_MEM[377]` | 0x6008C579 | 8 |  | - |
| `DS_RB_MEM[378]` | 0x6008C57A | 8 |  | - |
| `DS_RB_MEM[379]` | 0x6008C57B | 8 |  | - |
| `DS_RB_MEM[37]` | 0x6008C425 | 8 |  | - |
| `DS_RB_MEM[380]` | 0x6008C57C | 8 |  | - |
| `DS_RB_MEM[381]` | 0x6008C57D | 8 |  | - |
| `DS_RB_MEM[382]` | 0x6008C57E | 8 |  | - |
| `DS_RB_MEM[383]` | 0x6008C57F | 8 |  | - |
| `DS_RB_MEM[384]` | 0x6008C580 | 8 |  | - |
| `DS_RB_MEM[385]` | 0x6008C581 | 8 |  | - |
| `DS_RB_MEM[386]` | 0x6008C582 | 8 |  | - |
| `DS_RB_MEM[387]` | 0x6008C583 | 8 |  | - |
| `DS_RB_MEM[388]` | 0x6008C584 | 8 |  | - |
| `DS_RB_MEM[389]` | 0x6008C585 | 8 |  | - |
| `DS_RB_MEM[38]` | 0x6008C426 | 8 |  | - |
| `DS_RB_MEM[390]` | 0x6008C586 | 8 |  | - |
| `DS_RB_MEM[391]` | 0x6008C587 | 8 |  | - |
| `DS_RB_MEM[392]` | 0x6008C588 | 8 |  | - |
| `DS_RB_MEM[393]` | 0x6008C589 | 8 |  | - |
| `DS_RB_MEM[394]` | 0x6008C58A | 8 |  | - |
| `DS_RB_MEM[395]` | 0x6008C58B | 8 |  | - |
| `DS_RB_MEM[396]` | 0x6008C58C | 8 |  | - |
| `DS_RB_MEM[397]` | 0x6008C58D | 8 |  | - |
| `DS_RB_MEM[398]` | 0x6008C58E | 8 |  | - |
| `DS_RB_MEM[399]` | 0x6008C58F | 8 |  | - |
| `DS_RB_MEM[39]` | 0x6008C427 | 8 |  | - |
| `DS_RB_MEM[3]` | 0x6008C403 | 8 |  | - |
| `DS_RB_MEM[400]` | 0x6008C590 | 8 |  | - |
| `DS_RB_MEM[401]` | 0x6008C591 | 8 |  | - |
| `DS_RB_MEM[402]` | 0x6008C592 | 8 |  | - |
| `DS_RB_MEM[403]` | 0x6008C593 | 8 |  | - |
| `DS_RB_MEM[404]` | 0x6008C594 | 8 |  | - |
| `DS_RB_MEM[405]` | 0x6008C595 | 8 |  | - |
| `DS_RB_MEM[406]` | 0x6008C596 | 8 |  | - |
| `DS_RB_MEM[407]` | 0x6008C597 | 8 |  | - |
| `DS_RB_MEM[408]` | 0x6008C598 | 8 |  | - |
| `DS_RB_MEM[409]` | 0x6008C599 | 8 |  | - |
| `DS_RB_MEM[40]` | 0x6008C428 | 8 |  | - |
| `DS_RB_MEM[410]` | 0x6008C59A | 8 |  | - |
| `DS_RB_MEM[411]` | 0x6008C59B | 8 |  | - |
| `DS_RB_MEM[412]` | 0x6008C59C | 8 |  | - |
| `DS_RB_MEM[413]` | 0x6008C59D | 8 |  | - |
| `DS_RB_MEM[414]` | 0x6008C59E | 8 |  | - |
| `DS_RB_MEM[415]` | 0x6008C59F | 8 |  | - |
| `DS_RB_MEM[416]` | 0x6008C5A0 | 8 |  | - |
| `DS_RB_MEM[417]` | 0x6008C5A1 | 8 |  | - |
| `DS_RB_MEM[418]` | 0x6008C5A2 | 8 |  | - |
| `DS_RB_MEM[419]` | 0x6008C5A3 | 8 |  | - |
| `DS_RB_MEM[41]` | 0x6008C429 | 8 |  | - |
| `DS_RB_MEM[420]` | 0x6008C5A4 | 8 |  | - |
| `DS_RB_MEM[421]` | 0x6008C5A5 | 8 |  | - |
| `DS_RB_MEM[422]` | 0x6008C5A6 | 8 |  | - |
| `DS_RB_MEM[423]` | 0x6008C5A7 | 8 |  | - |
| `DS_RB_MEM[424]` | 0x6008C5A8 | 8 |  | - |
| `DS_RB_MEM[425]` | 0x6008C5A9 | 8 |  | - |
| `DS_RB_MEM[426]` | 0x6008C5AA | 8 |  | - |
| `DS_RB_MEM[427]` | 0x6008C5AB | 8 |  | - |
| `DS_RB_MEM[428]` | 0x6008C5AC | 8 |  | - |
| `DS_RB_MEM[429]` | 0x6008C5AD | 8 |  | - |
| `DS_RB_MEM[42]` | 0x6008C42A | 8 |  | - |
| `DS_RB_MEM[430]` | 0x6008C5AE | 8 |  | - |
| `DS_RB_MEM[431]` | 0x6008C5AF | 8 |  | - |
| `DS_RB_MEM[432]` | 0x6008C5B0 | 8 |  | - |
| `DS_RB_MEM[433]` | 0x6008C5B1 | 8 |  | - |
| `DS_RB_MEM[434]` | 0x6008C5B2 | 8 |  | - |
| `DS_RB_MEM[435]` | 0x6008C5B3 | 8 |  | - |
| `DS_RB_MEM[436]` | 0x6008C5B4 | 8 |  | - |
| `DS_RB_MEM[437]` | 0x6008C5B5 | 8 |  | - |
| `DS_RB_MEM[438]` | 0x6008C5B6 | 8 |  | - |
| `DS_RB_MEM[439]` | 0x6008C5B7 | 8 |  | - |
| `DS_RB_MEM[43]` | 0x6008C42B | 8 |  | - |
| `DS_RB_MEM[440]` | 0x6008C5B8 | 8 |  | - |
| `DS_RB_MEM[441]` | 0x6008C5B9 | 8 |  | - |
| `DS_RB_MEM[442]` | 0x6008C5BA | 8 |  | - |
| `DS_RB_MEM[443]` | 0x6008C5BB | 8 |  | - |
| `DS_RB_MEM[444]` | 0x6008C5BC | 8 |  | - |
| `DS_RB_MEM[445]` | 0x6008C5BD | 8 |  | - |
| `DS_RB_MEM[446]` | 0x6008C5BE | 8 |  | - |
| `DS_RB_MEM[447]` | 0x6008C5BF | 8 |  | - |
| `DS_RB_MEM[448]` | 0x6008C5C0 | 8 |  | - |
| `DS_RB_MEM[449]` | 0x6008C5C1 | 8 |  | - |
| `DS_RB_MEM[44]` | 0x6008C42C | 8 |  | - |
| `DS_RB_MEM[450]` | 0x6008C5C2 | 8 |  | - |
| `DS_RB_MEM[451]` | 0x6008C5C3 | 8 |  | - |
| `DS_RB_MEM[452]` | 0x6008C5C4 | 8 |  | - |
| `DS_RB_MEM[453]` | 0x6008C5C5 | 8 |  | - |
| `DS_RB_MEM[454]` | 0x6008C5C6 | 8 |  | - |
| `DS_RB_MEM[455]` | 0x6008C5C7 | 8 |  | - |
| `DS_RB_MEM[456]` | 0x6008C5C8 | 8 |  | - |
| `DS_RB_MEM[457]` | 0x6008C5C9 | 8 |  | - |
| `DS_RB_MEM[458]` | 0x6008C5CA | 8 |  | - |
| `DS_RB_MEM[459]` | 0x6008C5CB | 8 |  | - |
| `DS_RB_MEM[45]` | 0x6008C42D | 8 |  | - |
| `DS_RB_MEM[460]` | 0x6008C5CC | 8 |  | - |
| `DS_RB_MEM[461]` | 0x6008C5CD | 8 |  | - |
| `DS_RB_MEM[462]` | 0x6008C5CE | 8 |  | - |
| `DS_RB_MEM[463]` | 0x6008C5CF | 8 |  | - |
| `DS_RB_MEM[464]` | 0x6008C5D0 | 8 |  | - |
| `DS_RB_MEM[465]` | 0x6008C5D1 | 8 |  | - |
| `DS_RB_MEM[466]` | 0x6008C5D2 | 8 |  | - |
| `DS_RB_MEM[467]` | 0x6008C5D3 | 8 |  | - |
| `DS_RB_MEM[468]` | 0x6008C5D4 | 8 |  | - |
| `DS_RB_MEM[469]` | 0x6008C5D5 | 8 |  | - |
| `DS_RB_MEM[46]` | 0x6008C42E | 8 |  | - |
| `DS_RB_MEM[470]` | 0x6008C5D6 | 8 |  | - |
| `DS_RB_MEM[471]` | 0x6008C5D7 | 8 |  | - |
| `DS_RB_MEM[472]` | 0x6008C5D8 | 8 |  | - |
| `DS_RB_MEM[473]` | 0x6008C5D9 | 8 |  | - |
| `DS_RB_MEM[474]` | 0x6008C5DA | 8 |  | - |
| `DS_RB_MEM[475]` | 0x6008C5DB | 8 |  | - |
| `DS_RB_MEM[476]` | 0x6008C5DC | 8 |  | - |
| `DS_RB_MEM[477]` | 0x6008C5DD | 8 |  | - |
| `DS_RB_MEM[478]` | 0x6008C5DE | 8 |  | - |
| `DS_RB_MEM[479]` | 0x6008C5DF | 8 |  | - |
| `DS_RB_MEM[47]` | 0x6008C42F | 8 |  | - |
| `DS_RB_MEM[480]` | 0x6008C5E0 | 8 |  | - |
| `DS_RB_MEM[481]` | 0x6008C5E1 | 8 |  | - |
| `DS_RB_MEM[482]` | 0x6008C5E2 | 8 |  | - |
| `DS_RB_MEM[483]` | 0x6008C5E3 | 8 |  | - |
| `DS_RB_MEM[484]` | 0x6008C5E4 | 8 |  | - |
| `DS_RB_MEM[485]` | 0x6008C5E5 | 8 |  | - |
| `DS_RB_MEM[486]` | 0x6008C5E6 | 8 |  | - |
| `DS_RB_MEM[487]` | 0x6008C5E7 | 8 |  | - |
| `DS_RB_MEM[488]` | 0x6008C5E8 | 8 |  | - |
| `DS_RB_MEM[489]` | 0x6008C5E9 | 8 |  | - |
| `DS_RB_MEM[48]` | 0x6008C430 | 8 |  | - |
| `DS_RB_MEM[490]` | 0x6008C5EA | 8 |  | - |
| `DS_RB_MEM[491]` | 0x6008C5EB | 8 |  | - |
| `DS_RB_MEM[492]` | 0x6008C5EC | 8 |  | - |
| `DS_RB_MEM[493]` | 0x6008C5ED | 8 |  | - |
| `DS_RB_MEM[494]` | 0x6008C5EE | 8 |  | - |
| `DS_RB_MEM[495]` | 0x6008C5EF | 8 |  | - |
| `DS_RB_MEM[496]` | 0x6008C5F0 | 8 |  | - |
| `DS_RB_MEM[497]` | 0x6008C5F1 | 8 |  | - |
| `DS_RB_MEM[498]` | 0x6008C5F2 | 8 |  | - |
| `DS_RB_MEM[499]` | 0x6008C5F3 | 8 |  | - |
| `DS_RB_MEM[49]` | 0x6008C431 | 8 |  | - |
| `DS_RB_MEM[4]` | 0x6008C404 | 8 |  | - |
| `DS_RB_MEM[500]` | 0x6008C5F4 | 8 |  | - |
| `DS_RB_MEM[501]` | 0x6008C5F5 | 8 |  | - |
| `DS_RB_MEM[502]` | 0x6008C5F6 | 8 |  | - |
| `DS_RB_MEM[503]` | 0x6008C5F7 | 8 |  | - |
| `DS_RB_MEM[504]` | 0x6008C5F8 | 8 |  | - |
| `DS_RB_MEM[505]` | 0x6008C5F9 | 8 |  | - |
| `DS_RB_MEM[506]` | 0x6008C5FA | 8 |  | - |
| `DS_RB_MEM[507]` | 0x6008C5FB | 8 |  | - |
| `DS_RB_MEM[508]` | 0x6008C5FC | 8 |  | - |
| `DS_RB_MEM[509]` | 0x6008C5FD | 8 |  | - |
| `DS_RB_MEM[50]` | 0x6008C432 | 8 |  | - |
| `DS_RB_MEM[510]` | 0x6008C5FE | 8 |  | - |
| `DS_RB_MEM[511]` | 0x6008C5FF | 8 |  | - |
| `DS_RB_MEM[51]` | 0x6008C433 | 8 |  | - |
| `DS_RB_MEM[52]` | 0x6008C434 | 8 |  | - |
| `DS_RB_MEM[53]` | 0x6008C435 | 8 |  | - |
| `DS_RB_MEM[54]` | 0x6008C436 | 8 |  | - |
| `DS_RB_MEM[55]` | 0x6008C437 | 8 |  | - |
| `DS_RB_MEM[56]` | 0x6008C438 | 8 |  | - |
| `DS_RB_MEM[57]` | 0x6008C439 | 8 |  | - |
| `DS_RB_MEM[58]` | 0x6008C43A | 8 |  | - |
| `DS_RB_MEM[59]` | 0x6008C43B | 8 |  | - |
| `DS_RB_MEM[5]` | 0x6008C405 | 8 |  | - |
| `DS_RB_MEM[60]` | 0x6008C43C | 8 |  | - |
| `DS_RB_MEM[61]` | 0x6008C43D | 8 |  | - |
| `DS_RB_MEM[62]` | 0x6008C43E | 8 |  | - |
| `DS_RB_MEM[63]` | 0x6008C43F | 8 |  | - |
| `DS_RB_MEM[64]` | 0x6008C440 | 8 |  | - |
| `DS_RB_MEM[65]` | 0x6008C441 | 8 |  | - |
| `DS_RB_MEM[66]` | 0x6008C442 | 8 |  | - |
| `DS_RB_MEM[67]` | 0x6008C443 | 8 |  | - |
| `DS_RB_MEM[68]` | 0x6008C444 | 8 |  | - |
| `DS_RB_MEM[69]` | 0x6008C445 | 8 |  | - |
| `DS_RB_MEM[6]` | 0x6008C406 | 8 |  | - |
| `DS_RB_MEM[70]` | 0x6008C446 | 8 |  | - |
| `DS_RB_MEM[71]` | 0x6008C447 | 8 |  | - |
| `DS_RB_MEM[72]` | 0x6008C448 | 8 |  | - |
| `DS_RB_MEM[73]` | 0x6008C449 | 8 |  | - |
| `DS_RB_MEM[74]` | 0x6008C44A | 8 |  | - |
| `DS_RB_MEM[75]` | 0x6008C44B | 8 |  | - |
| `DS_RB_MEM[76]` | 0x6008C44C | 8 |  | - |
| `DS_RB_MEM[77]` | 0x6008C44D | 8 |  | - |
| `DS_RB_MEM[78]` | 0x6008C44E | 8 |  | - |
| `DS_RB_MEM[79]` | 0x6008C44F | 8 |  | - |
| `DS_RB_MEM[7]` | 0x6008C407 | 8 |  | - |
| `DS_RB_MEM[80]` | 0x6008C450 | 8 |  | - |
| `DS_RB_MEM[81]` | 0x6008C451 | 8 |  | - |
| `DS_RB_MEM[82]` | 0x6008C452 | 8 |  | - |
| `DS_RB_MEM[83]` | 0x6008C453 | 8 |  | - |
| `DS_RB_MEM[84]` | 0x6008C454 | 8 |  | - |
| `DS_RB_MEM[85]` | 0x6008C455 | 8 |  | - |
| `DS_RB_MEM[86]` | 0x6008C456 | 8 |  | - |
| `DS_RB_MEM[87]` | 0x6008C457 | 8 |  | - |
| `DS_RB_MEM[88]` | 0x6008C458 | 8 |  | - |
| `DS_RB_MEM[89]` | 0x6008C459 | 8 |  | - |
| `DS_RB_MEM[8]` | 0x6008C408 | 8 |  | - |
| `DS_RB_MEM[90]` | 0x6008C45A | 8 |  | - |
| `DS_RB_MEM[91]` | 0x6008C45B | 8 |  | - |
| `DS_RB_MEM[92]` | 0x6008C45C | 8 |  | - |
| `DS_RB_MEM[93]` | 0x6008C45D | 8 |  | - |
| `DS_RB_MEM[94]` | 0x6008C45E | 8 |  | - |
| `DS_RB_MEM[95]` | 0x6008C45F | 8 |  | - |
| `DS_RB_MEM[96]` | 0x6008C460 | 8 |  | - |
| `DS_RB_MEM[97]` | 0x6008C461 | 8 |  | - |
| `DS_RB_MEM[98]` | 0x6008C462 | 8 |  | - |
| `DS_RB_MEM[99]` | 0x6008C463 | 8 |  | - |
| `DS_RB_MEM[9]` | 0x6008C409 | 8 |  | - |
| `DS_SET_CONTINUE` | 0x6008CE04 | 32 | SET_CONTINUE=0x1 | - |
| `DS_SET_FINISH` | 0x6008CE08 | 32 | SET_FINISH=0x1 | - |
| `DS_SET_START` | 0x6008CE00 | 32 | SET_START=0x1 | - |
| `DS_X_MEM[0]` | 0x6008C800 | 8 |  | - |
| `DS_X_MEM[100]` | 0x6008C864 | 8 |  | - |
| `DS_X_MEM[101]` | 0x6008C865 | 8 |  | - |
| `DS_X_MEM[102]` | 0x6008C866 | 8 |  | - |
| `DS_X_MEM[103]` | 0x6008C867 | 8 |  | - |
| `DS_X_MEM[104]` | 0x6008C868 | 8 |  | - |
| `DS_X_MEM[105]` | 0x6008C869 | 8 |  | - |
| `DS_X_MEM[106]` | 0x6008C86A | 8 |  | - |
| `DS_X_MEM[107]` | 0x6008C86B | 8 |  | - |
| `DS_X_MEM[108]` | 0x6008C86C | 8 |  | - |
| `DS_X_MEM[109]` | 0x6008C86D | 8 |  | - |
| `DS_X_MEM[10]` | 0x6008C80A | 8 |  | - |
| `DS_X_MEM[110]` | 0x6008C86E | 8 |  | - |
| `DS_X_MEM[111]` | 0x6008C86F | 8 |  | - |
| `DS_X_MEM[112]` | 0x6008C870 | 8 |  | - |
| `DS_X_MEM[113]` | 0x6008C871 | 8 |  | - |
| `DS_X_MEM[114]` | 0x6008C872 | 8 |  | - |
| `DS_X_MEM[115]` | 0x6008C873 | 8 |  | - |
| `DS_X_MEM[116]` | 0x6008C874 | 8 |  | - |
| `DS_X_MEM[117]` | 0x6008C875 | 8 |  | - |
| `DS_X_MEM[118]` | 0x6008C876 | 8 |  | - |
| `DS_X_MEM[119]` | 0x6008C877 | 8 |  | - |
| `DS_X_MEM[11]` | 0x6008C80B | 8 |  | - |
| `DS_X_MEM[120]` | 0x6008C878 | 8 |  | - |
| `DS_X_MEM[121]` | 0x6008C879 | 8 |  | - |
| `DS_X_MEM[122]` | 0x6008C87A | 8 |  | - |
| `DS_X_MEM[123]` | 0x6008C87B | 8 |  | - |
| `DS_X_MEM[124]` | 0x6008C87C | 8 |  | - |
| `DS_X_MEM[125]` | 0x6008C87D | 8 |  | - |
| `DS_X_MEM[126]` | 0x6008C87E | 8 |  | - |
| `DS_X_MEM[127]` | 0x6008C87F | 8 |  | - |
| `DS_X_MEM[128]` | 0x6008C880 | 8 |  | - |
| `DS_X_MEM[129]` | 0x6008C881 | 8 |  | - |
| `DS_X_MEM[12]` | 0x6008C80C | 8 |  | - |
| `DS_X_MEM[130]` | 0x6008C882 | 8 |  | - |
| `DS_X_MEM[131]` | 0x6008C883 | 8 |  | - |
| `DS_X_MEM[132]` | 0x6008C884 | 8 |  | - |
| `DS_X_MEM[133]` | 0x6008C885 | 8 |  | - |
| `DS_X_MEM[134]` | 0x6008C886 | 8 |  | - |
| `DS_X_MEM[135]` | 0x6008C887 | 8 |  | - |
| `DS_X_MEM[136]` | 0x6008C888 | 8 |  | - |
| `DS_X_MEM[137]` | 0x6008C889 | 8 |  | - |
| `DS_X_MEM[138]` | 0x6008C88A | 8 |  | - |
| `DS_X_MEM[139]` | 0x6008C88B | 8 |  | - |
| `DS_X_MEM[13]` | 0x6008C80D | 8 |  | - |
| `DS_X_MEM[140]` | 0x6008C88C | 8 |  | - |
| `DS_X_MEM[141]` | 0x6008C88D | 8 |  | - |
| `DS_X_MEM[142]` | 0x6008C88E | 8 |  | - |
| `DS_X_MEM[143]` | 0x6008C88F | 8 |  | - |
| `DS_X_MEM[144]` | 0x6008C890 | 8 |  | - |
| `DS_X_MEM[145]` | 0x6008C891 | 8 |  | - |
| `DS_X_MEM[146]` | 0x6008C892 | 8 |  | - |
| `DS_X_MEM[147]` | 0x6008C893 | 8 |  | - |
| `DS_X_MEM[148]` | 0x6008C894 | 8 |  | - |
| `DS_X_MEM[149]` | 0x6008C895 | 8 |  | - |
| `DS_X_MEM[14]` | 0x6008C80E | 8 |  | - |
| `DS_X_MEM[150]` | 0x6008C896 | 8 |  | - |
| `DS_X_MEM[151]` | 0x6008C897 | 8 |  | - |
| `DS_X_MEM[152]` | 0x6008C898 | 8 |  | - |
| `DS_X_MEM[153]` | 0x6008C899 | 8 |  | - |
| `DS_X_MEM[154]` | 0x6008C89A | 8 |  | - |
| `DS_X_MEM[155]` | 0x6008C89B | 8 |  | - |
| `DS_X_MEM[156]` | 0x6008C89C | 8 |  | - |
| `DS_X_MEM[157]` | 0x6008C89D | 8 |  | - |
| `DS_X_MEM[158]` | 0x6008C89E | 8 |  | - |
| `DS_X_MEM[159]` | 0x6008C89F | 8 |  | - |
| `DS_X_MEM[15]` | 0x6008C80F | 8 |  | - |
| `DS_X_MEM[160]` | 0x6008C8A0 | 8 |  | - |
| `DS_X_MEM[161]` | 0x6008C8A1 | 8 |  | - |
| `DS_X_MEM[162]` | 0x6008C8A2 | 8 |  | - |
| `DS_X_MEM[163]` | 0x6008C8A3 | 8 |  | - |
| `DS_X_MEM[164]` | 0x6008C8A4 | 8 |  | - |
| `DS_X_MEM[165]` | 0x6008C8A5 | 8 |  | - |
| `DS_X_MEM[166]` | 0x6008C8A6 | 8 |  | - |
| `DS_X_MEM[167]` | 0x6008C8A7 | 8 |  | - |
| `DS_X_MEM[168]` | 0x6008C8A8 | 8 |  | - |
| `DS_X_MEM[169]` | 0x6008C8A9 | 8 |  | - |
| `DS_X_MEM[16]` | 0x6008C810 | 8 |  | - |
| `DS_X_MEM[170]` | 0x6008C8AA | 8 |  | - |
| `DS_X_MEM[171]` | 0x6008C8AB | 8 |  | - |
| `DS_X_MEM[172]` | 0x6008C8AC | 8 |  | - |
| `DS_X_MEM[173]` | 0x6008C8AD | 8 |  | - |
| `DS_X_MEM[174]` | 0x6008C8AE | 8 |  | - |
| `DS_X_MEM[175]` | 0x6008C8AF | 8 |  | - |
| `DS_X_MEM[176]` | 0x6008C8B0 | 8 |  | - |
| `DS_X_MEM[177]` | 0x6008C8B1 | 8 |  | - |
| `DS_X_MEM[178]` | 0x6008C8B2 | 8 |  | - |
| `DS_X_MEM[179]` | 0x6008C8B3 | 8 |  | - |
| `DS_X_MEM[17]` | 0x6008C811 | 8 |  | - |
| `DS_X_MEM[180]` | 0x6008C8B4 | 8 |  | - |
| `DS_X_MEM[181]` | 0x6008C8B5 | 8 |  | - |
| `DS_X_MEM[182]` | 0x6008C8B6 | 8 |  | - |
| `DS_X_MEM[183]` | 0x6008C8B7 | 8 |  | - |
| `DS_X_MEM[184]` | 0x6008C8B8 | 8 |  | - |
| `DS_X_MEM[185]` | 0x6008C8B9 | 8 |  | - |
| `DS_X_MEM[186]` | 0x6008C8BA | 8 |  | - |
| `DS_X_MEM[187]` | 0x6008C8BB | 8 |  | - |
| `DS_X_MEM[188]` | 0x6008C8BC | 8 |  | - |
| `DS_X_MEM[189]` | 0x6008C8BD | 8 |  | - |
| `DS_X_MEM[18]` | 0x6008C812 | 8 |  | - |
| `DS_X_MEM[190]` | 0x6008C8BE | 8 |  | - |
| `DS_X_MEM[191]` | 0x6008C8BF | 8 |  | - |
| `DS_X_MEM[192]` | 0x6008C8C0 | 8 |  | - |
| `DS_X_MEM[193]` | 0x6008C8C1 | 8 |  | - |
| `DS_X_MEM[194]` | 0x6008C8C2 | 8 |  | - |
| `DS_X_MEM[195]` | 0x6008C8C3 | 8 |  | - |
| `DS_X_MEM[196]` | 0x6008C8C4 | 8 |  | - |
| `DS_X_MEM[197]` | 0x6008C8C5 | 8 |  | - |
| `DS_X_MEM[198]` | 0x6008C8C6 | 8 |  | - |
| `DS_X_MEM[199]` | 0x6008C8C7 | 8 |  | - |
| `DS_X_MEM[19]` | 0x6008C813 | 8 |  | - |
| `DS_X_MEM[1]` | 0x6008C801 | 8 |  | - |
| `DS_X_MEM[200]` | 0x6008C8C8 | 8 |  | - |
| `DS_X_MEM[201]` | 0x6008C8C9 | 8 |  | - |
| `DS_X_MEM[202]` | 0x6008C8CA | 8 |  | - |
| `DS_X_MEM[203]` | 0x6008C8CB | 8 |  | - |
| `DS_X_MEM[204]` | 0x6008C8CC | 8 |  | - |
| `DS_X_MEM[205]` | 0x6008C8CD | 8 |  | - |
| `DS_X_MEM[206]` | 0x6008C8CE | 8 |  | - |
| `DS_X_MEM[207]` | 0x6008C8CF | 8 |  | - |
| `DS_X_MEM[208]` | 0x6008C8D0 | 8 |  | - |
| `DS_X_MEM[209]` | 0x6008C8D1 | 8 |  | - |
| `DS_X_MEM[20]` | 0x6008C814 | 8 |  | - |
| `DS_X_MEM[210]` | 0x6008C8D2 | 8 |  | - |
| `DS_X_MEM[211]` | 0x6008C8D3 | 8 |  | - |
| `DS_X_MEM[212]` | 0x6008C8D4 | 8 |  | - |
| `DS_X_MEM[213]` | 0x6008C8D5 | 8 |  | - |
| `DS_X_MEM[214]` | 0x6008C8D6 | 8 |  | - |
| `DS_X_MEM[215]` | 0x6008C8D7 | 8 |  | - |
| `DS_X_MEM[216]` | 0x6008C8D8 | 8 |  | - |
| `DS_X_MEM[217]` | 0x6008C8D9 | 8 |  | - |
| `DS_X_MEM[218]` | 0x6008C8DA | 8 |  | - |
| `DS_X_MEM[219]` | 0x6008C8DB | 8 |  | - |
| `DS_X_MEM[21]` | 0x6008C815 | 8 |  | - |
| `DS_X_MEM[220]` | 0x6008C8DC | 8 |  | - |
| `DS_X_MEM[221]` | 0x6008C8DD | 8 |  | - |
| `DS_X_MEM[222]` | 0x6008C8DE | 8 |  | - |
| `DS_X_MEM[223]` | 0x6008C8DF | 8 |  | - |
| `DS_X_MEM[224]` | 0x6008C8E0 | 8 |  | - |
| `DS_X_MEM[225]` | 0x6008C8E1 | 8 |  | - |
| `DS_X_MEM[226]` | 0x6008C8E2 | 8 |  | - |
| `DS_X_MEM[227]` | 0x6008C8E3 | 8 |  | - |
| `DS_X_MEM[228]` | 0x6008C8E4 | 8 |  | - |
| `DS_X_MEM[229]` | 0x6008C8E5 | 8 |  | - |
| `DS_X_MEM[22]` | 0x6008C816 | 8 |  | - |
| `DS_X_MEM[230]` | 0x6008C8E6 | 8 |  | - |
| `DS_X_MEM[231]` | 0x6008C8E7 | 8 |  | - |
| `DS_X_MEM[232]` | 0x6008C8E8 | 8 |  | - |
| `DS_X_MEM[233]` | 0x6008C8E9 | 8 |  | - |
| `DS_X_MEM[234]` | 0x6008C8EA | 8 |  | - |
| `DS_X_MEM[235]` | 0x6008C8EB | 8 |  | - |
| `DS_X_MEM[236]` | 0x6008C8EC | 8 |  | - |
| `DS_X_MEM[237]` | 0x6008C8ED | 8 |  | - |
| `DS_X_MEM[238]` | 0x6008C8EE | 8 |  | - |
| `DS_X_MEM[239]` | 0x6008C8EF | 8 |  | - |
| `DS_X_MEM[23]` | 0x6008C817 | 8 |  | - |
| `DS_X_MEM[240]` | 0x6008C8F0 | 8 |  | - |
| `DS_X_MEM[241]` | 0x6008C8F1 | 8 |  | - |
| `DS_X_MEM[242]` | 0x6008C8F2 | 8 |  | - |
| `DS_X_MEM[243]` | 0x6008C8F3 | 8 |  | - |
| `DS_X_MEM[244]` | 0x6008C8F4 | 8 |  | - |
| `DS_X_MEM[245]` | 0x6008C8F5 | 8 |  | - |
| `DS_X_MEM[246]` | 0x6008C8F6 | 8 |  | - |
| `DS_X_MEM[247]` | 0x6008C8F7 | 8 |  | - |
| `DS_X_MEM[248]` | 0x6008C8F8 | 8 |  | - |
| `DS_X_MEM[249]` | 0x6008C8F9 | 8 |  | - |
| `DS_X_MEM[24]` | 0x6008C818 | 8 |  | - |
| `DS_X_MEM[250]` | 0x6008C8FA | 8 |  | - |
| `DS_X_MEM[251]` | 0x6008C8FB | 8 |  | - |
| `DS_X_MEM[252]` | 0x6008C8FC | 8 |  | - |
| `DS_X_MEM[253]` | 0x6008C8FD | 8 |  | - |
| `DS_X_MEM[254]` | 0x6008C8FE | 8 |  | - |
| `DS_X_MEM[255]` | 0x6008C8FF | 8 |  | - |
| `DS_X_MEM[256]` | 0x6008C900 | 8 |  | - |
| `DS_X_MEM[257]` | 0x6008C901 | 8 |  | - |
| `DS_X_MEM[258]` | 0x6008C902 | 8 |  | - |
| `DS_X_MEM[259]` | 0x6008C903 | 8 |  | - |
| `DS_X_MEM[25]` | 0x6008C819 | 8 |  | - |
| `DS_X_MEM[260]` | 0x6008C904 | 8 |  | - |
| `DS_X_MEM[261]` | 0x6008C905 | 8 |  | - |
| `DS_X_MEM[262]` | 0x6008C906 | 8 |  | - |
| `DS_X_MEM[263]` | 0x6008C907 | 8 |  | - |
| `DS_X_MEM[264]` | 0x6008C908 | 8 |  | - |
| `DS_X_MEM[265]` | 0x6008C909 | 8 |  | - |
| `DS_X_MEM[266]` | 0x6008C90A | 8 |  | - |
| `DS_X_MEM[267]` | 0x6008C90B | 8 |  | - |
| `DS_X_MEM[268]` | 0x6008C90C | 8 |  | - |
| `DS_X_MEM[269]` | 0x6008C90D | 8 |  | - |
| `DS_X_MEM[26]` | 0x6008C81A | 8 |  | - |
| `DS_X_MEM[270]` | 0x6008C90E | 8 |  | - |
| `DS_X_MEM[271]` | 0x6008C90F | 8 |  | - |
| `DS_X_MEM[272]` | 0x6008C910 | 8 |  | - |
| `DS_X_MEM[273]` | 0x6008C911 | 8 |  | - |
| `DS_X_MEM[274]` | 0x6008C912 | 8 |  | - |
| `DS_X_MEM[275]` | 0x6008C913 | 8 |  | - |
| `DS_X_MEM[276]` | 0x6008C914 | 8 |  | - |
| `DS_X_MEM[277]` | 0x6008C915 | 8 |  | - |
| `DS_X_MEM[278]` | 0x6008C916 | 8 |  | - |
| `DS_X_MEM[279]` | 0x6008C917 | 8 |  | - |
| `DS_X_MEM[27]` | 0x6008C81B | 8 |  | - |
| `DS_X_MEM[280]` | 0x6008C918 | 8 |  | - |
| `DS_X_MEM[281]` | 0x6008C919 | 8 |  | - |
| `DS_X_MEM[282]` | 0x6008C91A | 8 |  | - |
| `DS_X_MEM[283]` | 0x6008C91B | 8 |  | - |
| `DS_X_MEM[284]` | 0x6008C91C | 8 |  | - |
| `DS_X_MEM[285]` | 0x6008C91D | 8 |  | - |
| `DS_X_MEM[286]` | 0x6008C91E | 8 |  | - |
| `DS_X_MEM[287]` | 0x6008C91F | 8 |  | - |
| `DS_X_MEM[288]` | 0x6008C920 | 8 |  | - |
| `DS_X_MEM[289]` | 0x6008C921 | 8 |  | - |
| `DS_X_MEM[28]` | 0x6008C81C | 8 |  | - |
| `DS_X_MEM[290]` | 0x6008C922 | 8 |  | - |
| `DS_X_MEM[291]` | 0x6008C923 | 8 |  | - |
| `DS_X_MEM[292]` | 0x6008C924 | 8 |  | - |
| `DS_X_MEM[293]` | 0x6008C925 | 8 |  | - |
| `DS_X_MEM[294]` | 0x6008C926 | 8 |  | - |
| `DS_X_MEM[295]` | 0x6008C927 | 8 |  | - |
| `DS_X_MEM[296]` | 0x6008C928 | 8 |  | - |
| `DS_X_MEM[297]` | 0x6008C929 | 8 |  | - |
| `DS_X_MEM[298]` | 0x6008C92A | 8 |  | - |
| `DS_X_MEM[299]` | 0x6008C92B | 8 |  | - |
| `DS_X_MEM[29]` | 0x6008C81D | 8 |  | - |
| `DS_X_MEM[2]` | 0x6008C802 | 8 |  | - |
| `DS_X_MEM[300]` | 0x6008C92C | 8 |  | - |
| `DS_X_MEM[301]` | 0x6008C92D | 8 |  | - |
| `DS_X_MEM[302]` | 0x6008C92E | 8 |  | - |
| `DS_X_MEM[303]` | 0x6008C92F | 8 |  | - |
| `DS_X_MEM[304]` | 0x6008C930 | 8 |  | - |
| `DS_X_MEM[305]` | 0x6008C931 | 8 |  | - |
| `DS_X_MEM[306]` | 0x6008C932 | 8 |  | - |
| `DS_X_MEM[307]` | 0x6008C933 | 8 |  | - |
| `DS_X_MEM[308]` | 0x6008C934 | 8 |  | - |
| `DS_X_MEM[309]` | 0x6008C935 | 8 |  | - |
| `DS_X_MEM[30]` | 0x6008C81E | 8 |  | - |
| `DS_X_MEM[310]` | 0x6008C936 | 8 |  | - |
| `DS_X_MEM[311]` | 0x6008C937 | 8 |  | - |
| `DS_X_MEM[312]` | 0x6008C938 | 8 |  | - |
| `DS_X_MEM[313]` | 0x6008C939 | 8 |  | - |
| `DS_X_MEM[314]` | 0x6008C93A | 8 |  | - |
| `DS_X_MEM[315]` | 0x6008C93B | 8 |  | - |
| `DS_X_MEM[316]` | 0x6008C93C | 8 |  | - |
| `DS_X_MEM[317]` | 0x6008C93D | 8 |  | - |
| `DS_X_MEM[318]` | 0x6008C93E | 8 |  | - |
| `DS_X_MEM[319]` | 0x6008C93F | 8 |  | - |
| `DS_X_MEM[31]` | 0x6008C81F | 8 |  | - |
| `DS_X_MEM[320]` | 0x6008C940 | 8 |  | - |
| `DS_X_MEM[321]` | 0x6008C941 | 8 |  | - |
| `DS_X_MEM[322]` | 0x6008C942 | 8 |  | - |
| `DS_X_MEM[323]` | 0x6008C943 | 8 |  | - |
| `DS_X_MEM[324]` | 0x6008C944 | 8 |  | - |
| `DS_X_MEM[325]` | 0x6008C945 | 8 |  | - |
| `DS_X_MEM[326]` | 0x6008C946 | 8 |  | - |
| `DS_X_MEM[327]` | 0x6008C947 | 8 |  | - |
| `DS_X_MEM[328]` | 0x6008C948 | 8 |  | - |
| `DS_X_MEM[329]` | 0x6008C949 | 8 |  | - |
| `DS_X_MEM[32]` | 0x6008C820 | 8 |  | - |
| `DS_X_MEM[330]` | 0x6008C94A | 8 |  | - |
| `DS_X_MEM[331]` | 0x6008C94B | 8 |  | - |
| `DS_X_MEM[332]` | 0x6008C94C | 8 |  | - |
| `DS_X_MEM[333]` | 0x6008C94D | 8 |  | - |
| `DS_X_MEM[334]` | 0x6008C94E | 8 |  | - |
| `DS_X_MEM[335]` | 0x6008C94F | 8 |  | - |
| `DS_X_MEM[336]` | 0x6008C950 | 8 |  | - |
| `DS_X_MEM[337]` | 0x6008C951 | 8 |  | - |
| `DS_X_MEM[338]` | 0x6008C952 | 8 |  | - |
| `DS_X_MEM[339]` | 0x6008C953 | 8 |  | - |
| `DS_X_MEM[33]` | 0x6008C821 | 8 |  | - |
| `DS_X_MEM[340]` | 0x6008C954 | 8 |  | - |
| `DS_X_MEM[341]` | 0x6008C955 | 8 |  | - |
| `DS_X_MEM[342]` | 0x6008C956 | 8 |  | - |
| `DS_X_MEM[343]` | 0x6008C957 | 8 |  | - |
| `DS_X_MEM[344]` | 0x6008C958 | 8 |  | - |
| `DS_X_MEM[345]` | 0x6008C959 | 8 |  | - |
| `DS_X_MEM[346]` | 0x6008C95A | 8 |  | - |
| `DS_X_MEM[347]` | 0x6008C95B | 8 |  | - |
| `DS_X_MEM[348]` | 0x6008C95C | 8 |  | - |
| `DS_X_MEM[349]` | 0x6008C95D | 8 |  | - |
| `DS_X_MEM[34]` | 0x6008C822 | 8 |  | - |
| `DS_X_MEM[350]` | 0x6008C95E | 8 |  | - |
| `DS_X_MEM[351]` | 0x6008C95F | 8 |  | - |
| `DS_X_MEM[352]` | 0x6008C960 | 8 |  | - |
| `DS_X_MEM[353]` | 0x6008C961 | 8 |  | - |
| `DS_X_MEM[354]` | 0x6008C962 | 8 |  | - |
| `DS_X_MEM[355]` | 0x6008C963 | 8 |  | - |
| `DS_X_MEM[356]` | 0x6008C964 | 8 |  | - |
| `DS_X_MEM[357]` | 0x6008C965 | 8 |  | - |
| `DS_X_MEM[358]` | 0x6008C966 | 8 |  | - |
| `DS_X_MEM[359]` | 0x6008C967 | 8 |  | - |
| `DS_X_MEM[35]` | 0x6008C823 | 8 |  | - |
| `DS_X_MEM[360]` | 0x6008C968 | 8 |  | - |
| `DS_X_MEM[361]` | 0x6008C969 | 8 |  | - |
| `DS_X_MEM[362]` | 0x6008C96A | 8 |  | - |
| `DS_X_MEM[363]` | 0x6008C96B | 8 |  | - |
| `DS_X_MEM[364]` | 0x6008C96C | 8 |  | - |
| `DS_X_MEM[365]` | 0x6008C96D | 8 |  | - |
| `DS_X_MEM[366]` | 0x6008C96E | 8 |  | - |
| `DS_X_MEM[367]` | 0x6008C96F | 8 |  | - |
| `DS_X_MEM[368]` | 0x6008C970 | 8 |  | - |
| `DS_X_MEM[369]` | 0x6008C971 | 8 |  | - |
| `DS_X_MEM[36]` | 0x6008C824 | 8 |  | - |
| `DS_X_MEM[370]` | 0x6008C972 | 8 |  | - |
| `DS_X_MEM[371]` | 0x6008C973 | 8 |  | - |
| `DS_X_MEM[372]` | 0x6008C974 | 8 |  | - |
| `DS_X_MEM[373]` | 0x6008C975 | 8 |  | - |
| `DS_X_MEM[374]` | 0x6008C976 | 8 |  | - |
| `DS_X_MEM[375]` | 0x6008C977 | 8 |  | - |
| `DS_X_MEM[376]` | 0x6008C978 | 8 |  | - |
| `DS_X_MEM[377]` | 0x6008C979 | 8 |  | - |
| `DS_X_MEM[378]` | 0x6008C97A | 8 |  | - |
| `DS_X_MEM[379]` | 0x6008C97B | 8 |  | - |
| `DS_X_MEM[37]` | 0x6008C825 | 8 |  | - |
| `DS_X_MEM[380]` | 0x6008C97C | 8 |  | - |
| `DS_X_MEM[381]` | 0x6008C97D | 8 |  | - |
| `DS_X_MEM[382]` | 0x6008C97E | 8 |  | - |
| `DS_X_MEM[383]` | 0x6008C97F | 8 |  | - |
| `DS_X_MEM[384]` | 0x6008C980 | 8 |  | - |
| `DS_X_MEM[385]` | 0x6008C981 | 8 |  | - |
| `DS_X_MEM[386]` | 0x6008C982 | 8 |  | - |
| `DS_X_MEM[387]` | 0x6008C983 | 8 |  | - |
| `DS_X_MEM[388]` | 0x6008C984 | 8 |  | - |
| `DS_X_MEM[389]` | 0x6008C985 | 8 |  | - |
| `DS_X_MEM[38]` | 0x6008C826 | 8 |  | - |
| `DS_X_MEM[390]` | 0x6008C986 | 8 |  | - |
| `DS_X_MEM[391]` | 0x6008C987 | 8 |  | - |
| `DS_X_MEM[392]` | 0x6008C988 | 8 |  | - |
| `DS_X_MEM[393]` | 0x6008C989 | 8 |  | - |
| `DS_X_MEM[394]` | 0x6008C98A | 8 |  | - |
| `DS_X_MEM[395]` | 0x6008C98B | 8 |  | - |
| `DS_X_MEM[396]` | 0x6008C98C | 8 |  | - |
| `DS_X_MEM[397]` | 0x6008C98D | 8 |  | - |
| `DS_X_MEM[398]` | 0x6008C98E | 8 |  | - |
| `DS_X_MEM[399]` | 0x6008C98F | 8 |  | - |
| `DS_X_MEM[39]` | 0x6008C827 | 8 |  | - |
| `DS_X_MEM[3]` | 0x6008C803 | 8 |  | - |
| `DS_X_MEM[400]` | 0x6008C990 | 8 |  | - |
| `DS_X_MEM[401]` | 0x6008C991 | 8 |  | - |
| `DS_X_MEM[402]` | 0x6008C992 | 8 |  | - |
| `DS_X_MEM[403]` | 0x6008C993 | 8 |  | - |
| `DS_X_MEM[404]` | 0x6008C994 | 8 |  | - |
| `DS_X_MEM[405]` | 0x6008C995 | 8 |  | - |
| `DS_X_MEM[406]` | 0x6008C996 | 8 |  | - |
| `DS_X_MEM[407]` | 0x6008C997 | 8 |  | - |
| `DS_X_MEM[408]` | 0x6008C998 | 8 |  | - |
| `DS_X_MEM[409]` | 0x6008C999 | 8 |  | - |
| `DS_X_MEM[40]` | 0x6008C828 | 8 |  | - |
| `DS_X_MEM[410]` | 0x6008C99A | 8 |  | - |
| `DS_X_MEM[411]` | 0x6008C99B | 8 |  | - |
| `DS_X_MEM[412]` | 0x6008C99C | 8 |  | - |
| `DS_X_MEM[413]` | 0x6008C99D | 8 |  | - |
| `DS_X_MEM[414]` | 0x6008C99E | 8 |  | - |
| `DS_X_MEM[415]` | 0x6008C99F | 8 |  | - |
| `DS_X_MEM[416]` | 0x6008C9A0 | 8 |  | - |
| `DS_X_MEM[417]` | 0x6008C9A1 | 8 |  | - |
| `DS_X_MEM[418]` | 0x6008C9A2 | 8 |  | - |
| `DS_X_MEM[419]` | 0x6008C9A3 | 8 |  | - |
| `DS_X_MEM[41]` | 0x6008C829 | 8 |  | - |
| `DS_X_MEM[420]` | 0x6008C9A4 | 8 |  | - |
| `DS_X_MEM[421]` | 0x6008C9A5 | 8 |  | - |
| `DS_X_MEM[422]` | 0x6008C9A6 | 8 |  | - |
| `DS_X_MEM[423]` | 0x6008C9A7 | 8 |  | - |
| `DS_X_MEM[424]` | 0x6008C9A8 | 8 |  | - |
| `DS_X_MEM[425]` | 0x6008C9A9 | 8 |  | - |
| `DS_X_MEM[426]` | 0x6008C9AA | 8 |  | - |
| `DS_X_MEM[427]` | 0x6008C9AB | 8 |  | - |
| `DS_X_MEM[428]` | 0x6008C9AC | 8 |  | - |
| `DS_X_MEM[429]` | 0x6008C9AD | 8 |  | - |
| `DS_X_MEM[42]` | 0x6008C82A | 8 |  | - |
| `DS_X_MEM[430]` | 0x6008C9AE | 8 |  | - |
| `DS_X_MEM[431]` | 0x6008C9AF | 8 |  | - |
| `DS_X_MEM[432]` | 0x6008C9B0 | 8 |  | - |
| `DS_X_MEM[433]` | 0x6008C9B1 | 8 |  | - |
| `DS_X_MEM[434]` | 0x6008C9B2 | 8 |  | - |
| `DS_X_MEM[435]` | 0x6008C9B3 | 8 |  | - |
| `DS_X_MEM[436]` | 0x6008C9B4 | 8 |  | - |
| `DS_X_MEM[437]` | 0x6008C9B5 | 8 |  | - |
| `DS_X_MEM[438]` | 0x6008C9B6 | 8 |  | - |
| `DS_X_MEM[439]` | 0x6008C9B7 | 8 |  | - |
| `DS_X_MEM[43]` | 0x6008C82B | 8 |  | - |
| `DS_X_MEM[440]` | 0x6008C9B8 | 8 |  | - |
| `DS_X_MEM[441]` | 0x6008C9B9 | 8 |  | - |
| `DS_X_MEM[442]` | 0x6008C9BA | 8 |  | - |
| `DS_X_MEM[443]` | 0x6008C9BB | 8 |  | - |
| `DS_X_MEM[444]` | 0x6008C9BC | 8 |  | - |
| `DS_X_MEM[445]` | 0x6008C9BD | 8 |  | - |
| `DS_X_MEM[446]` | 0x6008C9BE | 8 |  | - |
| `DS_X_MEM[447]` | 0x6008C9BF | 8 |  | - |
| `DS_X_MEM[448]` | 0x6008C9C0 | 8 |  | - |
| `DS_X_MEM[449]` | 0x6008C9C1 | 8 |  | - |
| `DS_X_MEM[44]` | 0x6008C82C | 8 |  | - |
| `DS_X_MEM[450]` | 0x6008C9C2 | 8 |  | - |
| `DS_X_MEM[451]` | 0x6008C9C3 | 8 |  | - |
| `DS_X_MEM[452]` | 0x6008C9C4 | 8 |  | - |
| `DS_X_MEM[453]` | 0x6008C9C5 | 8 |  | - |
| `DS_X_MEM[454]` | 0x6008C9C6 | 8 |  | - |
| `DS_X_MEM[455]` | 0x6008C9C7 | 8 |  | - |
| `DS_X_MEM[456]` | 0x6008C9C8 | 8 |  | - |
| `DS_X_MEM[457]` | 0x6008C9C9 | 8 |  | - |
| `DS_X_MEM[458]` | 0x6008C9CA | 8 |  | - |
| `DS_X_MEM[459]` | 0x6008C9CB | 8 |  | - |
| `DS_X_MEM[45]` | 0x6008C82D | 8 |  | - |
| `DS_X_MEM[460]` | 0x6008C9CC | 8 |  | - |
| `DS_X_MEM[461]` | 0x6008C9CD | 8 |  | - |
| `DS_X_MEM[462]` | 0x6008C9CE | 8 |  | - |
| `DS_X_MEM[463]` | 0x6008C9CF | 8 |  | - |
| `DS_X_MEM[464]` | 0x6008C9D0 | 8 |  | - |
| `DS_X_MEM[465]` | 0x6008C9D1 | 8 |  | - |
| `DS_X_MEM[466]` | 0x6008C9D2 | 8 |  | - |
| `DS_X_MEM[467]` | 0x6008C9D3 | 8 |  | - |
| `DS_X_MEM[468]` | 0x6008C9D4 | 8 |  | - |
| `DS_X_MEM[469]` | 0x6008C9D5 | 8 |  | - |
| `DS_X_MEM[46]` | 0x6008C82E | 8 |  | - |
| `DS_X_MEM[470]` | 0x6008C9D6 | 8 |  | - |
| `DS_X_MEM[471]` | 0x6008C9D7 | 8 |  | - |
| `DS_X_MEM[472]` | 0x6008C9D8 | 8 |  | - |
| `DS_X_MEM[473]` | 0x6008C9D9 | 8 |  | - |
| `DS_X_MEM[474]` | 0x6008C9DA | 8 |  | - |
| `DS_X_MEM[475]` | 0x6008C9DB | 8 |  | - |
| `DS_X_MEM[476]` | 0x6008C9DC | 8 |  | - |
| `DS_X_MEM[477]` | 0x6008C9DD | 8 |  | - |
| `DS_X_MEM[478]` | 0x6008C9DE | 8 |  | - |
| `DS_X_MEM[479]` | 0x6008C9DF | 8 |  | - |
| `DS_X_MEM[47]` | 0x6008C82F | 8 |  | - |
| `DS_X_MEM[480]` | 0x6008C9E0 | 8 |  | - |
| `DS_X_MEM[481]` | 0x6008C9E1 | 8 |  | - |
| `DS_X_MEM[482]` | 0x6008C9E2 | 8 |  | - |
| `DS_X_MEM[483]` | 0x6008C9E3 | 8 |  | - |
| `DS_X_MEM[484]` | 0x6008C9E4 | 8 |  | - |
| `DS_X_MEM[485]` | 0x6008C9E5 | 8 |  | - |
| `DS_X_MEM[486]` | 0x6008C9E6 | 8 |  | - |
| `DS_X_MEM[487]` | 0x6008C9E7 | 8 |  | - |
| `DS_X_MEM[488]` | 0x6008C9E8 | 8 |  | - |
| `DS_X_MEM[489]` | 0x6008C9E9 | 8 |  | - |
| `DS_X_MEM[48]` | 0x6008C830 | 8 |  | - |
| `DS_X_MEM[490]` | 0x6008C9EA | 8 |  | - |
| `DS_X_MEM[491]` | 0x6008C9EB | 8 |  | - |
| `DS_X_MEM[492]` | 0x6008C9EC | 8 |  | - |
| `DS_X_MEM[493]` | 0x6008C9ED | 8 |  | - |
| `DS_X_MEM[494]` | 0x6008C9EE | 8 |  | - |
| `DS_X_MEM[495]` | 0x6008C9EF | 8 |  | - |
| `DS_X_MEM[496]` | 0x6008C9F0 | 8 |  | - |
| `DS_X_MEM[497]` | 0x6008C9F1 | 8 |  | - |
| `DS_X_MEM[498]` | 0x6008C9F2 | 8 |  | - |
| `DS_X_MEM[499]` | 0x6008C9F3 | 8 |  | - |
| `DS_X_MEM[49]` | 0x6008C831 | 8 |  | - |
| `DS_X_MEM[4]` | 0x6008C804 | 8 |  | - |
| `DS_X_MEM[500]` | 0x6008C9F4 | 8 |  | - |
| `DS_X_MEM[501]` | 0x6008C9F5 | 8 |  | - |
| `DS_X_MEM[502]` | 0x6008C9F6 | 8 |  | - |
| `DS_X_MEM[503]` | 0x6008C9F7 | 8 |  | - |
| `DS_X_MEM[504]` | 0x6008C9F8 | 8 |  | - |
| `DS_X_MEM[505]` | 0x6008C9F9 | 8 |  | - |
| `DS_X_MEM[506]` | 0x6008C9FA | 8 |  | - |
| `DS_X_MEM[507]` | 0x6008C9FB | 8 |  | - |
| `DS_X_MEM[508]` | 0x6008C9FC | 8 |  | - |
| `DS_X_MEM[509]` | 0x6008C9FD | 8 |  | - |
| `DS_X_MEM[50]` | 0x6008C832 | 8 |  | - |
| `DS_X_MEM[510]` | 0x6008C9FE | 8 |  | - |
| `DS_X_MEM[511]` | 0x6008C9FF | 8 |  | - |
| `DS_X_MEM[51]` | 0x6008C833 | 8 |  | - |
| `DS_X_MEM[52]` | 0x6008C834 | 8 |  | - |
| `DS_X_MEM[53]` | 0x6008C835 | 8 |  | - |
| `DS_X_MEM[54]` | 0x6008C836 | 8 |  | - |
| `DS_X_MEM[55]` | 0x6008C837 | 8 |  | - |
| `DS_X_MEM[56]` | 0x6008C838 | 8 |  | - |
| `DS_X_MEM[57]` | 0x6008C839 | 8 |  | - |
| `DS_X_MEM[58]` | 0x6008C83A | 8 |  | - |
| `DS_X_MEM[59]` | 0x6008C83B | 8 |  | - |
| `DS_X_MEM[5]` | 0x6008C805 | 8 |  | - |
| `DS_X_MEM[60]` | 0x6008C83C | 8 |  | - |
| `DS_X_MEM[61]` | 0x6008C83D | 8 |  | - |
| `DS_X_MEM[62]` | 0x6008C83E | 8 |  | - |
| `DS_X_MEM[63]` | 0x6008C83F | 8 |  | - |
| `DS_X_MEM[64]` | 0x6008C840 | 8 |  | - |
| `DS_X_MEM[65]` | 0x6008C841 | 8 |  | - |
| `DS_X_MEM[66]` | 0x6008C842 | 8 |  | - |
| `DS_X_MEM[67]` | 0x6008C843 | 8 |  | - |
| `DS_X_MEM[68]` | 0x6008C844 | 8 |  | - |
| `DS_X_MEM[69]` | 0x6008C845 | 8 |  | - |
| `DS_X_MEM[6]` | 0x6008C806 | 8 |  | - |
| `DS_X_MEM[70]` | 0x6008C846 | 8 |  | - |
| `DS_X_MEM[71]` | 0x6008C847 | 8 |  | - |
| `DS_X_MEM[72]` | 0x6008C848 | 8 |  | - |
| `DS_X_MEM[73]` | 0x6008C849 | 8 |  | - |
| `DS_X_MEM[74]` | 0x6008C84A | 8 |  | - |
| `DS_X_MEM[75]` | 0x6008C84B | 8 |  | - |
| `DS_X_MEM[76]` | 0x6008C84C | 8 |  | - |
| `DS_X_MEM[77]` | 0x6008C84D | 8 |  | - |
| `DS_X_MEM[78]` | 0x6008C84E | 8 |  | - |
| `DS_X_MEM[79]` | 0x6008C84F | 8 |  | - |
| `DS_X_MEM[7]` | 0x6008C807 | 8 |  | - |
| `DS_X_MEM[80]` | 0x6008C850 | 8 |  | - |
| `DS_X_MEM[81]` | 0x6008C851 | 8 |  | - |
| `DS_X_MEM[82]` | 0x6008C852 | 8 |  | - |
| `DS_X_MEM[83]` | 0x6008C853 | 8 |  | - |
| `DS_X_MEM[84]` | 0x6008C854 | 8 |  | - |
| `DS_X_MEM[85]` | 0x6008C855 | 8 |  | - |
| `DS_X_MEM[86]` | 0x6008C856 | 8 |  | - |
| `DS_X_MEM[87]` | 0x6008C857 | 8 |  | - |
| `DS_X_MEM[88]` | 0x6008C858 | 8 |  | - |
| `DS_X_MEM[89]` | 0x6008C859 | 8 |  | - |
| `DS_X_MEM[8]` | 0x6008C808 | 8 |  | - |
| `DS_X_MEM[90]` | 0x6008C85A | 8 |  | - |
| `DS_X_MEM[91]` | 0x6008C85B | 8 |  | - |
| `DS_X_MEM[92]` | 0x6008C85C | 8 |  | - |
| `DS_X_MEM[93]` | 0x6008C85D | 8 |  | - |
| `DS_X_MEM[94]` | 0x6008C85E | 8 |  | - |
| `DS_X_MEM[95]` | 0x6008C85F | 8 |  | - |
| `DS_X_MEM[96]` | 0x6008C860 | 8 |  | - |
| `DS_X_MEM[97]` | 0x6008C861 | 8 |  | - |
| `DS_X_MEM[98]` | 0x6008C862 | 8 |  | - |
| `DS_X_MEM[99]` | 0x6008C863 | 8 |  | - |
| `DS_X_MEM[9]` | 0x6008C809 | 8 |  | - |
| `DS_Y_MEM[0]` | 0x6008C000 | 8 |  | - |
| `DS_Y_MEM[100]` | 0x6008C064 | 8 |  | - |
| `DS_Y_MEM[101]` | 0x6008C065 | 8 |  | - |
| `DS_Y_MEM[102]` | 0x6008C066 | 8 |  | - |
| `DS_Y_MEM[103]` | 0x6008C067 | 8 |  | - |
| `DS_Y_MEM[104]` | 0x6008C068 | 8 |  | - |
| `DS_Y_MEM[105]` | 0x6008C069 | 8 |  | - |
| `DS_Y_MEM[106]` | 0x6008C06A | 8 |  | - |
| `DS_Y_MEM[107]` | 0x6008C06B | 8 |  | - |
| `DS_Y_MEM[108]` | 0x6008C06C | 8 |  | - |
| `DS_Y_MEM[109]` | 0x6008C06D | 8 |  | - |
| `DS_Y_MEM[10]` | 0x6008C00A | 8 |  | - |
| `DS_Y_MEM[110]` | 0x6008C06E | 8 |  | - |
| `DS_Y_MEM[111]` | 0x6008C06F | 8 |  | - |
| `DS_Y_MEM[112]` | 0x6008C070 | 8 |  | - |
| `DS_Y_MEM[113]` | 0x6008C071 | 8 |  | - |
| `DS_Y_MEM[114]` | 0x6008C072 | 8 |  | - |
| `DS_Y_MEM[115]` | 0x6008C073 | 8 |  | - |
| `DS_Y_MEM[116]` | 0x6008C074 | 8 |  | - |
| `DS_Y_MEM[117]` | 0x6008C075 | 8 |  | - |
| `DS_Y_MEM[118]` | 0x6008C076 | 8 |  | - |
| `DS_Y_MEM[119]` | 0x6008C077 | 8 |  | - |
| `DS_Y_MEM[11]` | 0x6008C00B | 8 |  | - |
| `DS_Y_MEM[120]` | 0x6008C078 | 8 |  | - |
| `DS_Y_MEM[121]` | 0x6008C079 | 8 |  | - |
| `DS_Y_MEM[122]` | 0x6008C07A | 8 |  | - |
| `DS_Y_MEM[123]` | 0x6008C07B | 8 |  | - |
| `DS_Y_MEM[124]` | 0x6008C07C | 8 |  | - |
| `DS_Y_MEM[125]` | 0x6008C07D | 8 |  | - |
| `DS_Y_MEM[126]` | 0x6008C07E | 8 |  | - |
| `DS_Y_MEM[127]` | 0x6008C07F | 8 |  | - |
| `DS_Y_MEM[128]` | 0x6008C080 | 8 |  | - |
| `DS_Y_MEM[129]` | 0x6008C081 | 8 |  | - |
| `DS_Y_MEM[12]` | 0x6008C00C | 8 |  | - |
| `DS_Y_MEM[130]` | 0x6008C082 | 8 |  | - |
| `DS_Y_MEM[131]` | 0x6008C083 | 8 |  | - |
| `DS_Y_MEM[132]` | 0x6008C084 | 8 |  | - |
| `DS_Y_MEM[133]` | 0x6008C085 | 8 |  | - |
| `DS_Y_MEM[134]` | 0x6008C086 | 8 |  | - |
| `DS_Y_MEM[135]` | 0x6008C087 | 8 |  | - |
| `DS_Y_MEM[136]` | 0x6008C088 | 8 |  | - |
| `DS_Y_MEM[137]` | 0x6008C089 | 8 |  | - |
| `DS_Y_MEM[138]` | 0x6008C08A | 8 |  | - |
| `DS_Y_MEM[139]` | 0x6008C08B | 8 |  | - |
| `DS_Y_MEM[13]` | 0x6008C00D | 8 |  | - |
| `DS_Y_MEM[140]` | 0x6008C08C | 8 |  | - |
| `DS_Y_MEM[141]` | 0x6008C08D | 8 |  | - |
| `DS_Y_MEM[142]` | 0x6008C08E | 8 |  | - |
| `DS_Y_MEM[143]` | 0x6008C08F | 8 |  | - |
| `DS_Y_MEM[144]` | 0x6008C090 | 8 |  | - |
| `DS_Y_MEM[145]` | 0x6008C091 | 8 |  | - |
| `DS_Y_MEM[146]` | 0x6008C092 | 8 |  | - |
| `DS_Y_MEM[147]` | 0x6008C093 | 8 |  | - |
| `DS_Y_MEM[148]` | 0x6008C094 | 8 |  | - |
| `DS_Y_MEM[149]` | 0x6008C095 | 8 |  | - |
| `DS_Y_MEM[14]` | 0x6008C00E | 8 |  | - |
| `DS_Y_MEM[150]` | 0x6008C096 | 8 |  | - |
| `DS_Y_MEM[151]` | 0x6008C097 | 8 |  | - |
| `DS_Y_MEM[152]` | 0x6008C098 | 8 |  | - |
| `DS_Y_MEM[153]` | 0x6008C099 | 8 |  | - |
| `DS_Y_MEM[154]` | 0x6008C09A | 8 |  | - |
| `DS_Y_MEM[155]` | 0x6008C09B | 8 |  | - |
| `DS_Y_MEM[156]` | 0x6008C09C | 8 |  | - |
| `DS_Y_MEM[157]` | 0x6008C09D | 8 |  | - |
| `DS_Y_MEM[158]` | 0x6008C09E | 8 |  | - |
| `DS_Y_MEM[159]` | 0x6008C09F | 8 |  | - |
| `DS_Y_MEM[15]` | 0x6008C00F | 8 |  | - |
| `DS_Y_MEM[160]` | 0x6008C0A0 | 8 |  | - |
| `DS_Y_MEM[161]` | 0x6008C0A1 | 8 |  | - |
| `DS_Y_MEM[162]` | 0x6008C0A2 | 8 |  | - |
| `DS_Y_MEM[163]` | 0x6008C0A3 | 8 |  | - |
| `DS_Y_MEM[164]` | 0x6008C0A4 | 8 |  | - |
| `DS_Y_MEM[165]` | 0x6008C0A5 | 8 |  | - |
| `DS_Y_MEM[166]` | 0x6008C0A6 | 8 |  | - |
| `DS_Y_MEM[167]` | 0x6008C0A7 | 8 |  | - |
| `DS_Y_MEM[168]` | 0x6008C0A8 | 8 |  | - |
| `DS_Y_MEM[169]` | 0x6008C0A9 | 8 |  | - |
| `DS_Y_MEM[16]` | 0x6008C010 | 8 |  | - |
| `DS_Y_MEM[170]` | 0x6008C0AA | 8 |  | - |
| `DS_Y_MEM[171]` | 0x6008C0AB | 8 |  | - |
| `DS_Y_MEM[172]` | 0x6008C0AC | 8 |  | - |
| `DS_Y_MEM[173]` | 0x6008C0AD | 8 |  | - |
| `DS_Y_MEM[174]` | 0x6008C0AE | 8 |  | - |
| `DS_Y_MEM[175]` | 0x6008C0AF | 8 |  | - |
| `DS_Y_MEM[176]` | 0x6008C0B0 | 8 |  | - |
| `DS_Y_MEM[177]` | 0x6008C0B1 | 8 |  | - |
| `DS_Y_MEM[178]` | 0x6008C0B2 | 8 |  | - |
| `DS_Y_MEM[179]` | 0x6008C0B3 | 8 |  | - |
| `DS_Y_MEM[17]` | 0x6008C011 | 8 |  | - |
| `DS_Y_MEM[180]` | 0x6008C0B4 | 8 |  | - |
| `DS_Y_MEM[181]` | 0x6008C0B5 | 8 |  | - |
| `DS_Y_MEM[182]` | 0x6008C0B6 | 8 |  | - |
| `DS_Y_MEM[183]` | 0x6008C0B7 | 8 |  | - |
| `DS_Y_MEM[184]` | 0x6008C0B8 | 8 |  | - |
| `DS_Y_MEM[185]` | 0x6008C0B9 | 8 |  | - |
| `DS_Y_MEM[186]` | 0x6008C0BA | 8 |  | - |
| `DS_Y_MEM[187]` | 0x6008C0BB | 8 |  | - |
| `DS_Y_MEM[188]` | 0x6008C0BC | 8 |  | - |
| `DS_Y_MEM[189]` | 0x6008C0BD | 8 |  | - |
| `DS_Y_MEM[18]` | 0x6008C012 | 8 |  | - |
| `DS_Y_MEM[190]` | 0x6008C0BE | 8 |  | - |
| `DS_Y_MEM[191]` | 0x6008C0BF | 8 |  | - |
| `DS_Y_MEM[192]` | 0x6008C0C0 | 8 |  | - |
| `DS_Y_MEM[193]` | 0x6008C0C1 | 8 |  | - |
| `DS_Y_MEM[194]` | 0x6008C0C2 | 8 |  | - |
| `DS_Y_MEM[195]` | 0x6008C0C3 | 8 |  | - |
| `DS_Y_MEM[196]` | 0x6008C0C4 | 8 |  | - |
| `DS_Y_MEM[197]` | 0x6008C0C5 | 8 |  | - |
| `DS_Y_MEM[198]` | 0x6008C0C6 | 8 |  | - |
| `DS_Y_MEM[199]` | 0x6008C0C7 | 8 |  | - |
| `DS_Y_MEM[19]` | 0x6008C013 | 8 |  | - |
| `DS_Y_MEM[1]` | 0x6008C001 | 8 |  | - |
| `DS_Y_MEM[200]` | 0x6008C0C8 | 8 |  | - |
| `DS_Y_MEM[201]` | 0x6008C0C9 | 8 |  | - |
| `DS_Y_MEM[202]` | 0x6008C0CA | 8 |  | - |
| `DS_Y_MEM[203]` | 0x6008C0CB | 8 |  | - |
| `DS_Y_MEM[204]` | 0x6008C0CC | 8 |  | - |
| `DS_Y_MEM[205]` | 0x6008C0CD | 8 |  | - |
| `DS_Y_MEM[206]` | 0x6008C0CE | 8 |  | - |
| `DS_Y_MEM[207]` | 0x6008C0CF | 8 |  | - |
| `DS_Y_MEM[208]` | 0x6008C0D0 | 8 |  | - |
| `DS_Y_MEM[209]` | 0x6008C0D1 | 8 |  | - |
| `DS_Y_MEM[20]` | 0x6008C014 | 8 |  | - |
| `DS_Y_MEM[210]` | 0x6008C0D2 | 8 |  | - |
| `DS_Y_MEM[211]` | 0x6008C0D3 | 8 |  | - |
| `DS_Y_MEM[212]` | 0x6008C0D4 | 8 |  | - |
| `DS_Y_MEM[213]` | 0x6008C0D5 | 8 |  | - |
| `DS_Y_MEM[214]` | 0x6008C0D6 | 8 |  | - |
| `DS_Y_MEM[215]` | 0x6008C0D7 | 8 |  | - |
| `DS_Y_MEM[216]` | 0x6008C0D8 | 8 |  | - |
| `DS_Y_MEM[217]` | 0x6008C0D9 | 8 |  | - |
| `DS_Y_MEM[218]` | 0x6008C0DA | 8 |  | - |
| `DS_Y_MEM[219]` | 0x6008C0DB | 8 |  | - |
| `DS_Y_MEM[21]` | 0x6008C015 | 8 |  | - |
| `DS_Y_MEM[220]` | 0x6008C0DC | 8 |  | - |
| `DS_Y_MEM[221]` | 0x6008C0DD | 8 |  | - |
| `DS_Y_MEM[222]` | 0x6008C0DE | 8 |  | - |
| `DS_Y_MEM[223]` | 0x6008C0DF | 8 |  | - |
| `DS_Y_MEM[224]` | 0x6008C0E0 | 8 |  | - |
| `DS_Y_MEM[225]` | 0x6008C0E1 | 8 |  | - |
| `DS_Y_MEM[226]` | 0x6008C0E2 | 8 |  | - |
| `DS_Y_MEM[227]` | 0x6008C0E3 | 8 |  | - |
| `DS_Y_MEM[228]` | 0x6008C0E4 | 8 |  | - |
| `DS_Y_MEM[229]` | 0x6008C0E5 | 8 |  | - |
| `DS_Y_MEM[22]` | 0x6008C016 | 8 |  | - |
| `DS_Y_MEM[230]` | 0x6008C0E6 | 8 |  | - |
| `DS_Y_MEM[231]` | 0x6008C0E7 | 8 |  | - |
| `DS_Y_MEM[232]` | 0x6008C0E8 | 8 |  | - |
| `DS_Y_MEM[233]` | 0x6008C0E9 | 8 |  | - |
| `DS_Y_MEM[234]` | 0x6008C0EA | 8 |  | - |
| `DS_Y_MEM[235]` | 0x6008C0EB | 8 |  | - |
| `DS_Y_MEM[236]` | 0x6008C0EC | 8 |  | - |
| `DS_Y_MEM[237]` | 0x6008C0ED | 8 |  | - |
| `DS_Y_MEM[238]` | 0x6008C0EE | 8 |  | - |
| `DS_Y_MEM[239]` | 0x6008C0EF | 8 |  | - |
| `DS_Y_MEM[23]` | 0x6008C017 | 8 |  | - |
| `DS_Y_MEM[240]` | 0x6008C0F0 | 8 |  | - |
| `DS_Y_MEM[241]` | 0x6008C0F1 | 8 |  | - |
| `DS_Y_MEM[242]` | 0x6008C0F2 | 8 |  | - |
| `DS_Y_MEM[243]` | 0x6008C0F3 | 8 |  | - |
| `DS_Y_MEM[244]` | 0x6008C0F4 | 8 |  | - |
| `DS_Y_MEM[245]` | 0x6008C0F5 | 8 |  | - |
| `DS_Y_MEM[246]` | 0x6008C0F6 | 8 |  | - |
| `DS_Y_MEM[247]` | 0x6008C0F7 | 8 |  | - |
| `DS_Y_MEM[248]` | 0x6008C0F8 | 8 |  | - |
| `DS_Y_MEM[249]` | 0x6008C0F9 | 8 |  | - |
| `DS_Y_MEM[24]` | 0x6008C018 | 8 |  | - |
| `DS_Y_MEM[250]` | 0x6008C0FA | 8 |  | - |
| `DS_Y_MEM[251]` | 0x6008C0FB | 8 |  | - |
| `DS_Y_MEM[252]` | 0x6008C0FC | 8 |  | - |
| `DS_Y_MEM[253]` | 0x6008C0FD | 8 |  | - |
| `DS_Y_MEM[254]` | 0x6008C0FE | 8 |  | - |
| `DS_Y_MEM[255]` | 0x6008C0FF | 8 |  | - |
| `DS_Y_MEM[256]` | 0x6008C100 | 8 |  | - |
| `DS_Y_MEM[257]` | 0x6008C101 | 8 |  | - |
| `DS_Y_MEM[258]` | 0x6008C102 | 8 |  | - |
| `DS_Y_MEM[259]` | 0x6008C103 | 8 |  | - |
| `DS_Y_MEM[25]` | 0x6008C019 | 8 |  | - |
| `DS_Y_MEM[260]` | 0x6008C104 | 8 |  | - |
| `DS_Y_MEM[261]` | 0x6008C105 | 8 |  | - |
| `DS_Y_MEM[262]` | 0x6008C106 | 8 |  | - |
| `DS_Y_MEM[263]` | 0x6008C107 | 8 |  | - |
| `DS_Y_MEM[264]` | 0x6008C108 | 8 |  | - |
| `DS_Y_MEM[265]` | 0x6008C109 | 8 |  | - |
| `DS_Y_MEM[266]` | 0x6008C10A | 8 |  | - |
| `DS_Y_MEM[267]` | 0x6008C10B | 8 |  | - |
| `DS_Y_MEM[268]` | 0x6008C10C | 8 |  | - |
| `DS_Y_MEM[269]` | 0x6008C10D | 8 |  | - |
| `DS_Y_MEM[26]` | 0x6008C01A | 8 |  | - |
| `DS_Y_MEM[270]` | 0x6008C10E | 8 |  | - |
| `DS_Y_MEM[271]` | 0x6008C10F | 8 |  | - |
| `DS_Y_MEM[272]` | 0x6008C110 | 8 |  | - |
| `DS_Y_MEM[273]` | 0x6008C111 | 8 |  | - |
| `DS_Y_MEM[274]` | 0x6008C112 | 8 |  | - |
| `DS_Y_MEM[275]` | 0x6008C113 | 8 |  | - |
| `DS_Y_MEM[276]` | 0x6008C114 | 8 |  | - |
| `DS_Y_MEM[277]` | 0x6008C115 | 8 |  | - |
| `DS_Y_MEM[278]` | 0x6008C116 | 8 |  | - |
| `DS_Y_MEM[279]` | 0x6008C117 | 8 |  | - |
| `DS_Y_MEM[27]` | 0x6008C01B | 8 |  | - |
| `DS_Y_MEM[280]` | 0x6008C118 | 8 |  | - |
| `DS_Y_MEM[281]` | 0x6008C119 | 8 |  | - |
| `DS_Y_MEM[282]` | 0x6008C11A | 8 |  | - |
| `DS_Y_MEM[283]` | 0x6008C11B | 8 |  | - |
| `DS_Y_MEM[284]` | 0x6008C11C | 8 |  | - |
| `DS_Y_MEM[285]` | 0x6008C11D | 8 |  | - |
| `DS_Y_MEM[286]` | 0x6008C11E | 8 |  | - |
| `DS_Y_MEM[287]` | 0x6008C11F | 8 |  | - |
| `DS_Y_MEM[288]` | 0x6008C120 | 8 |  | - |
| `DS_Y_MEM[289]` | 0x6008C121 | 8 |  | - |
| `DS_Y_MEM[28]` | 0x6008C01C | 8 |  | - |
| `DS_Y_MEM[290]` | 0x6008C122 | 8 |  | - |
| `DS_Y_MEM[291]` | 0x6008C123 | 8 |  | - |
| `DS_Y_MEM[292]` | 0x6008C124 | 8 |  | - |
| `DS_Y_MEM[293]` | 0x6008C125 | 8 |  | - |
| `DS_Y_MEM[294]` | 0x6008C126 | 8 |  | - |
| `DS_Y_MEM[295]` | 0x6008C127 | 8 |  | - |
| `DS_Y_MEM[296]` | 0x6008C128 | 8 |  | - |
| `DS_Y_MEM[297]` | 0x6008C129 | 8 |  | - |
| `DS_Y_MEM[298]` | 0x6008C12A | 8 |  | - |
| `DS_Y_MEM[299]` | 0x6008C12B | 8 |  | - |
| `DS_Y_MEM[29]` | 0x6008C01D | 8 |  | - |
| `DS_Y_MEM[2]` | 0x6008C002 | 8 |  | - |
| `DS_Y_MEM[300]` | 0x6008C12C | 8 |  | - |
| `DS_Y_MEM[301]` | 0x6008C12D | 8 |  | - |
| `DS_Y_MEM[302]` | 0x6008C12E | 8 |  | - |
| `DS_Y_MEM[303]` | 0x6008C12F | 8 |  | - |
| `DS_Y_MEM[304]` | 0x6008C130 | 8 |  | - |
| `DS_Y_MEM[305]` | 0x6008C131 | 8 |  | - |
| `DS_Y_MEM[306]` | 0x6008C132 | 8 |  | - |
| `DS_Y_MEM[307]` | 0x6008C133 | 8 |  | - |
| `DS_Y_MEM[308]` | 0x6008C134 | 8 |  | - |
| `DS_Y_MEM[309]` | 0x6008C135 | 8 |  | - |
| `DS_Y_MEM[30]` | 0x6008C01E | 8 |  | - |
| `DS_Y_MEM[310]` | 0x6008C136 | 8 |  | - |
| `DS_Y_MEM[311]` | 0x6008C137 | 8 |  | - |
| `DS_Y_MEM[312]` | 0x6008C138 | 8 |  | - |
| `DS_Y_MEM[313]` | 0x6008C139 | 8 |  | - |
| `DS_Y_MEM[314]` | 0x6008C13A | 8 |  | - |
| `DS_Y_MEM[315]` | 0x6008C13B | 8 |  | - |
| `DS_Y_MEM[316]` | 0x6008C13C | 8 |  | - |
| `DS_Y_MEM[317]` | 0x6008C13D | 8 |  | - |
| `DS_Y_MEM[318]` | 0x6008C13E | 8 |  | - |
| `DS_Y_MEM[319]` | 0x6008C13F | 8 |  | - |
| `DS_Y_MEM[31]` | 0x6008C01F | 8 |  | - |
| `DS_Y_MEM[320]` | 0x6008C140 | 8 |  | - |
| `DS_Y_MEM[321]` | 0x6008C141 | 8 |  | - |
| `DS_Y_MEM[322]` | 0x6008C142 | 8 |  | - |
| `DS_Y_MEM[323]` | 0x6008C143 | 8 |  | - |
| `DS_Y_MEM[324]` | 0x6008C144 | 8 |  | - |
| `DS_Y_MEM[325]` | 0x6008C145 | 8 |  | - |
| `DS_Y_MEM[326]` | 0x6008C146 | 8 |  | - |
| `DS_Y_MEM[327]` | 0x6008C147 | 8 |  | - |
| `DS_Y_MEM[328]` | 0x6008C148 | 8 |  | - |
| `DS_Y_MEM[329]` | 0x6008C149 | 8 |  | - |
| `DS_Y_MEM[32]` | 0x6008C020 | 8 |  | - |
| `DS_Y_MEM[330]` | 0x6008C14A | 8 |  | - |
| `DS_Y_MEM[331]` | 0x6008C14B | 8 |  | - |
| `DS_Y_MEM[332]` | 0x6008C14C | 8 |  | - |
| `DS_Y_MEM[333]` | 0x6008C14D | 8 |  | - |
| `DS_Y_MEM[334]` | 0x6008C14E | 8 |  | - |
| `DS_Y_MEM[335]` | 0x6008C14F | 8 |  | - |
| `DS_Y_MEM[336]` | 0x6008C150 | 8 |  | - |
| `DS_Y_MEM[337]` | 0x6008C151 | 8 |  | - |
| `DS_Y_MEM[338]` | 0x6008C152 | 8 |  | - |
| `DS_Y_MEM[339]` | 0x6008C153 | 8 |  | - |
| `DS_Y_MEM[33]` | 0x6008C021 | 8 |  | - |
| `DS_Y_MEM[340]` | 0x6008C154 | 8 |  | - |
| `DS_Y_MEM[341]` | 0x6008C155 | 8 |  | - |
| `DS_Y_MEM[342]` | 0x6008C156 | 8 |  | - |
| `DS_Y_MEM[343]` | 0x6008C157 | 8 |  | - |
| `DS_Y_MEM[344]` | 0x6008C158 | 8 |  | - |
| `DS_Y_MEM[345]` | 0x6008C159 | 8 |  | - |
| `DS_Y_MEM[346]` | 0x6008C15A | 8 |  | - |
| `DS_Y_MEM[347]` | 0x6008C15B | 8 |  | - |
| `DS_Y_MEM[348]` | 0x6008C15C | 8 |  | - |
| `DS_Y_MEM[349]` | 0x6008C15D | 8 |  | - |
| `DS_Y_MEM[34]` | 0x6008C022 | 8 |  | - |
| `DS_Y_MEM[350]` | 0x6008C15E | 8 |  | - |
| `DS_Y_MEM[351]` | 0x6008C15F | 8 |  | - |
| `DS_Y_MEM[352]` | 0x6008C160 | 8 |  | - |
| `DS_Y_MEM[353]` | 0x6008C161 | 8 |  | - |
| `DS_Y_MEM[354]` | 0x6008C162 | 8 |  | - |
| `DS_Y_MEM[355]` | 0x6008C163 | 8 |  | - |
| `DS_Y_MEM[356]` | 0x6008C164 | 8 |  | - |
| `DS_Y_MEM[357]` | 0x6008C165 | 8 |  | - |
| `DS_Y_MEM[358]` | 0x6008C166 | 8 |  | - |
| `DS_Y_MEM[359]` | 0x6008C167 | 8 |  | - |
| `DS_Y_MEM[35]` | 0x6008C023 | 8 |  | - |
| `DS_Y_MEM[360]` | 0x6008C168 | 8 |  | - |
| `DS_Y_MEM[361]` | 0x6008C169 | 8 |  | - |
| `DS_Y_MEM[362]` | 0x6008C16A | 8 |  | - |
| `DS_Y_MEM[363]` | 0x6008C16B | 8 |  | - |
| `DS_Y_MEM[364]` | 0x6008C16C | 8 |  | - |
| `DS_Y_MEM[365]` | 0x6008C16D | 8 |  | - |
| `DS_Y_MEM[366]` | 0x6008C16E | 8 |  | - |
| `DS_Y_MEM[367]` | 0x6008C16F | 8 |  | - |
| `DS_Y_MEM[368]` | 0x6008C170 | 8 |  | - |
| `DS_Y_MEM[369]` | 0x6008C171 | 8 |  | - |
| `DS_Y_MEM[36]` | 0x6008C024 | 8 |  | - |
| `DS_Y_MEM[370]` | 0x6008C172 | 8 |  | - |
| `DS_Y_MEM[371]` | 0x6008C173 | 8 |  | - |
| `DS_Y_MEM[372]` | 0x6008C174 | 8 |  | - |
| `DS_Y_MEM[373]` | 0x6008C175 | 8 |  | - |
| `DS_Y_MEM[374]` | 0x6008C176 | 8 |  | - |
| `DS_Y_MEM[375]` | 0x6008C177 | 8 |  | - |
| `DS_Y_MEM[376]` | 0x6008C178 | 8 |  | - |
| `DS_Y_MEM[377]` | 0x6008C179 | 8 |  | - |
| `DS_Y_MEM[378]` | 0x6008C17A | 8 |  | - |
| `DS_Y_MEM[379]` | 0x6008C17B | 8 |  | - |
| `DS_Y_MEM[37]` | 0x6008C025 | 8 |  | - |
| `DS_Y_MEM[380]` | 0x6008C17C | 8 |  | - |
| `DS_Y_MEM[381]` | 0x6008C17D | 8 |  | - |
| `DS_Y_MEM[382]` | 0x6008C17E | 8 |  | - |
| `DS_Y_MEM[383]` | 0x6008C17F | 8 |  | - |
| `DS_Y_MEM[384]` | 0x6008C180 | 8 |  | - |
| `DS_Y_MEM[385]` | 0x6008C181 | 8 |  | - |
| `DS_Y_MEM[386]` | 0x6008C182 | 8 |  | - |
| `DS_Y_MEM[387]` | 0x6008C183 | 8 |  | - |
| `DS_Y_MEM[388]` | 0x6008C184 | 8 |  | - |
| `DS_Y_MEM[389]` | 0x6008C185 | 8 |  | - |
| `DS_Y_MEM[38]` | 0x6008C026 | 8 |  | - |
| `DS_Y_MEM[390]` | 0x6008C186 | 8 |  | - |
| `DS_Y_MEM[391]` | 0x6008C187 | 8 |  | - |
| `DS_Y_MEM[392]` | 0x6008C188 | 8 |  | - |
| `DS_Y_MEM[393]` | 0x6008C189 | 8 |  | - |
| `DS_Y_MEM[394]` | 0x6008C18A | 8 |  | - |
| `DS_Y_MEM[395]` | 0x6008C18B | 8 |  | - |
| `DS_Y_MEM[396]` | 0x6008C18C | 8 |  | - |
| `DS_Y_MEM[397]` | 0x6008C18D | 8 |  | - |
| `DS_Y_MEM[398]` | 0x6008C18E | 8 |  | - |
| `DS_Y_MEM[399]` | 0x6008C18F | 8 |  | - |
| `DS_Y_MEM[39]` | 0x6008C027 | 8 |  | - |
| `DS_Y_MEM[3]` | 0x6008C003 | 8 |  | - |
| `DS_Y_MEM[400]` | 0x6008C190 | 8 |  | - |
| `DS_Y_MEM[401]` | 0x6008C191 | 8 |  | - |
| `DS_Y_MEM[402]` | 0x6008C192 | 8 |  | - |
| `DS_Y_MEM[403]` | 0x6008C193 | 8 |  | - |
| `DS_Y_MEM[404]` | 0x6008C194 | 8 |  | - |
| `DS_Y_MEM[405]` | 0x6008C195 | 8 |  | - |
| `DS_Y_MEM[406]` | 0x6008C196 | 8 |  | - |
| `DS_Y_MEM[407]` | 0x6008C197 | 8 |  | - |
| `DS_Y_MEM[408]` | 0x6008C198 | 8 |  | - |
| `DS_Y_MEM[409]` | 0x6008C199 | 8 |  | - |
| `DS_Y_MEM[40]` | 0x6008C028 | 8 |  | - |
| `DS_Y_MEM[410]` | 0x6008C19A | 8 |  | - |
| `DS_Y_MEM[411]` | 0x6008C19B | 8 |  | - |
| `DS_Y_MEM[412]` | 0x6008C19C | 8 |  | - |
| `DS_Y_MEM[413]` | 0x6008C19D | 8 |  | - |
| `DS_Y_MEM[414]` | 0x6008C19E | 8 |  | - |
| `DS_Y_MEM[415]` | 0x6008C19F | 8 |  | - |
| `DS_Y_MEM[416]` | 0x6008C1A0 | 8 |  | - |
| `DS_Y_MEM[417]` | 0x6008C1A1 | 8 |  | - |
| `DS_Y_MEM[418]` | 0x6008C1A2 | 8 |  | - |
| `DS_Y_MEM[419]` | 0x6008C1A3 | 8 |  | - |
| `DS_Y_MEM[41]` | 0x6008C029 | 8 |  | - |
| `DS_Y_MEM[420]` | 0x6008C1A4 | 8 |  | - |
| `DS_Y_MEM[421]` | 0x6008C1A5 | 8 |  | - |
| `DS_Y_MEM[422]` | 0x6008C1A6 | 8 |  | - |
| `DS_Y_MEM[423]` | 0x6008C1A7 | 8 |  | - |
| `DS_Y_MEM[424]` | 0x6008C1A8 | 8 |  | - |
| `DS_Y_MEM[425]` | 0x6008C1A9 | 8 |  | - |
| `DS_Y_MEM[426]` | 0x6008C1AA | 8 |  | - |
| `DS_Y_MEM[427]` | 0x6008C1AB | 8 |  | - |
| `DS_Y_MEM[428]` | 0x6008C1AC | 8 |  | - |
| `DS_Y_MEM[429]` | 0x6008C1AD | 8 |  | - |
| `DS_Y_MEM[42]` | 0x6008C02A | 8 |  | - |
| `DS_Y_MEM[430]` | 0x6008C1AE | 8 |  | - |
| `DS_Y_MEM[431]` | 0x6008C1AF | 8 |  | - |
| `DS_Y_MEM[432]` | 0x6008C1B0 | 8 |  | - |
| `DS_Y_MEM[433]` | 0x6008C1B1 | 8 |  | - |
| `DS_Y_MEM[434]` | 0x6008C1B2 | 8 |  | - |
| `DS_Y_MEM[435]` | 0x6008C1B3 | 8 |  | - |
| `DS_Y_MEM[436]` | 0x6008C1B4 | 8 |  | - |
| `DS_Y_MEM[437]` | 0x6008C1B5 | 8 |  | - |
| `DS_Y_MEM[438]` | 0x6008C1B6 | 8 |  | - |
| `DS_Y_MEM[439]` | 0x6008C1B7 | 8 |  | - |
| `DS_Y_MEM[43]` | 0x6008C02B | 8 |  | - |
| `DS_Y_MEM[440]` | 0x6008C1B8 | 8 |  | - |
| `DS_Y_MEM[441]` | 0x6008C1B9 | 8 |  | - |
| `DS_Y_MEM[442]` | 0x6008C1BA | 8 |  | - |
| `DS_Y_MEM[443]` | 0x6008C1BB | 8 |  | - |
| `DS_Y_MEM[444]` | 0x6008C1BC | 8 |  | - |
| `DS_Y_MEM[445]` | 0x6008C1BD | 8 |  | - |
| `DS_Y_MEM[446]` | 0x6008C1BE | 8 |  | - |
| `DS_Y_MEM[447]` | 0x6008C1BF | 8 |  | - |
| `DS_Y_MEM[448]` | 0x6008C1C0 | 8 |  | - |
| `DS_Y_MEM[449]` | 0x6008C1C1 | 8 |  | - |
| `DS_Y_MEM[44]` | 0x6008C02C | 8 |  | - |
| `DS_Y_MEM[450]` | 0x6008C1C2 | 8 |  | - |
| `DS_Y_MEM[451]` | 0x6008C1C3 | 8 |  | - |
| `DS_Y_MEM[452]` | 0x6008C1C4 | 8 |  | - |
| `DS_Y_MEM[453]` | 0x6008C1C5 | 8 |  | - |
| `DS_Y_MEM[454]` | 0x6008C1C6 | 8 |  | - |
| `DS_Y_MEM[455]` | 0x6008C1C7 | 8 |  | - |
| `DS_Y_MEM[456]` | 0x6008C1C8 | 8 |  | - |
| `DS_Y_MEM[457]` | 0x6008C1C9 | 8 |  | - |
| `DS_Y_MEM[458]` | 0x6008C1CA | 8 |  | - |
| `DS_Y_MEM[459]` | 0x6008C1CB | 8 |  | - |
| `DS_Y_MEM[45]` | 0x6008C02D | 8 |  | - |
| `DS_Y_MEM[460]` | 0x6008C1CC | 8 |  | - |
| `DS_Y_MEM[461]` | 0x6008C1CD | 8 |  | - |
| `DS_Y_MEM[462]` | 0x6008C1CE | 8 |  | - |
| `DS_Y_MEM[463]` | 0x6008C1CF | 8 |  | - |
| `DS_Y_MEM[464]` | 0x6008C1D0 | 8 |  | - |
| `DS_Y_MEM[465]` | 0x6008C1D1 | 8 |  | - |
| `DS_Y_MEM[466]` | 0x6008C1D2 | 8 |  | - |
| `DS_Y_MEM[467]` | 0x6008C1D3 | 8 |  | - |
| `DS_Y_MEM[468]` | 0x6008C1D4 | 8 |  | - |
| `DS_Y_MEM[469]` | 0x6008C1D5 | 8 |  | - |
| `DS_Y_MEM[46]` | 0x6008C02E | 8 |  | - |
| `DS_Y_MEM[470]` | 0x6008C1D6 | 8 |  | - |
| `DS_Y_MEM[471]` | 0x6008C1D7 | 8 |  | - |
| `DS_Y_MEM[472]` | 0x6008C1D8 | 8 |  | - |
| `DS_Y_MEM[473]` | 0x6008C1D9 | 8 |  | - |
| `DS_Y_MEM[474]` | 0x6008C1DA | 8 |  | - |
| `DS_Y_MEM[475]` | 0x6008C1DB | 8 |  | - |
| `DS_Y_MEM[476]` | 0x6008C1DC | 8 |  | - |
| `DS_Y_MEM[477]` | 0x6008C1DD | 8 |  | - |
| `DS_Y_MEM[478]` | 0x6008C1DE | 8 |  | - |
| `DS_Y_MEM[479]` | 0x6008C1DF | 8 |  | - |
| `DS_Y_MEM[47]` | 0x6008C02F | 8 |  | - |
| `DS_Y_MEM[480]` | 0x6008C1E0 | 8 |  | - |
| `DS_Y_MEM[481]` | 0x6008C1E1 | 8 |  | - |
| `DS_Y_MEM[482]` | 0x6008C1E2 | 8 |  | - |
| `DS_Y_MEM[483]` | 0x6008C1E3 | 8 |  | - |
| `DS_Y_MEM[484]` | 0x6008C1E4 | 8 |  | - |
| `DS_Y_MEM[485]` | 0x6008C1E5 | 8 |  | - |
| `DS_Y_MEM[486]` | 0x6008C1E6 | 8 |  | - |
| `DS_Y_MEM[487]` | 0x6008C1E7 | 8 |  | - |
| `DS_Y_MEM[488]` | 0x6008C1E8 | 8 |  | - |
| `DS_Y_MEM[489]` | 0x6008C1E9 | 8 |  | - |
| `DS_Y_MEM[48]` | 0x6008C030 | 8 |  | - |
| `DS_Y_MEM[490]` | 0x6008C1EA | 8 |  | - |
| `DS_Y_MEM[491]` | 0x6008C1EB | 8 |  | - |
| `DS_Y_MEM[492]` | 0x6008C1EC | 8 |  | - |
| `DS_Y_MEM[493]` | 0x6008C1ED | 8 |  | - |
| `DS_Y_MEM[494]` | 0x6008C1EE | 8 |  | - |
| `DS_Y_MEM[495]` | 0x6008C1EF | 8 |  | - |
| `DS_Y_MEM[496]` | 0x6008C1F0 | 8 |  | - |
| `DS_Y_MEM[497]` | 0x6008C1F1 | 8 |  | - |
| `DS_Y_MEM[498]` | 0x6008C1F2 | 8 |  | - |
| `DS_Y_MEM[499]` | 0x6008C1F3 | 8 |  | - |
| `DS_Y_MEM[49]` | 0x6008C031 | 8 |  | - |
| `DS_Y_MEM[4]` | 0x6008C004 | 8 |  | - |
| `DS_Y_MEM[500]` | 0x6008C1F4 | 8 |  | - |
| `DS_Y_MEM[501]` | 0x6008C1F5 | 8 |  | - |
| `DS_Y_MEM[502]` | 0x6008C1F6 | 8 |  | - |
| `DS_Y_MEM[503]` | 0x6008C1F7 | 8 |  | - |
| `DS_Y_MEM[504]` | 0x6008C1F8 | 8 |  | - |
| `DS_Y_MEM[505]` | 0x6008C1F9 | 8 |  | - |
| `DS_Y_MEM[506]` | 0x6008C1FA | 8 |  | - |
| `DS_Y_MEM[507]` | 0x6008C1FB | 8 |  | - |
| `DS_Y_MEM[508]` | 0x6008C1FC | 8 |  | - |
| `DS_Y_MEM[509]` | 0x6008C1FD | 8 |  | - |
| `DS_Y_MEM[50]` | 0x6008C032 | 8 |  | - |
| `DS_Y_MEM[510]` | 0x6008C1FE | 8 |  | - |
| `DS_Y_MEM[511]` | 0x6008C1FF | 8 |  | - |
| `DS_Y_MEM[51]` | 0x6008C033 | 8 |  | - |
| `DS_Y_MEM[52]` | 0x6008C034 | 8 |  | - |
| `DS_Y_MEM[53]` | 0x6008C035 | 8 |  | - |
| `DS_Y_MEM[54]` | 0x6008C036 | 8 |  | - |
| `DS_Y_MEM[55]` | 0x6008C037 | 8 |  | - |
| `DS_Y_MEM[56]` | 0x6008C038 | 8 |  | - |
| `DS_Y_MEM[57]` | 0x6008C039 | 8 |  | - |
| `DS_Y_MEM[58]` | 0x6008C03A | 8 |  | - |
| `DS_Y_MEM[59]` | 0x6008C03B | 8 |  | - |
| `DS_Y_MEM[5]` | 0x6008C005 | 8 |  | - |
| `DS_Y_MEM[60]` | 0x6008C03C | 8 |  | - |
| `DS_Y_MEM[61]` | 0x6008C03D | 8 |  | - |
| `DS_Y_MEM[62]` | 0x6008C03E | 8 |  | - |
| `DS_Y_MEM[63]` | 0x6008C03F | 8 |  | - |
| `DS_Y_MEM[64]` | 0x6008C040 | 8 |  | - |
| `DS_Y_MEM[65]` | 0x6008C041 | 8 |  | - |
| `DS_Y_MEM[66]` | 0x6008C042 | 8 |  | - |
| `DS_Y_MEM[67]` | 0x6008C043 | 8 |  | - |
| `DS_Y_MEM[68]` | 0x6008C044 | 8 |  | - |
| `DS_Y_MEM[69]` | 0x6008C045 | 8 |  | - |
| `DS_Y_MEM[6]` | 0x6008C006 | 8 |  | - |
| `DS_Y_MEM[70]` | 0x6008C046 | 8 |  | - |
| `DS_Y_MEM[71]` | 0x6008C047 | 8 |  | - |
| `DS_Y_MEM[72]` | 0x6008C048 | 8 |  | - |
| `DS_Y_MEM[73]` | 0x6008C049 | 8 |  | - |
| `DS_Y_MEM[74]` | 0x6008C04A | 8 |  | - |
| `DS_Y_MEM[75]` | 0x6008C04B | 8 |  | - |
| `DS_Y_MEM[76]` | 0x6008C04C | 8 |  | - |
| `DS_Y_MEM[77]` | 0x6008C04D | 8 |  | - |
| `DS_Y_MEM[78]` | 0x6008C04E | 8 |  | - |
| `DS_Y_MEM[79]` | 0x6008C04F | 8 |  | - |
| `DS_Y_MEM[7]` | 0x6008C007 | 8 |  | - |
| `DS_Y_MEM[80]` | 0x6008C050 | 8 |  | - |
| `DS_Y_MEM[81]` | 0x6008C051 | 8 |  | - |
| `DS_Y_MEM[82]` | 0x6008C052 | 8 |  | - |
| `DS_Y_MEM[83]` | 0x6008C053 | 8 |  | - |
| `DS_Y_MEM[84]` | 0x6008C054 | 8 |  | - |
| `DS_Y_MEM[85]` | 0x6008C055 | 8 |  | - |
| `DS_Y_MEM[86]` | 0x6008C056 | 8 |  | - |
| `DS_Y_MEM[87]` | 0x6008C057 | 8 |  | - |
| `DS_Y_MEM[88]` | 0x6008C058 | 8 |  | - |
| `DS_Y_MEM[89]` | 0x6008C059 | 8 |  | - |
| `DS_Y_MEM[8]` | 0x6008C008 | 8 |  | - |
| `DS_Y_MEM[90]` | 0x6008C05A | 8 |  | - |
| `DS_Y_MEM[91]` | 0x6008C05B | 8 |  | - |
| `DS_Y_MEM[92]` | 0x6008C05C | 8 |  | - |
| `DS_Y_MEM[93]` | 0x6008C05D | 8 |  | - |
| `DS_Y_MEM[94]` | 0x6008C05E | 8 |  | - |
| `DS_Y_MEM[95]` | 0x6008C05F | 8 |  | - |
| `DS_Y_MEM[96]` | 0x6008C060 | 8 |  | - |
| `DS_Y_MEM[97]` | 0x6008C061 | 8 |  | - |
| `DS_Y_MEM[98]` | 0x6008C062 | 8 |  | - |
| `DS_Y_MEM[99]` | 0x6008C063 | 8 |  | - |
| `DS_Y_MEM[9]` | 0x6008C009 | 8 |  | - |
| `DS_Z_MEM[0]` | 0x6008CA00 | 8 |  | - |
| `DS_Z_MEM[100]` | 0x6008CA64 | 8 |  | - |
| `DS_Z_MEM[101]` | 0x6008CA65 | 8 |  | - |
| `DS_Z_MEM[102]` | 0x6008CA66 | 8 |  | - |
| `DS_Z_MEM[103]` | 0x6008CA67 | 8 |  | - |
| `DS_Z_MEM[104]` | 0x6008CA68 | 8 |  | - |
| `DS_Z_MEM[105]` | 0x6008CA69 | 8 |  | - |
| `DS_Z_MEM[106]` | 0x6008CA6A | 8 |  | - |
| `DS_Z_MEM[107]` | 0x6008CA6B | 8 |  | - |
| `DS_Z_MEM[108]` | 0x6008CA6C | 8 |  | - |
| `DS_Z_MEM[109]` | 0x6008CA6D | 8 |  | - |
| `DS_Z_MEM[10]` | 0x6008CA0A | 8 |  | - |
| `DS_Z_MEM[110]` | 0x6008CA6E | 8 |  | - |
| `DS_Z_MEM[111]` | 0x6008CA6F | 8 |  | - |
| `DS_Z_MEM[112]` | 0x6008CA70 | 8 |  | - |
| `DS_Z_MEM[113]` | 0x6008CA71 | 8 |  | - |
| `DS_Z_MEM[114]` | 0x6008CA72 | 8 |  | - |
| `DS_Z_MEM[115]` | 0x6008CA73 | 8 |  | - |
| `DS_Z_MEM[116]` | 0x6008CA74 | 8 |  | - |
| `DS_Z_MEM[117]` | 0x6008CA75 | 8 |  | - |
| `DS_Z_MEM[118]` | 0x6008CA76 | 8 |  | - |
| `DS_Z_MEM[119]` | 0x6008CA77 | 8 |  | - |
| `DS_Z_MEM[11]` | 0x6008CA0B | 8 |  | - |
| `DS_Z_MEM[120]` | 0x6008CA78 | 8 |  | - |
| `DS_Z_MEM[121]` | 0x6008CA79 | 8 |  | - |
| `DS_Z_MEM[122]` | 0x6008CA7A | 8 |  | - |
| `DS_Z_MEM[123]` | 0x6008CA7B | 8 |  | - |
| `DS_Z_MEM[124]` | 0x6008CA7C | 8 |  | - |
| `DS_Z_MEM[125]` | 0x6008CA7D | 8 |  | - |
| `DS_Z_MEM[126]` | 0x6008CA7E | 8 |  | - |
| `DS_Z_MEM[127]` | 0x6008CA7F | 8 |  | - |
| `DS_Z_MEM[128]` | 0x6008CA80 | 8 |  | - |
| `DS_Z_MEM[129]` | 0x6008CA81 | 8 |  | - |
| `DS_Z_MEM[12]` | 0x6008CA0C | 8 |  | - |
| `DS_Z_MEM[130]` | 0x6008CA82 | 8 |  | - |
| `DS_Z_MEM[131]` | 0x6008CA83 | 8 |  | - |
| `DS_Z_MEM[132]` | 0x6008CA84 | 8 |  | - |
| `DS_Z_MEM[133]` | 0x6008CA85 | 8 |  | - |
| `DS_Z_MEM[134]` | 0x6008CA86 | 8 |  | - |
| `DS_Z_MEM[135]` | 0x6008CA87 | 8 |  | - |
| `DS_Z_MEM[136]` | 0x6008CA88 | 8 |  | - |
| `DS_Z_MEM[137]` | 0x6008CA89 | 8 |  | - |
| `DS_Z_MEM[138]` | 0x6008CA8A | 8 |  | - |
| `DS_Z_MEM[139]` | 0x6008CA8B | 8 |  | - |
| `DS_Z_MEM[13]` | 0x6008CA0D | 8 |  | - |
| `DS_Z_MEM[140]` | 0x6008CA8C | 8 |  | - |
| `DS_Z_MEM[141]` | 0x6008CA8D | 8 |  | - |
| `DS_Z_MEM[142]` | 0x6008CA8E | 8 |  | - |
| `DS_Z_MEM[143]` | 0x6008CA8F | 8 |  | - |
| `DS_Z_MEM[144]` | 0x6008CA90 | 8 |  | - |
| `DS_Z_MEM[145]` | 0x6008CA91 | 8 |  | - |
| `DS_Z_MEM[146]` | 0x6008CA92 | 8 |  | - |
| `DS_Z_MEM[147]` | 0x6008CA93 | 8 |  | - |
| `DS_Z_MEM[148]` | 0x6008CA94 | 8 |  | - |
| `DS_Z_MEM[149]` | 0x6008CA95 | 8 |  | - |
| `DS_Z_MEM[14]` | 0x6008CA0E | 8 |  | - |
| `DS_Z_MEM[150]` | 0x6008CA96 | 8 |  | - |
| `DS_Z_MEM[151]` | 0x6008CA97 | 8 |  | - |
| `DS_Z_MEM[152]` | 0x6008CA98 | 8 |  | - |
| `DS_Z_MEM[153]` | 0x6008CA99 | 8 |  | - |
| `DS_Z_MEM[154]` | 0x6008CA9A | 8 |  | - |
| `DS_Z_MEM[155]` | 0x6008CA9B | 8 |  | - |
| `DS_Z_MEM[156]` | 0x6008CA9C | 8 |  | - |
| `DS_Z_MEM[157]` | 0x6008CA9D | 8 |  | - |
| `DS_Z_MEM[158]` | 0x6008CA9E | 8 |  | - |
| `DS_Z_MEM[159]` | 0x6008CA9F | 8 |  | - |
| `DS_Z_MEM[15]` | 0x6008CA0F | 8 |  | - |
| `DS_Z_MEM[160]` | 0x6008CAA0 | 8 |  | - |
| `DS_Z_MEM[161]` | 0x6008CAA1 | 8 |  | - |
| `DS_Z_MEM[162]` | 0x6008CAA2 | 8 |  | - |
| `DS_Z_MEM[163]` | 0x6008CAA3 | 8 |  | - |
| `DS_Z_MEM[164]` | 0x6008CAA4 | 8 |  | - |
| `DS_Z_MEM[165]` | 0x6008CAA5 | 8 |  | - |
| `DS_Z_MEM[166]` | 0x6008CAA6 | 8 |  | - |
| `DS_Z_MEM[167]` | 0x6008CAA7 | 8 |  | - |
| `DS_Z_MEM[168]` | 0x6008CAA8 | 8 |  | - |
| `DS_Z_MEM[169]` | 0x6008CAA9 | 8 |  | - |
| `DS_Z_MEM[16]` | 0x6008CA10 | 8 |  | - |
| `DS_Z_MEM[170]` | 0x6008CAAA | 8 |  | - |
| `DS_Z_MEM[171]` | 0x6008CAAB | 8 |  | - |
| `DS_Z_MEM[172]` | 0x6008CAAC | 8 |  | - |
| `DS_Z_MEM[173]` | 0x6008CAAD | 8 |  | - |
| `DS_Z_MEM[174]` | 0x6008CAAE | 8 |  | - |
| `DS_Z_MEM[175]` | 0x6008CAAF | 8 |  | - |
| `DS_Z_MEM[176]` | 0x6008CAB0 | 8 |  | - |
| `DS_Z_MEM[177]` | 0x6008CAB1 | 8 |  | - |
| `DS_Z_MEM[178]` | 0x6008CAB2 | 8 |  | - |
| `DS_Z_MEM[179]` | 0x6008CAB3 | 8 |  | - |
| `DS_Z_MEM[17]` | 0x6008CA11 | 8 |  | - |
| `DS_Z_MEM[180]` | 0x6008CAB4 | 8 |  | - |
| `DS_Z_MEM[181]` | 0x6008CAB5 | 8 |  | - |
| `DS_Z_MEM[182]` | 0x6008CAB6 | 8 |  | - |
| `DS_Z_MEM[183]` | 0x6008CAB7 | 8 |  | - |
| `DS_Z_MEM[184]` | 0x6008CAB8 | 8 |  | - |
| `DS_Z_MEM[185]` | 0x6008CAB9 | 8 |  | - |
| `DS_Z_MEM[186]` | 0x6008CABA | 8 |  | - |
| `DS_Z_MEM[187]` | 0x6008CABB | 8 |  | - |
| `DS_Z_MEM[188]` | 0x6008CABC | 8 |  | - |
| `DS_Z_MEM[189]` | 0x6008CABD | 8 |  | - |
| `DS_Z_MEM[18]` | 0x6008CA12 | 8 |  | - |
| `DS_Z_MEM[190]` | 0x6008CABE | 8 |  | - |
| `DS_Z_MEM[191]` | 0x6008CABF | 8 |  | - |
| `DS_Z_MEM[192]` | 0x6008CAC0 | 8 |  | - |
| `DS_Z_MEM[193]` | 0x6008CAC1 | 8 |  | - |
| `DS_Z_MEM[194]` | 0x6008CAC2 | 8 |  | - |
| `DS_Z_MEM[195]` | 0x6008CAC3 | 8 |  | - |
| `DS_Z_MEM[196]` | 0x6008CAC4 | 8 |  | - |
| `DS_Z_MEM[197]` | 0x6008CAC5 | 8 |  | - |
| `DS_Z_MEM[198]` | 0x6008CAC6 | 8 |  | - |
| `DS_Z_MEM[199]` | 0x6008CAC7 | 8 |  | - |
| `DS_Z_MEM[19]` | 0x6008CA13 | 8 |  | - |
| `DS_Z_MEM[1]` | 0x6008CA01 | 8 |  | - |
| `DS_Z_MEM[200]` | 0x6008CAC8 | 8 |  | - |
| `DS_Z_MEM[201]` | 0x6008CAC9 | 8 |  | - |
| `DS_Z_MEM[202]` | 0x6008CACA | 8 |  | - |
| `DS_Z_MEM[203]` | 0x6008CACB | 8 |  | - |
| `DS_Z_MEM[204]` | 0x6008CACC | 8 |  | - |
| `DS_Z_MEM[205]` | 0x6008CACD | 8 |  | - |
| `DS_Z_MEM[206]` | 0x6008CACE | 8 |  | - |
| `DS_Z_MEM[207]` | 0x6008CACF | 8 |  | - |
| `DS_Z_MEM[208]` | 0x6008CAD0 | 8 |  | - |
| `DS_Z_MEM[209]` | 0x6008CAD1 | 8 |  | - |
| `DS_Z_MEM[20]` | 0x6008CA14 | 8 |  | - |
| `DS_Z_MEM[210]` | 0x6008CAD2 | 8 |  | - |
| `DS_Z_MEM[211]` | 0x6008CAD3 | 8 |  | - |
| `DS_Z_MEM[212]` | 0x6008CAD4 | 8 |  | - |
| `DS_Z_MEM[213]` | 0x6008CAD5 | 8 |  | - |
| `DS_Z_MEM[214]` | 0x6008CAD6 | 8 |  | - |
| `DS_Z_MEM[215]` | 0x6008CAD7 | 8 |  | - |
| `DS_Z_MEM[216]` | 0x6008CAD8 | 8 |  | - |
| `DS_Z_MEM[217]` | 0x6008CAD9 | 8 |  | - |
| `DS_Z_MEM[218]` | 0x6008CADA | 8 |  | - |
| `DS_Z_MEM[219]` | 0x6008CADB | 8 |  | - |
| `DS_Z_MEM[21]` | 0x6008CA15 | 8 |  | - |
| `DS_Z_MEM[220]` | 0x6008CADC | 8 |  | - |
| `DS_Z_MEM[221]` | 0x6008CADD | 8 |  | - |
| `DS_Z_MEM[222]` | 0x6008CADE | 8 |  | - |
| `DS_Z_MEM[223]` | 0x6008CADF | 8 |  | - |
| `DS_Z_MEM[224]` | 0x6008CAE0 | 8 |  | - |
| `DS_Z_MEM[225]` | 0x6008CAE1 | 8 |  | - |
| `DS_Z_MEM[226]` | 0x6008CAE2 | 8 |  | - |
| `DS_Z_MEM[227]` | 0x6008CAE3 | 8 |  | - |
| `DS_Z_MEM[228]` | 0x6008CAE4 | 8 |  | - |
| `DS_Z_MEM[229]` | 0x6008CAE5 | 8 |  | - |
| `DS_Z_MEM[22]` | 0x6008CA16 | 8 |  | - |
| `DS_Z_MEM[230]` | 0x6008CAE6 | 8 |  | - |
| `DS_Z_MEM[231]` | 0x6008CAE7 | 8 |  | - |
| `DS_Z_MEM[232]` | 0x6008CAE8 | 8 |  | - |
| `DS_Z_MEM[233]` | 0x6008CAE9 | 8 |  | - |
| `DS_Z_MEM[234]` | 0x6008CAEA | 8 |  | - |
| `DS_Z_MEM[235]` | 0x6008CAEB | 8 |  | - |
| `DS_Z_MEM[236]` | 0x6008CAEC | 8 |  | - |
| `DS_Z_MEM[237]` | 0x6008CAED | 8 |  | - |
| `DS_Z_MEM[238]` | 0x6008CAEE | 8 |  | - |
| `DS_Z_MEM[239]` | 0x6008CAEF | 8 |  | - |
| `DS_Z_MEM[23]` | 0x6008CA17 | 8 |  | - |
| `DS_Z_MEM[240]` | 0x6008CAF0 | 8 |  | - |
| `DS_Z_MEM[241]` | 0x6008CAF1 | 8 |  | - |
| `DS_Z_MEM[242]` | 0x6008CAF2 | 8 |  | - |
| `DS_Z_MEM[243]` | 0x6008CAF3 | 8 |  | - |
| `DS_Z_MEM[244]` | 0x6008CAF4 | 8 |  | - |
| `DS_Z_MEM[245]` | 0x6008CAF5 | 8 |  | - |
| `DS_Z_MEM[246]` | 0x6008CAF6 | 8 |  | - |
| `DS_Z_MEM[247]` | 0x6008CAF7 | 8 |  | - |
| `DS_Z_MEM[248]` | 0x6008CAF8 | 8 |  | - |
| `DS_Z_MEM[249]` | 0x6008CAF9 | 8 |  | - |
| `DS_Z_MEM[24]` | 0x6008CA18 | 8 |  | - |
| `DS_Z_MEM[250]` | 0x6008CAFA | 8 |  | - |
| `DS_Z_MEM[251]` | 0x6008CAFB | 8 |  | - |
| `DS_Z_MEM[252]` | 0x6008CAFC | 8 |  | - |
| `DS_Z_MEM[253]` | 0x6008CAFD | 8 |  | - |
| `DS_Z_MEM[254]` | 0x6008CAFE | 8 |  | - |
| `DS_Z_MEM[255]` | 0x6008CAFF | 8 |  | - |
| `DS_Z_MEM[256]` | 0x6008CB00 | 8 |  | - |
| `DS_Z_MEM[257]` | 0x6008CB01 | 8 |  | - |
| `DS_Z_MEM[258]` | 0x6008CB02 | 8 |  | - |
| `DS_Z_MEM[259]` | 0x6008CB03 | 8 |  | - |
| `DS_Z_MEM[25]` | 0x6008CA19 | 8 |  | - |
| `DS_Z_MEM[260]` | 0x6008CB04 | 8 |  | - |
| `DS_Z_MEM[261]` | 0x6008CB05 | 8 |  | - |
| `DS_Z_MEM[262]` | 0x6008CB06 | 8 |  | - |
| `DS_Z_MEM[263]` | 0x6008CB07 | 8 |  | - |
| `DS_Z_MEM[264]` | 0x6008CB08 | 8 |  | - |
| `DS_Z_MEM[265]` | 0x6008CB09 | 8 |  | - |
| `DS_Z_MEM[266]` | 0x6008CB0A | 8 |  | - |
| `DS_Z_MEM[267]` | 0x6008CB0B | 8 |  | - |
| `DS_Z_MEM[268]` | 0x6008CB0C | 8 |  | - |
| `DS_Z_MEM[269]` | 0x6008CB0D | 8 |  | - |
| `DS_Z_MEM[26]` | 0x6008CA1A | 8 |  | - |
| `DS_Z_MEM[270]` | 0x6008CB0E | 8 |  | - |
| `DS_Z_MEM[271]` | 0x6008CB0F | 8 |  | - |
| `DS_Z_MEM[272]` | 0x6008CB10 | 8 |  | - |
| `DS_Z_MEM[273]` | 0x6008CB11 | 8 |  | - |
| `DS_Z_MEM[274]` | 0x6008CB12 | 8 |  | - |
| `DS_Z_MEM[275]` | 0x6008CB13 | 8 |  | - |
| `DS_Z_MEM[276]` | 0x6008CB14 | 8 |  | - |
| `DS_Z_MEM[277]` | 0x6008CB15 | 8 |  | - |
| `DS_Z_MEM[278]` | 0x6008CB16 | 8 |  | - |
| `DS_Z_MEM[279]` | 0x6008CB17 | 8 |  | - |
| `DS_Z_MEM[27]` | 0x6008CA1B | 8 |  | - |
| `DS_Z_MEM[280]` | 0x6008CB18 | 8 |  | - |
| `DS_Z_MEM[281]` | 0x6008CB19 | 8 |  | - |
| `DS_Z_MEM[282]` | 0x6008CB1A | 8 |  | - |
| `DS_Z_MEM[283]` | 0x6008CB1B | 8 |  | - |
| `DS_Z_MEM[284]` | 0x6008CB1C | 8 |  | - |
| `DS_Z_MEM[285]` | 0x6008CB1D | 8 |  | - |
| `DS_Z_MEM[286]` | 0x6008CB1E | 8 |  | - |
| `DS_Z_MEM[287]` | 0x6008CB1F | 8 |  | - |
| `DS_Z_MEM[288]` | 0x6008CB20 | 8 |  | - |
| `DS_Z_MEM[289]` | 0x6008CB21 | 8 |  | - |
| `DS_Z_MEM[28]` | 0x6008CA1C | 8 |  | - |
| `DS_Z_MEM[290]` | 0x6008CB22 | 8 |  | - |
| `DS_Z_MEM[291]` | 0x6008CB23 | 8 |  | - |
| `DS_Z_MEM[292]` | 0x6008CB24 | 8 |  | - |
| `DS_Z_MEM[293]` | 0x6008CB25 | 8 |  | - |
| `DS_Z_MEM[294]` | 0x6008CB26 | 8 |  | - |
| `DS_Z_MEM[295]` | 0x6008CB27 | 8 |  | - |
| `DS_Z_MEM[296]` | 0x6008CB28 | 8 |  | - |
| `DS_Z_MEM[297]` | 0x6008CB29 | 8 |  | - |
| `DS_Z_MEM[298]` | 0x6008CB2A | 8 |  | - |
| `DS_Z_MEM[299]` | 0x6008CB2B | 8 |  | - |
| `DS_Z_MEM[29]` | 0x6008CA1D | 8 |  | - |
| `DS_Z_MEM[2]` | 0x6008CA02 | 8 |  | - |
| `DS_Z_MEM[300]` | 0x6008CB2C | 8 |  | - |
| `DS_Z_MEM[301]` | 0x6008CB2D | 8 |  | - |
| `DS_Z_MEM[302]` | 0x6008CB2E | 8 |  | - |
| `DS_Z_MEM[303]` | 0x6008CB2F | 8 |  | - |
| `DS_Z_MEM[304]` | 0x6008CB30 | 8 |  | - |
| `DS_Z_MEM[305]` | 0x6008CB31 | 8 |  | - |
| `DS_Z_MEM[306]` | 0x6008CB32 | 8 |  | - |
| `DS_Z_MEM[307]` | 0x6008CB33 | 8 |  | - |
| `DS_Z_MEM[308]` | 0x6008CB34 | 8 |  | - |
| `DS_Z_MEM[309]` | 0x6008CB35 | 8 |  | - |
| `DS_Z_MEM[30]` | 0x6008CA1E | 8 |  | - |
| `DS_Z_MEM[310]` | 0x6008CB36 | 8 |  | - |
| `DS_Z_MEM[311]` | 0x6008CB37 | 8 |  | - |
| `DS_Z_MEM[312]` | 0x6008CB38 | 8 |  | - |
| `DS_Z_MEM[313]` | 0x6008CB39 | 8 |  | - |
| `DS_Z_MEM[314]` | 0x6008CB3A | 8 |  | - |
| `DS_Z_MEM[315]` | 0x6008CB3B | 8 |  | - |
| `DS_Z_MEM[316]` | 0x6008CB3C | 8 |  | - |
| `DS_Z_MEM[317]` | 0x6008CB3D | 8 |  | - |
| `DS_Z_MEM[318]` | 0x6008CB3E | 8 |  | - |
| `DS_Z_MEM[319]` | 0x6008CB3F | 8 |  | - |
| `DS_Z_MEM[31]` | 0x6008CA1F | 8 |  | - |
| `DS_Z_MEM[320]` | 0x6008CB40 | 8 |  | - |
| `DS_Z_MEM[321]` | 0x6008CB41 | 8 |  | - |
| `DS_Z_MEM[322]` | 0x6008CB42 | 8 |  | - |
| `DS_Z_MEM[323]` | 0x6008CB43 | 8 |  | - |
| `DS_Z_MEM[324]` | 0x6008CB44 | 8 |  | - |
| `DS_Z_MEM[325]` | 0x6008CB45 | 8 |  | - |
| `DS_Z_MEM[326]` | 0x6008CB46 | 8 |  | - |
| `DS_Z_MEM[327]` | 0x6008CB47 | 8 |  | - |
| `DS_Z_MEM[328]` | 0x6008CB48 | 8 |  | - |
| `DS_Z_MEM[329]` | 0x6008CB49 | 8 |  | - |
| `DS_Z_MEM[32]` | 0x6008CA20 | 8 |  | - |
| `DS_Z_MEM[330]` | 0x6008CB4A | 8 |  | - |
| `DS_Z_MEM[331]` | 0x6008CB4B | 8 |  | - |
| `DS_Z_MEM[332]` | 0x6008CB4C | 8 |  | - |
| `DS_Z_MEM[333]` | 0x6008CB4D | 8 |  | - |
| `DS_Z_MEM[334]` | 0x6008CB4E | 8 |  | - |
| `DS_Z_MEM[335]` | 0x6008CB4F | 8 |  | - |
| `DS_Z_MEM[336]` | 0x6008CB50 | 8 |  | - |
| `DS_Z_MEM[337]` | 0x6008CB51 | 8 |  | - |
| `DS_Z_MEM[338]` | 0x6008CB52 | 8 |  | - |
| `DS_Z_MEM[339]` | 0x6008CB53 | 8 |  | - |
| `DS_Z_MEM[33]` | 0x6008CA21 | 8 |  | - |
| `DS_Z_MEM[340]` | 0x6008CB54 | 8 |  | - |
| `DS_Z_MEM[341]` | 0x6008CB55 | 8 |  | - |
| `DS_Z_MEM[342]` | 0x6008CB56 | 8 |  | - |
| `DS_Z_MEM[343]` | 0x6008CB57 | 8 |  | - |
| `DS_Z_MEM[344]` | 0x6008CB58 | 8 |  | - |
| `DS_Z_MEM[345]` | 0x6008CB59 | 8 |  | - |
| `DS_Z_MEM[346]` | 0x6008CB5A | 8 |  | - |
| `DS_Z_MEM[347]` | 0x6008CB5B | 8 |  | - |
| `DS_Z_MEM[348]` | 0x6008CB5C | 8 |  | - |
| `DS_Z_MEM[349]` | 0x6008CB5D | 8 |  | - |
| `DS_Z_MEM[34]` | 0x6008CA22 | 8 |  | - |
| `DS_Z_MEM[350]` | 0x6008CB5E | 8 |  | - |
| `DS_Z_MEM[351]` | 0x6008CB5F | 8 |  | - |
| `DS_Z_MEM[352]` | 0x6008CB60 | 8 |  | - |
| `DS_Z_MEM[353]` | 0x6008CB61 | 8 |  | - |
| `DS_Z_MEM[354]` | 0x6008CB62 | 8 |  | - |
| `DS_Z_MEM[355]` | 0x6008CB63 | 8 |  | - |
| `DS_Z_MEM[356]` | 0x6008CB64 | 8 |  | - |
| `DS_Z_MEM[357]` | 0x6008CB65 | 8 |  | - |
| `DS_Z_MEM[358]` | 0x6008CB66 | 8 |  | - |
| `DS_Z_MEM[359]` | 0x6008CB67 | 8 |  | - |
| `DS_Z_MEM[35]` | 0x6008CA23 | 8 |  | - |
| `DS_Z_MEM[360]` | 0x6008CB68 | 8 |  | - |
| `DS_Z_MEM[361]` | 0x6008CB69 | 8 |  | - |
| `DS_Z_MEM[362]` | 0x6008CB6A | 8 |  | - |
| `DS_Z_MEM[363]` | 0x6008CB6B | 8 |  | - |
| `DS_Z_MEM[364]` | 0x6008CB6C | 8 |  | - |
| `DS_Z_MEM[365]` | 0x6008CB6D | 8 |  | - |
| `DS_Z_MEM[366]` | 0x6008CB6E | 8 |  | - |
| `DS_Z_MEM[367]` | 0x6008CB6F | 8 |  | - |
| `DS_Z_MEM[368]` | 0x6008CB70 | 8 |  | - |
| `DS_Z_MEM[369]` | 0x6008CB71 | 8 |  | - |
| `DS_Z_MEM[36]` | 0x6008CA24 | 8 |  | - |
| `DS_Z_MEM[370]` | 0x6008CB72 | 8 |  | - |
| `DS_Z_MEM[371]` | 0x6008CB73 | 8 |  | - |
| `DS_Z_MEM[372]` | 0x6008CB74 | 8 |  | - |
| `DS_Z_MEM[373]` | 0x6008CB75 | 8 |  | - |
| `DS_Z_MEM[374]` | 0x6008CB76 | 8 |  | - |
| `DS_Z_MEM[375]` | 0x6008CB77 | 8 |  | - |
| `DS_Z_MEM[376]` | 0x6008CB78 | 8 |  | - |
| `DS_Z_MEM[377]` | 0x6008CB79 | 8 |  | - |
| `DS_Z_MEM[378]` | 0x6008CB7A | 8 |  | - |
| `DS_Z_MEM[379]` | 0x6008CB7B | 8 |  | - |
| `DS_Z_MEM[37]` | 0x6008CA25 | 8 |  | - |
| `DS_Z_MEM[380]` | 0x6008CB7C | 8 |  | - |
| `DS_Z_MEM[381]` | 0x6008CB7D | 8 |  | - |
| `DS_Z_MEM[382]` | 0x6008CB7E | 8 |  | - |
| `DS_Z_MEM[383]` | 0x6008CB7F | 8 |  | - |
| `DS_Z_MEM[384]` | 0x6008CB80 | 8 |  | - |
| `DS_Z_MEM[385]` | 0x6008CB81 | 8 |  | - |
| `DS_Z_MEM[386]` | 0x6008CB82 | 8 |  | - |
| `DS_Z_MEM[387]` | 0x6008CB83 | 8 |  | - |
| `DS_Z_MEM[388]` | 0x6008CB84 | 8 |  | - |
| `DS_Z_MEM[389]` | 0x6008CB85 | 8 |  | - |
| `DS_Z_MEM[38]` | 0x6008CA26 | 8 |  | - |
| `DS_Z_MEM[390]` | 0x6008CB86 | 8 |  | - |
| `DS_Z_MEM[391]` | 0x6008CB87 | 8 |  | - |
| `DS_Z_MEM[392]` | 0x6008CB88 | 8 |  | - |
| `DS_Z_MEM[393]` | 0x6008CB89 | 8 |  | - |
| `DS_Z_MEM[394]` | 0x6008CB8A | 8 |  | - |
| `DS_Z_MEM[395]` | 0x6008CB8B | 8 |  | - |
| `DS_Z_MEM[396]` | 0x6008CB8C | 8 |  | - |
| `DS_Z_MEM[397]` | 0x6008CB8D | 8 |  | - |
| `DS_Z_MEM[398]` | 0x6008CB8E | 8 |  | - |
| `DS_Z_MEM[399]` | 0x6008CB8F | 8 |  | - |
| `DS_Z_MEM[39]` | 0x6008CA27 | 8 |  | - |
| `DS_Z_MEM[3]` | 0x6008CA03 | 8 |  | - |
| `DS_Z_MEM[400]` | 0x6008CB90 | 8 |  | - |
| `DS_Z_MEM[401]` | 0x6008CB91 | 8 |  | - |
| `DS_Z_MEM[402]` | 0x6008CB92 | 8 |  | - |
| `DS_Z_MEM[403]` | 0x6008CB93 | 8 |  | - |
| `DS_Z_MEM[404]` | 0x6008CB94 | 8 |  | - |
| `DS_Z_MEM[405]` | 0x6008CB95 | 8 |  | - |
| `DS_Z_MEM[406]` | 0x6008CB96 | 8 |  | - |
| `DS_Z_MEM[407]` | 0x6008CB97 | 8 |  | - |
| `DS_Z_MEM[408]` | 0x6008CB98 | 8 |  | - |
| `DS_Z_MEM[409]` | 0x6008CB99 | 8 |  | - |
| `DS_Z_MEM[40]` | 0x6008CA28 | 8 |  | - |
| `DS_Z_MEM[410]` | 0x6008CB9A | 8 |  | - |
| `DS_Z_MEM[411]` | 0x6008CB9B | 8 |  | - |
| `DS_Z_MEM[412]` | 0x6008CB9C | 8 |  | - |
| `DS_Z_MEM[413]` | 0x6008CB9D | 8 |  | - |
| `DS_Z_MEM[414]` | 0x6008CB9E | 8 |  | - |
| `DS_Z_MEM[415]` | 0x6008CB9F | 8 |  | - |
| `DS_Z_MEM[416]` | 0x6008CBA0 | 8 |  | - |
| `DS_Z_MEM[417]` | 0x6008CBA1 | 8 |  | - |
| `DS_Z_MEM[418]` | 0x6008CBA2 | 8 |  | - |
| `DS_Z_MEM[419]` | 0x6008CBA3 | 8 |  | - |
| `DS_Z_MEM[41]` | 0x6008CA29 | 8 |  | - |
| `DS_Z_MEM[420]` | 0x6008CBA4 | 8 |  | - |
| `DS_Z_MEM[421]` | 0x6008CBA5 | 8 |  | - |
| `DS_Z_MEM[422]` | 0x6008CBA6 | 8 |  | - |
| `DS_Z_MEM[423]` | 0x6008CBA7 | 8 |  | - |
| `DS_Z_MEM[424]` | 0x6008CBA8 | 8 |  | - |
| `DS_Z_MEM[425]` | 0x6008CBA9 | 8 |  | - |
| `DS_Z_MEM[426]` | 0x6008CBAA | 8 |  | - |
| `DS_Z_MEM[427]` | 0x6008CBAB | 8 |  | - |
| `DS_Z_MEM[428]` | 0x6008CBAC | 8 |  | - |
| `DS_Z_MEM[429]` | 0x6008CBAD | 8 |  | - |
| `DS_Z_MEM[42]` | 0x6008CA2A | 8 |  | - |
| `DS_Z_MEM[430]` | 0x6008CBAE | 8 |  | - |
| `DS_Z_MEM[431]` | 0x6008CBAF | 8 |  | - |
| `DS_Z_MEM[432]` | 0x6008CBB0 | 8 |  | - |
| `DS_Z_MEM[433]` | 0x6008CBB1 | 8 |  | - |
| `DS_Z_MEM[434]` | 0x6008CBB2 | 8 |  | - |
| `DS_Z_MEM[435]` | 0x6008CBB3 | 8 |  | - |
| `DS_Z_MEM[436]` | 0x6008CBB4 | 8 |  | - |
| `DS_Z_MEM[437]` | 0x6008CBB5 | 8 |  | - |
| `DS_Z_MEM[438]` | 0x6008CBB6 | 8 |  | - |
| `DS_Z_MEM[439]` | 0x6008CBB7 | 8 |  | - |
| `DS_Z_MEM[43]` | 0x6008CA2B | 8 |  | - |
| `DS_Z_MEM[440]` | 0x6008CBB8 | 8 |  | - |
| `DS_Z_MEM[441]` | 0x6008CBB9 | 8 |  | - |
| `DS_Z_MEM[442]` | 0x6008CBBA | 8 |  | - |
| `DS_Z_MEM[443]` | 0x6008CBBB | 8 |  | - |
| `DS_Z_MEM[444]` | 0x6008CBBC | 8 |  | - |
| `DS_Z_MEM[445]` | 0x6008CBBD | 8 |  | - |
| `DS_Z_MEM[446]` | 0x6008CBBE | 8 |  | - |
| `DS_Z_MEM[447]` | 0x6008CBBF | 8 |  | - |
| `DS_Z_MEM[448]` | 0x6008CBC0 | 8 |  | - |
| `DS_Z_MEM[449]` | 0x6008CBC1 | 8 |  | - |
| `DS_Z_MEM[44]` | 0x6008CA2C | 8 |  | - |
| `DS_Z_MEM[450]` | 0x6008CBC2 | 8 |  | - |
| `DS_Z_MEM[451]` | 0x6008CBC3 | 8 |  | - |
| `DS_Z_MEM[452]` | 0x6008CBC4 | 8 |  | - |
| `DS_Z_MEM[453]` | 0x6008CBC5 | 8 |  | - |
| `DS_Z_MEM[454]` | 0x6008CBC6 | 8 |  | - |
| `DS_Z_MEM[455]` | 0x6008CBC7 | 8 |  | - |
| `DS_Z_MEM[456]` | 0x6008CBC8 | 8 |  | - |
| `DS_Z_MEM[457]` | 0x6008CBC9 | 8 |  | - |
| `DS_Z_MEM[458]` | 0x6008CBCA | 8 |  | - |
| `DS_Z_MEM[459]` | 0x6008CBCB | 8 |  | - |
| `DS_Z_MEM[45]` | 0x6008CA2D | 8 |  | - |
| `DS_Z_MEM[460]` | 0x6008CBCC | 8 |  | - |
| `DS_Z_MEM[461]` | 0x6008CBCD | 8 |  | - |
| `DS_Z_MEM[462]` | 0x6008CBCE | 8 |  | - |
| `DS_Z_MEM[463]` | 0x6008CBCF | 8 |  | - |
| `DS_Z_MEM[464]` | 0x6008CBD0 | 8 |  | - |
| `DS_Z_MEM[465]` | 0x6008CBD1 | 8 |  | - |
| `DS_Z_MEM[466]` | 0x6008CBD2 | 8 |  | - |
| `DS_Z_MEM[467]` | 0x6008CBD3 | 8 |  | - |
| `DS_Z_MEM[468]` | 0x6008CBD4 | 8 |  | - |
| `DS_Z_MEM[469]` | 0x6008CBD5 | 8 |  | - |
| `DS_Z_MEM[46]` | 0x6008CA2E | 8 |  | - |
| `DS_Z_MEM[470]` | 0x6008CBD6 | 8 |  | - |
| `DS_Z_MEM[471]` | 0x6008CBD7 | 8 |  | - |
| `DS_Z_MEM[472]` | 0x6008CBD8 | 8 |  | - |
| `DS_Z_MEM[473]` | 0x6008CBD9 | 8 |  | - |
| `DS_Z_MEM[474]` | 0x6008CBDA | 8 |  | - |
| `DS_Z_MEM[475]` | 0x6008CBDB | 8 |  | - |
| `DS_Z_MEM[476]` | 0x6008CBDC | 8 |  | - |
| `DS_Z_MEM[477]` | 0x6008CBDD | 8 |  | - |
| `DS_Z_MEM[478]` | 0x6008CBDE | 8 |  | - |
| `DS_Z_MEM[479]` | 0x6008CBDF | 8 |  | - |
| `DS_Z_MEM[47]` | 0x6008CA2F | 8 |  | - |
| `DS_Z_MEM[480]` | 0x6008CBE0 | 8 |  | - |
| `DS_Z_MEM[481]` | 0x6008CBE1 | 8 |  | - |
| `DS_Z_MEM[482]` | 0x6008CBE2 | 8 |  | - |
| `DS_Z_MEM[483]` | 0x6008CBE3 | 8 |  | - |
| `DS_Z_MEM[484]` | 0x6008CBE4 | 8 |  | - |
| `DS_Z_MEM[485]` | 0x6008CBE5 | 8 |  | - |
| `DS_Z_MEM[486]` | 0x6008CBE6 | 8 |  | - |
| `DS_Z_MEM[487]` | 0x6008CBE7 | 8 |  | - |
| `DS_Z_MEM[488]` | 0x6008CBE8 | 8 |  | - |
| `DS_Z_MEM[489]` | 0x6008CBE9 | 8 |  | - |
| `DS_Z_MEM[48]` | 0x6008CA30 | 8 |  | - |
| `DS_Z_MEM[490]` | 0x6008CBEA | 8 |  | - |
| `DS_Z_MEM[491]` | 0x6008CBEB | 8 |  | - |
| `DS_Z_MEM[492]` | 0x6008CBEC | 8 |  | - |
| `DS_Z_MEM[493]` | 0x6008CBED | 8 |  | - |
| `DS_Z_MEM[494]` | 0x6008CBEE | 8 |  | - |
| `DS_Z_MEM[495]` | 0x6008CBEF | 8 |  | - |
| `DS_Z_MEM[496]` | 0x6008CBF0 | 8 |  | - |
| `DS_Z_MEM[497]` | 0x6008CBF1 | 8 |  | - |
| `DS_Z_MEM[498]` | 0x6008CBF2 | 8 |  | - |
| `DS_Z_MEM[499]` | 0x6008CBF3 | 8 |  | - |
| `DS_Z_MEM[49]` | 0x6008CA31 | 8 |  | - |
| `DS_Z_MEM[4]` | 0x6008CA04 | 8 |  | - |
| `DS_Z_MEM[500]` | 0x6008CBF4 | 8 |  | - |
| `DS_Z_MEM[501]` | 0x6008CBF5 | 8 |  | - |
| `DS_Z_MEM[502]` | 0x6008CBF6 | 8 |  | - |
| `DS_Z_MEM[503]` | 0x6008CBF7 | 8 |  | - |
| `DS_Z_MEM[504]` | 0x6008CBF8 | 8 |  | - |
| `DS_Z_MEM[505]` | 0x6008CBF9 | 8 |  | - |
| `DS_Z_MEM[506]` | 0x6008CBFA | 8 |  | - |
| `DS_Z_MEM[507]` | 0x6008CBFB | 8 |  | - |
| `DS_Z_MEM[508]` | 0x6008CBFC | 8 |  | - |
| `DS_Z_MEM[509]` | 0x6008CBFD | 8 |  | - |
| `DS_Z_MEM[50]` | 0x6008CA32 | 8 |  | - |
| `DS_Z_MEM[510]` | 0x6008CBFE | 8 |  | - |
| `DS_Z_MEM[511]` | 0x6008CBFF | 8 |  | - |
| `DS_Z_MEM[51]` | 0x6008CA33 | 8 |  | - |
| `DS_Z_MEM[52]` | 0x6008CA34 | 8 |  | - |
| `DS_Z_MEM[53]` | 0x6008CA35 | 8 |  | - |
| `DS_Z_MEM[54]` | 0x6008CA36 | 8 |  | - |
| `DS_Z_MEM[55]` | 0x6008CA37 | 8 |  | - |
| `DS_Z_MEM[56]` | 0x6008CA38 | 8 |  | - |
| `DS_Z_MEM[57]` | 0x6008CA39 | 8 |  | - |
| `DS_Z_MEM[58]` | 0x6008CA3A | 8 |  | - |
| `DS_Z_MEM[59]` | 0x6008CA3B | 8 |  | - |
| `DS_Z_MEM[5]` | 0x6008CA05 | 8 |  | - |
| `DS_Z_MEM[60]` | 0x6008CA3C | 8 |  | - |
| `DS_Z_MEM[61]` | 0x6008CA3D | 8 |  | - |
| `DS_Z_MEM[62]` | 0x6008CA3E | 8 |  | - |
| `DS_Z_MEM[63]` | 0x6008CA3F | 8 |  | - |
| `DS_Z_MEM[64]` | 0x6008CA40 | 8 |  | - |
| `DS_Z_MEM[65]` | 0x6008CA41 | 8 |  | - |
| `DS_Z_MEM[66]` | 0x6008CA42 | 8 |  | - |
| `DS_Z_MEM[67]` | 0x6008CA43 | 8 |  | - |
| `DS_Z_MEM[68]` | 0x6008CA44 | 8 |  | - |
| `DS_Z_MEM[69]` | 0x6008CA45 | 8 |  | - |
| `DS_Z_MEM[6]` | 0x6008CA06 | 8 |  | - |
| `DS_Z_MEM[70]` | 0x6008CA46 | 8 |  | - |
| `DS_Z_MEM[71]` | 0x6008CA47 | 8 |  | - |
| `DS_Z_MEM[72]` | 0x6008CA48 | 8 |  | - |
| `DS_Z_MEM[73]` | 0x6008CA49 | 8 |  | - |
| `DS_Z_MEM[74]` | 0x6008CA4A | 8 |  | - |
| `DS_Z_MEM[75]` | 0x6008CA4B | 8 |  | - |
| `DS_Z_MEM[76]` | 0x6008CA4C | 8 |  | - |
| `DS_Z_MEM[77]` | 0x6008CA4D | 8 |  | - |
| `DS_Z_MEM[78]` | 0x6008CA4E | 8 |  | - |
| `DS_Z_MEM[79]` | 0x6008CA4F | 8 |  | - |
| `DS_Z_MEM[7]` | 0x6008CA07 | 8 |  | - |
| `DS_Z_MEM[80]` | 0x6008CA50 | 8 |  | - |
| `DS_Z_MEM[81]` | 0x6008CA51 | 8 |  | - |
| `DS_Z_MEM[82]` | 0x6008CA52 | 8 |  | - |
| `DS_Z_MEM[83]` | 0x6008CA53 | 8 |  | - |
| `DS_Z_MEM[84]` | 0x6008CA54 | 8 |  | - |
| `DS_Z_MEM[85]` | 0x6008CA55 | 8 |  | - |
| `DS_Z_MEM[86]` | 0x6008CA56 | 8 |  | - |
| `DS_Z_MEM[87]` | 0x6008CA57 | 8 |  | - |
| `DS_Z_MEM[88]` | 0x6008CA58 | 8 |  | - |
| `DS_Z_MEM[89]` | 0x6008CA59 | 8 |  | - |
| `DS_Z_MEM[8]` | 0x6008CA08 | 8 |  | - |
| `DS_Z_MEM[90]` | 0x6008CA5A | 8 |  | - |
| `DS_Z_MEM[91]` | 0x6008CA5B | 8 |  | - |
| `DS_Z_MEM[92]` | 0x6008CA5C | 8 |  | - |
| `DS_Z_MEM[93]` | 0x6008CA5D | 8 |  | - |
| `DS_Z_MEM[94]` | 0x6008CA5E | 8 |  | - |
| `DS_Z_MEM[95]` | 0x6008CA5F | 8 |  | - |
| `DS_Z_MEM[96]` | 0x6008CA60 | 8 |  | - |
| `DS_Z_MEM[97]` | 0x6008CA61 | 8 |  | - |
| `DS_Z_MEM[98]` | 0x6008CA62 | 8 |  | - |
| `DS_Z_MEM[99]` | 0x6008CA63 | 8 |  | - |
| `DS_Z_MEM[9]` | 0x6008CA09 | 8 |  | - |
| `ECC_K_MEM[0]` | 0x60039100 | 8 |  | - |
| `ECC_K_MEM[10]` | 0x6003910A | 8 |  | - |
| `ECC_K_MEM[11]` | 0x6003910B | 8 |  | - |
| `ECC_K_MEM[12]` | 0x6003910C | 8 |  | - |
| `ECC_K_MEM[13]` | 0x6003910D | 8 |  | - |
| `ECC_K_MEM[14]` | 0x6003910E | 8 |  | - |
| `ECC_K_MEM[15]` | 0x6003910F | 8 |  | - |
| `ECC_K_MEM[16]` | 0x60039110 | 8 |  | - |
| `ECC_K_MEM[17]` | 0x60039111 | 8 |  | - |
| `ECC_K_MEM[18]` | 0x60039112 | 8 |  | - |
| `ECC_K_MEM[19]` | 0x60039113 | 8 |  | - |
| `ECC_K_MEM[1]` | 0x60039101 | 8 |  | - |
| `ECC_K_MEM[20]` | 0x60039114 | 8 |  | - |
| `ECC_K_MEM[21]` | 0x60039115 | 8 |  | - |
| `ECC_K_MEM[22]` | 0x60039116 | 8 |  | - |
| `ECC_K_MEM[23]` | 0x60039117 | 8 |  | - |
| `ECC_K_MEM[24]` | 0x60039118 | 8 |  | - |
| `ECC_K_MEM[25]` | 0x60039119 | 8 |  | - |
| `ECC_K_MEM[26]` | 0x6003911A | 8 |  | - |
| `ECC_K_MEM[27]` | 0x6003911B | 8 |  | - |
| `ECC_K_MEM[28]` | 0x6003911C | 8 |  | - |
| `ECC_K_MEM[29]` | 0x6003911D | 8 |  | - |
| `ECC_K_MEM[2]` | 0x60039102 | 8 |  | - |
| `ECC_K_MEM[30]` | 0x6003911E | 8 |  | - |
| `ECC_K_MEM[31]` | 0x6003911F | 8 |  | - |
| `ECC_K_MEM[3]` | 0x60039103 | 8 |  | - |
| `ECC_K_MEM[4]` | 0x60039104 | 8 |  | - |
| `ECC_K_MEM[5]` | 0x60039105 | 8 |  | - |
| `ECC_K_MEM[6]` | 0x60039106 | 8 |  | - |
| `ECC_K_MEM[7]` | 0x60039107 | 8 |  | - |
| `ECC_K_MEM[8]` | 0x60039108 | 8 |  | - |
| `ECC_K_MEM[9]` | 0x60039109 | 8 |  | - |
| `ECC_MULT_CONF` | 0x6003901C | 32 | CLK_EN=0x40000000 KEY_LENGTH=0x4 MEM_CLOCK_GATE_FORCE_ON=0x80000000 MOD_BASE=0x8 RESET=0x2 SECURITY_MODE=0x100 START=0x1 VERIFICATION_RESULT=0x20000000 WORK_MODE=0xF0 | - |
| `ECC_MULT_DATE` | 0x600390FC | 32 | DATE=0xFFFFFFF | - |
| `ECC_MULT_INT_CLR` | 0x60039018 | 32 | CALC_DONE_INT_CLR=0x1 | - |
| `ECC_MULT_INT_ENA` | 0x60039014 | 32 | CALC_DONE_INT_ENA=0x1 | - |
| `ECC_MULT_INT_RAW` | 0x6003900C | 32 | CALC_DONE_INT_RAW=0x1 | - |
| `ECC_MULT_INT_ST` | 0x60039010 | 32 | CALC_DONE_INT_ST=0x1 | - |
| `ECC_PX_MEM[0]` | 0x60039120 | 8 |  | - |
| `ECC_PX_MEM[10]` | 0x6003912A | 8 |  | - |
| `ECC_PX_MEM[11]` | 0x6003912B | 8 |  | - |
| `ECC_PX_MEM[12]` | 0x6003912C | 8 |  | - |
| `ECC_PX_MEM[13]` | 0x6003912D | 8 |  | - |
| `ECC_PX_MEM[14]` | 0x6003912E | 8 |  | - |
| `ECC_PX_MEM[15]` | 0x6003912F | 8 |  | - |
| `ECC_PX_MEM[16]` | 0x60039130 | 8 |  | - |
| `ECC_PX_MEM[17]` | 0x60039131 | 8 |  | - |
| `ECC_PX_MEM[18]` | 0x60039132 | 8 |  | - |
| `ECC_PX_MEM[19]` | 0x60039133 | 8 |  | - |
| `ECC_PX_MEM[1]` | 0x60039121 | 8 |  | - |
| `ECC_PX_MEM[20]` | 0x60039134 | 8 |  | - |
| `ECC_PX_MEM[21]` | 0x60039135 | 8 |  | - |
| `ECC_PX_MEM[22]` | 0x60039136 | 8 |  | - |
| `ECC_PX_MEM[23]` | 0x60039137 | 8 |  | - |
| `ECC_PX_MEM[24]` | 0x60039138 | 8 |  | - |
| `ECC_PX_MEM[25]` | 0x60039139 | 8 |  | - |
| `ECC_PX_MEM[26]` | 0x6003913A | 8 |  | - |
| `ECC_PX_MEM[27]` | 0x6003913B | 8 |  | - |
| `ECC_PX_MEM[28]` | 0x6003913C | 8 |  | - |
| `ECC_PX_MEM[29]` | 0x6003913D | 8 |  | - |
| `ECC_PX_MEM[2]` | 0x60039122 | 8 |  | - |
| `ECC_PX_MEM[30]` | 0x6003913E | 8 |  | - |
| `ECC_PX_MEM[31]` | 0x6003913F | 8 |  | - |
| `ECC_PX_MEM[3]` | 0x60039123 | 8 |  | - |
| `ECC_PX_MEM[4]` | 0x60039124 | 8 |  | - |
| `ECC_PX_MEM[5]` | 0x60039125 | 8 |  | - |
| `ECC_PX_MEM[6]` | 0x60039126 | 8 |  | - |
| `ECC_PX_MEM[7]` | 0x60039127 | 8 |  | - |
| `ECC_PX_MEM[8]` | 0x60039128 | 8 |  | - |
| `ECC_PX_MEM[9]` | 0x60039129 | 8 |  | - |
| `ECC_PY_MEM[0]` | 0x60039140 | 8 |  | - |
| `ECC_PY_MEM[10]` | 0x6003914A | 8 |  | - |
| `ECC_PY_MEM[11]` | 0x6003914B | 8 |  | - |
| `ECC_PY_MEM[12]` | 0x6003914C | 8 |  | - |
| `ECC_PY_MEM[13]` | 0x6003914D | 8 |  | - |
| `ECC_PY_MEM[14]` | 0x6003914E | 8 |  | - |
| `ECC_PY_MEM[15]` | 0x6003914F | 8 |  | - |
| `ECC_PY_MEM[16]` | 0x60039150 | 8 |  | - |
| `ECC_PY_MEM[17]` | 0x60039151 | 8 |  | - |
| `ECC_PY_MEM[18]` | 0x60039152 | 8 |  | - |
| `ECC_PY_MEM[19]` | 0x60039153 | 8 |  | - |
| `ECC_PY_MEM[1]` | 0x60039141 | 8 |  | - |
| `ECC_PY_MEM[20]` | 0x60039154 | 8 |  | - |
| `ECC_PY_MEM[21]` | 0x60039155 | 8 |  | - |
| `ECC_PY_MEM[22]` | 0x60039156 | 8 |  | - |
| `ECC_PY_MEM[23]` | 0x60039157 | 8 |  | - |
| `ECC_PY_MEM[24]` | 0x60039158 | 8 |  | - |
| `ECC_PY_MEM[25]` | 0x60039159 | 8 |  | - |
| `ECC_PY_MEM[26]` | 0x6003915A | 8 |  | - |
| `ECC_PY_MEM[27]` | 0x6003915B | 8 |  | - |
| `ECC_PY_MEM[28]` | 0x6003915C | 8 |  | - |
| `ECC_PY_MEM[29]` | 0x6003915D | 8 |  | - |
| `ECC_PY_MEM[2]` | 0x60039142 | 8 |  | - |
| `ECC_PY_MEM[30]` | 0x6003915E | 8 |  | - |
| `ECC_PY_MEM[31]` | 0x6003915F | 8 |  | - |
| `ECC_PY_MEM[3]` | 0x60039143 | 8 |  | - |
| `ECC_PY_MEM[4]` | 0x60039144 | 8 |  | - |
| `ECC_PY_MEM[5]` | 0x60039145 | 8 |  | - |
| `ECC_PY_MEM[6]` | 0x60039146 | 8 |  | - |
| `ECC_PY_MEM[7]` | 0x60039147 | 8 |  | - |
| `ECC_PY_MEM[8]` | 0x60039148 | 8 |  | - |
| `ECC_PY_MEM[9]` | 0x60039149 | 8 |  | - |
| `EFUSE_CLK` | 0x6001A1C8 | 32 | EN=0x10000 MEM_CLK_FORCE_ON=0x2 MEM_FORCE_PD=0x1 MEM_FORCE_PU=0x4 | - |
| `EFUSE_CMD` | 0x6001A1D4 | 32 | BLK_NUM=0x3C PGM_CMD=0x2 READ_CMD=0x1 | - |
| `EFUSE_CONF` | 0x6001A1CC | 32 | CFG_ECDSA_BLK=0xF0000 OP_CODE=0xFFFF | - |
| `EFUSE_DAC_CONF` | 0x6001A1E8 | 32 | DAC_CLK_DIV=0xFF DAC_CLK_PAD_SEL=0x100 DAC_NUM=0x1FE00 OE_CLR=0x20000 | - |
| `EFUSE_DATE` | 0x6001A1FC | 32 | DATE=0xFFFFFFF | - |
| `EFUSE_INT_CLR` | 0x6001A1E4 | 32 | PGM_DONE_INT_CLR=0x2 READ_DONE_INT_CLR=0x1 | - |
| `EFUSE_INT_ENA` | 0x6001A1E0 | 32 | PGM_DONE_INT_ENA=0x2 READ_DONE_INT_ENA=0x1 | - |
| `EFUSE_INT_RAW` | 0x6001A1D8 | 32 | PGM_DONE_INT_RAW=0x2 READ_DONE_INT_RAW=0x1 | - |
| `EFUSE_INT_ST` | 0x6001A1DC | 32 | PGM_DONE_INT_ST=0x2 READ_DONE_INT_ST=0x1 | - |
| `EFUSE_PGM_CHECK_VALUE0` | 0x6001A020 | 32 | PGM_RS_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_CHECK_VALUE1` | 0x6001A024 | 32 | PGM_RS_DATA_1=0xFFFFFFFF | - |
| `EFUSE_PGM_CHECK_VALUE2` | 0x6001A028 | 32 | PGM_RS_DATA_2=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA0` | 0x6001A000 | 32 | PGM_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA1` | 0x6001A004 | 32 | PGM_DATA_1=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA2` | 0x6001A008 | 32 | PGM_DATA_2=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA3` | 0x6001A00C | 32 | PGM_DATA_3=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA4` | 0x6001A010 | 32 | PGM_DATA_4=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA5` | 0x6001A014 | 32 | PGM_DATA_5=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA6` | 0x6001A018 | 32 | PGM_DATA_6=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA7` | 0x6001A01C | 32 | PGM_DATA_7=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA0` | 0x6001A09C | 32 | KEY0_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA1` | 0x6001A0A0 | 32 | KEY0_DATA1=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA2` | 0x6001A0A4 | 32 | KEY0_DATA2=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA3` | 0x6001A0A8 | 32 | KEY0_DATA3=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA4` | 0x6001A0AC | 32 | KEY0_DATA4=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA5` | 0x6001A0B0 | 32 | KEY0_DATA5=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA6` | 0x6001A0B4 | 32 | KEY0_DATA6=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA7` | 0x6001A0B8 | 32 | KEY0_DATA7=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA0` | 0x6001A0BC | 32 | KEY1_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA1` | 0x6001A0C0 | 32 | KEY1_DATA1=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA2` | 0x6001A0C4 | 32 | KEY1_DATA2=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA3` | 0x6001A0C8 | 32 | KEY1_DATA3=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA4` | 0x6001A0CC | 32 | KEY1_DATA4=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA5` | 0x6001A0D0 | 32 | KEY1_DATA5=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA6` | 0x6001A0D4 | 32 | KEY1_DATA6=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA7` | 0x6001A0D8 | 32 | KEY1_DATA7=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA0` | 0x6001A0DC | 32 | KEY2_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA1` | 0x6001A0E0 | 32 | KEY2_DATA1=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA2` | 0x6001A0E4 | 32 | KEY2_DATA2=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA3` | 0x6001A0E8 | 32 | KEY2_DATA3=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA4` | 0x6001A0EC | 32 | KEY2_DATA4=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA5` | 0x6001A0F0 | 32 | KEY2_DATA5=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA6` | 0x6001A0F4 | 32 | KEY2_DATA6=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA7` | 0x6001A0F8 | 32 | KEY2_DATA7=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA0` | 0x6001A0FC | 32 | KEY3_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA1` | 0x6001A100 | 32 | KEY3_DATA1=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA2` | 0x6001A104 | 32 | KEY3_DATA2=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA3` | 0x6001A108 | 32 | KEY3_DATA3=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA4` | 0x6001A10C | 32 | KEY3_DATA4=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA5` | 0x6001A110 | 32 | KEY3_DATA5=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA6` | 0x6001A114 | 32 | KEY3_DATA6=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA7` | 0x6001A118 | 32 | KEY3_DATA7=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA0` | 0x6001A11C | 32 | KEY4_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA1` | 0x6001A120 | 32 | KEY4_DATA1=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA2` | 0x6001A124 | 32 | KEY4_DATA2=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA3` | 0x6001A128 | 32 | KEY4_DATA3=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA4` | 0x6001A12C | 32 | KEY4_DATA4=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA5` | 0x6001A130 | 32 | KEY4_DATA5=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA6` | 0x6001A134 | 32 | KEY4_DATA6=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA7` | 0x6001A138 | 32 | KEY4_DATA7=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA0` | 0x6001A13C | 32 | KEY5_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA1` | 0x6001A140 | 32 | KEY5_DATA1=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA2` | 0x6001A144 | 32 | KEY5_DATA2=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA3` | 0x6001A148 | 32 | KEY5_DATA3=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA4` | 0x6001A14C | 32 | KEY5_DATA4=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA5` | 0x6001A150 | 32 | KEY5_DATA5=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA6` | 0x6001A154 | 32 | KEY5_DATA6=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA7` | 0x6001A158 | 32 | KEY5_DATA7=0xFFFFFFFF | - |
| `EFUSE_RD_MAC_SYS_0` | 0x6001A044 | 32 | MAC_0=0xFFFFFFFF | - |
| `EFUSE_RD_MAC_SYS_1` | 0x6001A048 | 32 | MAC_1=0xFFFF MAC_EXT=0xFFFF0000 | - |
| `EFUSE_RD_MAC_SYS_2` | 0x6001A04C | 32 | MAC_RESERVED_0=0xFFFFC000 MAC_RESERVED_1=0x3FFF | - |
| `EFUSE_RD_MAC_SYS_3` | 0x6001A050 | 32 | MAC_RESERVED_2=0x3FFFF SYS_DATA_PART0_0=0xFFFC0000 | - |
| `EFUSE_RD_MAC_SYS_4` | 0x6001A054 | 32 | SYS_DATA_PART0_1=0xFFFFFFFF | - |
| `EFUSE_RD_MAC_SYS_5` | 0x6001A058 | 32 | SYS_DATA_PART0_2=0xFFFFFFFF | - |
| `EFUSE_RD_REPEAT_DATA0` | 0x6001A030 | 32 | DIS_CAN=0x4000 DIS_DOWNLOAD_MANUAL_ENCRYPT=0x100000 DIS_FORCE_DOWNLOAD=0x1000 DIS_ICACHE=0x100 DIS_PAD_JTAG=0x80000 DIS_USB_JTAG=0x200 DIS_USB_SERIAL_JTAG=0x800 JTAG_SEL_ENABLE=0x8000 POWERGLITCH_EN=0x400 RD_DIS=0x7F RPT4_RESERVED0_0=0xC000... | - |
| `EFUSE_RD_REPEAT_DATA1` | 0x6001A034 | 32 | KEY_PURPOSE_0=0xF000000 KEY_PURPOSE_1=0xF0000000 RPT4_RESERVED1_1=0xFFFF SECURE_BOOT_KEY_REVOKE0=0x200000 SECURE_BOOT_KEY_REVOKE1=0x400000 SECURE_BOOT_KEY_REVOKE2=0x800000 SPI_BOOT_CRYPT_CNT=0x1C0000 WDT_DELAY_SEL=0x30000 | - |
| `EFUSE_RD_REPEAT_DATA2` | 0x6001A038 | 32 | CRYPT_DPA_ENABLE=0x80000 ECDSA_FORCE_USE_HARDWARE_K=0x40000 FLASH_TPUW=0xF0000000 KEY_PURPOSE_2=0xF KEY_PURPOSE_3=0xF0 KEY_PURPOSE_4=0xF00 KEY_PURPOSE_5=0xF000 RPT4_RESERVED2_0=0xFC00000 SECURE_BOOT_AGGRESSIVE_REVOKE=0x200000 SECURE_BOOT_EN... | - |
| `EFUSE_RD_REPEAT_DATA3` | 0x6001A03C | 32 | DIS_DIRECT_BOOT=0x2 DIS_DOWNLOAD_MODE=0x1 DIS_USB_PRINT=0x4 DIS_USB_SERIAL_JTAG_DOWNLOAD_MODE=0x10 ENABLE_SECURITY_DOWNLOAD=0x20 FORCE_SEND_RESUME=0x100 HYS_EN_PAD0=0xFC000000 RPT4_RESERVED3_5=0x8 SECURE_BOOT_DISABLE_FAST_WAKE=0x2000000 SEC... | - |
| `EFUSE_RD_REPEAT_DATA4` | 0x6001A040 | 32 | HYS_EN_PAD1=0x3FFFFF RPT4_RESERVED4_0=0xFF000000 RPT4_RESERVED4_1=0xC00000 | - |
| `EFUSE_RD_REPEAT_ERR0` | 0x6001A17C | 32 | DIS_DOWNLOAD_MANUAL_ENCRYPT_ERR=0x100000 DIS_FORCE_DOWNLOAD_ERR=0x1000 DIS_ICACHE_ERR=0x100 DIS_PAD_JTAG_ERR=0x80000 DIS_TWAI_ERR=0x4000 DIS_USB_JTAG_ERR=0x200 DIS_USB_SERIAL_JTAG_ERR=0x800 JTAG_SEL_ENABLE_ERR=0x8000 POWERGLITCH_EN_ERR=0x40... | - |
| `EFUSE_RD_REPEAT_ERR1` | 0x6001A180 | 32 | KEY_PURPOSE_0_ERR=0xF000000 KEY_PURPOSE_1_ERR=0xF0000000 RPT4_RESERVED1_ERR_0=0xFFFF SECURE_BOOT_KEY_REVOKE0_ERR=0x200000 SECURE_BOOT_KEY_REVOKE1_ERR=0x400000 SECURE_BOOT_KEY_REVOKE2_ERR=0x800000 SPI_BOOT_CRYPT_CNT_ERR=0x1C0000 WDT_DELAY_SE... | - |
| `EFUSE_RD_REPEAT_ERR2` | 0x6001A184 | 32 | CRYPT_DPA_ENABLE_ERR=0x80000 FLASH_TPUW_ERR=0xF0000000 KEY_PURPOSE_2_ERR=0xF KEY_PURPOSE_3_ERR=0xF0 KEY_PURPOSE_4_ERR=0xF00 KEY_PURPOSE_5_ERR=0xF000 RPT4_RESERVED2_ERR_0=0xFC00000 RPT4_RESERVED2_ERR_1=0x40000 SECURE_BOOT_AGGRESSIVE_REVOKE_E... | - |
| `EFUSE_RD_REPEAT_ERR3` | 0x6001A188 | 32 | DIS_DIRECT_BOOT_ERR=0x2 DIS_DOWNLOAD_MODE_ERR=0x1 DIS_USB_SERIAL_JTAG_DOWNLOAD_MODE_ERR=0x10 ENABLE_SECURITY_DOWNLOAD_ERR=0x20 FORCE_SEND_RESUME_ERR=0x100 HYS_EN_PAD0_ERR=0xFC000000 RPT4_RESERVED3_ERR_5=0x8 SECURE_BOOT_DISABLE_FAST_WAKE_ERR... | - |
| `EFUSE_RD_REPEAT_ERR4` | 0x6001A18C | 32 | HYS_EN_PAD1_ERR=0x3FFFFF RPT4_RESERVED4_ERR_0=0xFF000000 RPT4_RESERVED4_ERR_1=0xC00000 | - |
| `EFUSE_RD_RS_ERR0` | 0x6001A1C0 | 32 | KEY0_ERR_NUM=0x7000 KEY0_FAIL=0x8000 KEY1_ERR_NUM=0x70000 KEY1_FAIL=0x80000 KEY2_ERR_NUM=0x700000 KEY2_FAIL=0x800000 KEY3_ERR_NUM=0x7000000 KEY3_FAIL=0x8000000 KEY4_ERR_NUM=0x70000000 KEY4_FAIL=0x80000000 MAC_SPI_8M_ERR_NUM=0x7 MAC_SPI_8M_F... | - |
| `EFUSE_RD_RS_ERR1` | 0x6001A1C4 | 32 | KEY5_ERR_NUM=0x7 KEY5_FAIL=0x8 SYS_PART2_ERR_NUM=0x70 SYS_PART2_FAIL=0x80 | - |
| `EFUSE_RD_SYS_PART1_DATA0` | 0x6001A05C | 32 | SYS_DATA_PART1_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART1_DATA1` | 0x6001A060 | 32 | SYS_DATA_PART1_1=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART1_DATA2` | 0x6001A064 | 32 | SYS_DATA_PART1_2=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART1_DATA3` | 0x6001A068 | 32 | SYS_DATA_PART1_3=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART1_DATA4` | 0x6001A06C | 32 | SYS_DATA_PART1_4=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART1_DATA5` | 0x6001A070 | 32 | SYS_DATA_PART1_5=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART1_DATA6` | 0x6001A074 | 32 | SYS_DATA_PART1_6=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART1_DATA7` | 0x6001A078 | 32 | SYS_DATA_PART1_7=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART2_DATA0` | 0x6001A15C | 32 | SYS_DATA_PART2_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART2_DATA1` | 0x6001A160 | 32 | SYS_DATA_PART2_1=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART2_DATA2` | 0x6001A164 | 32 | SYS_DATA_PART2_2=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART2_DATA3` | 0x6001A168 | 32 | SYS_DATA_PART2_3=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART2_DATA4` | 0x6001A16C | 32 | SYS_DATA_PART2_4=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART2_DATA5` | 0x6001A170 | 32 | SYS_DATA_PART2_5=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART2_DATA6` | 0x6001A174 | 32 | SYS_DATA_PART2_6=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_PART2_DATA7` | 0x6001A178 | 32 | SYS_DATA_PART2_7=0xFFFFFFFF | - |
| `EFUSE_RD_TIM_CONF` | 0x6001A1EC | 32 | READ_INIT_NUM=0xFF000000 THR_A=0xFF TRD=0xFF00 TSUR_A=0xFF0000 | - |
| `EFUSE_RD_USR_DATA0` | 0x6001A07C | 32 | USR_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA1` | 0x6001A080 | 32 | USR_DATA1=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA2` | 0x6001A084 | 32 | USR_DATA2=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA3` | 0x6001A088 | 32 | USR_DATA3=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA4` | 0x6001A08C | 32 | USR_DATA4=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA5` | 0x6001A090 | 32 | USR_DATA5=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA6` | 0x6001A094 | 32 | USR_DATA6=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA7` | 0x6001A098 | 32 | USR_DATA7=0xFFFFFFFF | - |
| `EFUSE_RD_WR_DIS` | 0x6001A02C | 32 | WR_DIS=0xFFFFFFFF | - |
| `EFUSE_STATUS` | 0x6001A1D0 | 32 | BLK0_VALID_BIT_CNT=0xFFC00 CUR_ECDSA_BLK=0xF00000 OTP_CSB_SW=0x80 OTP_LOAD_SW=0x10 OTP_PGENB_SW=0x100 OTP_STROBE_SW=0x40 OTP_VDDQ_C_SYNC2=0x20 OTP_VDDQ_IS_SW=0x200 STATE=0xF | - |
| `EFUSE_WR_TIM_CONF0_RS_BYPASS` | 0x6001A1F8 | 32 | BYPASS_RS_BLK_NUM=0xFFE BYPASS_RS_CORRECTION=0x1 TPGM_INACTIVE=0x1FE000 UPDATE=0x1000 | - |
| `EFUSE_WR_TIM_CONF1` | 0x6001A1F0 | 32 | PWR_ON_NUM=0xFFFF00 THP_A=0xFF000000 TSUP_A=0xFF | - |
| `EFUSE_WR_TIM_CONF2` | 0x6001A1F4 | 32 | PWR_OFF_NUM=0xFFFF TPGM=0xFFFF0000 | - |
| `GDMA_AHB_TEST` | 0x6003F060 | 32 | AHB_TESTADDR=0x30 AHB_TESTMODE=0x7 | - |
| `GDMA_DATE` | 0x6003F068 | 32 | DATE=0xFFFFFFFF | - |
| `GDMA_INFIFO_STATUS_CH0` | 0x6003F078 | 32 | INFIFO_CNT_CH=0xFC INFIFO_EMPTY_CH=0x2 INFIFO_FULL_CH=0x1 IN_BUF_HUNGRY_CH=0x8000000 IN_REMAIN_UNDER_1B_CH=0x800000 IN_REMAIN_UNDER_2B_CH=0x1000000 IN_REMAIN_UNDER_3B_CH=0x2000000 IN_REMAIN_UNDER_4B_CH=0x4000000 | - |
| `GDMA_INFIFO_STATUS_CH1` | 0x6003F138 | 32 | INFIFO_CNT_CH=0xFC INFIFO_EMPTY_CH=0x2 INFIFO_FULL_CH=0x1 IN_BUF_HUNGRY_CH=0x8000000 IN_REMAIN_UNDER_1B_CH=0x800000 IN_REMAIN_UNDER_2B_CH=0x1000000 IN_REMAIN_UNDER_3B_CH=0x2000000 IN_REMAIN_UNDER_4B_CH=0x4000000 | - |
| `GDMA_INFIFO_STATUS_CH2` | 0x6003F1F8 | 32 | INFIFO_CNT_CH=0xFC INFIFO_EMPTY_CH=0x2 INFIFO_FULL_CH=0x1 IN_BUF_HUNGRY_CH=0x8000000 IN_REMAIN_UNDER_1B_CH=0x800000 IN_REMAIN_UNDER_2B_CH=0x1000000 IN_REMAIN_UNDER_3B_CH=0x2000000 IN_REMAIN_UNDER_4B_CH=0x4000000 | - |
| `GDMA_IN_CONF0_CH0` | 0x6003F070 | 32 | INDSCR_BURST_EN_CH=0x4 IN_DATA_BURST_EN_CH=0x8 IN_ETM_EN_CH=0x20 IN_LOOP_TEST_CH=0x2 IN_RST_CH=0x1 MEM_TRANS_EN_CH=0x10 | - |
| `GDMA_IN_CONF0_CH1` | 0x6003F130 | 32 | INDSCR_BURST_EN_CH=0x4 IN_DATA_BURST_EN_CH=0x8 IN_ETM_EN_CH=0x20 IN_LOOP_TEST_CH=0x2 IN_RST_CH=0x1 MEM_TRANS_EN_CH=0x10 | - |
| `GDMA_IN_CONF0_CH2` | 0x6003F1F0 | 32 | INDSCR_BURST_EN_CH=0x4 IN_DATA_BURST_EN_CH=0x8 IN_ETM_EN_CH=0x20 IN_LOOP_TEST_CH=0x2 IN_RST_CH=0x1 MEM_TRANS_EN_CH=0x10 | - |
| `GDMA_IN_CONF1_CH0` | 0x6003F074 | 32 | IN_CHECK_OWNER_CH=0x1000 | - |
| `GDMA_IN_CONF1_CH1` | 0x6003F134 | 32 | IN_CHECK_OWNER_CH=0x1000 | - |
| `GDMA_IN_CONF1_CH2` | 0x6003F1F4 | 32 | IN_CHECK_OWNER_CH=0x1000 | - |
| `GDMA_IN_DSCR_BF0_CH0` | 0x6003F094 | 32 | INLINK_DSCR_BF0_CH=0xFFFFFFFF | - |
| `GDMA_IN_DSCR_BF0_CH1` | 0x6003F154 | 32 | INLINK_DSCR_BF0_CH=0xFFFFFFFF | - |
| `GDMA_IN_DSCR_BF0_CH2` | 0x6003F214 | 32 | INLINK_DSCR_BF0_CH=0xFFFFFFFF | - |
| `GDMA_IN_DSCR_BF1_CH0` | 0x6003F098 | 32 | INLINK_DSCR_BF1_CH=0xFFFFFFFF | - |
| `GDMA_IN_DSCR_BF1_CH1` | 0x6003F158 | 32 | INLINK_DSCR_BF1_CH=0xFFFFFFFF | - |
| `GDMA_IN_DSCR_BF1_CH2` | 0x6003F218 | 32 | INLINK_DSCR_BF1_CH=0xFFFFFFFF | - |
| `GDMA_IN_DSCR_CH0` | 0x6003F090 | 32 | INLINK_DSCR_CH=0xFFFFFFFF | - |
| `GDMA_IN_DSCR_CH1` | 0x6003F150 | 32 | INLINK_DSCR_CH=0xFFFFFFFF | - |
| `GDMA_IN_DSCR_CH2` | 0x6003F210 | 32 | INLINK_DSCR_CH=0xFFFFFFFF | - |
| `GDMA_IN_ERR_EOF_DES_ADDR_CH0` | 0x6003F08C | 32 | IN_ERR_EOF_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_IN_ERR_EOF_DES_ADDR_CH1` | 0x6003F14C | 32 | IN_ERR_EOF_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_IN_ERR_EOF_DES_ADDR_CH2` | 0x6003F20C | 32 | IN_ERR_EOF_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_IN_INT_CLR_CH0` | 0x6003F00C | 32 | INFIFO_OVF_CH_INT_CLR=0x20 INFIFO_UDF_CH_INT_CLR=0x40 IN_DONE_CH_INT_CLR=0x1 IN_DSCR_EMPTY_CH_INT_CLR=0x10 IN_DSCR_ERR_CH_INT_CLR=0x8 IN_ERR_EOF_CH_INT_CLR=0x4 IN_SUC_EOF_CH_INT_CLR=0x2 | - |
| `GDMA_IN_INT_CLR_CH1` | 0x6003F01C | 32 | INFIFO_OVF_CH_INT_CLR=0x20 INFIFO_UDF_CH_INT_CLR=0x40 IN_DONE_CH_INT_CLR=0x1 IN_DSCR_EMPTY_CH_INT_CLR=0x10 IN_DSCR_ERR_CH_INT_CLR=0x8 IN_ERR_EOF_CH_INT_CLR=0x4 IN_SUC_EOF_CH_INT_CLR=0x2 | - |
| `GDMA_IN_INT_CLR_CH2` | 0x6003F02C | 32 | INFIFO_OVF_CH_INT_CLR=0x20 INFIFO_UDF_CH_INT_CLR=0x40 IN_DONE_CH_INT_CLR=0x1 IN_DSCR_EMPTY_CH_INT_CLR=0x10 IN_DSCR_ERR_CH_INT_CLR=0x8 IN_ERR_EOF_CH_INT_CLR=0x4 IN_SUC_EOF_CH_INT_CLR=0x2 | - |
| `GDMA_IN_INT_ENA_CH0` | 0x6003F008 | 32 | INFIFO_OVF_CH_INT_ENA=0x20 INFIFO_UDF_CH_INT_ENA=0x40 IN_DONE_CH_INT_ENA=0x1 IN_DSCR_EMPTY_CH_INT_ENA=0x10 IN_DSCR_ERR_CH_INT_ENA=0x8 IN_ERR_EOF_CH_INT_ENA=0x4 IN_SUC_EOF_CH_INT_ENA=0x2 | - |
| `GDMA_IN_INT_ENA_CH1` | 0x6003F018 | 32 | INFIFO_OVF_CH_INT_ENA=0x20 INFIFO_UDF_CH_INT_ENA=0x40 IN_DONE_CH_INT_ENA=0x1 IN_DSCR_EMPTY_CH_INT_ENA=0x10 IN_DSCR_ERR_CH_INT_ENA=0x8 IN_ERR_EOF_CH_INT_ENA=0x4 IN_SUC_EOF_CH_INT_ENA=0x2 | - |
| `GDMA_IN_INT_ENA_CH2` | 0x6003F028 | 32 | INFIFO_OVF_CH_INT_ENA=0x20 INFIFO_UDF_CH_INT_ENA=0x40 IN_DONE_CH_INT_ENA=0x1 IN_DSCR_EMPTY_CH_INT_ENA=0x10 IN_DSCR_ERR_CH_INT_ENA=0x8 IN_ERR_EOF_CH_INT_ENA=0x4 IN_SUC_EOF_CH_INT_ENA=0x2 | - |
| `GDMA_IN_INT_RAW_CH0` | 0x6003F000 | 32 | INFIFO_OVF_CH_INT_RAW=0x20 INFIFO_UDF_CH_INT_RAW=0x40 IN_DONE_CH_INT_RAW=0x1 IN_DSCR_EMPTY_CH_INT_RAW=0x10 IN_DSCR_ERR_CH_INT_RAW=0x8 IN_ERR_EOF_CH_INT_RAW=0x4 IN_SUC_EOF_CH_INT_RAW=0x2 | - |
| `GDMA_IN_INT_RAW_CH1` | 0x6003F010 | 32 | INFIFO_OVF_CH_INT_RAW=0x20 INFIFO_UDF_CH_INT_RAW=0x40 IN_DONE_CH_INT_RAW=0x1 IN_DSCR_EMPTY_CH_INT_RAW=0x10 IN_DSCR_ERR_CH_INT_RAW=0x8 IN_ERR_EOF_CH_INT_RAW=0x4 IN_SUC_EOF_CH_INT_RAW=0x2 | - |
| `GDMA_IN_INT_RAW_CH2` | 0x6003F020 | 32 | INFIFO_OVF_CH_INT_RAW=0x20 INFIFO_UDF_CH_INT_RAW=0x40 IN_DONE_CH_INT_RAW=0x1 IN_DSCR_EMPTY_CH_INT_RAW=0x10 IN_DSCR_ERR_CH_INT_RAW=0x8 IN_ERR_EOF_CH_INT_RAW=0x4 IN_SUC_EOF_CH_INT_RAW=0x2 | - |
| `GDMA_IN_INT_ST_CH0` | 0x6003F004 | 32 | INFIFO_OVF_CH_INT_ST=0x20 INFIFO_UDF_CH_INT_ST=0x40 IN_DONE_CH_INT_ST=0x1 IN_DSCR_EMPTY_CH_INT_ST=0x10 IN_DSCR_ERR_CH_INT_ST=0x8 IN_ERR_EOF_CH_INT_ST=0x4 IN_SUC_EOF_CH_INT_ST=0x2 | - |
| `GDMA_IN_INT_ST_CH1` | 0x6003F014 | 32 | INFIFO_OVF_CH_INT_ST=0x20 INFIFO_UDF_CH_INT_ST=0x40 IN_DONE_CH_INT_ST=0x1 IN_DSCR_EMPTY_CH_INT_ST=0x10 IN_DSCR_ERR_CH_INT_ST=0x8 IN_ERR_EOF_CH_INT_ST=0x4 IN_SUC_EOF_CH_INT_ST=0x2 | - |
| `GDMA_IN_INT_ST_CH2` | 0x6003F024 | 32 | INFIFO_OVF_CH_INT_ST=0x20 INFIFO_UDF_CH_INT_ST=0x40 IN_DONE_CH_INT_ST=0x1 IN_DSCR_EMPTY_CH_INT_ST=0x10 IN_DSCR_ERR_CH_INT_ST=0x8 IN_ERR_EOF_CH_INT_ST=0x4 IN_SUC_EOF_CH_INT_ST=0x2 | - |
| `GDMA_IN_LINK_CH0` | 0x6003F080 | 32 | INLINK_ADDR_CH=0xFFFFF INLINK_AUTO_RET_CH=0x100000 INLINK_PARK_CH=0x1000000 INLINK_RESTART_CH=0x800000 INLINK_START_CH=0x400000 INLINK_STOP_CH=0x200000 | - |
| `GDMA_IN_LINK_CH1` | 0x6003F140 | 32 | INLINK_ADDR_CH=0xFFFFF INLINK_AUTO_RET_CH=0x100000 INLINK_PARK_CH=0x1000000 INLINK_RESTART_CH=0x800000 INLINK_START_CH=0x400000 INLINK_STOP_CH=0x200000 | - |
| `GDMA_IN_LINK_CH2` | 0x6003F200 | 32 | INLINK_ADDR_CH=0xFFFFF INLINK_AUTO_RET_CH=0x100000 INLINK_PARK_CH=0x1000000 INLINK_RESTART_CH=0x800000 INLINK_START_CH=0x400000 INLINK_STOP_CH=0x200000 | - |
| `GDMA_IN_PERI_SEL_CH0` | 0x6003F0A0 | 32 | PERI_IN_SEL_CH=0x3F | - |
| `GDMA_IN_PERI_SEL_CH1` | 0x6003F160 | 32 | PERI_IN_SEL_CH=0x3F | - |
| `GDMA_IN_PERI_SEL_CH2` | 0x6003F220 | 32 | PERI_IN_SEL_CH=0x3F | - |
| `GDMA_IN_POP_CH0` | 0x6003F07C | 32 | INFIFO_POP_CH=0x1000 INFIFO_RDATA_CH=0xFFF | - |
| `GDMA_IN_POP_CH1` | 0x6003F13C | 32 | INFIFO_POP_CH=0x1000 INFIFO_RDATA_CH=0xFFF | - |
| `GDMA_IN_POP_CH2` | 0x6003F1FC | 32 | INFIFO_POP_CH=0x1000 INFIFO_RDATA_CH=0xFFF | - |
| `GDMA_IN_PRI_CH0` | 0x6003F09C | 32 | RX_PRI_CH=0xF | - |
| `GDMA_IN_PRI_CH1` | 0x6003F15C | 32 | RX_PRI_CH=0xF | - |
| `GDMA_IN_PRI_CH2` | 0x6003F21C | 32 | RX_PRI_CH=0xF | - |
| `GDMA_IN_STATE_CH0` | 0x6003F084 | 32 | INLINK_DSCR_ADDR_CH=0x3FFFF IN_DSCR_STATE_CH=0xC0000 IN_STATE_CH=0x700000 | - |
| `GDMA_IN_STATE_CH1` | 0x6003F144 | 32 | INLINK_DSCR_ADDR_CH=0x3FFFF IN_DSCR_STATE_CH=0xC0000 IN_STATE_CH=0x700000 | - |
| `GDMA_IN_STATE_CH2` | 0x6003F204 | 32 | INLINK_DSCR_ADDR_CH=0x3FFFF IN_DSCR_STATE_CH=0xC0000 IN_STATE_CH=0x700000 | - |
| `GDMA_IN_SUC_EOF_DES_ADDR_CH0` | 0x6003F088 | 32 | IN_SUC_EOF_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_IN_SUC_EOF_DES_ADDR_CH1` | 0x6003F148 | 32 | IN_SUC_EOF_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_IN_SUC_EOF_DES_ADDR_CH2` | 0x6003F208 | 32 | IN_SUC_EOF_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_MISC_CONF` | 0x6003F064 | 32 | AHBM_RST_INTER=0x1 ARB_PRI_DIS=0x4 CLK_EN=0x8 | - |
| `GDMA_OUTFIFO_STATUS_CH0` | 0x6003F0D8 | 32 | OUTFIFO_CNT_CH=0xFC OUTFIFO_EMPTY_CH=0x2 OUTFIFO_FULL_CH=0x1 OUT_REMAIN_UNDER_1B_CH=0x800000 OUT_REMAIN_UNDER_2B_CH=0x1000000 OUT_REMAIN_UNDER_3B_CH=0x2000000 OUT_REMAIN_UNDER_4B_CH=0x4000000 | - |
| `GDMA_OUTFIFO_STATUS_CH1` | 0x6003F198 | 32 | OUTFIFO_CNT_CH=0xFC OUTFIFO_EMPTY_CH=0x2 OUTFIFO_FULL_CH=0x1 OUT_REMAIN_UNDER_1B_CH=0x800000 OUT_REMAIN_UNDER_2B_CH=0x1000000 OUT_REMAIN_UNDER_3B_CH=0x2000000 OUT_REMAIN_UNDER_4B_CH=0x4000000 | - |
| `GDMA_OUTFIFO_STATUS_CH2` | 0x6003F258 | 32 | OUTFIFO_CNT_CH=0xFC OUTFIFO_EMPTY_CH=0x2 OUTFIFO_FULL_CH=0x1 OUT_REMAIN_UNDER_1B_CH=0x800000 OUT_REMAIN_UNDER_2B_CH=0x1000000 OUT_REMAIN_UNDER_3B_CH=0x2000000 OUT_REMAIN_UNDER_4B_CH=0x4000000 | - |
| `GDMA_OUT_CONF0_CH0` | 0x6003F190 | 32 | OUTDSCR_BURST_EN_CH=0x10 OUT_AUTO_WRBACK_CH=0x4 OUT_DATA_BURST_EN_CH=0x20 OUT_EOF_MODE_CH=0x8 OUT_ETM_EN_CH=0x40 OUT_LOOP_TEST_CH=0x2 OUT_RST_CH=0x1 | - |
| `GDMA_OUT_CONF0_CH1` | 0x6003F250 | 32 | OUTDSCR_BURST_EN_CH=0x10 OUT_AUTO_WRBACK_CH=0x4 OUT_DATA_BURST_EN_CH=0x20 OUT_EOF_MODE_CH=0x8 OUT_ETM_EN_CH=0x40 OUT_LOOP_TEST_CH=0x2 OUT_RST_CH=0x1 | - |
| `GDMA_OUT_CONF1_CH0` | 0x6003F0D4 | 32 | OUT_CHECK_OWNER_CH=0x1000 | - |
| `GDMA_OUT_CONF1_CH1` | 0x6003F194 | 32 | OUT_CHECK_OWNER_CH=0x1000 | - |
| `GDMA_OUT_CONF1_CH2` | 0x6003F254 | 32 | OUT_CHECK_OWNER_CH=0x1000 | - |
| `GDMA_OUT_DSCR_BF0_CH0` | 0x6003F0F4 | 32 | OUTLINK_DSCR_BF0_CH=0xFFFFFFFF | - |
| `GDMA_OUT_DSCR_BF0_CH1` | 0x6003F1B4 | 32 | OUTLINK_DSCR_BF0_CH=0xFFFFFFFF | - |
| `GDMA_OUT_DSCR_BF0_CH2` | 0x6003F274 | 32 | OUTLINK_DSCR_BF0_CH=0xFFFFFFFF | - |
| `GDMA_OUT_DSCR_BF1_CH0` | 0x6003F0F8 | 32 | OUTLINK_DSCR_BF1_CH=0xFFFFFFFF | - |
| `GDMA_OUT_DSCR_BF1_CH1` | 0x6003F1B8 | 32 | OUTLINK_DSCR_BF1_CH=0xFFFFFFFF | - |
| `GDMA_OUT_DSCR_BF1_CH2` | 0x6003F278 | 32 | OUTLINK_DSCR_BF1_CH=0xFFFFFFFF | - |
| `GDMA_OUT_DSCR_CH0` | 0x6003F0F0 | 32 | OUTLINK_DSCR_CH=0xFFFFFFFF | - |
| `GDMA_OUT_DSCR_CH1` | 0x6003F1B0 | 32 | OUTLINK_DSCR_CH=0xFFFFFFFF | - |
| `GDMA_OUT_DSCR_CH2` | 0x6003F270 | 32 | OUTLINK_DSCR_CH=0xFFFFFFFF | - |
| `GDMA_OUT_EOF_BFR_DES_ADDR_CH0` | 0x6003F0EC | 32 | OUT_EOF_BFR_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_OUT_EOF_BFR_DES_ADDR_CH1` | 0x6003F1AC | 32 | OUT_EOF_BFR_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_OUT_EOF_BFR_DES_ADDR_CH2` | 0x6003F26C | 32 | OUT_EOF_BFR_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_OUT_EOF_DES_ADDR_CH0` | 0x6003F0E8 | 32 | OUT_EOF_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_OUT_EOF_DES_ADDR_CH1` | 0x6003F1A8 | 32 | OUT_EOF_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_OUT_EOF_DES_ADDR_CH2` | 0x6003F268 | 32 | OUT_EOF_DES_ADDR_CH=0xFFFFFFFF | - |
| `GDMA_OUT_INT_CLR_CH0` | 0x6003F03C | 32 | OUTFIFO_OVF_CH_INT_CLR=0x10 OUTFIFO_UDF_CH_INT_CLR=0x20 OUT_DONE_CH_INT_CLR=0x1 OUT_DSCR_ERR_CH_INT_CLR=0x4 OUT_EOF_CH_INT_CLR=0x2 OUT_TOTAL_EOF_CH_INT_CLR=0x8 | - |
| `GDMA_OUT_INT_CLR_CH1` | 0x6003F04C | 32 | OUTFIFO_OVF_CH_INT_CLR=0x10 OUTFIFO_UDF_CH_INT_CLR=0x20 OUT_DONE_CH_INT_CLR=0x1 OUT_DSCR_ERR_CH_INT_CLR=0x4 OUT_EOF_CH_INT_CLR=0x2 OUT_TOTAL_EOF_CH_INT_CLR=0x8 | - |
| `GDMA_OUT_INT_CLR_CH2` | 0x6003F05C | 32 | OUTFIFO_OVF_CH_INT_CLR=0x10 OUTFIFO_UDF_CH_INT_CLR=0x20 OUT_DONE_CH_INT_CLR=0x1 OUT_DSCR_ERR_CH_INT_CLR=0x4 OUT_EOF_CH_INT_CLR=0x2 OUT_TOTAL_EOF_CH_INT_CLR=0x8 | - |
| `GDMA_OUT_INT_ENA_CH0` | 0x6003F038 | 32 | OUTFIFO_OVF_CH_INT_ENA=0x10 OUTFIFO_UDF_CH_INT_ENA=0x20 OUT_DONE_CH_INT_ENA=0x1 OUT_DSCR_ERR_CH_INT_ENA=0x4 OUT_EOF_CH_INT_ENA=0x2 OUT_TOTAL_EOF_CH_INT_ENA=0x8 | - |
| `GDMA_OUT_INT_ENA_CH1` | 0x6003F048 | 32 | OUTFIFO_OVF_CH_INT_ENA=0x10 OUTFIFO_UDF_CH_INT_ENA=0x20 OUT_DONE_CH_INT_ENA=0x1 OUT_DSCR_ERR_CH_INT_ENA=0x4 OUT_EOF_CH_INT_ENA=0x2 OUT_TOTAL_EOF_CH_INT_ENA=0x8 | - |
| `GDMA_OUT_INT_ENA_CH2` | 0x6003F058 | 32 | OUTFIFO_OVF_CH_INT_ENA=0x10 OUTFIFO_UDF_CH_INT_ENA=0x20 OUT_DONE_CH_INT_ENA=0x1 OUT_DSCR_ERR_CH_INT_ENA=0x4 OUT_EOF_CH_INT_ENA=0x2 OUT_TOTAL_EOF_CH_INT_ENA=0x8 | - |
| `GDMA_OUT_INT_RAW_CH0` | 0x6003F030 | 32 | OUTFIFO_OVF_CH_INT_RAW=0x10 OUTFIFO_UDF_CH_INT_RAW=0x20 OUT_DONE_CH_INT_RAW=0x1 OUT_DSCR_ERR_CH_INT_RAW=0x4 OUT_EOF_CH_INT_RAW=0x2 OUT_TOTAL_EOF_CH_INT_RAW=0x8 | - |
| `GDMA_OUT_INT_RAW_CH1` | 0x6003F040 | 32 | OUTFIFO_OVF_CH_INT_RAW=0x10 OUTFIFO_UDF_CH_INT_RAW=0x20 OUT_DONE_CH_INT_RAW=0x1 OUT_DSCR_ERR_CH_INT_RAW=0x4 OUT_EOF_CH_INT_RAW=0x2 OUT_TOTAL_EOF_CH_INT_RAW=0x8 | - |
| `GDMA_OUT_INT_RAW_CH2` | 0x6003F050 | 32 | OUTFIFO_OVF_CH_INT_RAW=0x10 OUTFIFO_UDF_CH_INT_RAW=0x20 OUT_DONE_CH_INT_RAW=0x1 OUT_DSCR_ERR_CH_INT_RAW=0x4 OUT_EOF_CH_INT_RAW=0x2 OUT_TOTAL_EOF_CH_INT_RAW=0x8 | - |
| `GDMA_OUT_INT_ST_CH0` | 0x6003F034 | 32 | OUTFIFO_OVF_CH_INT_ST=0x10 OUTFIFO_UDF_CH_INT_ST=0x20 OUT_DONE_CH_INT_ST=0x1 OUT_DSCR_ERR_CH_INT_ST=0x4 OUT_EOF_CH_INT_ST=0x2 OUT_TOTAL_EOF_CH_INT_ST=0x8 | - |
| `GDMA_OUT_INT_ST_CH1` | 0x6003F044 | 32 | OUTFIFO_OVF_CH_INT_ST=0x10 OUTFIFO_UDF_CH_INT_ST=0x20 OUT_DONE_CH_INT_ST=0x1 OUT_DSCR_ERR_CH_INT_ST=0x4 OUT_EOF_CH_INT_ST=0x2 OUT_TOTAL_EOF_CH_INT_ST=0x8 | - |
| `GDMA_OUT_INT_ST_CH2` | 0x6003F054 | 32 | OUTFIFO_OVF_CH_INT_ST=0x10 OUTFIFO_UDF_CH_INT_ST=0x20 OUT_DONE_CH_INT_ST=0x1 OUT_DSCR_ERR_CH_INT_ST=0x4 OUT_EOF_CH_INT_ST=0x2 OUT_TOTAL_EOF_CH_INT_ST=0x8 | - |
| `GDMA_OUT_LINK_CH0` | 0x6003F0E0 | 32 | OUTLINK_ADDR_CH=0xFFFFF OUTLINK_PARK_CH=0x800000 OUTLINK_RESTART_CH=0x400000 OUTLINK_START_CH=0x200000 OUTLINK_STOP_CH=0x100000 | - |
| `GDMA_OUT_LINK_CH1` | 0x6003F1A0 | 32 | OUTLINK_ADDR_CH=0xFFFFF OUTLINK_PARK_CH=0x800000 OUTLINK_RESTART_CH=0x400000 OUTLINK_START_CH=0x200000 OUTLINK_STOP_CH=0x100000 | - |
| `GDMA_OUT_LINK_CH2` | 0x6003F260 | 32 | OUTLINK_ADDR_CH=0xFFFFF OUTLINK_PARK_CH=0x800000 OUTLINK_RESTART_CH=0x400000 OUTLINK_START_CH=0x200000 OUTLINK_STOP_CH=0x100000 | - |
| `GDMA_OUT_PERI_SEL_CH0` | 0x6003F100 | 32 | PERI_OUT_SEL_CH=0x3F | - |
| `GDMA_OUT_PERI_SEL_CH1` | 0x6003F1C0 | 32 | PERI_OUT_SEL_CH=0x3F | - |
| `GDMA_OUT_PERI_SEL_CH2` | 0x6003F280 | 32 | PERI_OUT_SEL_CH=0x3F | - |
| `GDMA_OUT_PRI_CH0` | 0x6003F0FC | 32 | TX_PRI_CH=0xF | - |
| `GDMA_OUT_PRI_CH1` | 0x6003F1BC | 32 | TX_PRI_CH=0xF | - |
| `GDMA_OUT_PRI_CH2` | 0x6003F27C | 32 | TX_PRI_CH=0xF | - |
| `GDMA_OUT_PUSH_CH0` | 0x6003F0DC | 32 | OUTFIFO_PUSH_CH=0x200 OUTFIFO_WDATA_CH=0x1FF | - |
| `GDMA_OUT_PUSH_CH1` | 0x6003F19C | 32 | OUTFIFO_PUSH_CH=0x200 OUTFIFO_WDATA_CH=0x1FF | - |
| `GDMA_OUT_PUSH_CH2` | 0x6003F25C | 32 | OUTFIFO_PUSH_CH=0x200 OUTFIFO_WDATA_CH=0x1FF | - |
| `GDMA_OUT_STATE_CH0` | 0x6003F0E4 | 32 | OUTLINK_DSCR_ADDR_CH=0x3FFFF OUT_DSCR_STATE_CH=0xC0000 OUT_STATE_CH=0x700000 | - |
| `GDMA_OUT_STATE_CH1` | 0x6003F1A4 | 32 | OUTLINK_DSCR_ADDR_CH=0x3FFFF OUT_DSCR_STATE_CH=0xC0000 OUT_STATE_CH=0x700000 | - |
| `GDMA_OUT_STATE_CH2` | 0x6003F264 | 32 | OUTLINK_DSCR_ADDR_CH=0x3FFFF OUT_DSCR_STATE_CH=0xC0000 OUT_STATE_CH=0x700000 | - |
| `GPIO_BT_SELECT` | 0x60004000 | 32 | BT_SEL=0xFFFFFFFF | - |
| `GPIO_CLOCK_GATE` | 0x6000462C | 32 | CLK_EN=0x1 | - |
| `GPIO_CPUSDIO_INT` | 0x60004064 | 32 | SDIO_INT=0xFFFFFFFF | - |
| `GPIO_DATE` | 0x600046FC | 32 | DATE=0xFFFFFFF | - |
| `GPIO_ENABLE` | 0x60004020 | 32 | DATA=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27 |
| `GPIO_ENABLE_W1TC` | 0x60004028 | 32 | ENABLE_W1TC=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27 |
| `GPIO_ENABLE_W1TS` | 0x60004024 | 32 | ENABLE_W1TS=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27 |
| `GPIO_FUNC0_IN_SEL_CFG` | 0x60004154 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC0_OUT_SEL_CFG` | 0x60004554 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC100_IN_SEL_CFG` | 0x600042E4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC101_IN_SEL_CFG` | 0x600042E8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC102_IN_SEL_CFG` | 0x600042EC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC103_IN_SEL_CFG` | 0x600042F0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC104_IN_SEL_CFG` | 0x600042F4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC105_IN_SEL_CFG` | 0x600042F8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC106_IN_SEL_CFG` | 0x600042FC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC107_IN_SEL_CFG` | 0x60004300 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC108_IN_SEL_CFG` | 0x60004304 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC109_IN_SEL_CFG` | 0x60004308 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC10_IN_SEL_CFG` | 0x6000417C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC10_OUT_SEL_CFG` | 0x6000457C | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC110_IN_SEL_CFG` | 0x6000430C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC111_IN_SEL_CFG` | 0x60004310 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC112_IN_SEL_CFG` | 0x60004314 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC113_IN_SEL_CFG` | 0x60004318 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC114_IN_SEL_CFG` | 0x6000431C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC115_IN_SEL_CFG` | 0x60004320 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC116_IN_SEL_CFG` | 0x60004324 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC117_IN_SEL_CFG` | 0x60004328 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC118_IN_SEL_CFG` | 0x6000432C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC119_IN_SEL_CFG` | 0x60004330 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC11_IN_SEL_CFG` | 0x60004180 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC11_OUT_SEL_CFG` | 0x60004580 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC120_IN_SEL_CFG` | 0x60004334 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC121_IN_SEL_CFG` | 0x60004338 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC122_IN_SEL_CFG` | 0x6000433C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC123_IN_SEL_CFG` | 0x60004340 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC124_IN_SEL_CFG` | 0x60004344 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC12_IN_SEL_CFG` | 0x60004184 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC12_OUT_SEL_CFG` | 0x60004584 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC13_IN_SEL_CFG` | 0x60004188 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC13_OUT_SEL_CFG` | 0x60004588 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC14_IN_SEL_CFG` | 0x6000418C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC14_OUT_SEL_CFG` | 0x6000458C | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC15_IN_SEL_CFG` | 0x60004190 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC15_OUT_SEL_CFG` | 0x60004590 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC16_IN_SEL_CFG` | 0x60004194 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC16_OUT_SEL_CFG` | 0x60004594 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC17_IN_SEL_CFG` | 0x60004198 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC17_OUT_SEL_CFG` | 0x60004598 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC18_OUT_SEL_CFG` | 0x6000459C | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC19_IN_SEL_CFG` | 0x600041A0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC19_OUT_SEL_CFG` | 0x600045A0 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC1_OUT_SEL_CFG` | 0x60004558 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC20_OUT_SEL_CFG` | 0x600045A4 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC21_OUT_SEL_CFG` | 0x600045A8 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC22_OUT_SEL_CFG` | 0x600045AC | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC23_OUT_SEL_CFG` | 0x600045B0 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC24_OUT_SEL_CFG` | 0x600045B4 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC25_OUT_SEL_CFG` | 0x600045B8 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC26_OUT_SEL_CFG` | 0x600045BC | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC27_OUT_SEL_CFG` | 0x600045C0 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC28_IN_SEL_CFG` | 0x600041C4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC28_OUT_SEL_CFG` | 0x600045C4 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC29_IN_SEL_CFG` | 0x600041C8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC29_OUT_SEL_CFG` | 0x600045C8 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC2_OUT_SEL_CFG` | 0x6000455C | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC30_IN_SEL_CFG` | 0x600041CC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC30_OUT_SEL_CFG` | 0x600045CC | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC31_IN_SEL_CFG` | 0x600041D0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC31_OUT_SEL_CFG` | 0x600045D0 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC32_IN_SEL_CFG` | 0x600041D4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC33_IN_SEL_CFG` | 0x600041D8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC34_IN_SEL_CFG` | 0x600041DC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC35_IN_SEL_CFG` | 0x600041E0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC3_OUT_SEL_CFG` | 0x60004560 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC40_IN_SEL_CFG` | 0x600041F4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC41_IN_SEL_CFG` | 0x600041F8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC42_IN_SEL_CFG` | 0x600041FC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC45_IN_SEL_CFG` | 0x60004208 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC46_IN_SEL_CFG` | 0x6000420C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC47_IN_SEL_CFG` | 0x60004210 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC48_IN_SEL_CFG` | 0x60004214 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC49_IN_SEL_CFG` | 0x60004218 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC4_OUT_SEL_CFG` | 0x60004564 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC50_IN_SEL_CFG` | 0x6000421C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC51_IN_SEL_CFG` | 0x60004220 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC52_IN_SEL_CFG` | 0x60004224 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC53_IN_SEL_CFG` | 0x60004228 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC54_IN_SEL_CFG` | 0x6000422C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC55_IN_SEL_CFG` | 0x60004230 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC56_IN_SEL_CFG` | 0x60004234 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC5_OUT_SEL_CFG` | 0x60004568 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC63_IN_SEL_CFG` | 0x60004250 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC64_IN_SEL_CFG` | 0x60004254 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC65_IN_SEL_CFG` | 0x60004258 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC66_IN_SEL_CFG` | 0x6000425C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC67_IN_SEL_CFG` | 0x60004260 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC68_IN_SEL_CFG` | 0x60004264 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC69_IN_SEL_CFG` | 0x60004268 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC6_IN_SEL_CFG` | 0x6000416C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC6_OUT_SEL_CFG` | 0x6000456C | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC70_IN_SEL_CFG` | 0x6000426C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC71_IN_SEL_CFG` | 0x60004270 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC72_IN_SEL_CFG` | 0x60004274 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC73_IN_SEL_CFG` | 0x60004278 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC7_IN_SEL_CFG` | 0x60004170 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC7_OUT_SEL_CFG` | 0x60004570 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC81_IN_SEL_CFG` | 0x60004298 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC82_IN_SEL_CFG` | 0x6000429C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC87_IN_SEL_CFG` | 0x600042B0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC88_IN_SEL_CFG` | 0x600042B4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC89_IN_SEL_CFG` | 0x600042B8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC8_IN_SEL_CFG` | 0x60004174 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC8_OUT_SEL_CFG` | 0x60004574 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_FUNC90_IN_SEL_CFG` | 0x600042BC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC91_IN_SEL_CFG` | 0x600042C0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC92_IN_SEL_CFG` | 0x600042C4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC93_IN_SEL_CFG` | 0x600042C8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC94_IN_SEL_CFG` | 0x600042CC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC95_IN_SEL_CFG` | 0x600042D0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC97_IN_SEL_CFG` | 0x600042D8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC98_IN_SEL_CFG` | 0x600042DC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC99_IN_SEL_CFG` | 0x600042E0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC9_IN_SEL_CFG` | 0x60004178 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC9_OUT_SEL_CFG` | 0x60004578 | 32 | INV_SEL=0x100 OEN_INV_SEL=0x400 OEN_SEL=0x200 OUT_SEL=0xFF | - |
| `GPIO_IN` | 0x6000403C | 32 | DATA_NEXT=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27 |
| `GPIO_OUT` | 0x60004004 | 32 | DATA_ORIG=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27 |
| `GPIO_OUT_W1TC` | 0x6000400C | 32 | OUT_W1TC=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27 |
| `GPIO_OUT_W1TS` | 0x60004008 | 32 | OUT_W1TS=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27 |
| `GPIO_PCPU_INT` | 0x6000405C | 32 | PROCPU_INT=0xFFFFFFFF | - |
| `GPIO_PCPU_NMI_INT` | 0x60004060 | 32 | PROCPU_NMI_INT=0xFFFFFFFF | - |
| `GPIO_PIN0` | 0x60004074 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN1` | 0x60004078 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN10` | 0x6000409C | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN11` | 0x600040A0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN12` | 0x600040A4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN13` | 0x600040A8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN14` | 0x600040AC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN15` | 0x600040B0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN16` | 0x600040B4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN17` | 0x600040B8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN18` | 0x600040BC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN19` | 0x600040C0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN2` | 0x6000407C | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN20` | 0x600040C4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN21` | 0x600040C8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN22` | 0x600040CC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN23` | 0x600040D0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN24` | 0x600040D4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN25` | 0x600040D8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN26` | 0x600040DC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN27` | 0x600040E0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN28` | 0x600040E4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN29` | 0x600040E8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN3` | 0x60004080 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN30` | 0x600040EC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN31` | 0x600040F0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN4` | 0x60004084 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN5` | 0x60004088 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN6` | 0x6000408C | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN7` | 0x60004090 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN8` | 0x60004094 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN9` | 0x60004098 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_SDIO_SELECT` | 0x6000401C | 32 | SDIO_SEL=0xFF | - |
| `GPIO_STATUS` | 0x60004044 | 32 | INTERRUPT=0xFFFFFFFF | - |
| `GPIO_STATUS_NEXT` | 0x6000414C | 32 | STATUS_INTERRUPT_NEXT=0xFFFFFFFF | - |
| `GPIO_STATUS_W1TC` | 0x6000404C | 32 | STATUS_W1TC=0xFFFFFFFF | - |
| `GPIO_STATUS_W1TS` | 0x60004048 | 32 | STATUS_W1TS=0xFFFFFFFF | - |
| `GPIO_STRAP` | 0x60004038 | 32 | STRAPPING=0xFFFF | - |
| `HMAC_DATE` | 0x6003E1FC | 32 | DATE=0x3FFFFFFF | - |
| `HMAC_ONE_BLOCK` | 0x6003E0F4 | 32 | SET_ONE_BLOCK=0x1 | - |
| `HMAC_QUERY_BUSY` | 0x6003E06C | 32 | BUSY_STATE=0x1 | - |
| `HMAC_QUERY_ERROR` | 0x6003E068 | 32 | QUREY_CHECK=0x1 | - |
| `HMAC_RD_RESULT_MEM[0]` | 0x6003E0C0 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[10]` | 0x6003E0CA | 8 |  | - |
| `HMAC_RD_RESULT_MEM[11]` | 0x6003E0CB | 8 |  | - |
| `HMAC_RD_RESULT_MEM[12]` | 0x6003E0CC | 8 |  | - |
| `HMAC_RD_RESULT_MEM[13]` | 0x6003E0CD | 8 |  | - |
| `HMAC_RD_RESULT_MEM[14]` | 0x6003E0CE | 8 |  | - |
| `HMAC_RD_RESULT_MEM[15]` | 0x6003E0CF | 8 |  | - |
| `HMAC_RD_RESULT_MEM[16]` | 0x6003E0D0 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[17]` | 0x6003E0D1 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[18]` | 0x6003E0D2 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[19]` | 0x6003E0D3 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[1]` | 0x6003E0C1 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[20]` | 0x6003E0D4 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[21]` | 0x6003E0D5 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[22]` | 0x6003E0D6 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[23]` | 0x6003E0D7 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[24]` | 0x6003E0D8 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[25]` | 0x6003E0D9 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[26]` | 0x6003E0DA | 8 |  | - |
| `HMAC_RD_RESULT_MEM[27]` | 0x6003E0DB | 8 |  | - |
| `HMAC_RD_RESULT_MEM[28]` | 0x6003E0DC | 8 |  | - |
| `HMAC_RD_RESULT_MEM[29]` | 0x6003E0DD | 8 |  | - |
| `HMAC_RD_RESULT_MEM[2]` | 0x6003E0C2 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[30]` | 0x6003E0DE | 8 |  | - |
| `HMAC_RD_RESULT_MEM[31]` | 0x6003E0DF | 8 |  | - |
| `HMAC_RD_RESULT_MEM[3]` | 0x6003E0C3 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[4]` | 0x6003E0C4 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[5]` | 0x6003E0C5 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[6]` | 0x6003E0C6 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[7]` | 0x6003E0C7 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[8]` | 0x6003E0C8 | 8 |  | - |
| `HMAC_RD_RESULT_MEM[9]` | 0x6003E0C9 | 8 |  | - |
| `HMAC_SET_INVALIDATE_DS` | 0x6003E064 | 32 | SET_INVALIDATE_DS=0x1 | - |
| `HMAC_SET_INVALIDATE_JTAG` | 0x6003E060 | 32 | SET_INVALIDATE_JTAG=0x1 | - |
| `HMAC_SET_MESSAGE_END` | 0x6003E058 | 32 | SET_TEXT_END=0x1 | - |
| `HMAC_SET_MESSAGE_ING` | 0x6003E054 | 32 | SET_TEXT_ING=0x1 | - |
| `HMAC_SET_MESSAGE_ONE` | 0x6003E050 | 32 | SET_TEXT_ONE=0x1 | - |
| `HMAC_SET_MESSAGE_PAD` | 0x6003E0F0 | 32 | SET_TEXT_PAD=0x1 | - |
| `HMAC_SET_PARA_FINISH` | 0x6003E04C | 32 | SET_PARA_END=0x1 | - |
| `HMAC_SET_PARA_KEY` | 0x6003E048 | 32 | KEY_SET=0x7 | - |
| `HMAC_SET_PARA_PURPOSE` | 0x6003E044 | 32 | PURPOSE_SET=0xF | - |
| `HMAC_SET_RESULT_FINISH` | 0x6003E05C | 32 | SET_RESULT_END=0x1 | - |
| `HMAC_SET_START` | 0x6003E040 | 32 | SET_START=0x1 | - |
| `HMAC_SOFT_JTAG_CTRL` | 0x6003E0F8 | 32 | SOFT_JTAG_CTRL=0x1 | - |
| `HMAC_WR_JTAG` | 0x6003E0FC | 32 | WR_JTAG=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_MEM[0]` | 0x6003E080 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[10]` | 0x6003E08A | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[11]` | 0x6003E08B | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[12]` | 0x6003E08C | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[13]` | 0x6003E08D | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[14]` | 0x6003E08E | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[15]` | 0x6003E08F | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[16]` | 0x6003E090 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[17]` | 0x6003E091 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[18]` | 0x6003E092 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[19]` | 0x6003E093 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[1]` | 0x6003E081 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[20]` | 0x6003E094 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[21]` | 0x6003E095 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[22]` | 0x6003E096 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[23]` | 0x6003E097 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[24]` | 0x6003E098 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[25]` | 0x6003E099 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[26]` | 0x6003E09A | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[27]` | 0x6003E09B | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[28]` | 0x6003E09C | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[29]` | 0x6003E09D | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[2]` | 0x6003E082 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[30]` | 0x6003E09E | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[31]` | 0x6003E09F | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[32]` | 0x6003E0A0 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[33]` | 0x6003E0A1 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[34]` | 0x6003E0A2 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[35]` | 0x6003E0A3 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[36]` | 0x6003E0A4 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[37]` | 0x6003E0A5 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[38]` | 0x6003E0A6 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[39]` | 0x6003E0A7 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[3]` | 0x6003E083 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[40]` | 0x6003E0A8 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[41]` | 0x6003E0A9 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[42]` | 0x6003E0AA | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[43]` | 0x6003E0AB | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[44]` | 0x6003E0AC | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[45]` | 0x6003E0AD | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[46]` | 0x6003E0AE | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[47]` | 0x6003E0AF | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[48]` | 0x6003E0B0 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[49]` | 0x6003E0B1 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[4]` | 0x6003E084 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[50]` | 0x6003E0B2 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[51]` | 0x6003E0B3 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[52]` | 0x6003E0B4 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[53]` | 0x6003E0B5 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[54]` | 0x6003E0B6 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[55]` | 0x6003E0B7 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[56]` | 0x6003E0B8 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[57]` | 0x6003E0B9 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[58]` | 0x6003E0BA | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[59]` | 0x6003E0BB | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[5]` | 0x6003E085 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[60]` | 0x6003E0BC | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[61]` | 0x6003E0BD | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[62]` | 0x6003E0BE | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[63]` | 0x6003E0BF | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[6]` | 0x6003E086 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[7]` | 0x6003E087 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[8]` | 0x6003E088 | 8 |  | - |
| `HMAC_WR_MESSAGE_MEM[9]` | 0x6003E089 | 8 |  | - |
| `I2C0_CLK_CONF` | 0x60013054 | 32 | SCLK_ACTIVE=0x200000 SCLK_DIV_A=0x3F00 SCLK_DIV_B=0xFC000 SCLK_DIV_NUM=0xFF SCLK_SEL=0x100000 | - |
| `I2C0_COMD0` | 0x60013058 | 32 | COMMAND0=0x3FFF COMMAND0_DONE=0x80000000 | - |
| `I2C0_COMD1` | 0x6001305C | 32 | COMMAND0=0x3FFF COMMAND0_DONE=0x80000000 | - |
| `I2C0_COMD2` | 0x60013060 | 32 | COMMAND0=0x3FFF COMMAND0_DONE=0x80000000 | - |
| `I2C0_COMD3` | 0x60013064 | 32 | COMMAND0=0x3FFF COMMAND0_DONE=0x80000000 | - |
| `I2C0_COMD4` | 0x60013068 | 32 | COMMAND0=0x3FFF COMMAND0_DONE=0x80000000 | - |
| `I2C0_COMD5` | 0x6001306C | 32 | COMMAND0=0x3FFF COMMAND0_DONE=0x80000000 | - |
| `I2C0_COMD6` | 0x60013070 | 32 | COMMAND0=0x3FFF COMMAND0_DONE=0x80000000 | - |
| `I2C0_COMD7` | 0x60013074 | 32 | COMMAND0=0x3FFF COMMAND0_DONE=0x80000000 | - |
| `I2C0_CTR` | 0x60013004 | 32 | ADDR_10BIT_RW_CHECK_EN=0x2000 ADDR_BROADCASTING_EN=0x4000 ARBITRATION_EN=0x200 CLK_EN=0x100 CONF_UPGATE=0x800 FSM_RST=0x400 MS_MODE=0x10 RX_FULL_ACK_LEVEL=0x8 RX_LSB_FIRST=0x80 SAMPLE_SCL_LEVEL=0x4 SCL_FORCE_OUT=0x2 SDA_FORCE_OUT=0x1 SLV_TX... | - |
| `I2C0_DATA` | 0x6001301C | 32 | FIFO_RDATA=0xFF | - |
| `I2C0_DATE` | 0x600130F8 | 32 | DATE=0xFFFFFFFF | - |
| `I2C0_FIFO_CONF` | 0x60013018 | 32 | FIFO_ADDR_CFG_EN=0x800 FIFO_PRT_EN=0x4000 NONFIFO_EN=0x400 RXFIFO_WM_THRHD=0x1F RX_FIFO_RST=0x1000 TXFIFO_WM_THRHD=0x3E0 TX_FIFO_RST=0x2000 | - |
| `I2C0_FIFO_ST` | 0x60013014 | 32 | RXFIFO_RADDR=0x1F RXFIFO_WADDR=0x3E0 SLAVE_RW_POINT=0x3FC00000 TXFIFO_RADDR=0x7C00 TXFIFO_WADDR=0xF8000 | - |
| `I2C0_FILTER_CFG` | 0x60013050 | 32 | SCL_FILTER_EN=0x100 SCL_FILTER_THRES=0xF SDA_FILTER_EN=0x200 SDA_FILTER_THRES=0xF0 | - |
| `I2C0_INT_CLR` | 0x60013024 | 32 | ARBITRATION_LOST_INT_CLR=0x20 BYTE_TRANS_DONE_INT_CLR=0x10 DET_START_INT_CLR=0x8000 END_DETECT_INT_CLR=0x8 GENERAL_CALL_INT_CLR=0x20000 MST_TXFIFO_UDF_INT_CLR=0x40 NACK_INT_CLR=0x400 RXFIFO_OVF_INT_CLR=0x4 RXFIFO_UDF_INT_CLR=0x1000 RXFIFO_W... | - |
| `I2C0_INT_ENA` | 0x60013028 | 32 | ARBITRATION_LOST_INT_ENA=0x20 BYTE_TRANS_DONE_INT_ENA=0x10 DET_START_INT_ENA=0x8000 END_DETECT_INT_ENA=0x8 GENERAL_CALL_INT_ENA=0x20000 MST_TXFIFO_UDF_INT_ENA=0x40 NACK_INT_ENA=0x400 RXFIFO_OVF_INT_ENA=0x4 RXFIFO_UDF_INT_ENA=0x1000 RXFIFO_W... | - |
| `I2C0_INT_RAW` | 0x60013020 | 32 | ARBITRATION_LOST_INT_RAW=0x20 BYTE_TRANS_DONE_INT_RAW=0x10 DET_START_INT_RAW=0x8000 END_DETECT_INT_RAW=0x8 GENERAL_CALL_INT_RAW=0x20000 MST_TXFIFO_UDF_INT_RAW=0x40 NACK_INT_RAW=0x400 RXFIFO_OVF_INT_RAW=0x4 RXFIFO_UDF_INT_RAW=0x1000 RXFIFO_W... | - |
| `I2C0_INT_STATUS` | 0x6001302C | 32 | ARBITRATION_LOST_INT_ST=0x20 BYTE_TRANS_DONE_INT_ST=0x10 DET_START_INT_ST=0x8000 END_DETECT_INT_ST=0x8 GENERAL_CALL_INT_ST=0x20000 MST_TXFIFO_UDF_INT_ST=0x40 NACK_INT_ST=0x400 RXFIFO_OVF_INT_ST=0x4 RXFIFO_UDF_INT_ST=0x1000 RXFIFO_WM_INT_ST=... | - |
| `I2C0_RXFIFO_START_ADDR` | 0x60013180 | 32 | RXFIFO_START_ADDR=0xFFFFFFFF | - |
| `I2C0_SCL_HIGH_PERIOD` | 0x60013038 | 32 | SCL_HIGH_PERIOD=0x1FF SCL_WAIT_HIGH_PERIOD=0xFE00 | - |
| `I2C0_SCL_LOW_PERIOD` | 0x60013000 | 32 | SCL_LOW_PERIOD=0x1FF | - |
| `I2C0_SCL_MAIN_ST_TIME_OUT` | 0x6001307C | 32 | SCL_MAIN_ST_TO_I2C=0x1F | - |
| `I2C0_SCL_RSTART_SETUP` | 0x60013044 | 32 | TIME=0x1FF | - |
| `I2C0_SCL_SP_CONF` | 0x60013080 | 32 | SCL_PD_EN=0x40 SCL_RST_SLV_EN=0x1 SCL_RST_SLV_NUM=0x3E SDA_PD_EN=0x80 | - |
| `I2C0_SCL_START_HOLD` | 0x60013040 | 32 | TIME=0x1FF | - |
| `I2C0_SCL_STOP_HOLD` | 0x60013048 | 32 | TIME=0x1FF | - |
| `I2C0_SCL_STOP_SETUP` | 0x6001304C | 32 | TIME=0x1FF | - |
| `I2C0_SCL_STRETCH_CONF` | 0x60013084 | 32 | SLAVE_BYTE_ACK_CTL_EN=0x1000 SLAVE_BYTE_ACK_LVL=0x2000 SLAVE_SCL_STRETCH_CLR=0x800 SLAVE_SCL_STRETCH_EN=0x400 STRETCH_PROTECT_NUM=0x3FF | - |
| `I2C0_SCL_ST_TIME_OUT` | 0x60013078 | 32 | SCL_ST_TO_I2C=0x1F | - |
| `I2C0_SDA_HOLD` | 0x60013030 | 32 | TIME=0x1FF | - |
| `I2C0_SDA_SAMPLE` | 0x60013034 | 32 | TIME=0x1FF | - |
| `I2C0_SLAVE_ADDR` | 0x60013010 | 32 | ADDR_10BIT_EN=0x80000000 SLAVE_ADDR=0x7FFF | - |
| `I2C0_SR` | 0x60013008 | 32 | ARB_LOST=0x8 BUS_BUSY=0x10 RESP_REC=0x1 RXFIFO_CNT=0x3F00 SCL_MAIN_STATE_LAST=0x7000000 SCL_STATE_LAST=0x70000000 SLAVE_ADDRESSED=0x20 SLAVE_RW=0x2 STRETCH_CAUSE=0xC000 TXFIFO_CNT=0xFC0000 | - |
| `I2C0_TO` | 0x6001300C | 32 | TIME_OUT_EN=0x20 TIME_OUT_VALUE=0x1F | - |
| `I2C0_TXFIFO_START_ADDR` | 0x60013100 | 32 | TXFIFO_START_ADDR=0xFFFFFFFF | - |
| `I2S0_CONF_SIGLE_DATA` | 0x6002D068 | 32 | SINGLE_DATA=0xFFFFFFFF | - |
| `I2S0_DATE` | 0x6002D080 | 32 | DATE=0xFFFFFFF | - |
| `I2S0_ETM_CONF` | 0x6002D070 | 32 | ETM_RX_RECEIVE_WORD_NUM=0xFFC00 ETM_TX_SEND_WORD_NUM=0x3FF | - |
| `I2S0_INT_CLR` | 0x6002D018 | 32 | RX_DONE_INT_CLR=0x1 RX_HUNG_INT_CLR=0x4 TX_DONE_INT_CLR=0x2 TX_HUNG_INT_CLR=0x8 | - |
| `I2S0_INT_ENA` | 0x6002D014 | 32 | RX_DONE_INT_ENA=0x1 RX_HUNG_INT_ENA=0x4 TX_DONE_INT_ENA=0x2 TX_HUNG_INT_ENA=0x8 | - |
| `I2S0_INT_RAW` | 0x6002D00C | 32 | RX_DONE_INT_RAW=0x1 RX_HUNG_INT_RAW=0x4 TX_DONE_INT_RAW=0x2 TX_HUNG_INT_RAW=0x8 | - |
| `I2S0_INT_ST` | 0x6002D010 | 32 | RX_DONE_INT_ST=0x1 RX_HUNG_INT_ST=0x4 TX_DONE_INT_ST=0x2 TX_HUNG_INT_ST=0x8 | - |
| `I2S0_LC_HUNG_CONF` | 0x6002D060 | 32 | LC_FIFO_TIMEOUT=0xFF LC_FIFO_TIMEOUT_ENA=0x800 LC_FIFO_TIMEOUT_SHIFT=0x700 | - |
| `I2S0_RXEOF_NUM` | 0x6002D064 | 32 | RX_EOF_NUM=0xFFF | - |
| `I2S0_RX_CLKM_CONF` | 0x6002D030 | 32 | MCLK_SEL=0x20000000 RX_CLKM_DIV_NUM=0xFF RX_CLK_ACTIVE=0x4000000 RX_CLK_SEL=0x18000000 | - |
| `I2S0_RX_CLKM_DIV_CONF` | 0x6002D038 | 32 | RX_CLKM_DIV_X=0x7FC0000 RX_CLKM_DIV_Y=0x3FE00 RX_CLKM_DIV_YN1=0x8000000 RX_CLKM_DIV_Z=0x1FF | - |
| `I2S0_RX_CONF` | 0x6002D020 | 32 | RX_24_FILL_EN=0x10000 RX_BCK_DIV_NUM=0x7E00000 RX_BIG_ENDIAN=0x80 RX_BIT_ORDER=0x40000 RX_FIFO_RESET=0x2 RX_LEFT_ALIGN=0x8000 RX_MONO=0x40 RX_MONO_FST_VLD=0x200 RX_MSB_SHIFT=0x2000 RX_PCM_BYPASS=0x1000 RX_PCM_CONF=0xC00 RX_PDM_EN=0x100000 R... | - |
| `I2S0_RX_CONF1` | 0x6002D028 | 32 | RX_BITS_MOD=0x7C000 RX_HALF_SAMPLE_BITS=0x7F80000 RX_TDM_CHAN_BITS=0xF8000000 RX_TDM_WS_WIDTH=0x1FF | - |
| `I2S0_RX_TDM_CTRL` | 0x6002D050 | 32 | RX_TDM_CHAN10_EN=0x400 RX_TDM_CHAN11_EN=0x800 RX_TDM_CHAN12_EN=0x1000 RX_TDM_CHAN13_EN=0x2000 RX_TDM_CHAN14_EN=0x4000 RX_TDM_CHAN15_EN=0x8000 RX_TDM_CHAN8_EN=0x100 RX_TDM_CHAN9_EN=0x200 RX_TDM_PDM_CHAN0_EN=0x1 RX_TDM_PDM_CHAN1_EN=0x2 RX_TDM... | - |
| `I2S0_RX_TIMING` | 0x6002D058 | 32 | RX_BCK_IN_DM=0x30000000 RX_BCK_OUT_DM=0x300000 RX_SD_IN_DM=0x3 RX_WS_IN_DM=0x3000000 RX_WS_OUT_DM=0x30000 | - |
| `I2S0_STATE` | 0x6002D06C | 32 | TX_IDLE=0x1 | - |
| `I2S0_TX_CLKM_CONF` | 0x6002D034 | 32 | CLK_EN=0x20000000 TX_CLKM_DIV_NUM=0xFF TX_CLK_ACTIVE=0x4000000 TX_CLK_SEL=0x18000000 | - |
| `I2S0_TX_CLKM_DIV_CONF` | 0x6002D03C | 32 | TX_CLKM_DIV_X=0x7FC0000 TX_CLKM_DIV_Y=0x3FE00 TX_CLKM_DIV_YN1=0x8000000 TX_CLKM_DIV_Z=0x1FF | - |
| `I2S0_TX_CONF` | 0x6002D024 | 32 | SIG_LOOPBACK=0x40000000 TX_24_FILL_EN=0x10000 TX_BCK_DIV_NUM=0x7E00000 TX_BCK_NO_DLY=0x4000 TX_BIG_ENDIAN=0x80 TX_BIT_ORDER=0x40000 TX_CHAN_EQUAL=0x20 TX_CHAN_MOD=0x38000000 TX_FIFO_RESET=0x2 TX_LEFT_ALIGN=0x8000 TX_MONO=0x40 TX_MONO_FST_VL... | - |
| `I2S0_TX_CONF1` | 0x6002D02C | 32 | TX_BITS_MOD=0x7C000 TX_HALF_SAMPLE_BITS=0x7F80000 TX_TDM_CHAN_BITS=0xF8000000 TX_TDM_WS_WIDTH=0x1FF | - |
| `I2S0_TX_PCM2PDM_CONF` | 0x6002D040 | 32 | PCM2PDM_CONV_EN=0x2000000 TX_PDM_DAC_2OUT_EN=0x800000 TX_PDM_DAC_MODE_EN=0x1000000 TX_PDM_HP_BYPASS=0x1 TX_PDM_HP_IN_SHIFT=0x6000 TX_PDM_LP_IN_SHIFT=0x18000 TX_PDM_PRESCALE=0x1FE0 TX_PDM_SIGMADELTA_DITHER2=0x200000 TX_PDM_SIGMADELTA_DITHER=... | - |
| `I2S0_TX_PCM2PDM_CONF1` | 0x6002D044 | 32 | TX_IIR_HP_MULT12_0=0x3800000 TX_IIR_HP_MULT12_5=0x700000 TX_PDM_FP=0x3FF TX_PDM_FS=0xFFC00 | - |
| `I2S0_TX_TDM_CTRL` | 0x6002D054 | 32 | TX_TDM_CHAN0_EN=0x1 TX_TDM_CHAN10_EN=0x400 TX_TDM_CHAN11_EN=0x800 TX_TDM_CHAN12_EN=0x1000 TX_TDM_CHAN13_EN=0x2000 TX_TDM_CHAN14_EN=0x4000 TX_TDM_CHAN15_EN=0x8000 TX_TDM_CHAN1_EN=0x2 TX_TDM_CHAN2_EN=0x4 TX_TDM_CHAN3_EN=0x8 TX_TDM_CHAN4_EN=0x... | - |
| `I2S0_TX_TIMING` | 0x6002D05C | 32 | TX_BCK_IN_DM=0x30000000 TX_BCK_OUT_DM=0x300000 TX_SD1_OUT_DM=0x30 TX_SD_OUT_DM=0x3 TX_WS_IN_DM=0x3000000 TX_WS_OUT_DM=0x30000 | - |
| `IO_MUX_DATE` | 0x600090FC | 32 | REG_DATE=0xFFFFFFF | - |
| `IO_MUX_GPIO0` | 0x60009004 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO1` | 0x60009008 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO10` | 0x6000902C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO11` | 0x60009030 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO12` | 0x60009034 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO13` | 0x60009038 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO14` | 0x6000903C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO15` | 0x60009040 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO16` | 0x60009044 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO17` | 0x60009048 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO18` | 0x6000904C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO19` | 0x60009050 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO2` | 0x6000900C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO20` | 0x60009054 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO21` | 0x60009058 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO22` | 0x6000905C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO23` | 0x60009060 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO24` | 0x60009064 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO25` | 0x60009068 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO26` | 0x6000906C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO27` | 0x60009070 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO3` | 0x60009010 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO4` | 0x60009014 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO5` | 0x60009018 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO6` | 0x6000901C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO7` | 0x60009020 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO8` | 0x60009024 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO9` | 0x60009028 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 HYS_EN=0x10000 HYS_SEL=0x20000 MCU_DRV=0x60 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_MODEM_DIAG_EN` | 0x600090BC | 32 | MODEM_DIAG_EN=0xFFFFFFFF | - |
| `IO_MUX_PIN_CTRL` | 0x60009000 | 32 | CLK_OUT1=0x1F CLK_OUT2=0x3E0 CLK_OUT3=0x7C00 | - |
| `LEDC_CH0_CONF0` | 0x60019000 | 32 | IDLE_LV_CH=0x8 OVF_CNT_EN_CH=0x8000 OVF_CNT_RESET_CH=0x10000 OVF_NUM_CH=0x7FE0 PARA_UP_CH=0x10 SIG_OUT_EN_CH=0x4 TIMER_SEL_CH=0x3 | - |
| `LEDC_CH0_CONF1` | 0x6001900C | 32 | DUTY_START_CH=0x80000000 | - |
| `LEDC_CH0_DUTY` | 0x60019008 | 32 | DUTY_CH=0x1FFFFFF | - |
| `LEDC_CH0_DUTY_R` | 0x60019010 | 32 | DUTY_CH_R=0x1FFFFFF | - |
| `LEDC_CH0_GAMMA_CONF` | 0x60019180 | 32 | CH_GAMMA_ENTRY_NUM=0x1F CH_GAMMA_PAUSE=0x20 CH_GAMMA_RESUME=0x40 | - |
| `LEDC_CH0_GAMMA_RD_ADDR` | 0x60019108 | 32 | CH_GAMMA_RD_ADDR=0xF | - |
| `LEDC_CH0_GAMMA_RD_DATA` | 0x6001910C | 32 | CH_GAMMA_RD_DATA=0x7FFFFFFF | - |
| `LEDC_CH0_GAMMA_WR` | 0x60019100 | 32 | CH_GAMMA_DUTY_CYCLE=0x7FE CH_GAMMA_DUTY_INC=0x1 CH_GAMMA_DUTY_NUM=0x7FE00000 CH_GAMMA_SCALE=0x1FF800 | - |
| `LEDC_CH0_GAMMA_WR_ADDR` | 0x60019104 | 32 | CH_GAMMA_WR_ADDR=0xF | - |
| `LEDC_CH0_HPOINT` | 0x60019004 | 32 | HPOINT_CH=0xFFFFF | - |
| `LEDC_CH1_CONF0` | 0x60019014 | 32 | IDLE_LV_CH=0x8 OVF_CNT_EN_CH=0x8000 OVF_CNT_RESET_CH=0x10000 OVF_NUM_CH=0x7FE0 PARA_UP_CH=0x10 SIG_OUT_EN_CH=0x4 TIMER_SEL_CH=0x3 | - |
| `LEDC_CH1_CONF1` | 0x60019020 | 32 | DUTY_START_CH=0x80000000 | - |
| `LEDC_CH1_DUTY` | 0x6001901C | 32 | DUTY_CH=0x1FFFFFF | - |
| `LEDC_CH1_DUTY_R` | 0x60019024 | 32 | DUTY_CH_R=0x1FFFFFF | - |
| `LEDC_CH1_GAMMA_CONF` | 0x60019184 | 32 | CH_GAMMA_ENTRY_NUM=0x1F CH_GAMMA_PAUSE=0x20 CH_GAMMA_RESUME=0x40 | - |
| `LEDC_CH1_GAMMA_RD_ADDR` | 0x60019118 | 32 | CH_GAMMA_RD_ADDR=0xF | - |
| `LEDC_CH1_GAMMA_RD_DATA` | 0x6001911C | 32 | CH_GAMMA_RD_DATA=0x7FFFFFFF | - |
| `LEDC_CH1_GAMMA_WR` | 0x60019110 | 32 | CH_GAMMA_DUTY_CYCLE=0x7FE CH_GAMMA_DUTY_INC=0x1 CH_GAMMA_DUTY_NUM=0x7FE00000 CH_GAMMA_SCALE=0x1FF800 | - |
| `LEDC_CH1_GAMMA_WR_ADDR` | 0x60019114 | 32 | CH_GAMMA_WR_ADDR=0xF | - |
| `LEDC_CH1_HPOINT` | 0x60019018 | 32 | HPOINT_CH=0xFFFFF | - |
| `LEDC_CH2_CONF0` | 0x60019028 | 32 | IDLE_LV_CH=0x8 OVF_CNT_EN_CH=0x8000 OVF_CNT_RESET_CH=0x10000 OVF_NUM_CH=0x7FE0 PARA_UP_CH=0x10 SIG_OUT_EN_CH=0x4 TIMER_SEL_CH=0x3 | - |
| `LEDC_CH2_CONF1` | 0x60019034 | 32 | DUTY_START_CH=0x80000000 | - |
| `LEDC_CH2_DUTY` | 0x60019030 | 32 | DUTY_CH=0x1FFFFFF | - |
| `LEDC_CH2_DUTY_R` | 0x60019038 | 32 | DUTY_CH_R=0x1FFFFFF | - |
| `LEDC_CH2_GAMMA_CONF` | 0x60019188 | 32 | CH_GAMMA_ENTRY_NUM=0x1F CH_GAMMA_PAUSE=0x20 CH_GAMMA_RESUME=0x40 | - |
| `LEDC_CH2_GAMMA_RD_ADDR` | 0x60019128 | 32 | CH_GAMMA_RD_ADDR=0xF | - |
| `LEDC_CH2_GAMMA_RD_DATA` | 0x6001912C | 32 | CH_GAMMA_RD_DATA=0x7FFFFFFF | - |
| `LEDC_CH2_GAMMA_WR` | 0x60019120 | 32 | CH_GAMMA_DUTY_CYCLE=0x7FE CH_GAMMA_DUTY_INC=0x1 CH_GAMMA_DUTY_NUM=0x7FE00000 CH_GAMMA_SCALE=0x1FF800 | - |
| `LEDC_CH2_GAMMA_WR_ADDR` | 0x60019124 | 32 | CH_GAMMA_WR_ADDR=0xF | - |
| `LEDC_CH2_HPOINT` | 0x6001902C | 32 | HPOINT_CH=0xFFFFF | - |
| `LEDC_CH3_CONF0` | 0x6001903C | 32 | IDLE_LV_CH=0x8 OVF_CNT_EN_CH=0x8000 OVF_CNT_RESET_CH=0x10000 OVF_NUM_CH=0x7FE0 PARA_UP_CH=0x10 SIG_OUT_EN_CH=0x4 TIMER_SEL_CH=0x3 | - |
| `LEDC_CH3_CONF1` | 0x60019048 | 32 | DUTY_START_CH=0x80000000 | - |
| `LEDC_CH3_DUTY` | 0x60019044 | 32 | DUTY_CH=0x1FFFFFF | - |
| `LEDC_CH3_DUTY_R` | 0x6001904C | 32 | DUTY_CH_R=0x1FFFFFF | - |
| `LEDC_CH3_GAMMA_CONF` | 0x6001918C | 32 | CH_GAMMA_ENTRY_NUM=0x1F CH_GAMMA_PAUSE=0x20 CH_GAMMA_RESUME=0x40 | - |
| `LEDC_CH3_GAMMA_RD_ADDR` | 0x60019138 | 32 | CH_GAMMA_RD_ADDR=0xF | - |
| `LEDC_CH3_GAMMA_RD_DATA` | 0x6001913C | 32 | CH_GAMMA_RD_DATA=0x7FFFFFFF | - |
| `LEDC_CH3_GAMMA_WR` | 0x60019130 | 32 | CH_GAMMA_DUTY_CYCLE=0x7FE CH_GAMMA_DUTY_INC=0x1 CH_GAMMA_DUTY_NUM=0x7FE00000 CH_GAMMA_SCALE=0x1FF800 | - |
| `LEDC_CH3_GAMMA_WR_ADDR` | 0x60019134 | 32 | CH_GAMMA_WR_ADDR=0xF | - |
| `LEDC_CH3_HPOINT` | 0x60019040 | 32 | HPOINT_CH=0xFFFFF | - |
| `LEDC_CH4_CONF0` | 0x60019050 | 32 | IDLE_LV_CH=0x8 OVF_CNT_EN_CH=0x8000 OVF_CNT_RESET_CH=0x10000 OVF_NUM_CH=0x7FE0 PARA_UP_CH=0x10 SIG_OUT_EN_CH=0x4 TIMER_SEL_CH=0x3 | - |
| `LEDC_CH4_CONF1` | 0x6001905C | 32 | DUTY_START_CH=0x80000000 | - |
| `LEDC_CH4_DUTY` | 0x60019058 | 32 | DUTY_CH=0x1FFFFFF | - |
| `LEDC_CH4_DUTY_R` | 0x60019060 | 32 | DUTY_CH_R=0x1FFFFFF | - |
| `LEDC_CH4_GAMMA_CONF` | 0x60019190 | 32 | CH_GAMMA_ENTRY_NUM=0x1F CH_GAMMA_PAUSE=0x20 CH_GAMMA_RESUME=0x40 | - |
| `LEDC_CH4_GAMMA_RD_ADDR` | 0x60019148 | 32 | CH_GAMMA_RD_ADDR=0xF | - |
| `LEDC_CH4_GAMMA_RD_DATA` | 0x6001914C | 32 | CH_GAMMA_RD_DATA=0x7FFFFFFF | - |
| `LEDC_CH4_GAMMA_WR` | 0x60019140 | 32 | CH_GAMMA_DUTY_CYCLE=0x7FE CH_GAMMA_DUTY_INC=0x1 CH_GAMMA_DUTY_NUM=0x7FE00000 CH_GAMMA_SCALE=0x1FF800 | - |
| `LEDC_CH4_GAMMA_WR_ADDR` | 0x60019144 | 32 | CH_GAMMA_WR_ADDR=0xF | - |
| `LEDC_CH4_HPOINT` | 0x60019054 | 32 | HPOINT_CH=0xFFFFF | - |
| `LEDC_CH5_CONF0` | 0x60019064 | 32 | IDLE_LV_CH=0x8 OVF_CNT_EN_CH=0x8000 OVF_CNT_RESET_CH=0x10000 OVF_NUM_CH=0x7FE0 PARA_UP_CH=0x10 SIG_OUT_EN_CH=0x4 TIMER_SEL_CH=0x3 | - |
| `LEDC_CH5_CONF1` | 0x60019070 | 32 | DUTY_START_CH=0x80000000 | - |
| `LEDC_CH5_DUTY` | 0x6001906C | 32 | DUTY_CH=0x1FFFFFF | - |
| `LEDC_CH5_DUTY_R` | 0x60019074 | 32 | DUTY_CH_R=0x1FFFFFF | - |
| `LEDC_CH5_GAMMA_CONF` | 0x60019194 | 32 | CH_GAMMA_ENTRY_NUM=0x1F CH_GAMMA_PAUSE=0x20 CH_GAMMA_RESUME=0x40 | - |
| `LEDC_CH5_GAMMA_RD_ADDR` | 0x60019158 | 32 | CH_GAMMA_RD_ADDR=0xF | - |
| `LEDC_CH5_GAMMA_RD_DATA` | 0x6001915C | 32 | CH_GAMMA_RD_DATA=0x7FFFFFFF | - |
| `LEDC_CH5_GAMMA_WR` | 0x60019150 | 32 | CH_GAMMA_DUTY_CYCLE=0x7FE CH_GAMMA_DUTY_INC=0x1 CH_GAMMA_DUTY_NUM=0x7FE00000 CH_GAMMA_SCALE=0x1FF800 | - |
| `LEDC_CH5_GAMMA_WR_ADDR` | 0x60019154 | 32 | CH_GAMMA_WR_ADDR=0xF | - |
| `LEDC_CH5_HPOINT` | 0x60019068 | 32 | HPOINT_CH=0xFFFFF | - |
| `LEDC_CONF` | 0x600191F0 | 32 | APB_CLK_SEL=0x3 CLK_EN=0x80000000 GAMMA_RAM_CLK_EN_CH0=0x4 GAMMA_RAM_CLK_EN_CH1=0x8 GAMMA_RAM_CLK_EN_CH2=0x10 GAMMA_RAM_CLK_EN_CH3=0x20 GAMMA_RAM_CLK_EN_CH4=0x40 GAMMA_RAM_CLK_EN_CH5=0x80 | - |
| `LEDC_DATE` | 0x600191FC | 32 | LEDC_DATE=0xFFFFFFF | - |
| `LEDC_EVT_TASK_EN0` | 0x600191A0 | 32 | EVT_DUTY_CHNG_END_CH0_EN=0x1 EVT_DUTY_CHNG_END_CH1_EN=0x2 EVT_DUTY_CHNG_END_CH2_EN=0x4 EVT_DUTY_CHNG_END_CH3_EN=0x8 EVT_DUTY_CHNG_END_CH4_EN=0x10 EVT_DUTY_CHNG_END_CH5_EN=0x20 EVT_OVF_CNT_PLS_CH0_EN=0x100 EVT_OVF_CNT_PLS_CH1_EN=0x200 EVT_OV... | - |
| `LEDC_EVT_TASK_EN1` | 0x600191A4 | 32 | TASK_OVF_CNT_RST_CH0_EN=0x10000 TASK_OVF_CNT_RST_CH1_EN=0x20000 TASK_OVF_CNT_RST_CH2_EN=0x40000 TASK_OVF_CNT_RST_CH3_EN=0x80000 TASK_OVF_CNT_RST_CH4_EN=0x100000 TASK_OVF_CNT_RST_CH5_EN=0x200000 TASK_SIG_OUT_DIS_CH0_EN=0x100 TASK_SIG_OUT_DIS... | - |
| `LEDC_EVT_TASK_EN2` | 0x600191A8 | 32 | TASK_GAMMA_PAUSE_CH0_EN=0x100 TASK_GAMMA_PAUSE_CH1_EN=0x200 TASK_GAMMA_PAUSE_CH2_EN=0x400 TASK_GAMMA_PAUSE_CH3_EN=0x800 TASK_GAMMA_PAUSE_CH4_EN=0x1000 TASK_GAMMA_PAUSE_CH5_EN=0x2000 TASK_GAMMA_RESTART_CH0_EN=0x1 TASK_GAMMA_RESTART_CH1_EN=0x... | - |
| `LEDC_INT_CLR` | 0x600190CC | 32 | DUTY_CHNG_END_CH0_INT_CLR=0x10 DUTY_CHNG_END_CH1_INT_CLR=0x20 DUTY_CHNG_END_CH2_INT_CLR=0x40 DUTY_CHNG_END_CH3_INT_CLR=0x80 DUTY_CHNG_END_CH4_INT_CLR=0x100 DUTY_CHNG_END_CH5_INT_CLR=0x200 OVF_CNT_CH0_INT_CLR=0x1000 OVF_CNT_CH1_INT_CLR=0x200... | - |
| `LEDC_INT_ENA` | 0x600190C8 | 32 | DUTY_CHNG_END_CH0_INT_ENA=0x10 DUTY_CHNG_END_CH1_INT_ENA=0x20 DUTY_CHNG_END_CH2_INT_ENA=0x40 DUTY_CHNG_END_CH3_INT_ENA=0x80 DUTY_CHNG_END_CH4_INT_ENA=0x100 DUTY_CHNG_END_CH5_INT_ENA=0x200 OVF_CNT_CH0_INT_ENA=0x1000 OVF_CNT_CH1_INT_ENA=0x200... | - |
| `LEDC_INT_RAW` | 0x600190C0 | 32 | DUTY_CHNG_END_CH0_INT_RAW=0x10 DUTY_CHNG_END_CH1_INT_RAW=0x20 DUTY_CHNG_END_CH2_INT_RAW=0x40 DUTY_CHNG_END_CH3_INT_RAW=0x80 DUTY_CHNG_END_CH4_INT_RAW=0x100 DUTY_CHNG_END_CH5_INT_RAW=0x200 OVF_CNT_CH0_INT_RAW=0x1000 OVF_CNT_CH1_INT_RAW=0x200... | - |
| `LEDC_INT_ST` | 0x600190C4 | 32 | DUTY_CHNG_END_CH0_INT_ST=0x10 DUTY_CHNG_END_CH1_INT_ST=0x20 DUTY_CHNG_END_CH2_INT_ST=0x40 DUTY_CHNG_END_CH3_INT_ST=0x80 DUTY_CHNG_END_CH4_INT_ST=0x100 DUTY_CHNG_END_CH5_INT_ST=0x200 OVF_CNT_CH0_INT_ST=0x1000 OVF_CNT_CH1_INT_ST=0x2000 OVF_CN... | - |
| `LEDC_TIMER0_CMP` | 0x600191B0 | 32 | TIMER_CMP=0xFFFFF | - |
| `LEDC_TIMER0_CNT_CAP` | 0x600191C0 | 32 | TIMER_CNT_CAP=0xFFFFF | - |
| `LEDC_TIMER0_CONF` | 0x600190A0 | 32 | CLK_DIV_TIMER=0x7FFFE0 TICK_SEL_TIMER=0x2000000 TIMER_DUTY_RES=0x1F TIMER_PARA_UP=0x4000000 TIMER_PAUSE=0x800000 TIMER_RST=0x1000000 | - |
| `LEDC_TIMER0_VALUE` | 0x600190A4 | 32 | TIMER_CNT=0xFFFFF | - |
| `LEDC_TIMER1_CMP` | 0x600191B4 | 32 | TIMER_CMP=0xFFFFF | - |
| `LEDC_TIMER1_CNT_CAP` | 0x600191C4 | 32 | TIMER_CNT_CAP=0xFFFFF | - |
| `LEDC_TIMER1_CONF` | 0x600190A8 | 32 | CLK_DIV_TIMER=0x7FFFE0 TICK_SEL_TIMER=0x2000000 TIMER_DUTY_RES=0x1F TIMER_PARA_UP=0x4000000 TIMER_PAUSE=0x800000 TIMER_RST=0x1000000 | - |
| `LEDC_TIMER1_VALUE` | 0x600190AC | 32 | TIMER_CNT=0xFFFFF | - |
| `LEDC_TIMER2_CMP` | 0x600191B8 | 32 | TIMER_CMP=0xFFFFF | - |
| `LEDC_TIMER2_CNT_CAP` | 0x600191C8 | 32 | TIMER_CNT_CAP=0xFFFFF | - |
| `LEDC_TIMER2_CONF` | 0x600190B0 | 32 | CLK_DIV_TIMER=0x7FFFE0 TICK_SEL_TIMER=0x2000000 TIMER_DUTY_RES=0x1F TIMER_PARA_UP=0x4000000 TIMER_PAUSE=0x800000 TIMER_RST=0x1000000 | - |
| `LEDC_TIMER2_VALUE` | 0x600190B4 | 32 | TIMER_CNT=0xFFFFF | - |
| `LEDC_TIMER3_CMP` | 0x600191BC | 32 | TIMER_CMP=0xFFFFF | - |
| `LEDC_TIMER3_CNT_CAP` | 0x600191CC | 32 | TIMER_CNT_CAP=0xFFFFF | - |
| `LEDC_TIMER3_CONF` | 0x600190B8 | 32 | CLK_DIV_TIMER=0x7FFFE0 TICK_SEL_TIMER=0x2000000 TIMER_DUTY_RES=0x1F TIMER_PARA_UP=0x4000000 TIMER_PAUSE=0x800000 TIMER_RST=0x1000000 | - |
| `LEDC_TIMER3_VALUE` | 0x600190BC | 32 | TIMER_CNT=0xFFFFF | - |
| `PCNT_CTRL` | 0x60017060 | 32 | CLK_EN=0x10000 CNT_PAUSE_U0=0x2 CNT_PAUSE_U1=0x8 CNT_PAUSE_U2=0x20 CNT_PAUSE_U3=0x80 PULSE_CNT_RST_U0=0x1 PULSE_CNT_RST_U1=0x4 PULSE_CNT_RST_U2=0x10 PULSE_CNT_RST_U3=0x40 | - |
| `PCNT_DATE` | 0x600170FC | 32 | DATE=0xFFFFFFFF | - |
| `PCNT_INT_CLR` | 0x6001704C | 32 | CNT_THR_EVENT_U0_INT_CLR=0x1 CNT_THR_EVENT_U1_INT_CLR=0x2 CNT_THR_EVENT_U2_INT_CLR=0x4 CNT_THR_EVENT_U3_INT_CLR=0x8 | - |
| `PCNT_INT_ENA` | 0x60017048 | 32 | CNT_THR_EVENT_U0_INT_ENA=0x1 CNT_THR_EVENT_U1_INT_ENA=0x2 CNT_THR_EVENT_U2_INT_ENA=0x4 CNT_THR_EVENT_U3_INT_ENA=0x8 | - |
| `PCNT_INT_RAW` | 0x60017040 | 32 | CNT_THR_EVENT_U0_INT_RAW=0x1 CNT_THR_EVENT_U1_INT_RAW=0x2 CNT_THR_EVENT_U2_INT_RAW=0x4 CNT_THR_EVENT_U3_INT_RAW=0x8 | - |
| `PCNT_INT_ST` | 0x60017044 | 32 | CNT_THR_EVENT_U0_INT_ST=0x1 CNT_THR_EVENT_U1_INT_ST=0x2 CNT_THR_EVENT_U2_INT_ST=0x4 CNT_THR_EVENT_U3_INT_ST=0x8 | - |
| `PCNT_U0_CNT` | 0x60017030 | 32 | PULSE_CNT_U=0xFFFF | - |
| `PCNT_U0_CONF0` | 0x60017000 | 32 | CH0_HCTRL_MODE_U=0x300000 CH0_LCTRL_MODE_U=0xC00000 CH0_NEG_MODE_U=0x30000 CH0_POS_MODE_U=0xC0000 CH1_HCTRL_MODE_U=0x30000000 CH1_LCTRL_MODE_U=0xC0000000 CH1_NEG_MODE_U=0x3000000 CH1_POS_MODE_U=0xC000000 FILTER_EN_U=0x400 FILTER_THRES_U=0x3... | - |
| `PCNT_U0_CONF1` | 0x60017004 | 32 | CNT_THRES0_U=0xFFFF CNT_THRES1_U=0xFFFF0000 | - |
| `PCNT_U0_CONF2` | 0x60017008 | 32 | CNT_H_LIM_U=0xFFFF CNT_L_LIM_U=0xFFFF0000 | - |
| `PCNT_U0_STATUS` | 0x60017050 | 32 | CNT_THR_H_LIM_LAT_U=0x20 CNT_THR_L_LIM_LAT_U=0x10 CNT_THR_THRES0_LAT_U=0x8 CNT_THR_THRES1_LAT_U=0x4 CNT_THR_ZERO_LAT_U=0x40 CNT_THR_ZERO_MODE_U=0x3 | - |
| `PCNT_U1_CNT` | 0x60017034 | 32 | PULSE_CNT_U=0xFFFF | - |
| `PCNT_U1_CONF0` | 0x6001700C | 32 | CH0_HCTRL_MODE_U=0x300000 CH0_LCTRL_MODE_U=0xC00000 CH0_NEG_MODE_U=0x30000 CH0_POS_MODE_U=0xC0000 CH1_HCTRL_MODE_U=0x30000000 CH1_LCTRL_MODE_U=0xC0000000 CH1_NEG_MODE_U=0x3000000 CH1_POS_MODE_U=0xC000000 FILTER_EN_U=0x400 FILTER_THRES_U=0x3... | - |
| `PCNT_U1_CONF1` | 0x60017010 | 32 | CNT_THRES0_U=0xFFFF CNT_THRES1_U=0xFFFF0000 | - |
| `PCNT_U1_CONF2` | 0x60017014 | 32 | CNT_H_LIM_U=0xFFFF CNT_L_LIM_U=0xFFFF0000 | - |
| `PCNT_U1_STATUS` | 0x60017054 | 32 | CNT_THR_H_LIM_LAT_U=0x20 CNT_THR_L_LIM_LAT_U=0x10 CNT_THR_THRES0_LAT_U=0x8 CNT_THR_THRES1_LAT_U=0x4 CNT_THR_ZERO_LAT_U=0x40 CNT_THR_ZERO_MODE_U=0x3 | - |
| `PCNT_U2_CNT` | 0x60017038 | 32 | PULSE_CNT_U=0xFFFF | - |
| `PCNT_U2_CONF0` | 0x60017018 | 32 | CH0_HCTRL_MODE_U=0x300000 CH0_LCTRL_MODE_U=0xC00000 CH0_NEG_MODE_U=0x30000 CH0_POS_MODE_U=0xC0000 CH1_HCTRL_MODE_U=0x30000000 CH1_LCTRL_MODE_U=0xC0000000 CH1_NEG_MODE_U=0x3000000 CH1_POS_MODE_U=0xC000000 FILTER_EN_U=0x400 FILTER_THRES_U=0x3... | - |
| `PCNT_U2_CONF1` | 0x6001701C | 32 | CNT_THRES0_U=0xFFFF CNT_THRES1_U=0xFFFF0000 | - |
| `PCNT_U2_CONF2` | 0x60017020 | 32 | CNT_H_LIM_U=0xFFFF CNT_L_LIM_U=0xFFFF0000 | - |
| `PCNT_U2_STATUS` | 0x60017058 | 32 | CNT_THR_H_LIM_LAT_U=0x20 CNT_THR_L_LIM_LAT_U=0x10 CNT_THR_THRES0_LAT_U=0x8 CNT_THR_THRES1_LAT_U=0x4 CNT_THR_ZERO_LAT_U=0x40 CNT_THR_ZERO_MODE_U=0x3 | - |
| `PCNT_U3_CNT` | 0x6001703C | 32 | PULSE_CNT_U=0xFFFF | - |
| `PCNT_U3_CONF0` | 0x60017024 | 32 | CH0_HCTRL_MODE_U=0x300000 CH0_LCTRL_MODE_U=0xC00000 CH0_NEG_MODE_U=0x30000 CH0_POS_MODE_U=0xC0000 CH1_HCTRL_MODE_U=0x30000000 CH1_LCTRL_MODE_U=0xC0000000 CH1_NEG_MODE_U=0x3000000 CH1_POS_MODE_U=0xC000000 FILTER_EN_U=0x400 FILTER_THRES_U=0x3... | - |
| `PCNT_U3_CONF1` | 0x60017028 | 32 | CNT_THRES0_U=0xFFFF CNT_THRES1_U=0xFFFF0000 | - |
| `PCNT_U3_CONF2` | 0x6001702C | 32 | CNT_H_LIM_U=0xFFFF CNT_L_LIM_U=0xFFFF0000 | - |
| `PCNT_U3_STATUS` | 0x6001705C | 32 | CNT_THR_H_LIM_LAT_U=0x20 CNT_THR_L_LIM_LAT_U=0x10 CNT_THR_THRES0_LAT_U=0x8 CNT_THR_THRES1_LAT_U=0x4 CNT_THR_ZERO_LAT_U=0x40 CNT_THR_ZERO_MODE_U=0x3 | - |
| `RMT_CH0CARRIER_DUTY` | 0x60016048 | 32 | CARRIER_HIGH_CH0=0xFFFF0000 CARRIER_LOW_CH0=0xFFFF | - |
| `RMT_CH0_RX_CARRIER_RM` | 0x60016050 | 32 | CARRIER_HIGH_THRES_CH2=0xFFFF0000 CARRIER_LOW_THRES_CH2=0xFFFF | - |
| `RMT_CH0_RX_LIM` | 0x60016060 | 32 | RMT_RX_LIM_CH2=0x1FF | - |
| `RMT_CH0_TX_LIM` | 0x60016058 | 32 | LOOP_COUNT_RESET_CH0=0x100000 LOOP_STOP_EN_CH0=0x200000 TX_LIM_CH0=0x1FF TX_LOOP_CNT_EN_CH0=0x80000 TX_LOOP_NUM_CH0=0x7FE00 | - |
| `RMT_CH1CARRIER_DUTY` | 0x6001604C | 32 | CARRIER_HIGH_CH0=0xFFFF0000 CARRIER_LOW_CH0=0xFFFF | - |
| `RMT_CH1_RX_CARRIER_RM` | 0x60016054 | 32 | CARRIER_HIGH_THRES_CH2=0xFFFF0000 CARRIER_LOW_THRES_CH2=0xFFFF | - |
| `RMT_CH1_RX_LIM` | 0x60016064 | 32 | RMT_RX_LIM_CH2=0x1FF | - |
| `RMT_CH1_TX_LIM` | 0x6001605C | 32 | LOOP_COUNT_RESET_CH0=0x100000 LOOP_STOP_EN_CH0=0x200000 TX_LIM_CH0=0x1FF TX_LOOP_CNT_EN_CH0=0x80000 TX_LOOP_NUM_CH0=0x7FE00 | - |
| `RMT_DATE` | 0x600160CC | 32 | RMT_DATE=0xFFFFFFF | - |
| `RMT_INT_CLR` | 0x60016044 | 32 | CH0_TX_END_INT_CLR=0x1 CH0_TX_LOOP_INT_CLR=0x1000 CH0_TX_THR_EVENT_INT_CLR=0x100 CH1_TX_END_INT_CLR=0x2 CH1_TX_LOOP_INT_CLR=0x2000 CH1_TX_THR_EVENT_INT_CLR=0x200 CH2_RX_END_INT_CLR=0x4 CH2_RX_THR_EVENT_INT_CLR=0x400 CH3_RX_END_INT_CLR=0x8 C... | - |
| `RMT_INT_ENA` | 0x60016040 | 32 | CH0_ERR_INT_ENA=0x10 CH0_TX_END_INT_ENA=0x1 CH0_TX_LOOP_INT_ENA=0x1000 CH0_TX_THR_EVENT_INT_ENA=0x100 CH1_ERR_INT_ENA=0x20 CH1_TX_END_INT_ENA=0x2 CH1_TX_LOOP_INT_ENA=0x2000 CH1_TX_THR_EVENT_INT_ENA=0x200 CH2_ERR_INT_ENA=0x40 CH2_RX_END_INT_... | - |
| `RMT_INT_RAW` | 0x60016038 | 32 | CH0_TX_END_INT_RAW=0x1 CH0_TX_LOOP_INT_RAW=0x1000 CH0_TX_THR_EVENT_INT_RAW=0x100 CH1_TX_END_INT_RAW=0x2 CH1_TX_LOOP_INT_RAW=0x2000 CH1_TX_THR_EVENT_INT_RAW=0x200 CH2_RX_END_INT_RAW=0x4 CH2_RX_THR_EVENT_INT_RAW=0x400 CH3_RX_END_INT_RAW=0x8 C... | - |
| `RMT_INT_ST` | 0x6001603C | 32 | CH0_TX_END_INT_ST=0x1 CH0_TX_LOOP_INT_ST=0x1000 CH0_TX_THR_EVENT_INT_ST=0x100 CH1_TX_END_INT_ST=0x2 CH1_TX_LOOP_INT_ST=0x2000 CH1_TX_THR_EVENT_INT_ST=0x200 CH2_RX_END_INT_ST=0x4 CH2_RX_THR_EVENT_INT_ST=0x400 CH3_RX_END_INT_ST=0x8 CH3_RX_THR... | - |
| `RMT_REF_CNT_RST` | 0x60016070 | 32 | RX_REF_CNT_RST_CH2=0x4 RX_REF_CNT_RST_CH3=0x8 TX_REF_CNT_RST_CH0=0x1 TX_REF_CNT_RST_CH1=0x2 | - |
| `RMT_RX_CH0CONF0` | 0x60016018 | 32 | CARRIER_EN_CH2=0x10000000 CARRIER_OUT_LV_CH2=0x20000000 DIV_CNT_CH2=0xFF IDLE_THRES_CH2=0x7FFF00 MEM_SIZE_CH2=0x3800000 | - |
| `RMT_RX_CH0CONF1` | 0x6001601C | 32 | AFIFO_RST_CH2=0x4000 APB_MEM_RST_CH2=0x4 CONF_UPDATE_CH2=0x8000 MEM_OWNER_CH2=0x8 MEM_RX_WRAP_EN_CH2=0x2000 MEM_WR_RST_CH2=0x2 RX_EN_CH2=0x1 RX_FILTER_EN_CH2=0x10 RX_FILTER_THRES_CH2=0x1FE0 | - |
| `RMT_RX_CH0STATUS` | 0x60016030 | 32 | APB_MEM_RADDR_CH2=0x1FF000 APB_MEM_RD_ERR_CH2=0x8000000 MEM_FULL_CH2=0x4000000 MEM_OWNER_ERR_CH2=0x2000000 MEM_WADDR_EX_CH2=0x1FF STATE_CH2=0x1C00000 | - |
| `RMT_RX_CH1CONF0` | 0x60016020 | 32 | CARRIER_EN_CH2=0x10000000 CARRIER_OUT_LV_CH2=0x20000000 DIV_CNT_CH2=0xFF IDLE_THRES_CH2=0x7FFF00 MEM_SIZE_CH2=0x3800000 | - |
| `RMT_RX_CH1CONF1` | 0x60016024 | 32 | AFIFO_RST_CH2=0x4000 APB_MEM_RST_CH2=0x4 CONF_UPDATE_CH2=0x8000 MEM_OWNER_CH2=0x8 MEM_RX_WRAP_EN_CH2=0x2000 MEM_WR_RST_CH2=0x2 RX_EN_CH2=0x1 RX_FILTER_EN_CH2=0x10 RX_FILTER_THRES_CH2=0x1FE0 | - |
| `RMT_RX_CH1STATUS` | 0x60016034 | 32 | APB_MEM_RADDR_CH2=0x1FF000 APB_MEM_RD_ERR_CH2=0x8000000 MEM_FULL_CH2=0x4000000 MEM_OWNER_ERR_CH2=0x2000000 MEM_WADDR_EX_CH2=0x1FF STATE_CH2=0x1C00000 | - |
| `RMT_SYS_CONF` | 0x60016068 | 32 | APB_FIFO_MASK=0x1 CLK_EN=0x80000000 MEM_CLK_FORCE_ON=0x2 RMT_MEM_FORCE_PD=0x4 RMT_MEM_FORCE_PU=0x8 RMT_SCLK_ACTIVE=0x4000000 RMT_SCLK_DIV_A=0x3F000 RMT_SCLK_DIV_B=0xFC0000 RMT_SCLK_DIV_NUM=0xFF0 RMT_SCLK_SEL=0x3000000 | - |
| `RMT_TX_CH0CONF0` | 0x60016010 | 32 | AFIFO_RST_CH0=0x800000 APB_MEM_RST_CH0=0x4 CARRIER_EFF_EN_CH0=0x100000 CARRIER_EN_CH0=0x200000 CARRIER_OUT_LV_CH0=0x400000 CONF_UPDATE_CH0=0x1000000 DIV_CNT_CH0=0xFF00 IDLE_OUT_EN_CH0=0x40 IDLE_OUT_LV_CH0=0x20 MEM_RD_RST_CH0=0x2 MEM_SIZE_CH... | - |
| `RMT_TX_CH0DATA` | 0x60016000 | 32 | CHDATA=0xFFFFFFFF | - |
| `RMT_TX_CH0STATUS` | 0x60016028 | 32 | APB_MEM_RADDR_CH0=0xFF000000 APB_MEM_RD_ERR_CH0=0x200000 APB_MEM_WADDR_CH0=0x1FF000 APB_MEM_WR_ERR_CH0=0x800000 MEM_EMPTY_CH0=0x400000 MEM_RADDR_EX_CH0=0x1FF STATE_CH0=0xE00 | - |
| `RMT_TX_CH1CONF0` | 0x60016014 | 32 | AFIFO_RST_CH0=0x800000 APB_MEM_RST_CH0=0x4 CARRIER_EFF_EN_CH0=0x100000 CARRIER_EN_CH0=0x200000 CARRIER_OUT_LV_CH0=0x400000 CONF_UPDATE_CH0=0x1000000 DIV_CNT_CH0=0xFF00 IDLE_OUT_EN_CH0=0x40 IDLE_OUT_LV_CH0=0x20 MEM_RD_RST_CH0=0x2 MEM_SIZE_CH... | - |
| `RMT_TX_CH1DATA` | 0x60016004 | 32 | CHDATA=0xFFFFFFFF | - |
| `RMT_TX_CH1STATUS` | 0x6001602C | 32 | APB_MEM_RADDR_CH0=0xFF000000 APB_MEM_RD_ERR_CH0=0x200000 APB_MEM_WADDR_CH0=0x1FF000 APB_MEM_WR_ERR_CH0=0x800000 MEM_EMPTY_CH0=0x400000 MEM_RADDR_EX_CH0=0x1FF STATE_CH0=0xE00 | - |
| `RMT_TX_CH2DATA` | 0x60016008 | 32 | CHDATA=0xFFFFFFFF | - |
| `RMT_TX_CH3DATA` | 0x6001600C | 32 | CHDATA=0xFFFFFFFF | - |
| `RMT_TX_SIM` | 0x6001606C | 32 | CH0=0x1 CH1=0x2 EN=0x4 | - |
| `RNG_DATA` | 0x600260B0 | 32 |  | - |
| `RSA_CONSTANT_TIME` | 0x6003C820 | 32 | CONSTANT_TIME=0x1 | - |
| `RSA_DATE` | 0x6003C830 | 32 | DATE=0x3FFFFFFF | - |
| `RSA_INT_CLR` | 0x6003C81C | 32 | CLEAR_INTERRUPT=0x1 | - |
| `RSA_INT_ENA` | 0x6003C82C | 32 | INT_ENA=0x1 | - |
| `RSA_MODE` | 0x6003C804 | 32 | MODE=0x7F | - |
| `RSA_M_MEM[0]` | 0x6003C000 | 8 |  | - |
| `RSA_M_MEM[10]` | 0x6003C00A | 8 |  | - |
| `RSA_M_MEM[11]` | 0x6003C00B | 8 |  | - |
| `RSA_M_MEM[12]` | 0x6003C00C | 8 |  | - |
| `RSA_M_MEM[13]` | 0x6003C00D | 8 |  | - |
| `RSA_M_MEM[14]` | 0x6003C00E | 8 |  | - |
| `RSA_M_MEM[15]` | 0x6003C00F | 8 |  | - |
| `RSA_M_MEM[1]` | 0x6003C001 | 8 |  | - |
| `RSA_M_MEM[2]` | 0x6003C002 | 8 |  | - |
| `RSA_M_MEM[3]` | 0x6003C003 | 8 |  | - |
| `RSA_M_MEM[4]` | 0x6003C004 | 8 |  | - |
| `RSA_M_MEM[5]` | 0x6003C005 | 8 |  | - |
| `RSA_M_MEM[6]` | 0x6003C006 | 8 |  | - |
| `RSA_M_MEM[7]` | 0x6003C007 | 8 |  | - |
| `RSA_M_MEM[8]` | 0x6003C008 | 8 |  | - |
| `RSA_M_MEM[9]` | 0x6003C009 | 8 |  | - |
| `RSA_M_PRIME` | 0x6003C800 | 32 | M_PRIME=0xFFFFFFFF | - |
| `RSA_QUERY_CLEAN` | 0x6003C808 | 32 | QUERY_CLEAN=0x1 | - |
| `RSA_QUERY_IDLE` | 0x6003C818 | 32 | QUERY_IDLE=0x1 | - |
| `RSA_SEARCH_ENABLE` | 0x6003C824 | 32 | SEARCH_ENABLE=0x1 | - |
| `RSA_SEARCH_POS` | 0x6003C828 | 32 | SEARCH_POS=0xFFF | - |
| `RSA_SET_START_MODEXP` | 0x6003C80C | 32 | SET_START_MODEXP=0x1 | - |
| `RSA_SET_START_MODMULT` | 0x6003C810 | 32 | SET_START_MODMULT=0x1 | - |
| `RSA_SET_START_MULT` | 0x6003C814 | 32 | SET_START_MULT=0x1 | - |
| `RSA_X_MEM[0]` | 0x6003C600 | 8 |  | - |
| `RSA_X_MEM[10]` | 0x6003C60A | 8 |  | - |
| `RSA_X_MEM[11]` | 0x6003C60B | 8 |  | - |
| `RSA_X_MEM[12]` | 0x6003C60C | 8 |  | - |
| `RSA_X_MEM[13]` | 0x6003C60D | 8 |  | - |
| `RSA_X_MEM[14]` | 0x6003C60E | 8 |  | - |
| `RSA_X_MEM[15]` | 0x6003C60F | 8 |  | - |
| `RSA_X_MEM[1]` | 0x6003C601 | 8 |  | - |
| `RSA_X_MEM[2]` | 0x6003C602 | 8 |  | - |
| `RSA_X_MEM[3]` | 0x6003C603 | 8 |  | - |
| `RSA_X_MEM[4]` | 0x6003C604 | 8 |  | - |
| `RSA_X_MEM[5]` | 0x6003C605 | 8 |  | - |
| `RSA_X_MEM[6]` | 0x6003C606 | 8 |  | - |
| `RSA_X_MEM[7]` | 0x6003C607 | 8 |  | - |
| `RSA_X_MEM[8]` | 0x6003C608 | 8 |  | - |
| `RSA_X_MEM[9]` | 0x6003C609 | 8 |  | - |
| `RSA_Y_MEM[0]` | 0x6003C400 | 8 |  | - |
| `RSA_Y_MEM[10]` | 0x6003C40A | 8 |  | - |
| `RSA_Y_MEM[11]` | 0x6003C40B | 8 |  | - |
| `RSA_Y_MEM[12]` | 0x6003C40C | 8 |  | - |
| `RSA_Y_MEM[13]` | 0x6003C40D | 8 |  | - |
| `RSA_Y_MEM[14]` | 0x6003C40E | 8 |  | - |
| `RSA_Y_MEM[15]` | 0x6003C40F | 8 |  | - |
| `RSA_Y_MEM[1]` | 0x6003C401 | 8 |  | - |
| `RSA_Y_MEM[2]` | 0x6003C402 | 8 |  | - |
| `RSA_Y_MEM[3]` | 0x6003C403 | 8 |  | - |
| `RSA_Y_MEM[4]` | 0x6003C404 | 8 |  | - |
| `RSA_Y_MEM[5]` | 0x6003C405 | 8 |  | - |
| `RSA_Y_MEM[6]` | 0x6003C406 | 8 |  | - |
| `RSA_Y_MEM[7]` | 0x6003C407 | 8 |  | - |
| `RSA_Y_MEM[8]` | 0x6003C408 | 8 |  | - |
| `RSA_Y_MEM[9]` | 0x6003C409 | 8 |  | - |
| `RSA_Z_MEM[0]` | 0x6003C200 | 8 |  | - |
| `RSA_Z_MEM[10]` | 0x6003C20A | 8 |  | - |
| `RSA_Z_MEM[11]` | 0x6003C20B | 8 |  | - |
| `RSA_Z_MEM[12]` | 0x6003C20C | 8 |  | - |
| `RSA_Z_MEM[13]` | 0x6003C20D | 8 |  | - |
| `RSA_Z_MEM[14]` | 0x6003C20E | 8 |  | - |
| `RSA_Z_MEM[15]` | 0x6003C20F | 8 |  | - |
| `RSA_Z_MEM[1]` | 0x6003C201 | 8 |  | - |
| `RSA_Z_MEM[2]` | 0x6003C202 | 8 |  | - |
| `RSA_Z_MEM[3]` | 0x6003C203 | 8 |  | - |
| `RSA_Z_MEM[4]` | 0x6003C204 | 8 |  | - |
| `RSA_Z_MEM[5]` | 0x6003C205 | 8 |  | - |
| `RSA_Z_MEM[6]` | 0x6003C206 | 8 |  | - |
| `RSA_Z_MEM[7]` | 0x6003C207 | 8 |  | - |
| `RSA_Z_MEM[8]` | 0x6003C208 | 8 |  | - |
| `RSA_Z_MEM[9]` | 0x6003C209 | 8 |  | - |
| `SHA_BUSY` | 0x6003B018 | 32 | STATE=0x1 | - |
| `SHA_CLEAR_IRQ` | 0x6003B024 | 32 | CLEAR_INTERRUPT=0x1 | - |
| `SHA_CONTINUE` | 0x6003B014 | 32 | CONTINUE=0xFFFFFFFE | - |
| `SHA_DATE` | 0x6003B02C | 32 | DATE=0x3FFFFFFF | - |
| `SHA_DMA_BLOCK_NUM` | 0x6003B00C | 32 | DMA_BLOCK_NUM=0x3F | - |
| `SHA_DMA_CONTINUE` | 0x6003B020 | 32 | DMA_CONTINUE=0x1 | - |
| `SHA_DMA_START` | 0x6003B01C | 32 | DMA_START=0x1 | - |
| `SHA_H_MEM[0]` | 0x6003B040 | 8 |  | - |
| `SHA_H_MEM[10]` | 0x6003B04A | 8 |  | - |
| `SHA_H_MEM[11]` | 0x6003B04B | 8 |  | - |
| `SHA_H_MEM[12]` | 0x6003B04C | 8 |  | - |
| `SHA_H_MEM[13]` | 0x6003B04D | 8 |  | - |
| `SHA_H_MEM[14]` | 0x6003B04E | 8 |  | - |
| `SHA_H_MEM[15]` | 0x6003B04F | 8 |  | - |
| `SHA_H_MEM[16]` | 0x6003B050 | 8 |  | - |
| `SHA_H_MEM[17]` | 0x6003B051 | 8 |  | - |
| `SHA_H_MEM[18]` | 0x6003B052 | 8 |  | - |
| `SHA_H_MEM[19]` | 0x6003B053 | 8 |  | - |
| `SHA_H_MEM[1]` | 0x6003B041 | 8 |  | - |
| `SHA_H_MEM[20]` | 0x6003B054 | 8 |  | - |
| `SHA_H_MEM[21]` | 0x6003B055 | 8 |  | - |
| `SHA_H_MEM[22]` | 0x6003B056 | 8 |  | - |
| `SHA_H_MEM[23]` | 0x6003B057 | 8 |  | - |
| `SHA_H_MEM[24]` | 0x6003B058 | 8 |  | - |
| `SHA_H_MEM[25]` | 0x6003B059 | 8 |  | - |
| `SHA_H_MEM[26]` | 0x6003B05A | 8 |  | - |
| `SHA_H_MEM[27]` | 0x6003B05B | 8 |  | - |
| `SHA_H_MEM[28]` | 0x6003B05C | 8 |  | - |
| `SHA_H_MEM[29]` | 0x6003B05D | 8 |  | - |
| `SHA_H_MEM[2]` | 0x6003B042 | 8 |  | - |
| `SHA_H_MEM[30]` | 0x6003B05E | 8 |  | - |
| `SHA_H_MEM[31]` | 0x6003B05F | 8 |  | - |
| `SHA_H_MEM[32]` | 0x6003B060 | 8 |  | - |
| `SHA_H_MEM[33]` | 0x6003B061 | 8 |  | - |
| `SHA_H_MEM[34]` | 0x6003B062 | 8 |  | - |
| `SHA_H_MEM[35]` | 0x6003B063 | 8 |  | - |
| `SHA_H_MEM[36]` | 0x6003B064 | 8 |  | - |
| `SHA_H_MEM[37]` | 0x6003B065 | 8 |  | - |
| `SHA_H_MEM[38]` | 0x6003B066 | 8 |  | - |
| `SHA_H_MEM[39]` | 0x6003B067 | 8 |  | - |
| `SHA_H_MEM[3]` | 0x6003B043 | 8 |  | - |
| `SHA_H_MEM[40]` | 0x6003B068 | 8 |  | - |
| `SHA_H_MEM[41]` | 0x6003B069 | 8 |  | - |
| `SHA_H_MEM[42]` | 0x6003B06A | 8 |  | - |
| `SHA_H_MEM[43]` | 0x6003B06B | 8 |  | - |
| `SHA_H_MEM[44]` | 0x6003B06C | 8 |  | - |
| `SHA_H_MEM[45]` | 0x6003B06D | 8 |  | - |
| `SHA_H_MEM[46]` | 0x6003B06E | 8 |  | - |
| `SHA_H_MEM[47]` | 0x6003B06F | 8 |  | - |
| `SHA_H_MEM[48]` | 0x6003B070 | 8 |  | - |
| `SHA_H_MEM[49]` | 0x6003B071 | 8 |  | - |
| `SHA_H_MEM[4]` | 0x6003B044 | 8 |  | - |
| `SHA_H_MEM[50]` | 0x6003B072 | 8 |  | - |
| `SHA_H_MEM[51]` | 0x6003B073 | 8 |  | - |
| `SHA_H_MEM[52]` | 0x6003B074 | 8 |  | - |
| `SHA_H_MEM[53]` | 0x6003B075 | 8 |  | - |
| `SHA_H_MEM[54]` | 0x6003B076 | 8 |  | - |
| `SHA_H_MEM[55]` | 0x6003B077 | 8 |  | - |
| `SHA_H_MEM[56]` | 0x6003B078 | 8 |  | - |
| `SHA_H_MEM[57]` | 0x6003B079 | 8 |  | - |
| `SHA_H_MEM[58]` | 0x6003B07A | 8 |  | - |
| `SHA_H_MEM[59]` | 0x6003B07B | 8 |  | - |
| `SHA_H_MEM[5]` | 0x6003B045 | 8 |  | - |
| `SHA_H_MEM[60]` | 0x6003B07C | 8 |  | - |
| `SHA_H_MEM[61]` | 0x6003B07D | 8 |  | - |
| `SHA_H_MEM[62]` | 0x6003B07E | 8 |  | - |
| `SHA_H_MEM[63]` | 0x6003B07F | 8 |  | - |
| `SHA_H_MEM[6]` | 0x6003B046 | 8 |  | - |
| `SHA_H_MEM[7]` | 0x6003B047 | 8 |  | - |
| `SHA_H_MEM[8]` | 0x6003B048 | 8 |  | - |
| `SHA_H_MEM[9]` | 0x6003B049 | 8 |  | - |
| `SHA_IRQ_ENA` | 0x6003B028 | 32 | INTERRUPT_ENA=0x1 | - |
| `SHA_MODE` | 0x6003B000 | 32 | MODE=0x7 | - |
| `SHA_M_MEM[0]` | 0x6003B080 | 8 |  | - |
| `SHA_M_MEM[10]` | 0x6003B08A | 8 |  | - |
| `SHA_M_MEM[11]` | 0x6003B08B | 8 |  | - |
| `SHA_M_MEM[12]` | 0x6003B08C | 8 |  | - |
| `SHA_M_MEM[13]` | 0x6003B08D | 8 |  | - |
| `SHA_M_MEM[14]` | 0x6003B08E | 8 |  | - |
| `SHA_M_MEM[15]` | 0x6003B08F | 8 |  | - |
| `SHA_M_MEM[16]` | 0x6003B090 | 8 |  | - |
| `SHA_M_MEM[17]` | 0x6003B091 | 8 |  | - |
| `SHA_M_MEM[18]` | 0x6003B092 | 8 |  | - |
| `SHA_M_MEM[19]` | 0x6003B093 | 8 |  | - |
| `SHA_M_MEM[1]` | 0x6003B081 | 8 |  | - |
| `SHA_M_MEM[20]` | 0x6003B094 | 8 |  | - |
| `SHA_M_MEM[21]` | 0x6003B095 | 8 |  | - |
| `SHA_M_MEM[22]` | 0x6003B096 | 8 |  | - |
| `SHA_M_MEM[23]` | 0x6003B097 | 8 |  | - |
| `SHA_M_MEM[24]` | 0x6003B098 | 8 |  | - |
| `SHA_M_MEM[25]` | 0x6003B099 | 8 |  | - |
| `SHA_M_MEM[26]` | 0x6003B09A | 8 |  | - |
| `SHA_M_MEM[27]` | 0x6003B09B | 8 |  | - |
| `SHA_M_MEM[28]` | 0x6003B09C | 8 |  | - |
| `SHA_M_MEM[29]` | 0x6003B09D | 8 |  | - |
| `SHA_M_MEM[2]` | 0x6003B082 | 8 |  | - |
| `SHA_M_MEM[30]` | 0x6003B09E | 8 |  | - |
| `SHA_M_MEM[31]` | 0x6003B09F | 8 |  | - |
| `SHA_M_MEM[32]` | 0x6003B0A0 | 8 |  | - |
| `SHA_M_MEM[33]` | 0x6003B0A1 | 8 |  | - |
| `SHA_M_MEM[34]` | 0x6003B0A2 | 8 |  | - |
| `SHA_M_MEM[35]` | 0x6003B0A3 | 8 |  | - |
| `SHA_M_MEM[36]` | 0x6003B0A4 | 8 |  | - |
| `SHA_M_MEM[37]` | 0x6003B0A5 | 8 |  | - |
| `SHA_M_MEM[38]` | 0x6003B0A6 | 8 |  | - |
| `SHA_M_MEM[39]` | 0x6003B0A7 | 8 |  | - |
| `SHA_M_MEM[3]` | 0x6003B083 | 8 |  | - |
| `SHA_M_MEM[40]` | 0x6003B0A8 | 8 |  | - |
| `SHA_M_MEM[41]` | 0x6003B0A9 | 8 |  | - |
| `SHA_M_MEM[42]` | 0x6003B0AA | 8 |  | - |
| `SHA_M_MEM[43]` | 0x6003B0AB | 8 |  | - |
| `SHA_M_MEM[44]` | 0x6003B0AC | 8 |  | - |
| `SHA_M_MEM[45]` | 0x6003B0AD | 8 |  | - |
| `SHA_M_MEM[46]` | 0x6003B0AE | 8 |  | - |
| `SHA_M_MEM[47]` | 0x6003B0AF | 8 |  | - |
| `SHA_M_MEM[48]` | 0x6003B0B0 | 8 |  | - |
| `SHA_M_MEM[49]` | 0x6003B0B1 | 8 |  | - |
| `SHA_M_MEM[4]` | 0x6003B084 | 8 |  | - |
| `SHA_M_MEM[50]` | 0x6003B0B2 | 8 |  | - |
| `SHA_M_MEM[51]` | 0x6003B0B3 | 8 |  | - |
| `SHA_M_MEM[52]` | 0x6003B0B4 | 8 |  | - |
| `SHA_M_MEM[53]` | 0x6003B0B5 | 8 |  | - |
| `SHA_M_MEM[54]` | 0x6003B0B6 | 8 |  | - |
| `SHA_M_MEM[55]` | 0x6003B0B7 | 8 |  | - |
| `SHA_M_MEM[56]` | 0x6003B0B8 | 8 |  | - |
| `SHA_M_MEM[57]` | 0x6003B0B9 | 8 |  | - |
| `SHA_M_MEM[58]` | 0x6003B0BA | 8 |  | - |
| `SHA_M_MEM[59]` | 0x6003B0BB | 8 |  | - |
| `SHA_M_MEM[5]` | 0x6003B085 | 8 |  | - |
| `SHA_M_MEM[60]` | 0x6003B0BC | 8 |  | - |
| `SHA_M_MEM[61]` | 0x6003B0BD | 8 |  | - |
| `SHA_M_MEM[62]` | 0x6003B0BE | 8 |  | - |
| `SHA_M_MEM[63]` | 0x6003B0BF | 8 |  | - |
| `SHA_M_MEM[6]` | 0x6003B086 | 8 |  | - |
| `SHA_M_MEM[7]` | 0x6003B087 | 8 |  | - |
| `SHA_M_MEM[8]` | 0x6003B088 | 8 |  | - |
| `SHA_M_MEM[9]` | 0x6003B089 | 8 |  | - |
| `SHA_START` | 0x6003B010 | 32 | START=0xFFFFFFFE | - |
| `SHA_T_LENGTH` | 0x6003B008 | 32 | T_LENGTH=0x3F | - |
| `SHA_T_STRING` | 0x6003B004 | 32 | T_STRING=0xFFFFFFFF | - |
| `SPI0_SPI_FMEM_PMS0_ADDR` | 0x60003110 | 32 | S=0x3FFFFFF | - |
| `SPI0_SPI_FMEM_PMS0_ATTR` | 0x60003100 | 32 | SPI_FMEM_PMS_ECC=0x4 SPI_FMEM_PMS_RD_ATTR=0x1 SPI_FMEM_PMS_WR_ATTR=0x2 | - |
| `SPI0_SPI_FMEM_PMS0_SIZE` | 0x60003120 | 32 | SPI_FMEM_PMS_SIZE=0x3FFF | - |
| `SPI0_SPI_FMEM_PMS1_ADDR` | 0x60003114 | 32 | S=0x3FFFFFF | - |
| `SPI0_SPI_FMEM_PMS1_ATTR` | 0x60003104 | 32 | SPI_FMEM_PMS_ECC=0x4 SPI_FMEM_PMS_RD_ATTR=0x1 SPI_FMEM_PMS_WR_ATTR=0x2 | - |
| `SPI0_SPI_FMEM_PMS1_SIZE` | 0x60003124 | 32 | SPI_FMEM_PMS_SIZE=0x3FFF | - |
| `SPI0_SPI_FMEM_PMS2_ADDR` | 0x60003118 | 32 | S=0x3FFFFFF | - |
| `SPI0_SPI_FMEM_PMS2_ATTR` | 0x60003108 | 32 | SPI_FMEM_PMS_ECC=0x4 SPI_FMEM_PMS_RD_ATTR=0x1 SPI_FMEM_PMS_WR_ATTR=0x2 | - |
| `SPI0_SPI_FMEM_PMS2_SIZE` | 0x60003128 | 32 | SPI_FMEM_PMS_SIZE=0x3FFF | - |
| `SPI0_SPI_FMEM_PMS3_ADDR` | 0x6000311C | 32 | S=0x3FFFFFF | - |
| `SPI0_SPI_FMEM_PMS3_ATTR` | 0x6000310C | 32 | SPI_FMEM_PMS_ECC=0x4 SPI_FMEM_PMS_RD_ATTR=0x1 SPI_FMEM_PMS_WR_ATTR=0x2 | - |
| `SPI0_SPI_FMEM_PMS3_SIZE` | 0x6000312C | 32 | SPI_FMEM_PMS_SIZE=0x3FFF | - |
| `SPI0_SPI_MEM_AXI_ERR_ADDR` | 0x60003170 | 32 | SPI_ALL_AXI_TRANS_AFIFO_EMPTY=0x80000000 SPI_MEM_ALL_FIFO_EMPTY=0x4000000 SPI_MEM_AXI_ERR_ADDR=0x3FFFFFF SPI_RADDR_AFIFO_REMPTY=0x10000000 SPI_RDATA_AFIFO_REMPTY=0x8000000 SPI_WBLEN_AFIFO_REMPTY=0x40000000 SPI_WDATA_AFIFO_REMPTY=0x20000000 | - |
| `SPI0_SPI_MEM_CACHE_FCTRL` | 0x6000303C | 32 | SPI_CLOSE_AXI_INF_EN=0x80000000 SPI_MEM_AXI_REQ_EN=0x1 SPI_MEM_CACHE_FLASH_USR_CMD=0x4 SPI_MEM_CACHE_USR_ADDR_4BYTE=0x2 SPI_MEM_FADDR_DUAL=0x20 SPI_MEM_FADDR_QUAD=0x100 SPI_MEM_FDIN_DUAL=0x8 SPI_MEM_FDIN_QUAD=0x40 SPI_MEM_FDOUT_DUAL=0x10 SP... | - |
| `SPI0_SPI_MEM_CACHE_SCTRL` | 0x60003040 | 32 | SPI_MEM_CACHE_SRAM_USR_RCMD=0x20 SPI_MEM_CACHE_SRAM_USR_WCMD=0x100000 SPI_MEM_CACHE_USR_SADDR_4BYTE=0x1 SPI_MEM_SRAM_ADDR_BITLEN=0xFC000 SPI_MEM_SRAM_OCT=0x200000 SPI_MEM_SRAM_RDUMMY_CYCLELEN=0xFC0 SPI_MEM_SRAM_WDUMMY_CYCLELEN=0xFC00000 SPI... | - |
| `SPI0_SPI_MEM_CLOCK` | 0x60003014 | 32 | SPI_MEM_CLKCNT_H=0xFF00 SPI_MEM_CLKCNT_L=0xFF SPI_MEM_CLKCNT_N=0xFF0000 SPI_MEM_CLK_EQU_SYSCLK=0x80000000 | - |
| `SPI0_SPI_MEM_CLOCK_GATE` | 0x60003200 | 32 | SPI_CLK_EN=0x1 | - |
| `SPI0_SPI_MEM_CMD` | 0x60003000 | 32 | SPI_MEM_MST_ST=0xF SPI_MEM_SLV_ST=0xF0 SPI_MEM_USR=0x40000 | - |
| `SPI0_SPI_MEM_CTRL` | 0x60003008 | 32 | SPI_MEM_DATA_IE_ALWAYS_ON=0x80000000 SPI_MEM_DQS_IE_ALWAYS_ON=0x40000000 SPI_MEM_D_POL=0x80000 SPI_MEM_FADDR_OCT=0x40 SPI_MEM_FASTRD_MODE=0x2000 SPI_MEM_FCMD_OCT=0x200 SPI_MEM_FCMD_QUAD=0x100 SPI_MEM_FDIN_OCT=0x20 SPI_MEM_FDOUT_OCT=0x10 SPI... | - |
| `SPI0_SPI_MEM_CTRL1` | 0x6000300C | 32 | SPI_AR_SIZE0_1_SUPPORT_EN=0x200000 SPI_AW_SIZE0_1_SUPPORT_EN=0x400000 SPI_AXI_RDATA_BACK_FAST=0x800000 SPI_MEM_AR_SPLICE_EN=0x2000000 SPI_MEM_AW_SPLICE_EN=0x4000000 SPI_MEM_CLK_MODE=0x3 SPI_MEM_DUAL_RAM_EN=0x10000000 SPI_MEM_FAST_WRITE_EN=0... | - |
| `SPI0_SPI_MEM_CTRL2` | 0x60003010 | 32 | SPI_MEM_CS_HOLD_DELAY=0x7E000000 SPI_MEM_CS_HOLD_TIME=0x3E0 SPI_MEM_CS_SETUP_TIME=0x1F SPI_MEM_ECC_16TO18_BYTE_EN=0x4000 SPI_MEM_ECC_CS_HOLD_TIME=0x1C00 SPI_MEM_ECC_SKIP_PAGE_CORNER=0x2000 SPI_MEM_SPLIT_TRANS_EN=0x1000000 SPI_MEM_SYNC_RESET... | - |
| `SPI0_SPI_MEM_DATE` | 0x600033FC | 32 | SPI_MEM_DATE=0xFFFFFFF | - |
| `SPI0_SPI_MEM_DDR` | 0x600030D4 | 32 | SPI_FMEM_CLK_DIFF_EN=0x1000000 SPI_FMEM_CLK_DIFF_INV=0x10000000 SPI_FMEM_DDR_CMD_DIS=0x10 SPI_FMEM_DDR_DQS_LOOP=0x200000 SPI_FMEM_DDR_EN=0x1 SPI_FMEM_DDR_RDAT_SWP=0x4 SPI_FMEM_DDR_WDAT_SWP=0x8 SPI_FMEM_DQS_CA_IN=0x4000000 SPI_FMEM_HYPERBUS_... | - |
| `SPI0_SPI_MEM_DIN_MODE` | 0x60003184 | 32 | SPI_MEM_DIN0_MODE=0x7 SPI_MEM_DIN1_MODE=0x38 SPI_MEM_DIN2_MODE=0x1C0 SPI_MEM_DIN3_MODE=0xE00 SPI_MEM_DIN4_MODE=0x7000 SPI_MEM_DIN5_MODE=0x38000 SPI_MEM_DIN6_MODE=0x1C0000 SPI_MEM_DIN7_MODE=0xE00000 SPI_MEM_DINS_MODE=0x7000000 | - |
| `SPI0_SPI_MEM_DIN_NUM` | 0x60003188 | 32 | SPI_MEM_DIN0_NUM=0x3 SPI_MEM_DIN1_NUM=0xC SPI_MEM_DIN2_NUM=0x30 SPI_MEM_DIN3_NUM=0xC0 SPI_MEM_DIN4_NUM=0x300 SPI_MEM_DIN5_NUM=0xC00 SPI_MEM_DIN6_NUM=0x3000 SPI_MEM_DIN7_NUM=0xC000 SPI_MEM_DINS_NUM=0x30000 | - |
| `SPI0_SPI_MEM_DOUT_MODE` | 0x6000318C | 32 | SPI_MEM_DOUT0_MODE=0x1 SPI_MEM_DOUT1_MODE=0x2 SPI_MEM_DOUT2_MODE=0x4 SPI_MEM_DOUT3_MODE=0x8 SPI_MEM_DOUT4_MODE=0x10 SPI_MEM_DOUT5_MODE=0x20 SPI_MEM_DOUT6_MODE=0x40 SPI_MEM_DOUT7_MODE=0x80 SPI_MEM_DOUTS_MODE=0x100 | - |
| `SPI0_SPI_MEM_DPA_CTRL` | 0x60003388 | 32 | SPI_CRYPT_CALC_D_DPA_EN=0x8 SPI_CRYPT_DPA_SELECT_REGISTER=0x10 SPI_CRYPT_SECURITY_LEVEL=0x7 | - |
| `SPI0_SPI_MEM_ECC_CTRL` | 0x60003168 | 32 | SPI_FMEM_ECC_ADDR_EN=0x100000 SPI_FMEM_ECC_ERR_INT_EN=0x20000 SPI_FMEM_ECC_ERR_INT_NUM=0x1F800 SPI_FMEM_PAGE_SIZE=0xC0000 SPI_MEM_ECC_CONTINUE_RECORD_ERR_EN=0x1000000 SPI_MEM_ECC_ERR_BITS=0xFE000000 SPI_MEM_USR_ECC_ADDR_EN=0x200000 | - |
| `SPI0_SPI_MEM_ECC_ERR_ADDR` | 0x6000316C | 32 | SPI_MEM_ECC_ERR_ADDR=0x3FFFFFF SPI_MEM_ECC_ERR_CNT=0xFC000000 | - |
| `SPI0_SPI_MEM_FSM` | 0x60003054 | 32 | SPI_MEM_LOCK_DELAY_TIME=0xF80 | - |
| `SPI0_SPI_MEM_INT_CLR` | 0x600030C4 | 32 | SPI_MEM_AXI_RADDR_ERR_INT_CLR=0x80 SPI_MEM_AXI_WADDR_ERR_INT_CLR=0x200 SPI_MEM_AXI_WR_FLASH_ERR_INT_CLR=0x100 SPI_MEM_ECC_ERR_INT_CLR=0x20 SPI_MEM_MST_ST_END_INT_CLR=0x10 SPI_MEM_PMS_REJECT_INT_CLR=0x40 SPI_MEM_SLV_ST_END_INT_CLR=0x8 | - |
| `SPI0_SPI_MEM_INT_ENA` | 0x600030C0 | 32 | SPI_MEM_AXI_RADDR_ERR_INT_ENA=0x80 SPI_MEM_AXI_WADDR_ERR_INT__ENA=0x200 SPI_MEM_AXI_WR_FLASH_ERR_INT_ENA=0x100 SPI_MEM_ECC_ERR_INT_ENA=0x20 SPI_MEM_MST_ST_END_INT_ENA=0x10 SPI_MEM_PMS_REJECT_INT_ENA=0x40 SPI_MEM_SLV_ST_END_INT_ENA=0x8 | - |
| `SPI0_SPI_MEM_INT_RAW` | 0x600030C8 | 32 | SPI_MEM_AXI_RADDR_ERR_INT_RAW=0x80 SPI_MEM_AXI_WADDR_ERR_INT_RAW=0x200 SPI_MEM_AXI_WR_FLASH_ERR_INT_RAW=0x100 SPI_MEM_ECC_ERR_INT_RAW=0x20 SPI_MEM_MST_ST_END_INT_RAW=0x10 SPI_MEM_PMS_REJECT_INT_RAW=0x40 SPI_MEM_SLV_ST_END_INT_RAW=0x8 | - |
| `SPI0_SPI_MEM_INT_ST` | 0x600030CC | 32 | SPI_MEM_AXI_RADDR_ERR_INT_ST=0x80 SPI_MEM_AXI_WADDR_ERR_INT_ST=0x200 SPI_MEM_AXI_WR_FLASH_ERR_INT_ST=0x100 SPI_MEM_ECC_ERR_INT_ST=0x20 SPI_MEM_MST_ST_END_INT_ST=0x10 SPI_MEM_PMS_REJECT_INT_ST=0x40 SPI_MEM_SLV_ST_END_INT_ST=0x8 | - |
| `SPI0_SPI_MEM_MISC` | 0x60003034 | 32 | SPI_MEM_CK_IDLE_EDGE=0x200 SPI_MEM_CS_KEEP_ACTIVE=0x400 SPI_MEM_FSUB_PIN=0x80 SPI_MEM_SSUB_PIN=0x100 | - |
| `SPI0_SPI_MEM_MMU_ITEM_CONTENT` | 0x6000337C | 32 | SPI_MMU_ITEM_CONTENT=0xFFFFFFFF | - |
| `SPI0_SPI_MEM_MMU_ITEM_INDEX` | 0x60003380 | 32 | SPI_MMU_ITEM_INDEX=0xFFFFFFFF | - |
| `SPI0_SPI_MEM_MMU_POWER_CTRL` | 0x60003384 | 32 | SPI_MEM_AUX_CTRL=0x3FFF0000 SPI_MEM_RDN_ENA=0x40000000 SPI_MEM_RDN_RESULT=0x80000000 SPI_MMU_MEM_FORCE_ON=0x1 SPI_MMU_MEM_FORCE_PD=0x2 SPI_MMU_MEM_FORCE_PU=0x4 SPI_MMU_PAGE_SIZE=0x18 | - |
| `SPI0_SPI_MEM_PMS_REJECT` | 0x60003164 | 32 | SPI_MEM_PMS_IVD=0x80000000 SPI_MEM_PMS_LD=0x10000000 SPI_MEM_PMS_MULTI_HIT=0x40000000 SPI_MEM_PMS_ST=0x20000000 SPI_MEM_PM_EN=0x4000000 SPI_MEM_REJECT_ADDR=0x3FFFFFF | - |
| `SPI0_SPI_MEM_RD_STATUS` | 0x6000302C | 32 | SPI_MEM_WB_MODE=0xFF0000 | - |
| `SPI0_SPI_MEM_REGISTERRND_ECO_HIGH` | 0x600033F0 | 32 | SPI_MEM_REGISTERRND_ECO_HIGH=0xFFFFFFFF | - |
| `SPI0_SPI_MEM_REGISTERRND_ECO_LOW` | 0x600033F4 | 32 | SPI_MEM_REGISTERRND_ECO_LOW=0xFFFFFFFF | - |
| `SPI0_SPI_MEM_SRAM_CLK` | 0x60003050 | 32 | SPI_MEM_SCLKCNT_H=0xFF00 SPI_MEM_SCLKCNT_L=0xFF SPI_MEM_SCLKCNT_N=0xFF0000 SPI_MEM_SCLK_EQU_SYSCLK=0x80000000 | - |
| `SPI0_SPI_MEM_SRAM_CMD` | 0x60003044 | 32 | SPI_MEM_SADDR_DUAL=0x1000 SPI_MEM_SADDR_OCT=0x100000 SPI_MEM_SADDR_QUAD=0x10000 SPI_MEM_SCLK_MODE=0x3 SPI_MEM_SCMD_OCT=0x200000 SPI_MEM_SCMD_QUAD=0x20000 SPI_MEM_SDIN_DUAL=0x400 SPI_MEM_SDIN_OCT=0x40000 SPI_MEM_SDIN_QUAD=0x4000 SPI_MEM_SDOU... | - |
| `SPI0_SPI_MEM_SRAM_DRD_CMD` | 0x60003048 | 32 | SPI_MEM_CACHE_SRAM_USR_RD_CMD_BITLEN=0xF0000000 SPI_MEM_CACHE_SRAM_USR_RD_CMD_VALUE=0xFFFF | - |
| `SPI0_SPI_MEM_SRAM_DWR_CMD` | 0x6000304C | 32 | SPI_MEM_CACHE_SRAM_USR_WR_CMD_BITLEN=0xF0000000 SPI_MEM_CACHE_SRAM_USR_WR_CMD_VALUE=0xFFFF | - |
| `SPI0_SPI_MEM_TIMING_CALI` | 0x60003180 | 32 | SPI_MEM_DLL_TIMING_CALI=0x20 SPI_MEM_EXTRA_DUMMY_CYCLELEN=0x1C SPI_MEM_TIMING_CALI=0x2 SPI_MEM_TIMING_CLK_ENA=0x1 UPDATE=0x40 | - |
| `SPI0_SPI_MEM_USER` | 0x60003018 | 32 | SPI_MEM_CK_OUT_EDGE=0x200 SPI_MEM_CS_HOLD=0x40 SPI_MEM_CS_SETUP=0x80 SPI_MEM_USR_DUMMY=0x20000000 SPI_MEM_USR_DUMMY_IDLE=0x4000000 | - |
| `SPI0_SPI_MEM_USER1` | 0x6000301C | 32 | SPI_MEM_USR_ADDR_BITLEN=0xFC000000 SPI_MEM_USR_DBYTELEN=0x1C0 SPI_MEM_USR_DUMMY_CYCLELEN=0x3F | - |
| `SPI0_SPI_MEM_USER2` | 0x60003020 | 32 | SPI_MEM_USR_COMMAND_BITLEN=0xF0000000 SPI_MEM_USR_COMMAND_VALUE=0xFFFF | - |
| `SPI0_SPI_MEM_XTS_DATE` | 0x6000335C | 32 | SPI_XTS_DATE=0x3FFFFFFF | - |
| `SPI0_SPI_MEM_XTS_DESTINATION` | 0x60003344 | 32 | SPI_XTS_DESTINATION=0x1 | - |
| `SPI0_SPI_MEM_XTS_DESTROY` | 0x60003354 | 32 | SPI_XTS_DESTROY=0x1 | - |
| `SPI0_SPI_MEM_XTS_LINESIZE` | 0x60003340 | 32 | SPI_XTS_LINESIZE=0x3 | - |
| `SPI0_SPI_MEM_XTS_PHYSICAL_ADDRESS` | 0x60003348 | 32 | SPI_XTS_PHYSICAL_ADDRESS=0x3FFFFFF | - |
| `SPI0_SPI_MEM_XTS_PLAIN_BASE` | 0x60003300 | 32 | SPI_XTS_PLAIN=0xFFFFFFFF | - |
| `SPI0_SPI_MEM_XTS_RELEASE` | 0x60003350 | 32 | SPI_XTS_RELEASE=0x1 | - |
| `SPI0_SPI_MEM_XTS_STATE` | 0x60003358 | 32 | SPI_XTS_STATE=0x3 | - |
| `SPI0_SPI_MEM_XTS_TRIGGER` | 0x6000334C | 32 | SPI_XTS_TRIGGER=0x1 | - |
| `SPI0_SPI_SMEM_AC` | 0x600031A0 | 32 | SPI_SMEM_CS_HOLD=0x2 SPI_SMEM_CS_HOLD_DELAY=0x7E000000 SPI_SMEM_CS_HOLD_TIME=0xF80 SPI_SMEM_CS_SETUP=0x1 SPI_SMEM_CS_SETUP_TIME=0x7C SPI_SMEM_ECC_16TO18_BYTE_EN=0x10000 SPI_SMEM_ECC_CS_HOLD_TIME=0x7000 SPI_SMEM_ECC_SKIP_PAGE_CORNER=0x8000 S... | - |
| `SPI0_SPI_SMEM_DDR` | 0x600030D8 | 32 | CMD_DIS=0x10 DQS_LOOP=0x200000 EN=0x1 RDAT_SWP=0x4 SPI_SMEM_CLK_DIFF_EN=0x1000000 SPI_SMEM_CLK_DIFF_INV=0x10000000 SPI_SMEM_DQS_CA_IN=0x4000000 SPI_SMEM_HYPERBUS_CA=0x40000000 SPI_SMEM_HYPERBUS_DUMMY_2X=0x8000000 SPI_SMEM_OCTA_RAM_ADDR=0x20... | - |
| `SPI0_SPI_SMEM_DIN_MODE` | 0x60003194 | 32 | SPI_SMEM_DIN0_MODE=0x7 SPI_SMEM_DIN1_MODE=0x38 SPI_SMEM_DIN2_MODE=0x1C0 SPI_SMEM_DIN3_MODE=0xE00 SPI_SMEM_DIN4_MODE=0x7000 SPI_SMEM_DIN5_MODE=0x38000 SPI_SMEM_DIN6_MODE=0x1C0000 SPI_SMEM_DIN7_MODE=0xE00000 SPI_SMEM_DINS_MODE=0x7000000 | - |
| `SPI0_SPI_SMEM_DIN_NUM` | 0x60003198 | 32 | SPI_SMEM_DIN0_NUM=0x3 SPI_SMEM_DIN1_NUM=0xC SPI_SMEM_DIN2_NUM=0x30 SPI_SMEM_DIN3_NUM=0xC0 SPI_SMEM_DIN4_NUM=0x300 SPI_SMEM_DIN5_NUM=0xC00 SPI_SMEM_DIN6_NUM=0x3000 SPI_SMEM_DIN7_NUM=0xC000 SPI_SMEM_DINS_NUM=0x30000 | - |
| `SPI0_SPI_SMEM_DOUT_MODE` | 0x6000319C | 32 | SPI_SMEM_DOUT0_MODE=0x1 SPI_SMEM_DOUT1_MODE=0x2 SPI_SMEM_DOUT2_MODE=0x4 SPI_SMEM_DOUT3_MODE=0x8 SPI_SMEM_DOUT4_MODE=0x10 SPI_SMEM_DOUT5_MODE=0x20 SPI_SMEM_DOUT6_MODE=0x40 SPI_SMEM_DOUT7_MODE=0x80 SPI_SMEM_DOUTS_MODE=0x100 | - |
| `SPI0_SPI_SMEM_ECC_CTRL` | 0x60003174 | 32 | SPI_SMEM_ECC_ADDR_EN=0x100000 SPI_SMEM_ECC_ERR_INT_EN=0x20000 SPI_SMEM_PAGE_SIZE=0xC0000 | - |
| `SPI0_SPI_SMEM_PMS0_ADDR` | 0x60003140 | 32 | S=0x3FFFFFF | - |
| `SPI0_SPI_SMEM_PMS0_ATTR` | 0x60003130 | 32 | SPI_SMEM_PMS_ECC=0x4 SPI_SMEM_PMS_RD_ATTR=0x1 SPI_SMEM_PMS_WR_ATTR=0x2 | - |
| `SPI0_SPI_SMEM_PMS0_SIZE` | 0x60003150 | 32 | SPI_SMEM_PMS_SIZE=0x3FFF | - |
| `SPI0_SPI_SMEM_PMS1_ADDR` | 0x60003144 | 32 | S=0x3FFFFFF | - |
| `SPI0_SPI_SMEM_PMS1_ATTR` | 0x60003134 | 32 | SPI_SMEM_PMS_ECC=0x4 SPI_SMEM_PMS_RD_ATTR=0x1 SPI_SMEM_PMS_WR_ATTR=0x2 | - |
| `SPI0_SPI_SMEM_PMS1_SIZE` | 0x60003154 | 32 | SPI_SMEM_PMS_SIZE=0x3FFF | - |
| `SPI0_SPI_SMEM_PMS2_ADDR` | 0x60003148 | 32 | S=0x3FFFFFF | - |
| `SPI0_SPI_SMEM_PMS2_ATTR` | 0x60003138 | 32 | SPI_SMEM_PMS_ECC=0x4 SPI_SMEM_PMS_RD_ATTR=0x1 SPI_SMEM_PMS_WR_ATTR=0x2 | - |
| `SPI0_SPI_SMEM_PMS2_SIZE` | 0x60003158 | 32 | SPI_SMEM_PMS_SIZE=0x3FFF | - |
| `SPI0_SPI_SMEM_PMS3_ADDR` | 0x6000314C | 32 | S=0x3FFFFFF | - |
| `SPI0_SPI_SMEM_PMS3_ATTR` | 0x6000313C | 32 | SPI_SMEM_PMS_ECC=0x4 SPI_SMEM_PMS_RD_ATTR=0x1 SPI_SMEM_PMS_WR_ATTR=0x2 | - |
| `SPI0_SPI_SMEM_PMS3_SIZE` | 0x6000315C | 32 | SPI_SMEM_PMS_SIZE=0x3FFF | - |
| `SPI0_SPI_SMEM_TIMING_CALI` | 0x60003190 | 32 | SPI_SMEM_DLL_TIMING_CALI=0x20 SPI_SMEM_EXTRA_DUMMY_CYCLELEN=0x1C SPI_SMEM_TIMING_CALI=0x2 SPI_SMEM_TIMING_CLK_ENA=0x1 | - |
| `SPI1_SPI_MEM_ADDR` | 0x60002004 | 32 | SPI_MEM_USR_ADDR_VALUE=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_CACHE_FCTRL` | 0x6000203C | 32 | SPI_MEM_CACHE_USR_ADDR_4BYTE=0x2 SPI_MEM_FADDR_DUAL=0x20 SPI_MEM_FADDR_QUAD=0x100 SPI_MEM_FDIN_DUAL=0x8 SPI_MEM_FDIN_QUAD=0x40 SPI_MEM_FDOUT_DUAL=0x10 SPI_MEM_FDOUT_QUAD=0x80 | - |
| `SPI1_SPI_MEM_CLOCK` | 0x60002014 | 32 | SPI_MEM_CLKCNT_H=0xFF00 SPI_MEM_CLKCNT_L=0xFF SPI_MEM_CLKCNT_N=0xFF0000 SPI_MEM_CLK_EQU_SYSCLK=0x80000000 | - |
| `SPI1_SPI_MEM_CLOCK_GATE` | 0x60002200 | 32 | SPI_MEM_CLK_EN=0x1 | - |
| `SPI1_SPI_MEM_CMD` | 0x60002000 | 32 | SPI_MEM_FLASH_BE=0x800000 SPI_MEM_FLASH_CE=0x400000 SPI_MEM_FLASH_DP=0x200000 SPI_MEM_FLASH_HPM=0x80000 SPI_MEM_FLASH_PE=0x20000 SPI_MEM_FLASH_PP=0x2000000 SPI_MEM_FLASH_RDID=0x10000000 SPI_MEM_FLASH_RDSR=0x8000000 SPI_MEM_FLASH_READ=0x8000... | - |
| `SPI1_SPI_MEM_CTRL` | 0x60002008 | 32 | SPI_MEM_D_POL=0x80000 SPI_MEM_FADDR_OCT=0x40 SPI_MEM_FASTRD_MODE=0x2000 SPI_MEM_FCMD_OCT=0x200 SPI_MEM_FCMD_QUAD=0x100 SPI_MEM_FCS_CRC_EN=0x400 SPI_MEM_FDIN_OCT=0x20 SPI_MEM_FDOUT_OCT=0x10 SPI_MEM_FDUMMY_RIN=0x4 SPI_MEM_FDUMMY_WOUT=0x8 SPI_... | - |
| `SPI1_SPI_MEM_CTRL1` | 0x6000200C | 32 | SPI_MEM_CLK_MODE=0x3 SPI_MEM_CS_HOLD_DLY_RES=0xFFC | - |
| `SPI1_SPI_MEM_CTRL2` | 0x60002010 | 32 | SPI_MEM_SYNC_RESET=0x80000000 | - |
| `SPI1_SPI_MEM_DATE` | 0x600023FC | 32 | SPI_MEM_DATE=0xFFFFFFF | - |
| `SPI1_SPI_MEM_DDR` | 0x600020D4 | 32 | SPI_FMEM_CLK_DIFF_EN=0x1000000 SPI_FMEM_CLK_DIFF_INV=0x10000000 SPI_FMEM_DDR_CMD_DIS=0x10 SPI_FMEM_DDR_DQS_LOOP=0x200000 SPI_FMEM_DDR_EN=0x1 SPI_FMEM_DDR_RDAT_SWP=0x4 SPI_FMEM_DDR_WDAT_SWP=0x8 SPI_FMEM_DQS_CA_IN=0x4000000 SPI_FMEM_HYPERBUS_... | - |
| `SPI1_SPI_MEM_FLASH_SUS_CMD` | 0x600020A0 | 32 | SPI_MEM_FLASH_PES_COMMAND=0xFFFF SPI_MEM_WAIT_PESR_COMMAND=0xFFFF0000 | - |
| `SPI1_SPI_MEM_FLASH_SUS_CTRL` | 0x6000209C | 32 | SPI_FMEM_RD_SUS_2B=0x400000 SPI_MEM_FLASH_PER=0x1 SPI_MEM_FLASH_PER_WAIT_EN=0x4 SPI_MEM_FLASH_PES=0x2 SPI_MEM_FLASH_PES_EN=0x20 SPI_MEM_FLASH_PES_WAIT_EN=0x8 SPI_MEM_PER_END_EN=0x800000 SPI_MEM_PESR_END_MSK=0x3FFFC0 SPI_MEM_PES_END_EN=0x100... | - |
| `SPI1_SPI_MEM_FLASH_WAITI_CTRL` | 0x60002098 | 32 | SPI_MEM_WAITI_ADDR_CYCLELEN=0x18 SPI_MEM_WAITI_ADDR_EN=0x4 SPI_MEM_WAITI_CMD=0xFFFF0000 SPI_MEM_WAITI_CMD_2B=0x200 SPI_MEM_WAITI_DUMMY=0x2 SPI_MEM_WAITI_DUMMY_CYCLELEN=0xFC00 SPI_MEM_WAITI_EN=0x1 | - |
| `SPI1_SPI_MEM_INT_CLR` | 0x600020C4 | 32 | SPI_MEM_BROWN_OUT_INT_CLR=0x400 SPI_MEM_MST_ST_END_INT_CLR=0x10 SPI_MEM_PER_END_INT_CLR=0x1 SPI_MEM_PES_END_INT_CLR=0x2 SPI_MEM_SLV_ST_END_INT_CLR=0x8 SPI_MEM_WPE_END_INT_CLR=0x4 | - |
| `SPI1_SPI_MEM_INT_ENA` | 0x600020C0 | 32 | SPI_MEM_BROWN_OUT_INT_ENA=0x400 SPI_MEM_MST_ST_END_INT_ENA=0x10 SPI_MEM_PER_END_INT_ENA=0x1 SPI_MEM_PES_END_INT_ENA=0x2 SPI_MEM_SLV_ST_END_INT_ENA=0x8 SPI_MEM_WPE_END_INT_ENA=0x4 | - |
| `SPI1_SPI_MEM_INT_RAW` | 0x600020C8 | 32 | SPI_MEM_BROWN_OUT_INT_RAW=0x400 SPI_MEM_MST_ST_END_INT_RAW=0x10 SPI_MEM_PER_END_INT_RAW=0x1 SPI_MEM_PES_END_INT_RAW=0x2 SPI_MEM_SLV_ST_END_INT_RAW=0x8 SPI_MEM_WPE_END_INT_RAW=0x4 | - |
| `SPI1_SPI_MEM_INT_ST` | 0x600020CC | 32 | SPI_MEM_BROWN_OUT_INT_ST=0x400 SPI_MEM_MST_ST_END_INT_ST=0x10 SPI_MEM_PER_END_INT_ST=0x1 SPI_MEM_PES_END_INT_ST=0x2 SPI_MEM_SLV_ST_END_INT_ST=0x8 SPI_MEM_WPE_END_INT_ST=0x4 | - |
| `SPI1_SPI_MEM_MISC` | 0x60002034 | 32 | SPI_MEM_CK_IDLE_EDGE=0x200 SPI_MEM_CS0_DIS=0x1 SPI_MEM_CS1_DIS=0x2 SPI_MEM_CS_KEEP_ACTIVE=0x400 | - |
| `SPI1_SPI_MEM_MISO_DLEN` | 0x60002028 | 32 | SPI_MEM_USR_MISO_DBITLEN=0x3FF | - |
| `SPI1_SPI_MEM_MOSI_DLEN` | 0x60002024 | 32 | SPI_MEM_USR_MOSI_DBITLEN=0x3FF | - |
| `SPI1_SPI_MEM_RD_STATUS` | 0x6000202C | 32 | SPI_MEM_STATUS=0xFFFF SPI_MEM_WB_MODE=0xFF0000 | - |
| `SPI1_SPI_MEM_SUS_STATUS` | 0x600020A4 | 32 | SPI_MEM_FLASH_DP_DLY_128=0x10 SPI_MEM_FLASH_HPM_DLY_128=0x4 SPI_MEM_FLASH_PER_COMMAND=0xFFFF0000 SPI_MEM_FLASH_PER_DLY_128=0x20 SPI_MEM_FLASH_PESR_CMD_2B=0x8000 SPI_MEM_FLASH_PES_DLY_128=0x40 SPI_MEM_FLASH_RES_DLY_128=0x8 SPI_MEM_FLASH_SUS=... | - |
| `SPI1_SPI_MEM_TIMING_CALI` | 0x60002180 | 32 | SPI_MEM_EXTRA_DUMMY_CYCLELEN=0x1C SPI_MEM_TIMING_CALI=0x2 | - |
| `SPI1_SPI_MEM_TX_CRC` | 0x60002038 | 32 | DATA=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_USER` | 0x60002018 | 32 | SPI_MEM_CK_OUT_EDGE=0x200 SPI_MEM_FWRITE_DIO=0x4000 SPI_MEM_FWRITE_DUAL=0x1000 SPI_MEM_FWRITE_QIO=0x8000 SPI_MEM_FWRITE_QUAD=0x2000 SPI_MEM_USR_ADDR=0x40000000 SPI_MEM_USR_COMMAND=0x80000000 SPI_MEM_USR_DUMMY=0x20000000 SPI_MEM_USR_DUMMY_ID... | - |
| `SPI1_SPI_MEM_USER1` | 0x6000201C | 32 | SPI_MEM_USR_ADDR_BITLEN=0xFC000000 SPI_MEM_USR_DUMMY_CYCLELEN=0x3F | - |
| `SPI1_SPI_MEM_USER2` | 0x60002020 | 32 | SPI_MEM_USR_COMMAND_BITLEN=0xF0000000 SPI_MEM_USR_COMMAND_VALUE=0xFFFF | - |
| `SPI1_SPI_MEM_W0` | 0x60002058 | 32 | SPI_MEM_BUF0=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W1` | 0x6000205C | 32 | SPI_MEM_BUF1=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W10` | 0x60002080 | 32 | SPI_MEM_BUF10=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W11` | 0x60002084 | 32 | SPI_MEM_BUF11=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W12` | 0x60002088 | 32 | SPI_MEM_BUF12=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W13` | 0x6000208C | 32 | SPI_MEM_BUF13=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W14` | 0x60002090 | 32 | SPI_MEM_BUF14=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W15` | 0x60002094 | 32 | SPI_MEM_BUF15=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W2` | 0x60002060 | 32 | SPI_MEM_BUF2=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W3` | 0x60002064 | 32 | SPI_MEM_BUF3=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W4` | 0x60002068 | 32 | SPI_MEM_BUF4=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W5` | 0x6000206C | 32 | SPI_MEM_BUF5=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W6` | 0x60002070 | 32 | SPI_MEM_BUF6=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W7` | 0x60002074 | 32 | SPI_MEM_BUF7=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W8` | 0x60002078 | 32 | SPI_MEM_BUF8=0xFFFFFFFF | - |
| `SPI1_SPI_MEM_W9` | 0x6000207C | 32 | SPI_MEM_BUF9=0xFFFFFFFF | - |
| `SPI2_SPI_ADDR` | 0x60024004 | 32 | SPI_USR_ADDR_VALUE=0xFFFFFFFF | - |
| `SPI2_SPI_CLK_GATE` | 0x600240E8 | 32 | SPI_CLK_EN=0x1 SPI_MST_CLK_ACTIVE=0x2 SPI_MST_CLK_SEL=0x4 | - |
| `SPI2_SPI_CLOCK` | 0x6002400C | 32 | SPI_CLKCNT_H=0xFC0 SPI_CLKCNT_L=0x3F SPI_CLKCNT_N=0x3F000 SPI_CLKDIV_PRE=0x3C0000 SPI_CLK_EQU_SYSCLK=0x80000000 | - |
| `SPI2_SPI_CMD` | 0x60024000 | 32 | SPI_CONF_BITLEN=0x3FFFF SPI_UPDATE=0x800000 SPI_USR=0x1000000 | - |
| `SPI2_SPI_CTRL` | 0x60024008 | 32 | SPI_DUMMY_OUT=0x8 SPI_D_POL=0x80000 SPI_FADDR_DUAL=0x20 SPI_FADDR_OCT=0x80 SPI_FADDR_QUAD=0x40 SPI_FCMD_DUAL=0x100 SPI_FCMD_OCT=0x400 SPI_FCMD_QUAD=0x200 SPI_FREAD_DUAL=0x4000 SPI_FREAD_OCT=0x10000 SPI_FREAD_QUAD=0x8000 SPI_HOLD_POL=0x10000... | - |
| `SPI2_SPI_DATE` | 0x600240F0 | 32 | SPI_DATE=0xFFFFFFF | - |
| `SPI2_SPI_DIN_MODE` | 0x60024024 | 32 | SPI_DIN0_MODE=0x3 SPI_DIN1_MODE=0xC SPI_DIN2_MODE=0x30 SPI_DIN3_MODE=0xC0 SPI_DIN4_MODE=0x300 SPI_DIN5_MODE=0xC00 SPI_DIN6_MODE=0x3000 SPI_DIN7_MODE=0xC000 SPI_TIMING_HCLK_ACTIVE=0x10000 | - |
| `SPI2_SPI_DIN_NUM` | 0x60024028 | 32 | SPI_DIN0_NUM=0x3 SPI_DIN1_NUM=0xC SPI_DIN2_NUM=0x30 SPI_DIN3_NUM=0xC0 SPI_DIN4_NUM=0x300 SPI_DIN5_NUM=0xC00 SPI_DIN6_NUM=0x3000 SPI_DIN7_NUM=0xC000 | - |
| `SPI2_SPI_DMA_CONF` | 0x60024030 | 32 | SPI_BUF_AFIFO_RST=0x40000000 SPI_DMA_AFIFO_RST=0x80000000 SPI_DMA_INFIFO_FULL=0x2 SPI_DMA_OUTFIFO_EMPTY=0x1 SPI_DMA_RX_ENA=0x8000000 SPI_DMA_SLV_SEG_TRANS_EN=0x40000 SPI_DMA_TX_ENA=0x10000000 SPI_RX_AFIFO_RST=0x20000000 SPI_RX_EOF_EN=0x2000... | - |
| `SPI2_SPI_DMA_INT_CLR` | 0x60024038 | 32 | SPI_APP1_INT_CLR=0x100000 SPI_APP2_INT_CLR=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_CLR=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_CLR=0x2 SPI_DMA_SEG_TRANS_DONE_INT_CLR=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_CLR=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_CL... | - |
| `SPI2_SPI_DMA_INT_ENA` | 0x60024034 | 32 | SPI_APP1_INT_ENA=0x100000 SPI_APP2_INT_ENA=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_ENA=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_ENA=0x2 SPI_DMA_SEG_TRANS_DONE_INT_ENA=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_ENA=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_EN... | - |
| `SPI2_SPI_DMA_INT_RAW` | 0x6002403C | 32 | SPI_APP1_INT_RAW=0x100000 SPI_APP2_INT_RAW=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_RAW=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_RAW=0x2 SPI_DMA_SEG_TRANS_DONE_INT_RAW=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_RAW=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_RA... | - |
| `SPI2_SPI_DMA_INT_SET` | 0x60024044 | 32 | SPI_APP1_INT_SET=0x100000 SPI_APP2_INT_SET=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_SET=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_SET=0x2 SPI_DMA_SEG_TRANS_DONE_INT_SET=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_SET=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_SE... | - |
| `SPI2_SPI_DMA_INT_ST` | 0x60024040 | 32 | SPI_APP1_INT_ST=0x100000 SPI_APP2_INT_ST=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_ST=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_ST=0x2 SPI_DMA_SEG_TRANS_DONE_INT_ST=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_ST=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_ST=0x400... | - |
| `SPI2_SPI_DOUT_MODE` | 0x6002402C | 32 | SPI_DOUT0_MODE=0x1 SPI_DOUT1_MODE=0x2 SPI_DOUT2_MODE=0x4 SPI_DOUT3_MODE=0x8 SPI_DOUT4_MODE=0x10 SPI_DOUT5_MODE=0x20 SPI_DOUT6_MODE=0x40 SPI_DOUT7_MODE=0x80 SPI_D_DQS_MODE=0x100 | - |
| `SPI2_SPI_MISC` | 0x60024020 | 32 | SPI_ADDR_DTR_EN=0x40000 SPI_CK_DIS=0x40 SPI_CK_IDLE_EDGE=0x20000000 SPI_CLK_DATA_DTR_EN=0x10000 SPI_CMD_DTR_EN=0x80000 SPI_CS0_DIS=0x1 SPI_CS1_DIS=0x2 SPI_CS2_DIS=0x4 SPI_CS3_DIS=0x8 SPI_CS4_DIS=0x10 SPI_CS5_DIS=0x20 SPI_CS_KEEP_ACTIVE=0x40... | - |
| `SPI2_SPI_MS_DLEN` | 0x6002401C | 32 | SPI_MS_DATA_BITLEN=0x3FFFF | - |
| `SPI2_SPI_SLAVE` | 0x600240E0 | 32 | MODE=0x4000000 SPI_CLK_MODE=0x3 SPI_CLK_MODE_13=0x4 SPI_DMA_SEG_MAGIC_VALUE=0x3C00000 SPI_MST_FD_WAIT_DMA_TX_DATA=0x20000000 SPI_RSCK_DATA_OUT=0x8 SPI_SLV_RDBUF_BITLEN_EN=0x400 SPI_SLV_RDDMA_BITLEN_EN=0x100 SPI_SLV_WRBUF_BITLEN_EN=0x800 SPI... | - |
| `SPI2_SPI_SLAVE1` | 0x600240E4 | 32 | SPI_SLV_DATA_BITLEN=0x3FFFF SPI_SLV_LAST_ADDR=0xFC000000 SPI_SLV_LAST_COMMAND=0x3FC0000 | - |
| `SPI2_SPI_USER` | 0x60024010 | 32 | SPI_CK_OUT_EDGE=0x200 SPI_CS_HOLD=0x40 SPI_CS_SETUP=0x80 SPI_DOUTDIN=0x1 SPI_FWRITE_DUAL=0x1000 SPI_FWRITE_OCT=0x4000 SPI_FWRITE_QUAD=0x2000 SPI_OPI_MODE=0x10 SPI_QPI_MODE=0x8 SPI_RSCK_I_EDGE=0x100 SPI_SIO=0x20000 SPI_TSCK_I_EDGE=0x20 SPI_U... | - |
| `SPI2_SPI_USER1` | 0x60024014 | 32 | SPI_CS_HOLD_TIME=0x7C00000 SPI_CS_SETUP_TIME=0x3E0000 SPI_MST_WFULL_ERR_END_EN=0x10000 SPI_USR_ADDR_BITLEN=0xF8000000 SPI_USR_DUMMY_CYCLELEN=0xFF | - |
| `SPI2_SPI_USER2` | 0x60024018 | 32 | SPI_MST_REMPTY_ERR_END_EN=0x8000000 SPI_USR_COMMAND_BITLEN=0xF0000000 SPI_USR_COMMAND_VALUE=0xFFFF | - |
| `SPI2_SPI_W0` | 0x60024098 | 32 | SPI_BUF0=0xFFFFFFFF | - |
| `SPI2_SPI_W1` | 0x6002409C | 32 | SPI_BUF1=0xFFFFFFFF | - |
| `SPI2_SPI_W10` | 0x600240C0 | 32 | SPI_BUF10=0xFFFFFFFF | - |
| `SPI2_SPI_W11` | 0x600240C4 | 32 | SPI_BUF11=0xFFFFFFFF | - |
| `SPI2_SPI_W12` | 0x600240C8 | 32 | SPI_BUF12=0xFFFFFFFF | - |
| `SPI2_SPI_W13` | 0x600240CC | 32 | SPI_BUF13=0xFFFFFFFF | - |
| `SPI2_SPI_W14` | 0x600240D0 | 32 | SPI_BUF14=0xFFFFFFFF | - |
| `SPI2_SPI_W15` | 0x600240D4 | 32 | SPI_BUF15=0xFFFFFFFF | - |
| `SPI2_SPI_W2` | 0x600240A0 | 32 | SPI_BUF2=0xFFFFFFFF | - |
| `SPI2_SPI_W3` | 0x600240A4 | 32 | SPI_BUF3=0xFFFFFFFF | - |
| `SPI2_SPI_W4` | 0x600240A8 | 32 | SPI_BUF4=0xFFFFFFFF | - |
| `SPI2_SPI_W5` | 0x600240AC | 32 | SPI_BUF5=0xFFFFFFFF | - |
| `SPI2_SPI_W6` | 0x600240B0 | 32 | SPI_BUF6=0xFFFFFFFF | - |
| `SPI2_SPI_W7` | 0x600240B4 | 32 | SPI_BUF7=0xFFFFFFFF | - |
| `SPI2_SPI_W8` | 0x600240B8 | 32 | SPI_BUF8=0xFFFFFFFF | - |
| `SPI2_SPI_W9` | 0x600240BC | 32 | SPI_BUF9=0xFFFFFFFF | - |
| `SPI3_SPI_ADDR` | 0x60025004 | 32 | SPI_USR_ADDR_VALUE=0xFFFFFFFF | - |
| `SPI3_SPI_CLK_GATE` | 0x600250E8 | 32 | SPI_CLK_EN=0x1 SPI_MST_CLK_ACTIVE=0x2 SPI_MST_CLK_SEL=0x4 | - |
| `SPI3_SPI_CLOCK` | 0x6002500C | 32 | SPI_CLKCNT_H=0xFC0 SPI_CLKCNT_L=0x3F SPI_CLKCNT_N=0x3F000 SPI_CLKDIV_PRE=0x3C0000 SPI_CLK_EQU_SYSCLK=0x80000000 | - |
| `SPI3_SPI_CMD` | 0x60025000 | 32 | SPI_CONF_BITLEN=0x3FFFF SPI_UPDATE=0x800000 SPI_USR=0x1000000 | - |
| `SPI3_SPI_CTRL` | 0x60025008 | 32 | SPI_DUMMY_OUT=0x8 SPI_D_POL=0x80000 SPI_FADDR_DUAL=0x20 SPI_FADDR_OCT=0x80 SPI_FADDR_QUAD=0x40 SPI_FCMD_DUAL=0x100 SPI_FCMD_OCT=0x400 SPI_FCMD_QUAD=0x200 SPI_FREAD_DUAL=0x4000 SPI_FREAD_OCT=0x10000 SPI_FREAD_QUAD=0x8000 SPI_HOLD_POL=0x10000... | - |
| `SPI3_SPI_DATE` | 0x600250F0 | 32 | SPI_DATE=0xFFFFFFF | - |
| `SPI3_SPI_DIN_MODE` | 0x60025024 | 32 | SPI_DIN0_MODE=0x3 SPI_DIN1_MODE=0xC SPI_DIN2_MODE=0x30 SPI_DIN3_MODE=0xC0 SPI_DIN4_MODE=0x300 SPI_DIN5_MODE=0xC00 SPI_DIN6_MODE=0x3000 SPI_DIN7_MODE=0xC000 SPI_TIMING_HCLK_ACTIVE=0x10000 | - |
| `SPI3_SPI_DIN_NUM` | 0x60025028 | 32 | SPI_DIN0_NUM=0x3 SPI_DIN1_NUM=0xC SPI_DIN2_NUM=0x30 SPI_DIN3_NUM=0xC0 SPI_DIN4_NUM=0x300 SPI_DIN5_NUM=0xC00 SPI_DIN6_NUM=0x3000 SPI_DIN7_NUM=0xC000 | - |
| `SPI3_SPI_DMA_CONF` | 0x60025030 | 32 | SPI_BUF_AFIFO_RST=0x40000000 SPI_DMA_AFIFO_RST=0x80000000 SPI_DMA_INFIFO_FULL=0x2 SPI_DMA_OUTFIFO_EMPTY=0x1 SPI_DMA_RX_ENA=0x8000000 SPI_DMA_SLV_SEG_TRANS_EN=0x40000 SPI_DMA_TX_ENA=0x10000000 SPI_RX_AFIFO_RST=0x20000000 SPI_RX_EOF_EN=0x2000... | - |
| `SPI3_SPI_DMA_INT_CLR` | 0x60025038 | 32 | SPI_APP1_INT_CLR=0x100000 SPI_APP2_INT_CLR=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_CLR=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_CLR=0x2 SPI_DMA_SEG_TRANS_DONE_INT_CLR=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_CLR=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_CL... | - |
| `SPI3_SPI_DMA_INT_ENA` | 0x60025034 | 32 | SPI_APP1_INT_ENA=0x100000 SPI_APP2_INT_ENA=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_ENA=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_ENA=0x2 SPI_DMA_SEG_TRANS_DONE_INT_ENA=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_ENA=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_EN... | - |
| `SPI3_SPI_DMA_INT_RAW` | 0x6002503C | 32 | SPI_APP1_INT_RAW=0x100000 SPI_APP2_INT_RAW=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_RAW=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_RAW=0x2 SPI_DMA_SEG_TRANS_DONE_INT_RAW=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_RAW=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_RA... | - |
| `SPI3_SPI_DMA_INT_SET` | 0x60025044 | 32 | SPI_APP1_INT_SET=0x100000 SPI_APP2_INT_SET=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_SET=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_SET=0x2 SPI_DMA_SEG_TRANS_DONE_INT_SET=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_SET=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_SE... | - |
| `SPI3_SPI_DMA_INT_ST` | 0x60025040 | 32 | SPI_APP1_INT_ST=0x100000 SPI_APP2_INT_ST=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_ST=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_ST=0x2 SPI_DMA_SEG_TRANS_DONE_INT_ST=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_ST=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_ST=0x400... | - |
| `SPI3_SPI_DOUT_MODE` | 0x6002502C | 32 | SPI_DOUT0_MODE=0x1 SPI_DOUT1_MODE=0x2 SPI_DOUT2_MODE=0x4 SPI_DOUT3_MODE=0x8 SPI_DOUT4_MODE=0x10 SPI_DOUT5_MODE=0x20 SPI_DOUT6_MODE=0x40 SPI_DOUT7_MODE=0x80 SPI_D_DQS_MODE=0x100 | - |
| `SPI3_SPI_MISC` | 0x60025020 | 32 | SPI_ADDR_DTR_EN=0x40000 SPI_CK_DIS=0x40 SPI_CK_IDLE_EDGE=0x20000000 SPI_CLK_DATA_DTR_EN=0x10000 SPI_CMD_DTR_EN=0x80000 SPI_CS0_DIS=0x1 SPI_CS1_DIS=0x2 SPI_CS2_DIS=0x4 SPI_CS3_DIS=0x8 SPI_CS4_DIS=0x10 SPI_CS5_DIS=0x20 SPI_CS_KEEP_ACTIVE=0x40... | - |
| `SPI3_SPI_MS_DLEN` | 0x6002501C | 32 | SPI_MS_DATA_BITLEN=0x3FFFF | - |
| `SPI3_SPI_SLAVE` | 0x600250E0 | 32 | MODE=0x4000000 SPI_CLK_MODE=0x3 SPI_CLK_MODE_13=0x4 SPI_DMA_SEG_MAGIC_VALUE=0x3C00000 SPI_MST_FD_WAIT_DMA_TX_DATA=0x20000000 SPI_RSCK_DATA_OUT=0x8 SPI_SLV_RDBUF_BITLEN_EN=0x400 SPI_SLV_RDDMA_BITLEN_EN=0x100 SPI_SLV_WRBUF_BITLEN_EN=0x800 SPI... | - |
| `SPI3_SPI_SLAVE1` | 0x600250E4 | 32 | SPI_SLV_DATA_BITLEN=0x3FFFF SPI_SLV_LAST_ADDR=0xFC000000 SPI_SLV_LAST_COMMAND=0x3FC0000 | - |
| `SPI3_SPI_USER` | 0x60025010 | 32 | SPI_CK_OUT_EDGE=0x200 SPI_CS_HOLD=0x40 SPI_CS_SETUP=0x80 SPI_DOUTDIN=0x1 SPI_FWRITE_DUAL=0x1000 SPI_FWRITE_OCT=0x4000 SPI_FWRITE_QUAD=0x2000 SPI_OPI_MODE=0x10 SPI_QPI_MODE=0x8 SPI_RSCK_I_EDGE=0x100 SPI_SIO=0x20000 SPI_TSCK_I_EDGE=0x20 SPI_U... | - |
| `SPI3_SPI_USER1` | 0x60025014 | 32 | SPI_CS_HOLD_TIME=0x7C00000 SPI_CS_SETUP_TIME=0x3E0000 SPI_MST_WFULL_ERR_END_EN=0x10000 SPI_USR_ADDR_BITLEN=0xF8000000 SPI_USR_DUMMY_CYCLELEN=0xFF | - |
| `SPI3_SPI_USER2` | 0x60025018 | 32 | SPI_MST_REMPTY_ERR_END_EN=0x8000000 SPI_USR_COMMAND_BITLEN=0xF0000000 SPI_USR_COMMAND_VALUE=0xFFFF | - |
| `SPI3_SPI_W0` | 0x60025098 | 32 | SPI_BUF0=0xFFFFFFFF | - |
| `SPI3_SPI_W1` | 0x6002509C | 32 | SPI_BUF1=0xFFFFFFFF | - |
| `SPI3_SPI_W10` | 0x600250C0 | 32 | SPI_BUF10=0xFFFFFFFF | - |
| `SPI3_SPI_W11` | 0x600250C4 | 32 | SPI_BUF11=0xFFFFFFFF | - |
| `SPI3_SPI_W12` | 0x600250C8 | 32 | SPI_BUF12=0xFFFFFFFF | - |
| `SPI3_SPI_W13` | 0x600250CC | 32 | SPI_BUF13=0xFFFFFFFF | - |
| `SPI3_SPI_W14` | 0x600250D0 | 32 | SPI_BUF14=0xFFFFFFFF | - |
| `SPI3_SPI_W15` | 0x600250D4 | 32 | SPI_BUF15=0xFFFFFFFF | - |
| `SPI3_SPI_W2` | 0x600250A0 | 32 | SPI_BUF2=0xFFFFFFFF | - |
| `SPI3_SPI_W3` | 0x600250A4 | 32 | SPI_BUF3=0xFFFFFFFF | - |
| `SPI3_SPI_W4` | 0x600250A8 | 32 | SPI_BUF4=0xFFFFFFFF | - |
| `SPI3_SPI_W5` | 0x600250AC | 32 | SPI_BUF5=0xFFFFFFFF | - |
| `SPI3_SPI_W6` | 0x600250B0 | 32 | SPI_BUF6=0xFFFFFFFF | - |
| `SPI3_SPI_W7` | 0x600250B4 | 32 | SPI_BUF7=0xFFFFFFFF | - |
| `SPI3_SPI_W8` | 0x600250B8 | 32 | SPI_BUF8=0xFFFFFFFF | - |
| `SPI3_SPI_W9` | 0x600250BC | 32 | SPI_BUF9=0xFFFFFFFF | - |
| `SPI4_SPI_ADDR` | 0x60037004 | 32 | SPI_USR_ADDR_VALUE=0xFFFFFFFF | - |
| `SPI4_SPI_CLK_GATE` | 0x600370E8 | 32 | SPI_CLK_EN=0x1 SPI_MST_CLK_ACTIVE=0x2 SPI_MST_CLK_SEL=0x4 | - |
| `SPI4_SPI_CLOCK` | 0x6003700C | 32 | SPI_CLKCNT_H=0xFC0 SPI_CLKCNT_L=0x3F SPI_CLKCNT_N=0x3F000 SPI_CLKDIV_PRE=0x3C0000 SPI_CLK_EQU_SYSCLK=0x80000000 | - |
| `SPI4_SPI_CMD` | 0x60037000 | 32 | SPI_CONF_BITLEN=0x3FFFF SPI_UPDATE=0x800000 SPI_USR=0x1000000 | - |
| `SPI4_SPI_CTRL` | 0x60037008 | 32 | SPI_DUMMY_OUT=0x8 SPI_D_POL=0x80000 SPI_FADDR_DUAL=0x20 SPI_FADDR_OCT=0x80 SPI_FADDR_QUAD=0x40 SPI_FCMD_DUAL=0x100 SPI_FCMD_OCT=0x400 SPI_FCMD_QUAD=0x200 SPI_FREAD_DUAL=0x4000 SPI_FREAD_OCT=0x10000 SPI_FREAD_QUAD=0x8000 SPI_HOLD_POL=0x10000... | - |
| `SPI4_SPI_DATE` | 0x600370F0 | 32 | SPI_DATE=0xFFFFFFF | - |
| `SPI4_SPI_DIN_MODE` | 0x60037024 | 32 | SPI_DIN0_MODE=0x3 SPI_DIN1_MODE=0xC SPI_DIN2_MODE=0x30 SPI_DIN3_MODE=0xC0 SPI_DIN4_MODE=0x300 SPI_DIN5_MODE=0xC00 SPI_DIN6_MODE=0x3000 SPI_DIN7_MODE=0xC000 SPI_TIMING_HCLK_ACTIVE=0x10000 | - |
| `SPI4_SPI_DIN_NUM` | 0x60037028 | 32 | SPI_DIN0_NUM=0x3 SPI_DIN1_NUM=0xC SPI_DIN2_NUM=0x30 SPI_DIN3_NUM=0xC0 SPI_DIN4_NUM=0x300 SPI_DIN5_NUM=0xC00 SPI_DIN6_NUM=0x3000 SPI_DIN7_NUM=0xC000 | - |
| `SPI4_SPI_DMA_CONF` | 0x60037030 | 32 | SPI_BUF_AFIFO_RST=0x40000000 SPI_DMA_AFIFO_RST=0x80000000 SPI_DMA_INFIFO_FULL=0x2 SPI_DMA_OUTFIFO_EMPTY=0x1 SPI_DMA_RX_ENA=0x8000000 SPI_DMA_SLV_SEG_TRANS_EN=0x40000 SPI_DMA_TX_ENA=0x10000000 SPI_RX_AFIFO_RST=0x20000000 SPI_RX_EOF_EN=0x2000... | - |
| `SPI4_SPI_DMA_INT_CLR` | 0x60037038 | 32 | SPI_APP1_INT_CLR=0x100000 SPI_APP2_INT_CLR=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_CLR=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_CLR=0x2 SPI_DMA_SEG_TRANS_DONE_INT_CLR=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_CLR=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_CL... | - |
| `SPI4_SPI_DMA_INT_ENA` | 0x60037034 | 32 | SPI_APP1_INT_ENA=0x100000 SPI_APP2_INT_ENA=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_ENA=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_ENA=0x2 SPI_DMA_SEG_TRANS_DONE_INT_ENA=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_ENA=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_EN... | - |
| `SPI4_SPI_DMA_INT_RAW` | 0x6003703C | 32 | SPI_APP1_INT_RAW=0x100000 SPI_APP2_INT_RAW=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_RAW=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_RAW=0x2 SPI_DMA_SEG_TRANS_DONE_INT_RAW=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_RAW=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_RA... | - |
| `SPI4_SPI_DMA_INT_SET` | 0x60037044 | 32 | SPI_APP1_INT_SET=0x100000 SPI_APP2_INT_SET=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_SET=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_SET=0x2 SPI_DMA_SEG_TRANS_DONE_INT_SET=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_SET=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_SE... | - |
| `SPI4_SPI_DMA_INT_ST` | 0x60037040 | 32 | SPI_APP1_INT_ST=0x100000 SPI_APP2_INT_ST=0x80000 SPI_DMA_INFIFO_FULL_ERR_INT_ST=0x1 SPI_DMA_OUTFIFO_EMPTY_ERR_INT_ST=0x2 SPI_DMA_SEG_TRANS_DONE_INT_ST=0x2000 SPI_MST_RX_AFIFO_WFULL_ERR_INT_ST=0x20000 SPI_MST_TX_AFIFO_REMPTY_ERR_INT_ST=0x400... | - |
| `SPI4_SPI_DOUT_MODE` | 0x6003702C | 32 | SPI_DOUT0_MODE=0x1 SPI_DOUT1_MODE=0x2 SPI_DOUT2_MODE=0x4 SPI_DOUT3_MODE=0x8 SPI_DOUT4_MODE=0x10 SPI_DOUT5_MODE=0x20 SPI_DOUT6_MODE=0x40 SPI_DOUT7_MODE=0x80 SPI_D_DQS_MODE=0x100 | - |
| `SPI4_SPI_MISC` | 0x60037020 | 32 | SPI_ADDR_DTR_EN=0x40000 SPI_CK_DIS=0x40 SPI_CK_IDLE_EDGE=0x20000000 SPI_CLK_DATA_DTR_EN=0x10000 SPI_CMD_DTR_EN=0x80000 SPI_CS0_DIS=0x1 SPI_CS1_DIS=0x2 SPI_CS2_DIS=0x4 SPI_CS3_DIS=0x8 SPI_CS4_DIS=0x10 SPI_CS5_DIS=0x20 SPI_CS_KEEP_ACTIVE=0x40... | - |
| `SPI4_SPI_MS_DLEN` | 0x6003701C | 32 | SPI_MS_DATA_BITLEN=0x3FFFF | - |
| `SPI4_SPI_SLAVE` | 0x600370E0 | 32 | MODE=0x4000000 SPI_CLK_MODE=0x3 SPI_CLK_MODE_13=0x4 SPI_DMA_SEG_MAGIC_VALUE=0x3C00000 SPI_MST_FD_WAIT_DMA_TX_DATA=0x20000000 SPI_RSCK_DATA_OUT=0x8 SPI_SLV_RDBUF_BITLEN_EN=0x400 SPI_SLV_RDDMA_BITLEN_EN=0x100 SPI_SLV_WRBUF_BITLEN_EN=0x800 SPI... | - |
| `SPI4_SPI_SLAVE1` | 0x600370E4 | 32 | SPI_SLV_DATA_BITLEN=0x3FFFF SPI_SLV_LAST_ADDR=0xFC000000 SPI_SLV_LAST_COMMAND=0x3FC0000 | - |
| `SPI4_SPI_USER` | 0x60037010 | 32 | SPI_CK_OUT_EDGE=0x200 SPI_CS_HOLD=0x40 SPI_CS_SETUP=0x80 SPI_DOUTDIN=0x1 SPI_FWRITE_DUAL=0x1000 SPI_FWRITE_OCT=0x4000 SPI_FWRITE_QUAD=0x2000 SPI_OPI_MODE=0x10 SPI_QPI_MODE=0x8 SPI_RSCK_I_EDGE=0x100 SPI_SIO=0x20000 SPI_TSCK_I_EDGE=0x20 SPI_U... | - |
| `SPI4_SPI_USER1` | 0x60037014 | 32 | SPI_CS_HOLD_TIME=0x7C00000 SPI_CS_SETUP_TIME=0x3E0000 SPI_MST_WFULL_ERR_END_EN=0x10000 SPI_USR_ADDR_BITLEN=0xF8000000 SPI_USR_DUMMY_CYCLELEN=0xFF | - |
| `SPI4_SPI_USER2` | 0x60037018 | 32 | SPI_MST_REMPTY_ERR_END_EN=0x8000000 SPI_USR_COMMAND_BITLEN=0xF0000000 SPI_USR_COMMAND_VALUE=0xFFFF | - |
| `SPI4_SPI_W0` | 0x60037098 | 32 | SPI_BUF0=0xFFFFFFFF | - |
| `SPI4_SPI_W1` | 0x6003709C | 32 | SPI_BUF1=0xFFFFFFFF | - |
| `SPI4_SPI_W10` | 0x600370C0 | 32 | SPI_BUF10=0xFFFFFFFF | - |
| `SPI4_SPI_W11` | 0x600370C4 | 32 | SPI_BUF11=0xFFFFFFFF | - |
| `SPI4_SPI_W12` | 0x600370C8 | 32 | SPI_BUF12=0xFFFFFFFF | - |
| `SPI4_SPI_W13` | 0x600370CC | 32 | SPI_BUF13=0xFFFFFFFF | - |
| `SPI4_SPI_W14` | 0x600370D0 | 32 | SPI_BUF14=0xFFFFFFFF | - |
| `SPI4_SPI_W15` | 0x600370D4 | 32 | SPI_BUF15=0xFFFFFFFF | - |
| `SPI4_SPI_W2` | 0x600370A0 | 32 | SPI_BUF2=0xFFFFFFFF | - |
| `SPI4_SPI_W3` | 0x600370A4 | 32 | SPI_BUF3=0xFFFFFFFF | - |
| `SPI4_SPI_W4` | 0x600370A8 | 32 | SPI_BUF4=0xFFFFFFFF | - |
| `SPI4_SPI_W5` | 0x600370AC | 32 | SPI_BUF5=0xFFFFFFFF | - |
| `SPI4_SPI_W6` | 0x600370B0 | 32 | SPI_BUF6=0xFFFFFFFF | - |
| `SPI4_SPI_W7` | 0x600370B4 | 32 | SPI_BUF7=0xFFFFFFFF | - |
| `SPI4_SPI_W8` | 0x600370B8 | 32 | SPI_BUF8=0xFFFFFFFF | - |
| `SPI4_SPI_W9` | 0x600370BC | 32 | SPI_BUF9=0xFFFFFFFF | - |
| `SYSTIMER_COMP0_LOAD` | 0x60023050 | 32 | TIMER_COMP0_LOAD=0x1 | - |
| `SYSTIMER_COMP1_LOAD` | 0x60023054 | 32 | TIMER_COMP1_LOAD=0x1 | - |
| `SYSTIMER_COMP2_LOAD` | 0x60023058 | 32 | TIMER_COMP2_LOAD=0x1 | - |
| `SYSTIMER_CONF` | 0x60023000 | 32 | CLK_EN=0x80000000 ETM_EN=0x2 SYSTIMER_CLK_FO=0x1 TARGET0_WORK_EN=0x1000000 TARGET1_WORK_EN=0x800000 TARGET2_WORK_EN=0x400000 TIMER_UNIT0_CORE0_STALL_EN=0x10000000 TIMER_UNIT0_CORE1_STALL_EN=0x8000000 TIMER_UNIT0_WORK_EN=0x40000000 TIMER_UNI... | - |
| `SYSTIMER_DATE` | 0x600230FC | 32 | DATE=0xFFFFFFFF | - |
| `SYSTIMER_INT_CLR` | 0x6002306C | 32 | TARGET0_INT_CLR=0x1 TARGET1_INT_CLR=0x2 TARGET2_INT_CLR=0x4 | - |
| `SYSTIMER_INT_ENA` | 0x60023064 | 32 | TARGET0_INT_ENA=0x1 TARGET1_INT_ENA=0x2 TARGET2_INT_ENA=0x4 | - |
| `SYSTIMER_INT_RAW` | 0x60023068 | 32 | TARGET0_INT_RAW=0x1 TARGET1_INT_RAW=0x2 TARGET2_INT_RAW=0x4 | - |
| `SYSTIMER_INT_ST` | 0x60023070 | 32 | TARGET0_INT_ST=0x1 TARGET1_INT_ST=0x2 TARGET2_INT_ST=0x4 | - |
| `SYSTIMER_REAL_TARGET0_HI` | 0x60023078 | 32 | TARGET0_HI_RO=0xFFFFF | - |
| `SYSTIMER_REAL_TARGET0_LO` | 0x60023074 | 32 | TARGET0_LO_RO=0xFFFFFFFF | - |
| `SYSTIMER_REAL_TARGET1_HI` | 0x60023080 | 32 | TARGET1_HI_RO=0xFFFFF | - |
| `SYSTIMER_REAL_TARGET1_LO` | 0x6002307C | 32 | TARGET1_LO_RO=0xFFFFFFFF | - |
| `SYSTIMER_REAL_TARGET2_HI` | 0x60023088 | 32 | TARGET2_HI_RO=0xFFFFF | - |
| `SYSTIMER_REAL_TARGET2_LO` | 0x60023084 | 32 | TARGET2_LO_RO=0xFFFFFFFF | - |
| `SYSTIMER_TARGET0_CONF` | 0x60023034 | 32 | TARGET0_PERIOD=0x3FFFFFF TARGET0_PERIOD_MODE=0x40000000 TARGET0_TIMER_UNIT_SEL=0x80000000 | - |
| `SYSTIMER_TARGET0_HI` | 0x6002301C | 32 | TIMER_TARGET0_HI=0xFFFFF | - |
| `SYSTIMER_TARGET0_LO` | 0x60023020 | 32 | TIMER_TARGET0_LO=0xFFFFFFFF | - |
| `SYSTIMER_TARGET1_CONF` | 0x60023038 | 32 | TARGET1_PERIOD=0x3FFFFFF TARGET1_PERIOD_MODE=0x40000000 TARGET1_TIMER_UNIT_SEL=0x80000000 | - |
| `SYSTIMER_TARGET1_HI` | 0x60023024 | 32 | TIMER_TARGET1_HI=0xFFFFF | - |
| `SYSTIMER_TARGET1_LO` | 0x60023028 | 32 | TIMER_TARGET1_LO=0xFFFFFFFF | - |
| `SYSTIMER_TARGET2_CONF` | 0x6002303C | 32 | TARGET2_PERIOD=0x3FFFFFF TARGET2_PERIOD_MODE=0x40000000 TARGET2_TIMER_UNIT_SEL=0x80000000 | - |
| `SYSTIMER_TARGET2_HI` | 0x6002302C | 32 | TIMER_TARGET2_HI=0xFFFFF | - |
| `SYSTIMER_TARGET2_LO` | 0x60023030 | 32 | TIMER_TARGET2_LO=0xFFFFFFFF | - |
| `SYSTIMER_UNIT0_LOAD` | 0x6002305C | 32 | TIMER_UNIT0_LOAD=0x1 | - |
| `SYSTIMER_UNIT0_LOAD_HI` | 0x6002300C | 32 | TIMER_UNIT0_LOAD_HI=0xFFFFF | - |
| `SYSTIMER_UNIT0_LOAD_LO` | 0x60023010 | 32 | TIMER_UNIT0_LOAD_LO=0xFFFFFFFF | - |
| `SYSTIMER_UNIT0_OP` | 0x60023004 | 32 | TIMER_UNIT0_UPDATE=0x40000000 TIMER_UNIT0_VALUE_VALID=0x20000000 | - |
| `SYSTIMER_UNIT0_VALUE_HI` | 0x60023040 | 32 | TIMER_UNIT0_VALUE_HI=0xFFFFF | - |
| `SYSTIMER_UNIT0_VALUE_LO` | 0x60023044 | 32 | TIMER_UNIT0_VALUE_LO=0xFFFFFFFF | - |
| `SYSTIMER_UNIT1_LOAD` | 0x60023060 | 32 | TIMER_UNIT1_LOAD=0x1 | - |
| `SYSTIMER_UNIT1_LOAD_HI` | 0x60023014 | 32 | TIMER_UNIT1_LOAD_HI=0xFFFFF | - |
| `SYSTIMER_UNIT1_LOAD_LO` | 0x60023018 | 32 | TIMER_UNIT1_LOAD_LO=0xFFFFFFFF | - |
| `SYSTIMER_UNIT1_OP` | 0x60023008 | 32 | TIMER_UNIT1_UPDATE=0x40000000 TIMER_UNIT1_VALUE_VALID=0x20000000 | - |
| `SYSTIMER_UNIT1_VALUE_HI` | 0x60023048 | 32 | TIMER_UNIT1_VALUE_HI=0xFFFFF | - |
| `SYSTIMER_UNIT1_VALUE_LO` | 0x6002304C | 32 | TIMER_UNIT1_VALUE_LO=0xFFFFFFFF | - |
| `TEE_CLOCK_GATE` | 0x60098080 | 32 | CLK_EN=0x1 | - |
| `TEE_DATE` | 0x60098FFC | 32 | DATE=0xFFFFFFF | - |
| `TEE_M0_MODE_CTRL` | 0x60098000 | 32 | M0_MODE=0x3 | - |
| `TEE_M10_MODE_CTRL` | 0x60098028 | 32 | M10_MODE=0x3 | - |
| `TEE_M11_MODE_CTRL` | 0x6009802C | 32 | M11_MODE=0x3 | - |
| `TEE_M12_MODE_CTRL` | 0x60098030 | 32 | M12_MODE=0x3 | - |
| `TEE_M13_MODE_CTRL` | 0x60098034 | 32 | M13_MODE=0x3 | - |
| `TEE_M14_MODE_CTRL` | 0x60098038 | 32 | M14_MODE=0x3 | - |
| `TEE_M15_MODE_CTRL` | 0x6009803C | 32 | M15_MODE=0x3 | - |
| `TEE_M16_MODE_CTRL` | 0x60098040 | 32 | M16_MODE=0x3 | - |
| `TEE_M17_MODE_CTRL` | 0x60098044 | 32 | M17_MODE=0x3 | - |
| `TEE_M18_MODE_CTRL` | 0x60098048 | 32 | M18_MODE=0x3 | - |
| `TEE_M19_MODE_CTRL` | 0x6009804C | 32 | M19_MODE=0x3 | - |
| `TEE_M1_MODE_CTRL` | 0x60098004 | 32 | M1_MODE=0x3 | - |
| `TEE_M20_MODE_CTRL` | 0x60098050 | 32 | M20_MODE=0x3 | - |
| `TEE_M21_MODE_CTRL` | 0x60098054 | 32 | M21_MODE=0x3 | - |
| `TEE_M22_MODE_CTRL` | 0x60098058 | 32 | M22_MODE=0x3 | - |
| `TEE_M23_MODE_CTRL` | 0x6009805C | 32 | M23_MODE=0x3 | - |
| `TEE_M24_MODE_CTRL` | 0x60098060 | 32 | M24_MODE=0x3 | - |
| `TEE_M25_MODE_CTRL` | 0x60098064 | 32 | M25_MODE=0x3 | - |
| `TEE_M26_MODE_CTRL` | 0x60098068 | 32 | M26_MODE=0x3 | - |
| `TEE_M27_MODE_CTRL` | 0x6009806C | 32 | M27_MODE=0x3 | - |
| `TEE_M28_MODE_CTRL` | 0x60098070 | 32 | M28_MODE=0x3 | - |
| `TEE_M29_MODE_CTRL` | 0x60098074 | 32 | M29_MODE=0x3 | - |
| `TEE_M2_MODE_CTRL` | 0x60098008 | 32 | M2_MODE=0x3 | - |
| `TEE_M30_MODE_CTRL` | 0x60098078 | 32 | M30_MODE=0x3 | - |
| `TEE_M31_MODE_CTRL` | 0x6009807C | 32 | M31_MODE=0x3 | - |
| `TEE_M3_MODE_CTRL` | 0x6009800C | 32 | M3_MODE=0x3 | - |
| `TEE_M4_MODE_CTRL` | 0x60098010 | 32 | M4_MODE=0x3 | - |
| `TEE_M5_MODE_CTRL` | 0x60098014 | 32 | M5_MODE=0x3 | - |
| `TEE_M6_MODE_CTRL` | 0x60098018 | 32 | M6_MODE=0x3 | - |
| `TEE_M7_MODE_CTRL` | 0x6009801C | 32 | M7_MODE=0x3 | - |
| `TEE_M8_MODE_CTRL` | 0x60098020 | 32 | M8_MODE=0x3 | - |
| `TEE_M9_MODE_CTRL` | 0x60098024 | 32 | M9_MODE=0x3 | - |
| `TIMG0_INT_CLR_TIMERS` | 0x6001F07C | 32 | T0_INT_CLR=0x1 WDT_INT_CLR=0x2 | - |
| `TIMG0_INT_ENA_TIMERS` | 0x6001F070 | 32 | T0_INT_ENA=0x1 WDT_INT_ENA=0x2 | - |
| `TIMG0_INT_RAW_TIMERS` | 0x6001F074 | 32 | T0_INT_RAW=0x1 WDT_INT_RAW=0x2 | - |
| `TIMG0_INT_ST_TIMERS` | 0x6001F078 | 32 | T0_INT_ST=0x1 WDT_INT_ST=0x2 | - |
| `TIMG0_NTIMERS_DATE` | 0x6001F0F8 | 32 | NTIMGS_DATE=0xFFFFFFF | - |
| `TIMG0_REGCLK` | 0x6001F0FC | 32 | CLK_EN=0x80000000 ETM_EN=0x10000000 TIMER_CLK_IS_ACTIVE=0x40000000 WDT_CLK_IS_ACTIVE=0x20000000 | - |
| `TIMG0_RTCCALICFG` | 0x6001F068 | 32 | RTC_CALI_CLK_SEL=0x6000 RTC_CALI_MAX=0x7FFF0000 RTC_CALI_RDY=0x8000 RTC_CALI_START=0x80000000 RTC_CALI_START_CYCLING=0x1000 | - |
| `TIMG0_RTCCALICFG1` | 0x6001F06C | 32 | RTC_CALI_CYCLING_DATA_VLD=0x1 RTC_CALI_VALUE=0xFFFFFF80 | - |
| `TIMG0_RTCCALICFG2` | 0x6001F080 | 32 | RTC_CALI_TIMEOUT=0x1 RTC_CALI_TIMEOUT_RST_CNT=0x78 RTC_CALI_TIMEOUT_THRES=0xFFFFFF80 | - |
| `TIMG0_T0ALARMHI` | 0x6001F014 | 32 | T_ALARM_HI=0x3FFFFF | - |
| `TIMG0_T0ALARMLO` | 0x6001F010 | 32 | T_ALARM_LO=0xFFFFFFFF | - |
| `TIMG0_T0CONFIG` | 0x6001F000 | 32 | T_ALARM_EN=0x400 T_AUTORELOAD=0x20000000 T_DIVCNT_RST=0x1000 T_DIVIDER=0x1FFFE000 T_EN=0x80000000 T_INCREASE=0x40000000 T_USE_XTAL=0x200 | - |
| `TIMG0_T0HI` | 0x6001F008 | 32 | T_HI=0x3FFFFF | - |
| `TIMG0_T0LO` | 0x6001F004 | 32 | T_LO=0xFFFFFFFF | - |
| `TIMG0_T0LOAD` | 0x6001F020 | 32 | T_LOAD=0xFFFFFFFF | - |
| `TIMG0_T0LOADHI` | 0x6001F01C | 32 | T_LOAD_HI=0x3FFFFF | - |
| `TIMG0_T0LOADLO` | 0x6001F018 | 32 | T_LOAD_LO=0xFFFFFFFF | - |
| `TIMG0_T0UPDATE` | 0x6001F00C | 32 | T_UPDATE=0x80000000 | - |
| `TIMG0_WDTCONFIG0` | 0x6001F048 | 32 | WDT_APPCPU_RESET_EN=0x1000 WDT_CONF_UPDATE_EN=0x400000 WDT_CPU_RESET_LENGTH=0x1C0000 WDT_EN=0x80000000 WDT_FLASHBOOT_MOD_EN=0x4000 WDT_PROCPU_RESET_EN=0x2000 WDT_STG0=0x60000000 WDT_STG1=0x18000000 WDT_STG2=0x6000000 WDT_STG3=0x1800000 WDT_... | - |
| `TIMG0_WDTCONFIG1` | 0x6001F04C | 32 | WDT_CLK_PRESCALE=0xFFFF0000 WDT_DIVCNT_RST=0x1 | - |
| `TIMG0_WDTCONFIG2` | 0x6001F050 | 32 | WDT_STG0_HOLD=0xFFFFFFFF | - |
| `TIMG0_WDTCONFIG3` | 0x6001F054 | 32 | WDT_STG1_HOLD=0xFFFFFFFF | - |
| `TIMG0_WDTCONFIG4` | 0x6001F058 | 32 | WDT_STG2_HOLD=0xFFFFFFFF | - |
| `TIMG0_WDTCONFIG5` | 0x6001F05C | 32 | WDT_STG3_HOLD=0xFFFFFFFF | - |
| `TIMG0_WDTFEED` | 0x6001F060 | 32 | WDT_FEED=0xFFFFFFFF | - |
| `TIMG0_WDTWPROTECT` | 0x6001F064 | 32 | WDT_WKEY=0xFFFFFFFF | - |
| `TIMG1_INT_CLR_TIMERS` | 0x6002007C | 32 | T0_INT_CLR=0x1 WDT_INT_CLR=0x2 | - |
| `TIMG1_INT_ENA_TIMERS` | 0x60020070 | 32 | T0_INT_ENA=0x1 WDT_INT_ENA=0x2 | - |
| `TIMG1_INT_RAW_TIMERS` | 0x60020074 | 32 | T0_INT_RAW=0x1 WDT_INT_RAW=0x2 | - |
| `TIMG1_INT_ST_TIMERS` | 0x60020078 | 32 | T0_INT_ST=0x1 WDT_INT_ST=0x2 | - |
| `TIMG1_NTIMERS_DATE` | 0x600200F8 | 32 | NTIMGS_DATE=0xFFFFFFF | - |
| `TIMG1_REGCLK` | 0x600200FC | 32 | CLK_EN=0x80000000 ETM_EN=0x10000000 TIMER_CLK_IS_ACTIVE=0x40000000 WDT_CLK_IS_ACTIVE=0x20000000 | - |
| `TIMG1_RTCCALICFG` | 0x60020068 | 32 | RTC_CALI_CLK_SEL=0x6000 RTC_CALI_MAX=0x7FFF0000 RTC_CALI_RDY=0x8000 RTC_CALI_START=0x80000000 RTC_CALI_START_CYCLING=0x1000 | - |
| `TIMG1_RTCCALICFG1` | 0x6002006C | 32 | RTC_CALI_CYCLING_DATA_VLD=0x1 RTC_CALI_VALUE=0xFFFFFF80 | - |
| `TIMG1_RTCCALICFG2` | 0x60020080 | 32 | RTC_CALI_TIMEOUT=0x1 RTC_CALI_TIMEOUT_RST_CNT=0x78 RTC_CALI_TIMEOUT_THRES=0xFFFFFF80 | - |
| `TIMG1_T0ALARMHI` | 0x60020014 | 32 | T_ALARM_HI=0x3FFFFF | - |
| `TIMG1_T0ALARMLO` | 0x60020010 | 32 | T_ALARM_LO=0xFFFFFFFF | - |
| `TIMG1_T0CONFIG` | 0x60020000 | 32 | T_ALARM_EN=0x400 T_AUTORELOAD=0x20000000 T_DIVCNT_RST=0x1000 T_DIVIDER=0x1FFFE000 T_EN=0x80000000 T_INCREASE=0x40000000 T_USE_XTAL=0x200 | - |
| `TIMG1_T0HI` | 0x60020008 | 32 | T_HI=0x3FFFFF | - |
| `TIMG1_T0LO` | 0x60020004 | 32 | T_LO=0xFFFFFFFF | - |
| `TIMG1_T0LOAD` | 0x60020020 | 32 | T_LOAD=0xFFFFFFFF | - |
| `TIMG1_T0LOADHI` | 0x6002001C | 32 | T_LOAD_HI=0x3FFFFF | - |
| `TIMG1_T0LOADLO` | 0x60020018 | 32 | T_LOAD_LO=0xFFFFFFFF | - |
| `TIMG1_T0UPDATE` | 0x6002000C | 32 | T_UPDATE=0x80000000 | - |
| `TIMG1_WDTCONFIG0` | 0x60020048 | 32 | WDT_APPCPU_RESET_EN=0x1000 WDT_CONF_UPDATE_EN=0x400000 WDT_CPU_RESET_LENGTH=0x1C0000 WDT_EN=0x80000000 WDT_FLASHBOOT_MOD_EN=0x4000 WDT_PROCPU_RESET_EN=0x2000 WDT_STG0=0x60000000 WDT_STG1=0x18000000 WDT_STG2=0x6000000 WDT_STG3=0x1800000 WDT_... | - |
| `TIMG1_WDTCONFIG1` | 0x6002004C | 32 | WDT_CLK_PRESCALE=0xFFFF0000 WDT_DIVCNT_RST=0x1 | - |
| `TIMG1_WDTCONFIG2` | 0x60020050 | 32 | WDT_STG0_HOLD=0xFFFFFFFF | - |
| `TIMG1_WDTCONFIG3` | 0x60020054 | 32 | WDT_STG1_HOLD=0xFFFFFFFF | - |
| `TIMG1_WDTCONFIG4` | 0x60020058 | 32 | WDT_STG2_HOLD=0xFFFFFFFF | - |
| `TIMG1_WDTCONFIG5` | 0x6002005C | 32 | WDT_STG3_HOLD=0xFFFFFFFF | - |
| `TIMG1_WDTFEED` | 0x60020060 | 32 | WDT_FEED=0xFFFFFFFF | - |
| `TIMG1_WDTWPROTECT` | 0x60020064 | 32 | WDT_WKEY=0xFFFFFFFF | - |
| `TWAI_ARB_LOST_CAP` | 0x6002B02C | 32 | ARBITRATION_LOST_CAPTURE=0x1F | - |
| `TWAI_BUS_TIMING_0` | 0x6002B018 | 32 | BAUD_PRESC=0x3FFF SYNC_JUMP_WIDTH=0xC000 | - |
| `TWAI_BUS_TIMING_1` | 0x6002B01C | 32 | TIME_SAMPLING=0x80 TIME_SEGMENT1=0xF TIME_SEGMENT2=0x70 | - |
| `TWAI_CLOCK_DIVIDER` | 0x6002B07C | 32 | CD=0xFF CLOCK_OFF=0x100 | - |
| `TWAI_CMD` | 0x6002B004 | 32 | ABORT_TX=0x2 CLEAR_DATA_OVERRUN=0x8 RELEASE_BUFFER=0x4 SELF_RX_REQUEST=0x10 TX_REQUEST=0x1 | - |
| `TWAI_DATA_0` | 0x6002B040 | 32 | DATA_0=0xFF | - |
| `TWAI_DATA_1` | 0x6002B044 | 32 | DATA_1=0xFF | - |
| `TWAI_DATA_10` | 0x6002B068 | 32 | DATA_10=0xFF | - |
| `TWAI_DATA_11` | 0x6002B06C | 32 | DATA_11=0xFF | - |
| `TWAI_DATA_12` | 0x6002B070 | 32 | DATA_12=0xFF | - |
| `TWAI_DATA_2` | 0x6002B048 | 32 | DATA_2=0xFF | - |
| `TWAI_DATA_3` | 0x6002B04C | 32 | DATA_3=0xFF | - |
| `TWAI_DATA_4` | 0x6002B050 | 32 | DATA_4=0xFF | - |
| `TWAI_DATA_5` | 0x6002B054 | 32 | DATA_5=0xFF | - |
| `TWAI_DATA_6` | 0x6002B058 | 32 | DATA_6=0xFF | - |
| `TWAI_DATA_7` | 0x6002B05C | 32 | DATA_7=0xFF | - |
| `TWAI_DATA_8` | 0x6002B060 | 32 | DATA_8=0xFF | - |
| `TWAI_DATA_9` | 0x6002B064 | 32 | DATA_9=0xFF | - |
| `TWAI_ECO_CFG` | 0x6002B090 | 32 | RDN_ENA=0x1 RDN_RESULT=0x2 | - |
| `TWAI_ERR_CODE_CAP` | 0x6002B030 | 32 | ERR_CAPTURE_CODE_DIRECTION=0x20 ERR_CAPTURE_CODE_SEGMENT=0x1F ERR_CAPTURE_CODE_TYPE=0xC0 | - |
| `TWAI_ERR_WARNING_LIMIT` | 0x6002B034 | 32 | ERR_WARNING_LIMIT=0xFF | - |
| `TWAI_HW_CFG` | 0x6002B084 | 32 | HW_STANDBY_EN=0x1 | - |
| `TWAI_HW_STANDBY_CNT` | 0x6002B088 | 32 | STANDBY_WAIT_CNT=0xFFFFFFFF | - |
| `TWAI_IDLE_INTR_CNT` | 0x6002B08C | 32 | IDLE_INTR_CNT=0xFFFFFFFF | - |
| `TWAI_INTERRUPT` | 0x6002B00C | 32 | ARBITRATION_LOST_INT_ST=0x40 BUS_ERR_INT_ST=0x80 DATA_OVERRUN_INT_ST=0x8 ERR_PASSIVE_INT_ST=0x20 ERR_WARNING_INT_ST=0x4 IDLE_INT_ST=0x100 RECEIVE_INT_ST=0x1 TRANSMIT_INT_ST=0x2 | - |
| `TWAI_INTERRUPT_ENABLE` | 0x6002B010 | 32 | ARBITRATION_LOST_INT_ENA=0x40 BUS_ERR_INT_ENA=0x80 ERR_PASSIVE_INT_ENA=0x20 EXT_DATA_OVERRUN_INT_ENA=0x8 EXT_ERR_WARNING_INT_ENA=0x4 EXT_RECEIVE_INT_ENA=0x1 EXT_TRANSMIT_INT_ENA=0x2 IDLE_INT_ENA=0x100 | - |
| `TWAI_MODE` | 0x6002B000 | 32 | ACCEPTANCE_FILTER_MODE=0x8 LISTEN_ONLY_MODE=0x2 RESET_MODE=0x1 SELF_TEST_MODE=0x4 | - |
| `TWAI_RX_ERR_CNT` | 0x6002B038 | 32 | RX_ERR_CNT=0xFF | - |
| `TWAI_RX_MESSAGE_COUNTER` | 0x6002B074 | 32 | RX_MESSAGE_COUNTER=0x7F | - |
| `TWAI_STATUS` | 0x6002B008 | 32 | ERR=0x40 MISS=0x100 NODE_BUS_OFF=0x80 OVERRUN=0x2 RECEIVE=0x10 RECEIVE_BUFFER=0x1 TRANSMISSION_COMPLETE=0x8 TRANSMIT=0x20 TRANSMIT_BUFFER=0x4 | - |
| `TWAI_SW_STANDBY_CFG` | 0x6002B080 | 32 | SW_STANDBY_CLR=0x2 SW_STANDBY_EN=0x1 | - |
| `TWAI_TX_ERR_CNT` | 0x6002B03C | 32 | TX_ERR_CNT=0xFF | - |
| `UART0_AFIFO_STATUS` | 0x60000090 | 32 | RX_AFIFO_EMPTY=0x8 RX_AFIFO_FULL=0x4 TX_AFIFO_EMPTY=0x2 TX_AFIFO_FULL=0x1 | - |
| `UART0_AT_CMD_CHAR_SYNC` | 0x6000005C | 32 | AT_CMD_CHAR=0xFF CHAR_NUM=0xFF00 | - |
| `UART0_AT_CMD_GAPTOUT_SYNC` | 0x60000058 | 32 | RX_GAP_TOUT=0xFFFF | - |
| `UART0_AT_CMD_POSTCNT_SYNC` | 0x60000054 | 32 | POST_IDLE_NUM=0xFFFF | - |
| `UART0_AT_CMD_PRECNT_SYNC` | 0x60000050 | 32 | PRE_IDLE_NUM=0xFFFF | - |
| `UART0_CLKDIV_SYNC` | 0x60000014 | 32 | CLKDIV=0xFFF CLKDIV_FRAG=0xF00000 | - |
| `UART0_CLK_CONF` | 0x60000088 | 32 | RX_RST_CORE=0x8000000 RX_SCLK_EN=0x2000000 TX_RST_CORE=0x4000000 TX_SCLK_EN=0x1000000 | - |
| `UART0_CONF0_SYNC` | 0x60000020 | 32 | AUTOBAUD_EN=0x80000 BIT_NUM=0xC DIS_RX_DAT_OVF=0x20000 ERR_WR_MASK=0x40000 IRDA_DPLX=0x80 IRDA_EN=0x4000 IRDA_RX_INV=0x800 IRDA_TX_EN=0x100 IRDA_TX_INV=0x400 IRDA_WCTL=0x200 LOOPBACK=0x1000 MEM_CLK_EN=0x100000 PARITY=0x1 PARITY_EN=0x2 RXD_I... | - |
| `UART0_CONF1` | 0x60000024 | 32 | CLK_EN=0x200000 CTS_INV=0x10000 DSR_INV=0x20000 DTR_INV=0x80000 RTS_INV=0x40000 RXFIFO_FULL_THRHD=0xFF SW_DTR=0x100000 TXFIFO_EMPTY_THRHD=0xFF00 | - |
| `UART0_DATE` | 0x6000008C | 32 | DATE=0xFFFFFFFF | - |
| `UART0_FIFO` | 0x60000000 | 32 | RXFIFO_RD_BYTE=0xFF | - |
| `UART0_FSM_STATUS` | 0x60000070 | 32 | ST_URX_OUT=0xF ST_UTX_OUT=0xF0 | - |
| `UART0_HIGHPULSE` | 0x60000080 | 32 | MIN_CNT=0xFFF | - |
| `UART0_HWFC_CONF_SYNC` | 0x6000002C | 32 | RX_FLOW_EN=0x100 RX_FLOW_THRHD=0xFF | - |
| `UART0_ID` | 0x6000009C | 32 | ID=0xFFFFFFFF | - |
| `UART0_IDLE_CONF_SYNC` | 0x60000048 | 32 | RX_IDLE_THRHD=0x3FF TX_IDLE_NUM=0xFFC00 | - |
| `UART0_INT_CLR` | 0x60000010 | 32 | AT_CMD_CHAR_DET_INT_CLR=0x40000 BRK_DET_INT_CLR=0x80 CTS_CHG_INT_CLR=0x40 DSR_CHG_INT_CLR=0x20 FRM_ERR_INT_CLR=0x8 GLITCH_DET_INT_CLR=0x800 PARITY_ERR_INT_CLR=0x4 RS485_CLASH_INT_CLR=0x20000 RS485_FRM_ERR_INT_CLR=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART0_INT_ENA` | 0x6000000C | 32 | AT_CMD_CHAR_DET_INT_ENA=0x40000 BRK_DET_INT_ENA=0x80 CTS_CHG_INT_ENA=0x40 DSR_CHG_INT_ENA=0x20 FRM_ERR_INT_ENA=0x8 GLITCH_DET_INT_ENA=0x800 PARITY_ERR_INT_ENA=0x4 RS485_CLASH_INT_ENA=0x20000 RS485_FRM_ERR_INT_ENA=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART0_INT_RAW` | 0x60000004 | 32 | AT_CMD_CHAR_DET_INT_RAW=0x40000 BRK_DET_INT_RAW=0x80 CTS_CHG_INT_RAW=0x40 DSR_CHG_INT_RAW=0x20 FRM_ERR_INT_RAW=0x8 GLITCH_DET_INT_RAW=0x800 PARITY_ERR_INT_RAW=0x4 RS485_CLASH_INT_RAW=0x20000 RS485_FRM_ERR_INT_RAW=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART0_INT_ST` | 0x60000008 | 32 | AT_CMD_CHAR_DET_INT_ST=0x40000 BRK_DET_INT_ST=0x80 CTS_CHG_INT_ST=0x40 DSR_CHG_INT_ST=0x20 FRM_ERR_INT_ST=0x8 GLITCH_DET_INT_ST=0x800 PARITY_ERR_INT_ST=0x4 RS485_CLASH_INT_ST=0x20000 RS485_FRM_ERR_INT_ST=0x10000 RS485_PARITY_ERR_INT_ST=0x80... | - |
| `UART0_LOWPULSE` | 0x6000007C | 32 | MIN_CNT=0xFFF | - |
| `UART0_MEM_CONF` | 0x60000060 | 32 | MEM_FORCE_PD=0x2000000 MEM_FORCE_PU=0x4000000 | - |
| `UART0_MEM_RX_STATUS` | 0x6000006C | 32 | RX_SRAM_RADDR=0xFF RX_SRAM_WADDR=0x1FE00 | - |
| `UART0_MEM_TX_STATUS` | 0x60000068 | 32 | TX_SRAM_RADDR=0x1FE00 TX_SRAM_WADDR=0xFF | - |
| `UART0_NEGPULSE` | 0x60000078 | 32 | NEGEDGE_MIN_CNT=0xFFF | - |
| `UART0_POSPULSE` | 0x60000074 | 32 | POSEDGE_MIN_CNT=0xFFF | - |
| `UART0_REG_UPDATE` | 0x60000098 | 32 | REG_UPDATE=0x1 | - |
| `UART0_RS485_CONF_SYNC` | 0x6000004C | 32 | DL0_EN=0x2 DL1_EN=0x4 RS485RXBY_TX_EN=0x10 RS485TX_RX_EN=0x8 RS485_EN=0x1 RS485_RX_DLY_NUM=0x20 RS485_TX_DLY_NUM=0x3C0 | - |
| `UART0_RXD_CNT` | 0x60000084 | 32 | RXD_EDGE_CNT=0x3FF | - |
| `UART0_RX_FILT` | 0x60000018 | 32 | GLITCH_FILT=0xFF GLITCH_FILT_EN=0x100 | - |
| `UART0_SLEEP_CONF0` | 0x60000030 | 32 | WK_CHAR1=0xFF WK_CHAR2=0xFF00 WK_CHAR3=0xFF0000 WK_CHAR4=0xFF000000 | - |
| `UART0_SLEEP_CONF1` | 0x60000034 | 32 | WK_CHAR0=0xFF | - |
| `UART0_SLEEP_CONF2` | 0x60000038 | 32 | ACTIVE_THRESHOLD=0x3FF RX_WAKE_UP_THRHD=0x3FC00 WK_CHAR_MASK=0x3E00000 WK_CHAR_NUM=0x1C0000 WK_MODE_SEL=0xC000000 | - |
| `UART0_STATUS` | 0x6000001C | 32 | CTSN=0x4000 DSRN=0x2000 DTRN=0x20000000 RTSN=0x40000000 RXD=0x8000 RXFIFO_CNT=0xFF TXD=0x80000000 TXFIFO_CNT=0xFF0000 | - |
| `UART0_SWFC_CONF0_SYNC` | 0x6000003C | 32 | FORCE_XOFF=0x100000 FORCE_XON=0x80000 SEND_XOFF=0x400000 SEND_XON=0x200000 SW_FLOW_CON_EN=0x20000 XOFF_CHAR=0xFF00 XONOFF_DEL=0x40000 XON_CHAR=0xFF XON_XOFF_STILL_SEND=0x10000 | - |
| `UART0_SWFC_CONF1` | 0x60000040 | 32 | XOFF_THRESHOLD=0xFF00 XON_THRESHOLD=0xFF | - |
| `UART0_TOUT_CONF_SYNC` | 0x60000064 | 32 | RX_TOUT_EN=0x1 RX_TOUT_FLOW_DIS=0x2 RX_TOUT_THRHD=0xFFC | - |
| `UART0_TXBRK_CONF_SYNC` | 0x60000044 | 32 | TX_BRK_NUM=0xFF | - |
| `UART1_AFIFO_STATUS` | 0x60010090 | 32 | RX_AFIFO_EMPTY=0x8 RX_AFIFO_FULL=0x4 TX_AFIFO_EMPTY=0x2 TX_AFIFO_FULL=0x1 | - |
| `UART1_AT_CMD_CHAR_SYNC` | 0x6001005C | 32 | AT_CMD_CHAR=0xFF CHAR_NUM=0xFF00 | - |
| `UART1_AT_CMD_GAPTOUT_SYNC` | 0x60010058 | 32 | RX_GAP_TOUT=0xFFFF | - |
| `UART1_AT_CMD_POSTCNT_SYNC` | 0x60010054 | 32 | POST_IDLE_NUM=0xFFFF | - |
| `UART1_AT_CMD_PRECNT_SYNC` | 0x60010050 | 32 | PRE_IDLE_NUM=0xFFFF | - |
| `UART1_CLKDIV_SYNC` | 0x60010014 | 32 | CLKDIV=0xFFF CLKDIV_FRAG=0xF00000 | - |
| `UART1_CLK_CONF` | 0x60010088 | 32 | RX_RST_CORE=0x8000000 RX_SCLK_EN=0x2000000 TX_RST_CORE=0x4000000 TX_SCLK_EN=0x1000000 | - |
| `UART1_CONF0_SYNC` | 0x60010020 | 32 | AUTOBAUD_EN=0x80000 BIT_NUM=0xC DIS_RX_DAT_OVF=0x20000 ERR_WR_MASK=0x40000 IRDA_DPLX=0x80 IRDA_EN=0x4000 IRDA_RX_INV=0x800 IRDA_TX_EN=0x100 IRDA_TX_INV=0x400 IRDA_WCTL=0x200 LOOPBACK=0x1000 MEM_CLK_EN=0x100000 PARITY=0x1 PARITY_EN=0x2 RXD_I... | - |
| `UART1_CONF1` | 0x60010024 | 32 | CLK_EN=0x200000 CTS_INV=0x10000 DSR_INV=0x20000 DTR_INV=0x80000 RTS_INV=0x40000 RXFIFO_FULL_THRHD=0xFF SW_DTR=0x100000 TXFIFO_EMPTY_THRHD=0xFF00 | - |
| `UART1_DATE` | 0x6001008C | 32 | DATE=0xFFFFFFFF | - |
| `UART1_FIFO` | 0x60010000 | 32 | RXFIFO_RD_BYTE=0xFF | - |
| `UART1_FSM_STATUS` | 0x60010070 | 32 | ST_URX_OUT=0xF ST_UTX_OUT=0xF0 | - |
| `UART1_HIGHPULSE` | 0x60010080 | 32 | MIN_CNT=0xFFF | - |
| `UART1_HWFC_CONF_SYNC` | 0x6001002C | 32 | RX_FLOW_EN=0x100 RX_FLOW_THRHD=0xFF | - |
| `UART1_ID` | 0x6001009C | 32 | ID=0xFFFFFFFF | - |
| `UART1_IDLE_CONF_SYNC` | 0x60010048 | 32 | RX_IDLE_THRHD=0x3FF TX_IDLE_NUM=0xFFC00 | - |
| `UART1_INT_CLR` | 0x60010010 | 32 | AT_CMD_CHAR_DET_INT_CLR=0x40000 BRK_DET_INT_CLR=0x80 CTS_CHG_INT_CLR=0x40 DSR_CHG_INT_CLR=0x20 FRM_ERR_INT_CLR=0x8 GLITCH_DET_INT_CLR=0x800 PARITY_ERR_INT_CLR=0x4 RS485_CLASH_INT_CLR=0x20000 RS485_FRM_ERR_INT_CLR=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART1_INT_ENA` | 0x6001000C | 32 | AT_CMD_CHAR_DET_INT_ENA=0x40000 BRK_DET_INT_ENA=0x80 CTS_CHG_INT_ENA=0x40 DSR_CHG_INT_ENA=0x20 FRM_ERR_INT_ENA=0x8 GLITCH_DET_INT_ENA=0x800 PARITY_ERR_INT_ENA=0x4 RS485_CLASH_INT_ENA=0x20000 RS485_FRM_ERR_INT_ENA=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART1_INT_RAW` | 0x60010004 | 32 | AT_CMD_CHAR_DET_INT_RAW=0x40000 BRK_DET_INT_RAW=0x80 CTS_CHG_INT_RAW=0x40 DSR_CHG_INT_RAW=0x20 FRM_ERR_INT_RAW=0x8 GLITCH_DET_INT_RAW=0x800 PARITY_ERR_INT_RAW=0x4 RS485_CLASH_INT_RAW=0x20000 RS485_FRM_ERR_INT_RAW=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART1_INT_ST` | 0x60010008 | 32 | AT_CMD_CHAR_DET_INT_ST=0x40000 BRK_DET_INT_ST=0x80 CTS_CHG_INT_ST=0x40 DSR_CHG_INT_ST=0x20 FRM_ERR_INT_ST=0x8 GLITCH_DET_INT_ST=0x800 PARITY_ERR_INT_ST=0x4 RS485_CLASH_INT_ST=0x20000 RS485_FRM_ERR_INT_ST=0x10000 RS485_PARITY_ERR_INT_ST=0x80... | - |
| `UART1_LOWPULSE` | 0x6001007C | 32 | MIN_CNT=0xFFF | - |
| `UART1_MEM_CONF` | 0x60010060 | 32 | MEM_FORCE_PD=0x2000000 MEM_FORCE_PU=0x4000000 | - |
| `UART1_MEM_RX_STATUS` | 0x6001006C | 32 | RX_SRAM_RADDR=0xFF RX_SRAM_WADDR=0x1FE00 | - |
| `UART1_MEM_TX_STATUS` | 0x60010068 | 32 | TX_SRAM_RADDR=0x1FE00 TX_SRAM_WADDR=0xFF | - |
| `UART1_NEGPULSE` | 0x60010078 | 32 | NEGEDGE_MIN_CNT=0xFFF | - |
| `UART1_POSPULSE` | 0x60010074 | 32 | POSEDGE_MIN_CNT=0xFFF | - |
| `UART1_REG_UPDATE` | 0x60010098 | 32 | REG_UPDATE=0x1 | - |
| `UART1_RS485_CONF_SYNC` | 0x6001004C | 32 | DL0_EN=0x2 DL1_EN=0x4 RS485RXBY_TX_EN=0x10 RS485TX_RX_EN=0x8 RS485_EN=0x1 RS485_RX_DLY_NUM=0x20 RS485_TX_DLY_NUM=0x3C0 | - |
| `UART1_RXD_CNT` | 0x60010084 | 32 | RXD_EDGE_CNT=0x3FF | - |
| `UART1_RX_FILT` | 0x60010018 | 32 | GLITCH_FILT=0xFF GLITCH_FILT_EN=0x100 | - |
| `UART1_SLEEP_CONF0` | 0x60010030 | 32 | WK_CHAR1=0xFF WK_CHAR2=0xFF00 WK_CHAR3=0xFF0000 WK_CHAR4=0xFF000000 | - |
| `UART1_SLEEP_CONF1` | 0x60010034 | 32 | WK_CHAR0=0xFF | - |
| `UART1_SLEEP_CONF2` | 0x60010038 | 32 | ACTIVE_THRESHOLD=0x3FF RX_WAKE_UP_THRHD=0x3FC00 WK_CHAR_MASK=0x3E00000 WK_CHAR_NUM=0x1C0000 WK_MODE_SEL=0xC000000 | - |
| `UART1_STATUS` | 0x6001001C | 32 | CTSN=0x4000 DSRN=0x2000 DTRN=0x20000000 RTSN=0x40000000 RXD=0x8000 RXFIFO_CNT=0xFF TXD=0x80000000 TXFIFO_CNT=0xFF0000 | - |
| `UART1_SWFC_CONF0_SYNC` | 0x6001003C | 32 | FORCE_XOFF=0x100000 FORCE_XON=0x80000 SEND_XOFF=0x400000 SEND_XON=0x200000 SW_FLOW_CON_EN=0x20000 XOFF_CHAR=0xFF00 XONOFF_DEL=0x40000 XON_CHAR=0xFF XON_XOFF_STILL_SEND=0x10000 | - |
| `UART1_SWFC_CONF1` | 0x60010040 | 32 | XOFF_THRESHOLD=0xFF00 XON_THRESHOLD=0xFF | - |
| `UART1_TOUT_CONF_SYNC` | 0x60010064 | 32 | RX_TOUT_EN=0x1 RX_TOUT_FLOW_DIS=0x2 RX_TOUT_THRHD=0xFFC | - |
| `UART1_TXBRK_CONF_SYNC` | 0x60010044 | 32 | TX_BRK_NUM=0xFF | - |
| `UHCI0_ACK_NUM` | 0x60014028 | 32 | ACK_NUM=0x7 LOAD=0x8 | - |
| `UHCI0_CONF0` | 0x60014000 | 32 | CLK_EN=0x800 CRC_REC_EN=0x80 ENCODE_CRC_EN=0x400 HEAD_EN=0x40 LEN_EOF_EN=0x200 RX_RST=0x2 SEPER_EN=0x20 TX_RST=0x1 UART0_CE=0x4 UART1_CE=0x8 UART_IDLE_EOF_EN=0x100 UART_RX_BRK_EOF_EN=0x1000 | - |
| `UHCI0_CONF1` | 0x60014014 | 32 | CHECK_SEQ_EN=0x2 CHECK_SUM_EN=0x1 CRC_DISABLE=0x4 SAVE_HEAD=0x8 SW_START=0x100 TX_ACK_NUM_RE=0x20 TX_CHECK_SUM_RE=0x10 WAIT_SW_START=0x80 | - |
| `UHCI0_DATE` | 0x60014080 | 32 | DATE=0xFFFFFFFF | - |
| `UHCI0_ESCAPE_CONF` | 0x60014020 | 32 | RX_11_ESC_EN=0x40 RX_13_ESC_EN=0x80 RX_C0_ESC_EN=0x10 RX_DB_ESC_EN=0x20 TX_11_ESC_EN=0x4 TX_13_ESC_EN=0x8 TX_C0_ESC_EN=0x1 TX_DB_ESC_EN=0x2 | - |
| `UHCI0_ESC_CONF0` | 0x6001406C | 32 | SEPER_CHAR=0xFF SEPER_ESC_CHAR0=0xFF00 SEPER_ESC_CHAR1=0xFF0000 | - |
| `UHCI0_ESC_CONF1` | 0x60014070 | 32 | ESC_SEQ0=0xFF ESC_SEQ0_CHAR0=0xFF00 ESC_SEQ0_CHAR1=0xFF0000 | - |
| `UHCI0_ESC_CONF2` | 0x60014074 | 32 | ESC_SEQ1=0xFF ESC_SEQ1_CHAR0=0xFF00 ESC_SEQ1_CHAR1=0xFF0000 | - |
| `UHCI0_ESC_CONF3` | 0x60014078 | 32 | ESC_SEQ2=0xFF ESC_SEQ2_CHAR0=0xFF00 ESC_SEQ2_CHAR1=0xFF0000 | - |
| `UHCI0_HUNG_CONF` | 0x60014024 | 32 | RXFIFO_TIMEOUT=0xFF000 RXFIFO_TIMEOUT_ENA=0x800000 RXFIFO_TIMEOUT_SHIFT=0x700000 TXFIFO_TIMEOUT=0xFF TXFIFO_TIMEOUT_ENA=0x800 TXFIFO_TIMEOUT_SHIFT=0x700 | - |
| `UHCI0_INT_CLR` | 0x60014010 | 32 | APP_CTRL0_INT_CLR=0x80 APP_CTRL1_INT_CLR=0x100 OUTLINK_EOF_ERR_INT_CLR=0x40 RX_HUNG_INT_CLR=0x4 RX_START_INT_CLR=0x1 SEND_A_REG_Q_INT_CLR=0x20 SEND_S_REG_Q_INT_CLR=0x10 TX_HUNG_INT_CLR=0x8 TX_START_INT_CLR=0x2 | - |
| `UHCI0_INT_ENA` | 0x6001400C | 32 | APP_CTRL0_INT_ENA=0x80 APP_CTRL1_INT_ENA=0x100 OUTLINK_EOF_ERR_INT_ENA=0x40 RX_HUNG_INT_ENA=0x4 RX_START_INT_ENA=0x1 SEND_A_REG_Q_INT_ENA=0x20 SEND_S_REG_Q_INT_ENA=0x10 TX_HUNG_INT_ENA=0x8 TX_START_INT_ENA=0x2 | - |
| `UHCI0_INT_RAW` | 0x60014004 | 32 | APP_CTRL0_INT_RAW=0x80 APP_CTRL1_INT_RAW=0x100 OUT_EOF_INT_RAW=0x40 RX_HUNG_INT_RAW=0x4 RX_START_INT_RAW=0x1 SEND_A_REG_Q_INT_RAW=0x20 SEND_S_REG_Q_INT_RAW=0x10 TX_HUNG_INT_RAW=0x8 TX_START_INT_RAW=0x2 | - |
| `UHCI0_INT_ST` | 0x60014008 | 32 | APP_CTRL0_INT_ST=0x80 APP_CTRL1_INT_ST=0x100 OUTLINK_EOF_ERR_INT_ST=0x40 RX_HUNG_INT_ST=0x4 RX_START_INT_ST=0x1 SEND_A_REG_Q_INT_ST=0x20 SEND_S_REG_Q_INT_ST=0x10 TX_HUNG_INT_ST=0x8 TX_START_INT_ST=0x2 | - |
| `UHCI0_PKT_THRES` | 0x6001407C | 32 | PKT_THRS=0x1FFF | - |
| `UHCI0_QUICK_SENT` | 0x60014030 | 32 | ALWAYS_SEND_EN=0x80 ALWAYS_SEND_NUM=0x70 SINGLE_SEND_EN=0x8 SINGLE_SEND_NUM=0x7 | - |
| `UHCI0_REG_Q0_WORD0` | 0x60014034 | 32 | SEND_Q0_WORD0=0xFFFFFFFF | - |
| `UHCI0_REG_Q0_WORD1` | 0x60014038 | 32 | SEND_Q0_WORD1=0xFFFFFFFF | - |
| `UHCI0_REG_Q1_WORD0` | 0x6001403C | 32 | SEND_Q1_WORD0=0xFFFFFFFF | - |
| `UHCI0_REG_Q1_WORD1` | 0x60014040 | 32 | SEND_Q1_WORD1=0xFFFFFFFF | - |
| `UHCI0_REG_Q2_WORD0` | 0x60014044 | 32 | SEND_Q2_WORD0=0xFFFFFFFF | - |
| `UHCI0_REG_Q2_WORD1` | 0x60014048 | 32 | SEND_Q2_WORD1=0xFFFFFFFF | - |
| `UHCI0_REG_Q3_WORD0` | 0x6001404C | 32 | SEND_Q3_WORD0=0xFFFFFFFF | - |
| `UHCI0_REG_Q3_WORD1` | 0x60014050 | 32 | SEND_Q3_WORD1=0xFFFFFFFF | - |
| `UHCI0_REG_Q4_WORD0` | 0x60014054 | 32 | SEND_Q4_WORD0=0xFFFFFFFF | - |
| `UHCI0_REG_Q4_WORD1` | 0x60014058 | 32 | SEND_Q4_WORD1=0xFFFFFFFF | - |
| `UHCI0_REG_Q5_WORD0` | 0x6001405C | 32 | SEND_Q5_WORD0=0xFFFFFFFF | - |
| `UHCI0_REG_Q5_WORD1` | 0x60014060 | 32 | SEND_Q5_WORD1=0xFFFFFFFF | - |
| `UHCI0_REG_Q6_WORD0` | 0x60014064 | 32 | SEND_Q6_WORD0=0xFFFFFFFF | - |
| `UHCI0_REG_Q6_WORD1` | 0x60014068 | 32 | SEND_Q6_WORD1=0xFFFFFFFF | - |
| `UHCI0_RX_HEAD` | 0x6001402C | 32 | RX_HEAD=0xFFFFFFFF | - |
| `UHCI0_STATE0` | 0x60014018 | 32 | DECODE_STATE=0x38 RX_ERR_CAUSE=0x7 | - |
| `UHCI0_STATE1` | 0x6001401C | 32 | ENCODE_STATE=0x7 | - |
| `UHCI1_ACK_NUM` | 0x6000C028 | 32 | ACK_NUM=0x7 LOAD=0x8 | - |
| `UHCI1_CONF0` | 0x6000C000 | 32 | CLK_EN=0x800 CRC_REC_EN=0x80 ENCODE_CRC_EN=0x400 HEAD_EN=0x40 LEN_EOF_EN=0x200 RX_RST=0x2 SEPER_EN=0x20 TX_RST=0x1 UART0_CE=0x4 UART1_CE=0x8 UART_IDLE_EOF_EN=0x100 UART_RX_BRK_EOF_EN=0x1000 | - |
| `UHCI1_CONF1` | 0x6000C014 | 32 | CHECK_SEQ_EN=0x2 CHECK_SUM_EN=0x1 CRC_DISABLE=0x4 SAVE_HEAD=0x8 SW_START=0x100 TX_ACK_NUM_RE=0x20 TX_CHECK_SUM_RE=0x10 WAIT_SW_START=0x80 | - |
| `UHCI1_DATE` | 0x6000C080 | 32 | DATE=0xFFFFFFFF | - |
| `UHCI1_ESCAPE_CONF` | 0x6000C020 | 32 | RX_11_ESC_EN=0x40 RX_13_ESC_EN=0x80 RX_C0_ESC_EN=0x10 RX_DB_ESC_EN=0x20 TX_11_ESC_EN=0x4 TX_13_ESC_EN=0x8 TX_C0_ESC_EN=0x1 TX_DB_ESC_EN=0x2 | - |
| `UHCI1_ESC_CONF0` | 0x6000C06C | 32 | SEPER_CHAR=0xFF SEPER_ESC_CHAR0=0xFF00 SEPER_ESC_CHAR1=0xFF0000 | - |
| `UHCI1_ESC_CONF1` | 0x6000C070 | 32 | ESC_SEQ0=0xFF ESC_SEQ0_CHAR0=0xFF00 ESC_SEQ0_CHAR1=0xFF0000 | - |
| `UHCI1_ESC_CONF2` | 0x6000C074 | 32 | ESC_SEQ1=0xFF ESC_SEQ1_CHAR0=0xFF00 ESC_SEQ1_CHAR1=0xFF0000 | - |
| `UHCI1_ESC_CONF3` | 0x6000C078 | 32 | ESC_SEQ2=0xFF ESC_SEQ2_CHAR0=0xFF00 ESC_SEQ2_CHAR1=0xFF0000 | - |
| `UHCI1_HUNG_CONF` | 0x6000C024 | 32 | RXFIFO_TIMEOUT=0xFF000 RXFIFO_TIMEOUT_ENA=0x800000 RXFIFO_TIMEOUT_SHIFT=0x700000 TXFIFO_TIMEOUT=0xFF TXFIFO_TIMEOUT_ENA=0x800 TXFIFO_TIMEOUT_SHIFT=0x700 | - |
| `UHCI1_INT_CLR` | 0x6000C010 | 32 | APP_CTRL0_INT_CLR=0x80 APP_CTRL1_INT_CLR=0x100 OUTLINK_EOF_ERR_INT_CLR=0x40 RX_HUNG_INT_CLR=0x4 RX_START_INT_CLR=0x1 SEND_A_REG_Q_INT_CLR=0x20 SEND_S_REG_Q_INT_CLR=0x10 TX_HUNG_INT_CLR=0x8 TX_START_INT_CLR=0x2 | - |
| `UHCI1_INT_ENA` | 0x6000C00C | 32 | APP_CTRL0_INT_ENA=0x80 APP_CTRL1_INT_ENA=0x100 OUTLINK_EOF_ERR_INT_ENA=0x40 RX_HUNG_INT_ENA=0x4 RX_START_INT_ENA=0x1 SEND_A_REG_Q_INT_ENA=0x20 SEND_S_REG_Q_INT_ENA=0x10 TX_HUNG_INT_ENA=0x8 TX_START_INT_ENA=0x2 | - |
| `UHCI1_INT_RAW` | 0x6000C004 | 32 | APP_CTRL0_INT_RAW=0x80 APP_CTRL1_INT_RAW=0x100 OUT_EOF_INT_RAW=0x40 RX_HUNG_INT_RAW=0x4 RX_START_INT_RAW=0x1 SEND_A_REG_Q_INT_RAW=0x20 SEND_S_REG_Q_INT_RAW=0x10 TX_HUNG_INT_RAW=0x8 TX_START_INT_RAW=0x2 | - |
| `UHCI1_INT_ST` | 0x6000C008 | 32 | APP_CTRL0_INT_ST=0x80 APP_CTRL1_INT_ST=0x100 OUTLINK_EOF_ERR_INT_ST=0x40 RX_HUNG_INT_ST=0x4 RX_START_INT_ST=0x1 SEND_A_REG_Q_INT_ST=0x20 SEND_S_REG_Q_INT_ST=0x10 TX_HUNG_INT_ST=0x8 TX_START_INT_ST=0x2 | - |
| `UHCI1_PKT_THRES` | 0x6000C07C | 32 | PKT_THRS=0x1FFF | - |
| `UHCI1_QUICK_SENT` | 0x6000C030 | 32 | ALWAYS_SEND_EN=0x80 ALWAYS_SEND_NUM=0x70 SINGLE_SEND_EN=0x8 SINGLE_SEND_NUM=0x7 | - |
| `UHCI1_REG_Q0_WORD0` | 0x6000C034 | 32 | SEND_Q0_WORD0=0xFFFFFFFF | - |
| `UHCI1_REG_Q0_WORD1` | 0x6000C038 | 32 | SEND_Q0_WORD1=0xFFFFFFFF | - |
| `UHCI1_REG_Q1_WORD0` | 0x6000C03C | 32 | SEND_Q1_WORD0=0xFFFFFFFF | - |
| `UHCI1_REG_Q1_WORD1` | 0x6000C040 | 32 | SEND_Q1_WORD1=0xFFFFFFFF | - |
| `UHCI1_REG_Q2_WORD0` | 0x6000C044 | 32 | SEND_Q2_WORD0=0xFFFFFFFF | - |
| `UHCI1_REG_Q2_WORD1` | 0x6000C048 | 32 | SEND_Q2_WORD1=0xFFFFFFFF | - |
| `UHCI1_REG_Q3_WORD0` | 0x6000C04C | 32 | SEND_Q3_WORD0=0xFFFFFFFF | - |
| `UHCI1_REG_Q3_WORD1` | 0x6000C050 | 32 | SEND_Q3_WORD1=0xFFFFFFFF | - |
| `UHCI1_REG_Q4_WORD0` | 0x6000C054 | 32 | SEND_Q4_WORD0=0xFFFFFFFF | - |
| `UHCI1_REG_Q4_WORD1` | 0x6000C058 | 32 | SEND_Q4_WORD1=0xFFFFFFFF | - |
| `UHCI1_REG_Q5_WORD0` | 0x6000C05C | 32 | SEND_Q5_WORD0=0xFFFFFFFF | - |
| `UHCI1_REG_Q5_WORD1` | 0x6000C060 | 32 | SEND_Q5_WORD1=0xFFFFFFFF | - |
| `UHCI1_REG_Q6_WORD0` | 0x6000C064 | 32 | SEND_Q6_WORD0=0xFFFFFFFF | - |
| `UHCI1_REG_Q6_WORD1` | 0x6000C068 | 32 | SEND_Q6_WORD1=0xFFFFFFFF | - |
| `UHCI1_RX_HEAD` | 0x6000C02C | 32 | RX_HEAD=0xFFFFFFFF | - |
| `UHCI1_STATE0` | 0x6000C018 | 32 | DECODE_STATE=0x38 RX_ERR_CAUSE=0x7 | - |
| `UHCI1_STATE1` | 0x6000C01C | 32 | ENCODE_STATE=0x7 | - |
| `USB_DEVICE_BUS_RESET_ST` | 0x60043068 | 32 | USB_SERIAL_JTAG_USB_BUS_RESET_ST=0x1 | - |
| `USB_DEVICE_CHIP_RST` | 0x6004304C | 32 | USB_SERIAL_JTAG_DTR=0x2 USB_SERIAL_JTAG_RTS=0x1 USB_SERIAL_JTAG_USB_UART_CHIP_RST_DIS=0x4 | - |
| `USB_DEVICE_CONF0` | 0x60043018 | 32 | USB_SERIAL_JTAG_DM_PULLDOWN=0x1000 USB_SERIAL_JTAG_DM_PULLUP=0x800 USB_SERIAL_JTAG_DP_PULLDOWN=0x400 USB_SERIAL_JTAG_DP_PULLUP=0x200 USB_SERIAL_JTAG_EXCHG_PINS=0x4 USB_SERIAL_JTAG_EXCHG_PINS_OVERRIDE=0x2 USB_SERIAL_JTAG_PAD_PULL_OVERRIDE=0x... | - |
| `USB_DEVICE_CONFIG_UPDATE` | 0x60043060 | 32 | USB_SERIAL_JTAG_CONFIG_UPDATE=0x1 | - |
| `USB_DEVICE_DATE` | 0x60043080 | 32 | USB_SERIAL_JTAG_DATE=0xFFFFFFFF | - |
| `USB_DEVICE_EP1` | 0x60043000 | 32 | USB_SERIAL_JTAG_RDWR_BYTE=0xFF | - |
| `USB_DEVICE_EP1_CONF` | 0x60043004 | 32 | USB_SERIAL_JTAG_SERIAL_IN_EP_DATA_FREE=0x2 USB_SERIAL_JTAG_SERIAL_OUT_EP_DATA_AVAIL=0x4 USB_SERIAL_JTAG_WR_DONE=0x1 | - |
| `USB_DEVICE_FRAM_NUM` | 0x60043024 | 32 | USB_SERIAL_JTAG_SOF_FRAME_INDEX=0x7FF | - |
| `USB_DEVICE_GET_LINE_CODE_W0` | 0x60043058 | 32 | USB_SERIAL_JTAG_GET_DW_DTE_RATE=0xFFFFFFFF | - |
| `USB_DEVICE_GET_LINE_CODE_W1` | 0x6004305C | 32 | USB_SERIAL_JTAG_GET_BCHAR_FORMAT=0xFF0000 USB_SERIAL_JTAG_GET_BDATA_BITS=0xFF USB_SERIAL_JTAG_GET_BPARITY_TYPE=0xFF00 | - |
| `USB_DEVICE_INT_CLR` | 0x60043014 | 32 | USB_SERIAL_JTAG_CRC16_ERR_INT_CLR=0x40 USB_SERIAL_JTAG_CRC5_ERR_INT_CLR=0x20 USB_SERIAL_JTAG_DTR_CHG_INT_CLR=0x2000 USB_SERIAL_JTAG_GET_LINE_CODE_INT_CLR=0x4000 USB_SERIAL_JTAG_IN_TOKEN_REC_IN_EP1_INT_CLR=0x100 USB_SERIAL_JTAG_JTAG_IN_FLUSH... | - |
| `USB_DEVICE_INT_ENA` | 0x60043010 | 32 | USB_SERIAL_JTAG_CRC16_ERR_INT_ENA=0x40 USB_SERIAL_JTAG_CRC5_ERR_INT_ENA=0x20 USB_SERIAL_JTAG_DTR_CHG_INT_ENA=0x2000 USB_SERIAL_JTAG_GET_LINE_CODE_INT_ENA=0x4000 USB_SERIAL_JTAG_IN_TOKEN_REC_IN_EP1_INT_ENA=0x100 USB_SERIAL_JTAG_JTAG_IN_FLUSH... | - |
| `USB_DEVICE_INT_RAW` | 0x60043008 | 32 | USB_SERIAL_JTAG_CRC16_ERR_INT_RAW=0x40 USB_SERIAL_JTAG_CRC5_ERR_INT_RAW=0x20 USB_SERIAL_JTAG_DTR_CHG_INT_RAW=0x2000 USB_SERIAL_JTAG_GET_LINE_CODE_INT_RAW=0x4000 USB_SERIAL_JTAG_IN_TOKEN_REC_IN_EP1_INT_RAW=0x100 USB_SERIAL_JTAG_JTAG_IN_FLUSH... | - |
| `USB_DEVICE_INT_ST` | 0x6004300C | 32 | USB_SERIAL_JTAG_CRC16_ERR_INT_ST=0x40 USB_SERIAL_JTAG_CRC5_ERR_INT_ST=0x20 USB_SERIAL_JTAG_DTR_CHG_INT_ST=0x2000 USB_SERIAL_JTAG_GET_LINE_CODE_INT_ST=0x4000 USB_SERIAL_JTAG_IN_TOKEN_REC_IN_EP1_INT_ST=0x100 USB_SERIAL_JTAG_JTAG_IN_FLUSH_INT_... | - |
| `USB_DEVICE_IN_EP0_ST` | 0x60043028 | 32 | USB_SERIAL_JTAG_IN_EP0_RD_ADDR=0xFE00 USB_SERIAL_JTAG_IN_EP0_STATE=0x3 USB_SERIAL_JTAG_IN_EP0_WR_ADDR=0x1FC | - |
| `USB_DEVICE_IN_EP1_ST` | 0x6004302C | 32 | USB_SERIAL_JTAG_IN_EP1_RD_ADDR=0xFE00 USB_SERIAL_JTAG_IN_EP1_STATE=0x3 USB_SERIAL_JTAG_IN_EP1_WR_ADDR=0x1FC | - |
| `USB_DEVICE_IN_EP2_ST` | 0x60043030 | 32 | USB_SERIAL_JTAG_IN_EP2_RD_ADDR=0xFE00 USB_SERIAL_JTAG_IN_EP2_STATE=0x3 USB_SERIAL_JTAG_IN_EP2_WR_ADDR=0x1FC | - |
| `USB_DEVICE_IN_EP3_ST` | 0x60043034 | 32 | USB_SERIAL_JTAG_IN_EP3_RD_ADDR=0xFE00 USB_SERIAL_JTAG_IN_EP3_STATE=0x3 USB_SERIAL_JTAG_IN_EP3_WR_ADDR=0x1FC | - |
| `USB_DEVICE_JFIFO_ST` | 0x60043020 | 32 | USB_SERIAL_JTAG_IN_FIFO_CNT=0x3 USB_SERIAL_JTAG_IN_FIFO_EMPTY=0x4 USB_SERIAL_JTAG_IN_FIFO_FULL=0x8 USB_SERIAL_JTAG_IN_FIFO_RESET=0x100 USB_SERIAL_JTAG_OUT_FIFO_CNT=0x30 USB_SERIAL_JTAG_OUT_FIFO_EMPTY=0x40 USB_SERIAL_JTAG_OUT_FIFO_FULL=0x80 ... | - |
| `USB_DEVICE_MEM_CONF` | 0x60043048 | 32 | USB_SERIAL_JTAG_USB_MEM_CLK_EN=0x2 USB_SERIAL_JTAG_USB_MEM_PD=0x1 | - |
| `USB_DEVICE_MISC_CONF` | 0x60043044 | 32 | USB_SERIAL_JTAG_CLK_EN=0x1 | - |
| `USB_DEVICE_OUT_EP0_ST` | 0x60043038 | 32 | USB_SERIAL_JTAG_OUT_EP0_RD_ADDR=0xFE00 USB_SERIAL_JTAG_OUT_EP0_STATE=0x3 USB_SERIAL_JTAG_OUT_EP0_WR_ADDR=0x1FC | - |
| `USB_DEVICE_OUT_EP1_ST` | 0x6004303C | 32 | USB_SERIAL_JTAG_OUT_EP1_RD_ADDR=0xFE00 USB_SERIAL_JTAG_OUT_EP1_REC_DATA_CNT=0x7F0000 USB_SERIAL_JTAG_OUT_EP1_STATE=0x3 USB_SERIAL_JTAG_OUT_EP1_WR_ADDR=0x1FC | - |
| `USB_DEVICE_OUT_EP2_ST` | 0x60043040 | 32 | USB_SERIAL_JTAG_OUT_EP2_RD_ADDR=0xFE00 USB_SERIAL_JTAG_OUT_EP2_STATE=0x3 USB_SERIAL_JTAG_OUT_EP2_WR_ADDR=0x1FC | - |
| `USB_DEVICE_SER_AFIFO_CONFIG` | 0x60043064 | 32 | USB_SERIAL_JTAG_SERIAL_IN_AFIFO_RESET_RD=0x2 USB_SERIAL_JTAG_SERIAL_IN_AFIFO_RESET_WR=0x1 USB_SERIAL_JTAG_SERIAL_IN_AFIFO_WFULL=0x20 USB_SERIAL_JTAG_SERIAL_OUT_AFIFO_REMPTY=0x10 USB_SERIAL_JTAG_SERIAL_OUT_AFIFO_RESET_RD=0x8 USB_SERIAL_JTAG_... | - |
| `USB_DEVICE_SET_LINE_CODE_W0` | 0x60043050 | 32 | USB_SERIAL_JTAG_DW_DTE_RATE=0xFFFFFFFF | - |
| `USB_DEVICE_SET_LINE_CODE_W1` | 0x60043054 | 32 | USB_SERIAL_JTAG_BCHAR_FORMAT=0xFF USB_SERIAL_JTAG_BDATA_BITS=0xFF0000 USB_SERIAL_JTAG_BPARITY_TYPE=0xFF00 | - |
| `USB_DEVICE_TEST` | 0x6004301C | 32 | USB_SERIAL_JTAG_TEST_ENABLE=0x1 USB_SERIAL_JTAG_TEST_RX_DM=0x40 USB_SERIAL_JTAG_TEST_RX_DP=0x20 USB_SERIAL_JTAG_TEST_RX_RCV=0x10 USB_SERIAL_JTAG_TEST_TX_DM=0x8 USB_SERIAL_JTAG_TEST_TX_DP=0x4 USB_SERIAL_JTAG_TEST_USB_OE=0x2 | - |
