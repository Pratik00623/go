package main

import "fmt"

func main() {
	var rad float64
	const PI float64 = 3.14 // Constant
	var area float64

	fmt.Print("Enter radius of Circle: ")
	fmt.Scanln(&rad)

	area = PI * rad * rad
	fmt.Println("Area of Circle is:", area)
}
