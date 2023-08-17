// connection/channel handler to consume messages from monolith bus
package produce

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

// Connection handler
type Master struct {
	exName          string
	connection      *amqp091.Connection
	done            chan struct{}
	notifyConnClose chan *amqp091.Error
	isReady         bool
	ackTag          chan uint64
}

// creates a new consumer state instance, and automatically
// attempts to connect to the server.
func NewMaster(ctx context.Context, exName, busHost, busUser, busPass string) *Master {
	session := Master{
		exName: exName,
		done:   make(chan struct{}),
	}

	go session.handleReconnect(ctx, busHost, busUser, busPass)

	return &session
}

// will wait for a connection error on
// notifyConnClose, and then continuously attempt to reconnect.
func (s *Master) handleReconnect(ctx context.Context, busHost, busUser, busPass string) {
	for {
		s.isReady = false

		err := s.connect(ctx, busHost, busUser, busPass)
		if err != nil {
			select {
			case <-s.done:
				return
			case <-time.After(ReconnectDelay):
			}
			continue
		}

		select {
		case <-s.done:
			return
		case <-s.notifyConnClose:
			log.Println("Producer: connection closed. Reconnecting...")
		}
	}
}

// will create a new AMQP connection
func (s *Master) connect(_ context.Context, busHost, busUser, busPass string) error {
	conn, err := amqp091.DialConfig(busHost, amqp091.Config{
		SASL: []amqp091.Authentication{&amqp091.PlainAuth{busUser, busPass}},
	})

	if err != nil {
		return err
	}

	s.changeConnection(conn)
	s.isReady = true
	s.done = make(chan struct{})
	log.Println("Producer: CONNECTED")

	return nil
}

// takes a new connection to the queue,
// and updates the close listener to reflect this.
func (s *Master) changeConnection(connection *amqp091.Connection) {
	s.connection = connection
	s.notifyConnClose = make(chan *amqp091.Error, 1)
	s.connection.NotifyClose(s.notifyConnClose)
}

// will cleanly shutdown the channel and connection.
func (s *Master) Close() error {
	if !s.isReady {
		return errors.New(fmt.Sprintf("Producer: connection not ready while closing"))
	}
	err := s.connection.Close()
	if err != nil {
		return err
	}
	s.isReady = false

	return nil
}

func (s *Master) Complete() {
	close(s.done)
}
