// Exercise: 4.9

package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  freq := make(map[string]uint)
  in := bufio.NewScanner(os.Stdin)
  in.Split(bufio.ScanWords)  // default splitFunc for Scanner is ScanLines
  for in.Scan() {
    freq[in.Text()]++
  }
  if in.Err() != nil {
    fmt.Fprintf(os.Stderr, "ex9: %v\n", in.Err())
  }
  fmt.Println("word\t\tfrequency")
  for word, fr := range freq {
    fmt.Printf("%q\t%d\n", word, fr)
  }
}
