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
	Customer string
	Price    int
	ItemName string
	Quantity int
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
			quantity := rand.Intn(5) + 1
			var name string
			var price int

			switch rand.Intn(2) + 1 {
			case 1:
				item := items2.NewBook()
				name = item.GetName()
				price = item.GetPrice() * quantity
			case 2:
				item := items2.NewTable()
				name = item.GetName()
				price = item.GetPrice() * quantity
			}

			order := Order{
				Customer: names[rand.Intn(len(names))],
				Price:    price,
				ItemName: name,
				Quantity: quantity,
			}

			fmt.Printf("New order: %v\n", order)
			orders <- order
		}
	}
}

func ProcessOrder(ctx context.Context, orders <-chan Order, wg *sync.WaitGroup) {
	total := make(map[string]int)
	var totalPrice int

	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Shop is closed")
			for itemName, quantity := range total {
				fmt.Printf("%s: %d\n", itemName, quantity)
			}
			fmt.Printf("Total price: %d\n", totalPrice)
			return
		case order, ok := <-orders:
			if !ok {
				fmt.Println("No more orders")
				return
			}
			fmt.Printf("Process oreder: %v\n", order)
			total[order.ItemName] += order.Quantity
			totalPrice += order.Price
		}
	}
}
