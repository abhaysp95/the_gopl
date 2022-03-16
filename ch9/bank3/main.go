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
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
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
