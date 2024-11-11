package controllers

import (
	"github.com/VSM1le/jwt2/database"
	"github.com/VSM1le/jwt2/models"
	repositorys "github.com/VSM1le/jwt2/repositorys"
	"github.com/gofiber/fiber/v2"
)

func SelectReceipt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer db.Close()
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		receipts, err := repositoryNew.SelectReceipt(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"massage": "successful selecting receipts.",
			"data":    receipts,
		})
	}
}

func CreateReceipt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer db.Close()
		var receipt models.ReceiptHeader
		err = c.BodyParser(&receipt)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		ids := make([]int64, len(receipt.ReceiptDetails))
		for i, v := range receipt.ReceiptDetails {
			ids[i] = v.Id
		}
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		checkStatus, err := repositoryNew.CheckInvoiceDetail(c, ids)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if checkStatus {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "there is an invoice that already paid or the invoice detail does not exist",
			})
		}
		err = repositoryNew.CreateReceipt(c, &receipt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  "success",
			"massage": "create receipt successful",
		})

	}
}
func CancleReceipt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		db.Close()
		pId, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		err = repositoryNew.CancelReceipt(c, pId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"massage": "successful cancel receipt",
		})
	}
}
