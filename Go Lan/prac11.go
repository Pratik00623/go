// Build a geometry calculator that computes the area and perimeter of various shapes. Define an 
// interface Shape with methods like Area() and Perimeter(). Implement this interface for shapes such 
// as rectangles, circles, and triangles, showcasing interface contracts. 

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	fmt.Println("Area of Rectangle:")
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	fmt.Println("Perimeter of Rectangle:")
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

const pi = 3.14159

func (c Circle) Area() float64 {
	fmt.Println("Area of Circle:")

	return pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	fmt.Println("Perimeter of Circle:")

	return 2 * pi * c.radius
}

type Triangle struct {
	base, height float64
}

func (t Triangle) Area() float64 {
	fmt.Println("Area of Triangle:")

	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	fmt.Println("Perimeter of Triangle:")

	return t.base + 2*math.Sqrt(t.height*t.height+t.base*t.base/4)
}
func printInfo(s Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}

func main() {
	r := Rectangle{width: 4, height: 5}
	printInfo(r)

	c := Circle{radius: 5}
	printInfo(c)

	t := Triangle{base: 3, height: 4}
	printInfo(t)
}
