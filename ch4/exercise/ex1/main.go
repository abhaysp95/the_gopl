package main

import (
  "crypto/sha256"
  "fmt"
  "os"
)

func shaBitDiff(a, b []byte) int {
  c1 := sha256.Sum256(a)
  c2 := sha256.Sum256(b)
  return bitDiff(c1[:], c2[:])
}

func bitDiff(a, b []byte) int {
  count := 0
  for i := 0; i < len(a) || i < len(b); i++ {
    switch {
    case i >= len(a):
      count += popCount(b[i])
    case i >= len(b):
      count += popCount(a[i])
    default:
      count += popCount(a[i] ^ b[i])  // this will send bits only that are different
    }
  }
  return count
}

func popCount(x uint8) int {
  count := 0
  for ; x != 0; count++ {
    x &= (x - 1)
  }
  return count
}

func main() {
  for i := 1; i < len(os.Args); i+=2 {
    if (i + 1 == len(os.Args)) {
      fmt.Println("Enter pair for:", os.Args[i])
      os.Exit(1)
    }
    fmt.Println(shaBitDiff([]byte(os.Args[i]), []byte(os.Args[i + 1])))
  }
}

func firstApproach() {
  c1 := sha256.Sum256([]byte("x"))
  c2 := sha256.Sum256([]byte("X"))
  fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
  // fmt.Println(math.Abs(float64(setBits(c1)) - float64(setBits(c2))))  // or, if this feels a little too much, try below
  x := setBits(c1)
  y := setBits(c2)
  if x > y {
    fmt.Println(x - y)
  } else {
    fmt.Println(y - x)
  }
}

func setBits(arr [32]byte) int {
  c := 0
  for i := 0; i < 32; i++ {
    x := arr[i]
    for x != 0 {
      x &= (x - 1)
      c++
    }
  }
  return c
}
