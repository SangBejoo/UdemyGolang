package palindrom

import (
	"fmt"
	"unicode"
)

func IsPalindrome(s string) bool {
	var filtered []rune
	for _, chara := range s {
		if unicode.IsLetter(chara) || unicode.IsDigit(chara) {
			filtered = append(filtered, unicode.ToLower(chara))
		}
	}

	// ini untuk ngecek palindrom
	left, right := 0, len(filtered)-1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome("Kasur Rusak"))
}
