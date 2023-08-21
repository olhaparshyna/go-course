package main

import (
	"context"
	"encoding/json"
	"go-course/home_assignment18/consume"
	"log"
	"sync"
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

	var (
		smallMu  sync.Mutex
		middleMu sync.Mutex
		bigMu    sync.Mutex
	)

	small := make([]map[string]int, 0)
	middle := make([]map[string]int, 0)
	big := make([]map[string]int, 0)

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			log.Printf("small: %d, middle: %d, big: %d", len(small), len(middle), len(big))
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
				smallMu.Lock()
				small = append(small, res)
				smallMu.Unlock()
			case size >= 300 && size < 400:
				middleMu.Lock()
				middle = append(middle, res)
				middleMu.Unlock()
			case size >= 400:
				bigMu.Lock()
				big = append(big, res)
				bigMu.Unlock()
			}

			if err := d.Ack(false); err != nil {
				log.Println(err)
			}
		}
	}
}
