package main

import (
  "fmt"
  "os"

  "the_gopl/ch5/links"
)

func test() {
  for _, url := range os.Args[1:] {
    l, err := links.Extract(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "test: %v\n", err)
      continue  // or maybe exit
    }
    fmt.Println(l)
  }
}

func main() {
  str := []string{"this"}
  count := 0
  for len(str) > 0 {
    item := str
    str = nil
    str = append(str, []string{"that", "these", "those"}...)
    fmt.Println(item, str)
    count++
  }
  fmt.Printf("count: %d\n", count)
}
