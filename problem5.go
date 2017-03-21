package main

import "fmt"

func solveProblem5() {
	for i := 1; ; i++ {
		if divisibleUpTo(i, 20) {
			fmt.Printf("The small number divisible by all of the numbers from 1 to 20 is: %d", i)
			break
		}
	}
}

func divisibleUpTo(candidate int, max int) bool {
	for divisor := 2; divisor <= max; divisor++ {
		if candidate%divisor > 0 {
			return false
		}
	}
	return true
}
