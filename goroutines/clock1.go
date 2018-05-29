/*
Clock1 is a TCP server that periodically writes the time

The below code can only handle one client. The second client must wait until the first client is finished
because the server is sequential; it deals with one client at a time. Just one small change is needed to
make the server concurrent: adding the go keyword to the call to handleConn causes each call to run in
its own goroutine.
*/
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return //e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) //handle one connection at a time
	}
}
