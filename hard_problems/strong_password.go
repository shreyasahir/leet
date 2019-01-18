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
	requiredChar := getRequiredChar(s)

	fmt.Println("requiredchar", requiredChar)
	if len(s) < 6 {
		return max(requiredChar, 6-len(s))
	}

	replace := 0
	oned := 0
	twod := 0

	for i := 0; i < len(s); {
		length := 1
		for i+length < len(s) && s[i+length] == s[i+length-i] {
			length++
		}
		if length >= 3 {
			replace += length / 3
			if length%3 == 0 {
				oned++
			}
			if length%3 == 1 {
				twod++
			}
		}
		i += length
	}

	if len(s) <= 20 {
		return max(requiredChar, replace)
	}
	deleteCount := len(s) - 20

	replace -= min(deleteCount, oned)
	replace -= min(max(deleteCount-oned, 0), twod) / 2
	replace -= max(deleteCount-oned-twod, 0) / 3

	return deleteCount + max(requiredChar, replace)
}

func getRequiredChar(s string) int {
	lowercase, uppercase, digit := 1, 1, 1

	for _, v := range s {
		if 0 <= int(v-'0') && int(v-'0') <= 9 {
			digit = 0
		} else if v >= 'a' && v <= 'z' {
			lowercase = 0
		} else if v >= 'A' && v <= 'Z' {
			uppercase = 0
		}
	}
	return lowercase + digit + uppercase

}
