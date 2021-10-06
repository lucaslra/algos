package quicksort

import "fmt"

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}

	return arr
}

func RunDemo() {
	fmt.Println("------------------------------\nRunning Quick sort demo...")
	dataSlice := []int{31, 342, 21, 312, 5, 20, 40, 2, 54, 2, 94}
	fmt.Println(dataSlice)
	quickSort(dataSlice, 0, len(dataSlice)-1)
	fmt.Println(dataSlice)
	fmt.Print("Quick sort demo completed\n------------------------------\n")
}
