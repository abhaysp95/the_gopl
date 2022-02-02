package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct { X, Y float64 }

type ColoredPoint struct {
  Point
  Color color.RGBA
}

// used reciever argument of pointer type just for the sake of example
func (p *Point) Distance(q *Point) float64 {
  return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(scale float64) {
  p.X *= scale
  p.Y *= scale
}

func (p Point) Add(q Point) Point {
  return Point {
    p.X + q.X,
    p.Y + q.Y,
  }
}

func (p Point) Sub(q Point) Point {
  return Point {
    p.X - q.X,
    p.Y - q.Y,
  }
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
  var op func(p, q Point) Point
  switch add {
  case true:
    op = Point.Add
  case false:
    op = Point.Sub
  // not using default case here
  }

  for i := range path {
    path[i] = op(path[i], offset)
  }
}

func main() {
  p := Point{1, 2}
  q := Point{4, 6}

  distance := (*Point).Distance
  fmt.Println(distance(&p, &q))
  fmt.Printf("Distance: %T\n", distance)

  scale := (*Point).ScaleBy
  scale(&p, 2)
  fmt.Println(p)
  fmt.Printf("ScaleBy: %T\n", scale)

  path := Path{
    Point{1, 2},
    Point{3, 4},
    Point{5, 6},
  }

  fmt.Println(path)

  path.TranslateBy(Point{1, 1}, true)
  // change is happening in original path because internal slice is passed as
  // reference

  fmt.Println(path)
}
