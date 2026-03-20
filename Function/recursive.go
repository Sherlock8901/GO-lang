package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println("Factorial:", factorial(x))
}
