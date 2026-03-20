package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	numbers := []int{10, 20, 30, 40}

	result := sum(numbers...) // slice passed using ...
	fmt.Println("Sum:", result)
}
