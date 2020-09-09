package domain

type (
	User struct {
		Username     string
		Password     string
		PasswordHash string
		Email        string
	}
)
