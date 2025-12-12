package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2025, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())
	fmt.Println(utc)

	formatter := "2006-01-02 15:04:05"

	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valueTime)
	}
	fmt.Println(formatter)

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())

	// duration
	duration := time.Second * 100
	fmt.Println(duration)
	duration2 := time.Minute * 10
	fmt.Println(duration2)
	duration3 := time.Second * 10
	fmt.Println(duration3)

	duration4 := duration3 - duration2*duration
	fmt.Println(duration4)
	fmt.Printf("duration : %d\n", duration4)
}
