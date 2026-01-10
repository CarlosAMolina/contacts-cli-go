package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type Contact struct {
	ID              int             `json:"id"`
	Addresses       []string        `json:"addresses"`
	Categories      []string        `json:"categories"`
	Emails          []string        `json:"emails"`
	Name            string          `json:"name"`
	Nicknames       []string        `json:"nicknames"`
	Note            string          `json:"note"`
	Phones          []Phone         `json:"phones"`
	SocialNetwork   SocialNetwork   `json:"socialNetwork"`
	Surname         string          `json:"surname"`
	Urls            []string        `json:"urls"`
}

type Phone struct {
	Description string `json:"description"`
	Number      string `json:"number"`
}

type SocialNetwork struct {
	DiscordAccounts   []Discord   `json:"discordAccounts"`
	FacebookAccounts  []string    `json:"facebookAccounts"`
	GithubAccounts    []string    `json:"githubAccounts"`
	InstagramAccounts []string    `json:"instagramAccounts"`
	LinkedinAccounts  []string    `json:"linkedinAccounts"`
	TelegramAccounts  []string    `json:"telegramAccounts"`
	TiktokAccounts    []string    `json:"tiktokAccounts"`
	TwitterAccounts   []string    `json:"twitterAccounts"`
	WallapopAccounts  []Wallapop  `json:"wallapopAccounts"`
}

type Discord struct {
	Alias         string `json:"alias"`
	Discriminator int    `json:"discriminator"`
	GlobalName    string `json:"globalName"`
	LegacyUserName string `json:"legacyUserName"`
	UserName      string `json:"userName"`
}

type Wallapop struct {
	URL  string `json:"url"`
	Note string `json:"note"`
}

type Data struct {
	Contacts []Contact `json:"contacts"`
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

func searchContacts(term string, data Data) []int {
	var matchedIDs []int
	for _, contact := range data.Contacts {
		if searchInValue(reflect.ValueOf(contact), term) {
			matchedIDs = append(matchedIDs, contact.ID)
		}
	}
	return matchedIDs
}

func main() {
	jsonFile, err := ioutil.ReadFile("src/fake.json")
	if err != nil {
		fmt.Println(err)
	}

	var data Data
	json.Unmarshal(jsonFile, &data)

	searchTerm := "foo"
	matchedIDs := searchContacts(searchTerm, data)
	fmt.Printf("Contacts matching '%s': %v\n", searchTerm, matchedIDs)
}

func HelloWorld() string {
	return "Hello, World!"
}