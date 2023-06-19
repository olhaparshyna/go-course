package vehicles

import (
	"fmt"
	"math/rand"
)

type Plane struct {
	Capacity    int
	Speed       int
	MaxCapacity int
}

func NewPlanes(number int) []Plane {
	planes := make([]Plane, number)
	index := 0
	for i := number; i > 0; i-- {
		plane := Plane{
			MaxCapacity: 100,
			Capacity:    0,
			Speed:       rand.Intn(900),
		}

		planes[index] = plane
		index++
	}

	return planes
}

func (p Plane) GetName() string {
	return "plane"
}

func (p Plane) Move() {
	fmt.Println("Let's go!")
}

func (p Plane) Stop() {
	fmt.Println("Let's stop and drop/pick up passangers")
}

func (p Plane) ChangeSpeed(speed string) {
	fmt.Printf("Sorry, %s has a fixed speed\n", p.GetName())
}

func (p Plane) Drop() {
	p.Capacity = 0
	fmt.Printf("All passangers left %s\n", p.GetName())
}

func (p Plane) PickUp() {
	passengers := rand.Intn(100)
	if p.Capacity+passengers > p.MaxCapacity {
		fmt.Printf("Sorry, %s is full", p.GetName())
	} else {
		p.Capacity += passengers
		fmt.Printf("%d passangers are on board\n", passengers)
	}
}
