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
			matchedIDs := searchContacts(tc.term, data)
			if len(matchedIDs) == 0 && len(tc.expectedIDs) == 0 {
				return
			}
			if !reflect.DeepEqual(matchedIDs, tc.expectedIDs) {
				t.Errorf("searchContacts(%q) = %v, want %v", tc.term, matchedIDs, tc.expectedIDs)
			}
		})
	}
}

func TestGetContactByID(t *testing.T) {
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
		{1, []string{"123456789 John. Johnny, JD. (university, friend). ID 1", "123123123 John. Johnny, JD. (university, friend). ID 1"}},
		{2, []string{"123456789 Peter. ID 2"}},
		{3, nil},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("ID_%d", tc.id), func(t *testing.T) {
			result := getContactByID(tc.id, data)
			if !reflect.DeepEqual(result, tc.expectedResult) {
				t.Errorf("getContactByID(%d) = %v, want %v", tc.id, result, tc.expectedResult)
			}
		})
	}
}
