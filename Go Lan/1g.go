// Write a Go program that reads user input, performs a simple calculation (e.g., addition or 
// 	multiplication), and outputs the result. Focus on organizing the code into functions and using proper 
// 	program structure. 


package main

import (
	"fmt"
)

// function to perform addition
func add(a, b int) int {
	return a + b
}

// function to perform multiplication
func multiply(a, b int) int {
	return a * b
}

func main() {
	// ask user to select operation
	fmt.Println("Select operation:")
	fmt.Println("1) Add")
	fmt.Println("2) Multiply")

	var operation int
	fmt.Scan(&operation)

	// ask user to enter the first integer
	var num1 int
	fmt.Println("Enter first integer:")
	fmt.Scan(&num1)

	// ask user to enter the second integer
	var num2 int
	fmt.Println("Enter second integer:")
	fmt.Scan(&num2)

	var result int

	switch operation {
	case 1:
		// perform addition
		result = add(num1, num2)
	case 2:
		// perform multiplication
		result = multiply(num1, num2)
	default:
		fmt.Println("Invalid operation")
		return
	}

	// output the result
	fmt.Println("Result:", result)
}
