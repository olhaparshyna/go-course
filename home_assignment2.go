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

func (zk *Zookeeper) ChangeHome(t *Tiger, l *Lion, li *Lioness, d *Dolphin, b *Bear, lake Lake, c Cage) {
	t.ChangeHome(lake, c)
	l.ChangeHome(lake, c)
	li.ChangeHome(lake, c)
	d.ChangeHome(lake, c)
	b.ChangeHome(lake, c)

	fmt.Println("Check animals")
}

type Cage struct {
	Size string `json:"size"`
	Name string `json:"name"`
}

type Lake struct {
	Size  string  `json:"size"`
	Depth float64 `json:"depth"`
	Name  string  `json:"name"`
}

type Animal struct {
	Name string      `json:"name"`
	Home interface{} `json:"home"`
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

func (t *Tiger) ChangeHome(l Lake, c Cage) {
	if t.Home == l {
		t.Home = c
	} else if t.Home == c {
		t.Home = l
	}
}

type Lion struct {
	Cat
}

func (lion *Lion) SetData() {
	lion.Name = "LION"
}

func (lion *Lion) ChangeHome(l Lake, c Cage) {
	if lion.Home == l {
		lion.Home = c
	} else if lion.Home == c {
		lion.Home = l
	}
}

type Lioness struct {
	Cat
}

func (li *Lioness) SetData() {
	li.Name = "LIONESS"
	li.Mane = false
}

func (li *Lioness) ChangeHome(l Lake, c Cage) {
	if li.Home == l {
		li.Home = c
	} else if li.Home == c {
		li.Home = l
	}
}

type Bear struct {
	Animal
	Color string `json:"color"`
}

func (b *Bear) ChangeHome(l Lake, c Cage) {
	if b.Home == l {
		b.Home = c
	} else if b.Home == c {
		b.Home = l
	}
}

type Dolphin struct {
	Animal
	Ability string `json:"ability"`
}

func (d *Dolphin) ChangeHome(l Lake, c Cage) {
	if d.Home == l {
		d.Home = c
	} else if d.Home == c {
		d.Home = l
	}
}

func main() {
	zk := Zookeeper{
		name: "Greg",
		age:  40,
	}

	cage := Cage{
		Size: big,
		Name: "cage",
	}

	lake := Lake{
		Size:  small,
		Depth: 5.6,
		Name:  "lake",
	}

	cat := Cat{
		Animal: Animal{
			Home: cage,
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
			Home: cage,
		},
		Color: "brown",
	}

	dolphin := Dolphin{
		Animal: Animal{
			Name: "DOLPHIN",
			Home: lake,
		},
		Ability: "swim",
	}

	tiger.SetData()
	lion.SetData()
	lioness.SetData()

	fmt.Println(zk)

	zk.ChangeHome(&tiger, &lion, &lioness, &dolphin, &bear, lake, cage)
	fmt.Println("Animals might got lost")
	fmt.Println(tiger, lion, lioness, dolphin, bear)
	zk.ChangeHome(&tiger, &lion, &lioness, &dolphin, &bear, lake, cage)
	fmt.Println(tiger, lion, lioness, dolphin, bear)
	fmt.Println("Animals are at home")
}
