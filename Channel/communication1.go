package main

import (
	"fmt"
	"time"
)

func printNum(number chan int){
	fmt.Println("The processed number is:",<-number)
}


func main() {
	messageChan := make(chan int)
	go printNum(messageChan)
	messageChan <- 5
	time.Sleep(time.Second)
	
}