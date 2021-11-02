package main

func main() {
	defer causePanic()

	defer (func() {
		defer (func() {
			defer causeAnotherPanic()
			panic("E: Fatal Error")
		})()
		panic("C: Error")
	})()
	panic("A")
}

func causePanic() {
	defer (func() {
		panic("D: Fatal Err")
	})()
	panic("B: Err")
}

func causeAnotherPanic() {
	panic("F: Fatal Err")
}
