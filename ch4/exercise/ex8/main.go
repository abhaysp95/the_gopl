// Exercise: 4.8

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
  count := make(map[string]map[rune]int)
  // a map whose key is a rune and value is a map whose key is type of rune and
  // value is count
  var utfLen [utf8.UTFMax + 1]int
  var invalid uint
  in := bufio.NewReader(os.Stdin)
  for {
    r, s, e := in.ReadRune()
    if e == io.EOF {
      break
    }
    if e != nil {
      fmt.Fprintf(os.Stderr, "ex8: %v\n", e)
      os.Exit(1)
    }
    if r == unicode.ReplacementChar && s == 1 {
      invalid++
      continue
    }
    for name, table := range unicode.Properties {
      if unicode.In(r, table) {
        runeCount := count[name]
        if runeCount == nil {
          runeCount = make(map[rune]int)
          count[name] = runeCount
        }
        runeCount[r]++
        break  // I don't think a rune will belong in more than one table, can it ? (check that)
      }
    }
    utfLen[s]++
  }
  fmt.Println("Type of Rune\tRune\tRune Count")
  for name, runeCount := range count {
    for r, c := range runeCount {
      fmt.Printf("%s\t%q\t%d\n", name, r, c)
    }
  }
  fmt.Println("\nRune len\tCount")
  for i, s := range utfLen {
    fmt.Printf("%d\t%d\n", i, s)
  }
  fmt.Printf("\nInvalid: %d\n", invalid)
}
