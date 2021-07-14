package main

import (
	"fmt"
)

func main() {
	var a [5]int

	fmt.Println("Elements: ", a)

	a[3] = 12
	a[4] = 10

	fmt.Println("Elements: ", a)

	fmt.Println("Get index 4 data: ", a[4])

	fmt.Println("Length of array: ", len(a))

	//-------------------------------

	arr__elements := [5]int{1, 3, 5, 7, 9}
	fmt.Println("Array is : ", arr__elements)

	for i := 0; i < len(arr__elements); i++ {
		if arr__elements[i] == 3 {
			arr__elements[i] = 500 //Value change using condition
		}
		fmt.Println("Serial is: ", arr__elements[i])
	}
	
	//------------------2D Array---------------------
	var twoD [2][3]int

	for p := 0; p < 2; p++ {
		for q := 0; q < 3; q++ {
			twoD[p][q] = p + q
		}
	}
	fmt.Println("2D Result : ", twoD)
}
