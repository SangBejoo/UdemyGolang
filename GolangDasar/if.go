package main

import "fmt"

func main() {
	name := "Ayub"
	address := "Indonesia"

	if name == "Ayub" {
		println("Hello Ayub")
	} else {
		println("Hello, siapa kamu?")
	}

	if address == "Indonesia" {
		println("Kamu dari Indonesia")
	} else if address == "Malaysia" {
		println("Kamu dari Malaysia")
	} else {
		println("Kamu dari negara mana?")
	}

	fmt.Println("")

	murid := make(map[string]string)
	murid["name"] = "Budi Ayub"
	murid["address"] = "Jakarta"
	murid["school"] = "SMA 1 Jakarta"

	if murid["school"] == "SMA 1 Jakarta" {
		println("Murid ini bersekolah di SMA 1 Jakarta")
	} else {
		println("Murid ini bukan bersekolah di SMA 1 Jakarta")
	}
	if murid["address"] == "Jakarta" {
		println("Murid ini tinggal di Jakarta")
	} else {
		println("Murid ini bukan tinggal di Jakarta")
	}

	if length := len(murid["name"]); length > 5 {
		println("Nama murid terlalu panjang")
	} else {
		println("Nama murid sudah benar")
	}

}
