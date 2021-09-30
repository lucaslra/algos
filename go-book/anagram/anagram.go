package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(areAnagram("Clint Eastwood", "Old West Action"))
	fmt.Println(areAnagram("abraxas", "xabaras"))
	fmt.Println(areAnagram("parliament", "partial men"))
	fmt.Println(areAnagram("abc", "zyx"))
	fmt.Println(areAnagram("Bug Cherish Lips", "Schuberg Philis"))
}

func areAnagram(w1, w2 string) bool {
	w1Clean := strings.ToLower(strings.ReplaceAll(w1, " ", ""))
	w2Clean := strings.ToLower(strings.ReplaceAll(w2, " ", ""))

	if len(w1Clean) != len(w2Clean) {
		return false
	}

	w1Map := mapWord(w1Clean)
	w2Map := mapWord(w2Clean)

	return reflect.DeepEqual(w1Map, w2Map)
}

func mapWord(s string) map[rune]int {
	mappedWord := make(map[rune]int)
	for _, v := range s {
		mappedWord[v]++
	}

	return mappedWord
}
