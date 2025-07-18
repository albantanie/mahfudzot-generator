package database

import (
	"database/sql"
	"fmt"

	"github.com/albantanie/mahfudzot-generator/internal/config"
	"github.com/albantanie/mahfudzot-generator/internal/models"
	_ "github.com/lib/pq"
)

// DB represents the database connection
type DB struct {
	*sql.DB
}

// New creates a new database connection
func New(cfg *config.DatabaseConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{db}, nil
}

// QuoteRepository defines the interface for quote operations
type QuoteRepository interface {
	GetAll(limit, offset int) ([]*models.Quote, error)
	GetByID(id int) (*models.Quote, error)
	GetRandom() (*models.Quote, error)
	GetByAuthor(author string, limit, offset int) ([]*models.Quote, error)
	GetByCategory(category string, limit, offset int) ([]*models.Quote, error)
	Create(quote *models.QuoteRequest) (*models.Quote, error)
	Update(id int, quote *models.QuoteRequest) (*models.Quote, error)
	Delete(id int) error
	Count() (int, error)
}

// GetAll retrieves all quotes with pagination
func (db *DB) GetAll(limit, offset int) ([]*models.Quote, error) {
	query := `
		SELECT id, text_arabic, text_latin, translation, author, category, source, created_at, updated_at
		FROM quotes
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []*models.Quote
	for rows.Next() {
		quote := &models.Quote{}
		err := rows.Scan(
			&quote.ID,
			&quote.TextArabic,
			&quote.TextLatin,
			&quote.Translation,
			&quote.Author,
			&quote.Category,
			&quote.Source,
			&quote.CreatedAt,
			&quote.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// GetByID retrieves a quote by its ID
func (db *DB) GetByID(id int) (*models.Quote, error) {
	query := `
		SELECT id, text_arabic, text_latin, translation, author, category, source, created_at, updated_at
		FROM quotes
		WHERE id = $1
	`

	quote := &models.Quote{}
	err := db.QueryRow(query, id).Scan(
		&quote.ID,
		&quote.TextArabic,
		&quote.TextLatin,
		&quote.Translation,
		&quote.Author,
		&quote.Category,
		&quote.Source,
		&quote.CreatedAt,
		&quote.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return quote, nil
}

// GetRandom retrieves a random quote
func (db *DB) GetRandom() (*models.Quote, error) {
	query := `
		SELECT id, text_arabic, text_latin, translation, author, category, source, created_at, updated_at
		FROM quotes
		ORDER BY RANDOM()
		LIMIT 1
	`

	quote := &models.Quote{}
	err := db.QueryRow(query).Scan(
		&quote.ID,
		&quote.TextArabic,
		&quote.TextLatin,
		&quote.Translation,
		&quote.Author,
		&quote.Category,
		&quote.Source,
		&quote.CreatedAt,
		&quote.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return quote, nil
}

// GetByAuthor retrieves quotes by author with pagination
func (db *DB) GetByAuthor(author string, limit, offset int) ([]*models.Quote, error) {
	query := `
		SELECT id, text_arabic, text_latin, translation, author, category, source, created_at, updated_at
		FROM quotes
		WHERE author ILIKE $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := db.Query(query, "%"+author+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []*models.Quote
	for rows.Next() {
		quote := &models.Quote{}
		err := rows.Scan(
			&quote.ID,
			&quote.TextArabic,
			&quote.TextLatin,
			&quote.Translation,
			&quote.Author,
			&quote.Category,
			&quote.Source,
			&quote.CreatedAt,
			&quote.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// GetByCategory retrieves quotes by category with pagination
func (db *DB) GetByCategory(category string, limit, offset int) ([]*models.Quote, error) {
	query := `
		SELECT id, text_arabic, text_latin, translation, author, category, source, created_at, updated_at
		FROM quotes
		WHERE category ILIKE $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := db.Query(query, "%"+category+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []*models.Quote
	for rows.Next() {
		quote := &models.Quote{}
		err := rows.Scan(
			&quote.ID,
			&quote.TextArabic,
			&quote.TextLatin,
			&quote.Translation,
			&quote.Author,
			&quote.Category,
			&quote.Source,
			&quote.CreatedAt,
			&quote.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}

// Create creates a new quote
func (db *DB) Create(req *models.QuoteRequest) (*models.Quote, error) {
	query := `
		INSERT INTO quotes (text_arabic, text_latin, translation, author, category, source)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, text_arabic, text_latin, translation, author, category, source, created_at, updated_at
	`

	quote := &models.Quote{}
	err := db.QueryRow(query,
		req.TextArabic,
		req.TextLatin,
		req.Translation,
		req.Author,
		req.Category,
		req.Source,
	).Scan(
		&quote.ID,
		&quote.TextArabic,
		&quote.TextLatin,
		&quote.Translation,
		&quote.Author,
		&quote.Category,
		&quote.Source,
		&quote.CreatedAt,
		&quote.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return quote, nil
}

// Update updates an existing quote
func (db *DB) Update(id int, req *models.QuoteRequest) (*models.Quote, error) {
	query := `
		UPDATE quotes
		SET text_arabic = $2, text_latin = $3, translation = $4, author = $5, category = $6, source = $7
		WHERE id = $1
		RETURNING id, text_arabic, text_latin, translation, author, category, source, created_at, updated_at
	`

	quote := &models.Quote{}
	err := db.QueryRow(query,
		id,
		req.TextArabic,
		req.TextLatin,
		req.Translation,
		req.Author,
		req.Category,
		req.Source,
	).Scan(
		&quote.ID,
		&quote.TextArabic,
		&quote.TextLatin,
		&quote.Translation,
		&quote.Author,
		&quote.Category,
		&quote.Source,
		&quote.CreatedAt,
		&quote.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return quote, nil
}

// Delete deletes a quote by ID
func (db *DB) Delete(id int) error {
	query := "DELETE FROM quotes WHERE id = $1"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("quote with id %d not found", id)
	}

	return nil
}

// Count returns the total number of quotes
func (db *DB) Count() (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM quotes").Scan(&count)
	return count, err
}
