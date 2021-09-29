package main

import (
	"fmt"
	"pcount/popcount"
	"time"
)

func main() {
	start := time.Now()
	var x uint64
	for i := 0; i <= 100000000000; i = i + 100 {
		x = uint64(i)
		popcount.PopCount(x)
	}
	fmt.Printf("Single Expression took %.2fs", time.Since(start).Seconds())
}
