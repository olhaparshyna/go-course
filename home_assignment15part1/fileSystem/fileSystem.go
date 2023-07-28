package fileSystem

import (
	"fmt"
	"github.com/google/uuid"
	"io/fs"
	"path/filepath"
	"time"
)

type Subscriber interface {
	GetId() uuid.UUID
	GetNotified(event any)
}

type Service struct {
	subscribers chan Subscriber
	events      chan any
	stop        chan struct{}
}

func NewService() *Service {
	s := &Service{
		subscribers: make(chan Subscriber),
		events:      make(chan any),
		stop:        make(chan struct{}),
	}

	go s.processEvents()

	return s
}

func (s *Service) processEvents() {
	subscribers := make(map[uuid.UUID]Subscriber)

	for {
		select {
		case e := <-s.events:
			fmt.Println("Got event to process")
			for _, sub := range subscribers {
				sub.GetNotified(e)
			}
			fmt.Println("Finished event processing")
		case s := <-s.subscribers:
			subscribers[s.GetId()] = s
		case <-s.stop:
			fmt.Println("Got stop signal")
			s.stop <- struct{}{}
			return
		}
	}
}

func (s *Service) AddSubscriber(sub Subscriber) {
	s.subscribers <- sub
}

func (s *Service) FileMonitor() {
	directoryPath := "../home_assignment15part2"

	initialFiles := make(map[string]struct{})
	initialFiles = getFileStructure(directoryPath)

	for {
		currentFiles := getFileStructure(directoryPath)

		for path := range currentFiles {
			if _, ok := initialFiles[path]; !ok {
				event := fmt.Sprintf("New addition in %s: %s\n", directoryPath, path)
				s.events <- event
			}
		}

		for path := range initialFiles {
			if _, ok := currentFiles[path]; !ok {
				event := fmt.Sprintf("Deletion in %s: %s\n", directoryPath, path)
				s.events <- event
			}
		}

		initialFiles = currentFiles

		select {
		case <-s.stop:
			return
		default:
			time.Sleep(time.Second)
		}
	}
}

func (s *Service) Stop() {
	fmt.Println("Sending stop signal")
	s.stop <- struct{}{}
	<-s.stop
}

func getFileStructure(directoryPath string) map[string]struct{} {
	files := make(map[string]struct{})

	filepath.Walk(directoryPath, func(path string, info fs.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			files[path] = struct{}{}
		}
		return nil
	})

	return files
}
