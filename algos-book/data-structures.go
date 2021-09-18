package algosbook

import (
	"container/list"
	"fmt"
)

func RunDataStructuresDemo() {
	runListDemo()
	runTupleDemo()
	runHeapDemo()
}

func runHeapDemo() {
	fmt.Println("------------------------------\nRunning heap demo...")
	executeHeapDemo()
	fmt.Print("Heap demo completed\n------------------------------\n")
}

func runTupleDemo() {
	fmt.Println("------------------------------\nRunning tuple demo...")
	powerSeries := func(n int) (int, int) {
		return n * n, n * n * n
	}

	var square, cube int
	square, cube = powerSeries(3)

	fmt.Println("Square", square, "Cube", cube)
	fmt.Print("Tuple demo completed\n------------------------------\n")
}

func runListDemo() {
	fmt.Println("------------------------------\nRunning List demo...")
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

	fmt.Print("List demo completed\n------------------------------\n")
}
