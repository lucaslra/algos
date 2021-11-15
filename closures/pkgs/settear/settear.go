package settear

import "fmt"

func Setup() func() {
	fmt.Println("Setup")

	fmt.Println("Do Stuff")

	return func() {
		fmt.Println("Teardown")
	}
}
