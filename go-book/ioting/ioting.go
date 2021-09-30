package main

import (
	"fmt"
)

const (
	KB = 1e3
	MB = 1e6
	GB = 1e9
	TB = 1e12
	PB = 1e15
	EB = 1e18
	ZB = 1e21
	YB = 1e24
)

func main() {
	fmt.Printf("%.0f\n%.0f\n%.0f\n%.0f\n%.0f\n%.0f\n%.0f\n%.0f", KB, MB, GB, TB, PB, EB, ZB, YB)
}
