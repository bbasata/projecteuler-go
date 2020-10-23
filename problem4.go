package main

import "fmt"

func solveProblem4() {
	answer := maxPalindromicProductOfIntegersBetween(100, 999)
	fmt.Printf("Largest palindromic product of 3-digit numbers is %d\n", answer)
}

func maxPalindromicProductOfIntegersBetween(min, max int) int {
	maxPalindromicProduct := 0

	for i := min; i <= max; i++ {
		for j := min; j <= max; j++ {
			product := i * j
			productString := fmt.Sprintf("%d", product)

			if productString == reverse(productString) && product > maxPalindromicProduct {
				maxPalindromicProduct = product
			}
		}
	}

	return maxPalindromicProduct
}

// https://github.com/golang/example/blob/master/stringutil/reverse.go
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
