package main

// ReverseStringV2 reverses a slice of bytes in place.
// jika n lebih dari k maka di reverse semua

func ReverseStringV2(s string, k int) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n; i += 2 * k {
		end := i + k
		if end > n {
			end = n
		}
		reverse(runes[i:end])
	}

	return string(runes)
}

// pembalikkan
func reverse(runes []rune) {
	size := len(runes)

	for i := 0; i < size/2; i++ {
		runes[i], runes[size-1-i] = runes[size-1-i], runes[i]
	}
}

func main() {
	// Example usage
	s := "abcdefghij"
	k := 3
	result := ReverseStringV2(s, k)
	println(result) // Output: "cbadefihgj"
}
