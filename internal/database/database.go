package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx" //Base de donnees avec golang
	_ "github.com/lib/pq" //Driver postgres pour gloang
)

var DB *sqlx.DB;

func Connect() {
	// Recuperer l'url de la bd dans env
	dbUrl:= os.Getenv("DATABASE_URL")
	if dbUrl == ""{
		log.Fatal("L'url de la base est introuvable, Vérifiez votre fichier .env")
	}

	fmt.Println("L'url de la base est :",dbUrl)

	// Création de la connexion avec postgresql
	db, err := sqlx.Connect("postgres",dbUrl)
	if err != nil{
		log.Fatal("Erreur de connexion avec la base de donnee :",err)
	}

	// Verifier si la bd répond
	if err := db.Ping(); err != nil{
		log.Fatal("La base de donnee ne répond pas !")
	}

	DB = db
	log.Println("Connexion avec PostgreSQL réussi !")
}