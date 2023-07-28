package main

import "go-course/home_assignment15part2/game"

func main() {
	gameRoom := game.NewGameRoom(1)

	player1 := game.NewPlayer("Player1")
	player2 := player1.Invite("Player2", gameRoom)
	player3 := player1.Invite("Player3", gameRoom)

	gameRoom.AddObserver(player1)
	gameRoom.AddObserver(player2)
	gameRoom.AddObserver(player3)

	player2.Move(gameRoom)
	player3.Move(gameRoom)
}
