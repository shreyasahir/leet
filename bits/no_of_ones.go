package main

import (
	"fmt"
	//"strconv"
)

func main() {
	var input uint32
	input = 11
	fmt.Println(ones(input))
}

func ones(input uint32) int {

	//n := int64(input)
	var count int
	// bin := fmt.Sprintf("%032b", input)
	// n, _ := strconv.Atoi(bin)
	// fmt.Println(n)
	n := input
	for n > 0 {
		fmt.Println("n", n)
		n = n & (n - 1)
		count++
		fmt.Println("count", count)
	}
	return count
}
