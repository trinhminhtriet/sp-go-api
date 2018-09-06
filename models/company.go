package models

type Company struct {
	Model
	Name        string `json:"name"`
	Type        string `json:"type"`
	Size        string `json:"size"`
	Logo        string `json:"logo"`
	Slogan      string `json:"slogan"`
	Description string `json:"description"`
	Website     string `json:"website"`
	Address     string `json:"address"`
	FoundedAt   string `json:"founded_at"`
	Status      int    `json:"status"`
	Jobs        []Job  `json:"jobs"`
}
