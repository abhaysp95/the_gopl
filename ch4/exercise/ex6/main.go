// Exercise: 4.6

package main

import (
  "unicode/utf8"
  "unicode"
)

func squashSpace(arr []byte) []byte {
  out := arr[:0]
  var last rune
  for i := 0; i < len(arr); i++ {
    r, s := utf8.DecodeRune(arr[i:])
    if !unicode.IsSpace(r) {
      out = append(out, arr[i:i+s]...)  // append uses byte here
    } else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
      // ASCII space literal, ('', which is an untyped constant) will change
      // into byte when transferred as append's parameter
      // ' ' is rune/int32, character in Go is unicode, but string will be utf8
      // encoded
      // thus, you can just put append(out, arr[i:i+s]...) here too and it'll
      // work
      out = append(out, ' ')
    }
    last = r
    i += s
  }
  return out
}
