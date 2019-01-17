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

import "fmt"

func main() {
	input := []int{3, 3, 5, 0, 0, 3, 1, 4}
	k := 2
	fmt.Println(maxProfit(k, input))
}

func maxProfit(k int, prices []int) int {
	var result []int
	for i := 1; i < len(prices); i++ {
		result = append(result, prices[i]-prices[i-1])
	}

	fmt.Println("result", result)
	count := 0
	value := 0
	for _, v := range result {
		if k*2 == count {
			break
		}
		if v > 0 {
			value += v
			count++
		}

	}
	return value
}
