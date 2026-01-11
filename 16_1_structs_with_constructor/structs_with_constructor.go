package main

import "fmt"

type vehicle struct {
	wheels int
	brakes bool
	headlights int
}

func newVehicle(wheels int, brakes bool, headlights int) *vehicle {
	myVehicle := vehicle {
		wheels: wheels,
		brakes: brakes,
		headlights: headlights,
	}

	return &myVehicle
}

func main() {
	myVehicle := newVehicle(4, true, 2)
	fmt.Println(myVehicle)
	fmt.Println(myVehicle.wheels) 
}