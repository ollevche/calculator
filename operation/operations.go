package operation

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func Multiply(a, b int) int {
	return a * b
}
