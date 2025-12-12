package GolangContext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)

}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextG.Value("g"))
	fmt.Println(contextF.Value("e"))

	fmt.Println(contextA.Value("f"))
}

func TestContextWithCancel1(t *testing.T) {
	fmt.Println("Total gorutine:", runtime.NumGoroutine())

	//parent := context.Background()
	//ctx, cancel := context.WithCancel(parent) // tambah ini

	destination := CreateCounter1()
	fmt.Println("Total gorutine:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println(n)
		if n == 10 {
			break

		}
	}
	//cancel() // mengirim sinyal cancel ke context

	//time.Sleep(2 * time.Second)
	fmt.Println("Total gorutine:", runtime.NumGoroutine())

}

// ini namanya goroutine leak dimana ada zombie goroutine
func TestContextWithCancel2(t *testing.T) {
	fmt.Println("Total gorutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent) // tambah ini

	destination := CreateCounter(ctx)
	fmt.Println("Total gorutine:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println(n)
		if n == 10 {
			break

		}
	}
	cancel() // mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)
	fmt.Println("Total gorutine:", runtime.NumGoroutine())

}

// ini sangat penting
func TestContextWithTimeOut(t *testing.T) {

	fmt.Println("Total gorutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second) // tambah ini
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total gorutine:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println(n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total gorutine:", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {

	fmt.Println("Total gorutine:", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second)) // tambah ini
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total gorutine:", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println(n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total gorutine:", runtime.NumGoroutine())
}
