package main

import "fmt"

type blacklist func(string) bool

func regisUser(name string, filter blacklist) {
	if filter(name) {
		fmt.Println("You Are Blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	filterName := func(name string) bool {
		return name == "Anjing"
	}
	regisUser("Anjing", filterName)

}
