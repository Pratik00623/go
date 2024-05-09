// Create a program that performs a lengthy calculation and can be canceled on user demand. Use the 
// context package to handle cancellation gracefully, demonstrating the ability to stop ongoing 
// operations.


package main

import (
	"fmt"
	"sync"
)

var counter int

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go incrementCounter(&wg)
	go incrementCounter(&wg)

	wg.Wait()

	fmt.Printf("Final counter value (without synchronization): %d\n", counter)
}
