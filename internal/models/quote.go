package models

import (
	"time"
)

// Quote represents a mahfudzot (Arabic wisdom quote)
type Quote struct {
	ID          int       `json:"id" db:"id"`
	TextArabic  string    `json:"text_arabic" db:"text_arabic"`
	TextLatin   string    `json:"text_latin,omitempty" db:"text_latin"`
	Translation string    `json:"translation,omitempty" db:"translation"`
	Author      string    `json:"author" db:"author"`
	Category    string    `json:"category,omitempty" db:"category"`
	Source      string    `json:"source,omitempty" db:"source"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// QuoteRequest represents the request structure for creating/updating quotes
type QuoteRequest struct {
	TextArabic  string `json:"text_arabic" validate:"required"`
	TextLatin   string `json:"text_latin,omitempty"`
	Translation string `json:"translation,omitempty"`
	Author      string `json:"author" validate:"required"`
	Category    string `json:"category,omitempty"`
	Source      string `json:"source,omitempty"`
}

// QuoteResponse represents the response structure for API calls
type QuoteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    *Quote `json:"data,omitempty"`
}

// QuotesResponse represents the response structure for multiple quotes
type QuotesResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"message,omitempty"`
	Data    []*Quote `json:"data,omitempty"`
	Total   int      `json:"total,omitempty"`
	Page    int      `json:"page,omitempty"`
	Limit   int      `json:"limit,omitempty"`
}

// ErrorResponse represents error response structure
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}
