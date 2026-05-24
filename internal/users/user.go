package users
// Le modèle user

import "time"

// User représente un utilisateur dans la base de donnee
// db: représent le nom de chaque colonne
type User struct{
	ID string 				`json: "id"`
	Username string 		`json: "username"`
	Email string 			`json: "email"`
	Name string 			`json: "name"`
	Bio string 				`json: "bio"`
	AvatarUrl string 		`json: "avatarUrl"`
	Role string 			`json: "role"`
	GithubUsername string 	`json: "githubUsername"`
	LinkedinUrl string 		`json: "LinkedinUrl"`
	PortfolioUrl string 	`json: "portfolioUrl"`
	City string 			`json: "city"`
	TotalPoints int64 		`json: "totalPoints"`
	SeasonPoints int64 		`json: "seasonPoints"`
	CreatedAt time.Time 	`json: "createdAt"`
	UpdatedAt time.Time 	`json: "updatedAt"`
}