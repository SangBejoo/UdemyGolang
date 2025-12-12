package main

import (
	"math"
)

func Fibonacci(n int) []int {
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

func FibbonacciV2(n int) int {
	sqrt5 := math.Sqrt(5)
	return int((math.Round((math.Pow(1+sqrt5, float64(n)) - math.Pow(1-sqrt5, float64(n))) / (math.Pow(2, float64(n)) * sqrt5))))

}
