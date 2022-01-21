// basename functionality

package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  for _, arg := range os.Args[1:] {
    fmt.Println(basename2(arg))
  }
}

func basename2(s string) string {
  idx := strings.LastIndex(s, "/")
  s = s[idx + 1:]
  // return s[:strings.LastIndex(s, ".")]  // note that, go doesn't support negative slicing like python [:-1] etc.
  if idx = strings.LastIndex(s, "."); idx >= 0 {  // it isn't necessary to always declare the variable for a scope (you can use existing ones too)
    s = s[:idx]
  }
  return s
}
