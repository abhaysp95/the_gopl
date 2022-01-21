// intsToString is like fmt.Sprintf(values) but add commas

package main

import (
  "bytes"
  "fmt"
)

func main() {
  fmt.Println(intsToString([]int{1, 2, 3, 4, 5}))
}

func intsToString(arr []int) string {
  var buf bytes.Buffer
  buf.WriteByte('[')
  for i, v := range arr {
    if i > 0 {
      buf.WriteByte(',')
    }
    fmt.Fprintf(&buf, "%d", v)
  }
  buf.WriteByte(']')
  return buf.String()
}
