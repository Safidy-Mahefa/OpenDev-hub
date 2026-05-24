package server

import (
	"MainApp/internal/database"
	"MainApp/internal/users"
	"fmt"
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

	// Route pour afficher tt les utilisateurs dans la base
	app.Get("/users",func(c *fiber.Ctx) error{
		newUser,err := users.Create(database.DB,"Safidy06","safidymahefa05@gmail.com","admin")
		fmt.Println("Nouveau utilisateur crée: ",*newUser)
		tabUsers, err := users.GetAll(database.DB)
		if err != nil{
			fmt.Println("Erreur :",err)
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		// Afficher la liste
		return c.JSON(fiber.Map{"Utilisateurs":tabUsers})
	})
}