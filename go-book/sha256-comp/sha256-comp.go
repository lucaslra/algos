package main

import (
	"crypto/sha256"
	"fmt"
)

func diffSize(a1, a2 [32]uint8) int {
	size := 0
	for i := 0; i < 32; i++ {
		if a1[i] != a2[i] {
			size++
		}
	}

	fmt.Printf("%v\n%v\n", a1, a2)
	return size
}

func main() {
	c1 := sha256.Sum256([]byte("X"))
	c2 := sha256.Sum256([]byte("x"))

	size := diffSize(c1, c2)

	fmt.Printf("%x diff from %x = %d bits", c1, c2, size)
}
