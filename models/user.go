package models

type User struct {
	Model
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Status    int `json:"status"`
}
