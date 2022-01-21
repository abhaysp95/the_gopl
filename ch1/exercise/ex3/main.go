// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	method1()
	t := time.Now()
	fmt.Println("Time elapsed:", t.Sub(start))

	start = time.Now()
	method2()
	t = time.Now()
	fmt.Println("Time elapsed:", t.Sub(start))
}

// inefficient version
func method1() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// comparatively efficient version
func method2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
