package users
// Les requetes pour la manipulation de Users

import (
	"github.com/jmoiron/sqlx"
)

// Récup tt les users dans la base de donnee 
func GetAll(db *sqlx.DB) ([]User,error){
	var tabUsers []User
	err := db.Select(&tabUsers,"SELECT * FROM users")

	return tabUsers, err
}

// Creer et inserer un nouveau utilisateur dans la base
func Create(db *sqlx.DB, id, username, email, role string) (*User,error){
	var user User
	// Inserer les valeurs dans la base de donnees et dans la variable user
	err := db.QueryRowx(
		`INSERT INTO users
		VALUES($1,$2,$3,$4)
		RETURNING *`,id,username,email,role).StructScan(&user)

	return &user,err
}
