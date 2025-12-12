package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)
	if age <= 0 {
		fmt.Println("Age doesn't make sense—try again?")
	} else {
		fmt.Printf("Hello, %s! You're %d years old.\nin 10 years, you'll be %d\n", name, age, age+10)
	}

}
