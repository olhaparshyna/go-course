package post

type Parcel interface {
	Send(PostService string)
	GetPostServiceType() string
}

type PostServices map[string]string

func SortAndSend(parcels []Parcel, postServices PostServices) {
	for _, p := range parcels {
		if "big" == p.GetPostServiceType() {
			p.Send(postServices["big"])
		} else {
			p.Send(postServices["small"])
		}
	}
}
