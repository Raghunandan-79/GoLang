package main

import "fmt"

const age = 30

func main() {
	const name string = "Raghu"
	fmt.Println(name)

	fmt.Println(age)

	// constant grouping
	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port)
	fmt.Println(host)
}