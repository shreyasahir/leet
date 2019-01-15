package main

import (
	"fmt"
)

func main() {
	coin := []int{1, 2, 5}
	amount := 11

	fmt.Println(coinChange(coin, amount))
}

func coinChange(coins []int, amount int) int {
	m := len(coins)
	//return coinChanges(coins, m, amount)
	return dpCoin(coins, m, amount)
}
func dpCoin(coins []int, m, n int) int {
	var table = make([]int, n+1)
	table[0] = 1
	for i := 0; i < m; i++ {
		for j := coins[i]; j < n+1; j++ {
			table[j] += table[j-coins[i]]
		}
	}
	return table[n]
}

func coinChanges(coins []int, m, n int) int {
	if m <= 0 {
		return 0
	}

	if n < 0 {
		return -1
	}

	if n == 0 {
		return 1
	}

	return coinChanges(coins, m, n-coins[m-1]) + coinChanges(coins, m-1, n)
}
