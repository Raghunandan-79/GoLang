package main

import (
	"fmt"
	"sync"
)

// func task(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Executing task", id)
// }

func main() {
	var wg sync.WaitGroup

	// using function outside main
	// for i := 0; i <= 10; i++ {
	// 	wg.Add(1)
	// 	go task(i, &wg)
	// }

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Executing task", i)
		}(i)
	}

	wg.Wait()
}
