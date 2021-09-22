package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))

	for i, v := range os.Args {
		fmt.Printf("%d = %v\n", i, v)
	}
}
