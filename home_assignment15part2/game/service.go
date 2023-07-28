package game

import (
	"fmt"
)

type Player struct {
	name string
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) Invite(name string, room *GameRoom) *Player {
	return NewPlayer(name)
}

func (p *Player) Notify(message string) {
	fmt.Printf("%s got notification: %s\n", p.name, message)
}

func (p *Player) Move(room *GameRoom) {
	message := fmt.Sprintf("%s moved \n", p.name)
	room.NotifyAll(message, p)
}

func NewPlayer(name string) *Player {
	return &Player{
		name: name,
	}
}
