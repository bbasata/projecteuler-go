package main

import "testing"

type Problem2Test struct {
	limit          int
	expectedFibSum int
}

var problem2Tests = []Problem2Test{
	{0, 0},
	{1, 0},
	{2, 2},
	{8, 10},
	{34, 44},
}

func TestProblem2(t *testing.T) {
	for _, test := range problem2Tests {
		if actualFibSum := sumOfEvenValuedFibTermsUpTo(test.limit); actualFibSum != test.expectedFibSum {
			t.Errorf(`sumOfEvenValuedFibTermsUpTo(%d) = %d, want: %d`, test.limit, actualFibSum, test.expectedFibSum)
		}
	}
}
