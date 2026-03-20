package main
import "fmt"

import "slices"

func main(){
	var nums1 = []int {1,2}
	var nums2 = [] int{1,2}

fmt.Println(slices.Equal(nums1,nums2)) //It return true or false
}