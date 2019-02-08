package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/text-justification/
// http://buttercola.blogspot.com/2015/10/leetcode-text-justification.html

func main() {
	arr := []string{"This", "is", "an", "example", "of", "text", "justification."}
	// arr := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth := 16
	fmt.Println("arr", arr)

	textJustification(arr, maxWidth)
}

func textJustification(words []string, maxWidth int) {

	var subtext string
	var width int
	var result []string
	var cnt int
	for i := 0; i < len(words); i++ {
		// fmt.Println("words", words[i])
		if width < maxWidth && len(words[i]) <= (maxWidth-width) {
			subtext += words[i] + " "
			cnt++
			width = len(subtext)
			// fmt.Println("subtext1", subtext)

		} else {
			if maxWidth-len(subtext) > 0 {
				subtext = addSpace(subtext, maxWidth-width, cnt)
			}
			result = append(result, subtext)
			fmt.Println("subtext2", subtext)
			subtext = words[i] + " "
			width = len(subtext)

		}

	}

	if len(subtext) > 0 {
		left := maxWidth - len(subtext)
		for left > 0 {
			subtext += " "
			left--
		}
		result = append(result, subtext)
	}
	fmt.Println("result ", result)
	printArray(result)
}

func printArray(result []string) {
	for _, v := range result {
		fmt.Println(v)
	}
}

func addSpace(subtext string, gap int, cnt int) string {

	var space int
	if gap%cnt == 0 {
		space = gap / cnt
	} else {
		space = (gap / cnt) + 1
	}

	arr := strings.Split(subtext, " ")

	var sub, s string
	var j int
	for i := 0; i < len(arr); i = i + 1 {
		sub += arr[i]
		fmt.Println("arr[i]", arr[i])
		for j < space && gap >= j {
			s += " "
			j++
			gap = gap - space
		}
		sub += s
	}
	return sub
}
