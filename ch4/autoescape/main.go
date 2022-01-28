package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
  const templ = `<p>A: {{.A}}</p><br><p>B: {{.B}}</p>`

  var data struct {
    A string
    B template.HTML
  }
  data.A = "<b>Hello!</b>"
  data.B = "<b>Hello!</b>"

  t := template.Must(template.New("escape").Parse(templ))
  if err := t.Execute(os.Stdout, data); err != nil {
    log.Fatal(err)
  }
}
