package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch Case Operations:")

	today := time.Now().Weekday()

	fmt.Println(today + 1)

	var num = 5

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("More than Three ...")
	}

}
