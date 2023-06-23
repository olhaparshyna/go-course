package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Яка створює 3 горутини.
//	Перша горутина генерує випадкові числа й надсилає їх через канал у другу горутину.
//	Друга горутина отримує випадкові числа та знаходить їх середнє значення, п
//ісля чого надсилає його в третю горутину через канал. Третя горутина виводить середнє значення на екран.

func main() {
	var randomNumberStorage chan int
	randomNumberStorage = make(chan int)

	go func() {
		for i := rand.Intn(100) + 1; i > 0; i-- {
			number := rand.Intn(100)
			randomNumberStorage <- number
			fmt.Printf("number %d added \n", number)
		}

		close(randomNumberStorage)
	}()

	var avrgNumberStorage chan int
	avrgNumberStorage = make(chan int)

	go func() {
		sum := 0
		count := 0

		for number := range randomNumberStorage {
			fmt.Println(number)
			sum += number
			fmt.Println(sum)
			count++
		}

		avrgNumberStorage <- sum / count
		fmt.Printf("we have %d numbers\n", count)
		fmt.Printf("Sum is %d\n", sum)
	}()

	go func() {
		result := <-avrgNumberStorage
		fmt.Println(result)
	}()

	time.Sleep(2 * time.Second)
}
