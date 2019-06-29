package five

// SmallestMultiple ...
func SmallestMultiple(upperBound int) int {
	for i := 2520; ; i += upperBound {
		if isDivisibleByAll(i, upperBound) {
			return i
		}
	}
}

func isDivisibleByAll(n int, m int) bool {
	for j := m; j > 0; j-- {
		if n%j != 0 {
			return false
		}
	}
	return true
}
