package test

import (
	"fmt"
	"testing"
)

func Change(original *int, value int) {
	*original = value
}

func TestPointer(t *testing.T) {
	var number = 4
	Change(&number, 10)
	fmt.Println(number)
}

func TestReferencingAndDereferencing(t *testing.T) {
	var numberA = 4
	var numberB = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("")

	numberA = 10

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)
}
