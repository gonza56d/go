package channels

import (
	"fmt"
)

func producer(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // must close so the receiver knows there are no more values
}

func main() {
	ch := make(chan int)

	go producer(ch)

	// Range will automatically stop when channel is closed
	for val := range ch {
		fmt.Println("Received:", val)
	}

	fmt.Println("All values received, channel closed")
}
