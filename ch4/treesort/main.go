package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type tree struct  {
  value int
  left, right *tree
}

// sorts values in place
func sort(values []int) {
  var root *tree
  for _, node := range values {
    root = add(root, node)
  }
  fmt.Printf("(up): %s\n", getSliceHeader(&values))
  // fmt.Printf("%s\n", getSliceHeader(&values[:0]))
  values = appendValues(values, root)
  fmt.Printf("(down): %s\n", getSliceHeader(&values))
}

// appends elements of t to values in order and returns the resulting slice
// (inorder traversal)
func appendValues(values []int, t *tree) []int {
  if t != nil {
    values = appendValues(values, t.left)
    fmt.Printf("(after left): %s\n", getSliceHeader(&values))
    values = append(values, t.value)  // append will change values address
    fmt.Printf("(after append): %s\n", getSliceHeader(&values))
    values = appendValues(values, t.right)
    fmt.Printf("(after right): %s\n", getSliceHeader(&values))
  }
  fmt.Printf("(before return): %s\n", getSliceHeader(&values))
  return values
}


func getSliceHeader(slice *[]int) string {
  sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
  return fmt.Sprintf("%#x, %+v", sh.Data, sh)
}

// these extra prints are for debug only (I need to learn of go debuggin soon)
func add(t *tree, value int) *tree {
  if t == nil {
    // Equivalent to return &tree{value: value}
    t = new(tree)
    t.value = value
    // fmt.Println("returning with value(up):", t.value)
    return t
  }
  if t.value > value {
    // fmt.Println("going left")
    t.left = add(t.left, value)
    // fmt.Println("set value(left):", t.left.value)
  } else {
    // fmt.Println("going left")
    t.right = add(t.right, value)
    // fmt.Println("set value(right):", t.right.value)
  }
  // fmt.Println("returning with value(down):", t.value)
  return t
}

func main() {
  arr := []int{10, 3, 5, 6, 9, 7, 2, 8}
  sort(arr)
  fmt.Println(arr)
}
