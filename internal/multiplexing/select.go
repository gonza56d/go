package multiplexing

import (
	"fmt"
	"time"
)

func worker(id string, ch chan string, delay time.Duration) {
	time.Sleep(delay)
	ch <- fmt.Sprintf("Worker %s finished after %v", id, delay)
}

func RunSelectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go worker("A", ch1, 2*time.Second)
	go worker("B", ch2, 1*time.Second)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}

	fmt.Println("All workers done")
}
