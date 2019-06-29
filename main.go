package main

import (
	"fmt"
	"project-euler/five"
	"project-euler/four"
	"project-euler/one"

	"project-euler/two"
)

func main() {
	fmt.Printf("Problem one: %v\n", one.ProblemOne())
	fmt.Printf("Problem two: %v\n", two.FibsBelow(4000000))
	//fmt.Printf("Problem three: %v\n", three.LargestPrimeFactorOf(600851475143))
	fmt.Printf("Problem four: %v\n", four.LargestPalindromicProduct())
	fmt.Printf("Problem five: %v\n", five.SmallestMultiple(20))
}
