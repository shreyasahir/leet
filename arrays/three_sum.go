// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

// Note:

// The solution set must not contain duplicate triplets.

// Example:

// Given array nums = [-1, 0, 1, 2, -1, -4],

// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	//threeSumBrute(arr)
	optimal(arr)
}

func optimal(arr []int) {
	sort.Ints(arr)
	var result [][]int
	fmt.Println("sorted arr", arr)
	n := len(arr)

	for i := 0; i < n; i++ {
		if i != 0 && arr[i] == arr[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			if arr[j]+arr[i]+arr[k] == 0 {
				result = append(result, []int{arr[i], arr[j], arr[k]})
				j++
				for j < k && arr[j] == arr[j-1] {
					j++
				}
			} else if arr[j]+arr[i]+arr[k] < 0 {
				j++
			} else {
				k--
			}

		}
	}

	fmt.Println("result", result)
}

func threeSumBrute(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i]+arr[j]+arr[k] == 0 {
					fmt.Println("true")
					fmt.Println("i,j,k", arr[i], arr[j], arr[k])
				}
			}
		}
	}
}
