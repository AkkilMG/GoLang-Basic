package main

import (
	"fmt"
)

func main() {
	var sum, n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i*i
	}
	fmt.Println(sum)
}
