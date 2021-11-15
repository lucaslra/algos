package main

import (
	"io"
	"log"
	"os"
	"strings"

	"azurecoder.dev/algos-book/limitreader/limitreader"
)

func main() {
	data := "This is my string"
	reader := strings.NewReader(data)

	limitedReader := limitreader.LimitReader(reader, 8)

	if _, err := io.Copy(os.Stdout, limitedReader); err != nil {
		log.Fatal(err)
	}
}
