package main

import "fmt"

func main() {
	/*
		Chapter ini membahas mengenai macam operator yang bisa digunakan di Go.
		Secara umum terdapat 3 kategori operator:
		aritmatika, perbandingan, dan logika.
	*/
	// operator Aritmatika
	a := 99
	b := 1
	c := a + b
	d := c - a
	e := d * c
	f := e / d
	g := f % b
	fmt.Println(a, b, c, d, e, f, g)

	// operator Perbandingan
	str := "ayub"
	srt := "ayub"
	isSame := str == srt
	isNotSame := str != srt
	lowThan := a < b
	greaterThan := a > b
	lowThanEqual := a <= b
	greaterThanEqual := a >= b
	fmt.Println(isSame, isNotSame, lowThan, greaterThan, lowThanEqual, greaterThanEqual)

	// Operator Logika
	kanan := true
	kiri := false

	leftAndRight := kanan && kiri
	leftOrRight := leftAndRight && kiri
	notLeft := !kiri
	fmt.Println(leftAndRight, leftOrRight, notLeft)
	fmt.Print(str)

	// fizz buzz
	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 && i%5 == 0 {
	//		fmt.Println("FizzBuzz")
	//	} else if i%3 == 0 {
	//		fmt.Println("Fizz")
	//	} else if i%5 == 0 {
	//		fmt.Println("Buzz")
	//	} else {
	//		fmt.Println(i)
	//	}
	//}
}
