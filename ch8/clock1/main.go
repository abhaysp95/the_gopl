package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
  listener, err := net.Listen("tcp", "localhost:8000")
  if err != nil {
    log.Fatal(err)
  }
  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Print(err)  // e.g., connection aborted
      continue
    }
    handleConn(conn)  // handle one connection at a time
  }
}

func handleConn(c net.Conn) {
  defer c.Close()
  for {
    // if first argument satisfies StringWriter interface, its WriteString
    // method is called, else Write is called exactly one
    _, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
    if err != nil {
      return  // e.g., client disconnected
    }
    time.Sleep(1 * time.Second)
  }
}
