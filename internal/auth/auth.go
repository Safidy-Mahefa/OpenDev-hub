package auth
// Tout ce qui concerne l'authentification : hashage mdp, manipulation token JWT

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"  // bcrypt est un sous-package de crypto qui sert à la sécurité des mdp

)

// fonction pour hasher un mot de passe
func HashPassword(password string)([]byte,error){
	hash,err := bcrypt.GenerateFromPassword([]byte(password),12) //la deuxieme paramètre 12 representa la puissance de calcul lors du hash du mdp
	if err != nil{
		return nil,err
	}
	fmt.Println("Hashage du mot de passe réussi. Hash :\n",string(hash))
	return hash,nil
}

// Fonction pour verifier si un mdp correspond à un hash
func VerifyPassword(hash []byte, password string)(bool){
	err := bcrypt.CompareHashAndPassword(hash,[]byte(password))
	if err != nil{
		fmt.Println("Erreur lors de la comparaison hash/password :",err)
		return false
	}
	return true
}