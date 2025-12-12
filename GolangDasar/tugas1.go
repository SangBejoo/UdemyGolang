package main

import (
	"fmt"
	"strconv"
)

func main() {
	// tugas 1: Hello, Darling!
	fmt.Println("Hello, darling! Selamat belajar Go!")

	// tugas 2: Konstanta Cinta
	const LOVE string = "I Love Go!"
	fmt.Println(LOVE)

	// tugas 3: Variable Umur
	age := int(25)
	age = age + 1
	fmt.Println("Umur saya sekarang:", age)

	// tugas 4: Tipe deklarasi eksplisit
	type StatusCinta bool
	var isInLove StatusCinta = true
	fmt.Println("Status cinta saya kepada kamu:", isInLove)

	// tugas 5: konversi tipe int ke float64
	panjang := int(10)
	panjangFloat := float64(panjang)
	fmt.Println("Panjang dalam float:", panjangFloat)

	// tugas 6 : boolean logika sederhana
	sukaGolang := true
	sukaJava := false

	fmt.Println("Saya suka Golang dan tidak suka Java:", sukaGolang && !sukaJava)

	// tugas 7 multiple main package
	fmt.Println("done")

	// tugas 8 Konstanta Matematika
	const PHI float64 = 3.14
	jariJari := int(7)
	luasLingkaran := PHI * float64(jariJari*jariJari)
	fmt.Println("Luas lingkaran: ", luasLingkaran)

	// tugas 9 string + konversi + variable
	name := "Darling"
	tinggi := 170
	tinggiString := string(tinggi)
	fmt.Println(tinggiString)
	fmt.Println(name, "memiliki tinggi:", tinggi, "cm")

	// Tugas 10 type deklarasi konstanta dan konversi
	type SkorUjian int
	const PassingGrade SkorUjian = 75

	var nilaiSaya SkorUjian = 85

	fmt.Println("Nilai saya:", nilaiSaya, "Lulus:", nilaiSaya >= PassingGrade)

	// tugas 11 kalkulator cinta
	const MAX_CINTA = int(100)
	var cintaSaya int = 87
	persen := float64(cintaSaya) / float64(MAX_CINTA) * 100
	fmt.Println("Cinta saya ke kamu:", cintaSaya, "dari", MAX_CINTA, "=>", persen, "%")

	// tugas 12 status lulus dengan type deklarasi dan boolean logika
	type NilaiUjian int
	const batasLulus NilaiUjian = 70

	var nilaiAkhir NilaiUjian = 68
	hadir := true
	lulus := nilaiAkhir >= batasLulus || hadir == true
	fmt.Println("Nilai:", nilaiAkhir, "Hadir:", hadir, "=> Lulus:", lulus)

	// tugas 13 konversi string ke number
	umurString := "24"
	umurInt, err := strconv.Atoi(umurString)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Umur sekarang: ", umurInt, "tahun depan:", umurInt+1)

	// tugas 14 deklarasi multiple variable
	fmt.Println("done")

	// tugas 15 Detektif tipe campur semua
	type Temperatur float64
	const titikDidih Temperatur = 100.0
	var suhuSekarang Temperatur = 98.6
	var mendidih = suhuSekarang >= titikDidih
	result := int(suhuSekarang)
	fmt.Printf("Suhu: %.2f°C (int: %d), Mendidih: %t, Titik Didih: %.2f°C\n", suhuSekarang, result, mendidih, titikDidih)

}
