package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

// returning multiple values
func getLanguages() (string, string, string) {
	return "Golang", "Java", "C++"
}

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	fmt.Println(add(5, 2))
	fmt.Println(subtract(5, 3))
	fmt.Println(getLanguages())
	lang1, lang2, lang3 := getLanguages()
	fmt.Println(lang1, lang2, lang3)

	fn := processIt()
	fmt.Println(fn(6))
}
