package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Ayub", "Subagiya", "Belajar", "Golang"}
	values := []int{10, 20, 30, 40}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Contains(names, "Ayub"))
	fmt.Println(slices.Contains(values, 50))
	fmt.Println(slices.Index(names, "Golang"))
	fmt.Println(slices.Index(values, 30))
}
