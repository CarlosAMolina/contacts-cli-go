package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestSearchContacts(t *testing.T) {
	jsonFile, err := os.ReadFile("fake.json")
	if err != nil {
		t.Fatalf("Failed to read fake.json: %v", err)
	}

	var data Data
	if err := json.Unmarshal(jsonFile, &data); err != nil {
		t.Fatalf("Failed to unmarshal fake.json: %v", err)
	}

	testCases := []struct {
		term        string
		expectedIDs []int
	}{
		{"FOO", []int{1}},
		{"BAR", []int{1}},
		{"John", []int{1}},
		{"Peter", []int{2}},
		{"university", []int{1}},
		{"nonexistent", []int{}},
		{"234", []int{1, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.term, func(t *testing.T) {
			matchedContacts := searchContacts(tc.term, data.Contacts)
			var matchedIDs []int
			for _, contact := range matchedContacts {
				matchedIDs = append(matchedIDs, contact.ID)
			}
			if len(matchedIDs) == 0 && len(tc.expectedIDs) == 0 {
				return
			}
			if !reflect.DeepEqual(matchedIDs, tc.expectedIDs) {
				t.Errorf("searchContacts(%q) = %v, want %v", tc.term, matchedIDs, tc.expectedIDs)
			}
		})
	}
}

func TestGetSummary(t *testing.T) {
	jsonFile, err := os.ReadFile("fake.json")
	if err != nil {
		t.Fatalf("Failed to read fake.json: %v", err)
	}

	var data Data
	if err := json.Unmarshal(jsonFile, &data); err != nil {
		t.Fatalf("Failed to unmarshal fake.json: %v", err)
	}

	testCases := []struct {
		id             int
		expectedResult []string
	}{
		{1, []string{"123456789 (Personal) John Doe. Johnny, JD. University, friend. ID 1", "123123123 (Work) John Doe. Johnny, JD. University, friend. ID 1"}},
		{2, []string{"123456789 Peter. ID 2"}},
		{3, nil},
		{4, []string{"111222333 Jane. Friend. ID 4"}},
	}

	onlySomeFieldsContact := Contact{
		ID:         4,
		Name:       "Jane",
		Categories: []string{"friend"},
		Phones: []Phone{
			{Number: "111222333"},
		},
		Addresses:     []string{},
		Emails:        []string{},
		Nicknames:     []string{},
		Note:          "",
		SocialNetwork: SocialNetwork{},
		Surname:       "",
		Urls:          []string{},
	}
	data.Contacts = append(data.Contacts, onlySomeFieldsContact)

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("ID_%d", tc.id), func(t *testing.T) {
			contact, err := getConcatById(tc.id, data.Contacts)
			if err != nil {
				if tc.expectedResult == nil {
					return // Expected error
				}
				t.Errorf("getConcatById(%d) returned an error: %v", tc.id, err)
			}
			result := getSummary(*contact)
			if !reflect.DeepEqual(result, tc.expectedResult) {
				t.Errorf("getSummary(%d) = %v, want %v", tc.id, result, tc.expectedResult)
			}
		})
	}
}
