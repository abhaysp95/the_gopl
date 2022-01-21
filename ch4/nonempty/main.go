package main

import "fmt"

// nonempty returns the slice holding only the non-empty strings
// the underlying array is modified during the call
func nonempty(strings []string) []string {
  i := 0
  for _, s := range strings {
    if s != "" {
      strings[i] = s  // go doesn't support strings[i++] (just telling)
      i++
    }
  }
  return strings[:i]
}

// nonempty2 uses append
func nonempty2(strings []string) []string {
  out := strings[:0]  // zero-length slice of original
  for _, s := range strings {
    if s != "" {
      out = append(out, s)
    }
  }
  return out
}

// Above function raises the question, that how would "range" be affected if we
// are modifying the slice (same underlying array) on which "range" is
// iterating

func main() {
  data := []string{"one", "", "three"}
  fmt.Printf("%q\n", nonempty(data))
  fmt.Printf("%q\n", data)
}
