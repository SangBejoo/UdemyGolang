package GolangGoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// pool untuk menyimpan data yang bisa digunakan ulang
// New berfungsi untuk mengembalikan data awal ketika pool kosong
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New Data"
		},
	}

	pool.Put("Ayub")
	pool.Put("Subagiya")
	pool.Put("Golang")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("Data dari pool", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
