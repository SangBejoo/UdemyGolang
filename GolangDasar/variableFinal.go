package main

import "fmt"

func main() {
	// ini adalah variable
	var namaPertama string = "Ayub Subagiya"
	var namaKedua string
	namaKedua = "Ayub Juga"

	fmt.Printf("halo %s dan %s\n", namaPertama, namaKedua)

	// var tanpaTipe data
	usia := 20
	fmt.Printf("usia saya adalah %d tahun\n", usia)

	// multiple variable
	var first, second, third string = "satu", "dua", "tiga"
	fourth, fifth, sixth := "empat", "lima", "enam"
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(first, second, third)
	fmt.Println(fourth, fifth, sixth)
	fmt.Println(one, isFriday, twoPointTwo, say)

	// variable underscore
	/*
		Golang tidak mengizinkan adanya variable yang tidak digunakan.
	*/
	_ = "ini adalah variable underscore"
	name, _ := "Ayub", "Subagiya"
	fmt.Println(name)

	// menggunakan new
	/*
		ini biasanya untuk pointer
	*/
	name1 := new(string)
	fmt.Println(name)
	fmt.Println(*name1)

	// Variable dengan Make
	/*
		ini biasanya untuk channel, slice, dan map
	*/

	// Tugas dari AI

}
