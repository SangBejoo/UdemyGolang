package main

import "fmt"

func main() {
	// integer bilangan bulat
	// float bilangan desimal

	// int 8 -128 sampai 127
	// int 16 -32.768 sampai 32.767
	// int 32 -2.147.483.648 sampai 2.147.483.647
	// int 64 -9.223.372.036.854.775.808 sampai 9.223.372.036.854.775.807

	// ini tipe data tidak negatif
	// uint 8 0 sampai 255
	// uint 16 0 sampai 65.535
	// uint 32 0 sampai 4.294.967.295
	// uint 64 0 sampai 18.446.744.073.709.551.615

	// float32 1.18E-38 sampai 3.4E+38
	// float64 2.23E-308 sampai 1.8E+308

	// complex64 real  dan imaginer float32
	// complex128 real  dan imaginer float64

	// byte alias untuk uint8
	// rune alias untuk int32, menyimpan kode unicode

	// sesuaikan jenis kasus data untuk tipe data yang digunakan

	fmt.Println("satu", 1)
	fmt.Println("dua", 2)
	fmt.Println("tiga", 3)
	fmt.Println("tiga koma lima", float32(3.59))
	fmt.Println(1.1212121)
}
