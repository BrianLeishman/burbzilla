package main

import (
	"fmt"
	"strings"
)

type config struct {
	Boards []board `yaml:"boards"`
}

type boardType int

const (
	boardTypeADS1115 boardType = iota + 1
)

func (t boardType) String() string {
	switch t {
	case boardTypeADS1115:
		return "ADS1115"
	}
	panic(fmt.Errorf(`unhandled string value for board type "%d"`, t))
}

type board struct {
	Address int
	Type    boardType
	Sensors map[string]sensor
}

func (brd *board) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp struct {
		Address int               `yaml:"address"`
		Type    string            `yaml:"type"`
		Sensors map[string]sensor `yaml:"sensors"`
	}
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	brd.Address = tmp.Address
	switch strings.ToLower(tmp.Type) {
	case "ads1115":
		brd.Type = boardTypeADS1115
	default:
		return fmt.Errorf(`board type of "%s" is not supported`, tmp.Type)
	}
	brd.Sensors = tmp.Sensors

	return nil
}

type sensorType int

const (
	sensorTypeVolts sensorType = iota + 1
)

func (t sensorType) String() string {
	switch t {
	case sensorTypeVolts:
		return "Volts"
	}
	panic(fmt.Errorf(`unhandled string value for sensor type "%d"`, t))
}

type sensor struct {
	Channel int
	Type    sensorType
}

func (snsr *sensor) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var tmp struct {
		Channel int    `yaml:"channel"`
		Type    string `yaml:"type"`
	}
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	snsr.Channel = tmp.Channel
	switch strings.ToLower(tmp.Type) {
	case "volts":
		snsr.Type = sensorTypeVolts
	default:
		return fmt.Errorf(`sensor type of "%s" is not supported`, tmp.Type)
	}

	return nil
}
