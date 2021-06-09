package models

type Book struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Author      string  `json:"author"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}
