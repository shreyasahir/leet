package main

import (
	"fmt"
)

func main() {
	input := 5
	fmt.Println(countBits(input))
}

func countBits(num int) []int {
	var output []int
	for i := 0; i <= num; i++ {
		output = append(output, ones(i))
	}
	return output
}

func ones(input int) int {

	//n := int64(input)
	var count int
	// bin := fmt.Sprintf("%032b", input)
	// n, _ := strconv.Atoi(bin)
	// fmt.Println(n)
	n := input
	for n > 0 {
		//fmt.Println("n", n)
		n = n & (n - 1)
		count++
		//fmt.Println("count", count)
	}
	return count
}
