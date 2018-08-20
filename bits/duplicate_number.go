package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 1, 3, 4, 2}
	fmt.Println(findDuplicate(nums))
}

func findDuplicate(nums []int) int {
	tortoise := nums[0]
	hare := nums[0]
	for {
		fmt.Println("tortoise", tortoise)
		fmt.Println("hare", hare)
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]

		if hare == tortoise {
			break
		}
	}
	ptr1 := nums[0]
	ptr2 := tortoise
	fmt.Println("tortoisessss", tortoise)
	fmt.Println("haressss", hare)
	for ptr1 != ptr2 {
		ptr1 = nums[ptr1]
		ptr2 = nums[ptr2]
	}
	return ptr1
}
