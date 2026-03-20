package main
import "fmt"




// It is for the string type
type OrderStatus2 string

const(
	Received OrderStatus2 = "received"
	Confirmed = "confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
)
type OrderStatus1 int

const(
	Received1 OrderStatus1 = iota
	Confirmed1
	Prepared2
	Delivered3
)

func changestatus1(status OrderStatus2){
	fmt.Println("Changing the status into",status)
}
func changestatus2(status OrderStatus1){
	fmt.Println("Changing the status into",status)
}

func main(){
	changestatus1(Confirmed)
	changestatus2(Confirmed1)
}