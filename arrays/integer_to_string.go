package main

import "fmt"

func main() {
	input := "    -42"
	fmt.Println("myatoi", myAtoi(input))
}

const (
	maxInt32 = 1<<31 - 1
	minInt32 = -1 << 31
)

func myAtoi(str string) int {
	var out int64
	// isParsing marked as we're parsing the number.
	// isParsing is set to true whenever we encountered the sign (+/-) or digits rune (0..9).
	// Note: if isParsing == true, then all runes will be discarded, except digits rune (0-9).
	sign, isParsing := "", false

	for _, r := range str {
		if r == ' ' {
			if isParsing { // "  1", " 1 123"
				break
			}
			continue
		}
		if r == '+' || r == '-' {
			if isParsing || sign != "" { // "--0", "++0", "1-123"
				break
			}
			isParsing = true
			sign = string(r)
			continue
		}
		if !(r >= 48 && r <= 57) { // If rune also not the digit then stop parsing
			break
		}

		isParsing = true
		out *= 10
		out += int64(r - 48)                     // Convert digit's code point to its actual integer value
		if out >= maxInt32 || -out <= minInt32 { // Stop parsing when overflow
			break
		}
	}

	if sign == "-" {
		out = -out
		if out < minInt32 {
			return minInt32
		}
	}
	if out > maxInt32 {
		return maxInt32
	}
	return int(out)
}

func myAtoi(str string) int {

	var newStr, result string

	count := 0
	for _, v := range str {
		// fmt.Println("v", v)
		if v != 32 && v < 97 {
			newStr += string(v)
		} else {
			if v == 32 {
				continue
			} else if count == 0 {
				return 0
			}
			continue
		}
		count++
	}

	var sign bool
	count = 0
	// fmt.Println("newStr", newStr)

	for _, v := range newStr {
		if string(v) == "-" && count == 0 {
			sign = true
		} else {
			if v < '0' || v > '9' {
				return 0
			} else {
				result += string(v)
			}

		}
		count++
	}

	// fmt.Println("result", result)

	var d rune
	var n int64
	for _, v := range result {
		if '0' <= v && v <= '9' {
			d = v - '0'
		}
		n *= 10
		n += int64(d)                        // Convert digit's code point to its actual integer value
		if n >= maxInt32 || -n <= minInt32 { // Stop parsing when overflow
			break
		}
	}

	if sign {
		n *= -1

		if n < minInt32 {
			return minInt32
		}
	}

	if n > maxInt32 {
		return maxInt32
	}
	return int(n)

}
