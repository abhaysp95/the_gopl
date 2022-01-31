package main

import (
	"fmt"
	"os"
	"runtime"
)

func f(x int) {
  fmt.Printf("f(%d)\n", x+0/x)
  defer fmt.Printf("defer %d\n", x)
  f(x - 1)
}

func printStack() func() {
  fmt.Println("==> printStack")
  var buf [4096]byte
  n := runtime.Stack(buf[:], false)
  return func() {
    fmt.Println("==> write to Stdout")
    os.Stdout.Write(buf[:n])
  }
}

func main() {
  // defer printStack()  // made the change because of the test I wanted to do
  defer printStack()()
  f(3)
}
