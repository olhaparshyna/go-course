package main

import (
	"go-course/home_assignment6/post"
	parcels2 "go-course/home_assignment6/post/parcels"
)

func main() {
	postServices := post.PostServices{
		"big":   "for big parcels",
		"small": "for small parcels",
	}

	var parcels = []post.Parcel{
		parcels2.Box{
			Type: "big",
			From: "Lviv",
			To:   "Odessa",
		},
		parcels2.Envelop{
			Type: "small",
			From: "London",
			To:   "Paris",
		},
	}

	post.SortAndSend(parcels, postServices)

}
