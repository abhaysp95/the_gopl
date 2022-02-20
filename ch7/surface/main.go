package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"

	"the_gopl/ch7/eval"
)

const (
  width, height = 1280, 720            // canvas size in pixels
  cells         = 200                 // number of grid cells
  xyrange       = 30.0                // x, y axis range(-xyrange..+xyrange)
  xyscale       = width / 2 / xyrange // pixels per x or y unit
  zscale        = height * 0.4        // pixels per z unit
)

var sin30, cos30 = 0.5, math.Sqrt(3.0 / 4.0)

func corner(f func(x, y float64) float64, i, j int) (float64, float64) {
  x := xyrange * (float64(i) / cells - 0.5)
  y := xyrange * (float64(j) / cells - 0.5)

  z := f(x, y)

  // project (x,y,z) isometrically onto 2-D SVG canvas(sx,sy)
  sx := width/2 + (x-y)*cos30*xyscale
  sy := height/2 + (x+y)*sin30*xyscale-z*zscale
  return sx, sy
}

func surface(w io.Writer, f func(x, y float64) float64) {
  fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' " +
    "style='stroke: grey; fill: white; stroke-width: 0.7' " +
    "width='%d' height='%d'>", width, height)
  for i := 0; i < cells; i++ {
    for j := 0; j < cells; j++ {
      ax, ay := corner(f, i+1, j)
      bx, by := corner(f, i, j)
      cx, cy := corner(f, i, j+1)
      dx, dy := corner(f, i+1, j+1)
      fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
      ax, ay, bx, by, cx, cy, dx, dy)
    }
  }
  fmt.Fprintf(w, "</svg>")
}

//!+parseAndCheck
func parseAndCheck(s string) (eval.Expr, error) {
  if s == "" {
    return nil, fmt.Errorf("empty expression")
  }
  expr, err := eval.Parse(s)
  if err != nil {
    return nil, err
  }
  vars := make(map[eval.Var]bool)
  if err := expr.Check(vars); err != nil {
    return nil, err
  }
  for v := range vars {
    if v != "x" && v != "y" && v != "r" {
      return nil, fmt.Errorf("undefined variable: %s", v)
    }
  }
  return expr, nil
}
//!-parseAndCheck

//!+plot
func plot(w http.ResponseWriter, req *http.Request) {
  req.ParseForm()
  expr, err := parseAndCheck(req.Form.Get("expr"))
  if err != nil {
    http.Error(w, "bad expr: " + err.Error(), http.StatusBadRequest)
    return
  }
  w.Header().Set("Content-Type", "image/svg+xml")
  surface(w, func(x, y float64) float64 {
    r := math.Hypot(x, y)  // distance from (0,0)
    return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
  })
}
//!-plot

//!+main
func main() {
  http.HandleFunc("/plot", plot)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
//!-main
