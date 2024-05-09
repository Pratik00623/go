package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 3, 8, 4, 2}
	fmt.Println("Original Array :", arr)

	var order string
	fmt.Print("Enter 'asc' for ascending order or 'desc' for descending order :")
	fmt.Scan(&order)

	switch order {
	case "asc":
		sort.Ints(arr)
		fmt.Println("Sorted Array in Ascending Order :", arr)
	case "desc":
		sort.Sort(sort.Reverse(sort.IntSlice(arr)))
		fmt.Println("Sorted Array in Descending Order: ", arr)
	default:
		fmt.Println("Invalid Order Specified ! Please Enter 'asc' or 'desc'. ")
	}
}
