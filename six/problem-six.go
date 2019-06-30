package six

// ProblemSix ...
func ProblemSix() int {
	return SquareOfSums(100) - SumOfSquares(100)
}

// SumOfSquares ...
func SumOfSquares(limit int) int {
	sum := 0
	for i := 0; i <= limit; i++ {
		sum += i * i
	}
	return sum
}

// SquareOfSums ...
func SquareOfSums(limit int) int {
	sum := 0
	for i := 0; i <= limit; i++ {
		sum += i
	}
	return sum * sum
}
