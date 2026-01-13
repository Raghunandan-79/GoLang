package main

import "fmt"

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Processing...")
}

func main() {
	done := make(chan bool)
	go task(done)
	<- done
}