package palindrom

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Kasur Rusak", true},
		{"Anmapana", false},
	}
	for _, test := range tests {
		result := IsPalindrome(test.input)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", test.input, result, test.expected)
		} else {
			t.Logf("IsPalindrome(%q) = %v; as expected", test.input, result)
		}
	}

}

func TestIsPalindromeNumber(t *testing.T) {
	tests := []struct {
		x        int
		expected bool
	}{
		{12321, false},
		{12345, false},
	}
	for _, test := range tests {
		result := isPalindromeNumber(test.x)
		if result != test.expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", test.x, result, test.expected)
		}
	}
}
