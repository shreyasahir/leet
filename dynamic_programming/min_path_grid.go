package main

import (
	"fmt"
)

func main() {
	// input := [][]int{
	// 	{1, 3, 1},
	// 	{1, 5, 1},
	// 	{4, 2, 1},
	// }
	input := [][]int{{9, 1, 4, 8}}
	//var m = len(input[0])
	//var n = len(input)

	fmt.Println(minPathSum(input))
}

// func minPathSum(grid [][]int) int {
// 	//return minPathUtils(grid, 0, 0, len(grid)-1, len(grid[0])-1)
// 	return minPathDP(grid, 0, 0, len(grid), len(grid[0]))

// }

func minPathUtils(grid [][]int, start, end, m, n int) int {
	//fmt.Println("start,end", start, end)
	if start == m && end == n {
		return grid[m][n]
	}
	if start == m {
		return grid[start][end] + minPathUtils(grid, start, end+1, m, n)
	}
	if end == n {
		return grid[start][end] + minPathUtils(grid, start+1, end, m, n)
	}
	return min(grid[start][end]+minPathUtils(grid, start+1, end, m, n), grid[start][end]+minPathUtils(grid, start, end+1, m, n))

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minPathDP(grid [][]int, start, end, m, n int) int {
	fmt.Println("m,n", m, n)
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	fmt.Println(dp)
	// return 0
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	//fmt.Println(dp)
	for i := 1; i < n; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	//fmt.Println(dp)

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])

		}
	}
	//fmt.Println(dp[m][n])
	return dp[m][n]
}

func minPathSum(grid [][]int) int {
	rowLen := len(grid)
	colLen := len(grid[0])

	for i := 1; i < colLen; i++ {
		grid[0][i] += grid[0][i-1]
	}

	for i := 1; i < rowLen; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			left := grid[i][j-1] + grid[i][j]
			up := grid[i-1][j] + grid[i][j]
			if left > up {
				grid[i][j] = up
			} else {
				grid[i][j] = left
			}
		}
	}

	return grid[rowLen-1][colLen-1]
}
