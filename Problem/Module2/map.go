package main

import "fmt"

func main() {
	/*
		Output:
		Value of apple :  1
		Value of Banana :  2
		Updated value of apple :  5
		After deleting orange :  map[apple:5 banana:2]
		Value of banana :  2
		Key :  apple  Value :  5
		Key :  banana  Value :  2
		Length of the map :  2
	*/
	myMap := make(map[string]int)

	myMap["apple"] = 1
	myMap["banana"] = 2
	myMap["orange"] = 3

	appleValue := myMap["apple"]
	bananaValue := myMap["banana"]
	fmt.Println("Value of apple : ", appleValue)
	fmt.Println("Value of Banana : ", bananaValue)

	myMap["apple"] = 5

	fmt.Println("Updated value of apple : ", myMap["apple"])

	delete(myMap, "orange")
	fmt.Println("After deleting orange : ", myMap)

	value, exists := myMap["banana"]
	if exists {
		fmt.Println("Value of banana : ", value)
	} else {
		fmt.Println("Banana not found in the map.")
	}

	for key, value := range myMap {
		fmt.Println("Key : ", key, " Value : ", value)
	}

	fmt.Println("Length of the map : ", len(myMap))
}
