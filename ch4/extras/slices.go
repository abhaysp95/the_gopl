// https://www.tugberkugurlu.com/archive/working-with-slices-in-go-golang-understanding-how-append-copy-and-slicing-syntax-work

package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

// How append and copy works

func getSliceHeader(slice *[]int) string {
  sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
  return fmt.Sprintf("%#x, %+v", sh.Data, sh)
}

func appendBehindTheScene() {
  var result []int
  for i := 0; i < 10; i++ {
    if i % 2 == 0 {
      fmt.Printf("appending '%d': %s\n", i, getSliceHeader(&result))
      result = append(result, i)
      fmt.Printf("appended '%d': %s\n", i, getSliceHeader(&result))
    }
  }
  fmt.Println(result)
}

func appendBehindTheScene2() {
  maxValue := 10
  result := make([]int, 0, maxValue)
  for i := 0; i < maxValue; i++ {
    if i  % 2 == 0 {
      fmt.Printf("appending '%d': %s\n", i, getSliceHeader(&result))
      result = append(result, i)
      fmt.Printf("appended '%d': %s\n", i, getSliceHeader(&result))
    }
  }
  fmt.Println(result)
}

// How Slice Expressions work

func sliceExpression() {
  maxValue := 10
  result := make([]int, 0, maxValue)
  for i := 0; i < maxValue; i++ {
    if i % 2 == 0 {
      result = append(result, i)
      // since result already has maxValue capacity, no new backend array will
      // be created and result will point to the same backend address
    }
  }
  for i := range result {
    fmt.Printf("%d: %v\n", i, &result[i])
  }
  newSlice := result[1:3]
  newSlice2 := result[2:4]
  fmt.Printf("[:]: %s\n", getSliceHeader(&result))
  fmt.Printf("[1:3]: %s\n", getSliceHeader(&newSlice))
  fmt.Printf("[2:4]: %s\n", getSliceHeader(&newSlice2))
}

func copyWork() {
  a := make([]int, 5, 6)  // both a & b are slice
  b := []int{1, 2, 3, 4, 5}
  fmt.Println(copy(a, b))
  fmt.Printf("a: %v, cap: %d\n", a, cap(a))
}


// Modifying a Sliced-slice Modifies the Original Slice

func sliceChange() {
  a := []int{1, 2, 3, 4, 5}
  b := a[2:4]
  b[0] = 10
  fmt.Println(b)
  fmt.Println(a)
}

func sliceModificationDriver() {
  belgium2020Race := NewRace("belgian", []string{
    "Hamilton",
    "Bottas",
    "Verstappen",
    "Ricciardo",
    "Ocon",
		"Albon",
    "Norris",
    "Gasly",
    "Stroll",
    "Perez",
		"Kvyat",
    "Räikkönen",
    "Vettel",
    "Leclerc",
    "Grosjean",
		"Latifi",
    "Magnussen",
    "Giovinazzi",
    "Russell",
    "Sainz",
  })
  top10Finishers := belgium2020Race.Top10FinishersCopied()
  fmt.Printf("%s GP top 10 finishers: %v\n", belgium2020Race.Name(), belgium2020Race.Result())
  sort.Strings(top10Finishers)  // sort modifying the slice returned from Top10Finishers()
  fmt.Printf("%s GP top 10 finishers, in alphabetical order: %v\n", belgium2020Race.Name(), top10Finishers)
  fmt.Printf("%s GP result: %v\n", belgium2020Race.Name(), belgium2020Race.Result())
}

type race struct {
  name string
  result []string
}

func (r race) Name() string {
  return r.name
}

func (r race) Result() []string {
  return r.result
}

func (r race) Top10FinishersDeprecated() []string {
  return r.result[:10]
  // original result is getting modified by sort.Strings (as the slice getting
  // passed have same underlying array)
}

func (r race) Top10FinishersCopied() []string {
  top10 := make([]string, 10)
  copy(top10, r.result[:10])
  return top10
  // original r.result is not getting modified as the slice returned as new
  // underlying array
}

