package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(1))
}

var (
	result []string
)

func generateParenthesis(n int) []string {
	var str string
	permute(n, n, str)
	return result
}

func permute(left int, right int, s string) {
	if left == 0 && right == 0 {
		//fmt.Println(s)
		// return s
		result = append(result, s)
	}
	if left > right {
		return
	}

	if left > 0 {
		permute(left-1, right, s+"(")
	}

	if right > 0 {
		permute(left, right-1, s+")")
	}
}
