package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC  Celsius = -273.15
	FreesingC      Celsius = 0
	BoilingC       Celsius = 100
	AblsoluteZeroK Kelvin  = 0
	FreesingK      Kelvin  = 273.15
	BoilingK       Kelvin  = 373.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g °F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g °K", k)
}
