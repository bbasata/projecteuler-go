package main

import "testing"

type Problem1Test struct {
	limit       int
	expectedSum int
}

var tests = []Problem1Test{
	{0, 0},
	{1, 0},
	{3, 0},
	{4, 3},
	{5, 3},
	{6, 8},
	{7, 14},
	{8, 14},
	{9, 14},
	{10, 23},
}

func TestWithLimit10(t *testing.T) {
	for _, test := range tests {
		actualSum := sumMultiplesOf3Or5Below(test.limit)
		if actualSum != test.expectedSum {
			t.Errorf(`sumMultiplesOf3Or5Below(%d) = %d, want: %d`, test.limit, actualSum, test.expectedSum)
		}
	}
}
