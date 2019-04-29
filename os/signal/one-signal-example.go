package main

import (
	"fmt"
	"os"
	"os/signal"
)

func InterruptHandler(c <-chan os.Signal, done chan<- struct{}) {
	for i := 0; i < 3; i++ {
		select {
		case <-c:
			fmt.Println("Got Signal")
		}
	}
	done <- struct{}{}
}

func main() {
	c := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(c, os.Interrupt)
	go InterruptHandler(c, done)
	<-done
}
