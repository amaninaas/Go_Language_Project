package main

import "fmt"

func main() {
	//declare array
	//var arrayName = [size of array/length]dataType{element of array}
	var scores = [4]int{88, 71, 53, 91}
	fmt.Println(scores)
	fmt.Println(len(scores))

	//{index:value}
	var lifes = [12]int{0: 8, 5: 9, 11: 88}
	fmt.Println(lifes)
	fmt.Println(len(lifes))

	//use for range to iterate on arrays forr
	//declare the index
	for index, value := range scores {
		fmt.Println("Index: ", index, "Value: ", value)

	}

	//without declare an index : _ using underscore
	// _ ignore the value
	for _, value := range scores {
		fmt.Println("Value: ", value)

	}

	for index, _ := range scores {
		fmt.Println("Index: ", index)

	}
}
