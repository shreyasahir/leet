package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	input := []int{2, 7, 1, 8, 2, 8, 4, 5, 9, 0, 4, 5}
	fmt.Println("input", input)
	fmt.Println(DynamicProgrammingSingleSellProfit(input))
}

func bruteForceSingleSellProfit(arr []int) int {
	bestProfit := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			bestProfit = max(bestProfit, arr[j]-arr[i])
			fmt.Println("bestProfit", bestProfit)

		}
	}
	return bestProfit
}

func DynamicProgrammingSingleSellProfit(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	profit := 0
	cheapest := arr[0]
	for i := 1; i < len(arr); i++ {
		cheapest = min(cheapest, arr[i])
		profit = max(profit, arr[i]-cheapest)
	}
	return profit
}
