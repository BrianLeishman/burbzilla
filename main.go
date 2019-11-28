package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/mrmorphic/hwio"
)

func readADC(device hwio.I2CDevice, channel int, gain float64, dataRate int) int {
	if channel < 0 || channel > 3 {
		panic("channel must be a value within 0-3")
	}

	// return self._read(channel + 0x04, gain, data_rate, ADS1x15_CONFIG_MODE_SINGLE)
	mux := channel + 0x04

	config := ADS1x15_CONFIG_OS_SINGLE

	config |= (mux & 0x07) << ADS1x15_CONFIG_MUX_OFFSET

	if _, ok := ADS1x15_CONFIG_GAIN[gain]; !ok {
		panic("gain must be one of: 2/3, 1, 2, 4, 8, 16")
	}
	config |= ADS1x15_CONFIG_GAIN[gain]

	config |= ADS1x15_CONFIG_MODE_SINGLE

	if _, ok := ADS1115_CONFIG_DR[dataRate]; !ok {
		panic("data rate must be one of: 8, 16, 32, 64, 128, 250, 475, 860")
	}
	config |= ADS1115_CONFIG_DR[dataRate]

	config |= ADS1x15_CONFIG_COMP_QUE_DISABLE

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32((config>>8)&0xFF))
	if err := device.Write(ADS1x15_POINTER_CONFIG, buf[:2]); err != nil {
		panic(err)
	}
	binary.LittleEndian.PutUint32(buf, uint32(config&0xFF))
	if err := device.Write(ADS1x15_POINTER_CONFIG, buf[:2]); err != nil {
		panic(err)
	}

	time.Sleep(time.Millisecond)

	result0, err := device.Read(ADS1x15_POINTER_CONVERSION, 2)
	if err != nil {
		panic(err)
	}
	result1, err := device.Read(ADS1x15_POINTER_CONVERSION, 2)
	if err != nil {
		panic(err)
	}

	return conversionValue(int(binary.LittleEndian.Uint16(result0)), int(binary.LittleEndian.Uint16(result1)))
}

func conversionValue(low, high int) int {
	value := ((high & 0xFF) << 8) | (low & 0xFF)
	if value&0x8000 != 0 {
		value -= 1 << 16
	}
	return value
}

func main() {
	m, e := hwio.GetModule("i2c")
	if e != nil {
		fmt.Printf("could not get i2c module: %s\n", e)
		return
	}
	i2c := m.(hwio.I2CModule)

	// Uncomment on Raspberry pi, which doesn't automatically enable i2c bus. BeagleBone does,
	// as the default device tree enables it.

	i2c.Enable()
	defer i2c.Disable()

	device := i2c.GetDevice(0x4a)

	for {
		for i := 0; i < 4; i++ {
			log.Println(i, float32(readADC(device, i, 2/3, 128))/float32(1<<15)*6.144)
		}

		time.Sleep(time.Second)
	}
}
