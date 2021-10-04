package main

import (
	"fmt"
)

func removeAdjacents(s *[]string) {
	for i := 0; i < len(*s); i++ {
		if i == 0 {
			continue
		}
		if (*s)[i] == (*s)[i-1] {
			*s = append((*s)[:i], (*s)[i+1:]...)
		}
	}
}

func main() {
	fmt.Println(string('A' + 13))
	strs := []string{"L", "N", "K", "J", "J", "N", "W", "W", "S"}

	removeAdjacents(&strs)

	fmt.Println(strs)
}
