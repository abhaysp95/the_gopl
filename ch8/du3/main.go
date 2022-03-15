package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done()
	for _, entry := range dirent(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSize)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du3 [walkDir]: %v\n", err)
				return
			}
			fileSize <- info.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents
var sema = make(chan struct{}, 20)

func dirent(dir string) []os.DirEntry {
	sema <- struct{}{}  // acquire a token
	defer func() {
		<-sema
	}()

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3 [dirent]: %v\n", err)
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

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var n sync.WaitGroup

	fileSize := make(chan int64)
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSize)
	}

	// closer
	go func() {
		n.Wait()
		close(fileSize)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSize:
			if !ok {
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
