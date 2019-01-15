package main

import (
	"fmt"
)

func main() {
	coins := []int{2}
	n := 3
	fmt.Println(coinChange(coins, n))
}

func coinChange(coins []int, amount int) int {
	return dpMinCoin(coins, len(coins), amount)
}

func dpMinCoin(coins []int, m int, n int) int {
	var table = make([]int, n+1)
	var minCoin = make([]int, len(coins))
	table[0] = 0

	for i := 1; i <= n; i++ {
		for j := 0; j < len(coins); j++ {
			minCoin[j] = -1
		}

		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				minCoin[j] = 1 + table[i-coins[j]]
			}
		}

		table[i] = -1

		for j := 0; j < len(minCoin); j++ {
			if minCoin[j] > 0 && (table[i] == -1 || table[i] > minCoin[j]) {
				table[i] = minCoin[j]
			}
		}
	}
	return table[n]
}
