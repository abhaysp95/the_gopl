// Exercise: 4.5

package main

import "fmt"

func rmAdjDuplicate(arr []string) []string {
  c := 0
  for _, s := range arr {
    fmt.Println(c, arr[c], s)
    if arr[c] == s {
      continue
    }
    c++
    arr[c] = s
  }
  return arr[:c + 1]
}

func main() {
  str := []string{"this", "is", "this", "this", "good", "good", "good", "but", "bad", "bad", "too"}
  str = rmAdjDuplicate(str)
  fmt.Println(str)
}
