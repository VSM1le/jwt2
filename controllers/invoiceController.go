package controllers

import (
	"math"

	"github.com/VSM1le/jwt2/database"
	"github.com/VSM1le/jwt2/models"
	"github.com/VSM1le/jwt2/repositorys"
	"github.com/gofiber/fiber/v2"
)

func formatToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}

func SelectInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer db.Close()
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		invoices, err := repositoryNew.SelectInvoice(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"massage": "successful selecting invoice.",
			"data":    invoices,
		})
	}
}

func CreateInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer db.Close()
		var invoice models.InvoiceHeader
		err = c.BodyParser(&invoice)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		userId := c.Locals("id")
		user, ok := userId.(int64)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "something went wrong.",
			})
		}

		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		customer, err := repositoryNew.GetCustomer(c, invoice.CustomerId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		invoice.InvCustName = customer.CustName
		invoice.InvCustAddress1 = customer.CustAddress1
		invoice.InvCustAddress2 = customer.CustAddress2
		invoice.InvCustZipcode = customer.CustZipcode
		invoice.InvCustbranch = customer.CustBranch
		invoice.CreatedBy = user

		for i := range invoice.InvoiceDetail {
			product, err := repositoryNew.GetProduct(c, invoice.InvoiceDetail[i].ProductId)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			invoice.InvoiceDetail[i].InvdPsCode = product.PsCode
			invoice.InvoiceDetail[i].InvdPsNameTh = product.PsNameTh
			invoice.InvoiceDetail[i].InvdPsNameEn = product.PsNameEn

			invoice.InvoiceDetail[i].InvdVatAmt = formatToTwoDecimals(
				(invoice.InvoiceDetail[i].InvdAmt * float64(invoice.InvoiceDetail[i].InvdVat)) / 100)
			invoice.InvoiceDetail[i].InvdWhtaxAmt = formatToTwoDecimals(
				(invoice.InvoiceDetail[i].InvdAmt * float64(invoice.InvoiceDetail[i].InvdWhtax)) / 100)
			invoice.InvoiceDetail[i].InvdNetAmt =
				invoice.InvoiceDetail[i].InvdAmt + invoice.InvoiceDetail[i].InvdVatAmt

		}

		if err = repositoryNew.CreateInvoice(c, &invoice); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"massage": "successful creating invoice.",
			"data":    invoice,
		})

	}
}
func CancelInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		defer db.Close()
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		userId := c.Locals("id")
		user, ok := userId.(int64)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "something went wrong.",
			})
		}
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		err = repositoryNew.CheckInvoice(c, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		err = repositoryNew.CancelInvoice(c, user, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"massage": "cancel invoice successful.",
		})
	}
}
func UpdateInvoice() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		defer db.Close()
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		var invoice models.InvoiceHeader
		err = c.BodyParser(&invoice)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		userId := c.Locals("id")
		user, ok := userId.(int64)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "something went wrong.",
			})
		}
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		err = repositoryNew.CheckInvoice(c, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		customer, err := repositoryNew.GetCustomer(c, invoice.CustomerId)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		invoice.InvCustName = customer.CustName
		invoice.InvCustAddress1 = customer.CustAddress1
		invoice.InvCustAddress2 = customer.CustAddress2
		invoice.InvCustZipcode = customer.CustZipcode
		invoice.InvCustbranch = customer.CustBranch
		invoice.UpdatedBy = &user

		for i := range invoice.InvoiceDetail {
			product, err := repositoryNew.GetProduct(c, invoice.InvoiceDetail[i].ProductId)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			invoice.InvoiceDetail[i].InvdPsCode = product.PsCode
			invoice.InvoiceDetail[i].InvdPsNameTh = product.PsNameTh
			invoice.InvoiceDetail[i].InvdPsNameEn = product.PsNameEn

			invoice.InvoiceDetail[i].InvdVatAmt = formatToTwoDecimals(
				(invoice.InvoiceDetail[i].InvdAmt * float64(invoice.InvoiceDetail[i].InvdVat)) / 100)
			invoice.InvoiceDetail[i].InvdWhtaxAmt = formatToTwoDecimals(
				(invoice.InvoiceDetail[i].InvdAmt * float64(invoice.InvoiceDetail[i].InvdWhtax)) / 100)
			invoice.InvoiceDetail[i].InvdNetAmt =
				invoice.InvoiceDetail[i].InvdAmt + invoice.InvoiceDetail[i].InvdVatAmt
		}
		if err = repositoryNew.UpdateInvoice(c, &invoice, id); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"massage": "successful creating invoice.",
			"data":    invoice,
		})

	}
}
