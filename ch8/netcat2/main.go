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
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dest io.Writer, src io.Reader) {
	// fmt.Printf("%T, %[1]v\n", src)  // just for info
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
	fmt.Println("End:", src)
}


// From godoc net Conn
/* Conn is a generic stream-oriented network connection.
Multiple goroutines may invoke methods on a Conn simultaneously.   */

/* In both threads io.Copy() haven't finished execution and are still working,
* how ? I don't understand right now */
