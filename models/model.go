package models

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}

type Error struct {
	Message string `json:"message"`
}
