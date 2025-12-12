package main

import (
	"fmt"
	"strconv"
)

func main() {
	iniString := "42"
	jadiInt, err := strconv.Atoi(iniString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jadiInt + 22)

	result, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	var stringInt string = strconv.Itoa(1000)
	fmt.Println(stringInt)
}
