package src

import "fmt"

type person struct {
	Name string
	Age  int
}

func NewPerson(
	name string,
	age int,
) *person {
	return &person{
		Name: name,
		Age:  age,
	}
}

func (p *person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (p *person) SayHello() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

func (p *person) Aging() {
	p.Age++
}

func (p *person) SetName(name string) {
	p.Name = name
}

//go:generate mockery --name=Person --case=underscore --inpackage --testonly
type Person interface {
	String() string
	SayHello()
	Aging()
	SetName(name string)
}
