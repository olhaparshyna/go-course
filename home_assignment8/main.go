package main

import (
	"context"
	"flag"
	"fmt"
	"go-course/home_assignment8/part1_shop"
	"sync"
	"time"
)

func main() {
	var openTime int
	flag.IntVar(&openTime, "openTime", 10, "How long should shop work")
	flag.Parse()

	fmt.Println(openTime)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(openTime)*time.Second)
	defer cancel()

	orders := make(chan part1_shop.Order)
	var wg sync.WaitGroup

	wg.Add(1)
	go part1_shop.Generate(ctx, orders, &wg)

	wg.Add(1)
	go part1_shop.ProcessOrder(ctx, orders, &wg)

	wg.Wait()
}
