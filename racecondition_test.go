package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 1; i <= 10000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Counter: ", x)
}

func TestMutex(t *testing.T) {
	x := 0

	var mutex sync.Mutex
	for i := 1; i <= 10000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Counter: ", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.RWMutex.Lock()
	defer b.RWMutex.Unlock()
	b.Balance += amount
}

func (b *BankAccount) GetBalance() int {
	b.RWMutex.RLock()
	defer b.RWMutex.RUnlock()
	return b.Balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 1; i <= 100; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				account.Deposit(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Balance: ", account.GetBalance())
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Hello, World!")
		}()
	}

	wg.Wait()
	fmt.Println("Done")
}
