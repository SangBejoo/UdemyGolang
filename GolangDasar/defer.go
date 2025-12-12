package main

import "fmt"

func logging() {
	fmt.Println("Logging Berjalan")
	message := recover()
	fmt.Println("terjadi Error", message)
}

func runApplication() {
	defer logging()
	fmt.Println("Running Application")
}

func App(error bool) {
	defer logging()
	if error {
		panic("Error terjadi")
	}
}

func main() {
	//runApplication()
	App(true)
	fmt.Println("Running Application")
}
