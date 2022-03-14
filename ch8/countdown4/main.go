package main

import (
	"fmt"
	"os"
	"time"
)

/** Problem with countdown3 was the use of time.Tick(). Even though goroutine
* (containing it/could be a function in a thread) exited and recieving events
* from it has been stopped, but the ticker is still there, trying in vain to
* send on a channel from which no goroutine is recieving - a goroutine leak */

// The Tick() is convinient but appropriate only when the ticks will be needed
// throughout the lifetime of the application
func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing Launch. Press Return to abort countdown")

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()  // I suppose this is right

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-ticker.C:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("Aborting launch")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Congratulations! Launch Successful")
}
