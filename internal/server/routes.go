package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

// Fonction setup les routes
func setupRoutes(app *fiber.App){
	// route pour verifier si le serveur marche bien
	app.Get("/health", func(c *fiber.Ctx)error{
		return c.JSON(fiber.Map{
			"Statue" : "Ok",
			"env" : os.Getenv("ENV"),
		})
	})

	// Route de test
	app.Get("/ping",func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message" : "pong",
		})
	})

	// Route pour la page home
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Page" : "Home",
		})
	})
}