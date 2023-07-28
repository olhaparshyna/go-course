package game

import (
	"sync"
)

type Observer interface {
	Notify(message string)
}

type GameRoom struct {
	number    int
	observers []Observer
	mutex     sync.Mutex
}

func (g *GameRoom) AddObserver(observer Observer) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.observers = append(g.observers, observer)
}

func (g *GameRoom) RemoveObserver(observer Observer) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	for i, o := range g.observers {
		if o == observer {
			g.observers = append(g.observers[:i], g.observers[i+1:]...)
			break
		}
	}
}

func (g *GameRoom) NotifyAll(message string, sender Observer) {
	for _, observer := range g.observers {
		if observer != sender {
			observer.Notify(message)
		}
	}
}

func NewGameRoom(number int) *GameRoom {
	return &GameRoom{
		number:    number,
		observers: make([]Observer, 0),
	}
}
