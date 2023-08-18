package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-course/home_assignment18/produce"
	"log"
	"math/rand"
)

type Orange struct {
	Size int `json:"orange"`
}

func CreateOranges(number int) []Orange {
	var oranges []Orange
	for i := 0; i < 30; i++ {
		randomSize := rand.Intn(300) + 200 // Випадкове число від 1 до 10

		orange := Orange{
			Size: randomSize,
		}

		oranges = append(oranges, orange)
	}

	return oranges
}

func main() {
	pp := produce.NewPool(
		context.Background(),
		1,
		`main`,
		`amqp://localhost`,
		`guest`,
		`guest`,
	)

	oranges := CreateOranges(5)
	for _, p := range pp.Producers() {
		for _, orange := range oranges {
			orangeJSON, err := json.Marshal(orange)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			if err = p.Push(
				context.Background(),
				`key-product`,
				orangeJSON,
			); err != nil {
				log.Println(err)
			}
		}
	}
}
