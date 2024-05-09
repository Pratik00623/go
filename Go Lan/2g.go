// Create a program that converts temperatures between Fahrenheit and Celsius. Use variables to 
// store temperature values and ensure proper data type conversions during calculations. 



package main

import (
	"fmt"
)

// Constant for the number of degrees to convert between Fahrenheit and Celsius
const conversionFactor = 1.8
const offset = 32

// Function to convert Fahrenheit to Celsius
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - offset) / conversionFactor
}

// Function to convert Celsius to Fahrenheit
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * conversionFactor) + offset
}

func main() {
	// Ask user to select operation
	fmt.Println("Select operation:")
	fmt.Println("1) Convert Fahrenheit to Celsius")
	fmt.Println("2) Convert Celsius to Fahrenheit")

	var operation int
	fmt.Scan(&operation)

	// Ask user to enter the temperature
	var temp float64
	fmt.Println("Enter temperature:")
	fmt.Scan(&temp)

	var result float64

	switch operation {
	case 1:
		// Convert Fahrenheit to Celsius
		result = fahrenheitToCelsius(temp)
	case 2:
		// Convert Celsius to Fahrenheit
		result = celsiusToFahrenheit(temp)
	default:
		fmt.Println("Invalid operation")
		return
	}

	// Output the result
	fmt.Println("Result:", result)
}