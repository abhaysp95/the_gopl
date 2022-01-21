package main

import "fmt"

func reverse(arr *[5]int) {
  for i, j := 0, len(*arr) - 1; i < j; i, j = i + 1, j - 1 {
    arr[i], arr[j] = arr[j], arr[i]
  }
}

func main() {
  a := [5]int{1, 2, 3, 4, 5}
  reverse(&a)  // even though this is an array, to pass by reference you have to use & (unlike C/C++)
  fmt.Println(a)
}
