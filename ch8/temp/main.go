// temp is supposed to clear confusion regarding "done" channel from du4 example

package main

import (
	"fmt"
	"os"
	"time"
)

var done = make(chan struct{})

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		// done <- struct{}{}
		close(done)
	}()

loop:
	for {
		select {
		case <-done:
			fmt.Println("done closed")
			break loop
		default:
			fmt.Println("from default")
			time.Sleep(1 * time.Second)
		}
	}
}
