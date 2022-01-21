package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		in := input.Text()
		if in == "q" {
			return
		}
		x, _ := strconv.ParseInt(in, 0, 64)
		start := time.Now()
		fmt.Println(popCount(x), time.Since(start))
	}
}

func popCount(x int64) int {
	c := 0
	for x != 0 {
		x &= (x - 1)
		c++
	}
	return c
}
