// Build a concurrent program where multiple Goroutines can read simultaneously and only one 
// Goroutine can write at a time. Model a shared resource, such as a database or configuration, and 
// use sync.RWMutex to implement read and write access.



package main

import (
	"fmt"
	"sync"
	"time"
)

// Database represents our shared resource (in-memory database).
type Database struct {
	data map[string]string
	mu   sync.RWMutex
}

// Initialize initializes the database.
func (db *Database) Initialize() {
	db.data = make(map[string]string)
}

// Read retrieves the value associated with a key.
func (db *Database) Read(key string) string {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.data[key]
}

// Write sets the value associated with a key.
func (db *Database) Write(key, value string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[key] = value
}

func main() {
	db := &Database{}
	db.Initialize()
~
	// Simulate concurrent read and write operations.
	for i := 0; i < 5; i++ {
		go func(i int) {
			// Read operation
			fmt.Printf("Read Goroutine %d: Value for key 'MrX': %s\n", i, db.Read("MrX"))

			// Write operation
			db.Write("MrY", fmt.Sprintf("New value from Goroutine %d", i))
			fmt.Printf("Write Goroutine %d: Updated value for key 'MrY'\n", i)
		}(i)
	}

	// Wait for goroutines to finish
	time.Sleep(2 * time.Second)

	// Final state of the database
	fmt.Println("\nFinal database state:")
	for key, value := range db.data {
		fmt.Printf("%s: %s\n", key, value)
	}
}
