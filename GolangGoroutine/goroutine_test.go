package GolangGoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups, I am not a goroutine!")

	time.Sleep(1 * time.Second) // Wait for the goroutine to finish
}

func DisplayNumber(number int) {
	fmt.Printf("Number: %d\n", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second) // Wait for goroutines to finish

}
