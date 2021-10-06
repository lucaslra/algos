package mergesort2

import (
	"fmt"
)

func merge(a, b []int) []int {
	final := []int{}
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

func mergeSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	first := mergeSort(s[:len(s)/2])
	second := mergeSort(s[len(s)/2:])

	return merge(first, second)
}

func RunDemo() {
	fmt.Println("------------------------------\nRunning Merge sort 2 demo...")
	unsortedSlice := []int{31, 342, 21, 312, 5, 20, 40, 2, 54, 2, 94}
	result := mergeSort(unsortedSlice)
	fmt.Printf("Unsorted array: %v\n", unsortedSlice)

	fmt.Printf("Result array: %v\n", result)
	fmt.Print("Merge sort 2 demo completed\n------------------------------\n")
}
