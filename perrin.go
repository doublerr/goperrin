package goperrin

import "fmt"

func perrin() func() int {
	a, b, c := 3, 0, 2

	return func() int {
		a, b, c = b, c, a+b
		return a
	}
}

func perrin_max(max int) func() int {
	a, b, c := 3, 0, 2

	return func() int {
		if a >= max {
			return a
		}

		a, b, c = b, c, a+b
		return a
	}
}

func perrin_reset(max int) func() int {
	a, b, c := 3, 0, 2

	return func() int {
		if a >= max {
			a, b, c = 3, 0, 2
		}

		a, b, c = b, c, a+b
		return a
	}
}
