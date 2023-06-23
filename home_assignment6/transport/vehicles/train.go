package vehicles

import (
	"fmt"
	"go-course/home_assignment6/transport"
	"math/rand"
)

type Train struct {
	MaxCapacity int
	Capacity    int
	Speed       int
}

func NewTrains(number int) []Train {
	trains := make([]Train, number)
	index := 0
	for i := number; i > 0; i-- {
		train := Train{
			MaxCapacity: 800,
			Capacity:    0,
			Speed:       rand.Intn(200),
		}

		trains[index] = train
		index++
	}

	return trains
}

func (t Train) GetName() string {
	return "train"
}

func (t Train) Move() {
	fmt.Println("Let's go!")
}

func (t Train) Stop() bool {
	fmt.Println("Let's stop and drop/pick up passangers")
	return true
}

func (t Train) ChangeSpeed(speed string) {
	fmt.Printf("%s's speed was %d! ", t.GetName(), t.Speed)
	if speed == transport.SpeedFaster {
		t.Speed++
	}

	if speed == transport.SpeedSlower {
		if t.Speed-1 == 0 {
			t.Speed--
			t.Stop()
		} else if t.Speed-1 < 0 {
			fmt.Printf("The %s is alreade stopped", t.GetName())
		} else {
			t.Speed--
		}
	}
	fmt.Printf("now %d\n", t.Speed)
}

func (t Train) Drop() {
	if t.Capacity-1 < 0 {
		fmt.Printf("All passangers left %s\n", t.GetName())
	} else {
		t.Capacity--
		fmt.Printf("1 passanger left %s\n", t.GetName())
	}
}

func (t Train) PickUp() {
	passengers := rand.Intn(800) + 1
	if t.Capacity+passengers > t.MaxCapacity {
		fmt.Printf("Sorry, %s is full", t.GetName())
	} else {
		t.Capacity += passengers
		fmt.Printf("%d passanger get in to the %s\n", passengers, t.GetName())
	}
}
