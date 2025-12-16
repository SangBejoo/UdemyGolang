package randomLearn

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSetList(t *testing.T) {
	set := []string{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			set = append(set, "ini buku genap")
		}
		set = append(set, "ini buku ke "+strconv.Itoa(i))
	}
	fmt.Println(set)
}

func TestFindMax(t *testing.T) {
	data := []int{13213, 223131, 304, 500, 6121, 72332, 3122318, 912313213}
	maxValue := data[0]
	index := 0
	for i, value := range data {
		if value > maxValue {
			maxValue = value
			index = i
		}
	}
	fmt.Println(maxValue, index)
}

func TestFrekuensi(t *testing.T) {
	data := []string{"kucing", "anjing", "kucing", "ikan", "anjing", "kucing"}
	seen := make(map[string]int)
	for _, item := range data {
		seen[item]++
	}
	fmt.Println(seen)
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Println(i, j)
			}
		}
	}
}
