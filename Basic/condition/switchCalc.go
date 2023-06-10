package main

import (
	"fmt"
)

func main() {
	var a, b int
	var ch int = 0
	fmt.Println("1. Add\n2. Subtract\n3. Multiply\n4. Division\n5. Exit")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&ch)
	switch ch {
	case 1:
		fmt.Print("Enter a & b value: ")
		fmt.Scan(&a, &b)
		fmt.Println("The sum is", a+b)
		break
	case 2:
		fmt.Print("Enter a & b value: ")
		fmt.Scan(&a, &b)
		fmt.Println("The difference is", a-b)
		break
	case 3:
		fmt.Print("Enter a & b value: ")
		fmt.Scan(&a, &b)
		fmt.Println("The product is", a*b)
		break
	case 4:
		fmt.Print("Enter a & b value: ")
		fmt.Scan(&a, &b)
		fmt.Println("The quotient is", a/b)
		break
	case 5:
		return
	default:
		fmt.Println("Invalid choice.")
	}
}
