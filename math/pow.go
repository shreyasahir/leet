package main

import (
	"fmt"
)

func main() {
	x := 2.00000
	n := -2
	fmt.Println(myPow(x, n))
}

func myPow(x float64, n int) float64 {
	var signBit bool
	if n < 0 {
		signBit = true
		n = -1 * n
	}
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		x *= x
		n = n / 2
		if signBit {
			return 1.0 / myPow(x, n)
		}
		return myPow(x, n)
	}
	// x *= x
	// n = n - 1
	if signBit {
		return 1.0 / float64(x*myPow(x, n-1))
	}
	return float64(x * myPow(x, n-1))
}
