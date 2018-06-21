package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Blast off!")
}

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) //read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			//Do nothing
		case <-abort:
			fmt.Println("launch aborted!")
			return
		}
	}
	launch()
}
