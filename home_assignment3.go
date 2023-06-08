package main

import (
	"fmt"
	"math/rand"
)

//Розробити гру-текстовий квест «Новий світ».
//Ваш персонаж прокидається в невідомому місці з деякими речами.
//	Він нічого не памʼятає. У нього є можливість піти одним з кількох шляхів (усі перелічені сутності — структури).
//	Ситуація розвивається залежно від обраного рішення.
//Ігровий режим: текстом пишеться ситуація і пропонуються текстові варіанти, які може обрати гравець.
//	Гравець пише один з варіантів і читає, як у цьому випадку розвивається ситуація.
//Можливий сценарій:
//	Стівен прокинувся біля входу в печеру.
//	Він лише памʼятає своє імʼя. Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж.
//	У печері темно, тому Стівен іде стежкою, яка веде від печери в ліс.
//	У лісі Стівен натикається на мертве тіло дивної тварини.
//	Він обирає нічого з цим не робити й іти далі.
//	Через деякий час Стівен приходить до безлюдного табору.
//	Він вже втомлений і вирішує відпочити, а не йти далі.
//	У найближчому наметі він знаходить сейф з кодовим замком з двох чисел.
//	Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає.
//	Стівен непритомніє. А все могло бути зовсім інакше.

type Survivor struct {
	Name string
	Age  int
}

type Matches struct {
	Name     string
	Quantity int
}

type Knife struct {
	Name  string
	Sharp bool
}

type Flashlight struct {
	Name       string
	Brightness string
	Battery    int
}

type Animal struct {
	Name string
}

type Box struct {
	Name string
	Code int
}

func (b *Box) Find() {
	fmt.Printf("Here is a %s\n\n", b.Name)
}

func (b *Box) GuessCode() {
	fmt.Println("Please guess a code")
	for {
		number := CheckInput()

		if b.Code == number {
			fmt.Println("Congratulation!!!!! Now you are free!!!!")
			break
		}
		if b.Code > number {
			fmt.Println("Enter bigger number")
		}
		if b.Code < number {
			fmt.Println("Enter lower number")
		}
	}
}

func (b *Box) Events() {
	b.Find()
	b.GuessCode()
}

func FindAminal(a *Animal) {
	fmt.Printf("Here is a %s\n", a.Name)
}

func (k *Knife) Use() {
	fmt.Printf("Let's use a %s\n\n", k.Name)
	k.Sharp = false
}

func (k *Knife) Guess(yesNoOptions [2]string) {
	fmt.Printf("How do you think, is my %s still sharp\n", k.Name)
	ChooseOption(yesNoOptions[:])
	k.Check(CheckInput())
}

func (k *Knife) Check(input int) {
	if input == 1 && k.Sharp == true {
		fmt.Printf("You are right, it is %t", k.Sharp)
	} else if input == 2 && k.Sharp == false {
		fmt.Printf("You are right, it is %t", k.Sharp)
	} else {
		fmt.Println("You are wrong\n")
	}
}

func EatAnimal(yesNoOptions [2]string, k *Knife) {
	fmt.Println("I am so hungry, maybe i can cook it?")
	ChooseOption(yesNoOptions[:])
	if CheckInput() == 1 {
		k.Use()
		k.Guess(yesNoOptions)
	} else {
		fmt.Println("Ok, let's go futher!\n")
	}
}

func (f *Flashlight) Use(yesNoOptions [2]string) {
	fmt.Printf("It's too dark here!\nShould I use a %s\n", f.Name)
	ChooseOption(yesNoOptions[:])
	if CheckInput() == 1 {
		f.ReduceBattery()
		fmt.Printf("Oh, no %s's battary is now %d\n\n", f.Name, f.Battery)
	} else {
		fmt.Println("Well, save the battery!\n\n")
	}
}

func (f *Flashlight) ReduceBattery() {
	f.Battery--
}

type Belongings struct {
	*Matches
	*Knife
	*Flashlight
}

func (b *Belongings) Show() {
	fmt.Printf("Look, I'v got %d %s.\nAnd also a %s.\nAnd %s %s but with low battery only %d\n\n",
		b.Quantity, b.Matches.Name, b.Knife.Name, b.Brightness, b.Flashlight.Name, b.Battery)
}

func WakeUp() {
	fmt.Println("Uhhh, what is going on? Where am I?")
}

func Greetings(survivor Survivor) {
	fmt.Printf("Hi!\nWelcome on board and have fun.\nPoor %s got lost on the uninhabited island...\nHelp him to survive\n",
		survivor.Name)
}

func ChooseOption(options []string) {
	for i, option := range options {
		i++
		fmt.Println(i, " - ", option)
	}
	fmt.Println("Please, make your choice and enter a number!")
}

func CheckInput() int {
	var input int
	fmt.Scan(&input)
	return input
}

func FindBag() {
	fmt.Println("It seems to be my bag over there! Let's check what is in?\n")
}

func GoTo(place string) {
	fmt.Printf("Let's go to %s\n\n", place)
}

func main() {
	survivor := Survivor{
		Name: "Steven",
		Age:  30,
	}

	matches := Matches{
		Name:     "matches",
		Quantity: 3,
	}

	knife := Knife{
		Name:  "knife",
		Sharp: true,
	}

	flashlight := Flashlight{
		Name:       "flashlight",
		Brightness: "bright",
		Battery:    3,
	}

	belongings := Belongings{
		&matches,
		&knife,
		&flashlight,
	}

	deer := Animal{Name: "deer"}

	box := Box{
		Name: "box",
		Code: rand.Intn(15),
	}

	Greetings(survivor)
	WakeUp()

	placeOptions := [3]string{"Lake", "Cage", "Beach"}
	yesNoOptions := [2]string{"yes", "no"}
	ChooseOption(placeOptions[:])

Game:
	for {
		switch CheckInput() {
		case 1:
			fmt.Println("Oh my God, I can't swim!\n")
			FindBag()
			belongings.Show()
			GoTo("look around")
			box.Events()
			break Game
		case 2:
			fmt.Println("I hope there is no wild animal next to me!")
			FindBag()
			belongings.Show()
			flashlight.Use(yesNoOptions)
			GoTo("the forest")
			FindAminal(&deer)
			EatAnimal(yesNoOptions, &knife)
			box.Events()
			break Game
		case 3:
			fmt.Println("Oh, what a lovely place")
			FindBag()
			belongings.Show()
			GoTo("....")
			fmt.Println("No I will stay here and have a rest")
			break Game
		default:
			fmt.Println("No-no-no only:")
			ChooseOption(placeOptions[:])
		}
	}
}
