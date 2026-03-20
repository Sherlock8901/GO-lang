package main
import "fmt"

func main(){
	// var arr[]int
	// fmt.Println(len(arr))

    // make function (it is use to initialised a slice)
	// var arr = make([]int, 2,5) // for 2 one sice with the 0 value is added in the front. Here always put 0
	// fmt.Println(arr) //[0 0]
	// fmt.Println(cap(arr)) //used find out initial capacity of the slice

	// Proper way of writing:
	var arr = make([] int, 0,5) //We can write it in one another way:  nums:= []int{}  Then use append

	//Append function(used to add elements in the slice)
    arr = append(arr, 1, 2, 3, 4, 5)
    fmt.Println(arr)
	fmt.Println(cap(arr)) 
	// If the capacity is exceed then the slice will double from their current size 
	arr = append(arr, 1)  
	fmt.Println(arr)


	// 2D slice example

	var nums = [][]int {{1,2,3},{4,5,6}}
	fmt.Println(nums)

}