package binarysearch

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(array []int, value, offset int) int {
	mid := len(array) / 2
	if value < array[mid] {
		return binarySearch(array[0:mid], value, offset)
	} else if value > array[mid] {
		return binarySearch(array[(mid+1):], value, offset+mid+1)
	} else {
		return offset + mid
	}
}

func RunDemo() {
	fmt.Println("------------------------------\nRunning Binary search demo...")
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(50)[:20]
	value := arr[rand.Intn(len(arr))]
	sort.Ints(arr)
	result := binarySearch(arr, value, 0)
	fmt.Printf("Position of %v in %v is: %v\n", value, arr, result)
	fmt.Print("Binary search demo completed\n------------------------------\n")
}
