// Implement a program that calculates the average of a set of exam scores using an array. Allow the 
// user to input the scores and use an array to store and process the data

package main

import (
	"fmt"
)

func main() {
	// Ask user to enter the number of exam scores
	var numScores int
	fmt.Println("Enter the number of exam scores:")
	fmt.Scan(&numScores)

	// Create an array to store the exam scores
	var scores [10]int

	// Ask user to enter the exam scores
	fmt.Println("Enter the exam scores:")
	for i := 0; i < numScores; i++ {
		fmt.Scan(&scores[i])
	}

	// Calculate the average of the exam scores
	var sum int
	for i := 0; i < numScores; i++ {
		sum += scores[i]
	}
	average := float64(sum) / float64(numScores)

	// Output the average
	fmt.Println("The average is:", average)
}