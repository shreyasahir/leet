package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{-4, -1, 0, 3, 10}
	fmt.Println("output", sortedSquares(arr))
}

func sortedSquares(A []int) []int {
	for k, v := range A {
		A[k] = v * v
	}
	sort.Ints(A)
	return A
}
