package product

import "time"

type Product struct {
	ID             int64     `json:"id"`
	Name           string    `json:"name"`
	Quantity       int64     `json:"quantity"`
	CodeValue      string    `json:"code_value"`
	ExpirationDate time.Time `json:"expiration_date"`
	IsPublished    bool      `json:"is_published"`
	Price          float64   `json:"price"`
}

type ReuqestProduct struct {
	Name           string    `json:"name"`
	Quantity       int64     `json:"quantity"`
	CodeValue      string    `json:"code_value"`
	ExpirationDate time.Time `json:"expiration_date"`
	IsPublished    bool      `json:"is_published"`
	Price          float64   `json:"price"`
}
