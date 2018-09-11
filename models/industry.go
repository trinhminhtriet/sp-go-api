package models

type Industry struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}
