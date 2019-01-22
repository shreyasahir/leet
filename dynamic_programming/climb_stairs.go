package main

import (
	"fmt"
)

func main() {
	input := 2
	fmt.Println(climbStairs(input))
}
func climbStairs(n int) int {
	//return fib(n + 1)
	return climbStairsUtils(n, 2)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func climbStairsUtils(n, m int) int {
	var table = make([]int, n+1)
	table[0], table[1] = 1, 1

	for i := 2; i <= n; i++ {
		j := 1
		for j <= m && j <= i {
			table[i] = table[i] + table[i-j]
			j++
		}
	}
	return table[n]

}
