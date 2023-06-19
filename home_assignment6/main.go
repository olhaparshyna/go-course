package main

import (
	"go-course/home_assignment6/post"
	parcels2 "go-course/home_assignment6/post/parcels"
)

func main() {
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

}
