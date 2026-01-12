package main

import "fmt"

func printSlice[T comparable, V string] (items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	languages := []string{"golang", "java", "c++"}
	printSlice(languages, "Raghu")

	stackElements := []int{1, 2, 3, 4, 5}
	myStack := stack[int]{
		elements: stackElements,
	}
	fmt.Println(myStack)
}