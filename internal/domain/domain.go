package domain

type (
	User struct {
		Username     string `json:"username"`
		Password     string `json:"password,omitempty"` // raw password used for validation, not stored
		PasswordHash string `json:"-"`
		Email        string `json:"email"`
	}
)
