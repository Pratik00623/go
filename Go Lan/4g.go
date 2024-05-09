// Develop a program that manages a to-do list. Allow users to add, remove, and display tasks using 
// slices. Demonstrate the flexibility of slices in handling dynamic data. 



package main

import (
	"fmt"
)

func main() {
	tasks := []string{}

	for {
		fmt.Println("\nPlease choose an option:")
		fmt.Println("1. Add a task")
		fmt.Println("2. Remove a task")
		fmt.Println("3. Display tasks")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var task string
			fmt.Print("Enter the task: ")
			fmt.Scanln(&task)
			tasks = append(tasks, task)
			fmt.Println("Task added successfully.")
		case 2:
			if len(tasks) == 0 {
				fmt.Println("No tasks to remove.")
				continue
			}
			var index int
			fmt.Print("Enter the index of the task to remove: ")
			fmt.Scanln(&index)
			if index < 0 || index >= len(tasks) {
				fmt.Println("Invalid index.")
				continue
			}
			tasks = append(tasks[:index], tasks[index+1:]...)
			fmt.Println("Task removed successfully.")
		case 3:
			if len(tasks) == 0 {
				fmt.Println("No tasks to display.")
				continue
			}
			fmt.Println("Tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		case 4:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}