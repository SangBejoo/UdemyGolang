package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToLower("Ayub SUbagiya"))
	fmt.Println(strings.ToUpper("Ayub SUbagiya"))
	fmt.Println(strings.Trim("         Ayub SUbagiya", " "))
	fmt.Println(strings.Contains("Ayub Subagiya", "Ayub"))
	fmt.Println(strings.Split("Ayub Subagiya", " "))
	fmt.Println(strings.Count("Ayub Subagiya", "Ayub"))
	fmt.Println(strings.ReplaceAll("Ayub Subagiya", "Ayub", "Subagiya"))

}
