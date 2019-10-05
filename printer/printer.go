package printer

import "fmt"

// Hello funct
func Hello() {
	fmt.Println("hello")
}

type person struct {
	name string
	age  int
}

// NewPerson constructor returns person
func NewPerson(name string, age int) *person {
	return &person{name, age}
}

func (p *person) GetPersonName() string {
	return p.name
}

func (p *person) GetPersonAge() int {
	return p.age
}
