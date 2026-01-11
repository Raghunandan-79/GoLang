package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}

	sum := 0
	for i, num := range nums {
		fmt.Println(i, num)
		sum += num
	}
	fmt.Println(sum)

	m := map[string]string{
		"John":"Doe",
		"Raghunandan":"Sharma",
		"Daksh":".",
		"Pankaj":"Singh",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for i, c := range "Golang" {
		fmt.Println(i, string(c))
	}
}
