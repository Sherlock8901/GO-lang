package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	maxi:=0
	for i:=0;i<n;i++{
		if arr[i]>maxi{
			maxi=arr[i]
		}
	}
	fmt.Println("Maximum element of the array is: ",maxi)
}
