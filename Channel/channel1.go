package main

import "fmt"

func main() {
	messageChan := make(chan string)
	messageChan <- "new"
	res := <-messageChan
	fmt.Println(res)

	// fatal error: all goroutines are asleep - deadlock!(Output)
}