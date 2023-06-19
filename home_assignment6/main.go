package main

import (
	"fmt"
	"go-course/home_assignment6/post"
	parcels2 "go-course/home_assignment6/post/parcels"
	"go-course/home_assignment6/transport"
	"go-course/home_assignment6/transport/vehicles"
	"math/rand"
)

func main() {
	//post
	var parcels = []post.Parcel{
		parcels2.Box{
			Type: "small",
			From: "Lviv",
			To:   "Odessa",
		},
		parcels2.Envelop{
			From: "London",
			To:   "Paris",
		},
	}

	post.SortAndSend(parcels)

	//transport
	cars := vehicles.NewCars(4)
	planes := vehicles.NewPlanes(2)
	trains := vehicles.NewTraines(5)

	vehicles := make(map[string]transport.Vehicle)

	for i, car := range cars {
		i++
		vehicles[car.GetName()+fmt.Sprintf("%d", i)] = car
	}

	for i, plane := range planes {
		i++
		vehicles[plane.GetName()+fmt.Sprintf("%d", i)] = plane
	}

	for i, train := range trains {
		i++
		vehicles[train.GetName()+fmt.Sprintf("%d", i)] = train
	}

	route := transport.Route{}

	for _, vehicle := range vehicles {
		switch rand.Intn(2) + 1 {
		case 1:
			vehicle.ChangeSpeed(transport.SpeedFaster)
		case 2:
			vehicle.ChangeSpeed(transport.SpeedSlower)
		}
		route.AddVehicle(vehicle)
	}

	route.ShowRouteTrace()
}
