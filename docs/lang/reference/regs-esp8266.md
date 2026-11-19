# Registers - esp8266

| Name | Address | Bits | Fields | Drives pins |
|---|---|---|---|---|
| `DPORT_DPORT_CTL` | 0x3FF00014 | 32 | DPORT_CTL_DOUBLE_CLK=0x1 | - |
| `DPORT_EDGE_INT_ENABLE` | 0x3FF00004 | 32 | Register=0xFFFFFFFF | - |
| `EFUSE_EFUSE_DATA0` | 0x3FF00050 | 32 | Register=0xFFFFFFFF | - |
| `EFUSE_EFUSE_DATA1` | 0x3FF00054 | 32 | Register=0xFFFFFFFF | - |
| `EFUSE_EFUSE_DATA2` | 0x3FF00058 | 32 | Register=0xFFFFFFFF | - |
| `EFUSE_EFUSE_DATA3` | 0x3FF0005C | 32 | Register=0xFFFFFFFF | - |
| `GPIO_GPIO_ENABLE` | 0x6000030C | 32 | GPIO_ENABLE_DATA=0xFFFF GPIO_SDIO_SEL=0x3F0000 | - |
| `GPIO_GPIO_ENABLE_W1TC` | 0x60000314 | 32 | GPIO_ENABLE_DATA_W1TC=0xFFFF | - |
| `GPIO_GPIO_ENABLE_W1TS` | 0x60000310 | 32 | GPIO_ENABLE_DATA_W1TS=0xFFFF | - |
| `GPIO_GPIO_IN` | 0x60000318 | 32 | GPIO_IN_DATA=0xFFFF GPIO_STRAPPING=0xFFFF0000 | - |
| `GPIO_GPIO_OUT` | 0x60000300 | 32 | GPIO_BT_SEL=0xFFFF0000 GPIO_OUT_DATA=0xFFFF | - |
| `GPIO_GPIO_OUT_W1TC` | 0x60000308 | 32 | GPIO_OUT_DATA_W1TC=0xFFFF | - |
| `GPIO_GPIO_OUT_W1TS` | 0x60000304 | 32 | GPIO_OUT_DATA_W1TS=0xFFFF | - |
| `GPIO_GPIO_PIN0` | 0x60000328 | 32 | GPIO_PIN0_DRIVER=0x4 GPIO_PIN0_INT_TYPE=0x380 GPIO_PIN0_SOURCE=0x1 GPIO_PIN0_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN1` | 0x6000032C | 32 | GPIO_PIN1_DRIVER=0x4 GPIO_PIN1_INT_TYPE=0x380 GPIO_PIN1_SOURCE=0x1 GPIO_PIN1_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN10` | 0x60000350 | 32 | GPIO_PIN10_DRIVER=0x4 GPIO_PIN10_INT_TYPE=0x380 GPIO_PIN10_SOURCE=0x1 GPIO_PIN10_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN11` | 0x60000354 | 32 | GPIO_PIN11_DRIVER=0x4 GPIO_PIN11_INT_TYPE=0x380 GPIO_PIN11_SOURCE=0x1 GPIO_PIN11_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN12` | 0x60000358 | 32 | GPIO_PIN12_DRIVER=0x4 GPIO_PIN12_INT_TYPE=0x380 GPIO_PIN12_SOURCE=0x1 GPIO_PIN12_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN13` | 0x6000035C | 32 | GPIO_PIN13_DRIVER=0x4 GPIO_PIN13_INT_TYPE=0x380 GPIO_PIN13_SOURCE=0x1 GPIO_PIN13_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN14` | 0x60000360 | 32 | GPIO_PIN14_DRIVER=0x4 GPIO_PIN14_INT_TYPE=0x380 GPIO_PIN14_SOURCE=0x1 GPIO_PIN14_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN15` | 0x60000364 | 32 | GPIO_PIN15_DRIVER=0x4 GPIO_PIN15_INT_TYPE=0x380 GPIO_PIN15_SOURCE=0x1 GPIO_PIN15_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN2` | 0x60000330 | 32 | GPIO_PIN2_DRIVER=0x4 GPIO_PIN2_INT_TYPE=0x380 GPIO_PIN2_SOURCE=0x1 GPIO_PIN2_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN3` | 0x60000334 | 32 | GPIO_PIN3_DRIVER=0x4 GPIO_PIN3_INT_TYPE=0x380 GPIO_PIN3_SOURCE=0x1 GPIO_PIN3_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN4` | 0x60000338 | 32 | GPIO_PIN4_DRIVER=0x4 GPIO_PIN4_INT_TYPE=0x380 GPIO_PIN4_SOURCE=0x1 GPIO_PIN4_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN5` | 0x6000033C | 32 | GPIO_PIN5_DRIVER=0x4 GPIO_PIN5_INT_TYPE=0x380 GPIO_PIN5_SOURCE=0x1 GPIO_PIN5_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN6` | 0x60000340 | 32 | GPIO_PIN6_DRIVER=0x4 GPIO_PIN6_INT_TYPE=0x380 GPIO_PIN6_SOURCE=0x1 GPIO_PIN6_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN7` | 0x60000344 | 32 | GPIO_PIN7_DRIVER=0x4 GPIO_PIN7_INT_TYPE=0x380 GPIO_PIN7_SOURCE=0x1 GPIO_PIN7_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN8` | 0x60000348 | 32 | GPIO_PIN8_DRIVER=0x4 GPIO_PIN8_INT_TYPE=0x380 GPIO_PIN8_SOURCE=0x1 GPIO_PIN8_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_PIN9` | 0x6000034C | 32 | GPIO_PIN9_DRIVER=0x4 GPIO_PIN9_INT_TYPE=0x380 GPIO_PIN9_SOURCE=0x1 GPIO_PIN9_WAKEUP_ENABLE=0x400 | - |
| `GPIO_GPIO_RTC_CALIB_SYNC` | 0x6000036C | 32 | RTC_CALIB_START=0x80000000 RTC_PERIOD_NUM=0x3FF | - |
| `GPIO_GPIO_RTC_CALIB_VALUE` | 0x60000370 | 32 | RTC_CALIB_RDY=0x80000000 RTC_CALIB_RDY_REAL=0x40000000 RTC_CALIB_VALUE=0xFFFFF | - |
| `GPIO_GPIO_SIGMA_DELTA` | 0x60000368 | 32 | SIGMA_DELTA_ENABLE=0x10000 SIGMA_DELTA_PRESCALAR=0xFF00 SIGMA_DELTA_TARGET=0xFF | - |
| `GPIO_GPIO_STATUS` | 0x6000031C | 32 | GPIO_STATUS_INTERRUPT=0xFFFF | - |
| `GPIO_GPIO_STATUS_W1TC` | 0x60000324 | 32 | GPIO_STATUS_INTERRUPT_W1TC=0xFFFF | - |
| `GPIO_GPIO_STATUS_W1TS` | 0x60000320 | 32 | GPIO_STATUS_INTERRUPT_W1TS=0xFFFF | - |
| `I2S_I2SCONF` | 0x60000E08 | 32 | I2S_BCK_DIV_NUM=0xFC00000 I2S_BITS_MOD=0xF000 I2S_CLKM_DIV_NUM=0x3F0000 I2S_I2S_RX_FIFO_RESET=0x8 I2S_I2S_RX_RESET=0x2 I2S_I2S_RX_START=0x200 I2S_I2S_TX_FIFO_RESET=0x4 I2S_I2S_TX_RESET=0x1 I2S_I2S_TX_START=0x100 I2S_MSB_RIGHT=0x80 I2S_RECE_... | - |
| `I2S_I2SCONF_SIGLE_DATA` | 0x60000E28 | 32 | I2S_I2S_SIGLE_DATA=0xFFFFFFFF | - |
| `I2S_I2SINT_CLR` | 0x60000E18 | 32 | I2S_I2S_PUT_DATA_INT_CLR=0x2 I2S_I2S_RX_REMPTY_INT_CLR=0x8 I2S_I2S_RX_WFULL_INT_CLR=0x4 I2S_I2S_TAKE_DATA_INT_CLR=0x1 I2S_I2S_TX_REMPTY_INT_CLR=0x20 I2S_I2S_TX_WFULL_INT_CLR=0x10 | - |
| `I2S_I2SINT_ENA` | 0x60000E14 | 32 | I2S_I2S_RX_REMPTY_INT_ENA=0x8 I2S_I2S_RX_TAKE_DATA_INT_ENA=0x1 I2S_I2S_RX_WFULL_INT_ENA=0x4 I2S_I2S_TX_PUT_DATA_INT_ENA=0x2 I2S_I2S_TX_REMPTY_INT_ENA=0x20 I2S_I2S_TX_WFULL_INT_ENA=0x10 | - |
| `I2S_I2SINT_RAW` | 0x60000E0C | 32 | I2S_I2S_RX_REMPTY_INT_RAW=0x8 I2S_I2S_RX_TAKE_DATA_INT_RAW=0x1 I2S_I2S_RX_WFULL_INT_RAW=0x4 I2S_I2S_TX_PUT_DATA_INT_RAW=0x2 I2S_I2S_TX_REMPTY_INT_RAW=0x20 I2S_I2S_TX_WFULL_INT_RAW=0x10 | - |
| `I2S_I2SINT_ST` | 0x60000E10 | 32 | I2S_I2S_RX_REMPTY_INT_ST=0x8 I2S_I2S_RX_TAKE_DATA_INT_ST=0x1 I2S_I2S_RX_WFULL_INT_ST=0x4 I2S_I2S_TX_PUT_DATA_INT_ST=0x2 I2S_I2S_TX_REMPTY_INT_ST=0x20 I2S_I2S_TX_WFULL_INT_ST=0x10 | - |
| `I2S_I2SRXEOF_NUM` | 0x60000E24 | 32 | I2S_I2S_RX_EOF_NUM=0xFFFFFFFF | - |
| `I2S_I2SRXFIFO` | 0x60000E04 | 32 | Register=0xFFFFFFFF | - |
| `I2S_I2STIMING` | 0x60000E1C | 32 | I2S_RECE_BCK_IN_DELAY=0x30 I2S_RECE_BCK_OUT_DELAY=0xC0000 I2S_RECE_DSYNC_SW=0x200000 I2S_RECE_SD_IN_DELAY=0x300 I2S_RECE_WS_IN_DELAY=0xC0 I2S_RECE_WS_OUT_DELAY=0x30000 I2S_TRANS_BCK_IN_DELAY=0x3 I2S_TRANS_BCK_IN_INV=0x400000 I2S_TRANS_BCK_O... | - |
| `I2S_I2STXFIFO` | 0x60000E00 | 32 | Register=0xFFFFFFFF | - |
| `I2S_I2S_FIFO_CONF` | 0x60000E20 | 32 | I2S_I2S_DSCR_EN=0x1000 I2S_I2S_RX_DATA_NUM=0x3F I2S_I2S_RX_FIFO_MOD=0x70000 I2S_I2S_TX_DATA_NUM=0xFC0 I2S_I2S_TX_FIFO_MOD=0xE000 | - |
| `IO_MUX_IO_MUX_CONF` | 0x60000800 | 32 | SPI0_CLK_EQU_SYS_CLK=0x100 SPI1_CLK_EQU_SYS_CLK=0x200 | - |
| `IO_MUX_IO_MUX_GPIO0` | 0x60000834 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_GPIO2` | 0x60000838 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_GPIO4` | 0x6000083C | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_GPIO5` | 0x60000840 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_MTCK` | 0x60000808 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_MTDI` | 0x60000804 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_MTDO` | 0x60000810 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_MTMS` | 0x6000080C | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_SD_CLK` | 0x6000081C | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_SD_CMD` | 0x60000830 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_SD_DATA0` | 0x60000820 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_SD_DATA1` | 0x60000824 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_SD_DATA2` | 0x60000828 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_SD_DATA3` | 0x6000082C | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_U0RXD` | 0x60000814 | 32 | Register=0xFFFFFFFF | - |
| `IO_MUX_IO_MUX_U0TXD` | 0x60000818 | 32 | Register=0xFFFFFFFF | - |
| `RTC_RTC_STATE1` | 0x60000714 | 32 | Register=0xFFFFFFFF | - |
| `RTC_RTC_STORE0` | 0x60000730 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_AHB_TEST` | 0x60000B54 | 32 | SLC_AHB_TESTADDR=0x30 SLC_AHB_TESTMODE=0x7 | - |
| `SLC_SLC_BRIDGE_CONF` | 0x60000B44 | 32 | SLC_FIFO_MAP_ENA=0xF00 SLC_TXEOF_ENA=0x3F SLC_TX_DUMMY_MODE=0x1000 SLC_TX_PUSH_IDLE_NUM=0xFFFF0000 | - |
| `SLC_SLC_CONF0` | 0x60000B00 | 32 | SLC_AHBM_FIFO_RST=0x4 SLC_AHBM_RST=0x8 SLC_DATA_BURST_EN=0x200 SLC_DSCR_BURST_EN=0x100 SLC_MODE=0x3000 SLC_RXLINK_RST=0x2 SLC_RX_AUTO_WRBACK=0x40 SLC_RX_LOOP_TEST=0x20 SLC_RX_NO_RESTART_CLR=0x80 SLC_TXLINK_RST=0x1 SLC_TX_LOOP_TEST=0x10 | - |
| `SLC_SLC_CONF1` | 0x60000B38 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_DATE` | 0x60000B78 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_ID` | 0x60000B7C | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_INTVEC_TOHOST` | 0x60000B2C | 32 | SLC_TOHOST_INTVEC=0xFF | - |
| `SLC_SLC_INT_CLR` | 0x60000B10 | 32 | SLC_FRHOST_BIT0_INT_CLR=0x1 SLC_FRHOST_BIT1_INT_CLR=0x2 SLC_FRHOST_BIT2_INT_CLR=0x4 SLC_FRHOST_BIT3_INT_CLR=0x8 SLC_FRHOST_BIT4_INT_CLR=0x10 SLC_FRHOST_BIT5_INT_CLR=0x20 SLC_FRHOST_BIT6_INT_CLR=0x40 SLC_FRHOST_BIT7_INT_CLR=0x80 SLC_RX_DONE_... | - |
| `SLC_SLC_INT_ENA` | 0x60000B0C | 32 | SLC_FRHOST_BIT0_INT_ENA=0x1 SLC_FRHOST_BIT1_INT_ENA=0x2 SLC_FRHOST_BIT2_INT_ENA=0x4 SLC_FRHOST_BIT3_INT_ENA=0x8 SLC_FRHOST_BIT4_INT_ENA=0x10 SLC_FRHOST_BIT5_INT_ENA=0x20 SLC_FRHOST_BIT6_INT_ENA=0x40 SLC_FRHOST_BIT7_INT_ENA=0x80 SLC_RX_DONE_... | - |
| `SLC_SLC_INT_RAW` | 0x60000B04 | 32 | SLC_FRHOST_BIT0_INT_RAW=0x1 SLC_FRHOST_BIT1_INT_RAW=0x2 SLC_FRHOST_BIT2_INT_RAW=0x4 SLC_FRHOST_BIT3_INT_RAW=0x8 SLC_FRHOST_BIT4_INT_RAW=0x10 SLC_FRHOST_BIT5_INT_RAW=0x20 SLC_FRHOST_BIT6_INT_RAW=0x40 SLC_FRHOST_BIT7_INT_RAW=0x80 SLC_RX_DONE_... | - |
| `SLC_SLC_INT_STATUS` | 0x60000B08 | 32 | SLC_FRHOST_BIT0_INT_ST=0x1 SLC_FRHOST_BIT1_INT_ST=0x2 SLC_FRHOST_BIT2_INT_ST=0x4 SLC_FRHOST_BIT3_INT_ST=0x8 SLC_FRHOST_BIT4_INT_ST=0x10 SLC_FRHOST_BIT5_INT_ST=0x20 SLC_FRHOST_BIT6_INT_ST=0x40 SLC_FRHOST_BIT7_INT_ST=0x80 SLC_RX_DONE_INT_ST=0... | - |
| `SLC_SLC_RXLINK_DSCR` | 0x60000B6C | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_RXLINK_DSCR_BF0` | 0x60000B70 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_RXLINK_DSCR_BF1` | 0x60000B74 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_RX_DSCR_CONF` | 0x60000B5C | 32 | SLC_INFOR_NO_REPLACE=0x200 SLC_TOKEN_NO_REPLACE=0x100 | - |
| `SLC_SLC_RX_EOF_BFR_DES_ADDR` | 0x60000B50 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_RX_EOF_DES_ADDR` | 0x60000B48 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_RX_FIFO_PUSH` | 0x60000B18 | 32 | SLC_RXFIFO_PUSH=0x10000 SLC_RXFIFO_WDATA=0x1FF | - |
| `SLC_SLC_RX_LINK` | 0x60000B24 | 32 | SLC_RXLINK_ADDR=0xFFFFF SLC_RXLINK_PARK=0x80000000 SLC_RXLINK_RESTART=0x40000000 SLC_RXLINK_START=0x20000000 SLC_RXLINK_STOP=0x10000000 | - |
| `SLC_SLC_RX_STATUS` | 0x60000B14 | 32 | SLC_RX_EMPTY=0x2 SLC_RX_FULL=0x1 | - |
| `SLC_SLC_SDIO_ST` | 0x60000B58 | 32 | SLC_BUS_ST=0x7000 SLC_CMD_ST=0x7 SLC_FUNC_ST=0xF0 SLC_SDIO_WAKEUP=0x100 | - |
| `SLC_SLC_STATE0` | 0x60000B3C | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_STATE1` | 0x60000B40 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_TOKEN0` | 0x60000B30 | 32 | SLC_TOKEN0=0xFFF0000 SLC_TOKEN0_LOCAL_INC=0x2000 SLC_TOKEN0_LOCAL_INC_MORE=0x4000 SLC_TOKEN0_LOCAL_WDATA=0xFFF SLC_TOKEN0_LOCAL_WR=0x1000 | - |
| `SLC_SLC_TOKEN1` | 0x60000B34 | 32 | SLC_TOKEN1=0xFFF0000 SLC_TOKEN1_LOCAL_INC=0x2000 SLC_TOKEN1_LOCAL_INC_MORE=0x4000 SLC_TOKEN1_LOCAL_WDATA=0xFFF SLC_TOKEN1_LOCAL_WR=0x1000 | - |
| `SLC_SLC_TXLINK_DSCR` | 0x60000B60 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_TXLINK_DSCR_BF0` | 0x60000B64 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_TXLINK_DSCR_BF1` | 0x60000B68 | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_TX_EOF_DES_ADDR` | 0x60000B4C | 32 | Register=0xFFFFFFFF | - |
| `SLC_SLC_TX_FIFO_POP` | 0x60000B20 | 32 | SLC_TXFIFO_POP=0x10000 SLC_TXFIFO_RDATA=0x7FF | - |
| `SLC_SLC_TX_LINK` | 0x60000B28 | 32 | SLC_TXLINK_ADDR=0xFFFFF SLC_TXLINK_PARK=0x80000000 SLC_TXLINK_RESTART=0x40000000 SLC_TXLINK_START=0x20000000 SLC_TXLINK_STOP=0x10000000 | - |
| `SLC_SLC_TX_STATUS` | 0x60000B1C | 32 | SLC_TX_EMPTY=0x2 SLC_TX_FULL=0x1 | - |
| `SPI0_SPI_ADDR` | 0x60000204 | 32 | iodata_start_addr=0xFFFFFFFF | - |
| `SPI0_SPI_CLOCK` | 0x60000218 | 32 | spi_clk_equ_sysclk=0x80000000 spi_clkcnt_H=0xFC0 spi_clkcnt_L=0x3F spi_clkcnt_N=0x3F000 spi_clkdiv_pre=0x7FFC0000 | - |
| `SPI0_SPI_CMD` | 0x60000200 | 32 | spi_usr=0x40000 | - |
| `SPI0_SPI_CTRL` | 0x60000208 | 32 | spi_dio_mode=0x800000 spi_dout_mode=0x4000 spi_fastrd_mode=0x2000 spi_qio_mode=0x1000000 spi_qout_mode=0x100000 spi_rd_bit_order=0x2000000 spi_wr_bit_order=0x4000000 | - |
| `SPI0_SPI_CTRL2` | 0x60000214 | 32 | spi_cs_delay_mode=0xC000000 spi_cs_delay_num=0xF0000000 spi_miso_delay_mode=0x30000 spi_miso_delay_num=0x1C0000 spi_mosi_delay_mode=0x600000 spi_mosi_delay_num=0x3800000 | - |
| `SPI0_SPI_EXT3` | 0x600002FC | 32 | reg_int_hold_ena=0x3 | - |
| `SPI0_SPI_PIN` | 0x6000022C | 32 | spi_cs0_dis=0x1 spi_cs1_dis=0x2 spi_cs2_dis=0x4 | - |
| `SPI0_SPI_RD_STATUS` | 0x60000210 | 32 | slv_rd_status=0xFFFFFFFF | - |
| `SPI0_SPI_SLAVE` | 0x60000230 | 32 | slv_cmd_define=0x8000000 slv_rd_buf_done=0x1 slv_rd_sta_done=0x4 slv_wr_buf_done=0x2 slv_wr_sta_done=0x8 spi_int_en=0x3E0 spi_slave_mode=0x40000000 spi_sync_reset=0x80000000 spi_trans_cnt=0x7800000 spi_trans_done=0x10 | - |
| `SPI0_SPI_SLAVE1` | 0x60000234 | 32 | slv_buf_bitlen=0x1FF0000 slv_rd_addr_bitlen=0xFC00 slv_rdbuf_dummy_en=0x1 slv_rdsta_dummy_en=0x4 slv_status_bitlen=0xF8000000 slv_wr_addr_bitlen=0x3F0 slv_wrbuf_dummy_en=0x2 slv_wrsta_dummy_en=0x8 | - |
| `SPI0_SPI_SLAVE2` | 0x60000238 | 32 | slv_rdbuf_dummy_cyclelen=0xFF0000 slv_rdsta_dummy_cyclelen=0xFF slv_wrbuf_dummy_cyclelen=0xFF000000 slv_wrsta_dummy_cyclelen=0xFF00 | - |
| `SPI0_SPI_SLAVE3` | 0x6000023C | 32 | slv_rdbuf_cmd_value=0xFF slv_rdsta_cmd_value=0xFF0000 slv_wrbuf_cmd_value=0xFF00 slv_wrsta_cmd_value=0xFF000000 | - |
| `SPI0_SPI_USER` | 0x6000021C | 32 | reg_usr_miso_highpart=0x1000000 reg_usr_mosi_highpart=0x2000000 spi_ck_i_edge=0x40 spi_fwrite_dio=0x4000 spi_fwrite_dual=0x1000 spi_fwrite_qio=0x8000 spi_fwrite_quad=0x2000 spi_rd_byte_order=0x400 spi_sio=0x10000 spi_usr_addr=0x40000000 spi... | - |
| `SPI0_SPI_USER1` | 0x60000220 | 32 | reg_usr_addr_bitlen=0xFC000000 reg_usr_dummy_cyclelen=0xFF reg_usr_miso_bitlen=0x1FF00 reg_usr_mosi_bitlen=0x3FE0000 | - |
| `SPI0_SPI_USER2` | 0x60000224 | 32 | reg_usr_command_bitlen=0xF0000000 reg_usr_command_value=0xFFFF | - |
| `SPI0_SPI_W0` | 0x60000240 | 32 | spi_w0=0xFFFFFFFF | - |
| `SPI0_SPI_W1` | 0x60000244 | 32 | spi_w1=0xFFFFFFFF | - |
| `SPI0_SPI_W10` | 0x60000268 | 32 | spi_w10=0xFFFFFFFF | - |
| `SPI0_SPI_W11` | 0x6000026C | 32 | spi_w11=0xFFFFFFFF | - |
| `SPI0_SPI_W12` | 0x60000270 | 32 | spi_w12=0xFFFFFFFF | - |
| `SPI0_SPI_W13` | 0x60000274 | 32 | spi_w13=0xFFFFFFFF | - |
| `SPI0_SPI_W14` | 0x60000278 | 32 | spi_w14=0xFFFFFFFF | - |
| `SPI0_SPI_W15` | 0x6000027C | 32 | spi_w15=0xFFFFFFFF | - |
| `SPI0_SPI_W2` | 0x60000248 | 32 | spi_w2=0xFFFFFFFF | - |
| `SPI0_SPI_W3` | 0x6000024C | 32 | spi_w3=0xFFFFFFFF | - |
| `SPI0_SPI_W4` | 0x60000250 | 32 | spi_w4=0xFFFFFFFF | - |
| `SPI0_SPI_W5` | 0x60000254 | 32 | spi_w5=0xFFFFFFFF | - |
| `SPI0_SPI_W6` | 0x60000258 | 32 | spi_w6=0xFFFFFFFF | - |
| `SPI0_SPI_W7` | 0x6000025C | 32 | spi_w7=0xFFFFFFFF | - |
| `SPI0_SPI_W8` | 0x60000260 | 32 | spi_w8=0xFFFFFFFF | - |
| `SPI0_SPI_W9` | 0x60000264 | 32 | spi_w9=0xFFFFFFFF | - |
| `SPI0_SPI_WR_STATUS` | 0x60000228 | 32 | slv_wr_status=0xFFFFFFFF | - |
| `SPI1_SPI_ADDR` | 0x60000104 | 32 | iodata_start_addr=0xFFFFFFFF | - |
| `SPI1_SPI_CLOCK` | 0x60000118 | 32 | spi_clk_equ_sysclk=0x80000000 spi_clkcnt_H=0xFC0 spi_clkcnt_L=0x3F spi_clkcnt_N=0x3F000 spi_clkdiv_pre=0x7FFC0000 | - |
| `SPI1_SPI_CMD` | 0x60000100 | 32 | spi_usr=0x40000 | - |
| `SPI1_SPI_CTRL` | 0x60000108 | 32 | spi_dio_mode=0x800000 spi_dout_mode=0x4000 spi_fastrd_mode=0x2000 spi_qio_mode=0x1000000 spi_qout_mode=0x100000 spi_rd_bit_order=0x2000000 spi_wr_bit_order=0x4000000 | - |
| `SPI1_SPI_CTRL2` | 0x60000114 | 32 | spi_cs_delay_mode=0xC000000 spi_cs_delay_num=0xF0000000 spi_miso_delay_mode=0x30000 spi_miso_delay_num=0x1C0000 spi_mosi_delay_mode=0x600000 spi_mosi_delay_num=0x3800000 | - |
| `SPI1_SPI_EXT3` | 0x600001FC | 32 | reg_int_hold_ena=0x3 | - |
| `SPI1_SPI_PIN` | 0x6000012C | 32 | spi_cs0_dis=0x1 spi_cs1_dis=0x2 spi_cs2_dis=0x4 | - |
| `SPI1_SPI_RD_STATUS` | 0x60000110 | 32 | slv_rd_status=0xFFFFFFFF | - |
| `SPI1_SPI_SLAVE` | 0x60000130 | 32 | slv_cmd_define=0x8000000 slv_rd_buf_done=0x1 slv_rd_sta_done=0x4 slv_wr_buf_done=0x2 slv_wr_sta_done=0x8 spi_int_en=0x3E0 spi_slave_mode=0x40000000 spi_sync_reset=0x80000000 spi_trans_cnt=0x7800000 spi_trans_done=0x10 | - |
| `SPI1_SPI_SLAVE1` | 0x60000134 | 32 | slv_buf_bitlen=0x1FF0000 slv_rd_addr_bitlen=0xFC00 slv_rdbuf_dummy_en=0x1 slv_rdsta_dummy_en=0x4 slv_status_bitlen=0xF8000000 slv_wr_addr_bitlen=0x3F0 slv_wrbuf_dummy_en=0x2 slv_wrsta_dummy_en=0x8 | - |
| `SPI1_SPI_SLAVE2` | 0x60000138 | 32 | slv_rdbuf_dummy_cyclelen=0xFF0000 slv_rdsta_dummy_cyclelen=0xFF slv_wrbuf_dummy_cyclelen=0xFF000000 slv_wrsta_dummy_cyclelen=0xFF00 | - |
| `SPI1_SPI_SLAVE3` | 0x6000013C | 32 | slv_rdbuf_cmd_value=0xFF slv_rdsta_cmd_value=0xFF0000 slv_wrbuf_cmd_value=0xFF00 slv_wrsta_cmd_value=0xFF000000 | - |
| `SPI1_SPI_USER` | 0x6000011C | 32 | reg_usr_miso_highpart=0x1000000 reg_usr_mosi_highpart=0x2000000 spi_ck_i_edge=0x40 spi_fwrite_dio=0x4000 spi_fwrite_dual=0x1000 spi_fwrite_qio=0x8000 spi_fwrite_quad=0x2000 spi_rd_byte_order=0x400 spi_sio=0x10000 spi_usr_addr=0x40000000 spi... | - |
| `SPI1_SPI_USER1` | 0x60000120 | 32 | reg_usr_addr_bitlen=0xFC000000 reg_usr_dummy_cyclelen=0xFF reg_usr_miso_bitlen=0x1FF00 reg_usr_mosi_bitlen=0x3FE0000 | - |
| `SPI1_SPI_USER2` | 0x60000124 | 32 | reg_usr_command_bitlen=0xF0000000 reg_usr_command_value=0xFFFF | - |
| `SPI1_SPI_W0` | 0x60000140 | 32 | spi_w0=0xFFFFFFFF | - |
| `SPI1_SPI_W1` | 0x60000144 | 32 | spi_w1=0xFFFFFFFF | - |
| `SPI1_SPI_W10` | 0x60000168 | 32 | spi_w10=0xFFFFFFFF | - |
| `SPI1_SPI_W11` | 0x6000016C | 32 | spi_w11=0xFFFFFFFF | - |
| `SPI1_SPI_W12` | 0x60000170 | 32 | spi_w12=0xFFFFFFFF | - |
| `SPI1_SPI_W13` | 0x60000174 | 32 | spi_w13=0xFFFFFFFF | - |
| `SPI1_SPI_W14` | 0x60000178 | 32 | spi_w14=0xFFFFFFFF | - |
| `SPI1_SPI_W15` | 0x6000017C | 32 | spi_w15=0xFFFFFFFF | - |
| `SPI1_SPI_W2` | 0x60000148 | 32 | spi_w2=0xFFFFFFFF | - |
| `SPI1_SPI_W3` | 0x6000014C | 32 | spi_w3=0xFFFFFFFF | - |
| `SPI1_SPI_W4` | 0x60000150 | 32 | spi_w4=0xFFFFFFFF | - |
| `SPI1_SPI_W5` | 0x60000154 | 32 | spi_w5=0xFFFFFFFF | - |
| `SPI1_SPI_W6` | 0x60000158 | 32 | spi_w6=0xFFFFFFFF | - |
| `SPI1_SPI_W7` | 0x6000015C | 32 | spi_w7=0xFFFFFFFF | - |
| `SPI1_SPI_W8` | 0x60000160 | 32 | spi_w8=0xFFFFFFFF | - |
| `SPI1_SPI_W9` | 0x60000164 | 32 | spi_w9=0xFFFFFFFF | - |
| `SPI1_SPI_WR_STATUS` | 0x60000128 | 32 | slv_wr_status=0xFFFFFFFF | - |
| `TIMER_FRC1_COUNT` | 0x60000604 | 32 | frc1_count=0x7FFFFF | - |
| `TIMER_FRC1_CTRL` | 0x60000608 | 32 | frc1_ctrl=0xFF frc1_int=0x100 | - |
| `TIMER_FRC1_INT` | 0x6000060C | 32 | frc1_int_clr_mask=0x1 | - |
| `TIMER_FRC1_LOAD` | 0x60000600 | 32 | frc1_load_value=0x7FFFFF | - |
| `TIMER_FRC2_ALARM` | 0x60000630 | 32 | frc2_alarm=0xFFFFFFFF | - |
| `TIMER_FRC2_COUNT` | 0x60000624 | 32 | frc2_count=0xFFFFFFFF | - |
| `TIMER_FRC2_CTRL` | 0x60000628 | 32 | frc2_ctrl=0xFF frc2_int=0x100 | - |
| `TIMER_FRC2_INT` | 0x6000062C | 32 | frc2_int_clr_mask=0x1 | - |
| `TIMER_FRC2_LOAD` | 0x60000620 | 32 | frc2_load_value=0xFFFFFFFF | - |
| `UART0_UART_AUTOBAUD` | 0x60000018 | 32 | autobaud_en=0x1 glitch_filt=0xFF00 | - |
| `UART0_UART_CLKDIV` | 0x60000014 | 32 | uart_clkdiv=0xFFFFF | - |
| `UART0_UART_CONF0` | 0x60000020 | 32 | bit_num=0xC parity=0x1 parity_en=0x2 rxfifo_rst=0x20000 stop_bit_num=0x30 sw_dtr=0x80 sw_rts=0x40 tx_flow_en=0x8000 txd_brk=0x100 txfifo_rst=0x40000 uart_cts_inv=0x100000 uart_dsr_inv=0x200000 uart_dtr_inv=0x1000000 uart_loopback=0x4000 uar... | - |
| `UART0_UART_CONF1` | 0x60000024 | 32 | rx_flow_en=0x800000 rx_flow_thrhd=0x7F0000 rx_tout_en=0x80000000 rx_tout_thrhd=0x7F000000 rxfifo_full_thrhd=0x7F txfifo_empty_thrhd=0x7F00 | - |
| `UART0_UART_DATE` | 0x60000078 | 32 | uart_date=0xFFFFFFFF | - |
| `UART0_UART_FIFO` | 0x60000000 | 32 | rxfifo_rd_byte=0xFF | - |
| `UART0_UART_HIGHPULSE` | 0x6000002C | 32 | highpulse_min_cnt=0xFFFFF | - |
| `UART0_UART_ID` | 0x6000007C | 32 | uart_id=0xFFFFFFFF | - |
| `UART0_UART_INT_CLR` | 0x60000010 | 32 | brk_det_int_clr=0x80 cts_chg_int_clr=0x40 dsr_chg_int_clr=0x20 frm_err_int_clr=0x8 parity_err_int_clr=0x4 rxfifo_full_int_clr=0x1 rxfifo_ovf_int_clr=0x10 rxfifo_tout_int_clr=0x100 txfifo_empty_int_clr=0x2 | - |
| `UART0_UART_INT_ENA` | 0x6000000C | 32 | brk_det_int_ena=0x80 cts_chg_int_ena=0x40 dsr_chg_int_ena=0x20 frm_err_int_ena=0x8 parity_err_int_ena=0x4 rxfifo_full_int_ena=0x1 rxfifo_ovf_int_ena=0x10 rxfifo_tout_int_ena=0x100 txfifo_empty_int_ena=0x2 | - |
| `UART0_UART_INT_RAW` | 0x60000004 | 32 | brk_det_int_raw=0x80 cts_chg_int_raw=0x40 dsr_chg_int_raw=0x20 frm_err_int_raw=0x8 parity_err_int_raw=0x4 rxfifo_full_int_raw=0x1 rxfifo_ovf_int_raw=0x10 rxfifo_tout_int_raw=0x100 txfifo_empty_int_raw=0x2 | - |
| `UART0_UART_INT_ST` | 0x60000008 | 32 | brk_det_int_st=0x80 cts_chg_int_st=0x40 dsr_chg_int_st=0x20 frm_err_int_st=0x8 parity_err_int_st=0x4 rxfifo_full_int_st=0x1 rxfifo_ovf_int_st=0x10 rxfifo_tout_int_st=0x100 txfifo_empty_int_st=0x2 | - |
| `UART0_UART_LOWPULSE` | 0x60000028 | 32 | lowpulse_min_cnt=0xFFFFF | - |
| `UART0_UART_RXD_CNT` | 0x60000030 | 32 | rxd_edge_cnt=0x3FF | - |
| `UART0_UART_STATUS` | 0x6000001C | 32 | ctsn=0x4000 dsrn=0x2000 dtrn=0x20000000 rtsn=0x40000000 rxd=0x8000 rxfifo_cnt=0xFF txd=0x80000000 txfifo_cnt=0xFF0000 | - |
| `UART1_UART_AUTOBAUD` | 0x60000F18 | 32 | autobaud_en=0x1 glitch_filt=0xFF00 | - |
| `UART1_UART_CLKDIV` | 0x60000F14 | 32 | uart_clkdiv=0xFFFFF | - |
| `UART1_UART_CONF0` | 0x60000F20 | 32 | bit_num=0xC parity=0x1 parity_en=0x2 rxfifo_rst=0x20000 stop_bit_num=0x30 sw_dtr=0x80 sw_rts=0x40 tx_flow_en=0x8000 txd_brk=0x100 txfifo_rst=0x40000 uart_cts_inv=0x100000 uart_dsr_inv=0x200000 uart_dtr_inv=0x1000000 uart_loopback=0x4000 uar... | - |
| `UART1_UART_CONF1` | 0x60000F24 | 32 | rx_flow_en=0x800000 rx_flow_thrhd=0x7F0000 rx_tout_en=0x80000000 rx_tout_thrhd=0x7F000000 rxfifo_full_thrhd=0x7F txfifo_empty_thrhd=0x7F00 | - |
| `UART1_UART_DATE` | 0x60000F78 | 32 | uart_date=0xFFFFFFFF | - |
| `UART1_UART_FIFO` | 0x60000F00 | 32 | rxfifo_rd_byte=0xFF | - |
| `UART1_UART_HIGHPULSE` | 0x60000F2C | 32 | highpulse_min_cnt=0xFFFFF | - |
| `UART1_UART_ID` | 0x60000F7C | 32 | uart_id=0xFFFFFFFF | - |
| `UART1_UART_INT_CLR` | 0x60000F10 | 32 | brk_det_int_clr=0x80 cts_chg_int_clr=0x40 dsr_chg_int_clr=0x20 frm_err_int_clr=0x8 parity_err_int_clr=0x4 rxfifo_full_int_clr=0x1 rxfifo_ovf_int_clr=0x10 rxfifo_tout_int_clr=0x100 txfifo_empty_int_clr=0x2 | - |
| `UART1_UART_INT_ENA` | 0x60000F0C | 32 | brk_det_int_ena=0x80 cts_chg_int_ena=0x40 dsr_chg_int_ena=0x20 frm_err_int_ena=0x8 parity_err_int_ena=0x4 rxfifo_full_int_ena=0x1 rxfifo_ovf_int_ena=0x10 rxfifo_tout_int_ena=0x100 txfifo_empty_int_ena=0x2 | - |
| `UART1_UART_INT_RAW` | 0x60000F04 | 32 | brk_det_int_raw=0x80 cts_chg_int_raw=0x40 dsr_chg_int_raw=0x20 frm_err_int_raw=0x8 parity_err_int_raw=0x4 rxfifo_full_int_raw=0x1 rxfifo_ovf_int_raw=0x10 rxfifo_tout_int_raw=0x100 txfifo_empty_int_raw=0x2 | - |
| `UART1_UART_INT_ST` | 0x60000F08 | 32 | brk_det_int_st=0x80 cts_chg_int_st=0x40 dsr_chg_int_st=0x20 frm_err_int_st=0x8 parity_err_int_st=0x4 rxfifo_full_int_st=0x1 rxfifo_ovf_int_st=0x10 rxfifo_tout_int_st=0x100 txfifo_empty_int_st=0x2 | - |
| `UART1_UART_LOWPULSE` | 0x60000F28 | 32 | lowpulse_min_cnt=0xFFFFF | - |
| `UART1_UART_RXD_CNT` | 0x60000F30 | 32 | rxd_edge_cnt=0x3FF | - |
| `UART1_UART_STATUS` | 0x60000F1C | 32 | ctsn=0x4000 dsrn=0x2000 dtrn=0x20000000 rtsn=0x40000000 rxd=0x8000 rxfifo_cnt=0xFF txd=0x80000000 txfifo_cnt=0xFF0000 | - |
| `WDT_WDT_CTL` | 0x60000900 | 32 | Register=0xFFFFFFFF | - |
| `WDT_WDT_OP` | 0x60000904 | 32 | Register=0xFFFFFFFF | - |
| `WDT_WDT_OP_ND` | 0x60000908 | 32 | Register=0xFFFFFFFF | - |
| `WDT_WDT_RST` | 0x60000914 | 32 | Register=0xFFFFFFFF | - |
