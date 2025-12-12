package main

import "fmt"

func main() {
	/*
		Seleksi kondisi digunakan untuk mengontrol alur eksekusi flow program.
		Analoginya mirip seperti fungsi rambu lalu lintas di jalan raya.
		Kapan kendaraan diperbolehkan melaju dan kapan harus berhenti diatur oleh rambu tersebut.
		Seleksi kondisi pada program juga kurang lebih sama, kapan sebuah blok kode dieksekusi dikontrol.
	*/
	nilai := 75

	if nilai >= 80 {
		fmt.Printf("Nilai Kamu %d dibawah rata rata yaitu 80\n", nilai)
	} else if nilai >= 60 {
		fmt.Println("Nilai Kamu sudah cukup Baik")
	} else {
		fmt.Println("Belajar lebih giat lagi ya!")
	}

	// temporary variable dari if else

	point := 8888.400

	if percent := point / 100; percent >= 100 {
		fmt.Printf("NIceeee Point %.2f\n", percent)
	} else if percent >= 50 {
		fmt.Printf("Not a BadPoint %.2f\n", percent)
	} else {
		fmt.Printf("Bad Point %.2f", percent)
	}

	// penggunaan switch case
	kata := "satu"

	switch kata {
	case "satu":
		fmt.Println("Satu")
	case "dua":
		fmt.Println("dua")
	case "tiga":
		fmt.Println("tiga")
	default:
		fmt.Println("lainnya")
	}
	// case jika banyak kondisi
	angka := 2

	switch angka {
	case 10:
		fmt.Println("angka kamu 10")
	case 9, 8, 7, 6, 5:
		fmt.Println("angka kamu diantara 5 hingga 9")
	default:
		{
			fmt.Println("angka kamu dibawah 5")
			fmt.Println(angka)
		}
	}

	// switch dengan gaya if else
	switch {
	case angka == 10:
		fmt.Println("angka kamu 10")
	case (angka <= 5) && (angka >= 3):
		fmt.Println("angka kamu diantara 3 dan 5")
	default:
		{
			fmt.Println("angka kamu dibawah 2")
			fmt.Println(angka)
		}
	}

	// penggunaan switch dengan fallthrough
	nilaiAkhir := 0

	switch {
	case nilaiAkhir == 99:
		fmt.Println("Nilai Kamu perfect")
	case nilaiAkhir >= 80:
		fmt.Println("Sudah Hampir Sempurna")
		fallthrough
	case nilaiAkhir >= 70:
		fmt.Println("Nilai Kamu diatas Rata rata")
	default:
		{
			fmt.Println("SUdah Hampir Sempurna")
			fmt.Println(nilaiAkhir)
		}
	}

	// nested selection
	if nilaiAkhir >= 70 {
		switch nilaiAkhir {
		case 90:
			fmt.Println("Nilai Kamu A")
		case 80:
			fmt.Println("Nilai Kamu B")
		default:
			fmt.Println("Nilai Kamu C")
		}
	} else {
		if nilaiAkhir == 60 {
			fmt.Println("Nilai Kamu D")
		} else if nilaiAkhir == 50 {
			fmt.Println("Nilai Kamu E")
		} else {
			fmt.Println("Try Harder Next time")
			if nilaiAkhir == 0 {
				fmt.Println("Nilai Kamu bodoh")
			}

		}
	}

}
