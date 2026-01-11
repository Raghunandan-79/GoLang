package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating
	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"])

	/*
		if key doesnt exists then it returns 0 value
	*/
	fmt.Println("phone")

	fmt.Println(len(m))

	fmt.Println(m["area"])

	// deleting element
	delete(m, "area")
	fmt.Println(m["area"])

	fmt.Println(m)
	clear(m)
	fmt.Println(m)

	// making map without make
	m1 := map[string]int{"price": 40, "age": 18, "phones": 3}
	fmt.Println(m1)

	// checking doesnt exist
	k, ok := m1["price"]

	fmt.Println(k)

	if ok {
		fmt.Println("All ok")
	} else {
		fmt.Println("Not ok")
	}

	m2 := map[string]int{"price": 40, "age": 18, "phones": 3}
	
	fmt.Println(maps.Equal(m1, m2))
}