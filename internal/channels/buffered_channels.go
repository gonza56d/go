package channels

import (
	"fmt"
	"time"
)

func buffered(ch chan string) {
	for msg := range ch {
		fmt.Println("Processing:", msg)
		time.Sleep(500 * time.Millisecond)
	}
}

func RunBufferedChannelsGoroutines() {
	ch := make(chan string, 2)
	go buffered(ch)

	ch <- "Task 1"
	ch <- "Task 2"
	ch <- "Task 3" // this will block until worker processes one

	close(ch)
	time.Sleep(2 * time.Second)
}
