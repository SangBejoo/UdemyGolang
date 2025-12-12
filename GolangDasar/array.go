package main

import "fmt"

func main() {
	// Array here
	var names [3]string

	names[0] = "Ayub"
	names[1] = "Subagiya"
	names[2] = "Salhaa"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{
		90,
		80,
		70,
		60,
		50,
	}

	values[1] = 12
	fmt.Println(len(values))
	fmt.Println(values[1])
	fmt.Println(values[2])
}
