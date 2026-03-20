package main
import "fmt"

func main() {
	var num = 123
    for i:=num; i>0; i/=10 {
      var remainder int =  i%10
	//   fmt.Println(remainder) It is used for printing in the different lines
	     fmt.Print(remainder) //For single line printing
	}

}
