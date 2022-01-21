// dup3 prints count and text of line which occurred more than once in all the read files

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// NOTE: from v1.6 and above "ioutil" is supposed to be deprecated and you have
// to use io or os module now. All ioutil now contains is wrapper from these
// two. And example for so is shown in "dup4"(not mentioned in book)
