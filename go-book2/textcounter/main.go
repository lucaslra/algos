package main

import (
	"fmt"

	"azurecoder.dev/algos-book/textcounter/pkgs/counters"
)

func main() {
	bytesCounter := counters.NewBytesCounter()
	fmt.Fprint(&bytesCounter, "Hello World")

	fmt.Println(bytesCounter)

	wordsCounter := counters.NewWordsCounter()
	fmt.Fprint(&wordsCounter, "Hello World")
	fmt.Println(wordsCounter)

	linesCounter := counters.NewLinesCounter()
	fmt.Fprint(&linesCounter, "Hello\nWorld\nHello\nWorld")
	fmt.Println(linesCounter)
}
