package main

import (
	"fmt"
	"sync"
)

func task(id int, W * sync.WaitGroup){
	defer W.Done()
	fmt.Println("Doing task", id)
}

func main(){
	var Wg sync.WaitGroup //why we are use wait here ? -> The reason is written in the notebook properly
	for i:=0;i<=10;i++ {
		Wg.Add(1)
		go task(i, &Wg)
	}
	Wg.Wait()
}