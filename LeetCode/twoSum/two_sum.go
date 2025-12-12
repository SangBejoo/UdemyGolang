package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9), "with Target 9")
	fmt.Println(twoSum([]int{3, 2, 4}, 6), "with Target 6")
	fmt.Println(twoSum([]int{3, 3}, 6), "with Target 6")
}

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)
	for i, num := range nums {
		ans := target - num
		if j, found := mapping[ans]; found {
			return []int{j, i}
		}
		mapping[num] = i
	}
	return []int{}
}

// two nums
func twoNums(nums []int, target int) []int {
	mapping := make(map[int]int)

	for i, num := range nums {
		ans := target - num
		if j, found := mapping[ans]; found {
			return []int{j, i}
		}
		mapping[num] = i
	}
	return []int{}
}

// best time to buy and sell stock
func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, price := range prices {
		minPrice = min(price, minPrice)
		maxProfit = max(maxProfit, price-minPrice)
	}
	return maxProfit
}
