package main

import "fmt"

func main() {
	fmt.Println("Conditions Operations:")

	// var num = 8
	// num2 := 5

	// if num < 5 {
	// 	fmt.Println("This is less than ", num2)
	// } else {
	// 	fmt.Println("This is gratter than ", num2)
	// }

	var num = 2

	if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else if num == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("More than Three")
	}
}
