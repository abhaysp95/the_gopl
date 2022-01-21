// Exercise: 4.2

package main

import (
	"flag"
	"fmt"

	// "crypto/sha256"
	"crypto/sha256"
	"crypto/sha512"
)

var mode = flag.String("m", "256", "Enter hash mode 384, 512 (256 is default)")

func main() {
  flag.Parse()
  var hash func(v []byte) []byte
  for _, arg := range flag.Args() {
    switch {
    case *mode == "384":
      hash = func(v []byte) []byte {
        h := sha512.Sum384(v)
        return h[:]
      }
    case *mode == "512":
      hash = func(v []byte) []byte {
        h := sha512.Sum512(v)
        return h[:]
      }
    default:
      hash = func(v []byte) []byte {
        h := sha256.Sum256(v)
        return h[:]
      }
    }
    fmt.Printf("%x, %[1]T\n", hash([]byte(arg)))
  }
}
