package main

import (
	"fmt"
	"strconv"
)

func CountNumbers(s string) int {
	count := 0
	for _, char := range s {
		if _, err := strconv.Atoi(string(char)); err == nil {
			count++
		}

	}
	return count
}

func main() {
	input := "abc123def456"
	result := CountNumbers(input)
	fmt.Println("Number of numeric characters:", result)
}
