package goperrin

func Perrin() func() int {
	a, b, c := 3, 0, 2

	return func() int {
		a, b, c = b, c, a+b
		return a
	}
}

func PerrinMax(max int) func() int {
	a, b, c := 3, 0, 2

	return func() int {
		if a >= max {
			return a
		}

		a, b, c = b, c, a+b
		return a
	}
}

func PerrinReset(max int) func() int {
	a, b, c := 3, 0, 2

	return func() int {
		if a >= max {
			a, b, c = 3, 0, 2
		}

		a, b, c = b, c, a+b
		return a
	}
}
