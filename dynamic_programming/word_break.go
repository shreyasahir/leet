package main

import (
	"fmt"
)

func main() {
	// s := "aaaaaaa"
	s := "catsanddog"

	// wordDict := []string{"aaaa", "aaa"}
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}

	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	//fmt.Println("value of s,size", s, size)

	if size == 0 {
		return true
	}

	var dp = make([]bool, size+1)

	for i := 1; i <= size; i++ {
		if dp[i] == false && dictContains(s[:i], wordDict) {
			dp[i] = true
			if dp[i] == true {
				if i == size {
					return true
				}
				for j := i + 1; j <= size; j++ {
					//fmt.Println("value of i,j", i, j)
					if dp[j] == false && dictContains(s[i:j], wordDict) {
						i = j
						dp[j] = true
					}
					if dp[j] == true && j == size {
						return true
					}
				}
			}
		}
	}
	return false
}

func dictContains(s string, wordDict []string) bool {
	//fmt.Println("value in dictcontains", s)
	for i := 0; i < len(wordDict); i++ {

		if wordDict[i] == s {
			return true
		}
	}
	return false
}
