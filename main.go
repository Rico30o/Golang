package main

import (
	"log"

	"sample/db"
	"sample/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

//@Title Sample
//@version 1.16.2

func main() {
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = db.InitDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}

	routes.SetupRoutes(app)

	err = app.Listen(":5000")
	if err != nil {
		log.Fatal("Error starting the server")
	}
}
