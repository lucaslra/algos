package code_wars

import "fmt"

func FindUniq(arr []float32) float32 {
	fmt.Printf("Running Code Kata 585d7d5adb20cf33cb000235\n")
	counter := map[float32]int{}
	var result float32

	for _, v := range arr {
		_, exists := counter[v]

		if exists {
			counter[v]++
		} else {
			counter[v] = 1
		}
	}

	for k, v := range counter {
		if v == 1 {
			result = k
			break
		}
	}

	fmt.Printf("Code Kata 585d7d5adb20cf33cb000235 completed. Result %v\n", result)
	return result
}
