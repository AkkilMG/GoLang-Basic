package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	     *
	    * *
	   * * *
	    * *
	     *
	*/
	var n int
	fmt.Print("Enter n: ")
	fmt.Scanln(&n)
	for i := 0; i < n-1; i++ {
		fmt.Println(strings.Repeat(" ", n-i-1), strings.Repeat("* ", i+1))
	}
	for i := 0; i < n; i++ {
		fmt.Println(strings.Repeat(" ", i), strings.Repeat("* ", n-i))
	}
}
