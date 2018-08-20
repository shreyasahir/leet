package main

import (
	"fmt"
)

func main() {
	//arr := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	arr := []int{3, 0, 1}
	missingNumber(arr)
}

func missingNumber(nums []int) int {
	var result int
	result = len(nums)
	for i := 0; i < len(nums); i++ {
		result ^= nums[i] ^ i
	}
	fmt.Println(result)
	return result
}
