package parcels

import (
	"fmt"
	"go-course/home_assignment6/post"
)

type Envelop struct {
	From string
	To   string
}

func (e Envelop) GetTo() string {
	return e.To
}

func (e Envelop) GetPostServiceType() string {
	return post.SmallParcels
}

func (e Envelop) Send() {
	fmt.Printf("Send from: %s to %s via %s\n", e.From, e.To, e.GetPostServiceType())
}
