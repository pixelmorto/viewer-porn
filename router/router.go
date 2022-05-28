package router

import (
	"api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/all", handler.AllActress)
	api.Get("/delete/:id", handler.DeleteActress)
	api.Post("/create", handler.CreateAllActress)
}
