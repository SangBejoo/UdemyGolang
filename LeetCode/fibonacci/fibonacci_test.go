package main

import "testing"

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{5, []int{0, 1, 1, 2, 3, 5}},
		{0, []int{}},
		{1, []int{0, 1}},
		{2, []int{0, 1, 1}},
	}
	for _, tc := range testCases {
		result := Fibonacci(tc.input)
		if len(result) != len(tc.expected) {
			t.Errorf("Fibonacci(%d) = %v; want %v", tc.input, result, tc.expected)
			continue
		}
		for i := range result {
			if result[i] != tc.expected[i] {
				t.Errorf("Fibonacci(%d) = %v; want %v", tc.input, result, tc.expected)
				break
			}
		}
	}
}
