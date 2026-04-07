package main

import (
	"log"
	"os"

	"rfid-gateway/internal/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading system environment variables")
	}

	app := fiber.New()

	app.Use(logger.New())

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server Gateway berjalan di http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}