package main

import (
	"log"
	"os"

	"rfid-gateway/internal/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, reading system environment variables")
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_ANON_KEY") 
	
	sb := supabase.CreateClient(supabaseUrl, supabaseKey)
	app := fiber.New()

	app.Use(logger.New())

	routes.SetupRoutes(app, sb)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server Gateway berjalan di http://localhost:%s\n", port)
	
	log.Fatal(app.Listen(":" + port))
}