package main

import "testing"

func TestProblem4(t *testing.T) {
	desiredMaxProduct := 9009
	if actualMaxProduct := maxPalindromicProductOfIntegersBetween(10, 99); actualMaxProduct != desiredMaxProduct {
		t.Errorf("maxPalindromicProductOfIntegersBetween(10, 99) == %d; want %d", actualMaxProduct, desiredMaxProduct)
	}
}
