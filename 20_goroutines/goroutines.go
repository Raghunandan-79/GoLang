package main

import (
	"fmt"
	"time"
)

// func task(id int) {
// 	fmt.Println("Executing task", id)
// }

func main() {
	// usage of function made outside main
	// for i := 0; i <= 10; i++ {
	// 	go task(i)
	// }

	// using inline function
	for i := 0; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}