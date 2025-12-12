package main

import "fmt"

func main() {

	// constant adalah variable yang nilainya tidak bisa diubah setelah di deklarasi
	// cara membuat constant
	// const namaConstant tipeData = nilai

	const name = "Ayub"
	name1 := "Subagiya"
	fmt.Println(name1)
	fmt.Println(name)

	name1 = name
	fmt.Println(name1 + " " + name)

	const (
		Menu1 = "Nasi Goreng"
		Menu2 = "Mie Goreng"
		Menu3 = "Kwetiaw"
	)
	fmt.Println(Menu1)
	fmt.Println(Menu2)
	fmt.Println(Menu3)

	const pi = 3.14
	fmt.Println(pi)

}
