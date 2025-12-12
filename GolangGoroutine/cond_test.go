package GolangGoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var lock = sync.Mutex{}
var cond = sync.NewCond(&lock)
var group = sync.WaitGroup{}

func waitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Goroutine", value, "berjalan")

	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go waitCondition(i)
	}
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	cond.Broadcast()
	//}()
	group.Wait()
}
