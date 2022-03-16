package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	sema = make(chan struct{}, 1)  // a binary semaphore guarding balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{}  // acquire token
	balance += amount
	<-sema  // release token
}

func Balance() int {
	sema <- struct{}{}  // acquire token
	b := balance
	<-sema  // release token
	return b
}

func main() {
	var n sync.WaitGroup

	n.Add(1)
	go func() {
		Deposit(200)
		fmt.Fprintln(os.Stdout, "=", Balance())
		n.Done()
	}()

	n.Add(1)
	go func() {
		Deposit(100)
		n.Done()
	}()

	n.Wait()
}
