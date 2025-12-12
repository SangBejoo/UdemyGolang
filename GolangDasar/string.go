package main

import "fmt"

func main() {

	// tipe data string adalah kumpulan karakter yang diapit dengan tanda petik dua ""
	// string di golang itu immutable, artinya tidak bisa diubah setelah dibuat
	// untuk mengubah string harus membuat string baru
	// string itu disimpan dalam bentuk byte
	// panjang string dihitung berdasarkan byte, bukan karakter
	// untuk mengakses karakter dalam string bisa menggunakan indeks, dimulai dari 0

	fmt.Println("Bulk Supplement Caffein")
	fmt.Println("Ayub")
	fmt.Println("Subagiya")
	fmt.Println(len("Ayub "))
	fmt.Println(len("razer mouse gaming ini sangan keren"))
	// pengambilan karakter berdasarkan indeks nanti jadi byte
	fmt.Println("razer mouse gaming ini sangan keren"[1])
	fmt.Println("Subagiya"[0])
}
