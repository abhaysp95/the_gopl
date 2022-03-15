package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func walkDir(dir string, fileSize chan <- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSize)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du2: %v\n", err)
				return
			}
			fileSize <- info.Size()
		}
	}
}

func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du2: %v\n", err)
		return nil
	}
	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSize)
		}
		close(fileSize)
	}()

	// Print the results periodically
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64

loop:  // label
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
				// labeled break statement (breaks out of select and loop both)
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
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
