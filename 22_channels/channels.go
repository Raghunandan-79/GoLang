package main

import (
	"fmt"
	"math/rand"
	"time"
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
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number", num)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	numChan := make(chan int)
	go processNum(numChan)
	for {
		numChan <- rand.Intn(100)
	}
}