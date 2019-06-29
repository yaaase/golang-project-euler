package four

import (
	"strconv"
	"strings"
)

// LargestPalindromicProduct ...
func LargestPalindromicProduct() int {
	largest := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j
			if largest > product {
				break
			}
			if IsPalindrome(product) {
				largest = product
			}
		}
	}
	return largest
}

// IsPalindrome ...
func IsPalindrome(n int) bool {
	chars := strings.Split(strconv.Itoa(n), "")

	for i, j := 0, len(chars)-1; j > i; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}
	return true
}
