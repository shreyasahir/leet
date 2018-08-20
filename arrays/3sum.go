package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	//found := false
	n := len(nums)
	for i := 0; i < n-2; i++ {
		l := i + 1
		r := n - 1
		x := nums[i]
		for l < r {
			if x+nums[l]+nums[r] == 0 {
				//fmt.Println(x, nums[l], nums[r])
				if uniqueSlice(result, x, nums[l], nums[r]) {
					result = append(result, []int{x, nums[l], nums[r]})
					//fmt.Println(x, nums[l], nums[r])
				}
				l = l + 1
				r = r - 1
				//found = true
			} else if x+nums[l]+nums[r] > 0 {
				r = r - 1
			} else {
				l = l + 1
			}
		}
	}
	return result
}

func uniqueSlice(result [][]int, x, y, z int) bool {
	var appendedSlice map[int]bool
	for _, v := range result {
		appendedSlice = nil
		appendedSlice = make(map[int]bool)

		for _, val := range v {
			appendedSlice[val] = true
		}
		//fmt.Println("appendedSlice", appendedSlice)
		_, ok := appendedSlice[x]
		_, ok1 := appendedSlice[y]
		_, ok2 := appendedSlice[z]
		//fmt.Println("ok,ok1,ok2", ok, ok1, ok2)

		if ok && ok1 && ok2 {
			if x == y && y == z {
				return true
			}
			return false
		}

	}
	return true
}
