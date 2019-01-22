// Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

// Your algorithm should run in O(n) complexity.

// Example:

// Input: [100, 4, 200, 1, 3, 2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

// class Solution:
//     def longestConsecutive(self, nums):
//         if not nums:
//             return 0

//         nums.sort()

//         longest_streak = 1
//         current_streak = 1

//         for i in range(1, len(nums)):
//             if nums[i] != nums[i-1]:
//                 if nums[i] == nums[i-1]+1:
//                     current_streak += 1
//                 else:
//                     longest_streak = max(longest_streak, current_streak)
//                     current_streak = 1

//         return max(longest_streak, current_streak)

package main

import "sort"

func main() {

}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	longestStreak := 1
	currentStreak := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				currentStreak += 1
			} else {
				longestStreak = max(longestStreak, currentStreak)
				currentStreak = 1
			}
		}
	}
	return max(longestStreak, currentStreak)
}
