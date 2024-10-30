package routes

import (
	"github.com/VSM1le/jwt2/controllers"
	"github.com/VSM1le/jwt2/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	// userGroup := app.Group("/users", middlewares.Authenticate())
	// userGroup.Get(/)
	product := app.Group("/products", middlewares.Authenticate())
	product.Get("/", controllers.SelectProduct())
	product.Post("/", controllers.CreateProduct())
	product.Get("/:id", controllers.GetProduct())
}
