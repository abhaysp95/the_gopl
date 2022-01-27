package main

import "fmt"

type Point struct {
  X, Y int
}

type Circle struct {
  Center Point
  Radius int
}

type Wheel struct {
  Circle Circle
  Spokes int
}

// anonymous fields
// the type of field must be a named type or a pointer to a named type

type Circle1 struct {
  Point  // name and visibility of the field is implicitly determined
  // int  // valid (because type and field name both are treated as int)
  Radius int
}

type Wheel1 struct {
  Circle1
  Spokes int
}

func main() {
  var w Wheel
  w.Circle.Center.X = 8
  w.Circle.Center.Y = 8
  w.Circle.Radius = 5
  w.Spokes = 20

  // because of embedding we can refer to the names at leaves of implicit tree without giving the intervening names
  var w1 Wheel1
  w1.X = 8        // equivalent to explicit form w1.Circle1.Point.X
  w1.Y = 8
  w1.Radius = 5
  w1.Spokes = 20

  fmt.Printf("%#v, %#v\n", w, w1)


  // but there's no corresponding shorthand for struct literal syntax
  /* w = Wheel{8, 8, 5, 20}
  w1 = Wheel1{8, 8, 5, 20} */

  w = Wheel{Circle{Point{8, 8}, 5}, 20}
  wAgain := Wheel{
    Circle: Circle{
      Center: Point{X: 8, Y: 8},
      Radius: 5,
    },
    Spokes: 20,
  }
  fmt.Printf("%#v, %#v\n", w, wAgain)

  w1 = Wheel1{Circle1{Point{8, 8}, 5}, 20}
  w1Again := Wheel1{
    Circle1: Circle1{
      Point: Point{X: 8, Y: 8},
      Radius: 5,
    },
    Spokes: 20,
  }
  fmt.Printf("%#v, %#v\n", w1, w1Again)
}
