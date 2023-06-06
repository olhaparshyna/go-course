package main

import (
	"encoding/json"
	"fmt"
)

var (
	white  = "white"
	orange = "orange"
)

type Cat struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	Age   int    `json:"age"`
}

func (c Cat) voice() {
	fmt.Println("MEOW")
}

type Tiger struct {
	Cat *Cat `json:"cat"`
	//Cat
	Size string `json:"size"`
	Age  int    `json:"age"`
}

func (t Tiger) voice() {
	fmt.Println("ROAR")
}

func main() {
	cat := Cat{
		Name:  "Tom",
		Color: white,
		Age:   3,
	}

	tiger := Tiger{
		&cat,
		//cat,
		"big",
		10,
	}

	msg, _ := json.Marshal(cat)
	msg1, _ := json.Marshal(tiger)
	fmt.Println(string(msg), string(msg1))
	//fmt.Println(string(msg1))
	cat.voice()
	tiger.voice()
}
