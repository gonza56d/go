package timeouts

import (
	"fmt"
	"time"
)

func RunTimeoutExample() {
	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Result ready!"
	}()

	select {
	case res := <-ch:
		fmt.Println("Received:", res)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! Took too long.")
	}
}
