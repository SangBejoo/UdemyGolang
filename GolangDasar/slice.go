package main

import "fmt"

func main() {
	//slice
	name := [...]string{
		"John",
		"Jane",
		"Jack",
		"Jill",
	}
	slice1 := name[3:4]
	fmt.Println(slice1)
	slice2 := name[1:4]
	fmt.Println(slice2)
	slice3 := name[2:]
	fmt.Println(slice3)
	slice4 := name[:]
	fmt.Println(slice4)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice2))
	fmt.Println(append(slice3, "James"))

	fmt.Println("")

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	daySlices1 := days[5:]
	fmt.Println(daySlices1)
	daySlices1[0] = "Sabtu Baru"
	daySlices1[1] = "Minggu Baru"
	fmt.Println(daySlices1)
	fmt.Println(days)

	daySlice2 := append(daySlices1, "Libur Baru")
	daySlice2[0] = "Sabtu Lagi"
	fmt.Println(daySlice2)
	fmt.Println(cap(daySlice2))
	fmt.Println(days)

	fmt.Println("")

	// make slice
	var newSlice = make([]string, 2, 5)
	newSlice[0] = "ayub"
	newSlice[1] = "budi"
	// newSlice[2] = "caca" // akan error karena index 2 belum ada
	newSlice = append(newSlice, "caca")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice[2] = "Budi"
	fmt.Println(newSlice)

	fmt.Println("")

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// literal slice
	fmt.Println("")
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
