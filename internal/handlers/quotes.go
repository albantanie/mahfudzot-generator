package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/albantanie/mahfudzot-generator/internal/database"
	"github.com/albantanie/mahfudzot-generator/internal/models"
	"github.com/gorilla/mux"
)

// QuoteHandler handles quote-related HTTP requests
type QuoteHandler struct {
	db database.QuoteRepository
}

// NewQuoteHandler creates a new quote handler
func NewQuoteHandler(db database.QuoteRepository) *QuoteHandler {
	return &QuoteHandler{db: db}
}

// GetQuotes handles GET /api/v1/quotes
func (h *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	limit := 10 // default
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	page := 1 // default
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	offset := (page - 1) * limit

	quotes, err := h.db.GetAll(limit, offset)
	if err != nil {
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve quotes", err.Error())
		return
	}

	total, err := h.db.Count()
	if err != nil {
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to count quotes", err.Error())
		return
	}

	response := models.QuotesResponse{
		Success: true,
		Data:    quotes,
		Total:   total,
		Page:    page,
		Limit:   limit,
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// GetRandomQuote handles GET /api/v1/quotes/random
func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.db.GetRandom()
	if err != nil {
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve random quote", err.Error())
		return
	}

	response := models.QuoteResponse{
		Success: true,
		Data:    quote,
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// GetQuoteByID handles GET /api/v1/quotes/{id}
func (h *QuoteHandler) GetQuoteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.sendErrorResponse(w, http.StatusBadRequest, "Invalid quote ID", "ID must be a number")
		return
	}

	quote, err := h.db.GetByID(id)
	if err != nil {
		h.sendErrorResponse(w, http.StatusNotFound, "Quote not found", err.Error())
		return
	}

	response := models.QuoteResponse{
		Success: true,
		Data:    quote,
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// HealthCheck handles GET /health
func (h *QuoteHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":  "healthy",
		"service": "mahfudzot-generator",
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// sendJSONResponse sends a JSON response
func (h *QuoteHandler) sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// GetQuotesByAuthor handles GET /api/v1/quotes/author/{author}
func (h *QuoteHandler) GetQuotesByAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	author := vars["author"]

	if author == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Author parameter is required", "Missing author parameter")
		return
	}

	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	limit := 10 // default
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	page := 1 // default
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	offset := (page - 1) * limit

	quotes, err := h.db.GetByAuthor(author, limit, offset)
	if err != nil {
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve quotes by author", err.Error())
		return
	}

	response := models.QuotesResponse{
		Success: true,
		Data:    quotes,
		Total:   len(quotes), // Note: This is not the total count, just current page count
		Page:    page,
		Limit:   limit,
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// GetQuotesByCategory handles GET /api/v1/quotes/category/{category}
func (h *QuoteHandler) GetQuotesByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	if category == "" {
		h.sendErrorResponse(w, http.StatusBadRequest, "Category parameter is required", "Missing category parameter")
		return
	}

	// Parse query parameters
	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	limit := 10 // default
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	page := 1 // default
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	offset := (page - 1) * limit

	quotes, err := h.db.GetByCategory(category, limit, offset)
	if err != nil {
		h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve quotes by category", err.Error())
		return
	}

	response := models.QuotesResponse{
		Success: true,
		Data:    quotes,
		Total:   len(quotes), // Note: This is not the total count, just current page count
		Page:    page,
		Limit:   limit,
	}

	h.sendJSONResponse(w, http.StatusOK, response)
}

// sendErrorResponse sends an error response
func (h *QuoteHandler) sendErrorResponse(w http.ResponseWriter, statusCode int, message, error string) {
	response := models.ErrorResponse{
		Success: false,
		Error:   error,
		Message: message,
	}

	h.sendJSONResponse(w, statusCode, response)
}
