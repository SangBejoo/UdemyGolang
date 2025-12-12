package main

import "testing"

func TestBestTimeToBuy(t *testing.T) {
	testCases := []struct {
		prices []int
		expect int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{3, 2, 6, 5, 0, 3}, 4},
	}

	for _, tc := range testCases {
		result := maxProfit(tc.prices)
		if result != tc.expect {
			t.Errorf("maxProfit(%v) = %d; want %d", tc.prices, result, tc.expect)
		} else {
			t.Logf("maxProfit(%v) = %d; as expected", tc.prices, result)
		}
	}
}
