package main

import "fmt"

func main() {
	// Tugas 1 kenalan tipe dasar
	umur := 20
	populasi := int64(8000000000000)
	pi := float64(3.14159)
	isStudent := true
	kata := "Golang"
	huruf := 'A'
	emoji := rune('😊')

	fmt.Printf("Nilai umur: %d, dengan tipe data %T\n", umur, umur)
	fmt.Printf("Nilai populasi: %d, dengan tipe data %T\n", populasi, populasi)
	fmt.Printf("Nilai pi: %.5f, dengan tipe data %T\n", pi, pi)
	fmt.Printf("Nilai isStudent: %t, dengan tipe data %T\n", isStudent, isStudent)
	fmt.Printf("Nilai kata: %s, dengan tipe data %T\n", kata, kata)
	fmt.Printf("Nilai huruf: %c, dengan tipe data %T\n", huruf, huruf)
	fmt.Printf("Nilai emoji: %c, dengan tipe data %T\n", emoji, emoji)

	// tugas 2 float bikin bingung
	a := float32(3.14)
	var b float64 = 3.14
	fmt.Printf("a: %.6f\n", a)
	fmt.Printf("b: %.6f\n", b)

	// tugas 3 string vs byte vs rune babee tugasnya ini berlebihan aku belum paham konsep array
	perkataan := "Ayub 💖💖"
	fmt.Println(len(perkataan))
	bytes := []byte(perkataan)
	fmt.Println(len(bytes))
	runes := []rune(perkataan)
	fmt.Println(len(runes))

	// tugas 4 integer overflow
	maxInt8 := int8(127)
	maxInt8++
	fmt.Println("Nilai setelah overflow int8:", maxInt8)

	// tipe data komplex >:D
	c := complex64(1 + 2i)
	d := complex128(3.14 + 1i)
	fmt.Println(real(c), imag(c))
	fmt.Println(real(d), imag(d))

	// tugas 6 pembulatan
	f := float64(3.99)
	i := int(f)
	fmt.Println("Nilai setelah pembulatan:", i)
}
