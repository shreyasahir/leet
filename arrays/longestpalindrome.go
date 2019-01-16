package main

import "fmt"

func longestPalindrome(s string) string {

	if s == "" {
		return ""
	}
	if len(s) == 1 {
		return s
	}

	longest := 1
	var str string
	str = string(s[0])
	n := len(s)

	for i := 0; i < n; i++ {
		tmp := helper(s, i, i)
		if len(tmp) > longest {
			longest = len(tmp)
			str = tmp
		}
		fmt.Println("tmp", tmp)
		tmp = helper(s, i, i+1)
		if len(tmp) > longest {
			longest = len(tmp)
			str = tmp
		}
		fmt.Println("tmp", tmp)

	}
	return str
}

func helper(s string, begin int, end int) string {
	for begin >= 0 && end < len(s) && s[begin] == s[end] {
		begin--
		end++
	}
	return s[begin+1 : end]
}
func main() {
	fmt.Println(longestPalindrome("ac"))
}
