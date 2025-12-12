package main

import "fmt"

func main() {
	/*
		Dianjurkan untuk tidak sembarangan dalam menentukan tipe data variabel,
		sebisa mungkin tipe yang dipilih harus disesuaikan dengan nilainya,
		karena efeknya adalah ke alokasi memori variabel.
		Pemilihan tipe data yang tepat akan membuat pemakaian memori lebih optimal, tidak berlebihan.
	*/

	var positiveNumber uint8 = 89
	var negativeNumber = -100
	positiveNumber1 := uint8(89)

	fmt.Printf("bilangan positive: %d\n", positiveNumber)
	fmt.Printf("bilangan negative: %d\n", negativeNumber)
	fmt.Printf("bilangan positive number: %d\n", positiveNumber1)

	// ini untuk desimal dan cara manggilnya menggunakan %.(angka desimalnya)f
	decimalNumber := 2.2
	fmt.Printf("bilangan decimal: %.2f\n", decimalNumber)
	nomorDesimal := float32(2.3)
	fmt.Printf("bilangan nomor desimal: %.2f\n", nomorDesimal)

	// ini untuk boolean menggunakan %t
	isActive := true
	var isInactive bool = false
	fmt.Printf("is active: %t\n", isActive)
	fmt.Printf("is inactive: %t\n", isInactive)

	// ini untuk string menggunakan %s
	var iniString string = "halo ayub"
	salam := "selamat pagi"
	fmt.Printf("ini string: %s\n", iniString)
	fmt.Printf("salam: %s\n", salam)

}
