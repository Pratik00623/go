package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	// Load an HTML template from a file
	htmlTemplate, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal("Error parsing HTML template:", err)
	}

	// Load a text template from a file
	textTemplate, err := template.ParseFiles("text.txt")
	if err != nil {
		log.Fatal("Error parsing text template:", err)
	}

	// Load a JSON template from a file
	jsonTemplate, err := template.ParseFiles("data.json")
	if err != nil {
		log.Fatal("Error parsing JSON template:", err)
	}

	// Data for the templates (you can customize this)
	data := struct {
		Name  string
		Value int
	}{
		Name:  "Mihir",
		Value: 21,
	}

	// Execute and display the HTML template
	fmt.Println("HTML Template:")
	err = htmlTemplate.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal("Error executing HTML template:", err)
	}

	// Execute and display the text template
	fmt.Println("\nText Template:")
	err = textTemplate.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal("Error executing text template:", err)
	}

	// Execute and display the JSON template
	fmt.Println("\nJSON Template:")
	err = jsonTemplate.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal("Error executing JSON template:", err)
	}
}
