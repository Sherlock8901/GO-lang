package main
import ("fmt"
"time"
)

type Order struct{
	id string
	amount float32
	status string
	createAt time.Time
}
// It is like work consturctor in OOPs
func newOrder(id string, amount float32, status string) *Order {
	myorder:= Order{
		id : id,
		amount : amount,
		status : status,
	}
	return &myorder
}

func main(){
	myorder:= newOrder("1", 30.50, "received")
	fmt.Println(myorder)
}