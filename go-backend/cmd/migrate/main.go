package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Database configuration
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASS", "password")
	dbName := getEnv("DB_NAME", "virtual_lab")

	dataSourceName := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	// Connect to database
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Get migration files
	migrationDir := "migrations"
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		log.Fatal("Failed to read migrations directory:", err)
	}

	// Create migrations table if it doesn't exist
	createMigrationsTable := `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		version VARCHAR(255) PRIMARY KEY,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	if _, err := db.Exec(createMigrationsTable); err != nil {
		log.Fatal("Failed to create migrations table:", err)
	}

	// Run migrations
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			if err := runMigration(db, filepath.Join(migrationDir, file.Name())); err != nil {
				log.Fatal("Migration failed:", err)
			}
		}
	}

	fmt.Println("All migrations completed successfully!")
}

func runMigration(db *sql.DB, filePath string) error {
	// Check if migration has already been run
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM schema_migrations WHERE version = ?", filePath).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		fmt.Printf("Skipping already applied migration: %s\n", filePath)
		return nil
	}

	// Read migration file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Execute migration
	if _, err := tx.Exec(string(content)); err != nil {
		return fmt.Errorf("failed to execute %s: %v", filePath, err)
	}

	// Record migration
	if _, err := tx.Exec("INSERT INTO schema_migrations (version) VALUES (?)", filePath); err != nil {
		return err
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return err
	}

	fmt.Printf("Applied migration: %s\n", filePath)
	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
