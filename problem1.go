package main

import "fmt"

func solveProblem1() {
	limit := 1000
	fmt.Printf("The sum of all the multiples of 3 or 5 below %d is: %d\n", limit, sumMultiplesOf3Or5Below(limit))
}

func sumMultiplesOf3Or5Below(limit int) int {
	acc := 0
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			acc += i
		}
	}
	return acc
}
