package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// NOTE: I tried printing on my own (not included in book for this example)

// walkDir recursively walks the file tree rooted at dir and sends the size of
// each found file on filesizes
func walkDir(dir string, fileSizes chan<- int64, level int) {
	for _, entry := range dirents(dir) {
		fmt.Printf("%s. %s\n", strings.Repeat(" ", level), entry.Name())
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes, level + 1)
		} else {
			info, err := entry.Info()
			if err != nil {
				// log.Fatalf("du1: %v\n", err)
				fmt.Fprintf(os.Stderr, "du1: %v\n", err)
			}
			fileSizes <- info.Size()
		}
	}
}

// dirents returns the entries of directory dir
func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	// Determine the initial directories
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// traverse the file tree (I suppose, I can use buffered channel here too)
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes, 0)
		}
		close(fileSizes)
	}()

	// print the result
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	var size float64 = float64(nbytes)/1e6
	unit := "MB"
	if size >= 1e3 {
		size = float64(nbytes)/1e9
		unit = "GB"
	}
	fmt.Printf("%d files %.3f %s\n", nfiles, size, unit)
}
