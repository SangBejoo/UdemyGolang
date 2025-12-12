package GolangGoroutine

import (
	"fmt"
	"strconv"
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

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello from OnlyIn"
}

func OnlyOut(channel <-chan string) {
	time.Sleep(2 * time.Second)
	data := <-channel
	fmt.Println("Received in OnlyOut:", data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// buffer capacity buffered channel

func TestBUfferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Ayub Subagiya"
		channel <- "Subagiya sent"
		channel <- "Buffered Channel"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Message 1 sent to buffered channel")
}

// range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel) // harus close channel agar range berhenti
	}()

	for data := range channel {
		fmt.Println("Received:", data)
	}

	fmt.Println("All data received from channel")
}

// select channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2:", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

// default select channel
func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2:", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}
		if counter == 2 {
			break
		}
	}
}

/// masalah race condition pada channel ketika mengirim data ke channel lebih cepat dari pada mengambil data dari channel
