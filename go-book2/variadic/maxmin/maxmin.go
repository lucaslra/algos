package maxmin

import (
	"errors"
	"sort"
)

func Max(n ...int) (int, error) {
	if len(n) == 0 {
		return 0, errors.New("at least 1 input value required")
	}

	sort.Ints(n)

	return n[len(n)-1], nil
}

func Min(n ...int) (int, error) {
	if len(n) == 0 {
		return 0, errors.New("at least 1 input value required")
	}

	sort.Ints(n)

	return n[0], nil
}

func Max1(n int, ns ...int) int {
	ns = append(ns, n)
	sort.Ints(ns)

	return ns[len(ns)-1]
}

func Min1(n int, ns ...int) int {
	ns = append(ns, n)
	sort.Ints(ns)

	return ns[0]
}
