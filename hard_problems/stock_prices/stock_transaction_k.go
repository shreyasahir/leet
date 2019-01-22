package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 2, 6, 5, 0, 3}
	k := 2
	fmt.Println("arr", arr)
	fmt.Println(maxProfit(k, arr))
}

// // Function to find out maximum profit by buying &
// // selling/ a share atmost k times given stock price
// // of n days
// int maxProfit(int price[], int n, int k)
// {
//     // table to store results of subproblems
//     // profit[t][i] stores maximum profit using atmost
//     // t transactions up to day i (including day i)
//     int profit[k + 1][n + 1];

//     // For day 0, you can't earn money
//     // irrespective of how many times you trade
//     for (int i = 0; i <= k; i++)
//         profit[i][0] = 0;

//     // profit is 0 if we don't do any transation
//     // (i.e. k =0)
//     for (int j = 0; j <= n; j++)
//         profit[0][j] = 0;

//     // fill the table in bottom-up fashion
//     for (int i = 1; i <= k; i++) {
//         int prevDiff = INT_MIN;
//         for (int j = 1; j < n; j++) {
//             prevDiff = max(prevDiff,
//                            profit[i - 1][j - 1] - price[j - 1]);
//             profit[i][j] = max(profit[i][j - 1],
//                                price[j] + prevDiff);
//         }
//     }

//     return profit[k][n - 1];
// }

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
		for j := 1; j < 2*k; j += 2 {
			stock[j] = max(stock[j], stock[j-1]+prices[i])
			if j+1 < 2*k {
				stock[j+1] = max(stock[j+1], stock[j]-prices[i])
			}
		}
	}
	return max(0, stock[2*k-1])
}
