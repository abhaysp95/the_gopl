// Modify the echo program to also print os.Args[0], the name of the command
// that invoked it.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Printf("Running \"%s\"\n", os.Args[0])
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
