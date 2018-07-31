package main

import (
	"fmt"
	"math"
)

func main() {
	s := "  0000000000012345678"
	fmt.Println(myAtoi(s))
}

func myAtoi(str string) int {
	var signBit bool
	var res int
	for _, v := range str {
		if string(v) != " " {
			//fmt.Println("inside for", string(v))
			if string(v) == "-" {
				signBit = true
			} else if string(v) == "+" {
				signBit = false
			} else if isNumber(v) {
				//fmt.Println("res", res)
				//fmt.Println("res", int(v)-'0')

				res = res*10 + (int(v) - '0')
			} else {
				break
			}
		}
	}

	if signBit {
		res = -1 * res
		if res < math.MinInt32 {
			return math.MinInt32
		}
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}

func isNumber(s rune) bool {
	//fmt.Println(int(s) - 0)
	//fmt.Println("return bool", int(s)-'0' > 0 && int(s)-'0' < 10)
	return int(s)-'0' >= 0 && int(s)-'0' < 10
}
