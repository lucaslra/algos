package main

import (
	"fmt"

	binarysearch "github.com/lucaslra/go-algos/binary-search"
	linkedlist "github.com/lucaslra/go-algos/linked-list"
	mergesort "github.com/lucaslra/go-algos/merge-sort"
)

func main() {
	fmt.Println("Running Algos...")
	binarysearch.RunDemo()
	mergesort.RunDemo()
	linkedlist.RunDemo()
	fmt.Println("Algos completed")
}
