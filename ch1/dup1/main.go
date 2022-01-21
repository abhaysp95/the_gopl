// dup1 prints the text of each line which appear more than once in the
// standard input, preceeded by its count

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)  // map with key "string" and value "int"
	input := bufio.NewScanner(os.Stdin)
	//fmt.Printf("type[input]: %T\n", input)
	for input.Scan() {  // ignoring potential errors from input.Err()
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// NOTE: press Ctrl+d to cancel input(without terminating program) in
// linux/unix, and Ctrl+z in windows
