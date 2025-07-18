package database

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/albantanie/mahfudzot-generator/internal/models"
)

// MockDB represents a mock database for demo purposes
type MockDB struct {
	quotes []*models.Quote
}

// NewMockDB creates a new mock database with comprehensive seed data
func NewMockDB() *MockDB {
	seedData := GetSeedData()
	quotes := make([]*models.Quote, len(seedData))

	// Convert seed data to Quote models
	for i, seed := range seedData {
		quotes[i] = &models.Quote{
			ID:          i + 1,
			TextArabic:  seed.TextArabic,
			TextLatin:   seed.TextLatin,
			Translation: seed.Translation,
			Author:      seed.Author,
			Category:    seed.Category,
			Source:      seed.Source,
			CreatedAt:   time.Now().Add(-time.Duration(i) * time.Hour),
			UpdatedAt:   time.Now().Add(-time.Duration(i) * time.Hour),
		}
	}

	return &MockDB{quotes: quotes}
}

// GetAll retrieves all quotes with pagination
func (m *MockDB) GetAll(limit, offset int) ([]*models.Quote, error) {
	if offset >= len(m.quotes) {
		return []*models.Quote{}, nil
	}

	end := offset + limit
	if end > len(m.quotes) {
		end = len(m.quotes)
	}

	return m.quotes[offset:end], nil
}

// GetByID retrieves a quote by its ID
func (m *MockDB) GetByID(id int) (*models.Quote, error) {
	for _, quote := range m.quotes {
		if quote.ID == id {
			return quote, nil
		}
	}
	return nil, fmt.Errorf("quote with id %d not found", id)
}

// GetRandom retrieves a random quote
func (m *MockDB) GetRandom() (*models.Quote, error) {
	if len(m.quotes) == 0 {
		return nil, fmt.Errorf("no quotes available")
	}

	randomIndex := rand.Intn(len(m.quotes))
	return m.quotes[randomIndex], nil
}

// GetByAuthor retrieves quotes by author with pagination
func (m *MockDB) GetByAuthor(author string, limit, offset int) ([]*models.Quote, error) {
	var filtered []*models.Quote
	for _, quote := range m.quotes {
		if quote.Author == author {
			filtered = append(filtered, quote)
		}
	}

	if offset >= len(filtered) {
		return []*models.Quote{}, nil
	}

	end := offset + limit
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[offset:end], nil
}

// GetByCategory retrieves quotes by category with pagination
func (m *MockDB) GetByCategory(category string, limit, offset int) ([]*models.Quote, error) {
	var filtered []*models.Quote
	for _, quote := range m.quotes {
		if quote.Category == category {
			filtered = append(filtered, quote)
		}
	}

	if offset >= len(filtered) {
		return []*models.Quote{}, nil
	}

	end := offset + limit
	if end > len(filtered) {
		end = len(filtered)
	}

	return filtered[offset:end], nil
}

// Create creates a new quote (mock implementation)
func (m *MockDB) Create(req *models.QuoteRequest) (*models.Quote, error) {
	newID := len(m.quotes) + 1
	quote := &models.Quote{
		ID:          newID,
		TextArabic:  req.TextArabic,
		TextLatin:   req.TextLatin,
		Translation: req.Translation,
		Author:      req.Author,
		Category:    req.Category,
		Source:      req.Source,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.quotes = append(m.quotes, quote)
	return quote, nil
}

// Update updates an existing quote (mock implementation)
func (m *MockDB) Update(id int, req *models.QuoteRequest) (*models.Quote, error) {
	for i, quote := range m.quotes {
		if quote.ID == id {
			m.quotes[i].TextArabic = req.TextArabic
			m.quotes[i].TextLatin = req.TextLatin
			m.quotes[i].Translation = req.Translation
			m.quotes[i].Author = req.Author
			m.quotes[i].Category = req.Category
			m.quotes[i].Source = req.Source
			m.quotes[i].UpdatedAt = time.Now()
			return m.quotes[i], nil
		}
	}
	return nil, fmt.Errorf("quote with id %d not found", id)
}

// Delete deletes a quote by ID (mock implementation)
func (m *MockDB) Delete(id int) error {
	for i, quote := range m.quotes {
		if quote.ID == id {
			m.quotes = append(m.quotes[:i], m.quotes[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("quote with id %d not found", id)
}

// Count returns the total number of quotes
func (m *MockDB) Count() (int, error) {
	return len(m.quotes), nil
}
