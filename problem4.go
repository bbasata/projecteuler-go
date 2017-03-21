package main

import "fmt"

func solveProblem4() {
	maxPalindromicProduct := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := i * j
			productString := fmt.Sprintf("%d", product)

			if productString == reverse(productString) && product > maxPalindromicProduct {
				maxPalindromicProduct = product
			}
		}
	}

	fmt.Printf("Largest palindromic product of 3-digit numbers is %d\n", maxPalindromicProduct)
}

// https://github.com/golang/example/blob/master/stringutil/reverse.go
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
