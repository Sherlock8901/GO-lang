package main
import "fmt"

func main(){
nums := []int{}
nums = append(nums, 1)

nums1 := make([]int, len(nums)) // allocate length
copy(nums1, nums)

fmt.Println(nums, nums1)

}