package main

import (
	"fmt"
	"os"

	"azurecoder.dev/go-book/variadic/maxmin"
	"azurecoder.dev/go-book/variadic/strjoin"
)

func main() {

	fmt.Printf("Max: %d\n", maxmin.Max1(39, 49, 2, 59, 25, 96, 9, 91))
	fmt.Printf("Min: %d\n", maxmin.Min1(39, 49, 2, 59, 25, 96, 9, 91))

	res, err := maxmin.Max(39, 49, 2, 59, 25, 96, 9, 91)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	fmt.Printf("Max: %d\n", res)

	res, err = maxmin.Min(39, 49, 2, 59, 25, 96, 9, 91)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	fmt.Printf("Min: %d\n", res)

	res, err = maxmin.Max()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	fmt.Printf("Max: %d\n", res)

	res, err = maxmin.Min()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	fmt.Printf("Min: %d\n", res)

	fmt.Printf("Variadic Join: %s", strjoin.Join(";", "Lucas", "Nirvana", "Kate", "Nala", "Winter", "Summer"))

	// test := maxmin.Max1() // Invalid!
}
