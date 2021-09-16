package mergesort

import (
	"fmt"
	"math/rand"
	"time"
)

func shift(arr *[]int) int {
	r := (*arr)[0]
	*arr = (*arr)[1:]

	return r
}

func sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	midpoint := len(arr) / 2
	left := (arr)[0:midpoint]
	right := (arr)[midpoint:]

	return merge(sort(left), sort(right))
}

func merge(left, right []int) []int {
	var mergedArr []int

	for len(left) != 0 || len(right) != 0 {
		switch {
		case len(left) == 0:
			mergedArr = append(mergedArr, shift(&right))
		case len(right) == 0:
			mergedArr = append(mergedArr, shift(&left))
		case left[0] <= right[0]:
			mergedArr = append(mergedArr, shift(&left))
		case right[0] <= left[0]:
			mergedArr = append(mergedArr, shift(&right))
		}
	}

	return mergedArr
}

func RunDemo() {
	fmt.Println("------------------------------\nRunning Merge sort demo...")
	rand.Seed(time.Now().UnixNano())
	arr1 := rand.Perm(50)[:20]
	fmt.Printf("Original arrays: %v\n", arr1)
	result := sort(arr1)
	fmt.Printf("Result array: %v\n", result)
	fmt.Print("Merge sort demo completed\n------------------------------\n")
}
