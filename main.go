package main

import (
	"log"
	"os"

	"github.com/VSM1le/jwt2/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("no port in .env file")
	}
	router := fiber.New()
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	router.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	router.Listen(":" + port)

}
