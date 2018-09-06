package models

type Job struct {
	Model
	CompanyID   int     `json:"company_id"`
	Title       string  `json:"title"`
	Salary      string  `json:"salary"`
	Description string  `json:"description"`
	Requirement string  `json:"requirement"`
	Benefit     string  `json:"benefit"`
	Status      int     `json:"status"`
}
