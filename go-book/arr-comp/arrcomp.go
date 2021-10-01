package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("Lucas"))
	c2 := sha256.Sum256([]byte("Nirvana"))
	c3 := sha256.Sum256([]byte("Lucas"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println()
	fmt.Printf("%x\n%x\n%t\n%T\n", c3, c1, c1 == c3, c3)
}
