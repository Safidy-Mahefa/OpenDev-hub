package main

import (
	"fmt"
	"os"

	"MainApp/internal/database"
	"MainApp/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	// Charger le fichier env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erreur lors du chargement des fichiers env ! ")
	}
	
	// connexion avec la bd
	database.Connect()

	// Récuperer le port
	port := os.Getenv("PORT")
	if port == "" { //si port introuvable,
		port = "3000"
	}
	fmt.Println("Le port est :", port)

	// Créer le serveur
	app := server.New()

	// Lancer le serveur sur un port
	if err := app.Listen(":"+port) ; err != nil{
		fmt.Printf("Erreur lors du lancement du serveur !")
	}

}
