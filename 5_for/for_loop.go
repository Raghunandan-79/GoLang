package main

import "fmt"

func main() {
	// while loop using for
	i := 1

	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println()

	// for loop
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	fmt.Println()

	// range
	for j := range 3 + 1 {
		fmt.Println(j)
	}
}