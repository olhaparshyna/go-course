package main

import (
	"fmt"
	"github.com/google/uuid"
	"go-course/home_assignment15part1/fileSystem"
)

type FileSub struct {
}

func (f FileSub) GetNotified(subject any) {
	fmt.Printf("FileSub got notified about %v\n", subject)
}

func (f FileSub) GetId() uuid.UUID {
	return uuid.New()
}

func main() {
	fileMonitoringService := fileSystem.NewService()

	fileMonitoringService.AddSubscriber(FileSub{})

	fileMonitoringService.FileMonitor()

	fileMonitoringService.Stop()
}
