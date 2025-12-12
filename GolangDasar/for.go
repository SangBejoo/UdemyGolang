package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	fmt.Println("Selesai")
	fmt.Println("")

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke", i)
	}

	names := []string{
		"Asep",
		"Budi",
		"Caca",
		"Deni",
		"Eka",
	}
	for j := 0; j < cap(names); j++ {
		fmt.Println("Index", j, "=", names[j])
	}

	for index, name := range names[2:4] {
		fmt.Println("Index", index, "=", name)
	}
	for _, name := range names {
		fmt.Println("Name", name)
	}
}
