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
	winNumber := 1
	numberOfPlayers := 2
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	roundsCh := make(chan int)
	answersCh := make(chan []int, 0)
	playersCh := make([]chan int, numberOfPlayers)
	gameOverCh := make(chan int)

	for i := range playersCh {
		playersCh[i] = make(chan int)
	}

	wg.Add(1)
	go generateRound(ctx, roundsCh, &wg)

	for i := 0; i < numberOfPlayers; i++ {
		wg.Add(1)
		go player(ctx, i, playersCh, roundsCh, &wg)
	}

	wg.Add(1)
	go recordAnswers(ctx, numberOfPlayers, playersCh, answersCh, roundsCh, &wg)

	wg.Add(1)
	go counter(ctx, roundsCh, answersCh, winNumber, &wg, cancel, numberOfPlayers, gameOverCh)

	<-gameOverCh
	cancel()
	fmt.Printf("GAME OVER. WinNumber was %d\n", winNumber)

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

func counter(ctx context.Context, roundsCh <-chan int, answersCh chan []int, winNumber int, wg *sync.WaitGroup, cancel context.CancelFunc, numberOfPlayers int, gameOverCh chan int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case <-roundsCh:
			for i := 0; i <= numberOfPlayers; i++ {
				answers := <-answersCh
				for _, num := range answers {
					if num == winNumber {
						gameOverCh <- 1
						return
					}
				}
			}
		}
	}
}

func player(ctx context.Context, id int, playersCh []chan int, roundsCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			round := <-roundsCh
			answer := rand.Intn(3) + 1
			fmt.Printf("Player %d in round %d, entered number %d: ", id, round, answer)
			playersCh[id] <- answer
		}
	}
}

func recordAnswers(ctx context.Context, numberOfPlayers int, playersCh []chan int, answersCh chan []int, roundsCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	answers := make([]int, numberOfPlayers)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			round := <-roundsCh

			for i := 0; i < numberOfPlayers; i++ {
				answer := <-playersCh[i]
				fmt.Printf("Player %d in round %d, entered number %d\n", i, round, answer)
				answers[i] = answer
			}

			answersCh <- answers
		}
	}
}
