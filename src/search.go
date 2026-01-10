package main

import (
	"fmt"
	"reflect"
	"strings"
)

func getConcatById(id int, contacts []Contact) (*Contact, error) {
	for _, contact := range contacts {
		if contact.ID == id {
			return &contact, nil
		}
	}
	return nil, fmt.Errorf("ID not found %d", id)
}

func searchContacts(term string, contacts []Contact) []Contact {
	var matchedContacts []Contact
	for _, contact := range contacts {
		if searchInValue(reflect.ValueOf(contact), term) {
			matchedContacts = append(matchedContacts, contact)
		}
	}
	return matchedContacts
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

func searchInValue(v reflect.Value, term string) bool {
	switch v.Kind() {
	case reflect.String:
		return strings.Contains(strings.ToLower(v.String()), strings.ToLower(term))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strings.Contains(strings.ToLower(fmt.Sprintf("%d", v.Int())), strings.ToLower(term))
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if searchInValue(v.Field(i), term) {
				return true
			}
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if searchInValue(v.Index(i), term) {
				return true
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			if searchInValue(v.MapIndex(key), term) {
				return true
			}
		}
	}
	return false
}
