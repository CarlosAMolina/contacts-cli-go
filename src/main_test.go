package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestSearchContacts(t *testing.T) {
	jsonFile, err := ioutil.ReadFile("fake.json")
	if err != nil {
		t.Fatalf("Failed to read fake.json: %v", err)
	}

	var data Data
	if err := json.Unmarshal(jsonFile, &data); err != nil {
		t.Fatalf("Failed to unmarshal fake.json: %v", err)
	}

	testCases := []struct {
		term         string
		expectedIDs  []int
	}{
		{"foo", []int{1}},
		{"bar", []int{1}},
		{"John", []int{1}},
		{"Peter", []int{2}},
		{"university", []int{1}},
		{"nonexistent", []int{}},
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

func TestHelloWorld(t *testing.T) {
	want := "Hello, World!"
	if got := HelloWorld(); got != want {
		t.Errorf("HelloWorld() = %q, want %q", got, want)
	}
}