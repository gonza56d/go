package goroutines

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func RunGoroutines() {
	// Start a goroutine
	go sayHello("Alice")
	// Runs in main goroutine
	sayHello("Main")
	// Give goroutines time to finish
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Done")
}
