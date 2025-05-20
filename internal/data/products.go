package data

import "time"

type Product struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price"`
	Currency    string    `json:"currency"`
	Category    string    `json:"category"`    // e.g., ebook, course, etc.
	FileURL     string    `json:"fileUrl"` // Link to digital file
	Thumbnail   string    `json:"thumbnail,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	SellerID    string    `json:"sellerId"`
}
