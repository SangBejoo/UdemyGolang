package GolangContext

import (
	"context"
	"time"
)

// function membuat context

/*
context.Background() membuat context kosong tidak pernah dibatalkan, tidak pernah timeout dan tidak memiliki value apapun
biasanya di gunakan pada main function atau dalam test atau awal proses request terjadi

context.Todo() membuat context kosong seperti background(), namun biasanya menggunakan ini ketika belum jelas context apa yang diinginkan

*/

// goroutine Leak di perbaiki dengan context penambahan select
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // simulasi slow
			}
		}
	}()
	return destination
}

func CreateCounter1() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
			time.Sleep(1 * time.Second) // simulasi slow
		}

	}()
	return destination
}
