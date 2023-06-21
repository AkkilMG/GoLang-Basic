package main

import (
	"fmt"
	"math"
)

func largestPrimeFactor(n int) int {
	largest := 1
	for n%2 == 0 {
		largest = 2
		n /= 2
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		for n%i == 0 {
			largest = i
			n /= i
		}
	}
	if n > 2 {
		largest = n
	}
	return largest
}

func main() {
	/* 1 b
	Output:
	24
	The largest prime factor of 24 is 3.
	*/
	var n int
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Println("It is not an prime number.")
	} else {
		fmt.Printf("The largest prime factor of %d is %d.\n",n,largestPrimeFactor(n))
	}
}
