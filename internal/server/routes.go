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

	usersRoute := app.Group("/users") //Pour faciliter le regoupement des suffixes de /users

	// Route pour afficher tt les utilisateurs dans la base
	usersRoute.Get("/",func(c *fiber.Ctx) error{
		tabUsers, err := users.GetAll(database.DB)
		if err != nil{
			fmt.Println("Erreur :",err)
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		// Afficher la liste
		return c.JSON(fiber.Map{"Utilisateurs":tabUsers})
	})

	usersRoute.Get("/add", func(c *fiber.Ctx)error{
		// Recuperer les nom,email et role
		nom := c.Query("name")
		email := c.Query("email")
		role := c.Query("role")
		newUser,err := users.Create(database.DB,nom,email,role)
		fmt.Println("Nouveau utilisateur crée: ",*newUser)
		return err
	})

	usersRoute.Get("/delete", func (c *fiber.Ctx)error{
		// Récuperation du nom en paramètre de la route & suppression de l'utilisateur par son nom
		name := c.Query("name")
		err := users.Delete(database.DB,name)
		if err != nil {
			fmt.Println("Erreur lors de la suppression de l'utilisateur:",err.Error())
		}
		return err
	})
}