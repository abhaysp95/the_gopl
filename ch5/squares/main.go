package main

import "fmt"

func square() func() int {
  var x int
  return func() int {
    x++
    return x * x
  }
}

func main() {
  f := square()
  for i := 0; i < 5; i++ {
    fmt.Println(f())
  }
}
