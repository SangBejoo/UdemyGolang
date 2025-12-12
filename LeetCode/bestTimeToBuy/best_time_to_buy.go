package main

import "math"

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, price := range prices {
		minPrice = min(price, minPrice)
		maxProfit = max(maxProfit, price-minPrice)
	}
	return maxProfit
}
