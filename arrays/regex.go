package main

import (
	"fmt"
)

func main() {
	s := "aab"
	p := "c*a*b"

	fmt.Println(isMatch(s, p))
}

func isMatch(s string, p string) bool {

	var prev string
	var count, i, j int

	for i, j = 0, 0; i < len(s) && j < len(p); {
		currP := string(p[j])
		currS := string(s[i])
		count++
		fmt.Println("prev", prev)
		fmt.Println("currS", currS)
		fmt.Println("currP", currP)

		if currP == "*" && prev != "" && (prev == currS || prev == ".") {
			fmt.Println("first")
			i++
			j++
			//continue
		} else if currP == "." {
			fmt.Println("second")

			i++
			j++
			prev = "."
			//continue
		} else if currP == currS {
			prev = currS
			i++
			j++
			//continue
		} else {
			if string(p[j+1]) == "*" {
				j++
			} else {
				return false
			}
		}
	}
	fmt.Println("i,j,count", i, j, count)
	if count >= len(s) && count < len(p) && string(p[count]) != "*" {
		return false
	} else if count < len(s) && count >= len(p) && string(p[len(p)-1]) != "*" {
		return false
	}
	return true
}
