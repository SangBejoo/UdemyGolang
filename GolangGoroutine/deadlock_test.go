package GolangGoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}
func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}
func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user", user1.Name)
	user1.Change(-amount)

	time.Sleep(1) // simulate delay

	user2.Lock()
	fmt.Println("Lock user", user2.Name)
	user2.Change(amount)

	time.Sleep(1) // simulate delay

	user2.Unlock()
	user1.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Ayub Subagiya",
		Balance: 100000,
	}
	user2 := UserBalance{
		Name:    "Subagiya Ayubs",
		Balance: 100000,
	}

	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 20000)

	time.Sleep(3 * time.Second)

	fmt.Println("user ", user1.Name, "balance", user1.Balance)
	fmt.Println("user ", user2.Name, "balance", user2.Balance)
}
