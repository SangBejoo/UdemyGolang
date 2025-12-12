package palindrom

func isPalindromeNumber(x int) bool {
	reversed := 0
	original := x

	if x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x = x / 10
	}

	return reversed == original
}
