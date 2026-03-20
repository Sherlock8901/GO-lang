package main
import "fmt"

func main() {
 reverse:= 0
  var num int
  fmt.Scan(&num)

  for i:=num; i>0 ; i/=10{
	remainder:= i%10
    reverse = reverse * 10 + remainder
  }
    fmt.Println(reverse)
  	if num == reverse {
		fmt.Println("Number is palindrome")
	} else {
		fmt.Println("Number is not palindrome")
	}

}