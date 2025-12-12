package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"Ayub Subagiya", "Belajar Golang"})
	_ = writer.Write([]string{"Ayub Subagiya", "Belajar Golang"})
	_ = writer.Write([]string{"Ayub Subagiya", "Belajar Golang"})
	_ = writer.Write([]string{"Ayub Subagiya", "Belajar Golang"})
	writer.Flush()
}
