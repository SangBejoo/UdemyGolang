package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Ayub subagiya, ini benar benar, belajar golang" +
		"\nAyub subagiya, ini benar benar, belajar golang" +
		"\nAyub subagiya, ini benar benar, belajar golang"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
