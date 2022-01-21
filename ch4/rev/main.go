package main

import "fmt"

func reverse(a []int) {  // recieved slice
  for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
    a[i], a[j] = a[j], a[i]
  }
}

func main() {
  a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}  // implicit array of right size and yields slice that points to it
  reverse(a[:])  // passing as slice
  fmt.Println(a)
}
