package main

import (
	"log"
	"os"

	"github.com/josephOsemba/go-backend/internal/app/api"
	"github.com/josephOsemba/go-backend/internal/app/store"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASS", "password")
	dbName := getEnv("DB_NAME", "virtual_lab")

	dataSourceName := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	dbStore, err := store.NewMySQLStore(dataSourceName)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbStore.Close()

	app := fiber.New(fiber.Config{
		AppName: "Virtual Lab API",
	})

	api.SetupMiddleware(app)

	api.SetupRouter(app, dbStore)

	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
