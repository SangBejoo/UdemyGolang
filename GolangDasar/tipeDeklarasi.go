package main

import "fmt"

func main() {
	type NoKTP string
	var ktpAyub NoKTP = "23131231321"
	var contoh string = "231321321"

	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpAyub)
	fmt.Println(contohKTP)

	type Married bool
	var married Married = true
	fmt.Println(married)
}
