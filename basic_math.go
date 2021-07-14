package main

import (
	"fmt"
	"math"
)

func main() {
	var numData float64 = 15

	var result = math.Sqrt(numData)

	// var result__new = math.Round(result)
	// var result__new = math.Ceil(result)
	var result__new = math.Floor(result)

	fmt.Printf("The result is %.2f Thank you.", result)
	fmt.Println()
	fmt.Printf("The result is %.2g Thank you.", result)
	fmt.Println()
	fmt.Printf("The result is %.2f Thank you.", result__new)
}
