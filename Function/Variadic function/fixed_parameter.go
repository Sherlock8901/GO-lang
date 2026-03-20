package main

import "fmt"

func printMsg(msg string, nums ...int) {
	fmt.Println(msg)
	for _, n := range nums {
		fmt.Println(n)
	}
}

func main() {
	printMsg("Numbers are:", 1, 2, 3)
}
