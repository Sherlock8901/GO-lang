package main
import "fmt"
func main() {
	age:= 10

	if age >=18{
		fmt.Println("person is an adult")
	} else if age >= 12{
		fmt.Println("person is a teeneger")
	} else {
		fmt.Println("person is a kid")
	}

}
