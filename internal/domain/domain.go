package domain

type (
	User struct {
		ID           int // id is passed to domain model for simplicity
		Username     string
		Password     string
		PasswordHash string
		Email        string
	}
	Task struct {
		ID    int // id is passed to domain model for simplicity
		Title string
		Type  string // task type code
		User  string // username
	}
	TaskType struct {
		ID   int
		Code string
	}
)
