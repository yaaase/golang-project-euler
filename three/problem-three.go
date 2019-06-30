package three

import (
	"math"
)

// LargestPrimeFactorOf ...
func LargestPrimeFactorOf(n int) int {
	var largest int
	for i := int(math.Sqrt(float64(n))); i >= 1; i-- {
		if n%i == 0 && IsPrime(i) {
			largest = i
			break
		}
	}
	return largest
}

// IsPrime ...
func IsPrime(candidate int) bool {
	if candidate <= 1 {
		return false
	}

	for i := int(math.Sqrt(float64(candidate))); i >= 1; i-- {
		if i == 1 {
			return true
		}
		if candidate%i == 0 {
			return false
		}
	}
	return true
}
