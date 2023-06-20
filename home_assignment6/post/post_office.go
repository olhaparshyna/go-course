package post

import "fmt"

type Parcel interface {
	Send()
	GetPostServiceType() string
	GetTo() string
}

const (
	BigParcels   = "service for big parcels"
	SmallParcels = "service for small parcels"
	BigSize      = "big"
	SmallSize    = "small"
)

func SortAndSend(parcels []Parcel) {
	for _, p := range parcels {
		p.Send()
	}
}

func SortByDestination(parcels []Parcel) {
	//unique cities
	destinations := make(map[string]bool)
	for _, p := range parcels {
		if _, ok := destinations[p.GetTo()]; !ok {
			destinations[p.GetTo()] = true
		}
	}

	destinationsSlice := make([]string, 0, len(destinations))

	for c := range destinations {
		destinationsSlice = append(destinationsSlice, c)
	}

	sorted := make(map[string][]Parcel, 0)

	for _, city := range destinationsSlice {
		parcelsToCity := make([]Parcel, 0)
		for _, parcel := range parcels {
			if parcel.GetTo() == city {
				parcelsToCity = append(parcelsToCity, parcel)
			}
		}
		sorted[city] = parcelsToCity
	}

	fmt.Println(sorted)
}
