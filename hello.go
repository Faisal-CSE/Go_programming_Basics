package main

import (
	"fmt"
)

func main() {
	// var num1 int
	// var num2 int
	// num1 = 15
	// num2 = 8
	// var result = num1 + num2

	// var num1 = 15
	// var num2 = 8
	// var result = num1 + num2

	var num1, num2 int

	num1, num2 = 2, 3
	var result = num1 + num2

	const PI__Val = 3.14

	result__new := 6 //Means >> var result__new = 6 || create and assign a variable. This is short form

	fmt.Println("Summation is: ")
	fmt.Println(result)
	fmt.Println(result__new)
	fmt.Println(PI__Val)

	fmt.Println("-------------------")

	var name = "Faisal Porag"
	fmt.Print(name)
}
