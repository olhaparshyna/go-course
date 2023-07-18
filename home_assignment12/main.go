package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// strategy
type TextProcessor interface {
	ProcessText(text string) int
}

type WordCounter struct{}

func (w WordCounter) ProcessText(text string) int {
	words := strings.Fields(text)
	return len(words)
}

type LineCounter struct{}

func (lc LineCounter) ProcessText(text string) int {
	lines := strings.Count(text, "\n") + 1
	return lines
}

// decorator
type TextModifier interface {
	ModifyText(text string) string
}

type BaseText struct{}

func (bt BaseText) ModifyText(text string) string {
	return text
}

type TextWithoutTags struct {
	parent TextModifier
}

func (t TextWithoutTags) ModifyText(text string) string {
	re := regexp.MustCompile("<[^>]+>")

	return t.parent.ModifyText(re.ReplaceAllString(text, ":replaced tag:"))
}

type TextUpperCase struct {
	parent TextModifier
}

func (t TextUpperCase) ModifyText(text string) string {
	return t.parent.ModifyText(strings.ToUpper(text))
}

func createFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, err
}

func putToFile(file *os.File, text string) error {
	_, err := file.WriteString(text)

	return err
}

func main() {
	contentByte, err := os.ReadFile("text.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	content := string(contentByte)

	var answer string

	fmt.Println("Should I count lines or words? Type \"l\" or \"w\"!")
	fmt.Scanln(&answer)

	if answer == "l" {
		lineCounter := LineCounter{}
		fmt.Println(lineCounter.ProcessText(content))
	}

	if answer == "w" {
		wordCounter := WordCounter{}
		fmt.Println(wordCounter.ProcessText(content))
	}

	var modifier TextModifier = BaseText{}

	fmt.Println("Remove html tags from text?")
	fmt.Scanln(&answer)

	if answer == "y" {
		modifier = TextWithoutTags{
			parent: modifier,
		}
	}

	fmt.Println("Make text uppercase?")
	fmt.Scanln(&answer)

	if answer == "y" {
		modifier = TextUpperCase{
			parent: modifier,
		}
	}

	content = modifier.ModifyText(content)

	fmt.Println(content)

	file, err := createFile("modified.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = putToFile(file, content)

	if err != nil {
		fmt.Println(err.Error())
	}
}
