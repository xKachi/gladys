package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xKachi/gladys/internal/data"
)

func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new product")
}

func (app *application) showProductHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	product := data.Product{
		ID:          "a1b2c3",
		Title:       "Mastering Go APIs",
		Description: "A complete guide to building APIs in Go.",
		Price:       49.99,
		Currency:    "USD",
		Category:    "ebook",
		FileURL:     "https://cdn.example.com/files/mastering-go.pdf",
		Thumbnail:   "https://cdn.example.com/thumbs/go.png",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		SellerID:    "seller-1234",
	}

	// Encode the struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, envelope{"product": product}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
