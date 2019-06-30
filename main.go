package main

import (
	"fmt"
	"project-euler/five"
	"project-euler/four"
	"project-euler/one"
	"project-euler/seven"
	"project-euler/six"
	"project-euler/sixteen"
	"project-euler/three"
	"project-euler/twenty"
	"project-euler/two"
)

func main() {
	fmt.Printf("Problem one: %v\n", one.ProblemOne())
	fmt.Printf("Problem two: %v\n", two.FibsBelow(4000000))
	fmt.Printf("Problem three: %v\n", three.LargestPrimeFactorOf(600851475143))
	fmt.Printf("Problem four: %v\n", four.LargestPalindromicProduct())
	fmt.Printf("Problem five: %v\n", five.SmallestMultiple(20))
	fmt.Printf("Problem six: %v\n", six.ProblemSix())
	fmt.Printf("Problem seven: %v\n", seven.Sieve())
	fmt.Printf("Problem sixteen: %v\n", sixteen.ProblemSixteen())
	fmt.Printf("Problem twenty: %v\n", twenty.ProblemTwenty())
}
