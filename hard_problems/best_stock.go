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
)

func main() {
	//input := []int{6, 1, 3, 2, 4, 7}
	input := []int{3, 3, 5, 0, 0, 3, 1, 4}
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

// if(prices==null||prices.length<=1)
// return 0;

// int min=prices[0]; // min so far
// int result=0;

// for(int i=1; i<prices.length; i++){
// result = Math.max(result, prices[i]-min);
// min = Math.min(min, prices[i]);
// }

// return result;

func maxProfit(k int, prices []int) int {

	if len(prices) == 0 {
		return 0
	}
	var result int
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		result = max(result, prices[i]-min)
		min = minimum(min, prices[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
