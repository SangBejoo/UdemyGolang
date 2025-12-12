package main

import (
	"fmt"
	"math"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 20, 6, 7, 8, 9, 10}
	fmt.Println(data)

	maxVal := float64(data[0])
	for _, v := range data {
		maxVal = math.Max(maxVal, float64(v))
	}
	fmt.Println(maxVal)

	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.40))
	fmt.Println(math.Max(1.40, 1.20))
	fmt.Println(math.Min(1.40, 1.90))

}
