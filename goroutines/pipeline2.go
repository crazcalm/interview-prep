package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//Squarer
	go func() {
		for {
			for x := range naturals {
				squares <- x * x
			}
			// close(squares) // errors out. Not share how to properly exit
		}
	}()

	//Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}
