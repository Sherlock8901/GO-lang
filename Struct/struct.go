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

func main(){
	Order1 := Order {
        id : "1",
		amount: 100.0,
		status : "Delivered",
		createAt : time.Now(),
	}

		Order2 := Order {
        id : "2",
		amount: 150.0,
		status : "In transit",
		createAt : time.Now(),
	}


	fmt.Println(Order1)
	fmt.Println(Order2)
	fmt.Println(Order1.status)
	Order1.status = "Paid"
	fmt.Println(Order1.status)

}