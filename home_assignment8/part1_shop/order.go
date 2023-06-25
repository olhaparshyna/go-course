package part1_shop

import (
	"context"
	"fmt"
	items2 "go-course/home_assignment8/part1_shop/items"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	Name     string
	Item     Item
	Quantity int
}

type Item interface {
	GetPrice() int
	GetName() string
}

const PeriodToWaitBetweenGenerateNewOrder = 2 * time.Second

func Generate(ctx context.Context, orders chan<- Order, wg *sync.WaitGroup) {
	names := []string{"John", "Mary", "Bob", "Alice", "Tom", "Kate"}

	defer wg.Done()

	ticker := time.NewTicker(PeriodToWaitBetweenGenerateNewOrder)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			var item Item
			switch rand.Intn(2) + 1 {
			case 1:
				item = items2.NewBook()
			case 2:
				item = items2.NewTable()
			}

			order := Order{
				Name:     names[rand.Intn(len(names))],
				Item:     item,
				Quantity: rand.Intn(5) + 1,
			}

			fmt.Printf("New order: %v\n", order)
			orders <- order
		}
	}
}

func ProcessOrder(ctx context.Context, orders <-chan Order, wg *sync.WaitGroup) {
	total := make(map[string]int)

	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Shop is closed")
			for itemName, quantity := range total {
				fmt.Printf("%s: %d\n", itemName, quantity)
			}
			return
		case order, ok := <-orders:
			if !ok {
				fmt.Println("No more orders")
				return
			}
			fmt.Printf("Process oreder: %v\n", order)
			total[order.Item.GetName()] += order.Quantity
		}
	}
}
