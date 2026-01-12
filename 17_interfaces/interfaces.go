package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using Razorpay", amount)
}

type stripe struct {}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using Stripe", amount)
}

type fakePayment struct {}

func (f fakePayment) pay(amount float32) {
	fmt.Println("Making payment using Fake Payment", amount)
}

func main() {
	razorpayGateway := razorpay{}
	razorpayPayment := payment{
		gateway: razorpayGateway,
	}
	razorpayPayment.makePayment(500)

	stripeGateway := stripe{}
	stripePayment := payment{
		gateway: stripeGateway,
	}
	stripePayment.makePayment(200)

	fakeGateway := fakePayment{}
	fakePaymentGatewayPay := payment{
		gateway: fakeGateway,
	}
	fakePaymentGatewayPay.makePayment(1000)
}