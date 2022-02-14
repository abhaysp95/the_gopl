package main

import (
  "flag"
  "fmt"
  "time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
  flag.Parse()
  fmt.Println("Sleeping for...", *period)
  time.Sleep(*period)
  fmt.Println()
}
