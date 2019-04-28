package main

import (
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signal.Ignore(syscall.SIGKILL, syscall.SIGINT, syscall.SIGSTOP)
	for {
		time.Sleep(1 * time.Second)
	}
}
