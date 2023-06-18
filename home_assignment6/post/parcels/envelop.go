package parcels

import "fmt"

type Envelop struct {
	Type string
	From string
	To   string
}

func (e Envelop) GetPostServiceType() string {
	return "small"
}

func (e Envelop) Send(service string) {
	fmt.Printf("Send from: %s to %s via Service %s", e.From, e.To, service)
}
