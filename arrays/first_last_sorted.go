package main

import (
	"fmt"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	fmt.Println(searchRange(nums, target))
}

func bSearch(nums []int, target int, left bool) int {

	low := 0
	high := len(nums)

	for low < high {
		mid := (low + high) / 2
		if nums[mid] > target || (left && target == nums[mid]) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func searchRange(nums []int, target int) []int {

	leftIndex := bSearch(nums, target, true)

	if leftIndex == len(nums) || nums[leftIndex] != target {
		return []int{-1, -1}
	}

	return []int{leftIndex, bSearch(nums, target, false) - 1}
}
