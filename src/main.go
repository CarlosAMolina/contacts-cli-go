package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	jsonFile, err := os.ReadFile("src/fake.json")
	if err != nil {
		fmt.Println(err)
	}

	var data Data
	json.Unmarshal(jsonFile, &data)

	searchTerm := "foo"
	matchedIDs := searchContacts(searchTerm, data)
	fmt.Printf("Contacts matching '%s': %v\n", searchTerm, matchedIDs)
}
