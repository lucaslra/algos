package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func init() {
	fmt.Println("PACKAGE STARTED 1", AbsoluteZeroC)
}
func init() {
	fmt.Println("PACKAGE STARTED 2")
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	ZeroK         Kelvin  = Kelvin(AbsoluteZeroC)
	FreezingK     Kelvin  = 273.15
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
func KToC(k Kelvin) Celsius     { return Celsius(k - 273.15) }
func KToF(k Kelvin) Fahrenheit  { return Fahrenheit((k-273.15)*9/5 + 32) }
func FToK(f Fahrenheit) Kelvin  { return Kelvin((f-32)*5/9 + 273.15) }

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
