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

	jsonFile, err := os.ReadFile("src/fake.json")
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
	fmt.Printf("Contacts matching '%s': %v\n", searchTerm, matchedIDs)
}
