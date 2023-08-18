package main

import (
	"context"
	"encoding/json"
	"go-course/home_assignment18/consume"
	"log"
	"time"
)

func main() {
	cp := consume.NewPool(
		context.Background(),
		1,
		2,
		`q1`,
		`amqp://localhost`,
		`guest`,
		`guest`,
	)

	small := make([]map[string]int, 0)
	middle := make([]map[string]int, 0)
	big := make([]map[string]int, 0)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			log.Printf("small: %v, middle: %v, big: %v", small, middle, big)
		}
	}()

	for _, c := range cp.Consumers() {
		if err := c.InitStream(context.Background()); err != nil {
			log.Println(err)
		}

		for {
			for {
				if !c.IsDeliveryReady {
					log.Println(`Waiting...`)
					time.Sleep(consume.ReconnectDelay)
				} else {
					break
				}
			}

			d := <-c.GetStream()

			res := make(map[string]int)
			if err := json.Unmarshal(d.Body, &res); err != nil {
				log.Println(err)
			}

			switch size := res["orange"]; {
			case size < 300:
				small = append(small, res)
			case size >= 300 && size < 400:
				middle = append(middle, res)
			case size >= 400:
				big = append(big, res)
			}

			if err := d.Ack(false); err != nil {
				log.Println(err)
			}
		}
	}
}
