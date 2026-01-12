package main

import "fmt"

func printSlice[T comparable, V string] (items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

func main() {
	languages := []string{"golang", "java", "c++"}
	printSlice(languages, "Raghu")
}