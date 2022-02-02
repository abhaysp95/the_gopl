package main

import (
	"fmt"
	"image/color"
	"math"
	"time"
)

type Point struct { X, Y float64 }

type ColoredPoint struct {
  Point
  Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
  return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(scale float64) {
  p.X *= scale
  p.Y *= scale
}

type Rocket struct { /*....*/ }

func (r *Rocket) Launch() { /*....*/ }

func rocketLaunch() {
  r := new(Rocket)
  // long syntax
  // time.AfterFunc(10 * time.Second, func() { r.Launch() })
  time.AfterFunc(10 * time.Second, r.Launch)  // r.Launch being the method value
}

func main() {
  p := Point{1, 2}
  q := Point{4, 6}

  distanceFromP := p.Distance  // method value
  fmt.Println(distanceFromP(q))
  var origin Point
  fmt.Println(distanceFromP(origin))

  scaleP := p.ScaleBy
  for i := 2; i < 5; i++ {
    scaleP(float64(i))
    fmt.Println(p)
  }
}
