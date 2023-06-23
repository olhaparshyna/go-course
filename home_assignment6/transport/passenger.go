package transport

import (
	"math/rand"
)

type Passenger struct {
	ID int
}

func GeneratePassengers(number int) []Passenger {
	var result []Passenger
	IDs := make(map[int]bool)
	for i := number; i > 0; i-- {
		id := rand.Int()
		for _, ok := IDs[id]; ok; {
			id = rand.Int()
		}
		IDs[id] = true
		result = append(result, Passenger{ID: id})
	}

	return result
}
