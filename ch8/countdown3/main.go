package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))  // this goroutine waits for input
		abort <- struct{}{}  // then waits for the abort channel to be empty again by select
	}()

	fmt.Println("Commencing countdown. Press Return to Abort")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		select {
		// case <-time.Tick(1 * time.Second):
		case <-tick:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("Aborting Launch sequence")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Congratulations! Launch Successful")
}
