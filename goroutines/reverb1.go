/*
Notice that the third shout from the client is not dealt with until the second shout has petered out,
which is not very realistic. A real echo would consist of the composition of the three independent shouts.
To simulate it, we'll need more goroutines. Again, all we need to do is add the go keyword, this time to
the call to echo
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprint(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprint(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprint(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	//Note: ignoring potential errors from input.Err()
	c.Close()
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
		go handleConn(conn)
	}
}
