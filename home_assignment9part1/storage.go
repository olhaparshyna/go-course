package main

import "time"

type Storage struct {
	tasks []Task
}

func (s *Storage) GetTasks() []Task {
	return s.tasks
}

func (s *Storage) StoreTasks() {
	tasks := []Task{
		{
			ID:    1,
			Title: "Task 1",
			Date:  time.Date(2023, time.July, 29, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:    2,
			Title: "Task 2",
			Date:  time.Date(2023, time.July, 29, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:    3,
			Title: "Task 3",
			Date:  time.Date(2023, time.July, 28, 0, 0, 0, 0, time.UTC),
		},
	}

	s.tasks = tasks
}
