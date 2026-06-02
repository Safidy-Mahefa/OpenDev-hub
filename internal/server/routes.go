package server

import (
	"MainApp/internal/auth"
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

	app.Get("/login", func(c *fiber.Ctx)error{
		// Recuperer les nom,email et role
		nom := c.Query("name")
		email := c.Query("email")
		role := c.Query("role")
		password := c.Query("password")

		// Hasher le mot de passe
		hashedPassword,err := auth.HashPassword(password)
		if err != nil{
			fmt.Println("Erreur lors du hash du mot de passe")
			return err
		}
		newUser,err := users.Create(database.DB,nom,email,role,string(hashedPassword))
		fmt.Println("Nouveau utilisateur crée: ",*newUser)
		return err
	})

	app.Get("/register",func(c *fiber.Ctx)error{
		// nom := c.Query("name")
		// password := c.Query("password")
		return fmt.Errorf("Implementation en cours...")
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