package goroutines

import (
	"fmt"
	"sync"
	"time"
)

// var wg sync.WaitGroup // not recommended to use it in this way to avoid undefined behavior because of mutability.
// it's always better to pass its pointer as an argument.

func countNumbers(n int, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		fmt.Println("Counting numbers:", i, name)
		time.Sleep(100 * time.Millisecond)
	}
}

func RunGoroutinesHomework() {
	var wg sync.WaitGroup
	wg.Add(3)
	go countNumbers(5, "A", &wg)
	go countNumbers(5, "B", &wg)
	go countNumbers(5, "C", &wg)
	wg.Wait()
}
