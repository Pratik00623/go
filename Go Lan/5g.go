// Build a simple dictionary program where users can look up word definitions. Utilize a map to store 
// word-definition pairs and implement operations like adding new words and looking up definitions. 


package main

import (
	"fmt"
)

func main() {
	// Initialize a map to store word-definition pairs
	wordDefs := make(map[string]string)

	// Add some word-definition pairs to the map
	wordDefs["algorithm"] = "a set of rules for solving a problem"
	wordDefs["data structure"] = "a way of organizing and storing data"

	// Look up a definition
	definition, ok := wordDefs["algorithm"]
	if ok {
		fmt.Println("algorithm:", definition)
	} else {
		fmt.Println("Definition not found")
	}

	// Add a new word-definition pair
	wordDefs["function"] = "a named piece of code"
	definition, ok = wordDefs["function"]
	if ok {
		fmt.Println("function:", definition)
	} else {
		fmt.Println("Definition not found")
	}
}
