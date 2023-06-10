package main

import (
	"fmt"
)

func main() {
	/*
		1 2 3 4
		1 2 3
		1 2
		1
	*/
	var n int
	fmt.Print("Enter n: ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		for j := 1; j < n+1-i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
