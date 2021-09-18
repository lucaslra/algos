package main

import (
	"fmt"

	algosbook "github.com/lucaslra/go-algos/algos-book"
	binarysearch "github.com/lucaslra/go-algos/binary-search"
	linkedlist "github.com/lucaslra/go-algos/linked-list"
	mergesort "github.com/lucaslra/go-algos/merge-sort"
)

func main() {
	fmt.Println("Running Algos...")
	binarysearch.RunDemo()
	mergesort.RunDemo()
	linkedlist.RunDemo()
	algosbook.RunDataStructuresDemo()
	algosbook.RunAdapterDemo()
	fmt.Println("Algos completed")
}
