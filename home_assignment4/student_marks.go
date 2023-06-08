package main

import "fmt"

// 2. Завдання про студентські оцінки.
// Створити slice float з оцінками студентів з певного предмету.
// Написати функцію, яка приймає на вхід slice оцінок та повертає середню оцінку з цього предмету.
// Використовуючи цю функцію, обчислити середній бал з предмету та вивести його на екран.
func average(marks []float64) {
	var total float64
	for _, mark := range marks {
		total += mark
	}

	fmt.Println(total / float64(len(marks)))
}

func main() {
	marks := []float64{12.5, 8.4, 6.8, 9.2}
	average(marks)
}
