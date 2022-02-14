package tempconv

import (
	"flag"
	"fmt"
	"the_gopl/ch2/tempconv"
)

// *celciusFlag statisfies Flag.Value interface
type celciusFlag struct { tempconv.Celsius }

func (f *celciusFlag) Set(s string) error {
  var unit string
  var value float64
  fmt.Sscanf(s, "%f%s", &value, &unit)
  switch unit {
  case "C", "°C":
    f.Celsius = tempconv.Celsius(value)
    return nil
  case "F", "°F":
    f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
    return nil
  case "K", "°K":
    f.Celsius = tempconv.KToC(tempconv.Kelvin(value))
  }
  return fmt.Errorf("invalid temperature %q", s)
}

// CelciusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and unit, e.g., "100C"
func CelciusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
  f := celciusFlag{value}
  flag.CommandLine.Var(&f, name, usage)
  return &f.Celsius
}
