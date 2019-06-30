package sixteen

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ProblemSixteen ...
func ProblemSixteen() int {
	candidate := math.Pow(2, 1000)

	return SumOfDigits(candidate)
}

// SumOfDigits ...
func SumOfDigits(x float64) int {
	digits := strings.Split(fmt.Sprintf("%f", x), "")
	sum := 0
	for _, digit := range digits {
		dig, err := strconv.Atoi(digit)
		if err == nil {
			sum += dig
		}
	}
	return sum
}
