package main

import "fmt"

// Interface
type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

// Struct that uses the interface
type payment struct {
	gateway paymenter
}

// Method using interface (Open/Closed Principle)
func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

// Razorpay implementation
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Paid using Razorpay:", amount)
}

func (r razorpay) refund(amount float32, account string) {
	fmt.Println("Refunded", amount, "to account", account, "via Razorpay")
}

// Stripe implementation
type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Paid using Stripe:", amount)
}

func (s stripe) refund(amount float32, account string) {
	fmt.Println("Refunded", amount, "to account", account, "via Stripe")
}

func main() {
	// Using Razorpay
	r := razorpay{}
	p1 := payment{
		gateway: r,
	}
	p1.makePayment(100.5)

	// Using Stripe
	s := stripe{}
	p2 := payment{
		gateway: s,
	}
	p2.makePayment(200.75)

	// Refund example
	s.refund(50, "ACC123")
}