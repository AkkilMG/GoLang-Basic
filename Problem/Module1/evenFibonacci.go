package main

import "fmt"

func main() {
	// 1 a
	var sum, prev, current int
	prev = 0
	current = 1
	sum = 0
	for prev+current < 40 {
		next := prev + current
		if next%2 == 0 {
			sum += next
		}
		prev = current
		current = next
	}
	fmt.Println("Sum of Fibonacci: ", sum)
}
