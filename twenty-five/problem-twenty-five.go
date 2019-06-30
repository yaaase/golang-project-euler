package twentyfive

import (
	"math/big"
)

// ProblemTwentyFive ...
func ProblemTwentyFive() int {
	index := 2
	a, b := big.NewInt(1), big.NewInt(1)
	for {
		a, b = b, a.Add(a, b)
		index++
		if len(b.String()) > 999 {
			break
		}
	}
	return index
}
