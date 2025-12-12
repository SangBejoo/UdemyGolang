package main

import "testing"

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte("hello"), []byte("olleh")},
	}

	for _, tc := range testCases {
		t.Run(string(tc.input), func(t *testing.T) {
			ReverseString(tc.input)
			for i := range tc.input {
				if tc.input[i] != tc.expected[i] {
					t.Errorf("ReverseString() = %s; want %s", string(tc.input), string(tc.expected))
					return
				}
			}
			t.Logf("ReverseString() = %s; as expected", string(tc.input))
		})
	}
}
