// Exercise: 3.10

package main

import (
  "bytes"
  "fmt"
  "os"
)

func main() {
  for _, arg := range os.Args[1:] {
    fmt.Println(comma(arg))
  }
}

func comma(s string) string {
  var buf bytes.Buffer
  // for i, c := range s[] {  // read the doc for range and see if it can go reverse like happens in python
  n := len(s)
  rem := n % 3
  buf.WriteString(s[:rem])
  for i := 0; i < n - rem; i++ {
    if i % 3 == 0  {
      if buf.Len() != 0 {
        buf.WriteByte(',')
      }
    }
    buf.WriteByte(s[i + rem])
  }
  return buf.String()
}


// range splits string in Rune, but indexing returns byte
