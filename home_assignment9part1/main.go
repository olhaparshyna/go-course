package main

//1. Створити API для перегляду списку прав.
//Користувач повинен мати можливість переглядати список завдань за конкретну дату.
//Додаткові вимоги:
//• список завдань має бути збережений в оперативній пам'яті та бути доступним під час кожного запиту;
//• отримання списку завдань має здійснюватись методом GET на адресі "/tasks".

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
	"time"
)

func main() {
	r := mux.NewRouter()
	s := &Storage{}
	s.StoreTasks()

	fmt.Println(s.GetTasks())

	t := &taskHandler{s: s}

	r.HandleFunc("/tasks", t.getTasksByDate).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}

var taskCache sync.Map

type Task struct {
	ID    int       `json:"id"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
}

type taskHandler struct {
	s *Storage
}

func (t *taskHandler) getTasksByDate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling request")

	dateStr := r.URL.Query().Get("date")
	fmt.Println(dateStr)

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format - YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	date.Date()
	fmt.Println(date)

	if cachedTasks, ok := taskCache.Load(date); ok {
		response, err := json.Marshal(cachedTasks)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Write(response)
		return
	}

	filteredTasks := []Task{}
	for _, task := range t.s.GetTasks() {
		if task.Date.Equal(date) {
			filteredTasks = append(filteredTasks, task)
		}
	}

	taskCache.Store(date, filteredTasks)
	response, err := json.Marshal(filteredTasks)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
