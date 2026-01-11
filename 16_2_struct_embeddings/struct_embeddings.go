package main

import "fmt"

type vehicle struct {
	wheels int
	brakes bool
	headlights int
}

type fourWheeler struct {
	doors int
	seats int
	sunroof bool
	fourByFour bool
	vehicle
}

type twoWheeler struct {
	maxSpeed int
	company string
	twoWheelerType string
	vehicle
}

func newFourWheeler(doors int, seats int, sunroof bool, fourByFour bool) *fourWheeler {
	myFourWheeler := fourWheeler{
		doors: doors,
		seats: seats,
		sunroof: sunroof,
		fourByFour: fourByFour,
		vehicle: vehicle{
			wheels: 4,
			brakes: true,
			headlights: 2,
		},
	}

	return &myFourWheeler
}

func main() {
	myFourWheeler := newFourWheeler(4, 5, true, false)
	fmt.Println(myFourWheeler)

	newVehicle := vehicle{
		wheels: 4,
		brakes: true,
		headlights: 2,
	}

	newFourWheeler1 := fourWheeler{
		doors: 2,
		seats: 2,
		sunroof: false,
		fourByFour: false,
		vehicle: newVehicle,
	}

	fmt.Println(newFourWheeler1)

	myTwoWheeler := twoWheeler{
		maxSpeed: 240,
		company: "BMW",
		twoWheelerType: "Bike",
		vehicle: vehicle{
			wheels: 2,
			brakes: true,
			headlights: 1,
		},
	}

	fmt.Println(myTwoWheeler)
}
