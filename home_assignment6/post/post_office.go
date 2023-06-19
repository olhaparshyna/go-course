package post

type Parcel interface {
	Send()
	GetPostServiceType() string
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
