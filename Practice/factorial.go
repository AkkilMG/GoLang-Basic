package main

import "fmt"

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(fact(n))
}
