package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/happynet78/goblogbackend/database"
	"github.com/happynet78/goblogbackend/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // this is the default limit of 100MB
	})
	routes.Setup(app)
	app.Listen(":" + port)
}
