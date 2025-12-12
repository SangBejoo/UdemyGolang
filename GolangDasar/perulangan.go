package main

import "fmt"

func main() {
	/*
		Perulangan adalah proses mengulang-ulang eksekusi blok kode tanpa henti,
		selama kondisi yang dijadikan acuan terpenuhi. Biasanya disiapkan
		variabel untuk iterasi atau variabel penanda kapan perulangan akan diberhentikan.

		biasanya perulangan digunakan ketika kita ingin mengakses data yang berjumlah banyak,
		seperti data pada array, slice, map, atau string.
	*/

	// perulangan menggunakan for
	for i := 0; i < 5; i++ {
		fmt.Printf("angka %d\n", i)
	}

	// Penggunaan for dengan argumen kondisi
	j := 0

	for j < 5 {
		fmt.Println("Angka selanjutnya", j)
		j++
	}

	// for tanpa argumen

	k := 0
	for {
		fmt.Println("Angka setelahnya", k)
		k++

		if k == 5 {
			break
		}
	}

	// for dengan range
	/*
		perulangan dengan kombinasi keyword for dan range
		untuk me-looping data gabungan misalnya string, array, slice, map, atau channel.
	*/
	xs := "123" // string
	for i, v := range xs {
		fmt.Printf("index: %d, value: %c\n", i, v)
	}

	ys := [5]int{1, 2, 3, 4, 5} // array
	for _, v := range ys {
		fmt.Println("value:", v)
	}

	zs := map[byte]int{'a': 1, 'b': 2, 'c': 3} // map
	for k, v := range zs {
		fmt.Printf("key: %c, value: %d\n", k, v)
	}

	for range zs {
		fmt.Println("hanya ingin loop sebanyak data di map tanpa butuh key dan value")
	}

	for i := range 5 {
		fmt.Println(i)
	}

	// penggunaan break and continue
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue // skip angka genap
		}
		if i > 6 {
			break // hentikan perulangan jika i lebih dari 6
		}
		fmt.Println("angka ganjil:", i)
	}

	// perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, "")
		}
		fmt.Println()
	}

	// label pada perulangan
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break outerLoop
			}
			fmt.Print("Matriks [", i, "][", j, "]", "\n")
		}
	}

}
