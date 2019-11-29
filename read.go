package main

import (
	"time"

	cmap "github.com/orcaman/concurrent-map"

	"github.com/MichaelS11/go-ads"
)

var sensors = cmap.New()

func read() {
	for {
		for _, brd := range conf.Boards {
			for name, snsr := range brd.Sensors {
				var ads1 *ads.ADS
				ads1, err := ads.NewADS("I2C1", uint16(brd.Address), "")
				check(err)

				var gain ads.ConfigGain
				var dataRate ads.ConfigDataRate
				switch brd.Type {
				case boardTypeADS1115:
					gain = ads.ConfigGain2_3
					dataRate = ads.ConfigDataRate860
				}
				ads1.SetConfigGain(gain)
				ads1.SetConfigDataRate(dataRate)

				var mux ads.ConfigInputMultiplexer
				switch snsr.Channel {
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
					check(err)
				}

				err = ads1.Close()
				check(err)

				sensors.Set(name, result)
			}
		}
		time.Sleep(time.Millisecond)
	}
}
