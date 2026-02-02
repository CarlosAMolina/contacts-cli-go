package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func showSummary(contacts []Contact) {
	for _, contact := range contacts {
		contactDetails := getSummary(contact)
		for _, detail := range contactDetails {
			fmt.Println(detail)
		}
	}
}

// TODO show pretty contact info (drop nulls, not json, etc)
func showContact(contactId int, contacts []Contact) {
	contact, err := searchContactById(contactId, contacts)
	if err != nil {
		panic(err)
	}
	prettyJSON, err := json.MarshalIndent(contact, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prettyJSON))
}

func getSummary(contact Contact) []string {
	var result []string
	name := contact.Name
	if contact.Surname != "" {
		name = name + " " + contact.Surname
	}
	nicknamesStr := ""
	if len(contact.Nicknames) > 0 {
		nicknamesStr = ". " + strings.Join(contact.Nicknames, ", ")
	}
	categoriesStr := ""
	if len(contact.Categories) > 0 {
		categoriesStr = strings.Join(contact.Categories, ", ")
		categoriesStr = strings.ToUpper(string(categoriesStr[0])) + categoriesStr[1:]
		categoriesStr = ". " + categoriesStr
	}
	for _, phone := range contact.Phones {
		description := phone.Description
		if description != "" {
			description = strings.ToUpper(string(description[0])) + description[1:]
			description = "(" + description + ") "
		}
		result = append(result, fmt.Sprintf("%s %s%s%s%s. ID %d", phone.Number, description, name, nicknamesStr, categoriesStr, contact.ID))
	}
	return result
}
