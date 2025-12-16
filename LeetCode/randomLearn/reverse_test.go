package randomLearn

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	data := []byte("ayub subagiya")
	size := len(data)

	for i := 0; i < size/2; i++ {
		data[i], data[size-1-i] = data[size-1-i], data[i]
	}
	fmt.Println(string(data))
}
