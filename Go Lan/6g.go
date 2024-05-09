// Design a program that models a library system. Use structs to represent books with attributes like 
// title, author, and publication year. Implement functions to check out and return books. 


package main

import (
	"fmt"
)

// Book represents a book with attributes like title, author, and publication year.
type Book struct {
	Title           string
	Author          string
	PublicationYear int
	IsCheckedOut    bool // Indicates whether the book is currently checked out
}

// Library manages a collection of books.
type Library struct {
	Books []Book
}

// AddBook adds a book to the library.
func (l *Library) AddBook(title, author string, publicationYear int) {
	newBook := Book{
		Title:           title,
		Author:          author,
		PublicationYear: publicationYear,
		IsCheckedOut:    false, // Initially not checked out
	}
	l.Books = append(l.Books, newBook)
}

// CheckOutBook checks out a book.
func (l *Library) CheckOutBook(title string) bool {
	for i, book := range l.Books {
		if book.Title == title && !book.IsCheckedOut {
			l.Books[i].IsCheckedOut = true
			return true // Successfully checked out
		}
	}
	return false // Book not found or already checked out
}

// ReturnBook returns a book.
func (l *Library) ReturnBook(title string) bool {
	for i, book := range l.Books {
		if book.Title == title && book.IsCheckedOut {
			l.Books[i].IsCheckedOut = false
			return true // Successfully returned
		}
	}
	return false // Book not found or not checked out
}

func main() {
	myLibrary := Library{}

	// Add some books to the library
	myLibrary.AddBook("The Great Gatsby", "F. Scott Fitzgerald", 1925)
	myLibrary.AddBook("To Kill a Mockingbird", "Harper Lee", 1960)
	myLibrary.AddBook("1984", "George Orwell", 1949)

	// Check out a book
	if myLibrary.CheckOutBook("The Great Gatsby") {
		fmt.Println("Successfully checked out 'The Great Gatsby.'")
	} else {
		fmt.Println("Book not available for checkout.")
	}

	// Return a book
	if myLibrary.ReturnBook("The Great Gatsby") {
		fmt.Println("Successfully returned 'The Great Gatsby.'")
	} else {
		fmt.Println("Book not found or not checked out.")
	}
}
