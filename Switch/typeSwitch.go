package main

import "fmt"

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Boolean")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	checkType(10)
	checkType("Hello")
	checkType(true)
	checkType(3.14)
}