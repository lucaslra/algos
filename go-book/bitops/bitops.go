package main

import "fmt"

func main() {
	var x uint8 = 1<<1 | 1<<5 | 1<<3
	var y uint8 = 1<<1 | 1<<3 | 1<<4

	fmt.Printf("%08b\n", x)            // 00101010
	fmt.Printf("%08b\n", y)            // 00011010
	fmt.Printf("& op  : %08b\n", x&y)  // Intersection 00001010
	fmt.Printf("| op  : %08b\n", x|y)  // Union 00111010
	fmt.Printf("^ op  : %08b\n", x^y)  // Symetric Diff 00110000
	fmt.Printf("&^ op : %08b\n", x&^y) // Diff 00100000

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i)
		}
	}

	fmt.Printf("%08b\n", x<<1) // x shift left: 01010100
	fmt.Printf("%08b\n", x>>1) // x shift right: 00010101
}
