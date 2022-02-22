package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)  // NOTE:ignoring errors
		fmt.Println("done")
		done <- struct{}{}  // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done  // wait for background goroutine to finish
}

func mustCopy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}

// Communication over unbuffered channel causes the sending and recieving
// goroutines to synchronize. Because of this, unbuffered channels are
// sometimes called 'Synchronous Channels'.
