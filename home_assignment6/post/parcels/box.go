package parcels

import "fmt"

type Box struct {
	Type string
	From string
	To   string
}

func (b Box) GetPostServiceType() string {
	return "big"
}

func (b Box) Send(service string) {
	fmt.Printf("Send from: %s to %s via Service %s", b.From, b.To, service)
}
