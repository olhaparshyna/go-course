package main

import (
	"fmt"
)

//Розробити програму «Зоопарк».
// Завдання: 5 чи більше звірів повтікали,
//наглядач повинен їх зібрати.
// Кожну сутність (наглядач, звір, клітка тощо)
//представляти окремою структурою (zookeeper, animal, cage).
// Користуємось ембдінгом і методами.

var (
	big   = "big"
	small = "small"
)

type Zookeeper struct {
	name string
	age  int
}

func (zk Zookeeper) Info() string {
	return "Zookeeper's name is " + zk.name + fmt.Sprintf("\nHe is %d\n", zk.age)

}

func (zk Zookeeper) String() string {
	return zk.Info()
}

func (zk *Zookeeper) ChangeHome(t *Tiger, l *Lion, li *Lioness, d *Dolphin, b *Bear) {
	t.ChangeHome()
	l.ChangeHome()
	li.ChangeHome()
	d.ChangeHome()
	b.ChangeHome()

	fmt.Println("Check animals")
}

type Animal struct {
	Name string `json:"name"`
	Home string `json:"home"`
}

type Cat struct {
	Animal
	Paws int  `json:"paws"`
	Mane bool `json:"mane"`
}

type Tiger struct {
	Cat
	Voice string `json:"voice"`
}

func (t *Tiger) SetData() {
	t.Name = "TIGER"
}

func (t *Tiger) ChangeHome() {
	if t.Home == "lake" {
		t.Home = "cage"
	} else {
		t.Home = "lake"
	}
}

type Lion struct {
	Cat
}

func (lion *Lion) SetData() {
	lion.Name = "LION"
}

func (lion *Lion) ChangeHome() {
	if lion.Home == "lake" {
		lion.Home = "cage"
	} else {
		lion.Home = "lake"
	}
}

type Lioness struct {
	Cat
}

func (li *Lioness) SetData() {
	li.Name = "LIONESS"
	li.Mane = false
}

func (li *Lioness) ChangeHome() {
	if li.Home == "lake" {
		li.Home = "cage"
	} else {
		li.Home = "lake"
	}
}

type Bear struct {
	Animal
	Color string `json:"color"`
}

func (b *Bear) ChangeHome() {
	if b.Home == "lake" {
		b.Home = "cage"
	} else {
		b.Home = "lake"
	}
}

type Dolphin struct {
	Animal
	Ability string `json:"ability"`
}

func (d *Dolphin) ChangeHome() {
	if d.Home == "lake" {
		d.Home = "cage"
	} else {
		d.Home = "lake"
	}
}

func main() {
	zk := Zookeeper{
		name: "Greg",
		age:  40,
	}

	cat := Cat{
		Animal: Animal{
			Home: "cage",
		},
		Paws: 4,
		Mane: true,
	}

	tiger := Tiger{
		Cat:   cat,
		Voice: "roar",
	}

	lion := Lion{
		Cat: cat,
	}

	lioness := Lioness{
		Cat: cat,
	}

	bear := Bear{
		Animal: Animal{
			Name: "BEAR",
			Home: "cage",
		},
		Color: "brown",
	}

	dolphin := Dolphin{
		Animal: Animal{
			Name: "DOLPHIN",
			Home: "lake",
		},
		Ability: "swim",
	}

	tiger.SetData()
	lion.SetData()
	lioness.SetData()

	fmt.Println(zk)

	zk.ChangeHome(&tiger, &lion, &lioness, &dolphin, &bear)
	fmt.Println("Animals might got lost")
	fmt.Println(tiger, lion, lioness, dolphin, bear)
	zk.ChangeHome(&tiger, &lion, &lioness, &dolphin, &bear)
	fmt.Println(tiger, lion, lioness, dolphin, bear)
	fmt.Println("Animals are at home")
}
