package isolation

func MakeFibonacciGen() func() int {
	f1, f2 := 0, 1

	return func() int {
		f1, f2 = f2, f1+f2

		return f1
	}
}
