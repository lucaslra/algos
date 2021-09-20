package code_wars

import (
	"fmt"
)

func InAscOrder(numbers []int) bool {
	fmt.Printf("Running Code Kata 56b7f2f3f18876033f000307\n")

	if len(numbers) == 1 {
		return true
	}
	result := true
	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			continue
		}

		if numbers[i] < numbers[i-1] {
			result = false
			break
		}
	}

	fmt.Printf("Code Kata 56b7f2f3f18876033f000307 completed. Result %v\n", result)

	return result
}
