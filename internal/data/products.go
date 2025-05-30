package data

import (
	"time"

	"github.com/xKachi/gladys/internal/validator"
)

type Product struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price"`
	Currency    string    `json:"currency"`
	Category    string    `json:"category"` // e.g., ebook, course, etc.
	FileURL     string    `json:"fileUrl"`  // Link to digital file
	Thumbnail   string    `json:"thumbnail,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	SellerID    string    `json:"sellerId"`
}

func ValidateMovie(v *validator.Validator, product *Product) {
	v.Check(product.Title != "", "title", "must be provided")
	v.Check(len(product.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(product.Description != "", "description", "must be provided")
	v.Check(len(product.Description) <= 1000, "description", "must not be more than 1000 bytes long")
	v.Check(product.Price > float64(0.0), "price", "must be greater than 0")
	v.Check(validator.IsValidCurrency(product.Currency), "currency", "invalid currency code")
	v.Check(validator.IsCategoryType(product.Category), "category", "invalid product type: must be one of ebook, course, software, template, other")
}
