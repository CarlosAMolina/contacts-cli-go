package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <searchTerm>")
		os.Exit(1)
	}

	searchTerm := os.Args[1]

	const jsonPath = "src/fake.json"
	//const jsonPath= "/tmp/contacts.json"

	var data Data = getData(jsonPath)

	matchedIDs := searchContacts(searchTerm, data)
	fmt.Printf("Found %d contacts matching '%s'\n", len(matchedIDs), searchTerm)

	for _, id := range matchedIDs {
		contactDetails := getContactByID(id, data)
		for _, detail := range contactDetails {
			fmt.Println(detail)
		}
	}
}
