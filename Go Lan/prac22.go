// Develop a program where lazy initialization is crucial, such as initializing a configuration or database 
// connection only when needed. Use sync.Once to ensure that the initialization code is executed only 
// once, even if accessed concurrently. 


package main

import (
	"fmt"
	"sync"
)

// Config represents a configuration struct
type Config struct {
	APIKey string
}

// globalConfig represents the singleton instance of Config
var globalConfig *Config
var once sync.Once

// GetConfig initializes and returns the global configuration
func GetConfig() *Config {
	once.Do(func() {
		// Simulate some heavy initialization process, like reading from a file or database
		fmt.Println("Initializing configuration...")
		globalConfig = &Config{
			APIKey: "secret-api-key",
		}
	})
	return globalConfig
}

func main() {
	// Accessing the configuration from multiple goroutines concurrently
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			config := GetConfig()
			fmt.Printf("API Key: %s\n", config.APIKey)
		}()
	}
	wg.Wait()
}
