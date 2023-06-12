package models

import (
	"encoding/json"
	"time"
)

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	Discount	int
	CreatedAt   time.Time
	UpdatedAt	time.Time
}

type BookInput struct {
	Title 			string      `json:"title" binding:"required"`
	Price 			json.Number `json:"price" binding:"required,number"`
	Description 	string 		`json:"description" binding:"required"`
	Rating			int			`json:"rating" binding:"required"`
	Discount 		int 		`json:"discount" binding:"required"`
}

type BookResponse struct {
	ID          int		`json:"id"`
	Title       string	`json:"title"`
	Description string	`json:"description"`
	Price       int		`json:"price"`
	Rating      int		`json:"rating"`
	Discount	int		`json:"discount"`
}