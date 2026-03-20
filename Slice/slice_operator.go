package main
import "fmt"

func main(){
	nums:= []int {}
	nums = append(nums, 1,2,3)

	fmt.Println(nums[0:2])
	fmt.Println(nums[:2])
}