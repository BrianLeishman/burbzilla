package main

import (
	"fmt"
	"time"

	"github.com/MichaelS11/go-ads"
)

func main() {
	err := ads.HostInit()
	if err != nil {
		fmt.Println(err)
	}

	for {
		for i := 0; i < 4; i++ {
			// create new ads with wanted busName and address.
			var ads1 *ads.ADS
			ads1, err = ads.NewADS("I2C1", 0x4a, "")
			if err != nil {
				fmt.Println(err)
			}

			// example changing config gain (2/3 is default, so only an example)
			ads1.SetConfigGain(ads.ConfigGain2_3)
			ads1.SetConfigDataRate(ads.ConfigDataRate860)

			var mux ads.ConfigInputMultiplexer
			switch i {
			case 0:
				mux = ads.ConfigInputMultiplexerSingle0
			case 1:
				mux = ads.ConfigInputMultiplexerSingle1
			case 2:
				mux = ads.ConfigInputMultiplexerSingle2
			case 3:
				mux = ads.ConfigInputMultiplexerSingle3
			}
			ads1.SetConfigInputMultiplexer(mux)

			var result uint16
			result, err = ads1.Read()
			if err != nil {
				ads1.Close()
				fmt.Println(err)
			}

			// close ads bus
			err = ads1.Close()
			if err != nil {
				fmt.Println(err)
			}

			// print results
			fmt.Printf("% 12f", float64(result)/((1<<15)-1)*6.144)
		}
		time.Sleep(time.Millisecond)

		fmt.Println()
	}
}
