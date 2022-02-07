package test

import (
	"fmt"
	"testing"
)

type LivingThings interface {
	Walk() string
	Run() string
	Eat() string
}

type Human struct {
	Name string
	Age  int
}

type Animal struct {
	Name string
	Age  int
	Type string
}

type Teritorial interface {
}

func (a Animal) Walk() string {
	return "Animal Walk"
}

func (h Human) Walk() string {
	return "Human Walk"
}

func (a Animal) Run() string {
	return a.Name + " Run"
}

func TestAnimalWalk(t *testing.T) {
	tiger := Animal{
		Name: "Tiger",
		Type: "Predator",
	}

	fmt.Println(tiger.Walk())
}

func TestAnimalRun(t *testing.T) {
	rabbit := Animal{
		Name: "Rabbit",
	}
	fmt.Println(rabbit.Run())
}

func TestHumanWalk(t *testing.T) {
	arnold := Human{
		Name: "Arnold",
		Age:  20,
	}

	fmt.Println(arnold.Walk())
}
