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
	if cancelled() {
		return
	}
	for _, entry := range dirent(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(subdir, n, fileSize)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du4 [walkDir]: %v\n", err)
				return
			}
			fileSize <- info.Size()
		}
	}
}

// counting semaphore
var sema = make(chan struct{}, 20)

func dirent(dir string) []os.DirEntry {
	select {
	case sema <- struct{}{}:  // acquire token
		// do nothing
	case <-done:
		return nil  // cancelled
	}
	defer func() {
		<-sema
	}()

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du4 [dirent]: %v\n", err)
		return nil
	}
	return entries
}

var done = make(chan struct{})

/** when done is not closed, this is false because case is expecting to recieve
* some value from the channel, but when done channel is closed, it becomes true
* as "after a channel is has been closed and drained of all sent values,
* subseqent recieve operation proceed immediately, yielding zero values" */
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

var verbose = flag.Bool("v", false, "show verbose progress message")

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

	// cancel when input is detected
	go func() {
		os.Stdin.Read(make([]byte, 1))  // read a single byte
		close(done)
	}()

	// closer
	go func() {
		n.Wait()
		close(fileSize)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			// drain fileSize to allow existing goroutines to finish
			for range fileSize {
				// do nothing
			}
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
