package main

import "fmt"

func fibbo(n int) []int {
	if n <= 0 {
		return []int{}
	}
	a, b := 0, 1
	result := []int{0}

	result = append(result, 1)
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
		result = append(result, b)
	}
	return result
}

func main() {
	fmt.Print(fibbo(10))
}
