package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Name  string
	Value string
}

func (p *Player) PutValue(f *Field) {
	fmt.Println("Where do you want to put your value?")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	location := strings.Fields(input)

	row, _ := strconv.Atoi(location[0])
	number, _ := strconv.Atoi(location[1])
	if (*f.Line[row])[number] == "" {
		(*f.Line[row])[number] = p.Value
	}

	for _, f := range f.Line {
		fmt.Println(f)
	}
}

func SetValueToPlayers(p1, p2 *Player) {
	fmt.Println(`Player1 choose your value "x" or "0"`)
	var value string
	fmt.Scan(&value)
	p1.Value = value

	if value == "x" {
		p2.Value = "0"
	} else {
		p2.Value = "x"
	}
}

type Line map[int]string

func (l *Line) CheckResult() bool {
	if (*l)[1] != "" && (*l)[1] == (*l)[2] && (*l)[2] == (*l)[3] {
		return true
	}

	return false
}

type Field struct {
	Line map[int]*Line
}

func (f *Field) CheckResult() bool {
	for _, line := range f.Line {
		if line.CheckResult() {
			return true
		}
	}

	if (*f.Line[1])[1] != "" && (*f.Line[1])[1] == (*f.Line[2])[2] && (*f.Line[1])[1] == (*f.Line[3])[3] {
		return true
	}

	if (*f.Line[1])[3] != "" && (*f.Line[1])[3] == (*f.Line[2])[2] && (*f.Line[1])[3] == (*f.Line[3])[1] {
		return true
	}

	if (*f.Line[1])[1] != "" && (*f.Line[1])[1] == (*f.Line[2])[1] && (*f.Line[1])[1] == (*f.Line[3])[1] {
		return true
	}

	if (*f.Line[1])[3] != "" && (*f.Line[1])[3] == (*f.Line[2])[3] && (*f.Line[1])[3] == (*f.Line[3])[3] {
		return true
	}

	if (*f.Line[1])[2] != "" && (*f.Line[1])[2] == (*f.Line[2])[2] && (*f.Line[1])[2] == (*f.Line[3])[2] {
		return true
	}

	return false
}

func main() {
	p1 := Player{Name: "Player 1"}

	p2 := Player{Name: "Player 2"}

	SetValueToPlayers(&p1, &p2)

	l1 := Line{
		1: "",
		2: "",
		3: "",
	}

	l2 := Line{
		1: "",
		2: "",
		3: "",
	}

	l3 := Line{
		1: "",
		2: "",
		3: "",
	}

	field := Field{
		make(map[int]*Line, 3),
	}

	field.Line[1] = &l1
	field.Line[2] = &l2
	field.Line[3] = &l3

	fmt.Println(field)
	for _, f := range field.Line {
		fmt.Println(f)
	}

	for !field.CheckResult() {
		fmt.Println("Player1 your turn")
		p1.PutValue(&field)
		fmt.Println("Player2 your turn")
		p2.PutValue(&field)
	}

	fmt.Println("Game over")
}
