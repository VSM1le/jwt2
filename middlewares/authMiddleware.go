package middlewares

import (
	"strings"

	"github.com/VSM1le/jwt2/helpers"
	"github.com/gofiber/fiber/v2"
)

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientToken := c.Get("Authorization")
		if clientToken == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "No Authorization header provided",
			})
		}
		token := strings.TrimPrefix(clientToken, "Bearer ")
		claims, err := helpers.ValidateToken(token)
		if err != "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		c.Locals("email", claims.Email)
		c.Locals("first_name", claims.FirstName)
		c.Locals("last_name", claims.LastName)
		c.Locals("id", claims.ID)

		return c.Next()

	}
}
