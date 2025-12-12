package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, ayubs!")
}
func sayGoodbye(firstName string, lastName string) {
	fmt.Println("Goodbye, " + firstName + " " + lastName)
}

func array(numbers []int) {
	for _, number := range numbers {
		fmt.Println(number)
	}
}

func getHello(name string) string {
	hello := "Hello, " + name
	return hello
}

func getFullName() (string, string) {
	return "John", "Doe"
}

func getMenu() (string, string, string) {
	return "Nasi Goreng", "Mie Goreng", "Sate"
}

func getCompletedName() (firstName1, middleName1, lastName1 string) {
	firstName1 = "John"
	middleName1 = "Middle"
	lastName1 = "Doe"

	return firstName1, middleName1, lastName1

}

func main() {

	a, b, c := getCompletedName()
	fmt.Println(a, b, c)

	menu1, menu2, _ := getMenu()
	fmt.Println(menu1)
	fmt.Println(menu2)
	fmt.Println(getMenu())

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	result := getHello("James")
	fmt.Println(result)
	fmt.Println(getHello("James1"))

	sayHello()

	for i := 0; i < 5; i++ {
		sayHello()
		sayGoodbye("John", "Doe")
	}

	sayGoodbye("Ayubs", "Subagiya")

	name := "Ayubsss"
	if name == "Ayubs" {
		sayGoodbye("Ayubs", "Subagiya")
	} else if name == "Ayubss" {
		sayGoodbye("Ayubs", "Subagiya")
	} else {
		sayHello()
	}

	numbers := []int{1, 2, 3, 4, 5}
	array(numbers[3:4])

}
