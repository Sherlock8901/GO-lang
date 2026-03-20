package main

import (
	"fmt" )

type customer struct {
	name string
	ph   string
}

type Order struct {
	id        string
	amount    float32
	status    string
	// createdAt time.Time
	customer
}

func main() {
	newOrder := Order{
	id:        "1",
	amount:    30.40,
	status:    "Paid",
	customer: customer{
		name: "dip",
		ph:   "81018987",
	},
}

	fmt.Println(newOrder)
	fmt.Println(newOrder.customer.name)
}