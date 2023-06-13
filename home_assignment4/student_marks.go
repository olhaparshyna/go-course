package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 2. Завдання про студентські оцінки.
// Створити slice float з оцінками студентів з певного предмету.
// Написати функцію, яка приймає на вхід slice оцінок та повертає середню оцінку з цього предмету.
// Використовуючи цю функцію, обчислити середній бал з предмету та вивести його на екран.

func average(marks []float64, subject string) {
	var total float64
	for _, mark := range marks {
		total += mark
	}

	fmt.Printf("%s: %.2f\n", subject, total/float64(len(marks)))
}

func main() {
	fmt.Println("Enter subject")
	var subject string
	fmt.Scan(&subject)

	fmt.Println("Enter marks")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numbers := strings.Fields(input)
	fmt.Println(numbers)

	marks := make([]float64, 0)
	for _, number := range numbers {
		mark, _ := strconv.ParseFloat(number, 64)
		marks = append(marks, mark)
	}

	average(marks, subject)
}
