package users

// Les requetes pour la manipulation de Users

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Récup tt les users dans la base de donnee
func GetAll(db *sqlx.DB) ([]User,error){
	var tabUsers []User
	err := db.Select(&tabUsers,"SELECT * FROM users")

	return tabUsers, err
}

// Creer et inserer un nouveau utilisateur dans la base
func Create(db *sqlx.DB, username, email, role string) (*User,error){
	var user User
	// Inserer les valeurs dans la base de donnees et dans la variable user
	err := db.QueryRowx(
		`INSERT INTO users (username,email,role)
		VALUES($1,$2,$3)
		RETURNING *`,username,email,role).StructScan(&user)

	return &user,err
}

// Supprimer un utilisateur par son id/nom
func Delete(db *sqlx.DB, name string)error{
	// db.Exec(query) pour executer la requete mis en param
	query := "DELETE FROM users WHERE username = $1";
	res, err := db.Exec(query,name); //result = interface contenant le nb de lignes modifiés
	if err != nil{
		return err;
	}

	// Récup le nb de lignes supprimés
	row,_ := res.RowsAffected()
	if row <= 0{
		fmt.Println("Aucun utilisateur supprimé; lignes supprimées:",row)
	}else{
		fmt.Printf("%v utilisateurs supprimés.\n",row)
	}
	
	return nil
}
