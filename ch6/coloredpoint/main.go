package main

import (
  "fmt"
  "image/color"
  "math"
)

type Point struct {
  X, Y float64
}

type ColoredPoint struct {
  Point
  Color color.RGBA
}

func main() {
  red := color.RGBA{255, 0, 0, 255}
  blue := color.RGBA{0, 0, 255, 255}
  p := ColoredPoint{Point{1, 1}, red}
  q := ColoredPoint{Point{5, 4}, blue}

  fmt.Println(p.Distance(q.Point))
  p.ScaleBy(2)
  q.ScaleBy(2)
  fmt.Println(p.Distance(q.Point))
}

func (p Point) Distance(q Point) float64 {
  return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(scale float64) {
  p.X *= scale
  p.Y *= scale
}
