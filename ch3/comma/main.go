// comma inserts comma in non-negative decimal integer strings

package main

import (
	"fmt"
	"os"
)

func main() {
  for _, arg := range os.Args[1:] {
    fmt.Println(comma(arg))
  }
}

func comma(s string) string {
  n := len(s)
  if n <= 3 {
    return s
  }
  return comma(s[:n - 3]) + "," + s[n - 3:]
}
