package three

import (
	"math"
)

// LargestPrimeFactorOf ...
func LargestPrimeFactorOf(n int) int {
	largestPrimeFactor := 2
	for i := int(float64(n) / 2); i > 1; i-- {
		if n%i == 0 && IsPrime(i) {
			largestPrimeFactor = i
			break
		}
	}
	return largestPrimeFactor
}

// IsPrime ...
func IsPrime(candidate int) bool {
	for i := 2; i <= int(math.Floor(float64(candidate)/2)); i++ {
		if candidate%i == 0 {
			return false
		}
	}
	return true
}
