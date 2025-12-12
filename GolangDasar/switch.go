package main

import "fmt"

func main() {
	name := "AYUBS"

	switch name {
	case "Eko":
		fmt.Println("Eko")
	case "Budi":
		fmt.Println("Budi")
	case "AYUBsS":
		fmt.Println("AYUBS")
	default:
		fmt.Println("Tidak ada yang cocok")
	}

	paketMenu := []string{
		"Menu 1",
		"Menu 2",
		"Menu 3",
	}
	i := 2
	selectedMenu := paketMenu[i]

	switch selectedMenu {
	case "Menu 1":
		fmt.Println("Paket Menu 1")
	case paketMenu[1]:
		fmt.Println("Paket Menu 2")
	case paketMenu[2]:
		fmt.Println("Paket Menu 3")
	default:
		fmt.Println("Tidak ada yang cocok")

	}

	legnth := len(paketMenu)
	switch {
	case legnth > 3:
		fmt.Println("Banyak Menu")
	case legnth == 3:
		fmt.Println("Cukup Menu")
	default:
		fmt.Println("Sedikit Menu")
	}

	name = "Ayub Subagiya"
	length := len(name)
	switch {
	case length > 3:
		fmt.Println("Nama Terlalu Panjang")
	case length == 3:
		fmt.Println("Nama Pas")
	default:
		fmt.Println("Nama Terlalu Pendek")
	}

}
