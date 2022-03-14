package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))  // ignoring error
		abort <- struct{}{}
	}()

	fmt.Println("Commencing Launch. Press Return to Abort")
	select {
	case <-time.After(10 * time.Second):
		// do nothing
	case <-abort:
		fmt.Println("Launch Abort Succesful")
		multiplexingExample()
		return
	}

	launch()
}

func launch() {
	fmt.Println("Congratulations! Launch Successful")
}

func multiplexingExample() {
	ch := make(chan int, 1)  // either full or empty

	/**
	 *  If multiple cases are ready, select chooses one randomly so that every
	 *  channel has equal chance. Increasing the size of the above buffer makes
	 *  the output non-deterministic because the buffer the when the buffer is
	 *  neither full nor empty, select will figuratively toss a coin
	 */
	// ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)  // 0 2 4 6 8
		case ch <- i:
		}
	}
}
