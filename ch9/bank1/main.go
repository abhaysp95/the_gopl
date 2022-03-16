// one of the ways to avoid race condition is "variable confinement"
package main

import (
	"fmt"
	"sync"
)

var deposites = make(chan int)  // send amount to deposit
var balances = make(chan int)  // recieve balance

func Deposite(amount int) {
	deposites <- amount
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int  // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposites:
			fmt.Println("deposite balance: ", amount)
			balance += amount
		case balances <- balance:
			fmt.Println("balance updated")
		}
	}
}

func init() {
	go teller()
}

func main() {
	var n sync.WaitGroup
	go func() {
		defer n.Done()
		n.Add(1)
		Deposite(200)
		fmt.Println("=", Balance())
	}()

	go func() {
		defer n.Done()
		n.Add(1)
		Deposite(100)
	}()

	n.Wait()
	close(deposites)
	close(balances)
}
