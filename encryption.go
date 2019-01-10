package main

import (
	"math"
)

func main() {
	s := "chillout"
	encryption(s)
}

func encryption(s string) string {
	var str string
	for _, v := range s {
		newString := string(v)
		if newString != " " {
			str += newString
		}
	}
	n := len(str)
	rootN := math.Sqrt(float64(n))
	row := int(math.Floor(rootN))
	col := int(math.Ceil(rootN))
	// fmt.Println("row", row)
	// fmt.Println("col", col)

	for row*col < n {
		row++
	}

	var output string
	for i := 0; i < col; i++ {
		var temp string
		k := i
		for k < n {
			// fmt.Println("kinner", k)
			temp += string(str[k])
			k += col
		}
		// fmt.Println("kouter", k)
		// fmt.Println("row", row)
		// fmt.Println("col", col)
		output += temp + " "
		// fmt.Println("temp", temp)

	}
	// fmt.Println("output", output)
	return output
}
