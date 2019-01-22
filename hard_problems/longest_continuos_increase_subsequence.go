// Given an unsorted array of integers, find the length of longest continuous increasing subsequence (subarray).

// Example 1:
// Input: [1,3,5,4,7]
// Output: 3
// Explanation: The longest continuous increasing subsequence is [1,3,5], its length is 3.
// Even though [1,3,5,7] is also an increasing subsequence, it's not a continuous one where 5 and 7 are separated by 4.
// Example 2:
// Input: [2,2,2,2,2]
// Output: 1
// Explanation: The longest continuous increasing subsequence is [2], its length is 1.

// func findLengthOfLCIS(nums []int) int {

// }
// https://leetcode.com/problems/longest-continuous-increasing-subsequence/

package main

import "fmt"

func main() {

	arr := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(arr))
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func findLengthOfLCIS(nums []int) int {
	maxLen := 0
	i := 0

	for i < len(nums) {
		curr := 1
		for i+1 < len(nums) && nums[i] < nums[i+1] {
			curr += 1
			i++
		}
		if maxLen <= curr {

			maxLen = curr
		}
		i += 1
	}
	return maxLen
}
