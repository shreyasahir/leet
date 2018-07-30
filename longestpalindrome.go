package main

import (
	"fmt"
)

func longestPalindrome(s string) {
	//start := 0
	end := len(s) - 1
	//count := 0
	for i := 0; i <= end/2; i++ {
		fmt.Println("i,mid,end", i, end/2, end-i)
		fmt.Println("left,right", s[i:end/2], s[end/2:end+1-i])

		if s[i:end/2] == s[end/2:end+1-i] {
			fmt.Println("palindrome", s[i:end+1-i])
			//count += 2
		}
	}
	//print(count)
	//return count
}

func main() {
	longestPalindrome("cbbd")
}
