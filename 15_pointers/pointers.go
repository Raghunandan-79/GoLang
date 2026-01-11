package main

import "fmt"

// pass by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	var num int = 1
	
	fmt.Println("Before changeNum in main", num)

	changeNum(&num)

	fmt.Println("After changeNum in main", num)
}