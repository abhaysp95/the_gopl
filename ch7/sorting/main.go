package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
  Title string
  Artist string
  Album string
  Year int
  Length time.Duration
}

var tracks = []*Track{
  {"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
  {"Go", "Moby", "Moby", 1992, length("3m37s")},
  {"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
  {"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
  d, err := time.ParseDuration(s)
  if err != nil {
    panic(s)
  }
  return d
}

// print in Tabular form
func printTracks(tracks []*Track) {
  const format = "%v\t%v\t%v\t%v\t%v\t\n"
  // you can update flag (last arg of Init), like AlignRight | Debug
  tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
  fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
  fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
  for _, t := range tracks {
    fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
  }
  tw.Flush()
}

// to sort by artist
type byArtist []*Track

func (p byArtist) Len() int { return len(p) }
func (p byArtist) Less(i, j int) bool { return p[i].Artist < p[j].Artist }
func (p byArtist) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type byYear []*Track

func (p byYear) Len() int { return len(p) }
func (p byYear) Less(i, j int) bool { return p[i].Year < p[j].Year }
func (p byYear) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type customSort struct {
  tracks []*Track
  less func(x, y *Track) bool
}

func (p customSort) Len() int { return len(p.tracks) }
func (p customSort) Less(i, j int) bool { return p.less(p.tracks[i], p.tracks[j]) }
func (p customSort) Swap(i, j int) { p.tracks[i], p.tracks[j] = p.tracks[j], p.tracks[i] }

func main() {
  fmt.Println("By Aritst:")
  sort.Sort(byArtist(tracks))
  printTracks(tracks)

  fmt.Println("By Artist(Reverse):")
  sort.Sort(sort.Reverse(byArtist(tracks)))
  printTracks(tracks)

  fmt.Println("By Year:")
  sort.Sort(byYear(tracks))
  printTracks(tracks)

  // primary sort key is Title, secondary key is Year and tertiary is the running time length
  sort.Sort(customSort{tracks, func(x, y *Track) bool {
    if x.Title != y.Title {
      return x.Title < y.Title
    }
    if x.Year != y.Year {
      return x.Year < y.Year
    }
    if x.Length != y.Length {
      return x.Length < y.Length
    }
    return false
  }})
  printTracks(tracks)
}

func init() {
  values := []int{3, 1, 4, 1}
  fmt.Println(sort.IntsAreSorted(values))
  sort.Ints(values)
  fmt.Println(sort.IntsAreSorted(values))
  sort.Sort(sort.Reverse(sort.IntSlice(values)))
  fmt.Println(values)
  fmt.Println(sort.IntsAreSorted(values))
}
