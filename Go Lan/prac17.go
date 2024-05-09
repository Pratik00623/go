// Build a program that manages concurrent tasks using Goroutines and channels. Tasks can include 
// fetching data from multiple URLs or processing a set of concurrent calculations. Utilize Goroutines 
// for parallelism and channels for communication. 



package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchData(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching data from %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// Process the response (you can customize this part)
	// For demonstration purposes, we'll just print the status code.
	ch <- fmt.Sprintf("Fetched data from %s. Status code: %d", url, resp.StatusCode)
}

func main() {
	urls := []string{
		"https://youtube.com",
		"https://google.com",
		"https://github.com",
		// Add more URLs as needed
	}

	// Create a channel to communicate results
	resultCh := make(chan string)

	// Launch Goroutines for each URL
	for _, url := range urls {
		go fetchData(url, resultCh)
	}

	// Wait for all Goroutines to finish
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-resultCh)
	}

	// Optional: Add a delay to allow Goroutines to complete
	time.Sleep(1 * time.Second)
}
