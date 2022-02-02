package main

import (
	"fmt"
	"math"
)

type Point struct {
  X, Y float64
}

// traditional function

func Distance(p, q Point) float64 {
  return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Distance(q Point) float64 {
  return math.Hypot(p.X-q.X, p.Y-q.Y)
}

type Path []Point

// Distance retuns the distance travelled along the path
func (p Path) Distance() float64 {
  sum := 0.0
  for i := range p {
    if i > 0 {
      sum += p[i-1].Distance(p[i])
    }
  }
  return sum
}

func main() {
  perim := Path{
    {1, 1},
    {5, 1},
    {5, 4},
    {1, 1},
  }
  fmt.Println(perim.Distance())
}