func NewRace(name string, result []string) race {
  return race {
    name: name,
    result: result,
  }
}

// make a read-only version ( not including already written things below )
func (r race) Top10Finishers() ReadOnlyStringCollection {
  return readOnlyStringCollection{r.result[:10]}
}

type readOnlyStringCollection struct {
  value []string
}

func (r readOnlyStringCollection) Each(f func(i int, value string)) {
  for i, v := range r.value {
    f(i, v)
  }
}

func (r readOnlyStringCollection) Len() int {
  return len(r.value)
}

type ReadOnlyStringCollection interface {
  Each(f func(i int, value string))
  Len() int
}

func sliceModificationDriverReadOnly() {
  belgium2020Race := NewRace("belgian", []string{
    "Hamilton",
    "Bottas",
    "Verstappen",
    "Ricciardo",
    "Ocon",
		"Albon",
    "Norris",
    "Gasly",
    "Stroll",
    "Perez",
		"Kvyat",
    "Räikkönen",
    "Vettel",
    "Leclerc",
    "Grosjean",
		"Latifi",
    "Magnussen",
    "Giovinazzi",
    "Russell",
    "Sainz",
  })
  top10Finishers := func() []string {
    result := make([]string , 10)
    top10 := belgium2020Race.Top10Finishers()
    top10.Each(func(i int, val string) {
      result[i] = val
    })
    return result
  }()
  fmt.Printf("%s GP top 10 finishers: %v\n", belgium2020Race.Name(), top10Finishers)
  sort.Strings(top10Finishers)
  fmt.Printf("%s GP top 10 finishers, in alphabetical order: %v\n", belgium2020Race.Name(), top10Finishers)
  fmt.Printf("%s GP result: %v\n", belgium2020Race.Name(), belgium2020Race.Result())
}


// calling append on a sliced slice may modify the original slice

func slicedSliceAppend() {
  // a := []int{1, 2, 3, 4, 5}
  a := make([]int, 5, 6)
  copy(a, []int{1, 2, 3, 4, 5})
  b := a[2:4]
  fmt.Printf("a: %v, cap: %d\n", a, cap(a))
  fmt.Printf("b: %v, cap: %d\n", b, cap(b))

  // b = append(b, 20)
  a = append(a, 20)  // Experiment
  // this will also modify a (because b currently has capacity of 3 but only 2
  // are filled and this will not cause to point to new array)
  fmt.Printf("b: %v, cap: %d\n", b, cap(b))
  fmt.Printf("a: %v, cap: %d\n", a, cap(a))
}

func slicedSliceAppendNoChange() {
  a := []int{1, 2, 3, 4, 5}
  b := a[2:]  // capacity 3
  fmt.Printf("a: %v, cap: %d\n", a, cap(a))
  fmt.Printf("b: %v, cap: %d\n", b, cap(b))

  b = append(b, 20)  // will point to new underlying array now
  fmt.Printf("b: %v, cap: %d\n", b, cap(b))
  fmt.Printf("a: %v, cap: %d\n", a, cap(a))
}

func sliceAppendChangeInSlicedSlice() {
  a := make([]int, 5, 6)
  copy(a, []int{1, 2, 3, 4, 5})
  fmt.Printf("a: %v, cap: %d\n", a, cap(a))

  b := a[3:]  // cap 3
  fmt.Printf("b: %v, cap: %d\n", b, cap(b))
  b = append(b, 10)
  fmt.Printf("a: %v, cap: %d\n", a, cap(a))
  fmt.Printf("b: %v, cap: %d\n", b, cap(b))

  a = append(a, 20)
  fmt.Printf("a: %v, cap: %d\n", a, cap(a))
  fmt.Printf("b: %v, cap: %d\n", b, cap(b))
}

/* NOTE: if there's capacity left and you're appending to slice, then if the
* sliced slice have capacity left but len filled (and data upto the less than
* len index hasn't modified), it'll not change sliced slice. Same stands true
* for vice-versa. Or see it this way, reference does change (whether it's
* original or not) but struct.len is not updated */

func main() {
  slicedSliceAppend()
}
