package main

import "fmt"

func solveProblem2() {
	limit := 4000000
	fmt.Printf("Sum of even-valued Fibonacci terms up to %d is %d\n", limit, sumOfEvenValuedFibTermsUpTo(limit))
}

func sumOfEvenValuedFibTermsUpTo(limit int) int {
	ch := make(chan int)
	acc := 0

	go fibGenerator(ch)

	for {
		nextFib := <-ch
		if nextFib > limit {
			break
		}
		if nextFib%2 == 0 {
			acc += nextFib
		}
	}

	return acc
}

func fibGenerator(ch chan int) {
	a, b := 1, 1
	for {
		a, b = b, a+b
		ch <- a
	}
}
