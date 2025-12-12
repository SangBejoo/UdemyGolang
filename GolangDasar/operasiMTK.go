package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = 30

	var d = a + b*c
	fmt.Println(d)

	// augmented assignment
	var i = 10
	i += 30 // i = i + 30
	i += 20
	fmt.Println(i)

	// unary operator
	var j = 2
	j++ // j = j + 1
	fmt.Println(j)
	j++ // j = j + 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)
}
