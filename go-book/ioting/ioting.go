package main

import (
	"fmt"
)

const (
	KB float64 = 1000e0
	MB float64 = 1000e1
	GB float64 = 1000e2
	TB float64 = 1000e3
	PB float64 = 1000e4
	EB float64 = 1000e5
	ZB float64 = 1000e6
	YB float64 = 1000e7
)

func main() {
	fmt.Printf("%.0f\n%.0f\n%.0f\n%.0f\n%.0f\n%.0f\n%.0f\n%.0f", KB, MB, GB, TB, PB, EB, ZB, YB)
}
