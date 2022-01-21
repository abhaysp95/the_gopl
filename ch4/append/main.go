package main

import "fmt"

/**
  x: int slice
  y: an int
  returns: int slice
*/
func appendInt(x []int, y int) []int {
  var z []int
  zlen := len(x) + 1
  if zlen <= cap(x) {
    // there is room to grow, extend the slice
    z = x[:zlen]
  } else {
    // there is insufficient space, allocating new array
    // grow by doubling, for amortized linear complexity
    zcap := zlen
    if zcap < 2 * len(x) {
      zcap = 2 * len(x)
    }
    z = make([]int, zlen, zcap)
    copy(z, x)
  }
  z[len(x)] = y
  return z
}

func appendIntVariadic(x []int, y ...int) []int {
  var z []int
  zlen := len(x) + len(y)
  if zlen <= cap(x) {
    z = x[:zlen]
  } else {
    zcap := zlen
    if zcap < 2 * len(x) {
      zcap = 2 * len(x)
    }
    z = make([]int, zlen, zcap)
  }
  copy(z[len(x):], y)
  return z
}

func main() {
  var x, y []int
  for i := 0; i < 10; i++ {
    y = appendInt(x, i)
    fmt.Printf("%d cap=%d, len=%d\t%v\n", i, len(y), cap(y), y)
    x = y
  }
}
