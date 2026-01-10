package main

import (
	"fmt"
	"os"
)

func processTerm(data Data, searchTerm string) {
	matchedContacts := searchContacts(searchTerm, data.Contacts)
	fmt.Printf("Found %d contacts matching '%s'\n", len(matchedContacts), searchTerm)
	for _, contact := range matchedContacts {
		contactDetails := getSummary(contact)
		for _, detail := range contactDetails {
			fmt.Println(detail)
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <jsonPath> <searchTerm>")
		fmt.Println("Example: go run main.go /tmp/contacts.json Carlos")
		os.Exit(1)
	}

	jsonPath := os.Args[1]
	searchTerm := os.Args[2]

	var data Data = getData(jsonPath)
	processTerm(data, searchTerm)
}
