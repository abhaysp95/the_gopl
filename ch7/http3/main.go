package main

import (
  "fmt"
  "net/http"
)

type dollar float32

func (d dollar) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollar
