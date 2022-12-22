package models

type Product struct {
	Id          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"`
}

type Id struct {
	Id int32 `json:"id"`
}

type Error struct {
	Message string `json:"message"`
}
