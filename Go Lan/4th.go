package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World!"
	fmt.Println("Original String: ", str)

	// Convert the string to uppercase
	fmt.Println("ToUpper:", strings.ToUpper(str))

	// Convert the string to lowercase
	fmt.Println("ToLower:", strings.ToLower(str))

	// Check if the string contains the substring 'World'
	fmt.Println("Contains 'World':", strings.Contains(str, "World"))

	// Replace all occurrences of the substring 'World' with 'L'
	fmt.Println("ReplaceAll with 'L':", strings.Replace(str, "World", "L", -1))

	// Split the string by spaces
	fmt.Println("Split by space:", strings.Split(str, " "))

	// Check if the string starts with 'Hello '
	fmt.Println("Hasprefix 'Hello ':", strings.HasPrefix(str, "Hello "))

	// Check if the string ends with 'World'
	fmt.Println("Hassuffix 'World':", strings.HasSuffix(str, "World"))

	// Find the index of the first occurrence of 'O' in the string
	fmt.Println("Index of 'O':", strings.Index(str, "o"))

	// Repeat the string 3 times
	fmt.Println("Repeat 3 lines:", strings.Repeat(str, 3))
}
