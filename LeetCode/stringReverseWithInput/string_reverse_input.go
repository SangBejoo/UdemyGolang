package main

import "fmt"

func main() {
	userInput := []rune("Ayub USbaghiadasda")
	fmt.Println(string(checkUserInput(userInput)))

}

func checkUserInput(userInput []rune) []rune {
	if len(userInput) > 3 {
		return reverseString(userInput)
	} else {
		return userInput
	}
}

func reverseString(s []rune) []rune {
	size := len(s)
	for i := 0; i < size/2; i++ {

		s[i], s[size-1-i] = s[size-1-i], s[i]
		fmt.Println(string(s[size-1-i]))
	}
	return s
}
