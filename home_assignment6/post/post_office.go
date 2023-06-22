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

	sorted := make(map[string][]Parcel, 0)

	for _, parcel := range parcels {
		sorted[parcel.GetTo()] = append(sorted[parcel.GetTo()], parcel)
	}

	fmt.Println(sorted)
}
