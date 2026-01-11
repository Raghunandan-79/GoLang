package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5

	// simple switch
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// Multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// Type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a String")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("Other")
		}
	}

	whoAmI(50)
}