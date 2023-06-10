package main

import (
	"fmt"
)

func binarySearch(a []int, key int, low int, high int) int {
	if high >= low {
		mid := int(low + (high-low)/2)
		if a[mid] == key {
			return 1
		} else if a[mid] > key {
			return binarySearch(a, key, low, mid-1)
		}
		return binarySearch(a, key, mid+1, high)
	}
	return 0
}

func main() {
	/*

		Enter list: 4 5 6 7
		Enter key: 5
		Found the key.
	*/
	var n, key int
	n = 1
	fmt.Print("Enter no of elements: ")
	fmt.Scan(&n)
	a := [n]int{}
	fmt.Print("Enter list: ")
	for i := 1; i <= n; i++ {
		fmt.Print(i)
		fmt.Scan(&a[i])
	}
	fmt.Print("Enter the key: ")
	fmt.Scan(&key)
	if binarySearch(a, key, 0, n) == 1 {
		fmt.Print("Found the key.")
	} else {
		fmt.Print("Not found the key.")
	}
}
