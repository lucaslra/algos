package main

import "fmt"

func main() {
	var i interface{} = 1
	var j interface{} = "test"

	var i1 = i.(int)
	var j1 = j.(string)

	fmt.Printf("type of i == %T | type of i1 == %T\n", i, i1)
	fmt.Printf("type of j == %T | type of j1 == %T\n", j, j1)
}
