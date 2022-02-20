package main

import (
	"database/sql"
	"fmt"
)

func listTracks(db sql.DB, artist string, minYear, maxYear int) {
  result, err := db.Exec(
    "SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
    artist, minYear, maxYear)
    // ...
}

// within Exec, we might find a function like the one below, which converts each argument value to its literal SQL notation
func sqlQuote(x interface{}) string {
  if x == nil {
    return "NULL"
  } else if _, ok := x.(int); ok {
    return fmt.Sprintf("%d", x)
  } else if _, ok := x.(uint); ok {
    return fmt.Sprintf("%d", x)
  } else if b, ok := x.(bool); ok {
    if b {
      return "TRUE"
    }
    return "FALSE"
  } else if s, ok := x.(string); ok {
    return sqlQuoteString(s)  // not shown
  } else {
    panic(fmt.Sprintf("unexpected type %T: %v", x, x))
  }
}

// A switch statement simplifies an if-else chain that performs a series of value equality tests. An analogous type switch statement simplifies an if-else chain of type assertions

// Using extended switch form (rewrite)
func sqlQuote2(x interface{}) string {
  switch x := x.(type) {  // notice here
  case nil:
    return "NULL"
  case int, uint:  // x has new type here (result of Type Assertion)
    return fmt.Sprintf("%d", x)  // x has type interface{} here
  case bool:
    if x {
      return "TRUE"
    }
    return "FALSE"
  case string:
    return sqlQuoteString(x)  // not shown
  default:
    panic(fmt.Sprintf("unexpected type %T: %v", x, x))
  }
}
