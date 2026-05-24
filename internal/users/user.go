package users
// Le modèle user

import "time"

// User représente un utilisateur dans la base de donnee
// db: représent le nom de chaque colonne (utile pour sqlx)
type User struct{
	ID string 				`db:"id" json:"id"`
	Username string 		`db:"username" json:"username"`
	Email string 			`db:"email" json:"email"`
	Name *string 			`db:"name" json:"name"`
	Bio *string 				`db:"bio" json:"bio"`
	AvatarUrl *string 		`db:"avatarurl" json:"avatarurl"`
	Role string 			`db:"role" json:"role"`
	GithubUsername *string 	`db:"githubusername" json:"githubusername"`
	LinkedinUrl *string 		`db:"linkedinurl" json:"linkedinurl"`
	PortfolioUrl *string 	`db:"portfoliourl" json:"portfoliourl"`
	City *string 			`db:"city" json:"city"`
	TotalPoints int64 		`db:"totalpoints" json:"totalpoints"`
	SeasonPoints int64 		`db:"seasonpoints" json:"seasonpoints"`
	CreatedAt time.Time 	`db:"createdat" json:"createdat"`
	UpdatedAt time.Time 	`db:"updatedat" json:"updatedat"`
}