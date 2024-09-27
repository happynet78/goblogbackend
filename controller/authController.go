package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/happynet78/goblogbackend/database"
	"github.com/happynet78/goblogbackend/models"
	"log"
	"regexp"
	"strings"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9. %+\-]+@[a-z0-9. %+\-.]+\.[a-z]{2,4}`)
	return Re.MatchString(email)
}

func Register(c fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	// if err := c.BodyParser(&data); err
	if err := c.Bind().Body(&data); err != nil {
		fmt.Println("Unable to parse body")
	}
	// Check if password is less than 6 characters
	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 characters",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}
	// Check if email already exist in database
	database.DB.Where("email = ?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exists",
		})
	}
	user := models.User{
		FirstName: data["first_name"].(string),
		LatName:   data["last_name"].(string),
		Phone:     data["phone"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}
	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	c.Status(200)
	return c.JSON(fiber.Map{
		"user":    user,
		"message": "Account created successfullys",
	})
}
