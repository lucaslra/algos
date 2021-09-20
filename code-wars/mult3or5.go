package code_wars

import "fmt"

func Multiple3And5(number int) int {
	fmt.Printf("Running Code Kata 514b92a657cdc65150000006\n")

	result := 0
	for i := 0; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			result += i
		}
	}

	fmt.Printf("Code Kata 514b92a657cdc65150000006 completed. Result %v\n", result)

	return result
}
