package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func addToFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	// createNewFile("test.log", "Hello world")

	//result, _ := readFile("test.log")
	//fmt.Println(result)

	addToFile("test.log", "\n ths is log line 1\n")
}
