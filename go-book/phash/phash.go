package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

var hash = flag.Int("f", 256, "Valid Values: 256, 384, 512")

func main() {
	flag.Parse()

	value := flag.Args()

	switch *hash {
	case 256:
		fmt.Printf("%v = %x", value, sha256.Sum256([]byte(strings.Join(value, " "))))
	case 384:
		fmt.Printf("%v = %x", value, sha512.Sum384([]byte(strings.Join(value, " "))))
	case 512:
		fmt.Printf("%v = %x", value, sha512.Sum512([]byte(strings.Join(value, " "))))
	default:
		fmt.Println("Invalid Format. Valid formats are: 256, 384, 512")
	}
}
