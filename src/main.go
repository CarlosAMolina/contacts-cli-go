package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <contactsJsonPath>")
		fmt.Println("Example: go run main.go /tmp/contacts.json")
		os.Exit(1)
	}
	jsonPath := os.Args[1]
	var data Data = readFile(jsonPath)
	fmt.Println("Welcome to the contacts CLI!")
	showHelp()
	var choice string
	for {
		fmt.Print(">> ")
		fmt.Scan(&choice)
		switch choice {
		case "1":
			processTerm(data)
			fmt.Println()
			showHelp()
		case "2":
			processId(data)
			fmt.Println()
			showHelp()
		case "3":
			processAll(data)
			fmt.Println()
			showHelp()
		case "e":
			fmt.Println("Bye!")
			os.Exit(0)
		case "h":
			showHelp()
		default:
			fmt.Println("Invalid input")
		}
	}
}

func showHelp() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Search term")
	fmt.Println("2. Show contact by ID")
	fmt.Println("3. Show all contacts")
	fmt.Println("e. Exit")
	fmt.Println("h. Show help")
}

func processAll(data Data) {
	fmt.Printf("Number of contacts: %d\n", len(data.Contacts))
	showSummary(data.Contacts)
}

func processId(data Data) {
	fmt.Println("Write the ID to show:")
	fmt.Print(">> ")
	var id int
	_, err := fmt.Scan(&id)
	for err != nil {
		fmt.Println("Write a number")
		fmt.Print(">> ")
		_, err = fmt.Scan(&id)

	}
	showContact(id, data.Contacts)
}

func processTerm(data Data) {
	fmt.Println("Write the term to search:")
	fmt.Print(">> ")
	var searchTerm string
	fmt.Scan(&searchTerm)
	matchedContacts := searchContacts(searchTerm, data.Contacts)
	fmt.Printf("Found %d contacts matching '%s'\n", len(matchedContacts), searchTerm)
	showSummary(matchedContacts)
}
