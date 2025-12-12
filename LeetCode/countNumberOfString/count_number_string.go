package main

import "strconv"

func countNumeberOfStrings(s []string) int {
	count := 0
	for _, str := range s {
		if _, err := strconv.Atoi(str); err == nil {
			count++
		}
	}
	return count
}
func main() {
	data := []string{"1", "2", "hello", "world", "3"}
	result := countNumeberOfStrings(data)
	println(result) // Output: 3
}
