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
	if f.Line[row][number] == "" {
		f.Line[row][number] = p.Value
	}

	for _, f := range f.Line {
		fmt.Println(f)
	}
}

func SetValueToPlayers(p1, p2 *Player) {
	fmt.Println(`Player1 choose your value "x" or "0"`)
	var value string
	fmt.Scan(&value)

	switch value {
	case "x":
		p1.Value = "x"
		p2.Value = "0"
	case "0":
		p1.Value = "0"
		p2.Value = "x"
	default:
		p1.Value = "x"
		p2.Value = "0"
		fmt.Printf("Ok, I will make a choice for you! %s you get - %s\n %s you get - %s\n",
			p1.Name, p1.Value, p2.Name, p2.Value)
	}
}

type Line map[int]string

func (line Line) SetValue(value string) {
	for i := 1; i <= 3; i++ {
		line[i] = value
	}
}

func (l Line) CheckResult() bool {
	if l[1] != "" && l[1] == l[2] && l[2] == l[3] {
		return true
	}

	return false
}

type Field struct {
	Line map[int]Line
}

func (f *Field) CheckResult() bool {
	for _, line := range f.Line {
		if line.CheckResult() {
			return true
		}
	}

	if f.Line[1][1] != "" && f.Line[1][1] == f.Line[2][2] && f.Line[1][1] == f.Line[3][3] {
		return true
	}

	if f.Line[1][3] != "" && f.Line[1][3] == f.Line[2][2] && f.Line[1][3] == f.Line[3][1] {
		return true
	}

	if f.Line[1][1] != "" && f.Line[1][1] == f.Line[2][1] && f.Line[1][1] == f.Line[3][1] {
		return true
	}

	if f.Line[1][3] != "" && f.Line[1][3] == f.Line[2][3] && f.Line[1][3] == f.Line[3][3] {
		return true
	}

	if f.Line[1][2] != "" && f.Line[1][2] == f.Line[2][2] && f.Line[1][2] == f.Line[3][2] {
		return true
	}

	return false
}

func main() {
	p1 := Player{Name: "Player 1"}

	p2 := Player{Name: "Player 2"}

	SetValueToPlayers(&p1, &p2)

	field := Field{
		map[int]Line{
			1: {},
			2: {},
			3: {},
		},
	}

	for _, f := range field.Line {
		f.SetValue("")
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
