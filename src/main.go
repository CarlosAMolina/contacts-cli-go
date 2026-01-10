package main

import (
	"encoding/json"
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

	jsonFile, err := os.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var data Data
	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		os.Exit(1)
	}

	matchedIDs := searchContacts(searchTerm, data)
	fmt.Printf("Found %d contacts matching '%s'\n", len(matchedIDs), searchTerm)

	for _, id := range matchedIDs {
		contactDetails := getContactByID(id, data)
		for _, detail := range contactDetails {
			fmt.Println(detail)
		}
	}
}
