package insertionsort

import (
	"fmt"
)

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func RunDemo() {
	fmt.Println("------------------------------\nRunning Insertion sort demo...")
	dataSlice := []int{31, 342, 21, 312, 5, 20, 40, 2, 54, 2, 94}
	fmt.Printf("Unsorted array: %v\n", dataSlice)
	insertionSort(dataSlice)
	fmt.Printf("Result array: %v\n", dataSlice)
	fmt.Print("Insertion sort demo completed\n------------------------------\n")
}
