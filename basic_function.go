package main

import "fmt"

func add(x int, y int) int { //Last int is the return type
	output := x + y
	return output
}

func add2(x, y int) int { //If all the parameters are same type
	output := x + y
	return output
}

func main() {
	num1 := 5
	num2 := 6

	result := add(num1, num2)

	fmt.Println("The result is: ", result)
}
