package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func readFile(path string) Data {
	jsonFile, err := os.ReadFile(path)
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
	return data
}
