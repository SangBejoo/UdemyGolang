package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Hostname")
	username := flag.String("username", "admin", "Username")
	password := flag.String("password", "admin", "Password")

	namaDepan := flag.String("namaDepan", "Depan", "Depan")
	namaBelakang := flag.String("namaBelakang", "Belakang", "Belakang")
	namaTengah := flag.String("namaTengah", "Tengah", "Tengah")

	flag.Parse()

	fmt.Println(*host, *username, *password)
	fmt.Println(*namaDepan, *namaBelakang, *namaTengah)
}
