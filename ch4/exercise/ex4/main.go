package main

import "fmt"

func rotateLeft(arr []int, x int) {
  x = x % len(arr)
  if x > 0 && x < len(arr) {
    temp := make([]int, x)
    copy(temp, arr[:x])
    copy(arr, arr[x:])
    copy(arr[len(arr) - x:], temp)
  }
  // return arr
}

// another version (this will need return)
func rotateLeft2(arr []int, x int) []int {
  x = x % len(arr)
  if x > 0 && x < len(arr) {
    arr = append(arr[x:], arr[:x]...)
    // because you are changing the
  }
  return arr
}

// in the same way, rotateRight can be written using both methods
func rotateRight(arr []int, x int) {
  x = x % len(arr)
  if x > 0 && x < len(arr) {
    r := len(arr) - x
    temp := make([]int, x)
    copy(temp, arr[r:])
    copy(arr[x:], arr[:r])
    copy(arr, temp)
  }
}

func rotateRight2(arr []int, x int) []int {
  x = x % len(arr)
  if x > 0 && x < len(arr) {
    arr = append(arr[len(arr) - x:], arr[:len(arr) - x]...)
  }
  return arr
}

func main() {
  a := []int{1, 2, 3, 4, 5, 6}
  // rotateLeft(a, 2)
  a = rotateLeft2(a, 2)
  fmt.Printf("%v\n", a)
  // a = rotateRight2(a, 2)
  rotateRight(a, 2)
  fmt.Printf("%v\n", a)  // you should get back the original array
}
