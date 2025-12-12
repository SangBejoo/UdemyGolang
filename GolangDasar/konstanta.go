package main

import "fmt"

func main() {
	/*
		Konstanta adalah jenis variabel yang nilainya tidak bisa diubah setelah dideklarasikan.
		Inisialisasi nilai konstanta hanya dilakukan sekali saja di awal,
		setelah itu variabel tidak bisa diubah nilainya.
	*/
	const firstName string = "John"
	const fullName = firstName + " subagiya"
	fmt.Println(firstName)

	// Deklarasi Multi Konstanta
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 8
		floatNum       = 2.33
		// contoh jika tidak diisi maka nilainya akan mengikuti dari atasnya
		b
		c
		d
		e
	)

}
