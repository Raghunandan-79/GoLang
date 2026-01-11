package main

import "fmt"

func main() {
	age := 18

	// if else
	if age >= 18 {
		fmt.Println("Person is adult")
	} else {
		fmt.Println("Person is not an adult")
	}

	wheels := 4

	// if else if
	if wheels >= 4 {
		fmt.Println("Not a bike")
	} else if wheels == 2 {
		fmt.Println("It is a bike")
	} else if wheels == 3 {
		fmt.Println("It is a tricycle")
	} else {
		fmt.Println("Can't recognize vehicle")
	}

	// declare variables in if
	if role := "Admin"; role == "Admin" {
		fmt.Println("Admin access granted")
	} else {
		fmt.Println("No adming access")
	}
}