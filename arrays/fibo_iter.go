package main

import "fmt"

func main() {

	fmt.Println(fib(11))
}

func fib(N int) int {
	n1 := 0
	n2 := 1
	result := 0
	for i := 2; i <= N; i++ {
		result = n1 + n2
		n1 = n2
		n2 = result
	}
	return result
}
