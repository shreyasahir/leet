// https: //leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
// Say you have an array for which the ith element is the price of a given stock on day i.

// Design an algorithm to find the maximum profit. You may complete at most k transactions.

// Note:
// You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

// Example 1:

// Input: [2,4,1], k = 2
// Output: 2
// Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
// Example 2:

// Input: [3,2,6,5,0,3], k = 2
// Output: 7
// Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
//              Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.

package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{6, 1, 3, 2, 4, 7}
	k := 2
	fmt.Println("input", input)
	fmt.Println(maxProfit(k, input))
	// fmt.Println(maxDiff(k, input))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if k == 0 || n == 0 {
		return 0
	}

	if k > n/2 {
		k = n + 1
	}

	stock := make([]int, k*2)
	for k := range stock {
		stock[k] = int(math.MinInt32)
	}
	stock[0] = -prices[0]
	for i := 1; i < n; i++ {
		stock[0] = max(stock[0], -prices[i])
		fmt.Println("stock[0]", stock[0])

		for j := 1; j < 2*k; j += 2 {
			stock[j] = max(stock[j], stock[j-1]+prices[i])
			if j+1 < 2*k {
				stock[j+1] = max(stock[j+1], stock[j]-prices[i])
			}
		}
		fmt.Println(stock)
	}
	return max(0, stock[2*k-1])
}
