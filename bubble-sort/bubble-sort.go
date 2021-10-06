package bubblesort

import "fmt"

func bubbleSort(s []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(s); i++ {
			if s[i-1] > s[i] {
				s[i], s[i-1] = s[i-1], s[i]
				swapped = true
			}
		}
	}

	return s
}

func RunDemo() {
	fmt.Println("------------------------------\nRunning Bubble sort demo...")
	dataSlice := []int{31, 342, 21, 312, 5, 20, 40, 2, 54, 2, 94}
	fmt.Println(dataSlice)
	bubbleSort(dataSlice)
	fmt.Println(dataSlice)
	fmt.Print("Bubble sort demo completed\n------------------------------\n")
}
