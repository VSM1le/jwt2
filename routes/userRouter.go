package routes

import (
	"github.com/VSM1le/jwt2/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/login", controllers.Login())
	app.Post("/register", controllers.Signup())
	// app.Post("/products", controllers.CreateProduct())
	// app.Post("/logout", logoutHandler)
}
