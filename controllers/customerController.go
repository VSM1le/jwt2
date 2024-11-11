package controllers

import (
	"log"
	"strconv"

	"github.com/VSM1le/jwt2/database"
	"github.com/VSM1le/jwt2/models"
	repositorys "github.com/VSM1le/jwt2/repositorys"
	"github.com/gofiber/fiber/v2"
)

func SelectCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			log.Fatal("Could not connect to database")
		}
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		customers, err := repositoryNew.SelectAllCustomer(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"massage": "successful selecting customer.",
			"data":    customers,
		})
	}
}

func CreateCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var customer models.Customer
		err := c.BodyParser(&customer)
		if err != nil {
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

		validationErr := validate.Struct(customer)
		if validationErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": validationErr.Error(),
			})
		}

		customer.CreatedBy = userId.(int64)
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		err = repositoryNew.CreateCustomer(c, &customer)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"massage": "create customer successful.",
		})

	}
}
func UpdateCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var customer models.Customer
		err := c.BodyParser(&customer)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		customerId := c.Params("id")
		p, err := strconv.ParseInt(customerId, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
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
		validationErr := validate.Struct(customer)
		if validationErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": validationErr.Error(),
			})
		}
		user, ok := userId.(int64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: user not authenticated",
			})
		}
		customer.UpdatedBy = &user
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		err = repositoryNew.UpdateCustomer(c, &customer, p)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "successful",
			"massage": "update customer successful.",
			"data":    customer,
		})
	}
}
func GetCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		customerId := c.Params("id")
		db, err := database.DBinstance()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		p, err := strconv.ParseInt(customerId, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		customer, err := repositoryNew.GetCustomer(c, p)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"massage": "successful getting customer.",
			"data":    customer,
		})
	}
}
