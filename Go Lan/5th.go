package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	City  string
	Class string
}

func AddStudent(Students *[]Student) {
	var Student Student

	fmt.Print("Enter Student ID :")
	fmt.Scan(&Student.ID)

	fmt.Print("Enter Student Name :")
	fmt.Scan(&Student.Name)

	fmt.Print("Enter Student City :")
	fmt.Scan(&Student.City)

	fmt.Print("Enter Student Class :")
	fmt.Scan(&Student.Class)

	*Students = append(*Students, Student)

	fmt.Println("New student added!")

}

func DeleteStudent(Students *[]Student) {
	var StudID int

	fmt.Print("Enter the ID of the student you want to delete: ")
	fmt.Scan(&StudID)

	for i, Student := range *Students {
		if Student.ID == StudID {
			*Students = append((*Students)[:i], (*Students)[i+1:]...)
			fmt.Println("The selected student has been deleted.")
			return
		}
	}

	fmt.Println("No such student found with given ID.")
}

func SearchStudent(Students *[]Student) {
	var StudentID int
	fmt.Print("Enter the ID of the student you are looking for: ")
	fmt.Scan(&StudentID)

	for _, Student := range *Students {
		if Student.ID == StudentID {

			fmt.Printf("ID:%s\nName:%d\nCity:%v\nClass:%s", Student.ID, Student.Name, Student.City, Student.Class)
			return
		}
	}

	fmt.Println("No record found with provided ID number.\nPlease check and try again.")
}

func main() {

	var Student []Student
	for {

		fmt.Println("\n1.Add a new student \n2.Delete a student \n3.Search a student \n4.Exit")

		fmt.Print("Enter your choice : ")
		var Choice int
		fmt.Scan(&Choice)
		switch Choice {
		case 1:
			AddStudent(&Student)
		case 2:
			DeleteStudent(&Student)
		case 3:
			SearchStudent(&Student)
		case 4:
			fmt.Println("Exiting Program")
			return
		default:
			fmt.Println("Invalid Option! Please enter a valid option from 1 to 4.")
		}
	}
}
