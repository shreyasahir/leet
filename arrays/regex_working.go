package main

import (
	"fmt"
)

func main() {
	s := "ab"
	p := ".*"
	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {
	fmt.Println("s,p", s, p)
	if len(p) == 0 {
		return false
	}

	firstMatch := len(s) > 0 && (p[0] == s[0] || string(p[0]) == ".")
	fmt.Println("firstmatch", firstMatch)
	if len(p) >= 2 && string(p[1]) == "*" {
		return isMatch(s, p[2:]) || firstMatch && isMatch(s[1:], p)
	}
	return firstMatch && isMatch(s[1:], p[1:])
}

// class Solution(object):
//     def isMatch(self, text, pattern):
//         if not pattern:
//             return not text

//         first_match = bool(text) and pattern[0] in {text[0], '.'}

//         if len(pattern) >= 2 and pattern[1] == '*':
//             return (self.isMatch(text, pattern[2:]) or
//                     first_match and self.isMatch(text[1:], pattern))
//         else:
//             return first_match and self.isMatch(text[1:], pattern[1:])
