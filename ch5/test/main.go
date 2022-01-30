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

func loopTest() {
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

func scopeTestRangeForInvalid() {
  var lc []func()
  arr := []int{1, 2, 3, 4, 5}
  for _, v := range arr {
    fmt.Println("before:", v)
    lc = append(lc, func() {
      fmt.Println("after:", v)
    })
  }

  for _, c := range lc {
    c()
  }
}

func scopeTestRangeForValid() {
  var lc []func()
  arr := []int{1, 2, 3, 4, 5, 6}
  for _, v := range arr {
    i := v
    fmt.Println("before:", i)
    lc = append(lc, func() {
      fmt.Println("after:", i)
    })
  }

  for _, c := range lc {
    c()
  }
}

func main() {
  scopeTestRangeForInvalid()
}

/** same can be said for other types of for loop, Range based for loop is just an example */
