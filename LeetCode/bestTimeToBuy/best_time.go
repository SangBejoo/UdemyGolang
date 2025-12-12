package main

import (
	"fmt"
	"math"
)

func MaxProfit(prices []int) int {
	minPrice := math.MaxInt32
	profit := 0

	for _, price := range prices {
		minPrice = min(price, minPrice)
		profit = max(profit, price-minPrice)
	}
	return profit
}

func timeToBuy(prices []int) int {
	minPrice := math.MaxInt32
	profit := 0

	for _, price := range prices {
		minPrice = min(price, minPrice)
		fmt.Println("Current Min Price:", minPrice)
		profit = max(profit, price-minPrice)
		fmt.Println("Current Profit:", profit)

	}
	return profit
}

func main() {
	fmt.Print(MaxProfit([]int{7, 1, 5, 3, 6, 4}), "\n")
	fmt.Print(timeToBuy([]int{7, 1, 5, 3, 6, 4, 9, 10, 2}))
}
