package transport

import (
	"fmt"
)

type Vehicle interface {
	Stop
	Move()
	Stop()
	ChangeSpeed(speed string)
	GetName() string
}

const (
	SpeedSlower = "slower"
	SpeedFaster = "faster"
)

type Stop interface {
	Drop()
	PickUp()
}

type Route struct {
	Vehicles []Vehicle
}

func (r *Route) AddVehicle(vehicle Vehicle) {
	r.Vehicles = append(r.Vehicles, vehicle)
}

func (r *Route) ShowVehicles() {
	for _, v := range r.Vehicles {
		fmt.Printf("Please take %s\n", v.GetName())
	}
}

func (r *Route) ShowRouteTrace() {
	for _, v := range r.Vehicles {
		fmt.Printf("Please take %s\n", v.GetName())
		v.Stop()
		v.Drop()
		v.PickUp()
		fmt.Println("Let's go \n")
	}
}
