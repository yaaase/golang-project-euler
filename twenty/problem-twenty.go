package twenty

import (
	"math/big"
	"strconv"
	"strings"
)

// ProblemTwenty ...
func ProblemTwenty() int {
	factorial := Factorial(100)
	return sumOfDigits(factorial.String())
}

// Factorial ...
func Factorial(n int) big.Int {
	sum := big.NewInt(1)
	for i := 2; i <= n; i++ {
		sum.Mul(sum, big.NewInt(int64(i)))
	}
	return *sum
}

func sumOfDigits(x string) int {
	sum := 0
	digits := strings.Split(x, "")
	for _, digit := range digits {
		dig, err := strconv.Atoi(digit)
		if err == nil {
			sum += dig
		}
	}
	return sum
}
