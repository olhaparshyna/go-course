package main

//2. Створити програму для симуляції групи людей, які одночасно грають в ігри на великому екрані.
//Програма має використовувати горутину-генератор,
//який кожні 10 секунд генерує новий ігровий раунд та відправляє його до горутин-гравців через канал.
//Гравці отримують новий ігровий раунд та вводять свої відповіді через окремий канал.
//Далі горутина-лічильник перевіряє правильність відповідей та
//повертає результат у головну горутину через окремий канал.
//Якщо в програмі виникає помилка або користувач перериває програму,
//то програма має коректно завершувати роботу з використанням контексту.

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	winNumber := 12
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	roundsCh := make(chan int)
	answersCh := make(chan int)

	wg.Add(2)
	go generateRound(ctx, roundsCh, &wg)

	go counter(ctx, roundsCh, answersCh, winNumber, &wg, cancel)

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go player(ctx, i, answersCh, roundsCh, &wg)
	}

	wg.Wait()

	fmt.Println("Press Enter")
	fmt.Scanln()
}

func generateRound(ctx context.Context, roundsCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker := time.NewTicker(10 * time.Second)
	round := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			round += 1
			fmt.Printf("New round: %d\n", round)
			roundsCh <- round
		}
	}
}

func counter(ctx context.Context, roundsCh <-chan int, answersCh chan int, winNumber int, wg *sync.WaitGroup, cancel context.CancelFunc) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case <-roundsCh:
			for i := 1; i <= 5; i++ {
				answer := <-answersCh
				if answer == winNumber {
					fmt.Printf("GAME OVER. WinNumber was %d\n", winNumber)
					cancel()
				}
			}
		}
	}
}

func player(ctx context.Context, id int, answersCh chan int, roundsCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			round := <-roundsCh
			answer := rand.Intn(15) + 1
			fmt.Printf("Player %d in round %d, entered number %d: ", id, round, answer)
			answersCh <- answer
		}
	}
}
