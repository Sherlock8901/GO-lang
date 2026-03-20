package main
import "fmt"

func main(){
	var n int
	fmt.Println("Enter the size of the array:")
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter the elements:")
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	maxi:=0
	for i:=0;i<n;i++{
		if arr[i]>maxi{
			maxi=arr[i]
		}
	}
	fmt.Println("Maximum element of the array is:",maxi)
    sec:=0
	for i:=0;i<n;i++{
		if arr[i]==maxi{
			continue
		} else if sec<arr[i]{
			sec=arr[i]
		}
	}
	fmt.Println("Second max element of the array is:",sec)
}
