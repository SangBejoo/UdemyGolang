package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Printf("%s\n", arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Printf("%s\n", hostname)
	} else {
		fmt.Printf("%s\n", err)
	}
}
