package main

import (
	"fmt"
)

func main() {
	var sum, n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
