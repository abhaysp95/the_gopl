// Exercise: 3.12

package main

import (
  "fmt"
  "os"
  "strings"
)

func main() {
  n := len(os.Args)
  if n < 2 {
    return
  }
  for i := 1; i < n; i+=2 {
    s1 := os.Args[i]
    if i + 1 == n {
      fmt.Printf("Input pair for %q\n", s1)
      return
    }
    s2 := os.Args[i + 1]
    fmt.Println(s1, s2, anagram2(s1, s2))
  }
}

// for the sake of using strings package
func anagram(s1 string, s2 string) bool {
  if (len(s1) != len(s2)) {
    return false
  }
  for _, r := range s2 {  // range returns rune
    if !strings.ContainsRune(s1, r) {  // if using tradition (C-style for index loop), use ContainsAny()
      return false
    }
  }
  return true
}

// using maps
func anagram2(s1 string, s2 string) bool {
  if len(s1) != len(s2) {
    return false
  }
  sfreq := make(map[rune]int)
  for _, r := range s1 {
    sfreq[r]++
  }
  for _, r := range s2 {
    sfreq[r]--
  }
  // for k := range sfreq {  // this returns key(I need value only)
  for _, v := range sfreq {
    if v != 0 {
      return false
    }
  }
  return true
}
