package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		// fmt.Println(1)
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	/* go func() {
		for {
			x, ok := <- naturals; !ok {
				break  // naturals is closed
			}
			squares <- x*x
		}
		close(squares)
	}() */
	// above syntax (inside function) is clumsy but pattern is common, the
	// language lets us use 'range' loop over channel too
	go func() {
		// fmt.Println(2)
		for x := range naturals {
			squares <- x*x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	// fmt.Println(3)
	for x := range squares {
		fmt.Println(x)
	}
	// time.Sleep(1 * time.Second)  // use this by uncommenting, above 'fmt's and commenting out Printer loop
}
