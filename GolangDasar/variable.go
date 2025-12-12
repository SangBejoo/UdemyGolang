package main

import "fmt"

func main() {

	// variable adalah tempat menyimpan data sementara di memory
	// variable itu punya tipe data
	// cara membuat variable
	// var namaVariable tipeData =
	// atau dengan variable := nilai

	var supplement string

	supplement = "Bulk Supplement Caffein"
	fmt.Println(supplement)

	supplement = "Creatine Monohydrate"
	fmt.Println(supplement)

	// ini := deklarasi awal
	name := "Ayub"
	fmt.Println(name)

	name = "Subagiya"
	fmt.Println(name)

	// deklarasi multiple variable
	var (
		firstName = "Ayub"
		lastNmae  = "Subagiya"
	)

	var (
		a, b, c = 10, "hello", true
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// swap variable
	firstName, lastNmae = lastNmae, firstName

	fmt.Println(lastNmae)
	fmt.Println(firstName)
	fmt.Println(firstName, lastNmae)

}
