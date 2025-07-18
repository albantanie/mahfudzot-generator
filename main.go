package main

import (
	"log"
	"net/http"

	"github.com/albantanie/mahfudzot-generator/internal/config"
	"github.com/albantanie/mahfudzot-generator/internal/database"
	"github.com/albantanie/mahfudzot-generator/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database connection
	var quoteHandler *handlers.QuoteHandler

	// Try to connect to database, fallback to mock if it fails
	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Printf("Warning: Database connection failed: %v", err)
		log.Println("Running in demo mode with mock data")
		// Create mock database for demo
		mockDB := database.NewMockDB()
		quoteHandler = handlers.NewQuoteHandler(mockDB)
	} else {
		log.Println("Connected to database successfully")
		// Check if quotes table exists, if not use mock data
		_, err = db.Query("SELECT 1 FROM quotes LIMIT 1")
		if err != nil {
			log.Printf("Warning: Quotes table not found: %v", err)
			log.Println("Running in demo mode with mock data")
			db.Close()
			mockDB := database.NewMockDB()
			quoteHandler = handlers.NewQuoteHandler(mockDB)
		} else {
			quoteHandler = handlers.NewQuoteHandler(db)
			defer db.Close()
		}
	}

	// Create router
	router := mux.NewRouter()

	// Add CORS middleware
	router.Use(corsMiddleware)

	// Health check endpoint
	router.HandleFunc("/health", quoteHandler.HealthCheck).Methods("GET")

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/quotes", quoteHandler.GetQuotes).Methods("GET")
	api.HandleFunc("/quotes/random", quoteHandler.GetRandomQuote).Methods("GET")
	api.HandleFunc("/quotes/{id:[0-9]+}", quoteHandler.GetQuoteByID).Methods("GET")
	api.HandleFunc("/quotes/author/{author}", quoteHandler.GetQuotesByAuthor).Methods("GET")
	api.HandleFunc("/quotes/category/{category}", quoteHandler.GetQuotesByCategory).Methods("GET")

	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}

// corsMiddleware adds CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
