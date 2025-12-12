package main

import "fmt"

type HasName interface {
	GetName() string
}

type hasType interface {
	GetType() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

type TipeAnimal struct {
	Tipe string
}

func (animal TipeAnimal) GetName() string {
	return animal.Tipe
}

func sayHello1(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	person := Person{"Eko"}
	sayHello1(person)

	animal := Animal{"Anjing"}
	sayHello1(animal)

	animal1 := TipeAnimal{Tipe: "Kadal"}
	sayHello1(animal1)
}
