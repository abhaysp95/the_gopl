package main

import (
	"fmt"
	"unicode/utf8"
)

// version1
func revExtraSpace(arr []byte) []byte {
  out := make([]byte, 0, len(arr))
  sz := len(arr)
  for sz > 0 {
    _, s := utf8.DecodeLastRune(arr[:sz])
    out = append(out, arr[sz-s:sz]...)
    sz -= s
  }
  // now either return out or copy out to arr
  // I'll return it
  return out
}

// version2
func revNoExtraSpace(arr []byte) {
  // first, treat it as regular byte slice and reverse it
  for i, j := 0, len(arr) - 1; i < j; i, j = i + 1, j - 1 {
    arr[i], arr[j] = arr[j], arr[i]
  }

  // Now, encode them to rune
  for i := 0; i < len(arr); {
    var twoRune, threeRune, fourRune bool
    for {
      // check within the loop if you are getting valid rune, first decode if
      // not right than swap the next byte and then check and decode if not
      // right, then do for 3 and again if not right, then do for four
      r, s := utf8.DecodeRune(arr[i:])
      if r != utf8.RuneError {
        i += s
        break
      } else {
        // check for twoRune
        if !twoRune {
          twoRune = true
          arr[i], arr[i + 1] = arr[i + 1], arr[i]
          continue
        }
        if !threeRune {
          // make effect of twoRune from previous iteration null
          arr[i], arr[i + 1] = arr[i + 1], arr[i]
          threeRune = true
          arr[i], arr[i + 2] = arr[i + 2], arr[i]
          continue
        }
        if !fourRune {
          // make effect of threeRune from previous iteration null
          arr[i], arr[i + 2] = arr[i + 2], arr[i]
          fourRune = true
          arr[i], arr[i+1], arr[i+2], arr[i+3] = arr[i+3], arr[i+2], arr[i+1], arr[i]
          continue
        }
        panic("Found RuneError")
      }
    }
  }
}

func rev(arr []byte) {
  l := len(arr)
  for i := 0; i < l / 2; i++ {
    arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
  }
}

// version 3
func revRune(arr []byte) {
  for i := 0; i < len(arr); {
    _, s := utf8.DecodeRune(arr[i:])
    rev(arr[i:i+s])
    i += s
  }
  rev(arr)
}

func main() {
  str := []byte("hello 世界 world!")
  fmt.Println("Original: ", string(str))
  str = revExtraSpace(str)
  fmt.Println("Reverse (extra space):", string(str))
  revNoExtraSpace(str)
  fmt.Println("Reverse (no extra space):", string(str))
  revRune(str)
  fmt.Println("Reverse (revRune):", string(str))
}
