// basename removes directory component and a . suffix

package main

import (
  "fmt"
  "os"
)

func main() {
  for _, arg := range os.Args[1:] {
    fmt.Println(basename(arg))
  }
}


func basename(s string) string {
  // preserve everything before last '/'
  for i := len(s) - 1; i >= 0; i-- {
    if s[i] == '/' {
      s = s[i + 1:]
      break
    }
  }
  // preserve everything before last '.'
  for i:= len(s) - 1; i >= 0; i-- {
    if s[i] == '.' {
      s = s[:i]
    }
  }
  return s
}
