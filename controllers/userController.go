package controllers

import (
	"fmt"
	"log"

	"github.com/VSM1le/jwt2/database"
	"github.com/VSM1le/jwt2/helpers"
	"github.com/VSM1le/jwt2/models"
	repositorys "github.com/VSM1le/jwt2/repositorys"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = fmt.Sprint("email or password is incorrect")
		check = false
	}
	return check, msg
}

func Signup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": validationErr.Error(),
			})
		}

		// Check if email already exists
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		count, err := repositoryNew.GetEmail(c, user.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check email",
			})
		}

		password := HashPassword(user.Password)
		user.Password = password

		if count > 0 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Email already exists",
			})
		}

		err = repositoryNew.CreateUser(c, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User registered successfully",
		})
	}
}
func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		db, err := database.DBinstance()
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		repositoryNew := repositorys.NewPostgreSQLRepository(db)
		foundUser, err := repositoryNew.GetUserByEmail(c, user.Email)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "email or password is incorrect",
			})
		}
		passwordIsValid, msg := VerifyPassword(user.Password, foundUser.Password)
		if !passwordIsValid {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": msg,
			})
		}
		token, refresh_token, _ := helpers.GenerateAllTokens(foundUser.ID, foundUser.Email, foundUser.FirstName, foundUser.LastName)
		err = helpers.UpdateAllTokens(token, refresh_token, foundUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "User updated",
			"data":    foundUser,
		})

	}
}
