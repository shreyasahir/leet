// Max profit with at most two transactions =
//        MAX {max profit with one transaction and subarray price[0..i] +
//             max profit with one transaction and aubarray price[i+1..n-1]  }
// i varies from 0 to n-1.

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
	fmt.Println(ktransaction(input))
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

// We can do this O(n) using following Efficient Solution. The idea is to store maximum possible profit of every subarray and solve the problem in following two phases.

// 1) Create a table profit[0..n-1] and initialize all values in it 0.

// 2) Traverse price[] from right to left and update profit[i] such that profit[i] stores maximum profit achievable from one transaction in subarray price[i..n-1]

// 3) Traverse price[] from left to right and update profit[i] such that profit[i] stores maximum profit such that profit[i] contains maximum achievable profit from two transactions in subarray price[0..i].

// 4) Return profit[n-1]

// To do step 1, we need to keep track of maximum price from right to left side and to do step 2, we need to keep track of minimum price from left to right. Why we traverse in reverse directions? The idea is to save space, in second step, we use same array for both purposes, maximum with 1 transaction and maximum with 2 transactions. After an iteration i, the array profit[0..i] contains maximum profit with 2 transactions and profit[i+1..n-1] contains profit with two transactions.

// # Returns maximum profit with two transactions on a given
// # list of stock prices price[0..n-1]
// def maxProfit(price,n):

//     # Create profit array and initialize it as 0
//     profit = [0]*n

//     # Get the maximum profit with only one transaction
//     # allowed. After this loop, profit[i] contains maximum
//     # profit from price[i..n-1] using at most one trans.
//     max_price=price[n-1]

//     for i in range( n-2, 0 ,-1):

//         if price[i]> max_price:
//             max_price = price[i]

//         # we can get profit[i] by taking maximum of:
//         # a) previous maximum, i.e., profit[i+1]
//         # b) profit by buying at price[i] and selling at
//         #    max_price
//         profit[i] = max(profit[i+1], max_price - price[i])

//     # Get the maximum profit with two transactions allowed
//     # After this loop, profit[n-1] contains the result
//     min_price=price[0]

//     for i in range(1,n):

//         if price[i] < min_price:
//             min_price = price[i]

//         # Maximum profit is maximum of:
//         # a) previous maximum, i.e., profit[i-1]
//         # b) (Buy, Sell) at (min_price, A[i]) and add
//         #    profit of other trans. stored in profit[i]
//         profit[i] = max(profit[i-1], profit[i]+(price[i]-min_price))

//     result = profit[n-1]

//     return result
func ktransaction(arr []int) int {
	n := len(arr)
	var maxDiff int
	for i := 0; i < n; i++ {
		maxDiff = max(maxDiff, DynamicProgrammingSingleSellProfit(arr[0:i])+DynamicProgrammingSingleSellProfit(arr[i+1:n]))
	}
	return maxDiff
}
