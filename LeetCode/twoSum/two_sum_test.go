package main

import "testing"

func TestTwoSum(t *testing.T) {
	test := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, value := range test {
		result := TwoSum(value.nums, value.target)
		if len(result) != len(value.expected) {
			t.Errorf("TwoSum(%v, %d) = %v; want %v", value.nums, value.target, result, value.expected)
			continue
		}
		for i := range result {
			if result[i] != value.expected[i] {
				t.Errorf("TwoSum(%v, %d) = %v; want %v", value.nums, value.target, result, value.expected)
				break
			}
		}
	}

}
