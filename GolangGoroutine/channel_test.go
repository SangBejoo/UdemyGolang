package GolangGoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello from channel"
		fmt.Println("Message sent to channel")
	}()

	// cara ambil
	data := <-channel
	fmt.Println("Received:", data)

	// beri waktu untuk goroutine selesai sebelum fungsi utama selesai
	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello from GiveMeResponse"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println("Received:", data)
	// beri waktu untuk goroutine selesai sebelum fungsi utama selesai
	time.Sleep(5 * time.Second)
}
