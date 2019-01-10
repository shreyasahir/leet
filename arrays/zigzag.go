package main

import (
	"fmt"
)

func zigzag(s string, n int) string {
	var result string
	if n == 1 {
		return s
	}
	var arr = make([]string, n)
	var down = true
	row := 0

	for i := 0; i < len(s); i++ {
		arr[row] += string(s[i])
		if row == n-1 {
			down = false
		}
		if row == 0 {
			down = true
		}
		if down {
			row++
		} else {
			row--
		}
	}

	//fmt.Println("arr", arr)
	for i := 0; i < n; i++ {
		//fmt.Printf("%s", arr[i])
		result += arr[i]
	}
	return result
}
func main() {
	s := "AB"
	n := 1
	fmt.Println(zigzag(s, n))
	fmt.Println("expected")
	fmt.Println("PINALSIGYAHRPI")
}
