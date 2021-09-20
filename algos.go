package main

import (
	"fmt"
	code_wars "github.com/lucaslra/go-algos/code-wars"

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

	code_wars.Multiple3And5(10)
	code_wars.RowSumOddNumbers(13)
	code_wars.InAscOrder([]int{1, 6, 10, 18, 2, 4, 20})
	code_wars.FindUniq([]float32{ 0, 0, 0.55, 0, 0  })
	fmt.Println("Algos completed")
}
