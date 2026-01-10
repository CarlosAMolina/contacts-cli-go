package main

import (
	"fmt"
	"reflect"
	"strings"
)

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

func searchContacts(term string, data Data) []int {
	var matchedIDs []int
	for _, contact := range data.Contacts {
		if searchInValue(reflect.ValueOf(contact), term) {
			matchedIDs = append(matchedIDs, contact.ID)
		}
	}
	return matchedIDs
}
