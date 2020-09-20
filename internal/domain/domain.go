package domain

import "time"

type (
	User struct {
		ID           int  // id is passed to domain model for simplicity
		Service      bool // if user is a service // TODO: only service users create items
		Username     string
		Password     string
		PasswordHash string
		Email        string // TODO: add ent validation
	}
	Task struct {
		ID          int    // id is passed to domain model for simplicity
		Code        string // unique task code for external systems
		User        string // username
		Type        string // task type code
		Slug        string
		Title       string
		Description string
		Active      bool
		Items       []Item
	}
	TaskType struct {
		ID   int
		Code string
	}
	Item struct {
		ID     int
		Hash   string
		Source string
		Task   string // task code (from external system)
		Data   ItemData
	}
	ItemData struct {
		Price    float64 `json:"price,omitempty"`
		Product  string  `json:"product,omitempty"`
		Seller   string  `json:"seller,omitempty"`
		Word     string  `json:"word,omitempty"`
		Rate     float64 `json:"rate,omitempty"`
		Platform string  `json:"platform,omitempty"`
	}
	SystemSymmary struct {
		ID              int
		Users           int
		Tasks           int
		ActiveTasks     int
		Items           int
		AvgItemsPerTask float64
		CreateTime      time.Time
	}
)
