package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value" + strconv.Itoa(i)
		data = data.Next()
	}

	data.Value = "Value 1"
	data = data.Next()
	data.Value = "Value 2"
	data = data.Next()
	data.Value = "Value 3"
	data = data.Next()
	data.Value = "Value 4"

	data.Do(func(a any) {
		fmt.Println(a)
	})
}
