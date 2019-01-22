// A password is considered strong if below conditions are all met:

// It has at least 6 characters and at most 20 characters.
// It must contain at least one lowercase letter, at least one uppercase letter, and at least one digit.
// It must NOT contain three repeating characters in a row ("...aaa..." is weak, but "...aa...a..." is strong, assuming other conditions are met).
// Write a function strongPasswordChecker(s), that takes a string s as input, and return the MINIMUM change required to make s a strong password. If s is already strong, return 0.

// Insertion, deletion or replace of any one character are all considered as one change.
// https://leetcode.com/problems/strong-password-checker/
// http://massivealgorithms.blogspot.com/2016/10/leetcode-420-strong-password-checker.html

package main

import (
	"fmt"
)

func main() {

	password := "aoveeA1983"
	fmt.Println("strongPasswordChecker", strongPasswordChecker(password))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func strongPasswordChecker(s string) int {
	n := len(s)
	if n < 6 {
		return 6 - n
	}
	end := s[0]
	upper := end >= 'A' && end <= 'Z'
	lower := end >= 'a' && end <= 'z'
	digit := end >= '0' && end <= '9'

	endRep := 1
	change := 0
	delete := [3]int{}

	for i := 0; i < n; i++ {
		if s[i] == end {
			endRep++
		} else {
			change += endRep / 3
			if endRep/3 > 0 {
				delete[endRep%3]++
			}
			end = s[i]
			upper = end >= 'A' && end <= 'Z'
			lower = end >= 'a' && end <= 'z'
			digit = end >= '0' && end <= '9'

			endRep = 1
		}
	}
	change += endRep / 3
	if endRep/3 > 0 {
		delete[endRep%3]++
	}
	checkReq := 0
	if !upper {
		checkReq += 1
	}
	if !lower {
		checkReq += 1
	}
	if !digit {
		checkReq += 1
	}
	if n > 20 {
		dele := n - 20

		if dele <= delete[0] {
			change -= dele
		} else if dele-delete[0] <= 2*delete[1] {
			change -= delete[0] + (dele - delete[0]/2)
		} else {
			change -= delete[0] + delete[1] + (dele-delete[0-2*delete[1]])/3
		}
		return dele + max(checkReq, change)
	} else {
		return max(6-n, max(checkReq, change))
	}
}
