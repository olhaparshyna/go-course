package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//1. Завдання про текстовий редактор.
//Створити slice string з текстом, який користувач вводить у текстовий редактор.
//Написати функцію, яка приймає на вхід рядок та знаходить у текстовому редакторі всі рядки, які містять цей рядок.
//Використовуючи цю функцію, додати можливість пошуку тексту в текстовому редакторі
//та вивести на екран усі відповідні результати.

func Search(text []string) []string {
	fmt.Println("Enter search")

	var search string
	fmt.Scan(&search)

	result := make([]string, 0)
	for _, part := range text {
		if strings.Contains(part, search) {
			result = append(result, part)
		}
	}

	return result
}

func main() {
	fmt.Println("Enter text")

	text := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		text = append(text, line)
	}

	fmt.Println(Search(text))
}
