package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Fonction pour la création et la configuration du serveur
func New() *fiber.App{
	// Création du serveur avec configuration
	app := fiber.New(fiber.Config{
		ErrorHandler: func (c *fiber.Ctx, err error) error{  //executée en cas d'erreur
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Erreur" : err.Error(),
			}) //L'erreur retourné
		},
	})

	// Creer le middleware logger pour afficher tt les requetes dans les logs
	app.Use(logger.New())

	// Creer le middleware recover pour empecher les crash serveur
	app.Use(recover.New())

	// Ajouter tt les routes
	setupRoutes(app)

	return app
}