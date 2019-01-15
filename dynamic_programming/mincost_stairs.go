package main

import (
	"fmt"
)

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	//cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	var dp1, dp2 int
	n := len(cost)
	for i := 0; i < n; i++ {
		dp0 := cost[i] + min(dp1, dp2)
		dp2 = dp1
		dp1 = dp0
	}
	return min(dp1, dp2)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
