package main

import (
	"fmt"
	tc "tempconv/tempconv"
)

func main() {
	fever := tc.Celsius(37)

	f := tc.CToF(fever)
	c := tc.FToC(f)
	k := tc.CToK(c)

	zeroC := tc.Celsius(0)
	zeroF := tc.Fahrenheit(zeroC)
	zeroK := tc.Kelvin(zeroF)

	fmt.Printf("C Fever: %v\nF Fever: %v\nK Fever: %v\n", c, f, k)

	fmt.Println(zeroK)
}
