package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("validation error")
	notFoundError   = errors.New("not found")
)

func check(s string) error {
	if s == "" {
		return validationError
	}
	if s != "Ayubs" {
		return notFoundError
	}
	return nil
}

func main() {
	err := check("Ayubs")
	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println(validationError)
		} else if errors.Is(err, notFoundError) {
			fmt.Println(notFoundError)
		} else {
			fmt.Println("Unknown Error")
		}
	}
}
