package main

import "fmt"

func main() {
	counter := 0

	increment := func() int {
		fmt.Println("increment")
		counter++

		return counter
	}

	increment()
	increment()
	increment()

	fmt.Println("Counter:", counter)
}
