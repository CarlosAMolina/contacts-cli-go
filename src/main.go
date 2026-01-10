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
	searchTerm := os.Args[2]
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
