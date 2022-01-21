package main

import (
	"fmt"
	"time"

	"the_gopl/ch2/popcount"
)

func PopCountRewrite(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(popcount.PC[byte(x >> (i * 8))])
	}
	return sum
}

func main() {
	var n uint64 = 84
	//fmt.Printf("%v, %v, %v\n", time.Now(), time.Now().UTC(), time.Now().UTC().UnixNano())
	start := time.Now()
	fmt.Println(popcount.PopCount(n))
	fmt.Printf("Time cosumed, (book program): %v\n", time.Since(start))
	start = time.Now()
	fmt.Println(PopCountRewrite(n))
	fmt.Printf("Time consumed, (exercise): %v\n", time.Since(start))
}
