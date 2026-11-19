# Registers - esp32s2

| Name | Address | Bits | Fields | Drives pins |
|---|---|---|---|---|
| `AES_AAD_BLOCK_NUM` | 0x6003A0A0 | 32 | AAD_BLOCK_NUM=0xFFFFFFFF | - |
| `AES_BLOCK_MODE` | 0x6003A094 | 32 | BLOCK_MODE=0x7 | - |
| `AES_BLOCK_NUM` | 0x6003A098 | 32 | BLOCK_NUM=0xFFFFFFFF | - |
| `AES_CONTINUE_OP` | 0x6003A0A8 | 32 | CONTINUE_OP=0x1 | - |
| `AES_DATE` | 0x6003A0B4 | 32 | DATE=0x3FFFFFFF | - |
| `AES_DMA_ENABLE` | 0x6003A090 | 32 | DMA_ENABLE=0x1 | - |
| `AES_DMA_EXIT` | 0x6003A0B8 | 32 | DMA_EXIT=0x1 | - |
| `AES_ENDIAN` | 0x6003A044 | 32 | ENDIAN=0x3F | - |
| `AES_H_0` | 0x6003A060 | 32 | H_0=0xFFFFFFFF | - |
| `AES_H_1` | 0x6003A064 | 32 | H_0=0xFFFFFFFF | - |
| `AES_H_2` | 0x6003A068 | 32 | H_0=0xFFFFFFFF | - |
| `AES_H_3` | 0x6003A06C | 32 | H_0=0xFFFFFFFF | - |
| `AES_INC_SEL` | 0x6003A09C | 32 | INC_SEL=0x1 | - |
| `AES_INT_CLR` | 0x6003A0AC | 32 | INT_CLR=0x1 | - |
| `AES_INT_ENA` | 0x6003A0B0 | 32 | INT_ENA=0x1 | - |
| `AES_IV_0` | 0x6003A050 | 32 | IV_0=0xFFFFFFFF | - |
| `AES_IV_1` | 0x6003A054 | 32 | IV_0=0xFFFFFFFF | - |
| `AES_IV_2` | 0x6003A058 | 32 | IV_0=0xFFFFFFFF | - |
| `AES_IV_3` | 0x6003A05C | 32 | IV_0=0xFFFFFFFF | - |
| `AES_J0_0` | 0x6003A070 | 32 | J0_0=0xFFFFFFFF | - |
| `AES_J0_1` | 0x6003A074 | 32 | J0_0=0xFFFFFFFF | - |
| `AES_J0_2` | 0x6003A078 | 32 | J0_0=0xFFFFFFFF | - |
| `AES_J0_3` | 0x6003A07C | 32 | J0_0=0xFFFFFFFF | - |
| `AES_KEY_0` | 0x6003A000 | 32 | KEY_0=0xFFFFFFFF | - |
| `AES_KEY_1` | 0x6003A004 | 32 | KEY_0=0xFFFFFFFF | - |
| `AES_KEY_2` | 0x6003A008 | 32 | KEY_0=0xFFFFFFFF | - |
| `AES_KEY_3` | 0x6003A00C | 32 | KEY_0=0xFFFFFFFF | - |
| `AES_KEY_4` | 0x6003A010 | 32 | KEY_0=0xFFFFFFFF | - |
| `AES_KEY_5` | 0x6003A014 | 32 | KEY_0=0xFFFFFFFF | - |
| `AES_KEY_6` | 0x6003A018 | 32 | KEY_0=0xFFFFFFFF | - |
| `AES_KEY_7` | 0x6003A01C | 32 | KEY_0=0xFFFFFFFF | - |
| `AES_MODE` | 0x6003A040 | 32 | MODE=0x7 | - |
| `AES_REMAINDER_BIT_NUM` | 0x6003A0A4 | 32 | REMAINDER_BIT_NUM=0x7F | - |
| `AES_STATE` | 0x6003A04C | 32 | STATE=0x3 | - |
| `AES_T0_0` | 0x6003A080 | 32 | T0_0=0xFFFFFFFF | - |
| `AES_T0_1` | 0x6003A084 | 32 | T0_0=0xFFFFFFFF | - |
| `AES_T0_2` | 0x6003A088 | 32 | T0_0=0xFFFFFFFF | - |
| `AES_T0_3` | 0x6003A08C | 32 | T0_0=0xFFFFFFFF | - |
| `AES_TEXT_IN_0` | 0x6003A020 | 32 | TEXT_IN_0=0xFFFFFFFF | - |
| `AES_TEXT_IN_1` | 0x6003A024 | 32 | TEXT_IN_0=0xFFFFFFFF | - |
| `AES_TEXT_IN_2` | 0x6003A028 | 32 | TEXT_IN_0=0xFFFFFFFF | - |
| `AES_TEXT_IN_3` | 0x6003A02C | 32 | TEXT_IN_0=0xFFFFFFFF | - |
| `AES_TEXT_OUT_0` | 0x6003A030 | 32 | TEXT_OUT_0=0xFFFFFFFF | - |
| `AES_TEXT_OUT_1` | 0x6003A034 | 32 | TEXT_OUT_0=0xFFFFFFFF | - |
| `AES_TEXT_OUT_2` | 0x6003A038 | 32 | TEXT_OUT_0=0xFFFFFFFF | - |
| `AES_TEXT_OUT_3` | 0x6003A03C | 32 | TEXT_OUT_0=0xFFFFFFFF | - |
| `AES_TRIGGER` | 0x6003A048 | 32 | TRIGGER=0x1 | - |
| `APB_SARADC_APB_CTRL_DATE` | 0x3F4403FC | 32 | APB_CTRL_DATE=0xFFFFFFFF | - |
| `APB_SARADC_APB_DAC_CTRL` | 0x3F440060 | 32 | APB_DAC_ALTER_MODE=0x2000 APB_DAC_RST=0x10000 APB_DAC_TRANS=0x4000 DAC_RESET_FIFO=0x8000 DAC_TIMER_EN=0x1000 DAC_TIMER_TARGET=0xFFF | - |
| `APB_SARADC_ARB_CTRL` | 0x3F440038 | 32 | ADC_ARB_APB_FORCE=0x4 ADC_ARB_APB_PRIORITY=0xC0 ADC_ARB_FIX_PRIORITY=0x1000 ADC_ARB_GRANT_FORCE=0x20 ADC_ARB_RTC_FORCE=0x8 ADC_ARB_RTC_PRIORITY=0x300 ADC_ARB_WIFI_FORCE=0x10 ADC_ARB_WIFI_PRIORITY=0xC00 | - |
| `APB_SARADC_CLKM_CONF` | 0x3F44005C | 32 | CLKM_DIV_A=0xFC000 CLKM_DIV_B=0x3F00 CLKM_DIV_NUM=0xFF CLK_SEL=0x600000 | - |
| `APB_SARADC_CTRL` | 0x3F440000 | 32 | DATA_SAR_SEL=0x2000000 DATA_TO_I2S=0x4000000 SAR1_PATT_LEN=0x78000 SAR1_PATT_P_CLEAR=0x800000 SAR2_PATT_LEN=0x780000 SAR2_PATT_P_CLEAR=0x1000000 SAR_CLK_DIV=0x7F80 SAR_CLK_GATED=0x40 SAR_SEL=0x20 START=0x2 START_FORCE=0x1 WAIT_ARB_CYCLE=0xC... | - |
| `APB_SARADC_CTRL2` | 0x3F440004 | 32 | MAX_MEAS_NUM=0x1FE MEAS_NUM_LIMIT=0x1 SAR1_INV=0x200 SAR2_INV=0x400 TIMER_EN=0x1000000 TIMER_SEL=0x800 TIMER_TARGET=0xFFF000 | - |
| `APB_SARADC_DMA_CONF` | 0x3F440058 | 32 | APB_ADC_EOF_NUM=0xFFFF APB_ADC_RESET_FSM=0x40000000 APB_ADC_TRANS=0x80000000 | - |
| `APB_SARADC_FILTER_CTRL` | 0x3F44003C | 32 | ADC1_FILTER_EN=0x80000000 ADC1_FILTER_FACTOR=0x3F800000 ADC1_FILTER_RESET=0x2 ADC2_FILTER_EN=0x40000000 ADC2_FILTER_FACTOR=0x7F0000 ADC2_FILTER_RESET=0x1 | - |
| `APB_SARADC_FILTER_STATUS` | 0x3F440040 | 32 | ADC1_FILTER_DATA=0xFFFF0000 ADC2_FILTER_DATA=0xFFFF | - |
| `APB_SARADC_FSM` | 0x3F440008 | 32 | SAMPLE_CYCLE=0xFF000000 SAMPLE_NUM=0xFF0000 | - |
| `APB_SARADC_FSM_WAIT` | 0x3F44000C | 32 | RSTB_WAIT=0xFF00 STANDBY_WAIT=0xFF0000 XPD_WAIT=0xFF | - |
| `APB_SARADC_INT_CLR` | 0x3F440054 | 32 | ADC1_DONE_INT_CLR=0x80000000 ADC1_THRES_INT_CLR=0x20000000 ADC2_DONE_INT_CLR=0x40000000 ADC2_THRES_INT_CLR=0x10000000 | - |
| `APB_SARADC_INT_ENA` | 0x3F440048 | 32 | ADC1_DONE_INT_ENA=0x80000000 ADC1_THRES_INT_ENA=0x20000000 ADC2_DONE_INT_ENA=0x40000000 ADC2_THRES_INT_ENA=0x10000000 | - |
| `APB_SARADC_INT_RAW` | 0x3F44004C | 32 | ADC1_DONE_INT_RAW=0x80000000 ADC1_THRES_INT_RAW=0x20000000 ADC2_DONE_INT_RAW=0x40000000 ADC2_THRES_INT_RAW=0x10000000 | - |
| `APB_SARADC_INT_ST` | 0x3F440050 | 32 | ADC1_DONE_INT_ST=0x80000000 ADC1_THRES_INT_ST=0x20000000 ADC2_DONE_INT_ST=0x40000000 ADC2_THRES_INT_ST=0x10000000 | - |
| `APB_SARADC_SAR1_PATT_TAB1` | 0x3F440018 | 32 | SAR1_PATT_TAB1=0xFFFFFFFF | - |
| `APB_SARADC_SAR1_PATT_TAB2` | 0x3F44001C | 32 | SAR1_PATT_TAB2=0xFFFFFFFF | - |
| `APB_SARADC_SAR1_PATT_TAB3` | 0x3F440020 | 32 | SAR1_PATT_TAB3=0xFFFFFFFF | - |
| `APB_SARADC_SAR1_PATT_TAB4` | 0x3F440024 | 32 | SAR1_PATT_TAB4=0xFFFFFFFF | - |
| `APB_SARADC_SAR1_STATUS` | 0x3F440010 | 32 | SAR1_STATUS=0xFFFFFFFF | - |
| `APB_SARADC_SAR2_PATT_TAB1` | 0x3F440028 | 32 | SAR2_PATT_TAB1=0xFFFFFFFF | - |
| `APB_SARADC_SAR2_PATT_TAB2` | 0x3F44002C | 32 | SAR2_PATT_TAB2=0xFFFFFFFF | - |
| `APB_SARADC_SAR2_PATT_TAB3` | 0x3F440030 | 32 | SAR2_PATT_TAB3=0xFFFFFFFF | - |
| `APB_SARADC_SAR2_PATT_TAB4` | 0x3F440034 | 32 | SAR2_PATT_TAB4=0xFFFFFFFF | - |
| `APB_SARADC_SAR2_STATUS` | 0x3F440014 | 32 | SAR2_STATUS=0xFFFFFFFF | - |
| `APB_SARADC_THRES_CTRL` | 0x3F440044 | 32 | ADC1_THRES=0x3FFE0000 ADC1_THRES_EN=0x80000000 ADC1_THRES_MODE=0x8 ADC2_THRES=0x1FFF0 ADC2_THRES_EN=0x40000000 ADC2_THRES_MODE=0x4 CLK_EN=0x1 | - |
| `DEDICATED_GPIO_INTR_CLR` | 0x3F4CF02C | 32 | GPIO0_INT_CLR=0x1 GPIO1_INT_CLR=0x2 GPIO2_INT_CLR=0x4 GPIO3_INT_CLR=0x8 GPIO4_INT_CLR=0x10 GPIO5_INT_CLR=0x20 GPIO6_INT_CLR=0x40 GPIO7_INT_CLR=0x80 | - |
| `DEDICATED_GPIO_INTR_RAW` | 0x3F4CF020 | 32 | GPIO0=0x1 GPIO1=0x2 GPIO2=0x4 GPIO3=0x8 GPIO4=0x10 GPIO5=0x20 GPIO6=0x40 GPIO7=0x80 | - |
| `DEDICATED_GPIO_INTR_RCGN` | 0x3F4CF01C | 32 | INTR_MODE_CH0=0x7 INTR_MODE_CH1=0x38 INTR_MODE_CH2=0x1C0 INTR_MODE_CH3=0xE00 INTR_MODE_CH4=0x7000 INTR_MODE_CH5=0x38000 INTR_MODE_CH6=0x1C0000 INTR_MODE_CH7=0xE00000 | - |
| `DEDICATED_GPIO_INTR_RLS` | 0x3F4CF024 | 32 | GPIO0_INT_ENA=0x1 GPIO1_INT_ENA=0x2 GPIO2_INT_ENA=0x4 GPIO3_INT_ENA=0x8 GPIO4_INT_ENA=0x10 GPIO5_INT_ENA=0x20 GPIO6_INT_ENA=0x40 GPIO7_INT_ENA=0x80 | - |
| `DEDICATED_GPIO_INTR_ST` | 0x3F4CF028 | 32 | GPIO0_INT_ST=0x1 GPIO1_INT_ST=0x2 GPIO2_INT_ST=0x4 GPIO3_INT_ST=0x8 GPIO4_INT_ST=0x10 GPIO5_INT_ST=0x20 GPIO6_INT_ST=0x40 GPIO7_INT_ST=0x80 | - |
| `DEDICATED_GPIO_IN_DLY` | 0x3F4CF014 | 32 | CH0=0x3 CH1=0xC CH2=0x30 CH3=0xC0 CH4=0x300 CH5=0xC00 CH6=0x3000 CH7=0xC000 | - |
| `DEDICATED_GPIO_IN_SCAN` | 0x3F4CF018 | 32 | IN_STATUS=0xFF | - |
| `DEDICATED_GPIO_OUT_CPU` | 0x3F4CF010 | 32 | SEL0=0x1 SEL1=0x2 SEL2=0x4 SEL3=0x8 SEL4=0x10 SEL5=0x20 SEL6=0x40 SEL7=0x80 | - |
| `DEDICATED_GPIO_OUT_DRT` | 0x3F4CF000 | 32 | VLAUE=0xFF | - |
| `DEDICATED_GPIO_OUT_IDV` | 0x3F4CF008 | 32 | CH0=0x3 CH1=0xC CH2=0x30 CH3=0xC0 CH4=0x300 CH5=0xC00 CH6=0x3000 CH7=0xC000 | - |
| `DEDICATED_GPIO_OUT_MSK` | 0x3F4CF004 | 32 | OUT_MSK=0xFF00 OUT_VALUE=0xFF | - |
| `DEDICATED_GPIO_OUT_SCAN` | 0x3F4CF00C | 32 | OUT_STATUS=0xFF | - |
| `DS_C_MEM[0]` | 0x6003D000 | 8 |  | - |
| `DS_C_MEM[1000]` | 0x6003D3E8 | 8 |  | - |
| `DS_C_MEM[1001]` | 0x6003D3E9 | 8 |  | - |
| `DS_C_MEM[1002]` | 0x6003D3EA | 8 |  | - |
| `DS_C_MEM[1003]` | 0x6003D3EB | 8 |  | - |
| `DS_C_MEM[1004]` | 0x6003D3EC | 8 |  | - |
| `DS_C_MEM[1005]` | 0x6003D3ED | 8 |  | - |
| `DS_C_MEM[1006]` | 0x6003D3EE | 8 |  | - |
| `DS_C_MEM[1007]` | 0x6003D3EF | 8 |  | - |
| `DS_C_MEM[1008]` | 0x6003D3F0 | 8 |  | - |
| `DS_C_MEM[1009]` | 0x6003D3F1 | 8 |  | - |
| `DS_C_MEM[100]` | 0x6003D064 | 8 |  | - |
| `DS_C_MEM[1010]` | 0x6003D3F2 | 8 |  | - |
| `DS_C_MEM[1011]` | 0x6003D3F3 | 8 |  | - |
| `DS_C_MEM[1012]` | 0x6003D3F4 | 8 |  | - |
| `DS_C_MEM[1013]` | 0x6003D3F5 | 8 |  | - |
| `DS_C_MEM[1014]` | 0x6003D3F6 | 8 |  | - |
| `DS_C_MEM[1015]` | 0x6003D3F7 | 8 |  | - |
| `DS_C_MEM[1016]` | 0x6003D3F8 | 8 |  | - |
| `DS_C_MEM[1017]` | 0x6003D3F9 | 8 |  | - |
| `DS_C_MEM[1018]` | 0x6003D3FA | 8 |  | - |
| `DS_C_MEM[1019]` | 0x6003D3FB | 8 |  | - |
| `DS_C_MEM[101]` | 0x6003D065 | 8 |  | - |
| `DS_C_MEM[1020]` | 0x6003D3FC | 8 |  | - |
| `DS_C_MEM[1021]` | 0x6003D3FD | 8 |  | - |
| `DS_C_MEM[1022]` | 0x6003D3FE | 8 |  | - |
| `DS_C_MEM[1023]` | 0x6003D3FF | 8 |  | - |
| `DS_C_MEM[1024]` | 0x6003D400 | 8 |  | - |
| `DS_C_MEM[1025]` | 0x6003D401 | 8 |  | - |
| `DS_C_MEM[1026]` | 0x6003D402 | 8 |  | - |
| `DS_C_MEM[1027]` | 0x6003D403 | 8 |  | - |
| `DS_C_MEM[1028]` | 0x6003D404 | 8 |  | - |
| `DS_C_MEM[1029]` | 0x6003D405 | 8 |  | - |
| `DS_C_MEM[102]` | 0x6003D066 | 8 |  | - |
| `DS_C_MEM[1030]` | 0x6003D406 | 8 |  | - |
| `DS_C_MEM[1031]` | 0x6003D407 | 8 |  | - |
| `DS_C_MEM[1032]` | 0x6003D408 | 8 |  | - |
| `DS_C_MEM[1033]` | 0x6003D409 | 8 |  | - |
| `DS_C_MEM[1034]` | 0x6003D40A | 8 |  | - |
| `DS_C_MEM[1035]` | 0x6003D40B | 8 |  | - |
| `DS_C_MEM[1036]` | 0x6003D40C | 8 |  | - |
| `DS_C_MEM[1037]` | 0x6003D40D | 8 |  | - |
| `DS_C_MEM[1038]` | 0x6003D40E | 8 |  | - |
| `DS_C_MEM[1039]` | 0x6003D40F | 8 |  | - |
| `DS_C_MEM[103]` | 0x6003D067 | 8 |  | - |
| `DS_C_MEM[1040]` | 0x6003D410 | 8 |  | - |
| `DS_C_MEM[1041]` | 0x6003D411 | 8 |  | - |
| `DS_C_MEM[1042]` | 0x6003D412 | 8 |  | - |
| `DS_C_MEM[1043]` | 0x6003D413 | 8 |  | - |
| `DS_C_MEM[1044]` | 0x6003D414 | 8 |  | - |
| `DS_C_MEM[1045]` | 0x6003D415 | 8 |  | - |
| `DS_C_MEM[1046]` | 0x6003D416 | 8 |  | - |
| `DS_C_MEM[1047]` | 0x6003D417 | 8 |  | - |
| `DS_C_MEM[1048]` | 0x6003D418 | 8 |  | - |
| `DS_C_MEM[1049]` | 0x6003D419 | 8 |  | - |
| `DS_C_MEM[104]` | 0x6003D068 | 8 |  | - |
| `DS_C_MEM[1050]` | 0x6003D41A | 8 |  | - |
| `DS_C_MEM[1051]` | 0x6003D41B | 8 |  | - |
| `DS_C_MEM[1052]` | 0x6003D41C | 8 |  | - |
| `DS_C_MEM[1053]` | 0x6003D41D | 8 |  | - |
| `DS_C_MEM[1054]` | 0x6003D41E | 8 |  | - |
| `DS_C_MEM[1055]` | 0x6003D41F | 8 |  | - |
| `DS_C_MEM[1056]` | 0x6003D420 | 8 |  | - |
| `DS_C_MEM[1057]` | 0x6003D421 | 8 |  | - |
| `DS_C_MEM[1058]` | 0x6003D422 | 8 |  | - |
| `DS_C_MEM[1059]` | 0x6003D423 | 8 |  | - |
| `DS_C_MEM[105]` | 0x6003D069 | 8 |  | - |
| `DS_C_MEM[1060]` | 0x6003D424 | 8 |  | - |
| `DS_C_MEM[1061]` | 0x6003D425 | 8 |  | - |
| `DS_C_MEM[1062]` | 0x6003D426 | 8 |  | - |
| `DS_C_MEM[1063]` | 0x6003D427 | 8 |  | - |
| `DS_C_MEM[1064]` | 0x6003D428 | 8 |  | - |
| `DS_C_MEM[1065]` | 0x6003D429 | 8 |  | - |
| `DS_C_MEM[1066]` | 0x6003D42A | 8 |  | - |
| `DS_C_MEM[1067]` | 0x6003D42B | 8 |  | - |
| `DS_C_MEM[1068]` | 0x6003D42C | 8 |  | - |
| `DS_C_MEM[1069]` | 0x6003D42D | 8 |  | - |
| `DS_C_MEM[106]` | 0x6003D06A | 8 |  | - |
| `DS_C_MEM[1070]` | 0x6003D42E | 8 |  | - |
| `DS_C_MEM[1071]` | 0x6003D42F | 8 |  | - |
| `DS_C_MEM[1072]` | 0x6003D430 | 8 |  | - |
| `DS_C_MEM[1073]` | 0x6003D431 | 8 |  | - |
| `DS_C_MEM[1074]` | 0x6003D432 | 8 |  | - |
| `DS_C_MEM[1075]` | 0x6003D433 | 8 |  | - |
| `DS_C_MEM[1076]` | 0x6003D434 | 8 |  | - |
| `DS_C_MEM[1077]` | 0x6003D435 | 8 |  | - |
| `DS_C_MEM[1078]` | 0x6003D436 | 8 |  | - |
| `DS_C_MEM[1079]` | 0x6003D437 | 8 |  | - |
| `DS_C_MEM[107]` | 0x6003D06B | 8 |  | - |
| `DS_C_MEM[1080]` | 0x6003D438 | 8 |  | - |
| `DS_C_MEM[1081]` | 0x6003D439 | 8 |  | - |
| `DS_C_MEM[1082]` | 0x6003D43A | 8 |  | - |
| `DS_C_MEM[1083]` | 0x6003D43B | 8 |  | - |
| `DS_C_MEM[1084]` | 0x6003D43C | 8 |  | - |
| `DS_C_MEM[1085]` | 0x6003D43D | 8 |  | - |
| `DS_C_MEM[1086]` | 0x6003D43E | 8 |  | - |
| `DS_C_MEM[1087]` | 0x6003D43F | 8 |  | - |
| `DS_C_MEM[1088]` | 0x6003D440 | 8 |  | - |
| `DS_C_MEM[1089]` | 0x6003D441 | 8 |  | - |
| `DS_C_MEM[108]` | 0x6003D06C | 8 |  | - |
| `DS_C_MEM[1090]` | 0x6003D442 | 8 |  | - |
| `DS_C_MEM[1091]` | 0x6003D443 | 8 |  | - |
| `DS_C_MEM[1092]` | 0x6003D444 | 8 |  | - |
| `DS_C_MEM[1093]` | 0x6003D445 | 8 |  | - |
| `DS_C_MEM[1094]` | 0x6003D446 | 8 |  | - |
| `DS_C_MEM[1095]` | 0x6003D447 | 8 |  | - |
| `DS_C_MEM[1096]` | 0x6003D448 | 8 |  | - |
| `DS_C_MEM[1097]` | 0x6003D449 | 8 |  | - |
| `DS_C_MEM[1098]` | 0x6003D44A | 8 |  | - |
| `DS_C_MEM[1099]` | 0x6003D44B | 8 |  | - |
| `DS_C_MEM[109]` | 0x6003D06D | 8 |  | - |
| `DS_C_MEM[10]` | 0x6003D00A | 8 |  | - |
| `DS_C_MEM[1100]` | 0x6003D44C | 8 |  | - |
| `DS_C_MEM[1101]` | 0x6003D44D | 8 |  | - |
| `DS_C_MEM[1102]` | 0x6003D44E | 8 |  | - |
| `DS_C_MEM[1103]` | 0x6003D44F | 8 |  | - |
| `DS_C_MEM[1104]` | 0x6003D450 | 8 |  | - |
| `DS_C_MEM[1105]` | 0x6003D451 | 8 |  | - |
| `DS_C_MEM[1106]` | 0x6003D452 | 8 |  | - |
| `DS_C_MEM[1107]` | 0x6003D453 | 8 |  | - |
| `DS_C_MEM[1108]` | 0x6003D454 | 8 |  | - |
| `DS_C_MEM[1109]` | 0x6003D455 | 8 |  | - |
| `DS_C_MEM[110]` | 0x6003D06E | 8 |  | - |
| `DS_C_MEM[1110]` | 0x6003D456 | 8 |  | - |
| `DS_C_MEM[1111]` | 0x6003D457 | 8 |  | - |
| `DS_C_MEM[1112]` | 0x6003D458 | 8 |  | - |
| `DS_C_MEM[1113]` | 0x6003D459 | 8 |  | - |
| `DS_C_MEM[1114]` | 0x6003D45A | 8 |  | - |
| `DS_C_MEM[1115]` | 0x6003D45B | 8 |  | - |
| `DS_C_MEM[1116]` | 0x6003D45C | 8 |  | - |
| `DS_C_MEM[1117]` | 0x6003D45D | 8 |  | - |
| `DS_C_MEM[1118]` | 0x6003D45E | 8 |  | - |
| `DS_C_MEM[1119]` | 0x6003D45F | 8 |  | - |
| `DS_C_MEM[111]` | 0x6003D06F | 8 |  | - |
| `DS_C_MEM[1120]` | 0x6003D460 | 8 |  | - |
| `DS_C_MEM[1121]` | 0x6003D461 | 8 |  | - |
| `DS_C_MEM[1122]` | 0x6003D462 | 8 |  | - |
| `DS_C_MEM[1123]` | 0x6003D463 | 8 |  | - |
| `DS_C_MEM[1124]` | 0x6003D464 | 8 |  | - |
| `DS_C_MEM[1125]` | 0x6003D465 | 8 |  | - |
| `DS_C_MEM[1126]` | 0x6003D466 | 8 |  | - |
| `DS_C_MEM[1127]` | 0x6003D467 | 8 |  | - |
| `DS_C_MEM[1128]` | 0x6003D468 | 8 |  | - |
| `DS_C_MEM[1129]` | 0x6003D469 | 8 |  | - |
| `DS_C_MEM[112]` | 0x6003D070 | 8 |  | - |
| `DS_C_MEM[1130]` | 0x6003D46A | 8 |  | - |
| `DS_C_MEM[1131]` | 0x6003D46B | 8 |  | - |
| `DS_C_MEM[1132]` | 0x6003D46C | 8 |  | - |
| `DS_C_MEM[1133]` | 0x6003D46D | 8 |  | - |
| `DS_C_MEM[1134]` | 0x6003D46E | 8 |  | - |
| `DS_C_MEM[1135]` | 0x6003D46F | 8 |  | - |
| `DS_C_MEM[1136]` | 0x6003D470 | 8 |  | - |
| `DS_C_MEM[1137]` | 0x6003D471 | 8 |  | - |
| `DS_C_MEM[1138]` | 0x6003D472 | 8 |  | - |
| `DS_C_MEM[1139]` | 0x6003D473 | 8 |  | - |
| `DS_C_MEM[113]` | 0x6003D071 | 8 |  | - |
| `DS_C_MEM[1140]` | 0x6003D474 | 8 |  | - |
| `DS_C_MEM[1141]` | 0x6003D475 | 8 |  | - |
| `DS_C_MEM[1142]` | 0x6003D476 | 8 |  | - |
| `DS_C_MEM[1143]` | 0x6003D477 | 8 |  | - |
| `DS_C_MEM[1144]` | 0x6003D478 | 8 |  | - |
| `DS_C_MEM[1145]` | 0x6003D479 | 8 |  | - |
| `DS_C_MEM[1146]` | 0x6003D47A | 8 |  | - |
| `DS_C_MEM[1147]` | 0x6003D47B | 8 |  | - |
| `DS_C_MEM[1148]` | 0x6003D47C | 8 |  | - |
| `DS_C_MEM[1149]` | 0x6003D47D | 8 |  | - |
| `DS_C_MEM[114]` | 0x6003D072 | 8 |  | - |
| `DS_C_MEM[1150]` | 0x6003D47E | 8 |  | - |
| `DS_C_MEM[1151]` | 0x6003D47F | 8 |  | - |
| `DS_C_MEM[1152]` | 0x6003D480 | 8 |  | - |
| `DS_C_MEM[1153]` | 0x6003D481 | 8 |  | - |
| `DS_C_MEM[1154]` | 0x6003D482 | 8 |  | - |
| `DS_C_MEM[1155]` | 0x6003D483 | 8 |  | - |
| `DS_C_MEM[1156]` | 0x6003D484 | 8 |  | - |
| `DS_C_MEM[1157]` | 0x6003D485 | 8 |  | - |
| `DS_C_MEM[1158]` | 0x6003D486 | 8 |  | - |
| `DS_C_MEM[1159]` | 0x6003D487 | 8 |  | - |
| `DS_C_MEM[115]` | 0x6003D073 | 8 |  | - |
| `DS_C_MEM[1160]` | 0x6003D488 | 8 |  | - |
| `DS_C_MEM[1161]` | 0x6003D489 | 8 |  | - |
| `DS_C_MEM[1162]` | 0x6003D48A | 8 |  | - |
| `DS_C_MEM[1163]` | 0x6003D48B | 8 |  | - |
| `DS_C_MEM[1164]` | 0x6003D48C | 8 |  | - |
| `DS_C_MEM[1165]` | 0x6003D48D | 8 |  | - |
| `DS_C_MEM[1166]` | 0x6003D48E | 8 |  | - |
| `DS_C_MEM[1167]` | 0x6003D48F | 8 |  | - |
| `DS_C_MEM[1168]` | 0x6003D490 | 8 |  | - |
| `DS_C_MEM[1169]` | 0x6003D491 | 8 |  | - |
| `DS_C_MEM[116]` | 0x6003D074 | 8 |  | - |
| `DS_C_MEM[1170]` | 0x6003D492 | 8 |  | - |
| `DS_C_MEM[1171]` | 0x6003D493 | 8 |  | - |
| `DS_C_MEM[1172]` | 0x6003D494 | 8 |  | - |
| `DS_C_MEM[1173]` | 0x6003D495 | 8 |  | - |
| `DS_C_MEM[1174]` | 0x6003D496 | 8 |  | - |
| `DS_C_MEM[1175]` | 0x6003D497 | 8 |  | - |
| `DS_C_MEM[1176]` | 0x6003D498 | 8 |  | - |
| `DS_C_MEM[1177]` | 0x6003D499 | 8 |  | - |
| `DS_C_MEM[1178]` | 0x6003D49A | 8 |  | - |
| `DS_C_MEM[1179]` | 0x6003D49B | 8 |  | - |
| `DS_C_MEM[117]` | 0x6003D075 | 8 |  | - |
| `DS_C_MEM[1180]` | 0x6003D49C | 8 |  | - |
| `DS_C_MEM[1181]` | 0x6003D49D | 8 |  | - |
| `DS_C_MEM[1182]` | 0x6003D49E | 8 |  | - |
| `DS_C_MEM[1183]` | 0x6003D49F | 8 |  | - |
| `DS_C_MEM[1184]` | 0x6003D4A0 | 8 |  | - |
| `DS_C_MEM[1185]` | 0x6003D4A1 | 8 |  | - |
| `DS_C_MEM[1186]` | 0x6003D4A2 | 8 |  | - |
| `DS_C_MEM[1187]` | 0x6003D4A3 | 8 |  | - |
| `DS_C_MEM[1188]` | 0x6003D4A4 | 8 |  | - |
| `DS_C_MEM[1189]` | 0x6003D4A5 | 8 |  | - |
| `DS_C_MEM[118]` | 0x6003D076 | 8 |  | - |
| `DS_C_MEM[1190]` | 0x6003D4A6 | 8 |  | - |
| `DS_C_MEM[1191]` | 0x6003D4A7 | 8 |  | - |
| `DS_C_MEM[1192]` | 0x6003D4A8 | 8 |  | - |
| `DS_C_MEM[1193]` | 0x6003D4A9 | 8 |  | - |
| `DS_C_MEM[1194]` | 0x6003D4AA | 8 |  | - |
| `DS_C_MEM[1195]` | 0x6003D4AB | 8 |  | - |
| `DS_C_MEM[1196]` | 0x6003D4AC | 8 |  | - |
| `DS_C_MEM[1197]` | 0x6003D4AD | 8 |  | - |
| `DS_C_MEM[1198]` | 0x6003D4AE | 8 |  | - |
| `DS_C_MEM[1199]` | 0x6003D4AF | 8 |  | - |
| `DS_C_MEM[119]` | 0x6003D077 | 8 |  | - |
| `DS_C_MEM[11]` | 0x6003D00B | 8 |  | - |
| `DS_C_MEM[1200]` | 0x6003D4B0 | 8 |  | - |
| `DS_C_MEM[1201]` | 0x6003D4B1 | 8 |  | - |
| `DS_C_MEM[1202]` | 0x6003D4B2 | 8 |  | - |
| `DS_C_MEM[1203]` | 0x6003D4B3 | 8 |  | - |
| `DS_C_MEM[1204]` | 0x6003D4B4 | 8 |  | - |
| `DS_C_MEM[1205]` | 0x6003D4B5 | 8 |  | - |
| `DS_C_MEM[1206]` | 0x6003D4B6 | 8 |  | - |
| `DS_C_MEM[1207]` | 0x6003D4B7 | 8 |  | - |
| `DS_C_MEM[1208]` | 0x6003D4B8 | 8 |  | - |
| `DS_C_MEM[1209]` | 0x6003D4B9 | 8 |  | - |
| `DS_C_MEM[120]` | 0x6003D078 | 8 |  | - |
| `DS_C_MEM[1210]` | 0x6003D4BA | 8 |  | - |
| `DS_C_MEM[1211]` | 0x6003D4BB | 8 |  | - |
| `DS_C_MEM[1212]` | 0x6003D4BC | 8 |  | - |
| `DS_C_MEM[1213]` | 0x6003D4BD | 8 |  | - |
| `DS_C_MEM[1214]` | 0x6003D4BE | 8 |  | - |
| `DS_C_MEM[1215]` | 0x6003D4BF | 8 |  | - |
| `DS_C_MEM[1216]` | 0x6003D4C0 | 8 |  | - |
| `DS_C_MEM[1217]` | 0x6003D4C1 | 8 |  | - |
| `DS_C_MEM[1218]` | 0x6003D4C2 | 8 |  | - |
| `DS_C_MEM[1219]` | 0x6003D4C3 | 8 |  | - |
| `DS_C_MEM[121]` | 0x6003D079 | 8 |  | - |
| `DS_C_MEM[1220]` | 0x6003D4C4 | 8 |  | - |
| `DS_C_MEM[1221]` | 0x6003D4C5 | 8 |  | - |
| `DS_C_MEM[1222]` | 0x6003D4C6 | 8 |  | - |
| `DS_C_MEM[1223]` | 0x6003D4C7 | 8 |  | - |
| `DS_C_MEM[1224]` | 0x6003D4C8 | 8 |  | - |
| `DS_C_MEM[1225]` | 0x6003D4C9 | 8 |  | - |
| `DS_C_MEM[1226]` | 0x6003D4CA | 8 |  | - |
| `DS_C_MEM[1227]` | 0x6003D4CB | 8 |  | - |
| `DS_C_MEM[1228]` | 0x6003D4CC | 8 |  | - |
| `DS_C_MEM[1229]` | 0x6003D4CD | 8 |  | - |
| `DS_C_MEM[122]` | 0x6003D07A | 8 |  | - |
| `DS_C_MEM[1230]` | 0x6003D4CE | 8 |  | - |
| `DS_C_MEM[1231]` | 0x6003D4CF | 8 |  | - |
| `DS_C_MEM[1232]` | 0x6003D4D0 | 8 |  | - |
| `DS_C_MEM[1233]` | 0x6003D4D1 | 8 |  | - |
| `DS_C_MEM[1234]` | 0x6003D4D2 | 8 |  | - |
| `DS_C_MEM[1235]` | 0x6003D4D3 | 8 |  | - |
| `DS_C_MEM[1236]` | 0x6003D4D4 | 8 |  | - |
| `DS_C_MEM[1237]` | 0x6003D4D5 | 8 |  | - |
| `DS_C_MEM[1238]` | 0x6003D4D6 | 8 |  | - |
| `DS_C_MEM[1239]` | 0x6003D4D7 | 8 |  | - |
| `DS_C_MEM[123]` | 0x6003D07B | 8 |  | - |
| `DS_C_MEM[1240]` | 0x6003D4D8 | 8 |  | - |
| `DS_C_MEM[1241]` | 0x6003D4D9 | 8 |  | - |
| `DS_C_MEM[1242]` | 0x6003D4DA | 8 |  | - |
| `DS_C_MEM[1243]` | 0x6003D4DB | 8 |  | - |
| `DS_C_MEM[1244]` | 0x6003D4DC | 8 |  | - |
| `DS_C_MEM[1245]` | 0x6003D4DD | 8 |  | - |
| `DS_C_MEM[1246]` | 0x6003D4DE | 8 |  | - |
| `DS_C_MEM[1247]` | 0x6003D4DF | 8 |  | - |
| `DS_C_MEM[1248]` | 0x6003D4E0 | 8 |  | - |
| `DS_C_MEM[1249]` | 0x6003D4E1 | 8 |  | - |
| `DS_C_MEM[124]` | 0x6003D07C | 8 |  | - |
| `DS_C_MEM[1250]` | 0x6003D4E2 | 8 |  | - |
| `DS_C_MEM[1251]` | 0x6003D4E3 | 8 |  | - |
| `DS_C_MEM[1252]` | 0x6003D4E4 | 8 |  | - |
| `DS_C_MEM[1253]` | 0x6003D4E5 | 8 |  | - |
| `DS_C_MEM[1254]` | 0x6003D4E6 | 8 |  | - |
| `DS_C_MEM[1255]` | 0x6003D4E7 | 8 |  | - |
| `DS_C_MEM[1256]` | 0x6003D4E8 | 8 |  | - |
| `DS_C_MEM[1257]` | 0x6003D4E9 | 8 |  | - |
| `DS_C_MEM[1258]` | 0x6003D4EA | 8 |  | - |
| `DS_C_MEM[1259]` | 0x6003D4EB | 8 |  | - |
| `DS_C_MEM[125]` | 0x6003D07D | 8 |  | - |
| `DS_C_MEM[1260]` | 0x6003D4EC | 8 |  | - |
| `DS_C_MEM[1261]` | 0x6003D4ED | 8 |  | - |
| `DS_C_MEM[1262]` | 0x6003D4EE | 8 |  | - |
| `DS_C_MEM[1263]` | 0x6003D4EF | 8 |  | - |
| `DS_C_MEM[1264]` | 0x6003D4F0 | 8 |  | - |
| `DS_C_MEM[1265]` | 0x6003D4F1 | 8 |  | - |
| `DS_C_MEM[1266]` | 0x6003D4F2 | 8 |  | - |
| `DS_C_MEM[1267]` | 0x6003D4F3 | 8 |  | - |
| `DS_C_MEM[1268]` | 0x6003D4F4 | 8 |  | - |
| `DS_C_MEM[1269]` | 0x6003D4F5 | 8 |  | - |
| `DS_C_MEM[126]` | 0x6003D07E | 8 |  | - |
| `DS_C_MEM[1270]` | 0x6003D4F6 | 8 |  | - |
| `DS_C_MEM[1271]` | 0x6003D4F7 | 8 |  | - |
| `DS_C_MEM[1272]` | 0x6003D4F8 | 8 |  | - |
| `DS_C_MEM[1273]` | 0x6003D4F9 | 8 |  | - |
| `DS_C_MEM[1274]` | 0x6003D4FA | 8 |  | - |
| `DS_C_MEM[1275]` | 0x6003D4FB | 8 |  | - |
| `DS_C_MEM[1276]` | 0x6003D4FC | 8 |  | - |
| `DS_C_MEM[1277]` | 0x6003D4FD | 8 |  | - |
| `DS_C_MEM[1278]` | 0x6003D4FE | 8 |  | - |
| `DS_C_MEM[1279]` | 0x6003D4FF | 8 |  | - |
| `DS_C_MEM[127]` | 0x6003D07F | 8 |  | - |
| `DS_C_MEM[1280]` | 0x6003D500 | 8 |  | - |
| `DS_C_MEM[1281]` | 0x6003D501 | 8 |  | - |
| `DS_C_MEM[1282]` | 0x6003D502 | 8 |  | - |
| `DS_C_MEM[1283]` | 0x6003D503 | 8 |  | - |
| `DS_C_MEM[1284]` | 0x6003D504 | 8 |  | - |
| `DS_C_MEM[1285]` | 0x6003D505 | 8 |  | - |
| `DS_C_MEM[1286]` | 0x6003D506 | 8 |  | - |
| `DS_C_MEM[1287]` | 0x6003D507 | 8 |  | - |
| `DS_C_MEM[1288]` | 0x6003D508 | 8 |  | - |
| `DS_C_MEM[1289]` | 0x6003D509 | 8 |  | - |
| `DS_C_MEM[128]` | 0x6003D080 | 8 |  | - |
| `DS_C_MEM[1290]` | 0x6003D50A | 8 |  | - |
| `DS_C_MEM[1291]` | 0x6003D50B | 8 |  | - |
| `DS_C_MEM[1292]` | 0x6003D50C | 8 |  | - |
| `DS_C_MEM[1293]` | 0x6003D50D | 8 |  | - |
| `DS_C_MEM[1294]` | 0x6003D50E | 8 |  | - |
| `DS_C_MEM[1295]` | 0x6003D50F | 8 |  | - |
| `DS_C_MEM[1296]` | 0x6003D510 | 8 |  | - |
| `DS_C_MEM[1297]` | 0x6003D511 | 8 |  | - |
| `DS_C_MEM[1298]` | 0x6003D512 | 8 |  | - |
| `DS_C_MEM[1299]` | 0x6003D513 | 8 |  | - |
| `DS_C_MEM[129]` | 0x6003D081 | 8 |  | - |
| `DS_C_MEM[12]` | 0x6003D00C | 8 |  | - |
| `DS_C_MEM[1300]` | 0x6003D514 | 8 |  | - |
| `DS_C_MEM[1301]` | 0x6003D515 | 8 |  | - |
| `DS_C_MEM[1302]` | 0x6003D516 | 8 |  | - |
| `DS_C_MEM[1303]` | 0x6003D517 | 8 |  | - |
| `DS_C_MEM[1304]` | 0x6003D518 | 8 |  | - |
| `DS_C_MEM[1305]` | 0x6003D519 | 8 |  | - |
| `DS_C_MEM[1306]` | 0x6003D51A | 8 |  | - |
| `DS_C_MEM[1307]` | 0x6003D51B | 8 |  | - |
| `DS_C_MEM[1308]` | 0x6003D51C | 8 |  | - |
| `DS_C_MEM[1309]` | 0x6003D51D | 8 |  | - |
| `DS_C_MEM[130]` | 0x6003D082 | 8 |  | - |
| `DS_C_MEM[1310]` | 0x6003D51E | 8 |  | - |
| `DS_C_MEM[1311]` | 0x6003D51F | 8 |  | - |
| `DS_C_MEM[1312]` | 0x6003D520 | 8 |  | - |
| `DS_C_MEM[1313]` | 0x6003D521 | 8 |  | - |
| `DS_C_MEM[1314]` | 0x6003D522 | 8 |  | - |
| `DS_C_MEM[1315]` | 0x6003D523 | 8 |  | - |
| `DS_C_MEM[1316]` | 0x6003D524 | 8 |  | - |
| `DS_C_MEM[1317]` | 0x6003D525 | 8 |  | - |
| `DS_C_MEM[1318]` | 0x6003D526 | 8 |  | - |
| `DS_C_MEM[1319]` | 0x6003D527 | 8 |  | - |
| `DS_C_MEM[131]` | 0x6003D083 | 8 |  | - |
| `DS_C_MEM[1320]` | 0x6003D528 | 8 |  | - |
| `DS_C_MEM[1321]` | 0x6003D529 | 8 |  | - |
| `DS_C_MEM[1322]` | 0x6003D52A | 8 |  | - |
| `DS_C_MEM[1323]` | 0x6003D52B | 8 |  | - |
| `DS_C_MEM[1324]` | 0x6003D52C | 8 |  | - |
| `DS_C_MEM[1325]` | 0x6003D52D | 8 |  | - |
| `DS_C_MEM[1326]` | 0x6003D52E | 8 |  | - |
| `DS_C_MEM[1327]` | 0x6003D52F | 8 |  | - |
| `DS_C_MEM[1328]` | 0x6003D530 | 8 |  | - |
| `DS_C_MEM[1329]` | 0x6003D531 | 8 |  | - |
| `DS_C_MEM[132]` | 0x6003D084 | 8 |  | - |
| `DS_C_MEM[1330]` | 0x6003D532 | 8 |  | - |
| `DS_C_MEM[1331]` | 0x6003D533 | 8 |  | - |
| `DS_C_MEM[1332]` | 0x6003D534 | 8 |  | - |
| `DS_C_MEM[1333]` | 0x6003D535 | 8 |  | - |
| `DS_C_MEM[1334]` | 0x6003D536 | 8 |  | - |
| `DS_C_MEM[1335]` | 0x6003D537 | 8 |  | - |
| `DS_C_MEM[1336]` | 0x6003D538 | 8 |  | - |
| `DS_C_MEM[1337]` | 0x6003D539 | 8 |  | - |
| `DS_C_MEM[1338]` | 0x6003D53A | 8 |  | - |
| `DS_C_MEM[1339]` | 0x6003D53B | 8 |  | - |
| `DS_C_MEM[133]` | 0x6003D085 | 8 |  | - |
| `DS_C_MEM[1340]` | 0x6003D53C | 8 |  | - |
| `DS_C_MEM[1341]` | 0x6003D53D | 8 |  | - |
| `DS_C_MEM[1342]` | 0x6003D53E | 8 |  | - |
| `DS_C_MEM[1343]` | 0x6003D53F | 8 |  | - |
| `DS_C_MEM[1344]` | 0x6003D540 | 8 |  | - |
| `DS_C_MEM[1345]` | 0x6003D541 | 8 |  | - |
| `DS_C_MEM[1346]` | 0x6003D542 | 8 |  | - |
| `DS_C_MEM[1347]` | 0x6003D543 | 8 |  | - |
| `DS_C_MEM[1348]` | 0x6003D544 | 8 |  | - |
| `DS_C_MEM[1349]` | 0x6003D545 | 8 |  | - |
| `DS_C_MEM[134]` | 0x6003D086 | 8 |  | - |
| `DS_C_MEM[1350]` | 0x6003D546 | 8 |  | - |
| `DS_C_MEM[1351]` | 0x6003D547 | 8 |  | - |
| `DS_C_MEM[1352]` | 0x6003D548 | 8 |  | - |
| `DS_C_MEM[1353]` | 0x6003D549 | 8 |  | - |
| `DS_C_MEM[1354]` | 0x6003D54A | 8 |  | - |
| `DS_C_MEM[1355]` | 0x6003D54B | 8 |  | - |
| `DS_C_MEM[1356]` | 0x6003D54C | 8 |  | - |
| `DS_C_MEM[1357]` | 0x6003D54D | 8 |  | - |
| `DS_C_MEM[1358]` | 0x6003D54E | 8 |  | - |
| `DS_C_MEM[1359]` | 0x6003D54F | 8 |  | - |
| `DS_C_MEM[135]` | 0x6003D087 | 8 |  | - |
| `DS_C_MEM[1360]` | 0x6003D550 | 8 |  | - |
| `DS_C_MEM[1361]` | 0x6003D551 | 8 |  | - |
| `DS_C_MEM[1362]` | 0x6003D552 | 8 |  | - |
| `DS_C_MEM[1363]` | 0x6003D553 | 8 |  | - |
| `DS_C_MEM[1364]` | 0x6003D554 | 8 |  | - |
| `DS_C_MEM[1365]` | 0x6003D555 | 8 |  | - |
| `DS_C_MEM[1366]` | 0x6003D556 | 8 |  | - |
| `DS_C_MEM[1367]` | 0x6003D557 | 8 |  | - |
| `DS_C_MEM[1368]` | 0x6003D558 | 8 |  | - |
| `DS_C_MEM[1369]` | 0x6003D559 | 8 |  | - |
| `DS_C_MEM[136]` | 0x6003D088 | 8 |  | - |
| `DS_C_MEM[1370]` | 0x6003D55A | 8 |  | - |
| `DS_C_MEM[1371]` | 0x6003D55B | 8 |  | - |
| `DS_C_MEM[1372]` | 0x6003D55C | 8 |  | - |
| `DS_C_MEM[1373]` | 0x6003D55D | 8 |  | - |
| `DS_C_MEM[1374]` | 0x6003D55E | 8 |  | - |
| `DS_C_MEM[1375]` | 0x6003D55F | 8 |  | - |
| `DS_C_MEM[1376]` | 0x6003D560 | 8 |  | - |
| `DS_C_MEM[1377]` | 0x6003D561 | 8 |  | - |
| `DS_C_MEM[1378]` | 0x6003D562 | 8 |  | - |
| `DS_C_MEM[1379]` | 0x6003D563 | 8 |  | - |
| `DS_C_MEM[137]` | 0x6003D089 | 8 |  | - |
| `DS_C_MEM[1380]` | 0x6003D564 | 8 |  | - |
| `DS_C_MEM[1381]` | 0x6003D565 | 8 |  | - |
| `DS_C_MEM[1382]` | 0x6003D566 | 8 |  | - |
| `DS_C_MEM[1383]` | 0x6003D567 | 8 |  | - |
| `DS_C_MEM[1384]` | 0x6003D568 | 8 |  | - |
| `DS_C_MEM[1385]` | 0x6003D569 | 8 |  | - |
| `DS_C_MEM[1386]` | 0x6003D56A | 8 |  | - |
| `DS_C_MEM[1387]` | 0x6003D56B | 8 |  | - |
| `DS_C_MEM[1388]` | 0x6003D56C | 8 |  | - |
| `DS_C_MEM[1389]` | 0x6003D56D | 8 |  | - |
| `DS_C_MEM[138]` | 0x6003D08A | 8 |  | - |
| `DS_C_MEM[1390]` | 0x6003D56E | 8 |  | - |
| `DS_C_MEM[1391]` | 0x6003D56F | 8 |  | - |
| `DS_C_MEM[1392]` | 0x6003D570 | 8 |  | - |
| `DS_C_MEM[1393]` | 0x6003D571 | 8 |  | - |
| `DS_C_MEM[1394]` | 0x6003D572 | 8 |  | - |
| `DS_C_MEM[1395]` | 0x6003D573 | 8 |  | - |
| `DS_C_MEM[1396]` | 0x6003D574 | 8 |  | - |
| `DS_C_MEM[1397]` | 0x6003D575 | 8 |  | - |
| `DS_C_MEM[1398]` | 0x6003D576 | 8 |  | - |
| `DS_C_MEM[1399]` | 0x6003D577 | 8 |  | - |
| `DS_C_MEM[139]` | 0x6003D08B | 8 |  | - |
| `DS_C_MEM[13]` | 0x6003D00D | 8 |  | - |
| `DS_C_MEM[1400]` | 0x6003D578 | 8 |  | - |
| `DS_C_MEM[1401]` | 0x6003D579 | 8 |  | - |
| `DS_C_MEM[1402]` | 0x6003D57A | 8 |  | - |
| `DS_C_MEM[1403]` | 0x6003D57B | 8 |  | - |
| `DS_C_MEM[1404]` | 0x6003D57C | 8 |  | - |
| `DS_C_MEM[1405]` | 0x6003D57D | 8 |  | - |
| `DS_C_MEM[1406]` | 0x6003D57E | 8 |  | - |
| `DS_C_MEM[1407]` | 0x6003D57F | 8 |  | - |
| `DS_C_MEM[1408]` | 0x6003D580 | 8 |  | - |
| `DS_C_MEM[1409]` | 0x6003D581 | 8 |  | - |
| `DS_C_MEM[140]` | 0x6003D08C | 8 |  | - |
| `DS_C_MEM[1410]` | 0x6003D582 | 8 |  | - |
| `DS_C_MEM[1411]` | 0x6003D583 | 8 |  | - |
| `DS_C_MEM[1412]` | 0x6003D584 | 8 |  | - |
| `DS_C_MEM[1413]` | 0x6003D585 | 8 |  | - |
| `DS_C_MEM[1414]` | 0x6003D586 | 8 |  | - |
| `DS_C_MEM[1415]` | 0x6003D587 | 8 |  | - |
| `DS_C_MEM[1416]` | 0x6003D588 | 8 |  | - |
| `DS_C_MEM[1417]` | 0x6003D589 | 8 |  | - |
| `DS_C_MEM[1418]` | 0x6003D58A | 8 |  | - |
| `DS_C_MEM[1419]` | 0x6003D58B | 8 |  | - |
| `DS_C_MEM[141]` | 0x6003D08D | 8 |  | - |
| `DS_C_MEM[1420]` | 0x6003D58C | 8 |  | - |
| `DS_C_MEM[1421]` | 0x6003D58D | 8 |  | - |
| `DS_C_MEM[1422]` | 0x6003D58E | 8 |  | - |
| `DS_C_MEM[1423]` | 0x6003D58F | 8 |  | - |
| `DS_C_MEM[1424]` | 0x6003D590 | 8 |  | - |
| `DS_C_MEM[1425]` | 0x6003D591 | 8 |  | - |
| `DS_C_MEM[1426]` | 0x6003D592 | 8 |  | - |
| `DS_C_MEM[1427]` | 0x6003D593 | 8 |  | - |
| `DS_C_MEM[1428]` | 0x6003D594 | 8 |  | - |
| `DS_C_MEM[1429]` | 0x6003D595 | 8 |  | - |
| `DS_C_MEM[142]` | 0x6003D08E | 8 |  | - |
| `DS_C_MEM[1430]` | 0x6003D596 | 8 |  | - |
| `DS_C_MEM[1431]` | 0x6003D597 | 8 |  | - |
| `DS_C_MEM[1432]` | 0x6003D598 | 8 |  | - |
| `DS_C_MEM[1433]` | 0x6003D599 | 8 |  | - |
| `DS_C_MEM[1434]` | 0x6003D59A | 8 |  | - |
| `DS_C_MEM[1435]` | 0x6003D59B | 8 |  | - |
| `DS_C_MEM[1436]` | 0x6003D59C | 8 |  | - |
| `DS_C_MEM[1437]` | 0x6003D59D | 8 |  | - |
| `DS_C_MEM[1438]` | 0x6003D59E | 8 |  | - |
| `DS_C_MEM[1439]` | 0x6003D59F | 8 |  | - |
| `DS_C_MEM[143]` | 0x6003D08F | 8 |  | - |
| `DS_C_MEM[1440]` | 0x6003D5A0 | 8 |  | - |
| `DS_C_MEM[1441]` | 0x6003D5A1 | 8 |  | - |
| `DS_C_MEM[1442]` | 0x6003D5A2 | 8 |  | - |
| `DS_C_MEM[1443]` | 0x6003D5A3 | 8 |  | - |
| `DS_C_MEM[1444]` | 0x6003D5A4 | 8 |  | - |
| `DS_C_MEM[1445]` | 0x6003D5A5 | 8 |  | - |
| `DS_C_MEM[1446]` | 0x6003D5A6 | 8 |  | - |
| `DS_C_MEM[1447]` | 0x6003D5A7 | 8 |  | - |
| `DS_C_MEM[1448]` | 0x6003D5A8 | 8 |  | - |
| `DS_C_MEM[1449]` | 0x6003D5A9 | 8 |  | - |
| `DS_C_MEM[144]` | 0x6003D090 | 8 |  | - |
| `DS_C_MEM[1450]` | 0x6003D5AA | 8 |  | - |
| `DS_C_MEM[1451]` | 0x6003D5AB | 8 |  | - |
| `DS_C_MEM[1452]` | 0x6003D5AC | 8 |  | - |
| `DS_C_MEM[1453]` | 0x6003D5AD | 8 |  | - |
| `DS_C_MEM[1454]` | 0x6003D5AE | 8 |  | - |
| `DS_C_MEM[1455]` | 0x6003D5AF | 8 |  | - |
| `DS_C_MEM[1456]` | 0x6003D5B0 | 8 |  | - |
| `DS_C_MEM[1457]` | 0x6003D5B1 | 8 |  | - |
| `DS_C_MEM[1458]` | 0x6003D5B2 | 8 |  | - |
| `DS_C_MEM[1459]` | 0x6003D5B3 | 8 |  | - |
| `DS_C_MEM[145]` | 0x6003D091 | 8 |  | - |
| `DS_C_MEM[1460]` | 0x6003D5B4 | 8 |  | - |
| `DS_C_MEM[1461]` | 0x6003D5B5 | 8 |  | - |
| `DS_C_MEM[1462]` | 0x6003D5B6 | 8 |  | - |
| `DS_C_MEM[1463]` | 0x6003D5B7 | 8 |  | - |
| `DS_C_MEM[1464]` | 0x6003D5B8 | 8 |  | - |
| `DS_C_MEM[1465]` | 0x6003D5B9 | 8 |  | - |
| `DS_C_MEM[1466]` | 0x6003D5BA | 8 |  | - |
| `DS_C_MEM[1467]` | 0x6003D5BB | 8 |  | - |
| `DS_C_MEM[1468]` | 0x6003D5BC | 8 |  | - |
| `DS_C_MEM[1469]` | 0x6003D5BD | 8 |  | - |
| `DS_C_MEM[146]` | 0x6003D092 | 8 |  | - |
| `DS_C_MEM[1470]` | 0x6003D5BE | 8 |  | - |
| `DS_C_MEM[1471]` | 0x6003D5BF | 8 |  | - |
| `DS_C_MEM[1472]` | 0x6003D5C0 | 8 |  | - |
| `DS_C_MEM[1473]` | 0x6003D5C1 | 8 |  | - |
| `DS_C_MEM[1474]` | 0x6003D5C2 | 8 |  | - |
| `DS_C_MEM[1475]` | 0x6003D5C3 | 8 |  | - |
| `DS_C_MEM[1476]` | 0x6003D5C4 | 8 |  | - |
| `DS_C_MEM[1477]` | 0x6003D5C5 | 8 |  | - |
| `DS_C_MEM[1478]` | 0x6003D5C6 | 8 |  | - |
| `DS_C_MEM[1479]` | 0x6003D5C7 | 8 |  | - |
| `DS_C_MEM[147]` | 0x6003D093 | 8 |  | - |
| `DS_C_MEM[1480]` | 0x6003D5C8 | 8 |  | - |
| `DS_C_MEM[1481]` | 0x6003D5C9 | 8 |  | - |
| `DS_C_MEM[1482]` | 0x6003D5CA | 8 |  | - |
| `DS_C_MEM[1483]` | 0x6003D5CB | 8 |  | - |
| `DS_C_MEM[1484]` | 0x6003D5CC | 8 |  | - |
| `DS_C_MEM[1485]` | 0x6003D5CD | 8 |  | - |
| `DS_C_MEM[1486]` | 0x6003D5CE | 8 |  | - |
| `DS_C_MEM[1487]` | 0x6003D5CF | 8 |  | - |
| `DS_C_MEM[1488]` | 0x6003D5D0 | 8 |  | - |
| `DS_C_MEM[1489]` | 0x6003D5D1 | 8 |  | - |
| `DS_C_MEM[148]` | 0x6003D094 | 8 |  | - |
| `DS_C_MEM[1490]` | 0x6003D5D2 | 8 |  | - |
| `DS_C_MEM[1491]` | 0x6003D5D3 | 8 |  | - |
| `DS_C_MEM[1492]` | 0x6003D5D4 | 8 |  | - |
| `DS_C_MEM[1493]` | 0x6003D5D5 | 8 |  | - |
| `DS_C_MEM[1494]` | 0x6003D5D6 | 8 |  | - |
| `DS_C_MEM[1495]` | 0x6003D5D7 | 8 |  | - |
| `DS_C_MEM[1496]` | 0x6003D5D8 | 8 |  | - |
| `DS_C_MEM[1497]` | 0x6003D5D9 | 8 |  | - |
| `DS_C_MEM[1498]` | 0x6003D5DA | 8 |  | - |
| `DS_C_MEM[1499]` | 0x6003D5DB | 8 |  | - |
| `DS_C_MEM[149]` | 0x6003D095 | 8 |  | - |
| `DS_C_MEM[14]` | 0x6003D00E | 8 |  | - |
| `DS_C_MEM[1500]` | 0x6003D5DC | 8 |  | - |
| `DS_C_MEM[1501]` | 0x6003D5DD | 8 |  | - |
| `DS_C_MEM[1502]` | 0x6003D5DE | 8 |  | - |
| `DS_C_MEM[1503]` | 0x6003D5DF | 8 |  | - |
| `DS_C_MEM[1504]` | 0x6003D5E0 | 8 |  | - |
| `DS_C_MEM[1505]` | 0x6003D5E1 | 8 |  | - |
| `DS_C_MEM[1506]` | 0x6003D5E2 | 8 |  | - |
| `DS_C_MEM[1507]` | 0x6003D5E3 | 8 |  | - |
| `DS_C_MEM[1508]` | 0x6003D5E4 | 8 |  | - |
| `DS_C_MEM[1509]` | 0x6003D5E5 | 8 |  | - |
| `DS_C_MEM[150]` | 0x6003D096 | 8 |  | - |
| `DS_C_MEM[1510]` | 0x6003D5E6 | 8 |  | - |
| `DS_C_MEM[1511]` | 0x6003D5E7 | 8 |  | - |
| `DS_C_MEM[1512]` | 0x6003D5E8 | 8 |  | - |
| `DS_C_MEM[1513]` | 0x6003D5E9 | 8 |  | - |
| `DS_C_MEM[1514]` | 0x6003D5EA | 8 |  | - |
| `DS_C_MEM[1515]` | 0x6003D5EB | 8 |  | - |
| `DS_C_MEM[1516]` | 0x6003D5EC | 8 |  | - |
| `DS_C_MEM[1517]` | 0x6003D5ED | 8 |  | - |
| `DS_C_MEM[1518]` | 0x6003D5EE | 8 |  | - |
| `DS_C_MEM[1519]` | 0x6003D5EF | 8 |  | - |
| `DS_C_MEM[151]` | 0x6003D097 | 8 |  | - |
| `DS_C_MEM[1520]` | 0x6003D5F0 | 8 |  | - |
| `DS_C_MEM[1521]` | 0x6003D5F1 | 8 |  | - |
| `DS_C_MEM[1522]` | 0x6003D5F2 | 8 |  | - |
| `DS_C_MEM[1523]` | 0x6003D5F3 | 8 |  | - |
| `DS_C_MEM[1524]` | 0x6003D5F4 | 8 |  | - |
| `DS_C_MEM[1525]` | 0x6003D5F5 | 8 |  | - |
| `DS_C_MEM[1526]` | 0x6003D5F6 | 8 |  | - |
| `DS_C_MEM[1527]` | 0x6003D5F7 | 8 |  | - |
| `DS_C_MEM[1528]` | 0x6003D5F8 | 8 |  | - |
| `DS_C_MEM[1529]` | 0x6003D5F9 | 8 |  | - |
| `DS_C_MEM[152]` | 0x6003D098 | 8 |  | - |
| `DS_C_MEM[1530]` | 0x6003D5FA | 8 |  | - |
| `DS_C_MEM[1531]` | 0x6003D5FB | 8 |  | - |
| `DS_C_MEM[1532]` | 0x6003D5FC | 8 |  | - |
| `DS_C_MEM[1533]` | 0x6003D5FD | 8 |  | - |
| `DS_C_MEM[1534]` | 0x6003D5FE | 8 |  | - |
| `DS_C_MEM[1535]` | 0x6003D5FF | 8 |  | - |
| `DS_C_MEM[1536]` | 0x6003D600 | 8 |  | - |
| `DS_C_MEM[1537]` | 0x6003D601 | 8 |  | - |
| `DS_C_MEM[1538]` | 0x6003D602 | 8 |  | - |
| `DS_C_MEM[1539]` | 0x6003D603 | 8 |  | - |
| `DS_C_MEM[153]` | 0x6003D099 | 8 |  | - |
| `DS_C_MEM[1540]` | 0x6003D604 | 8 |  | - |
| `DS_C_MEM[1541]` | 0x6003D605 | 8 |  | - |
| `DS_C_MEM[1542]` | 0x6003D606 | 8 |  | - |
| `DS_C_MEM[1543]` | 0x6003D607 | 8 |  | - |
| `DS_C_MEM[1544]` | 0x6003D608 | 8 |  | - |
| `DS_C_MEM[1545]` | 0x6003D609 | 8 |  | - |
| `DS_C_MEM[1546]` | 0x6003D60A | 8 |  | - |
| `DS_C_MEM[1547]` | 0x6003D60B | 8 |  | - |
| `DS_C_MEM[1548]` | 0x6003D60C | 8 |  | - |
| `DS_C_MEM[1549]` | 0x6003D60D | 8 |  | - |
| `DS_C_MEM[154]` | 0x6003D09A | 8 |  | - |
| `DS_C_MEM[1550]` | 0x6003D60E | 8 |  | - |
| `DS_C_MEM[1551]` | 0x6003D60F | 8 |  | - |
| `DS_C_MEM[1552]` | 0x6003D610 | 8 |  | - |
| `DS_C_MEM[1553]` | 0x6003D611 | 8 |  | - |
| `DS_C_MEM[1554]` | 0x6003D612 | 8 |  | - |
| `DS_C_MEM[1555]` | 0x6003D613 | 8 |  | - |
| `DS_C_MEM[1556]` | 0x6003D614 | 8 |  | - |
| `DS_C_MEM[1557]` | 0x6003D615 | 8 |  | - |
| `DS_C_MEM[1558]` | 0x6003D616 | 8 |  | - |
| `DS_C_MEM[1559]` | 0x6003D617 | 8 |  | - |
| `DS_C_MEM[155]` | 0x6003D09B | 8 |  | - |
| `DS_C_MEM[1560]` | 0x6003D618 | 8 |  | - |
| `DS_C_MEM[1561]` | 0x6003D619 | 8 |  | - |
| `DS_C_MEM[1562]` | 0x6003D61A | 8 |  | - |
| `DS_C_MEM[1563]` | 0x6003D61B | 8 |  | - |
| `DS_C_MEM[1564]` | 0x6003D61C | 8 |  | - |
| `DS_C_MEM[1565]` | 0x6003D61D | 8 |  | - |
| `DS_C_MEM[1566]` | 0x6003D61E | 8 |  | - |
| `DS_C_MEM[1567]` | 0x6003D61F | 8 |  | - |
| `DS_C_MEM[1568]` | 0x6003D620 | 8 |  | - |
| `DS_C_MEM[1569]` | 0x6003D621 | 8 |  | - |
| `DS_C_MEM[156]` | 0x6003D09C | 8 |  | - |
| `DS_C_MEM[1570]` | 0x6003D622 | 8 |  | - |
| `DS_C_MEM[1571]` | 0x6003D623 | 8 |  | - |
| `DS_C_MEM[1572]` | 0x6003D624 | 8 |  | - |
| `DS_C_MEM[1573]` | 0x6003D625 | 8 |  | - |
| `DS_C_MEM[1574]` | 0x6003D626 | 8 |  | - |
| `DS_C_MEM[1575]` | 0x6003D627 | 8 |  | - |
| `DS_C_MEM[1576]` | 0x6003D628 | 8 |  | - |
| `DS_C_MEM[1577]` | 0x6003D629 | 8 |  | - |
| `DS_C_MEM[1578]` | 0x6003D62A | 8 |  | - |
| `DS_C_MEM[1579]` | 0x6003D62B | 8 |  | - |
| `DS_C_MEM[157]` | 0x6003D09D | 8 |  | - |
| `DS_C_MEM[1580]` | 0x6003D62C | 8 |  | - |
| `DS_C_MEM[1581]` | 0x6003D62D | 8 |  | - |
| `DS_C_MEM[1582]` | 0x6003D62E | 8 |  | - |
| `DS_C_MEM[1583]` | 0x6003D62F | 8 |  | - |
| `DS_C_MEM[158]` | 0x6003D09E | 8 |  | - |
| `DS_C_MEM[159]` | 0x6003D09F | 8 |  | - |
| `DS_C_MEM[15]` | 0x6003D00F | 8 |  | - |
| `DS_C_MEM[160]` | 0x6003D0A0 | 8 |  | - |
| `DS_C_MEM[161]` | 0x6003D0A1 | 8 |  | - |
| `DS_C_MEM[162]` | 0x6003D0A2 | 8 |  | - |
| `DS_C_MEM[163]` | 0x6003D0A3 | 8 |  | - |
| `DS_C_MEM[164]` | 0x6003D0A4 | 8 |  | - |
| `DS_C_MEM[165]` | 0x6003D0A5 | 8 |  | - |
| `DS_C_MEM[166]` | 0x6003D0A6 | 8 |  | - |
| `DS_C_MEM[167]` | 0x6003D0A7 | 8 |  | - |
| `DS_C_MEM[168]` | 0x6003D0A8 | 8 |  | - |
| `DS_C_MEM[169]` | 0x6003D0A9 | 8 |  | - |
| `DS_C_MEM[16]` | 0x6003D010 | 8 |  | - |
| `DS_C_MEM[170]` | 0x6003D0AA | 8 |  | - |
| `DS_C_MEM[171]` | 0x6003D0AB | 8 |  | - |
| `DS_C_MEM[172]` | 0x6003D0AC | 8 |  | - |
| `DS_C_MEM[173]` | 0x6003D0AD | 8 |  | - |
| `DS_C_MEM[174]` | 0x6003D0AE | 8 |  | - |
| `DS_C_MEM[175]` | 0x6003D0AF | 8 |  | - |
| `DS_C_MEM[176]` | 0x6003D0B0 | 8 |  | - |
| `DS_C_MEM[177]` | 0x6003D0B1 | 8 |  | - |
| `DS_C_MEM[178]` | 0x6003D0B2 | 8 |  | - |
| `DS_C_MEM[179]` | 0x6003D0B3 | 8 |  | - |
| `DS_C_MEM[17]` | 0x6003D011 | 8 |  | - |
| `DS_C_MEM[180]` | 0x6003D0B4 | 8 |  | - |
| `DS_C_MEM[181]` | 0x6003D0B5 | 8 |  | - |
| `DS_C_MEM[182]` | 0x6003D0B6 | 8 |  | - |
| `DS_C_MEM[183]` | 0x6003D0B7 | 8 |  | - |
| `DS_C_MEM[184]` | 0x6003D0B8 | 8 |  | - |
| `DS_C_MEM[185]` | 0x6003D0B9 | 8 |  | - |
| `DS_C_MEM[186]` | 0x6003D0BA | 8 |  | - |
| `DS_C_MEM[187]` | 0x6003D0BB | 8 |  | - |
| `DS_C_MEM[188]` | 0x6003D0BC | 8 |  | - |
| `DS_C_MEM[189]` | 0x6003D0BD | 8 |  | - |
| `DS_C_MEM[18]` | 0x6003D012 | 8 |  | - |
| `DS_C_MEM[190]` | 0x6003D0BE | 8 |  | - |
| `DS_C_MEM[191]` | 0x6003D0BF | 8 |  | - |
| `DS_C_MEM[192]` | 0x6003D0C0 | 8 |  | - |
| `DS_C_MEM[193]` | 0x6003D0C1 | 8 |  | - |
| `DS_C_MEM[194]` | 0x6003D0C2 | 8 |  | - |
| `DS_C_MEM[195]` | 0x6003D0C3 | 8 |  | - |
| `DS_C_MEM[196]` | 0x6003D0C4 | 8 |  | - |
| `DS_C_MEM[197]` | 0x6003D0C5 | 8 |  | - |
| `DS_C_MEM[198]` | 0x6003D0C6 | 8 |  | - |
| `DS_C_MEM[199]` | 0x6003D0C7 | 8 |  | - |
| `DS_C_MEM[19]` | 0x6003D013 | 8 |  | - |
| `DS_C_MEM[1]` | 0x6003D001 | 8 |  | - |
| `DS_C_MEM[200]` | 0x6003D0C8 | 8 |  | - |
| `DS_C_MEM[201]` | 0x6003D0C9 | 8 |  | - |
| `DS_C_MEM[202]` | 0x6003D0CA | 8 |  | - |
| `DS_C_MEM[203]` | 0x6003D0CB | 8 |  | - |
| `DS_C_MEM[204]` | 0x6003D0CC | 8 |  | - |
| `DS_C_MEM[205]` | 0x6003D0CD | 8 |  | - |
| `DS_C_MEM[206]` | 0x6003D0CE | 8 |  | - |
| `DS_C_MEM[207]` | 0x6003D0CF | 8 |  | - |
| `DS_C_MEM[208]` | 0x6003D0D0 | 8 |  | - |
| `DS_C_MEM[209]` | 0x6003D0D1 | 8 |  | - |
| `DS_C_MEM[20]` | 0x6003D014 | 8 |  | - |
| `DS_C_MEM[210]` | 0x6003D0D2 | 8 |  | - |
| `DS_C_MEM[211]` | 0x6003D0D3 | 8 |  | - |
| `DS_C_MEM[212]` | 0x6003D0D4 | 8 |  | - |
| `DS_C_MEM[213]` | 0x6003D0D5 | 8 |  | - |
| `DS_C_MEM[214]` | 0x6003D0D6 | 8 |  | - |
| `DS_C_MEM[215]` | 0x6003D0D7 | 8 |  | - |
| `DS_C_MEM[216]` | 0x6003D0D8 | 8 |  | - |
| `DS_C_MEM[217]` | 0x6003D0D9 | 8 |  | - |
| `DS_C_MEM[218]` | 0x6003D0DA | 8 |  | - |
| `DS_C_MEM[219]` | 0x6003D0DB | 8 |  | - |
| `DS_C_MEM[21]` | 0x6003D015 | 8 |  | - |
| `DS_C_MEM[220]` | 0x6003D0DC | 8 |  | - |
| `DS_C_MEM[221]` | 0x6003D0DD | 8 |  | - |
| `DS_C_MEM[222]` | 0x6003D0DE | 8 |  | - |
| `DS_C_MEM[223]` | 0x6003D0DF | 8 |  | - |
| `DS_C_MEM[224]` | 0x6003D0E0 | 8 |  | - |
| `DS_C_MEM[225]` | 0x6003D0E1 | 8 |  | - |
| `DS_C_MEM[226]` | 0x6003D0E2 | 8 |  | - |
| `DS_C_MEM[227]` | 0x6003D0E3 | 8 |  | - |
| `DS_C_MEM[228]` | 0x6003D0E4 | 8 |  | - |
| `DS_C_MEM[229]` | 0x6003D0E5 | 8 |  | - |
| `DS_C_MEM[22]` | 0x6003D016 | 8 |  | - |
| `DS_C_MEM[230]` | 0x6003D0E6 | 8 |  | - |
| `DS_C_MEM[231]` | 0x6003D0E7 | 8 |  | - |
| `DS_C_MEM[232]` | 0x6003D0E8 | 8 |  | - |
| `DS_C_MEM[233]` | 0x6003D0E9 | 8 |  | - |
| `DS_C_MEM[234]` | 0x6003D0EA | 8 |  | - |
| `DS_C_MEM[235]` | 0x6003D0EB | 8 |  | - |
| `DS_C_MEM[236]` | 0x6003D0EC | 8 |  | - |
| `DS_C_MEM[237]` | 0x6003D0ED | 8 |  | - |
| `DS_C_MEM[238]` | 0x6003D0EE | 8 |  | - |
| `DS_C_MEM[239]` | 0x6003D0EF | 8 |  | - |
| `DS_C_MEM[23]` | 0x6003D017 | 8 |  | - |
| `DS_C_MEM[240]` | 0x6003D0F0 | 8 |  | - |
| `DS_C_MEM[241]` | 0x6003D0F1 | 8 |  | - |
| `DS_C_MEM[242]` | 0x6003D0F2 | 8 |  | - |
| `DS_C_MEM[243]` | 0x6003D0F3 | 8 |  | - |
| `DS_C_MEM[244]` | 0x6003D0F4 | 8 |  | - |
| `DS_C_MEM[245]` | 0x6003D0F5 | 8 |  | - |
| `DS_C_MEM[246]` | 0x6003D0F6 | 8 |  | - |
| `DS_C_MEM[247]` | 0x6003D0F7 | 8 |  | - |
| `DS_C_MEM[248]` | 0x6003D0F8 | 8 |  | - |
| `DS_C_MEM[249]` | 0x6003D0F9 | 8 |  | - |
| `DS_C_MEM[24]` | 0x6003D018 | 8 |  | - |
| `DS_C_MEM[250]` | 0x6003D0FA | 8 |  | - |
| `DS_C_MEM[251]` | 0x6003D0FB | 8 |  | - |
| `DS_C_MEM[252]` | 0x6003D0FC | 8 |  | - |
| `DS_C_MEM[253]` | 0x6003D0FD | 8 |  | - |
| `DS_C_MEM[254]` | 0x6003D0FE | 8 |  | - |
| `DS_C_MEM[255]` | 0x6003D0FF | 8 |  | - |
| `DS_C_MEM[256]` | 0x6003D100 | 8 |  | - |
| `DS_C_MEM[257]` | 0x6003D101 | 8 |  | - |
| `DS_C_MEM[258]` | 0x6003D102 | 8 |  | - |
| `DS_C_MEM[259]` | 0x6003D103 | 8 |  | - |
| `DS_C_MEM[25]` | 0x6003D019 | 8 |  | - |
| `DS_C_MEM[260]` | 0x6003D104 | 8 |  | - |
| `DS_C_MEM[261]` | 0x6003D105 | 8 |  | - |
| `DS_C_MEM[262]` | 0x6003D106 | 8 |  | - |
| `DS_C_MEM[263]` | 0x6003D107 | 8 |  | - |
| `DS_C_MEM[264]` | 0x6003D108 | 8 |  | - |
| `DS_C_MEM[265]` | 0x6003D109 | 8 |  | - |
| `DS_C_MEM[266]` | 0x6003D10A | 8 |  | - |
| `DS_C_MEM[267]` | 0x6003D10B | 8 |  | - |
| `DS_C_MEM[268]` | 0x6003D10C | 8 |  | - |
| `DS_C_MEM[269]` | 0x6003D10D | 8 |  | - |
| `DS_C_MEM[26]` | 0x6003D01A | 8 |  | - |
| `DS_C_MEM[270]` | 0x6003D10E | 8 |  | - |
| `DS_C_MEM[271]` | 0x6003D10F | 8 |  | - |
| `DS_C_MEM[272]` | 0x6003D110 | 8 |  | - |
| `DS_C_MEM[273]` | 0x6003D111 | 8 |  | - |
| `DS_C_MEM[274]` | 0x6003D112 | 8 |  | - |
| `DS_C_MEM[275]` | 0x6003D113 | 8 |  | - |
| `DS_C_MEM[276]` | 0x6003D114 | 8 |  | - |
| `DS_C_MEM[277]` | 0x6003D115 | 8 |  | - |
| `DS_C_MEM[278]` | 0x6003D116 | 8 |  | - |
| `DS_C_MEM[279]` | 0x6003D117 | 8 |  | - |
| `DS_C_MEM[27]` | 0x6003D01B | 8 |  | - |
| `DS_C_MEM[280]` | 0x6003D118 | 8 |  | - |
| `DS_C_MEM[281]` | 0x6003D119 | 8 |  | - |
| `DS_C_MEM[282]` | 0x6003D11A | 8 |  | - |
| `DS_C_MEM[283]` | 0x6003D11B | 8 |  | - |
| `DS_C_MEM[284]` | 0x6003D11C | 8 |  | - |
| `DS_C_MEM[285]` | 0x6003D11D | 8 |  | - |
| `DS_C_MEM[286]` | 0x6003D11E | 8 |  | - |
| `DS_C_MEM[287]` | 0x6003D11F | 8 |  | - |
| `DS_C_MEM[288]` | 0x6003D120 | 8 |  | - |
| `DS_C_MEM[289]` | 0x6003D121 | 8 |  | - |
| `DS_C_MEM[28]` | 0x6003D01C | 8 |  | - |
| `DS_C_MEM[290]` | 0x6003D122 | 8 |  | - |
| `DS_C_MEM[291]` | 0x6003D123 | 8 |  | - |
| `DS_C_MEM[292]` | 0x6003D124 | 8 |  | - |
| `DS_C_MEM[293]` | 0x6003D125 | 8 |  | - |
| `DS_C_MEM[294]` | 0x6003D126 | 8 |  | - |
| `DS_C_MEM[295]` | 0x6003D127 | 8 |  | - |
| `DS_C_MEM[296]` | 0x6003D128 | 8 |  | - |
| `DS_C_MEM[297]` | 0x6003D129 | 8 |  | - |
| `DS_C_MEM[298]` | 0x6003D12A | 8 |  | - |
| `DS_C_MEM[299]` | 0x6003D12B | 8 |  | - |
| `DS_C_MEM[29]` | 0x6003D01D | 8 |  | - |
| `DS_C_MEM[2]` | 0x6003D002 | 8 |  | - |
| `DS_C_MEM[300]` | 0x6003D12C | 8 |  | - |
| `DS_C_MEM[301]` | 0x6003D12D | 8 |  | - |
| `DS_C_MEM[302]` | 0x6003D12E | 8 |  | - |
| `DS_C_MEM[303]` | 0x6003D12F | 8 |  | - |
| `DS_C_MEM[304]` | 0x6003D130 | 8 |  | - |
| `DS_C_MEM[305]` | 0x6003D131 | 8 |  | - |
| `DS_C_MEM[306]` | 0x6003D132 | 8 |  | - |
| `DS_C_MEM[307]` | 0x6003D133 | 8 |  | - |
| `DS_C_MEM[308]` | 0x6003D134 | 8 |  | - |
| `DS_C_MEM[309]` | 0x6003D135 | 8 |  | - |
| `DS_C_MEM[30]` | 0x6003D01E | 8 |  | - |
| `DS_C_MEM[310]` | 0x6003D136 | 8 |  | - |
| `DS_C_MEM[311]` | 0x6003D137 | 8 |  | - |
| `DS_C_MEM[312]` | 0x6003D138 | 8 |  | - |
| `DS_C_MEM[313]` | 0x6003D139 | 8 |  | - |
| `DS_C_MEM[314]` | 0x6003D13A | 8 |  | - |
| `DS_C_MEM[315]` | 0x6003D13B | 8 |  | - |
| `DS_C_MEM[316]` | 0x6003D13C | 8 |  | - |
| `DS_C_MEM[317]` | 0x6003D13D | 8 |  | - |
| `DS_C_MEM[318]` | 0x6003D13E | 8 |  | - |
| `DS_C_MEM[319]` | 0x6003D13F | 8 |  | - |
| `DS_C_MEM[31]` | 0x6003D01F | 8 |  | - |
| `DS_C_MEM[320]` | 0x6003D140 | 8 |  | - |
| `DS_C_MEM[321]` | 0x6003D141 | 8 |  | - |
| `DS_C_MEM[322]` | 0x6003D142 | 8 |  | - |
| `DS_C_MEM[323]` | 0x6003D143 | 8 |  | - |
| `DS_C_MEM[324]` | 0x6003D144 | 8 |  | - |
| `DS_C_MEM[325]` | 0x6003D145 | 8 |  | - |
| `DS_C_MEM[326]` | 0x6003D146 | 8 |  | - |
| `DS_C_MEM[327]` | 0x6003D147 | 8 |  | - |
| `DS_C_MEM[328]` | 0x6003D148 | 8 |  | - |
| `DS_C_MEM[329]` | 0x6003D149 | 8 |  | - |
| `DS_C_MEM[32]` | 0x6003D020 | 8 |  | - |
| `DS_C_MEM[330]` | 0x6003D14A | 8 |  | - |
| `DS_C_MEM[331]` | 0x6003D14B | 8 |  | - |
| `DS_C_MEM[332]` | 0x6003D14C | 8 |  | - |
| `DS_C_MEM[333]` | 0x6003D14D | 8 |  | - |
| `DS_C_MEM[334]` | 0x6003D14E | 8 |  | - |
| `DS_C_MEM[335]` | 0x6003D14F | 8 |  | - |
| `DS_C_MEM[336]` | 0x6003D150 | 8 |  | - |
| `DS_C_MEM[337]` | 0x6003D151 | 8 |  | - |
| `DS_C_MEM[338]` | 0x6003D152 | 8 |  | - |
| `DS_C_MEM[339]` | 0x6003D153 | 8 |  | - |
| `DS_C_MEM[33]` | 0x6003D021 | 8 |  | - |
| `DS_C_MEM[340]` | 0x6003D154 | 8 |  | - |
| `DS_C_MEM[341]` | 0x6003D155 | 8 |  | - |
| `DS_C_MEM[342]` | 0x6003D156 | 8 |  | - |
| `DS_C_MEM[343]` | 0x6003D157 | 8 |  | - |
| `DS_C_MEM[344]` | 0x6003D158 | 8 |  | - |
| `DS_C_MEM[345]` | 0x6003D159 | 8 |  | - |
| `DS_C_MEM[346]` | 0x6003D15A | 8 |  | - |
| `DS_C_MEM[347]` | 0x6003D15B | 8 |  | - |
| `DS_C_MEM[348]` | 0x6003D15C | 8 |  | - |
| `DS_C_MEM[349]` | 0x6003D15D | 8 |  | - |
| `DS_C_MEM[34]` | 0x6003D022 | 8 |  | - |
| `DS_C_MEM[350]` | 0x6003D15E | 8 |  | - |
| `DS_C_MEM[351]` | 0x6003D15F | 8 |  | - |
| `DS_C_MEM[352]` | 0x6003D160 | 8 |  | - |
| `DS_C_MEM[353]` | 0x6003D161 | 8 |  | - |
| `DS_C_MEM[354]` | 0x6003D162 | 8 |  | - |
| `DS_C_MEM[355]` | 0x6003D163 | 8 |  | - |
| `DS_C_MEM[356]` | 0x6003D164 | 8 |  | - |
| `DS_C_MEM[357]` | 0x6003D165 | 8 |  | - |
| `DS_C_MEM[358]` | 0x6003D166 | 8 |  | - |
| `DS_C_MEM[359]` | 0x6003D167 | 8 |  | - |
| `DS_C_MEM[35]` | 0x6003D023 | 8 |  | - |
| `DS_C_MEM[360]` | 0x6003D168 | 8 |  | - |
| `DS_C_MEM[361]` | 0x6003D169 | 8 |  | - |
| `DS_C_MEM[362]` | 0x6003D16A | 8 |  | - |
| `DS_C_MEM[363]` | 0x6003D16B | 8 |  | - |
| `DS_C_MEM[364]` | 0x6003D16C | 8 |  | - |
| `DS_C_MEM[365]` | 0x6003D16D | 8 |  | - |
| `DS_C_MEM[366]` | 0x6003D16E | 8 |  | - |
| `DS_C_MEM[367]` | 0x6003D16F | 8 |  | - |
| `DS_C_MEM[368]` | 0x6003D170 | 8 |  | - |
| `DS_C_MEM[369]` | 0x6003D171 | 8 |  | - |
| `DS_C_MEM[36]` | 0x6003D024 | 8 |  | - |
| `DS_C_MEM[370]` | 0x6003D172 | 8 |  | - |
| `DS_C_MEM[371]` | 0x6003D173 | 8 |  | - |
| `DS_C_MEM[372]` | 0x6003D174 | 8 |  | - |
| `DS_C_MEM[373]` | 0x6003D175 | 8 |  | - |
| `DS_C_MEM[374]` | 0x6003D176 | 8 |  | - |
| `DS_C_MEM[375]` | 0x6003D177 | 8 |  | - |
| `DS_C_MEM[376]` | 0x6003D178 | 8 |  | - |
| `DS_C_MEM[377]` | 0x6003D179 | 8 |  | - |
| `DS_C_MEM[378]` | 0x6003D17A | 8 |  | - |
| `DS_C_MEM[379]` | 0x6003D17B | 8 |  | - |
| `DS_C_MEM[37]` | 0x6003D025 | 8 |  | - |
| `DS_C_MEM[380]` | 0x6003D17C | 8 |  | - |
| `DS_C_MEM[381]` | 0x6003D17D | 8 |  | - |
| `DS_C_MEM[382]` | 0x6003D17E | 8 |  | - |
| `DS_C_MEM[383]` | 0x6003D17F | 8 |  | - |
| `DS_C_MEM[384]` | 0x6003D180 | 8 |  | - |
| `DS_C_MEM[385]` | 0x6003D181 | 8 |  | - |
| `DS_C_MEM[386]` | 0x6003D182 | 8 |  | - |
| `DS_C_MEM[387]` | 0x6003D183 | 8 |  | - |
| `DS_C_MEM[388]` | 0x6003D184 | 8 |  | - |
| `DS_C_MEM[389]` | 0x6003D185 | 8 |  | - |
| `DS_C_MEM[38]` | 0x6003D026 | 8 |  | - |
| `DS_C_MEM[390]` | 0x6003D186 | 8 |  | - |
| `DS_C_MEM[391]` | 0x6003D187 | 8 |  | - |
| `DS_C_MEM[392]` | 0x6003D188 | 8 |  | - |
| `DS_C_MEM[393]` | 0x6003D189 | 8 |  | - |
| `DS_C_MEM[394]` | 0x6003D18A | 8 |  | - |
| `DS_C_MEM[395]` | 0x6003D18B | 8 |  | - |
| `DS_C_MEM[396]` | 0x6003D18C | 8 |  | - |
| `DS_C_MEM[397]` | 0x6003D18D | 8 |  | - |
| `DS_C_MEM[398]` | 0x6003D18E | 8 |  | - |
| `DS_C_MEM[399]` | 0x6003D18F | 8 |  | - |
| `DS_C_MEM[39]` | 0x6003D027 | 8 |  | - |
| `DS_C_MEM[3]` | 0x6003D003 | 8 |  | - |
| `DS_C_MEM[400]` | 0x6003D190 | 8 |  | - |
| `DS_C_MEM[401]` | 0x6003D191 | 8 |  | - |
| `DS_C_MEM[402]` | 0x6003D192 | 8 |  | - |
| `DS_C_MEM[403]` | 0x6003D193 | 8 |  | - |
| `DS_C_MEM[404]` | 0x6003D194 | 8 |  | - |
| `DS_C_MEM[405]` | 0x6003D195 | 8 |  | - |
| `DS_C_MEM[406]` | 0x6003D196 | 8 |  | - |
| `DS_C_MEM[407]` | 0x6003D197 | 8 |  | - |
| `DS_C_MEM[408]` | 0x6003D198 | 8 |  | - |
| `DS_C_MEM[409]` | 0x6003D199 | 8 |  | - |
| `DS_C_MEM[40]` | 0x6003D028 | 8 |  | - |
| `DS_C_MEM[410]` | 0x6003D19A | 8 |  | - |
| `DS_C_MEM[411]` | 0x6003D19B | 8 |  | - |
| `DS_C_MEM[412]` | 0x6003D19C | 8 |  | - |
| `DS_C_MEM[413]` | 0x6003D19D | 8 |  | - |
| `DS_C_MEM[414]` | 0x6003D19E | 8 |  | - |
| `DS_C_MEM[415]` | 0x6003D19F | 8 |  | - |
| `DS_C_MEM[416]` | 0x6003D1A0 | 8 |  | - |
| `DS_C_MEM[417]` | 0x6003D1A1 | 8 |  | - |
| `DS_C_MEM[418]` | 0x6003D1A2 | 8 |  | - |
| `DS_C_MEM[419]` | 0x6003D1A3 | 8 |  | - |
| `DS_C_MEM[41]` | 0x6003D029 | 8 |  | - |
| `DS_C_MEM[420]` | 0x6003D1A4 | 8 |  | - |
| `DS_C_MEM[421]` | 0x6003D1A5 | 8 |  | - |
| `DS_C_MEM[422]` | 0x6003D1A6 | 8 |  | - |
| `DS_C_MEM[423]` | 0x6003D1A7 | 8 |  | - |
| `DS_C_MEM[424]` | 0x6003D1A8 | 8 |  | - |
| `DS_C_MEM[425]` | 0x6003D1A9 | 8 |  | - |
| `DS_C_MEM[426]` | 0x6003D1AA | 8 |  | - |
| `DS_C_MEM[427]` | 0x6003D1AB | 8 |  | - |
| `DS_C_MEM[428]` | 0x6003D1AC | 8 |  | - |
| `DS_C_MEM[429]` | 0x6003D1AD | 8 |  | - |
| `DS_C_MEM[42]` | 0x6003D02A | 8 |  | - |
| `DS_C_MEM[430]` | 0x6003D1AE | 8 |  | - |
| `DS_C_MEM[431]` | 0x6003D1AF | 8 |  | - |
| `DS_C_MEM[432]` | 0x6003D1B0 | 8 |  | - |
| `DS_C_MEM[433]` | 0x6003D1B1 | 8 |  | - |
| `DS_C_MEM[434]` | 0x6003D1B2 | 8 |  | - |
| `DS_C_MEM[435]` | 0x6003D1B3 | 8 |  | - |
| `DS_C_MEM[436]` | 0x6003D1B4 | 8 |  | - |
| `DS_C_MEM[437]` | 0x6003D1B5 | 8 |  | - |
| `DS_C_MEM[438]` | 0x6003D1B6 | 8 |  | - |
| `DS_C_MEM[439]` | 0x6003D1B7 | 8 |  | - |
| `DS_C_MEM[43]` | 0x6003D02B | 8 |  | - |
| `DS_C_MEM[440]` | 0x6003D1B8 | 8 |  | - |
| `DS_C_MEM[441]` | 0x6003D1B9 | 8 |  | - |
| `DS_C_MEM[442]` | 0x6003D1BA | 8 |  | - |
| `DS_C_MEM[443]` | 0x6003D1BB | 8 |  | - |
| `DS_C_MEM[444]` | 0x6003D1BC | 8 |  | - |
| `DS_C_MEM[445]` | 0x6003D1BD | 8 |  | - |
| `DS_C_MEM[446]` | 0x6003D1BE | 8 |  | - |
| `DS_C_MEM[447]` | 0x6003D1BF | 8 |  | - |
| `DS_C_MEM[448]` | 0x6003D1C0 | 8 |  | - |
| `DS_C_MEM[449]` | 0x6003D1C1 | 8 |  | - |
| `DS_C_MEM[44]` | 0x6003D02C | 8 |  | - |
| `DS_C_MEM[450]` | 0x6003D1C2 | 8 |  | - |
| `DS_C_MEM[451]` | 0x6003D1C3 | 8 |  | - |
| `DS_C_MEM[452]` | 0x6003D1C4 | 8 |  | - |
| `DS_C_MEM[453]` | 0x6003D1C5 | 8 |  | - |
| `DS_C_MEM[454]` | 0x6003D1C6 | 8 |  | - |
| `DS_C_MEM[455]` | 0x6003D1C7 | 8 |  | - |
| `DS_C_MEM[456]` | 0x6003D1C8 | 8 |  | - |
| `DS_C_MEM[457]` | 0x6003D1C9 | 8 |  | - |
| `DS_C_MEM[458]` | 0x6003D1CA | 8 |  | - |
| `DS_C_MEM[459]` | 0x6003D1CB | 8 |  | - |
| `DS_C_MEM[45]` | 0x6003D02D | 8 |  | - |
| `DS_C_MEM[460]` | 0x6003D1CC | 8 |  | - |
| `DS_C_MEM[461]` | 0x6003D1CD | 8 |  | - |
| `DS_C_MEM[462]` | 0x6003D1CE | 8 |  | - |
| `DS_C_MEM[463]` | 0x6003D1CF | 8 |  | - |
| `DS_C_MEM[464]` | 0x6003D1D0 | 8 |  | - |
| `DS_C_MEM[465]` | 0x6003D1D1 | 8 |  | - |
| `DS_C_MEM[466]` | 0x6003D1D2 | 8 |  | - |
| `DS_C_MEM[467]` | 0x6003D1D3 | 8 |  | - |
| `DS_C_MEM[468]` | 0x6003D1D4 | 8 |  | - |
| `DS_C_MEM[469]` | 0x6003D1D5 | 8 |  | - |
| `DS_C_MEM[46]` | 0x6003D02E | 8 |  | - |
| `DS_C_MEM[470]` | 0x6003D1D6 | 8 |  | - |
| `DS_C_MEM[471]` | 0x6003D1D7 | 8 |  | - |
| `DS_C_MEM[472]` | 0x6003D1D8 | 8 |  | - |
| `DS_C_MEM[473]` | 0x6003D1D9 | 8 |  | - |
| `DS_C_MEM[474]` | 0x6003D1DA | 8 |  | - |
| `DS_C_MEM[475]` | 0x6003D1DB | 8 |  | - |
| `DS_C_MEM[476]` | 0x6003D1DC | 8 |  | - |
| `DS_C_MEM[477]` | 0x6003D1DD | 8 |  | - |
| `DS_C_MEM[478]` | 0x6003D1DE | 8 |  | - |
| `DS_C_MEM[479]` | 0x6003D1DF | 8 |  | - |
| `DS_C_MEM[47]` | 0x6003D02F | 8 |  | - |
| `DS_C_MEM[480]` | 0x6003D1E0 | 8 |  | - |
| `DS_C_MEM[481]` | 0x6003D1E1 | 8 |  | - |
| `DS_C_MEM[482]` | 0x6003D1E2 | 8 |  | - |
| `DS_C_MEM[483]` | 0x6003D1E3 | 8 |  | - |
| `DS_C_MEM[484]` | 0x6003D1E4 | 8 |  | - |
| `DS_C_MEM[485]` | 0x6003D1E5 | 8 |  | - |
| `DS_C_MEM[486]` | 0x6003D1E6 | 8 |  | - |
| `DS_C_MEM[487]` | 0x6003D1E7 | 8 |  | - |
| `DS_C_MEM[488]` | 0x6003D1E8 | 8 |  | - |
| `DS_C_MEM[489]` | 0x6003D1E9 | 8 |  | - |
| `DS_C_MEM[48]` | 0x6003D030 | 8 |  | - |
| `DS_C_MEM[490]` | 0x6003D1EA | 8 |  | - |
| `DS_C_MEM[491]` | 0x6003D1EB | 8 |  | - |
| `DS_C_MEM[492]` | 0x6003D1EC | 8 |  | - |
| `DS_C_MEM[493]` | 0x6003D1ED | 8 |  | - |
| `DS_C_MEM[494]` | 0x6003D1EE | 8 |  | - |
| `DS_C_MEM[495]` | 0x6003D1EF | 8 |  | - |
| `DS_C_MEM[496]` | 0x6003D1F0 | 8 |  | - |
| `DS_C_MEM[497]` | 0x6003D1F1 | 8 |  | - |
| `DS_C_MEM[498]` | 0x6003D1F2 | 8 |  | - |
| `DS_C_MEM[499]` | 0x6003D1F3 | 8 |  | - |
| `DS_C_MEM[49]` | 0x6003D031 | 8 |  | - |
| `DS_C_MEM[4]` | 0x6003D004 | 8 |  | - |
| `DS_C_MEM[500]` | 0x6003D1F4 | 8 |  | - |
| `DS_C_MEM[501]` | 0x6003D1F5 | 8 |  | - |
| `DS_C_MEM[502]` | 0x6003D1F6 | 8 |  | - |
| `DS_C_MEM[503]` | 0x6003D1F7 | 8 |  | - |
| `DS_C_MEM[504]` | 0x6003D1F8 | 8 |  | - |
| `DS_C_MEM[505]` | 0x6003D1F9 | 8 |  | - |
| `DS_C_MEM[506]` | 0x6003D1FA | 8 |  | - |
| `DS_C_MEM[507]` | 0x6003D1FB | 8 |  | - |
| `DS_C_MEM[508]` | 0x6003D1FC | 8 |  | - |
| `DS_C_MEM[509]` | 0x6003D1FD | 8 |  | - |
| `DS_C_MEM[50]` | 0x6003D032 | 8 |  | - |
| `DS_C_MEM[510]` | 0x6003D1FE | 8 |  | - |
| `DS_C_MEM[511]` | 0x6003D1FF | 8 |  | - |
| `DS_C_MEM[512]` | 0x6003D200 | 8 |  | - |
| `DS_C_MEM[513]` | 0x6003D201 | 8 |  | - |
| `DS_C_MEM[514]` | 0x6003D202 | 8 |  | - |
| `DS_C_MEM[515]` | 0x6003D203 | 8 |  | - |
| `DS_C_MEM[516]` | 0x6003D204 | 8 |  | - |
| `DS_C_MEM[517]` | 0x6003D205 | 8 |  | - |
| `DS_C_MEM[518]` | 0x6003D206 | 8 |  | - |
| `DS_C_MEM[519]` | 0x6003D207 | 8 |  | - |
| `DS_C_MEM[51]` | 0x6003D033 | 8 |  | - |
| `DS_C_MEM[520]` | 0x6003D208 | 8 |  | - |
| `DS_C_MEM[521]` | 0x6003D209 | 8 |  | - |
| `DS_C_MEM[522]` | 0x6003D20A | 8 |  | - |
| `DS_C_MEM[523]` | 0x6003D20B | 8 |  | - |
| `DS_C_MEM[524]` | 0x6003D20C | 8 |  | - |
| `DS_C_MEM[525]` | 0x6003D20D | 8 |  | - |
| `DS_C_MEM[526]` | 0x6003D20E | 8 |  | - |
| `DS_C_MEM[527]` | 0x6003D20F | 8 |  | - |
| `DS_C_MEM[528]` | 0x6003D210 | 8 |  | - |
| `DS_C_MEM[529]` | 0x6003D211 | 8 |  | - |
| `DS_C_MEM[52]` | 0x6003D034 | 8 |  | - |
| `DS_C_MEM[530]` | 0x6003D212 | 8 |  | - |
| `DS_C_MEM[531]` | 0x6003D213 | 8 |  | - |
| `DS_C_MEM[532]` | 0x6003D214 | 8 |  | - |
| `DS_C_MEM[533]` | 0x6003D215 | 8 |  | - |
| `DS_C_MEM[534]` | 0x6003D216 | 8 |  | - |
| `DS_C_MEM[535]` | 0x6003D217 | 8 |  | - |
| `DS_C_MEM[536]` | 0x6003D218 | 8 |  | - |
| `DS_C_MEM[537]` | 0x6003D219 | 8 |  | - |
| `DS_C_MEM[538]` | 0x6003D21A | 8 |  | - |
| `DS_C_MEM[539]` | 0x6003D21B | 8 |  | - |
| `DS_C_MEM[53]` | 0x6003D035 | 8 |  | - |
| `DS_C_MEM[540]` | 0x6003D21C | 8 |  | - |
| `DS_C_MEM[541]` | 0x6003D21D | 8 |  | - |
| `DS_C_MEM[542]` | 0x6003D21E | 8 |  | - |
| `DS_C_MEM[543]` | 0x6003D21F | 8 |  | - |
| `DS_C_MEM[544]` | 0x6003D220 | 8 |  | - |
| `DS_C_MEM[545]` | 0x6003D221 | 8 |  | - |
| `DS_C_MEM[546]` | 0x6003D222 | 8 |  | - |
| `DS_C_MEM[547]` | 0x6003D223 | 8 |  | - |
| `DS_C_MEM[548]` | 0x6003D224 | 8 |  | - |
| `DS_C_MEM[549]` | 0x6003D225 | 8 |  | - |
| `DS_C_MEM[54]` | 0x6003D036 | 8 |  | - |
| `DS_C_MEM[550]` | 0x6003D226 | 8 |  | - |
| `DS_C_MEM[551]` | 0x6003D227 | 8 |  | - |
| `DS_C_MEM[552]` | 0x6003D228 | 8 |  | - |
| `DS_C_MEM[553]` | 0x6003D229 | 8 |  | - |
| `DS_C_MEM[554]` | 0x6003D22A | 8 |  | - |
| `DS_C_MEM[555]` | 0x6003D22B | 8 |  | - |
| `DS_C_MEM[556]` | 0x6003D22C | 8 |  | - |
| `DS_C_MEM[557]` | 0x6003D22D | 8 |  | - |
| `DS_C_MEM[558]` | 0x6003D22E | 8 |  | - |
| `DS_C_MEM[559]` | 0x6003D22F | 8 |  | - |
| `DS_C_MEM[55]` | 0x6003D037 | 8 |  | - |
| `DS_C_MEM[560]` | 0x6003D230 | 8 |  | - |
| `DS_C_MEM[561]` | 0x6003D231 | 8 |  | - |
| `DS_C_MEM[562]` | 0x6003D232 | 8 |  | - |
| `DS_C_MEM[563]` | 0x6003D233 | 8 |  | - |
| `DS_C_MEM[564]` | 0x6003D234 | 8 |  | - |
| `DS_C_MEM[565]` | 0x6003D235 | 8 |  | - |
| `DS_C_MEM[566]` | 0x6003D236 | 8 |  | - |
| `DS_C_MEM[567]` | 0x6003D237 | 8 |  | - |
| `DS_C_MEM[568]` | 0x6003D238 | 8 |  | - |
| `DS_C_MEM[569]` | 0x6003D239 | 8 |  | - |
| `DS_C_MEM[56]` | 0x6003D038 | 8 |  | - |
| `DS_C_MEM[570]` | 0x6003D23A | 8 |  | - |
| `DS_C_MEM[571]` | 0x6003D23B | 8 |  | - |
| `DS_C_MEM[572]` | 0x6003D23C | 8 |  | - |
| `DS_C_MEM[573]` | 0x6003D23D | 8 |  | - |
| `DS_C_MEM[574]` | 0x6003D23E | 8 |  | - |
| `DS_C_MEM[575]` | 0x6003D23F | 8 |  | - |
| `DS_C_MEM[576]` | 0x6003D240 | 8 |  | - |
| `DS_C_MEM[577]` | 0x6003D241 | 8 |  | - |
| `DS_C_MEM[578]` | 0x6003D242 | 8 |  | - |
| `DS_C_MEM[579]` | 0x6003D243 | 8 |  | - |
| `DS_C_MEM[57]` | 0x6003D039 | 8 |  | - |
| `DS_C_MEM[580]` | 0x6003D244 | 8 |  | - |
| `DS_C_MEM[581]` | 0x6003D245 | 8 |  | - |
| `DS_C_MEM[582]` | 0x6003D246 | 8 |  | - |
| `DS_C_MEM[583]` | 0x6003D247 | 8 |  | - |
| `DS_C_MEM[584]` | 0x6003D248 | 8 |  | - |
| `DS_C_MEM[585]` | 0x6003D249 | 8 |  | - |
| `DS_C_MEM[586]` | 0x6003D24A | 8 |  | - |
| `DS_C_MEM[587]` | 0x6003D24B | 8 |  | - |
| `DS_C_MEM[588]` | 0x6003D24C | 8 |  | - |
| `DS_C_MEM[589]` | 0x6003D24D | 8 |  | - |
| `DS_C_MEM[58]` | 0x6003D03A | 8 |  | - |
| `DS_C_MEM[590]` | 0x6003D24E | 8 |  | - |
| `DS_C_MEM[591]` | 0x6003D24F | 8 |  | - |
| `DS_C_MEM[592]` | 0x6003D250 | 8 |  | - |
| `DS_C_MEM[593]` | 0x6003D251 | 8 |  | - |
| `DS_C_MEM[594]` | 0x6003D252 | 8 |  | - |
| `DS_C_MEM[595]` | 0x6003D253 | 8 |  | - |
| `DS_C_MEM[596]` | 0x6003D254 | 8 |  | - |
| `DS_C_MEM[597]` | 0x6003D255 | 8 |  | - |
| `DS_C_MEM[598]` | 0x6003D256 | 8 |  | - |
| `DS_C_MEM[599]` | 0x6003D257 | 8 |  | - |
| `DS_C_MEM[59]` | 0x6003D03B | 8 |  | - |
| `DS_C_MEM[5]` | 0x6003D005 | 8 |  | - |
| `DS_C_MEM[600]` | 0x6003D258 | 8 |  | - |
| `DS_C_MEM[601]` | 0x6003D259 | 8 |  | - |
| `DS_C_MEM[602]` | 0x6003D25A | 8 |  | - |
| `DS_C_MEM[603]` | 0x6003D25B | 8 |  | - |
| `DS_C_MEM[604]` | 0x6003D25C | 8 |  | - |
| `DS_C_MEM[605]` | 0x6003D25D | 8 |  | - |
| `DS_C_MEM[606]` | 0x6003D25E | 8 |  | - |
| `DS_C_MEM[607]` | 0x6003D25F | 8 |  | - |
| `DS_C_MEM[608]` | 0x6003D260 | 8 |  | - |
| `DS_C_MEM[609]` | 0x6003D261 | 8 |  | - |
| `DS_C_MEM[60]` | 0x6003D03C | 8 |  | - |
| `DS_C_MEM[610]` | 0x6003D262 | 8 |  | - |
| `DS_C_MEM[611]` | 0x6003D263 | 8 |  | - |
| `DS_C_MEM[612]` | 0x6003D264 | 8 |  | - |
| `DS_C_MEM[613]` | 0x6003D265 | 8 |  | - |
| `DS_C_MEM[614]` | 0x6003D266 | 8 |  | - |
| `DS_C_MEM[615]` | 0x6003D267 | 8 |  | - |
| `DS_C_MEM[616]` | 0x6003D268 | 8 |  | - |
| `DS_C_MEM[617]` | 0x6003D269 | 8 |  | - |
| `DS_C_MEM[618]` | 0x6003D26A | 8 |  | - |
| `DS_C_MEM[619]` | 0x6003D26B | 8 |  | - |
| `DS_C_MEM[61]` | 0x6003D03D | 8 |  | - |
| `DS_C_MEM[620]` | 0x6003D26C | 8 |  | - |
| `DS_C_MEM[621]` | 0x6003D26D | 8 |  | - |
| `DS_C_MEM[622]` | 0x6003D26E | 8 |  | - |
| `DS_C_MEM[623]` | 0x6003D26F | 8 |  | - |
| `DS_C_MEM[624]` | 0x6003D270 | 8 |  | - |
| `DS_C_MEM[625]` | 0x6003D271 | 8 |  | - |
| `DS_C_MEM[626]` | 0x6003D272 | 8 |  | - |
| `DS_C_MEM[627]` | 0x6003D273 | 8 |  | - |
| `DS_C_MEM[628]` | 0x6003D274 | 8 |  | - |
| `DS_C_MEM[629]` | 0x6003D275 | 8 |  | - |
| `DS_C_MEM[62]` | 0x6003D03E | 8 |  | - |
| `DS_C_MEM[630]` | 0x6003D276 | 8 |  | - |
| `DS_C_MEM[631]` | 0x6003D277 | 8 |  | - |
| `DS_C_MEM[632]` | 0x6003D278 | 8 |  | - |
| `DS_C_MEM[633]` | 0x6003D279 | 8 |  | - |
| `DS_C_MEM[634]` | 0x6003D27A | 8 |  | - |
| `DS_C_MEM[635]` | 0x6003D27B | 8 |  | - |
| `DS_C_MEM[636]` | 0x6003D27C | 8 |  | - |
| `DS_C_MEM[637]` | 0x6003D27D | 8 |  | - |
| `DS_C_MEM[638]` | 0x6003D27E | 8 |  | - |
| `DS_C_MEM[639]` | 0x6003D27F | 8 |  | - |
| `DS_C_MEM[63]` | 0x6003D03F | 8 |  | - |
| `DS_C_MEM[640]` | 0x6003D280 | 8 |  | - |
| `DS_C_MEM[641]` | 0x6003D281 | 8 |  | - |
| `DS_C_MEM[642]` | 0x6003D282 | 8 |  | - |
| `DS_C_MEM[643]` | 0x6003D283 | 8 |  | - |
| `DS_C_MEM[644]` | 0x6003D284 | 8 |  | - |
| `DS_C_MEM[645]` | 0x6003D285 | 8 |  | - |
| `DS_C_MEM[646]` | 0x6003D286 | 8 |  | - |
| `DS_C_MEM[647]` | 0x6003D287 | 8 |  | - |
| `DS_C_MEM[648]` | 0x6003D288 | 8 |  | - |
| `DS_C_MEM[649]` | 0x6003D289 | 8 |  | - |
| `DS_C_MEM[64]` | 0x6003D040 | 8 |  | - |
| `DS_C_MEM[650]` | 0x6003D28A | 8 |  | - |
| `DS_C_MEM[651]` | 0x6003D28B | 8 |  | - |
| `DS_C_MEM[652]` | 0x6003D28C | 8 |  | - |
| `DS_C_MEM[653]` | 0x6003D28D | 8 |  | - |
| `DS_C_MEM[654]` | 0x6003D28E | 8 |  | - |
| `DS_C_MEM[655]` | 0x6003D28F | 8 |  | - |
| `DS_C_MEM[656]` | 0x6003D290 | 8 |  | - |
| `DS_C_MEM[657]` | 0x6003D291 | 8 |  | - |
| `DS_C_MEM[658]` | 0x6003D292 | 8 |  | - |
| `DS_C_MEM[659]` | 0x6003D293 | 8 |  | - |
| `DS_C_MEM[65]` | 0x6003D041 | 8 |  | - |
| `DS_C_MEM[660]` | 0x6003D294 | 8 |  | - |
| `DS_C_MEM[661]` | 0x6003D295 | 8 |  | - |
| `DS_C_MEM[662]` | 0x6003D296 | 8 |  | - |
| `DS_C_MEM[663]` | 0x6003D297 | 8 |  | - |
| `DS_C_MEM[664]` | 0x6003D298 | 8 |  | - |
| `DS_C_MEM[665]` | 0x6003D299 | 8 |  | - |
| `DS_C_MEM[666]` | 0x6003D29A | 8 |  | - |
| `DS_C_MEM[667]` | 0x6003D29B | 8 |  | - |
| `DS_C_MEM[668]` | 0x6003D29C | 8 |  | - |
| `DS_C_MEM[669]` | 0x6003D29D | 8 |  | - |
| `DS_C_MEM[66]` | 0x6003D042 | 8 |  | - |
| `DS_C_MEM[670]` | 0x6003D29E | 8 |  | - |
| `DS_C_MEM[671]` | 0x6003D29F | 8 |  | - |
| `DS_C_MEM[672]` | 0x6003D2A0 | 8 |  | - |
| `DS_C_MEM[673]` | 0x6003D2A1 | 8 |  | - |
| `DS_C_MEM[674]` | 0x6003D2A2 | 8 |  | - |
| `DS_C_MEM[675]` | 0x6003D2A3 | 8 |  | - |
| `DS_C_MEM[676]` | 0x6003D2A4 | 8 |  | - |
| `DS_C_MEM[677]` | 0x6003D2A5 | 8 |  | - |
| `DS_C_MEM[678]` | 0x6003D2A6 | 8 |  | - |
| `DS_C_MEM[679]` | 0x6003D2A7 | 8 |  | - |
| `DS_C_MEM[67]` | 0x6003D043 | 8 |  | - |
| `DS_C_MEM[680]` | 0x6003D2A8 | 8 |  | - |
| `DS_C_MEM[681]` | 0x6003D2A9 | 8 |  | - |
| `DS_C_MEM[682]` | 0x6003D2AA | 8 |  | - |
| `DS_C_MEM[683]` | 0x6003D2AB | 8 |  | - |
| `DS_C_MEM[684]` | 0x6003D2AC | 8 |  | - |
| `DS_C_MEM[685]` | 0x6003D2AD | 8 |  | - |
| `DS_C_MEM[686]` | 0x6003D2AE | 8 |  | - |
| `DS_C_MEM[687]` | 0x6003D2AF | 8 |  | - |
| `DS_C_MEM[688]` | 0x6003D2B0 | 8 |  | - |
| `DS_C_MEM[689]` | 0x6003D2B1 | 8 |  | - |
| `DS_C_MEM[68]` | 0x6003D044 | 8 |  | - |
| `DS_C_MEM[690]` | 0x6003D2B2 | 8 |  | - |
| `DS_C_MEM[691]` | 0x6003D2B3 | 8 |  | - |
| `DS_C_MEM[692]` | 0x6003D2B4 | 8 |  | - |
| `DS_C_MEM[693]` | 0x6003D2B5 | 8 |  | - |
| `DS_C_MEM[694]` | 0x6003D2B6 | 8 |  | - |
| `DS_C_MEM[695]` | 0x6003D2B7 | 8 |  | - |
| `DS_C_MEM[696]` | 0x6003D2B8 | 8 |  | - |
| `DS_C_MEM[697]` | 0x6003D2B9 | 8 |  | - |
| `DS_C_MEM[698]` | 0x6003D2BA | 8 |  | - |
| `DS_C_MEM[699]` | 0x6003D2BB | 8 |  | - |
| `DS_C_MEM[69]` | 0x6003D045 | 8 |  | - |
| `DS_C_MEM[6]` | 0x6003D006 | 8 |  | - |
| `DS_C_MEM[700]` | 0x6003D2BC | 8 |  | - |
| `DS_C_MEM[701]` | 0x6003D2BD | 8 |  | - |
| `DS_C_MEM[702]` | 0x6003D2BE | 8 |  | - |
| `DS_C_MEM[703]` | 0x6003D2BF | 8 |  | - |
| `DS_C_MEM[704]` | 0x6003D2C0 | 8 |  | - |
| `DS_C_MEM[705]` | 0x6003D2C1 | 8 |  | - |
| `DS_C_MEM[706]` | 0x6003D2C2 | 8 |  | - |
| `DS_C_MEM[707]` | 0x6003D2C3 | 8 |  | - |
| `DS_C_MEM[708]` | 0x6003D2C4 | 8 |  | - |
| `DS_C_MEM[709]` | 0x6003D2C5 | 8 |  | - |
| `DS_C_MEM[70]` | 0x6003D046 | 8 |  | - |
| `DS_C_MEM[710]` | 0x6003D2C6 | 8 |  | - |
| `DS_C_MEM[711]` | 0x6003D2C7 | 8 |  | - |
| `DS_C_MEM[712]` | 0x6003D2C8 | 8 |  | - |
| `DS_C_MEM[713]` | 0x6003D2C9 | 8 |  | - |
| `DS_C_MEM[714]` | 0x6003D2CA | 8 |  | - |
| `DS_C_MEM[715]` | 0x6003D2CB | 8 |  | - |
| `DS_C_MEM[716]` | 0x6003D2CC | 8 |  | - |
| `DS_C_MEM[717]` | 0x6003D2CD | 8 |  | - |
| `DS_C_MEM[718]` | 0x6003D2CE | 8 |  | - |
| `DS_C_MEM[719]` | 0x6003D2CF | 8 |  | - |
| `DS_C_MEM[71]` | 0x6003D047 | 8 |  | - |
| `DS_C_MEM[720]` | 0x6003D2D0 | 8 |  | - |
| `DS_C_MEM[721]` | 0x6003D2D1 | 8 |  | - |
| `DS_C_MEM[722]` | 0x6003D2D2 | 8 |  | - |
| `DS_C_MEM[723]` | 0x6003D2D3 | 8 |  | - |
| `DS_C_MEM[724]` | 0x6003D2D4 | 8 |  | - |
| `DS_C_MEM[725]` | 0x6003D2D5 | 8 |  | - |
| `DS_C_MEM[726]` | 0x6003D2D6 | 8 |  | - |
| `DS_C_MEM[727]` | 0x6003D2D7 | 8 |  | - |
| `DS_C_MEM[728]` | 0x6003D2D8 | 8 |  | - |
| `DS_C_MEM[729]` | 0x6003D2D9 | 8 |  | - |
| `DS_C_MEM[72]` | 0x6003D048 | 8 |  | - |
| `DS_C_MEM[730]` | 0x6003D2DA | 8 |  | - |
| `DS_C_MEM[731]` | 0x6003D2DB | 8 |  | - |
| `DS_C_MEM[732]` | 0x6003D2DC | 8 |  | - |
| `DS_C_MEM[733]` | 0x6003D2DD | 8 |  | - |
| `DS_C_MEM[734]` | 0x6003D2DE | 8 |  | - |
| `DS_C_MEM[735]` | 0x6003D2DF | 8 |  | - |
| `DS_C_MEM[736]` | 0x6003D2E0 | 8 |  | - |
| `DS_C_MEM[737]` | 0x6003D2E1 | 8 |  | - |
| `DS_C_MEM[738]` | 0x6003D2E2 | 8 |  | - |
| `DS_C_MEM[739]` | 0x6003D2E3 | 8 |  | - |
| `DS_C_MEM[73]` | 0x6003D049 | 8 |  | - |
| `DS_C_MEM[740]` | 0x6003D2E4 | 8 |  | - |
| `DS_C_MEM[741]` | 0x6003D2E5 | 8 |  | - |
| `DS_C_MEM[742]` | 0x6003D2E6 | 8 |  | - |
| `DS_C_MEM[743]` | 0x6003D2E7 | 8 |  | - |
| `DS_C_MEM[744]` | 0x6003D2E8 | 8 |  | - |
| `DS_C_MEM[745]` | 0x6003D2E9 | 8 |  | - |
| `DS_C_MEM[746]` | 0x6003D2EA | 8 |  | - |
| `DS_C_MEM[747]` | 0x6003D2EB | 8 |  | - |
| `DS_C_MEM[748]` | 0x6003D2EC | 8 |  | - |
| `DS_C_MEM[749]` | 0x6003D2ED | 8 |  | - |
| `DS_C_MEM[74]` | 0x6003D04A | 8 |  | - |
| `DS_C_MEM[750]` | 0x6003D2EE | 8 |  | - |
| `DS_C_MEM[751]` | 0x6003D2EF | 8 |  | - |
| `DS_C_MEM[752]` | 0x6003D2F0 | 8 |  | - |
| `DS_C_MEM[753]` | 0x6003D2F1 | 8 |  | - |
| `DS_C_MEM[754]` | 0x6003D2F2 | 8 |  | - |
| `DS_C_MEM[755]` | 0x6003D2F3 | 8 |  | - |
| `DS_C_MEM[756]` | 0x6003D2F4 | 8 |  | - |
| `DS_C_MEM[757]` | 0x6003D2F5 | 8 |  | - |
| `DS_C_MEM[758]` | 0x6003D2F6 | 8 |  | - |
| `DS_C_MEM[759]` | 0x6003D2F7 | 8 |  | - |
| `DS_C_MEM[75]` | 0x6003D04B | 8 |  | - |
| `DS_C_MEM[760]` | 0x6003D2F8 | 8 |  | - |
| `DS_C_MEM[761]` | 0x6003D2F9 | 8 |  | - |
| `DS_C_MEM[762]` | 0x6003D2FA | 8 |  | - |
| `DS_C_MEM[763]` | 0x6003D2FB | 8 |  | - |
| `DS_C_MEM[764]` | 0x6003D2FC | 8 |  | - |
| `DS_C_MEM[765]` | 0x6003D2FD | 8 |  | - |
| `DS_C_MEM[766]` | 0x6003D2FE | 8 |  | - |
| `DS_C_MEM[767]` | 0x6003D2FF | 8 |  | - |
| `DS_C_MEM[768]` | 0x6003D300 | 8 |  | - |
| `DS_C_MEM[769]` | 0x6003D301 | 8 |  | - |
| `DS_C_MEM[76]` | 0x6003D04C | 8 |  | - |
| `DS_C_MEM[770]` | 0x6003D302 | 8 |  | - |
| `DS_C_MEM[771]` | 0x6003D303 | 8 |  | - |
| `DS_C_MEM[772]` | 0x6003D304 | 8 |  | - |
| `DS_C_MEM[773]` | 0x6003D305 | 8 |  | - |
| `DS_C_MEM[774]` | 0x6003D306 | 8 |  | - |
| `DS_C_MEM[775]` | 0x6003D307 | 8 |  | - |
| `DS_C_MEM[776]` | 0x6003D308 | 8 |  | - |
| `DS_C_MEM[777]` | 0x6003D309 | 8 |  | - |
| `DS_C_MEM[778]` | 0x6003D30A | 8 |  | - |
| `DS_C_MEM[779]` | 0x6003D30B | 8 |  | - |
| `DS_C_MEM[77]` | 0x6003D04D | 8 |  | - |
| `DS_C_MEM[780]` | 0x6003D30C | 8 |  | - |
| `DS_C_MEM[781]` | 0x6003D30D | 8 |  | - |
| `DS_C_MEM[782]` | 0x6003D30E | 8 |  | - |
| `DS_C_MEM[783]` | 0x6003D30F | 8 |  | - |
| `DS_C_MEM[784]` | 0x6003D310 | 8 |  | - |
| `DS_C_MEM[785]` | 0x6003D311 | 8 |  | - |
| `DS_C_MEM[786]` | 0x6003D312 | 8 |  | - |
| `DS_C_MEM[787]` | 0x6003D313 | 8 |  | - |
| `DS_C_MEM[788]` | 0x6003D314 | 8 |  | - |
| `DS_C_MEM[789]` | 0x6003D315 | 8 |  | - |
| `DS_C_MEM[78]` | 0x6003D04E | 8 |  | - |
| `DS_C_MEM[790]` | 0x6003D316 | 8 |  | - |
| `DS_C_MEM[791]` | 0x6003D317 | 8 |  | - |
| `DS_C_MEM[792]` | 0x6003D318 | 8 |  | - |
| `DS_C_MEM[793]` | 0x6003D319 | 8 |  | - |
| `DS_C_MEM[794]` | 0x6003D31A | 8 |  | - |
| `DS_C_MEM[795]` | 0x6003D31B | 8 |  | - |
| `DS_C_MEM[796]` | 0x6003D31C | 8 |  | - |
| `DS_C_MEM[797]` | 0x6003D31D | 8 |  | - |
| `DS_C_MEM[798]` | 0x6003D31E | 8 |  | - |
| `DS_C_MEM[799]` | 0x6003D31F | 8 |  | - |
| `DS_C_MEM[79]` | 0x6003D04F | 8 |  | - |
| `DS_C_MEM[7]` | 0x6003D007 | 8 |  | - |
| `DS_C_MEM[800]` | 0x6003D320 | 8 |  | - |
| `DS_C_MEM[801]` | 0x6003D321 | 8 |  | - |
| `DS_C_MEM[802]` | 0x6003D322 | 8 |  | - |
| `DS_C_MEM[803]` | 0x6003D323 | 8 |  | - |
| `DS_C_MEM[804]` | 0x6003D324 | 8 |  | - |
| `DS_C_MEM[805]` | 0x6003D325 | 8 |  | - |
| `DS_C_MEM[806]` | 0x6003D326 | 8 |  | - |
| `DS_C_MEM[807]` | 0x6003D327 | 8 |  | - |
| `DS_C_MEM[808]` | 0x6003D328 | 8 |  | - |
| `DS_C_MEM[809]` | 0x6003D329 | 8 |  | - |
| `DS_C_MEM[80]` | 0x6003D050 | 8 |  | - |
| `DS_C_MEM[810]` | 0x6003D32A | 8 |  | - |
| `DS_C_MEM[811]` | 0x6003D32B | 8 |  | - |
| `DS_C_MEM[812]` | 0x6003D32C | 8 |  | - |
| `DS_C_MEM[813]` | 0x6003D32D | 8 |  | - |
| `DS_C_MEM[814]` | 0x6003D32E | 8 |  | - |
| `DS_C_MEM[815]` | 0x6003D32F | 8 |  | - |
| `DS_C_MEM[816]` | 0x6003D330 | 8 |  | - |
| `DS_C_MEM[817]` | 0x6003D331 | 8 |  | - |
| `DS_C_MEM[818]` | 0x6003D332 | 8 |  | - |
| `DS_C_MEM[819]` | 0x6003D333 | 8 |  | - |
| `DS_C_MEM[81]` | 0x6003D051 | 8 |  | - |
| `DS_C_MEM[820]` | 0x6003D334 | 8 |  | - |
| `DS_C_MEM[821]` | 0x6003D335 | 8 |  | - |
| `DS_C_MEM[822]` | 0x6003D336 | 8 |  | - |
| `DS_C_MEM[823]` | 0x6003D337 | 8 |  | - |
| `DS_C_MEM[824]` | 0x6003D338 | 8 |  | - |
| `DS_C_MEM[825]` | 0x6003D339 | 8 |  | - |
| `DS_C_MEM[826]` | 0x6003D33A | 8 |  | - |
| `DS_C_MEM[827]` | 0x6003D33B | 8 |  | - |
| `DS_C_MEM[828]` | 0x6003D33C | 8 |  | - |
| `DS_C_MEM[829]` | 0x6003D33D | 8 |  | - |
| `DS_C_MEM[82]` | 0x6003D052 | 8 |  | - |
| `DS_C_MEM[830]` | 0x6003D33E | 8 |  | - |
| `DS_C_MEM[831]` | 0x6003D33F | 8 |  | - |
| `DS_C_MEM[832]` | 0x6003D340 | 8 |  | - |
| `DS_C_MEM[833]` | 0x6003D341 | 8 |  | - |
| `DS_C_MEM[834]` | 0x6003D342 | 8 |  | - |
| `DS_C_MEM[835]` | 0x6003D343 | 8 |  | - |
| `DS_C_MEM[836]` | 0x6003D344 | 8 |  | - |
| `DS_C_MEM[837]` | 0x6003D345 | 8 |  | - |
| `DS_C_MEM[838]` | 0x6003D346 | 8 |  | - |
| `DS_C_MEM[839]` | 0x6003D347 | 8 |  | - |
| `DS_C_MEM[83]` | 0x6003D053 | 8 |  | - |
| `DS_C_MEM[840]` | 0x6003D348 | 8 |  | - |
| `DS_C_MEM[841]` | 0x6003D349 | 8 |  | - |
| `DS_C_MEM[842]` | 0x6003D34A | 8 |  | - |
| `DS_C_MEM[843]` | 0x6003D34B | 8 |  | - |
| `DS_C_MEM[844]` | 0x6003D34C | 8 |  | - |
| `DS_C_MEM[845]` | 0x6003D34D | 8 |  | - |
| `DS_C_MEM[846]` | 0x6003D34E | 8 |  | - |
| `DS_C_MEM[847]` | 0x6003D34F | 8 |  | - |
| `DS_C_MEM[848]` | 0x6003D350 | 8 |  | - |
| `DS_C_MEM[849]` | 0x6003D351 | 8 |  | - |
| `DS_C_MEM[84]` | 0x6003D054 | 8 |  | - |
| `DS_C_MEM[850]` | 0x6003D352 | 8 |  | - |
| `DS_C_MEM[851]` | 0x6003D353 | 8 |  | - |
| `DS_C_MEM[852]` | 0x6003D354 | 8 |  | - |
| `DS_C_MEM[853]` | 0x6003D355 | 8 |  | - |
| `DS_C_MEM[854]` | 0x6003D356 | 8 |  | - |
| `DS_C_MEM[855]` | 0x6003D357 | 8 |  | - |
| `DS_C_MEM[856]` | 0x6003D358 | 8 |  | - |
| `DS_C_MEM[857]` | 0x6003D359 | 8 |  | - |
| `DS_C_MEM[858]` | 0x6003D35A | 8 |  | - |
| `DS_C_MEM[859]` | 0x6003D35B | 8 |  | - |
| `DS_C_MEM[85]` | 0x6003D055 | 8 |  | - |
| `DS_C_MEM[860]` | 0x6003D35C | 8 |  | - |
| `DS_C_MEM[861]` | 0x6003D35D | 8 |  | - |
| `DS_C_MEM[862]` | 0x6003D35E | 8 |  | - |
| `DS_C_MEM[863]` | 0x6003D35F | 8 |  | - |
| `DS_C_MEM[864]` | 0x6003D360 | 8 |  | - |
| `DS_C_MEM[865]` | 0x6003D361 | 8 |  | - |
| `DS_C_MEM[866]` | 0x6003D362 | 8 |  | - |
| `DS_C_MEM[867]` | 0x6003D363 | 8 |  | - |
| `DS_C_MEM[868]` | 0x6003D364 | 8 |  | - |
| `DS_C_MEM[869]` | 0x6003D365 | 8 |  | - |
| `DS_C_MEM[86]` | 0x6003D056 | 8 |  | - |
| `DS_C_MEM[870]` | 0x6003D366 | 8 |  | - |
| `DS_C_MEM[871]` | 0x6003D367 | 8 |  | - |
| `DS_C_MEM[872]` | 0x6003D368 | 8 |  | - |
| `DS_C_MEM[873]` | 0x6003D369 | 8 |  | - |
| `DS_C_MEM[874]` | 0x6003D36A | 8 |  | - |
| `DS_C_MEM[875]` | 0x6003D36B | 8 |  | - |
| `DS_C_MEM[876]` | 0x6003D36C | 8 |  | - |
| `DS_C_MEM[877]` | 0x6003D36D | 8 |  | - |
| `DS_C_MEM[878]` | 0x6003D36E | 8 |  | - |
| `DS_C_MEM[879]` | 0x6003D36F | 8 |  | - |
| `DS_C_MEM[87]` | 0x6003D057 | 8 |  | - |
| `DS_C_MEM[880]` | 0x6003D370 | 8 |  | - |
| `DS_C_MEM[881]` | 0x6003D371 | 8 |  | - |
| `DS_C_MEM[882]` | 0x6003D372 | 8 |  | - |
| `DS_C_MEM[883]` | 0x6003D373 | 8 |  | - |
| `DS_C_MEM[884]` | 0x6003D374 | 8 |  | - |
| `DS_C_MEM[885]` | 0x6003D375 | 8 |  | - |
| `DS_C_MEM[886]` | 0x6003D376 | 8 |  | - |
| `DS_C_MEM[887]` | 0x6003D377 | 8 |  | - |
| `DS_C_MEM[888]` | 0x6003D378 | 8 |  | - |
| `DS_C_MEM[889]` | 0x6003D379 | 8 |  | - |
| `DS_C_MEM[88]` | 0x6003D058 | 8 |  | - |
| `DS_C_MEM[890]` | 0x6003D37A | 8 |  | - |
| `DS_C_MEM[891]` | 0x6003D37B | 8 |  | - |
| `DS_C_MEM[892]` | 0x6003D37C | 8 |  | - |
| `DS_C_MEM[893]` | 0x6003D37D | 8 |  | - |
| `DS_C_MEM[894]` | 0x6003D37E | 8 |  | - |
| `DS_C_MEM[895]` | 0x6003D37F | 8 |  | - |
| `DS_C_MEM[896]` | 0x6003D380 | 8 |  | - |
| `DS_C_MEM[897]` | 0x6003D381 | 8 |  | - |
| `DS_C_MEM[898]` | 0x6003D382 | 8 |  | - |
| `DS_C_MEM[899]` | 0x6003D383 | 8 |  | - |
| `DS_C_MEM[89]` | 0x6003D059 | 8 |  | - |
| `DS_C_MEM[8]` | 0x6003D008 | 8 |  | - |
| `DS_C_MEM[900]` | 0x6003D384 | 8 |  | - |
| `DS_C_MEM[901]` | 0x6003D385 | 8 |  | - |
| `DS_C_MEM[902]` | 0x6003D386 | 8 |  | - |
| `DS_C_MEM[903]` | 0x6003D387 | 8 |  | - |
| `DS_C_MEM[904]` | 0x6003D388 | 8 |  | - |
| `DS_C_MEM[905]` | 0x6003D389 | 8 |  | - |
| `DS_C_MEM[906]` | 0x6003D38A | 8 |  | - |
| `DS_C_MEM[907]` | 0x6003D38B | 8 |  | - |
| `DS_C_MEM[908]` | 0x6003D38C | 8 |  | - |
| `DS_C_MEM[909]` | 0x6003D38D | 8 |  | - |
| `DS_C_MEM[90]` | 0x6003D05A | 8 |  | - |
| `DS_C_MEM[910]` | 0x6003D38E | 8 |  | - |
| `DS_C_MEM[911]` | 0x6003D38F | 8 |  | - |
| `DS_C_MEM[912]` | 0x6003D390 | 8 |  | - |
| `DS_C_MEM[913]` | 0x6003D391 | 8 |  | - |
| `DS_C_MEM[914]` | 0x6003D392 | 8 |  | - |
| `DS_C_MEM[915]` | 0x6003D393 | 8 |  | - |
| `DS_C_MEM[916]` | 0x6003D394 | 8 |  | - |
| `DS_C_MEM[917]` | 0x6003D395 | 8 |  | - |
| `DS_C_MEM[918]` | 0x6003D396 | 8 |  | - |
| `DS_C_MEM[919]` | 0x6003D397 | 8 |  | - |
| `DS_C_MEM[91]` | 0x6003D05B | 8 |  | - |
| `DS_C_MEM[920]` | 0x6003D398 | 8 |  | - |
| `DS_C_MEM[921]` | 0x6003D399 | 8 |  | - |
| `DS_C_MEM[922]` | 0x6003D39A | 8 |  | - |
| `DS_C_MEM[923]` | 0x6003D39B | 8 |  | - |
| `DS_C_MEM[924]` | 0x6003D39C | 8 |  | - |
| `DS_C_MEM[925]` | 0x6003D39D | 8 |  | - |
| `DS_C_MEM[926]` | 0x6003D39E | 8 |  | - |
| `DS_C_MEM[927]` | 0x6003D39F | 8 |  | - |
| `DS_C_MEM[928]` | 0x6003D3A0 | 8 |  | - |
| `DS_C_MEM[929]` | 0x6003D3A1 | 8 |  | - |
| `DS_C_MEM[92]` | 0x6003D05C | 8 |  | - |
| `DS_C_MEM[930]` | 0x6003D3A2 | 8 |  | - |
| `DS_C_MEM[931]` | 0x6003D3A3 | 8 |  | - |
| `DS_C_MEM[932]` | 0x6003D3A4 | 8 |  | - |
| `DS_C_MEM[933]` | 0x6003D3A5 | 8 |  | - |
| `DS_C_MEM[934]` | 0x6003D3A6 | 8 |  | - |
| `DS_C_MEM[935]` | 0x6003D3A7 | 8 |  | - |
| `DS_C_MEM[936]` | 0x6003D3A8 | 8 |  | - |
| `DS_C_MEM[937]` | 0x6003D3A9 | 8 |  | - |
| `DS_C_MEM[938]` | 0x6003D3AA | 8 |  | - |
| `DS_C_MEM[939]` | 0x6003D3AB | 8 |  | - |
| `DS_C_MEM[93]` | 0x6003D05D | 8 |  | - |
| `DS_C_MEM[940]` | 0x6003D3AC | 8 |  | - |
| `DS_C_MEM[941]` | 0x6003D3AD | 8 |  | - |
| `DS_C_MEM[942]` | 0x6003D3AE | 8 |  | - |
| `DS_C_MEM[943]` | 0x6003D3AF | 8 |  | - |
| `DS_C_MEM[944]` | 0x6003D3B0 | 8 |  | - |
| `DS_C_MEM[945]` | 0x6003D3B1 | 8 |  | - |
| `DS_C_MEM[946]` | 0x6003D3B2 | 8 |  | - |
| `DS_C_MEM[947]` | 0x6003D3B3 | 8 |  | - |
| `DS_C_MEM[948]` | 0x6003D3B4 | 8 |  | - |
| `DS_C_MEM[949]` | 0x6003D3B5 | 8 |  | - |
| `DS_C_MEM[94]` | 0x6003D05E | 8 |  | - |
| `DS_C_MEM[950]` | 0x6003D3B6 | 8 |  | - |
| `DS_C_MEM[951]` | 0x6003D3B7 | 8 |  | - |
| `DS_C_MEM[952]` | 0x6003D3B8 | 8 |  | - |
| `DS_C_MEM[953]` | 0x6003D3B9 | 8 |  | - |
| `DS_C_MEM[954]` | 0x6003D3BA | 8 |  | - |
| `DS_C_MEM[955]` | 0x6003D3BB | 8 |  | - |
| `DS_C_MEM[956]` | 0x6003D3BC | 8 |  | - |
| `DS_C_MEM[957]` | 0x6003D3BD | 8 |  | - |
| `DS_C_MEM[958]` | 0x6003D3BE | 8 |  | - |
| `DS_C_MEM[959]` | 0x6003D3BF | 8 |  | - |
| `DS_C_MEM[95]` | 0x6003D05F | 8 |  | - |
| `DS_C_MEM[960]` | 0x6003D3C0 | 8 |  | - |
| `DS_C_MEM[961]` | 0x6003D3C1 | 8 |  | - |
| `DS_C_MEM[962]` | 0x6003D3C2 | 8 |  | - |
| `DS_C_MEM[963]` | 0x6003D3C3 | 8 |  | - |
| `DS_C_MEM[964]` | 0x6003D3C4 | 8 |  | - |
| `DS_C_MEM[965]` | 0x6003D3C5 | 8 |  | - |
| `DS_C_MEM[966]` | 0x6003D3C6 | 8 |  | - |
| `DS_C_MEM[967]` | 0x6003D3C7 | 8 |  | - |
| `DS_C_MEM[968]` | 0x6003D3C8 | 8 |  | - |
| `DS_C_MEM[969]` | 0x6003D3C9 | 8 |  | - |
| `DS_C_MEM[96]` | 0x6003D060 | 8 |  | - |
| `DS_C_MEM[970]` | 0x6003D3CA | 8 |  | - |
| `DS_C_MEM[971]` | 0x6003D3CB | 8 |  | - |
| `DS_C_MEM[972]` | 0x6003D3CC | 8 |  | - |
| `DS_C_MEM[973]` | 0x6003D3CD | 8 |  | - |
| `DS_C_MEM[974]` | 0x6003D3CE | 8 |  | - |
| `DS_C_MEM[975]` | 0x6003D3CF | 8 |  | - |
| `DS_C_MEM[976]` | 0x6003D3D0 | 8 |  | - |
| `DS_C_MEM[977]` | 0x6003D3D1 | 8 |  | - |
| `DS_C_MEM[978]` | 0x6003D3D2 | 8 |  | - |
| `DS_C_MEM[979]` | 0x6003D3D3 | 8 |  | - |
| `DS_C_MEM[97]` | 0x6003D061 | 8 |  | - |
| `DS_C_MEM[980]` | 0x6003D3D4 | 8 |  | - |
| `DS_C_MEM[981]` | 0x6003D3D5 | 8 |  | - |
| `DS_C_MEM[982]` | 0x6003D3D6 | 8 |  | - |
| `DS_C_MEM[983]` | 0x6003D3D7 | 8 |  | - |
| `DS_C_MEM[984]` | 0x6003D3D8 | 8 |  | - |
| `DS_C_MEM[985]` | 0x6003D3D9 | 8 |  | - |
| `DS_C_MEM[986]` | 0x6003D3DA | 8 |  | - |
| `DS_C_MEM[987]` | 0x6003D3DB | 8 |  | - |
| `DS_C_MEM[988]` | 0x6003D3DC | 8 |  | - |
| `DS_C_MEM[989]` | 0x6003D3DD | 8 |  | - |
| `DS_C_MEM[98]` | 0x6003D062 | 8 |  | - |
| `DS_C_MEM[990]` | 0x6003D3DE | 8 |  | - |
| `DS_C_MEM[991]` | 0x6003D3DF | 8 |  | - |
| `DS_C_MEM[992]` | 0x6003D3E0 | 8 |  | - |
| `DS_C_MEM[993]` | 0x6003D3E1 | 8 |  | - |
| `DS_C_MEM[994]` | 0x6003D3E2 | 8 |  | - |
| `DS_C_MEM[995]` | 0x6003D3E3 | 8 |  | - |
| `DS_C_MEM[996]` | 0x6003D3E4 | 8 |  | - |
| `DS_C_MEM[997]` | 0x6003D3E5 | 8 |  | - |
| `DS_C_MEM[998]` | 0x6003D3E6 | 8 |  | - |
| `DS_C_MEM[999]` | 0x6003D3E7 | 8 |  | - |
| `DS_C_MEM[99]` | 0x6003D063 | 8 |  | - |
| `DS_C_MEM[9]` | 0x6003D009 | 8 |  | - |
| `DS_DATE` | 0x6003DE20 | 32 | DATE=0x3FFFFFFF | - |
| `DS_IV_0` | 0x6003D630 | 32 | IV_0=0xFFFFFFFF | - |
| `DS_IV_1` | 0x6003D634 | 32 | IV_0=0xFFFFFFFF | - |
| `DS_IV_2` | 0x6003D638 | 32 | IV_0=0xFFFFFFFF | - |
| `DS_IV_3` | 0x6003D63C | 32 | IV_0=0xFFFFFFFF | - |
| `DS_QUERY_BUSY` | 0x6003DE0C | 32 | QUERY_BUSY=0x1 | - |
| `DS_QUERY_CHECK` | 0x6003DE14 | 32 | MD_ERROR=0x1 PADDING_BAD=0x2 | - |
| `DS_QUERY_KEY_WRONG` | 0x6003DE10 | 32 | QUERY_KEY_WRONG=0xF | - |
| `DS_SET_FINISH` | 0x6003DE08 | 32 | SET_FINISH=0x1 | - |
| `DS_SET_ME` | 0x6003DE04 | 32 | SET_ME=0x1 | - |
| `DS_SET_START` | 0x6003DE00 | 32 | SET_START=0x1 | - |
| `DS_X_MEM[0]` | 0x6003D800 | 8 |  | - |
| `DS_X_MEM[100]` | 0x6003D864 | 8 |  | - |
| `DS_X_MEM[101]` | 0x6003D865 | 8 |  | - |
| `DS_X_MEM[102]` | 0x6003D866 | 8 |  | - |
| `DS_X_MEM[103]` | 0x6003D867 | 8 |  | - |
| `DS_X_MEM[104]` | 0x6003D868 | 8 |  | - |
| `DS_X_MEM[105]` | 0x6003D869 | 8 |  | - |
| `DS_X_MEM[106]` | 0x6003D86A | 8 |  | - |
| `DS_X_MEM[107]` | 0x6003D86B | 8 |  | - |
| `DS_X_MEM[108]` | 0x6003D86C | 8 |  | - |
| `DS_X_MEM[109]` | 0x6003D86D | 8 |  | - |
| `DS_X_MEM[10]` | 0x6003D80A | 8 |  | - |
| `DS_X_MEM[110]` | 0x6003D86E | 8 |  | - |
| `DS_X_MEM[111]` | 0x6003D86F | 8 |  | - |
| `DS_X_MEM[112]` | 0x6003D870 | 8 |  | - |
| `DS_X_MEM[113]` | 0x6003D871 | 8 |  | - |
| `DS_X_MEM[114]` | 0x6003D872 | 8 |  | - |
| `DS_X_MEM[115]` | 0x6003D873 | 8 |  | - |
| `DS_X_MEM[116]` | 0x6003D874 | 8 |  | - |
| `DS_X_MEM[117]` | 0x6003D875 | 8 |  | - |
| `DS_X_MEM[118]` | 0x6003D876 | 8 |  | - |
| `DS_X_MEM[119]` | 0x6003D877 | 8 |  | - |
| `DS_X_MEM[11]` | 0x6003D80B | 8 |  | - |
| `DS_X_MEM[120]` | 0x6003D878 | 8 |  | - |
| `DS_X_MEM[121]` | 0x6003D879 | 8 |  | - |
| `DS_X_MEM[122]` | 0x6003D87A | 8 |  | - |
| `DS_X_MEM[123]` | 0x6003D87B | 8 |  | - |
| `DS_X_MEM[124]` | 0x6003D87C | 8 |  | - |
| `DS_X_MEM[125]` | 0x6003D87D | 8 |  | - |
| `DS_X_MEM[126]` | 0x6003D87E | 8 |  | - |
| `DS_X_MEM[127]` | 0x6003D87F | 8 |  | - |
| `DS_X_MEM[128]` | 0x6003D880 | 8 |  | - |
| `DS_X_MEM[129]` | 0x6003D881 | 8 |  | - |
| `DS_X_MEM[12]` | 0x6003D80C | 8 |  | - |
| `DS_X_MEM[130]` | 0x6003D882 | 8 |  | - |
| `DS_X_MEM[131]` | 0x6003D883 | 8 |  | - |
| `DS_X_MEM[132]` | 0x6003D884 | 8 |  | - |
| `DS_X_MEM[133]` | 0x6003D885 | 8 |  | - |
| `DS_X_MEM[134]` | 0x6003D886 | 8 |  | - |
| `DS_X_MEM[135]` | 0x6003D887 | 8 |  | - |
| `DS_X_MEM[136]` | 0x6003D888 | 8 |  | - |
| `DS_X_MEM[137]` | 0x6003D889 | 8 |  | - |
| `DS_X_MEM[138]` | 0x6003D88A | 8 |  | - |
| `DS_X_MEM[139]` | 0x6003D88B | 8 |  | - |
| `DS_X_MEM[13]` | 0x6003D80D | 8 |  | - |
| `DS_X_MEM[140]` | 0x6003D88C | 8 |  | - |
| `DS_X_MEM[141]` | 0x6003D88D | 8 |  | - |
| `DS_X_MEM[142]` | 0x6003D88E | 8 |  | - |
| `DS_X_MEM[143]` | 0x6003D88F | 8 |  | - |
| `DS_X_MEM[144]` | 0x6003D890 | 8 |  | - |
| `DS_X_MEM[145]` | 0x6003D891 | 8 |  | - |
| `DS_X_MEM[146]` | 0x6003D892 | 8 |  | - |
| `DS_X_MEM[147]` | 0x6003D893 | 8 |  | - |
| `DS_X_MEM[148]` | 0x6003D894 | 8 |  | - |
| `DS_X_MEM[149]` | 0x6003D895 | 8 |  | - |
| `DS_X_MEM[14]` | 0x6003D80E | 8 |  | - |
| `DS_X_MEM[150]` | 0x6003D896 | 8 |  | - |
| `DS_X_MEM[151]` | 0x6003D897 | 8 |  | - |
| `DS_X_MEM[152]` | 0x6003D898 | 8 |  | - |
| `DS_X_MEM[153]` | 0x6003D899 | 8 |  | - |
| `DS_X_MEM[154]` | 0x6003D89A | 8 |  | - |
| `DS_X_MEM[155]` | 0x6003D89B | 8 |  | - |
| `DS_X_MEM[156]` | 0x6003D89C | 8 |  | - |
| `DS_X_MEM[157]` | 0x6003D89D | 8 |  | - |
| `DS_X_MEM[158]` | 0x6003D89E | 8 |  | - |
| `DS_X_MEM[159]` | 0x6003D89F | 8 |  | - |
| `DS_X_MEM[15]` | 0x6003D80F | 8 |  | - |
| `DS_X_MEM[160]` | 0x6003D8A0 | 8 |  | - |
| `DS_X_MEM[161]` | 0x6003D8A1 | 8 |  | - |
| `DS_X_MEM[162]` | 0x6003D8A2 | 8 |  | - |
| `DS_X_MEM[163]` | 0x6003D8A3 | 8 |  | - |
| `DS_X_MEM[164]` | 0x6003D8A4 | 8 |  | - |
| `DS_X_MEM[165]` | 0x6003D8A5 | 8 |  | - |
| `DS_X_MEM[166]` | 0x6003D8A6 | 8 |  | - |
| `DS_X_MEM[167]` | 0x6003D8A7 | 8 |  | - |
| `DS_X_MEM[168]` | 0x6003D8A8 | 8 |  | - |
| `DS_X_MEM[169]` | 0x6003D8A9 | 8 |  | - |
| `DS_X_MEM[16]` | 0x6003D810 | 8 |  | - |
| `DS_X_MEM[170]` | 0x6003D8AA | 8 |  | - |
| `DS_X_MEM[171]` | 0x6003D8AB | 8 |  | - |
| `DS_X_MEM[172]` | 0x6003D8AC | 8 |  | - |
| `DS_X_MEM[173]` | 0x6003D8AD | 8 |  | - |
| `DS_X_MEM[174]` | 0x6003D8AE | 8 |  | - |
| `DS_X_MEM[175]` | 0x6003D8AF | 8 |  | - |
| `DS_X_MEM[176]` | 0x6003D8B0 | 8 |  | - |
| `DS_X_MEM[177]` | 0x6003D8B1 | 8 |  | - |
| `DS_X_MEM[178]` | 0x6003D8B2 | 8 |  | - |
| `DS_X_MEM[179]` | 0x6003D8B3 | 8 |  | - |
| `DS_X_MEM[17]` | 0x6003D811 | 8 |  | - |
| `DS_X_MEM[180]` | 0x6003D8B4 | 8 |  | - |
| `DS_X_MEM[181]` | 0x6003D8B5 | 8 |  | - |
| `DS_X_MEM[182]` | 0x6003D8B6 | 8 |  | - |
| `DS_X_MEM[183]` | 0x6003D8B7 | 8 |  | - |
| `DS_X_MEM[184]` | 0x6003D8B8 | 8 |  | - |
| `DS_X_MEM[185]` | 0x6003D8B9 | 8 |  | - |
| `DS_X_MEM[186]` | 0x6003D8BA | 8 |  | - |
| `DS_X_MEM[187]` | 0x6003D8BB | 8 |  | - |
| `DS_X_MEM[188]` | 0x6003D8BC | 8 |  | - |
| `DS_X_MEM[189]` | 0x6003D8BD | 8 |  | - |
| `DS_X_MEM[18]` | 0x6003D812 | 8 |  | - |
| `DS_X_MEM[190]` | 0x6003D8BE | 8 |  | - |
| `DS_X_MEM[191]` | 0x6003D8BF | 8 |  | - |
| `DS_X_MEM[192]` | 0x6003D8C0 | 8 |  | - |
| `DS_X_MEM[193]` | 0x6003D8C1 | 8 |  | - |
| `DS_X_MEM[194]` | 0x6003D8C2 | 8 |  | - |
| `DS_X_MEM[195]` | 0x6003D8C3 | 8 |  | - |
| `DS_X_MEM[196]` | 0x6003D8C4 | 8 |  | - |
| `DS_X_MEM[197]` | 0x6003D8C5 | 8 |  | - |
| `DS_X_MEM[198]` | 0x6003D8C6 | 8 |  | - |
| `DS_X_MEM[199]` | 0x6003D8C7 | 8 |  | - |
| `DS_X_MEM[19]` | 0x6003D813 | 8 |  | - |
| `DS_X_MEM[1]` | 0x6003D801 | 8 |  | - |
| `DS_X_MEM[200]` | 0x6003D8C8 | 8 |  | - |
| `DS_X_MEM[201]` | 0x6003D8C9 | 8 |  | - |
| `DS_X_MEM[202]` | 0x6003D8CA | 8 |  | - |
| `DS_X_MEM[203]` | 0x6003D8CB | 8 |  | - |
| `DS_X_MEM[204]` | 0x6003D8CC | 8 |  | - |
| `DS_X_MEM[205]` | 0x6003D8CD | 8 |  | - |
| `DS_X_MEM[206]` | 0x6003D8CE | 8 |  | - |
| `DS_X_MEM[207]` | 0x6003D8CF | 8 |  | - |
| `DS_X_MEM[208]` | 0x6003D8D0 | 8 |  | - |
| `DS_X_MEM[209]` | 0x6003D8D1 | 8 |  | - |
| `DS_X_MEM[20]` | 0x6003D814 | 8 |  | - |
| `DS_X_MEM[210]` | 0x6003D8D2 | 8 |  | - |
| `DS_X_MEM[211]` | 0x6003D8D3 | 8 |  | - |
| `DS_X_MEM[212]` | 0x6003D8D4 | 8 |  | - |
| `DS_X_MEM[213]` | 0x6003D8D5 | 8 |  | - |
| `DS_X_MEM[214]` | 0x6003D8D6 | 8 |  | - |
| `DS_X_MEM[215]` | 0x6003D8D7 | 8 |  | - |
| `DS_X_MEM[216]` | 0x6003D8D8 | 8 |  | - |
| `DS_X_MEM[217]` | 0x6003D8D9 | 8 |  | - |
| `DS_X_MEM[218]` | 0x6003D8DA | 8 |  | - |
| `DS_X_MEM[219]` | 0x6003D8DB | 8 |  | - |
| `DS_X_MEM[21]` | 0x6003D815 | 8 |  | - |
| `DS_X_MEM[220]` | 0x6003D8DC | 8 |  | - |
| `DS_X_MEM[221]` | 0x6003D8DD | 8 |  | - |
| `DS_X_MEM[222]` | 0x6003D8DE | 8 |  | - |
| `DS_X_MEM[223]` | 0x6003D8DF | 8 |  | - |
| `DS_X_MEM[224]` | 0x6003D8E0 | 8 |  | - |
| `DS_X_MEM[225]` | 0x6003D8E1 | 8 |  | - |
| `DS_X_MEM[226]` | 0x6003D8E2 | 8 |  | - |
| `DS_X_MEM[227]` | 0x6003D8E3 | 8 |  | - |
| `DS_X_MEM[228]` | 0x6003D8E4 | 8 |  | - |
| `DS_X_MEM[229]` | 0x6003D8E5 | 8 |  | - |
| `DS_X_MEM[22]` | 0x6003D816 | 8 |  | - |
| `DS_X_MEM[230]` | 0x6003D8E6 | 8 |  | - |
| `DS_X_MEM[231]` | 0x6003D8E7 | 8 |  | - |
| `DS_X_MEM[232]` | 0x6003D8E8 | 8 |  | - |
| `DS_X_MEM[233]` | 0x6003D8E9 | 8 |  | - |
| `DS_X_MEM[234]` | 0x6003D8EA | 8 |  | - |
| `DS_X_MEM[235]` | 0x6003D8EB | 8 |  | - |
| `DS_X_MEM[236]` | 0x6003D8EC | 8 |  | - |
| `DS_X_MEM[237]` | 0x6003D8ED | 8 |  | - |
| `DS_X_MEM[238]` | 0x6003D8EE | 8 |  | - |
| `DS_X_MEM[239]` | 0x6003D8EF | 8 |  | - |
| `DS_X_MEM[23]` | 0x6003D817 | 8 |  | - |
| `DS_X_MEM[240]` | 0x6003D8F0 | 8 |  | - |
| `DS_X_MEM[241]` | 0x6003D8F1 | 8 |  | - |
| `DS_X_MEM[242]` | 0x6003D8F2 | 8 |  | - |
| `DS_X_MEM[243]` | 0x6003D8F3 | 8 |  | - |
| `DS_X_MEM[244]` | 0x6003D8F4 | 8 |  | - |
| `DS_X_MEM[245]` | 0x6003D8F5 | 8 |  | - |
| `DS_X_MEM[246]` | 0x6003D8F6 | 8 |  | - |
| `DS_X_MEM[247]` | 0x6003D8F7 | 8 |  | - |
| `DS_X_MEM[248]` | 0x6003D8F8 | 8 |  | - |
| `DS_X_MEM[249]` | 0x6003D8F9 | 8 |  | - |
| `DS_X_MEM[24]` | 0x6003D818 | 8 |  | - |
| `DS_X_MEM[250]` | 0x6003D8FA | 8 |  | - |
| `DS_X_MEM[251]` | 0x6003D8FB | 8 |  | - |
| `DS_X_MEM[252]` | 0x6003D8FC | 8 |  | - |
| `DS_X_MEM[253]` | 0x6003D8FD | 8 |  | - |
| `DS_X_MEM[254]` | 0x6003D8FE | 8 |  | - |
| `DS_X_MEM[255]` | 0x6003D8FF | 8 |  | - |
| `DS_X_MEM[256]` | 0x6003D900 | 8 |  | - |
| `DS_X_MEM[257]` | 0x6003D901 | 8 |  | - |
| `DS_X_MEM[258]` | 0x6003D902 | 8 |  | - |
| `DS_X_MEM[259]` | 0x6003D903 | 8 |  | - |
| `DS_X_MEM[25]` | 0x6003D819 | 8 |  | - |
| `DS_X_MEM[260]` | 0x6003D904 | 8 |  | - |
| `DS_X_MEM[261]` | 0x6003D905 | 8 |  | - |
| `DS_X_MEM[262]` | 0x6003D906 | 8 |  | - |
| `DS_X_MEM[263]` | 0x6003D907 | 8 |  | - |
| `DS_X_MEM[264]` | 0x6003D908 | 8 |  | - |
| `DS_X_MEM[265]` | 0x6003D909 | 8 |  | - |
| `DS_X_MEM[266]` | 0x6003D90A | 8 |  | - |
| `DS_X_MEM[267]` | 0x6003D90B | 8 |  | - |
| `DS_X_MEM[268]` | 0x6003D90C | 8 |  | - |
| `DS_X_MEM[269]` | 0x6003D90D | 8 |  | - |
| `DS_X_MEM[26]` | 0x6003D81A | 8 |  | - |
| `DS_X_MEM[270]` | 0x6003D90E | 8 |  | - |
| `DS_X_MEM[271]` | 0x6003D90F | 8 |  | - |
| `DS_X_MEM[272]` | 0x6003D910 | 8 |  | - |
| `DS_X_MEM[273]` | 0x6003D911 | 8 |  | - |
| `DS_X_MEM[274]` | 0x6003D912 | 8 |  | - |
| `DS_X_MEM[275]` | 0x6003D913 | 8 |  | - |
| `DS_X_MEM[276]` | 0x6003D914 | 8 |  | - |
| `DS_X_MEM[277]` | 0x6003D915 | 8 |  | - |
| `DS_X_MEM[278]` | 0x6003D916 | 8 |  | - |
| `DS_X_MEM[279]` | 0x6003D917 | 8 |  | - |
| `DS_X_MEM[27]` | 0x6003D81B | 8 |  | - |
| `DS_X_MEM[280]` | 0x6003D918 | 8 |  | - |
| `DS_X_MEM[281]` | 0x6003D919 | 8 |  | - |
| `DS_X_MEM[282]` | 0x6003D91A | 8 |  | - |
| `DS_X_MEM[283]` | 0x6003D91B | 8 |  | - |
| `DS_X_MEM[284]` | 0x6003D91C | 8 |  | - |
| `DS_X_MEM[285]` | 0x6003D91D | 8 |  | - |
| `DS_X_MEM[286]` | 0x6003D91E | 8 |  | - |
| `DS_X_MEM[287]` | 0x6003D91F | 8 |  | - |
| `DS_X_MEM[288]` | 0x6003D920 | 8 |  | - |
| `DS_X_MEM[289]` | 0x6003D921 | 8 |  | - |
| `DS_X_MEM[28]` | 0x6003D81C | 8 |  | - |
| `DS_X_MEM[290]` | 0x6003D922 | 8 |  | - |
| `DS_X_MEM[291]` | 0x6003D923 | 8 |  | - |
| `DS_X_MEM[292]` | 0x6003D924 | 8 |  | - |
| `DS_X_MEM[293]` | 0x6003D925 | 8 |  | - |
| `DS_X_MEM[294]` | 0x6003D926 | 8 |  | - |
| `DS_X_MEM[295]` | 0x6003D927 | 8 |  | - |
| `DS_X_MEM[296]` | 0x6003D928 | 8 |  | - |
| `DS_X_MEM[297]` | 0x6003D929 | 8 |  | - |
| `DS_X_MEM[298]` | 0x6003D92A | 8 |  | - |
| `DS_X_MEM[299]` | 0x6003D92B | 8 |  | - |
| `DS_X_MEM[29]` | 0x6003D81D | 8 |  | - |
| `DS_X_MEM[2]` | 0x6003D802 | 8 |  | - |
| `DS_X_MEM[300]` | 0x6003D92C | 8 |  | - |
| `DS_X_MEM[301]` | 0x6003D92D | 8 |  | - |
| `DS_X_MEM[302]` | 0x6003D92E | 8 |  | - |
| `DS_X_MEM[303]` | 0x6003D92F | 8 |  | - |
| `DS_X_MEM[304]` | 0x6003D930 | 8 |  | - |
| `DS_X_MEM[305]` | 0x6003D931 | 8 |  | - |
| `DS_X_MEM[306]` | 0x6003D932 | 8 |  | - |
| `DS_X_MEM[307]` | 0x6003D933 | 8 |  | - |
| `DS_X_MEM[308]` | 0x6003D934 | 8 |  | - |
| `DS_X_MEM[309]` | 0x6003D935 | 8 |  | - |
| `DS_X_MEM[30]` | 0x6003D81E | 8 |  | - |
| `DS_X_MEM[310]` | 0x6003D936 | 8 |  | - |
| `DS_X_MEM[311]` | 0x6003D937 | 8 |  | - |
| `DS_X_MEM[312]` | 0x6003D938 | 8 |  | - |
| `DS_X_MEM[313]` | 0x6003D939 | 8 |  | - |
| `DS_X_MEM[314]` | 0x6003D93A | 8 |  | - |
| `DS_X_MEM[315]` | 0x6003D93B | 8 |  | - |
| `DS_X_MEM[316]` | 0x6003D93C | 8 |  | - |
| `DS_X_MEM[317]` | 0x6003D93D | 8 |  | - |
| `DS_X_MEM[318]` | 0x6003D93E | 8 |  | - |
| `DS_X_MEM[319]` | 0x6003D93F | 8 |  | - |
| `DS_X_MEM[31]` | 0x6003D81F | 8 |  | - |
| `DS_X_MEM[320]` | 0x6003D940 | 8 |  | - |
| `DS_X_MEM[321]` | 0x6003D941 | 8 |  | - |
| `DS_X_MEM[322]` | 0x6003D942 | 8 |  | - |
| `DS_X_MEM[323]` | 0x6003D943 | 8 |  | - |
| `DS_X_MEM[324]` | 0x6003D944 | 8 |  | - |
| `DS_X_MEM[325]` | 0x6003D945 | 8 |  | - |
| `DS_X_MEM[326]` | 0x6003D946 | 8 |  | - |
| `DS_X_MEM[327]` | 0x6003D947 | 8 |  | - |
| `DS_X_MEM[328]` | 0x6003D948 | 8 |  | - |
| `DS_X_MEM[329]` | 0x6003D949 | 8 |  | - |
| `DS_X_MEM[32]` | 0x6003D820 | 8 |  | - |
| `DS_X_MEM[330]` | 0x6003D94A | 8 |  | - |
| `DS_X_MEM[331]` | 0x6003D94B | 8 |  | - |
| `DS_X_MEM[332]` | 0x6003D94C | 8 |  | - |
| `DS_X_MEM[333]` | 0x6003D94D | 8 |  | - |
| `DS_X_MEM[334]` | 0x6003D94E | 8 |  | - |
| `DS_X_MEM[335]` | 0x6003D94F | 8 |  | - |
| `DS_X_MEM[336]` | 0x6003D950 | 8 |  | - |
| `DS_X_MEM[337]` | 0x6003D951 | 8 |  | - |
| `DS_X_MEM[338]` | 0x6003D952 | 8 |  | - |
| `DS_X_MEM[339]` | 0x6003D953 | 8 |  | - |
| `DS_X_MEM[33]` | 0x6003D821 | 8 |  | - |
| `DS_X_MEM[340]` | 0x6003D954 | 8 |  | - |
| `DS_X_MEM[341]` | 0x6003D955 | 8 |  | - |
| `DS_X_MEM[342]` | 0x6003D956 | 8 |  | - |
| `DS_X_MEM[343]` | 0x6003D957 | 8 |  | - |
| `DS_X_MEM[344]` | 0x6003D958 | 8 |  | - |
| `DS_X_MEM[345]` | 0x6003D959 | 8 |  | - |
| `DS_X_MEM[346]` | 0x6003D95A | 8 |  | - |
| `DS_X_MEM[347]` | 0x6003D95B | 8 |  | - |
| `DS_X_MEM[348]` | 0x6003D95C | 8 |  | - |
| `DS_X_MEM[349]` | 0x6003D95D | 8 |  | - |
| `DS_X_MEM[34]` | 0x6003D822 | 8 |  | - |
| `DS_X_MEM[350]` | 0x6003D95E | 8 |  | - |
| `DS_X_MEM[351]` | 0x6003D95F | 8 |  | - |
| `DS_X_MEM[352]` | 0x6003D960 | 8 |  | - |
| `DS_X_MEM[353]` | 0x6003D961 | 8 |  | - |
| `DS_X_MEM[354]` | 0x6003D962 | 8 |  | - |
| `DS_X_MEM[355]` | 0x6003D963 | 8 |  | - |
| `DS_X_MEM[356]` | 0x6003D964 | 8 |  | - |
| `DS_X_MEM[357]` | 0x6003D965 | 8 |  | - |
| `DS_X_MEM[358]` | 0x6003D966 | 8 |  | - |
| `DS_X_MEM[359]` | 0x6003D967 | 8 |  | - |
| `DS_X_MEM[35]` | 0x6003D823 | 8 |  | - |
| `DS_X_MEM[360]` | 0x6003D968 | 8 |  | - |
| `DS_X_MEM[361]` | 0x6003D969 | 8 |  | - |
| `DS_X_MEM[362]` | 0x6003D96A | 8 |  | - |
| `DS_X_MEM[363]` | 0x6003D96B | 8 |  | - |
| `DS_X_MEM[364]` | 0x6003D96C | 8 |  | - |
| `DS_X_MEM[365]` | 0x6003D96D | 8 |  | - |
| `DS_X_MEM[366]` | 0x6003D96E | 8 |  | - |
| `DS_X_MEM[367]` | 0x6003D96F | 8 |  | - |
| `DS_X_MEM[368]` | 0x6003D970 | 8 |  | - |
| `DS_X_MEM[369]` | 0x6003D971 | 8 |  | - |
| `DS_X_MEM[36]` | 0x6003D824 | 8 |  | - |
| `DS_X_MEM[370]` | 0x6003D972 | 8 |  | - |
| `DS_X_MEM[371]` | 0x6003D973 | 8 |  | - |
| `DS_X_MEM[372]` | 0x6003D974 | 8 |  | - |
| `DS_X_MEM[373]` | 0x6003D975 | 8 |  | - |
| `DS_X_MEM[374]` | 0x6003D976 | 8 |  | - |
| `DS_X_MEM[375]` | 0x6003D977 | 8 |  | - |
| `DS_X_MEM[376]` | 0x6003D978 | 8 |  | - |
| `DS_X_MEM[377]` | 0x6003D979 | 8 |  | - |
| `DS_X_MEM[378]` | 0x6003D97A | 8 |  | - |
| `DS_X_MEM[379]` | 0x6003D97B | 8 |  | - |
| `DS_X_MEM[37]` | 0x6003D825 | 8 |  | - |
| `DS_X_MEM[380]` | 0x6003D97C | 8 |  | - |
| `DS_X_MEM[381]` | 0x6003D97D | 8 |  | - |
| `DS_X_MEM[382]` | 0x6003D97E | 8 |  | - |
| `DS_X_MEM[383]` | 0x6003D97F | 8 |  | - |
| `DS_X_MEM[384]` | 0x6003D980 | 8 |  | - |
| `DS_X_MEM[385]` | 0x6003D981 | 8 |  | - |
| `DS_X_MEM[386]` | 0x6003D982 | 8 |  | - |
| `DS_X_MEM[387]` | 0x6003D983 | 8 |  | - |
| `DS_X_MEM[388]` | 0x6003D984 | 8 |  | - |
| `DS_X_MEM[389]` | 0x6003D985 | 8 |  | - |
| `DS_X_MEM[38]` | 0x6003D826 | 8 |  | - |
| `DS_X_MEM[390]` | 0x6003D986 | 8 |  | - |
| `DS_X_MEM[391]` | 0x6003D987 | 8 |  | - |
| `DS_X_MEM[392]` | 0x6003D988 | 8 |  | - |
| `DS_X_MEM[393]` | 0x6003D989 | 8 |  | - |
| `DS_X_MEM[394]` | 0x6003D98A | 8 |  | - |
| `DS_X_MEM[395]` | 0x6003D98B | 8 |  | - |
| `DS_X_MEM[396]` | 0x6003D98C | 8 |  | - |
| `DS_X_MEM[397]` | 0x6003D98D | 8 |  | - |
| `DS_X_MEM[398]` | 0x6003D98E | 8 |  | - |
| `DS_X_MEM[399]` | 0x6003D98F | 8 |  | - |
| `DS_X_MEM[39]` | 0x6003D827 | 8 |  | - |
| `DS_X_MEM[3]` | 0x6003D803 | 8 |  | - |
| `DS_X_MEM[400]` | 0x6003D990 | 8 |  | - |
| `DS_X_MEM[401]` | 0x6003D991 | 8 |  | - |
| `DS_X_MEM[402]` | 0x6003D992 | 8 |  | - |
| `DS_X_MEM[403]` | 0x6003D993 | 8 |  | - |
| `DS_X_MEM[404]` | 0x6003D994 | 8 |  | - |
| `DS_X_MEM[405]` | 0x6003D995 | 8 |  | - |
| `DS_X_MEM[406]` | 0x6003D996 | 8 |  | - |
| `DS_X_MEM[407]` | 0x6003D997 | 8 |  | - |
| `DS_X_MEM[408]` | 0x6003D998 | 8 |  | - |
| `DS_X_MEM[409]` | 0x6003D999 | 8 |  | - |
| `DS_X_MEM[40]` | 0x6003D828 | 8 |  | - |
| `DS_X_MEM[410]` | 0x6003D99A | 8 |  | - |
| `DS_X_MEM[411]` | 0x6003D99B | 8 |  | - |
| `DS_X_MEM[412]` | 0x6003D99C | 8 |  | - |
| `DS_X_MEM[413]` | 0x6003D99D | 8 |  | - |
| `DS_X_MEM[414]` | 0x6003D99E | 8 |  | - |
| `DS_X_MEM[415]` | 0x6003D99F | 8 |  | - |
| `DS_X_MEM[416]` | 0x6003D9A0 | 8 |  | - |
| `DS_X_MEM[417]` | 0x6003D9A1 | 8 |  | - |
| `DS_X_MEM[418]` | 0x6003D9A2 | 8 |  | - |
| `DS_X_MEM[419]` | 0x6003D9A3 | 8 |  | - |
| `DS_X_MEM[41]` | 0x6003D829 | 8 |  | - |
| `DS_X_MEM[420]` | 0x6003D9A4 | 8 |  | - |
| `DS_X_MEM[421]` | 0x6003D9A5 | 8 |  | - |
| `DS_X_MEM[422]` | 0x6003D9A6 | 8 |  | - |
| `DS_X_MEM[423]` | 0x6003D9A7 | 8 |  | - |
| `DS_X_MEM[424]` | 0x6003D9A8 | 8 |  | - |
| `DS_X_MEM[425]` | 0x6003D9A9 | 8 |  | - |
| `DS_X_MEM[426]` | 0x6003D9AA | 8 |  | - |
| `DS_X_MEM[427]` | 0x6003D9AB | 8 |  | - |
| `DS_X_MEM[428]` | 0x6003D9AC | 8 |  | - |
| `DS_X_MEM[429]` | 0x6003D9AD | 8 |  | - |
| `DS_X_MEM[42]` | 0x6003D82A | 8 |  | - |
| `DS_X_MEM[430]` | 0x6003D9AE | 8 |  | - |
| `DS_X_MEM[431]` | 0x6003D9AF | 8 |  | - |
| `DS_X_MEM[432]` | 0x6003D9B0 | 8 |  | - |
| `DS_X_MEM[433]` | 0x6003D9B1 | 8 |  | - |
| `DS_X_MEM[434]` | 0x6003D9B2 | 8 |  | - |
| `DS_X_MEM[435]` | 0x6003D9B3 | 8 |  | - |
| `DS_X_MEM[436]` | 0x6003D9B4 | 8 |  | - |
| `DS_X_MEM[437]` | 0x6003D9B5 | 8 |  | - |
| `DS_X_MEM[438]` | 0x6003D9B6 | 8 |  | - |
| `DS_X_MEM[439]` | 0x6003D9B7 | 8 |  | - |
| `DS_X_MEM[43]` | 0x6003D82B | 8 |  | - |
| `DS_X_MEM[440]` | 0x6003D9B8 | 8 |  | - |
| `DS_X_MEM[441]` | 0x6003D9B9 | 8 |  | - |
| `DS_X_MEM[442]` | 0x6003D9BA | 8 |  | - |
| `DS_X_MEM[443]` | 0x6003D9BB | 8 |  | - |
| `DS_X_MEM[444]` | 0x6003D9BC | 8 |  | - |
| `DS_X_MEM[445]` | 0x6003D9BD | 8 |  | - |
| `DS_X_MEM[446]` | 0x6003D9BE | 8 |  | - |
| `DS_X_MEM[447]` | 0x6003D9BF | 8 |  | - |
| `DS_X_MEM[448]` | 0x6003D9C0 | 8 |  | - |
| `DS_X_MEM[449]` | 0x6003D9C1 | 8 |  | - |
| `DS_X_MEM[44]` | 0x6003D82C | 8 |  | - |
| `DS_X_MEM[450]` | 0x6003D9C2 | 8 |  | - |
| `DS_X_MEM[451]` | 0x6003D9C3 | 8 |  | - |
| `DS_X_MEM[452]` | 0x6003D9C4 | 8 |  | - |
| `DS_X_MEM[453]` | 0x6003D9C5 | 8 |  | - |
| `DS_X_MEM[454]` | 0x6003D9C6 | 8 |  | - |
| `DS_X_MEM[455]` | 0x6003D9C7 | 8 |  | - |
| `DS_X_MEM[456]` | 0x6003D9C8 | 8 |  | - |
| `DS_X_MEM[457]` | 0x6003D9C9 | 8 |  | - |
| `DS_X_MEM[458]` | 0x6003D9CA | 8 |  | - |
| `DS_X_MEM[459]` | 0x6003D9CB | 8 |  | - |
| `DS_X_MEM[45]` | 0x6003D82D | 8 |  | - |
| `DS_X_MEM[460]` | 0x6003D9CC | 8 |  | - |
| `DS_X_MEM[461]` | 0x6003D9CD | 8 |  | - |
| `DS_X_MEM[462]` | 0x6003D9CE | 8 |  | - |
| `DS_X_MEM[463]` | 0x6003D9CF | 8 |  | - |
| `DS_X_MEM[464]` | 0x6003D9D0 | 8 |  | - |
| `DS_X_MEM[465]` | 0x6003D9D1 | 8 |  | - |
| `DS_X_MEM[466]` | 0x6003D9D2 | 8 |  | - |
| `DS_X_MEM[467]` | 0x6003D9D3 | 8 |  | - |
| `DS_X_MEM[468]` | 0x6003D9D4 | 8 |  | - |
| `DS_X_MEM[469]` | 0x6003D9D5 | 8 |  | - |
| `DS_X_MEM[46]` | 0x6003D82E | 8 |  | - |
| `DS_X_MEM[470]` | 0x6003D9D6 | 8 |  | - |
| `DS_X_MEM[471]` | 0x6003D9D7 | 8 |  | - |
| `DS_X_MEM[472]` | 0x6003D9D8 | 8 |  | - |
| `DS_X_MEM[473]` | 0x6003D9D9 | 8 |  | - |
| `DS_X_MEM[474]` | 0x6003D9DA | 8 |  | - |
| `DS_X_MEM[475]` | 0x6003D9DB | 8 |  | - |
| `DS_X_MEM[476]` | 0x6003D9DC | 8 |  | - |
| `DS_X_MEM[477]` | 0x6003D9DD | 8 |  | - |
| `DS_X_MEM[478]` | 0x6003D9DE | 8 |  | - |
| `DS_X_MEM[479]` | 0x6003D9DF | 8 |  | - |
| `DS_X_MEM[47]` | 0x6003D82F | 8 |  | - |
| `DS_X_MEM[480]` | 0x6003D9E0 | 8 |  | - |
| `DS_X_MEM[481]` | 0x6003D9E1 | 8 |  | - |
| `DS_X_MEM[482]` | 0x6003D9E2 | 8 |  | - |
| `DS_X_MEM[483]` | 0x6003D9E3 | 8 |  | - |
| `DS_X_MEM[484]` | 0x6003D9E4 | 8 |  | - |
| `DS_X_MEM[485]` | 0x6003D9E5 | 8 |  | - |
| `DS_X_MEM[486]` | 0x6003D9E6 | 8 |  | - |
| `DS_X_MEM[487]` | 0x6003D9E7 | 8 |  | - |
| `DS_X_MEM[488]` | 0x6003D9E8 | 8 |  | - |
| `DS_X_MEM[489]` | 0x6003D9E9 | 8 |  | - |
| `DS_X_MEM[48]` | 0x6003D830 | 8 |  | - |
| `DS_X_MEM[490]` | 0x6003D9EA | 8 |  | - |
| `DS_X_MEM[491]` | 0x6003D9EB | 8 |  | - |
| `DS_X_MEM[492]` | 0x6003D9EC | 8 |  | - |
| `DS_X_MEM[493]` | 0x6003D9ED | 8 |  | - |
| `DS_X_MEM[494]` | 0x6003D9EE | 8 |  | - |
| `DS_X_MEM[495]` | 0x6003D9EF | 8 |  | - |
| `DS_X_MEM[496]` | 0x6003D9F0 | 8 |  | - |
| `DS_X_MEM[497]` | 0x6003D9F1 | 8 |  | - |
| `DS_X_MEM[498]` | 0x6003D9F2 | 8 |  | - |
| `DS_X_MEM[499]` | 0x6003D9F3 | 8 |  | - |
| `DS_X_MEM[49]` | 0x6003D831 | 8 |  | - |
| `DS_X_MEM[4]` | 0x6003D804 | 8 |  | - |
| `DS_X_MEM[500]` | 0x6003D9F4 | 8 |  | - |
| `DS_X_MEM[501]` | 0x6003D9F5 | 8 |  | - |
| `DS_X_MEM[502]` | 0x6003D9F6 | 8 |  | - |
| `DS_X_MEM[503]` | 0x6003D9F7 | 8 |  | - |
| `DS_X_MEM[504]` | 0x6003D9F8 | 8 |  | - |
| `DS_X_MEM[505]` | 0x6003D9F9 | 8 |  | - |
| `DS_X_MEM[506]` | 0x6003D9FA | 8 |  | - |
| `DS_X_MEM[507]` | 0x6003D9FB | 8 |  | - |
| `DS_X_MEM[508]` | 0x6003D9FC | 8 |  | - |
| `DS_X_MEM[509]` | 0x6003D9FD | 8 |  | - |
| `DS_X_MEM[50]` | 0x6003D832 | 8 |  | - |
| `DS_X_MEM[510]` | 0x6003D9FE | 8 |  | - |
| `DS_X_MEM[511]` | 0x6003D9FF | 8 |  | - |
| `DS_X_MEM[51]` | 0x6003D833 | 8 |  | - |
| `DS_X_MEM[52]` | 0x6003D834 | 8 |  | - |
| `DS_X_MEM[53]` | 0x6003D835 | 8 |  | - |
| `DS_X_MEM[54]` | 0x6003D836 | 8 |  | - |
| `DS_X_MEM[55]` | 0x6003D837 | 8 |  | - |
| `DS_X_MEM[56]` | 0x6003D838 | 8 |  | - |
| `DS_X_MEM[57]` | 0x6003D839 | 8 |  | - |
| `DS_X_MEM[58]` | 0x6003D83A | 8 |  | - |
| `DS_X_MEM[59]` | 0x6003D83B | 8 |  | - |
| `DS_X_MEM[5]` | 0x6003D805 | 8 |  | - |
| `DS_X_MEM[60]` | 0x6003D83C | 8 |  | - |
| `DS_X_MEM[61]` | 0x6003D83D | 8 |  | - |
| `DS_X_MEM[62]` | 0x6003D83E | 8 |  | - |
| `DS_X_MEM[63]` | 0x6003D83F | 8 |  | - |
| `DS_X_MEM[64]` | 0x6003D840 | 8 |  | - |
| `DS_X_MEM[65]` | 0x6003D841 | 8 |  | - |
| `DS_X_MEM[66]` | 0x6003D842 | 8 |  | - |
| `DS_X_MEM[67]` | 0x6003D843 | 8 |  | - |
| `DS_X_MEM[68]` | 0x6003D844 | 8 |  | - |
| `DS_X_MEM[69]` | 0x6003D845 | 8 |  | - |
| `DS_X_MEM[6]` | 0x6003D806 | 8 |  | - |
| `DS_X_MEM[70]` | 0x6003D846 | 8 |  | - |
| `DS_X_MEM[71]` | 0x6003D847 | 8 |  | - |
| `DS_X_MEM[72]` | 0x6003D848 | 8 |  | - |
| `DS_X_MEM[73]` | 0x6003D849 | 8 |  | - |
| `DS_X_MEM[74]` | 0x6003D84A | 8 |  | - |
| `DS_X_MEM[75]` | 0x6003D84B | 8 |  | - |
| `DS_X_MEM[76]` | 0x6003D84C | 8 |  | - |
| `DS_X_MEM[77]` | 0x6003D84D | 8 |  | - |
| `DS_X_MEM[78]` | 0x6003D84E | 8 |  | - |
| `DS_X_MEM[79]` | 0x6003D84F | 8 |  | - |
| `DS_X_MEM[7]` | 0x6003D807 | 8 |  | - |
| `DS_X_MEM[80]` | 0x6003D850 | 8 |  | - |
| `DS_X_MEM[81]` | 0x6003D851 | 8 |  | - |
| `DS_X_MEM[82]` | 0x6003D852 | 8 |  | - |
| `DS_X_MEM[83]` | 0x6003D853 | 8 |  | - |
| `DS_X_MEM[84]` | 0x6003D854 | 8 |  | - |
| `DS_X_MEM[85]` | 0x6003D855 | 8 |  | - |
| `DS_X_MEM[86]` | 0x6003D856 | 8 |  | - |
| `DS_X_MEM[87]` | 0x6003D857 | 8 |  | - |
| `DS_X_MEM[88]` | 0x6003D858 | 8 |  | - |
| `DS_X_MEM[89]` | 0x6003D859 | 8 |  | - |
| `DS_X_MEM[8]` | 0x6003D808 | 8 |  | - |
| `DS_X_MEM[90]` | 0x6003D85A | 8 |  | - |
| `DS_X_MEM[91]` | 0x6003D85B | 8 |  | - |
| `DS_X_MEM[92]` | 0x6003D85C | 8 |  | - |
| `DS_X_MEM[93]` | 0x6003D85D | 8 |  | - |
| `DS_X_MEM[94]` | 0x6003D85E | 8 |  | - |
| `DS_X_MEM[95]` | 0x6003D85F | 8 |  | - |
| `DS_X_MEM[96]` | 0x6003D860 | 8 |  | - |
| `DS_X_MEM[97]` | 0x6003D861 | 8 |  | - |
| `DS_X_MEM[98]` | 0x6003D862 | 8 |  | - |
| `DS_X_MEM[99]` | 0x6003D863 | 8 |  | - |
| `DS_X_MEM[9]` | 0x6003D809 | 8 |  | - |
| `DS_Z_MEM[0]` | 0x6003DA00 | 8 |  | - |
| `DS_Z_MEM[100]` | 0x6003DA64 | 8 |  | - |
| `DS_Z_MEM[101]` | 0x6003DA65 | 8 |  | - |
| `DS_Z_MEM[102]` | 0x6003DA66 | 8 |  | - |
| `DS_Z_MEM[103]` | 0x6003DA67 | 8 |  | - |
| `DS_Z_MEM[104]` | 0x6003DA68 | 8 |  | - |
| `DS_Z_MEM[105]` | 0x6003DA69 | 8 |  | - |
| `DS_Z_MEM[106]` | 0x6003DA6A | 8 |  | - |
| `DS_Z_MEM[107]` | 0x6003DA6B | 8 |  | - |
| `DS_Z_MEM[108]` | 0x6003DA6C | 8 |  | - |
| `DS_Z_MEM[109]` | 0x6003DA6D | 8 |  | - |
| `DS_Z_MEM[10]` | 0x6003DA0A | 8 |  | - |
| `DS_Z_MEM[110]` | 0x6003DA6E | 8 |  | - |
| `DS_Z_MEM[111]` | 0x6003DA6F | 8 |  | - |
| `DS_Z_MEM[112]` | 0x6003DA70 | 8 |  | - |
| `DS_Z_MEM[113]` | 0x6003DA71 | 8 |  | - |
| `DS_Z_MEM[114]` | 0x6003DA72 | 8 |  | - |
| `DS_Z_MEM[115]` | 0x6003DA73 | 8 |  | - |
| `DS_Z_MEM[116]` | 0x6003DA74 | 8 |  | - |
| `DS_Z_MEM[117]` | 0x6003DA75 | 8 |  | - |
| `DS_Z_MEM[118]` | 0x6003DA76 | 8 |  | - |
| `DS_Z_MEM[119]` | 0x6003DA77 | 8 |  | - |
| `DS_Z_MEM[11]` | 0x6003DA0B | 8 |  | - |
| `DS_Z_MEM[120]` | 0x6003DA78 | 8 |  | - |
| `DS_Z_MEM[121]` | 0x6003DA79 | 8 |  | - |
| `DS_Z_MEM[122]` | 0x6003DA7A | 8 |  | - |
| `DS_Z_MEM[123]` | 0x6003DA7B | 8 |  | - |
| `DS_Z_MEM[124]` | 0x6003DA7C | 8 |  | - |
| `DS_Z_MEM[125]` | 0x6003DA7D | 8 |  | - |
| `DS_Z_MEM[126]` | 0x6003DA7E | 8 |  | - |
| `DS_Z_MEM[127]` | 0x6003DA7F | 8 |  | - |
| `DS_Z_MEM[128]` | 0x6003DA80 | 8 |  | - |
| `DS_Z_MEM[129]` | 0x6003DA81 | 8 |  | - |
| `DS_Z_MEM[12]` | 0x6003DA0C | 8 |  | - |
| `DS_Z_MEM[130]` | 0x6003DA82 | 8 |  | - |
| `DS_Z_MEM[131]` | 0x6003DA83 | 8 |  | - |
| `DS_Z_MEM[132]` | 0x6003DA84 | 8 |  | - |
| `DS_Z_MEM[133]` | 0x6003DA85 | 8 |  | - |
| `DS_Z_MEM[134]` | 0x6003DA86 | 8 |  | - |
| `DS_Z_MEM[135]` | 0x6003DA87 | 8 |  | - |
| `DS_Z_MEM[136]` | 0x6003DA88 | 8 |  | - |
| `DS_Z_MEM[137]` | 0x6003DA89 | 8 |  | - |
| `DS_Z_MEM[138]` | 0x6003DA8A | 8 |  | - |
| `DS_Z_MEM[139]` | 0x6003DA8B | 8 |  | - |
| `DS_Z_MEM[13]` | 0x6003DA0D | 8 |  | - |
| `DS_Z_MEM[140]` | 0x6003DA8C | 8 |  | - |
| `DS_Z_MEM[141]` | 0x6003DA8D | 8 |  | - |
| `DS_Z_MEM[142]` | 0x6003DA8E | 8 |  | - |
| `DS_Z_MEM[143]` | 0x6003DA8F | 8 |  | - |
| `DS_Z_MEM[144]` | 0x6003DA90 | 8 |  | - |
| `DS_Z_MEM[145]` | 0x6003DA91 | 8 |  | - |
| `DS_Z_MEM[146]` | 0x6003DA92 | 8 |  | - |
| `DS_Z_MEM[147]` | 0x6003DA93 | 8 |  | - |
| `DS_Z_MEM[148]` | 0x6003DA94 | 8 |  | - |
| `DS_Z_MEM[149]` | 0x6003DA95 | 8 |  | - |
| `DS_Z_MEM[14]` | 0x6003DA0E | 8 |  | - |
| `DS_Z_MEM[150]` | 0x6003DA96 | 8 |  | - |
| `DS_Z_MEM[151]` | 0x6003DA97 | 8 |  | - |
| `DS_Z_MEM[152]` | 0x6003DA98 | 8 |  | - |
| `DS_Z_MEM[153]` | 0x6003DA99 | 8 |  | - |
| `DS_Z_MEM[154]` | 0x6003DA9A | 8 |  | - |
| `DS_Z_MEM[155]` | 0x6003DA9B | 8 |  | - |
| `DS_Z_MEM[156]` | 0x6003DA9C | 8 |  | - |
| `DS_Z_MEM[157]` | 0x6003DA9D | 8 |  | - |
| `DS_Z_MEM[158]` | 0x6003DA9E | 8 |  | - |
| `DS_Z_MEM[159]` | 0x6003DA9F | 8 |  | - |
| `DS_Z_MEM[15]` | 0x6003DA0F | 8 |  | - |
| `DS_Z_MEM[160]` | 0x6003DAA0 | 8 |  | - |
| `DS_Z_MEM[161]` | 0x6003DAA1 | 8 |  | - |
| `DS_Z_MEM[162]` | 0x6003DAA2 | 8 |  | - |
| `DS_Z_MEM[163]` | 0x6003DAA3 | 8 |  | - |
| `DS_Z_MEM[164]` | 0x6003DAA4 | 8 |  | - |
| `DS_Z_MEM[165]` | 0x6003DAA5 | 8 |  | - |
| `DS_Z_MEM[166]` | 0x6003DAA6 | 8 |  | - |
| `DS_Z_MEM[167]` | 0x6003DAA7 | 8 |  | - |
| `DS_Z_MEM[168]` | 0x6003DAA8 | 8 |  | - |
| `DS_Z_MEM[169]` | 0x6003DAA9 | 8 |  | - |
| `DS_Z_MEM[16]` | 0x6003DA10 | 8 |  | - |
| `DS_Z_MEM[170]` | 0x6003DAAA | 8 |  | - |
| `DS_Z_MEM[171]` | 0x6003DAAB | 8 |  | - |
| `DS_Z_MEM[172]` | 0x6003DAAC | 8 |  | - |
| `DS_Z_MEM[173]` | 0x6003DAAD | 8 |  | - |
| `DS_Z_MEM[174]` | 0x6003DAAE | 8 |  | - |
| `DS_Z_MEM[175]` | 0x6003DAAF | 8 |  | - |
| `DS_Z_MEM[176]` | 0x6003DAB0 | 8 |  | - |
| `DS_Z_MEM[177]` | 0x6003DAB1 | 8 |  | - |
| `DS_Z_MEM[178]` | 0x6003DAB2 | 8 |  | - |
| `DS_Z_MEM[179]` | 0x6003DAB3 | 8 |  | - |
| `DS_Z_MEM[17]` | 0x6003DA11 | 8 |  | - |
| `DS_Z_MEM[180]` | 0x6003DAB4 | 8 |  | - |
| `DS_Z_MEM[181]` | 0x6003DAB5 | 8 |  | - |
| `DS_Z_MEM[182]` | 0x6003DAB6 | 8 |  | - |
| `DS_Z_MEM[183]` | 0x6003DAB7 | 8 |  | - |
| `DS_Z_MEM[184]` | 0x6003DAB8 | 8 |  | - |
| `DS_Z_MEM[185]` | 0x6003DAB9 | 8 |  | - |
| `DS_Z_MEM[186]` | 0x6003DABA | 8 |  | - |
| `DS_Z_MEM[187]` | 0x6003DABB | 8 |  | - |
| `DS_Z_MEM[188]` | 0x6003DABC | 8 |  | - |
| `DS_Z_MEM[189]` | 0x6003DABD | 8 |  | - |
| `DS_Z_MEM[18]` | 0x6003DA12 | 8 |  | - |
| `DS_Z_MEM[190]` | 0x6003DABE | 8 |  | - |
| `DS_Z_MEM[191]` | 0x6003DABF | 8 |  | - |
| `DS_Z_MEM[192]` | 0x6003DAC0 | 8 |  | - |
| `DS_Z_MEM[193]` | 0x6003DAC1 | 8 |  | - |
| `DS_Z_MEM[194]` | 0x6003DAC2 | 8 |  | - |
| `DS_Z_MEM[195]` | 0x6003DAC3 | 8 |  | - |
| `DS_Z_MEM[196]` | 0x6003DAC4 | 8 |  | - |
| `DS_Z_MEM[197]` | 0x6003DAC5 | 8 |  | - |
| `DS_Z_MEM[198]` | 0x6003DAC6 | 8 |  | - |
| `DS_Z_MEM[199]` | 0x6003DAC7 | 8 |  | - |
| `DS_Z_MEM[19]` | 0x6003DA13 | 8 |  | - |
| `DS_Z_MEM[1]` | 0x6003DA01 | 8 |  | - |
| `DS_Z_MEM[200]` | 0x6003DAC8 | 8 |  | - |
| `DS_Z_MEM[201]` | 0x6003DAC9 | 8 |  | - |
| `DS_Z_MEM[202]` | 0x6003DACA | 8 |  | - |
| `DS_Z_MEM[203]` | 0x6003DACB | 8 |  | - |
| `DS_Z_MEM[204]` | 0x6003DACC | 8 |  | - |
| `DS_Z_MEM[205]` | 0x6003DACD | 8 |  | - |
| `DS_Z_MEM[206]` | 0x6003DACE | 8 |  | - |
| `DS_Z_MEM[207]` | 0x6003DACF | 8 |  | - |
| `DS_Z_MEM[208]` | 0x6003DAD0 | 8 |  | - |
| `DS_Z_MEM[209]` | 0x6003DAD1 | 8 |  | - |
| `DS_Z_MEM[20]` | 0x6003DA14 | 8 |  | - |
| `DS_Z_MEM[210]` | 0x6003DAD2 | 8 |  | - |
| `DS_Z_MEM[211]` | 0x6003DAD3 | 8 |  | - |
| `DS_Z_MEM[212]` | 0x6003DAD4 | 8 |  | - |
| `DS_Z_MEM[213]` | 0x6003DAD5 | 8 |  | - |
| `DS_Z_MEM[214]` | 0x6003DAD6 | 8 |  | - |
| `DS_Z_MEM[215]` | 0x6003DAD7 | 8 |  | - |
| `DS_Z_MEM[216]` | 0x6003DAD8 | 8 |  | - |
| `DS_Z_MEM[217]` | 0x6003DAD9 | 8 |  | - |
| `DS_Z_MEM[218]` | 0x6003DADA | 8 |  | - |
| `DS_Z_MEM[219]` | 0x6003DADB | 8 |  | - |
| `DS_Z_MEM[21]` | 0x6003DA15 | 8 |  | - |
| `DS_Z_MEM[220]` | 0x6003DADC | 8 |  | - |
| `DS_Z_MEM[221]` | 0x6003DADD | 8 |  | - |
| `DS_Z_MEM[222]` | 0x6003DADE | 8 |  | - |
| `DS_Z_MEM[223]` | 0x6003DADF | 8 |  | - |
| `DS_Z_MEM[224]` | 0x6003DAE0 | 8 |  | - |
| `DS_Z_MEM[225]` | 0x6003DAE1 | 8 |  | - |
| `DS_Z_MEM[226]` | 0x6003DAE2 | 8 |  | - |
| `DS_Z_MEM[227]` | 0x6003DAE3 | 8 |  | - |
| `DS_Z_MEM[228]` | 0x6003DAE4 | 8 |  | - |
| `DS_Z_MEM[229]` | 0x6003DAE5 | 8 |  | - |
| `DS_Z_MEM[22]` | 0x6003DA16 | 8 |  | - |
| `DS_Z_MEM[230]` | 0x6003DAE6 | 8 |  | - |
| `DS_Z_MEM[231]` | 0x6003DAE7 | 8 |  | - |
| `DS_Z_MEM[232]` | 0x6003DAE8 | 8 |  | - |
| `DS_Z_MEM[233]` | 0x6003DAE9 | 8 |  | - |
| `DS_Z_MEM[234]` | 0x6003DAEA | 8 |  | - |
| `DS_Z_MEM[235]` | 0x6003DAEB | 8 |  | - |
| `DS_Z_MEM[236]` | 0x6003DAEC | 8 |  | - |
| `DS_Z_MEM[237]` | 0x6003DAED | 8 |  | - |
| `DS_Z_MEM[238]` | 0x6003DAEE | 8 |  | - |
| `DS_Z_MEM[239]` | 0x6003DAEF | 8 |  | - |
| `DS_Z_MEM[23]` | 0x6003DA17 | 8 |  | - |
| `DS_Z_MEM[240]` | 0x6003DAF0 | 8 |  | - |
| `DS_Z_MEM[241]` | 0x6003DAF1 | 8 |  | - |
| `DS_Z_MEM[242]` | 0x6003DAF2 | 8 |  | - |
| `DS_Z_MEM[243]` | 0x6003DAF3 | 8 |  | - |
| `DS_Z_MEM[244]` | 0x6003DAF4 | 8 |  | - |
| `DS_Z_MEM[245]` | 0x6003DAF5 | 8 |  | - |
| `DS_Z_MEM[246]` | 0x6003DAF6 | 8 |  | - |
| `DS_Z_MEM[247]` | 0x6003DAF7 | 8 |  | - |
| `DS_Z_MEM[248]` | 0x6003DAF8 | 8 |  | - |
| `DS_Z_MEM[249]` | 0x6003DAF9 | 8 |  | - |
| `DS_Z_MEM[24]` | 0x6003DA18 | 8 |  | - |
| `DS_Z_MEM[250]` | 0x6003DAFA | 8 |  | - |
| `DS_Z_MEM[251]` | 0x6003DAFB | 8 |  | - |
| `DS_Z_MEM[252]` | 0x6003DAFC | 8 |  | - |
| `DS_Z_MEM[253]` | 0x6003DAFD | 8 |  | - |
| `DS_Z_MEM[254]` | 0x6003DAFE | 8 |  | - |
| `DS_Z_MEM[255]` | 0x6003DAFF | 8 |  | - |
| `DS_Z_MEM[256]` | 0x6003DB00 | 8 |  | - |
| `DS_Z_MEM[257]` | 0x6003DB01 | 8 |  | - |
| `DS_Z_MEM[258]` | 0x6003DB02 | 8 |  | - |
| `DS_Z_MEM[259]` | 0x6003DB03 | 8 |  | - |
| `DS_Z_MEM[25]` | 0x6003DA19 | 8 |  | - |
| `DS_Z_MEM[260]` | 0x6003DB04 | 8 |  | - |
| `DS_Z_MEM[261]` | 0x6003DB05 | 8 |  | - |
| `DS_Z_MEM[262]` | 0x6003DB06 | 8 |  | - |
| `DS_Z_MEM[263]` | 0x6003DB07 | 8 |  | - |
| `DS_Z_MEM[264]` | 0x6003DB08 | 8 |  | - |
| `DS_Z_MEM[265]` | 0x6003DB09 | 8 |  | - |
| `DS_Z_MEM[266]` | 0x6003DB0A | 8 |  | - |
| `DS_Z_MEM[267]` | 0x6003DB0B | 8 |  | - |
| `DS_Z_MEM[268]` | 0x6003DB0C | 8 |  | - |
| `DS_Z_MEM[269]` | 0x6003DB0D | 8 |  | - |
| `DS_Z_MEM[26]` | 0x6003DA1A | 8 |  | - |
| `DS_Z_MEM[270]` | 0x6003DB0E | 8 |  | - |
| `DS_Z_MEM[271]` | 0x6003DB0F | 8 |  | - |
| `DS_Z_MEM[272]` | 0x6003DB10 | 8 |  | - |
| `DS_Z_MEM[273]` | 0x6003DB11 | 8 |  | - |
| `DS_Z_MEM[274]` | 0x6003DB12 | 8 |  | - |
| `DS_Z_MEM[275]` | 0x6003DB13 | 8 |  | - |
| `DS_Z_MEM[276]` | 0x6003DB14 | 8 |  | - |
| `DS_Z_MEM[277]` | 0x6003DB15 | 8 |  | - |
| `DS_Z_MEM[278]` | 0x6003DB16 | 8 |  | - |
| `DS_Z_MEM[279]` | 0x6003DB17 | 8 |  | - |
| `DS_Z_MEM[27]` | 0x6003DA1B | 8 |  | - |
| `DS_Z_MEM[280]` | 0x6003DB18 | 8 |  | - |
| `DS_Z_MEM[281]` | 0x6003DB19 | 8 |  | - |
| `DS_Z_MEM[282]` | 0x6003DB1A | 8 |  | - |
| `DS_Z_MEM[283]` | 0x6003DB1B | 8 |  | - |
| `DS_Z_MEM[284]` | 0x6003DB1C | 8 |  | - |
| `DS_Z_MEM[285]` | 0x6003DB1D | 8 |  | - |
| `DS_Z_MEM[286]` | 0x6003DB1E | 8 |  | - |
| `DS_Z_MEM[287]` | 0x6003DB1F | 8 |  | - |
| `DS_Z_MEM[288]` | 0x6003DB20 | 8 |  | - |
| `DS_Z_MEM[289]` | 0x6003DB21 | 8 |  | - |
| `DS_Z_MEM[28]` | 0x6003DA1C | 8 |  | - |
| `DS_Z_MEM[290]` | 0x6003DB22 | 8 |  | - |
| `DS_Z_MEM[291]` | 0x6003DB23 | 8 |  | - |
| `DS_Z_MEM[292]` | 0x6003DB24 | 8 |  | - |
| `DS_Z_MEM[293]` | 0x6003DB25 | 8 |  | - |
| `DS_Z_MEM[294]` | 0x6003DB26 | 8 |  | - |
| `DS_Z_MEM[295]` | 0x6003DB27 | 8 |  | - |
| `DS_Z_MEM[296]` | 0x6003DB28 | 8 |  | - |
| `DS_Z_MEM[297]` | 0x6003DB29 | 8 |  | - |
| `DS_Z_MEM[298]` | 0x6003DB2A | 8 |  | - |
| `DS_Z_MEM[299]` | 0x6003DB2B | 8 |  | - |
| `DS_Z_MEM[29]` | 0x6003DA1D | 8 |  | - |
| `DS_Z_MEM[2]` | 0x6003DA02 | 8 |  | - |
| `DS_Z_MEM[300]` | 0x6003DB2C | 8 |  | - |
| `DS_Z_MEM[301]` | 0x6003DB2D | 8 |  | - |
| `DS_Z_MEM[302]` | 0x6003DB2E | 8 |  | - |
| `DS_Z_MEM[303]` | 0x6003DB2F | 8 |  | - |
| `DS_Z_MEM[304]` | 0x6003DB30 | 8 |  | - |
| `DS_Z_MEM[305]` | 0x6003DB31 | 8 |  | - |
| `DS_Z_MEM[306]` | 0x6003DB32 | 8 |  | - |
| `DS_Z_MEM[307]` | 0x6003DB33 | 8 |  | - |
| `DS_Z_MEM[308]` | 0x6003DB34 | 8 |  | - |
| `DS_Z_MEM[309]` | 0x6003DB35 | 8 |  | - |
| `DS_Z_MEM[30]` | 0x6003DA1E | 8 |  | - |
| `DS_Z_MEM[310]` | 0x6003DB36 | 8 |  | - |
| `DS_Z_MEM[311]` | 0x6003DB37 | 8 |  | - |
| `DS_Z_MEM[312]` | 0x6003DB38 | 8 |  | - |
| `DS_Z_MEM[313]` | 0x6003DB39 | 8 |  | - |
| `DS_Z_MEM[314]` | 0x6003DB3A | 8 |  | - |
| `DS_Z_MEM[315]` | 0x6003DB3B | 8 |  | - |
| `DS_Z_MEM[316]` | 0x6003DB3C | 8 |  | - |
| `DS_Z_MEM[317]` | 0x6003DB3D | 8 |  | - |
| `DS_Z_MEM[318]` | 0x6003DB3E | 8 |  | - |
| `DS_Z_MEM[319]` | 0x6003DB3F | 8 |  | - |
| `DS_Z_MEM[31]` | 0x6003DA1F | 8 |  | - |
| `DS_Z_MEM[320]` | 0x6003DB40 | 8 |  | - |
| `DS_Z_MEM[321]` | 0x6003DB41 | 8 |  | - |
| `DS_Z_MEM[322]` | 0x6003DB42 | 8 |  | - |
| `DS_Z_MEM[323]` | 0x6003DB43 | 8 |  | - |
| `DS_Z_MEM[324]` | 0x6003DB44 | 8 |  | - |
| `DS_Z_MEM[325]` | 0x6003DB45 | 8 |  | - |
| `DS_Z_MEM[326]` | 0x6003DB46 | 8 |  | - |
| `DS_Z_MEM[327]` | 0x6003DB47 | 8 |  | - |
| `DS_Z_MEM[328]` | 0x6003DB48 | 8 |  | - |
| `DS_Z_MEM[329]` | 0x6003DB49 | 8 |  | - |
| `DS_Z_MEM[32]` | 0x6003DA20 | 8 |  | - |
| `DS_Z_MEM[330]` | 0x6003DB4A | 8 |  | - |
| `DS_Z_MEM[331]` | 0x6003DB4B | 8 |  | - |
| `DS_Z_MEM[332]` | 0x6003DB4C | 8 |  | - |
| `DS_Z_MEM[333]` | 0x6003DB4D | 8 |  | - |
| `DS_Z_MEM[334]` | 0x6003DB4E | 8 |  | - |
| `DS_Z_MEM[335]` | 0x6003DB4F | 8 |  | - |
| `DS_Z_MEM[336]` | 0x6003DB50 | 8 |  | - |
| `DS_Z_MEM[337]` | 0x6003DB51 | 8 |  | - |
| `DS_Z_MEM[338]` | 0x6003DB52 | 8 |  | - |
| `DS_Z_MEM[339]` | 0x6003DB53 | 8 |  | - |
| `DS_Z_MEM[33]` | 0x6003DA21 | 8 |  | - |
| `DS_Z_MEM[340]` | 0x6003DB54 | 8 |  | - |
| `DS_Z_MEM[341]` | 0x6003DB55 | 8 |  | - |
| `DS_Z_MEM[342]` | 0x6003DB56 | 8 |  | - |
| `DS_Z_MEM[343]` | 0x6003DB57 | 8 |  | - |
| `DS_Z_MEM[344]` | 0x6003DB58 | 8 |  | - |
| `DS_Z_MEM[345]` | 0x6003DB59 | 8 |  | - |
| `DS_Z_MEM[346]` | 0x6003DB5A | 8 |  | - |
| `DS_Z_MEM[347]` | 0x6003DB5B | 8 |  | - |
| `DS_Z_MEM[348]` | 0x6003DB5C | 8 |  | - |
| `DS_Z_MEM[349]` | 0x6003DB5D | 8 |  | - |
| `DS_Z_MEM[34]` | 0x6003DA22 | 8 |  | - |
| `DS_Z_MEM[350]` | 0x6003DB5E | 8 |  | - |
| `DS_Z_MEM[351]` | 0x6003DB5F | 8 |  | - |
| `DS_Z_MEM[352]` | 0x6003DB60 | 8 |  | - |
| `DS_Z_MEM[353]` | 0x6003DB61 | 8 |  | - |
| `DS_Z_MEM[354]` | 0x6003DB62 | 8 |  | - |
| `DS_Z_MEM[355]` | 0x6003DB63 | 8 |  | - |
| `DS_Z_MEM[356]` | 0x6003DB64 | 8 |  | - |
| `DS_Z_MEM[357]` | 0x6003DB65 | 8 |  | - |
| `DS_Z_MEM[358]` | 0x6003DB66 | 8 |  | - |
| `DS_Z_MEM[359]` | 0x6003DB67 | 8 |  | - |
| `DS_Z_MEM[35]` | 0x6003DA23 | 8 |  | - |
| `DS_Z_MEM[360]` | 0x6003DB68 | 8 |  | - |
| `DS_Z_MEM[361]` | 0x6003DB69 | 8 |  | - |
| `DS_Z_MEM[362]` | 0x6003DB6A | 8 |  | - |
| `DS_Z_MEM[363]` | 0x6003DB6B | 8 |  | - |
| `DS_Z_MEM[364]` | 0x6003DB6C | 8 |  | - |
| `DS_Z_MEM[365]` | 0x6003DB6D | 8 |  | - |
| `DS_Z_MEM[366]` | 0x6003DB6E | 8 |  | - |
| `DS_Z_MEM[367]` | 0x6003DB6F | 8 |  | - |
| `DS_Z_MEM[368]` | 0x6003DB70 | 8 |  | - |
| `DS_Z_MEM[369]` | 0x6003DB71 | 8 |  | - |
| `DS_Z_MEM[36]` | 0x6003DA24 | 8 |  | - |
| `DS_Z_MEM[370]` | 0x6003DB72 | 8 |  | - |
| `DS_Z_MEM[371]` | 0x6003DB73 | 8 |  | - |
| `DS_Z_MEM[372]` | 0x6003DB74 | 8 |  | - |
| `DS_Z_MEM[373]` | 0x6003DB75 | 8 |  | - |
| `DS_Z_MEM[374]` | 0x6003DB76 | 8 |  | - |
| `DS_Z_MEM[375]` | 0x6003DB77 | 8 |  | - |
| `DS_Z_MEM[376]` | 0x6003DB78 | 8 |  | - |
| `DS_Z_MEM[377]` | 0x6003DB79 | 8 |  | - |
| `DS_Z_MEM[378]` | 0x6003DB7A | 8 |  | - |
| `DS_Z_MEM[379]` | 0x6003DB7B | 8 |  | - |
| `DS_Z_MEM[37]` | 0x6003DA25 | 8 |  | - |
| `DS_Z_MEM[380]` | 0x6003DB7C | 8 |  | - |
| `DS_Z_MEM[381]` | 0x6003DB7D | 8 |  | - |
| `DS_Z_MEM[382]` | 0x6003DB7E | 8 |  | - |
| `DS_Z_MEM[383]` | 0x6003DB7F | 8 |  | - |
| `DS_Z_MEM[384]` | 0x6003DB80 | 8 |  | - |
| `DS_Z_MEM[385]` | 0x6003DB81 | 8 |  | - |
| `DS_Z_MEM[386]` | 0x6003DB82 | 8 |  | - |
| `DS_Z_MEM[387]` | 0x6003DB83 | 8 |  | - |
| `DS_Z_MEM[388]` | 0x6003DB84 | 8 |  | - |
| `DS_Z_MEM[389]` | 0x6003DB85 | 8 |  | - |
| `DS_Z_MEM[38]` | 0x6003DA26 | 8 |  | - |
| `DS_Z_MEM[390]` | 0x6003DB86 | 8 |  | - |
| `DS_Z_MEM[391]` | 0x6003DB87 | 8 |  | - |
| `DS_Z_MEM[392]` | 0x6003DB88 | 8 |  | - |
| `DS_Z_MEM[393]` | 0x6003DB89 | 8 |  | - |
| `DS_Z_MEM[394]` | 0x6003DB8A | 8 |  | - |
| `DS_Z_MEM[395]` | 0x6003DB8B | 8 |  | - |
| `DS_Z_MEM[396]` | 0x6003DB8C | 8 |  | - |
| `DS_Z_MEM[397]` | 0x6003DB8D | 8 |  | - |
| `DS_Z_MEM[398]` | 0x6003DB8E | 8 |  | - |
| `DS_Z_MEM[399]` | 0x6003DB8F | 8 |  | - |
| `DS_Z_MEM[39]` | 0x6003DA27 | 8 |  | - |
| `DS_Z_MEM[3]` | 0x6003DA03 | 8 |  | - |
| `DS_Z_MEM[400]` | 0x6003DB90 | 8 |  | - |
| `DS_Z_MEM[401]` | 0x6003DB91 | 8 |  | - |
| `DS_Z_MEM[402]` | 0x6003DB92 | 8 |  | - |
| `DS_Z_MEM[403]` | 0x6003DB93 | 8 |  | - |
| `DS_Z_MEM[404]` | 0x6003DB94 | 8 |  | - |
| `DS_Z_MEM[405]` | 0x6003DB95 | 8 |  | - |
| `DS_Z_MEM[406]` | 0x6003DB96 | 8 |  | - |
| `DS_Z_MEM[407]` | 0x6003DB97 | 8 |  | - |
| `DS_Z_MEM[408]` | 0x6003DB98 | 8 |  | - |
| `DS_Z_MEM[409]` | 0x6003DB99 | 8 |  | - |
| `DS_Z_MEM[40]` | 0x6003DA28 | 8 |  | - |
| `DS_Z_MEM[410]` | 0x6003DB9A | 8 |  | - |
| `DS_Z_MEM[411]` | 0x6003DB9B | 8 |  | - |
| `DS_Z_MEM[412]` | 0x6003DB9C | 8 |  | - |
| `DS_Z_MEM[413]` | 0x6003DB9D | 8 |  | - |
| `DS_Z_MEM[414]` | 0x6003DB9E | 8 |  | - |
| `DS_Z_MEM[415]` | 0x6003DB9F | 8 |  | - |
| `DS_Z_MEM[416]` | 0x6003DBA0 | 8 |  | - |
| `DS_Z_MEM[417]` | 0x6003DBA1 | 8 |  | - |
| `DS_Z_MEM[418]` | 0x6003DBA2 | 8 |  | - |
| `DS_Z_MEM[419]` | 0x6003DBA3 | 8 |  | - |
| `DS_Z_MEM[41]` | 0x6003DA29 | 8 |  | - |
| `DS_Z_MEM[420]` | 0x6003DBA4 | 8 |  | - |
| `DS_Z_MEM[421]` | 0x6003DBA5 | 8 |  | - |
| `DS_Z_MEM[422]` | 0x6003DBA6 | 8 |  | - |
| `DS_Z_MEM[423]` | 0x6003DBA7 | 8 |  | - |
| `DS_Z_MEM[424]` | 0x6003DBA8 | 8 |  | - |
| `DS_Z_MEM[425]` | 0x6003DBA9 | 8 |  | - |
| `DS_Z_MEM[426]` | 0x6003DBAA | 8 |  | - |
| `DS_Z_MEM[427]` | 0x6003DBAB | 8 |  | - |
| `DS_Z_MEM[428]` | 0x6003DBAC | 8 |  | - |
| `DS_Z_MEM[429]` | 0x6003DBAD | 8 |  | - |
| `DS_Z_MEM[42]` | 0x6003DA2A | 8 |  | - |
| `DS_Z_MEM[430]` | 0x6003DBAE | 8 |  | - |
| `DS_Z_MEM[431]` | 0x6003DBAF | 8 |  | - |
| `DS_Z_MEM[432]` | 0x6003DBB0 | 8 |  | - |
| `DS_Z_MEM[433]` | 0x6003DBB1 | 8 |  | - |
| `DS_Z_MEM[434]` | 0x6003DBB2 | 8 |  | - |
| `DS_Z_MEM[435]` | 0x6003DBB3 | 8 |  | - |
| `DS_Z_MEM[436]` | 0x6003DBB4 | 8 |  | - |
| `DS_Z_MEM[437]` | 0x6003DBB5 | 8 |  | - |
| `DS_Z_MEM[438]` | 0x6003DBB6 | 8 |  | - |
| `DS_Z_MEM[439]` | 0x6003DBB7 | 8 |  | - |
| `DS_Z_MEM[43]` | 0x6003DA2B | 8 |  | - |
| `DS_Z_MEM[440]` | 0x6003DBB8 | 8 |  | - |
| `DS_Z_MEM[441]` | 0x6003DBB9 | 8 |  | - |
| `DS_Z_MEM[442]` | 0x6003DBBA | 8 |  | - |
| `DS_Z_MEM[443]` | 0x6003DBBB | 8 |  | - |
| `DS_Z_MEM[444]` | 0x6003DBBC | 8 |  | - |
| `DS_Z_MEM[445]` | 0x6003DBBD | 8 |  | - |
| `DS_Z_MEM[446]` | 0x6003DBBE | 8 |  | - |
| `DS_Z_MEM[447]` | 0x6003DBBF | 8 |  | - |
| `DS_Z_MEM[448]` | 0x6003DBC0 | 8 |  | - |
| `DS_Z_MEM[449]` | 0x6003DBC1 | 8 |  | - |
| `DS_Z_MEM[44]` | 0x6003DA2C | 8 |  | - |
| `DS_Z_MEM[450]` | 0x6003DBC2 | 8 |  | - |
| `DS_Z_MEM[451]` | 0x6003DBC3 | 8 |  | - |
| `DS_Z_MEM[452]` | 0x6003DBC4 | 8 |  | - |
| `DS_Z_MEM[453]` | 0x6003DBC5 | 8 |  | - |
| `DS_Z_MEM[454]` | 0x6003DBC6 | 8 |  | - |
| `DS_Z_MEM[455]` | 0x6003DBC7 | 8 |  | - |
| `DS_Z_MEM[456]` | 0x6003DBC8 | 8 |  | - |
| `DS_Z_MEM[457]` | 0x6003DBC9 | 8 |  | - |
| `DS_Z_MEM[458]` | 0x6003DBCA | 8 |  | - |
| `DS_Z_MEM[459]` | 0x6003DBCB | 8 |  | - |
| `DS_Z_MEM[45]` | 0x6003DA2D | 8 |  | - |
| `DS_Z_MEM[460]` | 0x6003DBCC | 8 |  | - |
| `DS_Z_MEM[461]` | 0x6003DBCD | 8 |  | - |
| `DS_Z_MEM[462]` | 0x6003DBCE | 8 |  | - |
| `DS_Z_MEM[463]` | 0x6003DBCF | 8 |  | - |
| `DS_Z_MEM[464]` | 0x6003DBD0 | 8 |  | - |
| `DS_Z_MEM[465]` | 0x6003DBD1 | 8 |  | - |
| `DS_Z_MEM[466]` | 0x6003DBD2 | 8 |  | - |
| `DS_Z_MEM[467]` | 0x6003DBD3 | 8 |  | - |
| `DS_Z_MEM[468]` | 0x6003DBD4 | 8 |  | - |
| `DS_Z_MEM[469]` | 0x6003DBD5 | 8 |  | - |
| `DS_Z_MEM[46]` | 0x6003DA2E | 8 |  | - |
| `DS_Z_MEM[470]` | 0x6003DBD6 | 8 |  | - |
| `DS_Z_MEM[471]` | 0x6003DBD7 | 8 |  | - |
| `DS_Z_MEM[472]` | 0x6003DBD8 | 8 |  | - |
| `DS_Z_MEM[473]` | 0x6003DBD9 | 8 |  | - |
| `DS_Z_MEM[474]` | 0x6003DBDA | 8 |  | - |
| `DS_Z_MEM[475]` | 0x6003DBDB | 8 |  | - |
| `DS_Z_MEM[476]` | 0x6003DBDC | 8 |  | - |
| `DS_Z_MEM[477]` | 0x6003DBDD | 8 |  | - |
| `DS_Z_MEM[478]` | 0x6003DBDE | 8 |  | - |
| `DS_Z_MEM[479]` | 0x6003DBDF | 8 |  | - |
| `DS_Z_MEM[47]` | 0x6003DA2F | 8 |  | - |
| `DS_Z_MEM[480]` | 0x6003DBE0 | 8 |  | - |
| `DS_Z_MEM[481]` | 0x6003DBE1 | 8 |  | - |
| `DS_Z_MEM[482]` | 0x6003DBE2 | 8 |  | - |
| `DS_Z_MEM[483]` | 0x6003DBE3 | 8 |  | - |
| `DS_Z_MEM[484]` | 0x6003DBE4 | 8 |  | - |
| `DS_Z_MEM[485]` | 0x6003DBE5 | 8 |  | - |
| `DS_Z_MEM[486]` | 0x6003DBE6 | 8 |  | - |
| `DS_Z_MEM[487]` | 0x6003DBE7 | 8 |  | - |
| `DS_Z_MEM[488]` | 0x6003DBE8 | 8 |  | - |
| `DS_Z_MEM[489]` | 0x6003DBE9 | 8 |  | - |
| `DS_Z_MEM[48]` | 0x6003DA30 | 8 |  | - |
| `DS_Z_MEM[490]` | 0x6003DBEA | 8 |  | - |
| `DS_Z_MEM[491]` | 0x6003DBEB | 8 |  | - |
| `DS_Z_MEM[492]` | 0x6003DBEC | 8 |  | - |
| `DS_Z_MEM[493]` | 0x6003DBED | 8 |  | - |
| `DS_Z_MEM[494]` | 0x6003DBEE | 8 |  | - |
| `DS_Z_MEM[495]` | 0x6003DBEF | 8 |  | - |
| `DS_Z_MEM[496]` | 0x6003DBF0 | 8 |  | - |
| `DS_Z_MEM[497]` | 0x6003DBF1 | 8 |  | - |
| `DS_Z_MEM[498]` | 0x6003DBF2 | 8 |  | - |
| `DS_Z_MEM[499]` | 0x6003DBF3 | 8 |  | - |
| `DS_Z_MEM[49]` | 0x6003DA31 | 8 |  | - |
| `DS_Z_MEM[4]` | 0x6003DA04 | 8 |  | - |
| `DS_Z_MEM[500]` | 0x6003DBF4 | 8 |  | - |
| `DS_Z_MEM[501]` | 0x6003DBF5 | 8 |  | - |
| `DS_Z_MEM[502]` | 0x6003DBF6 | 8 |  | - |
| `DS_Z_MEM[503]` | 0x6003DBF7 | 8 |  | - |
| `DS_Z_MEM[504]` | 0x6003DBF8 | 8 |  | - |
| `DS_Z_MEM[505]` | 0x6003DBF9 | 8 |  | - |
| `DS_Z_MEM[506]` | 0x6003DBFA | 8 |  | - |
| `DS_Z_MEM[507]` | 0x6003DBFB | 8 |  | - |
| `DS_Z_MEM[508]` | 0x6003DBFC | 8 |  | - |
| `DS_Z_MEM[509]` | 0x6003DBFD | 8 |  | - |
| `DS_Z_MEM[50]` | 0x6003DA32 | 8 |  | - |
| `DS_Z_MEM[510]` | 0x6003DBFE | 8 |  | - |
| `DS_Z_MEM[511]` | 0x6003DBFF | 8 |  | - |
| `DS_Z_MEM[51]` | 0x6003DA33 | 8 |  | - |
| `DS_Z_MEM[52]` | 0x6003DA34 | 8 |  | - |
| `DS_Z_MEM[53]` | 0x6003DA35 | 8 |  | - |
| `DS_Z_MEM[54]` | 0x6003DA36 | 8 |  | - |
| `DS_Z_MEM[55]` | 0x6003DA37 | 8 |  | - |
| `DS_Z_MEM[56]` | 0x6003DA38 | 8 |  | - |
| `DS_Z_MEM[57]` | 0x6003DA39 | 8 |  | - |
| `DS_Z_MEM[58]` | 0x6003DA3A | 8 |  | - |
| `DS_Z_MEM[59]` | 0x6003DA3B | 8 |  | - |
| `DS_Z_MEM[5]` | 0x6003DA05 | 8 |  | - |
| `DS_Z_MEM[60]` | 0x6003DA3C | 8 |  | - |
| `DS_Z_MEM[61]` | 0x6003DA3D | 8 |  | - |
| `DS_Z_MEM[62]` | 0x6003DA3E | 8 |  | - |
| `DS_Z_MEM[63]` | 0x6003DA3F | 8 |  | - |
| `DS_Z_MEM[64]` | 0x6003DA40 | 8 |  | - |
| `DS_Z_MEM[65]` | 0x6003DA41 | 8 |  | - |
| `DS_Z_MEM[66]` | 0x6003DA42 | 8 |  | - |
| `DS_Z_MEM[67]` | 0x6003DA43 | 8 |  | - |
| `DS_Z_MEM[68]` | 0x6003DA44 | 8 |  | - |
| `DS_Z_MEM[69]` | 0x6003DA45 | 8 |  | - |
| `DS_Z_MEM[6]` | 0x6003DA06 | 8 |  | - |
| `DS_Z_MEM[70]` | 0x6003DA46 | 8 |  | - |
| `DS_Z_MEM[71]` | 0x6003DA47 | 8 |  | - |
| `DS_Z_MEM[72]` | 0x6003DA48 | 8 |  | - |
| `DS_Z_MEM[73]` | 0x6003DA49 | 8 |  | - |
| `DS_Z_MEM[74]` | 0x6003DA4A | 8 |  | - |
| `DS_Z_MEM[75]` | 0x6003DA4B | 8 |  | - |
| `DS_Z_MEM[76]` | 0x6003DA4C | 8 |  | - |
| `DS_Z_MEM[77]` | 0x6003DA4D | 8 |  | - |
| `DS_Z_MEM[78]` | 0x6003DA4E | 8 |  | - |
| `DS_Z_MEM[79]` | 0x6003DA4F | 8 |  | - |
| `DS_Z_MEM[7]` | 0x6003DA07 | 8 |  | - |
| `DS_Z_MEM[80]` | 0x6003DA50 | 8 |  | - |
| `DS_Z_MEM[81]` | 0x6003DA51 | 8 |  | - |
| `DS_Z_MEM[82]` | 0x6003DA52 | 8 |  | - |
| `DS_Z_MEM[83]` | 0x6003DA53 | 8 |  | - |
| `DS_Z_MEM[84]` | 0x6003DA54 | 8 |  | - |
| `DS_Z_MEM[85]` | 0x6003DA55 | 8 |  | - |
| `DS_Z_MEM[86]` | 0x6003DA56 | 8 |  | - |
| `DS_Z_MEM[87]` | 0x6003DA57 | 8 |  | - |
| `DS_Z_MEM[88]` | 0x6003DA58 | 8 |  | - |
| `DS_Z_MEM[89]` | 0x6003DA59 | 8 |  | - |
| `DS_Z_MEM[8]` | 0x6003DA08 | 8 |  | - |
| `DS_Z_MEM[90]` | 0x6003DA5A | 8 |  | - |
| `DS_Z_MEM[91]` | 0x6003DA5B | 8 |  | - |
| `DS_Z_MEM[92]` | 0x6003DA5C | 8 |  | - |
| `DS_Z_MEM[93]` | 0x6003DA5D | 8 |  | - |
| `DS_Z_MEM[94]` | 0x6003DA5E | 8 |  | - |
| `DS_Z_MEM[95]` | 0x6003DA5F | 8 |  | - |
| `DS_Z_MEM[96]` | 0x6003DA60 | 8 |  | - |
| `DS_Z_MEM[97]` | 0x6003DA61 | 8 |  | - |
| `DS_Z_MEM[98]` | 0x6003DA62 | 8 |  | - |
| `DS_Z_MEM[99]` | 0x6003DA63 | 8 |  | - |
| `DS_Z_MEM[9]` | 0x6003DA09 | 8 |  | - |
| `EFUSE_CLK` | 0x3F41A1C8 | 32 | EFUSE_MEM_FORCE_PD=0x1 EFUSE_MEM_FORCE_PU=0x4 EN=0x10000 MEM_CLK_FORCE_ON=0x2 | - |
| `EFUSE_CMD` | 0x3F41A1D4 | 32 | BLK_NUM=0x3C PGM_CMD=0x2 READ_CMD=0x1 | - |
| `EFUSE_CONF` | 0x3F41A1CC | 32 | OP_CODE=0xFFFF | - |
| `EFUSE_DAC_CONF` | 0x3F41A1E8 | 32 | DAC_CLK_DIV=0xFF DAC_CLK_PAD_SEL=0x100 DAC_NUM=0x1FE00 OE_CLR=0x20000 | - |
| `EFUSE_DATE` | 0x3F41A1FC | 32 | DATE=0xFFFFFFFF | - |
| `EFUSE_INT_CLR` | 0x3F41A1E4 | 32 | PGM_DONE_INT_CLR=0x2 READ_DONE_INT_CLR=0x1 | - |
| `EFUSE_INT_ENA` | 0x3F41A1E0 | 32 | PGM_DONE_INT_ENA=0x2 READ_DONE_INT_ENA=0x1 | - |
| `EFUSE_INT_RAW` | 0x3F41A1D8 | 32 | PGM_DONE_INT_RAW=0x2 READ_DONE_INT_RAW=0x1 | - |
| `EFUSE_INT_ST` | 0x3F41A1DC | 32 | PGM_DONE_INT_ST=0x2 READ_DONE_INT_ST=0x1 | - |
| `EFUSE_PGM_CHECK_VALUE0` | 0x3F41A020 | 32 | PGM_RS_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_CHECK_VALUE1` | 0x3F41A024 | 32 | PGM_RS_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_CHECK_VALUE2` | 0x3F41A028 | 32 | PGM_RS_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA0` | 0x3F41A000 | 32 | PGM_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA1` | 0x3F41A004 | 32 | PGM_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA2` | 0x3F41A008 | 32 | PGM_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA3` | 0x3F41A00C | 32 | PGM_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA4` | 0x3F41A010 | 32 | PGM_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA5` | 0x3F41A014 | 32 | PGM_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA6` | 0x3F41A018 | 32 | PGM_DATA_0=0xFFFFFFFF | - |
| `EFUSE_PGM_DATA7` | 0x3F41A01C | 32 | PGM_DATA_0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA0` | 0x3F41A09C | 32 | KEY0_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA1` | 0x3F41A0A0 | 32 | KEY0_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA2` | 0x3F41A0A4 | 32 | KEY0_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA3` | 0x3F41A0A8 | 32 | KEY0_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA4` | 0x3F41A0AC | 32 | KEY0_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA5` | 0x3F41A0B0 | 32 | KEY0_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA6` | 0x3F41A0B4 | 32 | KEY0_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY0_DATA7` | 0x3F41A0B8 | 32 | KEY0_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA0` | 0x3F41A0BC | 32 | KEY1_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA1` | 0x3F41A0C0 | 32 | KEY1_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA2` | 0x3F41A0C4 | 32 | KEY1_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA3` | 0x3F41A0C8 | 32 | KEY1_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA4` | 0x3F41A0CC | 32 | KEY1_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA5` | 0x3F41A0D0 | 32 | KEY1_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA6` | 0x3F41A0D4 | 32 | KEY1_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY1_DATA7` | 0x3F41A0D8 | 32 | KEY1_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA0` | 0x3F41A0DC | 32 | KEY2_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA1` | 0x3F41A0E0 | 32 | KEY2_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA2` | 0x3F41A0E4 | 32 | KEY2_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA3` | 0x3F41A0E8 | 32 | KEY2_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA4` | 0x3F41A0EC | 32 | KEY2_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA5` | 0x3F41A0F0 | 32 | KEY2_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA6` | 0x3F41A0F4 | 32 | KEY2_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY2_DATA7` | 0x3F41A0F8 | 32 | KEY2_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA0` | 0x3F41A0FC | 32 | KEY3_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA1` | 0x3F41A100 | 32 | KEY3_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA2` | 0x3F41A104 | 32 | KEY3_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA3` | 0x3F41A108 | 32 | KEY3_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA4` | 0x3F41A10C | 32 | KEY3_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA5` | 0x3F41A110 | 32 | KEY3_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA6` | 0x3F41A114 | 32 | KEY3_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY3_DATA7` | 0x3F41A118 | 32 | KEY3_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA0` | 0x3F41A11C | 32 | KEY4_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA1` | 0x3F41A120 | 32 | KEY4_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA2` | 0x3F41A124 | 32 | KEY4_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA3` | 0x3F41A128 | 32 | KEY4_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA4` | 0x3F41A12C | 32 | KEY4_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA5` | 0x3F41A130 | 32 | KEY4_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA6` | 0x3F41A134 | 32 | KEY4_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY4_DATA7` | 0x3F41A138 | 32 | KEY4_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA0` | 0x3F41A13C | 32 | KEY5_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA1` | 0x3F41A140 | 32 | KEY5_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA2` | 0x3F41A144 | 32 | KEY5_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA3` | 0x3F41A148 | 32 | KEY5_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA4` | 0x3F41A14C | 32 | KEY5_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA5` | 0x3F41A150 | 32 | KEY5_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA6` | 0x3F41A154 | 32 | KEY5_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_KEY5_DATA7` | 0x3F41A158 | 32 | KEY5_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_MAC_SPI_SYS_0` | 0x3F41A044 | 32 | MAC_0=0xFFFFFFFF | - |
| `EFUSE_RD_MAC_SPI_SYS_1` | 0x3F41A048 | 32 | MAC_1=0xFFFF SPI_PAD_CONF_0=0xFFFF0000 | - |
| `EFUSE_RD_MAC_SPI_SYS_2` | 0x3F41A04C | 32 | SPI_PAD_CONF_1=0xFFFFFFFF | - |
| `EFUSE_RD_MAC_SPI_SYS_3` | 0x3F41A050 | 32 | SPI_PAD_CONF_2=0x3FFFF SYS_DATA_PART0_0=0xFFFC0000 | - |
| `EFUSE_RD_MAC_SPI_SYS_4` | 0x3F41A054 | 32 | SYS_DATA_PART0_1=0xFFFFFFFF | - |
| `EFUSE_RD_MAC_SPI_SYS_5` | 0x3F41A058 | 32 | SYS_DATA_PART0_2=0xFFFFFFFF | - |
| `EFUSE_RD_REPEAT_DATA0` | 0x3F41A030 | 32 | DIS_BOOT_REMAP=0x8000 DIS_CAN=0x4000 DIS_DCACHE=0x200 DIS_DOWNLOAD_DCACHE=0x800 DIS_DOWNLOAD_ICACHE=0x400 DIS_DOWNLOAD_MANUAL_ENCRYPT=0x80000 DIS_FORCE_DOWNLOAD=0x1000 DIS_ICACHE=0x100 DIS_RTC_RAM_BOOT=0x80 DIS_USB=0x2000 EXT_PHY_ENABLE=0x2... | - |
| `EFUSE_RD_REPEAT_DATA1` | 0x3F41A034 | 32 | KEY_PURPOSE_0=0xF000000 KEY_PURPOSE_1=0xF0000000 SECURE_BOOT_KEY_REVOKE0=0x200000 SECURE_BOOT_KEY_REVOKE1=0x400000 SECURE_BOOT_KEY_REVOKE2=0x800000 SPI_BOOT_CRYPT_CNT=0x1C0000 VDD_SPI_DCAP=0xC000 VDD_SPI_DCURLIM=0xE00 VDD_SPI_DREFL=0xC VDD_... | - |
| `EFUSE_RD_REPEAT_DATA2` | 0x3F41A038 | 32 | FLASH_TPUW=0xF0000000 KEY_PURPOSE_2=0xF KEY_PURPOSE_3=0xF0 KEY_PURPOSE_4=0xF00 KEY_PURPOSE_5=0xF000 KEY_PURPOSE_6=0xF0000 RPT4_RESERVED1=0xFC00000 SECURE_BOOT_AGGRESSIVE_REVOKE=0x200000 SECURE_BOOT_EN=0x100000 | - |
| `EFUSE_RD_REPEAT_DATA3` | 0x3F41A03C | 32 | DIS_DOWNLOAD_MODE=0x1 DIS_LEGACY_SPI_BOOT=0x2 DIS_USB_DOWNLOAD_MODE=0x10 ENABLE_SECURITY_DOWNLOAD=0x20 FLASH_TYPE=0x200 FORCE_SEND_RESUME=0x400 PIN_POWER_SELECTION=0x100 RPT4_RESERVED2=0xF8000000 RPT4_RESERVED3=0x8 SECURE_VERSION=0x7FFF800 ... | - |
| `EFUSE_RD_REPEAT_DATA4` | 0x3F41A040 | 32 | RPT4_RESERVED4=0xFFFFFF | - |
| `EFUSE_RD_REPEAT_ERR0` | 0x3F41A17C | 32 | DIS_BOOT_REMAP_ERR=0x8000 DIS_CAN_ERR=0x4000 DIS_DCACHE_ERR=0x200 DIS_DOWNLOAD_DCACHE_ERR=0x800 DIS_DOWNLOAD_ICACHE_ERR=0x400 DIS_DOWNLOAD_MANUAL_ENCRYPT_ERR=0x80000 DIS_FORCE_DOWNLOAD_ERR=0x1000 DIS_ICACHE_ERR=0x100 DIS_RTC_RAM_BOOT_ERR=0x... | - |
| `EFUSE_RD_REPEAT_ERR1` | 0x3F41A180 | 32 | KEY_PURPOSE_0_ERR=0xF000000 KEY_PURPOSE_1_ERR=0xF0000000 SECURE_BOOT_KEY_REVOKE0_ERR=0x200000 SECURE_BOOT_KEY_REVOKE1_ERR=0x400000 SECURE_BOOT_KEY_REVOKE2_ERR=0x800000 SPI_BOOT_CRYPT_CNT_ERR=0x1C0000 VDD_SPI_DCAP_ERR=0xC000 VDD_SPI_DCURLIM_... | - |
| `EFUSE_RD_REPEAT_ERR2` | 0x3F41A184 | 32 | FLASH_TPUW_ERR=0xF0000000 KEY_PURPOSE_2_ERR=0xF KEY_PURPOSE_3_ERR=0xF0 KEY_PURPOSE_4_ERR=0xF00 KEY_PURPOSE_5_ERR=0xF000 KEY_PURPOSE_6_ERR=0xF0000 RPT4_RESERVED1_ERR=0xFC00000 SECURE_BOOT_AGGRESSIVE_REVOKE_ERR=0x200000 SECURE_BOOT_EN_ERR=0x1... | - |
| `EFUSE_RD_REPEAT_ERR3` | 0x3F41A188 | 32 | DIS_DOWNLOAD_MODE_ERR=0x1 DIS_LEGACY_SPI_BOOT_ERR=0x2 DIS_USB_DOWNLOAD_MODE_ERR=0x10 ENABLE_SECURITY_DOWNLOAD_ERR=0x20 FLASH_TYPE_ERR=0x200 FORCE_SEND_RESUME_ERR=0x400 PIN_POWER_SELECTION_ERR=0x100 RPT4_RESERVED2_ERR=0xF8000000 RPT4_RESERVE... | - |
| `EFUSE_RD_REPEAT_ERR4` | 0x3F41A190 | 32 | RPT4_RESERVED4_ERR=0xFFFFFF | - |
| `EFUSE_RD_RS_ERR0` | 0x3F41A1C0 | 32 | KEY0_ERR_NUM=0x7000 KEY0_FAIL=0x8000 KEY1_ERR_NUM=0x70000 KEY1_FAIL=0x80000 KEY2_ERR_NUM=0x700000 KEY2_FAIL=0x800000 KEY3_ERR_NUM=0x7000000 KEY3_FAIL=0x8000000 KEY4_ERR_NUM=0x70000000 KEY4_FAIL=0x80000000 MAC_SPI_8M_ERR_NUM=0x7 MAC_SPI_8M_F... | - |
| `EFUSE_RD_RS_ERR1` | 0x3F41A1C4 | 32 | KEY5_ERR_NUM=0x7 KEY5_FAIL=0x8 SYS_PART2_ERR_NUM=0x70 SYS_PART2_FAIL=0x80 | - |
| `EFUSE_RD_SYS_DATA_PART1_0` | 0x3F41A05C | 32 | SYS_DATA_PART1_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART1_1` | 0x3F41A060 | 32 | SYS_DATA_PART1_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART1_2` | 0x3F41A064 | 32 | SYS_DATA_PART1_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART1_3` | 0x3F41A068 | 32 | SYS_DATA_PART1_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART1_4` | 0x3F41A06C | 32 | SYS_DATA_PART1_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART1_5` | 0x3F41A070 | 32 | SYS_DATA_PART1_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART1_6` | 0x3F41A074 | 32 | SYS_DATA_PART1_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART1_7` | 0x3F41A078 | 32 | SYS_DATA_PART1_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART2_0` | 0x3F41A15C | 32 | SYS_DATA_PART2_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART2_1` | 0x3F41A160 | 32 | SYS_DATA_PART2_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART2_2` | 0x3F41A164 | 32 | SYS_DATA_PART2_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART2_3` | 0x3F41A168 | 32 | SYS_DATA_PART2_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART2_4` | 0x3F41A16C | 32 | SYS_DATA_PART2_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART2_5` | 0x3F41A170 | 32 | SYS_DATA_PART2_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART2_6` | 0x3F41A174 | 32 | SYS_DATA_PART2_0=0xFFFFFFFF | - |
| `EFUSE_RD_SYS_DATA_PART2_7` | 0x3F41A178 | 32 | SYS_DATA_PART2_0=0xFFFFFFFF | - |
| `EFUSE_RD_TIM_CONF` | 0x3F41A1EC | 32 | READ_INIT_NUM=0xFF000000 THR_A=0xFF TRD=0xFF00 TSUR_A=0xFF0000 | - |
| `EFUSE_RD_USR_DATA0` | 0x3F41A07C | 32 | USR_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA1` | 0x3F41A080 | 32 | USR_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA2` | 0x3F41A084 | 32 | USR_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA3` | 0x3F41A088 | 32 | USR_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA4` | 0x3F41A08C | 32 | USR_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA5` | 0x3F41A090 | 32 | USR_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA6` | 0x3F41A094 | 32 | USR_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_USR_DATA7` | 0x3F41A098 | 32 | USR_DATA0=0xFFFFFFFF | - |
| `EFUSE_RD_WR_DIS` | 0x3F41A02C | 32 | WR_DIS=0xFFFFFFFF | - |
| `EFUSE_STATUS` | 0x3F41A1D0 | 32 | OTP_CSB_SW=0x80 OTP_LOAD_SW=0x10 OTP_PGENB_SW=0x100 OTP_STROBE_SW=0x40 OTP_VDDQ_C_SYNC2=0x20 OTP_VDDQ_IS_SW=0x200 REPEAT_ERR_CNT=0x3FC00 STATE=0xF | - |
| `EFUSE_WR_TIM_CONF0` | 0x3F41A1F0 | 32 | THP_A=0xFF TPGM=0xFFFF0000 TPGM_INACTIVE=0xFF00 | - |
| `EFUSE_WR_TIM_CONF1` | 0x3F41A1F4 | 32 | PWR_ON_NUM=0xFFFF00 TSUP_A=0xFF | - |
| `EFUSE_WR_TIM_CONF2` | 0x3F41A1F8 | 32 | PWR_OFF_NUM=0xFFFF | - |
| `EXTMEM_CACHE_BRIDGE_ARBITER_CTRL` | 0x61800128 | 32 | ALLOC_WB_HOLD_ARBITER=0x1 | - |
| `EXTMEM_CACHE_CONF_MISC` | 0x61800134 | 32 | PRO_CACHE_IGNORE_PRELOAD_MMU_ENTRY_FAULT=0x1 PRO_CACHE_IGNORE_SYNC_MMU_ENTRY_FAULT=0x2 | - |
| `EXTMEM_CACHE_DBG_INT_CLR` | 0x618000EC | 32 | DBUS_ACS_MSK_DC_INT_CLR=0x40 DBUS_CNT_OVF_INT_CLR=0x80 DCACHE_REJECT_INT_CLR=0x800 DCACHE_SET_ILG_INT_CLR=0x1000 DCACHE_WRITE_FLASH_INT_CLR=0x400 DC_PRELOAD_SIZE_FAULT_INT_CLR=0x200 DC_SYNC_SIZE_FAULT_INT_CLR=0x100 IBUS_ACS_MSK_IC_INT_CLR=0... | - |
| `EXTMEM_CACHE_DBG_INT_ENA` | 0x618000E8 | 32 | CACHE_DBG_EN=0x1 DBUS_ACS_MSK_DC_INT_ENA=0x400 DBUS_CNT_OVF_INT_ENA=0x800 DCACHE_REJECT_INT_ENA=0x8000 DCACHE_SET_LOCK_ILG_INT_ENA=0x40000 DCACHE_SET_PRELOAD_ILG_INT_ENA=0x10000 DCACHE_SET_SYNC_ILG_INT_ENA=0x20000 DCACHE_WRITE_FLASH_INT_ENA... | - |
| `EXTMEM_CACHE_DBG_STATUS0` | 0x618000F0 | 32 | IBUS0_ABANDON_CNT_OVF_ST=0x1000 IBUS0_ACS_CNT_OVF_ST=0x10 IBUS0_ACS_MISS_CNT_OVF_ST=0x100 IBUS0_ACS_MSK_ICACHE_ST=0x1 IBUS1_ABANDON_CNT_OVF_ST=0x2000 IBUS1_ACS_CNT_OVF_ST=0x20 IBUS1_ACS_MISS_CNT_OVF_ST=0x200 IBUS1_ACS_MSK_ICACHE_ST=0x2 IBUS... | - |
| `EXTMEM_CACHE_DBG_STATUS1` | 0x618000F4 | 32 | DBUS0_ABANDON_CNT_OVF_ST=0x10000 DBUS0_ACS_CNT_OVF_ST=0x10 DBUS0_ACS_MISS_CNT_OVF_ST=0x100 DBUS0_ACS_MSK_DCACHE_ST=0x1 DBUS0_ACS_WB_CNT_OVF_ST=0x1000 DBUS1_ABANDON_CNT_OVF_ST=0x20000 DBUS1_ACS_CNT_OVF_ST=0x20 DBUS1_ACS_MISS_CNT_OVF_ST=0x200... | - |
| `EXTMEM_CACHE_ENCRYPT_DECRYPT_CLK_FORCE_ON` | 0x61800124 | 32 | CLK_FORCE_ON_AUTOMATIC_ENCRYPT_DECRYPT=0x4 CLK_FORCE_ON_DB_ENCRYPT=0x1 CLK_FORCE_ON_G0CB_DECRYPT=0x2 | - |
| `EXTMEM_CACHE_ENCRYPT_DECRYPT_RECORD_DISABLE` | 0x61800120 | 32 | RECORD_DISABLE_DB_ENCRYPT=0x1 RECORD_DISABLE_G0CB_DECRYPT=0x2 | - |
| `EXTMEM_CACHE_PRELOAD_INT_CTRL` | 0x6180012C | 32 | PRO_DCACHE_PRELOAD_INT_CLR=0x20 PRO_DCACHE_PRELOAD_INT_ENA=0x10 PRO_DCACHE_PRELOAD_INT_ST=0x8 PRO_ICACHE_PRELOAD_INT_CLR=0x4 PRO_ICACHE_PRELOAD_INT_ENA=0x2 PRO_ICACHE_PRELOAD_INT_ST=0x1 | - |
| `EXTMEM_CACHE_SYNC_INT_CTRL` | 0x61800130 | 32 | PRO_DCACHE_SYNC_INT_CLR=0x20 PRO_DCACHE_SYNC_INT_ENA=0x10 PRO_DCACHE_SYNC_INT_ST=0x8 PRO_ICACHE_SYNC_INT_CLR=0x4 PRO_ICACHE_SYNC_INT_ENA=0x2 PRO_ICACHE_SYNC_INT_ST=0x1 | - |
| `EXTMEM_CLOCK_GATE` | 0x61800138 | 32 | CLK_EN=0x1 | - |
| `EXTMEM_DBUS0_ABANDON_CNT` | 0x618000C0 | 32 | DBUS0_ABANDON_CNT=0xFFFF | - |
| `EXTMEM_DBUS0_ACS_CNT` | 0x618000E4 | 32 | DBUS0_ACS_CNT=0xFFFFFFFF | - |
| `EXTMEM_DBUS0_ACS_MISS_CNT` | 0x618000D8 | 32 | DBUS0_ACS_MISS_CNT=0xFFFFFFFF | - |
| `EXTMEM_DBUS0_ACS_WB_CNT` | 0x618000CC | 32 | DBUS0_ACS_WB_CNT=0xFFFFF | - |
| `EXTMEM_DBUS1_ABANDON_CNT` | 0x618000BC | 32 | DBUS1_ABANDON_CNT=0xFFFF | - |
| `EXTMEM_DBUS1_ACS_CNT` | 0x618000E0 | 32 | DBUS1_ACS_CNT=0xFFFFFFFF | - |
| `EXTMEM_DBUS1_ACS_MISS_CNT` | 0x618000D4 | 32 | DBUS1_ACS_MISS_CNT=0xFFFFFFFF | - |
| `EXTMEM_DBUS1_ACS_WB_CNT` | 0x618000C8 | 32 | DBUS1_ACS_WB_CNT=0xFFFFF | - |
| `EXTMEM_DBUS2_ABANDON_CNT` | 0x618000B8 | 32 | DBUS2_ABANDON_CNT=0xFFFF | - |
| `EXTMEM_DBUS2_ACS_CNT` | 0x618000DC | 32 | DBUS2_ACS_CNT=0xFFFFFFFF | - |
| `EXTMEM_DBUS2_ACS_MISS_CNT` | 0x618000D0 | 32 | DBUS2_ACS_MISS_CNT=0xFFFFFFFF | - |
| `EXTMEM_DBUS2_ACS_WB_CNT` | 0x618000C4 | 32 | DBUS2_ACS_WB_CNT=0xFFFFF | - |
| `EXTMEM_DC_PRELOAD_CNT` | 0x618000AC | 32 | DC_PRELOAD_CNT=0xFFFF | - |
| `EXTMEM_DC_PRELOAD_EVICT_CNT` | 0x618000B0 | 32 | DC_PRELOAD_EVICT_CNT=0xFFFF | - |
| `EXTMEM_DC_PRELOAD_MISS_CNT` | 0x618000B4 | 32 | DC_PRELOAD_MISS_CNT=0xFFFF | - |
| `EXTMEM_IBUS0_ABANDON_CNT` | 0x61800090 | 32 | IBUS0_ABANDON_CNT=0xFFFF | - |
| `EXTMEM_IBUS0_ACS_CNT` | 0x618000A8 | 32 | IBUS0_ACS_CNT=0xFFFFFFFF | - |
| `EXTMEM_IBUS0_ACS_MISS_CNT` | 0x6180009C | 32 | IBUS0_ACS_MISS_CNT=0xFFFFFFFF | - |
| `EXTMEM_IBUS1_ABANDON_CNT` | 0x6180008C | 32 | IBUS1_ABANDON_CNT=0xFFFF | - |
| `EXTMEM_IBUS1_ACS_CNT` | 0x618000A4 | 32 | IBUS1_ACS_CNT=0xFFFFFFFF | - |
| `EXTMEM_IBUS1_ACS_MISS_CNT` | 0x61800098 | 32 | IBUS1_ACS_MISS_CNT=0xFFFFFFFF | - |
| `EXTMEM_IBUS2_ABANDON_CNT` | 0x61800088 | 32 | IBUS2_ABANDON_CNT=0xFFFF | - |
| `EXTMEM_IBUS2_ACS_CNT` | 0x618000A0 | 32 | IBUS2_ACS_CNT=0xFFFFFFFF | - |
| `EXTMEM_IBUS2_ACS_MISS_CNT` | 0x61800094 | 32 | IBUS2_ACS_MISS_CNT=0xFFFFFFFF | - |
| `EXTMEM_IC_PRELOAD_CNT` | 0x61800080 | 32 | IC_PRELOAD_CNT=0xFFFF | - |
| `EXTMEM_IC_PRELOAD_MISS_CNT` | 0x61800084 | 32 | IC_PRELOAD_MISS_CNT=0xFFFF | - |
| `EXTMEM_PRO_CACHE_ACS_CNT_CLR` | 0x618000F8 | 32 | PRO_DCACHE_ACS_CNT_CLR=0x1 PRO_ICACHE_ACS_CNT_CLR=0x2 | - |
| `EXTMEM_PRO_CACHE_MMU_FAULT_CONTENT` | 0x6180010C | 32 | PRO_CACHE_MMU_FAULT_CODE=0xE0000 PRO_CACHE_MMU_FAULT_CONTENT=0x1FFFF | - |
| `EXTMEM_PRO_CACHE_MMU_FAULT_VADDR` | 0x61800110 | 32 | PRO_CACHE_MMU_FAULT_VADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_CACHE_MMU_POWER_CTRL` | 0x61800118 | 32 | PRO_CACHE_MMU_MEM_FORCE_ON=0x1 PRO_CACHE_MMU_MEM_FORCE_PD=0x2 PRO_CACHE_MMU_MEM_FORCE_PU=0x4 | - |
| `EXTMEM_PRO_CACHE_STATE` | 0x6180011C | 32 | PRO_DCACHE_STATE=0xFFF000 PRO_ICACHE_STATE=0xFFF | - |
| `EXTMEM_PRO_CACHE_WRAP_AROUND_CTRL` | 0x61800114 | 32 | PRO_CACHE_FLASH_WRAP_AROUND=0x1 PRO_CACHE_SRAM_RD_WRAP_AROUND=0x2 | - |
| `EXTMEM_PRO_DCACHE_AUTOLOAD_CFG` | 0x6180002C | 32 | PRO_DCACHE_AUTOLOAD_MODE=0x1 PRO_DCACHE_AUTOLOAD_ORDER=0x8 PRO_DCACHE_AUTOLOAD_RQST=0x30 PRO_DCACHE_AUTOLOAD_SCT0_ENA=0x100 PRO_DCACHE_AUTOLOAD_SCT1_ENA=0x200 PRO_DCACHE_AUTOLOAD_SIZE=0xC0 PRO_DCACHE_AUTOLOAD_STEP=0x6 | - |
| `EXTMEM_PRO_DCACHE_AUTOLOAD_SECTION0_ADDR` | 0x61800030 | 32 | PRO_DCACHE_AUTOLOAD_SCT0_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_DCACHE_AUTOLOAD_SECTION0_SIZE` | 0x61800034 | 32 | PRO_DCACHE_AUTOLOAD_SCT0_SIZE=0xFFFFFF | - |
| `EXTMEM_PRO_DCACHE_AUTOLOAD_SECTION1_ADDR` | 0x61800038 | 32 | PRO_DCACHE_AUTOLOAD_SCT1_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_DCACHE_AUTOLOAD_SECTION1_SIZE` | 0x6180003C | 32 | PRO_DCACHE_AUTOLOAD_SCT1_SIZE=0xFFFFFF | - |
| `EXTMEM_PRO_DCACHE_CTRL` | 0x61800000 | 32 | PRO_DCACHE_AUTOLOAD_DONE=0x80000 PRO_DCACHE_AUTOLOAD_ENA=0x40000 PRO_DCACHE_BLOCKSIZE_MODE=0x8 PRO_DCACHE_CLEAN_DONE=0x2000 PRO_DCACHE_CLEAN_ENA=0x1000 PRO_DCACHE_ENABLE=0x1 PRO_DCACHE_FLUSH_DONE=0x800 PRO_DCACHE_FLUSH_ENA=0x400 PRO_DCACHE_... | - |
| `EXTMEM_PRO_DCACHE_CTRL1` | 0x61800004 | 32 | PRO_DCACHE_MASK_BUS0=0x1 PRO_DCACHE_MASK_BUS1=0x2 PRO_DCACHE_MASK_BUS2=0x4 | - |
| `EXTMEM_PRO_DCACHE_LOCK0_ADDR` | 0x6180000C | 32 | PRO_DCACHE_LOCK0_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_DCACHE_LOCK0_SIZE` | 0x61800010 | 32 | PRO_DCACHE_LOCK0_SIZE=0xFFFF | - |
| `EXTMEM_PRO_DCACHE_LOCK1_ADDR` | 0x61800014 | 32 | PRO_DCACHE_LOCK1_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_DCACHE_LOCK1_SIZE` | 0x61800018 | 32 | PRO_DCACHE_LOCK1_SIZE=0xFFFF | - |
| `EXTMEM_PRO_DCACHE_MEM_SYNC0` | 0x6180001C | 32 | PRO_DCACHE_MEMSYNC_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_DCACHE_MEM_SYNC1` | 0x61800020 | 32 | PRO_DCACHE_MEMSYNC_SIZE=0x7FFFF | - |
| `EXTMEM_PRO_DCACHE_PRELOAD_ADDR` | 0x61800024 | 32 | PRO_DCACHE_PRELOAD_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_DCACHE_PRELOAD_SIZE` | 0x61800028 | 32 | PRO_DCACHE_PRELOAD_ORDER=0x400 PRO_DCACHE_PRELOAD_SIZE=0x3FF | - |
| `EXTMEM_PRO_DCACHE_REJECT_ST` | 0x618000FC | 32 | PRO_DCACHE_CPU_ATTR=0x38 PRO_DCACHE_TAG_ATTR=0x7 | - |
| `EXTMEM_PRO_DCACHE_REJECT_VADDR` | 0x61800100 | 32 | PRO_DCACHE_CPU_VADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_DCACHE_TAG_POWER_CTRL` | 0x61800008 | 32 | PRO_DCACHE_TAG_MEM_FORCE_ON=0x1 PRO_DCACHE_TAG_MEM_FORCE_PD=0x2 PRO_DCACHE_TAG_MEM_FORCE_PU=0x4 | - |
| `EXTMEM_PRO_EXTMEM_REG_DATE` | 0x618003FC | 32 | PRO_EXTMEM_REG_DATE=0xFFFFFFF | - |
| `EXTMEM_PRO_ICACHE_AUTOLOAD_CFG` | 0x6180006C | 32 | PRO_ICACHE_AUTOLOAD_MODE=0x1 PRO_ICACHE_AUTOLOAD_ORDER=0x8 PRO_ICACHE_AUTOLOAD_RQST=0x30 PRO_ICACHE_AUTOLOAD_SCT0_ENA=0x100 PRO_ICACHE_AUTOLOAD_SCT1_ENA=0x200 PRO_ICACHE_AUTOLOAD_SIZE=0xC0 PRO_ICACHE_AUTOLOAD_STEP=0x6 | - |
| `EXTMEM_PRO_ICACHE_AUTOLOAD_SECTION0_ADDR` | 0x61800070 | 32 | PRO_ICACHE_AUTOLOAD_SCT0_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_ICACHE_AUTOLOAD_SECTION0_SIZE` | 0x61800074 | 32 | PRO_ICACHE_AUTOLOAD_SCT0_SIZE=0xFFFFFF | - |
| `EXTMEM_PRO_ICACHE_AUTOLOAD_SECTION1_ADDR` | 0x61800078 | 32 | PRO_ICACHE_AUTOLOAD_SCT1_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_ICACHE_AUTOLOAD_SECTION1_SIZE` | 0x6180007C | 32 | PRO_ICACHE_AUTOLOAD_SCT1_SIZE=0xFFFFFF | - |
| `EXTMEM_PRO_ICACHE_CTRL` | 0x61800040 | 32 | PRO_ICACHE_AUTOLOAD_DONE=0x80000 PRO_ICACHE_AUTOLOAD_ENA=0x40000 PRO_ICACHE_BLOCKSIZE_MODE=0x8 PRO_ICACHE_ENABLE=0x1 PRO_ICACHE_INVALIDATE_DONE=0x200 PRO_ICACHE_INVALIDATE_ENA=0x100 PRO_ICACHE_LOCK0_EN=0x4000 PRO_ICACHE_LOCK1_EN=0x8000 PRO_... | - |
| `EXTMEM_PRO_ICACHE_CTRL1` | 0x61800044 | 32 | PRO_ICACHE_MASK_BUS0=0x1 PRO_ICACHE_MASK_BUS1=0x2 PRO_ICACHE_MASK_BUS2=0x4 | - |
| `EXTMEM_PRO_ICACHE_LOCK0_ADDR` | 0x6180004C | 32 | PRO_ICACHE_LOCK0_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_ICACHE_LOCK0_SIZE` | 0x61800050 | 32 | PRO_ICACHE_LOCK0_SIZE=0xFFFF | - |
| `EXTMEM_PRO_ICACHE_LOCK1_ADDR` | 0x61800054 | 32 | PRO_ICACHE_LOCK1_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_ICACHE_LOCK1_SIZE` | 0x61800058 | 32 | PRO_ICACHE_LOCK1_SIZE=0xFFFF | - |
| `EXTMEM_PRO_ICACHE_MEM_SYNC0` | 0x6180005C | 32 | PRO_ICACHE_MEMSYNC_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_ICACHE_MEM_SYNC1` | 0x61800060 | 32 | PRO_ICACHE_MEMSYNC_SIZE=0x7FFFF | - |
| `EXTMEM_PRO_ICACHE_PRELOAD_ADDR` | 0x61800064 | 32 | PRO_ICACHE_PRELOAD_ADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_ICACHE_PRELOAD_SIZE` | 0x61800068 | 32 | PRO_ICACHE_PRELOAD_ORDER=0x400 PRO_ICACHE_PRELOAD_SIZE=0x3FF | - |
| `EXTMEM_PRO_ICACHE_REJECT_ST` | 0x61800104 | 32 | PRO_ICACHE_CPU_ATTR=0x38 PRO_ICACHE_TAG_ATTR=0x7 | - |
| `EXTMEM_PRO_ICACHE_REJECT_VADDR` | 0x61800108 | 32 | PRO_ICACHE_CPU_VADDR=0xFFFFFFFF | - |
| `EXTMEM_PRO_ICACHE_TAG_POWER_CTRL` | 0x61800048 | 32 | PRO_ICACHE_TAG_MEM_FORCE_ON=0x1 PRO_ICACHE_TAG_MEM_FORCE_PD=0x2 PRO_ICACHE_TAG_MEM_FORCE_PU=0x4 | - |
| `GPIO_BT_SELECT` | 0x3F404000 | 32 | BT_SEL=0xFFFFFFFF | - |
| `GPIO_CLOCK_GATE` | 0x3F40462C | 32 | CLK_EN=0x1 | - |
| `GPIO_CPUSDIO_INT` | 0x3F404064 | 32 | SDIO_INT=0xFFFFFFFF | - |
| `GPIO_CPUSDIO_INT1` | 0x3F404070 | 32 | SDIO1_INT=0x3FFFFF | - |
| `GPIO_ENABLE` | 0x3F404020 | 32 | DATA=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_ENABLE1` | 0x3F40402C | 32 | DATA=0x3FFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_ENABLE1_W1TC` | 0x3F404034 | 32 | ENABLE1_W1TC=0x3FFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_ENABLE1_W1TS` | 0x3F404030 | 32 | ENABLE1_W1TS=0x3FFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_ENABLE_W1TC` | 0x3F404028 | 32 | ENABLE_W1TC=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_ENABLE_W1TS` | 0x3F404024 | 32 | ENABLE_W1TS=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_FUNC0_IN_SEL_CFG` | 0x3F404154 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC0_OUT_SEL_CFG` | 0x3F404554 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC100_IN_SEL_CFG` | 0x3F4042E4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC101_IN_SEL_CFG` | 0x3F4042E8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC102_IN_SEL_CFG` | 0x3F4042EC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC103_IN_SEL_CFG` | 0x3F4042F0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC104_IN_SEL_CFG` | 0x3F4042F4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC105_IN_SEL_CFG` | 0x3F4042F8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC106_IN_SEL_CFG` | 0x3F4042FC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC107_IN_SEL_CFG` | 0x3F404300 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC108_IN_SEL_CFG` | 0x3F404304 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC109_IN_SEL_CFG` | 0x3F404308 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC10_IN_SEL_CFG` | 0x3F40417C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC10_OUT_SEL_CFG` | 0x3F40457C | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC110_IN_SEL_CFG` | 0x3F40430C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC111_IN_SEL_CFG` | 0x3F404310 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC112_IN_SEL_CFG` | 0x3F404314 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC113_IN_SEL_CFG` | 0x3F404318 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC114_IN_SEL_CFG` | 0x3F40431C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC115_IN_SEL_CFG` | 0x3F404320 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC116_IN_SEL_CFG` | 0x3F404324 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC117_IN_SEL_CFG` | 0x3F404328 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC118_IN_SEL_CFG` | 0x3F40432C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC119_IN_SEL_CFG` | 0x3F404330 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC11_IN_SEL_CFG` | 0x3F404180 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC11_OUT_SEL_CFG` | 0x3F404580 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC120_IN_SEL_CFG` | 0x3F404334 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC121_IN_SEL_CFG` | 0x3F404338 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC122_IN_SEL_CFG` | 0x3F40433C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC123_IN_SEL_CFG` | 0x3F404340 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC124_IN_SEL_CFG` | 0x3F404344 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC125_IN_SEL_CFG` | 0x3F404348 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC126_IN_SEL_CFG` | 0x3F40434C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC127_IN_SEL_CFG` | 0x3F404350 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC128_IN_SEL_CFG` | 0x3F404354 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC129_IN_SEL_CFG` | 0x3F404358 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC12_IN_SEL_CFG` | 0x3F404184 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC12_OUT_SEL_CFG` | 0x3F404584 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC130_IN_SEL_CFG` | 0x3F40435C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC131_IN_SEL_CFG` | 0x3F404360 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC132_IN_SEL_CFG` | 0x3F404364 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC133_IN_SEL_CFG` | 0x3F404368 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC134_IN_SEL_CFG` | 0x3F40436C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC135_IN_SEL_CFG` | 0x3F404370 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC136_IN_SEL_CFG` | 0x3F404374 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC137_IN_SEL_CFG` | 0x3F404378 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC138_IN_SEL_CFG` | 0x3F40437C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC139_IN_SEL_CFG` | 0x3F404380 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC13_IN_SEL_CFG` | 0x3F404188 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC13_OUT_SEL_CFG` | 0x3F404588 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC140_IN_SEL_CFG` | 0x3F404384 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC141_IN_SEL_CFG` | 0x3F404388 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC142_IN_SEL_CFG` | 0x3F40438C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC143_IN_SEL_CFG` | 0x3F404390 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC144_IN_SEL_CFG` | 0x3F404394 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC145_IN_SEL_CFG` | 0x3F404398 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC146_IN_SEL_CFG` | 0x3F40439C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC147_IN_SEL_CFG` | 0x3F4043A0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC148_IN_SEL_CFG` | 0x3F4043A4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC149_IN_SEL_CFG` | 0x3F4043A8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC14_IN_SEL_CFG` | 0x3F40418C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC14_OUT_SEL_CFG` | 0x3F40458C | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC150_IN_SEL_CFG` | 0x3F4043AC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC151_IN_SEL_CFG` | 0x3F4043B0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC152_IN_SEL_CFG` | 0x3F4043B4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC153_IN_SEL_CFG` | 0x3F4043B8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC154_IN_SEL_CFG` | 0x3F4043BC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC155_IN_SEL_CFG` | 0x3F4043C0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC156_IN_SEL_CFG` | 0x3F4043C4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC157_IN_SEL_CFG` | 0x3F4043C8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC158_IN_SEL_CFG` | 0x3F4043CC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC159_IN_SEL_CFG` | 0x3F4043D0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC15_IN_SEL_CFG` | 0x3F404190 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC15_OUT_SEL_CFG` | 0x3F404590 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC160_IN_SEL_CFG` | 0x3F4043D4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC161_IN_SEL_CFG` | 0x3F4043D8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC162_IN_SEL_CFG` | 0x3F4043DC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC163_IN_SEL_CFG` | 0x3F4043E0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC164_IN_SEL_CFG` | 0x3F4043E4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC165_IN_SEL_CFG` | 0x3F4043E8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC166_IN_SEL_CFG` | 0x3F4043EC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC167_IN_SEL_CFG` | 0x3F4043F0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC168_IN_SEL_CFG` | 0x3F4043F4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC169_IN_SEL_CFG` | 0x3F4043F8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC16_IN_SEL_CFG` | 0x3F404194 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC16_OUT_SEL_CFG` | 0x3F404594 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC170_IN_SEL_CFG` | 0x3F4043FC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC171_IN_SEL_CFG` | 0x3F404400 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC172_IN_SEL_CFG` | 0x3F404404 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC173_IN_SEL_CFG` | 0x3F404408 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC174_IN_SEL_CFG` | 0x3F40440C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC175_IN_SEL_CFG` | 0x3F404410 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC176_IN_SEL_CFG` | 0x3F404414 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC177_IN_SEL_CFG` | 0x3F404418 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC178_IN_SEL_CFG` | 0x3F40441C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC179_IN_SEL_CFG` | 0x3F404420 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC17_IN_SEL_CFG` | 0x3F404198 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC17_OUT_SEL_CFG` | 0x3F404598 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC180_IN_SEL_CFG` | 0x3F404424 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC181_IN_SEL_CFG` | 0x3F404428 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC182_IN_SEL_CFG` | 0x3F40442C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC183_IN_SEL_CFG` | 0x3F404430 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC184_IN_SEL_CFG` | 0x3F404434 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC185_IN_SEL_CFG` | 0x3F404438 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC186_IN_SEL_CFG` | 0x3F40443C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC187_IN_SEL_CFG` | 0x3F404440 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC188_IN_SEL_CFG` | 0x3F404444 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC189_IN_SEL_CFG` | 0x3F404448 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC18_IN_SEL_CFG` | 0x3F40419C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC18_OUT_SEL_CFG` | 0x3F40459C | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC190_IN_SEL_CFG` | 0x3F40444C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC191_IN_SEL_CFG` | 0x3F404450 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC192_IN_SEL_CFG` | 0x3F404454 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC193_IN_SEL_CFG` | 0x3F404458 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC194_IN_SEL_CFG` | 0x3F40445C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC195_IN_SEL_CFG` | 0x3F404460 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC196_IN_SEL_CFG` | 0x3F404464 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC197_IN_SEL_CFG` | 0x3F404468 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC198_IN_SEL_CFG` | 0x3F40446C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC199_IN_SEL_CFG` | 0x3F404470 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC19_IN_SEL_CFG` | 0x3F4041A0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC19_OUT_SEL_CFG` | 0x3F4045A0 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC1_IN_SEL_CFG` | 0x3F404158 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC1_OUT_SEL_CFG` | 0x3F404558 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC200_IN_SEL_CFG` | 0x3F404474 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC201_IN_SEL_CFG` | 0x3F404478 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC202_IN_SEL_CFG` | 0x3F40447C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC203_IN_SEL_CFG` | 0x3F404480 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC204_IN_SEL_CFG` | 0x3F404484 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC205_IN_SEL_CFG` | 0x3F404488 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC206_IN_SEL_CFG` | 0x3F40448C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC207_IN_SEL_CFG` | 0x3F404490 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC208_IN_SEL_CFG` | 0x3F404494 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC209_IN_SEL_CFG` | 0x3F404498 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC20_IN_SEL_CFG` | 0x3F4041A4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC20_OUT_SEL_CFG` | 0x3F4045A4 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC210_IN_SEL_CFG` | 0x3F40449C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC211_IN_SEL_CFG` | 0x3F4044A0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC212_IN_SEL_CFG` | 0x3F4044A4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC213_IN_SEL_CFG` | 0x3F4044A8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC214_IN_SEL_CFG` | 0x3F4044AC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC215_IN_SEL_CFG` | 0x3F4044B0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC216_IN_SEL_CFG` | 0x3F4044B4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC217_IN_SEL_CFG` | 0x3F4044B8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC218_IN_SEL_CFG` | 0x3F4044BC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC219_IN_SEL_CFG` | 0x3F4044C0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC21_IN_SEL_CFG` | 0x3F4041A8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC21_OUT_SEL_CFG` | 0x3F4045A8 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC220_IN_SEL_CFG` | 0x3F4044C4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC221_IN_SEL_CFG` | 0x3F4044C8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC222_IN_SEL_CFG` | 0x3F4044CC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC223_IN_SEL_CFG` | 0x3F4044D0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC224_IN_SEL_CFG` | 0x3F4044D4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC225_IN_SEL_CFG` | 0x3F4044D8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC226_IN_SEL_CFG` | 0x3F4044DC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC227_IN_SEL_CFG` | 0x3F4044E0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC228_IN_SEL_CFG` | 0x3F4044E4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC229_IN_SEL_CFG` | 0x3F4044E8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC22_IN_SEL_CFG` | 0x3F4041AC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC22_OUT_SEL_CFG` | 0x3F4045AC | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC230_IN_SEL_CFG` | 0x3F4044EC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC231_IN_SEL_CFG` | 0x3F4044F0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC232_IN_SEL_CFG` | 0x3F4044F4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC233_IN_SEL_CFG` | 0x3F4044F8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC234_IN_SEL_CFG` | 0x3F4044FC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC235_IN_SEL_CFG` | 0x3F404500 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC236_IN_SEL_CFG` | 0x3F404504 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC237_IN_SEL_CFG` | 0x3F404508 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC238_IN_SEL_CFG` | 0x3F40450C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC239_IN_SEL_CFG` | 0x3F404510 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC23_IN_SEL_CFG` | 0x3F4041B0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC23_OUT_SEL_CFG` | 0x3F4045B0 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC240_IN_SEL_CFG` | 0x3F404514 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC241_IN_SEL_CFG` | 0x3F404518 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC242_IN_SEL_CFG` | 0x3F40451C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC243_IN_SEL_CFG` | 0x3F404520 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC244_IN_SEL_CFG` | 0x3F404524 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC245_IN_SEL_CFG` | 0x3F404528 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC246_IN_SEL_CFG` | 0x3F40452C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC247_IN_SEL_CFG` | 0x3F404530 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC248_IN_SEL_CFG` | 0x3F404534 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC249_IN_SEL_CFG` | 0x3F404538 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC24_IN_SEL_CFG` | 0x3F4041B4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC24_OUT_SEL_CFG` | 0x3F4045B4 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC250_IN_SEL_CFG` | 0x3F40453C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC251_IN_SEL_CFG` | 0x3F404540 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC252_IN_SEL_CFG` | 0x3F404544 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC253_IN_SEL_CFG` | 0x3F404548 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC254_IN_SEL_CFG` | 0x3F40454C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC255_IN_SEL_CFG` | 0x3F404550 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC25_IN_SEL_CFG` | 0x3F4041B8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC25_OUT_SEL_CFG` | 0x3F4045B8 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC26_IN_SEL_CFG` | 0x3F4041BC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC26_OUT_SEL_CFG` | 0x3F4045BC | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC27_IN_SEL_CFG` | 0x3F4041C0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC27_OUT_SEL_CFG` | 0x3F4045C0 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC28_IN_SEL_CFG` | 0x3F4041C4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC28_OUT_SEL_CFG` | 0x3F4045C4 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC29_IN_SEL_CFG` | 0x3F4041C8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC29_OUT_SEL_CFG` | 0x3F4045C8 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC2_IN_SEL_CFG` | 0x3F40415C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC2_OUT_SEL_CFG` | 0x3F40455C | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC30_IN_SEL_CFG` | 0x3F4041CC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC30_OUT_SEL_CFG` | 0x3F4045CC | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC31_IN_SEL_CFG` | 0x3F4041D0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC31_OUT_SEL_CFG` | 0x3F4045D0 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC32_IN_SEL_CFG` | 0x3F4041D4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC32_OUT_SEL_CFG` | 0x3F4045D4 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC33_IN_SEL_CFG` | 0x3F4041D8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC33_OUT_SEL_CFG` | 0x3F4045D8 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC34_IN_SEL_CFG` | 0x3F4041DC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC34_OUT_SEL_CFG` | 0x3F4045DC | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC35_IN_SEL_CFG` | 0x3F4041E0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC35_OUT_SEL_CFG` | 0x3F4045E0 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC36_IN_SEL_CFG` | 0x3F4041E4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC36_OUT_SEL_CFG` | 0x3F4045E4 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC37_IN_SEL_CFG` | 0x3F4041E8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC37_OUT_SEL_CFG` | 0x3F4045E8 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC38_IN_SEL_CFG` | 0x3F4041EC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC38_OUT_SEL_CFG` | 0x3F4045EC | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC39_IN_SEL_CFG` | 0x3F4041F0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC39_OUT_SEL_CFG` | 0x3F4045F0 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC3_IN_SEL_CFG` | 0x3F404160 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC3_OUT_SEL_CFG` | 0x3F404560 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC40_IN_SEL_CFG` | 0x3F4041F4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC40_OUT_SEL_CFG` | 0x3F4045F4 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC41_IN_SEL_CFG` | 0x3F4041F8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC41_OUT_SEL_CFG` | 0x3F4045F8 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC42_IN_SEL_CFG` | 0x3F4041FC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC42_OUT_SEL_CFG` | 0x3F4045FC | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC43_IN_SEL_CFG` | 0x3F404200 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC43_OUT_SEL_CFG` | 0x3F404600 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC44_IN_SEL_CFG` | 0x3F404204 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC44_OUT_SEL_CFG` | 0x3F404604 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC45_IN_SEL_CFG` | 0x3F404208 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC45_OUT_SEL_CFG` | 0x3F404608 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC46_IN_SEL_CFG` | 0x3F40420C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC46_OUT_SEL_CFG` | 0x3F40460C | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC47_IN_SEL_CFG` | 0x3F404210 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC47_OUT_SEL_CFG` | 0x3F404610 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC48_IN_SEL_CFG` | 0x3F404214 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC48_OUT_SEL_CFG` | 0x3F404614 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC49_IN_SEL_CFG` | 0x3F404218 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC49_OUT_SEL_CFG` | 0x3F404618 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC4_IN_SEL_CFG` | 0x3F404164 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC4_OUT_SEL_CFG` | 0x3F404564 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC50_IN_SEL_CFG` | 0x3F40421C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC50_OUT_SEL_CFG` | 0x3F40461C | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC51_IN_SEL_CFG` | 0x3F404220 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC51_OUT_SEL_CFG` | 0x3F404620 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC52_IN_SEL_CFG` | 0x3F404224 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC52_OUT_SEL_CFG` | 0x3F404624 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC53_IN_SEL_CFG` | 0x3F404228 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC53_OUT_SEL_CFG` | 0x3F404628 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC54_IN_SEL_CFG` | 0x3F40422C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC55_IN_SEL_CFG` | 0x3F404230 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC56_IN_SEL_CFG` | 0x3F404234 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC57_IN_SEL_CFG` | 0x3F404238 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC58_IN_SEL_CFG` | 0x3F40423C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC59_IN_SEL_CFG` | 0x3F404240 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC5_IN_SEL_CFG` | 0x3F404168 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC5_OUT_SEL_CFG` | 0x3F404568 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC60_IN_SEL_CFG` | 0x3F404244 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC61_IN_SEL_CFG` | 0x3F404248 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC62_IN_SEL_CFG` | 0x3F40424C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC63_IN_SEL_CFG` | 0x3F404250 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC64_IN_SEL_CFG` | 0x3F404254 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC65_IN_SEL_CFG` | 0x3F404258 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC66_IN_SEL_CFG` | 0x3F40425C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC67_IN_SEL_CFG` | 0x3F404260 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC68_IN_SEL_CFG` | 0x3F404264 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC69_IN_SEL_CFG` | 0x3F404268 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC6_IN_SEL_CFG` | 0x3F40416C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC6_OUT_SEL_CFG` | 0x3F40456C | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC70_IN_SEL_CFG` | 0x3F40426C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC71_IN_SEL_CFG` | 0x3F404270 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC72_IN_SEL_CFG` | 0x3F404274 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC73_IN_SEL_CFG` | 0x3F404278 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC74_IN_SEL_CFG` | 0x3F40427C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC75_IN_SEL_CFG` | 0x3F404280 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC76_IN_SEL_CFG` | 0x3F404284 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC77_IN_SEL_CFG` | 0x3F404288 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC78_IN_SEL_CFG` | 0x3F40428C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC79_IN_SEL_CFG` | 0x3F404290 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC7_IN_SEL_CFG` | 0x3F404170 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC7_OUT_SEL_CFG` | 0x3F404570 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC80_IN_SEL_CFG` | 0x3F404294 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC81_IN_SEL_CFG` | 0x3F404298 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC82_IN_SEL_CFG` | 0x3F40429C | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC83_IN_SEL_CFG` | 0x3F4042A0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC84_IN_SEL_CFG` | 0x3F4042A4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC85_IN_SEL_CFG` | 0x3F4042A8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC86_IN_SEL_CFG` | 0x3F4042AC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC87_IN_SEL_CFG` | 0x3F4042B0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC88_IN_SEL_CFG` | 0x3F4042B4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC89_IN_SEL_CFG` | 0x3F4042B8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC8_IN_SEL_CFG` | 0x3F404174 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC8_OUT_SEL_CFG` | 0x3F404574 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_FUNC90_IN_SEL_CFG` | 0x3F4042BC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC91_IN_SEL_CFG` | 0x3F4042C0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC92_IN_SEL_CFG` | 0x3F4042C4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC93_IN_SEL_CFG` | 0x3F4042C8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC94_IN_SEL_CFG` | 0x3F4042CC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC95_IN_SEL_CFG` | 0x3F4042D0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC96_IN_SEL_CFG` | 0x3F4042D4 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC97_IN_SEL_CFG` | 0x3F4042D8 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC98_IN_SEL_CFG` | 0x3F4042DC | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC99_IN_SEL_CFG` | 0x3F4042E0 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC9_IN_SEL_CFG` | 0x3F404178 | 32 | IN_INV_SEL=0x40 IN_SEL=0x3F SEL=0x80 | - |
| `GPIO_FUNC9_OUT_SEL_CFG` | 0x3F404578 | 32 | INV_SEL=0x200 OEN_INV_SEL=0x800 OEN_SEL=0x400 OUT_SEL=0x1FF | - |
| `GPIO_IN` | 0x3F40403C | 32 | DATA_NEXT=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_IN1` | 0x3F404040 | 32 | IN_DATA1_NEXT=0x3FFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_OUT` | 0x3F404004 | 32 | DATA_ORIG=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_OUT1` | 0x3F404010 | 32 | DATA_ORIG=0x3FFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_OUT1_W1TC` | 0x3F404018 | 32 | OUT1_W1TC=0x3FFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_OUT1_W1TS` | 0x3F404014 | 32 | OUT1_W1TS=0x3FFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_OUT_W1TC` | 0x3F40400C | 32 | OUT_W1TC=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_OUT_W1TS` | 0x3F404008 | 32 | OUT_W1TS=0xFFFFFFFF | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31 |
| `GPIO_PCPU_INT` | 0x3F40405C | 32 | PROCPU_INT=0xFFFFFFFF | - |
| `GPIO_PCPU_INT1` | 0x3F404068 | 32 | PROCPU1_INT=0x3FFFFF | - |
| `GPIO_PCPU_NMI_INT` | 0x3F404060 | 32 | PROCPU_NMI_INT=0xFFFFFFFF | - |
| `GPIO_PCPU_NMI_INT1` | 0x3F40406C | 32 | PROCPU_NMI1_INT=0x3FFFFF | - |
| `GPIO_PIN0` | 0x3F404074 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN1` | 0x3F404078 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN10` | 0x3F40409C | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN11` | 0x3F4040A0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN12` | 0x3F4040A4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN13` | 0x3F4040A8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN14` | 0x3F4040AC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN15` | 0x3F4040B0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN16` | 0x3F4040B4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN17` | 0x3F4040B8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN18` | 0x3F4040BC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN19` | 0x3F4040C0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN2` | 0x3F40407C | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN20` | 0x3F4040C4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN21` | 0x3F4040C8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN22` | 0x3F4040CC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN23` | 0x3F4040D0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN24` | 0x3F4040D4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN25` | 0x3F4040D8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN26` | 0x3F4040DC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN27` | 0x3F4040E0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN28` | 0x3F4040E4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN29` | 0x3F4040E8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN3` | 0x3F404080 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN30` | 0x3F4040EC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN31` | 0x3F4040F0 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN32` | 0x3F4040F4 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN33` | 0x3F4040F8 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN34` | 0x3F4040FC | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN35` | 0x3F404100 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN36` | 0x3F404104 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN37` | 0x3F404108 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN38` | 0x3F40410C | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN39` | 0x3F404110 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN4` | 0x3F404084 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN40` | 0x3F404114 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN41` | 0x3F404118 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN42` | 0x3F40411C | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN43` | 0x3F404120 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN44` | 0x3F404124 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN45` | 0x3F404128 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN46` | 0x3F40412C | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN47` | 0x3F404130 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN48` | 0x3F404134 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN49` | 0x3F404138 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN5` | 0x3F404088 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN50` | 0x3F40413C | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN51` | 0x3F404140 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN52` | 0x3F404144 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN53` | 0x3F404148 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN6` | 0x3F40408C | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN7` | 0x3F404090 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN8` | 0x3F404094 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_PIN9` | 0x3F404098 | 32 | CONFIG=0x1800 INT_ENA=0x3E000 INT_TYPE=0x380 PAD_DRIVER=0x4 SYNC1_BYPASS=0x18 SYNC2_BYPASS=0x3 WAKEUP_ENABLE=0x400 | - |
| `GPIO_REG_DATE` | 0x3F4046FC | 32 | DATE=0xFFFFFFF | - |
| `GPIO_SDIO_SELECT` | 0x3F40401C | 32 | SDIO_SEL=0xFF | - |
| `GPIO_SD_SIGMADELTA0` | 0x3F404F00 | 32 | SD0_IN=0xFF SD0_PRESCALE=0xFF00 | - |
| `GPIO_SD_SIGMADELTA1` | 0x3F404F04 | 32 | SD0_IN=0xFF SD0_PRESCALE=0xFF00 | - |
| `GPIO_SD_SIGMADELTA2` | 0x3F404F08 | 32 | SD0_IN=0xFF SD0_PRESCALE=0xFF00 | - |
| `GPIO_SD_SIGMADELTA3` | 0x3F404F0C | 32 | SD0_IN=0xFF SD0_PRESCALE=0xFF00 | - |
| `GPIO_SD_SIGMADELTA4` | 0x3F404F10 | 32 | SD0_IN=0xFF SD0_PRESCALE=0xFF00 | - |
| `GPIO_SD_SIGMADELTA5` | 0x3F404F14 | 32 | SD0_IN=0xFF SD0_PRESCALE=0xFF00 | - |
| `GPIO_SD_SIGMADELTA6` | 0x3F404F18 | 32 | SD0_IN=0xFF SD0_PRESCALE=0xFF00 | - |
| `GPIO_SD_SIGMADELTA7` | 0x3F404F1C | 32 | SD0_IN=0xFF SD0_PRESCALE=0xFF00 | - |
| `GPIO_SD_SIGMADELTA_CG` | 0x3F404F20 | 32 | CLK_EN=0x80000000 | - |
| `GPIO_SD_SIGMADELTA_MISC` | 0x3F404F24 | 32 | FUNCTION_CLK_EN=0x40000000 SPI_SWAP=0x80000000 | - |
| `GPIO_SD_SIGMADELTA_VERSION` | 0x3F404F28 | 32 | GPIO_SD_DATE=0xFFFFFFF | - |
| `GPIO_STATUS` | 0x3F404044 | 32 | INTERRUPT=0xFFFFFFFF | - |
| `GPIO_STATUS1` | 0x3F404050 | 32 | INTERRUPT=0x3FFFFF | - |
| `GPIO_STATUS1_W1TC` | 0x3F404058 | 32 | STATUS1_W1TC=0x3FFFFF | - |
| `GPIO_STATUS1_W1TS` | 0x3F404054 | 32 | STATUS1_W1TS=0x3FFFFF | - |
| `GPIO_STATUS_NEXT` | 0x3F40414C | 32 | STATUS_INTERRUPT_NEXT=0xFFFFFFFF | - |
| `GPIO_STATUS_NEXT1` | 0x3F404150 | 32 | STATUS1_INTERRUPT_NEXT=0x3FFFFF | - |
| `GPIO_STATUS_W1TC` | 0x3F40404C | 32 | STATUS_W1TC=0xFFFFFFFF | - |
| `GPIO_STATUS_W1TS` | 0x3F404048 | 32 | STATUS_W1TS=0xFFFFFFFF | - |
| `GPIO_STRAP` | 0x3F404038 | 32 | STRAPPING=0xFFFF | - |
| `HMAC_DATE` | 0x6003E0F8 | 32 | DATE=0x3FFFFFFF | - |
| `HMAC_ONE_BLOCK` | 0x6003E0F4 | 32 | SET_ONE_BLOCK=0x1 | - |
| `HMAC_QUERY_BUSY` | 0x6003E06C | 32 | BUSY_STATE=0x1 | - |
| `HMAC_QUERY_ERROR` | 0x6003E068 | 32 | QUREY_CHECK=0x1 | - |
| `HMAC_RD_RESULT_0` | 0x6003E0C0 | 32 | RDATA=0xFFFFFFFF | - |
| `HMAC_RD_RESULT_1` | 0x6003E0C4 | 32 | RDATA=0xFFFFFFFF | - |
| `HMAC_RD_RESULT_2` | 0x6003E0C8 | 32 | RDATA=0xFFFFFFFF | - |
| `HMAC_RD_RESULT_3` | 0x6003E0CC | 32 | RDATA=0xFFFFFFFF | - |
| `HMAC_RD_RESULT_4` | 0x6003E0D0 | 32 | RDATA=0xFFFFFFFF | - |
| `HMAC_RD_RESULT_5` | 0x6003E0D4 | 32 | RDATA=0xFFFFFFFF | - |
| `HMAC_RD_RESULT_6` | 0x6003E0D8 | 32 | RDATA=0xFFFFFFFF | - |
| `HMAC_RD_RESULT_7` | 0x6003E0DC | 32 | RDATA=0xFFFFFFFF | - |
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
| `HMAC_WR_MESSAGE_0` | 0x6003E080 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_1` | 0x6003E084 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_10` | 0x6003E0A8 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_11` | 0x6003E0AC | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_12` | 0x6003E0B0 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_13` | 0x6003E0B4 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_14` | 0x6003E0B8 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_15` | 0x6003E0BC | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_2` | 0x6003E088 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_3` | 0x6003E08C | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_4` | 0x6003E090 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_5` | 0x6003E094 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_6` | 0x6003E098 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_7` | 0x6003E09C | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_8` | 0x6003E0A0 | 32 | WDATA=0xFFFFFFFF | - |
| `HMAC_WR_MESSAGE_9` | 0x6003E0A4 | 32 | WDATA=0xFFFFFFFF | - |
| `I2C0_COMD0` | 0x3F413058 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD1` | 0x3F41305C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD10` | 0x3F413080 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD11` | 0x3F413084 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD12` | 0x3F413088 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD13` | 0x3F41308C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD14` | 0x3F413090 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD15` | 0x3F413094 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD2` | 0x3F413060 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD3` | 0x3F413064 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD4` | 0x3F413068 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD5` | 0x3F41306C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD6` | 0x3F413070 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD7` | 0x3F413074 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD8` | 0x3F413078 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_COMD9` | 0x3F41307C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C0_CTR` | 0x3F413004 | 32 | ARBITRATION_EN=0x200 CLK_EN=0x100 FSM_RST=0x400 MS_MODE=0x10 REF_ALWAYS_ON=0x800 RX_FULL_ACK_LEVEL=0x8 RX_LSB_FIRST=0x80 SAMPLE_SCL_LEVEL=0x4 SCL_FORCE_OUT=0x2 SDA_FORCE_OUT=0x1 TRANS_START=0x20 TX_LSB_FIRST=0x40 | - |
| `I2C0_DATA` | 0x3F41301C | 32 | FIFO_RDATA=0xFF | - |
| `I2C0_DATE` | 0x3F4130F8 | 32 | DATE=0xFFFFFFFF | - |
| `I2C0_FIFO_CONF` | 0x3F413018 | 32 | FIFO_ADDR_CFG_EN=0x800 FIFO_PRT_EN=0x4000000 NONFIFO_EN=0x400 NONFIFO_RX_THRES=0xFC000 NONFIFO_TX_THRES=0x3F00000 RXFIFO_WM_THRHD=0x1F RX_FIFO_RST=0x1000 TXFIFO_WM_THRHD=0x3E0 TX_FIFO_RST=0x2000 | - |
| `I2C0_FIFO_ST` | 0x3F413014 | 32 | RXFIFO_END_ADDR=0x3E0 RXFIFO_START_ADDR=0x1F RX_UPDATE=0x100000 SLAVE_RW_POINT=0x3FC00000 TXFIFO_END_ADDR=0xF8000 TXFIFO_START_ADDR=0x7C00 TX_UPDATE=0x200000 | - |
| `I2C0_INT_CLR` | 0x3F413024 | 32 | ARBITRATION_LOST_INT_CLR=0x20 BYTE_TRANS_DONE_INT_CLR=0x10 DET_START_INT_CLR=0x8000 END_DETECT_INT_CLR=0x8 MST_TXFIFO_UDF_INT_CLR=0x40 NACK_INT_CLR=0x400 RXFIFO_OVF_INT_CLR=0x4 RXFIFO_UDF_INT_CLR=0x1000 RXFIFO_WM_INT_CLR=0x1 SCL_MAIN_ST_TO_... | - |
| `I2C0_INT_ENA` | 0x3F413028 | 32 | ARBITRATION_LOST_INT_ENA=0x20 BYTE_TRANS_DONE_INT_ENA=0x10 DET_START_INT_ENA=0x8000 END_DETECT_INT_ENA=0x8 MST_TXFIFO_UDF_INT_ENA=0x40 NACK_INT_ENA=0x400 RXFIFO_OVF_INT_ENA=0x4 RXFIFO_UDF_INT_ENA=0x1000 RXFIFO_WM_INT_ENA=0x1 SCL_MAIN_ST_TO_... | - |
| `I2C0_INT_RAW` | 0x3F413020 | 32 | ARBITRATION_LOST_INT_RAW=0x20 BYTE_TRANS_DONE_INT_RAW=0x10 DET_START_INT_RAW=0x8000 END_DETECT_INT_RAW=0x8 MST_TXFIFO_UDF_INT_RAW=0x40 NACK_INT_RAW=0x400 RXFIFO_OVF_INT_RAW=0x4 RXFIFO_UDF_INT_RAW=0x1000 RXFIFO_WM_INT_RAW=0x1 SCL_MAIN_ST_TO_... | - |
| `I2C0_INT_STATUS` | 0x3F41302C | 32 | ARBITRATION_LOST_INT_ST=0x20 BYTE_TRANS_DONE_INT_ST=0x10 DET_START_INT_ST=0x8000 END_DETECT_INT_ST=0x8 MST_TXFIFO_UDF_INT_ST=0x40 NACK_INT_ST=0x400 RXFIFO_OVF_INT_ST=0x4 RXFIFO_UDF_INT_ST=0x1000 RXFIFO_WM_INT_ST=0x1 SCL_MAIN_ST_TO_INT_ST=0x... | - |
| `I2C0_SCL_FILTER_CFG` | 0x3F413050 | 32 | SCL_FILTER_EN=0x10 SCL_FILTER_THRES=0xF | - |
| `I2C0_SCL_HIGH_PERIOD` | 0x3F413038 | 32 | SCL_HIGH_PERIOD=0x3FFF SCL_WAIT_HIGH_PERIOD=0xFFFC000 | - |
| `I2C0_SCL_LOW_PERIOD` | 0x3F413000 | 32 | SCL_LOW_PERIOD=0x3FFF | - |
| `I2C0_SCL_MAIN_ST_TIME_OUT` | 0x3F41309C | 32 | SCL_MAIN_ST_TO=0xFFFFFF | - |
| `I2C0_SCL_RSTART_SETUP` | 0x3F413044 | 32 | TIME=0x3FF | - |
| `I2C0_SCL_SP_CONF` | 0x3F4130A0 | 32 | SCL_PD_EN=0x40 SCL_RST_SLV_EN=0x1 SCL_RST_SLV_NUM=0x3E SDA_PD_EN=0x80 | - |
| `I2C0_SCL_START_HOLD` | 0x3F413040 | 32 | TIME=0x3FF | - |
| `I2C0_SCL_STOP_HOLD` | 0x3F413048 | 32 | TIME=0x3FFF | - |
| `I2C0_SCL_STOP_SETUP` | 0x3F41304C | 32 | TIME=0x3FF | - |
| `I2C0_SCL_STRETCH_CONF` | 0x3F4130A4 | 32 | SLAVE_SCL_STRETCH_CLR=0x800 SLAVE_SCL_STRETCH_EN=0x400 STRETCH_PROTECT_NUM=0x3FF | - |
| `I2C0_SCL_ST_TIME_OUT` | 0x3F413098 | 32 | SCL_ST_TO=0xFFFFFF | - |
| `I2C0_SDA_FILTER_CFG` | 0x3F413054 | 32 | SDA_FILTER_EN=0x10 SDA_FILTER_THRES=0xF | - |
| `I2C0_SDA_HOLD` | 0x3F413030 | 32 | TIME=0x3FF | - |
| `I2C0_SDA_SAMPLE` | 0x3F413034 | 32 | TIME=0x3FF | - |
| `I2C0_SLAVE_ADDR` | 0x3F413010 | 32 | ADDR_10BIT_EN=0x80000000 SLAVE_ADDR=0x7FFF | - |
| `I2C0_SR` | 0x3F413008 | 32 | ARB_LOST=0x8 BUS_BUSY=0x10 BYTE_TRANS=0x40 RESP_REC=0x1 RXFIFO_CNT=0x3F00 SCL_MAIN_STATE_LAST=0x7000000 SCL_STATE_LAST=0x70000000 SLAVE_ADDRESSED=0x20 SLAVE_RW=0x2 STRETCH_CAUSE=0xC000 TIME_OUT=0x4 TXFIFO_CNT=0xFC0000 | - |
| `I2C0_TO` | 0x3F41300C | 32 | TIME_OUT_EN=0x1000000 TIME_OUT_VALUE=0xFFFFFF | - |
| `I2C1_COMD0` | 0x3F427058 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD1` | 0x3F42705C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD10` | 0x3F427080 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD11` | 0x3F427084 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD12` | 0x3F427088 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD13` | 0x3F42708C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD14` | 0x3F427090 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD15` | 0x3F427094 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD2` | 0x3F427060 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD3` | 0x3F427064 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD4` | 0x3F427068 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD5` | 0x3F42706C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD6` | 0x3F427070 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD7` | 0x3F427074 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD8` | 0x3F427078 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_COMD9` | 0x3F42707C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `I2C1_CTR` | 0x3F427004 | 32 | ARBITRATION_EN=0x200 CLK_EN=0x100 FSM_RST=0x400 MS_MODE=0x10 REF_ALWAYS_ON=0x800 RX_FULL_ACK_LEVEL=0x8 RX_LSB_FIRST=0x80 SAMPLE_SCL_LEVEL=0x4 SCL_FORCE_OUT=0x2 SDA_FORCE_OUT=0x1 TRANS_START=0x20 TX_LSB_FIRST=0x40 | - |
| `I2C1_DATA` | 0x3F42701C | 32 | FIFO_RDATA=0xFF | - |
| `I2C1_DATE` | 0x3F4270F8 | 32 | DATE=0xFFFFFFFF | - |
| `I2C1_FIFO_CONF` | 0x3F427018 | 32 | FIFO_ADDR_CFG_EN=0x800 FIFO_PRT_EN=0x4000000 NONFIFO_EN=0x400 NONFIFO_RX_THRES=0xFC000 NONFIFO_TX_THRES=0x3F00000 RXFIFO_WM_THRHD=0x1F RX_FIFO_RST=0x1000 TXFIFO_WM_THRHD=0x3E0 TX_FIFO_RST=0x2000 | - |
| `I2C1_FIFO_ST` | 0x3F427014 | 32 | RXFIFO_END_ADDR=0x3E0 RXFIFO_START_ADDR=0x1F RX_UPDATE=0x100000 SLAVE_RW_POINT=0x3FC00000 TXFIFO_END_ADDR=0xF8000 TXFIFO_START_ADDR=0x7C00 TX_UPDATE=0x200000 | - |
| `I2C1_INT_CLR` | 0x3F427024 | 32 | ARBITRATION_LOST_INT_CLR=0x20 BYTE_TRANS_DONE_INT_CLR=0x10 DET_START_INT_CLR=0x8000 END_DETECT_INT_CLR=0x8 MST_TXFIFO_UDF_INT_CLR=0x40 NACK_INT_CLR=0x400 RXFIFO_OVF_INT_CLR=0x4 RXFIFO_UDF_INT_CLR=0x1000 RXFIFO_WM_INT_CLR=0x1 SCL_MAIN_ST_TO_... | - |
| `I2C1_INT_ENA` | 0x3F427028 | 32 | ARBITRATION_LOST_INT_ENA=0x20 BYTE_TRANS_DONE_INT_ENA=0x10 DET_START_INT_ENA=0x8000 END_DETECT_INT_ENA=0x8 MST_TXFIFO_UDF_INT_ENA=0x40 NACK_INT_ENA=0x400 RXFIFO_OVF_INT_ENA=0x4 RXFIFO_UDF_INT_ENA=0x1000 RXFIFO_WM_INT_ENA=0x1 SCL_MAIN_ST_TO_... | - |
| `I2C1_INT_RAW` | 0x3F427020 | 32 | ARBITRATION_LOST_INT_RAW=0x20 BYTE_TRANS_DONE_INT_RAW=0x10 DET_START_INT_RAW=0x8000 END_DETECT_INT_RAW=0x8 MST_TXFIFO_UDF_INT_RAW=0x40 NACK_INT_RAW=0x400 RXFIFO_OVF_INT_RAW=0x4 RXFIFO_UDF_INT_RAW=0x1000 RXFIFO_WM_INT_RAW=0x1 SCL_MAIN_ST_TO_... | - |
| `I2C1_INT_STATUS` | 0x3F42702C | 32 | ARBITRATION_LOST_INT_ST=0x20 BYTE_TRANS_DONE_INT_ST=0x10 DET_START_INT_ST=0x8000 END_DETECT_INT_ST=0x8 MST_TXFIFO_UDF_INT_ST=0x40 NACK_INT_ST=0x400 RXFIFO_OVF_INT_ST=0x4 RXFIFO_UDF_INT_ST=0x1000 RXFIFO_WM_INT_ST=0x1 SCL_MAIN_ST_TO_INT_ST=0x... | - |
| `I2C1_SCL_FILTER_CFG` | 0x3F427050 | 32 | SCL_FILTER_EN=0x10 SCL_FILTER_THRES=0xF | - |
| `I2C1_SCL_HIGH_PERIOD` | 0x3F427038 | 32 | SCL_HIGH_PERIOD=0x3FFF SCL_WAIT_HIGH_PERIOD=0xFFFC000 | - |
| `I2C1_SCL_LOW_PERIOD` | 0x3F427000 | 32 | SCL_LOW_PERIOD=0x3FFF | - |
| `I2C1_SCL_MAIN_ST_TIME_OUT` | 0x3F42709C | 32 | SCL_MAIN_ST_TO=0xFFFFFF | - |
| `I2C1_SCL_RSTART_SETUP` | 0x3F427044 | 32 | TIME=0x3FF | - |
| `I2C1_SCL_SP_CONF` | 0x3F4270A0 | 32 | SCL_PD_EN=0x40 SCL_RST_SLV_EN=0x1 SCL_RST_SLV_NUM=0x3E SDA_PD_EN=0x80 | - |
| `I2C1_SCL_START_HOLD` | 0x3F427040 | 32 | TIME=0x3FF | - |
| `I2C1_SCL_STOP_HOLD` | 0x3F427048 | 32 | TIME=0x3FFF | - |
| `I2C1_SCL_STOP_SETUP` | 0x3F42704C | 32 | TIME=0x3FF | - |
| `I2C1_SCL_STRETCH_CONF` | 0x3F4270A4 | 32 | SLAVE_SCL_STRETCH_CLR=0x800 SLAVE_SCL_STRETCH_EN=0x400 STRETCH_PROTECT_NUM=0x3FF | - |
| `I2C1_SCL_ST_TIME_OUT` | 0x3F427098 | 32 | SCL_ST_TO=0xFFFFFF | - |
| `I2C1_SDA_FILTER_CFG` | 0x3F427054 | 32 | SDA_FILTER_EN=0x10 SDA_FILTER_THRES=0xF | - |
| `I2C1_SDA_HOLD` | 0x3F427030 | 32 | TIME=0x3FF | - |
| `I2C1_SDA_SAMPLE` | 0x3F427034 | 32 | TIME=0x3FF | - |
| `I2C1_SLAVE_ADDR` | 0x3F427010 | 32 | ADDR_10BIT_EN=0x80000000 SLAVE_ADDR=0x7FFF | - |
| `I2C1_SR` | 0x3F427008 | 32 | ARB_LOST=0x8 BUS_BUSY=0x10 BYTE_TRANS=0x40 RESP_REC=0x1 RXFIFO_CNT=0x3F00 SCL_MAIN_STATE_LAST=0x7000000 SCL_STATE_LAST=0x70000000 SLAVE_ADDRESSED=0x20 SLAVE_RW=0x2 STRETCH_CAUSE=0xC000 TIME_OUT=0x4 TXFIFO_CNT=0xFC0000 | - |
| `I2C1_TO` | 0x3F42700C | 32 | TIME_OUT_EN=0x1000000 TIME_OUT_VALUE=0xFFFFFF | - |
| `I2S_CLKM_CONF` | 0x3F40F0AC | 32 | CLKM_DIV_A=0xFC000 CLKM_DIV_B=0x3F00 CLKM_DIV_NUM=0xFF CLK_EN=0x100000 CLK_SEL=0x600000 | - |
| `I2S_CONF` | 0x3F40F008 | 32 | PRE_REQ_EN=0x4000000 RX_BIG_ENDIAN=0x10000000 RX_DMA_EQUAL=0x2000000 RX_FIFO_RESET=0x8 RX_FIFO_RESET_ST=0x400000 RX_LSB_FIRST_DMA=0x80000 RX_MONO=0x8000 RX_MSB_RIGHT=0x20000 RX_MSB_SHIFT=0x800 RX_RESET=0x2 RX_RESET_ST=0x20000000 RX_RIGHT_FI... | - |
| `I2S_CONF1` | 0x3F40F0A0 | 32 | RX_PCM_BYPASS=0x80 RX_PCM_CONF=0x70 TX_PCM_BYPASS=0x8 TX_PCM_CONF=0x7 TX_STOP_EN=0x100 TX_ZEROS_RM_EN=0x200 | - |
| `I2S_CONF2` | 0x3F40F0A8 | 32 | CAMERA_EN=0x1 CAM_CLK_LOOPBACK=0x200 CAM_SYNC_FIFO_RESET=0x100 DATA_ENABLE=0x10 DATA_ENABLE_TEST_EN=0x8 EXT_ADC_START_EN=0x40 INTER_VALID_EN=0x80 LCD_EN=0x20 LCD_TX_SDX2_EN=0x4 LCD_TX_WRX2_EN=0x2 VSYNC_FILTER_EN=0x400 VSYNC_FILTER_THRES=0x3... | - |
| `I2S_CONF_CHAN` | 0x3F40F02C | 32 | RX_CHAN_MOD=0x18 TX_CHAN_MOD=0x7 | - |
| `I2S_CONF_SIGLE_DATA` | 0x3F40F028 | 32 | SIGLE_DATA=0xFFFFFFFF | - |
| `I2S_DATE` | 0x3F40F0FC | 32 | DATE=0xFFFFFFFF | - |
| `I2S_FIFO_CONF` | 0x3F40F020 | 32 | DSCR_EN=0x1000 RX_24MSB_EN=0x400000 RX_DATA_NUM=0x3F RX_FIFO_MOD=0x70000 RX_FIFO_MOD_FORCE_EN=0x100000 RX_FIFO_SYNC=0x200000 TX_24MSB_EN=0x800000 TX_DATA_NUM=0xFC0 TX_FIFO_MOD=0xE000 TX_FIFO_MOD_FORCE_EN=0x80000 | - |
| `I2S_INFIFO_POP` | 0x3F40F068 | 32 | INFIFO_POP=0x10000 INFIFO_RDATA=0xFFF | - |
| `I2S_INLINK_DSCR` | 0x3F40F048 | 32 | INLINK_DSCR=0xFFFFFFFF | - |
| `I2S_INLINK_DSCR_BF0` | 0x3F40F04C | 32 | INLINK_DSCR_BF0=0xFFFFFFFF | - |
| `I2S_INLINK_DSCR_BF1` | 0x3F40F050 | 32 | INLINK_DSCR_BF1=0xFFFFFFFF | - |
| `I2S_INT_CLR` | 0x3F40F018 | 32 | IN_DONE_INT_CLR=0x100 IN_DSCR_EMPTY_INT_CLR=0x8000 IN_DSCR_ERR_INT_CLR=0x2000 IN_ERR_EOF_INT_CLR=0x400 IN_SUC_EOF_INT_CLR=0x200 OUT_DONE_INT_CLR=0x800 OUT_DSCR_ERR_INT_CLR=0x4000 OUT_EOF_INT_CLR=0x1000 OUT_TOTAL_EOF_INT_CLR=0x10000 PUT_DATA... | - |
| `I2S_INT_ENA` | 0x3F40F014 | 32 | IN_DONE_INT_ENA=0x100 IN_DSCR_EMPTY_INT_ENA=0x8000 IN_DSCR_ERR_INT_ENA=0x2000 IN_ERR_EOF_INT_ENA=0x400 IN_SUC_EOF_INT_ENA=0x200 OUT_DONE_INT_ENA=0x800 OUT_DSCR_ERR_INT_ENA=0x4000 OUT_EOF_INT_ENA=0x1000 OUT_TOTAL_EOF_INT_ENA=0x10000 RX_HUNG_... | - |
| `I2S_INT_RAW` | 0x3F40F00C | 32 | IN_DONE_INT_RAW=0x100 IN_DSCR_EMPTY_INT_RAW=0x8000 IN_DSCR_ERR_INT_RAW=0x2000 IN_ERR_EOF_INT_RAW=0x400 IN_SUC_EOF_INT_RAW=0x200 OUT_DONE_INT_RAW=0x800 OUT_DSCR_ERR_INT_RAW=0x4000 OUT_EOF_INT_RAW=0x1000 OUT_TOTAL_EOF_INT_RAW=0x10000 RX_HUNG_... | - |
| `I2S_INT_ST` | 0x3F40F010 | 32 | IN_DONE_INT_ST=0x100 IN_DSCR_EMPTY_INT_ST=0x8000 IN_DSCR_ERR_INT_ST=0x2000 IN_ERR_EOF_INT_ST=0x400 IN_SUC_EOF_INT_ST=0x200 OUT_DONE_INT_ST=0x800 OUT_DSCR_ERR_INT_ST=0x4000 OUT_EOF_INT_ST=0x1000 OUT_TOTAL_EOF_INT_ST=0x10000 RX_HUNG_INT_ST=0x... | - |
| `I2S_IN_EOF_DES_ADDR` | 0x3F40F03C | 32 | IN_SUC_EOF_DES_ADDR=0xFFFFFFFF | - |
| `I2S_IN_LINK` | 0x3F40F034 | 32 | INLINK_ADDR=0xFFFFF INLINK_PARK=0x80000000 INLINK_RESTART=0x40000000 INLINK_START=0x20000000 INLINK_STOP=0x10000000 | - |
| `I2S_LC_CONF` | 0x3F40F060 | 32 | AHBM_FIFO_RST=0x4 AHBM_RST=0x8 CHECK_OWNER=0x1000 EXT_MEM_BK_SIZE=0xC000 INDSCR_BURST_EN=0x400 IN_LOOP_TEST=0x20 IN_RST=0x1 MEM_TRANS_EN=0x2000 OUTDSCR_BURST_EN=0x200 OUT_AUTO_WRBACK=0x40 OUT_DATA_BURST_EN=0x800 OUT_EOF_MODE=0x100 OUT_LOOP_... | - |
| `I2S_LC_HUNG_CONF` | 0x3F40F074 | 32 | LC_FIFO_TIMEOUT=0xFF LC_FIFO_TIMEOUT_ENA=0x800 LC_FIFO_TIMEOUT_SHIFT=0x700 | - |
| `I2S_LC_STATE0` | 0x3F40F06C | 32 | OUTFIFO_CNT=0x3F800000 OUTLINK_DSCR_ADDR=0x3FFFF OUT_DSCR_STATE=0xC0000 OUT_EMPTY=0x80000000 OUT_FULL=0x40000000 OUT_STATE=0x700000 | - |
| `I2S_LC_STATE1` | 0x3F40F070 | 32 | INFIFO_CNT_DEBUG=0x3F800000 INLINK_DSCR_ADDR=0x3FFFF IN_DSCR_STATE=0xC0000 IN_EMPTY=0x80000000 IN_FULL=0x40000000 IN_STATE=0x700000 | - |
| `I2S_OUTFIFO_PUSH` | 0x3F40F064 | 32 | OUTFIFO_PUSH=0x10000 OUTFIFO_WDATA=0x1FF | - |
| `I2S_OUTLINK_DSCR` | 0x3F40F054 | 32 | OUTLINK_DSCR=0xFFFFFFFF | - |
| `I2S_OUTLINK_DSCR_BF0` | 0x3F40F058 | 32 | OUTLINK_DSCR_BF0=0xFFFFFFFF | - |
| `I2S_OUTLINK_DSCR_BF1` | 0x3F40F05C | 32 | OUTLINK_DSCR_BF1=0xFFFFFFFF | - |
| `I2S_OUT_EOF_BFR_DES_ADDR` | 0x3F40F040 | 32 | OUT_EOF_BFR_DES_ADDR=0xFFFFFFFF | - |
| `I2S_OUT_EOF_DES_ADDR` | 0x3F40F038 | 32 | OUT_EOF_DES_ADDR=0xFFFFFFFF | - |
| `I2S_OUT_LINK` | 0x3F40F030 | 32 | OUTLINK_ADDR=0xFFFFF OUTLINK_PARK=0x80000000 OUTLINK_RESTART=0x40000000 OUTLINK_START=0x20000000 OUTLINK_STOP=0x10000000 | - |
| `I2S_PD_CONF` | 0x3F40F0A4 | 32 | DMA_RAM_CLK_FO=0x40 DMA_RAM_FORCE_PD=0x10 DMA_RAM_FORCE_PU=0x20 FIFO_FORCE_PD=0x1 FIFO_FORCE_PU=0x2 PLC_MEM_FORCE_PD=0x4 PLC_MEM_FORCE_PU=0x8 | - |
| `I2S_RXEOF_NUM` | 0x3F40F024 | 32 | RX_EOF_NUM=0xFFFFFFFF | - |
| `I2S_SAMPLE_RATE_CONF` | 0x3F40F0B0 | 32 | RX_BCK_DIV_NUM=0xFC0 RX_BITS_MOD=0xFC0000 TX_BCK_DIV_NUM=0x3F TX_BITS_MOD=0x3F000 | - |
| `I2S_STATE` | 0x3F40F0BC | 32 | TX_IDLE=0x1 | - |
| `I2S_TIMING` | 0x3F40F01C | 32 | DATA_ENABLE_DELAY=0xC00000 RX_BCK_IN_DELAY=0x30 RX_BCK_OUT_DELAY=0xC0000 RX_DSYNC_SW=0x200000 RX_SD_IN_DELAY=0x300 RX_WS_IN_DELAY=0xC0 RX_WS_OUT_DELAY=0x30000 TX_BCK_IN_DELAY=0x3 TX_BCK_IN_INV=0x1000000 TX_BCK_OUT_DELAY=0xC00 TX_DSYNC_SW=0x... | - |
| `INTERRUPT_CLOCK_GATE` | 0x3F4C2188 | 32 | CLK_EN=0x1 PRO_NMI_MASK_HW=0x2 | - |
| `INTERRUPT_PRO_AES_INTR_MAP` | 0x3F4C20E0 | 32 | PRO_AES_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_APB_ADC_INT_MAP` | 0x3F4C2164 | 32 | PRO_APB_ADC_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_APB_PERI_ERROR_INT_MAP` | 0x3F4C2170 | 32 | PRO_APB_PERI_ERROR_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_ASSIST_DEBUG_INTR_MAP` | 0x3F4C2128 | 32 | PRO_ASSIST_DEBUG_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_BB_INT_MAP` | 0x3F4C200C | 32 | PRO_BB_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_BT_BB_INT_MAP` | 0x3F4C2014 | 32 | PRO_BT_BB_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_BT_BB_NMI_MAP` | 0x3F4C2018 | 32 | PRO_BT_BB_NMI_MAP=0x1F | - |
| `INTERRUPT_PRO_BT_MAC_INT_MAP` | 0x3F4C2010 | 32 | PRO_BT_MAC_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_CACHE_IA_INT_MAP` | 0x3F4C2118 | 32 | PRO_CACHE_IA_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_CAN_INT_MAP` | 0x3F4C20BC | 32 | PRO_CAN_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_CPU_INTR_FROM_CPU_0_MAP` | 0x3F4C2070 | 32 | PRO_CPU_INTR_FROM_CPU_0_MAP=0x1F | - |
| `INTERRUPT_PRO_CPU_INTR_FROM_CPU_1_MAP` | 0x3F4C2074 | 32 | PRO_CPU_INTR_FROM_CPU_1_MAP=0x1F | - |
| `INTERRUPT_PRO_CPU_INTR_FROM_CPU_2_MAP` | 0x3F4C2078 | 32 | PRO_CPU_INTR_FROM_CPU_2_MAP=0x1F | - |
| `INTERRUPT_PRO_CPU_INTR_FROM_CPU_3_MAP` | 0x3F4C207C | 32 | PRO_CPU_INTR_FROM_CPU_3_MAP=0x1F | - |
| `INTERRUPT_PRO_CPU_PERI_ERROR_INT_MAP` | 0x3F4C216C | 32 | PRO_CPU_PERI_ERROR_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_CRYPTO_DMA_INT_MAP` | 0x3F4C2168 | 32 | PRO_CRYPTO_DMA_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_DCACHE_PRELOAD_INT_MAP` | 0x3F4C215C | 32 | PRO_DCACHE_PRELOAD_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_DCACHE_SYNC_INT_MAP` | 0x3F4C2174 | 32 | PRO_DCACHE_SYNC_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_DEDICATED_GPIO_IN_INTR_MAP` | 0x3F4C206C | 32 | PRO_DEDICATED_GPIO_IN_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_DMA_COPY_INTR_MAP` | 0x3F4C2150 | 32 | PRO_DMA_COPY_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_EFUSE_INT_MAP` | 0x3F4C20B8 | 32 | PRO_EFUSE_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_GPIO_INTERRUPT_APP_MAP` | 0x3F4C2064 | 32 | PRO_GPIO_INTERRUPT_APP_MAP=0x1F | - |
| `INTERRUPT_PRO_GPIO_INTERRUPT_APP_NMI_MAP` | 0x3F4C2068 | 32 | PRO_GPIO_INTERRUPT_APP_NMI_MAP=0x1F | - |
| `INTERRUPT_PRO_GPIO_INTERRUPT_PRO_MAP` | 0x3F4C205C | 32 | PRO_GPIO_INTERRUPT_PRO_MAP=0x1F | - |
| `INTERRUPT_PRO_GPIO_INTERRUPT_PRO_NMI_MAP` | 0x3F4C2060 | 32 | PRO_GPIO_INTERRUPT_PRO_NMI_MAP=0x1F | - |
| `INTERRUPT_PRO_I2C_EXT0_INTR_MAP` | 0x3F4C20D0 | 32 | PRO_I2C_EXT0_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_I2C_EXT1_INTR_MAP` | 0x3F4C20D4 | 32 | PRO_I2C_EXT1_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_I2S0_INT_MAP` | 0x3F4C208C | 32 | PRO_I2S0_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_I2S1_INT_MAP` | 0x3F4C2090 | 32 | PRO_I2S1_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_ICACHE_PRELOAD_INT_MAP` | 0x3F4C2160 | 32 | PRO_ICACHE_PRELOAD_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_ICACHE_SYNC_INT_MAP` | 0x3F4C2178 | 32 | PRO_ICACHE_SYNC_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_INTR_STATUS_0` | 0x3F4C217C | 32 | PRO_INTR_STATUS_0=0xFFFFFFFF | - |
| `INTERRUPT_PRO_INTR_STATUS_1` | 0x3F4C2180 | 32 | PRO_INTR_STATUS_1=0xFFFFFFFF | - |
| `INTERRUPT_PRO_INTR_STATUS_2` | 0x3F4C2184 | 32 | PRO_INTR_STATUS_2=0xFFFFFFFF | - |
| `INTERRUPT_PRO_LEDC_INT_MAP` | 0x3F4C20B4 | 32 | PRO_LEDC_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_MAC_INTR_MAP` | 0x3F4C2000 | 32 | PRO_MAC_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_MAC_NMI_MAP` | 0x3F4C2004 | 32 | PRO_MAC_NMI_MAP=0x1F | - |
| `INTERRUPT_PRO_PCNT_INTR_MAP` | 0x3F4C20CC | 32 | PRO_PCNT_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PMS_DMA_APB_I_ILG_INTR_MAP` | 0x3F4C2140 | 32 | PRO_PMS_DMA_APB_I_ILG_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PMS_DMA_RX_I_ILG_INTR_MAP` | 0x3F4C2144 | 32 | PRO_PMS_DMA_RX_I_ILG_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PMS_DMA_TX_I_ILG_INTR_MAP` | 0x3F4C2148 | 32 | PRO_PMS_DMA_TX_I_ILG_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PMS_PRO_AHB_ILG_INTR_MAP` | 0x3F4C2138 | 32 | PRO_PMS_PRO_AHB_ILG_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PMS_PRO_CACHE_ILG_INTR_MAP` | 0x3F4C213C | 32 | PRO_PMS_PRO_CACHE_ILG_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PMS_PRO_DPORT_ILG_INTR_MAP` | 0x3F4C2134 | 32 | PRO_PMS_PRO_DPORT_ILG_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PMS_PRO_DRAM0_ILG_INTR_MAP` | 0x3F4C2130 | 32 | PRO_PMS_PRO_DRAM0_ILG_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PMS_PRO_IRAM0_ILG_INTR_MAP` | 0x3F4C212C | 32 | PRO_PMS_PRO_IRAM0_ILG_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PWM0_INTR_MAP` | 0x3F4C20A4 | 32 | PRO_PWM0_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PWM1_INTR_MAP` | 0x3F4C20A8 | 32 | PRO_PWM1_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PWM2_INTR_MAP` | 0x3F4C20AC | 32 | PRO_PWM2_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PWM3_INTR_MAP` | 0x3F4C20B0 | 32 | PRO_PWM3_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_PWR_INTR_MAP` | 0x3F4C2008 | 32 | PRO_PWR_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_RMT_INTR_MAP` | 0x3F4C20C8 | 32 | PRO_RMT_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_RSA_INTR_MAP` | 0x3F4C20D8 | 32 | PRO_RSA_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_RTC_CORE_INTR_MAP` | 0x3F4C20C4 | 32 | PRO_RTC_CORE_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_RWBLE_IRQ_MAP` | 0x3F4C2020 | 32 | PRO_RWBLE_IRQ_MAP=0x1F | - |
| `INTERRUPT_PRO_RWBLE_NMI_MAP` | 0x3F4C2028 | 32 | PRO_RWBLE_NMI_MAP=0x1F | - |
| `INTERRUPT_PRO_RWBT_IRQ_MAP` | 0x3F4C201C | 32 | PRO_RWBT_IRQ_MAP=0x1F | - |
| `INTERRUPT_PRO_RWBT_NMI_MAP` | 0x3F4C2024 | 32 | PRO_RWBT_NMI_MAP=0x1F | - |
| `INTERRUPT_PRO_SDIO_HOST_INTERRUPT_MAP` | 0x3F4C20A0 | 32 | PRO_SDIO_HOST_INTERRUPT_MAP=0x1F | - |
| `INTERRUPT_PRO_SHA_INTR_MAP` | 0x3F4C20DC | 32 | PRO_SHA_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_SLC0_INTR_MAP` | 0x3F4C202C | 32 | PRO_SLC0_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_SLC1_INTR_MAP` | 0x3F4C2030 | 32 | PRO_SLC1_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_SPI2_DMA_INT_MAP` | 0x3F4C20E4 | 32 | PRO_SPI2_DMA_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_SPI3_DMA_INT_MAP` | 0x3F4C20E8 | 32 | PRO_SPI3_DMA_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_SPI4_DMA_INT_MAP` | 0x3F4C2154 | 32 | PRO_SPI4_DMA_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_SPI_INTR_1_MAP` | 0x3F4C2080 | 32 | PRO_SPI_INTR_1_MAP=0x1F | - |
| `INTERRUPT_PRO_SPI_INTR_2_MAP` | 0x3F4C2084 | 32 | PRO_SPI_INTR_2_MAP=0x1F | - |
| `INTERRUPT_PRO_SPI_INTR_3_MAP` | 0x3F4C2088 | 32 | PRO_SPI_INTR_3_MAP=0x1F | - |
| `INTERRUPT_PRO_SPI_INTR_4_MAP` | 0x3F4C2158 | 32 | PRO_SPI_INTR_4_MAP=0x1F | - |
| `INTERRUPT_PRO_SPI_MEM_REJECT_INTR_MAP` | 0x3F4C214C | 32 | PRO_SPI_MEM_REJECT_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_SYSTIMER_TARGET0_INT_MAP` | 0x3F4C211C | 32 | PRO_SYSTIMER_TARGET0_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_SYSTIMER_TARGET1_INT_MAP` | 0x3F4C2120 | 32 | PRO_SYSTIMER_TARGET1_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_SYSTIMER_TARGET2_INT_MAP` | 0x3F4C2124 | 32 | PRO_SYSTIMER_TARGET2_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG1_LACT_EDGE_INT_MAP` | 0x3F4C2114 | 32 | PRO_TG1_LACT_EDGE_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG1_LACT_LEVEL_INT_MAP` | 0x3F4C2058 | 32 | PRO_TG1_LACT_LEVEL_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG1_T0_EDGE_INT_MAP` | 0x3F4C2108 | 32 | PRO_TG1_T0_EDGE_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG1_T0_LEVEL_INT_MAP` | 0x3F4C204C | 32 | PRO_TG1_T0_LEVEL_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG1_T1_EDGE_INT_MAP` | 0x3F4C210C | 32 | PRO_TG1_T1_EDGE_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG1_T1_LEVEL_INT_MAP` | 0x3F4C2050 | 32 | PRO_TG1_T1_LEVEL_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG1_WDT_EDGE_INT_MAP` | 0x3F4C2110 | 32 | PRO_TG1_WDT_EDGE_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG1_WDT_LEVEL_INT_MAP` | 0x3F4C2054 | 32 | PRO_TG1_WDT_LEVEL_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG_LACT_EDGE_INT_MAP` | 0x3F4C2104 | 32 | PRO_TG_LACT_EDGE_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG_LACT_LEVEL_INT_MAP` | 0x3F4C2048 | 32 | PRO_TG_LACT_LEVEL_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG_T0_EDGE_INT_MAP` | 0x3F4C20F8 | 32 | PRO_TG_T0_EDGE_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG_T0_LEVEL_INT_MAP` | 0x3F4C203C | 32 | PRO_TG_T0_LEVEL_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG_T1_EDGE_INT_MAP` | 0x3F4C20FC | 32 | PRO_TG_T1_EDGE_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG_T1_LEVEL_INT_MAP` | 0x3F4C2040 | 32 | PRO_TG_T1_LEVEL_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG_WDT_EDGE_INT_MAP` | 0x3F4C2100 | 32 | PRO_TG_WDT_EDGE_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TG_WDT_LEVEL_INT_MAP` | 0x3F4C2044 | 32 | PRO_TG_WDT_LEVEL_INT_MAP=0x1F | - |
| `INTERRUPT_PRO_TIMER_INT1_MAP` | 0x3F4C20F0 | 32 | PRO_TIMER_INT1_MAP=0x1F | - |
| `INTERRUPT_PRO_TIMER_INT2_MAP` | 0x3F4C20F4 | 32 | PRO_TIMER_INT2_MAP=0x1F | - |
| `INTERRUPT_PRO_UART1_INTR_MAP` | 0x3F4C2098 | 32 | PRO_UART1_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_UART2_INTR_MAP` | 0x3F4C209C | 32 | PRO_UART2_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_UART_INTR_MAP` | 0x3F4C2094 | 32 | PRO_UART_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_UHCI0_INTR_MAP` | 0x3F4C2034 | 32 | PRO_UHCI0_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_UHCI1_INTR_MAP` | 0x3F4C2038 | 32 | PRO_UHCI1_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_USB_INTR_MAP` | 0x3F4C20C0 | 32 | PRO_USB_INTR_MAP=0x1F | - |
| `INTERRUPT_PRO_WDG_INT_MAP` | 0x3F4C20EC | 32 | PRO_WDG_INT_MAP=0x1F | - |
| `INTERRUPT_REG_DATE` | 0x3F4C2FFC | 32 | INTERRUPT_REG_DATE=0xFFFFFFF | - |
| `IO_MUX_DATE` | 0x3F4090FC | 32 | VERSION=0xFFFFFFF | - |
| `IO_MUX_GPIO0` | 0x3F409004 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO1` | 0x3F409008 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO10` | 0x3F40902C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO11` | 0x3F409030 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO12` | 0x3F409034 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO13` | 0x3F409038 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO14` | 0x3F40903C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO15` | 0x3F409040 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO16` | 0x3F409044 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO17` | 0x3F409048 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO18` | 0x3F40904C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO19` | 0x3F409050 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO2` | 0x3F40900C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO20` | 0x3F409054 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO21` | 0x3F409058 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO26` | 0x3F40906C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO27` | 0x3F409070 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO28` | 0x3F409074 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO29` | 0x3F409078 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO3` | 0x3F409010 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO30` | 0x3F40907C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO31` | 0x3F409080 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO32` | 0x3F409084 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO33` | 0x3F409088 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO34` | 0x3F40908C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO35` | 0x3F409090 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO36` | 0x3F409094 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO37` | 0x3F409098 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO38` | 0x3F40909C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO39` | 0x3F4090A0 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO4` | 0x3F409014 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO40` | 0x3F4090A4 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO41` | 0x3F4090A8 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO42` | 0x3F4090AC | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO43` | 0x3F4090B0 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO44` | 0x3F4090B4 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO45` | 0x3F4090B8 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO46` | 0x3F4090BC | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO5` | 0x3F409018 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO6` | 0x3F40901C | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO7` | 0x3F409020 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO8` | 0x3F409024 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_GPIO9` | 0x3F409028 | 32 | FILTER_EN=0x8000 FUN_DRV=0xC00 FUN_IE=0x200 FUN_WPD=0x80 FUN_WPU=0x100 MCU_IE=0x10 MCU_OE=0x1 MCU_SEL=0x7000 MCU_WPD=0x4 MCU_WPU=0x8 SLP_SEL=0x2 | - |
| `IO_MUX_PIN_CTRL` | 0x3F409000 | 32 | PAD_POWER_CTRL=0x8000 PIN_CLK_OUT1=0xF PIN_CLK_OUT2=0xF0 PIN_CLK_OUT3=0xF00 SWITCH_PRT_NUM=0x7000 | - |
| `LEDC_CH0_CONF0` | 0x3F419000 | 32 | IDLE_LV=0x8 OVF_CNT_EN=0x8000 OVF_CNT_RESET=0x10000 OVF_CNT_RESET_ST=0x20000 OVF_NUM=0x7FE0 PARA_UP=0x10 SIG_OUT_EN=0x4 TIMER_SEL=0x3 | - |
| `LEDC_CH0_CONF1` | 0x3F41900C | 32 | DUTY_CYCLE=0xFFC00 DUTY_INC=0x40000000 DUTY_NUM=0x3FF00000 DUTY_SCALE=0x3FF DUTY_START=0x80000000 | - |
| `LEDC_CH0_DUTY` | 0x3F419008 | 32 | DUTY=0x7FFFF | - |
| `LEDC_CH0_DUTY_R` | 0x3F419010 | 32 | DUTY_R=0x7FFFF | - |
| `LEDC_CH0_HPOINT` | 0x3F419004 | 32 | HPOINT=0x3FFF | - |
| `LEDC_CH1_CONF0` | 0x3F419014 | 32 | IDLE_LV=0x8 OVF_CNT_EN=0x8000 OVF_CNT_RESET=0x10000 OVF_CNT_RESET_ST=0x20000 OVF_NUM=0x7FE0 PARA_UP=0x10 SIG_OUT_EN=0x4 TIMER_SEL=0x3 | - |
| `LEDC_CH1_CONF1` | 0x3F419020 | 32 | DUTY_CYCLE=0xFFC00 DUTY_INC=0x40000000 DUTY_NUM=0x3FF00000 DUTY_SCALE=0x3FF DUTY_START=0x80000000 | - |
| `LEDC_CH1_DUTY` | 0x3F41901C | 32 | DUTY=0x7FFFF | - |
| `LEDC_CH1_DUTY_R` | 0x3F419024 | 32 | DUTY_R=0x7FFFF | - |
| `LEDC_CH1_HPOINT` | 0x3F419018 | 32 | HPOINT=0x3FFF | - |
| `LEDC_CH2_CONF0` | 0x3F419028 | 32 | IDLE_LV=0x8 OVF_CNT_EN=0x8000 OVF_CNT_RESET=0x10000 OVF_CNT_RESET_ST=0x20000 OVF_NUM=0x7FE0 PARA_UP=0x10 SIG_OUT_EN=0x4 TIMER_SEL=0x3 | - |
| `LEDC_CH2_CONF1` | 0x3F419034 | 32 | DUTY_CYCLE=0xFFC00 DUTY_INC=0x40000000 DUTY_NUM=0x3FF00000 DUTY_SCALE=0x3FF DUTY_START=0x80000000 | - |
| `LEDC_CH2_DUTY` | 0x3F419030 | 32 | DUTY=0x7FFFF | - |
| `LEDC_CH2_DUTY_R` | 0x3F419038 | 32 | DUTY_R=0x7FFFF | - |
| `LEDC_CH2_HPOINT` | 0x3F41902C | 32 | HPOINT=0x3FFF | - |
| `LEDC_CH3_CONF0` | 0x3F41903C | 32 | IDLE_LV=0x8 OVF_CNT_EN=0x8000 OVF_CNT_RESET=0x10000 OVF_CNT_RESET_ST=0x20000 OVF_NUM=0x7FE0 PARA_UP=0x10 SIG_OUT_EN=0x4 TIMER_SEL=0x3 | - |
| `LEDC_CH3_CONF1` | 0x3F419048 | 32 | DUTY_CYCLE=0xFFC00 DUTY_INC=0x40000000 DUTY_NUM=0x3FF00000 DUTY_SCALE=0x3FF DUTY_START=0x80000000 | - |
| `LEDC_CH3_DUTY` | 0x3F419044 | 32 | DUTY=0x7FFFF | - |
| `LEDC_CH3_DUTY_R` | 0x3F41904C | 32 | DUTY_R=0x7FFFF | - |
| `LEDC_CH3_HPOINT` | 0x3F419040 | 32 | HPOINT=0x3FFF | - |
| `LEDC_CH4_CONF0` | 0x3F419050 | 32 | IDLE_LV=0x8 OVF_CNT_EN=0x8000 OVF_CNT_RESET=0x10000 OVF_CNT_RESET_ST=0x20000 OVF_NUM=0x7FE0 PARA_UP=0x10 SIG_OUT_EN=0x4 TIMER_SEL=0x3 | - |
| `LEDC_CH4_CONF1` | 0x3F41905C | 32 | DUTY_CYCLE=0xFFC00 DUTY_INC=0x40000000 DUTY_NUM=0x3FF00000 DUTY_SCALE=0x3FF DUTY_START=0x80000000 | - |
| `LEDC_CH4_DUTY` | 0x3F419058 | 32 | DUTY=0x7FFFF | - |
| `LEDC_CH4_DUTY_R` | 0x3F419060 | 32 | DUTY_R=0x7FFFF | - |
| `LEDC_CH4_HPOINT` | 0x3F419054 | 32 | HPOINT=0x3FFF | - |
| `LEDC_CH5_CONF0` | 0x3F419064 | 32 | IDLE_LV=0x8 OVF_CNT_EN=0x8000 OVF_CNT_RESET=0x10000 OVF_CNT_RESET_ST=0x20000 OVF_NUM=0x7FE0 PARA_UP=0x10 SIG_OUT_EN=0x4 TIMER_SEL=0x3 | - |
| `LEDC_CH5_CONF1` | 0x3F419070 | 32 | DUTY_CYCLE=0xFFC00 DUTY_INC=0x40000000 DUTY_NUM=0x3FF00000 DUTY_SCALE=0x3FF DUTY_START=0x80000000 | - |
| `LEDC_CH5_DUTY` | 0x3F41906C | 32 | DUTY=0x7FFFF | - |
| `LEDC_CH5_DUTY_R` | 0x3F419074 | 32 | DUTY_R=0x7FFFF | - |
| `LEDC_CH5_HPOINT` | 0x3F419068 | 32 | HPOINT=0x3FFF | - |
| `LEDC_CH6_CONF0` | 0x3F419078 | 32 | IDLE_LV=0x8 OVF_CNT_EN=0x8000 OVF_CNT_RESET=0x10000 OVF_CNT_RESET_ST=0x20000 OVF_NUM=0x7FE0 PARA_UP=0x10 SIG_OUT_EN=0x4 TIMER_SEL=0x3 | - |
| `LEDC_CH6_CONF1` | 0x3F419084 | 32 | DUTY_CYCLE=0xFFC00 DUTY_INC=0x40000000 DUTY_NUM=0x3FF00000 DUTY_SCALE=0x3FF DUTY_START=0x80000000 | - |
| `LEDC_CH6_DUTY` | 0x3F419080 | 32 | DUTY=0x7FFFF | - |
| `LEDC_CH6_DUTY_R` | 0x3F419088 | 32 | DUTY_R=0x7FFFF | - |
| `LEDC_CH6_HPOINT` | 0x3F41907C | 32 | HPOINT=0x3FFF | - |
| `LEDC_CH7_CONF0` | 0x3F41908C | 32 | IDLE_LV=0x8 OVF_CNT_EN=0x8000 OVF_CNT_RESET=0x10000 OVF_CNT_RESET_ST=0x20000 OVF_NUM=0x7FE0 PARA_UP=0x10 SIG_OUT_EN=0x4 TIMER_SEL=0x3 | - |
| `LEDC_CH7_CONF1` | 0x3F419098 | 32 | DUTY_CYCLE=0xFFC00 DUTY_INC=0x40000000 DUTY_NUM=0x3FF00000 DUTY_SCALE=0x3FF DUTY_START=0x80000000 | - |
| `LEDC_CH7_DUTY` | 0x3F419094 | 32 | DUTY=0x7FFFF | - |
| `LEDC_CH7_DUTY_R` | 0x3F41909C | 32 | DUTY_R=0x7FFFF | - |
| `LEDC_CH7_HPOINT` | 0x3F419090 | 32 | HPOINT=0x3FFF | - |
| `LEDC_CONF` | 0x3F4190D0 | 32 | APB_CLK_SEL=0x3 CLK_EN=0x80000000 | - |
| `LEDC_DATE` | 0x3F4190FC | 32 | DATE=0xFFFFFFFF | - |
| `LEDC_INT_CLR` | 0x3F4190CC | 32 | DUTY_CHNG_END_CH0_INT_CLR=0x10 DUTY_CHNG_END_CH1_INT_CLR=0x20 DUTY_CHNG_END_CH2_INT_CLR=0x40 DUTY_CHNG_END_CH3_INT_CLR=0x80 DUTY_CHNG_END_CH4_INT_CLR=0x100 DUTY_CHNG_END_CH5_INT_CLR=0x200 DUTY_CHNG_END_CH6_INT_CLR=0x400 DUTY_CHNG_END_CH7_IN... | - |
| `LEDC_INT_ENA` | 0x3F4190C8 | 32 | DUTY_CHNG_END_CH0_INT_ENA=0x10 DUTY_CHNG_END_CH1_INT_ENA=0x20 DUTY_CHNG_END_CH2_INT_ENA=0x40 DUTY_CHNG_END_CH3_INT_ENA=0x80 DUTY_CHNG_END_CH4_INT_ENA=0x100 DUTY_CHNG_END_CH5_INT_ENA=0x200 DUTY_CHNG_END_CH6_INT_ENA=0x400 DUTY_CHNG_END_CH7_IN... | - |
| `LEDC_INT_RAW` | 0x3F4190C0 | 32 | DUTY_CHNG_END_CH0_INT_RAW=0x10 DUTY_CHNG_END_CH1_INT_RAW=0x20 DUTY_CHNG_END_CH2_INT_RAW=0x40 DUTY_CHNG_END_CH3_INT_RAW=0x80 DUTY_CHNG_END_CH4_INT_RAW=0x100 DUTY_CHNG_END_CH5_INT_RAW=0x200 DUTY_CHNG_END_CH6_INT_RAW=0x400 DUTY_CHNG_END_CH7_IN... | - |
| `LEDC_INT_ST` | 0x3F4190C4 | 32 | DUTY_CHNG_END_CH0_INT_ST=0x10 DUTY_CHNG_END_CH1_INT_ST=0x20 DUTY_CHNG_END_CH2_INT_ST=0x40 DUTY_CHNG_END_CH3_INT_ST=0x80 DUTY_CHNG_END_CH4_INT_ST=0x100 DUTY_CHNG_END_CH5_INT_ST=0x200 DUTY_CHNG_END_CH6_INT_ST=0x400 DUTY_CHNG_END_CH7_INT_ST=0x... | - |
| `LEDC_TIMER0_CONF` | 0x3F4190A0 | 32 | CLK_DIV=0x3FFFF0 DUTY_RES=0xF PARA_UP=0x2000000 PAUSE=0x400000 RST=0x800000 TICK_SEL=0x1000000 | - |
| `LEDC_TIMER0_VALUE` | 0x3F4190A4 | 32 | CNT=0x3FFF | - |
| `LEDC_TIMER1_CONF` | 0x3F4190A8 | 32 | CLK_DIV=0x3FFFF0 DUTY_RES=0xF PARA_UP=0x2000000 PAUSE=0x400000 RST=0x800000 TICK_SEL=0x1000000 | - |
| `LEDC_TIMER1_VALUE` | 0x3F4190AC | 32 | CNT=0x3FFF | - |
| `LEDC_TIMER2_CONF` | 0x3F4190B0 | 32 | CLK_DIV=0x3FFFF0 DUTY_RES=0xF PARA_UP=0x2000000 PAUSE=0x400000 RST=0x800000 TICK_SEL=0x1000000 | - |
| `LEDC_TIMER2_VALUE` | 0x3F4190B4 | 32 | CNT=0x3FFF | - |
| `LEDC_TIMER3_CONF` | 0x3F4190B8 | 32 | CLK_DIV=0x3FFFF0 DUTY_RES=0xF PARA_UP=0x2000000 PAUSE=0x400000 RST=0x800000 TICK_SEL=0x1000000 | - |
| `LEDC_TIMER3_VALUE` | 0x3F4190BC | 32 | CNT=0x3FFF | - |
| `PCNT_CTRL` | 0x3F417060 | 32 | CLK_EN=0x10000 CNT_PAUSE_U0=0x2 CNT_PAUSE_U1=0x8 CNT_PAUSE_U2=0x20 CNT_PAUSE_U3=0x80 PULSE_CNT_RST_U0=0x1 PULSE_CNT_RST_U1=0x4 PULSE_CNT_RST_U2=0x10 PULSE_CNT_RST_U3=0x40 | - |
| `PCNT_DATE` | 0x3F4170FC | 32 | DATE=0xFFFFFFFF | - |
| `PCNT_INT_CLR` | 0x3F41704C | 32 | CNT_THR_EVENT_U0_INT_CLR=0x1 CNT_THR_EVENT_U1_INT_CLR=0x2 CNT_THR_EVENT_U2_INT_CLR=0x4 CNT_THR_EVENT_U3_INT_CLR=0x8 | - |
| `PCNT_INT_ENA` | 0x3F417048 | 32 | CNT_THR_EVENT_U0_INT_ENA=0x1 CNT_THR_EVENT_U1_INT_ENA=0x2 CNT_THR_EVENT_U2_INT_ENA=0x4 CNT_THR_EVENT_U3_INT_ENA=0x8 | - |
| `PCNT_INT_RAW` | 0x3F417040 | 32 | CNT_THR_EVENT_U0_INT_RAW=0x1 CNT_THR_EVENT_U1_INT_RAW=0x2 CNT_THR_EVENT_U2_INT_RAW=0x4 CNT_THR_EVENT_U3_INT_RAW=0x8 | - |
| `PCNT_INT_ST` | 0x3F417044 | 32 | CNT_THR_EVENT_U0_INT_ST=0x1 CNT_THR_EVENT_U1_INT_ST=0x2 CNT_THR_EVENT_U2_INT_ST=0x4 CNT_THR_EVENT_U3_INT_ST=0x8 | - |
| `PCNT_U0_CNT` | 0x3F417030 | 32 | PULSE_CNT_U0=0xFFFF | - |
| `PCNT_U0_CONF0` | 0x3F417000 | 32 | CH0_HCTRL_MODE_U0=0x300000 CH0_LCTRL_MODE_U0=0xC00000 CH0_NEG_MODE_U0=0x30000 CH0_POS_MODE_U0=0xC0000 CH1_HCTRL_MODE_U0=0x30000000 CH1_LCTRL_MODE_U0=0xC0000000 CH1_NEG_MODE_U0=0x3000000 CH1_POS_MODE_U0=0xC000000 FILTER_EN_U0=0x400 FILTER_TH... | - |
| `PCNT_U0_CONF1` | 0x3F417004 | 32 | CNT_THRES0_U0=0xFFFF CNT_THRES1_U0=0xFFFF0000 | - |
| `PCNT_U0_CONF2` | 0x3F417008 | 32 | CNT_H_LIM_U0=0xFFFF CNT_L_LIM_U0=0xFFFF0000 | - |
| `PCNT_U0_STATUS` | 0x3F417050 | 32 | CNT_THR_H_LIM_LAT_U0=0x20 CNT_THR_L_LIM_LAT_U0=0x10 CNT_THR_THRES0_LAT_U0=0x8 CNT_THR_THRES1_LAT_U0=0x4 CNT_THR_ZERO_LAT_U0=0x40 CNT_THR_ZERO_MODE_U0=0x3 | - |
| `PCNT_U1_CNT` | 0x3F417034 | 32 | PULSE_CNT_U0=0xFFFF | - |
| `PCNT_U1_CONF0` | 0x3F41700C | 32 | CH0_HCTRL_MODE_U0=0x300000 CH0_LCTRL_MODE_U0=0xC00000 CH0_NEG_MODE_U0=0x30000 CH0_POS_MODE_U0=0xC0000 CH1_HCTRL_MODE_U0=0x30000000 CH1_LCTRL_MODE_U0=0xC0000000 CH1_NEG_MODE_U0=0x3000000 CH1_POS_MODE_U0=0xC000000 FILTER_EN_U0=0x400 FILTER_TH... | - |
| `PCNT_U1_CONF1` | 0x3F417010 | 32 | CNT_THRES0_U0=0xFFFF CNT_THRES1_U0=0xFFFF0000 | - |
| `PCNT_U1_CONF2` | 0x3F417014 | 32 | CNT_H_LIM_U0=0xFFFF CNT_L_LIM_U0=0xFFFF0000 | - |
| `PCNT_U1_STATUS` | 0x3F417054 | 32 | CNT_THR_H_LIM_LAT_U0=0x20 CNT_THR_L_LIM_LAT_U0=0x10 CNT_THR_THRES0_LAT_U0=0x8 CNT_THR_THRES1_LAT_U0=0x4 CNT_THR_ZERO_LAT_U0=0x40 CNT_THR_ZERO_MODE_U0=0x3 | - |
| `PCNT_U2_CNT` | 0x3F417038 | 32 | PULSE_CNT_U0=0xFFFF | - |
| `PCNT_U2_CONF0` | 0x3F417018 | 32 | CH0_HCTRL_MODE_U0=0x300000 CH0_LCTRL_MODE_U0=0xC00000 CH0_NEG_MODE_U0=0x30000 CH0_POS_MODE_U0=0xC0000 CH1_HCTRL_MODE_U0=0x30000000 CH1_LCTRL_MODE_U0=0xC0000000 CH1_NEG_MODE_U0=0x3000000 CH1_POS_MODE_U0=0xC000000 FILTER_EN_U0=0x400 FILTER_TH... | - |
| `PCNT_U2_CONF1` | 0x3F41701C | 32 | CNT_THRES0_U0=0xFFFF CNT_THRES1_U0=0xFFFF0000 | - |
| `PCNT_U2_CONF2` | 0x3F417020 | 32 | CNT_H_LIM_U0=0xFFFF CNT_L_LIM_U0=0xFFFF0000 | - |
| `PCNT_U2_STATUS` | 0x3F417058 | 32 | CNT_THR_H_LIM_LAT_U0=0x20 CNT_THR_L_LIM_LAT_U0=0x10 CNT_THR_THRES0_LAT_U0=0x8 CNT_THR_THRES1_LAT_U0=0x4 CNT_THR_ZERO_LAT_U0=0x40 CNT_THR_ZERO_MODE_U0=0x3 | - |
| `PCNT_U3_CNT` | 0x3F41703C | 32 | PULSE_CNT_U0=0xFFFF | - |
| `PCNT_U3_CONF0` | 0x3F417024 | 32 | CH0_HCTRL_MODE_U0=0x300000 CH0_LCTRL_MODE_U0=0xC00000 CH0_NEG_MODE_U0=0x30000 CH0_POS_MODE_U0=0xC0000 CH1_HCTRL_MODE_U0=0x30000000 CH1_LCTRL_MODE_U0=0xC0000000 CH1_NEG_MODE_U0=0x3000000 CH1_POS_MODE_U0=0xC000000 FILTER_EN_U0=0x400 FILTER_TH... | - |
| `PCNT_U3_CONF1` | 0x3F417028 | 32 | CNT_THRES0_U0=0xFFFF CNT_THRES1_U0=0xFFFF0000 | - |
| `PCNT_U3_CONF2` | 0x3F41702C | 32 | CNT_H_LIM_U0=0xFFFF CNT_L_LIM_U0=0xFFFF0000 | - |
| `PCNT_U3_STATUS` | 0x3F41705C | 32 | CNT_THR_H_LIM_LAT_U0=0x20 CNT_THR_L_LIM_LAT_U0=0x10 CNT_THR_THRES0_LAT_U0=0x8 CNT_THR_THRES1_LAT_U0=0x4 CNT_THR_ZERO_LAT_U0=0x40 CNT_THR_ZERO_MODE_U0=0x3 | - |
| `PMS_APB_PERIPHERAL_0` | 0x3F4C10CC | 32 | APB_PERIPHERAL_LOCK=0x1 | - |
| `PMS_APB_PERIPHERAL_1` | 0x3F4C10D0 | 32 | APB_PERIPHERAL_SPLIT_BURST=0x1 | - |
| `PMS_APB_PERIPHERAL_INTR` | 0x3F4C10F4 | 32 | APB_PERI_BYTE_ERROR_CLR=0x1 APB_PERI_BYTE_ERROR_EN=0x2 APB_PERI_BYTE_ERROR_INTR=0x4 | - |
| `PMS_APB_PERIPHERAL_STATUS` | 0x3F4C10F8 | 32 | APB_PERI_BYTE_ERROR_ADDR=0xFFFFFFFF | - |
| `PMS_CACHE_MMU_ACCESS_0` | 0x3F4C10EC | 32 | CACHE_MMU_ACCESS_LOCK=0x1 | - |
| `PMS_CACHE_MMU_ACCESS_1` | 0x3F4C10F0 | 32 | PRO_MMU_RD_ACS=0x1 PRO_MMU_WR_ACS=0x2 | - |
| `PMS_CACHE_SOURCE_0` | 0x3F4C10C4 | 32 | CACHE_SOURCE_LOCK=0x1 | - |
| `PMS_CACHE_SOURCE_1` | 0x3F4C10C8 | 32 | PRO_CACHE_D_SOURCE_PRO_DPORT=0x10 PRO_CACHE_D_SOURCE_PRO_DRAM0=0x8 PRO_CACHE_D_SOURCE_PRO_DROM0=0x20 PRO_CACHE_I_SOURCE_PRO_DROM0=0x4 PRO_CACHE_I_SOURCE_PRO_IRAM1=0x1 PRO_CACHE_I_SOURCE_PRO_IROM0=0x2 | - |
| `PMS_CACHE_TAG_ACCESS_0` | 0x3F4C10E4 | 32 | CACHE_TAG_ACCESS_LOCK=0x1 | - |
| `PMS_CACHE_TAG_ACCESS_1` | 0x3F4C10E8 | 32 | PRO_D_TAG_RD_ACS=0x4 PRO_D_TAG_WR_ACS=0x8 PRO_I_TAG_RD_ACS=0x1 PRO_I_TAG_WR_ACS=0x2 | - |
| `PMS_CLOCK_GATE_REG` | 0x3F4C1104 | 32 | CLK_EN=0x1 | - |
| `PMS_CPU_PERIPHERAL_INTR` | 0x3F4C10FC | 32 | CPU_PERI_BYTE_ERROR_CLR=0x1 CPU_PERI_BYTE_ERROR_EN=0x2 CPU_PERI_BYTE_ERROR_INTR=0x4 | - |
| `PMS_CPU_PERIPHERAL_STATUS` | 0x3F4C1100 | 32 | CPU_PERI_BYTE_ERROR_ADDR=0xFFFFFFFF | - |
| `PMS_DATE` | 0x3F4C1FFC | 32 | DATE=0xFFFFFFF | - |
| `PMS_DMA_APB_I_0` | 0x3F4C108C | 32 | DMA_APB_I_LOCK=0x1 | - |
| `PMS_DMA_APB_I_1` | 0x3F4C1090 | 32 | DMA_APB_I_SRAM_0_R=0x1 DMA_APB_I_SRAM_0_W=0x2 DMA_APB_I_SRAM_1_R=0x4 DMA_APB_I_SRAM_1_W=0x8 DMA_APB_I_SRAM_2_R=0x10 DMA_APB_I_SRAM_2_W=0x20 DMA_APB_I_SRAM_3_R=0x40 DMA_APB_I_SRAM_3_W=0x80 DMA_APB_I_SRAM_4_H_R=0x8000000 DMA_APB_I_SRAM_4_H_W=... | - |
| `PMS_DMA_APB_I_2` | 0x3F4C1094 | 32 | DMA_APB_I_ILG_CLR=0x1 DMA_APB_I_ILG_EN=0x2 DMA_APB_I_ILG_INTR=0x4 | - |
| `PMS_DMA_APB_I_3` | 0x3F4C1098 | 32 | DMA_APB_I_ILG_ST=0x7FFFFF | - |
| `PMS_DMA_RX_I_0` | 0x3F4C109C | 32 | DMA_RX_I_LOCK=0x1 | - |
| `PMS_DMA_RX_I_1` | 0x3F4C10A0 | 32 | DMA_RX_I_SRAM_0_R=0x1 DMA_RX_I_SRAM_0_W=0x2 DMA_RX_I_SRAM_1_R=0x4 DMA_RX_I_SRAM_1_W=0x8 DMA_RX_I_SRAM_2_R=0x10 DMA_RX_I_SRAM_2_W=0x20 DMA_RX_I_SRAM_3_R=0x40 DMA_RX_I_SRAM_3_W=0x80 DMA_RX_I_SRAM_4_H_R=0x8000000 DMA_RX_I_SRAM_4_H_W=0x10000000... | - |
| `PMS_DMA_RX_I_2` | 0x3F4C10A4 | 32 | DMA_RX_I_ILG_CLR=0x1 DMA_RX_I_ILG_EN=0x2 DMA_RX_I_ILG_INTR=0x4 | - |
| `PMS_DMA_RX_I_3` | 0x3F4C10A8 | 32 | DMA_RX_I_ILG_ST=0x7FFFFF | - |
| `PMS_DMA_TX_I_0` | 0x3F4C10AC | 32 | DMA_TX_I_LOCK=0x1 | - |
| `PMS_DMA_TX_I_1` | 0x3F4C10B0 | 32 | DMA_TX_I_SRAM_0_R=0x1 DMA_TX_I_SRAM_0_W=0x2 DMA_TX_I_SRAM_1_R=0x4 DMA_TX_I_SRAM_1_W=0x8 DMA_TX_I_SRAM_2_R=0x10 DMA_TX_I_SRAM_2_W=0x20 DMA_TX_I_SRAM_3_R=0x40 DMA_TX_I_SRAM_3_W=0x80 DMA_TX_I_SRAM_4_H_R=0x8000000 DMA_TX_I_SRAM_4_H_W=0x10000000... | - |
| `PMS_DMA_TX_I_2` | 0x3F4C10B4 | 32 | DMA_TX_I_ILG_CLR=0x1 DMA_TX_I_ILG_EN=0x2 DMA_TX_I_ILG_INTR=0x4 | - |
| `PMS_DMA_TX_I_3` | 0x3F4C10B8 | 32 | DMA_TX_I_ILG_ST=0x7FFFFF | - |
| `PMS_MAC_DUMP_0` | 0x3F4C1008 | 32 | MAC_DUMP_LOCK=0x1 | - |
| `PMS_MAC_DUMP_1` | 0x3F4C100C | 32 | MAC_DUMP_CONNECT=0xFFF | - |
| `PMS_OCCUPY_0` | 0x3F4C10D4 | 32 | OCCUPY_LOCK=0x1 | - |
| `PMS_OCCUPY_1` | 0x3F4C10D8 | 32 | OCCUPY_CACHE=0xF | - |
| `PMS_OCCUPY_2` | 0x3F4C10DC | 32 | OCCUPY_MAC_DUMP=0xF | - |
| `PMS_OCCUPY_3` | 0x3F4C10E0 | 32 | OCCUPY_PRO_TRACE=0x3FFFF | - |
| `PMS_PRO_AHB_0` | 0x3F4C105C | 32 | PRO_AHB_LOCK=0x1 | - |
| `PMS_PRO_AHB_1` | 0x3F4C1060 | 32 | PRO_AHB_RTCSLOW_0_H_F=0x4000 PRO_AHB_RTCSLOW_0_H_R=0x8000 PRO_AHB_RTCSLOW_0_H_W=0x10000 PRO_AHB_RTCSLOW_0_L_F=0x800 PRO_AHB_RTCSLOW_0_L_R=0x1000 PRO_AHB_RTCSLOW_0_L_W=0x2000 PRO_AHB_RTCSLOW_0_SPLTADDR=0x7FF | - |
| `PMS_PRO_AHB_2` | 0x3F4C1064 | 32 | PRO_AHB_RTCSLOW_1_H_F=0x4000 PRO_AHB_RTCSLOW_1_H_R=0x8000 PRO_AHB_RTCSLOW_1_H_W=0x10000 PRO_AHB_RTCSLOW_1_L_F=0x800 PRO_AHB_RTCSLOW_1_L_R=0x1000 PRO_AHB_RTCSLOW_1_L_W=0x2000 PRO_AHB_RTCSLOW_1_SPLTADDR=0x7FF | - |
| `PMS_PRO_AHB_3` | 0x3F4C1068 | 32 | PRO_AHB_ILG_CLR=0x1 PRO_AHB_ILG_EN=0x2 PRO_AHB_ILG_INTR=0x4 | - |
| `PMS_PRO_AHB_4` | 0x3F4C106C | 32 | PRO_AHB_ILG_ST=0xFFFFFFFF | - |
| `PMS_PRO_BOOT_LOCATION_0` | 0x3F4C10BC | 32 | PRO_BOOT_LOCATION_LOCK=0x1 | - |
| `PMS_PRO_BOOT_LOCATION_1` | 0x3F4C10C0 | 32 | PRO_BOOT_REMAP=0x1 | - |
| `PMS_PRO_CACHE_0` | 0x3F4C1078 | 32 | PRO_CACHE_LOCK=0x1 | - |
| `PMS_PRO_CACHE_1` | 0x3F4C107C | 32 | PRO_CACHE_CONNECT=0xFFFF | - |
| `PMS_PRO_CACHE_2` | 0x3F4C1080 | 32 | PRO_CACHE_ILG_CLR=0x1 PRO_CACHE_ILG_EN=0x2 PRO_CACHE_ILG_INTR=0x4 | - |
| `PMS_PRO_CACHE_3` | 0x3F4C1084 | 32 | PRO_CACHE_ILG_ST_I=0x1FFFF | - |
| `PMS_PRO_CACHE_4` | 0x3F4C1088 | 32 | PRO_CACHE_ILG_ST_D=0x1FFFF | - |
| `PMS_PRO_DPORT_0` | 0x3F4C103C | 32 | PRO_DPORT_LOCK=0x1 | - |
| `PMS_PRO_DPORT_1` | 0x3F4C1040 | 32 | PRO_DPORT_APB_PERIPHERAL_FORBID=0x1 PRO_DPORT_RESERVE_FIFO_VALID=0xF0000 PRO_DPORT_RTCSLOW_H_R=0x4000 PRO_DPORT_RTCSLOW_H_W=0x8000 PRO_DPORT_RTCSLOW_L_R=0x1000 PRO_DPORT_RTCSLOW_L_W=0x2000 PRO_DPORT_RTCSLOW_SPLTADDR=0xFFE | - |
| `PMS_PRO_DPORT_2` | 0x3F4C1044 | 32 | PRO_DPORT_RESERVE_FIFO_0=0x3FFFF | - |
| `PMS_PRO_DPORT_3` | 0x3F4C1048 | 32 | PRO_DPORT_RESERVE_FIFO_1=0x3FFFF | - |
| `PMS_PRO_DPORT_4` | 0x3F4C104C | 32 | PRO_DPORT_RESERVE_FIFO_2=0x3FFFF | - |
| `PMS_PRO_DPORT_5` | 0x3F4C1050 | 32 | PRO_DPORT_RESERVE_FIFO_3=0x3FFFF | - |
| `PMS_PRO_DPORT_6` | 0x3F4C1054 | 32 | PRO_DPORT_ILG_CLR=0x1 PRO_DPORT_ILG_EN=0x2 PRO_DPORT_ILG_INTR=0x4 | - |
| `PMS_PRO_DPORT_7` | 0x3F4C1058 | 32 | PRO_DPORT_ILG_ST=0x3FFFFFF | - |
| `PMS_PRO_DRAM0_0` | 0x3F4C1028 | 32 | PRO_DRAM0_LOCK=0x1 | - |
| `PMS_PRO_DRAM0_1` | 0x3F4C102C | 32 | PRO_DRAM0_SRAM_0_R=0x1 PRO_DRAM0_SRAM_0_W=0x2 PRO_DRAM0_SRAM_1_R=0x4 PRO_DRAM0_SRAM_1_W=0x8 PRO_DRAM0_SRAM_2_R=0x10 PRO_DRAM0_SRAM_2_W=0x20 PRO_DRAM0_SRAM_3_R=0x40 PRO_DRAM0_SRAM_3_W=0x80 PRO_DRAM0_SRAM_4_H_R=0x8000000 PRO_DRAM0_SRAM_4_H_W=... | - |
| `PMS_PRO_DRAM0_2` | 0x3F4C1030 | 32 | PRO_DRAM0_RTCFAST_H_R=0x2000 PRO_DRAM0_RTCFAST_H_W=0x4000 PRO_DRAM0_RTCFAST_L_R=0x800 PRO_DRAM0_RTCFAST_L_W=0x1000 PRO_DRAM0_RTCFAST_SPLTADDR=0x7FF | - |
| `PMS_PRO_DRAM0_3` | 0x3F4C1034 | 32 | PRO_DRAM0_ILG_CLR=0x1 PRO_DRAM0_ILG_EN=0x2 PRO_DRAM0_ILG_INTR=0x4 | - |
| `PMS_PRO_DRAM0_4` | 0x3F4C1038 | 32 | PRO_DRAM0_ILG_ST=0x3FFFFFF | - |
| `PMS_PRO_IRAM0_0` | 0x3F4C1010 | 32 | PRO_IRAM0_LOCK=0x1 | - |
| `PMS_PRO_IRAM0_1` | 0x3F4C1014 | 32 | PRO_IRAM0_SRAM_0_F=0x1 PRO_IRAM0_SRAM_0_R=0x2 PRO_IRAM0_SRAM_0_W=0x4 PRO_IRAM0_SRAM_1_F=0x8 PRO_IRAM0_SRAM_1_R=0x10 PRO_IRAM0_SRAM_1_W=0x20 PRO_IRAM0_SRAM_2_F=0x40 PRO_IRAM0_SRAM_2_R=0x80 PRO_IRAM0_SRAM_2_W=0x100 PRO_IRAM0_SRAM_3_F=0x200 PR... | - |
| `PMS_PRO_IRAM0_2` | 0x3F4C1018 | 32 | PRO_IRAM0_SRAM_4_H_F=0x100000 PRO_IRAM0_SRAM_4_H_R=0x200000 PRO_IRAM0_SRAM_4_H_W=0x400000 PRO_IRAM0_SRAM_4_L_F=0x20000 PRO_IRAM0_SRAM_4_L_R=0x40000 PRO_IRAM0_SRAM_4_L_W=0x80000 PRO_IRAM0_SRAM_4_SPLTADDR=0x1FFFF | - |
| `PMS_PRO_IRAM0_3` | 0x3F4C101C | 32 | PRO_IRAM0_RTCFAST_H_F=0x4000 PRO_IRAM0_RTCFAST_H_R=0x8000 PRO_IRAM0_RTCFAST_H_W=0x10000 PRO_IRAM0_RTCFAST_L_F=0x800 PRO_IRAM0_RTCFAST_L_R=0x1000 PRO_IRAM0_RTCFAST_L_W=0x2000 PRO_IRAM0_RTCFAST_SPLTADDR=0x7FF | - |
| `PMS_PRO_IRAM0_4` | 0x3F4C1020 | 32 | PRO_IRAM0_ILG_CLR=0x1 PRO_IRAM0_ILG_EN=0x2 PRO_IRAM0_ILG_INTR=0x4 | - |
| `PMS_PRO_IRAM0_5` | 0x3F4C1024 | 32 | PRO_IRAM0_ILG_ST=0x3FFFFF | - |
| `PMS_PRO_TRACE_0` | 0x3F4C1070 | 32 | PRO_TRACE_LOCK=0x1 | - |
| `PMS_PRO_TRACE_1` | 0x3F4C1074 | 32 | PRO_TRACE_DISABLE=0x1 | - |
| `PMS_SDIO_0` | 0x3F4C1000 | 32 | SDIO_LOCK=0x1 | - |
| `PMS_SDIO_1` | 0x3F4C1004 | 32 | SDIO_DISABLE=0x1 | - |
| `RMT_APB_CONF` | 0x3F416080 | 32 | APB_FIFO_MASK=0x1 CLK_EN=0x80000000 MEM_CLK_FORCE_ON=0x4 MEM_FORCE_PD=0x8 MEM_FORCE_PU=0x10 MEM_TX_WRAP_EN=0x2 | - |
| `RMT_CH0ADDR` | 0x3F416040 | 32 | APB_MEM_RADDR=0x7FC00 APB_MEM_WADDR=0x1FF | - |
| `RMT_CH0CARRIER_DUTY` | 0x3F416060 | 32 | CARRIER_HIGH=0xFFFF0000 CARRIER_LOW=0xFFFF | - |
| `RMT_CH0CONF0` | 0x3F416010 | 32 | CARRIER_EFF_EN=0x8000000 CARRIER_EN=0x10000000 CARRIER_OUT_LV=0x20000000 DIV_CNT=0xFF IDLE_THRES=0xFFFF00 MEM_SIZE=0x7000000 | - |
| `RMT_CH0CONF1` | 0x3F416014 | 32 | APB_MEM_RST=0x10 CHK_RX_CARRIER_EN=0x10000 IDLE_OUT_EN=0x80000 IDLE_OUT_LV=0x40000 MEM_OWNER=0x20 MEM_RD_RST=0x8 MEM_WR_RST=0x4 REF_ALWAYS_ON=0x20000 RX_EN=0x2 RX_FILTER_EN=0x80 RX_FILTER_THRES=0xFF00 TX_CONTI_MODE=0x40 TX_START=0x1 TX_STOP... | - |
| `RMT_CH0DATA` | 0x3F416000 | 32 | CH_DATA=0xFFFFFFFF | - |
| `RMT_CH0STATUS` | 0x3F416030 | 32 | APB_MEM_RD_ERR=0x8000000 APB_MEM_WR_ERR=0x4000000 MEM_EMPTY=0x2000000 MEM_FULL=0x1000000 MEM_OWNER_ERR=0x800000 MEM_RADDR_EX=0x7FC00 MEM_WADDR_EX=0x1FF STATE=0x700000 | - |
| `RMT_CH0_RX_CARRIER_RM` | 0x3F41608C | 32 | CARRIER_HIGH_THRES=0xFFFF0000 CARRIER_LOW_THRES=0xFFFF | - |
| `RMT_CH0_TX_LIM` | 0x3F416070 | 32 | LOOP_COUNT_RESET=0x100000 TX_LIM=0x1FF TX_LOOP_CNT_EN=0x80000 TX_LOOP_NUM=0x7FE00 | - |
| `RMT_CH1ADDR` | 0x3F416044 | 32 | APB_MEM_RADDR=0x7FC00 APB_MEM_WADDR=0x1FF | - |
| `RMT_CH1CARRIER_DUTY` | 0x3F416064 | 32 | CARRIER_HIGH=0xFFFF0000 CARRIER_LOW=0xFFFF | - |
| `RMT_CH1CONF0` | 0x3F416018 | 32 | CARRIER_EFF_EN=0x8000000 CARRIER_EN=0x10000000 CARRIER_OUT_LV=0x20000000 DIV_CNT=0xFF IDLE_THRES=0xFFFF00 MEM_SIZE=0x7000000 | - |
| `RMT_CH1CONF1` | 0x3F41601C | 32 | APB_MEM_RST=0x10 CHK_RX_CARRIER_EN=0x10000 IDLE_OUT_EN=0x80000 IDLE_OUT_LV=0x40000 MEM_OWNER=0x20 MEM_RD_RST=0x8 MEM_WR_RST=0x4 REF_ALWAYS_ON=0x20000 RX_EN=0x2 RX_FILTER_EN=0x80 RX_FILTER_THRES=0xFF00 TX_CONTI_MODE=0x40 TX_START=0x1 TX_STOP... | - |
| `RMT_CH1DATA` | 0x3F416004 | 32 | CH_DATA=0xFFFFFFFF | - |
| `RMT_CH1STATUS` | 0x3F416034 | 32 | APB_MEM_RD_ERR=0x8000000 APB_MEM_WR_ERR=0x4000000 MEM_EMPTY=0x2000000 MEM_FULL=0x1000000 MEM_OWNER_ERR=0x800000 MEM_RADDR_EX=0x7FC00 MEM_WADDR_EX=0x1FF STATE=0x700000 | - |
| `RMT_CH1_RX_CARRIER_RM` | 0x3F416090 | 32 | CARRIER_HIGH_THRES=0xFFFF0000 CARRIER_LOW_THRES=0xFFFF | - |
| `RMT_CH1_TX_LIM` | 0x3F416074 | 32 | LOOP_COUNT_RESET=0x100000 TX_LIM=0x1FF TX_LOOP_CNT_EN=0x80000 TX_LOOP_NUM=0x7FE00 | - |
| `RMT_CH2ADDR` | 0x3F416048 | 32 | APB_MEM_RADDR=0x7FC00 APB_MEM_WADDR=0x1FF | - |
| `RMT_CH2CARRIER_DUTY` | 0x3F416068 | 32 | CARRIER_HIGH=0xFFFF0000 CARRIER_LOW=0xFFFF | - |
| `RMT_CH2CONF0` | 0x3F416020 | 32 | CARRIER_EFF_EN=0x8000000 CARRIER_EN=0x10000000 CARRIER_OUT_LV=0x20000000 DIV_CNT=0xFF IDLE_THRES=0xFFFF00 MEM_SIZE=0x7000000 | - |
| `RMT_CH2CONF1` | 0x3F416024 | 32 | APB_MEM_RST=0x10 CHK_RX_CARRIER_EN=0x10000 IDLE_OUT_EN=0x80000 IDLE_OUT_LV=0x40000 MEM_OWNER=0x20 MEM_RD_RST=0x8 MEM_WR_RST=0x4 REF_ALWAYS_ON=0x20000 RX_EN=0x2 RX_FILTER_EN=0x80 RX_FILTER_THRES=0xFF00 TX_CONTI_MODE=0x40 TX_START=0x1 TX_STOP... | - |
| `RMT_CH2DATA` | 0x3F416008 | 32 | CH_DATA=0xFFFFFFFF | - |
| `RMT_CH2STATUS` | 0x3F416038 | 32 | APB_MEM_RD_ERR=0x8000000 APB_MEM_WR_ERR=0x4000000 MEM_EMPTY=0x2000000 MEM_FULL=0x1000000 MEM_OWNER_ERR=0x800000 MEM_RADDR_EX=0x7FC00 MEM_WADDR_EX=0x1FF STATE=0x700000 | - |
| `RMT_CH2_RX_CARRIER_RM` | 0x3F416094 | 32 | CARRIER_HIGH_THRES=0xFFFF0000 CARRIER_LOW_THRES=0xFFFF | - |
| `RMT_CH2_TX_LIM` | 0x3F416078 | 32 | LOOP_COUNT_RESET=0x100000 TX_LIM=0x1FF TX_LOOP_CNT_EN=0x80000 TX_LOOP_NUM=0x7FE00 | - |
| `RMT_CH3ADDR` | 0x3F41604C | 32 | APB_MEM_RADDR=0x7FC00 APB_MEM_WADDR=0x1FF | - |
| `RMT_CH3CARRIER_DUTY` | 0x3F41606C | 32 | CARRIER_HIGH=0xFFFF0000 CARRIER_LOW=0xFFFF | - |
| `RMT_CH3CONF0` | 0x3F416028 | 32 | CARRIER_EFF_EN=0x8000000 CARRIER_EN=0x10000000 CARRIER_OUT_LV=0x20000000 DIV_CNT=0xFF IDLE_THRES=0xFFFF00 MEM_SIZE=0x7000000 | - |
| `RMT_CH3CONF1` | 0x3F41602C | 32 | APB_MEM_RST=0x10 CHK_RX_CARRIER_EN=0x10000 IDLE_OUT_EN=0x80000 IDLE_OUT_LV=0x40000 MEM_OWNER=0x20 MEM_RD_RST=0x8 MEM_WR_RST=0x4 REF_ALWAYS_ON=0x20000 RX_EN=0x2 RX_FILTER_EN=0x80 RX_FILTER_THRES=0xFF00 TX_CONTI_MODE=0x40 TX_START=0x1 TX_STOP... | - |
| `RMT_CH3DATA` | 0x3F41600C | 32 | CH_DATA=0xFFFFFFFF | - |
| `RMT_CH3STATUS` | 0x3F41603C | 32 | APB_MEM_RD_ERR=0x8000000 APB_MEM_WR_ERR=0x4000000 MEM_EMPTY=0x2000000 MEM_FULL=0x1000000 MEM_OWNER_ERR=0x800000 MEM_RADDR_EX=0x7FC00 MEM_WADDR_EX=0x1FF STATE=0x700000 | - |
| `RMT_CH3_RX_CARRIER_RM` | 0x3F416098 | 32 | CARRIER_HIGH_THRES=0xFFFF0000 CARRIER_LOW_THRES=0xFFFF | - |
| `RMT_CH3_TX_LIM` | 0x3F41607C | 32 | LOOP_COUNT_RESET=0x100000 TX_LIM=0x1FF TX_LOOP_CNT_EN=0x80000 TX_LOOP_NUM=0x7FE00 | - |
| `RMT_DATE` | 0x3F4160FC | 32 | DATE=0xFFFFFFFF | - |
| `RMT_INT_CLR` | 0x3F41605C | 32 | CH%s_ERR_INT_CLR=0x4 CH%s_RX_END_INT_CLR=0x2 CH%s_TX_END_INT_CLR=0x1 CH%s_TX_LOOP_INT_CLR=0x10000 CH%s_TX_THR_EVENT_INT_CLR=0x1000 | - |
| `RMT_INT_ENA` | 0x3F416058 | 32 | CH%s_ERR_INT_ENA=0x4 CH%s_RX_END_INT_ENA=0x2 CH%s_TX_END_INT_ENA=0x1 CH%s_TX_LOOP_INT_ENA=0x10000 CH%s_TX_THR_EVENT_INT_ENA=0x1000 | - |
| `RMT_INT_RAW` | 0x3F416050 | 32 | CH%s_ERR_INT_RAW=0x4 CH%s_RX_END_INT_RAW=0x2 CH%s_TX_END_INT_RAW=0x1 CH%s_TX_LOOP_INT_RAW=0x10000 CH%s_TX_THR_EVENT_INT_RAW=0x1000 | - |
| `RMT_INT_ST` | 0x3F416054 | 32 | CH%s_ERR_INT_ST=0x4 CH%s_RX_END_INT_ST=0x2 CH%s_TX_END_INT_ST=0x1 CH%s_TX_LOOP_INT_ST=0x10000 CH%s_TX_THR_EVENT_INT_ST=0x1000 | - |
| `RMT_REF_CNT_RST` | 0x3F416088 | 32 | CH0=0x1 CH1=0x2 CH2=0x4 CH3=0x8 | - |
| `RMT_TX_SIM` | 0x3F416084 | 32 | CH0=0x1 CH1=0x2 CH2=0x4 CH3=0x8 EN=0x10 | - |
| `RNG_DATA` | 0x60035110 | 32 |  | - |
| `RSA_CLEAN` | 0x6003C808 | 32 | CLEAN=0x1 | - |
| `RSA_CLEAR_INTERRUPT` | 0x6003C81C | 32 | CLEAR_INTERRUPT=0x1 | - |
| `RSA_CONSTANT_TIME` | 0x6003C820 | 32 | CONSTANT_TIME=0x1 | - |
| `RSA_DATE` | 0x6003C830 | 32 | DATE=0x3FFFFFFF | - |
| `RSA_IDLE` | 0x6003C818 | 32 | IDLE=0x1 | - |
| `RSA_INTERRUPT_ENA` | 0x6003C82C | 32 | INTERRUPT_ENA=0x1 | - |
| `RSA_MODE` | 0x6003C804 | 32 | MODE=0x7F | - |
| `RSA_MODEXP_START` | 0x6003C80C | 32 | MODEXP_START=0x1 | - |
| `RSA_MODMULT_START` | 0x6003C810 | 32 | MODMULT_START=0x1 | - |
| `RSA_MULT_START` | 0x6003C814 | 32 | MULT_START=0x1 | - |
| `RSA_M_PRIME` | 0x6003C800 | 32 | M_PRIME=0xFFFFFFFF | - |
| `RSA_SEARCH_ENABLE` | 0x6003C824 | 32 | SEARCH_ENABLE=0x1 | - |
| `RSA_SEARCH_POS` | 0x6003C828 | 32 | SEARCH_POS=0xFFF | - |
| `RTCIO_ENABLE_W1TC` | 0x3F408414 | 32 | ENABLE_W1TC=0xFFFFFC00 | - |
| `RTCIO_EXT_WAKEUP0` | 0x3F4084DC | 32 | SEL=0xF8000000 | - |
| `RTCIO_PAD_DAC1` | 0x3F4084C8 | 32 | PDAC1_DAC=0x7F8 PDAC1_DAC_XPD_FORCE=0x1000 PDAC1_DRV=0x60000000 PDAC1_FUN_IE=0x2000 PDAC1_FUN_SEL=0x60000 PDAC1_MUX_SEL=0x80000 PDAC1_RDE=0x10000000 PDAC1_RUE=0x8000000 PDAC1_SLP_IE=0x8000 PDAC1_SLP_OE=0x4000 PDAC1_SLP_SEL=0x10000 PDAC1_XPD... | - |
| `RTCIO_PAD_DAC2` | 0x3F4084CC | 32 | PDAC2_DAC=0x7F8 PDAC2_DAC_XPD_FORCE=0x1000 PDAC2_DRV=0x60000000 PDAC2_FUN_IE=0x2000 PDAC2_FUN_SEL=0x60000 PDAC2_MUX_SEL=0x80000 PDAC2_RDE=0x10000000 PDAC2_RUE=0x8000000 PDAC2_SLP_IE=0x8000 PDAC2_SLP_OE=0x4000 PDAC2_SLP_SEL=0x10000 PDAC2_XPD... | - |
| `RTCIO_PIN0` | 0x3F408428 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN1` | 0x3F40842C | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN10` | 0x3F408450 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN11` | 0x3F408454 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN12` | 0x3F408458 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN13` | 0x3F40845C | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN14` | 0x3F408460 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN15` | 0x3F408464 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN16` | 0x3F408468 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN17` | 0x3F40846C | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN18` | 0x3F408470 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN19` | 0x3F408474 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN2` | 0x3F408430 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN20` | 0x3F408478 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN21` | 0x3F40847C | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN3` | 0x3F408434 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN4` | 0x3F408438 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN5` | 0x3F40843C | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN6` | 0x3F408440 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN7` | 0x3F408444 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN8` | 0x3F408448 | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_PIN9` | 0x3F40844C | 32 | GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_WAKEUP_ENABLE=0x400 PAD_DRIVER=0x4 | - |
| `RTCIO_RTC_DEBUG_SEL` | 0x3F408480 | 32 | RTC_DEBUG_12M_NO_GATING=0x2000000 RTC_DEBUG_SEL0=0x1F RTC_DEBUG_SEL1=0x3E0 RTC_DEBUG_SEL2=0x7C00 RTC_DEBUG_SEL3=0xF8000 RTC_DEBUG_SEL4=0x1F00000 | - |
| `RTCIO_RTC_GPIO_ENABLE` | 0x3F40840C | 32 | REG_RTCIO_REG_GPIO_ENABLE=0xFFFFFC00 | - |
| `RTCIO_RTC_GPIO_ENABLE_W1TS` | 0x3F408410 | 32 | REG_RTCIO_REG_GPIO_ENABLE_W1TS=0xFFFFFC00 | - |
| `RTCIO_RTC_GPIO_IN` | 0x3F408424 | 32 | GPIO_IN_NEXT=0xFFFFFC00 | - |
| `RTCIO_RTC_GPIO_OUT` | 0x3F408400 | 32 | GPIO_OUT_DATA=0xFFFFFC00 | - |
| `RTCIO_RTC_GPIO_OUT_W1TC` | 0x3F408408 | 32 | GPIO_OUT_DATA_W1TC=0xFFFFFC00 | - |
| `RTCIO_RTC_GPIO_OUT_W1TS` | 0x3F408404 | 32 | GPIO_OUT_DATA_W1TS=0xFFFFFC00 | - |
| `RTCIO_RTC_GPIO_STATUS` | 0x3F408418 | 32 | GPIO_STATUS_INT=0xFFFFFC00 | - |
| `RTCIO_RTC_GPIO_STATUS_W1TC` | 0x3F408420 | 32 | GPIO_STATUS_INT_W1TC=0xFFFFFC00 | - |
| `RTCIO_RTC_GPIO_STATUS_W1TS` | 0x3F40841C | 32 | GPIO_STATUS_INT_W1TS=0xFFFFFC00 | - |
| `RTCIO_RTC_IO_DATE` | 0x3F4085FC | 32 | IO_DATE=0xFFFFFFF | - |
| `RTCIO_RTC_IO_TOUCH_CTRL` | 0x3F4084E8 | 32 | IO_TOUCH_BUFMODE=0x10 IO_TOUCH_BUFSEL=0xF | - |
| `RTCIO_RTC_PAD19` | 0x3F4084D0 | 32 | DRV=0x60000000 FUN_IE=0x2000 FUN_SEL=0x60000 MUX_SEL=0x80000 RDE=0x10000000 RUE=0x8000000 SLP_IE=0x8000 SLP_OE=0x4000 SLP_SEL=0x10000 | - |
| `RTCIO_RTC_PAD20` | 0x3F4084D4 | 32 | DRV=0x60000000 FUN_IE=0x2000 FUN_SEL=0x60000 MUX_SEL=0x80000 RDE=0x10000000 RUE=0x8000000 SLP_IE=0x8000 SLP_OE=0x4000 SLP_SEL=0x10000 | - |
| `RTCIO_RTC_PAD21` | 0x3F4084D8 | 32 | DRV=0x60000000 FUN_IE=0x2000 FUN_SEL=0x60000 MUX_SEL=0x80000 RDE=0x10000000 RUE=0x8000000 SLP_IE=0x8000 SLP_OE=0x4000 SLP_SEL=0x10000 | - |
| `RTCIO_SAR_I2C_IO` | 0x3F4084E4 | 32 | SAR_DEBUG_BIT_SEL=0xF800000 SAR_I2C_SCL_SEL=0x30000000 SAR_I2C_SDA_SEL=0xC0000000 | - |
| `RTCIO_TOUCH_PAD0` | 0x3F408484 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD1` | 0x3F408488 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD10` | 0x3F4084AC | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD11` | 0x3F4084B0 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD12` | 0x3F4084B4 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD13` | 0x3F4084B8 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD14` | 0x3F4084BC | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD2` | 0x3F40848C | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD3` | 0x3F408490 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD4` | 0x3F408494 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD5` | 0x3F408498 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD6` | 0x3F40849C | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD7` | 0x3F4084A0 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD8` | 0x3F4084A4 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_TOUCH_PAD9` | 0x3F4084A8 | 32 | TOUCH_PAD0_DAC=0x3800000 TOUCH_PAD0_DRV=0x60000000 TOUCH_PAD0_FUN_IE=0x2000 TOUCH_PAD0_FUN_SEL=0x60000 TOUCH_PAD0_MUX_SEL=0x80000 TOUCH_PAD0_RDE=0x10000000 TOUCH_PAD0_RUE=0x8000000 TOUCH_PAD0_SLP_IE=0x8000 TOUCH_PAD0_SLP_OE=0x4000 TOUCH_PAD... | - |
| `RTCIO_XTAL_32N_PAD` | 0x3F4084C4 | 32 | X32N_DRV=0x60000000 X32N_FUN_IE=0x2000 X32N_FUN_SEL=0x60000 X32N_MUX_SEL=0x80000 X32N_RDE=0x10000000 X32N_RUE=0x8000000 X32N_SLP_IE=0x8000 X32N_SLP_OE=0x4000 X32N_SLP_SEL=0x10000 | - |
| `RTCIO_XTAL_32P_PAD` | 0x3F4084C0 | 32 | X32P_DRV=0x60000000 X32P_FUN_IE=0x2000 X32P_FUN_SEL=0x60000 X32P_MUX_SEL=0x80000 X32P_RDE=0x10000000 X32P_RUE=0x8000000 X32P_SLP_IE=0x8000 X32P_SLP_OE=0x4000 X32P_SLP_SEL=0x10000 | - |
| `RTCIO_XTL_EXT_CTR` | 0x3F4084E0 | 32 | SEL=0xF8000000 | - |
| `RTC_CNTL_ANA_CONF` | 0x3F408034 | 32 | BBPLL_CAL_SLP_START=0x2000000 CKGEN_I2C_PU=0x40000000 GLITCH_RST_EN=0x100000 I2C_RESET_POR_FORCE_PD=0x40000 I2C_RESET_POR_FORCE_PU=0x80000 PLLA_FORCE_PD=0x800000 PLLA_FORCE_PU=0x1000000 PLL_I2C_PU=0x80000000 PVTMON_PU=0x4000000 RFRX_PBUS_PU... | - |
| `RTC_CNTL_BIAS_CONF` | 0x3F408080 | 32 | BIAS_BUF_DEEP_SLP=0x1000 BIAS_BUF_IDLE=0x400 BIAS_BUF_MONITOR=0x2000 BIAS_BUF_WAKE=0x800 BIAS_SLEEP_DEEP_SLP=0x10000 BIAS_SLEEP_MONITOR=0x20000 DBG_ATTEN_DEEP_SLP=0x3C0000 DBG_ATTEN_MONITOR=0x3C00000 DEC_HEARTBEAT_PERIOD=0x10000000 DEC_HEAR... | - |
| `RTC_CNTL_BROWN_OUT` | 0x3F4080E4 | 32 | BROWN_OUT2_ENA=0x1 CLOSE_FLASH_ENA=0x4000 CNT_CLR=0x20000000 DET=0x80000000 ENA=0x40000000 INT_WAIT=0x3FF0 PD_RF_ENA=0x8000 RST_ENA=0x4000000 RST_SEL=0x8000000 RST_WAIT=0x3FF0000 | - |
| `RTC_CNTL_CLK_CONF` | 0x3F408074 | 32 | ANA_CLK_RTC_SEL=0xC0000000 CK8M_DFREQ=0x1FE0000 CK8M_DIV=0x30 CK8M_DIV_SEL=0x7000 CK8M_DIV_SEL_VLD=0x8 CK8M_FORCE_NOGATING=0x10000 CK8M_FORCE_PD=0x2000000 CK8M_FORCE_PU=0x4000000 DIG_CLK8M_D256_EN=0x200 DIG_CLK8M_EN=0x400 DIG_XTAL32K_EN=0x1... | - |
| `RTC_CNTL_COCPU_CTRL` | 0x3F408100 | 32 | COCPU_CLK_FO=0x1 COCPU_DONE=0x2000000 COCPU_DONE_FORCE=0x1000000 COCPU_SEL=0x800000 COCPU_SHUT=0x2000 COCPU_SHUT_2_CLK_DIS=0x3FC000 COCPU_SHUT_RESET_EN=0x400000 COCPU_START_2_INTR_EN=0x1F80 COCPU_START_2_RESET_DIS=0x7E COCPU_SW_INT_TRIGGER=... | - |
| `RTC_CNTL_CPU_PERIOD_CONF` | 0x3F40806C | 32 | CPUPERIOD_SEL=0xC0000000 CPUSEL_CONF=0x20000000 | - |
| `RTC_CNTL_DATE` | 0x3F408138 | 32 | CNTL_DATE=0xFFFFFFF | - |
| `RTC_CNTL_DIAG0` | 0x3F4080D0 | 32 | LOW_POWER_DIAG1=0xFFFFFFFF | - |
| `RTC_CNTL_DIG_ISO` | 0x3F408090 | 32 | CLR_DG_PAD_AUTOHOLD=0x400 DG_PAD_AUTOHOLD=0x200 DG_PAD_AUTOHOLD_EN=0x800 DG_PAD_FORCE_HOLD=0x8000 DG_PAD_FORCE_ISO=0x2000 DG_PAD_FORCE_NOISO=0x1000 DG_PAD_FORCE_UNHOLD=0x4000 DG_WRAP_FORCE_ISO=0x40000000 DG_WRAP_FORCE_NOISO=0x80000000 FORCE... | - |
| `RTC_CNTL_DIG_PAD_HOLD` | 0x3F4080D8 | 32 | DIG_PAD_HOLD=0xFFFFFFFF | - |
| `RTC_CNTL_DIG_PWC` | 0x3F40808C | 32 | DG_DCDC_FORCE_PD=0x200000 DG_DCDC_FORCE_PU=0x400000 DG_DCDC_PD_EN=0x800000 DG_WRAP_FORCE_PD=0x80000 DG_WRAP_FORCE_PU=0x100000 DG_WRAP_PD_EN=0x80000000 INTER_RAM0_FORCE_PD=0x80 INTER_RAM0_FORCE_PU=0x100 INTER_RAM0_PD_EN=0x2000000 INTER_RAM1_... | - |
| `RTC_CNTL_EXT_WAKEUP1` | 0x3F4080DC | 32 | SEL=0x3FFFFF STATUS_CLR=0x400000 | - |
| `RTC_CNTL_EXT_WAKEUP1_STATUS` | 0x3F4080E0 | 32 | EXT_WAKEUP1_STATUS=0x3FFFFF | - |
| `RTC_CNTL_EXT_WAKEUP_CONF` | 0x3F408064 | 32 | EXT_WAKEUP0_LV=0x40000000 EXT_WAKEUP1_LV=0x80000000 GPIO_WAKEUP_FILTER=0x20000000 | - |
| `RTC_CNTL_EXT_XTL_CONF` | 0x3F408060 | 32 | DAC_XTAL_32K=0xE0000 DBUF_XTAL_32K=0x200 DGM_XTAL_32K=0x1C00 DRES_XTAL_32K=0xE000 ENCKINIT_XTAL_32K=0x100 WDT_STATE=0x700000 XPD_XTAL_32K=0x10000 XTAL32K_AUTO_BACKUP=0x10 XTAL32K_AUTO_RESTART=0x20 XTAL32K_AUTO_RETURN=0x40 XTAL32K_EXT_CLK_FO... | - |
| `RTC_CNTL_INT_CLR_RTC` | 0x3F40804C | 32 | BROWN_OUT_INT_CLR=0x200 COCPU_INT_CLR=0x2000 COCPU_TRAP_INT_CLR=0x20000 GLITCH_DET_INT_CLR=0x80000 MAIN_TIMER_INT_CLR=0x400 SARADC1_INT_CLR=0x800 SARADC2_INT_CLR=0x4000 SDIO_IDLE_INT_CLR=0x4 SLP_REJECT_INT_CLR=0x2 SLP_WAKEUP_INT_CLR=0x1 SWD... | - |
| `RTC_CNTL_INT_ENA_RTC` | 0x3F408040 | 32 | BROWN_OUT_INT_ENA=0x200 COCPU_INT_ENA=0x2000 COCPU_TRAP_INT_ENA=0x20000 GLITCH_DET_INT_ENA=0x80000 MAIN_TIMER_INT_ENA=0x400 SARADC1_INT_ENA=0x800 SARADC2_INT_ENA=0x4000 SDIO_IDLE_INT_ENA=0x4 SLP_REJECT_INT_ENA=0x2 SLP_WAKEUP_INT_ENA=0x1 SWD... | - |
| `RTC_CNTL_INT_RAW_RTC` | 0x3F408044 | 32 | BROWN_OUT_INT_RAW=0x200 COCPU_INT_RAW=0x2000 COCPU_TRAP_INT_RAW=0x20000 GLITCH_DET_INT_RAW=0x80000 MAIN_TIMER_INT_RAW=0x400 SARADC1_INT_RAW=0x800 SARADC2_INT_RAW=0x4000 SDIO_IDLE_INT_RAW=0x4 SLP_REJECT_INT_RAW=0x2 SLP_WAKEUP_INT_RAW=0x1 SWD... | - |
| `RTC_CNTL_INT_ST_RTC` | 0x3F408048 | 32 | BROWN_OUT_INT_ST=0x200 COCPU_INT_ST=0x2000 COCPU_TRAP_INT_ST=0x20000 GLITCH_DET_INT_ST=0x80000 MAIN_TIMER_INT_ST=0x400 SARADC1_INT_ST=0x800 SARADC2_INT_ST=0x4000 SDIO_IDLE_INT_ST=0x4 SLP_REJECT_INT_ST=0x2 SLP_WAKEUP_INT_ST=0x1 SWD_INT_ST=0x... | - |
| `RTC_CNTL_LOW_POWER_ST` | 0x3F4080CC | 32 | COCPU_STATE_DONE=0x10000 COCPU_STATE_SLP=0x8000 COCPU_STATE_START=0x2000 COCPU_STATE_SWITCH=0x4000 DIG_ISO=0x80 IN_LOW_POWER_STATE=0x400000 IN_WAKEUP_STATE=0x200000 MAIN_STATE=0xF0000000 MAIN_STATE_IN_IDLE=0x8000000 MAIN_STATE_IN_SLP=0x4000... | - |
| `RTC_CNTL_OPTIONS0` | 0x3F408000 | 32 | ANALOG_FORCE_ISO=0x2000000 ANALOG_FORCE_NOISO=0x10000000 BBPLL_FORCE_PD=0x400 BBPLL_FORCE_PU=0x800 BBPLL_I2C_FORCE_PD=0x100 BBPLL_I2C_FORCE_PU=0x200 BB_I2C_FORCE_PD=0x40 BB_I2C_FORCE_PU=0x80 DG_WRAP_FORCE_NORST=0x40000000 DG_WRAP_FORCE_RST=... | - |
| `RTC_CNTL_OPTIONS1` | 0x3F408128 | 32 | FORCE_DOWNLOAD_BOOT=0x1 | - |
| `RTC_CNTL_PAD_HOLD` | 0x3F4080D4 | 32 | PAD19_HOLD=0x80000 PAD20_HOLD=0x100000 PAD21_HOLD=0x200000 PDAC1_HOLD=0x20000 PDAC2_HOLD=0x40000 TOUCH_PAD0_HOLD=0x1 TOUCH_PAD10_HOLD=0x400 TOUCH_PAD11_HOLD=0x800 TOUCH_PAD12_HOLD=0x1000 TOUCH_PAD13_HOLD=0x2000 TOUCH_PAD14_HOLD=0x4000 TOUCH... | - |
| `RTC_CNTL_PWC` | 0x3F408088 | 32 | FASTMEM_FOLW_CPU=0x40 FASTMEM_FORCE_ISO=0x2 FASTMEM_FORCE_LPD=0x80 FASTMEM_FORCE_LPU=0x100 FASTMEM_FORCE_NOISO=0x1 FASTMEM_FORCE_PD=0x1000 FASTMEM_FORCE_PU=0x2000 FASTMEM_PD_EN=0x4000 FORCE_ISO=0x10 FORCE_NOISO=0x20 FORCE_PD=0x40000 FORCE_P... | - |
| `RTC_CNTL_REG` | 0x3F408084 | 32 | DBIAS_SLP=0x1C00000 DBIAS_WAK=0xE000000 DBOOST_FORCE_PD=0x10000000 DBOOST_FORCE_PU=0x20000000 DIG_REG_DBIAS_SLP=0x700 DIG_REG_DBIAS_WAK=0x3800 REGULATOR_FORCE_PD=0x40000000 REGULATOR_FORCE_PU=0x80000000 SCK_DCAP=0x3FC000 | - |
| `RTC_CNTL_RESET_STATE` | 0x3F408038 | 32 | APPCPU_STAT_VECTOR_SEL=0x1000 PROCPU_STAT_VECTOR_SEL=0x2000 RESET_CAUSE_APPCPU=0xFC0 RESET_CAUSE_PROCPU=0x3F | - |
| `RTC_CNTL_SDIO_ACT_CONF` | 0x3F408070 | 32 | SDIO_ACT_DNUM=0xFFC00000 | - |
| `RTC_CNTL_SDIO_CONF` | 0x3F40807C | 32 | DREFH_SDIO=0x60000000 DREFL_SDIO=0x6000000 DREFM_SDIO=0x18000000 REG1P8_READY=0x1000000 SDIO_DCAP=0x1800 SDIO_DCURLIM=0x70000 SDIO_DTHDRV=0x600 SDIO_ENCURLIM=0x100000 SDIO_EN_INITI=0x8000 SDIO_FORCE=0x400000 SDIO_INITI=0x6000 SDIO_MODECURLI... | - |
| `RTC_CNTL_SLOW_CLK_CONF` | 0x3F408078 | 32 | ANA_CLK_DIV=0x7F800000 ANA_CLK_DIV_VLD=0x400000 SLOW_CLK_NEXT_EDGE=0x80000000 | - |
| `RTC_CNTL_SLP_REJECT_CAUSE` | 0x3F408124 | 32 | REJECT_CAUSE=0x1FFFF | - |
| `RTC_CNTL_SLP_REJECT_CONF` | 0x3F408068 | 32 | DEEP_SLP_REJECT_EN=0x80000000 LIGHT_SLP_REJECT_EN=0x40000000 SLEEP_REJECT_ENA=0x3FFFE000 | - |
| `RTC_CNTL_SLP_TIMER0` | 0x3F408004 | 32 | SLP_VAL_LO=0xFFFFFFFF | - |
| `RTC_CNTL_SLP_TIMER1` | 0x3F408008 | 32 | MAIN_TIMER_ALARM_EN=0x10000 SLP_VAL_HI=0xFFFF | - |
| `RTC_CNTL_SLP_WAKEUP_CAUSE` | 0x3F40812C | 32 | WAKEUP_CAUSE=0x1FFFF | - |
| `RTC_CNTL_STATE0` | 0x3F408018 | 32 | APB2RTC_BRIDGE_SEL=0x400000 SDIO_ACTIVE_IND=0x10000000 SLEEP_EN=0x80000000 SLP_REJECT=0x40000000 SLP_REJECT_CAUSE_CLR=0x2 SLP_WAKEUP=0x20000000 SW_CPU_INT=0x1 | - |
| `RTC_CNTL_STORE0` | 0x3F408050 | 32 | SCRATCH0=0xFFFFFFFF | - |
| `RTC_CNTL_STORE1` | 0x3F408054 | 32 | SCRATCH1=0xFFFFFFFF | - |
| `RTC_CNTL_STORE2` | 0x3F408058 | 32 | SCRATCH2=0xFFFFFFFF | - |
| `RTC_CNTL_STORE3` | 0x3F40805C | 32 | SCRATCH3=0xFFFFFFFF | - |
| `RTC_CNTL_STORE4` | 0x3F4080BC | 32 | SCRATCH4=0xFFFFFFFF | - |
| `RTC_CNTL_STORE5` | 0x3F4080C0 | 32 | SCRATCH5=0xFFFFFFFF | - |
| `RTC_CNTL_STORE6` | 0x3F4080C4 | 32 | SCRATCH6=0xFFFFFFFF | - |
| `RTC_CNTL_STORE7` | 0x3F4080C8 | 32 | SCRATCH7=0xFFFFFFFF | - |
| `RTC_CNTL_SWD_CONF` | 0x3F4080B0 | 32 | SWD_AUTO_FEED_EN=0x80000000 SWD_DISABLE=0x40000000 SWD_FEED=0x20000000 SWD_FEED_INT=0x2 SWD_RESET_FLAG=0x1 SWD_RST_FLAG_CLR=0x10000000 SWD_SIGNAL_WIDTH=0xFFC0000 | - |
| `RTC_CNTL_SWD_WPROTECT` | 0x3F4080B4 | 32 | SWD_WKEY=0xFFFFFFFF | - |
| `RTC_CNTL_SW_CPU_STALL` | 0x3F4080B8 | 32 | SW_STALL_APPCPU_C1=0x3F00000 SW_STALL_PROCPU_C1=0xFC000000 | - |
| `RTC_CNTL_TIMER1` | 0x3F40801C | 32 | CK8M_WAIT=0x3FC0 CPU_STALL_EN=0x1 CPU_STALL_WAIT=0x3E PLL_BUF_WAIT=0xFF000000 XTL_BUF_WAIT=0xFFC000 | - |
| `RTC_CNTL_TIMER2` | 0x3F408020 | 32 | MIN_TIME_CK8M_OFF=0xFF000000 ULPCP_TOUCH_START_WAIT=0xFF8000 | - |
| `RTC_CNTL_TIMER3` | 0x3F408024 | 32 | ROM_RAM_POWERUP_TIMER=0xFE000000 ROM_RAM_WAIT_TIMER=0x1FF0000 WIFI_POWERUP_TIMER=0xFE00 WIFI_WAIT_TIMER=0x1FF | - |
| `RTC_CNTL_TIMER4` | 0x3F408028 | 32 | DG_WRAP_POWERUP_TIMER=0xFE000000 DG_WRAP_WAIT_TIMER=0x1FF0000 POWERUP_TIMER=0xFE00 WAIT_TIMER=0x1FF | - |
| `RTC_CNTL_TIMER5` | 0x3F40802C | 32 | MIN_SLP_VAL=0xFF00 RTCMEM_POWERUP_TIMER=0xFE000000 RTCMEM_WAIT_TIMER=0x1FF0000 | - |
| `RTC_CNTL_TIMER6` | 0x3F408030 | 32 | DG_DCDC_POWERUP_TIMER=0xFE000000 DG_DCDC_WAIT_TIMER=0x1FF0000 | - |
| `RTC_CNTL_TIME_HIGH0` | 0x3F408014 | 32 | TIMER_VALUE0_HIGH=0xFFFF | - |
| `RTC_CNTL_TIME_HIGH1` | 0x3F4080EC | 32 | TIMER_VALUE1_HIGH=0xFFFF | - |
| `RTC_CNTL_TIME_LOW0` | 0x3F408010 | 32 | TIMER_VALUE0_LOW=0xFFFFFFFF | - |
| `RTC_CNTL_TIME_LOW1` | 0x3F4080E8 | 32 | TIMER_VALUE1_LOW=0xFFFFFFFF | - |
| `RTC_CNTL_TIME_UPDATE` | 0x3F40800C | 32 | TIMER_SYS_RST=0x20000000 TIMER_SYS_STALL=0x8000000 TIMER_XTL_OFF=0x10000000 TIME_UPDATE=0x80000000 | - |
| `RTC_CNTL_TOUCH_APPROACH` | 0x3F408114 | 32 | MEAS_TIME=0xFF000000 TOUCH_SLP_CHANNEL_CLR=0x800000 | - |
| `RTC_CNTL_TOUCH_CTRL1` | 0x3F408104 | 32 | TOUCH_MEAS_NUM=0xFFFF0000 TOUCH_SLEEP_CYCLES=0xFFFF | - |
| `RTC_CNTL_TOUCH_CTRL2` | 0x3F408108 | 32 | TOUCH_CLKGATE_EN=0x80000000 TOUCH_CLK_FO=0x40000000 TOUCH_DBIAS=0x1000 TOUCH_DRANGE=0xC TOUCH_DREFH=0xC0 TOUCH_DREFL=0x30 TOUCH_REFC=0xE00 TOUCH_RESET=0x20000000 TOUCH_SLP_CYC_DIV=0x6000000 TOUCH_SLP_TIMER_EN=0x2000 TOUCH_START_EN=0x8000 TO... | - |
| `RTC_CNTL_TOUCH_FILTER_CTRL` | 0x3F408118 | 32 | TOUCH_DEBOUNCE=0xE000000 TOUCH_FILTER_EN=0x80000000 TOUCH_FILTER_MODE=0x70000000 TOUCH_HYSTERESIS=0x1800000 TOUCH_JITTER_STEP=0x7800 TOUCH_NEG_NOISE_LIMIT=0x78000 TOUCH_NEG_NOISE_THRES=0x180000 TOUCH_NOISE_THRES=0x600000 TOUCH_SMOOTH_LVL=0x... | - |
| `RTC_CNTL_TOUCH_SCAN_CTRL` | 0x3F40810C | 32 | TOUCH_BUFDRV=0xE000000 TOUCH_DENOISE_EN=0x4 TOUCH_DENOISE_RES=0x3 TOUCH_INACTIVE_CONNECTION=0x100 TOUCH_OUT_RING=0xF0000000 TOUCH_SCAN_PAD_MAP=0x1FFFC00 TOUCH_SHIELD_PAD_EN=0x200 | - |
| `RTC_CNTL_TOUCH_SLP_THRES` | 0x3F408110 | 32 | TOUCH_SLP_APPROACH_EN=0x4000000 TOUCH_SLP_PAD=0xF8000000 TOUCH_SLP_TH=0x3FFFFF | - |
| `RTC_CNTL_TOUCH_TIMEOUT_CTRL` | 0x3F408120 | 32 | TOUCH_TIMEOUT_EN=0x400000 TOUCH_TIMEOUT_NUM=0x3FFFFF | - |
| `RTC_CNTL_ULP_CP_CTRL` | 0x3F4080FC | 32 | ULP_CP_CLK_FO=0x10000000 ULP_CP_FORCE_START_TOP=0x40000000 ULP_CP_MEM_ADDR_INIT=0x7FF ULP_CP_MEM_ADDR_SIZE=0x3FF800 ULP_CP_MEM_OFFSET_CLR=0x400000 ULP_CP_RESET=0x20000000 ULP_CP_START_TOP=0x80000000 | - |
| `RTC_CNTL_ULP_CP_TIMER` | 0x3F4080F8 | 32 | ULP_CP_GPIO_WAKEUP_CLR=0x40000000 ULP_CP_GPIO_WAKEUP_ENA=0x20000000 ULP_CP_PC_INIT=0x7FF ULP_CP_SLP_TIMER_EN=0x80000000 | - |
| `RTC_CNTL_ULP_CP_TIMER_1` | 0x3F408130 | 32 | ULP_CP_TIMER_SLP_CYCLE=0xFFFFFF00 | - |
| `RTC_CNTL_USB_CONF` | 0x3F40811C | 32 | IO_MUX_RESET_DISABLE=0x40000 USB_DM_PULLDOWN=0x200 USB_DM_PULLUP=0x100 USB_DP_PULLDOWN=0x80 USB_DP_PULLUP=0x40 USB_PAD_ENABLE=0x1000 USB_PAD_ENABLE_OVERRIDE=0x800 USB_PAD_PULL_OVERRIDE=0x20 USB_PULLUP_VALUE=0x400 USB_RESET_DISABLE=0x20000 U... | - |
| `RTC_CNTL_WAKEUP_STATE` | 0x3F40803C | 32 | WAKEUP_ENA=0xFFFF8000 | - |
| `RTC_CNTL_WDTCONFIG0` | 0x3F408094 | 32 | WDT_APPCPU_RESET_EN=0x400 WDT_CHIP_RESET_EN=0x100 WDT_CHIP_RESET_WIDTH=0xFF WDT_CPU_RESET_LENGTH=0x70000 WDT_EN=0x80000000 WDT_FLASHBOOT_MOD_EN=0x1000 WDT_PAUSE_IN_SLP=0x200 WDT_PROCPU_RESET_EN=0x800 WDT_STG0=0x70000000 WDT_STG1=0xE000000 W... | - |
| `RTC_CNTL_WDTCONFIG1` | 0x3F408098 | 32 | WDT_STG0_HOLD=0xFFFFFFFF | - |
| `RTC_CNTL_WDTCONFIG2` | 0x3F40809C | 32 | WDT_STG1_HOLD=0xFFFFFFFF | - |
| `RTC_CNTL_WDTCONFIG3` | 0x3F4080A0 | 32 | WDT_STG2_HOLD=0xFFFFFFFF | - |
| `RTC_CNTL_WDTCONFIG4` | 0x3F4080A4 | 32 | WDT_STG3_HOLD=0xFFFFFFFF | - |
| `RTC_CNTL_WDTFEED` | 0x3F4080A8 | 32 | WDT_FEED=0x80000000 | - |
| `RTC_CNTL_WDTWPROTECT` | 0x3F4080AC | 32 | WDT_WKEY=0xFFFFFFFF | - |
| `RTC_CNTL_XTAL32K_CLK_FACTOR` | 0x3F4080F0 | 32 | XTAL32K_CLK_FACTOR=0xFFFFFFFF | - |
| `RTC_CNTL_XTAL32K_CONF` | 0x3F4080F4 | 32 | XTAL32K_RESTART_WAIT=0xFFFF0 XTAL32K_RETURN_WAIT=0xF XTAL32K_STABLE_THRES=0xF0000000 XTAL32K_WDT_TIMEOUT=0xFF00000 | - |
| `RTC_I2C_CMD0` | 0x3F408C38 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD1` | 0x3F408C3C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD10` | 0x3F408C60 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD11` | 0x3F408C64 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD12` | 0x3F408C68 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD13` | 0x3F408C6C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD14` | 0x3F408C70 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD15` | 0x3F408C74 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD2` | 0x3F408C40 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD3` | 0x3F408C44 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD4` | 0x3F408C48 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD5` | 0x3F408C4C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD6` | 0x3F408C50 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD7` | 0x3F408C54 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD8` | 0x3F408C58 | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CMD9` | 0x3F408C5C | 32 | COMMAND=0x3FFF COMMAND_DONE=0x80000000 | - |
| `RTC_I2C_CTRL` | 0x3F408C04 | 32 | CLK_EN=0x80000000 CLK_GATE_EN=0x20000000 MS_MODE=0x4 RESET=0x40000000 RX_LSB_FIRST=0x20 SCL_FORCE_OUT=0x2 SDA_FORCE_OUT=0x1 TRANS_START=0x8 TX_LSB_FIRST=0x10 | - |
| `RTC_I2C_DATA` | 0x3F408C34 | 32 | DONE=0x80000000 RDATA=0xFF SLAVE_TX_DATA=0xFF00 | - |
| `RTC_I2C_DATE` | 0x3F408CFC | 32 | DATE=0xFFFFFFF | - |
| `RTC_I2C_INT_CLR` | 0x3F408C24 | 32 | ACK_ERR_INT_CLR=0x20 ARBITRATION_LOST_INT_CLR=0x2 DETECT_START_INT_CLR=0x100 MASTER_TRAN_COMP_INT_CLR=0x4 RX_DATA_INT_CLR=0x40 SLAVE_TRAN_COMP_INT_CLR=0x1 TIME_OUT_INT_CLR=0x10 TRANS_COMPLETE_INT_CLR=0x8 TX_DATA_INT_CLR=0x80 | - |
| `RTC_I2C_INT_ENA` | 0x3F408C30 | 32 | ACK_ERR_INT_ENA=0x20 ARBITRATION_LOST_INT_ENA=0x2 DETECT_START_INT_ENA=0x100 MASTER_TRAN_COMP_INT_ENA=0x4 RX_DATA_INT_ENA=0x40 SLAVE_TRAN_COMP_INT_ENA=0x1 TIME_OUT_INT_ENA=0x10 TRANS_COMPLETE_INT_ENA=0x8 TX_DATA_INT_ENA=0x80 | - |
| `RTC_I2C_INT_RAW` | 0x3F408C28 | 32 | ACK_ERR_INT_RAW=0x20 ARBITRATION_LOST_INT_RAW=0x2 DETECT_START_INT_RAW=0x100 MASTER_TRAN_COMP_INT_RAW=0x4 RX_DATA_INT_RAW=0x40 SLAVE_TRAN_COMP_INT_RAW=0x1 TIME_OUT_INT_RAW=0x10 TRANS_COMPLETE_INT_RAW=0x8 TX_DATA_INT_RAW=0x80 | - |
| `RTC_I2C_INT_ST` | 0x3F408C2C | 32 | ACK_ERR_INT_ST=0x20 ARBITRATION_LOST_INT_ST=0x2 DETECT_START_INT_ST=0x100 MASTER_TRAN_COMP_INT_ST=0x4 RX_DATA_INT_ST=0x40 SLAVE_TRAN_COMP_INT_ST=0x1 TIME_OUT_INT_ST=0x10 TRANS_COMPLETE_INT_ST=0x8 TX_DATA_INT_ST=0x80 | - |
| `RTC_I2C_SCL_HIGH` | 0x3F408C14 | 32 | PERIOD=0xFFFFF | - |
| `RTC_I2C_SCL_LOW` | 0x3F408C00 | 32 | PERIOD=0xFFFFF | - |
| `RTC_I2C_SCL_START_PERIOD` | 0x3F408C1C | 32 | SCL_START_PERIOD=0xFFFFF | - |
| `RTC_I2C_SCL_STOP_PERIOD` | 0x3F408C20 | 32 | SCL_STOP_PERIOD=0xFFFFF | - |
| `RTC_I2C_SDA_DUTY` | 0x3F408C18 | 32 | NUM=0xFFFFF | - |
| `RTC_I2C_SLAVE_ADDR` | 0x3F408C10 | 32 | ADDR_10BIT_EN=0x80000000 SLAVE_ADDR=0x7FFF | - |
| `RTC_I2C_STATUS` | 0x3F408C08 | 32 | ACK_REC=0x1 ARB_LOST=0x4 BUS_BUSY=0x8 BYTE_TRANS=0x20 OP_CNT=0xC0 SCL_MAIN_STATE_LAST=0x7000000 SCL_STATE_LAST=0x70000000 SHIFT=0xFF0000 SLAVE_ADDRESSED=0x10 SLAVE_RW=0x2 | - |
| `RTC_I2C_TO` | 0x3F408C0C | 32 | TIME_OUT=0xFFFFF | - |
| `SENS_SARDATE` | 0x3F408948 | 32 | SAR_DATE=0xFFFFFFF | - |
| `SENS_SAR_AMP_CTRL1` | 0x3F408818 | 32 | SAR_AMP_WAIT1=0xFFFF SAR_AMP_WAIT2=0xFFFF0000 | - |
| `SENS_SAR_AMP_CTRL2` | 0x3F40881C | 32 | AMP_RST_FB_FSM_IDLE=0x4 AMP_SHORT_REF_FSM_IDLE=0x8 AMP_SHORT_REF_GND_FSM_IDLE=0x10 SAR1_DAC_XPD_FSM_IDLE=0x1 SAR_AMP_WAIT3=0xFFFF0000 SAR_RSTB_FSM_IDLE=0x40 XPD_SAR_AMP_FSM_IDLE=0x2 XPD_SAR_FSM_IDLE=0x20 | - |
| `SENS_SAR_AMP_CTRL3` | 0x3F408820 | 32 | AMP_RST_FB_FSM=0xF00 AMP_SHORT_REF_FSM=0xF000 AMP_SHORT_REF_GND_FSM=0xF0000 SAR1_DAC_XPD_FSM=0xF SAR_RSTB_FSM=0xF000000 XPD_SAR_AMP_FSM=0xF0 XPD_SAR_FSM=0xF00000 | - |
| `SENS_SAR_ATTEN1` | 0x3F408814 | 32 | SAR1_ATTEN=0xFFFFFFFF | - |
| `SENS_SAR_ATTEN2` | 0x3F408838 | 32 | SAR2_ATTEN=0xFFFFFFFF | - |
| `SENS_SAR_COCPU_DEBUG` | 0x3F408938 | 32 | COCPU_MEM_ADDR=0xFFF80000 COCPU_MEM_RDY=0x4000 COCPU_MEM_VLD=0x2000 COCPU_MEM_WEN=0x78000 COCPU_PC=0x1FFF | - |
| `SENS_SAR_COCPU_INT_CLR` | 0x3F408934 | 32 | COCPU_SARADC1_INT_CLR=0x8 COCPU_SARADC2_INT_CLR=0x10 COCPU_START_INT_CLR=0x40 COCPU_SWD_INT_CLR=0x100 COCPU_SW_INT_CLR=0x80 COCPU_TOUCH_ACTIVE_INT_CLR=0x4 COCPU_TOUCH_DONE_INT_CLR=0x1 COCPU_TOUCH_INACTIVE_INT_CLR=0x2 COCPU_TSENS_INT_CLR=0x2... | - |
| `SENS_SAR_COCPU_INT_ENA` | 0x3F40892C | 32 | COCPU_SARADC1_INT_ENA=0x8 COCPU_SARADC2_INT_ENA=0x10 COCPU_START_INT_ENA=0x40 COCPU_SWD_INT_ENA=0x100 COCPU_SW_INT_ENA=0x80 COCPU_TOUCH_ACTIVE_INT_ENA=0x4 COCPU_TOUCH_DONE_INT_ENA=0x1 COCPU_TOUCH_INACTIVE_INT_ENA=0x2 COCPU_TSENS_INT_ENA=0x2... | - |
| `SENS_SAR_COCPU_INT_RAW` | 0x3F408928 | 32 | COCPU_SARADC1_INT_RAW=0x8 COCPU_SARADC2_INT_RAW=0x10 COCPU_START_INT_RAW=0x40 COCPU_SWD_INT_RAW=0x100 COCPU_SW_INT_RAW=0x80 COCPU_TOUCH_ACTIVE_INT_RAW=0x4 COCPU_TOUCH_DONE_INT_RAW=0x1 COCPU_TOUCH_INACTIVE_INT_RAW=0x2 COCPU_TSENS_INT_RAW=0x2... | - |
| `SENS_SAR_COCPU_INT_ST` | 0x3F408930 | 32 | COCPU_SARADC1_INT_ST=0x8 COCPU_SARADC2_INT_ST=0x10 COCPU_START_INT_ST=0x40 COCPU_SWD_INT_ST=0x100 COCPU_SW_INT_ST=0x80 COCPU_TOUCH_ACTIVE_INT_ST=0x4 COCPU_TOUCH_DONE_INT_ST=0x1 COCPU_TOUCH_INACTIVE_INT_ST=0x2 COCPU_TSENS_INT_ST=0x20 | - |
| `SENS_SAR_COCPU_STATE` | 0x3F408924 | 32 | COCPU_CLK_EN=0x4000000 COCPU_DBG_TRIGGER=0x2000000 COCPU_EBREAK=0x40000000 COCPU_EOI=0x10000000 COCPU_RESET_N=0x8000000 COCPU_TRAP=0x20000000 | - |
| `SENS_SAR_DAC_CTRL1` | 0x3F40891C | 32 | DAC_CLKGATE_EN=0x8000000 DAC_CLK_FORCE_HIGH=0x1000000 DAC_CLK_FORCE_LOW=0x800000 DAC_CLK_INV=0x2000000 DAC_DIG_FORCE=0x400000 DAC_RESET=0x4000000 DEBUG_BIT_SEL=0x3E0000 SW_FSTEP=0xFFFF SW_TONE_EN=0x10000 | - |
| `SENS_SAR_DAC_CTRL2` | 0x3F408920 | 32 | DAC_CW_EN1=0x1000000 DAC_CW_EN2=0x2000000 DAC_DC1=0xFF DAC_DC2=0xFF00 DAC_INV1=0x300000 DAC_INV2=0xC00000 DAC_SCALE1=0x30000 DAC_SCALE2=0xC0000 | - |
| `SENS_SAR_HALL_CTRL` | 0x3F40893C | 32 | HALL_PHASE=0x40000000 HALL_PHASE_FORCE=0x80000000 XPD_HALL=0x10000000 XPD_HALL_FORCE=0x20000000 | - |
| `SENS_SAR_I2C_CTRL` | 0x3F408858 | 32 | SAR_I2C_CTRL=0xFFFFFFF SAR_I2C_START=0x10000000 SAR_I2C_START_FORCE=0x20000000 | - |
| `SENS_SAR_IO_MUX_CONF` | 0x3F408944 | 32 | IOMUX_CLK_GATE_EN=0x80000000 IOMUX_RESET=0x40000000 | - |
| `SENS_SAR_MEAS1_CTRL1` | 0x3F408808 | 32 | AMP_RST_FB_FORCE=0xC000000 AMP_SHORT_REF_FORCE=0x30000000 AMP_SHORT_REF_GND_FORCE=0xC0000000 FORCE_XPD_AMP=0x3000000 RTC_SARADC_CLKGATE_EN=0x800000 RTC_SARADC_RESET=0x400000 | - |
| `SENS_SAR_MEAS1_CTRL2` | 0x3F40880C | 32 | MEAS1_DATA_SAR=0xFFFF MEAS1_DONE_SAR=0x10000 MEAS1_START_FORCE=0x40000 MEAS1_START_SAR=0x20000 SAR1_EN_PAD=0x7FF80000 SAR1_EN_PAD_FORCE=0x80000000 | - |
| `SENS_SAR_MEAS1_MUX` | 0x3F408810 | 32 | SAR1_DIG_FORCE=0x80000000 | - |
| `SENS_SAR_MEAS2_CTRL1` | 0x3F40882C | 32 | SAR2_CNTL_STATE=0x7 SAR2_EN_TEST=0x20 SAR2_PKDET_CAL_EN=0x10 SAR2_PWDET_CAL_EN=0x8 SAR2_RSTB_FORCE=0xC0 SAR2_RSTB_WAIT=0xFF0000 SAR2_STANDBY_WAIT=0xFF00 SAR2_XPD_WAIT=0xFF000000 | - |
| `SENS_SAR_MEAS2_CTRL2` | 0x3F408830 | 32 | MEAS2_DATA_SAR=0xFFFF MEAS2_DONE_SAR=0x10000 MEAS2_START_FORCE=0x40000 MEAS2_START_SAR=0x20000 SAR2_EN_PAD=0x7FF80000 SAR2_EN_PAD_FORCE=0x80000000 | - |
| `SENS_SAR_MEAS2_MUX` | 0x3F408834 | 32 | SAR2_PWDET_CCT=0x70000000 SAR2_RTC_FORCE=0x80000000 | - |
| `SENS_SAR_NOUSE` | 0x3F408940 | 32 | SAR_NOUSE=0xFFFFFFFF | - |
| `SENS_SAR_POWER_XPD_SAR` | 0x3F40883C | 32 | FORCE_XPD_SAR=0x60000000 SARCLK_EN=0x80000000 | - |
| `SENS_SAR_READER1_CTRL` | 0x3F408800 | 32 | SAR1_CLK_DIV=0xFF SAR1_CLK_GATED=0x40000 SAR1_DATA_INV=0x10000000 SAR1_INT_EN=0x20000000 SAR1_SAMPLE_NUM=0x7F80000 | - |
| `SENS_SAR_READER1_STATUS` | 0x3F408804 | 32 | SAR1_READER_STATUS=0xFFFFFFFF | - |
| `SENS_SAR_READER2_CTRL` | 0x3F408824 | 32 | SAR2_CLK_DIV=0xFF SAR2_CLK_GATED=0x40000 SAR2_DATA_INV=0x20000000 SAR2_INT_EN=0x40000000 SAR2_SAMPLE_NUM=0x7F80000 SAR2_WAIT_ARB_CYCLE=0x30000 | - |
| `SENS_SAR_READER2_STATUS` | 0x3F408828 | 32 | SAR2_READER_STATUS=0xFFFFFFFF | - |
| `SENS_SAR_SLAVE_ADDR1` | 0x3F408840 | 32 | I2C_SLAVE_ADDR0=0x3FF800 I2C_SLAVE_ADDR1=0x7FF MEAS_STATUS=0x3FC00000 | - |
| `SENS_SAR_SLAVE_ADDR2` | 0x3F408844 | 32 | I2C_SLAVE_ADDR2=0x3FF800 I2C_SLAVE_ADDR3=0x7FF | - |
| `SENS_SAR_SLAVE_ADDR3` | 0x3F408848 | 32 | I2C_SLAVE_ADDR4=0x3FF800 I2C_SLAVE_ADDR5=0x7FF | - |
| `SENS_SAR_SLAVE_ADDR4` | 0x3F40884C | 32 | I2C_SLAVE_ADDR6=0x3FF800 I2C_SLAVE_ADDR7=0x7FF | - |
| `SENS_SAR_TOUCH_CHN_ST` | 0x3F4088D4 | 32 | TOUCH_CHANNEL_CLR=0x3FFF8000 TOUCH_MEAS_DONE=0x80000000 TOUCH_PAD_ACTIVE=0x7FFF | - |
| `SENS_SAR_TOUCH_CONF` | 0x3F40885C | 32 | TOUCH_APPROACH_PAD0=0xF0000000 TOUCH_APPROACH_PAD1=0xF000000 TOUCH_APPROACH_PAD2=0xF00000 TOUCH_DATA_SEL=0x30000 TOUCH_DENOISE_END=0x40000 TOUCH_OUTEN=0x7FFF TOUCH_STATUS_CLR=0x8000 TOUCH_UNIT_END=0x80000 | - |
| `SENS_SAR_TOUCH_STATUS0` | 0x3F4088D8 | 32 | TOUCH_DENOISE_DATA=0x3FFFFF TOUCH_SCAN_CURR=0x3C00000 | - |
| `SENS_SAR_TOUCH_STATUS1` | 0x3F4088DC | 32 | TOUCH_PAD1_DATA=0x3FFFFF TOUCH_PAD1_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS10` | 0x3F408900 | 32 | TOUCH_PAD10_DATA=0x3FFFFF TOUCH_PAD10_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS11` | 0x3F408904 | 32 | TOUCH_PAD11_DATA=0x3FFFFF TOUCH_PAD11_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS12` | 0x3F408908 | 32 | TOUCH_PAD12_DATA=0x3FFFFF TOUCH_PAD12_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS13` | 0x3F40890C | 32 | TOUCH_PAD13_DATA=0x3FFFFF TOUCH_PAD13_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS14` | 0x3F408910 | 32 | TOUCH_PAD14_DATA=0x3FFFFF TOUCH_PAD14_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS15` | 0x3F408914 | 32 | TOUCH_SLP_DATA=0x3FFFFF TOUCH_SLP_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS16` | 0x3F408918 | 32 | TOUCH_APPROACH_PAD0_CNT=0xFF0000 TOUCH_APPROACH_PAD1_CNT=0xFF00 TOUCH_APPROACH_PAD2_CNT=0xFF TOUCH_SLP_APPROACH_CNT=0xFF000000 | - |
| `SENS_SAR_TOUCH_STATUS2` | 0x3F4088E0 | 32 | TOUCH_PAD2_DATA=0x3FFFFF TOUCH_PAD2_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS3` | 0x3F4088E4 | 32 | TOUCH_PAD3_DATA=0x3FFFFF TOUCH_PAD3_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS4` | 0x3F4088E8 | 32 | TOUCH_PAD4_DATA=0x3FFFFF TOUCH_PAD4_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS5` | 0x3F4088EC | 32 | TOUCH_PAD5_DATA=0x3FFFFF TOUCH_PAD5_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS6` | 0x3F4088F0 | 32 | TOUCH_PAD6_DATA=0x3FFFFF TOUCH_PAD6_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS7` | 0x3F4088F4 | 32 | TOUCH_PAD7_DATA=0x3FFFFF TOUCH_PAD7_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS8` | 0x3F4088F8 | 32 | TOUCH_PAD8_DATA=0x3FFFFF TOUCH_PAD8_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_STATUS9` | 0x3F4088FC | 32 | TOUCH_PAD9_DATA=0x3FFFFF TOUCH_PAD9_DEBOUNCE=0xE0000000 | - |
| `SENS_SAR_TOUCH_THRES1` | 0x3F408860 | 32 | TOUCH_OUT_TH1=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES10` | 0x3F408884 | 32 | TOUCH_OUT_TH10=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES11` | 0x3F408888 | 32 | TOUCH_OUT_TH11=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES12` | 0x3F40888C | 32 | TOUCH_OUT_TH12=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES13` | 0x3F408890 | 32 | TOUCH_OUT_TH13=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES14` | 0x3F408894 | 32 | TOUCH_OUT_TH14=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES2` | 0x3F408864 | 32 | TOUCH_OUT_TH2=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES3` | 0x3F408868 | 32 | TOUCH_OUT_TH3=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES4` | 0x3F40886C | 32 | TOUCH_OUT_TH4=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES5` | 0x3F408870 | 32 | TOUCH_OUT_TH5=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES6` | 0x3F408874 | 32 | TOUCH_OUT_TH6=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES7` | 0x3F408878 | 32 | TOUCH_OUT_TH7=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES8` | 0x3F40887C | 32 | TOUCH_OUT_TH8=0x3FFFFF | - |
| `SENS_SAR_TOUCH_THRES9` | 0x3F408880 | 32 | TOUCH_OUT_TH9=0x3FFFFF | - |
| `SENS_SAR_TSENS_CTRL` | 0x3F408850 | 32 | TSENS_CLK_DIV=0x3FC000 TSENS_DUMP_OUT=0x1000000 TSENS_INT_EN=0x1000 TSENS_IN_INV=0x2000 TSENS_OUT=0xFF TSENS_POWER_UP=0x400000 TSENS_POWER_UP_FORCE=0x800000 TSENS_READY=0x100 | - |
| `SENS_SAR_TSENS_CTRL2` | 0x3F408854 | 32 | TSENS_CLKGATE_EN=0x8000 TSENS_CLK_INV=0x4000 TSENS_RESET=0x10000 TSENS_XPD_FORCE=0x3000 TSENS_XPD_WAIT=0xFFF | - |
| `SHA_BUSY` | 0x6003B018 | 32 | STATE=0x1 | - |
| `SHA_CONTINUE_OP` | 0x6003B014 | 32 | CONTINUE_OP=0x1 | - |
| `SHA_DATE` | 0x6003B02C | 32 | DATE=0x3FFFFFFF | - |
| `SHA_DMA_BLOCK_NUM` | 0x6003B00C | 32 | DMA_BLOCK_NUM=0x3F | - |
| `SHA_DMA_CONTINUE` | 0x6003B020 | 32 | DMA_CONTINUE=0x1 | - |
| `SHA_DMA_START` | 0x6003B01C | 32 | DMA_START=0x1 | - |
| `SHA_H_0` | 0x6003B040 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_1` | 0x6003B044 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_10` | 0x6003B068 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_11` | 0x6003B06C | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_12` | 0x6003B070 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_13` | 0x6003B074 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_14` | 0x6003B078 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_15` | 0x6003B07C | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_2` | 0x6003B048 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_3` | 0x6003B04C | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_4` | 0x6003B050 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_5` | 0x6003B054 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_6` | 0x6003B058 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_7` | 0x6003B05C | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_8` | 0x6003B060 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_H_9` | 0x6003B064 | 32 | H_0=0xFFFFFFFF | - |
| `SHA_INT_CLEAR` | 0x6003B024 | 32 | CLEAR_INTERRUPT=0x1 | - |
| `SHA_INT_ENA` | 0x6003B028 | 32 | INTERRUPT_ENA=0x1 | - |
| `SHA_MODE` | 0x6003B000 | 32 | MODE=0x7 | - |
| `SHA_M_0` | 0x6003B080 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_1` | 0x6003B084 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_10` | 0x6003B0A8 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_11` | 0x6003B0AC | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_12` | 0x6003B0B0 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_13` | 0x6003B0B4 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_14` | 0x6003B0B8 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_15` | 0x6003B0BC | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_16` | 0x6003B0C0 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_17` | 0x6003B0C4 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_18` | 0x6003B0C8 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_19` | 0x6003B0CC | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_2` | 0x6003B088 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_20` | 0x6003B0D0 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_21` | 0x6003B0D4 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_22` | 0x6003B0D8 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_23` | 0x6003B0DC | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_24` | 0x6003B0E0 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_25` | 0x6003B0E4 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_26` | 0x6003B0E8 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_27` | 0x6003B0EC | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_28` | 0x6003B0F0 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_29` | 0x6003B0F4 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_3` | 0x6003B08C | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_30` | 0x6003B0F8 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_31` | 0x6003B0FC | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_4` | 0x6003B090 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_5` | 0x6003B094 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_6` | 0x6003B098 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_7` | 0x6003B09C | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_8` | 0x6003B0A0 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_M_9` | 0x6003B0A4 | 32 | M_0=0xFFFFFFFF | - |
| `SHA_START` | 0x6003B010 | 32 | START=0x1 | - |
| `SHA_T_LENGTH` | 0x6003B008 | 32 | T_LENGTH=0x3F | - |
| `SHA_T_STRING` | 0x6003B004 | 32 | T_STRING=0xFFFFFFFF | - |
| `SPI0_ADDR` | 0x3F403004 | 32 | USR_ADDR_VALUE=0xFFFFFFFF | - |
| `SPI0_CLOCK` | 0x3F403014 | 32 | CLKCNT_H=0xFC0 CLKCNT_L=0x3F CLKCNT_N=0x3F000 CLKDIV_PRE=0x7FFC0000 CLK_EQU_SYSCLK=0x80000000 | - |
| `SPI0_CMD` | 0x3F403000 | 32 | CONF_BITLEN=0x7FFFFF USR=0x1000000 | - |
| `SPI0_CTRL` | 0x3F403008 | 32 | DUMMY_OUT=0x8 D_POL=0x80000 EXT_HOLD_EN=0x4 FADDR_DUAL=0x20 FADDR_OCT=0x80 FADDR_QUAD=0x40 FCMD_DUAL=0x100 FCMD_OCT=0x400 FCMD_QUAD=0x200 FREAD_DUAL=0x4000 FREAD_OCT=0x10000 FREAD_QUAD=0x8000 Q_POL=0x40000 RD_BIT_ORDER=0x2000000 WP=0x200000... | - |
| `SPI0_CTRL1` | 0x3F40300C | 32 | CLK_MODE=0x3 CLK_MODE_13=0x4 CS_HOLD_DELAY=0xFC000 RSCK_DATA_OUT=0x8 W16_17_WR_ENA=0x10 | - |
| `SPI0_CTRL2` | 0x3F403010 | 32 | CS_DELAY_MODE=0x1C000000 CS_DELAY_NUM=0x60000000 CS_HOLD_TIME=0x3FFE000 CS_SETUP_TIME=0x1FFF | - |
| `SPI0_DIN_MODE` | 0x3F4030E0 | 32 | DIN0_MODE=0x7 DIN1_MODE=0x38 DIN2_MODE=0x1C0 DIN3_MODE=0xE00 DIN4_MODE=0x7000 DIN5_MODE=0x38000 DIN6_MODE=0x1C0000 DIN7_MODE=0xE00000 TIMING_CLK_ENA=0x1000000 | - |
| `SPI0_DIN_NUM` | 0x3F4030E4 | 32 | DIN0_NUM=0x3 DIN1_NUM=0xC DIN2_NUM=0x30 DIN3_NUM=0xC0 DIN4_NUM=0x300 DIN5_NUM=0xC00 DIN6_NUM=0x3000 DIN7_NUM=0xC000 | - |
| `SPI0_DMA_CONF` | 0x3F40304C | 32 | AHBM_FIFO_RST=0x10 AHBM_RST=0x20 DMA_CONTINUE=0x10000 DMA_INFIFO_FULL_CLR=0x400000 DMA_OUTFIFO_EMPTY_CLR=0x800000 DMA_RX_STOP=0x4000 DMA_SEG_TRANS_CLR=0x10000000 DMA_SLV_SEG_TRANS_EN=0x40000 DMA_TX_STOP=0x8000 EXT_MEM_BK_SIZE=0xC000000 INDS... | - |
| `SPI0_DMA_INSTATUS` | 0x3F403094 | 32 | DMA_INDSCR_ADDR=0x3FFFF DMA_INDSCR_STATE=0xC0000 DMA_INFIFO_CNT=0x3F800000 DMA_INFIFO_EMPTY=0x80000000 DMA_INFIFO_FULL=0x40000000 DMA_IN_STATE=0x700000 | - |
| `SPI0_DMA_INT_CLR` | 0x3F403064 | 32 | INFIFO_FULL_ERR_INT_CLR=0x200 INLINK_DSCR_EMPTY_INT_CLR=0x1 INLINK_DSCR_ERROR_INT_CLR=0x4 IN_DONE_INT_CLR=0x8 IN_ERR_EOF_INT_CLR=0x10 IN_SUC_EOF_INT_CLR=0x20 OUTFIFO_EMPTY_ERR_INT_CLR=0x400 OUTLINK_DSCR_ERROR_INT_CLR=0x2 OUT_DONE_INT_CLR=0x... | - |
| `SPI0_DMA_INT_ENA` | 0x3F403058 | 32 | INFIFO_FULL_ERR_INT_ENA=0x200 INLINK_DSCR_EMPTY_INT_ENA=0x1 INLINK_DSCR_ERROR_INT_ENA=0x4 IN_DONE_INT_ENA=0x8 IN_ERR_EOF_INT_ENA=0x10 IN_SUC_EOF_INT_ENA=0x20 OUTFIFO_EMPTY_ERR_INT_ENA=0x400 OUTLINK_DSCR_ERROR_INT_ENA=0x2 OUT_DONE_INT_ENA=0x... | - |
| `SPI0_DMA_INT_RAW` | 0x3F40305C | 32 | INFIFO_FULL_ERR_INT_RAW=0x200 INLINK_DSCR_EMPTY_INT_RAW=0x1 INLINK_DSCR_ERROR_INT_RAW=0x4 IN_DONE_INT_RAW=0x8 IN_ERR_EOF_INT_RAW=0x10 IN_SUC_EOF_INT_RAW=0x20 OUTFIFO_EMPTY_ERR_INT_RAW=0x400 OUTLINK_DSCR_ERROR_INT_RAW=0x2 OUT_DONE_INT_RAW=0x... | - |
| `SPI0_DMA_INT_ST` | 0x3F403060 | 32 | INFIFO_FULL_ERR_INT_ST=0x200 INLINK_DSCR_EMPTY_INT_ST=0x1 INLINK_DSCR_ERROR_INT_ST=0x4 IN_DONE_INT_ST=0x8 IN_ERR_EOF_INT_ST=0x10 IN_SUC_EOF_INT_ST=0x20 OUTFIFO_EMPTY_ERR_INT_ST=0x400 OUTLINK_DSCR_ERROR_INT_ST=0x2 OUT_DONE_INT_ST=0x40 OUT_EO... | - |
| `SPI0_DMA_IN_LINK` | 0x3F403054 | 32 | DMA_RX_ENA=0x80000000 INLINK_ADDR=0xFFFFF INLINK_AUTO_RET=0x100000 INLINK_RESTART=0x40000000 INLINK_START=0x20000000 INLINK_STOP=0x10000000 | - |
| `SPI0_DMA_OUTSTATUS` | 0x3F403090 | 32 | DMA_OUTDSCR_ADDR=0x3FFFF DMA_OUTDSCR_STATE=0xC0000 DMA_OUTFIFO_CNT=0x3F800000 DMA_OUTFIFO_EMPTY=0x80000000 DMA_OUTFIFO_FULL=0x40000000 DMA_OUT_STATE=0x700000 | - |
| `SPI0_DMA_OUT_LINK` | 0x3F403050 | 32 | DMA_TX_ENA=0x80000000 OUTLINK_ADDR=0xFFFFF OUTLINK_RESTART=0x40000000 OUTLINK_START=0x20000000 OUTLINK_STOP=0x10000000 | - |
| `SPI0_DOUT_MODE` | 0x3F4030E8 | 32 | DOUT0_MODE=0x7 DOUT1_MODE=0x38 DOUT2_MODE=0x1C0 DOUT3_MODE=0xE00 DOUT4_MODE=0x7000 DOUT5_MODE=0x38000 DOUT6_MODE=0x1C0000 DOUT7_MODE=0xE00000 | - |
| `SPI0_DOUT_NUM` | 0x3F4030EC | 32 | DOUT0_NUM=0x3 DOUT1_NUM=0xC DOUT2_NUM=0x30 DOUT3_NUM=0xC0 DOUT4_NUM=0x300 DOUT5_NUM=0xC00 DOUT6_NUM=0x3000 DOUT7_NUM=0xC000 | - |
| `SPI0_FSM` | 0x3F403044 | 32 | MST_DMA_RD_BYTELEN=0xFFFFF000 ST=0xF | - |
| `SPI0_HOLD` | 0x3F403048 | 32 | DMA_SEG_TRANS_DONE=0x80 INT_HOLD_ENA=0x3 OUT_EN=0x8 OUT_TIME=0x70 VAL=0x4 | - |
| `SPI0_INLINK_DSCR` | 0x3F403070 | 32 | DMA_INLINK_DSCR=0xFFFFFFFF | - |
| `SPI0_INLINK_DSCR_BF0` | 0x3F403074 | 32 | DMA_INLINK_DSCR_BF0=0xFFFFFFFF | - |
| `SPI0_INLINK_DSCR_BF1` | 0x3F403078 | 32 | DMA_INLINK_DSCR_BF1=0xFFFFFFFF | - |
| `SPI0_IN_ERR_EOF_DES_ADDR` | 0x3F403068 | 32 | DMA_IN_ERR_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI0_IN_SUC_EOF_DES_ADDR` | 0x3F40306C | 32 | DMA_IN_SUC_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI0_LCD_CTRL` | 0x3F4030F0 | 32 | LCD_HB_FRONT=0x7FF LCD_MODE_EN=0x80000000 LCD_VA_HEIGHT=0x1FF800 LCD_VT_HEIGHT=0x7FE00000 | - |
| `SPI0_LCD_CTRL1` | 0x3F4030F4 | 32 | LCD_HA_WIDTH=0xFFF00 LCD_HT_WIDTH=0xFFF00000 LCD_VB_FRONT=0xFF | - |
| `SPI0_LCD_CTRL2` | 0x3F4030F8 | 32 | HSYNC_IDLE_POL=0x800000 LCD_HSYNC_POSITION=0xFF000000 LCD_HSYNC_WIDTH=0x7F0000 LCD_VSYNC_WIDTH=0x7F VSYNC_IDLE_POL=0x80 | - |
| `SPI0_LCD_D_MODE` | 0x3F4030FC | 32 | DE_IDLE_POL=0x8000 D_CD_MODE=0x38 D_DE_MODE=0x1C0 D_DQS_MODE=0x7 D_HSYNC_MODE=0xE00 D_VSYNC_MODE=0x7000 HS_BLANK_EN=0x10000 | - |
| `SPI0_LCD_D_NUM` | 0x3F403100 | 32 | D_CD_NUM=0xC D_DE_NUM=0x30 D_DQS_NUM=0x3 D_HSYNC_NUM=0xC0 D_VSYNC_NUM=0x300 | - |
| `SPI0_MISC` | 0x3F40302C | 32 | ADDR_DTR_EN=0x40000 CD_ADDR_SET=0x400000 CD_CMD_SET=0x2000000 CD_DATA_SET=0x100000 CD_DUMMY_SET=0x200000 CD_IDLE_EDGE=0x4000000 CK_DIS=0x40 CK_IDLE_EDGE=0x20000000 CLK_DATA_DTR_EN=0x10000 CMD_DTR_EN=0x80000 CS0_DIS=0x1 CS1_DIS=0x2 CS2_DIS=0... | - |
| `SPI0_MISO_DLEN` | 0x3F403028 | 32 | USR_MISO_DBITLEN=0x7FFFFF | - |
| `SPI0_MOSI_DLEN` | 0x3F403024 | 32 | USR_MOSI_DBITLEN=0x7FFFFF | - |
| `SPI0_OUTLINK_DSCR` | 0x3F403084 | 32 | DMA_OUTLINK_DSCR=0xFFFFFFFF | - |
| `SPI0_OUTLINK_DSCR_BF0` | 0x3F403088 | 32 | DMA_OUTLINK_DSCR_BF0=0xFFFFFFFF | - |
| `SPI0_OUTLINK_DSCR_BF1` | 0x3F40308C | 32 | DMA_OUTLINK_DSCR_BF1=0xFFFFFFFF | - |
| `SPI0_OUT_EOF_BFR_DES_ADDR` | 0x3F40307C | 32 | DMA_OUT_EOF_BFR_DES_ADDR=0xFFFFFFFF | - |
| `SPI0_OUT_EOF_DES_ADDR` | 0x3F403080 | 32 | DMA_OUT_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI0_REG_DATE` | 0x3F4033FC | 32 | DATE=0xFFFFFFF | - |
| `SPI0_SLAVE` | 0x3F403030 | 32 | INT_DMA_SEG_TRANS_EN=0x400 INT_RD_BUF_DONE_EN=0x20 INT_RD_DMA_DONE_EN=0x80 INT_TRANS_DONE_EN=0x200 INT_WR_BUF_DONE_EN=0x40 INT_WR_DMA_DONE_EN=0x100 MODE=0x40000000 SEG_MAGIC_ERR_INT_EN=0x800 SOFT_RESET=0x80000000 TRANS_CNT=0x7800000 TRANS_D... | - |
| `SPI0_SLAVE1` | 0x3F403034 | 32 | SLV_ADDR_ERR=0x2000 SLV_ADDR_ERR_CLR=0x400 SLV_CMD_ERR=0x4000 SLV_CMD_ERR_CLR=0x800 SLV_LAST_ADDR=0xFF000000 SLV_LAST_COMMAND=0xFF0000 SLV_NO_QPI_EN=0x1000 SLV_WR_DMA_DONE=0x8000 | - |
| `SPI0_SLV_RDBUF_DLEN` | 0x3F40303C | 32 | SEG_MAGIC_ERR=0x2000000 SLV_DMA_RD_BYTELEN=0xFFFFF SLV_RD_BUF_DONE=0x1000000 | - |
| `SPI0_SLV_RD_BYTE` | 0x3F403040 | 32 | DMA_SEG_MAGIC_VALUE=0xF000000 SLV_DATA_BYTELEN=0xFFFFF SLV_RDBUF_BYTELEN_EN=0x400000 SLV_RDDMA_BYTELEN_EN=0x100000 SLV_RD_DMA_DONE=0x40000000 SLV_WRBUF_BYTELEN_EN=0x800000 SLV_WRDMA_BYTELEN_EN=0x200000 USR_CONF=0x80000000 | - |
| `SPI0_SLV_WRBUF_DLEN` | 0x3F403038 | 32 | CONF_BASE_BITLEN=0xFE000000 SLV_WR_BUF_DONE=0x1000000 | - |
| `SPI0_USER` | 0x3F403018 | 32 | CK_OUT_EDGE=0x200 CS_HOLD=0x40 CS_SETUP=0x80 DOUTDIN=0x1 FWRITE_DUAL=0x1000 FWRITE_OCT=0x4000 FWRITE_QUAD=0x2000 OPI_MODE=0x10 QPI_MODE=0x8 RD_BYTE_ORDER=0x400 RSCK_I_EDGE=0x100 SIO=0x10000 TSCK_I_EDGE=0x20 USR_ADDR=0x40000000 USR_ADDR_HOLD... | - |
| `SPI0_USER1` | 0x3F40301C | 32 | USR_ADDR_BITLEN=0xF8000000 USR_DUMMY_CYCLELEN=0xFF | - |
| `SPI0_USER2` | 0x3F403020 | 32 | USR_COMMAND_BITLEN=0xF0000000 USR_COMMAND_VALUE=0xFFFF | - |
| `SPI0_W0` | 0x3F403098 | 32 | BUF0=0xFFFFFFFF | - |
| `SPI0_W1` | 0x3F40309C | 32 | BUF1=0xFFFFFFFF | - |
| `SPI0_W10` | 0x3F4030C0 | 32 | BUF10=0xFFFFFFFF | - |
| `SPI0_W11` | 0x3F4030C4 | 32 | BUF11=0xFFFFFFFF | - |
| `SPI0_W12` | 0x3F4030C8 | 32 | BUF12=0xFFFFFFFF | - |
| `SPI0_W13` | 0x3F4030CC | 32 | BUF13=0xFFFFFFFF | - |
| `SPI0_W14` | 0x3F4030D0 | 32 | BUF14=0xFFFFFFFF | - |
| `SPI0_W15` | 0x3F4030D4 | 32 | BUF15=0xFFFFFFFF | - |
| `SPI0_W16` | 0x3F4030D8 | 32 | BUF16=0xFFFFFFFF | - |
| `SPI0_W17` | 0x3F4030DC | 32 | BUF17=0xFFFFFFFF | - |
| `SPI0_W2` | 0x3F4030A0 | 32 | BUF2=0xFFFFFFFF | - |
| `SPI0_W3` | 0x3F4030A4 | 32 | BUF3=0xFFFFFFFF | - |
| `SPI0_W4` | 0x3F4030A8 | 32 | BUF4=0xFFFFFFFF | - |
| `SPI0_W5` | 0x3F4030AC | 32 | BUF5=0xFFFFFFFF | - |
| `SPI0_W6` | 0x3F4030B0 | 32 | BUF6=0xFFFFFFFF | - |
| `SPI0_W7` | 0x3F4030B4 | 32 | BUF7=0xFFFFFFFF | - |
| `SPI0_W8` | 0x3F4030B8 | 32 | BUF8=0xFFFFFFFF | - |
| `SPI0_W9` | 0x3F4030BC | 32 | BUF9=0xFFFFFFFF | - |
| `SPI1_ADDR` | 0x3F402004 | 32 | USR_ADDR_VALUE=0xFFFFFFFF | - |
| `SPI1_CLOCK` | 0x3F402014 | 32 | CLKCNT_H=0xFC0 CLKCNT_L=0x3F CLKCNT_N=0x3F000 CLKDIV_PRE=0x7FFC0000 CLK_EQU_SYSCLK=0x80000000 | - |
| `SPI1_CMD` | 0x3F402000 | 32 | CONF_BITLEN=0x7FFFFF USR=0x1000000 | - |
| `SPI1_CTRL` | 0x3F402008 | 32 | DUMMY_OUT=0x8 D_POL=0x80000 EXT_HOLD_EN=0x4 FADDR_DUAL=0x20 FADDR_OCT=0x80 FADDR_QUAD=0x40 FCMD_DUAL=0x100 FCMD_OCT=0x400 FCMD_QUAD=0x200 FREAD_DUAL=0x4000 FREAD_OCT=0x10000 FREAD_QUAD=0x8000 Q_POL=0x40000 RD_BIT_ORDER=0x2000000 WP=0x200000... | - |
| `SPI1_CTRL1` | 0x3F40200C | 32 | CLK_MODE=0x3 CLK_MODE_13=0x4 CS_HOLD_DELAY=0xFC000 RSCK_DATA_OUT=0x8 W16_17_WR_ENA=0x10 | - |
| `SPI1_CTRL2` | 0x3F402010 | 32 | CS_DELAY_MODE=0x1C000000 CS_DELAY_NUM=0x60000000 CS_HOLD_TIME=0x3FFE000 CS_SETUP_TIME=0x1FFF | - |
| `SPI1_DIN_MODE` | 0x3F4020E0 | 32 | DIN0_MODE=0x7 DIN1_MODE=0x38 DIN2_MODE=0x1C0 DIN3_MODE=0xE00 DIN4_MODE=0x7000 DIN5_MODE=0x38000 DIN6_MODE=0x1C0000 DIN7_MODE=0xE00000 TIMING_CLK_ENA=0x1000000 | - |
| `SPI1_DIN_NUM` | 0x3F4020E4 | 32 | DIN0_NUM=0x3 DIN1_NUM=0xC DIN2_NUM=0x30 DIN3_NUM=0xC0 DIN4_NUM=0x300 DIN5_NUM=0xC00 DIN6_NUM=0x3000 DIN7_NUM=0xC000 | - |
| `SPI1_DMA_CONF` | 0x3F40204C | 32 | AHBM_FIFO_RST=0x10 AHBM_RST=0x20 DMA_CONTINUE=0x10000 DMA_INFIFO_FULL_CLR=0x400000 DMA_OUTFIFO_EMPTY_CLR=0x800000 DMA_RX_STOP=0x4000 DMA_SEG_TRANS_CLR=0x10000000 DMA_SLV_SEG_TRANS_EN=0x40000 DMA_TX_STOP=0x8000 EXT_MEM_BK_SIZE=0xC000000 INDS... | - |
| `SPI1_DMA_INSTATUS` | 0x3F402094 | 32 | DMA_INDSCR_ADDR=0x3FFFF DMA_INDSCR_STATE=0xC0000 DMA_INFIFO_CNT=0x3F800000 DMA_INFIFO_EMPTY=0x80000000 DMA_INFIFO_FULL=0x40000000 DMA_IN_STATE=0x700000 | - |
| `SPI1_DMA_INT_CLR` | 0x3F402064 | 32 | INFIFO_FULL_ERR_INT_CLR=0x200 INLINK_DSCR_EMPTY_INT_CLR=0x1 INLINK_DSCR_ERROR_INT_CLR=0x4 IN_DONE_INT_CLR=0x8 IN_ERR_EOF_INT_CLR=0x10 IN_SUC_EOF_INT_CLR=0x20 OUTFIFO_EMPTY_ERR_INT_CLR=0x400 OUTLINK_DSCR_ERROR_INT_CLR=0x2 OUT_DONE_INT_CLR=0x... | - |
| `SPI1_DMA_INT_ENA` | 0x3F402058 | 32 | INFIFO_FULL_ERR_INT_ENA=0x200 INLINK_DSCR_EMPTY_INT_ENA=0x1 INLINK_DSCR_ERROR_INT_ENA=0x4 IN_DONE_INT_ENA=0x8 IN_ERR_EOF_INT_ENA=0x10 IN_SUC_EOF_INT_ENA=0x20 OUTFIFO_EMPTY_ERR_INT_ENA=0x400 OUTLINK_DSCR_ERROR_INT_ENA=0x2 OUT_DONE_INT_ENA=0x... | - |
| `SPI1_DMA_INT_RAW` | 0x3F40205C | 32 | INFIFO_FULL_ERR_INT_RAW=0x200 INLINK_DSCR_EMPTY_INT_RAW=0x1 INLINK_DSCR_ERROR_INT_RAW=0x4 IN_DONE_INT_RAW=0x8 IN_ERR_EOF_INT_RAW=0x10 IN_SUC_EOF_INT_RAW=0x20 OUTFIFO_EMPTY_ERR_INT_RAW=0x400 OUTLINK_DSCR_ERROR_INT_RAW=0x2 OUT_DONE_INT_RAW=0x... | - |
| `SPI1_DMA_INT_ST` | 0x3F402060 | 32 | INFIFO_FULL_ERR_INT_ST=0x200 INLINK_DSCR_EMPTY_INT_ST=0x1 INLINK_DSCR_ERROR_INT_ST=0x4 IN_DONE_INT_ST=0x8 IN_ERR_EOF_INT_ST=0x10 IN_SUC_EOF_INT_ST=0x20 OUTFIFO_EMPTY_ERR_INT_ST=0x400 OUTLINK_DSCR_ERROR_INT_ST=0x2 OUT_DONE_INT_ST=0x40 OUT_EO... | - |
| `SPI1_DMA_IN_LINK` | 0x3F402054 | 32 | DMA_RX_ENA=0x80000000 INLINK_ADDR=0xFFFFF INLINK_AUTO_RET=0x100000 INLINK_RESTART=0x40000000 INLINK_START=0x20000000 INLINK_STOP=0x10000000 | - |
| `SPI1_DMA_OUTSTATUS` | 0x3F402090 | 32 | DMA_OUTDSCR_ADDR=0x3FFFF DMA_OUTDSCR_STATE=0xC0000 DMA_OUTFIFO_CNT=0x3F800000 DMA_OUTFIFO_EMPTY=0x80000000 DMA_OUTFIFO_FULL=0x40000000 DMA_OUT_STATE=0x700000 | - |
| `SPI1_DMA_OUT_LINK` | 0x3F402050 | 32 | DMA_TX_ENA=0x80000000 OUTLINK_ADDR=0xFFFFF OUTLINK_RESTART=0x40000000 OUTLINK_START=0x20000000 OUTLINK_STOP=0x10000000 | - |
| `SPI1_DOUT_MODE` | 0x3F4020E8 | 32 | DOUT0_MODE=0x7 DOUT1_MODE=0x38 DOUT2_MODE=0x1C0 DOUT3_MODE=0xE00 DOUT4_MODE=0x7000 DOUT5_MODE=0x38000 DOUT6_MODE=0x1C0000 DOUT7_MODE=0xE00000 | - |
| `SPI1_DOUT_NUM` | 0x3F4020EC | 32 | DOUT0_NUM=0x3 DOUT1_NUM=0xC DOUT2_NUM=0x30 DOUT3_NUM=0xC0 DOUT4_NUM=0x300 DOUT5_NUM=0xC00 DOUT6_NUM=0x3000 DOUT7_NUM=0xC000 | - |
| `SPI1_FSM` | 0x3F402044 | 32 | MST_DMA_RD_BYTELEN=0xFFFFF000 ST=0xF | - |
| `SPI1_HOLD` | 0x3F402048 | 32 | DMA_SEG_TRANS_DONE=0x80 INT_HOLD_ENA=0x3 OUT_EN=0x8 OUT_TIME=0x70 VAL=0x4 | - |
| `SPI1_INLINK_DSCR` | 0x3F402070 | 32 | DMA_INLINK_DSCR=0xFFFFFFFF | - |
| `SPI1_INLINK_DSCR_BF0` | 0x3F402074 | 32 | DMA_INLINK_DSCR_BF0=0xFFFFFFFF | - |
| `SPI1_INLINK_DSCR_BF1` | 0x3F402078 | 32 | DMA_INLINK_DSCR_BF1=0xFFFFFFFF | - |
| `SPI1_IN_ERR_EOF_DES_ADDR` | 0x3F402068 | 32 | DMA_IN_ERR_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI1_IN_SUC_EOF_DES_ADDR` | 0x3F40206C | 32 | DMA_IN_SUC_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI1_LCD_CTRL` | 0x3F4020F0 | 32 | LCD_HB_FRONT=0x7FF LCD_MODE_EN=0x80000000 LCD_VA_HEIGHT=0x1FF800 LCD_VT_HEIGHT=0x7FE00000 | - |
| `SPI1_LCD_CTRL1` | 0x3F4020F4 | 32 | LCD_HA_WIDTH=0xFFF00 LCD_HT_WIDTH=0xFFF00000 LCD_VB_FRONT=0xFF | - |
| `SPI1_LCD_CTRL2` | 0x3F4020F8 | 32 | HSYNC_IDLE_POL=0x800000 LCD_HSYNC_POSITION=0xFF000000 LCD_HSYNC_WIDTH=0x7F0000 LCD_VSYNC_WIDTH=0x7F VSYNC_IDLE_POL=0x80 | - |
| `SPI1_LCD_D_MODE` | 0x3F4020FC | 32 | DE_IDLE_POL=0x8000 D_CD_MODE=0x38 D_DE_MODE=0x1C0 D_DQS_MODE=0x7 D_HSYNC_MODE=0xE00 D_VSYNC_MODE=0x7000 HS_BLANK_EN=0x10000 | - |
| `SPI1_LCD_D_NUM` | 0x3F402100 | 32 | D_CD_NUM=0xC D_DE_NUM=0x30 D_DQS_NUM=0x3 D_HSYNC_NUM=0xC0 D_VSYNC_NUM=0x300 | - |
| `SPI1_MISC` | 0x3F40202C | 32 | ADDR_DTR_EN=0x40000 CD_ADDR_SET=0x400000 CD_CMD_SET=0x2000000 CD_DATA_SET=0x100000 CD_DUMMY_SET=0x200000 CD_IDLE_EDGE=0x4000000 CK_DIS=0x40 CK_IDLE_EDGE=0x20000000 CLK_DATA_DTR_EN=0x10000 CMD_DTR_EN=0x80000 CS0_DIS=0x1 CS1_DIS=0x2 CS2_DIS=0... | - |
| `SPI1_MISO_DLEN` | 0x3F402028 | 32 | USR_MISO_DBITLEN=0x7FFFFF | - |
| `SPI1_MOSI_DLEN` | 0x3F402024 | 32 | USR_MOSI_DBITLEN=0x7FFFFF | - |
| `SPI1_OUTLINK_DSCR` | 0x3F402084 | 32 | DMA_OUTLINK_DSCR=0xFFFFFFFF | - |
| `SPI1_OUTLINK_DSCR_BF0` | 0x3F402088 | 32 | DMA_OUTLINK_DSCR_BF0=0xFFFFFFFF | - |
| `SPI1_OUTLINK_DSCR_BF1` | 0x3F40208C | 32 | DMA_OUTLINK_DSCR_BF1=0xFFFFFFFF | - |
| `SPI1_OUT_EOF_BFR_DES_ADDR` | 0x3F40207C | 32 | DMA_OUT_EOF_BFR_DES_ADDR=0xFFFFFFFF | - |
| `SPI1_OUT_EOF_DES_ADDR` | 0x3F402080 | 32 | DMA_OUT_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI1_REG_DATE` | 0x3F4023FC | 32 | DATE=0xFFFFFFF | - |
| `SPI1_SLAVE` | 0x3F402030 | 32 | INT_DMA_SEG_TRANS_EN=0x400 INT_RD_BUF_DONE_EN=0x20 INT_RD_DMA_DONE_EN=0x80 INT_TRANS_DONE_EN=0x200 INT_WR_BUF_DONE_EN=0x40 INT_WR_DMA_DONE_EN=0x100 MODE=0x40000000 SEG_MAGIC_ERR_INT_EN=0x800 SOFT_RESET=0x80000000 TRANS_CNT=0x7800000 TRANS_D... | - |
| `SPI1_SLAVE1` | 0x3F402034 | 32 | SLV_ADDR_ERR=0x2000 SLV_ADDR_ERR_CLR=0x400 SLV_CMD_ERR=0x4000 SLV_CMD_ERR_CLR=0x800 SLV_LAST_ADDR=0xFF000000 SLV_LAST_COMMAND=0xFF0000 SLV_NO_QPI_EN=0x1000 SLV_WR_DMA_DONE=0x8000 | - |
| `SPI1_SLV_RDBUF_DLEN` | 0x3F40203C | 32 | SEG_MAGIC_ERR=0x2000000 SLV_DMA_RD_BYTELEN=0xFFFFF SLV_RD_BUF_DONE=0x1000000 | - |
| `SPI1_SLV_RD_BYTE` | 0x3F402040 | 32 | DMA_SEG_MAGIC_VALUE=0xF000000 SLV_DATA_BYTELEN=0xFFFFF SLV_RDBUF_BYTELEN_EN=0x400000 SLV_RDDMA_BYTELEN_EN=0x100000 SLV_RD_DMA_DONE=0x40000000 SLV_WRBUF_BYTELEN_EN=0x800000 SLV_WRDMA_BYTELEN_EN=0x200000 USR_CONF=0x80000000 | - |
| `SPI1_SLV_WRBUF_DLEN` | 0x3F402038 | 32 | CONF_BASE_BITLEN=0xFE000000 SLV_WR_BUF_DONE=0x1000000 | - |
| `SPI1_USER` | 0x3F402018 | 32 | CK_OUT_EDGE=0x200 CS_HOLD=0x40 CS_SETUP=0x80 DOUTDIN=0x1 FWRITE_DUAL=0x1000 FWRITE_OCT=0x4000 FWRITE_QUAD=0x2000 OPI_MODE=0x10 QPI_MODE=0x8 RD_BYTE_ORDER=0x400 RSCK_I_EDGE=0x100 SIO=0x10000 TSCK_I_EDGE=0x20 USR_ADDR=0x40000000 USR_ADDR_HOLD... | - |
| `SPI1_USER1` | 0x3F40201C | 32 | USR_ADDR_BITLEN=0xF8000000 USR_DUMMY_CYCLELEN=0xFF | - |
| `SPI1_USER2` | 0x3F402020 | 32 | USR_COMMAND_BITLEN=0xF0000000 USR_COMMAND_VALUE=0xFFFF | - |
| `SPI1_W0` | 0x3F402098 | 32 | BUF0=0xFFFFFFFF | - |
| `SPI1_W1` | 0x3F40209C | 32 | BUF1=0xFFFFFFFF | - |
| `SPI1_W10` | 0x3F4020C0 | 32 | BUF10=0xFFFFFFFF | - |
| `SPI1_W11` | 0x3F4020C4 | 32 | BUF11=0xFFFFFFFF | - |
| `SPI1_W12` | 0x3F4020C8 | 32 | BUF12=0xFFFFFFFF | - |
| `SPI1_W13` | 0x3F4020CC | 32 | BUF13=0xFFFFFFFF | - |
| `SPI1_W14` | 0x3F4020D0 | 32 | BUF14=0xFFFFFFFF | - |
| `SPI1_W15` | 0x3F4020D4 | 32 | BUF15=0xFFFFFFFF | - |
| `SPI1_W16` | 0x3F4020D8 | 32 | BUF16=0xFFFFFFFF | - |
| `SPI1_W17` | 0x3F4020DC | 32 | BUF17=0xFFFFFFFF | - |
| `SPI1_W2` | 0x3F4020A0 | 32 | BUF2=0xFFFFFFFF | - |
| `SPI1_W3` | 0x3F4020A4 | 32 | BUF3=0xFFFFFFFF | - |
| `SPI1_W4` | 0x3F4020A8 | 32 | BUF4=0xFFFFFFFF | - |
| `SPI1_W5` | 0x3F4020AC | 32 | BUF5=0xFFFFFFFF | - |
| `SPI1_W6` | 0x3F4020B0 | 32 | BUF6=0xFFFFFFFF | - |
| `SPI1_W7` | 0x3F4020B4 | 32 | BUF7=0xFFFFFFFF | - |
| `SPI1_W8` | 0x3F4020B8 | 32 | BUF8=0xFFFFFFFF | - |
| `SPI1_W9` | 0x3F4020BC | 32 | BUF9=0xFFFFFFFF | - |
| `SPI2_ADDR` | 0x3F424004 | 32 | USR_ADDR_VALUE=0xFFFFFFFF | - |
| `SPI2_CLOCK` | 0x3F424014 | 32 | CLKCNT_H=0xFC0 CLKCNT_L=0x3F CLKCNT_N=0x3F000 CLKDIV_PRE=0x7FFC0000 CLK_EQU_SYSCLK=0x80000000 | - |
| `SPI2_CMD` | 0x3F424000 | 32 | CONF_BITLEN=0x7FFFFF USR=0x1000000 | - |
| `SPI2_CTRL` | 0x3F424008 | 32 | DUMMY_OUT=0x8 D_POL=0x80000 EXT_HOLD_EN=0x4 FADDR_DUAL=0x20 FADDR_OCT=0x80 FADDR_QUAD=0x40 FCMD_DUAL=0x100 FCMD_OCT=0x400 FCMD_QUAD=0x200 FREAD_DUAL=0x4000 FREAD_OCT=0x10000 FREAD_QUAD=0x8000 Q_POL=0x40000 RD_BIT_ORDER=0x2000000 WP=0x200000... | - |
| `SPI2_CTRL1` | 0x3F42400C | 32 | CLK_MODE=0x3 CLK_MODE_13=0x4 CS_HOLD_DELAY=0xFC000 RSCK_DATA_OUT=0x8 W16_17_WR_ENA=0x10 | - |
| `SPI2_CTRL2` | 0x3F424010 | 32 | CS_DELAY_MODE=0x1C000000 CS_DELAY_NUM=0x60000000 CS_HOLD_TIME=0x3FFE000 CS_SETUP_TIME=0x1FFF | - |
| `SPI2_DIN_MODE` | 0x3F4240E0 | 32 | DIN0_MODE=0x7 DIN1_MODE=0x38 DIN2_MODE=0x1C0 DIN3_MODE=0xE00 DIN4_MODE=0x7000 DIN5_MODE=0x38000 DIN6_MODE=0x1C0000 DIN7_MODE=0xE00000 TIMING_CLK_ENA=0x1000000 | - |
| `SPI2_DIN_NUM` | 0x3F4240E4 | 32 | DIN0_NUM=0x3 DIN1_NUM=0xC DIN2_NUM=0x30 DIN3_NUM=0xC0 DIN4_NUM=0x300 DIN5_NUM=0xC00 DIN6_NUM=0x3000 DIN7_NUM=0xC000 | - |
| `SPI2_DMA_CONF` | 0x3F42404C | 32 | AHBM_FIFO_RST=0x10 AHBM_RST=0x20 DMA_CONTINUE=0x10000 DMA_INFIFO_FULL_CLR=0x400000 DMA_OUTFIFO_EMPTY_CLR=0x800000 DMA_RX_STOP=0x4000 DMA_SEG_TRANS_CLR=0x10000000 DMA_SLV_SEG_TRANS_EN=0x40000 DMA_TX_STOP=0x8000 EXT_MEM_BK_SIZE=0xC000000 INDS... | - |
| `SPI2_DMA_INSTATUS` | 0x3F424094 | 32 | DMA_INDSCR_ADDR=0x3FFFF DMA_INDSCR_STATE=0xC0000 DMA_INFIFO_CNT=0x3F800000 DMA_INFIFO_EMPTY=0x80000000 DMA_INFIFO_FULL=0x40000000 DMA_IN_STATE=0x700000 | - |
| `SPI2_DMA_INT_CLR` | 0x3F424064 | 32 | INFIFO_FULL_ERR_INT_CLR=0x200 INLINK_DSCR_EMPTY_INT_CLR=0x1 INLINK_DSCR_ERROR_INT_CLR=0x4 IN_DONE_INT_CLR=0x8 IN_ERR_EOF_INT_CLR=0x10 IN_SUC_EOF_INT_CLR=0x20 OUTFIFO_EMPTY_ERR_INT_CLR=0x400 OUTLINK_DSCR_ERROR_INT_CLR=0x2 OUT_DONE_INT_CLR=0x... | - |
| `SPI2_DMA_INT_ENA` | 0x3F424058 | 32 | INFIFO_FULL_ERR_INT_ENA=0x200 INLINK_DSCR_EMPTY_INT_ENA=0x1 INLINK_DSCR_ERROR_INT_ENA=0x4 IN_DONE_INT_ENA=0x8 IN_ERR_EOF_INT_ENA=0x10 IN_SUC_EOF_INT_ENA=0x20 OUTFIFO_EMPTY_ERR_INT_ENA=0x400 OUTLINK_DSCR_ERROR_INT_ENA=0x2 OUT_DONE_INT_ENA=0x... | - |
| `SPI2_DMA_INT_RAW` | 0x3F42405C | 32 | INFIFO_FULL_ERR_INT_RAW=0x200 INLINK_DSCR_EMPTY_INT_RAW=0x1 INLINK_DSCR_ERROR_INT_RAW=0x4 IN_DONE_INT_RAW=0x8 IN_ERR_EOF_INT_RAW=0x10 IN_SUC_EOF_INT_RAW=0x20 OUTFIFO_EMPTY_ERR_INT_RAW=0x400 OUTLINK_DSCR_ERROR_INT_RAW=0x2 OUT_DONE_INT_RAW=0x... | - |
| `SPI2_DMA_INT_ST` | 0x3F424060 | 32 | INFIFO_FULL_ERR_INT_ST=0x200 INLINK_DSCR_EMPTY_INT_ST=0x1 INLINK_DSCR_ERROR_INT_ST=0x4 IN_DONE_INT_ST=0x8 IN_ERR_EOF_INT_ST=0x10 IN_SUC_EOF_INT_ST=0x20 OUTFIFO_EMPTY_ERR_INT_ST=0x400 OUTLINK_DSCR_ERROR_INT_ST=0x2 OUT_DONE_INT_ST=0x40 OUT_EO... | - |
| `SPI2_DMA_IN_LINK` | 0x3F424054 | 32 | DMA_RX_ENA=0x80000000 INLINK_ADDR=0xFFFFF INLINK_AUTO_RET=0x100000 INLINK_RESTART=0x40000000 INLINK_START=0x20000000 INLINK_STOP=0x10000000 | - |
| `SPI2_DMA_OUTSTATUS` | 0x3F424090 | 32 | DMA_OUTDSCR_ADDR=0x3FFFF DMA_OUTDSCR_STATE=0xC0000 DMA_OUTFIFO_CNT=0x3F800000 DMA_OUTFIFO_EMPTY=0x80000000 DMA_OUTFIFO_FULL=0x40000000 DMA_OUT_STATE=0x700000 | - |
| `SPI2_DMA_OUT_LINK` | 0x3F424050 | 32 | DMA_TX_ENA=0x80000000 OUTLINK_ADDR=0xFFFFF OUTLINK_RESTART=0x40000000 OUTLINK_START=0x20000000 OUTLINK_STOP=0x10000000 | - |
| `SPI2_DOUT_MODE` | 0x3F4240E8 | 32 | DOUT0_MODE=0x7 DOUT1_MODE=0x38 DOUT2_MODE=0x1C0 DOUT3_MODE=0xE00 DOUT4_MODE=0x7000 DOUT5_MODE=0x38000 DOUT6_MODE=0x1C0000 DOUT7_MODE=0xE00000 | - |
| `SPI2_DOUT_NUM` | 0x3F4240EC | 32 | DOUT0_NUM=0x3 DOUT1_NUM=0xC DOUT2_NUM=0x30 DOUT3_NUM=0xC0 DOUT4_NUM=0x300 DOUT5_NUM=0xC00 DOUT6_NUM=0x3000 DOUT7_NUM=0xC000 | - |
| `SPI2_FSM` | 0x3F424044 | 32 | MST_DMA_RD_BYTELEN=0xFFFFF000 ST=0xF | - |
| `SPI2_HOLD` | 0x3F424048 | 32 | DMA_SEG_TRANS_DONE=0x80 INT_HOLD_ENA=0x3 OUT_EN=0x8 OUT_TIME=0x70 VAL=0x4 | - |
| `SPI2_INLINK_DSCR` | 0x3F424070 | 32 | DMA_INLINK_DSCR=0xFFFFFFFF | - |
| `SPI2_INLINK_DSCR_BF0` | 0x3F424074 | 32 | DMA_INLINK_DSCR_BF0=0xFFFFFFFF | - |
| `SPI2_INLINK_DSCR_BF1` | 0x3F424078 | 32 | DMA_INLINK_DSCR_BF1=0xFFFFFFFF | - |
| `SPI2_IN_ERR_EOF_DES_ADDR` | 0x3F424068 | 32 | DMA_IN_ERR_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI2_IN_SUC_EOF_DES_ADDR` | 0x3F42406C | 32 | DMA_IN_SUC_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI2_LCD_CTRL` | 0x3F4240F0 | 32 | LCD_HB_FRONT=0x7FF LCD_MODE_EN=0x80000000 LCD_VA_HEIGHT=0x1FF800 LCD_VT_HEIGHT=0x7FE00000 | - |
| `SPI2_LCD_CTRL1` | 0x3F4240F4 | 32 | LCD_HA_WIDTH=0xFFF00 LCD_HT_WIDTH=0xFFF00000 LCD_VB_FRONT=0xFF | - |
| `SPI2_LCD_CTRL2` | 0x3F4240F8 | 32 | HSYNC_IDLE_POL=0x800000 LCD_HSYNC_POSITION=0xFF000000 LCD_HSYNC_WIDTH=0x7F0000 LCD_VSYNC_WIDTH=0x7F VSYNC_IDLE_POL=0x80 | - |
| `SPI2_LCD_D_MODE` | 0x3F4240FC | 32 | DE_IDLE_POL=0x8000 D_CD_MODE=0x38 D_DE_MODE=0x1C0 D_DQS_MODE=0x7 D_HSYNC_MODE=0xE00 D_VSYNC_MODE=0x7000 HS_BLANK_EN=0x10000 | - |
| `SPI2_LCD_D_NUM` | 0x3F424100 | 32 | D_CD_NUM=0xC D_DE_NUM=0x30 D_DQS_NUM=0x3 D_HSYNC_NUM=0xC0 D_VSYNC_NUM=0x300 | - |
| `SPI2_MISC` | 0x3F42402C | 32 | ADDR_DTR_EN=0x40000 CD_ADDR_SET=0x400000 CD_CMD_SET=0x2000000 CD_DATA_SET=0x100000 CD_DUMMY_SET=0x200000 CD_IDLE_EDGE=0x4000000 CK_DIS=0x40 CK_IDLE_EDGE=0x20000000 CLK_DATA_DTR_EN=0x10000 CMD_DTR_EN=0x80000 CS0_DIS=0x1 CS1_DIS=0x2 CS2_DIS=0... | - |
| `SPI2_MISO_DLEN` | 0x3F424028 | 32 | USR_MISO_DBITLEN=0x7FFFFF | - |
| `SPI2_MOSI_DLEN` | 0x3F424024 | 32 | USR_MOSI_DBITLEN=0x7FFFFF | - |
| `SPI2_OUTLINK_DSCR` | 0x3F424084 | 32 | DMA_OUTLINK_DSCR=0xFFFFFFFF | - |
| `SPI2_OUTLINK_DSCR_BF0` | 0x3F424088 | 32 | DMA_OUTLINK_DSCR_BF0=0xFFFFFFFF | - |
| `SPI2_OUTLINK_DSCR_BF1` | 0x3F42408C | 32 | DMA_OUTLINK_DSCR_BF1=0xFFFFFFFF | - |
| `SPI2_OUT_EOF_BFR_DES_ADDR` | 0x3F42407C | 32 | DMA_OUT_EOF_BFR_DES_ADDR=0xFFFFFFFF | - |
| `SPI2_OUT_EOF_DES_ADDR` | 0x3F424080 | 32 | DMA_OUT_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI2_REG_DATE` | 0x3F4243FC | 32 | DATE=0xFFFFFFF | - |
| `SPI2_SLAVE` | 0x3F424030 | 32 | INT_DMA_SEG_TRANS_EN=0x400 INT_RD_BUF_DONE_EN=0x20 INT_RD_DMA_DONE_EN=0x80 INT_TRANS_DONE_EN=0x200 INT_WR_BUF_DONE_EN=0x40 INT_WR_DMA_DONE_EN=0x100 MODE=0x40000000 SEG_MAGIC_ERR_INT_EN=0x800 SOFT_RESET=0x80000000 TRANS_CNT=0x7800000 TRANS_D... | - |
| `SPI2_SLAVE1` | 0x3F424034 | 32 | SLV_ADDR_ERR=0x2000 SLV_ADDR_ERR_CLR=0x400 SLV_CMD_ERR=0x4000 SLV_CMD_ERR_CLR=0x800 SLV_LAST_ADDR=0xFF000000 SLV_LAST_COMMAND=0xFF0000 SLV_NO_QPI_EN=0x1000 SLV_WR_DMA_DONE=0x8000 | - |
| `SPI2_SLV_RDBUF_DLEN` | 0x3F42403C | 32 | SEG_MAGIC_ERR=0x2000000 SLV_DMA_RD_BYTELEN=0xFFFFF SLV_RD_BUF_DONE=0x1000000 | - |
| `SPI2_SLV_RD_BYTE` | 0x3F424040 | 32 | DMA_SEG_MAGIC_VALUE=0xF000000 SLV_DATA_BYTELEN=0xFFFFF SLV_RDBUF_BYTELEN_EN=0x400000 SLV_RDDMA_BYTELEN_EN=0x100000 SLV_RD_DMA_DONE=0x40000000 SLV_WRBUF_BYTELEN_EN=0x800000 SLV_WRDMA_BYTELEN_EN=0x200000 USR_CONF=0x80000000 | - |
| `SPI2_SLV_WRBUF_DLEN` | 0x3F424038 | 32 | CONF_BASE_BITLEN=0xFE000000 SLV_WR_BUF_DONE=0x1000000 | - |
| `SPI2_USER` | 0x3F424018 | 32 | CK_OUT_EDGE=0x200 CS_HOLD=0x40 CS_SETUP=0x80 DOUTDIN=0x1 FWRITE_DUAL=0x1000 FWRITE_OCT=0x4000 FWRITE_QUAD=0x2000 OPI_MODE=0x10 QPI_MODE=0x8 RD_BYTE_ORDER=0x400 RSCK_I_EDGE=0x100 SIO=0x10000 TSCK_I_EDGE=0x20 USR_ADDR=0x40000000 USR_ADDR_HOLD... | - |
| `SPI2_USER1` | 0x3F42401C | 32 | USR_ADDR_BITLEN=0xF8000000 USR_DUMMY_CYCLELEN=0xFF | - |
| `SPI2_USER2` | 0x3F424020 | 32 | USR_COMMAND_BITLEN=0xF0000000 USR_COMMAND_VALUE=0xFFFF | - |
| `SPI2_W0` | 0x3F424098 | 32 | BUF0=0xFFFFFFFF | - |
| `SPI2_W1` | 0x3F42409C | 32 | BUF1=0xFFFFFFFF | - |
| `SPI2_W10` | 0x3F4240C0 | 32 | BUF10=0xFFFFFFFF | - |
| `SPI2_W11` | 0x3F4240C4 | 32 | BUF11=0xFFFFFFFF | - |
| `SPI2_W12` | 0x3F4240C8 | 32 | BUF12=0xFFFFFFFF | - |
| `SPI2_W13` | 0x3F4240CC | 32 | BUF13=0xFFFFFFFF | - |
| `SPI2_W14` | 0x3F4240D0 | 32 | BUF14=0xFFFFFFFF | - |
| `SPI2_W15` | 0x3F4240D4 | 32 | BUF15=0xFFFFFFFF | - |
| `SPI2_W16` | 0x3F4240D8 | 32 | BUF16=0xFFFFFFFF | - |
| `SPI2_W17` | 0x3F4240DC | 32 | BUF17=0xFFFFFFFF | - |
| `SPI2_W2` | 0x3F4240A0 | 32 | BUF2=0xFFFFFFFF | - |
| `SPI2_W3` | 0x3F4240A4 | 32 | BUF3=0xFFFFFFFF | - |
| `SPI2_W4` | 0x3F4240A8 | 32 | BUF4=0xFFFFFFFF | - |
| `SPI2_W5` | 0x3F4240AC | 32 | BUF5=0xFFFFFFFF | - |
| `SPI2_W6` | 0x3F4240B0 | 32 | BUF6=0xFFFFFFFF | - |
| `SPI2_W7` | 0x3F4240B4 | 32 | BUF7=0xFFFFFFFF | - |
| `SPI2_W8` | 0x3F4240B8 | 32 | BUF8=0xFFFFFFFF | - |
| `SPI2_W9` | 0x3F4240BC | 32 | BUF9=0xFFFFFFFF | - |
| `SPI3_ADDR` | 0x3F425004 | 32 | USR_ADDR_VALUE=0xFFFFFFFF | - |
| `SPI3_CLOCK` | 0x3F425014 | 32 | CLKCNT_H=0xFC0 CLKCNT_L=0x3F CLKCNT_N=0x3F000 CLKDIV_PRE=0x7FFC0000 CLK_EQU_SYSCLK=0x80000000 | - |
| `SPI3_CMD` | 0x3F425000 | 32 | CONF_BITLEN=0x7FFFFF USR=0x1000000 | - |
| `SPI3_CTRL` | 0x3F425008 | 32 | DUMMY_OUT=0x8 D_POL=0x80000 EXT_HOLD_EN=0x4 FADDR_DUAL=0x20 FADDR_OCT=0x80 FADDR_QUAD=0x40 FCMD_DUAL=0x100 FCMD_OCT=0x400 FCMD_QUAD=0x200 FREAD_DUAL=0x4000 FREAD_OCT=0x10000 FREAD_QUAD=0x8000 Q_POL=0x40000 RD_BIT_ORDER=0x2000000 WP=0x200000... | - |
| `SPI3_CTRL1` | 0x3F42500C | 32 | CLK_MODE=0x3 CLK_MODE_13=0x4 CS_HOLD_DELAY=0xFC000 RSCK_DATA_OUT=0x8 W16_17_WR_ENA=0x10 | - |
| `SPI3_CTRL2` | 0x3F425010 | 32 | CS_DELAY_MODE=0x1C000000 CS_DELAY_NUM=0x60000000 CS_HOLD_TIME=0x3FFE000 CS_SETUP_TIME=0x1FFF | - |
| `SPI3_DIN_MODE` | 0x3F4250E0 | 32 | DIN0_MODE=0x7 DIN1_MODE=0x38 DIN2_MODE=0x1C0 DIN3_MODE=0xE00 DIN4_MODE=0x7000 DIN5_MODE=0x38000 DIN6_MODE=0x1C0000 DIN7_MODE=0xE00000 TIMING_CLK_ENA=0x1000000 | - |
| `SPI3_DIN_NUM` | 0x3F4250E4 | 32 | DIN0_NUM=0x3 DIN1_NUM=0xC DIN2_NUM=0x30 DIN3_NUM=0xC0 DIN4_NUM=0x300 DIN5_NUM=0xC00 DIN6_NUM=0x3000 DIN7_NUM=0xC000 | - |
| `SPI3_DMA_CONF` | 0x3F42504C | 32 | AHBM_FIFO_RST=0x10 AHBM_RST=0x20 DMA_CONTINUE=0x10000 DMA_INFIFO_FULL_CLR=0x400000 DMA_OUTFIFO_EMPTY_CLR=0x800000 DMA_RX_STOP=0x4000 DMA_SEG_TRANS_CLR=0x10000000 DMA_SLV_SEG_TRANS_EN=0x40000 DMA_TX_STOP=0x8000 EXT_MEM_BK_SIZE=0xC000000 INDS... | - |
| `SPI3_DMA_INSTATUS` | 0x3F425094 | 32 | DMA_INDSCR_ADDR=0x3FFFF DMA_INDSCR_STATE=0xC0000 DMA_INFIFO_CNT=0x3F800000 DMA_INFIFO_EMPTY=0x80000000 DMA_INFIFO_FULL=0x40000000 DMA_IN_STATE=0x700000 | - |
| `SPI3_DMA_INT_CLR` | 0x3F425064 | 32 | INFIFO_FULL_ERR_INT_CLR=0x200 INLINK_DSCR_EMPTY_INT_CLR=0x1 INLINK_DSCR_ERROR_INT_CLR=0x4 IN_DONE_INT_CLR=0x8 IN_ERR_EOF_INT_CLR=0x10 IN_SUC_EOF_INT_CLR=0x20 OUTFIFO_EMPTY_ERR_INT_CLR=0x400 OUTLINK_DSCR_ERROR_INT_CLR=0x2 OUT_DONE_INT_CLR=0x... | - |
| `SPI3_DMA_INT_ENA` | 0x3F425058 | 32 | INFIFO_FULL_ERR_INT_ENA=0x200 INLINK_DSCR_EMPTY_INT_ENA=0x1 INLINK_DSCR_ERROR_INT_ENA=0x4 IN_DONE_INT_ENA=0x8 IN_ERR_EOF_INT_ENA=0x10 IN_SUC_EOF_INT_ENA=0x20 OUTFIFO_EMPTY_ERR_INT_ENA=0x400 OUTLINK_DSCR_ERROR_INT_ENA=0x2 OUT_DONE_INT_ENA=0x... | - |
| `SPI3_DMA_INT_RAW` | 0x3F42505C | 32 | INFIFO_FULL_ERR_INT_RAW=0x200 INLINK_DSCR_EMPTY_INT_RAW=0x1 INLINK_DSCR_ERROR_INT_RAW=0x4 IN_DONE_INT_RAW=0x8 IN_ERR_EOF_INT_RAW=0x10 IN_SUC_EOF_INT_RAW=0x20 OUTFIFO_EMPTY_ERR_INT_RAW=0x400 OUTLINK_DSCR_ERROR_INT_RAW=0x2 OUT_DONE_INT_RAW=0x... | - |
| `SPI3_DMA_INT_ST` | 0x3F425060 | 32 | INFIFO_FULL_ERR_INT_ST=0x200 INLINK_DSCR_EMPTY_INT_ST=0x1 INLINK_DSCR_ERROR_INT_ST=0x4 IN_DONE_INT_ST=0x8 IN_ERR_EOF_INT_ST=0x10 IN_SUC_EOF_INT_ST=0x20 OUTFIFO_EMPTY_ERR_INT_ST=0x400 OUTLINK_DSCR_ERROR_INT_ST=0x2 OUT_DONE_INT_ST=0x40 OUT_EO... | - |
| `SPI3_DMA_IN_LINK` | 0x3F425054 | 32 | DMA_RX_ENA=0x80000000 INLINK_ADDR=0xFFFFF INLINK_AUTO_RET=0x100000 INLINK_RESTART=0x40000000 INLINK_START=0x20000000 INLINK_STOP=0x10000000 | - |
| `SPI3_DMA_OUTSTATUS` | 0x3F425090 | 32 | DMA_OUTDSCR_ADDR=0x3FFFF DMA_OUTDSCR_STATE=0xC0000 DMA_OUTFIFO_CNT=0x3F800000 DMA_OUTFIFO_EMPTY=0x80000000 DMA_OUTFIFO_FULL=0x40000000 DMA_OUT_STATE=0x700000 | - |
| `SPI3_DMA_OUT_LINK` | 0x3F425050 | 32 | DMA_TX_ENA=0x80000000 OUTLINK_ADDR=0xFFFFF OUTLINK_RESTART=0x40000000 OUTLINK_START=0x20000000 OUTLINK_STOP=0x10000000 | - |
| `SPI3_DOUT_MODE` | 0x3F4250E8 | 32 | DOUT0_MODE=0x7 DOUT1_MODE=0x38 DOUT2_MODE=0x1C0 DOUT3_MODE=0xE00 DOUT4_MODE=0x7000 DOUT5_MODE=0x38000 DOUT6_MODE=0x1C0000 DOUT7_MODE=0xE00000 | - |
| `SPI3_DOUT_NUM` | 0x3F4250EC | 32 | DOUT0_NUM=0x3 DOUT1_NUM=0xC DOUT2_NUM=0x30 DOUT3_NUM=0xC0 DOUT4_NUM=0x300 DOUT5_NUM=0xC00 DOUT6_NUM=0x3000 DOUT7_NUM=0xC000 | - |
| `SPI3_FSM` | 0x3F425044 | 32 | MST_DMA_RD_BYTELEN=0xFFFFF000 ST=0xF | - |
| `SPI3_HOLD` | 0x3F425048 | 32 | DMA_SEG_TRANS_DONE=0x80 INT_HOLD_ENA=0x3 OUT_EN=0x8 OUT_TIME=0x70 VAL=0x4 | - |
| `SPI3_INLINK_DSCR` | 0x3F425070 | 32 | DMA_INLINK_DSCR=0xFFFFFFFF | - |
| `SPI3_INLINK_DSCR_BF0` | 0x3F425074 | 32 | DMA_INLINK_DSCR_BF0=0xFFFFFFFF | - |
| `SPI3_INLINK_DSCR_BF1` | 0x3F425078 | 32 | DMA_INLINK_DSCR_BF1=0xFFFFFFFF | - |
| `SPI3_IN_ERR_EOF_DES_ADDR` | 0x3F425068 | 32 | DMA_IN_ERR_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI3_IN_SUC_EOF_DES_ADDR` | 0x3F42506C | 32 | DMA_IN_SUC_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI3_LCD_CTRL` | 0x3F4250F0 | 32 | LCD_HB_FRONT=0x7FF LCD_MODE_EN=0x80000000 LCD_VA_HEIGHT=0x1FF800 LCD_VT_HEIGHT=0x7FE00000 | - |
| `SPI3_LCD_CTRL1` | 0x3F4250F4 | 32 | LCD_HA_WIDTH=0xFFF00 LCD_HT_WIDTH=0xFFF00000 LCD_VB_FRONT=0xFF | - |
| `SPI3_LCD_CTRL2` | 0x3F4250F8 | 32 | HSYNC_IDLE_POL=0x800000 LCD_HSYNC_POSITION=0xFF000000 LCD_HSYNC_WIDTH=0x7F0000 LCD_VSYNC_WIDTH=0x7F VSYNC_IDLE_POL=0x80 | - |
| `SPI3_LCD_D_MODE` | 0x3F4250FC | 32 | DE_IDLE_POL=0x8000 D_CD_MODE=0x38 D_DE_MODE=0x1C0 D_DQS_MODE=0x7 D_HSYNC_MODE=0xE00 D_VSYNC_MODE=0x7000 HS_BLANK_EN=0x10000 | - |
| `SPI3_LCD_D_NUM` | 0x3F425100 | 32 | D_CD_NUM=0xC D_DE_NUM=0x30 D_DQS_NUM=0x3 D_HSYNC_NUM=0xC0 D_VSYNC_NUM=0x300 | - |
| `SPI3_MISC` | 0x3F42502C | 32 | ADDR_DTR_EN=0x40000 CD_ADDR_SET=0x400000 CD_CMD_SET=0x2000000 CD_DATA_SET=0x100000 CD_DUMMY_SET=0x200000 CD_IDLE_EDGE=0x4000000 CK_DIS=0x40 CK_IDLE_EDGE=0x20000000 CLK_DATA_DTR_EN=0x10000 CMD_DTR_EN=0x80000 CS0_DIS=0x1 CS1_DIS=0x2 CS2_DIS=0... | - |
| `SPI3_MISO_DLEN` | 0x3F425028 | 32 | USR_MISO_DBITLEN=0x7FFFFF | - |
| `SPI3_MOSI_DLEN` | 0x3F425024 | 32 | USR_MOSI_DBITLEN=0x7FFFFF | - |
| `SPI3_OUTLINK_DSCR` | 0x3F425084 | 32 | DMA_OUTLINK_DSCR=0xFFFFFFFF | - |
| `SPI3_OUTLINK_DSCR_BF0` | 0x3F425088 | 32 | DMA_OUTLINK_DSCR_BF0=0xFFFFFFFF | - |
| `SPI3_OUTLINK_DSCR_BF1` | 0x3F42508C | 32 | DMA_OUTLINK_DSCR_BF1=0xFFFFFFFF | - |
| `SPI3_OUT_EOF_BFR_DES_ADDR` | 0x3F42507C | 32 | DMA_OUT_EOF_BFR_DES_ADDR=0xFFFFFFFF | - |
| `SPI3_OUT_EOF_DES_ADDR` | 0x3F425080 | 32 | DMA_OUT_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI3_REG_DATE` | 0x3F4253FC | 32 | DATE=0xFFFFFFF | - |
| `SPI3_SLAVE` | 0x3F425030 | 32 | INT_DMA_SEG_TRANS_EN=0x400 INT_RD_BUF_DONE_EN=0x20 INT_RD_DMA_DONE_EN=0x80 INT_TRANS_DONE_EN=0x200 INT_WR_BUF_DONE_EN=0x40 INT_WR_DMA_DONE_EN=0x100 MODE=0x40000000 SEG_MAGIC_ERR_INT_EN=0x800 SOFT_RESET=0x80000000 TRANS_CNT=0x7800000 TRANS_D... | - |
| `SPI3_SLAVE1` | 0x3F425034 | 32 | SLV_ADDR_ERR=0x2000 SLV_ADDR_ERR_CLR=0x400 SLV_CMD_ERR=0x4000 SLV_CMD_ERR_CLR=0x800 SLV_LAST_ADDR=0xFF000000 SLV_LAST_COMMAND=0xFF0000 SLV_NO_QPI_EN=0x1000 SLV_WR_DMA_DONE=0x8000 | - |
| `SPI3_SLV_RDBUF_DLEN` | 0x3F42503C | 32 | SEG_MAGIC_ERR=0x2000000 SLV_DMA_RD_BYTELEN=0xFFFFF SLV_RD_BUF_DONE=0x1000000 | - |
| `SPI3_SLV_RD_BYTE` | 0x3F425040 | 32 | DMA_SEG_MAGIC_VALUE=0xF000000 SLV_DATA_BYTELEN=0xFFFFF SLV_RDBUF_BYTELEN_EN=0x400000 SLV_RDDMA_BYTELEN_EN=0x100000 SLV_RD_DMA_DONE=0x40000000 SLV_WRBUF_BYTELEN_EN=0x800000 SLV_WRDMA_BYTELEN_EN=0x200000 USR_CONF=0x80000000 | - |
| `SPI3_SLV_WRBUF_DLEN` | 0x3F425038 | 32 | CONF_BASE_BITLEN=0xFE000000 SLV_WR_BUF_DONE=0x1000000 | - |
| `SPI3_USER` | 0x3F425018 | 32 | CK_OUT_EDGE=0x200 CS_HOLD=0x40 CS_SETUP=0x80 DOUTDIN=0x1 FWRITE_DUAL=0x1000 FWRITE_OCT=0x4000 FWRITE_QUAD=0x2000 OPI_MODE=0x10 QPI_MODE=0x8 RD_BYTE_ORDER=0x400 RSCK_I_EDGE=0x100 SIO=0x10000 TSCK_I_EDGE=0x20 USR_ADDR=0x40000000 USR_ADDR_HOLD... | - |
| `SPI3_USER1` | 0x3F42501C | 32 | USR_ADDR_BITLEN=0xF8000000 USR_DUMMY_CYCLELEN=0xFF | - |
| `SPI3_USER2` | 0x3F425020 | 32 | USR_COMMAND_BITLEN=0xF0000000 USR_COMMAND_VALUE=0xFFFF | - |
| `SPI3_W0` | 0x3F425098 | 32 | BUF0=0xFFFFFFFF | - |
| `SPI3_W1` | 0x3F42509C | 32 | BUF1=0xFFFFFFFF | - |
| `SPI3_W10` | 0x3F4250C0 | 32 | BUF10=0xFFFFFFFF | - |
| `SPI3_W11` | 0x3F4250C4 | 32 | BUF11=0xFFFFFFFF | - |
| `SPI3_W12` | 0x3F4250C8 | 32 | BUF12=0xFFFFFFFF | - |
| `SPI3_W13` | 0x3F4250CC | 32 | BUF13=0xFFFFFFFF | - |
| `SPI3_W14` | 0x3F4250D0 | 32 | BUF14=0xFFFFFFFF | - |
| `SPI3_W15` | 0x3F4250D4 | 32 | BUF15=0xFFFFFFFF | - |
| `SPI3_W16` | 0x3F4250D8 | 32 | BUF16=0xFFFFFFFF | - |
| `SPI3_W17` | 0x3F4250DC | 32 | BUF17=0xFFFFFFFF | - |
| `SPI3_W2` | 0x3F4250A0 | 32 | BUF2=0xFFFFFFFF | - |
| `SPI3_W3` | 0x3F4250A4 | 32 | BUF3=0xFFFFFFFF | - |
| `SPI3_W4` | 0x3F4250A8 | 32 | BUF4=0xFFFFFFFF | - |
| `SPI3_W5` | 0x3F4250AC | 32 | BUF5=0xFFFFFFFF | - |
| `SPI3_W6` | 0x3F4250B0 | 32 | BUF6=0xFFFFFFFF | - |
| `SPI3_W7` | 0x3F4250B4 | 32 | BUF7=0xFFFFFFFF | - |
| `SPI3_W8` | 0x3F4250B8 | 32 | BUF8=0xFFFFFFFF | - |
| `SPI3_W9` | 0x3F4250BC | 32 | BUF9=0xFFFFFFFF | - |
| `SPI4_ADDR` | 0x3F437004 | 32 | USR_ADDR_VALUE=0xFFFFFFFF | - |
| `SPI4_CLOCK` | 0x3F437014 | 32 | CLKCNT_H=0xFC0 CLKCNT_L=0x3F CLKCNT_N=0x3F000 CLKDIV_PRE=0x7FFC0000 CLK_EQU_SYSCLK=0x80000000 | - |
| `SPI4_CMD` | 0x3F437000 | 32 | CONF_BITLEN=0x7FFFFF USR=0x1000000 | - |
| `SPI4_CTRL` | 0x3F437008 | 32 | DUMMY_OUT=0x8 D_POL=0x80000 EXT_HOLD_EN=0x4 FADDR_DUAL=0x20 FADDR_OCT=0x80 FADDR_QUAD=0x40 FCMD_DUAL=0x100 FCMD_OCT=0x400 FCMD_QUAD=0x200 FREAD_DUAL=0x4000 FREAD_OCT=0x10000 FREAD_QUAD=0x8000 Q_POL=0x40000 RD_BIT_ORDER=0x2000000 WP=0x200000... | - |
| `SPI4_CTRL1` | 0x3F43700C | 32 | CLK_MODE=0x3 CLK_MODE_13=0x4 CS_HOLD_DELAY=0xFC000 RSCK_DATA_OUT=0x8 W16_17_WR_ENA=0x10 | - |
| `SPI4_CTRL2` | 0x3F437010 | 32 | CS_DELAY_MODE=0x1C000000 CS_DELAY_NUM=0x60000000 CS_HOLD_TIME=0x3FFE000 CS_SETUP_TIME=0x1FFF | - |
| `SPI4_DIN_MODE` | 0x3F4370E0 | 32 | DIN0_MODE=0x7 DIN1_MODE=0x38 DIN2_MODE=0x1C0 DIN3_MODE=0xE00 DIN4_MODE=0x7000 DIN5_MODE=0x38000 DIN6_MODE=0x1C0000 DIN7_MODE=0xE00000 TIMING_CLK_ENA=0x1000000 | - |
| `SPI4_DIN_NUM` | 0x3F4370E4 | 32 | DIN0_NUM=0x3 DIN1_NUM=0xC DIN2_NUM=0x30 DIN3_NUM=0xC0 DIN4_NUM=0x300 DIN5_NUM=0xC00 DIN6_NUM=0x3000 DIN7_NUM=0xC000 | - |
| `SPI4_DMA_CONF` | 0x3F43704C | 32 | AHBM_FIFO_RST=0x10 AHBM_RST=0x20 DMA_CONTINUE=0x10000 DMA_INFIFO_FULL_CLR=0x400000 DMA_OUTFIFO_EMPTY_CLR=0x800000 DMA_RX_STOP=0x4000 DMA_SEG_TRANS_CLR=0x10000000 DMA_SLV_SEG_TRANS_EN=0x40000 DMA_TX_STOP=0x8000 EXT_MEM_BK_SIZE=0xC000000 INDS... | - |
| `SPI4_DMA_INSTATUS` | 0x3F437094 | 32 | DMA_INDSCR_ADDR=0x3FFFF DMA_INDSCR_STATE=0xC0000 DMA_INFIFO_CNT=0x3F800000 DMA_INFIFO_EMPTY=0x80000000 DMA_INFIFO_FULL=0x40000000 DMA_IN_STATE=0x700000 | - |
| `SPI4_DMA_INT_CLR` | 0x3F437064 | 32 | INFIFO_FULL_ERR_INT_CLR=0x200 INLINK_DSCR_EMPTY_INT_CLR=0x1 INLINK_DSCR_ERROR_INT_CLR=0x4 IN_DONE_INT_CLR=0x8 IN_ERR_EOF_INT_CLR=0x10 IN_SUC_EOF_INT_CLR=0x20 OUTFIFO_EMPTY_ERR_INT_CLR=0x400 OUTLINK_DSCR_ERROR_INT_CLR=0x2 OUT_DONE_INT_CLR=0x... | - |
| `SPI4_DMA_INT_ENA` | 0x3F437058 | 32 | INFIFO_FULL_ERR_INT_ENA=0x200 INLINK_DSCR_EMPTY_INT_ENA=0x1 INLINK_DSCR_ERROR_INT_ENA=0x4 IN_DONE_INT_ENA=0x8 IN_ERR_EOF_INT_ENA=0x10 IN_SUC_EOF_INT_ENA=0x20 OUTFIFO_EMPTY_ERR_INT_ENA=0x400 OUTLINK_DSCR_ERROR_INT_ENA=0x2 OUT_DONE_INT_ENA=0x... | - |
| `SPI4_DMA_INT_RAW` | 0x3F43705C | 32 | INFIFO_FULL_ERR_INT_RAW=0x200 INLINK_DSCR_EMPTY_INT_RAW=0x1 INLINK_DSCR_ERROR_INT_RAW=0x4 IN_DONE_INT_RAW=0x8 IN_ERR_EOF_INT_RAW=0x10 IN_SUC_EOF_INT_RAW=0x20 OUTFIFO_EMPTY_ERR_INT_RAW=0x400 OUTLINK_DSCR_ERROR_INT_RAW=0x2 OUT_DONE_INT_RAW=0x... | - |
| `SPI4_DMA_INT_ST` | 0x3F437060 | 32 | INFIFO_FULL_ERR_INT_ST=0x200 INLINK_DSCR_EMPTY_INT_ST=0x1 INLINK_DSCR_ERROR_INT_ST=0x4 IN_DONE_INT_ST=0x8 IN_ERR_EOF_INT_ST=0x10 IN_SUC_EOF_INT_ST=0x20 OUTFIFO_EMPTY_ERR_INT_ST=0x400 OUTLINK_DSCR_ERROR_INT_ST=0x2 OUT_DONE_INT_ST=0x40 OUT_EO... | - |
| `SPI4_DMA_IN_LINK` | 0x3F437054 | 32 | DMA_RX_ENA=0x80000000 INLINK_ADDR=0xFFFFF INLINK_AUTO_RET=0x100000 INLINK_RESTART=0x40000000 INLINK_START=0x20000000 INLINK_STOP=0x10000000 | - |
| `SPI4_DMA_OUTSTATUS` | 0x3F437090 | 32 | DMA_OUTDSCR_ADDR=0x3FFFF DMA_OUTDSCR_STATE=0xC0000 DMA_OUTFIFO_CNT=0x3F800000 DMA_OUTFIFO_EMPTY=0x80000000 DMA_OUTFIFO_FULL=0x40000000 DMA_OUT_STATE=0x700000 | - |
| `SPI4_DMA_OUT_LINK` | 0x3F437050 | 32 | DMA_TX_ENA=0x80000000 OUTLINK_ADDR=0xFFFFF OUTLINK_RESTART=0x40000000 OUTLINK_START=0x20000000 OUTLINK_STOP=0x10000000 | - |
| `SPI4_DOUT_MODE` | 0x3F4370E8 | 32 | DOUT0_MODE=0x7 DOUT1_MODE=0x38 DOUT2_MODE=0x1C0 DOUT3_MODE=0xE00 DOUT4_MODE=0x7000 DOUT5_MODE=0x38000 DOUT6_MODE=0x1C0000 DOUT7_MODE=0xE00000 | - |
| `SPI4_DOUT_NUM` | 0x3F4370EC | 32 | DOUT0_NUM=0x3 DOUT1_NUM=0xC DOUT2_NUM=0x30 DOUT3_NUM=0xC0 DOUT4_NUM=0x300 DOUT5_NUM=0xC00 DOUT6_NUM=0x3000 DOUT7_NUM=0xC000 | - |
| `SPI4_FSM` | 0x3F437044 | 32 | MST_DMA_RD_BYTELEN=0xFFFFF000 ST=0xF | - |
| `SPI4_HOLD` | 0x3F437048 | 32 | DMA_SEG_TRANS_DONE=0x80 INT_HOLD_ENA=0x3 OUT_EN=0x8 OUT_TIME=0x70 VAL=0x4 | - |
| `SPI4_INLINK_DSCR` | 0x3F437070 | 32 | DMA_INLINK_DSCR=0xFFFFFFFF | - |
| `SPI4_INLINK_DSCR_BF0` | 0x3F437074 | 32 | DMA_INLINK_DSCR_BF0=0xFFFFFFFF | - |
| `SPI4_INLINK_DSCR_BF1` | 0x3F437078 | 32 | DMA_INLINK_DSCR_BF1=0xFFFFFFFF | - |
| `SPI4_IN_ERR_EOF_DES_ADDR` | 0x3F437068 | 32 | DMA_IN_ERR_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI4_IN_SUC_EOF_DES_ADDR` | 0x3F43706C | 32 | DMA_IN_SUC_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI4_LCD_CTRL` | 0x3F4370F0 | 32 | LCD_HB_FRONT=0x7FF LCD_MODE_EN=0x80000000 LCD_VA_HEIGHT=0x1FF800 LCD_VT_HEIGHT=0x7FE00000 | - |
| `SPI4_LCD_CTRL1` | 0x3F4370F4 | 32 | LCD_HA_WIDTH=0xFFF00 LCD_HT_WIDTH=0xFFF00000 LCD_VB_FRONT=0xFF | - |
| `SPI4_LCD_CTRL2` | 0x3F4370F8 | 32 | HSYNC_IDLE_POL=0x800000 LCD_HSYNC_POSITION=0xFF000000 LCD_HSYNC_WIDTH=0x7F0000 LCD_VSYNC_WIDTH=0x7F VSYNC_IDLE_POL=0x80 | - |
| `SPI4_LCD_D_MODE` | 0x3F4370FC | 32 | DE_IDLE_POL=0x8000 D_CD_MODE=0x38 D_DE_MODE=0x1C0 D_DQS_MODE=0x7 D_HSYNC_MODE=0xE00 D_VSYNC_MODE=0x7000 HS_BLANK_EN=0x10000 | - |
| `SPI4_LCD_D_NUM` | 0x3F437100 | 32 | D_CD_NUM=0xC D_DE_NUM=0x30 D_DQS_NUM=0x3 D_HSYNC_NUM=0xC0 D_VSYNC_NUM=0x300 | - |
| `SPI4_MISC` | 0x3F43702C | 32 | ADDR_DTR_EN=0x40000 CD_ADDR_SET=0x400000 CD_CMD_SET=0x2000000 CD_DATA_SET=0x100000 CD_DUMMY_SET=0x200000 CD_IDLE_EDGE=0x4000000 CK_DIS=0x40 CK_IDLE_EDGE=0x20000000 CLK_DATA_DTR_EN=0x10000 CMD_DTR_EN=0x80000 CS0_DIS=0x1 CS1_DIS=0x2 CS2_DIS=0... | - |
| `SPI4_MISO_DLEN` | 0x3F437028 | 32 | USR_MISO_DBITLEN=0x7FFFFF | - |
| `SPI4_MOSI_DLEN` | 0x3F437024 | 32 | USR_MOSI_DBITLEN=0x7FFFFF | - |
| `SPI4_OUTLINK_DSCR` | 0x3F437084 | 32 | DMA_OUTLINK_DSCR=0xFFFFFFFF | - |
| `SPI4_OUTLINK_DSCR_BF0` | 0x3F437088 | 32 | DMA_OUTLINK_DSCR_BF0=0xFFFFFFFF | - |
| `SPI4_OUTLINK_DSCR_BF1` | 0x3F43708C | 32 | DMA_OUTLINK_DSCR_BF1=0xFFFFFFFF | - |
| `SPI4_OUT_EOF_BFR_DES_ADDR` | 0x3F43707C | 32 | DMA_OUT_EOF_BFR_DES_ADDR=0xFFFFFFFF | - |
| `SPI4_OUT_EOF_DES_ADDR` | 0x3F437080 | 32 | DMA_OUT_EOF_DES_ADDR=0xFFFFFFFF | - |
| `SPI4_REG_DATE` | 0x3F4373FC | 32 | DATE=0xFFFFFFF | - |
| `SPI4_SLAVE` | 0x3F437030 | 32 | INT_DMA_SEG_TRANS_EN=0x400 INT_RD_BUF_DONE_EN=0x20 INT_RD_DMA_DONE_EN=0x80 INT_TRANS_DONE_EN=0x200 INT_WR_BUF_DONE_EN=0x40 INT_WR_DMA_DONE_EN=0x100 MODE=0x40000000 SEG_MAGIC_ERR_INT_EN=0x800 SOFT_RESET=0x80000000 TRANS_CNT=0x7800000 TRANS_D... | - |
| `SPI4_SLAVE1` | 0x3F437034 | 32 | SLV_ADDR_ERR=0x2000 SLV_ADDR_ERR_CLR=0x400 SLV_CMD_ERR=0x4000 SLV_CMD_ERR_CLR=0x800 SLV_LAST_ADDR=0xFF000000 SLV_LAST_COMMAND=0xFF0000 SLV_NO_QPI_EN=0x1000 SLV_WR_DMA_DONE=0x8000 | - |
| `SPI4_SLV_RDBUF_DLEN` | 0x3F43703C | 32 | SEG_MAGIC_ERR=0x2000000 SLV_DMA_RD_BYTELEN=0xFFFFF SLV_RD_BUF_DONE=0x1000000 | - |
| `SPI4_SLV_RD_BYTE` | 0x3F437040 | 32 | DMA_SEG_MAGIC_VALUE=0xF000000 SLV_DATA_BYTELEN=0xFFFFF SLV_RDBUF_BYTELEN_EN=0x400000 SLV_RDDMA_BYTELEN_EN=0x100000 SLV_RD_DMA_DONE=0x40000000 SLV_WRBUF_BYTELEN_EN=0x800000 SLV_WRDMA_BYTELEN_EN=0x200000 USR_CONF=0x80000000 | - |
| `SPI4_SLV_WRBUF_DLEN` | 0x3F437038 | 32 | CONF_BASE_BITLEN=0xFE000000 SLV_WR_BUF_DONE=0x1000000 | - |
| `SPI4_USER` | 0x3F437018 | 32 | CK_OUT_EDGE=0x200 CS_HOLD=0x40 CS_SETUP=0x80 DOUTDIN=0x1 FWRITE_DUAL=0x1000 FWRITE_OCT=0x4000 FWRITE_QUAD=0x2000 OPI_MODE=0x10 QPI_MODE=0x8 RD_BYTE_ORDER=0x400 RSCK_I_EDGE=0x100 SIO=0x10000 TSCK_I_EDGE=0x20 USR_ADDR=0x40000000 USR_ADDR_HOLD... | - |
| `SPI4_USER1` | 0x3F43701C | 32 | USR_ADDR_BITLEN=0xF8000000 USR_DUMMY_CYCLELEN=0xFF | - |
| `SPI4_USER2` | 0x3F437020 | 32 | USR_COMMAND_BITLEN=0xF0000000 USR_COMMAND_VALUE=0xFFFF | - |
| `SPI4_W0` | 0x3F437098 | 32 | BUF0=0xFFFFFFFF | - |
| `SPI4_W1` | 0x3F43709C | 32 | BUF1=0xFFFFFFFF | - |
| `SPI4_W10` | 0x3F4370C0 | 32 | BUF10=0xFFFFFFFF | - |
| `SPI4_W11` | 0x3F4370C4 | 32 | BUF11=0xFFFFFFFF | - |
| `SPI4_W12` | 0x3F4370C8 | 32 | BUF12=0xFFFFFFFF | - |
| `SPI4_W13` | 0x3F4370CC | 32 | BUF13=0xFFFFFFFF | - |
| `SPI4_W14` | 0x3F4370D0 | 32 | BUF14=0xFFFFFFFF | - |
| `SPI4_W15` | 0x3F4370D4 | 32 | BUF15=0xFFFFFFFF | - |
| `SPI4_W16` | 0x3F4370D8 | 32 | BUF16=0xFFFFFFFF | - |
| `SPI4_W17` | 0x3F4370DC | 32 | BUF17=0xFFFFFFFF | - |
| `SPI4_W2` | 0x3F4370A0 | 32 | BUF2=0xFFFFFFFF | - |
| `SPI4_W3` | 0x3F4370A4 | 32 | BUF3=0xFFFFFFFF | - |
| `SPI4_W4` | 0x3F4370A8 | 32 | BUF4=0xFFFFFFFF | - |
| `SPI4_W5` | 0x3F4370AC | 32 | BUF5=0xFFFFFFFF | - |
| `SPI4_W6` | 0x3F4370B0 | 32 | BUF6=0xFFFFFFFF | - |
| `SPI4_W7` | 0x3F4370B4 | 32 | BUF7=0xFFFFFFFF | - |
| `SPI4_W8` | 0x3F4370B8 | 32 | BUF8=0xFFFFFFFF | - |
| `SPI4_W9` | 0x3F4370BC | 32 | BUF9=0xFFFFFFFF | - |
| `SYSTEM_BT_LPCK_DIV_FRAC` | 0x3F4C0054 | 32 | LPCLK_RTC_EN=0x10000000 LPCLK_SEL_8M=0x2000000 LPCLK_SEL_RTC_SLOW=0x1000000 LPCLK_SEL_XTAL32K=0x8000000 LPCLK_SEL_XTAL=0x4000000 | - |
| `SYSTEM_BUSTOEXTMEM_ENA` | 0x3F4C006C | 32 | BUSTOEXTMEM_ENA=0x1 | - |
| `SYSTEM_CACHE_CONTROL` | 0x3F4C0070 | 32 | PRO_CACHE_RESET=0x4 PRO_DCACHE_CLK_ON=0x2 PRO_ICACHE_CLK_ON=0x1 | - |
| `SYSTEM_CLOCK_GATE` | 0x3F4C0084 | 32 | CLK_EN=0x1 | - |
| `SYSTEM_CPU_INTR_FROM_CPU_0` | 0x3F4C0058 | 32 | CPU_INTR_FROM_CPU_0=0x1 | - |
| `SYSTEM_CPU_INTR_FROM_CPU_1` | 0x3F4C005C | 32 | CPU_INTR_FROM_CPU_1=0x1 | - |
| `SYSTEM_CPU_INTR_FROM_CPU_2` | 0x3F4C0060 | 32 | CPU_INTR_FROM_CPU_2=0x1 | - |
| `SYSTEM_CPU_INTR_FROM_CPU_3` | 0x3F4C0064 | 32 | CPU_INTR_FROM_CPU_3=0x1 | - |
| `SYSTEM_CPU_PERI_CLK_EN` | 0x3F4C0010 | 32 | CLK_EN_DEDICATED_GPIO=0x80 | - |
| `SYSTEM_CPU_PERI_RST_EN` | 0x3F4C0014 | 32 | RST_EN_DEDICATED_GPIO=0x80 | - |
| `SYSTEM_CPU_PER_CONF` | 0x3F4C0018 | 32 | CPUPERIOD_SEL=0x3 CPU_WAITI_DELAY_NUM=0xF0 CPU_WAIT_MODE_FORCE_ON=0x8 PLL_FREQ_SEL=0x4 | - |
| `SYSTEM_DATE` | 0x3F4C0FFC | 32 | DATE=0xFFFFFFF | - |
| `SYSTEM_EXTERNAL_DEVICE_ENCRYPT_DECRYPT_CONTROL` | 0x3F4C0074 | 32 | ENABLE_DOWNLOAD_DB_ENCRYPT=0x2 ENABLE_DOWNLOAD_G0CB_DECRYPT=0x4 ENABLE_DOWNLOAD_MANUAL_ENCRYPT=0x8 ENABLE_SPI_MANUAL_ENCRYPT=0x1 | - |
| `SYSTEM_JTAG_CTRL_0` | 0x3F4C001C | 32 | CANCEL_EFUSE_DISABLE_JTAG_TEMPORARY_0=0xFFFFFFFF | - |
| `SYSTEM_JTAG_CTRL_1` | 0x3F4C0020 | 32 | CANCEL_EFUSE_DISABLE_JTAG_TEMPORARY_1=0xFFFFFFFF | - |
| `SYSTEM_JTAG_CTRL_2` | 0x3F4C0024 | 32 | CANCEL_EFUSE_DISABLE_JTAG_TEMPORARY_2=0xFFFFFFFF | - |
| `SYSTEM_JTAG_CTRL_3` | 0x3F4C0028 | 32 | CANCEL_EFUSE_DISABLE_JTAG_TEMPORARY_3=0xFFFFFFFF | - |
| `SYSTEM_JTAG_CTRL_4` | 0x3F4C002C | 32 | CANCEL_EFUSE_DISABLE_JTAG_TEMPORARY_4=0xFFFFFFFF | - |
| `SYSTEM_JTAG_CTRL_5` | 0x3F4C0030 | 32 | CANCEL_EFUSE_DISABLE_JTAG_TEMPORARY_5=0xFFFFFFFF | - |
| `SYSTEM_JTAG_CTRL_6` | 0x3F4C0034 | 32 | CANCEL_EFUSE_DISABLE_JTAG_TEMPORARY_6=0xFFFFFFFF | - |
| `SYSTEM_JTAG_CTRL_7` | 0x3F4C0038 | 32 | CANCEL_EFUSE_DISABLE_JTAG_TEMPORARY_7=0xFFFFFFFF | - |
| `SYSTEM_LPCK_DIV_INT` | 0x3F4C0050 | 32 | LPCK_DIV_NUM=0xFFF | - |
| `SYSTEM_MEM_PD_MASK` | 0x3F4C003C | 32 | LSLP_MEM_PD_MASK=0x1 | - |
| `SYSTEM_PERIP_CLK_EN0` | 0x3F4C0040 | 32 | ADC2_ARB_CLK_EN=0x40000000 APB_SARADC_CLK_EN=0x10000000 CAN_CLK_EN=0x80000 EFUSE_CLK_EN=0x4000 I2C_EXT0_CLK_EN=0x80 I2C_EXT1_CLK_EN=0x40000 I2S0_CLK_EN=0x10 I2S1_CLK_EN=0x200000 LEDC_CLK_EN=0x800 PCNT_CLK_EN=0x400 PWM0_CLK_EN=0x20000 PWM1_C... | - |
| `SYSTEM_PERIP_CLK_EN1` | 0x3F4C0044 | 32 | CRYPTO_AES_CLK_EN=0x2 CRYPTO_DMA_CLK_EN=0x40 CRYPTO_DS_CLK_EN=0x10 CRYPTO_HMAC_CLK_EN=0x20 CRYPTO_RSA_CLK_EN=0x8 CRYPTO_SHA_CLK_EN=0x4 | - |
| `SYSTEM_PERIP_RST_EN0` | 0x3F4C0048 | 32 | ADC2_ARB_RST=0x40000000 APB_SARADC_RST=0x10000000 CAN_RST=0x80000 EFUSE_RST=0x4000 I2C_EXT0_RST=0x80 I2C_EXT1_RST=0x40000 I2S0_RST=0x10 I2S1_RST=0x200000 LEDC_RST=0x800 PCNT_RST=0x400 PWM0_RST=0x20000 PWM1_RST=0x100000 PWM2_RST=0x2000000 PW... | - |
| `SYSTEM_PERIP_RST_EN1` | 0x3F4C004C | 32 | CRYPTO_AES_RST=0x2 CRYPTO_DMA_RST=0x40 CRYPTO_DS_RST=0x10 CRYPTO_HMAC_RST=0x20 CRYPTO_RSA_RST=0x8 CRYPTO_SHA_RST=0x4 | - |
| `SYSTEM_ROM_CTRL_0` | 0x3F4C0000 | 32 | ROM_FO=0x3 | - |
| `SYSTEM_ROM_CTRL_1` | 0x3F4C0004 | 32 | ROM_FORCE_PD=0x3 ROM_FORCE_PU=0xC | - |
| `SYSTEM_RSA_PD_CTRL` | 0x3F4C0068 | 32 | RSA_MEM_FORCE_PD=0x4 RSA_MEM_FORCE_PU=0x2 RSA_MEM_PD=0x1 | - |
| `SYSTEM_RTC_FASTMEM_CONFIG` | 0x3F4C0078 | 32 | RTC_MEM_CRC_ADDR=0xFFE00 RTC_MEM_CRC_FINISH=0x80000000 RTC_MEM_CRC_LEN=0x7FF00000 RTC_MEM_CRC_START=0x100 | - |
| `SYSTEM_RTC_FASTMEM_CRC` | 0x3F4C007C | 32 | RTC_MEM_CRC_RES=0xFFFFFFFF | - |
| `SYSTEM_Redundant_ECO_Ctrl` | 0x3F4C0080 | 32 | REDUNDANT_ECO_DRIVE=0x1 REDUNDANT_ECO_RESULT=0x2 | - |
| `SYSTEM_SRAM_CTRL_0` | 0x3F4C0008 | 32 | SRAM_FO=0x3FFFFF | - |
| `SYSTEM_SRAM_CTRL_1` | 0x3F4C000C | 32 | SRAM_FORCE_PD=0x3FFFFF | - |
| `SYSTEM_SRAM_CTRL_2` | 0x3F4C0088 | 32 | SRAM_FORCE_PU=0x3FFFFF | - |
| `SYSTEM_SYSCLK_CONF` | 0x3F4C008C | 32 | CLK_DIV_EN=0x80000 CLK_XTAL_FREQ=0x7F000 PRE_DIV_CNT=0x3FF SOC_CLK_SEL=0xC00 | - |
| `SYSTIMER_CONF` | 0x3F423000 | 32 | CLK_EN=0x80000000 CLK_FO=0x1 | - |
| `SYSTIMER_DATE` | 0x3F4230FC | 32 | DATE=0xFFFFFFFF | - |
| `SYSTIMER_INT_CLR` | 0x3F42304C | 32 | TARGET0_INT_CLR=0x1 TARGET1_INT_CLR=0x2 TARGET2_INT_CLR=0x4 | - |
| `SYSTIMER_INT_ENA` | 0x3F423044 | 32 | TARGET0_INT_ENA=0x1 TARGET1_INT_ENA=0x2 TARGET2_INT_ENA=0x4 | - |
| `SYSTIMER_INT_RAW` | 0x3F423048 | 32 | INT0_RAW=0x1 INT1_RAW=0x2 INT2_RAW=0x4 | - |
| `SYSTIMER_LOAD` | 0x3F423004 | 32 | TIMER_LOAD=0x80000000 | - |
| `SYSTIMER_LOAD_HI` | 0x3F423008 | 32 | TIMER_LOAD_HI=0xFFFFFFFF | - |
| `SYSTIMER_LOAD_LO` | 0x3F42300C | 32 | TIMER_LOAD_LO=0xFFFFFFFF | - |
| `SYSTIMER_STEP` | 0x3F423010 | 32 | TIMER_PLL_STEP=0xFFC00 TIMER_XTAL_STEP=0x3FF | - |
| `SYSTIMER_TARGET0_CONF` | 0x3F42302C | 32 | TARGET0_PERIOD=0x3FFFFFFF TARGET0_PERIOD_MODE=0x40000000 TARGET0_WORK_EN=0x80000000 | - |
| `SYSTIMER_TARGET0_HI` | 0x3F423014 | 32 | TIMER_TARGET0_HI=0xFFFFFFFF | - |
| `SYSTIMER_TARGET0_LO` | 0x3F423018 | 32 | TIMER_TARGET0_LO=0xFFFFFFFF | - |
| `SYSTIMER_TARGET1_CONF` | 0x3F423030 | 32 | TARGET1_PERIOD=0x3FFFFFFF TARGET1_PERIOD_MODE=0x40000000 TARGET1_WORK_EN=0x80000000 | - |
| `SYSTIMER_TARGET1_HI` | 0x3F42301C | 32 | TIMER_TARGET1_HI=0xFFFFFFFF | - |
| `SYSTIMER_TARGET1_LO` | 0x3F423020 | 32 | TIMER_TARGET1_LO=0xFFFFFFFF | - |
| `SYSTIMER_TARGET2_CONF` | 0x3F423034 | 32 | TARGET2_PERIOD=0x3FFFFFFF TARGET2_PERIOD_MODE=0x40000000 TARGET2_WORK_EN=0x80000000 | - |
| `SYSTIMER_TARGET2_HI` | 0x3F423024 | 32 | TIMER_TARGET2_HI=0xFFFFFFFF | - |
| `SYSTIMER_TARGET2_LO` | 0x3F423028 | 32 | TIMER_TARGET2_LO=0xFFFFFFFF | - |
| `SYSTIMER_UNIT0_OP` | 0x3F423038 | 32 | TIMER_UNIT0_UPDATE=0x80000000 TIMER_UNIT0_VALUE_VALID=0x40000000 | - |
| `SYSTIMER_UNIT0_VALUE_HI` | 0x3F42303C | 32 | TIMER_VALUE_HI=0xFFFFFFFF | - |
| `SYSTIMER_UNIT0_VALUE_LO` | 0x3F423040 | 32 | TIMER_VALUE_LO=0xFFFFFFFF | - |
| `TIMG0_INT_CLR_TIMERS` | 0x3F41F0A4 | 32 | LACT_INT_CLR=0x8 T0_INT_CLR=0x1 T1_INT_CLR=0x2 WDT_INT_CLR=0x4 | - |
| `TIMG0_INT_ENA_TIMERS` | 0x3F41F098 | 32 | LACT_INT_ENA=0x8 T0_INT_ENA=0x1 T1_INT_ENA=0x2 WDT_INT_ENA=0x4 | - |
| `TIMG0_INT_RAW_TIMERS` | 0x3F41F09C | 32 | LACT_INT_RAW=0x8 T0_INT_RAW=0x1 T1_INT_RAW=0x2 WDT_INT_RAW=0x4 | - |
| `TIMG0_INT_ST_TIMERS` | 0x3F41F0A0 | 32 | LACT_INT_ST=0x8 T0_INT_ST=0x1 T1_INT_ST=0x2 WDT_INT_ST=0x4 | - |
| `TIMG0_LACTALARMHI` | 0x3F41F088 | 32 | LACT_ALARM_HI=0xFFFFFFFF | - |
| `TIMG0_LACTALARMLO` | 0x3F41F084 | 32 | LACT_ALARM_LO=0xFFFFFFFF | - |
| `TIMG0_LACTCONFIG` | 0x3F41F070 | 32 | LACT_ALARM_EN=0x400 LACT_AUTORELOAD=0x20000000 LACT_CPST_EN=0x100 LACT_DIVIDER=0x1FFFE000 LACT_EDGE_INT_EN=0x1000 LACT_EN=0x80000000 LACT_INCREASE=0x40000000 LACT_LAC_EN=0x200 LACT_LEVEL_INT_EN=0x800 LACT_RTC_ONLY=0x80 LACT_USE_REFTICK=0x40 | - |
| `TIMG0_LACTHI` | 0x3F41F07C | 32 | LACT_HI=0xFFFFFFFF | - |
| `TIMG0_LACTLO` | 0x3F41F078 | 32 | LACT_LO=0xFFFFFFFF | - |
| `TIMG0_LACTLOAD` | 0x3F41F094 | 32 | LACT_LOAD=0xFFFFFFFF | - |
| `TIMG0_LACTLOADHI` | 0x3F41F090 | 32 | LACT_LOAD_HI=0xFFFFFFFF | - |
| `TIMG0_LACTLOADLO` | 0x3F41F08C | 32 | LACT_LOAD_LO=0xFFFFFFFF | - |
| `TIMG0_LACTRTC` | 0x3F41F074 | 32 | LACT_RTC_STEP_LEN=0xFFFFFFC0 | - |
| `TIMG0_LACTUPDATE` | 0x3F41F080 | 32 | LACT_UPDATE=0xFFFFFFFF | - |
| `TIMG0_REGCLK` | 0x3F41F0FC | 32 | CLK_EN=0x80000000 | - |
| `TIMG0_RTCCALICFG` | 0x3F41F068 | 32 | RTC_CALI_CLK_SEL=0x6000 RTC_CALI_MAX=0x7FFF0000 RTC_CALI_RDY=0x8000 RTC_CALI_START=0x80000000 RTC_CALI_START_CYCLING=0x1000 | - |
| `TIMG0_RTCCALICFG1` | 0x3F41F06C | 32 | RTC_CALI_CYCLING_DATA_VLD=0x1 RTC_CALI_VALUE=0xFFFFFF80 | - |
| `TIMG0_RTCCALICFG2` | 0x3F41F0A8 | 32 | RTC_CALI_TIMEOUT=0x1 RTC_CALI_TIMEOUT_RST_CNT=0x78 RTC_CALI_TIMEOUT_THRES=0xFFFFFF80 | - |
| `TIMG0_T0ALARMHI` | 0x3F41F014 | 32 | ALARM_HI=0xFFFFFFFF | - |
| `TIMG0_T0ALARMLO` | 0x3F41F010 | 32 | ALARM_LO=0xFFFFFFFF | - |
| `TIMG0_T0CONFIG` | 0x3F41F000 | 32 | ALARM_EN=0x400 AUTORELOAD=0x20000000 DIVIDER=0x1FFFE000 EDGE_INT_EN=0x1000 EN=0x80000000 INCREASE=0x40000000 LEVEL_INT_EN=0x800 USE_XTAL=0x200 | - |
| `TIMG0_T0HI` | 0x3F41F008 | 32 | HI=0xFFFFFFFF | - |
| `TIMG0_T0LO` | 0x3F41F004 | 32 | LO=0xFFFFFFFF | - |
| `TIMG0_T0LOAD` | 0x3F41F020 | 32 | LOAD=0xFFFFFFFF | - |
| `TIMG0_T0LOADHI` | 0x3F41F01C | 32 | LOAD_HI=0xFFFFFFFF | - |
| `TIMG0_T0LOADLO` | 0x3F41F018 | 32 | LOAD_LO=0xFFFFFFFF | - |
| `TIMG0_T0UPDATE` | 0x3F41F00C | 32 | UPDATE=0x80000000 | - |
| `TIMG0_T1ALARMHI` | 0x3F41F038 | 32 | ALARM_HI=0xFFFFFFFF | - |
| `TIMG0_T1ALARMLO` | 0x3F41F034 | 32 | ALARM_LO=0xFFFFFFFF | - |
| `TIMG0_T1CONFIG` | 0x3F41F024 | 32 | ALARM_EN=0x400 AUTORELOAD=0x20000000 DIVIDER=0x1FFFE000 EDGE_INT_EN=0x1000 EN=0x80000000 INCREASE=0x40000000 LEVEL_INT_EN=0x800 USE_XTAL=0x200 | - |
| `TIMG0_T1HI` | 0x3F41F02C | 32 | HI=0xFFFFFFFF | - |
| `TIMG0_T1LO` | 0x3F41F028 | 32 | LO=0xFFFFFFFF | - |
| `TIMG0_T1LOAD` | 0x3F41F044 | 32 | LOAD=0xFFFFFFFF | - |
| `TIMG0_T1LOADHI` | 0x3F41F040 | 32 | LOAD_HI=0xFFFFFFFF | - |
| `TIMG0_T1LOADLO` | 0x3F41F03C | 32 | LOAD_LO=0xFFFFFFFF | - |
| `TIMG0_T1UPDATE` | 0x3F41F030 | 32 | UPDATE=0x80000000 | - |
| `TIMG0_TIMERS_DATE` | 0x3F41F0F8 | 32 | TIMERS_DATE=0xFFFFFFF | - |
| `TIMG0_WDTCONFIG0` | 0x3F41F048 | 32 | WDT_APPCPU_RESET_EN=0x1000 WDT_CPU_RESET_LENGTH=0x1C0000 WDT_EDGE_INT_EN=0x400000 WDT_EN=0x80000000 WDT_FLASHBOOT_MOD_EN=0x4000 WDT_LEVEL_INT_EN=0x200000 WDT_PROCPU_RESET_EN=0x2000 WDT_STG0=0x60000000 WDT_STG1=0x18000000 WDT_STG2=0x6000000 ... | - |
| `TIMG0_WDTCONFIG1` | 0x3F41F04C | 32 | WDT_CLK_PRESCALE=0xFFFF0000 | - |
| `TIMG0_WDTCONFIG2` | 0x3F41F050 | 32 | WDT_STG0_HOLD=0xFFFFFFFF | - |
| `TIMG0_WDTCONFIG3` | 0x3F41F054 | 32 | WDT_STG1_HOLD=0xFFFFFFFF | - |
| `TIMG0_WDTCONFIG4` | 0x3F41F058 | 32 | WDT_STG2_HOLD=0xFFFFFFFF | - |
| `TIMG0_WDTCONFIG5` | 0x3F41F05C | 32 | WDT_STG3_HOLD=0xFFFFFFFF | - |
| `TIMG0_WDTFEED` | 0x3F41F060 | 32 | WDT_FEED=0xFFFFFFFF | - |
| `TIMG0_WDTWPROTECT` | 0x3F41F064 | 32 | WDT_WKEY=0xFFFFFFFF | - |
| `TIMG1_INT_CLR_TIMERS` | 0x3F4200A4 | 32 | LACT_INT_CLR=0x8 T0_INT_CLR=0x1 T1_INT_CLR=0x2 WDT_INT_CLR=0x4 | - |
| `TIMG1_INT_ENA_TIMERS` | 0x3F420098 | 32 | LACT_INT_ENA=0x8 T0_INT_ENA=0x1 T1_INT_ENA=0x2 WDT_INT_ENA=0x4 | - |
| `TIMG1_INT_RAW_TIMERS` | 0x3F42009C | 32 | LACT_INT_RAW=0x8 T0_INT_RAW=0x1 T1_INT_RAW=0x2 WDT_INT_RAW=0x4 | - |
| `TIMG1_INT_ST_TIMERS` | 0x3F4200A0 | 32 | LACT_INT_ST=0x8 T0_INT_ST=0x1 T1_INT_ST=0x2 WDT_INT_ST=0x4 | - |
| `TIMG1_LACTALARMHI` | 0x3F420088 | 32 | LACT_ALARM_HI=0xFFFFFFFF | - |
| `TIMG1_LACTALARMLO` | 0x3F420084 | 32 | LACT_ALARM_LO=0xFFFFFFFF | - |
| `TIMG1_LACTCONFIG` | 0x3F420070 | 32 | LACT_ALARM_EN=0x400 LACT_AUTORELOAD=0x20000000 LACT_CPST_EN=0x100 LACT_DIVIDER=0x1FFFE000 LACT_EDGE_INT_EN=0x1000 LACT_EN=0x80000000 LACT_INCREASE=0x40000000 LACT_LAC_EN=0x200 LACT_LEVEL_INT_EN=0x800 LACT_RTC_ONLY=0x80 LACT_USE_REFTICK=0x40 | - |
| `TIMG1_LACTHI` | 0x3F42007C | 32 | LACT_HI=0xFFFFFFFF | - |
| `TIMG1_LACTLO` | 0x3F420078 | 32 | LACT_LO=0xFFFFFFFF | - |
| `TIMG1_LACTLOAD` | 0x3F420094 | 32 | LACT_LOAD=0xFFFFFFFF | - |
| `TIMG1_LACTLOADHI` | 0x3F420090 | 32 | LACT_LOAD_HI=0xFFFFFFFF | - |
| `TIMG1_LACTLOADLO` | 0x3F42008C | 32 | LACT_LOAD_LO=0xFFFFFFFF | - |
| `TIMG1_LACTRTC` | 0x3F420074 | 32 | LACT_RTC_STEP_LEN=0xFFFFFFC0 | - |
| `TIMG1_LACTUPDATE` | 0x3F420080 | 32 | LACT_UPDATE=0xFFFFFFFF | - |
| `TIMG1_REGCLK` | 0x3F4200FC | 32 | CLK_EN=0x80000000 | - |
| `TIMG1_RTCCALICFG` | 0x3F420068 | 32 | RTC_CALI_CLK_SEL=0x6000 RTC_CALI_MAX=0x7FFF0000 RTC_CALI_RDY=0x8000 RTC_CALI_START=0x80000000 RTC_CALI_START_CYCLING=0x1000 | - |
| `TIMG1_RTCCALICFG1` | 0x3F42006C | 32 | RTC_CALI_CYCLING_DATA_VLD=0x1 RTC_CALI_VALUE=0xFFFFFF80 | - |
| `TIMG1_RTCCALICFG2` | 0x3F4200A8 | 32 | RTC_CALI_TIMEOUT=0x1 RTC_CALI_TIMEOUT_RST_CNT=0x78 RTC_CALI_TIMEOUT_THRES=0xFFFFFF80 | - |
| `TIMG1_T0ALARMHI` | 0x3F420014 | 32 | ALARM_HI=0xFFFFFFFF | - |
| `TIMG1_T0ALARMLO` | 0x3F420010 | 32 | ALARM_LO=0xFFFFFFFF | - |
| `TIMG1_T0CONFIG` | 0x3F420000 | 32 | ALARM_EN=0x400 AUTORELOAD=0x20000000 DIVIDER=0x1FFFE000 EDGE_INT_EN=0x1000 EN=0x80000000 INCREASE=0x40000000 LEVEL_INT_EN=0x800 USE_XTAL=0x200 | - |
| `TIMG1_T0HI` | 0x3F420008 | 32 | HI=0xFFFFFFFF | - |
| `TIMG1_T0LO` | 0x3F420004 | 32 | LO=0xFFFFFFFF | - |
| `TIMG1_T0LOAD` | 0x3F420020 | 32 | LOAD=0xFFFFFFFF | - |
| `TIMG1_T0LOADHI` | 0x3F42001C | 32 | LOAD_HI=0xFFFFFFFF | - |
| `TIMG1_T0LOADLO` | 0x3F420018 | 32 | LOAD_LO=0xFFFFFFFF | - |
| `TIMG1_T0UPDATE` | 0x3F42000C | 32 | UPDATE=0x80000000 | - |
| `TIMG1_T1ALARMHI` | 0x3F420038 | 32 | ALARM_HI=0xFFFFFFFF | - |
| `TIMG1_T1ALARMLO` | 0x3F420034 | 32 | ALARM_LO=0xFFFFFFFF | - |
| `TIMG1_T1CONFIG` | 0x3F420024 | 32 | ALARM_EN=0x400 AUTORELOAD=0x20000000 DIVIDER=0x1FFFE000 EDGE_INT_EN=0x1000 EN=0x80000000 INCREASE=0x40000000 LEVEL_INT_EN=0x800 USE_XTAL=0x200 | - |
| `TIMG1_T1HI` | 0x3F42002C | 32 | HI=0xFFFFFFFF | - |
| `TIMG1_T1LO` | 0x3F420028 | 32 | LO=0xFFFFFFFF | - |
| `TIMG1_T1LOAD` | 0x3F420044 | 32 | LOAD=0xFFFFFFFF | - |
| `TIMG1_T1LOADHI` | 0x3F420040 | 32 | LOAD_HI=0xFFFFFFFF | - |
| `TIMG1_T1LOADLO` | 0x3F42003C | 32 | LOAD_LO=0xFFFFFFFF | - |
| `TIMG1_T1UPDATE` | 0x3F420030 | 32 | UPDATE=0x80000000 | - |
| `TIMG1_TIMERS_DATE` | 0x3F4200F8 | 32 | TIMERS_DATE=0xFFFFFFF | - |
| `TIMG1_WDTCONFIG0` | 0x3F420048 | 32 | WDT_APPCPU_RESET_EN=0x1000 WDT_CPU_RESET_LENGTH=0x1C0000 WDT_EDGE_INT_EN=0x400000 WDT_EN=0x80000000 WDT_FLASHBOOT_MOD_EN=0x4000 WDT_LEVEL_INT_EN=0x200000 WDT_PROCPU_RESET_EN=0x2000 WDT_STG0=0x60000000 WDT_STG1=0x18000000 WDT_STG2=0x6000000 ... | - |
| `TIMG1_WDTCONFIG1` | 0x3F42004C | 32 | WDT_CLK_PRESCALE=0xFFFF0000 | - |
| `TIMG1_WDTCONFIG2` | 0x3F420050 | 32 | WDT_STG0_HOLD=0xFFFFFFFF | - |
| `TIMG1_WDTCONFIG3` | 0x3F420054 | 32 | WDT_STG1_HOLD=0xFFFFFFFF | - |
| `TIMG1_WDTCONFIG4` | 0x3F420058 | 32 | WDT_STG2_HOLD=0xFFFFFFFF | - |
| `TIMG1_WDTCONFIG5` | 0x3F42005C | 32 | WDT_STG3_HOLD=0xFFFFFFFF | - |
| `TIMG1_WDTFEED` | 0x3F420060 | 32 | WDT_FEED=0xFFFFFFFF | - |
| `TIMG1_WDTWPROTECT` | 0x3F420064 | 32 | WDT_WKEY=0xFFFFFFFF | - |
| `TWAI_ARB_LOST_CAP` | 0x3F42B02C | 32 | ARB_LOST_CAP=0x1F | - |
| `TWAI_BUS_TIMING_0` | 0x3F42B018 | 32 | BAUD_PRESC=0x3FFF SYNC_JUMP_WIDTH=0xC000 | - |
| `TWAI_BUS_TIMING_1` | 0x3F42B01C | 32 | TIME_SAMP=0x80 TIME_SEG1=0xF TIME_SEG2=0x70 | - |
| `TWAI_CLOCK_DIVIDER` | 0x3F42B07C | 32 | CD=0xFF CLOCK_OFF=0x100 | - |
| `TWAI_CMD` | 0x3F42B004 | 32 | ABORT_TX=0x2 CLR_OVERRUN=0x8 RELEASE_BUF=0x4 SELF_RX_REQ=0x10 TX_REQ=0x1 | - |
| `TWAI_DATA_0` | 0x3F42B040 | 32 | TX_BYTE_0=0xFF | - |
| `TWAI_DATA_1` | 0x3F42B044 | 32 | TX_BYTE_1=0xFF | - |
| `TWAI_DATA_10` | 0x3F42B068 | 32 | TX_BYTE_10=0xFF | - |
| `TWAI_DATA_11` | 0x3F42B06C | 32 | TX_BYTE_11=0xFF | - |
| `TWAI_DATA_12` | 0x3F42B070 | 32 | TX_BYTE_12=0xFF | - |
| `TWAI_DATA_2` | 0x3F42B048 | 32 | TX_BYTE_2=0xFF | - |
| `TWAI_DATA_3` | 0x3F42B04C | 32 | TX_BYTE_3=0xFF | - |
| `TWAI_DATA_4` | 0x3F42B050 | 32 | TX_BYTE_4=0xFF | - |
| `TWAI_DATA_5` | 0x3F42B054 | 32 | TX_BYTE_5=0xFF | - |
| `TWAI_DATA_6` | 0x3F42B058 | 32 | TX_BYTE_6=0xFF | - |
| `TWAI_DATA_7` | 0x3F42B05C | 32 | TX_BYTE_7=0xFF | - |
| `TWAI_DATA_8` | 0x3F42B060 | 32 | TX_BYTE_8=0xFF | - |
| `TWAI_DATA_9` | 0x3F42B064 | 32 | TX_BYTE_9=0xFF | - |
| `TWAI_ERR_CODE_CAP` | 0x3F42B030 | 32 | ECC_DIRECTION=0x20 ECC_SEGMENT=0x1F ECC_TYPE=0xC0 | - |
| `TWAI_ERR_WARNING_LIMIT` | 0x3F42B034 | 32 | ERR_WARNING_LIMIT=0xFF | - |
| `TWAI_INT_ENA` | 0x3F42B010 | 32 | ARB_LOST_INT_ENA=0x40 BUS_ERR_INT_ENA=0x80 ERR_PASSIVE_INT_ENA=0x20 ERR_WARN_INT_ENA=0x4 OVERRUN_INT_ENA=0x8 RX_INT_ENA=0x1 TX_INT_ENA=0x2 | - |
| `TWAI_INT_RAW` | 0x3F42B00C | 32 | ARB_LOST_INT_ST=0x40 BUS_ERR_INT_ST=0x80 ERR_PASSIVE_INT_ST=0x20 ERR_WARN_INT_ST=0x4 OVERRUN_INT_ST=0x8 RX_INT_ST=0x1 TX_INT_ST=0x2 | - |
| `TWAI_MODE` | 0x3F42B000 | 32 | LISTEN_ONLY_MODE=0x2 RESET_MODE=0x1 RX_FILTER_MODE=0x8 SELF_TEST_MODE=0x4 | - |
| `TWAI_RX_ERR_CNT` | 0x3F42B038 | 32 | RX_ERR_CNT=0xFF | - |
| `TWAI_RX_MESSAGE_CNT` | 0x3F42B074 | 32 | RX_MESSAGE_COUNTER=0x7F | - |
| `TWAI_STATUS` | 0x3F42B008 | 32 | BUS_OFF_ST=0x80 ERR_ST=0x40 MISS_ST=0x100 OVERRUN_ST=0x2 RX_BUF_ST=0x1 RX_ST=0x10 TX_BUF_ST=0x4 TX_COMPLETE=0x8 TX_ST=0x20 | - |
| `TWAI_TX_ERR_CNT` | 0x3F42B03C | 32 | TX_ERR_CNT=0xFF | - |
| `UART0_AT_CMD_CHAR` | 0x3F400058 | 32 | AT_CMD_CHAR=0xFF CHAR_NUM=0xFF00 | - |
| `UART0_AT_CMD_GAPTOUT` | 0x3F400054 | 32 | RX_GAP_TOUT=0xFFFF | - |
| `UART0_AT_CMD_POSTCNT` | 0x3F400050 | 32 | POST_IDLE_NUM=0xFFFF | - |
| `UART0_AT_CMD_PRECNT` | 0x3F40004C | 32 | PRE_IDLE_NUM=0xFFFF | - |
| `UART0_AUTOBAUD` | 0x3F400018 | 32 | EN=0x1 GLITCH_FILT=0xFF00 | - |
| `UART0_CLKDIV` | 0x3F400014 | 32 | CLKDIV=0xFFFFF FRAG=0xF00000 | - |
| `UART0_CONF0` | 0x3F400020 | 32 | BIT_NUM=0xC CLK_EN=0x2000000 CTS_INV=0x100000 DSR_INV=0x200000 DTR_INV=0x1000000 ERR_WR_MASK=0x4000000 IRDA_DPLX=0x200 IRDA_EN=0x10000 IRDA_RX_INV=0x2000 IRDA_TX_EN=0x400 IRDA_TX_INV=0x1000 IRDA_WCTL=0x800 LOOPBACK=0x4000 MEM_CLK_EN=0x10000... | - |
| `UART0_CONF1` | 0x3F400024 | 32 | RXFIFO_FULL_THRHD=0x1FF RX_FLOW_EN=0x40000000 RX_TOUT_EN=0x80000000 RX_TOUT_FLOW_DIS=0x20000000 TXFIFO_EMPTY_THRHD=0x3FE00 | - |
| `UART0_DATE` | 0x3F400074 | 32 | DATE=0xFFFFFFFF | - |
| `UART0_FIFO` | 0x3F400000 | 32 | RXFIFO_RD_BYTE=0xFF | - |
| `UART0_FLOW_CONF` | 0x3F400034 | 32 | FORCE_XOFF=0x8 FORCE_XON=0x4 SEND_XOFF=0x20 SEND_XON=0x10 SW_FLOW_CON_EN=0x1 XONOFF_DEL=0x2 | - |
| `UART0_FSM_STATUS` | 0x3F400068 | 32 | ST_URX_OUT=0xF ST_UTX_OUT=0xF0 | - |
| `UART0_HIGHPULSE` | 0x3F40002C | 32 | MIN_CNT=0xFFFFF | - |
| `UART0_ID` | 0x3F400078 | 32 | ID=0xFFFFFFFF | - |
| `UART0_IDLE_CONF` | 0x3F400044 | 32 | RX_IDLE_THRHD=0x3FF TX_BRK_NUM=0xFF00000 TX_IDLE_NUM=0xFFC00 | - |
| `UART0_INT_CLR` | 0x3F400010 | 32 | AT_CMD_CHAR_DET_INT_CLR=0x40000 BRK_DET_INT_CLR=0x80 CTS_CHG_INT_CLR=0x40 DSR_CHG_INT_CLR=0x20 FRM_ERR_INT_CLR=0x8 GLITCH_DET_INT_CLR=0x800 PARITY_ERR_INT_CLR=0x4 RS485_CLASH_INT_CLR=0x20000 RS485_FRM_ERR_INT_CLR=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART0_INT_ENA` | 0x3F40000C | 32 | AT_CMD_CHAR_DET_INT_ENA=0x40000 BRK_DET_INT_ENA=0x80 CTS_CHG_INT_ENA=0x40 DSR_CHG_INT_ENA=0x20 FRM_ERR_INT_ENA=0x8 GLITCH_DET_INT_ENA=0x800 PARITY_ERR_INT_ENA=0x4 RS485_CLASH_INT_ENA=0x20000 RS485_FRM_ERR_INT_ENA=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART0_INT_RAW` | 0x3F400004 | 32 | AT_CMD_CHAR_DET_INT_RAW=0x40000 BRK_DET_INT_RAW=0x80 CTS_CHG_INT_RAW=0x40 DSR_CHG_INT_RAW=0x20 FRM_ERR_INT_RAW=0x8 GLITCH_DET_INT_RAW=0x800 PARITY_ERR_INT_RAW=0x4 RS485_CLASH_INT_RAW=0x20000 RS485_FRM_ERR_INT_RAW=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART0_INT_ST` | 0x3F400008 | 32 | AT_CMD_CHAR_DET_INT_ST=0x40000 BRK_DET_INT_ST=0x80 CTS_CHG_INT_ST=0x40 DSR_CHG_INT_ST=0x20 FRM_ERR_INT_ST=0x8 GLITCH_DET_INT_ST=0x800 PARITY_ERR_INT_ST=0x4 RS485_CLASH_INT_ST=0x20000 RS485_FRM_ERR_INT_ST=0x10000 RS485_PARITY_ERR_INT_ST=0x80... | - |
| `UART0_LOWPULSE` | 0x3F400028 | 32 | MIN_CNT=0xFFFFF | - |
| `UART0_MEM_CONF` | 0x3F40005C | 32 | MEM_FORCE_PD=0x4000000 MEM_FORCE_PU=0x8000000 RX_FLOW_THRHD=0xFF80 RX_SIZE=0xE RX_TOUT_THRHD=0x3FF0000 TX_SIZE=0x70 | - |
| `UART0_MEM_RX_STATUS` | 0x3F400064 | 32 | APB_RX_RADDR=0x3FF RX_WADDR=0x1FF800 | - |
| `UART0_MEM_TX_STATUS` | 0x3F400060 | 32 | APB_TX_WADDR=0x3FF TX_RADDR=0x1FF800 | - |
| `UART0_NEGPULSE` | 0x3F400070 | 32 | NEGEDGE_MIN_CNT=0xFFFFF | - |
| `UART0_POSPULSE` | 0x3F40006C | 32 | POSEDGE_MIN_CNT=0xFFFFF | - |
| `UART0_RS485_CONF` | 0x3F400048 | 32 | DL0_EN=0x2 DL1_EN=0x4 RS485RXBY_TX_EN=0x10 RS485TX_RX_EN=0x8 RS485_EN=0x1 RS485_RX_DLY_NUM=0x20 RS485_TX_DLY_NUM=0x3C0 | - |
| `UART0_RXD_CNT` | 0x3F400030 | 32 | RXD_EDGE_CNT=0x3FF | - |
| `UART0_SLEEP_CONF` | 0x3F400038 | 32 | ACTIVE_THRESHOLD=0x3FF | - |
| `UART0_STATUS` | 0x3F40001C | 32 | CTSN=0x4000 DSRN=0x2000 DTRN=0x20000000 RTSN=0x40000000 RXD=0x8000 RXFIFO_CNT=0x3FF TXD=0x80000000 TXFIFO_CNT=0x3FF0000 | - |
| `UART0_SWFC_CONF0` | 0x3F40003C | 32 | XOFF_CHAR=0x1FE00 XOFF_THRESHOLD=0x1FF | - |
| `UART0_SWFC_CONF1` | 0x3F400040 | 32 | XON_CHAR=0x1FE00 XON_THRESHOLD=0x1FF | - |
| `UART1_AT_CMD_CHAR` | 0x3F410058 | 32 | AT_CMD_CHAR=0xFF CHAR_NUM=0xFF00 | - |
| `UART1_AT_CMD_GAPTOUT` | 0x3F410054 | 32 | RX_GAP_TOUT=0xFFFF | - |
| `UART1_AT_CMD_POSTCNT` | 0x3F410050 | 32 | POST_IDLE_NUM=0xFFFF | - |
| `UART1_AT_CMD_PRECNT` | 0x3F41004C | 32 | PRE_IDLE_NUM=0xFFFF | - |
| `UART1_AUTOBAUD` | 0x3F410018 | 32 | EN=0x1 GLITCH_FILT=0xFF00 | - |
| `UART1_CLKDIV` | 0x3F410014 | 32 | CLKDIV=0xFFFFF FRAG=0xF00000 | - |
| `UART1_CONF0` | 0x3F410020 | 32 | BIT_NUM=0xC CLK_EN=0x2000000 CTS_INV=0x100000 DSR_INV=0x200000 DTR_INV=0x1000000 ERR_WR_MASK=0x4000000 IRDA_DPLX=0x200 IRDA_EN=0x10000 IRDA_RX_INV=0x2000 IRDA_TX_EN=0x400 IRDA_TX_INV=0x1000 IRDA_WCTL=0x800 LOOPBACK=0x4000 MEM_CLK_EN=0x10000... | - |
| `UART1_CONF1` | 0x3F410024 | 32 | RXFIFO_FULL_THRHD=0x1FF RX_FLOW_EN=0x40000000 RX_TOUT_EN=0x80000000 RX_TOUT_FLOW_DIS=0x20000000 TXFIFO_EMPTY_THRHD=0x3FE00 | - |
| `UART1_DATE` | 0x3F410074 | 32 | DATE=0xFFFFFFFF | - |
| `UART1_FIFO` | 0x3F410000 | 32 | RXFIFO_RD_BYTE=0xFF | - |
| `UART1_FLOW_CONF` | 0x3F410034 | 32 | FORCE_XOFF=0x8 FORCE_XON=0x4 SEND_XOFF=0x20 SEND_XON=0x10 SW_FLOW_CON_EN=0x1 XONOFF_DEL=0x2 | - |
| `UART1_FSM_STATUS` | 0x3F410068 | 32 | ST_URX_OUT=0xF ST_UTX_OUT=0xF0 | - |
| `UART1_HIGHPULSE` | 0x3F41002C | 32 | MIN_CNT=0xFFFFF | - |
| `UART1_ID` | 0x3F410078 | 32 | ID=0xFFFFFFFF | - |
| `UART1_IDLE_CONF` | 0x3F410044 | 32 | RX_IDLE_THRHD=0x3FF TX_BRK_NUM=0xFF00000 TX_IDLE_NUM=0xFFC00 | - |
| `UART1_INT_CLR` | 0x3F410010 | 32 | AT_CMD_CHAR_DET_INT_CLR=0x40000 BRK_DET_INT_CLR=0x80 CTS_CHG_INT_CLR=0x40 DSR_CHG_INT_CLR=0x20 FRM_ERR_INT_CLR=0x8 GLITCH_DET_INT_CLR=0x800 PARITY_ERR_INT_CLR=0x4 RS485_CLASH_INT_CLR=0x20000 RS485_FRM_ERR_INT_CLR=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART1_INT_ENA` | 0x3F41000C | 32 | AT_CMD_CHAR_DET_INT_ENA=0x40000 BRK_DET_INT_ENA=0x80 CTS_CHG_INT_ENA=0x40 DSR_CHG_INT_ENA=0x20 FRM_ERR_INT_ENA=0x8 GLITCH_DET_INT_ENA=0x800 PARITY_ERR_INT_ENA=0x4 RS485_CLASH_INT_ENA=0x20000 RS485_FRM_ERR_INT_ENA=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART1_INT_RAW` | 0x3F410004 | 32 | AT_CMD_CHAR_DET_INT_RAW=0x40000 BRK_DET_INT_RAW=0x80 CTS_CHG_INT_RAW=0x40 DSR_CHG_INT_RAW=0x20 FRM_ERR_INT_RAW=0x8 GLITCH_DET_INT_RAW=0x800 PARITY_ERR_INT_RAW=0x4 RS485_CLASH_INT_RAW=0x20000 RS485_FRM_ERR_INT_RAW=0x10000 RS485_PARITY_ERR_IN... | - |
| `UART1_INT_ST` | 0x3F410008 | 32 | AT_CMD_CHAR_DET_INT_ST=0x40000 BRK_DET_INT_ST=0x80 CTS_CHG_INT_ST=0x40 DSR_CHG_INT_ST=0x20 FRM_ERR_INT_ST=0x8 GLITCH_DET_INT_ST=0x800 PARITY_ERR_INT_ST=0x4 RS485_CLASH_INT_ST=0x20000 RS485_FRM_ERR_INT_ST=0x10000 RS485_PARITY_ERR_INT_ST=0x80... | - |
| `UART1_LOWPULSE` | 0x3F410028 | 32 | MIN_CNT=0xFFFFF | - |
| `UART1_MEM_CONF` | 0x3F41005C | 32 | MEM_FORCE_PD=0x4000000 MEM_FORCE_PU=0x8000000 RX_FLOW_THRHD=0xFF80 RX_SIZE=0xE RX_TOUT_THRHD=0x3FF0000 TX_SIZE=0x70 | - |
| `UART1_MEM_RX_STATUS` | 0x3F410064 | 32 | APB_RX_RADDR=0x3FF RX_WADDR=0x1FF800 | - |
| `UART1_MEM_TX_STATUS` | 0x3F410060 | 32 | APB_TX_WADDR=0x3FF TX_RADDR=0x1FF800 | - |
| `UART1_NEGPULSE` | 0x3F410070 | 32 | NEGEDGE_MIN_CNT=0xFFFFF | - |
| `UART1_POSPULSE` | 0x3F41006C | 32 | POSEDGE_MIN_CNT=0xFFFFF | - |
| `UART1_RS485_CONF` | 0x3F410048 | 32 | DL0_EN=0x2 DL1_EN=0x4 RS485RXBY_TX_EN=0x10 RS485TX_RX_EN=0x8 RS485_EN=0x1 RS485_RX_DLY_NUM=0x20 RS485_TX_DLY_NUM=0x3C0 | - |
| `UART1_RXD_CNT` | 0x3F410030 | 32 | RXD_EDGE_CNT=0x3FF | - |
| `UART1_SLEEP_CONF` | 0x3F410038 | 32 | ACTIVE_THRESHOLD=0x3FF | - |
| `UART1_STATUS` | 0x3F41001C | 32 | CTSN=0x4000 DSRN=0x2000 DTRN=0x20000000 RTSN=0x40000000 RXD=0x8000 RXFIFO_CNT=0x3FF TXD=0x80000000 TXFIFO_CNT=0x3FF0000 | - |
| `UART1_SWFC_CONF0` | 0x3F41003C | 32 | XOFF_CHAR=0x1FE00 XOFF_THRESHOLD=0x1FF | - |
| `UART1_SWFC_CONF1` | 0x3F410040 | 32 | XON_CHAR=0x1FE00 XON_THRESHOLD=0x1FF | - |
| `UHCI0_AHB_TEST` | 0x3F414048 | 32 | AHB_TESTADDR=0x30 AHB_TESTMODE=0x7 | - |
| `UHCI0_CONF0` | 0x3F414000 | 32 | AHBM_FIFO_RST=0x4 AHBM_RST=0x8 CLK_EN=0x400000 CRC_REC_EN=0x40000 ENCODE_CRC_EN=0x200000 HEAD_EN=0x20000 INDSCR_BURST_EN=0x2000 IN_LOOP_TEST=0x10 IN_RST=0x1 LEN_EOF_EN=0x100000 MEM_TRANS_EN=0x8000 OUTDSCR_BURST_EN=0x1000 OUT_AUTO_WRBACK=0x4... | - |
| `UHCI0_CONF1` | 0x3F41402C | 32 | CHECK_OWNER=0x40 CHECK_SEQ_EN=0x2 CHECK_SUM_EN=0x1 CRC_DISABLE=0x4 DMA_INFIFO_FULL_THRS=0x1FFE00 SAVE_HEAD=0x8 SW_START=0x100 TX_ACK_NUM_RE=0x20 TX_CHECK_SUM_RE=0x10 WAIT_SW_START=0x80 | - |
| `UHCI0_DATE` | 0x3F4140FC | 32 | DATE=0xFFFFFFFF | - |
| `UHCI0_DMA_IN_DSCR` | 0x3F41404C | 32 | INLINK_DSCR=0xFFFFFFFF | - |
| `UHCI0_DMA_IN_DSCR_BF0` | 0x3F414050 | 32 | INLINK_DSCR_BF0=0xFFFFFFFF | - |
| `UHCI0_DMA_IN_ERR_EOF_DES_ADDR` | 0x3F414040 | 32 | IN_ERR_EOF_DES_ADDR=0xFFFFFFFF | - |
| `UHCI0_DMA_IN_LINK` | 0x3F414028 | 32 | INLINK_ADDR=0xFFFFF INLINK_AUTO_RET=0x100000 INLINK_PARK=0x80000000 INLINK_RESTART=0x40000000 INLINK_START=0x20000000 INLINK_STOP=0x10000000 | - |
| `UHCI0_DMA_IN_POP` | 0x3F414020 | 32 | INFIFO_POP=0x10000 INFIFO_RDATA=0xFFF | - |
| `UHCI0_DMA_IN_STATUS` | 0x3F41401C | 32 | IN_EMPTY=0x2 IN_FULL=0x1 RX_ERR_CAUSE=0x70 | - |
| `UHCI0_DMA_IN_SUC_EOF_DES_ADDR` | 0x3F41403C | 32 | IN_SUC_EOF_DES_ADDR=0xFFFFFFFF | - |
| `UHCI0_DMA_OUT_DSCR` | 0x3F414058 | 32 | OUTLINK_DSCR=0xFFFFFFFF | - |
| `UHCI0_DMA_OUT_DSCR_BF0` | 0x3F41405C | 32 | OUTLINK_DSCR_BF0=0xFFFFFFFF | - |
| `UHCI0_DMA_OUT_EOF_BFR_DES_ADDR` | 0x3F414044 | 32 | OUT_EOF_BFR_DES_ADDR=0xFFFFFFFF | - |
| `UHCI0_DMA_OUT_EOF_DES_ADDR` | 0x3F414038 | 32 | OUT_EOF_DES_ADDR=0xFFFFFFFF | - |
| `UHCI0_DMA_OUT_LINK` | 0x3F414024 | 32 | OUTLINK_ADDR=0xFFFFF OUTLINK_PARK=0x80000000 OUTLINK_RESTART=0x40000000 OUTLINK_START=0x20000000 OUTLINK_STOP=0x10000000 | - |
| `UHCI0_DMA_OUT_PUSH` | 0x3F414018 | 32 | OUTFIFO_PUSH=0x10000 OUTFIFO_WDATA=0x1FF | - |
| `UHCI0_DMA_OUT_STATUS` | 0x3F414014 | 32 | OUT_EMPTY=0x2 OUT_FULL=0x1 | - |
| `UHCI0_ESCAPE_CONF` | 0x3F414064 | 32 | RX_11_ESC_EN=0x40 RX_13_ESC_EN=0x80 RX_C0_ESC_EN=0x10 RX_DB_ESC_EN=0x20 TX_11_ESC_EN=0x4 TX_13_ESC_EN=0x8 TX_C0_ESC_EN=0x1 TX_DB_ESC_EN=0x2 | - |
| `UHCI0_ESC_CONF0` | 0x3F4140B0 | 32 | SEPER_CHAR=0xFF SEPER_ESC_CHAR0=0xFF00 SEPER_ESC_CHAR1=0xFF0000 | - |
| `UHCI0_ESC_CONF1` | 0x3F4140B4 | 32 | ESC_SEQ0=0xFF ESC_SEQ0_CHAR0=0xFF00 ESC_SEQ0_CHAR1=0xFF0000 | - |
| `UHCI0_ESC_CONF2` | 0x3F4140B8 | 32 | ESC_SEQ1=0xFF ESC_SEQ1_CHAR0=0xFF00 ESC_SEQ1_CHAR1=0xFF0000 | - |
| `UHCI0_ESC_CONF3` | 0x3F4140BC | 32 | ESC_SEQ2=0xFF ESC_SEQ2_CHAR0=0xFF00 ESC_SEQ2_CHAR1=0xFF0000 | - |
| `UHCI0_HUNG_CONF` | 0x3F414068 | 32 | RXFIFO_TIMEOUT=0xFF000 RXFIFO_TIMEOUT_ENA=0x800000 RXFIFO_TIMEOUT_SHIFT=0x700000 TXFIFO_TIMEOUT=0xFF TXFIFO_TIMEOUT_ENA=0x800 TXFIFO_TIMEOUT_SHIFT=0x700 | - |
| `UHCI0_INT_CLR` | 0x3F414010 | 32 | DMA_INFIFO_FULL_WM_INT_CLR=0x10000 IN_DONE_INT_CLR=0x10 IN_DSCR_EMPTY_INT_CLR=0x800 IN_DSCR_ERR_INT_CLR=0x200 IN_ERR_EOF_INT_CLR=0x40 IN_SUC_EOF_INT_CLR=0x20 OUTLINK_EOF_ERR_INT_CLR=0x1000 OUT_DONE_INT_CLR=0x80 OUT_DSCR_ERR_INT_CLR=0x400 OU... | - |
| `UHCI0_INT_ENA` | 0x3F41400C | 32 | DMA_INFIFO_FULL_WM_INT_ENA=0x10000 IN_DONE_INT_ENA=0x10 IN_DSCR_EMPTY_INT_ENA=0x800 IN_DSCR_ERR_INT_ENA=0x200 IN_ERR_EOF_INT_ENA=0x40 IN_SUC_EOF_INT_ENA=0x20 OUTLINK_EOF_ERR_INT_ENA=0x1000 OUT_DONE_INT_ENA=0x80 OUT_DSCR_ERR_INT_ENA=0x400 OU... | - |
| `UHCI0_INT_RAW` | 0x3F414004 | 32 | DMA_INFIFO_FULL_WM_INT_RAW=0x10000 IN_DONE_INT_RAW=0x10 IN_DSCR_EMPTY_INT_RAW=0x800 IN_DSCR_ERR_INT_RAW=0x200 IN_ERR_EOF_INT_RAW=0x40 IN_SUC_EOF_INT_RAW=0x20 OUTLINK_EOF_ERR_INT_RAW=0x1000 OUT_DONE_INT_RAW=0x80 OUT_DSCR_ERR_INT_RAW=0x400 OU... | - |
| `UHCI0_INT_ST` | 0x3F414008 | 32 | DMA_INFIFO_FULL_WM_INT_ST=0x10000 IN_DONE_INT_ST=0x10 IN_DSCR_EMPTY_INT_ST=0x800 IN_DSCR_ERR_INT_ST=0x200 IN_ERR_EOF_INT_ST=0x40 IN_SUC_EOF_INT_ST=0x20 OUTLINK_EOF_ERR_INT_ST=0x1000 OUT_DONE_INT_ST=0x80 OUT_DSCR_ERR_INT_ST=0x400 OUT_EOF_INT... | - |
| `UHCI0_PKT_THRES` | 0x3F4140C0 | 32 | PKT_THRS=0x1FFF | - |
| `UHCI0_Q0_WORD0` | 0x3F414078 | 32 | SEND_Q0_WORD0=0xFFFFFFFF | - |
| `UHCI0_Q0_WORD1` | 0x3F41407C | 32 | SEND_Q0_WORD1=0xFFFFFFFF | - |
| `UHCI0_Q1_WORD0` | 0x3F414080 | 32 | SEND_Q1_WORD0=0xFFFFFFFF | - |
| `UHCI0_Q1_WORD1` | 0x3F414084 | 32 | SEND_Q1_WORD1=0xFFFFFFFF | - |
| `UHCI0_Q2_WORD0` | 0x3F414088 | 32 | SEND_Q2_WORD0=0xFFFFFFFF | - |
| `UHCI0_Q2_WORD1` | 0x3F41408C | 32 | SEND_Q2_WORD1=0xFFFFFFFF | - |
| `UHCI0_Q3_WORD0` | 0x3F414090 | 32 | SEND_Q3_WORD0=0xFFFFFFFF | - |
| `UHCI0_Q3_WORD1` | 0x3F414094 | 32 | SEND_Q3_WORD1=0xFFFFFFFF | - |
| `UHCI0_Q4_WORD0` | 0x3F414098 | 32 | SEND_Q4_WORD0=0xFFFFFFFF | - |
| `UHCI0_Q4_WORD1` | 0x3F41409C | 32 | SEND_Q4_WORD1=0xFFFFFFFF | - |
| `UHCI0_Q5_WORD0` | 0x3F4140A0 | 32 | SEND_Q5_WORD0=0xFFFFFFFF | - |
| `UHCI0_Q5_WORD1` | 0x3F4140A4 | 32 | SEND_Q5_WORD1=0xFFFFFFFF | - |
| `UHCI0_Q6_WORD0` | 0x3F4140A8 | 32 | SEND_Q6_WORD0=0xFFFFFFFF | - |
| `UHCI0_Q6_WORD1` | 0x3F4140AC | 32 | SEND_Q6_WORD1=0xFFFFFFFF | - |
| `UHCI0_QUICK_SENT` | 0x3F414074 | 32 | ALWAYS_SEND_EN=0x80 ALWAYS_SEND_NUM=0x70 SINGLE_SEND_EN=0x8 SINGLE_SEND_NUM=0x7 | - |
| `UHCI0_RX_HEAD` | 0x3F414070 | 32 | RX_HEAD=0xFFFFFFFF | - |
| `UHCI0_STATE0` | 0x3F414030 | 32 | DECODE_STATE=0x70000000 INFIFO_CNT_DEBUG=0xF800000 INLINK_DSCR_ADDR=0x3FFFF IN_DSCR_STATE=0xC0000 IN_STATE=0x700000 | - |
| `UHCI0_STATE1` | 0x3F414034 | 32 | ENCODE_STATE=0x70000000 OUTFIFO_CNT=0xF800000 OUTLINK_DSCR_ADDR=0x3FFFF OUT_DSCR_STATE=0xC0000 OUT_STATE=0x700000 | - |
| `USB0_DAINT` | 0x60080818 | 32 | INEPINT0=0x1 INEPINT1=0x2 INEPINT2=0x4 INEPINT3=0x8 INEPINT4=0x10 INEPINT5=0x20 INEPINT6=0x40 OUTEPINT0=0x10000 OUTEPINT1=0x20000 OUTEPINT2=0x40000 OUTEPINT3=0x80000 OUTEPINT4=0x100000 OUTEPINT5=0x200000 OUTEPINT6=0x400000 | - |
| `USB0_DAINTMSK` | 0x6008081C | 32 | INEPMSK0=0x1 INEPMSK1=0x2 INEPMSK2=0x4 INEPMSK3=0x8 INEPMSK4=0x10 INEPMSK5=0x20 INEPMSK6=0x40 OUTEPMSK0=0x10000 OUTEPMSK1=0x20000 OUTEPMSK2=0x40000 OUTEPMSK3=0x80000 OUTEPMSK4=0x100000 OUTEPMSK5=0x200000 OUTEPMSK6=0x400000 | - |
| `USB0_DCFG` | 0x60080800 | 32 | DESCDMA=0x800000 DEVADDR=0x7F0 ENDEVOUTNAK=0x2000 EPMISCNT=0x7C0000 ERRATICINTMSK=0x8000 NZSTSOUTHSHK=0x4 PERFRLINT=0x1800 PERSCHINTVL=0x3000000 RESVALID=0xFC000000 XCVRDLY=0x4000 | - |
| `USB0_DCTL` | 0x60080804 | 32 | CGNPINNAK=0x100 CGOUTNAK=0x400 DEEPSLEEPBESLREJECT=0x40000 ENCOUNTONBNA=0x20000 GMC=0x6000 GNPINNAKSTS=0x4 GOUTNAKSTS=0x8 IGNRFRMNUM=0x8000 NAKONBBLE=0x10000 PWRONPRGDONE=0x800 RMTWKUPSIG=0x1 SFTDISCON=0x2 SGNPINNAK=0x80 SGOUTNAK=0x200 TSTC... | - |
| `USB0_DIEPCTL0` | 0x60080900 | 32 | DI_SNAK0=0x8000000 D_CNAK0=0x4000000 D_EPDIS0=0x40000000 D_EPENA0=0x80000000 D_EPTYPE0=0xC0000 D_MPS0=0x3 D_NAKSTS0=0x20000 D_STALL0=0x200000 D_TXFNUM0=0x3C00000 D_USBACTEP0=0x8000 | - |
| `USB0_DIEPCTL1` | 0x60080920 | 32 | DI_SETD0PID1=0x10000000 DI_SETD1PID1=0x20000000 DI_SNAK1=0x8000000 D_CNAK1=0x4000000 D_EPDIS1=0x40000000 D_EPENA1=0x80000000 D_EPTYPE1=0xC0000 D_MPS1=0x3 D_NAKSTS1=0x20000 D_STALL1=0x200000 D_TXFNUM1=0x3C00000 D_USBACTEP1=0x8000 | - |
| `USB0_DIEPCTL2` | 0x60080940 | 32 | DI_SETD0PID2=0x10000000 DI_SETD1PID2=0x20000000 DI_SNAK2=0x8000000 D_CNAK2=0x4000000 D_EPDIS2=0x40000000 D_EPENA2=0x80000000 D_EPTYPE2=0xC0000 D_MPS2=0x3 D_NAKSTS2=0x20000 D_STALL2=0x200000 D_TXFNUM2=0x3C00000 D_USBACTEP2=0x8000 | - |
| `USB0_DIEPCTL3` | 0x60080960 | 32 | DI_CNAK3=0x4000000 DI_EPDIS3=0x40000000 DI_EPENA3=0x80000000 DI_EPTYPE3=0xC0000 DI_MPS3=0x3 DI_NAKSTS3=0x20000 DI_SETD0PID3=0x10000000 DI_SETD1PID3=0x20000000 DI_SNAK3=0x8000000 DI_STALL3=0x200000 DI_TXFNUM3=0x3C00000 DI_USBACTEP3=0x8000 | - |
| `USB0_DIEPCTL4` | 0x60080980 | 32 | DI_SETD0PID4=0x10000000 DI_SETD1PID4=0x20000000 DI_SNAK4=0x8000000 D_CNAK4=0x4000000 D_EPDIS4=0x40000000 D_EPENA4=0x80000000 D_EPTYPE4=0xC0000 D_MPS4=0x3 D_NAKSTS4=0x20000 D_STALL4=0x200000 D_TXFNUM4=0x3C00000 D_USBACTEP4=0x8000 | - |
| `USB0_DIEPCTL5` | 0x600809A0 | 32 | DI_CNAK5=0x4000000 DI_EPDIS5=0x40000000 DI_EPENA5=0x80000000 DI_EPTYPE5=0xC0000 DI_MPS5=0x3 DI_NAKSTS5=0x20000 DI_SETD0PID5=0x10000000 DI_SETD1PID5=0x20000000 DI_SNAK5=0x8000000 DI_STALL5=0x200000 DI_TXFNUM5=0x3C00000 DI_USBACTEP5=0x8000 | - |
| `USB0_DIEPCTL6` | 0x600809C0 | 32 | DI_SETD0PID6=0x10000000 DI_SETD1PID6=0x20000000 DI_SNAK6=0x8000000 D_CNAK6=0x4000000 D_EPDIS6=0x40000000 D_EPENA6=0x80000000 D_EPTYPE6=0xC0000 D_MPS6=0x3 D_NAKSTS6=0x20000 D_STALL6=0x200000 D_TXFNUM6=0x3C00000 D_USBACTEP6=0x8000 | - |
| `USB0_DIEPDMA0` | 0x60080914 | 32 | D_DMAADDR0=0xFFFFFFFF | - |
| `USB0_DIEPDMA1` | 0x60080934 | 32 | D_DMAADDR1=0xFFFFFFFF | - |
| `USB0_DIEPDMA2` | 0x60080954 | 32 | D_DMAADDR2=0xFFFFFFFF | - |
| `USB0_DIEPDMA3` | 0x60080974 | 32 | D_DMAADDR3=0xFFFFFFFF | - |
| `USB0_DIEPDMA4` | 0x60080994 | 32 | D_DMAADDR4=0xFFFFFFFF | - |
| `USB0_DIEPDMA5` | 0x600809B4 | 32 | D_DMAADDR5=0xFFFFFFFF | - |
| `USB0_DIEPDMA6` | 0x600809D4 | 32 | D_DMAADDR6=0xFFFFFFFF | - |
| `USB0_DIEPDMAB0` | 0x6008091C | 32 | D_DMABUFFERADDR0=0xFFFFFFFF | - |
| `USB0_DIEPDMAB1` | 0x6008093C | 32 | D_DMABUFFERADDR1=0xFFFFFFFF | - |
| `USB0_DIEPDMAB2` | 0x6008095C | 32 | D_DMABUFFERADDR2=0xFFFFFFFF | - |
| `USB0_DIEPDMAB3` | 0x6008097C | 32 | D_DMABUFFERADDR3=0xFFFFFFFF | - |
| `USB0_DIEPDMAB4` | 0x6008099C | 32 | D_DMABUFFERADDR4=0xFFFFFFFF | - |
| `USB0_DIEPDMAB5` | 0x600809BC | 32 | D_DMABUFFERADDR5=0xFFFFFFFF | - |
| `USB0_DIEPDMAB6` | 0x600809DC | 32 | D_DMABUFFERADDR6=0xFFFFFFFF | - |
| `USB0_DIEPEMPMSK` | 0x60080834 | 32 | D_INEPTXFEMPMSK=0xFFFF | - |
| `USB0_DIEPINT0` | 0x60080908 | 32 | D_AHBERR0=0x4 D_BBLEERR0=0x1000 D_BNAINTR0=0x200 D_EPDISBLD0=0x2 D_INEPNAKEFF0=0x40 D_INTKNEPMIS0=0x20 D_INTKNTXFEMP0=0x10 D_NAKINTRPT0=0x2000 D_NYETINTRPT0=0x4000 D_PKTDRPSTS0=0x800 D_TIMEOUT0=0x8 D_TXFEMP0=0x80 D_TXFIFOUNDRN0=0x100 D_XFER... | - |
| `USB0_DIEPINT1` | 0x60080928 | 32 | D_AHBERR1=0x4 D_BBLEERR1=0x1000 D_BNAINTR1=0x200 D_EPDISBLD1=0x2 D_INEPNAKEFF1=0x40 D_INTKNEPMIS1=0x20 D_INTKNTXFEMP1=0x10 D_NAKINTRPT1=0x2000 D_NYETINTRPT1=0x4000 D_PKTDRPSTS1=0x800 D_TIMEOUT1=0x8 D_TXFEMP1=0x80 D_TXFIFOUNDRN1=0x100 D_XFER... | - |
| `USB0_DIEPINT2` | 0x60080948 | 32 | D_AHBERR2=0x4 D_BBLEERR2=0x1000 D_BNAINTR2=0x200 D_EPDISBLD2=0x2 D_INEPNAKEFF2=0x40 D_INTKNEPMIS2=0x20 D_INTKNTXFEMP2=0x10 D_NAKINTRPT2=0x2000 D_NYETINTRPT2=0x4000 D_PKTDRPSTS2=0x800 D_TIMEOUT2=0x8 D_TXFEMP2=0x80 D_TXFIFOUNDRN2=0x100 D_XFER... | - |
| `USB0_DIEPINT3` | 0x60080968 | 32 | D_AHBERR3=0x4 D_BBLEERR3=0x1000 D_BNAINTR3=0x200 D_EPDISBLD3=0x2 D_INEPNAKEFF3=0x40 D_INTKNEPMIS3=0x20 D_INTKNTXFEMP3=0x10 D_NAKINTRPT3=0x2000 D_NYETINTRPT3=0x4000 D_PKTDRPSTS3=0x800 D_TIMEOUT3=0x8 D_TXFEMP3=0x80 D_TXFIFOUNDRN3=0x100 D_XFER... | - |
| `USB0_DIEPINT4` | 0x60080988 | 32 | D_AHBERR4=0x4 D_BBLEERR4=0x1000 D_BNAINTR4=0x200 D_EPDISBLD4=0x2 D_INEPNAKEFF4=0x40 D_INTKNEPMIS4=0x20 D_INTKNTXFEMP4=0x10 D_NAKINTRPT4=0x2000 D_NYETINTRPT4=0x4000 D_PKTDRPSTS4=0x800 D_TIMEOUT4=0x8 D_TXFEMP4=0x80 D_TXFIFOUNDRN4=0x100 D_XFER... | - |
| `USB0_DIEPINT5` | 0x600809A8 | 32 | D_AHBERR5=0x4 D_BBLEERR5=0x1000 D_BNAINTR5=0x200 D_EPDISBLD5=0x2 D_INEPNAKEFF5=0x40 D_INTKNEPMIS5=0x20 D_INTKNTXFEMP5=0x10 D_NAKINTRPT5=0x2000 D_NYETINTRPT5=0x4000 D_PKTDRPSTS5=0x800 D_TIMEOUT5=0x8 D_TXFEMP5=0x80 D_TXFIFOUNDRN5=0x100 D_XFER... | - |
| `USB0_DIEPINT6` | 0x600809C8 | 32 | D_AHBERR6=0x4 D_BBLEERR6=0x1000 D_BNAINTR6=0x200 D_EPDISBLD6=0x2 D_INEPNAKEFF6=0x40 D_INTKNEPMIS6=0x20 D_INTKNTXFEMP6=0x10 D_NAKINTRPT6=0x2000 D_NYETINTRPT6=0x4000 D_PKTDRPSTS6=0x800 D_TIMEOUT6=0x8 D_TXFEMP6=0x80 D_TXFIFOUNDRN6=0x100 D_XFER... | - |
| `USB0_DIEPMSK` | 0x60080810 | 32 | BNAININTRMSK=0x200 DI_AHBERMSK=0x4 DI_EPDISBLDMSK=0x2 DI_NAKMSK=0x2000 DI_XFERCOMPLMSK=0x1 INEPNAKEFFMSK=0x40 INTKNEPMISMSK=0x20 INTKNTXFEMPMSK=0x10 TIMEOUTMSK=0x8 TXFIFOUNDRNMSK=0x100 | - |
| `USB0_DIEPTSIZ0` | 0x60080910 | 32 | D_PKTCNT0=0x180000 D_XFERSIZE0=0x7F | - |
| `USB0_DIEPTSIZ1` | 0x60080930 | 32 | D_PKTCNT1=0x180000 D_XFERSIZE1=0x7F | - |
| `USB0_DIEPTSIZ2` | 0x60080950 | 32 | D_PKTCNT2=0x180000 D_XFERSIZE2=0x7F | - |
| `USB0_DIEPTSIZ3` | 0x60080970 | 32 | D_PKTCNT3=0x180000 D_XFERSIZE3=0x7F | - |
| `USB0_DIEPTSIZ4` | 0x60080990 | 32 | D_PKTCNT4=0x180000 D_XFERSIZE4=0x7F | - |
| `USB0_DIEPTSIZ5` | 0x600809B0 | 32 | D_PKTCNT5=0x180000 D_XFERSIZE5=0x7F | - |
| `USB0_DIEPTSIZ6` | 0x600809D0 | 32 | D_PKTCNT6=0x180000 D_XFERSIZE6=0x7F | - |
| `USB0_DIEPTXF1` | 0x60080104 | 32 | INEP1TXFDEP=0xFFFF0000 INEP1TXFSTADDR=0xFFFF | - |
| `USB0_DIEPTXF2` | 0x60080108 | 32 | INEP2TXFDEP=0xFFFF0000 INEP2TXFSTADDR=0xFFFF | - |
| `USB0_DIEPTXF3` | 0x6008010C | 32 | INEP3TXFDEP=0xFFFF0000 INEP3TXFSTADDR=0xFFFF | - |
| `USB0_DIEPTXF4` | 0x60080110 | 32 | INEP4TXFDEP=0xFFFF0000 INEP4TXFSTADDR=0xFFFF | - |
| `USB0_DOEPCTL0` | 0x60080B00 | 32 | CNAK0=0x4000000 DO_SNAK0=0x8000000 EPDIS0=0x40000000 EPENA0=0x80000000 EPTYPE0=0xC0000 MPS0=0x3 NAKSTS0=0x20000 SNP0=0x100000 STALL0=0x200000 USBACTEP0=0x8000 | - |
| `USB0_DOEPCTL1` | 0x60080B20 | 32 | CNAK1=0x4000000 DO_SETD0PID1=0x10000000 DO_SETD1PID1=0x20000000 DO_SNAK1=0x8000000 EPDIS1=0x40000000 EPENA1=0x80000000 EPTYPE1=0xC0000 MPS1=0x7FF NAKSTS1=0x20000 SNP1=0x100000 STALL1=0x200000 USBACTEP1=0x8000 | - |
| `USB0_DOEPCTL2` | 0x60080B40 | 32 | CNAK2=0x4000000 DO_SETD0PID2=0x10000000 DO_SETD1PID2=0x20000000 DO_SNAK2=0x8000000 EPDIS2=0x40000000 EPENA2=0x80000000 EPTYPE2=0xC0000 MPS2=0x7FF NAKSTS2=0x20000 SNP2=0x100000 STALL2=0x200000 USBACTEP2=0x8000 | - |
| `USB0_DOEPCTL3` | 0x60080B60 | 32 | CNAK3=0x4000000 DO_SETD0PID3=0x10000000 DO_SETD1PID3=0x20000000 DO_SNAK3=0x8000000 EPDIS3=0x40000000 EPENA3=0x80000000 EPTYPE3=0xC0000 MPS3=0x7FF NAKSTS3=0x20000 SNP3=0x100000 STALL3=0x200000 USBACTEP3=0x8000 | - |
| `USB0_DOEPCTL4` | 0x60080B80 | 32 | CNAK4=0x4000000 DO_SETD0PID4=0x10000000 DO_SETD1PID4=0x20000000 DO_SNAK4=0x8000000 EPDIS4=0x40000000 EPENA4=0x80000000 EPTYPE4=0xC0000 MPS4=0x7FF NAKSTS4=0x20000 SNP4=0x100000 STALL4=0x200000 USBACTEP4=0x8000 | - |
| `USB0_DOEPCTL5` | 0x60080BA0 | 32 | CNAK5=0x4000000 DO_SETD0PID5=0x10000000 DO_SETD1PID5=0x20000000 DO_SNAK5=0x8000000 EPDIS5=0x40000000 EPENA5=0x80000000 EPTYPE5=0xC0000 MPS5=0x7FF NAKSTS5=0x20000 SNP5=0x100000 STALL5=0x200000 USBACTEP5=0x8000 | - |
| `USB0_DOEPCTL6` | 0x60080BC0 | 32 | CNAK6=0x4000000 DO_SETD0PID6=0x10000000 DO_SETD1PID6=0x20000000 DO_SNAK6=0x8000000 EPDIS6=0x40000000 EPENA6=0x80000000 EPTYPE6=0xC0000 MPS6=0x7FF NAKSTS6=0x20000 SNP6=0x100000 STALL6=0x200000 USBACTEP6=0x8000 | - |
| `USB0_DOEPDMA0` | 0x60080B14 | 32 | DMAADDR0=0xFFFFFFFF | - |
| `USB0_DOEPDMA1` | 0x60080B34 | 32 | DMAADDR1=0xFFFFFFFF | - |
| `USB0_DOEPDMA2` | 0x60080B54 | 32 | DMAADDR2=0xFFFFFFFF | - |
| `USB0_DOEPDMA3` | 0x60080B74 | 32 | DMAADDR3=0xFFFFFFFF | - |
| `USB0_DOEPDMA4` | 0x60080B94 | 32 | DMAADDR4=0xFFFFFFFF | - |
| `USB0_DOEPDMA5` | 0x60080BB4 | 32 | DMAADDR5=0xFFFFFFFF | - |
| `USB0_DOEPDMA6` | 0x60080BD4 | 32 | DMAADDR6=0xFFFFFFFF | - |
| `USB0_DOEPDMAB0` | 0x60080B1C | 32 | DMABUFFERADDR0=0xFFFFFFFF | - |
| `USB0_DOEPDMAB1` | 0x60080B3C | 32 | DMABUFFERADDR1=0xFFFFFFFF | - |
| `USB0_DOEPDMAB2` | 0x60080B5C | 32 | DMABUFFERADDR2=0xFFFFFFFF | - |
| `USB0_DOEPDMAB3` | 0x60080B7C | 32 | DMABUFFERADDR3=0xFFFFFFFF | - |
| `USB0_DOEPDMAB4` | 0x60080B9C | 32 | DMABUFFERADDR4=0xFFFFFFFF | - |
| `USB0_DOEPDMAB5` | 0x60080BBC | 32 | DMABUFFERADDR5=0xFFFFFFFF | - |
| `USB0_DOEPDMAB6` | 0x60080BDC | 32 | DMABUFFERADDR6=0xFFFFFFFF | - |
| `USB0_DOEPINT0` | 0x60080B08 | 32 | AHBERR0=0x4 BACK2BACKSETUP0=0x40 BBLEERR0=0x1000 BNAINTR0=0x200 EPDISBLD0=0x2 NAKINTRPT0=0x2000 NYEPINTRPT0=0x4000 OUTPKTERR0=0x100 OUTTKNEPDIS0=0x10 PKTDRPSTS0=0x800 SETUP0=0x8 STSPHSERCVD0=0x20 STUPPKTRCVD0=0x8000 XFERCOMPL0=0x1 | - |
| `USB0_DOEPINT1` | 0x60080B28 | 32 | AHBERR1=0x4 BACK2BACKSETUP1=0x40 BBLEERR1=0x1000 BNAINTR1=0x200 EPDISBLD1=0x2 NAKINTRPT1=0x2000 NYEPINTRPT1=0x4000 OUTPKTERR1=0x100 OUTTKNEPDIS1=0x10 PKTDRPSTS1=0x800 SETUP1=0x8 STSPHSERCVD1=0x20 STUPPKTRCVD1=0x8000 XFERCOMPL1=0x1 | - |
| `USB0_DOEPINT2` | 0x60080B48 | 32 | AHBERR2=0x4 BACK2BACKSETUP2=0x40 BBLEERR2=0x1000 BNAINTR2=0x200 EPDISBLD2=0x2 NAKINTRPT2=0x2000 NYEPINTRPT2=0x4000 OUTPKTERR2=0x100 OUTTKNEPDIS2=0x10 PKTDRPSTS2=0x800 SETUP2=0x8 STSPHSERCVD2=0x20 STUPPKTRCVD2=0x8000 XFERCOMPL2=0x1 | - |
| `USB0_DOEPINT3` | 0x60080B68 | 32 | AHBERR3=0x4 BACK2BACKSETUP3=0x40 BBLEERR3=0x1000 BNAINTR3=0x200 EPDISBLD3=0x2 NAKINTRPT3=0x2000 NYEPINTRPT3=0x4000 OUTPKTERR3=0x100 OUTTKNEPDIS3=0x10 PKTDRPSTS3=0x800 SETUP3=0x8 STSPHSERCVD3=0x20 STUPPKTRCVD3=0x8000 XFERCOMPL3=0x1 | - |
| `USB0_DOEPINT4` | 0x60080B88 | 32 | AHBERR4=0x4 BACK2BACKSETUP4=0x40 BBLEERR4=0x1000 BNAINTR4=0x200 EPDISBLD4=0x2 NAKINTRPT4=0x2000 NYEPINTRPT4=0x4000 OUTPKTERR4=0x100 OUTTKNEPDIS4=0x10 PKTDRPSTS4=0x800 SETUP4=0x8 STSPHSERCVD4=0x20 STUPPKTRCVD4=0x8000 XFERCOMPL4=0x1 | - |
| `USB0_DOEPINT5` | 0x60080BA8 | 32 | AHBERR5=0x4 BACK2BACKSETUP5=0x40 BBLEERR5=0x1000 BNAINTR5=0x200 EPDISBLD5=0x2 NAKINTRPT5=0x2000 NYEPINTRPT5=0x4000 OUTPKTERR5=0x100 OUTTKNEPDIS5=0x10 PKTDRPSTS5=0x800 SETUP5=0x8 STSPHSERCVD5=0x20 STUPPKTRCVD5=0x8000 XFERCOMPL5=0x1 | - |
| `USB0_DOEPINT6` | 0x60080BC8 | 32 | AHBERR6=0x4 BACK2BACKSETUP6=0x40 BBLEERR6=0x1000 BNAINTR6=0x200 EPDISBLD6=0x2 NAKINTRPT6=0x2000 NYEPINTRPT6=0x4000 OUTPKTERR6=0x100 OUTTKNEPDIS6=0x10 PKTDRPSTS6=0x800 SETUP6=0x8 STSPHSERCVD6=0x20 STUPPKTRCVD6=0x8000 XFERCOMPL6=0x1 | - |
| `USB0_DOEPMSK` | 0x60080814 | 32 | AHBERMSK=0x4 BACK2BACKSETUP=0x40 BBLEERRMSK=0x1000 BNAOUTINTRMSK=0x200 EPDISBLDMSK=0x2 NAKMSK=0x2000 NYETMSK=0x4000 OUTPKTERRMSK=0x100 OUTTKNEPDISMSK=0x10 SETUPMSK=0x8 STSPHSERCVDMSK=0x20 XFERCOMPLMSK=0x1 | - |
| `USB0_DOEPTSIZ0` | 0x60080B10 | 32 | PKTCNT0=0x80000 SUPCNT0=0x60000000 XFERSIZE0=0x7F | - |
| `USB0_DOEPTSIZ1` | 0x60080B30 | 32 | PKTCNT1=0x80000 SUPCNT1=0x60000000 XFERSIZE1=0x7F | - |
| `USB0_DOEPTSIZ2` | 0x60080B50 | 32 | PKTCNT2=0x80000 SUPCNT2=0x60000000 XFERSIZE2=0x7F | - |
| `USB0_DOEPTSIZ3` | 0x60080B70 | 32 | PKTCNT3=0x80000 SUPCNT3=0x60000000 XFERSIZE3=0x7F | - |
| `USB0_DOEPTSIZ4` | 0x60080B90 | 32 | PKTCNT4=0x80000 SUPCNT4=0x60000000 XFERSIZE4=0x7F | - |
| `USB0_DOEPTSIZ5` | 0x60080BB0 | 32 | PKTCNT5=0x80000 SUPCNT5=0x60000000 XFERSIZE5=0x7F | - |
| `USB0_DOEPTSIZ6` | 0x60080BD0 | 32 | PKTCNT6=0x80000 SUPCNT6=0x60000000 XFERSIZE6=0x7F | - |
| `USB0_DSTS` | 0x60080808 | 32 | DEVLNSTS=0xC00000 ENUMSPD=0x6 ERRTICERR=0x8 SOFFN=0x3FFF00 SUSPSTS=0x1 | - |
| `USB0_DTHRCTL` | 0x60080830 | 32 | AHBTHRRATIO=0x1800 ARBPRKEN=0x8000000 ISOTHREN=0x2 NONISOTHREN=0x1 RXTHREN=0x10000 RXTHRLEN=0x3FE0000 TXTHRLEN=0x7FC | - |
| `USB0_DTXFSTS0` | 0x60080918 | 32 | D_INEPTXFSPCAVAIL0=0xFFFF | - |
| `USB0_DTXFSTS1` | 0x60080938 | 32 | D_INEPTXFSPCAVAIL1=0xFFFF | - |
| `USB0_DTXFSTS2` | 0x60080958 | 32 | D_INEPTXFSPCAVAIL2=0xFFFF | - |
| `USB0_DTXFSTS3` | 0x60080978 | 32 | D_INEPTXFSPCAVAIL3=0xFFFF | - |
| `USB0_DTXFSTS4` | 0x60080998 | 32 | D_INEPTXFSPCAVAIL4=0xFFFF | - |
| `USB0_DTXFSTS5` | 0x600809B8 | 32 | D_INEPTXFSPCAVAIL5=0xFFFF | - |
| `USB0_DTXFSTS6` | 0x600809D8 | 32 | D_INEPTXFSPCAVAIL6=0xFFFF | - |
| `USB0_DVBUSDIS` | 0x60080828 | 32 | DVBUSDIS=0xFFFF | - |
| `USB0_DVBUSPULSE` | 0x6008082C | 32 | DVBUSPULSE=0xFFF | - |
| `USB0_GAHBCFG` | 0x60080008 | 32 | AHBSINGLE=0x800000 DMAEN=0x20 GLBLLNTRMSK=0x1 HBSTLEN=0x1E INVDESCENDIANESS=0x1000000 NOTIALLDMAWRIT=0x400000 NPTXFEMPLVL=0x80 PTXFEMPLVL=0x100 REMMEMSUPP=0x200000 | - |
| `USB0_GDFIFOCFG` | 0x6008005C | 32 | EPINFOBASEADDR=0xFFFF0000 GDFIFOCFG=0xFFFF | - |
| `USB0_GHWCFG1` | 0x60080044 | 32 | EPDIR=0xFFFFFFFF | - |
| `USB0_GHWCFG2` | 0x60080048 | 32 | DYNFIFOSIZING=0x80000 FSPHYTYPE=0x300 HSPHYTYPE=0xC0 MULTIPROCINTRPT=0x100000 NPTXQDEPTH=0xC00000 NUMDEVEPS=0x3C00 NUMHSTCHNL=0x3C000 OTGARCH=0x18 OTGMODE=0x7 OTG_ENABLE_IC_USB=0x80000000 PERIOSUPPORT=0x40000 PTXQDEPTH=0x3000000 SINGPNT=0x2... | - |
| `USB0_GHWCFG3` | 0x6008004C | 32 | ADPSUPPORT=0x1000 BCSUPPORT=0x4000 DFIFODEPTH=0xFFFF0000 HSICMODE=0x2000 I2CINTSEL=0x100 LPMMODE=0x8000 OPTFEATURE=0x400 OTGEN=0x80 PKTSIZEWIDTH=0x70 RSTTYPE=0x800 VNDCTLSUPT=0x200 XFERSIZEWIDTH=0xF | - |
| `USB0_GHWCFG4` | 0x60080050 | 32 | G_ACGSUPT=0x1000 G_AHBFREQ=0x20 G_AVALIDFLTR=0x400000 G_BVALIDFLTR=0x800000 G_DEDFIFOMODE=0x2000000 G_DESCDMA=0x80000000 G_DESCDMAENABLED=0x40000000 G_ENHANCEDLPMSUPT=0x2000 G_EXTENDEDHIBERNATION=0x80 G_HIBERNATION=0x40 G_IDDQFLTR=0x100000 ... | - |
| `USB0_GINTMSK` | 0x60080018 | 32 | CONIDSTSCHNGMSK=0x10000000 DISCONNINTMSK=0x20000000 ENUMDONEMSK=0x2000 EOPFMSK=0x8000 EPMISMSK=0x20000 ERLYSUSPMSK=0x400 FETSUSPMSK=0x400000 GINNAKEFFMSK=0x40 GOUTNACKEFFMSK=0x80 HCHINTMSK=0x2000000 IEPINTMSK=0x40000 INCOMPIPMSK=0x200000 IN... | - |
| `USB0_GINTSTS` | 0x60080014 | 32 | CONIDSTSCHNG=0x10000000 CURMOD_INT=0x1 DISCONNINT=0x20000000 ENUMDONE=0x2000 EOPF=0x8000 EPMIS=0x20000 ERLYSUSP=0x400 FETSUSP=0x400000 GINNAKEFF=0x40 GOUTNAKEFF=0x80 HCHLNT=0x2000000 IEPINT=0x40000 INCOMPIP=0x200000 INCOMPISOIN=0x100000 ISO... | - |
| `USB0_GNPTXFSIZ` | 0x60080028 | 32 | NPTXFDEP=0xFFFF0000 NPTXFSTADDR=0xFFFF | - |
| `USB0_GNPTXSTS` | 0x6008002C | 32 | NPTXFSPCAVAIL=0xFFFF NPTXQSPCAVAIL=0xF0000 NPTXQTOP=0x7F000000 | - |
| `USB0_GOTGCTL` | 0x60080000 | 32 | ASESVLD=0x40000 AVALIDOVEN=0x10 AVALIDOVVAL=0x20 BSESVLD=0x80000 BVALIDOVEN=0x40 BVALIDOVVAL=0x80 CONIDSTS=0x10000 CURMOD=0x200000 DBNCEFLTRBYPASS=0x8000 DBNCTIME=0x20000 DEVHNPEN=0x800 EHEN=0x1000 HNPREQ=0x200 HSTNEGSCS=0x100 HSTSETHNPEN=0... | - |
| `USB0_GOTGINT` | 0x60080004 | 32 | ADEVTOUTCHG=0x40000 DBNCEDONE=0x80000 HSTNEGDET=0x20000 HSTNEGSUCSTSCHNG=0x200 SESENDDET=0x4 SESREQSUCSTSCHNG=0x100 | - |
| `USB0_GRSTCTL` | 0x60080010 | 32 | AHBIDLE=0x80000000 CSFTRST=0x1 DMAREQ=0x40000000 FRMCNTRRST=0x4 PIUFSSFTRST=0x2 RXFFLSH=0x10 TXFFLSH=0x20 TXFNUM=0x7C0 | - |
| `USB0_GRXFSIZ` | 0x60080024 | 32 | RXFDEP=0xFFFF | - |
| `USB0_GRXSTSP` | 0x60080020 | 32 | BCNT=0x7FF0 CHNUM=0xF DPID=0x18000 FN=0x1E00000 PKTSTS=0x1E0000 | - |
| `USB0_GRXSTSR` | 0x6008001C | 32 | G_BCNT=0x7FF0 G_CHNUM=0xF G_DPID=0x18000 G_FN=0x1E00000 G_PKTSTS=0x1E0000 | - |
| `USB0_GSNPSID` | 0x60080040 | 32 | SYNOPSYSID=0xFFFFFFFF | - |
| `USB0_GUSBCFG` | 0x6008000C | 32 | CORRUPTTXPKT=0x80000000 FORCEDEVMODE=0x40000000 FORCEHSTMODE=0x20000000 FSINTF=0x20 HNPCAP=0x200 PHYIF=0x8 PHYSEL=0x40 SRPCAP=0x100 TERMSELDLPULSE=0x400000 TOUTCAL=0x7 TXENDDELAY=0x10000000 ULPI_UTMI_SEL=0x10 USBTRDTIM=0x3C00 | - |
| `USB0_HAINT` | 0x60080414 | 32 | HAINT=0xFF | - |
| `USB0_HAINTMSK` | 0x60080418 | 32 | HAINTMSK=0xFF | - |
| `USB0_HCCHAR0` | 0x60080500 | 32 | H_CHDIS0=0x40000000 H_CHENA0=0x80000000 H_DEVADDR0=0x1FC00000 H_EC0=0x200000 H_EPDIR0=0x8000 H_EPNUM0=0x7800 H_EPTYPE0=0xC0000 H_LSPDDEV0=0x20000 H_MPS0=0x7FF H_ODDFRM0=0x20000000 | - |
| `USB0_HCCHAR1` | 0x60080520 | 32 | H_CHDIS1=0x40000000 H_CHENA1=0x80000000 H_DEVADDR1=0x1FC00000 H_EC1=0x200000 H_EPDIR1=0x8000 H_EPNUM1=0x7800 H_EPTYPE1=0xC0000 H_LSPDDEV1=0x20000 H_MPS1=0x7FF H_ODDFRM1=0x20000000 | - |
| `USB0_HCCHAR2` | 0x60080540 | 32 | H_CHDIS2=0x40000000 H_CHENA2=0x80000000 H_DEVADDR2=0x1FC00000 H_EC2=0x200000 H_EPDIR2=0x8000 H_EPNUM2=0x7800 H_EPTYPE2=0xC0000 H_LSPDDEV2=0x20000 H_MPS2=0x7FF H_ODDFRM2=0x20000000 | - |
| `USB0_HCCHAR3` | 0x60080560 | 32 | H_CHDIS3=0x40000000 H_CHENA3=0x80000000 H_DEVADDR3=0x1FC00000 H_EC3=0x200000 H_EPDIR3=0x8000 H_EPNUM3=0x7800 H_EPTYPE3=0xC0000 H_LSPDDEV3=0x20000 H_MPS3=0x7FF H_ODDFRM3=0x20000000 | - |
| `USB0_HCCHAR4` | 0x60080580 | 32 | H_CHDIS4=0x40000000 H_CHENA4=0x80000000 H_DEVADDR4=0x1FC00000 H_EC4=0x200000 H_EPDIR4=0x8000 H_EPNUM4=0x7800 H_EPTYPE4=0xC0000 H_LSPDDEV4=0x20000 H_MPS4=0x7FF H_ODDFRM4=0x20000000 | - |
| `USB0_HCCHAR5` | 0x600805A0 | 32 | H_CHDIS5=0x40000000 H_CHENA5=0x80000000 H_DEVADDR5=0x1FC00000 H_EC5=0x200000 H_EPDIR5=0x8000 H_EPNUM5=0x7800 H_EPTYPE5=0xC0000 H_LSPDDEV5=0x20000 H_MPS5=0x7FF H_ODDFRM5=0x20000000 | - |
| `USB0_HCCHAR6` | 0x600805C0 | 32 | H_CHDIS6=0x40000000 H_CHENA6=0x80000000 H_DEVADDR6=0x1FC00000 H_EC6=0x200000 H_EPDIR6=0x8000 H_EPNUM6=0x7800 H_EPTYPE6=0xC0000 H_LSPDDEV6=0x20000 H_MPS6=0x7FF H_ODDFRM6=0x20000000 | - |
| `USB0_HCCHAR7` | 0x600805E0 | 32 | H_CHDIS7=0x40000000 H_CHENA7=0x80000000 H_DEVADDR7=0x1FC00000 H_EC7=0x200000 H_EPDIR7=0x8000 H_EPNUM7=0x7800 H_EPTYPE7=0xC0000 H_LSPDDEV7=0x20000 H_MPS7=0x7FF H_ODDFRM7=0x20000000 | - |
| `USB0_HCDMA0` | 0x60080514 | 32 | H_DMAADDR0=0xFFFFFFFF | - |
| `USB0_HCDMA1` | 0x60080534 | 32 | H_DMAADDR1=0xFFFFFFFF | - |
| `USB0_HCDMA2` | 0x60080554 | 32 | H_DMAADDR2=0xFFFFFFFF | - |
| `USB0_HCDMA3` | 0x60080574 | 32 | H_DMAADDR3=0xFFFFFFFF | - |
| `USB0_HCDMA4` | 0x60080594 | 32 | H_DMAADDR4=0xFFFFFFFF | - |
| `USB0_HCDMA5` | 0x600805B4 | 32 | H_DMAADDR5=0xFFFFFFFF | - |
| `USB0_HCDMA6` | 0x600805D4 | 32 | H_DMAADDR6=0xFFFFFFFF | - |
| `USB0_HCDMA7` | 0x600805F4 | 32 | H_DMAADDR7=0xFFFFFFFF | - |
| `USB0_HCDMAB0` | 0x6008051C | 32 | H_HCDMAB0=0xFFFFFFFF | - |
| `USB0_HCDMAB1` | 0x6008053C | 32 | H_HCDMAB1=0xFFFFFFFF | - |
| `USB0_HCDMAB2` | 0x6008055C | 32 | H_HCDMAB2=0xFFFFFFFF | - |
| `USB0_HCDMAB3` | 0x6008057C | 32 | H_HCDMAB3=0xFFFFFFFF | - |
| `USB0_HCDMAB4` | 0x6008059C | 32 | H_HCDMAB4=0xFFFFFFFF | - |
| `USB0_HCDMAB5` | 0x600805BC | 32 | H_HCDMAB5=0xFFFFFFFF | - |
| `USB0_HCDMAB6` | 0x600805DC | 32 | H_HCDMAB6=0xFFFFFFFF | - |
| `USB0_HCDMAB7` | 0x600805FC | 32 | H_HCDMAB7=0xFFFFFFFF | - |
| `USB0_HCFG` | 0x60080400 | 32 | H_DESCDMA=0x800000 H_ENA32KHZS=0x80 H_FRLISTEN=0x3000000 H_FSLSPCLKSEL=0x3 H_FSLSSUPP=0x4 H_MODECHTIMEN=0x80000000 H_PERSCHEDENA=0x4000000 | - |
| `USB0_HCINT0` | 0x60080508 | 32 | H_ACK0=0x20 H_AHBERR0=0x4 H_BBLERR0=0x100 H_BNAINTR0=0x800 H_CHHLTD0=0x2 H_DATATGLERR0=0x400 H_DESC_LST_ROLLINTR0=0x2000 H_FRMOVRUN0=0x200 H_NACK0=0x10 H_NYET0=0x40 H_STALL0=0x8 H_XACTERR0=0x80 H_XCS_XACT_ERR0=0x1000 H_XFERCOMPL0=0x1 | - |
| `USB0_HCINT1` | 0x60080528 | 32 | H_ACK1=0x20 H_AHBERR1=0x4 H_BBLERR1=0x100 H_BNAINTR1=0x800 H_CHHLTD1=0x2 H_DATATGLERR1=0x400 H_DESC_LST_ROLLINTR1=0x2000 H_FRMOVRUN1=0x200 H_NACK1=0x10 H_NYET1=0x40 H_STALL1=0x8 H_XACTERR1=0x80 H_XCS_XACT_ERR1=0x1000 H_XFERCOMPL1=0x1 | - |
| `USB0_HCINT2` | 0x60080548 | 32 | H_ACK2=0x20 H_AHBERR2=0x4 H_BBLERR2=0x100 H_BNAINTR2=0x800 H_CHHLTD2=0x2 H_DATATGLERR2=0x400 H_DESC_LST_ROLLINTR2=0x2000 H_FRMOVRUN2=0x200 H_NACK2=0x10 H_NYET2=0x40 H_STALL2=0x8 H_XACTERR2=0x80 H_XCS_XACT_ERR2=0x1000 H_XFERCOMPL2=0x1 | - |
| `USB0_HCINT3` | 0x60080568 | 32 | H_ACK3=0x20 H_AHBERR3=0x4 H_BBLERR3=0x100 H_BNAINTR3=0x800 H_CHHLTD3=0x2 H_DATATGLERR3=0x400 H_DESC_LST_ROLLINTR3=0x2000 H_FRMOVRUN3=0x200 H_NACK3=0x10 H_NYET3=0x40 H_STALL3=0x8 H_XACTERR3=0x80 H_XCS_XACT_ERR3=0x1000 H_XFERCOMPL3=0x1 | - |
| `USB0_HCINT4` | 0x60080588 | 32 | H_ACK4=0x20 H_AHBERR4=0x4 H_BBLERR4=0x100 H_BNAINTR4=0x800 H_CHHLTD4=0x2 H_DATATGLERR4=0x400 H_DESC_LST_ROLLINTR4=0x2000 H_FRMOVRUN4=0x200 H_NACK4=0x10 H_NYET4=0x40 H_STALL4=0x8 H_XACTERR4=0x80 H_XCS_XACT_ERR4=0x1000 H_XFERCOMPL4=0x1 | - |
| `USB0_HCINT5` | 0x600805A8 | 32 | H_ACK5=0x20 H_AHBERR5=0x4 H_BBLERR5=0x100 H_BNAINTR5=0x800 H_CHHLTD5=0x2 H_DATATGLERR5=0x400 H_DESC_LST_ROLLINTR5=0x2000 H_FRMOVRUN5=0x200 H_NACK5=0x10 H_NYET5=0x40 H_STALL5=0x8 H_XACTERR5=0x80 H_XCS_XACT_ERR5=0x1000 H_XFERCOMPL5=0x1 | - |
| `USB0_HCINT6` | 0x600805C8 | 32 | H_ACK6=0x20 H_AHBERR6=0x4 H_BBLERR6=0x100 H_BNAINTR6=0x800 H_CHHLTD6=0x2 H_DATATGLERR6=0x400 H_DESC_LST_ROLLINTR6=0x2000 H_FRMOVRUN6=0x200 H_NACK6=0x10 H_NYET6=0x40 H_STALL6=0x8 H_XACTERR6=0x80 H_XCS_XACT_ERR6=0x1000 H_XFERCOMPL6=0x1 | - |
| `USB0_HCINT7` | 0x600805E8 | 32 | H_ACK7=0x20 H_AHBERR7=0x4 H_BBLERR7=0x100 H_BNAINTR7=0x800 H_CHHLTD7=0x2 H_DATATGLERR7=0x400 H_DESC_LST_ROLLINTR7=0x2000 H_FRMOVRUN7=0x200 H_NACK7=0x10 H_NYET7=0x40 H_STALL7=0x8 H_XACTERR7=0x80 H_XCS_XACT_ERR7=0x1000 H_XFERCOMPL7=0x1 | - |
| `USB0_HCINTMSK0` | 0x6008050C | 32 | H_ACKMSK0=0x20 H_AHBERRMSK0=0x4 H_BBLERRMSK0=0x100 H_BNAINTRMSK0=0x800 H_CHHLTDMSK0=0x2 H_DATATGLERRMSK0=0x400 H_DESC_LST_ROLLINTRMSK0=0x2000 H_FRMOVRUNMSK0=0x200 H_NAKMSK0=0x10 H_NYETMSK0=0x40 H_STALLMSK0=0x8 H_XACTERRMSK0=0x80 H_XFERCOMPL... | - |
| `USB0_HCINTMSK1` | 0x6008052C | 32 | H_ACKMSK1=0x20 H_AHBERRMSK1=0x4 H_BBLERRMSK1=0x100 H_BNAINTRMSK1=0x800 H_CHHLTDMSK1=0x2 H_DATATGLERRMSK1=0x400 H_DESC_LST_ROLLINTRMSK1=0x2000 H_FRMOVRUNMSK1=0x200 H_NAKMSK1=0x10 H_NYETMSK1=0x40 H_STALLMSK1=0x8 H_XACTERRMSK1=0x80 H_XFERCOMPL... | - |
| `USB0_HCINTMSK2` | 0x6008054C | 32 | H_ACKMSK2=0x20 H_AHBERRMSK2=0x4 H_BBLERRMSK2=0x100 H_BNAINTRMSK2=0x800 H_CHHLTDMSK2=0x2 H_DATATGLERRMSK2=0x400 H_DESC_LST_ROLLINTRMSK2=0x2000 H_FRMOVRUNMSK2=0x200 H_NAKMSK2=0x10 H_NYETMSK2=0x40 H_STALLMSK2=0x8 H_XACTERRMSK2=0x80 H_XFERCOMPL... | - |
| `USB0_HCINTMSK3` | 0x6008056C | 32 | H_ACKMSK3=0x20 H_AHBERRMSK3=0x4 H_BBLERRMSK3=0x100 H_BNAINTRMSK3=0x800 H_CHHLTDMSK3=0x2 H_DATATGLERRMSK3=0x400 H_DESC_LST_ROLLINTRMSK3=0x2000 H_FRMOVRUNMSK3=0x200 H_NAKMSK3=0x10 H_NYETMSK3=0x40 H_STALLMSK3=0x8 H_XACTERRMSK3=0x80 H_XFERCOMPL... | - |
| `USB0_HCINTMSK4` | 0x6008058C | 32 | H_ACKMSK4=0x20 H_AHBERRMSK4=0x4 H_BBLERRMSK4=0x100 H_BNAINTRMSK4=0x800 H_CHHLTDMSK4=0x2 H_DATATGLERRMSK4=0x400 H_DESC_LST_ROLLINTRMSK4=0x2000 H_FRMOVRUNMSK4=0x200 H_NAKMSK4=0x10 H_NYETMSK4=0x40 H_STALLMSK4=0x8 H_XACTERRMSK4=0x80 H_XFERCOMPL... | - |
| `USB0_HCINTMSK5` | 0x600805AC | 32 | H_ACKMSK5=0x20 H_AHBERRMSK5=0x4 H_BBLERRMSK5=0x100 H_BNAINTRMSK5=0x800 H_CHHLTDMSK5=0x2 H_DATATGLERRMSK5=0x400 H_DESC_LST_ROLLINTRMSK5=0x2000 H_FRMOVRUNMSK5=0x200 H_NAKMSK5=0x10 H_NYETMSK5=0x40 H_STALLMSK5=0x8 H_XACTERRMSK5=0x80 H_XFERCOMPL... | - |
| `USB0_HCINTMSK6` | 0x600805CC | 32 | H_ACKMSK6=0x20 H_AHBERRMSK6=0x4 H_BBLERRMSK6=0x100 H_BNAINTRMSK6=0x800 H_CHHLTDMSK6=0x2 H_DATATGLERRMSK6=0x400 H_DESC_LST_ROLLINTRMSK6=0x2000 H_FRMOVRUNMSK6=0x200 H_NAKMSK6=0x10 H_NYETMSK6=0x40 H_STALLMSK6=0x8 H_XACTERRMSK6=0x80 H_XFERCOMPL... | - |
| `USB0_HCINTMSK7` | 0x600805EC | 32 | H_ACKMSK7=0x20 H_AHBERRMSK7=0x4 H_BBLERRMSK7=0x100 H_BNAINTRMSK7=0x800 H_CHHLTDMSK7=0x2 H_DATATGLERRMSK7=0x400 H_DESC_LST_ROLLINTRMSK7=0x2000 H_FRMOVRUNMSK7=0x200 H_NAKMSK7=0x10 H_NYETMSK7=0x40 H_STALLMSK7=0x8 H_XACTERRMSK7=0x80 H_XFERCOMPL... | - |
| `USB0_HCTSIZ0` | 0x60080510 | 32 | H_DOPNG0=0x80000000 H_PID0=0x60000000 H_PKTCNT0=0x1FF80000 H_XFERSIZE0=0x7FFFF | - |
| `USB0_HCTSIZ1` | 0x60080530 | 32 | H_DOPNG1=0x80000000 H_PID1=0x60000000 H_PKTCNT1=0x1FF80000 H_XFERSIZE1=0x7FFFF | - |
| `USB0_HCTSIZ2` | 0x60080550 | 32 | H_DOPNG2=0x80000000 H_PID2=0x60000000 H_PKTCNT2=0x1FF80000 H_XFERSIZE2=0x7FFFF | - |
| `USB0_HCTSIZ3` | 0x60080570 | 32 | H_DOPNG3=0x80000000 H_PID3=0x60000000 H_PKTCNT3=0x1FF80000 H_XFERSIZE3=0x7FFFF | - |
| `USB0_HCTSIZ4` | 0x60080590 | 32 | H_DOPNG4=0x80000000 H_PID4=0x60000000 H_PKTCNT4=0x1FF80000 H_XFERSIZE4=0x7FFFF | - |
| `USB0_HCTSIZ5` | 0x600805B0 | 32 | H_DOPNG5=0x80000000 H_PID5=0x60000000 H_PKTCNT5=0x1FF80000 H_XFERSIZE5=0x7FFFF | - |
| `USB0_HCTSIZ6` | 0x600805D0 | 32 | H_DOPNG6=0x80000000 H_PID6=0x60000000 H_PKTCNT6=0x1FF80000 H_XFERSIZE6=0x7FFFF | - |
| `USB0_HCTSIZ7` | 0x600805F0 | 32 | H_DOPNG7=0x80000000 H_PID7=0x60000000 H_PKTCNT7=0x1FF80000 H_XFERSIZE7=0x7FFFF | - |
| `USB0_HFIR` | 0x60080404 | 32 | FRINT=0xFFFF HFIRRLDCTRL=0x10000 | - |
| `USB0_HFLBADDR` | 0x6008041C | 32 | HFLBADDR=0xFFFFFFFF | - |
| `USB0_HFNUM` | 0x60080408 | 32 | FRNUM=0x3FFF FRREM=0xFFFF0000 | - |
| `USB0_HPRT` | 0x60080440 | 32 | PRTCONNDET=0x2 PRTCONNSTS=0x1 PRTENA=0x4 PRTENCHNG=0x8 PRTLNSTS=0xC00 PRTOVRCURRACT=0x10 PRTOVRCURRCHNG=0x20 PRTPWR=0x1000 PRTRES=0x40 PRTRST=0x100 PRTSPD=0x60000 PRTSUSP=0x80 PRTTSTCTL=0x1E000 | - |
| `USB0_HPTXFSIZ` | 0x60080100 | 32 | PTXFSIZE=0xFFFF0000 PTXFSTADDR=0xFFFF | - |
| `USB0_HPTXSTS` | 0x60080410 | 32 | PTXFSPCAVAIL=0xFFFF PTXQSPCAVAIL=0x1F0000 PTXQTOP=0xFF000000 | - |
| `USB0_PCGCCTL` | 0x60080E00 | 32 | GATEHCLK=0x2 L1SUSPENDED=0x80 PHYSLEEP=0x40 PWRCLMP=0x4 RESETAFTERSUSP=0x100 RSTPDWNMODULE=0x8 STOPPCLK=0x1 | - |
| `USB_WRAP_DATE` | 0x3F4393FC | 32 | USB_WRAP_DATE=0xFFFFFFFF | - |
| `USB_WRAP_OTG_CONF` | 0x3F439000 | 32 | AHB_CLK_FORCE_ON=0x80000 CLK_EN=0x80000000 DBNCE_FLTR_BYPASS=0x10 DFIFO_FORCE_PD=0x8 DFIFO_FORCE_PU=0x400000 DM_PULLDOWN=0x10000 DM_PULLUP=0x8000 DP_PULLDOWN=0x4000 DP_PULLUP=0x2000 EXCHG_PINS=0x40 EXCHG_PINS_OVERRIDE=0x20 PAD_PULL_OVERRIDE... | - |
| `USB_WRAP_TEST_CONF` | 0x3F439004 | 32 | TEST_ENABLE=0x1 TEST_RX_DM=0x40 TEST_RX_DP=0x20 TEST_RX_RCV=0x10 TEST_TX_DM=0x8 TEST_TX_DP=0x4 TEST_USB_OE=0x2 | - |
| `XTS_AES_DATE` | 0x6003A25C | 32 | DATE=0x3FFFFFFF | - |
| `XTS_AES_DESTINATION` | 0x6003A244 | 32 | DESTINATION=0x1 | - |
| `XTS_AES_DESTROY` | 0x6003A254 | 32 | DESTROY=0x1 | - |
| `XTS_AES_LINESIZE` | 0x6003A240 | 32 | LINESIZE=0x3 | - |
| `XTS_AES_PHYSICAL_ADDRESS` | 0x6003A248 | 32 | PHYSICAL_ADDRESS=0x3FFFFFFF | - |
| `XTS_AES_PLAIN_0` | 0x6003A200 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_1` | 0x6003A204 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_10` | 0x6003A228 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_11` | 0x6003A22C | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_12` | 0x6003A230 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_13` | 0x6003A234 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_14` | 0x6003A238 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_15` | 0x6003A23C | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_2` | 0x6003A208 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_3` | 0x6003A20C | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_4` | 0x6003A210 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_5` | 0x6003A214 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_6` | 0x6003A218 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_7` | 0x6003A21C | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_8` | 0x6003A220 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_PLAIN_9` | 0x6003A224 | 32 | PLAIN_0=0xFFFFFFFF | - |
| `XTS_AES_RELEASE` | 0x6003A250 | 32 | RELEASE=0x1 | - |
| `XTS_AES_STATE` | 0x6003A258 | 32 | STATE=0x3 | - |
| `XTS_AES_TRIGGER` | 0x6003A24C | 32 | TRIGGER=0x1 | - |
