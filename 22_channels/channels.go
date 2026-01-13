package main

import (
	"fmt"
	// import below two when sending data
	// "math/rand"
	// "time"
)

// below two are for sending
// without loop
// func processNum(numChan chan int) {
// 	fmt.Println("Processing number", <-numChan)
// }

// func main() {
// 	numChan := make(chan int)
// 	go processNum(numChan)
// 	numChan <- 5
// 	time.Sleep(time.Second * 1)
// }

// with loop
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func main() {
// 	numChan := make(chan int)
// 	go processNum(numChan)
// 	for {
// 		numChan <- rand.Intn(100)
// 	}
// }

// receiving data
func sum(result chan int, num1 int, num2 int) {
	sumResult := num1 + num2
	result <- sumResult
}

func main() {
	result := make(chan int)
	go sum(result, 4, 5)
	res := <- result
	fmt.Println(res)
}