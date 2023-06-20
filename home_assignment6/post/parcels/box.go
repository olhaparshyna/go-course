package parcels

import (
	"fmt"
	"go-course/home_assignment6/post"
)

type Box struct {
	Type string
	From string
	To   string
}

func (b Box) GetTo() string {
	return b.To
}

func (b Box) GetPostServiceType() string {
	if b.Type == post.SmallSize {
		fmt.Printf("This box is %s and could be send as a small parcel. ", b.Type)
		return post.SmallParcels
	}

	return post.BigParcels
}

func (b Box) Send() {
	fmt.Printf("Send from: %s to %s via %s\n", b.From, b.To, b.GetPostServiceType())
}
