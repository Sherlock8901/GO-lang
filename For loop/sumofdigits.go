package main
import "fmt"

func main() {
 sum:= 0
  var num int
  fmt.Scan(&num)

  for i:=num; i>0 ; i/=10{
	remainder:= i%10
    sum = sum + remainder
  }
  fmt.Println("Sum of the digits of the number is:", sum)

}