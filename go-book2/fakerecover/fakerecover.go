package main

import "fmt"

func main() {
	fmt.Println(panics())
}

func panics() (panResult string) {
	defer func() {
		if p := recover(); p != nil {
			panResult = fmt.Sprintf("Don't Panic! %v", p)
		}
	}()

	panic("Panic!")
}
