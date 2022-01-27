package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
  Title string
  Year int `json:"released"`
  Color bool `json:"color,omitempty"`
  Actors []string
}
// only exported fields are marshaled (if you make fiels
// unexport(small), then will not be marshalled)

func main() {
  var movies = []Movie{
    {
      Title: "Casablanca",
      Year: 1942,
      // Color: false,  // just to check how omitempty works
      Actors: []string{ "Humphrey Bogart", "Ingrid Bergman" },
    },
    {
      Title: "Cool Hand Luke",
      Year: 1967,
      Color: true,
      Actors: []string{ "Paul Newman" },
    },
    {
      Title: "Bullitt",
      Year: 1968,
      Color: true,
      Actors: []string{ "Steve McQueen", "Jacqueline Bisset" },
    },
  }

  // data, err := json.Marshal(movies)
  // MarshalIndent is just like Marshal but applies indent to format
  // the output
  data, err := json.MarshalIndent(movies, "", "\t")
  if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
  }
  fmt.Printf("%s\n", data)

  var titles []struct{ Title string }
  if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSON unmarshaling failed: %s\n", err)
  }
  fmt.Println(titles)
}

// Note: Putting "=> " as arg for prefix in MarshalIndent caused problem in
// unmarshaling that JSON object/s
