package channels

import (
	"fmt"
	"time"
)

func channel_worker(output chan string, id string) {
	time.Sleep(100 * time.Millisecond)
	output <- "Worker " + id + " done!"
}

func ChannelsHomework() {
	result := make(chan string)
	go channel_worker(result, "A")
	go channel_worker(result, "B")
	go channel_worker(result, "C")
	for i := 0; i < 3; i++ {
		fmt.Println(<-result)
	}
}
