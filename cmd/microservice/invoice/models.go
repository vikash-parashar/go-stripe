package main

import "time"

type Order struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Items     []Products
}

type Products struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Amount   int    `json:"amount"`
	Product  string `json:"product"`
}
