package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// encoding
	value := "Ayub Subagiya"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(string(decoded))
	}
}
