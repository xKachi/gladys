package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xKachi/gladys/internal/data"
)

func (app *application) createProductHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Currency    string  `json:"currency"`
		Category    string  `json:"category"` // e.g., "ebook", "course"
		FileURL     string  `json:"fileUrl"`  // URL to the digital file
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Dump the contents of the input struct in a HTTP response
	fmt.Fprintf(w, "%+v\n", input)
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
		// Use the new serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
	}
}
