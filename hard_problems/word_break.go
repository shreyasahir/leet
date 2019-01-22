// https://leetcode.com/problems/word-break-ii/
// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences.

// Note:

// The same word in the dictionary may be reused multiple times in the segmentation.
// You may assume the dictionary does not contain duplicate words.
// Example 1:

// Input:
// s = "catsanddog"
// wordDict = ["cat", "cats", "and", "sand", "dog"]
// Output:
// [
//   "cats and dog",
//   "cat sand dog"
// ]
// Example 2:

// Input:
// s = "pineapplepenapple"
// wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
// Output:
// [
//   "pine apple pen apple",
//   "pineapple pen apple",
//   "pine applepen apple"
// ]
// Explanation: Note that you are allowed to reuse a dictionary word.
// Example 3:

// Input:
// s = "catsandog"
// wordDict = ["cats", "dog", "sand", "and", "cat"]
// Output:
// []

// func wordBreak(s string, wordDict []string) []string {

// }

package main

import (
	"fmt"
	"strings"
)

var (
	output []string
)

func main() {
	s := "catsanddog"
	fmt.Println("input string", s)
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}

	fmt.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) []string {

	if len(s) == 0 {
		return []string{}
	}

	alphabets := make(map[rune]bool)
	for _, w := range wordDict {
		// fmt.Println("w", w)
		for _, c := range w {
			// fmt.Println("c", c)

			alphabets[c] = true
		}
	}

	// fmt.Println("alphabests", alphabets)

	for _, c := range s {
		// fmt.Println("second c", c)
		if _, ok := alphabets[c]; !ok {
			return []string{}
		}
	}

	var record = make(map[string][]string)
	return wordBreak2(s, wordDict, &record)
}

func wordBreak2(s string, wordDict []string, rDict *map[string][]string) []string {
	res := make([]string, 0)
	for _, w := range wordDict {
		head := strings.Index(s, w)

		fmt.Println("head", head)
		if head == 0 {
			_s := s[len(w):]
			if len(_s) == 0 {
				res = append(res, w)
			} else {
				if _r, ok := (*rDict)[_s]; !ok {
					_r = wordBreak2(_s, wordDict, rDict)
					for _, rs := range _r {
						res = append(res, fmt.Sprintf("%s %s", w, rs))
					}

				} else {
					for _, rs := range _r {
						res = append(res, fmt.Sprintf("%s %s", w, rs))
					}
				}

			}

		}
	}
	(*rDict)[s] = res
	return res
}
