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

//Return multiple values
func calc(x, y int) (int, int) {
	var output1 = x + y
	var output2 = x - y

	return output1, output2
}

func main() {
	num1 := 5
	num2 := 6

	num3 := 10
	num4 := 4

	result := add(num1, num2)
	fmt.Println("The result is: ", result)

	result1, result2 := calc(num3, num4)
	fmt.Println("The result1 is: ", result1, " || The result2 is: ", result2)
}
