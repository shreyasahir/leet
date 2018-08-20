package main

import (
	"fmt"
)

func main() {
	inp := []int{1, 1}
	fmt.Println(maxArea(inp))
}

func maxArea(height []int) int {

	var l, r, maxArea int
	l = 0
	r = len(height) - 1

	for l < r {
		maxArea = max(maxArea, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
