package main

import (
	"fmt"
	"time"
)

func emailSender(emailChan <- chan string, done chan <- bool) {
	defer func() { done <- true}()

	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	emailChan := make(chan string, 100)
	done := make(chan bool)
	go emailSender(emailChan, done)

	for i := 1; i <= 100; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	} 

	fmt.Println("Done sending")
	close(emailChan)
	<- done
}