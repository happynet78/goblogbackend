package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/happynet78/goblogbackend/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
