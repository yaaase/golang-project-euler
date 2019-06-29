package two

// FibsBelow ...
func FibsBelow(limit int) int {
	a, b := 1, 1
	sum := 0

	for b < limit {
		fib := a + b
		if fib%2 == 0 {
			sum += fib
		}
		a, b = b, a+b
	}

	return sum
}
