package main

type Class struct {
	Students []Student
	Teacher  Teacher
	Id       string
}

type Student struct {
	Name string
	Id   string
}

type Teacher struct {
	Password string
	Username string
	ClassId  string
}
