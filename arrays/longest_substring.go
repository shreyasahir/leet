package main

import (
	"fmt"
)

func longSubSeq(s string) int {
	var ans, n int
	n = len(s)
	arr := make(map[rune]int)
	s1 := []rune(s)
	for i, j := 0, 0; j < n; j++ {
		if arr[s1[j]] != 0 {
			i = max(arr[s1[j]], i)
		}

		ans = max(ans, j-i+1)
		arr[s1[j]] = j + 1
	}
	return ans
}
func main() {
	fmt.Println("long", longSubSeq("dvdf"))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
