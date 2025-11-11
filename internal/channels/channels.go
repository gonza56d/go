package channels

import (
	"fmt"
	"time"
)

func worker(ch chan string) {
	fmt.Println("Worker started")
	time.Sleep(200 * time.Millisecond)
	ch <- "Work done!" // send a message
}

func Channels() {
	ch := make(chan string) // create a channel

	go worker(ch)

	msg := <-ch // receive message (blocks until available)
	fmt.Println("Received:", msg)
}
