package main

import (
	"flag"
	"fmt"
	"the_gopl/ch7/tempconv"
)

var temp = tempconv.CelciusFlag("temp", 20.0, "the temperature (with unit)")

func main() {
  flag.Parse()
  fmt.Println("Temp is:", *temp)
}
