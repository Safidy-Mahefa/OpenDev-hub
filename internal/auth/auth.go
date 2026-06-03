package auth

// Tout ce qui concerne l'authentification : hashage mdp, manipulation token JWT

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt" // bcrypt est un sous-package de crypto qui sert à la sécurité des mdp
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

// TOKENS JWT
// Claims represente les donnees qui seront stockés dans le token
type Claims struct{
	ID string	`json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims //RegisteredClaims est une structure contenant des infos supplementaires utiles pour le token (claims standards de la v5)
}

// Creer un token JWT valide
func GenerateToken(id,email string) (string,error){
	// Création du claim à mettre dans le token
	newClaim := Claims{
		ID : id,
		Email : email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)), //expire apres 15 minutes
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	key := os.Getenv("JWT_SECRET") //Récuperer le key secret dans .env

	// Création du token avec la methode HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,newClaim)

	// Convertir le token en string et finaliser sa signature.
	// SignedString() permet de signer le token (avec le key en paramètre) par la méthode SigningMethodHS256 : Forme finale-> header.payload.signature
	// la methode HS256 nécessite un key de type byte pour éviter les erreurs...
	tokenString,err := token.SignedString([]byte(key))
	if err != nil{
		return "",err
	}
	return tokenString,nil
}

// Vérifier si un token est valide
func VerifyToken(token string)(*Claims,error){
	key := os.Getenv("JWT_SECRET") //Récuperer le key secret dans .env
	claim := &Claims{}

	t,err := jwt.ParseWithClaims(
		token, //Le token en chaine de caracteres brute
		claim, //Pointeur vers la structure Claims (cette dernière doit implementer la structure jwr.Claims)
		func (t *jwt.Token) (interface{},error){ //Fonction callback qui prend en param le token en cours, vérifie si l'algo utilisé est correct et renvoyer la clé de signature si oui
			// Vérifier si l'algo utilisé est celui attendu
			if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil,fmt.Errorf("methode de signature innattendu : %v",t.Header["alg"])
			}
			
			// Retourner la clé
			return []byte(key),nil
		})

	if err != nil{ //erreur de parsing ou token invalide
		return nil,err
	}

	if t.Valid{
		fmt.Println("Token valide")
		return claim,nil //Retourner les donnees dans le token si il est valide
	}
	return nil,fmt.Errorf("token invalide !")
}