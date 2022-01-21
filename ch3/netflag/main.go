package main

import "fmt"

type Flags uint

const (
  FlagUp Flags = 1 << iota // is up
  FlagBroadcast            // supports broadcast access capability
  FlagLoopback             // is a loopback interface
  FlagPointToPoint         // belongs to a point-to-point link
  FlagMulticast            // supports multicast access capability
)

func IsUp(v Flags) bool { return v & FlagUp == FlagUp }
func TurnDown(v *Flags) { *v &^= FlagUp }  // or do *v = *v & (~y)  // ~ is not (I've to check if golang supports it)
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool { return v & (FlagBroadcast | FlagMulticast) != 0 }

func main() {
  var v Flags = FlagUp | FlagMulticast
  fmt.Printf("%b %t\n", v, IsUp(v))
  TurnDown(&v)
  fmt.Printf("%b %t\n", v, IsUp(v))
  SetBroadcast(&v)
  fmt.Printf("%b %t\n", v, IsUp(v))
  fmt.Printf("%b %t\n", v, IsCast(v))
}
