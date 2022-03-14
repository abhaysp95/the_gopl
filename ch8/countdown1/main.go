package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown")
	tick := time.Tick(2 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println("countdown")
		<-tick  // waits for 2 second (because of duration) and release the Ticker
	}
	// <-tick
	launch()
}

func launch() {
	fmt.Println("Congratulations! Launch Successful")
}
