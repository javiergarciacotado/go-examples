package model

type Pet struct {
	Name  string `json:"name"`
	Breed string `json:"breed,omitempty"`
}
