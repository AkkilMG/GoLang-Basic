package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		* * * *
		 * * *
		  * *
		   *
	*/
	var n int
	fmt.Print("Enter n: ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Println(strings.Repeat(" ", i), strings.Repeat("* ", n-i))
	}
}
