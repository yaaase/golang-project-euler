package five

// SmallestMultiple ...
func SmallestMultiple(upperBound int) int {
	for i := 2520; ; i += upperBound {
		if isDivisibleByAll(i, upperBound) {
			return i
		}
	}
}

func isDivisibleByAll(n, m int) bool {
	for j := m; j > 1; j-- {
		if n%j != 0 {
			return false
		}
	}
	return true
}
