package main

// Register and other configuration values:
const (
	ADS1x15_DEFAULT_ADDRESS        = 0x48
	ADS1x15_POINTER_CONVERSION     = 0x00
	ADS1x15_POINTER_CONFIG         = 0x01
	ADS1x15_POINTER_LOW_THRESHOLD  = 0x02
	ADS1x15_POINTER_HIGH_THRESHOLD = 0x03
	ADS1x15_CONFIG_OS_SINGLE       = 0x8000
	ADS1x15_CONFIG_MUX_OFFSET      = 12
)

// Maping of gain values to config register values.
var ADS1x15_CONFIG_GAIN = map[float64]int{
	2 / 3: 0x0000,
	1:     0x0200,
	2:     0x0400,
	4:     0x0600,
	8:     0x0800,
	16:    0x0A00,
}

const ADS1x15_CONFIG_MODE_CONTINUOUS = 0x0000
const ADS1x15_CONFIG_MODE_SINGLE = 0x0100

// Mapping of data/sample rate to config register values for ADS1015 (faster).
var ADS1015_CONFIG_DR = map[int]int{
	128:  0x0000,
	250:  0x0020,
	490:  0x0040,
	920:  0x0060,
	1600: 0x0080,
	2400: 0x00A0,
	3300: 0x00C0,
}

// Mapping of data/sample rate to config register values for ADS1115 (slower).
var ADS1115_CONFIG_DR = map[int]int{
	8:   0x0000,
	16:  0x0020,
	32:  0x0040,
	64:  0x0060,
	128: 0x0080,
	250: 0x00A0,
	475: 0x00C0,
	860: 0x00E0,
}

const ADS1x15_CONFIG_COMP_WINDOW = 0x0010
const ADS1x15_CONFIG_COMP_ACTIVE_HIGH = 0x0008
const ADS1x15_CONFIG_COMP_LATCHING = 0x0004

var ADS1x15_CONFIG_COMP_QUE = map[int]int{
	1: 0x0000,
	2: 0x0001,
	4: 0x0002,
}

const ADS1x15_CONFIG_COMP_QUE_DISABLE = 0x0003
