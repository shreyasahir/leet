package main

import (
	"fmt"
	//	"fmt"
)

func main() {
	s := ")()())"
	//s := ")"
	//s := "()(()"

	fmt.Println(longestValidParentheses(s))
}

func longestValidParentheses(s string) int {
	var left, right, maxLength int

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			left++
		} else {
			right++
		}

		if left == right {
			maxLength = max(maxLength, right*2)
		} else if right >= left {
			left = 0
			right = 0
		}

		//fmt.Println("maxlength_left", maxLength)
	}
	left = 0
	right = 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "(" {
			left++
		} else {
			right++
		}

		if left == right {
			maxLength = max(maxLength, left*2)
		} else if left >= right {
			left = 0
			right = 0
		}
		//fmt.Println("maxlength_right", maxLength)

	}

	return maxLength
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
