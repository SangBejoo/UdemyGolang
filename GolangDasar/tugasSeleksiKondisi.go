package main

import "fmt"

func main() {
	// tugas 1 if else dasar
	nilai := 95

	if nilai >= 90 {
		fmt.Println("Luar Biasa!")
	} else if nilai >= 70 && nilai < 90 {
		fmt.Println("Bagus")
	} else if nilai >= 50 && nilai < 70 {
		fmt.Println("Cukup")
	} else {
		fmt.Println("Perlu Belajar Lebih")
	}
	/*
		kenapa if-else chain lebih baik dari pada switch?
		karena penggunaan if-else chain lebih fleksibel untuk kondisi yang kompleks dan rentang nilai,
		sedangkan switch lebih cocok untuk kondisi yang sederhana dan nilai diskrit.
	*/

	// tugas 2 temporary var di if + multiple kondisi
	berat := 70.0
	tinggi := 1.75

	if bmi := berat / (tinggi * tinggi); bmi >= 30 {
		fmt.Println("Obesitas")
	} else if bmi >= 25 && bmi <= 29.9 {
		fmt.Println("Overweight")
	} else if bmi >= 18.5 && bmi <= 24.9 {
		fmt.Println("Normal")
	} else {
		fmt.Println("Uderweight")
	}

	// tugas 3 switch cases dan multiple cases
	hari := "Senin"

	switch hari {
	case "Senin":
		fmt.Println("Hari Ini Senin")
	case "Selasa":
		fmt.Println("Hari Ini Selasa")
	case "Rabu":
		fmt.Println("Hari Ini Rabu")
	case "Kamis":
		fmt.Println("Hari Ini Kamis")
	case "Jumat":
		fmt.Println("Hari Ini Jumat")
	case "Sabtu":
		fmt.Println("Hari Ini Sabtu")
	case "Minggu":
		fmt.Println("Hari Ini Minggu")
	}

	// tugas 4 switch gaya if-else dan fallthrough
	var umur = 17
	switch {
	case umur >= 21:
		fmt.Println("Dewasa Penuh")
	case umur >= 18:
		fmt.Println("Dewasa Muda")
		fmt.Println("bisa Nyoblos")
	case umur >= 13:
		fmt.Println("Remaja")
	default:
		fmt.Println("Anak-anak")
	}

	// tugas 5 Nested if/switch
	nilaiAkhir := 85

	if nilaiAkhir >= 70 {
		switch {
		case nilaiAkhir >= 90 && nilaiAkhir <= 100:
			fmt.Println("A")
		case nilaiAkhir >= 80 && nilaiAkhir <= 89:
			fmt.Println("B")
		default:
			{
				fmt.Println("C")
			}
		}
	} else if nilaiAkhir >= 50 && nilaiAkhir <= 79 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
		if nilaiAkhir == 0 {
			fmt.Println("Gagal Total")
		}
	}

	// tugas 6 kombinasi if-Switch dan error handling
	age := 15
	negara := "Indonesia"

	if age >= 0 {
		switch {
		case negara == "Indonesia":
			if age >= 18 {
				fmt.Println("Bisa Vote")
			} else {
				fmt.Println("Belum Bisa Vote")
			}
		case negara == "USA":
			if age >= 18 {
				fmt.Println("Bisa Vote")
			} else {
				fmt.Println("Belum Bisa Vote")
			}
		default:
			{
				fmt.Println("Negara Tidak Dikenal")
			}
		}
	} else {
		fmt.Println("Umur Invalid")
	}

}
