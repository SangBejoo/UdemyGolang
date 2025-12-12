package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.FindString("hello eyo ero ezo"))
	fmt.Println(regex.MatchString("Ayub Subagiya"))

	fmt.Println(regex.FindAllString("eko ayub subagiya eko edo EKKOO", 10))
	fmt.Println(len(regex.FindAllString("eko ayub subagiya eko edo EKKOO", 10)))

}
