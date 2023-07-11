package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	content, err := getFileContent("1689007675141_numbers.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	re := regexp.MustCompile(`(\(\d{3}\)\s?|\d{3}[-.\s]?)\d{3}[-.\s]?\d{4}`)
	result := re.FindAllString(content, -1)

	for _, number := range result {
		fmt.Println(number)

	}
}

func getFileContent(filePath string) (content string, err error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	content = string(fileContent)
	return content, nil
}
