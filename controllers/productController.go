package controllers

import (
	"log"
	"strconv"

	"github.com/VSM1le/jwt2/database"
	"github.com/VSM1le/jwt2/models"
	repositorys "github.com/VSM1le/jwt2/repositorys"
	"github.com/gofiber/fiber/v2"
)

func SelectProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		products, err := repositoryNew.SelectAllProduct(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "cannot get product service.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "succesful getting product",
			"data":    products,
		})

	}
}

func CreateProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var product models.Product
		if err := c.BodyParser(&product); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		db, err := database.DBinstance()
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		userId := c.Locals("id")
		if userId == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: user not authenticated",
			})
		}

		validationErr := validate.Struct(product)
		if validationErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": validationErr.Error(),
			})
		}

		product.CreatedBy = userId.(int64)
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		err = repositoryNew.CreateProduct(c, &product)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Create product successful",
		})
	}

}
func GetProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		userId := c.Locals("id")
		productId := c.Params("id")
		p, err := strconv.ParseInt(productId, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if userId == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: user not authenticated",
			})
		}
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		product, err := repositoryNew.GetProduct(c, p)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "succesful getting product",
			"data":    product,
		})
	}
}

func UpdateProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var product models.Product
		id := c.Params("id")
		if err := c.BodyParser(&product); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		db, err := database.DBinstance()
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		userId := c.Locals("id")
		if userId == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: user not authenticated",
			})
		}

		user, ok := userId.(int64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: user not authenticated",
			})
		}

		product.UpdatedBy = &user
		validationErr := validate.Struct(product)
		if validationErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": validationErr.Error(),
			})
		}
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		err = repositoryNew.UpdateProduct(c, &product, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "succesful update product",
			"data":    product,
		})
	}
}
