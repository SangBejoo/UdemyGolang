package main

import "fmt"

func main() {
	// boolean hanya punya dua nilai, true dan false
	// biasanya digunakan untuk kondisi

	fmt.Println("Benar =", true)
	fmt.Println("Salah =", false)

	var nilaiAkhir = 90
	var absesnsi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsesnsi bool = absesnsi > 80

	var lulus = lulusNilaiAkhir && lulusAbsesnsi
	fmt.Println(lulus)

}
