package model

type Student struct {
	ID     int
	Name   string
	Age    int
	Gender string
	Course string
}

type Course struct {
	Id    int
	Name  string
	Field string
}
