package main

import (
	"semaphorePattern/semaphore"
	"time"
)

func main() {
	tickets, timeout := 1, 3*time.Second
	s := semaphore.New(tickets, timeout)

	if err := s.Acquire(); err != nil {
		panic(err)
	}

	// Do important work

	if err := s.Release(); err != nil {
		panic(err)
	}
}
