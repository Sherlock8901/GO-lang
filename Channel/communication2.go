package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printNum2(number chan int) {
	for num := range number {
		fmt.Println("Processing number", num)
		time.Sleep(time.Second)
	}
}

func main() {
	numChan := make(chan int)

	go printNum2(numChan)

	for {
		numChan <- rand.Intn(100) // correct function
	}
}