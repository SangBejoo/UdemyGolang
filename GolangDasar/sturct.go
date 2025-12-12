package main

import "fmt"

type Customer struct {
	Name, Address, Alamat string
	Age, Kelahiran        int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Subang"
	eko.Age = 30
	eko.Kelahiran = 1990
	fmt.Println(eko.Name)

	// struct literal
	Ayub := Customer{
		Name:      "Ayub",
		Address:   "Bandung",
		Age:       25,
		Kelahiran: 1995,
	}
	fmt.Println(Ayub)

	Ayubs := Customer{"Ayubs", "Bandung", "Bekasi", 25, 1995}
	fmt.Println(Ayubs)

	Ayubs.sayHello("Subagiya")

}
