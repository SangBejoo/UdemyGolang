package main

import "fmt"

func main() {
	// Tugas 1: Deklarasi Dasar
	var _ string = "Ayub" // Bukan _
	namaLengkap := "Ayub Subagiya"
	var umur int
	umur = 20
	fmt.Printf("Nama saya %s, umur %d tahun\n", namaLengkap, umur)

	// Tugas 2 Short Declaration + Multiple Variable
	a, b, c, d := 10, "hello", true, 3.14
	fmt.Printf("a adalah %d\n", a)
	fmt.Printf("b adalah %s\n", b)
	fmt.Printf("c adalah %t\n", c)
	fmt.Printf("d adalah %.2f\n", d)

	// tugas 3 tipe data eksplicit vs inferensi
	var usia1 int = 25
	usia2 := 25
	fmt.Printf("usia1 %T\n", usia1)
	fmt.Printf("usia2 %T\n", usia2)
	/*
		kenapa usia1 dan usia2 bisa sama padahal deklarasi berbeda?
		Jawaban: Karena kedua variabel tersebut diinisialisasi dengan nilai yang sama (25) dan tipe data int,
		meskipun cara deklarasinya berbeda (ekspisit untuk usia1 dan inferensi untuk usia2),
		hasil akhirnya tetap sama yaitu keduanya bertipe data int dengan nilai 25.
	*/

	// Tugas 4 Variable Underscore
	_, namaBelakang := "Ayub", "Subagiya"
	fmt.Println(namaBelakang)
	/*
		kenapa kita butuh uderscore atau _ dalam deklarasi variable?
		Jawaban: Karena dalam Go, jika ada variabel yang dideklarasikan tetapi tidak digunakan,
		maka akan menyebabkan error kompilasi. Dengan menggunakan underscore (_) sebagai variabel
		penampung, kita dapat mengabaikan nilai yang tidak diperlukan tanpa menimbulkan error.
	*/
	// tugas 5 new() dan make()
	ptr := new(string)
	*ptr = "Ini adalah pointer"
	fmt.Println(*ptr)
	fmt.Println(ptr)
	/*
		perbedaan menggunakan new dengan var, mungkin ketika menggunakan new, jelas new untuk pointer
		kita mendapatkan alamat memori dari variable tersebut, sedangkan var
		kita mendapatkan nilai dari variable tersebut secara langsung.
	*/

	// Bonus Tugas 6 Make
	angka := make([]int, 3)
	angka[0] = 100
	fmt.Println(angka)
}
