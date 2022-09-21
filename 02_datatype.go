package main

import "fmt"

func main() {
	fmt.Println("Datatypes")

	//declare a variable in go
	// var variableName datatype
	var score int
	var cityName string
	var check bool
	var lifeline float32

	score = 88
	cityName = "KL"
	check = true
	lifeline = 88.88

	fmt.Println(score)
	fmt.Println(cityName)
	fmt.Println(check)

	fmt.Println("Score: ", score)
	fmt.Println("CityName: ", cityName)
	fmt.Println("Check: ", check)
	fmt.Println("Lifeline: ", lifeline)
	fmt.Println("My Score is ", score, ". I live in", cityName)

	//use fromat specifier
	//add backspace \n to add new line
	fmt.Printf("My Score is %v. I live in %v. \n", score, cityName)
	fmt.Printf("My lifeline is %v. \n", lifeline)

	//%V format specifier - get datatype of variable
	fmt.Printf("Datatype of lifeline variable is %V. \n", lifeline)

	//auto type inference
	//let go detect the datatype of the variable automatically
	//use :=
	autoScore := "PASS"
	fmt.Printf("My autoScore is %v. \n", autoScore)
	fmt.Printf("Datatype of autoscore variable is %V. \n", autoScore)
}
