package models

type Skill struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}
