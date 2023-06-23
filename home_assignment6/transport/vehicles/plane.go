package vehicles

import (
	"fmt"
	"go-course/home_assignment6/transport"
	"math/rand"
)

type Plane struct {
	Capacity     int
	Speed        int
	MaxCapacity  int
	Registration map[int]transport.Passenger
}

func (p *Plane) CheckRegistration(person transport.Passenger) bool {
	_, ok := p.Registration[person.ID]

	if ok {
		return true
	}

	return false
}

func (p *Plane) RegisterPassanger(person transport.Passenger) {
	if p.Registration == nil {
		p.Registration = make(map[int]transport.Passenger)
	}

	p.Registration[person.ID] = person
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

func (p Plane) Stop() bool {
	if p.Speed > 500 {
		fmt.Println("No-no-no plane is flying to fast to stop")
		return false
	}
	fmt.Println("Let's stop and drop/pick up passangers")
	return true
}

func (p Plane) ChangeSpeed(speed string) {
	fmt.Printf("Sorry, %s has a fixed speed\n", p.GetName())
}

func (p Plane) Drop() {
	p.Capacity = 0
	fmt.Printf("All passangers left %s\n", p.GetName())
}

func (p Plane) PickUp() {
	passangersForPlane := transport.GeneratePassengers(100)
	passangersToRegister := passangersForPlane[:95]

	for _, passenger := range passangersToRegister {
		p.RegisterPassanger(passenger)
	}

	onBoard := 0
	for _, passenger := range passangersForPlane {
		result := p.CheckRegistration(passenger)

		if result {
			if p.Capacity >= p.MaxCapacity {
				fmt.Printf("Sorry, %s is full", p.GetName())
			} else {
				p.Capacity++
				onBoard++
			}
		} else {
			fmt.Printf("Passenger with %d is not registered\n", passenger.ID)
		}
	}

	fmt.Printf("%d passangers are on board\n", onBoard)
}
