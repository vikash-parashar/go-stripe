package main

import (
	"net/http"
	"time"
)

type Invoice struct {
	ID        int       `json:"id"`
	Quantity  int       `json:"quantity"`
	Amount    int       `json:"amount"`
	Product   string    `json:"product"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Hit the handler")
}
func (app *application) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	//TODO:
}
