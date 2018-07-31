package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var res [][]int
	var l, r int
	n := len(nums)
	sort.Ints(nums)
	//fmt.Println("nums after sort", nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r = i+1, n-1

		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s < 0 {
				l = l + 1
			} else if s > 0 {
				r = r - 1
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l = l + 1
				}
				for l < r && nums[r] == nums[r-1] {
					r = r - 1
				}
				l++
				r--
			}
		}
	}
	return res
}
