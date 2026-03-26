package main
import "fmt"

// This are the writing style of the generics
// func printslice[T any](items[]T){
// func printslice[T interface{}](items[]T){
func printslice[T int|string](items[]T){
	for _, item := range items {
		fmt.Println(item)
	}
}

type stack [T any]struct{
	element []T
}

func main(){
	printslice([]int{1,2,3})
	printslice([] string{"Dip","Patra"})

	// mystack := stack[int]{
	// 	element: []int{1,2,3}, //Like that we can use generics for struct also.
	// }
	mystack := stack[string]{
		element: []string{"Dip","Power"}, //Like that we can use generics for struct also.
	}
	fmt.Println(mystack)
}