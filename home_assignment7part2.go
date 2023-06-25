package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//Яка створює 2 горутини.
//Перша горутина генерує випадкові числа в заданому діапазоні й надсилає їх через канал у другу горутину.
//Друга горутина отримує випадкові числа і знаходить найбільше й найменше число,
//після чого надсилає його назад у першу горутину через канал.
//Перша горутина виводить найбільше й найменше числа на екран.

func main() {
	randomNumbersCh := make(chan int)
	minMaxCh := make(chan int)
	finish := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			number := rand.Intn(50) + 51
			randomNumbersCh <- number
			fmt.Printf("number %d added \n", number)
		}
		close(randomNumbersCh)

		fmt.Printf("min %d\n", <-minMaxCh)
		fmt.Printf("max %d\n", <-minMaxCh)

		finish <- "finish"
	}()

	go func() {
		numbers := make([]int, 0)
		for number := range randomNumbersCh {
			numbers = append(numbers, number)
		}

		sort.Ints(numbers)

		min := numbers[0]
		minMaxCh <- min
		max := numbers[len(numbers)-1]
		minMaxCh <- max

		close(minMaxCh)
	}()

	message := <-finish
	fmt.Println(message)
}
