package main

import "fmt"

func main() {

	// konversi tipe data
	// konversi itu merubah dari satu tipe data ke tipe data lain
	// cara konversi
	// tipeDataBaru(nilai)

	// contoh konversi dari int32 ke int64
	var nilai32 int32 = 999999
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai16)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	// konversi dari float64 ke int
	var nilaiFloat64 float64 = 99.99
	var nilaiInt int = int(nilaiFloat64)

	fmt.Println(nilaiFloat64)
	fmt.Println(nilaiInt)

	// konversi dari string ke byte
	var name string = "Ayub"
	var nameByte byte = name[0]

	fmt.Println(name)
	fmt.Println(nameByte)

	// konversi dari byte ke string
	var nameString string = string(nameByte)

	fmt.Println(nameByte)
	fmt.Println(nameString)

	// konversi dari int ke string
	var angka int = 65
	var angkaString string = string(angka)

	fmt.Println(angka)
	fmt.Println(angkaString)

}
