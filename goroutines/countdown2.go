/*
Let's add the ability to abort the launch sequence by pressing the return key during the countdown.

Note: Regardless of when you abort the program, the program will still complete the countdown
*/

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

	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}

	select {
	case <-time.After(10 * time.Second):
		//Do nothing
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()
}
