package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount  // critical section (portion between mutex lock and unlock)
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {
	var n sync.WaitGroup

	n.Add(1)
	go func() {
		defer n.Done()
		Deposit(200)
		fmt.Println("=", Balance())
	}()

	n.Add(1)
	go func() {
		defer n.Done()
		Deposit(100)
	}()

	n.Wait()
	fmt.Println("Exiting")
}
