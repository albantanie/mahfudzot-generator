package main

import (
	"flag"
	"log"
	"os"

	"github.com/albantanie/mahfudzot-generator/internal/config"
	"github.com/albantanie/mahfudzot-generator/internal/database"
)

func main() {
	var (
		force = flag.Bool("force", false, "Force seed even if data exists")
		help  = flag.Bool("help", false, "Show help")
	)
	flag.Parse()

	if *help {
		showHelp()
		return
	}

	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := database.New(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Check if data already exists
	if !*force {
		count, err := db.Count()
		if err != nil {
			log.Fatalf("Failed to check existing data: %v", err)
		}
		
		if count > 0 {
			log.Printf("Database already contains %d quotes. Use -force flag to override.", count)
			return
		}
	}

	// Run seeder
	log.Println("Starting database seeding...")
	err = database.SeedDatabase(db)
	if err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	// Verify seeding
	count, err := db.Count()
	if err != nil {
		log.Printf("Warning: Failed to verify seeding: %v", err)
	} else {
		log.Printf("✅ Database seeding completed successfully! Total quotes: %d", count)
	}
}

func showHelp() {
	log.Println("Mahfudzot Generator Database Seeder")
	log.Println("")
	log.Println("Usage:")
	log.Printf("  %s [options]\n", os.Args[0])
	log.Println("")
	log.Println("Options:")
	log.Println("  -force    Force seed even if data already exists")
	log.Println("  -help     Show this help message")
	log.Println("")
	log.Println("Environment Variables:")
	log.Println("  DB_HOST     Database host (default: localhost)")
	log.Println("  DB_PORT     Database port (default: 5432)")
	log.Println("  DB_USER     Database user (default: postgres)")
	log.Println("  DB_PASSWORD Database password")
	log.Println("  DB_NAME     Database name (default: mahfudzot)")
	log.Println("  DB_SSLMODE  SSL mode (default: disable)")
	log.Println("")
	log.Printf("This will seed the database with %d comprehensive mahfudzot quotes.\n", database.GetQuoteCount())
}
