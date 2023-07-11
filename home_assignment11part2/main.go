package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	filePath := "1689007676028_text.txt"
	paragraphs, err := readParagraphsFromFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, paragraph := range paragraphs {
		//find all words beginning with vowels
		re := regexp.MustCompile(`(?:^|\s)([аеиоуыэюя][а-я]*)`)
		result := re.FindAllString(paragraph, -1)

		for _, number := range result {
			fmt.Println(number)

		}
	}
}

func readParagraphsFromFile(filePath string) ([]string, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	content := string(fileData)

	paragraphs := strings.Split(content, "\n\n")

	return paragraphs, err
}
