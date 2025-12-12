package main

import "fmt"

func sayThanks(name string) string {
	return "Thanks " + name
}

func main() {
	sayTerimakasih := sayThanks
	fmt.Println(sayTerimakasih("Ayub Subagiya"))
}
