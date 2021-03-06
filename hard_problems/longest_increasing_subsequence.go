// Given an unsorted array of integers, find the number of longest increasing subsequence.

// Example 1:
// Input: [1,3,5,4,7]
// Output: 2
// Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].
// Example 2:
// Input: [2,2,2,2,2]
// Output: 5
// Explanation: The length of longest continuous increasing subsequence is 1, and there are 5 subsequences' length is 1, so output 5.

// https://leetcode.com/problems/number-of-longest-increasing-subsequence/

package main

import (
	"fmt"
)

func main() {

	arr := []int{2, 2, 2, 2, 2}
	findNumberOfLIS(arr)
}

// def findNumberOfLIS(self, nums):
// N = len(nums)
// if N <= 1: return N
// lengths = [0] * N #lengths[i] = longest ending in nums[i]
// counts = [1] * N #count[i] = number of longest ending in nums[i]

// for j, num in enumerate(nums):
// 	for i in xrange(j):
// 		if nums[i] < nums[j]:
// 			if lengths[i] >= lengths[j]:
// 				lengths[j] = 1 + lengths[i]
// 				counts[j] = counts[i]
// 			elif lengths[i] + 1 == lengths[j]:
// 				counts[j] += counts[i]

// longest = max(lengths)
// return sum(c for i, c in enumerate(counts) if lengths[i] == longest)

func findNumberOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	longest := make([]int, len(nums))
	counts := make([]int, len(nums))

	for j := 0; j < len(nums); j++ {
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				if longest[i] >= longest[j] {
					longest[j] = 1 + longest[i]
					counts[j] = counts[i]
				} else if longest[i]+1 == longest[i] {
					counts[j] += counts[i]
				}
			}
		}
	}

	// longest = math.Max(longest)
	fmt.Println("longest", longest)
	return -1
}
