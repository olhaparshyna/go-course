package main

import (
	"encoding/json"
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
	size string `json:"size"`
	name string `json:"name"`
}

type Lake struct {
	size  string  `json:"size"`
	depth float64 `json:"depth"`
	name  string  `json:"name"`
}

type Animal struct {
	name      string      `json:"name"`
	neighbour interface{} `json:"neighbour"`
	home      interface{} `json:"home"`
}

type Cat struct {
	Animal
	paws int  `json:"paws"`
	mane bool `json:"mane"`
}

type Tiger struct {
	Cat
	voice string `json:"voice"`
	//neighbour Lion
}

func (t *Tiger) SetData(l Lion) {
	t.name = "TIGER"
	t.neighbour = l
}

func (t *Tiger) ChangeHome(l Lake, c Cage) {
	if t.home == l {
		t.home = c
	} else if t.home == c {
		t.home = l
	}
}

type Lion struct {
	Cat
	//neighbour Tiger
}

func (lion *Lion) SetData(li Lioness) {
	lion.name = "LION"
	lion.neighbour = li
}

func (lion *Lion) ChangeHome(l Lake, c Cage) {
	if lion.home == l {
		lion.home = c
	} else if lion.home == c {
		lion.home = l
	}
}

type Lioness struct {
	Cat
	//neighbour Lion
}

func (li *Lioness) SetData(l Lion) {
	li.name = "LIONESS"
	li.neighbour = l
	li.mane = false
}

func (li *Lioness) ChangeHome(l Lake, c Cage) {
	if li.home == l {
		li.home = c
	} else if li.home == c {
		li.home = l
	}
}

type Bear struct {
	Animal
	color string `json:"color"`
	//neighbour Lion
}

func (b *Bear) ChangeHome(l Lake, c Cage) {
	if b.home == l {
		b.home = c
	} else if b.home == c {
		b.home = l
	}
}

type Dolphin struct {
	Animal
	ability string `json:"ability"`
}

func (d *Dolphin) ChangeHome(l Lake, c Cage) {
	if d.home == l {
		d.home = c
	} else if d.home == c {
		d.home = l
	}
}

//func changeHome(animals []interface{}, lake Lake, cage Cage) {
//	for i := range animals {
//		if a, ok := animals[i].(Animal); ok {
//			if a.home == cage {
//				a.home = lake
//			} else if a.home == lake {
//				a.home = cage
//			}
//		}
//	}
//}

func main() {
	zk := Zookeeper{
		name: "Greg",
		age:  40,
	}

	cage := Cage{
		size: big,
		name: "cage",
	}

	lake := Lake{
		size:  small,
		depth: 5.6,
		name:  "lake",
	}

	cat := Cat{
		Animal: Animal{
			neighbour: nil,
			home:      cage,
		},
		paws: 4,
		mane: true,
	}

	tiger := Tiger{
		Cat:   cat,
		voice: "roar",
	}

	lion := Lion{
		Cat: cat,
	}

	lioness := Lioness{
		Cat: cat,
	}

	bear := Bear{
		Animal: Animal{
			name:      "BEAR",
			neighbour: nil,
			home:      cage,
		},
		color: "brown",
	}

	dolphin := Dolphin{
		Animal: Animal{
			name:      "DOLPHIN",
			neighbour: nil,
			home:      lake,
		},
		ability: "swim",
	}

	tiger.SetData(lion)
	lion.SetData(lioness)
	lioness.SetData(lion)

	//не вийшло таким чином змінити житло, ідеї закінчились як це правильно зробити(
	//прийшлось писати для кожної тварини окрему функцію
	//changeHome([]interface{}{tiger, lion, lioness, bear, dolphin}, lake, cage)

	animals := []interface{}{tiger, lion, lioness, dolphin, bear}
	msg, err := json.Marshal(animals)

	if err != nil {
		panic(err)
	}

	//[{},{},{},{},{}] чомусь ця фунцкія виводить пустий результат(
	fmt.Println(string(msg))
	//fmt.Println(animals)

	fmt.Println(zk)

	zk.ChangeHome(&tiger, &lion, &lioness, &dolphin, &bear, lake, cage)
	fmt.Println("Animals might got lost")
	fmt.Println(tiger, lion, lioness, dolphin, bear)
	zk.ChangeHome(&tiger, &lion, &lioness, &dolphin, &bear, lake, cage)
	fmt.Println(tiger, lion, lioness, dolphin, bear)
	fmt.Println("Animals are at home")
}
