package main

import (
	"api/database"
	"api/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	database.InitDatabase()
	log.Fatal(app.Listen(":3000"))

}
