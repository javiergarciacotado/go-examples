package main

import (
	"encoding/json"
	"fmt"
	"go-examples/14-json/domain/model"
)

func firstMain() { //rename to main for its execution
	petstoreList := make([]*model.Petstore, 5)
	petstore := &model.Petstore{
		Name:     "petstore-test",
		Location: "location-test",
	}

	petstore.Dogs = append(petstore.Dogs, &model.Pet{
		Name:  "pet-test",
		Breed: "breed-test",
	})

	petstore.Dogs = append(petstore.Dogs, &model.Pet{Name: "pet-other-test"})
	petstoreList = append(petstoreList, petstore)

	json, err := json.MarshalIndent(petstoreList, "", "\t")
	if err == nil {
		fmt.Printf(string(json))
	}
}
