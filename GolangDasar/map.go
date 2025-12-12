package main

import "fmt"

func main() {
	// Tipe data map
	menu := map[string]string{}
	menu["name"] = "Nasi Goreng"
	menu["price"] = "15000"
	menu["stock"] = "20"

	fmt.Println(menu)
	fmt.Println(menu["name"])
	fmt.Println(menu["price"])
	fmt.Println(menu["stock"])

	// Membuat cepat
	person := map[string]string{
		"name":    "John",
		"address": "New York",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	fmt.Println("===================")
	// Menghapus data pada map
	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "John Doe"
	book["year"] = "2020"
	book["price"] = "500000000"

	fmt.Println(book)
	delete(book, "price")
	fmt.Println(book)

}
