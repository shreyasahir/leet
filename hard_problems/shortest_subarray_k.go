// https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/

// Return the length of the shortest, non-empty, contiguous subarray of A with sum at least K.

// If there is no non-empty subarray with sum at least K, return -1.

// Example 1:

// Input: A = [1], K = 1
// Output: 1
// Example 2:

// Input: A = [1,2], K = 4
// Output: -1
// Example 3:

// Input: A = [2,-1,2], K = 3
// Output: 3

// Note:

// 1 <= A.length <= 50000
// -10 ^ 5 <= A[i] <= 10 ^ 5
// 1 <= K <= 10 ^ 9

// func shortestSubarray(A []int, K int) int {

// }

package main

import "fmt"

func main() {
	arr := []int{1, 2}
	k := 4
	fmt.Println(shortestSubarray(arr, k))
}

//brute force

func shortestSubarray(A []int, k int) int {
	n := len(A)
	minLength := n + 1
	for i := 0; i < n; i++ {
		currSum := A[i]
		if currSum >= k {
			return 1
		}

		for j := i + 1; j < n; j++ {
			currSum += A[j]
			if currSum >= k && j-i+1 < minLength {
				minLength = j - i + 1
			}
		}
	}

	if minLength > n {
		return -1
	}
	return minLength
}
