package main

import (
	"fmt"
	"strings"
)

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func menuAll(menus ...string) string {
	menu := strings.Join(menus, ", ")
	return menu
}

func sumAllSLice(number []int) int {
	total := 0
	for _, number1 := range number {
		total += number1
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 34, 5))
	fmt.Println(menuAll("Nasi Goreng", "Mie Goreng", "Sate"))
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumAll(number...))
	result := sumAllSLice(number)
	fmt.Println(result)

}
