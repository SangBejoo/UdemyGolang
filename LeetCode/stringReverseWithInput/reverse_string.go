package main

import "fmt"

func main() {
	userInput1 := []byte("Ayub Subagiya Selalau Happy")
	fmt.Println(string(checkUserInput1(userInput1)))
}

func checkUserInput1(u []byte) []byte {
	if len(u) > 5 {
		return reverseString1(u)
	} else {
		return u
	}
}

func reverseString1(s []byte) []byte {
	size := len(s)
	for i := 0; i < size/2; i++ {
		s[i], s[size-1-i] = s[size-1-i], s[i]
	}
	return s
}
