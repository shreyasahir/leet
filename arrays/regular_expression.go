// Regular Expression Matching
// Hard

// 2068

// 422

// Favorite

// Share
// Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

// '.' Matches any single character.
// '*' Matches zero or more of the preceding element.
// The matching should cover the entire input string (not partial).

// Note:

// s could be empty and contains only lowercase letters a-z.
// p could be empty and contains only lowercase letters a-z, and characters like . or *.
// Example 1:

// Input:
// s = "aa"
// p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".
// Example 2:

// Input:
// s = "aa"
// p = "a*"
// Output: true
// Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
// Example 3:

// Input:
// s = "ab"
// p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".
// Example 4:

// Input:
// s = "aab"
// p = "c*a*b"
// Output: true
// Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
// Example 5:

// Input:
// s = "mississippi"
// p = "mis*is*p*."
// Output: false

package main

import (
	"fmt"
)

func main() {
	s := "ab"
	p := ".*"
	fmt.Println("match", isMatch(s, p))
}

func isMatch(s string, p string) bool {

	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) > 1 && string(p[1]) == "*" {
		if isMatch(s, p[2:]) {
			return true
		}
		if len(s) > 0 && (string(p[0]) == "." || s[0] == p[0]) {
			return isMatch(s[1:], p)
		}
		return false
	} else {
		if len(s) > 0 && (string(p[0]) == "." || s[0] == p[0]) {
			return isMatch(s[1:], p[1:])
		}
		return false
	}
}
