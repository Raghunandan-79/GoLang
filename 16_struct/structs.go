package main

import (
	"fmt"
	"time"
)

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getStatus() string {
	return o.status
} 

func main() {
	var myOrder order = order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}

	fmt.Println("Order struct", myOrder)

	myOrder.createdAt = time.Now()

	fmt.Println(myOrder)

	fmt.Println(myOrder.status)

	myOrder2 := order {
		id: "2",
		amount: 100,
		status: "deliverd",
		createdAt: time.Now(),
	}

	fmt.Println(myOrder2)

	myOrder2.changeStatus("return")

	fmt.Println(myOrder2.getStatus())

	fmt.Println(myOrder2)

	// one time usage struct
	language := struct {
		name string
		isGood bool
	} {"Golang", true}

	fmt.Println(language)
}
