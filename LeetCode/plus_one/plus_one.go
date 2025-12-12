package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	return newDigits
}

func main() {
	// Example usage
	result := plusOne([]int{9, 9, 9})
	fmt.Println(result) // Output: [1, 0, 0, 0]

}
