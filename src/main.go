package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <jsonPath> <searchTerm>")
		fmt.Println("Example: go run main.go /tmp/contacts.json Carlos")
		os.Exit(1)
	}

	jsonPath := os.Args[1]

	var data Data = readFile(jsonPath)
	processAll(data)

	searchTerm := os.Args[2]
	processTerm(data, searchTerm)

	showContact(1, data.Contacts)
}

func processAll(data Data) {
	fmt.Printf("Number of contacts: %d\n", len(data.Contacts))
	showSummary(data.Contacts)
}

func processTerm(data Data, searchTerm string) {
	matchedContacts := searchContacts(searchTerm, data.Contacts)
	fmt.Printf("Found %d contacts matching '%s'\n", len(matchedContacts), searchTerm)
	showSummary(matchedContacts)
}
