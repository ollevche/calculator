package operation

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

// Divide will panic if enounters division by zero.
func Divide(a, b float64) float64 {
	return a / b
}

func Multiply(a, b int) int {
	return a * b
}
