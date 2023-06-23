package vehicles

import (
	"fmt"
	"go-course/home_assignment6/transport"
	"math/rand"
)

type Car struct {
	MaxCapacity int
	Capacity    int
	Speed       int
}

func NewCars(number int) []Car {
	cars := make([]Car, number)
	index := 0
	for i := number; i > 0; i-- {
		car := Car{
			MaxCapacity: 4,
			Capacity:    0,
			Speed:       rand.Intn(100),
		}

		cars[index] = car
		index++
	}

	return cars
}

func (c Car) GetName() string {
	return "car"
}

func (c Car) Stop() bool {
	fmt.Println("Let's stop and drop/pick up passangers")
	return true
}

func (c Car) Move() {
	fmt.Println("Let's go!")
}

func (c Car) ChangeSpeed(speed string) {
	fmt.Printf("%s's speed was %d! ", c.GetName(), c.Speed)
	if speed == transport.SpeedFaster {
		c.Speed++
	}

	if speed == transport.SpeedSlower {
		if c.Speed-1 == 0 {
			c.Speed--
		} else if c.Speed-1 < 0 {
			fmt.Printf("The %s is already stopped\n", c.GetName())
			c.Drop()
			c.PickUp()
		} else {
			c.Speed--
		}
	}
	fmt.Printf("now %d\n", c.Speed)
}

func (c Car) Drop() {
	if c.Capacity-1 < 0 {
		fmt.Printf("All passangers left %s\n", c.GetName())
	} else {
		c.Capacity--
		fmt.Printf("%d passangers remain %s\n", c.Capacity)
	}
}

func (c Car) PickUp() {
	passengers := rand.Intn(4) + 1
	if c.Capacity+passengers > c.MaxCapacity {
		fmt.Printf("Sorry, %s is full", c.GetName())
	} else {
		c.Capacity += passengers
		fmt.Printf("%d passanger get in to the %s\n", passengers, c.GetName())
	}
}
