package one

// ProblemOne ...
func ProblemOne() int {
	return sumOfMultiplesBelow(1000)
}

func sumOfMultiplesBelow(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
