package model

type (
	CreateTaskRequest struct {
		Title       string `json:"title"`
		Type        string `json:"type"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
	}
)
