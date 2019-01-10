package main

import (
	"fmt"
	"strings"
)

func main() {
	staircase(6)
}

func staircase(n int32) {
	var i int32
	for i = 1; i <= n; i++ {
		str := strings.Repeat(" ", int(n-i))
		str += strings.Repeat("#", int(i))
		fmt.Printf("%s\n", str)
	}
}
