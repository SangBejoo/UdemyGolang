package main

import "fmt"

func filterName(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "Nama Hewan"
	} else {
		return "Hello" + name
	}
}

func main() {
	filter := spamFilter
	name := filterName
	name("Budi", filter)
	filterName("Anjing", filter)

}
