package main

import (
	"encoding/json"
	"fmt"
	"go-examples/14-json/domain/model"
	"io/ioutil"
)

func main() {
	petstoreList := make([]*model.Petstore, 5)
	bytes, err := ioutil.ReadFile("resources/decoding-example.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	err = json.Unmarshal(bytes, &petstoreList)
	if err != nil {
		fmt.Printf("Error decoding json: %s\n", err)
	}

	fmt.Printf("%v\n", petstoreList)
}
